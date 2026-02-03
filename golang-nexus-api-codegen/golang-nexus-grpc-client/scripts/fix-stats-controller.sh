#!/bin/bash
#
# Fix for Stats Controllers (ItemAssociationStats and ItemStats)
# Fixes CompletableFuture return type issue
#

CONTROLLER_FILE="$1"
if [ -z "$CONTROLLER_FILE" ]; then
    echo "Usage: $0 <controller-file-path>"
    exit 1
fi

if [ ! -f "$CONTROLLER_FILE" ]; then
    echo "File not found: $CONTROLLER_FILE"
    exit 0
fi

# Use Python for reliable text processing
export CONTROLLER_FILE_PATH="$CONTROLLER_FILE"
python3 << 'PYEOF'
import re
import os

file_path = os.environ.get('CONTROLLER_FILE_PATH')
if not file_path:
    print("Error: CONTROLLER_FILE_PATH not set")
    exit(1)

with open(file_path, 'r') as f:
    content = f.read()

# Detect which controller we're fixing
is_item_stats = 'ItemStats' in file_path or 'listItemStats' in content
is_item_association_stats = 'ItemAssociationStats' in file_path or 'listItemAssociationStats' in content

if is_item_stats:
    method_name = 'listItemStats'
    service_method = 'listItemStats'
    log_message = 'NexusStatsItemStats Controller -- Request received for nexus.v4.stats:listItemStats'
    ret_type = 'ListItemStatsRet'
    api_response_type = 'ListItemStatsApiResponse'
    mapper_name = 'listItemStatsApiResponseMapper'
    error_method = 'Error in listItemStats'
elif is_item_association_stats:
    method_name = 'listItemAssociationStats'
    service_method = 'listItemAssociationStats'
    log_message = 'NexusStatsItemAssociationStats Controller -- Request received for nexus.v4.stats:listItemAssociationStats'
    ret_type = 'ListItemAssociationStatsRet'
    api_response_type = 'ListItemAssociationStatsApiResponse'
    mapper_name = 'listItemAssociationStatsApiResponseMapper'
    error_method = 'Error in listItemAssociationStats'
else:
    print(f"⚠️  Unknown controller type in {file_path}")
    exit(0)

# Step 1: Fix return type
content = re.sub(
    rf'public CompletableFuture<ResponseEntity<MappingJacksonValue>> {method_name}',
    rf'public ResponseEntity<MappingJacksonValue> {method_name}',
    content
)

# Step 2: Replace CompletableFuture pattern with synchronous call
# Pattern: from "return service.methodName(...)" to "});" (end of method)
pattern = rf'(    log\.debug\("{re.escape(log_message)}"\);\s*\n\s*)return service\.{service_method}\(([^)]+)\)\s*\n\s*\.thenApply\(response -> \{{.*?\n\s*}}\);'

replacement = rf'''\1try {{
      var response = service.{service_method}(\2).get();

      HttpStatus httpStatus = HttpStatus.valueOf(200);
      httpServletResponse.setStatus(httpStatus.value());

      Map<String, String> responseHeaders = response.getReservedMap();
      if (responseHeaders != null) {{
        responseHeaders.forEach((k, v) -> httpServletResponse.addHeader(k, v));
      }}

      if (response.getContent() == null) {{
        return ResponseEntity.noContent().build();
      }}

      if(response.getContent().hasErrorResponseData()) {{
        if (response.getContent().getErrorResponseData().getValue().getErrorCase().equals(
               nexus.v4.error.ErrorResponse.ErrorCase.APP_MESSAGE_ARRAY_ERROR)) {{
          String locale = null;

          if(httpServletRequest.getHeader(HttpHeaders.ACCEPT_LANGUAGE) != null) {{
            locale = httpServletRequest.getHeader(HttpHeaders.ACCEPT_LANGUAGE);
          }} else {{
            locale = defaultLocale;
          }}

          List<nexus.v4.error.AppMessage> appMessageList = response.getContent()
               .getErrorResponseData().getValue().getAppMessageArrayError().getValueList();
          List<nexus.v4.error.AppMessage> appMessageWithErrorMsg = new ArrayList<>();
          for (nexus.v4.error.AppMessage appMessage : appMessageList) {{
               Map<String, String> argumentsMap = new HashMap<>();
               argumentsMap = appMessage.getArgumentsMap().getValueMap();

            com.nutanix.devplatform.messages.models.AppMessage appMsgFromUtils  = AppMessageBuilder.buildAppMessage(locale, artifactsPath,
                   serviceName, servicePrefix, appMessage.getCode(), argumentsMap);

               appMessageWithErrorMsg.add(appMessage.newBuilder().setCode(appMsgFromUtils.getCode()).setMessage(appMsgFromUtils.getMessage()).setLocale(locale).setErrorGroup(appMsgFromUtils.getErrorGroup())
                     .build());
          }}

         nexus.v4.stats.{ret_type} protoWithAppMsg = nexus.v4.stats.{ret_type}.newBuilder()
           .setContent(response.getContent().toBuilder().setErrorResponseData(response.getContent().getErrorResponseData().toBuilder().setValue(
         response.getContent().getErrorResponseData().getValue().toBuilder().setAppMessageArrayError(
         response.getContent().getErrorResponseData().getValue().getAppMessageArrayError().toBuilder().clearValue().addAllValue(appMessageWithErrorMsg).build()).build()).build()).build()).build();

         dp1.mock.nexus.v4.stats.{api_response_type} contentWithAppMessage = {mapper_name}.protoToModel(protoWithAppMsg.getContent());
         return new ResponseEntity(JsonUtils.getMappingJacksonValue(contentWithAppMessage), httpStatus);
        }}
      }}

      dp1.mock.nexus.v4.stats.{api_response_type} content = {mapper_name}.protoToModel(response.getContent());
      MappingJacksonValue mappingJacksonValue = JsonUtils.getMappingJacksonValue(content);

      return new ResponseEntity(mappingJacksonValue, httpStatus);
    }} catch (Exception e) {{
      log.error("{error_method}", e);
      return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).build();
    }}'''

content = re.sub(pattern, replacement, content, flags=re.DOTALL)

# Write back
with open(file_path, 'w') as f:
    f.write(content)

controller_name = os.path.basename(file_path)
print(f"✅ Fixed {controller_name}")
PYEOF


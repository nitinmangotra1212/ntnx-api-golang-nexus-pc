#!/bin/bash
#
# Complete fix for MockConfigCatController.java
# Completely replaces the method body to fix CompletableFuture issue
#

CONTROLLER_FILE="$1"
if [ -z "$CONTROLLER_FILE" ]; then
    CONTROLLER_FILE="target/generated-sources/swagger/src/nexus/v4/config/NexusConfigItemController.java"
fi

if [ ! -f "$CONTROLLER_FILE" ]; then
    echo "File not found: $CONTROLLER_FILE"
    exit 0
fi

# Use Python for reliable text processing - pass file path as environment variable
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

# Step 1: Fix return type
content = re.sub(
    r'public CompletableFuture<ResponseEntity<MappingJacksonValue>> listItems',
    r'public ResponseEntity<MappingJacksonValue> listItems',
    content
)

# Step 2: Find the method body and replace it completely
# Pattern: from "return service.listItems(...)" to "});" (end of method)
# Match method with parameters: listItems($filter, $orderby)
pattern = r'(    log\.debug\("NexusConfigItem Controller -- Request received for nexus\.v4\.config:listItems"\);\s*\n\s*)return service\.listItems\(([^)]+)\)\s*\n\s*\.thenApply\(response -> \{.*?\n\s*\}\);'

replacement = r'''\1try {
      var response = service.listItems(\2).get();

      HttpStatus httpStatus = HttpStatus.valueOf(200);
      httpServletResponse.setStatus(httpStatus.value());

      Map<String, String> responseHeaders = response.getReservedMap();
      if (responseHeaders != null) {
        responseHeaders.forEach((k, v) -> httpServletResponse.addHeader(k, v));
      }

      if (response.getContent() == null) {
        return ResponseEntity.noContent().build();
      }

      if(response.getContent().hasErrorResponseData()) {
        if (response.getContent().getErrorResponseData().getValue().getErrorCase().equals(
               nexus.v4.error.ErrorResponse.ErrorCase.APP_MESSAGE_ARRAY_ERROR)) {
          String locale = null;

          if(httpServletRequest.getHeader(HttpHeaders.ACCEPT_LANGUAGE) != null) {
            locale = httpServletRequest.getHeader(HttpHeaders.ACCEPT_LANGUAGE);
          } else {
            locale = defaultLocale;
          }

          List<nexus.v4.error.AppMessage> appMessageList = response.getContent()
               .getErrorResponseData().getValue().getAppMessageArrayError().getValueList();
          List<nexus.v4.error.AppMessage> appMessageWithErrorMsg = new ArrayList<>();
          for (nexus.v4.error.AppMessage appMessage : appMessageList) {
               Map<String, String> argumentsMap = new HashMap<>();
               argumentsMap = appMessage.getArgumentsMap().getValueMap();

            com.nutanix.devplatform.messages.models.AppMessage appMsgFromUtils  = AppMessageBuilder.buildAppMessage(locale, artifactsPath,
                   serviceName, servicePrefix, appMessage.getCode(), argumentsMap);

               appMessageWithErrorMsg.add(appMessage.newBuilder().setCode(appMsgFromUtils.getCode()).setMessage(appMsgFromUtils.getMessage()).setLocale(locale).setErrorGroup(appMsgFromUtils.getErrorGroup())
                     .build());
          }

         nexus.v4.config.ListItemsRet protoWithAppMsg = nexus.v4.config.ListItemsRet.newBuilder()
           .setContent(response.getContent().toBuilder().setErrorResponseData(response.getContent().getErrorResponseData().toBuilder().setValue(
         response.getContent().getErrorResponseData().getValue().toBuilder().setAppMessageArrayError(
         response.getContent().getErrorResponseData().getValue().getAppMessageArrayError().toBuilder().clearValue().addAllValue(appMessageWithErrorMsg).build()).build()).build()).build()).build();

         dp1.mock.nexus.v4.config.ListItemsApiResponse contentWithAppMessage = listItemsApiResponseMapper.protoToModel(protoWithAppMsg.getContent());
         return new ResponseEntity(JsonUtils.getMappingJacksonValue(contentWithAppMessage), httpStatus);
        }
      }

      dp1.mock.nexus.v4.config.ListItemsApiResponse content = listItemsApiResponseMapper.protoToModel(response.getContent());
      MappingJacksonValue mappingJacksonValue = JsonUtils.getMappingJacksonValue(content);

      return new ResponseEntity(mappingJacksonValue, httpStatus);
    } catch (Exception e) {
      log.error("Error in listItems", e);
      return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).build();
    }'''

content = re.sub(pattern, replacement, content, flags=re.DOTALL)

# Write back
with open(file_path, 'w') as f:
    f.write(content)

print("✅ Fixed NexusConfigItemController.java")
PYEOF

# Adding Metadata to API Responses - Implementation Guide

## Overview

This guide explains how to add metadata (flags, links, totalAvailableResults) to API responses, following the pattern used in `az-manager`.

## How az-manager Implements Metadata

### 1. YAML Definition

**File**: `az-manager-pc-api-definitions/.../azManagerEndpoint.yaml`

```yaml
x-api-responses:
  responseModelName: "ListDomainsApiResponse"
  template: ext:common:/namespaces/common/versioned/v1/modules/response/released/models/apiResponse
```

**Key**: The `template` field automatically adds the `metadata` field to the generated proto.

### 2. Generated Proto Structure

After adding the template, the generated proto will have:

```proto
message ListDomainsApiResponse {
  oneof data {
    DomainArrayWrapper domain_array_data = 2001;
    ErrorResponseWrapper error_response_data = 400;
  }
  optional common.v1.response.ApiResponseMetadata metadata = 1001;  // ← Added by template
  optional ObjectMapWrapper _reserved = 900000;
}
```

### 3. Response Utilities (az-manager)

**File**: `az-manager-service/utils/response/response_utils.go`

Key functions:
- `CreateResponseMetadata()` - Creates metadata with flags and links
- `CreateMetadataFlags()` - Creates flags (hasError, isPaginated)
- `GetPaginationLinks()` - Gets pagination links (first, last, next, self)
- `GetSelfLink()` - Gets self link from gRPC context

### 4. Usage in Service

**File**: `az-manager-service/domain/domain_Services.go`

```go
// Get pagination links
apiUrl := responseUtils.GetApiUrl(ctx, queryParams.Filter, queryParams.Expand, 
    queryParams.Orderby, &queryParams.Limit, &queryParams.Page)
paginationLinks := responseUtils.GetPaginationLinks(totalEntityCount, apiUrl)

// Create response with metadata
isPaginated := listDomainsArg.GetXPage() > 0 || listDomainsArg.GetXLimit() > 0
return responseUtils.CreateListDomainsResponse(domains, paginationLinks, 
    isPaginated, int32(totalEntityCount)), nil
```

---

## Implementation Steps for golang-mock

### Step 1: Update YAML to Add Template

**File**: `ntnx-api-golang-mock-pc/golang-mock-api-definitions/defs/namespaces/mock/versioned/v4/modules/config/released/api/catEndpoint.yaml`

**Change**:
```yaml
x-api-responses:
  responseModelName: "ListCatsApiResponse"
  template: ext:common:/namespaces/common/versioned/v1/modules/response/released/models/apiResponse  # ← ADD THIS
```

### Step 2: Regenerate Code

```bash
cd ~/ntnx-api-golang-mock-pc
mvn clean install -DskipTests -s settings.xml
```

**Verify**: Check that `config.proto` now has:
```proto
message ListCatsApiResponse {
  oneof data { ... }
  optional common.v1.response.ApiResponseMetadata metadata = 1001;  // ← Should appear
  optional ObjectMapWrapper _reserved = 900000;
}
```

### Step 3: Add Dependencies

**File**: `ntnx-api-golang-mock/go.mod`

Add:
```go
require (
    // ... existing dependencies ...
    github.com/nutanix-core/ntnx-api-utils-go v1.0.38  // For pagination links
)
```

**File**: `ntnx-api-golang-mock/go.mod` (replace directives)

**⚠️ IMPORTANT**: You need access to `common/v1/response` protobufs. Check:

1. **If they exist in golang-mock-pc**:
   ```bash
   ls -R ~/ntnx-api-golang-mock-pc/generated-code/protobuf/common/v1/response/
   ```

2. **If not, you may need to**:
   - Reference from `az-manager-pc`:
     ```go
     replace github.com/nutanix-core/ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/response => ../ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/response
     ```
   - Or ensure they're generated during the Maven build
   - Or use a shared dependency if available

3. **Also need `common/v1/config` for FlagArrayWrapper**:
   ```bash
   ls -R ~/ntnx-api-golang-mock-pc/generated-code/protobuf/common/v1/config/
   ```

### Step 4: Create Response Utilities

**File**: `ntnx-api-golang-mock/golang-mock-service/utils/response/response_utils.go`

```go
package response

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	commonConfig "github.com/nutanix-core/ntnx-api-golang-mock-pc/generated-code/protobuf/common/v1/config"
	"github.com/nutanix-core/ntnx-api-golang-mock-pc/generated-code/protobuf/common/v1/response"
	pb "github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/config"
	"github.com/nutanix-core/ntnx-api-utils-go/responseutils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

const (
	HasError    = "hasError"
	IsPaginated = "isPaginated"
	EnvoyOriginalPath = "x-envoy-original-path"
)

// CreateResponseMetadata creates metadata with flags and links
func CreateResponseMetadata(hasError bool, isPaginated bool, paginationLinks []*response.ApiLink, url string, rel string) *response.ApiResponseMetadata {
	links := &response.ApiLinkArrayWrapper{
		Value: paginationLinks,
	}

	if url != "" {
		links = AddToHateOASLinks(links, url, rel)
	}
	return &response.ApiResponseMetadata{
		Flags: CreateMetadataFlags(hasError, isPaginated),
		Links: links,
	}
}

// AddToHateOASLinks adds a link to the HATEOAS links
func AddToHateOASLinks(linksWrapper *response.ApiLinkArrayWrapper, url string, rel string) *response.ApiLinkArrayWrapper {
	if linksWrapper == nil {
		return &response.ApiLinkArrayWrapper{
			Value: []*response.ApiLink{
				{
					Href: proto.String(url),
					Rel:  proto.String(rel),
				},
			},
		}
	}

	linksWrapper.Value = append(linksWrapper.Value, &response.ApiLink{
		Href: proto.String(url),
		Rel:  proto.String(rel),
	})
	return linksWrapper
}

// CreateMetadataFlags creates flags for metadata
func CreateMetadataFlags(hasError bool, isPaginated bool) *commonConfig.FlagArrayWrapper {
	return &commonConfig.FlagArrayWrapper{
		Value: []*commonConfig.Flag{
			{
				Name:  proto.String(HasError),
				Value: proto.Bool(hasError),
			},
			{
				Name:  proto.String(IsPaginated),
				Value: proto.Bool(isPaginated),
			},
		},
	}
}

// GetPathFromGrpcContext extracts the original path from gRPC context
func GetPathFromGrpcContext(ctx context.Context) string {
	uriPath := GetVariableFromGrpcContext(ctx, EnvoyOriginalPath)
	if len(uriPath) > 0 {
		return uriPath[0]
	}
	return ""
}

// GetVariableFromGrpcContext extracts a variable from gRPC context metadata
func GetVariableFromGrpcContext(ctx context.Context, varName string) []string {
	if ctx == nil {
		log.Error("gRPC context is nil")
		return []string{}
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Error("gRPC context doesn't have metadata")
		return []string{}
	}
	retVal, ok := md[varName]
	if !ok {
		log.Errorf("gRPC context metadata doesn't have %v", varName)
		return []string{}
	}
	return retVal
}

// GetSelfLink gets the self link from gRPC context
func GetSelfLink(ctx context.Context) *response.ApiLink {
	hostPort, err := responseutils.GetOriginHostPortFromGrpcContext(ctx)
	if err != nil {
		log.Errorf("Error in getting Host Port info from ctx %v", err)
	}
	if hostPort != "" && !strings.HasPrefix(hostPort, "https://") {
		hostPort = "https://" + hostPort
	}
	uriPath, err := url.Parse(GetPathFromGrpcContext(ctx))
	if err != nil {
		log.Errorf("Error in parsing URI path: %v", err)
	}
	uriPath.RawQuery = ""
	uriBase, err := url.Parse(hostPort)
	if err != nil {
		log.Errorf("Error in parsing Host Port for URI: %v", err)
	}
	selfUri := uriBase.ResolveReference(uriPath).String()

	selfLink := &response.ApiLink{
		Href: proto.String(selfUri),
		Rel:  proto.String("self"),
	}
	return selfLink
}

// GetPaginationLinks gets pagination links
func GetPaginationLinks(total int64, completeUrl string) []*response.ApiLink {
	paginationLinksList := []*response.ApiLink{}
	paginationLinks, err := responseutils.GetPaginationLinks(int(total), completeUrl)
	if err != nil {
		log.Errorf("Error while getting pagination links for %s: %v", completeUrl, err)
	}
	log.Debugf("paginationLinks: %+v", paginationLinks)

	for linkType, link := range paginationLinks {
		apiLink := &response.ApiLink{
			Href: proto.String(link),
			Rel:  proto.String(linkType),
		}
		paginationLinksList = append(paginationLinksList, apiLink)
	}
	return paginationLinksList
}

// GetApiUrl constructs the API URL with query parameters
func GetApiUrl(ctx context.Context, filter, expand, orderby string, limit, page *int32) string {
	apiUrl := *GetSelfLink(ctx).Href + "?"
	if limit == nil {
		limit = proto.Int32(50) // Default limit
	}
	if page == nil {
		page = proto.Int32(0)
	}
	apiUrl = apiUrl + "$limit=" + strconv.FormatInt(int64(*limit), 10) + "&"
	apiUrl = apiUrl + "$page=" + strconv.FormatInt(int64(*page), 10) + "&"
	if filter != "" {
		apiUrl = apiUrl + "$filter=" + filter + "&"
	}
	if expand != "" {
		apiUrl = apiUrl + "$expand=" + expand
	}

	return apiUrl
}

// CreateListCatsResponse creates a response for ListCats API with metadata
func CreateListCatsResponse(cats []*pb.Cat, paginationLinks []*response.ApiLink, isPaginated bool, totalAvailableResults int32) *pb.ListCatsRet {
	resp := &pb.ListCatsRet{
		Content: &pb.ListCatsApiResponse{
			Data: &pb.ListCatsApiResponse_CatArrayData{
				CatArrayData: &pb.CatArrayWrapper{
					Value: cats,
				},
			},
			Metadata: CreateResponseMetadata(false, isPaginated, paginationLinks, "", ""),
		},
	}
	resp.Content.Metadata.TotalAvailableResults = proto.Int32(totalAvailableResults)
	return resp
}
```

### Step 5: Update Service Implementation

**File**: `ntnx-api-golang-mock/golang-mock-service/grpc/cat_grpc_service.go`

**Update the `ListCats` method**:

```go
import (
	// ... existing imports ...
	responseUtils "github.com/nutanix/ntnx-api-golang-mock/golang-mock-service/utils/response"
)

// ListCats implements the gRPC ListCats RPC
func (s *CatGrpcService) ListCats(ctx context.Context, req *pb.ListCatsArg) (*pb.ListCatsRet, error) {
	log.Infof("gRPC: ListCats called")
	log.Debugf("gRPC: ListCats request details: %+v", req)

	s.catMutex.RLock()
	defer s.catMutex.RUnlock()

	log.Debugf("gRPC: Total cats in memory: %d", len(s.cats))

	// Collect all cats
	allCats := make([]*pb.Cat, 0, len(s.cats))
	for _, cat := range s.cats {
		allCats = append(allCats, cat)
	}

	totalCount := int32(len(allCats))
	
	// Determine if paginated (for now, always false since we don't have page/limit in ListCatsArg yet)
	isPaginated := false
	
	// Get pagination links (even if not paginated, we still want self link)
	apiUrl := responseUtils.GetApiUrl(ctx, "", "", "", nil, nil)
	paginationLinks := responseUtils.GetPaginationLinks(int64(totalCount), apiUrl)
	
	// Create response with metadata
	response := responseUtils.CreateListCatsResponse(allCats, paginationLinks, isPaginated, totalCount)

	log.Infof("✅ gRPC: Returning %d cats with metadata", totalCount)
	if log.GetLevel() == log.DebugLevel {
		log.Debugf("gRPC: Returning cats: %+v", allCats)
		log.Debugf("gRPC: Metadata: %+v", response.Content.Metadata)
	}

	return response, nil
}
```

### Step 6: Rebuild and Test

```bash
# Rebuild golang-mock-pc
cd ~/ntnx-api-golang-mock-pc
mvn clean install -DskipTests -s settings.xml

# Rebuild golang-mock
cd ~/ntnx-api-golang-mock
go mod tidy
make build

# Test locally
./golang-mock-server-local -port 9090 -log-level debug
```

---

## Expected Response Structure

After implementation, the response will look like:

```json
{
  "data": [
    {
      "catId": 1,
      "catName": "Cat-1",
      "$reserved": { "$fv": "v4.r1" },
      "$objectType": "mock.v4.config.Cat"
    }
  ],
  "$reserved": { "$fv": "v4.r1" },
  "$objectType": "mock.v4.config.ListCatsApiResponse",
  "metadata": {
    "flags": [
      { "name": "hasError", "value": false },
      { "name": "isPaginated", "value": false }
    ],
    "links": [
      {
        "href": "https://10.112.90.11:9440/api/mock/v4.1/config/cats",
        "rel": "self"
      }
    ],
    "totalAvailableResults": 100
  }
}
```

---

## Important Notes

1. **Common Response Protobufs**: You need access to `common/v1/response` protobufs. Check if they're generated in `golang-mock-pc` or if you need to reference them from `az-manager-pc`.

2. **ntnx-api-utils-go**: This package provides `GetPaginationLinks()` and `GetOriginHostPortFromGrpcContext()`. Ensure it's available in your Go module.

3. **gRPC Context Metadata**: The `x-envoy-original-path` header is set by Adonis/Envoy. If not available, you may need to construct the URL differently.

4. **Pagination**: Currently, `ListCatsArg` doesn't have page/limit fields. Once you add pagination support (as discussed earlier), update `isPaginated` logic.

5. **Template Regeneration**: After adding the `template` field to YAML, you MUST regenerate the code for the metadata field to appear in the proto.

---

## Troubleshooting

### Issue: `common/v1/response` not found

**Solution**: Check if the common response protobufs are generated. You may need to:
- Check if they exist in `golang-mock-pc/generated-code/protobuf/common/v1/response/`
- Or reference them from `az-manager-pc` using a replace directive

### Issue: `ntnx-api-utils-go` not found

**Solution**: 
```bash
go get github.com/nutanix-core/ntnx-api-utils-go@v1.0.38
```

### Issue: Metadata field not in generated proto

**Solution**: 
- Verify the `template` field is in YAML
- **IMPORTANT**: The build requires GitHub access to `https://github.com/nutanix-core/ntnx-api-dev-platform.git` to fetch the common repository
- Ensure `repositories.yaml` exists in `golang-mock-api-definitions/defs/metadata/` with the common repository entry
- Regenerate code: `mvn clean install -DskipTests -s settings.xml`
- Check `config.proto` for `metadata` field

**Note**: If you get "git-upload-pack not permitted" error, you need:
- GitHub authentication configured (SSH keys or token)
- Or access to a cached/mirrored version of the common repository

### Issue: Self link URL incorrect

**Solution**: 
- Check if `x-envoy-original-path` is in gRPC context
- May need to construct URL manually if context metadata is missing

---

## Next Steps

1. ✅ Add `template` to YAML (DONE)
2. ✅ Add `repositories.yaml` with common repository (DONE)
3. ⚠️ Regenerate code (REQUIRES GitHub access to common repository)
4. ✅ Add dependencies to go.mod (DONE)
5. ✅ Create response utils (DONE)
6. ✅ Update service implementation (DONE)
7. ⏳ Test and verify metadata appears in response (WAITING FOR BUILD)

**Current Status**: 
- All code changes are complete
- Build is blocked on GitHub access to fetch common repository
- Once build succeeds, the `metadata` field will be added to `ListCatsApiResponse`
- Then run `go mod tidy` and rebuild the Go service

---

**Last Updated**: 2025-01-XX  
**Related**: `COMPLETE_DEPLOYMENT_GUIDE.md`


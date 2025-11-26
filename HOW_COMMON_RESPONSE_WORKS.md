# How `common/v1/response` Works in az-manager-pc

## Overview

The `common/v1/response` protobuf definitions are **automatically generated** during the Maven build process in `az-manager-pc`. They are not manually defined in YAML files.

## Location in az-manager-pc

```
az-manager-pc/
└── generated-code/
    └── protobuf/
        └── common/
            └── v1/
                ├── config/
                │   ├── config.proto          # FlagArrayWrapper, KVPair, etc.
                │   └── config.pb.go
                └── response/
                    ├── response.proto        # ApiResponseMetadata, ApiLink, etc.
                    └── response.pb.go
```

## Generated Proto Structure

**File**: `az-manager-pc/generated-code/protobuf/common/v1/response/response.proto`

```proto
syntax = "proto2";
package common.v1.response;

import "common/v1/config/config.proto";

message ApiLink {
  optional string href = 1001;
  optional string rel = 1002;
}

message ApiLinkArrayWrapper {
  repeated common.v1.response.ApiLink value = 1000;
}

message ApiResponseMetadata {
  optional common.v1.config.FlagArrayWrapper flags = 1001;
  optional common.v1.response.ApiLinkArrayWrapper links = 1002;
  optional int32 total_available_results = 1003;
  optional common.v1.config.MessageArrayWrapper messages = 1004;
  optional common.v1.config.KVPairArrayWrapper extra_info = 1005;
}
```

## How It's Used

### 1. In Generated Proto Files

Other proto files import it:

```proto
// lifecycle/v4/config/config.proto
import "common/v1/response/response.proto";

message ListDomainsApiResponse {
  oneof data { ... }
  optional common.v1.response.ApiResponseMetadata metadata = 1001;  // ← Uses it here
}
```

### 2. In Go Service Code

**File**: `az-manager-service/utils/response/response_utils.go`

```go
import (
    commonConfig "github.com/nutanix-core/ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/config"
    "github.com/nutanix-core/ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/response"
)

func CreateResponseMetadata(...) *response.ApiResponseMetadata {
    return &response.ApiResponseMetadata{
        Flags: CreateMetadataFlags(hasError, isPaginated),
        Links: links,
    }
}
```

### 3. Import Paths

- **Proto import**: `import "common/v1/response/response.proto";`
- **Go import**: `github.com/nutanix-core/ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/response`
- **Package name**: `common.v1.response`

## How It Gets Generated

The `common/v1/response` protobufs are generated automatically when:

1. **Template is used in YAML**: When you add:
   ```yaml
   x-api-responses:
     template: ext:common:/namespaces/common/versioned/v1/modules/response/released/models/apiResponse
   ```

2. **Maven build runs**: The code generation process:
   - Detects the template reference
   - Generates `common/v1/response/response.proto` from the template
   - Generates `common/v1/config/config.proto` (for FlagArrayWrapper, etc.)
   - Compiles them to `.pb.go` files

3. **Dependencies**: The template system knows to generate these common types because they're referenced by the `template` field.

## For golang-mock-pc

### Option 1: Auto-Generated (Recommended)

When you add the `template` field to your YAML and run `mvn clean install`, the build system should automatically generate `common/v1/response` in your `golang-mock-pc` repository.

**Check after build**:
```bash
ls -R ~/ntnx-api-golang-mock-pc/generated-code/protobuf/common/v1/
```

**Expected structure**:
```
generated-code/protobuf/common/v1/
├── config/
│   ├── config.proto
│   └── config.pb.go
└── response/
    ├── response.proto
    └── response.pb.go
```

### Option 2: Reference from az-manager-pc

If auto-generation doesn't work, you can reference it from `az-manager-pc`:

**In `golang-mock/go.mod`**:
```go
replace github.com/nutanix-core/ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/response => ../ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/response

replace github.com/nutanix-core/ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/config => ../ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/config
```

**In Go code**:
```go
import (
    commonConfig "github.com/nutanix-core/ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/config"
    "github.com/nutanix-core/ntnx-api-az-manager-pc/generated-code/protobuf/common/v1/response"
)
```

### Option 3: Copy Generated Files

If neither works, you can manually copy the generated files:

```bash
# Copy from az-manager-pc to golang-mock-pc
cp -r ~/ntnx-api-az-manager-pc/generated-code/protobuf/common \
   ~/ntnx-api-golang-mock-pc/generated-code/protobuf/
```

**Note**: This is not recommended as it requires manual updates on each regeneration.

## Verification Steps

### 1. Check if Generated

```bash
cd ~/ntnx-api-golang-mock-pc
find generated-code/protobuf/common -name "*.proto" -o -name "*.pb.go" | head -10
```

### 2. Check Proto Import

After adding template and regenerating, check `config.proto`:

```bash
grep "import.*common/v1/response" ~/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/config/config.proto
```

Should show:
```proto
import "common/v1/response/response.proto";
```

### 3. Check Metadata Field

```bash
grep "metadata" ~/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/config/config.proto
```

Should show:
```proto
optional common.v1.response.ApiResponseMetadata metadata = 1001;
```

## Summary

1. **`common/v1/response` is generated automatically** when you use the `template` field in YAML
2. **It's part of the code generation process**, not manually maintained
3. **It provides**: `ApiResponseMetadata`, `ApiLink`, `ApiLinkArrayWrapper`
4. **It depends on**: `common/v1/config` (for `FlagArrayWrapper`, `KVPair`, etc.)
5. **For golang-mock**: Add template → Regenerate → Should auto-generate common protobufs

---

**Last Updated**: 2025-01-XX  
**Related**: `ADD_METADATA_IMPLEMENTATION.md`


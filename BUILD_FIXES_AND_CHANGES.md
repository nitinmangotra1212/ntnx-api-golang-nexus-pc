# Build Fixes and Configuration Changes

**Last Updated**: 2025-11-25  
**Purpose**: Document all fixes and configuration changes made during the golang-mock build and deployment setup

---

## 📋 Table of Contents

1. [Build Fixes](#build-fixes)
2. [Configuration Changes](#configuration-changes)
3. [Proto Generation Fixes](#proto-generation-fixes)
4. [Service Code Updates](#service-code-updates)
5. [Adonis Configuration](#adonis-configuration)
6. [Mercury Configuration](#mercury-configuration)
7. [Complete Build Process](#complete-build-process)

---

## 🔧 Build Fixes

### 1. Import Path Fixes in Generated DTOs

**Problem**: Generated Go DTOs had incorrect import paths like `models/mock/v4/error` instead of full module paths.

**Location**: `ntnx-api-golang-mock-pc/generated-code/dto/src/models/mock/v4/config/config_model.go`

**Fix Applied**:
- Updated `publishCode.sh` script to automatically fix import paths
- Manually fixed `config_model.go` import from:
  ```go
  import1 "models/mock/v4/error"
  ```
  to:
  ```go
  import1 "github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/mock/v4/error"
  ```

**File Modified**: `golang-mock-api-codegen/golang-mock-go-dto-definitions/scripts/publishCode.sh`

**Script Enhancement**:
```bash
# Fix import paths: models/... -> github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/...
sed -i '' 's|"models/|"github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/|g' "$file"
```

---

### 2. Proto Generation Script Fixes

**Problem**: `generate-grpc.sh` couldn't find imported proto files (like `error.proto`) and generated incorrect import paths.

**Location**: `ntnx-api-golang-mock-pc/generate-grpc.sh`

**Fixes Applied**:

#### 2.1 Added Proto Path Configuration
```bash
# Set proto paths - need to include swagger root for imports like mock/v4/error/error.proto
SWAGGER_PROTO_ROOT="${SCRIPT_DIR}/generated-code/protobuf/swagger"
PROTO_OUT_ROOT="${SCRIPT_DIR}/generated-code/protobuf"

# Change to swagger root so we can reference proto files with their full paths
cd "${SWAGGER_PROTO_ROOT}"
```

#### 2.2 Added Error Proto Generation
```bash
# First, generate error.proto (needed by config.proto)
echo "  → error.proto"
protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
    --proto_path="$(go env GOROOT)/src" \
    --go_out="${PROTO_OUT_ROOT}" \
    --go_opt=paths=source_relative \
    mock/v4/error/error.proto
```

#### 2.3 Fixed Import Path Post-Processing
```bash
# Post-process: Fix import paths in generated .pb.go files
for file in $(find "${PROTO_OUT_ROOT}" -name "*.pb.go" -type f); do
    # Fix import paths: mock/v4/... -> github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/...
    sed -i '' 's|"mock/v4/|"github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/|g' "$file"
    # Remove blank imports to non-existent packages
    sed -i '' '/^[[:space:]]*_[[:space:]]*"mock\/v4"$/d' "$file"
done
```

#### 2.4 Added Go Bin to PATH
```bash
# Add Go bin to PATH if not already there
export PATH="${PATH}:$(go env GOPATH)/bin"
```

#### 2.5 Created go.mod for Error Package
**Location**: `generated-code/protobuf/mock/v4/error/go.mod`

**Content**:
```go
module github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/error

go 1.24.0
```

---

### 3. Go Module Replace Directives

**Location**: `ntnx-api-golang-mock/go.mod`

**Changes**:
```go
// Replace directives to use local generated code from ntnx-api-golang-mock-pc
replace github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto => ../ntnx-api-golang-mock-pc/generated-code/dto/src

replace github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/config => ../ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/config

replace github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/error => ../ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/error
```

**Added Dependency**:
```go
require (
    // ... other dependencies ...
    github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/error v0.0.0-00010101000000-000000000000 // indirect
)
```

---

## 🔄 Service Code Updates

### 1. Updated Method Signatures

**Location**: `ntnx-api-golang-mock/golang-mock-service/grpc/cat_grpc_service.go`

**Change**: Updated `ListCats` method to use `Arg`/`Ret` pattern instead of `Request`/`Response`:

**Before**:
```go
func (s *CatGrpcService) ListCats(ctx context.Context, req *pb.ListCatsRequest) (*pb.ListCatsResponse, error)
```

**After**:
```go
func (s *CatGrpcService) ListCats(ctx context.Context, req *pb.ListCatsArg) (*pb.ListCatsRet, error)
```

### 2. Updated Response Structure

**Change**: Updated to use Nutanix API response pattern with `ListCatsApiResponse`:

**Before**:
```go
return &pb.ListCatsResponse{
    Cats:       paginatedCats,
    TotalCount: totalCount,
    Page:       page,
    Limit:      limit,
}, nil
```

**After**:
```go
// Create CatArrayWrapper with all cats
catArrayWrapper := &pb.CatArrayWrapper{
    Value: allCats,
}

// Create ListCatsApiResponse with CatArrayData
apiResponse := &pb.ListCatsApiResponse{
    Data: &pb.ListCatsApiResponse_CatArrayData{
        CatArrayData: catArrayWrapper,
    },
}

// Return ListCatsRet with Content
return &pb.ListCatsRet{
    Content: apiResponse,
}, nil
```

### 3. Fixed Pointer Types in Cat Initialization

**Change**: Updated to use pointers for protobuf fields:

**Before**:
```go
cat := &pb.Cat{
    CatId:       i,
    CatName:     fmt.Sprintf("Cat-%d", i),
    CatType:     "TYPE1",
    Description: "A fluffy cat",
}
```

**After**:
```go
catName := fmt.Sprintf("Cat-%d", i)
catType := "TYPE1"
description := "A fluffy cat"
cat := &pb.Cat{
    CatId:       &i,
    CatName:     &catName,
    CatType:     &catType,
    Description: &description,
}
```

### 4. Commented Out Unimplemented Methods

**Reason**: Only `listCats` is defined in the proto file. Other methods (`GetCat`, `CreateCat`, `UpdateCat`, `DeleteCat`, `GetCatAsync`) are not yet in the proto.

**Action**: Commented out these methods with a note that they should be uncommented once the proto definitions are added.

---

## ⚙️ Adonis Configuration

### 1. Updated `application.yaml`

**Location**: `ntnx-api-prism-service/src/main/resources/application.yaml`

**Changes**:

#### 1.1 Added Controller Packages
```yaml
controller:
  packages:
    onprem:
      - mock.v4.server.configuration
      - mock.v4.config.server.controllers
      - mock.v4.config.server.services
```

**Important Notes**:
- **DO NOT** include `com.nutanix.golangmock` - this package doesn't exist
- **DO NOT** include `com.nutanix.mock.controller` - this package doesn't exist
- **MUST include** `mock.v4.server.configuration` - required for `GolangmockGrpcConfiguration` class that creates the `ManagedChannel` bean
- Use exact package names: `mock.v4.config.server.controllers` (plural!), `mock.v4.config.server.services`

#### 1.2 Added gRPC Configuration
```yaml
grpc:
  Vcenter:
    host: localhost
    port: 2122
  azmanager:
    host: nutanix-infrastructure-manager.pc-platform-other
    port: 8500
  golangmock:
    host: localhost
    port: 9090
```

**Critical**: This configuration is **REQUIRED** for the `GolangmockGrpcConfiguration` class to create the `ManagedChannel` bean. The class reads `grpc.golangmock.host` and `grpc.golangmock.port` from this configuration.

**Note**: `golangmock:` should be at the same indentation level as other gRPC services like `azmanager:` and `Vcenter:`.

### 2. Updated `pom.xml`

**Location**: `ntnx-api-prism-service/pom.xml`

**Changes**:

#### 2.1 Added Version Property
```xml
<properties>
    <!-- ... other properties ... -->
    <golang-mock-api-controller.version>17.0.0-SNAPSHOT</golang-mock-api-controller.version>
</properties>
```

#### 2.2 Added Dependency
```xml
<dependencies>
    <!-- ... other dependencies ... -->
    <dependency>
        <groupId>com.nutanix.nutanix-core.ntnx-api.golang-mock-pc</groupId>
        <artifactId>golang-mock-grpc-client</artifactId>
        <version>${golang-mock-api-controller.version}</version>
        <exclusions>
            <exclusion>
                <artifactId>spring-webmvc</artifactId>
                <groupId>org.springframework</groupId>
            </exclusion>
            <exclusion>
                <artifactId>spring-test</artifactId>
                <groupId>org.springframework</groupId>
            </exclusion>
        </exclusions>
    </dependency>
</dependencies>
```

**Note**: Exclusions were added to resolve Maven Enforcer plugin errors.

#### 2.3 Fixed Spring Boot Plugin Configuration
```xml
<plugin>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-maven-plugin</artifactId>
    <configuration>
        <!-- ... other config ... -->
        <buildInfo>
            <properties>
                <additionalProperty>
                    <name>controller.golang-mock.version</name>
                    <value>${golang-mock-api-controller.version}</value>
                </additionalProperty>
            </properties>
        </buildInfo>
    </configuration>
</plugin>
```

**Fix**: Changed from `${golang-mock-grpc-client.version}` to `${golang-mock-api-controller.version}`.

---

## 🌐 Mercury Configuration

### 1. Created Mercury Config File

**Location on PC**: `~/config/mercury/mercury_request_handler_config_golangmock.json`

**Content**:
```json
{
  "api_path_config_list" : [
    {
      "api_path" : "/api/mock/v4.1",
      "handler_list" : [
        {
          "priority" : 1,
          "port" : 8888,
          "transport_options" : "kHttp",
          "external_request_auth_options" : "kAllowAnyAuthenticatedUserExt",
          "internal_request_auth_options" : "kAllowAnyAuthenticatedUserInt"
        }
      ]
    },
    {
      "api_path" : "/api/mock/unversioned",
      "handler_list" : [
        {
          "priority" : 1,
          "port" : 8888,
          "transport_options" : "kHttp",
          "external_request_auth_options" : "kAllowAnyAuthenticatedUserExt",
          "internal_request_auth_options" : "kAllowAnyAuthenticatedUserInt"
        }
      ]
    }
  ]
}
```

**Important**: 
- API path uses `v4.1` format (minor version) for client-facing URLs
- Port `8888` is Adonis (REST-to-gRPC gateway)
- If there's a conflict with `mercury_request_handler_config_apimock.json`, backup the old config

---

## 📝 PC Configuration Files

### 1. `lookup_cache.json`

**Location on PC**: `/home/nutanix/api_artifacts/lookup_cache.json`

**Entry Added**:
```json
{
  "routeMappings": [
    {
      "apiPath": "/mock/v4.1/config",
      "artifactPath": "mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT"
    }
  ]
}
```

**Notes**:
- `apiPath` uses `v4.1` format (minor version)
- `artifactPath` uses `v4.r1.a1` format (revision.release_type.release_number)

### 2. PC `application.yaml`

**Location on PC**: `/home/nutanix/adonis/config/application.yaml`

**Same changes as local `application.yaml`** (see [Adonis Configuration](#adonis-configuration) above).

**Critical**: Both local and PC configurations must be identical.

---

## 🔨 Complete Build Process

### Step 1: Build golang-mock-pc

```bash
cd ~/ntnx-api-golang-mock-pc
mvn clean install -s settings.xml -DskipTests
```

**Expected Output**:
- ✅ Generated proto files in `generated-code/protobuf/swagger/`
- ✅ Generated Java gRPC client in `golang-mock-api-codegen/golang-mock-grpc-client/target/`
- ✅ Generated artifacts in `golang-mock-api-definitions/target/generated-api-artifacts/`
- ✅ Generated Go DTOs in `generated-code/dto/src/` (with fixed import paths)

### Step 2: Generate gRPC Go Code

```bash
cd ~/ntnx-api-golang-mock-pc
./generate-grpc.sh
```

**What This Does**:
1. Generates `error.pb.go` from `error.proto`
2. Generates `config.pb.go` from `config.proto`
3. Generates `cat_service.pb.go` and `cat_service_grpc.pb.go` from `cat_service.proto`
4. Fixes import paths in all generated `.pb.go` files
5. Fixes method names to lowercase (camelCase) in `cat_service_grpc.pb.go`
6. Removes blank imports to non-existent packages

**Expected Output**:
- ✅ `generated-code/protobuf/mock/v4/error/error.pb.go`
- ✅ `generated-code/protobuf/mock/v4/config/config.pb.go`
- ✅ `generated-code/protobuf/mock/v4/config/cat_service.pb.go`
- ✅ `generated-code/protobuf/mock/v4/config/cat_service_grpc.pb.go`

### Step 3: Build golang-mock Server

```bash
cd ~/ntnx-api-golang-mock
GOOS=linux GOARCH=amd64 go build -o golang-mock-server golang-mock-service/server/main.go
```

**Expected Output**:
- ✅ Binary: `golang-mock-server` (Linux x86-64, ~15MB)

### Step 4: Build prism-service (Adonis)

```bash
cd ~/ntnx-api-prism-service
mvn clean install -DskipTests -s settings.xml
```

**Expected Output**:
- ✅ JAR file: `target/prism-service-17.6.0-SNAPSHOT.jar`
- ✅ Contains `MockConfigCatController.java` (gRPC client)
- ✅ Contains `GolangmockGrpcConfiguration.java` (gRPC channel configuration)

---

## 🚀 Deployment Checklist

### Files to Copy to PC

1. **golang-mock-server binary**
   ```bash
   scp golang-mock-server nutanix@<PC_IP>:~/golang-mock-build/
   ```

2. **prism-service JAR**
   ```bash
   scp target/prism-service-17.6.0-SNAPSHOT.jar nutanix@<PC_IP>:~/adonis/lib/
   ```

3. **Generated API artifacts**
   ```bash
   scp -r golang-mock-api-definitions/target/generated-api-artifacts/* \
     nutanix@<PC_IP>:~/api_artifacts/mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT/
   ```

### PC Configuration Steps

1. **Update `application.yaml`** (see [Adonis Configuration](#adonis-configuration))
2. **Update `lookup_cache.json`** (see [PC Configuration Files](#pc-configuration-files))
3. **Create Mercury config** (see [Mercury Configuration](#mercury-configuration))
4. **Start golang-mock-server**
   ```bash
   cd ~/golang-mock-build
   nohup ./golang-mock-server -port 9090 -log-level debug > ~/golang-mock-build/golang-mock-server.log 2>&1 &
   ```
5. **Restart Adonis/Mercury**
   ```bash
   genesis stop adonis mercury && cluster start
   ```

---

## 🐛 Troubleshooting

### Issue: Import Path Errors

**Symptom**: `package models/mock/v4/error is not in std`

**Solution**: 
1. Check that `publishCode.sh` ran successfully during Maven build
2. Manually fix import in `config_model.go` if needed
3. Rebuild `golang-mock-pc`

### Issue: Proto Import Errors

**Symptom**: `mock/v4/error/error.proto: File not found`

**Solution**:
1. Ensure `generate-grpc.sh` generates `error.proto` first
2. Check that `--proto_path` flags are set correctly
3. Verify `error.proto` exists in `generated-code/protobuf/swagger/mock/v4/error/`

### Issue: Undefined Types

**Symptom**: `undefined: pb.ListCatsRequest`

**Solution**:
1. Use `ListCatsArg` and `ListCatsRet` instead of `ListCatsRequest`/`ListCatsResponse`
2. Check that `generate-grpc.sh` ran successfully
3. Verify generated files exist in `generated-code/protobuf/mock/v4/config/`

### Issue: Method Name Mismatch

**Symptom**: `UNIMPLEMENTED: unknown method listCats`

**Solution**:
1. Ensure `generate-grpc.sh` fixes method names to lowercase
2. Check `cat_service_grpc.pb.go` has `MethodName: "listCats"` (not `"ListCats"`)
3. Rebuild the binary after fixing

---

## 📚 Related Documents

- **`SETUP_GOLANG_MOCK_IN_PC.md`**: Complete deployment guide
- **`REPOSITORY_ARCHITECTURE.md`**: Architecture and repository relationships
- **`RUN_LOCALLY.md`**: Local development and testing guide
- **`DEBUG_LOGGING.md`**: Debug logging configuration

---

## ✅ Summary of All Changes

### Files Modified

1. **`ntnx-api-golang-mock-pc/generate-grpc.sh`**
   - Added proto path configuration
   - Added error.proto generation
   - Added import path fixing
   - Added blank import removal
   - Added PATH configuration

2. **`ntnx-api-golang-mock-pc/golang-mock-api-codegen/golang-mock-go-dto-definitions/scripts/publishCode.sh`**
   - Enhanced import path fixing for DTOs

3. **`ntnx-api-golang-mock/go.mod`**
   - Added replace directives
   - Added error package dependency

4. **`ntnx-api-golang-mock/golang-mock-service/grpc/cat_grpc_service.go`**
   - Updated method signatures to use Arg/Ret pattern
   - Updated response structure
   - Fixed pointer types
   - Commented out unimplemented methods

5. **`ntnx-api-prism-service/pom.xml`**
   - Added golang-mock dependency
   - Added exclusions for Maven Enforcer
   - Fixed Spring Boot plugin configuration

6. **`ntnx-api-prism-service/src/main/resources/application.yaml`**
   - Added controller packages
   - Added gRPC configuration

### Files Created

1. **`ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/error/go.mod`**
   - Go module for error package

### Configuration Files (PC)

1. **`~/config/mercury/mercury_request_handler_config_golangmock.json`**
   - Mercury routing configuration

2. **`~/api_artifacts/lookup_cache.json`** (updated)
   - Added golang-mock route mapping

3. **`~/adonis/config/application.yaml`** (updated)
   - Added controller packages and gRPC config

---

## 🗑️ Cleanup

### Files Deleted

The following unnecessary files were removed:

1. **`ntnx-api-golang-mock/golang-linux-final`**
   - Old test binary (15MB)
   - Replaced by `golang-mock-server` (production binary)

2. **`ntnx-api-golang-mock/golang-mock-server.log`**
   - Temporary log file
   - Already in `.gitignore`

### Updated `.gitignore`

Added to `ntnx-api-golang-mock/.gitignore`:
```
golang-mock-server-local
golang-linux-final
```

These binaries are build artifacts and should not be committed to version control.

---

**Last Updated**: 2025-11-25  
**Status**: ✅ All fixes documented and verified


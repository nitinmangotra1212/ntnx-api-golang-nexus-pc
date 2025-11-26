# Metadata Implementation - Build Steps

## Prerequisites
- ✅ GitHub access configured
- ✅ `ntnx-api-utils-go` cloned locally at `~/ntnx-api-utils-go`
- ✅ All code changes completed

## Step-by-Step Commands

### Step 1: Update Go Module to Use Local ntnx-api-utils-go
```bash
cd ~/ntnx-api-golang-mock
# Replace directive already added to go.mod
```

### Step 2: Regenerate Protobufs with Metadata Field
```bash
cd ~/ntnx-api-golang-mock-pc
mvn clean install -DskipTests -s settings.xml
```

**Expected**: Build should succeed and generate protobufs with `metadata` field.

**Verify metadata field was added**:
```bash
grep "metadata" generated-code/protobuf/mock/v4/config/config.proto
# Should show: optional common.v1.response.ApiResponseMetadata metadata = 1001;
```

### Step 3: Regenerate gRPC Code (if needed)
```bash
cd ~/ntnx-api-golang-mock-pc
./generate-grpc.sh
```

### Step 4: Update Go Dependencies
```bash
cd ~/ntnx-api-golang-mock
go mod tidy
```

**Expected**: Should resolve all dependencies including local `ntnx-api-utils-go`.

### Step 5: Rebuild Go Service
```bash
cd ~/ntnx-api-golang-mock
make build-local
```

**Expected**: Build should succeed without errors.

### Step 6: Test Locally
```bash
cd ~/ntnx-api-golang-mock
./golang-mock-server-local -port 9090 -log-level debug
```

**In another terminal, test with grpcurl**:
```bash
grpcurl -plaintext -d '{}' localhost:9090 mock.v4.config.CatService/listCats
```

**Or test via Postman** (if Adonis is running):
```bash
curl -k -X GET "https://localhost:9440/api/mock/v4.1/config/cats" \
  -H "Authorization: Basic <your-token>" \
  -H "Content-Type: application/json"
```

### Step 7: Verify Metadata in Response

The response should now include metadata:
```json
{
  "data": {
    "catArrayData": {
      "value": [...]
    }
  },
  "metadata": {
    "flags": [
      { "name": "hasError", "value": false },
      { "name": "isPaginated", "value": false }
    ],
    "links": [
      { "href": "https://...", "rel": "self" }
    ],
    "totalAvailableResults": 100
  }
}
```

## Troubleshooting

### Issue: Maven build fails with "git-upload-pack not permitted"
**Solution**: Ensure GitHub authentication is configured:
```bash
# Check if SSH key is loaded
ssh -T git@github.com

# Or configure Git credentials
git config --global credential.helper store
```

### Issue: "metadata field not found" after build
**Solution**: 
1. Verify `template` field is in `catEndpoint.yaml`
2. Check `config.proto` for metadata field
3. If missing, check Maven build logs for template processing errors

### Issue: go mod tidy fails
**Solution**: 
1. Verify `ntnx-api-utils-go` exists at `~/ntnx-api-utils-go`
2. Check replace directive in `go.mod`
3. Try: `go clean -modcache && go mod tidy`

### Issue: Build errors about Metadata field
**Solution**: 
1. Ensure Maven build completed successfully
2. Verify `config.proto` has metadata field
3. Regenerate gRPC code: `./generate-grpc.sh`
4. Rebuild: `make build-local`

## Quick All-in-One Command Sequence

```bash
# 1. Regenerate protobufs
cd ~/ntnx-api-golang-mock-pc && \
mvn clean install -DskipTests -s settings.xml && \
echo "✅ Maven build complete"

# 2. Verify metadata field
grep "metadata" generated-code/protobuf/mock/v4/config/config.proto && \
echo "✅ Metadata field found"

# 3. Regenerate gRPC code
./generate-grpc.sh && \
echo "✅ gRPC code generated"

# 4. Update Go dependencies
cd ~/ntnx-api-golang-mock && \
go mod tidy && \
echo "✅ Go dependencies updated"

# 5. Rebuild Go service
make build-local && \
echo "✅ Go service built"

# 6. Test
./golang-mock-server-local -port 9090 -log-level debug &
echo "✅ Server started on port 9090"
```

---

**Last Updated**: 2025-01-26


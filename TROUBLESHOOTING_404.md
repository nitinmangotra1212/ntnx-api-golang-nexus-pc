# Troubleshooting 404 NOT_FOUND Error

**Error**: `404 NOT_FOUND` when accessing `/api/mock/v4.1/config/cats`

**Log Evidence**:
```
VersionAwareProxyRequestMappingHandlerMapping:getHandlerInternal:79  Request Path is /mock/v4.1/config/cats
AbstractHandlerExceptionResolver:resolveException:146  Resolved [org.springframework.web.server.ResponseStatusException: 404 NOT_FOUND "Not Found"]
```

This means Adonis cannot find a handler for the path `/mock/v4.1/config/cats`.

---

## 🔍 Root Cause Analysis

The 404 error indicates one of these issues:

1. **Controller not loaded** - Package scanning failed
2. **Controller not in JAR** - JAR doesn't contain the controller classes
3. **lookup_cache.json misconfigured** - Route mapping is incorrect
4. **Artifacts missing** - Generated artifacts not in correct location

---

## ✅ Step-by-Step Troubleshooting

### Step 1: Verify Controller is in JAR

**On PC**, check if the controller classes are in the JAR:

```bash
# SSH to PC
ssh nutanix@<PC_IP>

# Check if controller is in JAR
jar -tf /home/nutanix/adonis/ib/prism-service-17.6.0-SNAPSHOT.jar | grep MockConfigCatController

# Expected output:
# mock/v4/config/server/controllers/MockConfigCatController.class

# Also check for the service and configuration classes:
jar -tf /home/nutanix/adonis/lib/prism-service-17.6.0-SNAPSHOT.jar | grep -E "MockConfigCatService|GolangmockGrpcConfiguration|GolangMockGrpcConfiguration"

# Expected outputs:
# mock/v4/config/server/services/MockConfigCatServiceImpl.class
# mock/v4/server/configuration/GolangmockGrpcConfiguration.class
# OR
# mock/v4/server/configuration/GolangMockGrpcConfiguration.class
```

**If classes are missing**: Rebuild and redeploy the JAR (see Step 2).

---

### Step 2: Rebuild and Redeploy JAR

**On local machine**:

```bash
# 1. Rebuild golang-mock-pc (if API definitions changed)
cd ~/ntnx-api-golang-mock-pc
mvn clean install -s settings.xml -DskipTests

# 2. Rebuild prism-service
cd ~/ntnx-api-prism-service
mvn clean install -DskipTests -s settings.xml

# 3. Verify JAR contains controller
jar -tf target/prism-service-17.6.0-SNAPSHOT.jar | grep MockConfigCatController

# 4. Copy to PC
scp -O target/prism-service-17.6.0-SNAPSHOT.jar \
   nutanix@<PC_IP>:/home/nutanix/adonis/lib/

# 5. On PC, restart Adonis
ssh nutanix@<PC_IP>
genesis stop adonis mercury && cluster start
```

---

### Step 3: Verify application.yaml Configuration

**On PC**, check the `application.yaml`:

```bash
# SSH to PC
ssh nutanix@<PC_IP>

# Check controller packages
grep -A 5 "mock.v4" /home/nutanix/adonis/config/application.yaml

# Expected output should include:
# mock.v4.server.configuration, \
# mock.v4.config.server.controllers, \
# mock.v4.config.server.services

# Check gRPC configuration
grep -A 3 "golangmock:" /home/nutanix/adonis/config/application.yaml

# Expected output:
# golangmock:
#   host: localhost
#   port: 9090
```

**If missing or incorrect**: Update `application.yaml` and restart Adonis.

---

### Step 4: Verify lookup_cache.json

**On PC**, check the route mapping:

```bash
# SSH to PC
ssh nutanix@<PC_IP>

# Check lookup_cache.json
cat /home/nutanix/api_artifacts/lookup_cache.json | grep -A 2 "mock"

# Expected output:
# {
#   "apiPath": "/mock/v4.1/config",
#   "artifactPath": "mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT"
# }

# Verify JSON syntax is correct
cat /home/nutanix/api_artifacts/lookup_cache.json | python3 -m json.tool
```

**If missing or incorrect**: Add/update the entry and restart Adonis.

---

### Step 5: Verify Artifacts are Present

**On PC**, check if artifacts are in the correct location:

```bash
# SSH to PC
ssh nutanix@<PC_IP>

# Check artifacts directory
ls -la /home/nutanix/api_artifacts/mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT/

# Expected files:
# - swagger-mock-v4.r1-all.yaml
# - swagger-mock-v4.r1-all.pb
# - api-manifest-17.0.0-SNAPSHOT.json
# - mercury_request_handler_config_apimockv4.1.json
# - object-type-mapping-17.0.0-SNAPSHOT.yaml
# - validationScopes.yaml
# - etc.

# Verify swagger file has the correct path
grep "/mock/v4.1/config/cats" /home/nutanix/api_artifacts/mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT/swagger-mock-v4.r1-all.yaml

# Expected output:
# /mock/v4.1/config/cats:
```

**If missing**: Copy artifacts from local machine (see SETUP_GOLANG_MOCK_IN_PC.md Step 5.3).

---

### Step 6: Check Adonis Logs for Controller Loading

**On PC**, check if controllers are being loaded:

```bash
# SSH to PC
ssh nutanix@<PC_IP>

# Check Adonis logs for controller loading
grep -i "MockConfigCatController\|mock.v4.config.server.controllers\|ComponentScan" /home/nutanix/data/logs/adonis.out | tail -20

# Check for any errors during startup
grep -i "error\|exception\|failed" /home/nutanix/data/logs/adonis.out | grep -i "mock\|golang" | tail -20
```

**Look for**:
- `ComponentScan` messages showing package scanning
- `MockConfigCatController` being instantiated
- Any errors related to `mock.v4` packages

---

### Step 7: Verify golang-mock-server is Running

**On PC**, check if the gRPC server is running:

```bash
# SSH to PC
ssh nutanix@<PC_IP>

# Check if server is running
ps aux | grep golang-mock-server | grep -v grep

# Check if port 9090 is listening
netstat -tlnp | grep 9090

# Check server logs
tail -50 ~/golang-mock-build/golang-mock-server.log
```

**If not running**: Start the server (see SETUP_GOLANG_MOCK_IN_PC.md Step 6.2).

---

## 🔧 Common Fixes

### Fix 1: Rebuild and Redeploy Everything

If you made changes to API definitions or code:

```bash
# 1. Rebuild golang-mock-pc
cd ~/ntnx-api-golang-mock-pc
mvn clean install -s settings.xml -DskipTests

# 2. Generate gRPC code
./generate-grpc.sh

# 3. Rebuild golang-mock server
cd ~/ntnx-api-golang-mock
make build

# 4. Rebuild prism-service
cd ~/ntnx-api-prism-service
mvn clean install -DskipTests -s settings.xml

# 5. Deploy to PC (follow SETUP_GOLANG_MOCK_IN_PC.md Step 5)
```

### Fix 2: Verify Package Names are Correct

**In `application.yaml`**, ensure packages are listed correctly:

```yaml
adonis:
  controller:
    packages:
      onprem: com.nutanix.catalogserver,\
        # ... other packages ...
        mock.v4.server.configuration, \
        mock.v4.config.server.controllers, \
        mock.v4.config.server.services
```

**Important**:
- Use comma-separated format with backslashes (not YAML list format)
- No trailing comma after the last package
- Exact package names: `mock.v4.config.server.controllers` (plural!)

### Fix 3: Verify lookup_cache.json Format

**In `lookup_cache.json`**, ensure the entry is correct:

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

**Important**:
- `apiPath` uses `v4.1` format (minor version)
- `artifactPath` uses `v4.r1.a1` format (revision.release_type.release_number)
- JSON syntax must be valid

---

## 🎯 Quick Checklist

Before asking for help, verify:

- [ ] Controller classes are in JAR (`jar -tf prism-service-*.jar | grep MockConfigCatController`)
- [ ] `application.yaml` has correct packages (`mock.v4.server.configuration`, `mock.v4.config.server.controllers`, `mock.v4.config.server.services`)
- [ ] `application.yaml` has gRPC config (`grpc.golangmock.host` and `port`)
- [ ] `lookup_cache.json` has correct entry (`/mock/v4.1/config`)
- [ ] Artifacts are in correct location (`~/api_artifacts/mock/v4.r1.a1/...`)
- [ ] `golang-mock-server` is running on port 9090
- [ ] Adonis was restarted after configuration changes
- [ ] No errors in Adonis logs related to `mock.v4` packages

---

## 📚 Related Documents

- **`SETUP_GOLANG_MOCK_IN_PC.md`**: Complete deployment guide
- **`BUILD_FIXES_AND_CHANGES.md`**: Configuration changes documentation

---

**Last Updated**: 2025-11-25


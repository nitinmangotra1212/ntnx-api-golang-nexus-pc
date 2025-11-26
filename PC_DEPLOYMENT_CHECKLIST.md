# PC Deployment Checklist - Metadata Implementation

## ✅ Pre-Deployment Verification

### 1. Code Changes
- ✅ YAML updated with `template` field
- ✅ `repositories.yaml` created
- ✅ `response_utils.go` created
- ✅ `cat_grpc_service.go` updated with metadata
- ✅ `go.mod` updated with dependencies

### 2. Build Status
- ✅ Maven build successful (with metadata field)
- ✅ gRPC code generated
- ✅ Linux binary built (`golang-mock-server`)
- ✅ Local testing successful (metadata in response)

### 3. Generated Artifacts
- ✅ Proto files with metadata field
- ✅ `.pb.go` files generated
- ✅ Import paths fixed

## 📋 Deployment Steps

### Step 1: Rebuild prism-service (if needed)

**Check if prism-service needs rebuild:**
```bash
cd ~/ntnx-api-prism-service
# Check if golang-mock-grpc-client version matches
grep "golang-mock-api-controller.version" pom.xml
```

**If version is `17.0.0-SNAPSHOT`, rebuild:**
```bash
cd ~/ntnx-api-prism-service
mvn clean install -DskipTests -s settings.xml
```

**Verify JAR contains required classes:**
```bash
cd /tmp && rm -rf BOOT-INF
jar -xf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
    BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar
jar -tf BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar | \
    grep -E "MockConfigCatController|GolangmockGrpcConfiguration"
rm -rf BOOT-INF
```

### Step 2: Copy Files to PC

```bash
# 1. Copy Go binary
scp ~/ntnx-api-golang-mock/golang-mock-server \
    nutanix@<PC_IP>:~/golang-mock-build/

# 2. Copy generated API artifacts
scp -r ~/ntnx-api-golang-mock-pc/golang-mock-api-definitions/target/generated-api-artifacts/* \
    nutanix@<PC_IP>:~/api_artifacts/mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT/

# 3. Copy Adonis JAR (if rebuilt)
scp ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
    nutanix@<PC_IP>:/home/nutanix/adonis/lib/
```

### Step 3: Verify PC Configuration

**Check `application.yaml` has:**
```yaml
adonis:
  controller:
    packages:
      onprem: |
        mock.v4.server.configuration, \
        mock.v4.config.server.controllers, \
        mock.v4.config.server.services, \
```

**Check `lookup_cache.json` has:**
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

**Check `grpc` section has:**
```yaml
grpc:
  golangmock:
    host: localhost
    port: 9090
```

### Step 4: Start Services on PC

```bash
ssh nutanix@<PC_IP>

# Start golang-mock-server
cd ~/golang-mock-build
nohup ./golang-mock-server -port 9090 -log-level debug > \
    ~/golang-mock-build/golang-mock-server.log 2>&1 &

# Verify it's running
ps aux | grep golang-mock-server | grep -v grep
netstat -tlnp | grep 9090

# Restart Adonis
genesis stop adonis mercury
sleep 5
cluster start
sleep 120  # Wait for Adonis to start

# Check logs
tail -f /home/nutanix/data/logs/adonis.out | grep -i "golang\|mock"
```

### Step 5: Test API

```bash
# Get token
TOKEN=$(curl -k -X POST https://<PC_IP>:9440/api/nutanix/v3/users/login \
    -H "Content-Type: application/json" \
    -d '{"username":"admin","password":"<PASSWORD>"}' | \
    python3 -c "import sys, json; print(json.load(sys.stdin).get('service_account', {}).get('value', ''))")

# Test API
curl -k -X GET "https://<PC_IP>:9440/api/mock/v4.1/config/cats" \
    -H "Authorization: Basic $TOKEN" \
    -H "Content-Type: application/json" | jq '.metadata'
```

**Expected**: Response should include `metadata` with flags, links, and totalAvailableResults.

## ⚠️ Important Notes

1. **No code changes needed** - All metadata implementation is complete
2. **prism-service rebuild** - Only needed if `golang-mock-grpc-client` version changed
3. **Generated artifacts** - Should include `.pb` files (compiled protobuf) with metadata field
4. **Metadata will work automatically** - Once deployed, responses will include metadata

## 🎯 Summary

**Status**: ✅ **READY FOR PC DEPLOYMENT**

All code changes are complete. The metadata implementation will work automatically once:
1. Linux binary is deployed
2. Generated artifacts are copied
3. Services are restarted

No additional code changes are required!

---

**Last Updated**: 2025-01-26


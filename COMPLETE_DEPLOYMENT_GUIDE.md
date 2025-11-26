# Complete Deployment Guide: golang-mock Service

## Table of Contents
1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Build Process](#build-process)
4. [Deployment Steps](#deployment-steps)
5. [Verification](#verification)
6. [Troubleshooting](#troubleshooting)
7. [Common Issues and Solutions](#common-issues-and-solutions)

---

## Overview

This guide covers the complete deployment process for the **golang-mock** service, which migrates from Java-based mockrest to a Go-based gRPC service. The architecture follows:

```
REST Client → Mercury (Nginx) → Adonis (REST-to-gRPC Gateway) → golang-mock (gRPC Server)
```

**Key Components:**
- `ntnx-api-golang-mock-pc`: Code generation (DTOs, protobuf, gRPC client)
- `ntnx-api-golang-mock`: Go gRPC server implementation
- `ntnx-api-prism-service`: Adonis gateway (REST-to-gRPC conversion)

---

## Prerequisites

### Local Setup
- Java 21+
- Maven 3.6+
- Go 1.24+
- `protoc` (Protocol Buffer compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins
- Access to Nutanix Artifactory
- Git access to all three repositories

### PC Access
- SSH access to PC (IP: `10.112.90.11` or your PC IP)
- `nutanix` user credentials
- Permissions to modify `/home/nutanix/adonis/` and `/home/nutanix/config/`

---

## Build Process

### Step 1: Build golang-mock-pc (Code Generation)

```bash
cd ~/ntnx-api-golang-mock-pc
mvn clean install -DskipTests -s settings.xml
```

**What this generates:**
- Go DTOs in `generated-code/dto/src/models/`
- Protobuf definitions in `generated-code/protobuf/swagger/`
- Java gRPC client in `golang-mock-api-codegen/golang-mock-grpc-client/target/`

**Expected artifacts:**
- `golang-mock-grpc-client-17.0.0-SNAPSHOT.jar` (contains controllers, services, configuration)
- Generated `.proto` files
- Generated `.pb.go` files

**Verify build:**
```bash
# Check JAR was created
ls -lh golang-mock-api-codegen/golang-mock-grpc-client/target/*.jar

# Check generated code exists
ls -R generated-code/dto/src/models/
ls -R generated-code/protobuf/
```

### Step 2: Generate gRPC Go Code

```bash
cd ~/ntnx-api-golang-mock-pc
./generate-grpc.sh
```

**What this does:**
- Generates `.pb.go` files from `.proto` files
- Fixes import paths in generated code
- Fixes method names to lowercase (camelCase) for Java client compatibility

**Verify generation:**
```bash
ls -lh generated-code/protobuf/mock/v4/config/*.pb.go
# Should see: cat_service.pb.go, cat_service_grpc.pb.go, config.pb.go
```

### Step 3: Build golang-mock (Go Service)

```bash
cd ~/ntnx-api-golang-mock
make build
```

**Expected output:**
- `golang-mock-server` binary (Linux, amd64)

**Verify build:**
```bash
file golang-mock-server
# Should show: ELF 64-bit LSB executable, x86-64
```

### Step 4: Build prism-service (Adonis)

```bash
cd ~/ntnx-api-prism-service
mvn clean install -DskipTests -s settings.xml
```

**Expected output:**
- `target/prism-service-17.6.0-SNAPSHOT.jar`

**⚠️ CRITICAL: Verify JAR Contents Before Deployment**

```bash
# Extract and check nested JAR
cd /tmp
rm -rf BOOT-INF 2>/dev/null
jar -xf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
    BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar

# Verify required classes exist
jar -tf BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar | grep -E \
    "MockConfigCatController|GolangmockGrpcConfiguration|MockConfigCatService"

# Expected output:
# mock/v4/config/server/controllers/MockConfigCatController.class
# mock/v4/server/configuration/GolangmockGrpcConfiguration.class
# mock/v4/config/server/services/MockConfigCatServiceImpl.class
# mock/v4/config/server/services/MockConfigCatService.class

# Cleanup
rm -rf BOOT-INF
```

**If classes are missing:**
- Rebuild `golang-mock-pc` (Step 1)
- Rebuild `prism-service` (Step 4)
- Check `pom.xml` version matches: `${golang-mock-api-controller.version}` should be `17.0.0-SNAPSHOT`

---

## Deployment Steps

### Step 1: Prepare PC Directories

```bash
ssh nutanix@10.112.90.11

# Create directories
mkdir -p ~/golang-mock-build
mkdir -p ~/api_artifacts/mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT
```

### Step 2: Copy Generated Artifacts to PC

**From local machine:**

```bash
# Copy Go binary
scp ~/ntnx-api-golang-mock/golang-mock-server \
    nutanix@10.112.90.11:~/golang-mock-build/

# Copy generated API artifacts (YAML, .pb files)
scp -r ~/ntnx-api-golang-mock-pc/golang-mock-api-definitions/target/generated-api-artifacts/* \
    nutanix@10.112.90.11:~/api_artifacts/mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT/

# Copy Adonis JAR (backup existing first!)
ssh nutanix@10.112.90.11 "cp /home/nutanix/adonis/lib/prism-service-*.jar /home/nutanix/adonis/lib/prism-service-BACKUP-$(date +%Y%m%d).jar"

scp ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
    nutanix@10.112.90.11:/home/nutanix/adonis/lib/
```

### Step 3: Configure PC

**SSH to PC and configure:**

```bash
ssh nutanix@10.112.90.11
```

#### 3.1 Update application.yaml

**File:** `/home/nutanix/adonis/config/application.yaml`

**Add to `adonis.controller.packages.onprem`:**
```yaml
adonis:
  controller:
    packages:
      onprem: |
        mock.v4.server.configuration, \
        mock.v4.config.server.controllers, \
        mock.v4.config.server.services, \
        # ... other packages ...
```

**Add to `grpc:` section:**
```yaml
grpc:
  # ... other grpc services ...
  golangmock:
    host: localhost
    port: 9090
```

**Add validation bypass (if needed):**
```yaml
oas:
  validation:
    enabled: false
```

#### 3.2 Update lookup_cache.json

**File:** `/home/nutanix/api_artifacts/lookup_cache.json`

**Add entry:**
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

#### 3.3 Create Mercury Config

**File:** `~/config/mercury/mercury_request_handler_config_golangmock.json`

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

**⚠️ Check for conflicts:**
```bash
# Remove old mock config if exists
ls ~/config/mercury/mercury_request_handler_config_apimock*.json
# If exists, backup it:
mv ~/config/mercury/mercury_request_handler_config_apimock*.json \
   ~/config/mercury/mercury_request_handler_config_apimock-BACKUP.json
```

### Step 4: Start Services

#### 4.1 Start golang-mock-server

```bash
cd ~/golang-mock-build
nohup ./golang-mock-server -port 9090 -log-level debug > \
    ~/golang-mock-build/golang-mock-server.log 2>&1 &
```

**Verify it's running:**
```bash
ps aux | grep golang-mock-server | grep -v grep
netstat -tlnp | grep 9090
tail -f ~/golang-mock-build/golang-mock-server.log
```

#### 4.2 Restart Adonis

```bash
# Stop Adonis and Mercury
cluster stop adonis mercury

# Wait a few seconds
sleep 5

# Start services
cluster start

# Wait for Adonis to fully start (2-3 minutes)
sleep 120

# Check logs
tail -f /home/nutanix/data/logs/adonis.out | grep -i "golang\|mock\|error"
```

### Step 5: Verify Deployment

#### 5.1 Check Service Status

```bash
# Check golang-mock-server
ps aux | grep golang-mock-server | grep -v grep
netstat -tlnp | grep 9090

# Check Adonis
ps aux | grep adonis | grep java
netstat -tlnp | grep 8888

# Check Mercury
ps aux | grep mercury
netstat -tlnp | grep 9440
```

#### 5.2 Test API Endpoint

```bash
# Get authentication token (replace with your PC credentials)
TOKEN=$(curl -k -X POST https://10.112.90.11:9440/api/nutanix/v3/users/login \
    -H "Content-Type: application/json" \
    -d '{"username":"admin","password":"Nutanix.123"}' | \
    python3 -c "import sys, json; print(json.load(sys.stdin).get('service_account', {}).get('value', ''))")

# Test API
curl -k -X GET "https://10.112.90.11:9440/api/mock/v4.1/config/cats" \
    -H "Authorization: Basic $TOKEN" \
    -H "Content-Type: application/json"
```

**Expected response:**
```json
{
  "data": {
    "catArrayData": {
      "value": [
        {
          "catId": 1,
          "catName": "Cat-1",
          "catType": "TYPE1",
          "description": "A fluffy cat"
        },
        ...
      ]
    }
  }
}
```

---

## Verification

### Pre-Deployment Verification Checklist

- [ ] `golang-mock-grpc-client-17.0.0-SNAPSHOT.jar` exists in `golang-mock-pc/target/`
- [ ] Required classes in nested JAR:
  - [ ] `MockConfigCatController.class`
  - [ ] `GolangmockGrpcConfiguration.class`
  - [ ] `MockConfigCatServiceImpl.class`
  - [ ] `MockConfigCatService.class`
- [ ] `golang-mock-server` binary exists and is Linux amd64
- [ ] `prism-service-17.6.0-SNAPSHOT.jar` contains nested `golang-mock-grpc-client` JAR
- [ ] Generated artifacts copied to PC
- [ ] `application.yaml` updated with correct packages
- [ ] `lookup_cache.json` has correct entry
- [ ] Mercury config created (no conflicts)

### Post-Deployment Verification

- [ ] `golang-mock-server` process running on port 9090
- [ ] Adonis running on port 8888
- [ ] Mercury running on port 9440
- [ ] API endpoint returns 200 OK
- [ ] Response contains cat data
- [ ] No errors in logs

---

## Troubleshooting

### How to Check JAR Contents

#### Method 1: List All Classes

```bash
# For main JAR
jar -tf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar | grep "\.class$" | head -20

# For nested JAR (golang-mock-grpc-client)
cd /tmp
rm -rf BOOT-INF
jar -xf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
    BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar
jar -tf BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar | grep "\.class$"
rm -rf BOOT-INF
```

#### Method 2: Search for Specific Class

```bash
# Check if controller exists
jar -tf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar | \
    grep "MockConfigCatController"

# If not found, check nested JAR
cd /tmp && rm -rf BOOT-INF
jar -xf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
    BOOT-INF/lib/golang-mock-grpc-client-*.jar 2>/dev/null
find BOOT-INF -name "*MockConfigCatController*"
rm -rf BOOT-INF
```

#### Method 3: Extract and Inspect

```bash
# Extract entire JAR
cd /tmp
rm -rf prism-service-extract
mkdir prism-service-extract
cd prism-service-extract
jar -xf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar

# Find specific class
find . -name "MockConfigCatController.class"
find . -name "GolangmockGrpcConfiguration.class"

# Check nested JARs
ls -lh BOOT-INF/lib/golang-mock-grpc-client-*.jar

# Extract nested JAR
cd BOOT-INF/lib
jar -xf golang-mock-grpc-client-17.0.0-SNAPSHOT.jar
find . -name "*.class" | grep -E "Mock|Golangmock"
```

### How to Check Java Source Files

```bash
# Check if source JAR exists
ls -lh ~/ntnx-api-golang-mock-pc/golang-mock-api-codegen/golang-mock-grpc-client/target/*-sources.jar

# Extract and view source
cd /tmp
jar -xf ~/ntnx-api-golang-mock-pc/golang-mock-api-codegen/golang-mock-grpc-client/target/golang-mock-grpc-client-17.0.0-SNAPSHOT-sources.jar
find . -name "MockConfigCatController.java"
cat $(find . -name "MockConfigCatController.java")
```

---

## Common Issues and Solutions

### Issue 1: Missing Classes in JAR

**Symptoms:**
- 404 Not Found when calling API
- Adonis logs show: "No handler found"
- Controller not loaded

**Diagnosis:**
```bash
# Check if classes exist in JAR
cd /tmp && rm -rf BOOT-INF
jar -xf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
    BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar
jar -tf BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar | \
    grep -E "MockConfigCatController|GolangmockGrpcConfiguration"
rm -rf BOOT-INF
```

**Solution:**
1. Rebuild `golang-mock-pc`:
   ```bash
   cd ~/ntnx-api-golang-mock-pc
   mvn clean install -DskipTests -s settings.xml
   ```

2. Verify JAR was rebuilt:
   ```bash
   ls -lh golang-mock-api-codegen/golang-mock-grpc-client/target/*.jar
   ```

3. Rebuild `prism-service`:
   ```bash
   cd ~/ntnx-api-prism-service
   mvn clean install -DskipTests -s settings.xml
   ```

4. Verify nested JAR contains classes (use Method 1 above)

5. Redeploy to PC

### Issue 2: Wrong JAR Version

**Symptoms:**
- Classes exist but Adonis doesn't load them
- Version mismatch in logs

**Diagnosis:**
```bash
# Check pom.xml version
grep "golang-mock-api-controller.version" ~/ntnx-api-prism-service/pom.xml

# Check actual JAR version in target
ls -lh ~/ntnx-api-golang-mock-pc/golang-mock-api-codegen/golang-mock-grpc-client/target/*.jar
```

**Solution:**
1. Ensure `pom.xml` has:
   ```xml
   <golang-mock-api-controller.version>17.0.0-SNAPSHOT</golang-mock-api-controller.version>
   ```

2. Ensure dependency uses this version:
   ```xml
   <dependency>
       <groupId>com.nutanix.nutanix-core.ntnx-api.golang-mock-pc</groupId>
       <artifactId>golang-mock-grpc-client</artifactId>
       <version>${golang-mock-api-controller.version}</version>
   </dependency>
   ```

3. Rebuild both projects

### Issue 3: Adonis Not Loading Controller

**Symptoms:**
- 404 Not Found
- No controller logs in Adonis
- Controller class exists in JAR

**Diagnosis:**
```bash
# Check application.yaml on PC
ssh nutanix@10.112.90.11 "grep -A 5 'adonis.controller.packages.onprem' /home/nutanix/adonis/config/application.yaml"

# Check Adonis logs
ssh nutanix@10.112.90.11 "tail -100 /home/nutanix/data/logs/adonis.out | grep -i 'mock\|golang\|controller'"
```

**Solution:**
1. Verify `application.yaml` has:
   ```yaml
   mock.v4.server.configuration, \
   mock.v4.config.server.controllers, \
   mock.v4.config.server.services, \
   ```

2. Verify package names match exactly (case-sensitive)

3. Restart Adonis:
   ```bash
   cluster stop adonis
   cluster start
   ```

### Issue 4: gRPC Connection Failed (403 Forbidden)

**Symptoms:**
- 403 "Request to Aplos failed"
- Adonis logs show gRPC connection errors
- `golang-mock-server` running but not receiving requests

**Diagnosis:**
```bash
# Check if golang-mock-server is running
ssh nutanix@10.112.90.11 "ps aux | grep golang-mock-server | grep -v grep"

# Check if port 9090 is listening
ssh nutanix@10.112.90.11 "netstat -tlnp | grep 9090"

# Check golang-mock-server logs
ssh nutanix@10.112.90.11 "tail -50 ~/golang-mock-build/golang-mock-server.log"

# Check Adonis logs for gRPC errors
ssh nutanix@10.112.90.11 "tail -100 /home/nutanix/data/logs/adonis.out | grep -i 'grpc\|aplos\|golangmock'"
```

**Solution:**
1. Verify `GolangmockGrpcConfiguration` is in JAR (see Issue 1)

2. Verify `application.yaml` has `grpc.golangmock` config:
   ```yaml
   grpc:
     golangmock:
       host: localhost
       port: 9090
   ```

3. Verify `mock.v4.server.configuration` package is in `adonis.controller.packages.onprem`

4. Check if Adonis is using correct JAR:
   ```bash
   ssh nutanix@10.112.90.11 "ps aux | grep adonis | grep java"
   # Should show: prism-service-17.6.0-SNAPSHOT.jar
   # If shows LKG-RELEASE.jar, backup and remove it
   ```

5. Restart both services:
   ```bash
   # Kill golang-mock-server
   ssh nutanix@10.112.90.11 "pkill -f golang-mock-server"
   
   # Restart golang-mock-server
   ssh nutanix@10.112.90.11 "cd ~/golang-mock-build && nohup ./golang-mock-server -port 9090 -log-level debug > ~/golang-mock-build/golang-mock-server.log 2>&1 &"
   
   # Restart Adonis
   ssh nutanix@10.112.90.11 "cluster stop adonis && cluster start"
   ```

### Issue 5: Method Not Found (UNIMPLEMENTED)

**Symptoms:**
- gRPC error: "UNIMPLEMENTED: unknown method listCats"
- Method name mismatch

**Diagnosis:**
```bash
# Check generated gRPC code
grep "FullMethodName" ~/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/config/cat_service_grpc.pb.go

# Should show: /mock.v4.config.CatService/listCats (lowercase)
# NOT: /mock.v4.config.CatService/ListCats (PascalCase)
```

**Solution:**
1. Regenerate gRPC code:
   ```bash
   cd ~/ntnx-api-golang-mock-pc
   ./generate-grpc.sh
   ```

2. Verify method names are lowercase in `cat_service_grpc.pb.go`

3. Rebuild `golang-mock`:
   ```bash
   cd ~/ntnx-api-golang-mock
   make build
   ```

4. Redeploy `golang-mock-server` to PC

### Issue 6: Import Path Errors in Go Build

**Symptoms:**
- `make build` fails with: "package models/mock/v4/error is not in std"
- Import path errors

**Diagnosis:**
```bash
# Check import paths in generated code
grep "import.*models/" ~/ntnx-api-golang-mock-pc/generated-code/dto/src/models/mock/v4/config/config_model.go
```

**Solution:**
1. Run `publishCode.sh` to fix import paths:
   ```bash
   cd ~/ntnx-api-golang-mock-pc/golang-mock-api-codegen/golang-mock-go-dto-definitions
   ./scripts/publishCode.sh
   ```

2. Manually fix if needed:
   ```bash
   # Fix import paths
   find ~/ntnx-api-golang-mock-pc/generated-code/dto/src/models -name "*.go" -exec \
       sed -i '' 's|"models/|"github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/dto/models/|g' {} +
   ```

3. Rebuild

### Issue 7: Adonis Running with Multiple JARs

**Symptoms:**
- Adonis loads old JAR instead of new one
- Changes not taking effect

**Diagnosis:**
```bash
# Check which JARs Adonis is using
ssh nutanix@10.112.90.11 "ps aux | grep adonis | grep java"
# May show multiple JARs: LKG-RELEASE.jar and SNAPSHOT.jar
```

**Solution:**
1. Backup and remove old JAR:
   ```bash
   ssh nutanix@10.112.90.11 "
     cd /home/nutanix/adonis/lib
     cp prism-service-*-LKG-RELEASE.jar prism-service-LKG-RELEASE-BACKUP-$(date +%Y%m%d).jar
     rm prism-service-*-LKG-RELEASE.jar
   "
   ```

2. Ensure only SNAPSHOT JAR exists:
   ```bash
   ssh nutanix@10.112.90.11 "ls -lh /home/nutanix/adonis/lib/prism-service-*.jar"
   ```

3. Restart Adonis

### Issue 8: Permission Denied Errors

**Symptoms:**
- Cannot write log files
- Cannot start services

**Solution:**
```bash
# Use home directory for logs instead of /var/log
nohup ./golang-mock-server -port 9090 > ~/golang-mock-build/golang-mock-server.log 2>&1 &

# Or use sudo if /var/log is required
sudo nohup ./golang-mock-server -port 9090 > /var/log/golang-mock-server.log 2>&1 &
```

### Issue 9: Port Already in Use

**Symptoms:**
- "address already in use" error
- Cannot start golang-mock-server

**Solution:**
```bash
# Find and kill existing process
ssh nutanix@10.112.90.11 "pkill -f golang-mock-server"

# Verify port is free
ssh nutanix@10.112.90.11 "netstat -tlnp | grep 9090"

# Start server
ssh nutanix@10.112.90.11 "cd ~/golang-mock-build && nohup ./golang-mock-server -port 9090 > ~/golang-mock-build/golang-mock-server.log 2>&1 &"
```

### Issue 10: API Returns 401 Unauthorized

**Symptoms:**
- Authentication fails
- Cannot get token

**Solution:**
1. Use correct PC credentials (may differ from default)
2. Check if PC requires different authentication method
3. Verify token extraction:
   ```bash
   # Test login
   curl -k -X POST https://10.112.90.11:9440/api/nutanix/v3/users/login \
       -H "Content-Type: application/json" \
       -d '{"username":"admin","password":"YOUR_PC_PASSWORD"}'
   
   # Extract token manually if jq not available
   # Look for "value" field in "service_account" object
   ```

---

## Quick Reference: Diagnostic Commands

### Check JAR Contents
```bash
# One-liner to verify all required classes
cd /tmp && rm -rf BOOT-INF && \
jar -xf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
    BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar && \
jar -tf BOOT-INF/lib/golang-mock-grpc-client-17.0.0-SNAPSHOT.jar | \
    grep -E "MockConfigCatController|GolangmockGrpcConfiguration" && \
rm -rf BOOT-INF && echo "✅ All classes found"
```

### Check Service Status on PC
```bash
ssh nutanix@10.112.90.11 "
  echo '=== golang-mock-server ==='
  ps aux | grep golang-mock-server | grep -v grep
  netstat -tlnp | grep 9090
  echo ''
  echo '=== Adonis ==='
  ps aux | grep adonis | grep java
  netstat -tlnp | grep 8888
  echo ''
  echo '=== Mercury ==='
  ps aux | grep mercury
  netstat -tlnp | grep 9440
"
```

### Check Logs
```bash
# golang-mock-server logs
ssh nutanix@10.112.90.11 "tail -50 ~/golang-mock-build/golang-mock-server.log"

# Adonis logs
ssh nutanix@10.112.90.11 "tail -100 /home/nutanix/data/logs/adonis.out | grep -i 'mock\|golang\|error'"
```

---

## Summary

**Key Points:**
1. Always verify JAR contents before deployment
2. Ensure version consistency across all projects
3. Check `application.yaml` package names match exactly
4. Verify `golang-mock-server` is running before testing
5. Check logs for detailed error messages
6. Use lowercase method names in gRPC (camelCase)
7. Backup old JARs before replacing

**Deployment Order:**
1. Build `golang-mock-pc` → Generate code
2. Generate gRPC code → Fix imports
3. Build `golang-mock` → Create binary
4. Build `prism-service` → Create JAR with nested client
5. **Verify JAR contents** → Critical step!
6. Deploy to PC → Copy files
7. Configure PC → Update configs
8. Start services → golang-mock-server, then Adonis
9. Verify → Test API endpoint

---

**Last Updated:** 2025-01-XX  
**Maintained By:** [Your Name]  
**Related Docs:** `SETUP_GOLANG_MOCK_IN_PC.md`, `TROUBLESHOOTING_404.md`


# Complete Deployment Guide - Nexus Service on PC

## Table of Contents
1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Local Build Steps](#local-build-steps)
4. [Adonis Configuration](#adonis-configuration)
5. [PC Deployment Steps](#pc-deployment-steps)
6. [Troubleshooting](#troubleshooting)
7. [Common Issues and Fixes](#common-issues-and-fixes)

---

## Overview

This guide covers the complete deployment of the `ntnx-api-golang-nexus` service to Prism Central (PC), including:
- Building the Go service
- Building and configuring Adonis (REST-to-gRPC gateway)
- Deploying to PC
- Fixing common deployment issues

---

## Prerequisites

### Local Machine
- Go 1.21+
- Maven 3.6+
- Java 21
- Access to Nutanix internal Maven repositories (via VPN or network)
- SSH access to PC

### PC Requirements
- IDF (Insights Data Framework) running on port 2027
- Adonis service directory: `~/adonis/`
- API artifacts directory: `~/api_artifacts/`
- Mercury configuration directory: `~/config/mercury/`

---

## Local Build Steps

### 1. Build API Definitions (golang-nexus-pc)

```bash
cd ~/ntnx-api-golang-nexus-pc

# Build with GitHub authentication
./build-with-auth.sh

# Or manually with Maven
mvn clean install -DskipTests -s settings.xml
```

**Expected Output:**
- `golang-nexus-api-definitions/target/generated-api-artifacts/` - Contains API artifacts
- `golang-nexus-api-codegen/golang-nexus-grpc-client/target/golang-nexus-grpc-client-17.0.0-SNAPSHOT.jar`
- `golang-nexus-api-codegen/golang-nexus-springmvc-interfaces/target/golang-nexus-springmvc-interfaces-17.0.0-SNAPSHOT.jar`
- `golang-nexus-api-codegen/golang-nexus-protobuf-mappers/target/golang-nexus-protobuf-mappers-17.0.0-SNAPSHOT.jar`

### 2. Build Go Service (golang-nexus)

```bash
cd ~/ntnx-api-golang-nexus

# Build the Go binary
go build -o golang-nexus-server ./golang-nexus-service/server
```

**Expected Output:**
- `golang-nexus-server` - Executable binary

### 3. Build Adonis (prism-service)

**IMPORTANT:** Adonis must include these dependencies:
- `golang-nexus-springmvc-interfaces` (for `ItemApiControllerInterface`)
- `golang-nexus-protobuf-mappers` (for `ListItemsApiResponseMapper`)
- `golang-nexus-grpc-client` (for `NexusConfigItemController`)

**File:** `ntnx-api-prism-service/pom.xml`

Ensure these dependencies are present (around line 983):

```xml
<!-- Golang Nexus Spring MVC Interfaces (needed for ItemApiControllerInterface) -->
<dependency>
    <groupId>com.nutanix.nutanix-core.ntnx-api.golang-nexus-pc</groupId>
    <artifactId>golang-nexus-springmvc-interfaces</artifactId>
    <version>${golang-nexus-api-controller.version}</version>
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

<!-- Golang Nexus Protobuf Mappers (needed for ListItemsApiResponseMapper) -->
<dependency>
    <groupId>com.nutanix.nutanix-core.ntnx-api.golang-nexus-pc</groupId>
    <artifactId>golang-nexus-protobuf-mappers</artifactId>
    <version>${golang-nexus-api-controller.version}</version>
    <exclusions>
        <exclusion>
            <artifactId>spring-webmvc</artifactId>
            <groupId>org.springframework</groupId>
        </exclusion>
        <exclusion>
            <artifactId>spring-test</artifactId>
            <groupId>org.springframework</groupId>
        </exclusions>
    </exclusions>
</dependency>

<!-- Golang Nexus gRPC Client  -->
<dependency>
    <groupId>com.nutanix.nutanix-core.ntnx-api.golang-nexus-pc</groupId>
    <artifactId>golang-nexus-grpc-client</artifactId>
    <version>${golang-nexus-api-controller.version}</version>
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
```

**Build Adonis:**
```bash
cd ~/ntnx-api-prism-service
mvn clean install -DskipTests -s settings.xml
```

**Expected Output:**
- `target/prism-service-17.6.0-SNAPSHOT.jar`

**Verify Dependencies in JAR:**
```bash
mkdir -p ~/temp && cd ~/temp
rm -rf BOOT-INF 2>/dev/null
jar -xf ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar BOOT-INF/lib/

# Check interface
jar -tf BOOT-INF/lib/golang-nexus-springmvc-interfaces-17.0.0-SNAPSHOT.jar | grep -i "ItemApiControllerInterface"
# Should show: nexus/v4/config/ItemApiControllerInterface.class

# Check mapper
jar -tf BOOT-INF/lib/golang-nexus-protobuf-mappers-17.0.0-SNAPSHOT.jar | grep -i "ListItemsApiResponseMapper"
# Should show: dp1/mock/nexus/v4/config/mappers/ListItemsApiResponseMapper.class

# Check controller
jar -tf BOOT-INF/lib/golang-nexus-grpc-client-17.0.0-SNAPSHOT.jar | grep -i "NexusConfigItemController"
# Should show: nexus/v4/config/server/controllers/NexusConfigItemController.class

rm -rf BOOT-INF
```

---

## Adonis Configuration

### 1. Update application.yaml

**File:** `ntnx-api-prism-service/src/main/resources/application.yaml`

Ensure these controller packages are included:

```yaml
controller:
  packages:
    - nexus.v4.server.configuration
    - nexus.v4.config.server.controllers
    - nexus.v4.config.server.services
```

### 2. Update lookup_cache.json (on PC)

**File:** `/home/nutanix/api_artifacts/lookup_cache.json`

Add the Nexus API mapping:

```json
{
  "routeMappings": [
    // ... existing entries ...,
    {
      "apiPath": "/nexus/v4.1/config",
      "artifactPath": "nexus/v4.r1.a1/golang-nexus-api-definitions-17.0.0-SNAPSHOT"
    }
  ]
}
```

### 3. Create Mercury Configuration (on PC)

**File:** `~/config/mercury/mercury_request_handler_config_golangnexus.json`

```json
{
  "api_path_config_list" : [
    {
      "api_path" : "/api/nexus/v4.1",
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
      "api_path" : "/api/nexus/unversioned",
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

---

## PC Deployment Steps

### 1. Copy Go Service Binary

```bash
PC_IP="10.124.86.25"  # Update with your PC IP

# Create directory on PC
ssh nutanix@${PC_IP} "mkdir -p ~/golang-nexus-build"

# Copy binary
scp -O ~/ntnx-api-golang-nexus/golang-nexus-server nutanix@${PC_IP}:~/golang-nexus-build/
```

### 2. Copy API Artifacts

```bash
# Create directory structure on PC
ssh nutanix@${PC_IP} "mkdir -p ~/api_artifacts/nexus/v4.r1.a1/golang-nexus-api-definitions-17.0.0-SNAPSHOT"

# Copy artifacts
scp -r -O ~/ntnx-api-golang-nexus-pc/golang-nexus-api-definitions/target/generated-api-artifacts/* \
  nutanix@${PC_IP}:~/api_artifacts/nexus/v4.r1.a1/golang-nexus-api-definitions-17.0.0-SNAPSHOT/
```

**Verify artifacts include:**
- `swagger-all-17.0.0-SNAPSHOT.yaml`
- `*.proto` files (e.g., `config.proto`, `item_service.proto`)
- `api-manifest-17.0.0-SNAPSHOT.json`
- `object-type-mapping-17.0.0-SNAPSHOT.yaml`

### 3. Update lookup_cache.json on PC

```bash
ssh nutanix@${PC_IP}

# Backup existing file
cp ~/api_artifacts/lookup_cache.json ~/api_artifacts/lookup_cache.json.backup

# Edit file (use vi/nano)
vi ~/api_artifacts/lookup_cache.json

# Add entry in routeMappings array:
# {
#   "apiPath": "/nexus/v4.1/config",
#   "artifactPath": "nexus/v4.r1.a1/golang-nexus-api-definitions-17.0.0-SNAPSHOT"
# }
```

### 4. Create Mercury Config on PC

```bash
ssh nutanix@${PC_IP}

# Create directory
mkdir -p ~/config/mercury

# Create config file
cat > ~/config/mercury/mercury_request_handler_config_golangnexus.json << 'EOF'
{
  "api_path_config_list" : [
    {
      "api_path" : "/api/nexus/v4.1",
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
      "api_path" : "/api/nexus/unversioned",
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
EOF
```

### 5. Copy Adonis JAR to PC

```bash
# Stop Adonis first (on PC)
ssh nutanix@${PC_IP} "genesis stop adonis mercury"
sleep 10

# Remove backup JARs
ssh nutanix@${PC_IP} "rm -f ~/adonis/lib/*-BACKUP-*.jar"

# Copy new JAR
scp -O ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
  nutanix@${PC_IP}:~/adonis/lib/

# Verify only SNAPSHOT JAR exists
ssh nutanix@${PC_IP} "ls -lh ~/adonis/lib/prism-service-*.jar"
# Should show only: prism-service-17.6.0-SNAPSHOT.jar
```

### 6. Update application.yaml on PC

```bash
ssh nutanix@${PC_IP}

# Backup existing file
cp ~/adonis/config/application.yaml ~/adonis/config/application.yaml.backup

# Edit file (use vi/nano)
vi ~/adonis/config/application.yaml

# Ensure these packages are in controller.packages:
# - nexus.v4.server.configuration
# - nexus.v4.config.server.controllers
# - nexus.v4.config.server.services
```

### 7. Start Services on PC

```bash
ssh nutanix@${PC_IP}

# Start Go service (if not already running)
cd ~/golang-nexus-build
nohup ./golang-nexus-server > /dev/null 2>&1 &

# Start Adonis
cluster start
sleep 30

# Check status
genesis status adonis
```

### 8. Verify Deployment

```bash
ssh nutanix@${PC_IP}

# 1. Check Adonis logs for controller registration
tail -100 /home/nutanix/adonis/logs/prism-service.log | grep -i "nexus\|controller" | tail -10

# 2. Check Go service is running
ps aux | grep golang-nexus-server | grep -v grep

# 3. Test API endpoint
curl -k https://127.0.0.1:9440/api/nexus/v4.1/config/items
```

---

## Troubleshooting

### Verify JAR Contents on PC

```bash
ssh nutanix@${PC_IP}

mkdir -p ~/temp
cd ~/temp
rm -rf BOOT-INF 2>/dev/null

# Extract JAR
jar -xf ~/adonis/lib/prism-service-17.6.0-SNAPSHOT.jar BOOT-INF/lib/

# Check interface
jar -tf BOOT-INF/lib/golang-nexus-springmvc-interfaces-17.0.0-SNAPSHOT.jar | grep -i "ItemApiControllerInterface"
# Should show: nexus/v4/config/ItemApiControllerInterface.class

# Check mapper
jar -tf BOOT-INF/lib/golang-nexus-protobuf-mappers-17.0.0-SNAPSHOT.jar | grep -i "ListItemsApiResponseMapper"
# Should show: dp1/mock/nexus/v4/config/mappers/ListItemsApiResponseMapper.class

# Check controller
jar -tf BOOT-INF/lib/golang-nexus-grpc-client-17.0.0-SNAPSHOT.jar | grep -i "NexusConfigItemController"
# Should show: nexus/v4/config/server/controllers/NexusConfigItemController.class

rm -rf BOOT-INF
```

### Check Which JAR Adonis is Using

```bash
ssh nutanix@${PC_IP}
ps aux | grep -i 'prism-service.*jar' | grep -v grep
# Should show: prism-service-17.6.0-SNAPSHOT.jar
```

### Check Adonis Logs

```bash
ssh nutanix@${PC_IP}

# Recent errors
tail -100 /home/nutanix/adonis/logs/prism-service.log | grep -i "error\|exception" | tail -10

# Startup messages
tail -100 /home/nutanix/adonis/logs/prism-service.log | grep -i "started\|ready" | tail -5

# Controller registration
tail -200 /home/nutanix/adonis/logs/prism-service.log | grep -i "nexus\|controller" | tail -10
```

---

## Common Issues and Fixes

### Issue 1: Missing ItemApiControllerInterface

**Error:**
```
java.io.FileNotFoundException: class path resource [nexus/v4/config/ItemApiControllerInterface.class] cannot be opened because it does not exist
```

**Solution:**
1. Ensure `golang-nexus-springmvc-interfaces` dependency is in Adonis's `pom.xml`
2. Rebuild Adonis: `mvn clean install -DskipTests -s settings.xml`
3. Copy new JAR to PC
4. Restart Adonis

### Issue 2: Missing ListItemsApiResponseMapper

**Error:**
```
java.lang.ClassNotFoundException: dp1.mock.nexus.v4.config.mappers.ListItemsApiResponseMapper
```

**Solution:**
1. Ensure `golang-nexus-protobuf-mappers` dependency is in Adonis's `pom.xml`
2. Rebuild Adonis: `mvn clean install -DskipTests -s settings.xml`
3. Copy new JAR to PC
4. Restart Adonis

### Issue 3: Adonis Using Backup JAR

**Symptom:** Adonis is using `prism-service-*-BACKUP-*.jar` instead of SNAPSHOT

**Solution:**
```bash
ssh nutanix@${PC_IP}
genesis stop adonis mercury
sleep 10
rm -f ~/adonis/lib/*-BACKUP-*.jar
# Or move it: mkdir -p ~/adonis/lib/backup-jars && mv ~/adonis/lib/*-BACKUP-*.jar ~/adonis/lib/backup-jars/
cluster start
```

### Issue 4: Missing .proto Files in API Artifacts

**Symptom:** `.proto` files are missing from `api_artifacts` directory

**Solution:**
1. Check `golang-nexus-api-definitions/pom.xml` has `maven-antrun-plugin` to copy `.proto` files
2. Rebuild: `mvn clean install -DskipTests`
3. Verify `.proto` files are in `target/generated-api-artifacts/`
4. Copy artifacts to PC again

### Issue 5: 404 Error on API Endpoint

**Checklist:**
1. ✅ Controller in JAR: `jar -tf JAR | grep NexusConfigItemController`
2. ✅ `application.yaml` has controller packages
3. ✅ `lookup_cache.json` has API path mapping
4. ✅ API artifacts are in correct directory
5. ✅ Mercury config exists
6. ✅ Adonis restarted after changes

### Issue 6: /tmp Directory Full on PC

**Error:** `No space left on device` when extracting JAR

**Solution:**
```bash
# Use ~/temp instead of /tmp
mkdir -p ~/temp
cd ~/temp
jar -xf ~/adonis/lib/prism-service-17.6.0-SNAPSHOT.jar BOOT-INF/lib/
```

### Issue 7: Maven Build Fails - Network/Proxy Issues

**Error:** `proxy.nutanix.com: nodename nor servname provided, or not known`

**Solution:**
1. Connect to VPN (if available)
2. Use project's `settings.xml`: `mvn clean install -DskipTests -s settings.xml`
3. Or build on a machine with Nutanix network access

---

## Quick Reference

### Build Commands
```bash
# API Definitions
cd ~/ntnx-api-golang-nexus-pc && ./build-with-auth.sh

# Go Service
cd ~/ntnx-api-golang-nexus && go build -o golang-nexus-server ./golang-nexus-service/server

# Adonis
cd ~/ntnx-api-prism-service && mvn clean install -DskipTests -s settings.xml
```

### Copy to PC
```bash
PC_IP="10.124.86.25"

# Go binary
scp -O ~/ntnx-api-golang-nexus/golang-nexus-server nutanix@${PC_IP}:~/golang-nexus-build/

# API artifacts
scp -r -O ~/ntnx-api-golang-nexus-pc/golang-nexus-api-definitions/target/generated-api-artifacts/* \
  nutanix@${PC_IP}:~/api_artifacts/nexus/v4.r1.a1/golang-nexus-api-definitions-17.0.0-SNAPSHOT/

# Adonis JAR
scp -O ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
  nutanix@${PC_IP}:~/adonis/lib/
```

### Restart Services on PC
```bash
ssh nutanix@${PC_IP}
genesis stop adonis mercury
sleep 10
cluster start
sleep 30
genesis status adonis
```

### Test API
```bash
curl -k https://10.124.86.25:9440/api/nexus/v4.1/config/items
```

---

## Summary of Required Dependencies

For Adonis to work correctly, these dependencies must be in `pom.xml`:

1. **golang-nexus-springmvc-interfaces** - Provides `ItemApiControllerInterface`
2. **golang-nexus-protobuf-mappers** - Provides `ListItemsApiResponseMapper`
3. **golang-nexus-grpc-client** - Provides `NexusConfigItemController`

All three are required. Missing any one will cause Adonis to fail to start.

---

**Last Updated:** December 2025  
**Version:** 1.0

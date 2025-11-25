# Setup golang-mock in PC

**Last Updated**: 2025-11-25

---

## Prerequisites

- Java, Maven, Go installed
- Access to PC (SSH)
- GitHub access configured

---

## Step 1: Build golang-mock-pc

```bash
cd ~/ntnx-api-golang-mock-pc
mvn clean install -s settings.xml -DskipTests
```

**Expected**: JARs in `target/` directories, generated code in `generated-code/`

---

## Step 2: Generate gRPC Code

```bash
cd ~/ntnx-api-golang-mock-pc
./generate-grpc.sh
```

**Expected**: `.pb.go` files in `generated-code/protobuf/mock/v4/config/`

---

## Step 3: Build golang-mock Server

```bash
cd ~/ntnx-api-golang-mock
make build
```

**Expected**: `golang-mock-server` binary (Linux)

---

## Step 4: Build prism-service (Adonis)

### 4.1 Update pom.xml

Add to `<properties>`:
```xml
<golang-mock-api-controller.version>17.0.0-SNAPSHOT</golang-mock-api-controller.version>
```

Add to `<dependencies>`:
```xml
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
```

Add to `<executions>` in `spring-boot-maven-plugin`:
```xml
<controller.golang-mock.version>${golang-mock-api-controller.version}</controller.golang-mock.version>
```

### 4.2 Update application.yaml

Add to `adonis.controller.packages.onprem`:
```yaml
mock.v4.server.configuration, \
mock.v4.config.server.controllers, \
mock.v4.config.server.services
```

Add to `grpc:` section:
```yaml
golangmock:
  host: localhost
  port: 9090
```

### 4.3 Build

```bash
cd ~/ntnx-api-prism-service
mvn clean install -DskipTests -s settings.xml
```

**Expected**: `target/prism-service-17.6.0-SNAPSHOT.jar`

---

## Step 5: PC Configuration

### 5.1 Create Directories

```bash
ssh nutanix@10.112.90.11
mkdir -p ~/golang-mock-build
mkdir -p ~/api_artifacts/mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT
```

### 5.2 Update application.yaml

**File**: `/home/nutanix/adonis/config/application.yaml`

Add to `adonis.controller.packages.onprem`:
```yaml
mock.v4.server.configuration, \
mock.v4.config.server.controllers, \
mock.v4.config.server.services
```

Add to `grpc:` section:
```yaml
golangmock:
  host: localhost
  port: 9090
```

### 5.3 Update lookup_cache.json

**File**: `/home/nutanix/api_artifacts/lookup_cache.json`

Add entry:
```json
{
  "apiPath": "/mock/v4.1/config",
  "artifactPath": "mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT"
}
```

### 5.4 Create Mercury Config

**File**: `~/config/mercury/mercury_request_handler_config_golangmock.json`

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

---

## Step 6: Deploy Files to PC

### 6.1 Copy golang-mock-server Binary

```bash
# Stop old server first
ssh nutanix@10.112.90.11 "pkill -f golang-mock-server"

# Copy binary
scp -O ~/ntnx-api-golang-mock/golang-mock-server \
   nutanix@10.112.90.11:/home/nutanix/golang-mock-build/

# On PC: make executable and start
ssh nutanix@10.112.90.11 << 'EOF'
cd ~/golang-mock-build
chmod +x golang-mock-server
nohup ./golang-mock-server -port 9090 -log-level debug > ~/golang-mock-build/golang-mock-server.log 2>&1 &
EOF
```

### 6.2 Copy prism-service JAR

```bash
# Backup old JAR on PC
ssh nutanix@10.112.90.11 "cp /home/nutanix/adonis/lib/prism-service-*.jar /home/nutanix/adonis/lib/prism-service-*.jar.backup 2>/dev/null || true"

# Backup LKG-RELEASE JAR if exists
ssh nutanix@10.112.90.11 "mv /home/nutanix/adonis/lib/prism-service-*-LKG-RELEASE.jar /home/nutanix/adonis/lib/prism-service-*-LKG-RELEASE.jar.backup 2>/dev/null || true"

# Copy new JAR
scp -O ~/ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
   nutanix@10.112.90.11:/home/nutanix/adonis/lib/
```

### 6.3 Copy API Artifacts

```bash
scp -O -r ~/ntnx-api-golang-mock-pc/golang-mock-api-definitions/target/generated-api-artifacts/* \
   nutanix@10.112.90.11:/home/nutanix/api_artifacts/mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT/
```

---

## Step 7: Start Services on PC

```bash
ssh nutanix@10.112.90.11

# Verify golang-mock-server is running
ps aux | grep golang-mock-server | grep -v grep
netstat -tlnp | grep 9090

# Restart Adonis and Mercury
genesis stop adonis mercury && cluster start

# Wait 5 minutes for services to start
```

---

## Step 8: Verify Deployment

### 8.1 Check Services

```bash
# Check golang-mock-server
ps aux | grep golang-mock-server | grep -v grep
netstat -tlnp | grep 9090

# Check Adonis
ps aux | grep adonis | grep java | grep -v grep
netstat -tlnp | grep 8888

# Check Mercury
ps aux | grep mercury | grep -v grep
netstat -tlnp | grep 9440
```

### 8.2 Test API

```bash
# Get authentication token
TOKEN=$(curl -X POST https://10.112.90.11:9440/api/nutanix/v3/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"<password>"}' | jq -r .token)

# Test endpoint
curl -k -H "Authorization: Bearer $TOKEN" \
  https://10.112.90.11:9440/api/mock/v4.1/config/cats
```

**Expected**: JSON response with list of cats

---

## Troubleshooting

### 404 NOT_FOUND

1. **Check JAR has controller classes**:
   ```bash
   jar -tf /home/nutanix/adonis/lib/prism-service-17.6.0-SNAPSHOT.jar | grep "golang-mock-grpc-client"
   ```

2. **Check application.yaml has packages**:
   ```bash
   grep -A 3 "mock.v4" /home/nutanix/adonis/config/application.yaml
   ```

3. **Check which JAR Adonis is using**:
   ```bash
   ps aux | grep adonis | grep java | grep -o "prism-service.*jar"
   ```

4. **If LKG-RELEASE JAR exists, backup it**:
   ```bash
   mv /home/nutanix/adonis/lib/prism-service-*-LKG-RELEASE.jar \
      /home/nutanix/adonis/lib/prism-service-*-LKG-RELEASE.jar.backup
   ```

5. **Restart Adonis**:
   ```bash
   genesis stop adonis mercury && cluster start
   ```

### UNIMPLEMENTED Error

- Ensure `golang-mock-server` binary is rebuilt after running `generate-grpc.sh`
- Verify binary has lowercase method names: `MethodName: "listCats"` (not `"ListCats"`)

### Text File Busy

- Stop the running server before copying new binary:
  ```bash
  ssh nutanix@10.112.90.11 "pkill -f golang-mock-server"
  ```

---

## Quick Reference

### File Locations on PC

| File/Directory | Location |
|----------------|----------|
| golang-mock-server | `/home/nutanix/golang-mock-build/golang-mock-server` |
| prism-service JAR | `/home/nutanix/adonis/lib/prism-service-*.jar` |
| Adonis config | `/home/nutanix/adonis/config/application.yaml` |
| API artifacts | `/home/nutanix/api_artifacts/mock/v4.r1.a1/golang-mock-api-definitions-17.0.0-SNAPSHOT/` |
| lookup_cache.json | `/home/nutanix/api_artifacts/lookup_cache.json` |
| Mercury config | `~/config/mercury/mercury_request_handler_config_golangmock.json` |

### Key Commands

```bash
# Start golang-mock-server
cd /home/nutanix/golang-mock-build
nohup ./golang-mock-server -port 9090 -log-level debug > ~/golang-mock-build/golang-mock-server.log 2>&1 &

# Restart Adonis/Mercury
genesis stop adonis mercury && cluster start

# Test API
curl -k -H "Authorization: Bearer $TOKEN" \
  https://10.112.90.11:9440/api/mock/v4.1/config/cats
```

---

## Version Format

- **API URL**: `v4.1` (e.g., `/api/mock/v4.1/config/cats`)
- **Directory Path**: `v4.r1.a1` (e.g., `~/api_artifacts/mock/v4.r1.a1/...`)

---

**Last Updated**: 2025-11-25

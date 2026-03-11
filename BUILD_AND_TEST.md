# Nexus gRPC Service - Build and Test Guide

## Overview

This guide covers building, deploying, and testing the Nexus gRPC APIs (v4.1) on Prism Central.

**Architecture:**
```
Client (REST) → Mercury → Adonis (REST-to-gRPC) → Go Service (gRPC :9090) → IDF / StatsGW
```

**Key Components:**
| Component | Repository | Purpose |
|---|---|---|
| API Definitions & Codegen | `ntnx-api-golang-nexus-pc` | YAML models, proto generation, Java DTOs, mappers |
| Go Service | `ntnx-api-golang-nexus` | gRPC server implementation (items, stats, file transfer) |
| Adonis (Prism Service) | `ntnx-api-prism-service` | REST-to-gRPC gateway |
| Go Utilities | `ntnx-api-utils-go` | Shared Go utility libraries |

---

## Prerequisites

### Required Software
- Java 21+
- Go 1.21+
- Maven 3.6+
- Git

### Certificate Setup (Do This First)
Before cloning or building, install certificates to access GitHub and Caneveral on Mac.
This prevents PKIX failure errors during Maven builds.

> Follow: *Install Certificates to access Github and Caneveral on Mac*

### Clone Repositories
```bash
mkdir -p ~/groupby_aggregate_mock && cd ~/groupby_aggregate_mock

git clone https://github.com/nitinmangotra1212/ntnx-api-golang-nexus-pc.git
git clone https://github.com/nitinmangotra1212/ntnx-api-golang-nexus.git
git clone https://github.com/nutanix-core/ntnx-api-prism-service.git
git clone https://github.com/nutanix-core/ntnx-api-utils-go.git
```

---

## Build Steps

> **Build order matters:** Step 1 must complete before Steps 2 and 3 (which can run in parallel).

### Step 1: Build API Definitions (`ntnx-api-golang-nexus-pc`)

Generates API artifacts, proto files, Java DTOs, protobuf mappers, and gRPC client stubs.

```bash
cd ntnx-api-golang-nexus-pc

export MAVEN_OPTS="--add-opens java.base/java.lang=ALL-UNNAMED \
  --add-opens java.base/java.lang.reflect=ALL-UNNAMED \
  --add-opens java.base/java.util=ALL-UNNAMED"

mvn clean install -DskipTests -s settings.xml
```

**Generated artifacts:**
- `target/generated-api-artifacts/` — API YAML, proto files, manifest
- `generated-code/protobuf/` — Proto message definitions
- `golang-nexus-api-codegen/golang-nexus-grpc-client/target/*.jar` — gRPC client JAR
- `golang-nexus-api-codegen/golang-nexus-protobuf-mappers/target/*.jar` — Protobuf mapper JAR

### Step 2: Build Go Service (`ntnx-api-golang-nexus`)

```bash
cd ntnx-api-golang-nexus

GOOS=linux GOARCH=amd64 go build -o golang-nexus-server golang-nexus-service/server/main.go
```

Produces a Linux binary: `golang-nexus-server`

### Step 3: Build Prism Service / Adonis (`ntnx-api-prism-service`)

#### Pre-build Configuration

**a) Update `application.yaml`** (`src/main/resources/application.yaml`):
```yaml
adonis:
  controller:
    packages:
      onprem:
        # ... existing packages ...,
        nexus.v4.server.configuration,
        nexus.v4.config.server.controllers,
        nexus.v4.config.server.services

grpc:
  golangnexus:
    host: localhost
    port: 9090
```

**b) Update `pom.xml`** — Add version property:
```xml
<properties>
  <!-- ... existing properties ... -->
  <golang-nexus-api-controller.version>17.0.0-SNAPSHOT</golang-nexus-api-controller.version>
</properties>
```

Add controller version mapping (in `<controllerVersions>` section):
```xml
<controller.golang-nexus.version>${golang-nexus-api-controller.version}</controller.golang-nexus.version>
```

**c) Verify dependencies in `pom.xml`** — All three are required:
```xml
<dependency>
  <groupId>com.nutanix.nutanix-core.ntnx-api.golang-nexus-pc</groupId>
  <artifactId>golang-nexus-springmvc-interfaces</artifactId>
  <version>${golang-nexus-api-controller.version}</version>
</dependency>
<dependency>
  <groupId>com.nutanix.nutanix-core.ntnx-api.golang-nexus-pc</groupId>
  <artifactId>golang-nexus-protobuf-mappers</artifactId>
  <version>${golang-nexus-api-controller.version}</version>
</dependency>
<dependency>
  <groupId>com.nutanix.nutanix-core.ntnx-api.golang-nexus-pc</groupId>
  <artifactId>golang-nexus-grpc-client</artifactId>
  <version>${golang-nexus-api-controller.version}</version>
</dependency>
```

**d) Build:**
```bash
cd ntnx-api-prism-service

export MAVEN_OPTS="--add-opens java.base/java.lang=ALL-UNNAMED \
  --add-opens java.base/java.lang.reflect=ALL-UNNAMED \
  --add-opens java.base/java.util=ALL-UNNAMED"

mvn clean install -DskipTests -s settings.xml
```

Produces: `target/prism-service-17.6.0-SNAPSHOT.jar`

---

## Deployment Steps (On PC)

### Prerequisites on PC
- Prism Central VM with SSH access (`nutanix@<PC_IP>`)
- IDF service running
- Adonis service directory: `~/adonis/`

Set the PC IP for all subsequent commands:
```bash
export PC_IP="10.114.163.60"
```

### Step 1: Deploy Go Service Binary

```bash
ssh nutanix@${PC_IP} "mkdir -p ~/golang-nexus-build"

scp -O golang-nexus-server nutanix@${PC_IP}:~/golang-nexus-build/

ssh nutanix@${PC_IP} "chmod +x ~/golang-nexus-build/golang-nexus-server"
```

### Step 2: Deploy API Artifacts

```bash
ssh nutanix@${PC_IP} "mkdir -p ~/api_artifacts/nexus/v4.r1.a1/golang-nexus-api-definitions-17.0.0-SNAPSHOT"

scp -r -O ntnx-api-golang-nexus-pc/golang-nexus-api-definitions/target/generated-api-artifacts/* \
  nutanix@${PC_IP}:~/api_artifacts/nexus/v4.r1.a1/golang-nexus-api-definitions-17.0.0-SNAPSHOT/
```

Verify on PC:
```bash
ls -la ~/api_artifacts/nexus/v4.r1.a1/golang-nexus-api-definitions-17.0.0-SNAPSHOT/
# Should contain: swagger-all-*.yaml, *.proto files, api-manifest-*.json
```

### Step 3: Update `lookup_cache.json` on PC

```bash
ssh nutanix@${PC_IP}
cp ~/api_artifacts/lookup_cache.json ~/api_artifacts/lookup_cache.json.backup
vi ~/api_artifacts/lookup_cache.json
```

Add entry in `routeMappings` array:
```json
{
  "apiPath": "/nexus/v4.1/config",
  "artifactPath": "nexus/v4.r1.a1/golang-nexus-api-definitions-17.0.0-SNAPSHOT"
}
```

### Step 4: Create Mercury Configuration on PC

```bash
ssh nutanix@${PC_IP}
mkdir -p ~/config/mercury

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

### Step 5: Deploy Prism Service JAR

```bash
ssh nutanix@${PC_IP} "genesis stop adonis mercury"
sleep 10

ssh nutanix@${PC_IP} "rm -f ~/adonis/lib/*-BACKUP-*.jar"

scp -O ntnx-api-prism-service/target/prism-service-17.6.0-SNAPSHOT.jar \
  nutanix@${PC_IP}:~/adonis/lib/
```

### Step 6: Update `application.yaml` on PC

```bash
ssh nutanix@${PC_IP}
cp ~/adonis/config/application.yaml ~/adonis/config/application.yaml.backup
vi ~/adonis/config/application.yaml
```

Ensure controller packages and gRPC config are present (same as local config in Build Step 3a).

### Step 7: Register IDF Entity Types (One-time setup)

```bash
scp -O ntnx-api-golang-nexus-pc/register_create_item_idf.py \
  nutanix@${PC_IP}:~/ntnx-api-golang-nexus-pc/

ssh nutanix@${PC_IP} "cd ~/ntnx-api-golang-nexus-pc && python3 register_create_item_idf.py"
```

This registers entity types (`item`, `item_stats`) and populates test data (120 items + 120 item_stats).

### Step 8: Start Services

```bash
ssh nutanix@${PC_IP}

# Start Go service
cd ~/golang-nexus-build
pkill -f golang-nexus-server || true
nohup ./golang-nexus-server > golang-nexus-server.log 2>&1 &

# Start Adonis + Mercury
genesis stop adonis mercury && sleep 30 && cluster start && sleep 50
```

### Step 9: Verify Deployment

```bash
# Check Go service
ps aux | grep golang-nexus-server | grep -v grep

# Check Adonis logs for controller registration
tail -100 ~/adonis/logs/prism-service.log | grep -i "nexus\|controller" | tail -10

# Test API endpoint
curl -k https://127.0.0.1:9440/api/nexus/v4.1/config/items -u 'admin:Nutanix.123'
```

---

## Quick Re-deploy (Code Changes Only)

When you only changed Go service code:
```bash
# Local: rebuild
cd ntnx-api-golang-nexus
GOOS=linux GOARCH=amd64 go build -o golang-nexus-server golang-nexus-service/server/main.go

# Deploy and restart
scp -O golang-nexus-server nutanix@${PC_IP}:~/golang-nexus-build/
ssh nutanix@${PC_IP} "cd /home/nutanix/golang-nexus-build && pkill -f golang-nexus-server || true && nohup ./golang-nexus-server > golang-nexus-server.log 2>&1 &"
```

When you changed API definitions (YAML models, proto):
```bash
# Local: rebuild all three
cd ntnx-api-golang-nexus-pc && mvn clean install -DskipTests -s settings.xml
cd ntnx-api-golang-nexus && GOOS=linux GOARCH=amd64 go build -o golang-nexus-server golang-nexus-service/server/main.go
cd ntnx-api-prism-service && mvn clean install -DskipTests -s settings.xml

# Deploy artifacts + binary + JAR, then restart services
```

When you only changed Adonis/Prism configuration:
```bash
# Deploy new JAR
scp -O ntnx-api-prism-service/target/prism-service-*.jar nutanix@${PC_IP}:~/adonis/lib/

# Restart Adonis
ssh nutanix@${PC_IP} "genesis stop adonis mercury && sleep 30 && cluster start && sleep 50"
```

---

## API Endpoints

**Base URL:** `https://<PC_IP>:9440/api/nexus/v4.1`

### Items API

| Method | Endpoint | Description |
|---|---|---|
| GET | `/config/items` | List items with OData query support |
| GET | `/config/items/{extId}` | Get a single item by extId |

### OData Query Examples

```bash
# Basic list with pagination
curl -k "https://${PC_IP}:9440/api/nexus/v4.1/config/items?\$limit=10&\$page=0" -u 'admin:Nutanix.123'

# Filter + Select + OrderBy
curl -k "https://${PC_IP}:9440/api/nexus/v4.1/config/items?\$filter=isActive eq true&\$select=itemName,price&\$orderby=price" -u 'admin:Nutanix.123'

# Expand with time-series stats
curl -k "https://${PC_IP}:9440/api/nexus/v4.1/config/items?\$expand=itemStats(\$startTime=2026-02-01T01:00:00Z;\$endTime=2026-02-22T05:30:00Z;\$statType=AVG;\$samplingInterval=25200)&\$limit=10" -u 'admin:Nutanix.123'

# Expand with nested $select, $orderby
curl -k "https://${PC_IP}:9440/api/nexus/v4.1/config/items?\$expand=itemStats(\$startTime=2026-02-01T01:00:00Z;\$endTime=2026-02-22T05:30:00Z;\$statType=AVG;\$samplingInterval=25200;\$select=age;\$orderby=age desc)&\$limit=5" -u 'admin:Nutanix.123'

# GroupBy with aggregation
curl -k "https://${PC_IP}:9440/api/nexus/v4.1/config/items?\$apply=groupby((itemType),aggregate(itemId with count as totalCount))&\$limit=10" -u 'admin:Nutanix.123'

# GroupBy + Expand
curl -k "https://${PC_IP}:9440/api/nexus/v4.1/config/items?\$expand=itemStats(\$startTime=2026-02-01T01:00:00Z;\$endTime=2026-02-22T05:30:00Z;\$statType=AVG;\$samplingInterval=25200)&\$apply=groupby((itemType),aggregate(itemId with count as totalCount))&\$limit=10" -u 'admin:Nutanix.123'

# Lambda filter on time-series metric
curl -k "https://${PC_IP}:9440/api/nexus/v4.1/config/items?\$expand=itemStats(\$startTime=2026-02-01T01:00:00Z;\$endTime=2026-02-22T05:30:00Z;\$statType=AVG;\$samplingInterval=25200;\$filter=age/any(t:t eq 10))" -u 'admin:Nutanix.123'
```

### File Transfer API

| Method | Endpoint | Description |
|---|---|---|
| POST | `/config/upload` | Upload file (streaming, max 1GB) |
| GET | `/config/download/{extId}` | Download file by extId |

---

## Troubleshooting

| Issue | Possible Cause | Solution |
|---|---|---|
| Upload/List returns 503 | Go service not running, Mercury config missing | Check `ps aux \| grep golang-nexus-server`, verify Mercury config |
| 404 on API endpoint | `lookup_cache.json` missing entry, controller packages not configured | Verify lookup_cache, check `application.yaml` packages |
| Build fails with "cannot find symbol" | Missing Adonis dependency | Verify all 3 nexus dependencies in `pom.xml` |
| Adonis using BACKUP JAR | Old JAR present | `rm -f ~/adonis/lib/*-BACKUP-*.jar` then restart |
| Missing proto files | Codegen incomplete | Rebuild `ntnx-api-golang-nexus-pc`, re-deploy artifacts |
| `NullPointerException` in MapstructMapperGenerator | Proto field index gap (e.g., starting at 4002 leaves gap at 4001) | Start `x-codegen-hint` identifiers at index 2001 |

### Useful Log Locations

| Log | Path |
|---|---|
| Go service | `~/golang-nexus-build/golang-nexus-server.log` |
| Adonis | `~/adonis/logs/prism-service.log` |

### Restart Commands

```bash
# Restart Go service only
ssh nutanix@${PC_IP} "cd /home/nutanix/golang-nexus-build && pkill -f golang-nexus-server || true && nohup ./golang-nexus-server > golang-nexus-server.log 2>&1 &"

# Restart Adonis + Mercury
ssh nutanix@${PC_IP} "genesis stop adonis mercury && sleep 30 && cluster start && sleep 50"
```

---

## Deployment Checklist

### Before Deploy
- [ ] All three repositories built successfully
- [ ] Go binary created (`golang-nexus-server`)
- [ ] API artifacts generated (`target/generated-api-artifacts/`)
- [ ] Prism-service JAR includes nexus dependencies

### After Deploy
- [ ] Go service running (`ps aux | grep golang-nexus-server`)
- [ ] API artifacts copied to PC
- [ ] `lookup_cache.json` updated with nexus path
- [ ] Mercury config file created
- [ ] `application.yaml` has controller packages
- [ ] Prism-service JAR deployed (SNAPSHOT version, no BACKUP JARs)
- [ ] IDF entity types registered
- [ ] Adonis restarted successfully
- [ ] API endpoints responding (test with curl)

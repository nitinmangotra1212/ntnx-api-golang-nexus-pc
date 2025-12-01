# Deploy OData Parser Implementation to PC

## ✅ Pre-Deployment Checklist

- [x] OData parser implemented
- [x] Exception handler implemented
- [x] Build successful
- [x] Code aligned with az-manager

---

## Step 1: Rebuild Prism-Service JAR (if needed)

**Check:** Since we added OData dependency, the Java side might need updates. However, since OData parsing is only in Go code, the JAR might not need rebuilding. But let's check if protobuf changed.

**If protobuf didn't change:**
```bash
# Skip this step - JAR is fine
```

**If you want to be safe:**
```bash
cd ~/ntnx-api-prism-service
mvn clean install -DskipTests -s settings.xml
```

**Verify JAR:**
```bash
ls -lh target/prism-service-*.jar
```

---

## Step 2: Copy Binary to PC

**On your local machine:**
```bash
cd ~/ntnx-api-golang-nexus
scp golang-nexus-server nutanix@10.124.86.254:~/golang-nexus-build/
```

**Example:**
```bash
scp golang-nexus-server nutanix@10.124.86.254:~/golang-nexus-build/
```

---

## Step 3: Deploy on PC

**On PC, SSH into it:**
```bash
ssh nutanix@10.124.86.254
```

**Stop old service:**
```bash
pkill -f golang-nexus-server
```

**Verify it's stopped:**
```bash
ps aux | grep golang-nexus-server
# Should show only the grep command itself
```

**Start new service:**
```bash
cd ~/golang-nexus-build
chmod +x golang-nexus-server
nohup ./golang-nexus-server \
  -port 9090 \
  -idf-host 127.0.0.1 \
  -idf-port 2027 \
  -log-level debug \
  > golang-nexus-server.log 2>&1 &
```

**Verify it started:**
```bash
tail -20 golang-nexus-server.log
```

**Expected output:**
```
Starting Golang Nexus Service...
Initializing IDF client: 127.0.0.1:2027
✅ IDF client initialized via singleton
✅ IDF repository initialized
Starting GRPC Server on port 9090
Registered ItemService with the gRPC server
```

---

## Step 4: Test OData Queries

### Test 1: Basic List (No Filter)
```bash
# Get auth token
TOKEN=$(curl -k -u admin:Nutanix.123 -X POST \
  https://10.124.86.254:9440/api/nutanix/v3/users/me/list 2>/dev/null | \
  jq -r '.entities[0].metadata.uuid' || echo "")

# Test basic list
curl -k -H "Authorization: Bearer $TOKEN" \
  "https://10.124.86.254:9440/api/nexus/v4.1/config/items?$page=0&$limit=10" | jq
```

**Expected:** Returns 10 items with pagination metadata

---

### Test 2: Filter by itemName
```bash
curl -k -H "Authorization: Bearer $TOKEN" \
  "https://10.124.86.254:9440/api/nexus/v4.1/config/items?$filter=itemName eq 'test item 0'" | jq
```

**Expected:** Returns only items where itemName equals 'test item 0'

---

### Test 3: Sort by itemId
```bash
curl -k -H "Authorization: Bearer $TOKEN" \
  "https://10.124.86.254:9440/api/nexus/v4.1/config/items?$orderby=itemId asc&$limit=5" | jq
```

**Expected:** Returns items sorted by itemId in ascending order

---

### Test 4: Combined Filter and Sort
```bash
curl -k -H "Authorization: Bearer $TOKEN" \
  "https://10.124.86.254:9440/api/nexus/v4.1/config/items?$filter=itemType eq 'TYPE1'&$orderby=itemName asc&$page=0&$limit=10" | jq
```

**Expected:** Returns TYPE1 items, sorted by itemName, paginated

---

### Test 5: Error Handling (Invalid Query)
```bash
curl -k -H "Authorization: Bearer $TOKEN" \
  "https://10.124.86.254:9440/api/nexus/v4.1/config/items?$filter=invalidField eq 'test'" | jq
```

**Expected:** Returns 400 Bad Request with error message:
```json
{
  "error": {
    "message": "Unknown property in OData query: property 'invalidField' not found"
  }
}
```

---

## Step 5: Verify Logs

**Check Go server logs:**
```bash
tail -50 ~/golang-nexus-build/golang-nexus-server.log | grep -i "odata\|filter\|orderby"
```

**Expected logs:**
```
Using regular IDF query evaluator
Using OData $filter: ...
Using OData $orderby: ...
Final IDF Query: ...
```

**Check Adonis logs:**
```bash
tail -50 ~/adonis/logs/prism-service.log | grep -i "nexus\|item"
```

**Expected:** Successful gRPC calls

---

## Troubleshooting

### Issue: OData filter not working

**Check:**
1. Are logs showing OData parsing? Look for "Using OData $filter" in logs
2. Is EDM binding correct? Check logs for "EDM bindings count: 1"
3. Is IDF query correct? Check logs for "Final IDF Query"

**Debug:**
```bash
# On PC, check logs
tail -100 ~/golang-nexus-build/golang-nexus-server.log | grep -A 5 -B 5 "filter\|orderby"
```

### Issue: Error "Unknown property"

**Cause:** Field name in OData query doesn't match EDM binding

**Fix:** Check EDM bindings in `odata_parser.go` - ensure property names match

### Issue: Error "Invalid OData query syntax"

**Cause:** Malformed OData expression

**Fix:** Check OData syntax (use `eq`, `ne`, `gt`, `lt`, etc.)

---

## Quick Command Reference

```bash
# Local: Build
cd ~/ntnx-api-golang-nexus && make build

# Local: Copy to PC
scp ~/ntnx-api-golang-nexus/golang-nexus-server nutanix@10.124.86.254:~/golang-nexus-build/

# PC: Stop and start
pkill -f golang-nexus-server
cd ~/golang-nexus-build
chmod +x golang-nexus-server
nohup ./golang-nexus-server -port 9090 -idf-host 127.0.0.1 -idf-port 2027 -log-level debug > golang-nexus-server.log 2>&1 &

# PC: Test
curl -k -H "Authorization: Bearer $TOKEN" \
  "https://10.124.86.254:9440/api/nexus/v4.1/config/items?$filter=itemName eq 'test'" | jq
```

---

## What Changed in This Deployment

**New Features:**
- ✅ OData query parsing (`$filter`, `$orderby`, `$select`)
- ✅ Field name mapping (OData → IDF)
- ✅ Error handling for invalid queries
- ✅ Enhanced IDF query building

**Files Changed:**
- `golang-nexus-server` binary (includes OData parser)
- No changes to JAR needed (OData parsing is in Go code only)

---

## Next Steps After Deployment

1. **Test all OData queries** (filter, sort, select)
2. **Verify error handling** (invalid queries return proper errors)
3. **Check performance** (OData parsing adds minimal overhead)
4. **Monitor logs** for any issues

---

**Ready to deploy!** Follow the steps above to push the OData-enabled binary to PC.


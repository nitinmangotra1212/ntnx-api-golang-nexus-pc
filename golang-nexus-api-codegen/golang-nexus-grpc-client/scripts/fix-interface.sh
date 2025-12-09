#!/bin/bash
# Fix interface file to remove CompletableFuture wrapper

INTERFACE_FILE="$1"
if [ -z "$INTERFACE_FILE" ]; then
    INTERFACE_FILE="target/generated-sources/swagger/src/nexus/v4/config/ItemApiControllerInterface.java"
fi

if [ ! -f "$INTERFACE_FILE" ]; then
    echo "Interface file not found: $INTERFACE_FILE"
    exit 0
fi

# Fix CompletableFuture to ResponseEntity
sed -i.bak 's/CompletableFuture<ResponseEntity<MappingJacksonValue>> listItems/ResponseEntity<MappingJacksonValue> listItems/g' "$INTERFACE_FILE"
sed -i.bak 's/CompletableFuture<ResponseEntity<MappingJacksonValue>> listCats/ResponseEntity<MappingJacksonValue> listCats/g' "$INTERFACE_FILE"

echo "✅ Fixed interface file: $INTERFACE_FILE"


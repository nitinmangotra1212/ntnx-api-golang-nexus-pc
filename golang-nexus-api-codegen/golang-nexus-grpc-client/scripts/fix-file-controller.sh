#!/bin/bash
#
# Fix for NexusConfigFileController.java
# Removes error handling code that doesn't work with simple File entity model
#

CONTROLLER_FILE="$1"
if [ -z "$CONTROLLER_FILE" ]; then
    CONTROLLER_FILE="target/generated-sources/swagger/src/nexus/v4/config/NexusConfigFileController.java"
fi

if [ ! -f "$CONTROLLER_FILE" ]; then
    echo "File not found: $CONTROLLER_FILE"
    exit 0
fi

# Use Python for reliable text processing
export CONTROLLER_FILE_PATH="$CONTROLLER_FILE"
python3 << 'PYEOF'
import re
import os

file_path = os.environ.get('CONTROLLER_FILE_PATH')
if not file_path:
    print("Error: CONTROLLER_FILE_PATH not set")
    exit(1)

with open(file_path, 'r') as f:
    content = f.read()

# Remove error handling code for File model (File doesn't have error response methods)
# Pattern: from "if(response.getContent().hasErrorResponseData())" to the closing brace before "// map proto-compiled-java"
pattern = r'(\s+if \(response\.getContent\(\) == null\) \{\s+return ResponseEntity\.noContent\(\)\.build\(\);\s+\}\s+)if\(response\.getContent\(\)\.hasErrorResponseData\(\)\) \{.*?\}\s+(\s+// map proto-compiled-java response content to java DTO)'

replacement = r'\1\2'

content = re.sub(pattern, replacement, content, flags=re.DOTALL)

# Write back
with open(file_path, 'w') as f:
    f.write(content)

print("✅ Fixed NexusConfigFileController.java - removed error handling for File model")
PYEOF


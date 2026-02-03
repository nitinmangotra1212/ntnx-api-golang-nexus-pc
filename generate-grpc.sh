#!/bin/bash
#
# Generate gRPC protobuf code (.pb.go files)
# This creates the same .pb.go files that guru service has
#

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROTO_DIR="${SCRIPT_DIR}/generated-code/protobuf/swagger/nexus/v4/config"
OUT_DIR="${SCRIPT_DIR}/generated-code/protobuf/nexus/v4/config"

echo "🔧 Generating gRPC protobuf code..."
echo "================================================"

# Create output directory
mkdir -p "${OUT_DIR}"

# Check if protoc is installed
if ! command -v protoc &> /dev/null; then
    echo "❌ protoc not found. Installing..."
    echo ""
    echo "Please install protoc:"
    echo "  macOS: brew install protobuf"
    echo "  Linux: apt-get install protobuf-compiler"
    exit 1
fi

# Check if protoc-gen-go is installed
if ! command -v protoc-gen-go &> /dev/null; then
    echo "📦 Installing protoc-gen-go..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

# Check if protoc-gen-go-grpc is installed
if ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo "📦 Installing protoc-gen-go-grpc..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

# Add Go bin to PATH if not already there
export PATH="${PATH}:$(go env GOPATH)/bin"

echo ""
echo "✅ Prerequisites installed"
echo ""

# Generate protobuf Go code
echo "🔄 Generating .pb.go files from protobuf definitions..."
echo ""

# Set proto paths - need to include swagger root for imports like nexus/v4/error/error.proto
SWAGGER_PROTO_ROOT="${SCRIPT_DIR}/generated-code/protobuf/swagger"
PROTO_OUT_ROOT="${SCRIPT_DIR}/generated-code/protobuf"

# Change to swagger root so we can reference proto files with their full paths
cd "${SWAGGER_PROTO_ROOT}"

# First, generate error.proto (needed by config.proto)
echo "  → error.proto"
protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
    --proto_path="$(go env GOROOT)/src" \
    --go_out="${PROTO_OUT_ROOT}" \
    --go_opt=paths=source_relative \
    nexus/v4/error/error.proto

# Generate stats.proto (needed by config.proto for itemStats navigation property)
echo "  → stats.proto"
protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
    --proto_path="$(go env GOROOT)/src" \
    --go_out="${PROTO_OUT_ROOT}" \
    --go_opt=paths=source_relative \
    nexus/v4/stats/stats.proto 2>/dev/null || echo "  ⚠️  stats.proto not found (will be generated after Maven build)"

# Generate for config.proto
echo "  → config.proto"
protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
    --proto_path="$(go env GOROOT)/src" \
    --go_out="${PROTO_OUT_ROOT}" \
    --go_opt=paths=source_relative \
    nexus/v4/config/config.proto

# Generate for item_service.proto (includes gRPC service definitions)
echo "  → item_service.proto"
protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
    --proto_path="$(go env GOROOT)/src" \
    --go_out="${PROTO_OUT_ROOT}" \
    --go_opt=paths=source_relative \
    --go-grpc_out="${PROTO_OUT_ROOT}" \
    --go-grpc_opt=paths=source_relative \
    nexus/v4/config/item_service.proto

# Generate for file_service.proto (includes gRPC service definitions for file transfer)
echo "  → file_service.proto"
protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
    --proto_path="$(go env GOROOT)/src" \
    --go_out="${PROTO_OUT_ROOT}" \
    --go_opt=paths=source_relative \
    --go-grpc_out="${PROTO_OUT_ROOT}" \
    --go-grpc_opt=paths=source_relative \
    nexus/v4/config/file_service.proto 2>/dev/null || echo "  ⚠️  file_service.proto not found (will be generated after Maven build)"

# Generate for stats module protobufs
STATS_PROTO_DIR="${SCRIPT_DIR}/generated-code/protobuf/swagger/nexus/v4/stats"
STATS_OUT_DIR="${SCRIPT_DIR}/generated-code/protobuf/nexus/v4/stats"
if [ -d "${STATS_PROTO_DIR}" ]; then
    mkdir -p "${STATS_OUT_DIR}"
    echo "  → ItemAssociationStats_service.proto"
    protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
        --proto_path="$(go env GOROOT)/src" \
        --go_out="${PROTO_OUT_ROOT}" \
        --go_opt=paths=source_relative \
        --go-grpc_out="${PROTO_OUT_ROOT}" \
        --go-grpc_opt=paths=source_relative \
        nexus/v4/stats/ItemAssociationStats_service.proto 2>/dev/null || echo "  ⚠️  ItemAssociationStats_service.proto not found"
    
    echo "  → ItemStats_service.proto"
    protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
        --proto_path="$(go env GOROOT)/src" \
        --go_out="${PROTO_OUT_ROOT}" \
        --go_opt=paths=source_relative \
        --go-grpc_out="${PROTO_OUT_ROOT}" \
        --go-grpc_opt=paths=source_relative \
        nexus/v4/stats/ItemStats_service.proto 2>/dev/null || echo "  ⚠️  ItemStats_service.proto not found"
fi

# Post-process: Fix import paths in generated .pb.go files
# IMPORTANT: Only fix Go import statements, NOT the raw descriptor (binary data)
echo ""
echo "🔧 Fixing import paths in generated .pb.go files..."
PROTO_OUT_DIR="${SCRIPT_DIR}/generated-code/protobuf"
for file in $(find "${PROTO_OUT_DIR}" -name "*.pb.go" -type f); do
    if [[ -f "$file" ]]; then
        # Fix import statements (lines starting with "import" or in import block)
        # But NOT the raw descriptor constant (which contains binary data)
        # Use a more targeted approach: only fix actual import statements
        python3 << PYEOF
import re
with open("$file", "r") as f:
    content = f.read()

# Only replace in import statements, not in raw descriptor constants
# Match import statements (single line or block)
lines = content.split('\n')
in_import_block = False
result_lines = []
skip_next_rawdesc = False

for i, line in enumerate(lines):
    # Skip raw descriptor constants - they contain binary data and should not be modified
    if 'const file_' in line and '_rawDesc =' in line:
        skip_next_rawdesc = True
        result_lines.append(line)
        continue
    
    if skip_next_rawdesc:
        # Skip until we find the closing of the rawDesc constant
        result_lines.append(line)
        if line.strip().endswith('"') and not line.strip().endswith('\\"'):
            # Check if this looks like end of rawDesc (next line is usually "var" or "func")
            if i + 1 < len(lines) and ('var ' in lines[i+1] or 'func ' in lines[i+1]):
                skip_next_rawdesc = False
        continue
    
    # Fix import statements only
    if line.strip().startswith('import ') or (in_import_block and (line.strip().startswith('"') or line.strip() == ')')):
        # This is an import statement
        line = re.sub(r'"nexus/v4/', r'"github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/protobuf/nexus/v4/', line)
        if line.strip() == 'import (':
            in_import_block = True
        elif line.strip() == ')':
            in_import_block = False
    elif 'import (' in line:
        in_import_block = True
    
    result_lines.append(line)

content = '\n'.join(result_lines)

# Remove blank imports
content = re.sub(r'^[ \t]*_[ \t]*"github\.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/protobuf/nexus/v4"$', '', content, flags=re.MULTILINE)
content = re.sub(r'^[ \t]*_[ \t]*"nexus/v4"$', '', content, flags=re.MULTILINE)

with open("$file", "w") as f:
    f.write(content)
PYEOF
    fi
done

# Fix common/v1 and nexus/v4/error and nexus/v4/stats imports (these are in import statements, safe to fix)
echo "  🔧 Fixing common/v1, nexus/v4/error, and nexus/v4/stats import paths..."
for file in $(find "${PROTO_OUT_DIR}" -name "*.pb.go" -type f); do
    if [[ -f "$file" ]]; then
        sed -i '' 's|response "common/v1/response"|response "github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/protobuf/common/v1/response"|g' "$file"
        sed -i '' 's|config "common/v1/config"|config "github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/protobuf/common/v1/config"|g' "$file"
        sed -i '' 's|"nexus/v4/error"|"github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/protobuf/nexus/v4/error"|g' "$file"
        sed -i '' 's|stats "nexus/v4/stats"|stats "github.com/nutanix/ntnx-api-golang-nexus-pc/generated-code/protobuf/nexus/v4/stats"|g' "$file"
    fi
done

echo "  ✅ Fixed import paths in all .pb.go files (preserved raw descriptors)"

# Post-process: Fix FullMethodName to use lowercase method names from proto
# protoc-gen-go-grpc converts method names to PascalCase, but FullMethodName should preserve proto case
echo ""
echo "🔧 Fixing FullMethodName constants to match proto method names (lowercase)..."
GRPC_FILE="${OUT_DIR}/item_service_grpc.pb.go"
if [ -f "${GRPC_FILE}" ]; then
    # Note: Proto file already has lowercase method names (listItems, getItem, etc.)
    # protoc-gen-go-grpc preserves the proto case, so no changes needed
    # Just verify method names are lowercase
    if grep -q 'MethodName: "ListItems"' "${GRPC_FILE}"; then
        # If protoc generated PascalCase, convert to lowercase
        sed -i '' 's|MethodName: "ListItems"|MethodName: "listItems"|g' "${GRPC_FILE}"
        sed -i '' 's|MethodName: "GetItem"|MethodName: "getItem"|g' "${GRPC_FILE}"
        sed -i '' 's|MethodName: "CreateItem"|MethodName: "createItem"|g' "${GRPC_FILE}"
        sed -i '' 's|MethodName: "UpdateItem"|MethodName: "updateItem"|g' "${GRPC_FILE}"
        sed -i '' 's|MethodName: "DeleteItem"|MethodName: "deleteItem"|g' "${GRPC_FILE}"
        sed -i '' 's|MethodName: "GetItemAsync"|MethodName: "getItemAsync"|g' "${GRPC_FILE}"
        # Fix FullMethodName constants if needed
        sed -i '' 's|/nexus.v4.config.ItemService/ListItems|/nexus.v4.config.ItemService/listItems|g' "${GRPC_FILE}"
        sed -i '' 's|/nexus.v4.config.ItemService/GetItem|/nexus.v4.config.ItemService/getItem|g' "${GRPC_FILE}"
        sed -i '' 's|/nexus.v4.config.ItemService/CreateItem|/nexus.v4.config.ItemService/createItem|g' "${GRPC_FILE}"
        sed -i '' 's|/nexus.v4.config.ItemService/UpdateItem|/nexus.v4.config.ItemService/updateItem|g' "${GRPC_FILE}"
        sed -i '' 's|/nexus.v4.config.ItemService/DeleteItem|/nexus.v4.config.ItemService/deleteItem|g' "${GRPC_FILE}"
        sed -i '' 's|/nexus.v4.config.ItemService/GetItemAsync|/nexus.v4.config.ItemService/getItemAsync|g' "${GRPC_FILE}"
        echo "  ✅ Fixed ItemService method names to lowercase (listItems, getItem, createItem, updateItem, deleteItem, getItemAsync)"
    else
        echo "  ✅ Method names already lowercase (no changes needed)"
    fi
fi

# Fix FileService method names if file_service_grpc.pb.go exists
FILE_GRPC_FILE="${OUT_DIR}/file_service_grpc.pb.go"
if [ -f "${FILE_GRPC_FILE}" ]; then
    if grep -q 'MethodName: "UploadFile"' "${FILE_GRPC_FILE}"; then
        sed -i '' 's|MethodName: "UploadFile"|MethodName: "uploadFile"|g' "${FILE_GRPC_FILE}"
        sed -i '' 's|MethodName: "DownloadFile"|MethodName: "downloadFile"|g' "${FILE_GRPC_FILE}"
        sed -i '' 's|/nexus.v4.config.FileService/UploadFile|/nexus.v4.config.FileService/uploadFile|g' "${FILE_GRPC_FILE}"
        sed -i '' 's|/nexus.v4.config.FileService/DownloadFile|/nexus.v4.config.FileService/downloadFile|g' "${FILE_GRPC_FILE}"
        echo "  ✅ Fixed FileService method names to lowercase (uploadFile, downloadFile)"
    else
        echo "  ✅ FileService method names already lowercase (no changes needed)"
    fi
fi

echo ""
echo "================================================"
echo "✅ gRPC code generation complete!"
echo ""
echo "📁 Generated files:"
ls -lh "${OUT_DIR}"/*.pb.go 2>/dev/null || echo "  (checking...)"
echo ""
echo "📦 Generated files:"
echo "  - config.pb.go          (protobuf messages)"
echo "  - item_service.pb.go     (service messages)"
echo "  - item_service_grpc.pb.go (gRPC service stubs)"
if [ -f "${OUT_DIR}/file_service_grpc.pb.go" ]; then
    echo "  - file_service.pb.go     (file service messages)"
    echo "  - file_service_grpc.pb.go (file service gRPC stubs)"
fi
echo ""
echo "🎉 Ready to implement gRPC servers!"


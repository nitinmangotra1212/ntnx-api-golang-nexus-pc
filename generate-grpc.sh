#!/bin/bash
#
# Generate gRPC protobuf code (.pb.go files)
# This creates the same .pb.go files that guru service has
#

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROTO_DIR="${SCRIPT_DIR}/generated-code/protobuf/swagger/mock/v4/config"
OUT_DIR="${SCRIPT_DIR}/generated-code/protobuf/mock/v4/config"

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

# Set proto paths - need to include swagger root for imports like mock/v4/error/error.proto
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
    mock/v4/error/error.proto

# Generate for config.proto
echo "  → config.proto"
protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
    --proto_path="$(go env GOROOT)/src" \
    --go_out="${PROTO_OUT_ROOT}" \
    --go_opt=paths=source_relative \
    mock/v4/config/config.proto

# Generate for cat_service.proto (includes gRPC service definitions)
echo "  → cat_service.proto"
protoc --proto_path="${SWAGGER_PROTO_ROOT}" \
    --proto_path="$(go env GOROOT)/src" \
    --go_out="${PROTO_OUT_ROOT}" \
    --go_opt=paths=source_relative \
    --go-grpc_out="${PROTO_OUT_ROOT}" \
    --go-grpc_opt=paths=source_relative \
    mock/v4/config/cat_service.proto

# Post-process: Fix import paths in generated .pb.go files
echo ""
echo "🔧 Fixing import paths in generated .pb.go files..."
PROTO_OUT_DIR="${SCRIPT_DIR}/generated-code/protobuf"
for file in $(find "${PROTO_OUT_DIR}" -name "*.pb.go" -type f); do
    if [[ -f "$file" ]]; then
        # Fix import paths: mock/v4/... -> github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/...
        sed -i '' 's|"mock/v4/|"github.com/nutanix/ntnx-api-golang-mock-pc/generated-code/protobuf/mock/v4/|g' "$file"
        # Remove blank imports to non-existent packages (like mock/v4 for api_version.proto which doesn't generate Go code)
        # Match both the full path and the relative path versions
        sed -i '' '/^[[:space:]]*_[[:space:]]*"github\.com\/nutanix\/ntnx-api-golang-mock-pc\/generated-code\/protobuf\/mock\/v4"$/d' "$file"
        sed -i '' '/^[[:space:]]*_[[:space:]]*"mock\/v4"$/d' "$file"
    fi
done
echo "  ✅ Fixed import paths in all .pb.go files"

# Post-process: Fix FullMethodName to use lowercase method names from proto
# protoc-gen-go-grpc converts method names to PascalCase, but FullMethodName should preserve proto case
echo ""
echo "🔧 Fixing FullMethodName constants to match proto method names (lowercase)..."
GRPC_FILE="${OUT_DIR}/cat_service_grpc.pb.go"
if [ -f "${GRPC_FILE}" ]; then
    # Fix all method names to use lowercase (camelCase) to match proto definitions
    # Fix FullMethodName constants
    sed -i '' 's|/mock.v4.config.CatService/ListCats|/mock.v4.config.CatService/listCats|g' "${GRPC_FILE}"
    sed -i '' 's|/mock.v4.config.CatService/GetCat|/mock.v4.config.CatService/getCat|g' "${GRPC_FILE}"
    sed -i '' 's|/mock.v4.config.CatService/CreateCat|/mock.v4.config.CatService/createCat|g' "${GRPC_FILE}"
    sed -i '' 's|/mock.v4.config.CatService/UpdateCat|/mock.v4.config.CatService/updateCat|g' "${GRPC_FILE}"
    sed -i '' 's|/mock.v4.config.CatService/DeleteCat|/mock.v4.config.CatService/deleteCat|g' "${GRPC_FILE}"
    sed -i '' 's|/mock.v4.config.CatService/GetCatAsync|/mock.v4.config.CatService/getCatAsync|g' "${GRPC_FILE}"
    # Fix MethodName in ServiceDesc (gRPC uses this for method matching)
    sed -i '' 's|MethodName: "ListCats"|MethodName: "listCats"|g' "${GRPC_FILE}"
    sed -i '' 's|MethodName: "GetCat"|MethodName: "getCat"|g' "${GRPC_FILE}"
    sed -i '' 's|MethodName: "CreateCat"|MethodName: "createCat"|g' "${GRPC_FILE}"
    sed -i '' 's|MethodName: "UpdateCat"|MethodName: "updateCat"|g' "${GRPC_FILE}"
    sed -i '' 's|MethodName: "DeleteCat"|MethodName: "deleteCat"|g' "${GRPC_FILE}"
    sed -i '' 's|MethodName: "GetCatAsync"|MethodName: "getCatAsync"|g' "${GRPC_FILE}"
    echo "  ✅ Fixed all CatService FullMethodName constants and MethodName fields (listCats, getCat, createCat, updateCat, deleteCat, getCatAsync)"
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
echo "  - cat_service.pb.go     (service messages)"
echo "  - cat_service_grpc.pb.go (gRPC service stubs)"
echo ""
echo "🎉 Ready to implement gRPC servers!"


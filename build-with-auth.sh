#!/bin/bash

# Build script with GitHub authentication for Maven
# This script pre-clones the repository to cache, avoiding JGit authentication issues
# 
# Usage:
#   ./build-with-auth.sh
# 
# Or set environment variables:
#   export GIT_USERNAME=your-username
#   export GIT_PASSWORD=your-token
#   ./build-with-auth.sh

set -e

echo "🔧 Setting up GitHub authentication for Maven build..."

# Set Git credentials (from environment or use defaults)
# IMPORTANT: Do not commit credentials! Use environment variables instead.
GIT_USERNAME="${GIT_USERNAME:-your-username}"
GIT_PASSWORD="${GIT_PASSWORD:-your-token-or-password}"

echo "✅ Using Git credentials for user: $GIT_USERNAME"
echo ""

# Pre-clone repository to Maven cache location
# The plugin caches repos in ~/.m2/repository/.cache/dev-platform/
CACHE_BASE="$HOME/.m2/repository/.cache"
CACHE_DIR="$CACHE_BASE/dev-platform"
REPO_CACHE="$CACHE_DIR/ntnx-api-dev-platform"
REPO_URL="https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/nutanix-core/ntnx-api-dev-platform.git"
TAG="17.6.0.9581-RELEASE"

echo "📦 Pre-cloning repository to cache..."
mkdir -p "$CACHE_DIR"

PROJECT_DIR="$(cd "$(dirname "$0")" && pwd)"

if [ -d "$REPO_CACHE" ]; then
    echo "  Cache exists, updating..."
    cd "$REPO_CACHE"
    git fetch origin "refs/tags/${TAG}:refs/tags/${TAG}" --depth 1 2>/dev/null || \
    git fetch origin "${TAG}" --depth 1 2>/dev/null || true
    git checkout "${TAG}" 2>/dev/null || git checkout "tags/${TAG}" 2>/dev/null || true
    cd "$PROJECT_DIR"
else
    echo "  Cloning fresh copy..."
    cd "$CACHE_DIR"
    git clone --depth 1 --branch "${TAG}" "${REPO_URL}" ntnx-api-dev-platform 2>&1 | grep -v "^Cloning" || {
        # If branch clone fails, try tag checkout
        git clone --depth 1 "${REPO_URL}" ntnx-api-dev-platform
        cd ntnx-api-dev-platform
        git checkout "${TAG}" 2>/dev/null || git checkout "tags/${TAG}" 2>/dev/null || true
    }
    cd "$PROJECT_DIR"
fi

echo "✅ Repository cached at: $REPO_CACHE"
echo ""

# Run Maven build (plugin will use cached copy)
echo "🚀 Running Maven build..."

# Set MAVEN_OPTS to fix Java module access issues
export MAVEN_OPTS="--add-opens java.base/java.lang=ALL-UNNAMED --add-opens java.base/java.lang.reflect=ALL-UNNAMED --add-opens java.base/java.util=ALL-UNNAMED"

cd "$PROJECT_DIR"
mvn clean install -DskipTests -s settings.xml

echo ""
echo "✅ Build complete!"

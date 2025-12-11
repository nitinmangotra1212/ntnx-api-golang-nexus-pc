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

# Set Git credentials
# TODO: REMOVE CREDENTIALS BEFORE PUSHING TO GIT!
# This is temporary - credentials should be removed before committing

# Try to load from .env file if it exists (preferred)
if [ -f .env ]; then
    echo "📝 Loading credentials from .env file..."
    set -a
    source .env
    set +a
fi

# Default credentials (TEMPORARY - REMOVE BEFORE PUSHING!)
# These can be overridden by environment variables or .env file
GIT_USERNAME="${GIT_USERNAME:-USERNAME_HERE}"
GIT_PASSWORD="${GIT_PASSWORD:-PASSWORD_HERE}"

echo "✅ Using Git credentials for user: $GIT_USERNAME"
echo ""

# Export credentials as environment variables for Maven
export GIT_USERNAME
export GIT_PASSWORD

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

if [ -d "$REPO_CACHE" ] && [ -d "$REPO_CACHE/.git" ]; then
    echo "  Cache exists, updating..."
    cd "$REPO_CACHE"
    # Update remote URL with credentials
    git remote set-url origin "${REPO_URL}" 2>/dev/null || true
    git fetch origin "refs/tags/${TAG}:refs/tags/${TAG}" --depth 1 2>/dev/null || \
    git fetch origin "${TAG}" --depth 1 2>/dev/null || {
        echo "  ⚠️  Fetch failed, trying full clone..."
        cd "$PROJECT_DIR"
        rm -rf "$REPO_CACHE"
        cd "$CACHE_DIR"
        git clone --depth 1 --branch "${TAG}" "${REPO_URL}" ntnx-api-dev-platform || {
            echo "  ⚠️  Branch clone failed, trying tag checkout..."
            git clone --depth 1 "${REPO_URL}" ntnx-api-dev-platform
            cd ntnx-api-dev-platform
            git checkout "${TAG}" 2>/dev/null || git checkout "tags/${TAG}" 2>/dev/null || {
                echo "❌ ERROR: Failed to checkout tag ${TAG}"
                exit 1
            }
        }
    }
    git checkout "${TAG}" 2>/dev/null || git checkout "tags/${TAG}" 2>/dev/null || true
    cd "$PROJECT_DIR"
else
    echo "  Cloning fresh copy..."
    cd "$CACHE_DIR"
    # Remove any broken cache
    rm -rf ntnx-api-dev-platform
    git clone --depth 1 --branch "${TAG}" "${REPO_URL}" ntnx-api-dev-platform 2>&1 | grep -v "^Cloning" || {
        echo "  ⚠️  Branch clone failed, trying tag checkout..."
        # If branch clone fails, try tag checkout
        git clone --depth 1 "${REPO_URL}" ntnx-api-dev-platform || {
            echo "❌ ERROR: Failed to clone repository!"
            echo "   Check your GIT_USERNAME and GIT_PASSWORD in .env file"
            exit 1
        }
        cd ntnx-api-dev-platform
        git checkout "${TAG}" 2>/dev/null || git checkout "tags/${TAG}" 2>/dev/null || {
            echo "❌ ERROR: Failed to checkout tag ${TAG}"
            exit 1
        }
        cd ..
    }
    cd "$PROJECT_DIR"
fi

# Verify cache is valid
if [ ! -d "$REPO_CACHE/.git" ]; then
    echo "❌ ERROR: Cache directory is invalid!"
    echo "   Removing broken cache..."
    rm -rf "$REPO_CACHE"
    exit 1
fi

echo "✅ Repository cached at: $REPO_CACHE"
echo ""

# Run Maven build (plugin will use cached copy)
echo "🚀 Running Maven build..."

# Set MAVEN_OPTS to fix Java module access issues
export MAVEN_OPTS="--add-opens java.base/java.lang=ALL-UNNAMED --add-opens java.base/java.lang.reflect=ALL-UNNAMED --add-opens java.base/java.util=ALL-UNNAMED"

cd "$PROJECT_DIR"

# Ensure cache remote URL has credentials embedded
cd "$REPO_CACHE"
git remote set-url origin "https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/nutanix-core/ntnx-api-dev-platform.git" 2>/dev/null || true
git checkout "${TAG}" 2>/dev/null || true
cd "$PROJECT_DIR"

# Run Maven build with credentials
mvn clean install -DskipTests -s settings.xml \
  -Dgit.username="${GIT_USERNAME}" \
  -Dgit.password="${GIT_PASSWORD}" \
  -Dmaven.wagon.http.ssl.insecure=true

echo ""
echo "✅ Build complete!"

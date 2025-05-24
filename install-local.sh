#!/bin/bash

set -e

APP_NAME="zyracli"
INSTALL_PATH="/usr/local/bin"
PROJECT_DIR=$(pwd)
BUILD_DIR="$PROJECT_DIR/build"

echo "🔨 Building $APP_NAME from source..."

# Ensure Go is installed
if ! command -v go &> /dev/null; then
  echo "❌ Go is not installed. Please install Go first."
  exit 1
fi

# Clean old builds
go clean

# Create build directory
mkdir -p "$BUILD_DIR"

# Detect OS and architecture
GOOS=$(uname | tr '[:upper:]' '[:lower:]')
GOARCH=$(uname -m)

case "$GOARCH" in
  arm64|aarch64)   GOARCH="arm64" ;;
  x86_64|amd64)    GOARCH="amd64" ;;
  *) echo "Unsupported architecture: $GOARCH" && exit 1 ;;
esac

echo "📦 Targeting OS=$GOOS ARCH=$GOARCH"

# Build with flags
CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -o "$BUILD_DIR/$APP_NAME" .

echo "✅ Build successful."

echo "🚀 Installing to $INSTALL_PATH/$APP_NAME"
sudo mv "$BUILD_DIR/$APP_NAME" "$INSTALL_PATH/$APP_NAME"

if [ -f "$INSTALL_PATH/$APP_NAME" ]; then
  chmod +x "$INSTALL_PATH/$APP_NAME"
  echo "✅ Installed at $(which $APP_NAME)"
else
  echo "❌ Failed to install $APP_NAME"
  exit 1
fi

echo "🎉 Done! You can now run '$APP_NAME' globally."
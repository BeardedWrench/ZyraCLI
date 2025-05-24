#!/bin/bash

set -e

REPO_URL="https://github.com/beardedwrench/ZyraCLI"
INSTALL_DIR="/usr/local/bin"
TEMP_DIR=$(mktemp -d)
BINARY_NAME="zyracli"
ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

echo "Detected OS: $OS, Arch: $ARCH"

case "$OS" in
    linux)
        if [[ "$ARCH" == "x86_64" ]]; then
            BINARY_URL="$REPO_URL/releases/latest/download/zyracli-linux-amd64"
        elif [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then
            BINARY_URL="$REPO_URL/releases/latest/download/zyracli-linux-arm64"
        else
            echo "Unsupported architecture: $ARCH"
            exit 1
        fi
        ;;
    darwin)
        if [[ "$ARCH" == "x86_64" ]]; then
            BINARY_URL="$REPO_URL/releases/latest/download/zyracli-darwin-amd64"
        elif [[ "$ARCH" == "arm64" ]]; then
            BINARY_URL="$REPO_URL/releases/latest/download/zyracli-darwin-arm64"
        else
            echo "Unsupported architecture: $ARCH"
            exit 1
        fi
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

echo "Downloading ZyraCLI from $BINARY_URL"
curl -L "$BINARY_URL" -o "$TEMP_DIR/$BINARY_NAME"
chmod +x "$TEMP_DIR/$BINARY_NAME"

echo "Installing to $INSTALL_DIR..."
sudo mv "$TEMP_DIR/$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"

echo "Installed successfully! Run 'zyracli install' to set it up."
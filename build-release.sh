#!/bin/bash

set -e

APP_NAME="dashuni"
OUTPUT_DIR="release"

# Determine version from the latest Git tag
VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0-dev")

echo "Building $APP_NAME $VERSION..."

# Clean output directory
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

# Make a staging folder for packaging
STAGE_DIR="./.package_stage"

prepare_package() {
    PLATFORM_BINARY=$1
    FINAL_ARCHIVE=$2

    echo "➡️  Preparing package: $FINAL_ARCHIVE"

    # Clean stage
    rm -rf "$STAGE_DIR"
    mkdir -p "$STAGE_DIR"

    # Copy binary
    cp "${OUTPUT_DIR}/${PLATFORM_BINARY}" "${STAGE_DIR}/${APP_NAME}"

    # Copy mappings and README
    cp -r mappings "${STAGE_DIR}/"
    cp README.md "${STAGE_DIR}/"

    # Create archive
    (cd "$STAGE_DIR" && tar -czf "../${OUTPUT_DIR}/${FINAL_ARCHIVE}" *)
}

# macOS
echo "➡️  Building macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o "${OUTPUT_DIR}/${APP_NAME}-macos-amd64"
prepare_package "${APP_NAME}-macos-amd64" "${APP_NAME}-${VERSION}-macos-amd64.tar.gz"
rm "${OUTPUT_DIR}/${APP_NAME}-macos-amd64"

# Linux
echo "➡️  Building Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_DIR}/${APP_NAME}-linux-amd64"
prepare_package "${APP_NAME}-linux-amd64" "${APP_NAME}-${VERSION}-linux-amd64.tar.gz"
rm "${OUTPUT_DIR}/${APP_NAME}-linux-amd64"

# Windows
echo "➡️  Building Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o "${OUTPUT_DIR}/${APP_NAME}-windows-amd64.exe"

# Prepare Windows ZIP
echo "➡️  Preparing package: Windows zip"

rm -rf "$STAGE_DIR"
mkdir -p "$STAGE_DIR"

cp "${OUTPUT_DIR}/${APP_NAME}-windows-amd64.exe" "${STAGE_DIR}/${APP_NAME}.exe"
cp -r mappings "${STAGE_DIR}/"
cp README.md "${STAGE_DIR}/"

(cd "$STAGE_DIR" && zip -r "../${OUTPUT_DIR}/${APP_NAME}-${VERSION}-windows-amd64.zip" *)
rm "${OUTPUT_DIR}/${APP_NAME}-windows-amd64.exe"

# Clean up stage
rm -rf "$STAGE_DIR"

echo ""
echo "✅ All builds complete! Files in ${OUTPUT_DIR}:"
ls -lh "${OUTPUT_DIR}"

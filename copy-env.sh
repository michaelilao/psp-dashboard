#!/bin/bash

SOURCE_FILE="sample.env"
TARGET_FILES=(".env" "fe/.env" "be/.env")

if [ ! -f "$SOURCE_FILE" ]; then
  echo "Error: '$SOURCE_FILE' not found!"
  exit 1
fi

for TARGET in "${TARGET_FILES[@]}"; do
  cp "$SOURCE_FILE" "$TARGET"
  echo "Copied to $TARGET"
done

echo "All environment files have been updated."

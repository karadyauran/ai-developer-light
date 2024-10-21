#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIRECTORY="$(dirname "$(dirname "$SCRIPT_DIR")")"

GENERATED_PROJECTS_DIR="$ROOT_DIRECTORY/generated_projects"
PROJECT_DIRECTORY="$GENERATED_PROJECTS_DIR"

PROJECT_NAME="$1"
T_FILE_PATH="$2"

FULL_PATH="$PROJECT_DIRECTORY/$PROJECT_NAME/$T_FILE_PATH"

DIR_PATH="$(dirname "$FULL_PATH")"
if [ ! -d "$DIR_PATH" ]; then
    mkdir -p "$DIR_PATH"
fi

if [ ! -f "$FULL_PATH" ]; then
    touch "$FULL_PATH"
else
    echo "File $FULL_PATH already exists, skipping creation."
fi
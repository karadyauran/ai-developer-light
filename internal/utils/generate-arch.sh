#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIRECTORY="$(dirname "$(dirname "$SCRIPT_DIR")")"

PROJECT_NAME="$1"

if [ -z "$PROJECT_NAME" ]; then
    echo "No project name provided."
    exit 1
fi

GENERATED_PROJECTS_DIR="$ROOT_DIRECTORY/generated_projects"
PROJECT_DIRECTORY="$GENERATED_PROJECTS_DIR/$PROJECT_NAME"

DIRS=(
    ".github/workflows"
    "cmd/main"
    "internal/config"
    "internal/controller"
    "internal/model"
    "internal/router"
    "internal/service"
    "docs"
)

DIRS_WITH_FILES=(
    "cmd/main:main.go"
    "internal/config:config.go"
    "internal/controller:controller.go"
    "internal/router:router.go"
    "internal/service:service.go"
)

TOP_LEVEL_FILES=("docker-compose.yaml" "package.json" "README.md" ".gitignore")

create_directories() {
    for dir in "${DIRS[@]}"; do
        mkdir -p "$PROJECT_DIRECTORY/$dir"
    done
}

create_files_in_directories() {
    for entry in "${DIRS_WITH_FILES[@]}"; do
        IFS=":" read -r dir file <<< "$entry"
        touch "$PROJECT_DIRECTORY/$dir/$file"
    done
}

create_top_level_files() {
    for file in "${TOP_LEVEL_FILES[@]}"; do
        touch "$PROJECT_DIRECTORY/$file"
    done
}

if [ ! -d "$GENERATED_PROJECTS_DIR" ]; then
    mkdir -p "$GENERATED_PROJECTS_DIR"
fi

if [ ! -d "$PROJECT_DIRECTORY" ]; then
    mkdir -p "$PROJECT_DIRECTORY"
fi

create_directories
create_files_in_directories
create_top_level_files

"$SCRIPT_DIR/init-code.sh" "$PROJECT_NAME"

echo "Project $PROJECT_NAME successfully created in $PROJECT_DIRECTORY."
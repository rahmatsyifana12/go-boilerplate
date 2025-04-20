#!/bin/bash

echo "Generating mocks..."

# Define an array of commands
declare -a MOCK_COMMANDS=(
    "rm -rf ./src/mocks"
    "mockgen -source=./src/repositories/user_repository.go -destination=./src/mocks/mock_user_repository.go -package=mocks"
    "mockgen -source=./src/repositories/todo_repository.go -destination=./src/mocks/mock_todo_repository.go -package=mocks"
)

# Iterate over the commands and execute them
for cmd in "${MOCK_COMMANDS[@]}"; do
    echo "Executing: $cmd"
    eval "$cmd"

    if [[ $? -ne 0 ]]; then
        echo "Failed to execute: $cmd"
        exit 1
    fi

    echo "Success: $cmd"
done

echo "All mocks generated successfully!"

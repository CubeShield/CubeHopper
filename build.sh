#!/bin/bash

set -e

APP_NAME="CubeHopper"

mkdir -p builds

echo "Сборка для Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -o ./builds/${APP_NAME}.exe .

echo "Сборка для Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o ./builds/${APP_NAME}-linux .

echo "Сборка для macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o ./builds/${APP_NAME}-intel .

echo "Сборка для macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build -o ./builds/${APP_NAME}-msilicon .

echo "Сборка завершена. Файлы в папке builds/."
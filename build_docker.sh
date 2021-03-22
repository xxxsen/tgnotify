#!/bin/bash

export DOCKER_CLI_EXPERIMENTAL=enabled

docker buildx create --use --name mybuild

docker buildx build -t xxxsen/tgnotify:0.0.1 \
  --platform=linux/amd64,linux/arm64 . --push

docker buildx build -t xxxsen/tgnotify:latest \
  --platform=linux/amd64,linux/arm64 . --push
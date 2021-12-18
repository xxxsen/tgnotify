#!/bin/bash

export DOCKER_CLI_EXPERIMENTAL=enabled

docker buildx create --use --name mybuild

docker buildx build -t xxxsen/tgnotify:v0.0.2 -t xxxsen/tgnotify:latest  \
  --platform=linux/amd64,linux/arm64 . --push
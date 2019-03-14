#!/bin/sh
set -ex

build() {
    local dockerfile="docker/Dockerfile.$1"
    docker build \
        -f "$dockerfile" \
        -t "pets-api:latest" \
        .
}

compose() {
    COMPOSE_PROJECT_NAME='pets' \
    COMPOSE_FILE='docker/docker-compose.yml' \
        docker-compose "$@"
}

build api
compose up -d --remove-orphans --force-recreate --no-build api

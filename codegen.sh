#!/bin/bash

DIR=$(cd $(dirname $0); pwd)
DOCKER_DIR=$(cd $(dirname $(dirname $0)); pwd)/pkg/docker/options

# clean generated files
rm ${DOCKER_DIR}/*.gen.go
mkdir -p bin

set -e

go build -o ${DIR}/pkg/gen/bin/option-gen ${DIR}/pkg/gen/main.go

subcmds=(
    "run"
    "attatch"
    "build"
    "builder"
    #"checkpoint"
    "commit"
    "config"
    "container"
    "context"
    "cp"
    "create"
    #"deploy"
    #"diff"
    "engine"
    "events"
    "exec"
    "export"
    "history"
    "image"
    "images"
    "import"
    "info"
    "inspect"
    "kill"
    "load"
    "login"
    #"logout"
    "logs"
    #"manifest"
    "network"
    "node"
    #"pause"
    "plugin"
    #"port"
    "ps"
    "pull"
    "push"
    "restart"
    "rm"
    "rmi"
    "run"
    "save"
    "search"
    "secret"
    "service"
    "stack"
    "start"
    "stats"
    "stop"
    "swarm"
    "system"
    #"tag"
    #"top"
    "trust"
    #"unpause"
    "update"
    "version"
    "volume"
    #"wait"
)

for cmd in "${subcmds[@]}"; do
  camelized=`echo ${cmd} | gsed -r 's/[- ](.)/\U\1\E/g'`
  snaked=`echo ${cmd} | gsed -r 's/[- ]/_/g'`
  docker ${cmd} --help | ${DIR}/pkg/gen/bin/option-gen -o ${DOCKER_DIR}/option_${snaked}.gen.go -var ${camelized}
  goimports -w ${DOCKER_DIR}/option_${snaked}.gen.go
done
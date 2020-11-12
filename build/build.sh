#!/bin/bash

set -eo pipefail

echo "> building binary on the host..."

# set binary target directory
mkdir -p ${ROOT}/build/bin/${ARCH}/

# shellcheck disable=SC2089
LDFLAGS="-s -w -X '${MODULE}/version.REVISION=${REVISION}' -X '${MODULE}/version.BRANCH=${BRANCH}' -X '${MODULE}/version.TAG=${TAG}' -X '${MODULE}/version.GOVERSION=${GOVERSION}'  -X '${MODULE}/version.BUILDTIME=${BUILDTIME}'"

# go build
CGO_ENABLED=0 go build -mod vendor -o ${ROOT}/build/bin/${ARCH}/${APPNAME} -v -x -ldflags "${LDFLAGS}" ${ROOT}/cmd/nfs/nfs.go

echo "CGO_ENABLED=0 go build -mod vendor -o ${ROOT}/build/bin/${ARCH}/${APPNAME} -v -x -ldflags "${LDFLAGS}" ${ROOT}/cmd/nfs/nfs.go"

if [[ $? -eq 0 ]]; then
    echo "> build successfully."
fi

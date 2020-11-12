.DEFAULT_GOAL=build

# set default shell
SHELL=/bin/bash -o pipefail -o errexit

HOST_ARCH = $(shell which go >/dev/null 2>&1 && go env GOARCH)
ARCH ?= $(HOST_ARCH)

ROOT=$(shell pwd)
APPNAME=$(shell basename `pwd`)
MODULE=github.com/penglongli/$(APPNAME)

# build version variables
REVISION := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD | grep -v HEAD)
TAG := $(shell git branch | head -n 1 | awk '{print $$4}' | head -c-2)
GOVERSION := $(shell go version)
BUILDTIME := $(shell date "+%Y-%m-%d %H:%M:%S")

# build
.PHONY: build
build:
	ROOT=$(ROOT) \
	APPNAME=$(APPNAME) \
	MODULE=$(MODULE) \
	ARCH=$(ARCH) \
	REVISION=$(REVISION) \
	BRANCH=$(BRANCH) \
	TAG=$(TAG) \
	GOVERSION="$(GOVERSION)" \
	BUILDTIME="$(BUILDTIME)" \
	build/build.sh


# 当前目录
CUR_DIR=$(shell pwd)
OUT_DIR=$(CUR_DIR)/bin

# 命令
GO_BUILD = CGO_ENABLED=0 go build -trimpath

SERVER_VERSION	?= $(shell git describe --long --tags --dirty --always)
SERVER_VERSION	?= unkonwn
BUILD_TIME      ?= $(shell date "+%F_%T_%Z")
COMMIT_SHA1     ?= $(shell git show -s --format=%h)
VERSION_PACKAGE	?= github.com/zhangga/swagger/version
API_PROTO_FILES	?= $(shell find swagger_file_example -name *.proto)

VERSION_BUILD_LDFLAGS= \
-X "${VERSION_PACKAGE}.Version=${SERVER_VERSION}" \
-X "${VERSION_PACKAGE}.BuildTime=${BUILD_TIME}" \
-X "${VERSION_PACKAGE}.CommitHash=${COMMIT_SHA1}"
.PHONY: build
# build
build:
	$(GO_BUILD) \
	-ldflags '$(VERSION_BUILD_LDFLAGS)' \
	-o $(OUT_DIR)/ \
	.

.PHONY: run
# run
run:
	go run main.go

.PHONY: api
# generate api proto
api:
	protoc \
		--proto_path=./swagger_file_example:./third_party \
		--openapi_out=fq_schema_naming=true,default_response=false:./swagger_file_example \
		$(API_PROTO_FILES)

.PHONY: test
# run all test
test:
	go test -race ./...

.PHONY: lint
# run all lint
lint:
	golangci-lint run -c .golangci.yml ./...

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
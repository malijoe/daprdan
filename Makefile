export GO111MODULE ?= on
export GOPROXY ?= https://proxy.golang.org
export GOSUMDB ?= sum.golang.org

GIT_COMMIT = $(shell git rev-list -1 HEAD)
GIT_VERSION ?= $(shell git describe --always --abbrev=7 --dirty)
# By default, disable CGO_ENABLED. See the details on https://golang.org/cmd/cgo
CGO ?= 0
BINARIES ?= daprdd placement operator injector sentry

PROTOC ?= protoc
# name of protoc-gen-go when protoc-gen-go --version is run.
PROTOC_GEN_GO_NAME = "protoc-gen-go"

PROTOC_GEN_GO_VERSION = v1.28.0
PROTOC_GEN_GO_NAME += $(PROTOC_GEN_GO_VERSION)

echo-path:
	echo $(PATH)	

.PHONY: modtidy
modtidy:
	go mod tidy

.PHONY: init-proto
init-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

################################################################################
# Target: gen-proto                                                            #
################################################################################
GRPC_PROTOS := common internals operator placement runtime sentry components
PROTO_PREFIX := github.com/malijoe/daprdan

# Generate archive files for each binary
# $(1): the binary name to be archived
define genProtoc
.PHONY: gen-proto-$(1)
gen-proto-$(1):
	$(PROTOC) --go_out=. --go_opt=module=$(PROTO_PREFIX) --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false,module=$(PROTO_PREFIX) ./daprd/proto/$(1)/v1/*.proto
endef

$(foreach ITEM,$(GRPC_PROTOS),$(eval $(call genProtoc,$(ITEM))))

GEN_PROTOS := $(foreach ITEM,$(GRPC_PROTOS),gen-proto-$(ITEM))

.PHONY: gen-proto
gen-proto: check-proto-version $(GEN_PROTOS) modtidy

.PHONY: check-proto-version
check-proto-version: ## Checking the version of proto related tools
	@test "$(shell protoc --version)" = "libprotoc 3.21.1" \
	|| { echo "please use protoc 3.21.1 (protobuf 21.1) to generate proto, see https://github.com/malijoe/daperdan/blob/master/daprd/README.md#proto-client-generation"; exit 1; }

	@test "$(shell protoc-gen-go-grpc --version)" = "protoc-gen-go-grpc 1.2.0" \
	|| { echo "please use protoc-gen-go-grpc 1.2.0 to generate proto, see https://github.com/malijoe/daperdan/blob/master/daprd/README.md#proto-client-generation"; exit 1; }

	@test "$(shell protoc-gen-go --version 2>&1)" = "$(PROTOC_GEN_GO_NAME)" \
	|| { echo "please use protoc-gen-go v1.28.0 to generate proto, see https://github.com/malijoe/daperdan/blob/master/daprd/README.md#proto-client-generation"; exit 1; }
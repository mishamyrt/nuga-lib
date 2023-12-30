GOLANGCI_LINT_VERSION = v1.55.2
REVIVE_VERSION = v1.3.4
GO_BIN_PATH := $(shell go env GOPATH)/bin

define build_app
	cd "cmd/$1"; go build -o ../../build/
endef

.PHONY: setup
setup:
	curl -sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(GO_BIN_PATH) $(GOLANGCI_LINT_VERSION)
	go install github.com/mgechev/revive@$(REVIVE_VERSION)

.PHONY: lint
lint:
	golangci-lint run ./...
	revive -config ./revive.toml  ./...

.PHONY: test
test:
	go test -cover ./...

.PHONY: build
build:
	$(call build_app,dumper)

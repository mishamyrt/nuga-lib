GOLANGCI_LINT_VERSION = v1.55.2
REVIVE_VERSION = v1.3.4
GIT_CHGLOG = v0.15.4
GO_BIN_PATH := $(shell go env GOPATH)/bin

define build_app
	cd "cmd/$1"; go build -o "../../build/nuga-$1"
endef

.PHONY: setup
setup:
	curl -sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(GO_BIN_PATH) $(GOLANGCI_LINT_VERSION)
	go install github.com/mgechev/revive@$(REVIVE_VERSION)
	go install github.com/git-chglog/git-chglog/cmd/git-chglog@$(GIT_CHGLOG)

.PHONY: lint
lint:
	golangci-lint run ./...
	revive -config ./revive.toml  ./...

.PHONY: test
test:
	go test -cover ./internal... ./pkg...

.PHONY: build
build:
	$(call build_app,simulation)
	$(call build_app,describe)

.PHONY: changelog
changelog:
	git-chglog -o CHANGELOG.md

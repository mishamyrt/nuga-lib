GOLANGCI_LINT_VERSION = v1.56.2
REVIVE_VERSION = v1.3.7
GIT_CHGLOG_VERSION = v0.15.4
GO_BIN_PATH := $(shell go env GOPATH)/bin
TEST_MODULES := $(shell go list ./... | grep -v /cmd/)

define build_app
	cd "cmd/$1"; go build -o "../../build/nuga-$1"
endef

.PHONY: setup
setup:
	curl -sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(GO_BIN_PATH) $(GOLANGCI_LINT_VERSION)
	go install github.com/mgechev/revive@$(REVIVE_VERSION)
	go install github.com/git-chglog/git-chglog/cmd/git-chglog@$(GIT_CHGLOG_VERSION)

.PHONY: lint
lint:
	golangci-lint run ./...
	revive -config ./revive.toml  ./...

.PHONY: test
test:
	go test -cover $(TEST_MODULES)

.PHONY: build
build:
	$(call build_app,custom-effect)
	$(call build_app,dump)
	$(call build_app,describe)

.PHONY: clean
clean:
	rm -rf build

.PHONY: check
check:
	make lint
	make test
	make build

.PHONY: changelog
changelog:
	git-chglog -o CHANGELOG.md

GO := go

GO_BUILD_PACKAGES := ./cmd/cli
GO_BUILD_BINDIR :=./bin
GIT_COMMIT := $(or $(SOURCE_GIT_COMMIT),$(shell git rev-parse --short HEAD))
GIT_TAG :="$(shell git tag | sort -V | tail -1)"

GO_LD_EXTRAFLAGS :=-X github.com/jpower432/jpower432-go-cli/version.version="$(shell git tag | sort -V | tail -1)" \
				   -X github.com/jpower432/jpower432-go-cli/version.buildData="dev" \
				   -X github.com/jpower432/jpower432-go-cli/version.commit="$(GIT_COMMIT)" \
				   -X github.com/jpower432/jpower432-go-cli/version.buildDate="$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')"

build: prep-build-dir
	$(GO) build -o $(GO_BUILD_BINDIR)/ -ldflags="$(GO_LD_EXTRAFLAGS)" $(GO_BUILD_PACKAGES)
.PHONY: build

prep-build-dir:
	mkdir -p ${GO_BUILD_BINDIR}
.PHONY: prep-build-dir

vendor:
	$(GO) mod tidy
	$(GO) mod verify
	$(GO) mod vendor
.PHONY: vendor

clean:
	@rm -rf ./$(GO_BUILD_BINDIR)/*
.PHONY: clean

test-unit:
	$(GO) test $(GO_BUILD_FLAGS) -coverprofile=coverage.out -race -count=1 ./...
.PHONY: test-unit

sanity: vendor format vet
	git diff --exit-code
.PHONY: sanity

format:
	$(GO) fmt ./...
.PHONY: format

vet:
	$(GO) vet ./...
.PHONY: vet

all: clean vendor test-unit build
.PHONY: all
ROOT=$(realpath $(dir $(realpath $(firstword $(MAKEFILE_LIST)))))

VERSION ?= 1.0.0
IMG ?= ping-api:$(VERSION)

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	go test $$(go list ./... | grep -v /e2e) -coverprofile cover.out

GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint
GOLANGCI_LINT_VERSION ?= v1.54.2
golangci-lint:
	@[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLANGCI_LINT)) $(GOLANGCI_LINT_VERSION) ;\
	}

.PHONY: lint
lint: golangci-lint  ## Run golangci-lint linter
	$(GOLANGCI_LINT) run

.PHONY: lint-fix
lint-fix: golangci-lint ## Run golangci-lint linter and perform fixes
	$(GOLANGCI_LINT) run --fix

.PHONY: build
.build: fmt vet ## Build manager binary.
	go build -o bin/ping-api main.go

.PHONY: run
run: fmt vet ## Run a app from your host.
	go run main.go

.PHONY: docker-build
docker-build:  ## Build docker image.
	docker rmi -f ${IMG}
	docker build -t ${IMG} .

.PHONY: docker-run
docker-run:  ## Run docker image.
	docker rm -f ping-api
	docker run -d --name ping-api --restart unless-stopped -p 8005:8080 ${IMG}

#arxiv_tests:
	#go test -run ExampleGetComputerScienceCategoryArticles ./pkg/arxiv -count=1 -v

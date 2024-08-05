-include .env
BIN_DIR := $(GOPATH)/bin

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
INTEGRATION_DIR := ./test/integration
FIXTURES_DIR    := $(INTEGRATION_DIR)/fixtures

TEST_TIMEOUT := 2h

SKIP_DOCKER       ?= 0
GOLANGCILINT      := golangci-lint
GOLANGCILINT_IMG  := golangci/golangci-lint:v1.46.2-alpine
GOLANGCILINT_ARGS := run

LINODE_URL := https://api.linode.com/

PACKAGES := $(shell go list ./... | grep -v integration)

SKIP_LINT ?= 0

.PHONY: build vet test refresh-fixtures clean-fixtures lint run_fixtures sanitize fixtures godoc testint testunit tidy

test: build lint testunit testint

citest: lint test

testunit:
	go test -v $(PACKAGES) $(ARGS)

testint:
	cd test && make test

build: vet lint
	go build ./...
	cd k8s && go build ./...

vet:
	go vet ./...
	cd k8s && go vet ./...

lint:
ifeq ($(SKIP_LINT), 1)
	@echo Skipping lint stage
else ifeq ($(SKIP_DOCKER), 1)
	$(GOLANGCILINT) $(GOLANGCILINT_ARGS)
else
	docker run --rm -v $(shell pwd):/app -w /app $(GOLANGCILINT_IMG) $(GOLANGCILINT) $(GOLANGCILINT_ARGS)
endif

clean-fixtures:
	@-rm fixtures/*.yaml

refresh-fixtures: clean-fixtures fixtures

run_fixtures:
	@echo "* Running fixtures"
	cd $(INTEGRATION_DIR) && \
	LINODE_FIXTURE_MODE="record" \
	LINODE_TOKEN=$(LINODE_TOKEN) \
	LINODE_API_VERSION="v4beta" \
	LINODE_URL="$(LINODE_URL)" \
	GO111MODULE="on" \
	go test -timeout=$(TEST_TIMEOUT) -v $(ARGS)

sanitize:
	@echo "* Sanitizing fixtures"
	@for yaml in $(FIXTURES_DIR)/*yaml; do \
		sed -E -i.bak \
			-e 's_stats/20[0-9]{2}/[1-9][0-2]?_stats/2018/1_g' \
			-e 's/(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))/1234::5678/g' \
			$$yaml; \
	done
	@find $(FIXTURES_DIR) -name *yaml.bak -exec rm {} \;

fixtures: run_fixtures sanitize

godoc:
	@godoc -http=:6060

tidy:
	go mod tidy
	cd k8s && go mod tidy
	cd test && go mod tidy
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
INTEGRATION_DIR := ./test/integration
FIXTURES_DIR    := $(INTEGRATION_DIR)/fixtures

>>>>>>> 5ce8c7613 (update vendored files)
GOLANGCILINT      := golangci-lint
GOLANGCILINT_IMG  := golangci/golangci-lint:v1.38-alpine
GOLANGCILINT_ARGS := run

PACKAGES := $(shell go list ./... | grep -v integration)

SKIP_LINT ?= 0

.PHONY: build vet test refresh-fixtures clean-fixtures lint run_fixtures sanitize fixtures godoc testint testunit tidy

test: build lint testunit testint

citest: lint test

testunit:
	go test -v $(PACKAGES) $(ARGS)

testint:
	cd test && make test

build: vet lint
	go build ./...
	cd k8s && go build ./...

vet:
	go vet ./...
	cd k8s && go vet ./...

lint:
ifeq ($(SKIP_LINT), 1)
	@echo Skipping lint stage
else
	docker run --rm -v $(shell pwd):/app -w /app $(GOLANGCILINT_IMG) $(GOLANGCILINT) run
endif

clean-fixtures:
	@-rm fixtures/*.yaml

refresh-fixtures: clean-fixtures fixtures

run_fixtures:
	@echo "* Running fixtures"
	cd $(INTEGRATION_DIR) && \
	LINODE_FIXTURE_MODE="record" \
	LINODE_TOKEN=$(LINODE_TOKEN) \
	LINODE_API_VERSION="v4beta" \
	GO111MODULE="on" \
	go test -timeout=60m -v $(ARGS)

sanitize:
	@echo "* Sanitizing fixtures"
	@for yaml in $(FIXTURES_DIR)/*yaml; do \
		sed -E -i.bak \
			-e 's_stats/20[0-9]{2}/[1-9][0-2]?_stats/2018/1_g' \
			-e 's/(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))/1234::5678/g' \
			$$yaml; \
	done
	@find $(FIXTURES_DIR) -name *yaml.bak -exec rm {} \;

fixtures: run_fixtures sanitize

godoc:
	@godoc -http=:6060
<<<<<<< HEAD
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======

tidy:
	go mod tidy
	cd k8s && go mod tidy
	cd test && go mod tidy
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
INTEGRATION_DIR := ./test/integration
FIXTURES_DIR    := $(INTEGRATION_DIR)/fixtures

>>>>>>> 6b7ce455e (update vendored files)
GOLANGCILINT      := golangci-lint
GOLANGCILINT_IMG  := golangci/golangci-lint:v1.38-alpine
GOLANGCILINT_ARGS := run

PACKAGES := $(shell go list ./... | grep -v integration)

SKIP_LINT ?= 0

.PHONY: build vet test refresh-fixtures clean-fixtures lint run_fixtures sanitize fixtures godoc testint testunit tidy

test: build lint testunit testint

citest: lint test

testunit:
	go test -v $(PACKAGES) $(ARGS)

testint:
	cd test && make test

build: vet lint
	go build ./...
	cd k8s && go build ./...

vet:
	go vet ./...
	cd k8s && go vet ./...

lint:
ifeq ($(SKIP_LINT), 1)
	@echo Skipping lint stage
else
	docker run --rm -v $(shell pwd):/app -w /app $(GOLANGCILINT_IMG) $(GOLANGCILINT) run
endif

clean-fixtures:
	@-rm fixtures/*.yaml

refresh-fixtures: clean-fixtures fixtures

run_fixtures:
	@echo "* Running fixtures"
	cd $(INTEGRATION_DIR) && \
	LINODE_FIXTURE_MODE="record" \
	LINODE_TOKEN=$(LINODE_TOKEN) \
	LINODE_API_VERSION="v4beta" \
	GO111MODULE="on" \
	go test -timeout=60m -v $(ARGS)

sanitize:
	@echo "* Sanitizing fixtures"
	@for yaml in $(FIXTURES_DIR)/*yaml; do \
		sed -E -i.bak \
			-e 's_stats/20[0-9]{2}/[1-9][0-2]?_stats/2018/1_g' \
			-e 's/(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))/1234::5678/g' \
			$$yaml; \
	done
	@find $(FIXTURES_DIR) -name *yaml.bak -exec rm {} \;

fixtures: run_fixtures sanitize

godoc:
	@godoc -http=:6060
<<<<<<< HEAD
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======

tidy:
	go mod tidy
	cd k8s && go mod tidy
	cd test && go mod tidy
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
INTEGRATION_DIR := ./test/integration
FIXTURES_DIR    := $(INTEGRATION_DIR)/fixtures

>>>>>>> 4d7e5ad26 (update vendored files)
GOLANGCILINT      := golangci-lint
GOLANGCILINT_IMG  := golangci/golangci-lint:v1.38-alpine
GOLANGCILINT_ARGS := run

PACKAGES := $(shell go list ./... | grep -v integration)

SKIP_LINT ?= 0

.PHONY: build vet test refresh-fixtures clean-fixtures lint run_fixtures sanitize fixtures godoc testint testunit tidy

test: build lint testunit testint

citest: lint test

testunit:
	go test -v $(PACKAGES) $(ARGS)

testint:
	cd test && make test

build: vet lint
	go build ./...
	cd k8s && go build ./...

vet:
	go vet ./...
	cd k8s && go vet ./...

lint:
ifeq ($(SKIP_LINT), 1)
	@echo Skipping lint stage
else
	docker run --rm -v $(shell pwd):/app -w /app $(GOLANGCILINT_IMG) $(GOLANGCILINT) run
endif

clean-fixtures:
	@-rm fixtures/*.yaml

refresh-fixtures: clean-fixtures fixtures

run_fixtures:
	@echo "* Running fixtures"
	cd $(INTEGRATION_DIR) && \
	LINODE_FIXTURE_MODE="record" \
	LINODE_TOKEN=$(LINODE_TOKEN) \
	LINODE_API_VERSION="v4beta" \
	GO111MODULE="on" \
	go test -timeout=60m -v $(ARGS)

sanitize:
	@echo "* Sanitizing fixtures"
	@for yaml in $(FIXTURES_DIR)/*yaml; do \
		sed -E -i.bak \
			-e 's_stats/20[0-9]{2}/[1-9][0-2]?_stats/2018/1_g' \
			-e 's/(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))/1234::5678/g' \
			$$yaml; \
	done
	@find $(FIXTURES_DIR) -name *yaml.bak -exec rm {} \;

fixtures: run_fixtures sanitize

godoc:
	@godoc -http=:6060
<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======

tidy:
	go mod tidy
	cd k8s && go mod tidy
	cd test && go mod tidy
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
INTEGRATION_DIR := ./test/integration
FIXTURES_DIR    := $(INTEGRATION_DIR)/fixtures

TEST_TIMEOUT := 5h

SKIP_DOCKER       ?= 0
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
GOLANGCILINT      := golangci-lint
GOLANGCILINT_IMG  := golangci/golangci-lint:latest
GOLANGCILINT_ARGS := run

LINODE_URL ?= https://api.linode.com/

PACKAGES := $(shell go list ./... | grep -v integration)

SKIP_LINT ?= 0

.PHONY: build vet test refresh-fixtures clean clean-cov clean-fixtures lint run_fixtures sanitize fixtures godoc testint testunit testcov tidy

test: build lint testunit testint

citest: lint test

testunit:
	go test -v $(PACKAGES) $(ARGS)
	cd test && make testunit

testint:
	cd test && make testint

testcov-func:
	@go test -v -coverprofile="coverage.txt" . > /dev/null 2>&1
	@go tool cover -func coverage.txt

# Note for WSL Users, set BROWSER=wslview
testcov-html:
	@go test -v -coverprofile="coverage.txt" . > /dev/null 2>&1
	@go tool cover -html coverage.txt

smoketest:
	cd test && make smoketest

build: vet lint
	go build ./...
	cd k8s && go build ./...

vet:
	go vet ./...
	cd k8s && go vet ./...

lint:
ifeq ($(SKIP_LINT), 1)
	@echo Skipping lint stage
else ifeq ($(SKIP_DOCKER), 1)
	$(GOLANGCILINT) $(GOLANGCILINT_ARGS)
else
	docker run --rm -v $(shell pwd):/app -w /app $(GOLANGCILINT_IMG) $(GOLANGCILINT) $(GOLANGCILINT_ARGS)
endif

clean: clean-cov

clean-cov:
	@-rm -f coverage.txt

clean-fixtures:
	@-rm -f fixtures/*.yaml

refresh-fixtures: clean-fixtures fixtures

run_fixtures:
	@echo "* Running fixtures"
	cd $(INTEGRATION_DIR) && \
	LINODE_FIXTURE_MODE="record" \
	LINODE_TOKEN=$(LINODE_TOKEN) \
	LINODE_API_VERSION="v4beta" \
	LINODE_URL="$(LINODE_URL)" \
	GO111MODULE="on" \
	go test -timeout=$(TEST_TIMEOUT) -v $(ARGS)

sanitize:
	@echo "* Sanitizing fixtures"
	@for yaml in $(FIXTURES_DIR)/*yaml; do \
		sed -E -i.bak \
			-e 's_stats/20[0-9]{2}/[1-9][0-2]?_stats/2018/1_g' \
			-e 's/(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))/1234::5678/g' \
			$$yaml; \
	done
	@find $(FIXTURES_DIR) -name *yaml.bak -exec rm {} \;

fixtures: run_fixtures sanitize

godoc:
	@godoc -http=:6060
<<<<<<< HEAD
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======

tidy:
	go mod tidy
	cd k8s && go mod tidy
	cd test && go mod tidy
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

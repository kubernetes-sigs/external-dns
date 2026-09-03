# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#? cover: Creates coverage report for whole project excluding vendor and opens result in the default browser
.PHONY: cover cover-html
.DEFAULT_GOAL := build

# Use the Go pinned in mise.toml, never a toolchain go.mod would download.
export GOTOOLCHAIN := local

cover:
	@go test -cover -coverprofile=cover.out -v ./...

#? cover-html: Run tests with coverage and open coverage report in the browser
cover-html: cover
	@go tool cover -html=cover.out

#? check-tools: Compare the tools on PATH against the versions pinned in mise.toml
.PHONY: check-tools
check-tools:
	@scripts/check-tools.sh

#? go-lint: Run the golangci-lint tool
.PHONY: go-lint
go-lint:
	golangci-lint config verify
	gofmt -l -s -w .
	golangci-lint run --timeout=30m --fix ./...

#? licensecheck: Run the to check for license headers
.PHONY: licensecheck
licensecheck:
	@echo ">> checking license header"
	@licRes=$$(for file in $$(find . -type f -iname '*.go' ! -path './vendor/*') ; do \
			awk 'NR<=5' $$file | grep -Eq "(Copyright|generated|GENERATED)" || echo $$file; \
		done); \
		if [ -n "$${licRes}" ]; then \
			echo "license header checking failed:"; echo "$${licRes}"; \
			exit 1; \
		fi

#? lint: Run all the linters
.PHONY: lint
lint: check-tools licensecheck go-lint validate-json-yaml

#? validate-json-yaml: Validate JSON and YAML files
.PHONY: validate-json-yaml
validate-json-yaml:
	bash scripts/validate-json-yaml.sh

#? crd: Generates CRD using controller-gen and copy it into chart
.PHONY: crd
crd: check-tools
	@./scripts/generate-crd.sh

#? update-tools: Bump every tool in mise.toml to its latest version
.PHONY: update-tools
update-tools:
	@mise upgrade --bump
	@$(MAKE) check-tools

#? test: The verify target runs tasks similar to the CI tasks, but without code coverage
.PHONY: test
test:
	go test -race ./...


.PHONY: test
go-test:
	go test -race -coverprofile=profile.cov ./...
	go tool cover -func=profile.cov > coverage.summary
	@tail -n 1 coverage.summary

#? build: The build targets allow to build the binary and container image
.PHONY: build

BINARY        ?= external-dns
SOURCES        = $(shell find . -name '*.go')
IMAGE_STAGING  = gcr.io/k8s-staging-external-dns/$(BINARY)
REGISTRY      ?= us.gcr.io/k8s-artifacts-prod/external-dns
IMAGE         ?= $(REGISTRY)/$(BINARY)
VERSION       ?= $(shell git describe --tags --always --dirty --match "v*")
GIT_REVISION  ?= $(shell git rev-parse HEAD)
GIT_COMMIT    ?= $(shell git rev-parse --short HEAD)
GIT_COMMIT    := $(or $(GIT_COMMIT),$(shell echo "$(GIT_REVISION)" | cut -c1-7))
BUILD_FLAGS   ?= -v
LDFLAGS       ?= -X sigs.k8s.io/external-dns/pkg/apis/externaldns.Version=$(VERSION) -w -s
LDFLAGS       += -X sigs.k8s.io/external-dns/pkg/apis/externaldns.GitCommit=$(GIT_COMMIT)
ARCH          ?= amd64
SHELL          = /bin/bash
IMG_PLATFORM  ?= linux/amd64,linux/arm64,linux/arm/v7
IMG_PUSH      ?= true
IMG_SBOM      ?= none

build: check-tools build/$(BINARY)

build/$(BINARY): $(SOURCES)
	CGO_ENABLED=0 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

build.push/multiarch: ko
	KO_DOCKER_REPO=${IMAGE} \
	VERSION=${VERSION} \
	ko build --tags ${VERSION} --bare --sbom ${IMG_SBOM} \
		--image-label org.opencontainers.image.source="https://github.com/kubernetes-sigs/external-dns" \
		--image-label org.opencontainers.image.revision=$(GIT_REVISION) \
		--platform=${IMG_PLATFORM}  --push=${IMG_PUSH} .

build.image/multiarch:
	$(MAKE) IMG_PUSH=false build.push/multiarch

build.image:
	$(MAKE) IMG_PLATFORM=linux/$(ARCH) build.image/multiarch

build.image-amd64:
	$(MAKE) ARCH=amd64 build.image

build.image-arm64:
	$(MAKE) ARCH=arm64 build.image

build.image-arm/v7:
	$(MAKE) ARCH=arm/v7 build.image

.PHONY: build.image-fips
#? build.image-fips: Build a FIPS 140-3 image variant (tagged <version>-fips); see scripts/build-fips-image.sh --help
build.image-fips: ko
	scripts/build-fips-image.sh

build.push:
	$(MAKE) IMG_PLATFORM=linux/$(ARCH) build.push/multiarch

build.push-amd64:
	$(MAKE) ARCH=amd64 build.push

build.push-arm64:
	$(MAKE) ARCH=arm64 build.push

build.push-arm/v7:
	$(MAKE) ARCH=arm/v7 build.push

build.arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

build.amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

build.arm/v7:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

clean:
	@rm -rf build
	@go clean -cache

.PHONY: release.staging
#? release.staging: Builds and push container images to the staging bucket.
release.staging: test
	IMAGE=$(IMAGE_STAGING) $(MAKE) build.push/multiarch

release.prod: test
	$(MAKE) build.push/multiarch

# cloudbuild.yaml runs make inside golang:*-bookworm, which has no mise.
.PHONY: ko
ko:
	scripts/install-ko.sh

.PHONY: generate-flags-documentation
#? generate-flags-documentation: Generate documentation (docs/flags.md)
generate-flags-documentation:
	go run internal/gen/docs/flags/main.go

.PHONY: generate-metrics-documentation
#? generate-metrics-documentation: Generate documentation (docs/monitoring/metrics.md)
generate-metrics-documentation:
	go run internal/gen/docs/metrics/main.go

.PHONY: generate-sources-documentation
#? generate-sources-documentation: Generate documentation (docs/sources/index.md)
generate-sources-documentation:
	go run internal/gen/docs/sources/main.go

#? file-hygiene: Run repository file hygiene checks (replaces pre-commit)
file-hygiene:
	@./scripts/file-hygiene.sh

.PHONY: help
#? help: Get more info on available commands
help: Makefile
	@sed -n 's/^#?//p' $< | column -t -s ':' |  sort | sed -e 's/^/ /'

#? helm-test: Run unit tests
helm-test:
	scripts/helm-tools.sh --helm-unittest

#? helm-template: Run helm template
helm-template:
	scripts/helm-tools.sh --helm-template

#? helm-lint: Run helm linting (schema,docs)
helm-lint:
	scripts/helm-tools.sh --schema
	scripts/helm-tools.sh --docs

.PHONY: go-dependency
#? go-dependency: Dependency maintenance
go-dependency:
	go mod tidy

.PHONY: mkdocs-serve
#? mkdocs-serve: Run the builtin development server for mkdocs
mkdocs-serve:
	@$(info "contribute to documentation docs/contributing/dev-guide.md")
	@mkdocs serve

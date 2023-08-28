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

# cover-html creates coverage report for whole project excluding vendor and opens result in the default browser
.PHONY: cover cover-html
.DEFAULT_GOAL := build

cover:
	go get github.com/wadey/gocovmerge
	$(eval PKGS := $(shell go list ./... | grep -v /vendor/))
	$(eval PKGS_DELIM := $(shell echo $(PKGS) | tr / -'))
	go list -f '{{if or (len .TestGoFiles) (len .XTestGoFiles)}}go test -test.v -test.timeout=120s -covermode=count -coverprofile={{.Name}}_{{len .Imports}}_{{len .Deps}}.coverprofile -coverpkg $(PKGS_DELIM) {{.ImportPath}}{{end}}' $(PKGS) | xargs -0 sh -c
	gocovmerge `ls *.coverprofile` > cover.out
	rm *.coverprofile

cover-html: cover
	go tool cover -html cover.out

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.5.0 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

golangci-lint:
	@command -v golangci-lint > /dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.3

# Run the golangci-lint tool
.PHONY: go-lint
go-lint: golangci-lint
	golangci-lint run --timeout=30m ./...

# Run the licensecheck script to check for license headers
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

# Run all the linters
.PHONY: lint
lint: licensecheck go-lint

# generates CRD using controller-gen
.PHONY: crd
crd: controller-gen
	${CONTROLLER_GEN} crd:crdVersions=v1 paths="./endpoint/..." output:crd:stdout > docs/contributing/crd-source/crd-manifest.yaml

# The verify target runs tasks similar to the CI tasks, but without code coverage
.PHONY: test
test:
	go test -race -coverprofile=profile.cov ./...

# The build targets allow to build the binary and container image
.PHONY: build

BINARY        ?= external-dns
SOURCES        = $(shell find . -name '*.go')
IMAGE_STAGING  = gcr.io/k8s-staging-external-dns/$(BINARY)
REGISTRY      ?= us.gcr.io/k8s-artifacts-prod/external-dns
IMAGE         ?= $(REGISTRY)/$(BINARY)
VERSION       ?= $(shell git describe --tags --always --dirty --match "v*")
BUILD_FLAGS   ?= -v
LDFLAGS       ?= -X sigs.k8s.io/external-dns/pkg/apis/externaldns.Version=$(VERSION) -w -s
ARCH          ?= amd64
SHELL          = /bin/bash
IMG_PLATFORM  ?= linux/amd64,linux/arm64,linux/arm/v7
IMG_PUSH      ?= true
IMG_SBOM      ?= none

build: build/$(BINARY)

build/$(BINARY): $(SOURCES)
	CGO_ENABLED=0 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

build.push/multiarch: ko
	KO_DOCKER_REPO=${IMAGE} \
    VERSION=${VERSION} \
    ko build --tags ${VERSION} --platform=${IMG_PLATFORM} --bare --sbom ${IMG_SBOM} --push=${IMG_PUSH} .

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

 # Builds and push container images to the staging bucket.
.PHONY: release.staging

release.staging: test
	IMAGE=$(IMAGE_STAGING) $(MAKE) build.push/multiarch

release.prod: test
	$(MAKE) build.push/multiarch

.PHONY: ko
ko:
	scripts/install-ko.sh

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

.PHONY: go-lint

# Run the golangci-lint tool
go-lint:
	golangci-lint run --timeout=15m ./...

.PHONY: licensecheck

# Run the licensecheck script to check for license headers
licensecheck:
	@echo ">> checking license header"
	@licRes=$$(for file in $$(find . -type f -iname '*.go' ! -path './vendor/*') ; do \
               awk 'NR<=5' $$file | grep -Eq "(Copyright|generated|GENERATED)" || echo $$file; \
       done); \
       if [ -n "$${licRes}" ]; then \
               echo "license header checking failed:"; echo "$${licRes}"; \
               exit 1; \
       fi

.PHONY: lint

# Run all the linters
lint: licensecheck go-lint

.PHONY: crd

# generates CRD using controller-gen
crd: controller-gen
	${CONTROLLER_GEN} crd:crdVersions=v1 paths="./endpoint/..." output:crd:stdout > docs/contributing/crd-source/crd-manifest.yaml

# The verify target runs tasks similar to the CI tasks, but without code coverage
.PHONY: verify test

test:
	go test -race -coverprofile=profile.cov ./...

# The build targets allow to build the binary and docker image
.PHONY: build build.docker build.mini

BINARY        ?= external-dns
SOURCES        = $(shell find . -name '*.go')
IMAGE_STAGING  = gcr.io/k8s-staging-external-dns/$(BINARY)
IMAGE         ?= us.gcr.io/k8s-artifacts-prod/external-dns/$(BINARY)
VERSION       ?= $(shell git describe --tags --always --dirty --match "v*")
BUILD_FLAGS   ?= -v
LDFLAGS       ?= -X sigs.k8s.io/external-dns/pkg/apis/externaldns.Version=$(VERSION) -w -s
ARCHS          = amd64 arm64 arm/v7
ARCH          ?= amd64
DEFAULT_ARCH   = amd64
SHELL          = /bin/bash
OUTPUT_TYPE   ?= docker


build: build/$(BINARY)

build/$(BINARY): $(SOURCES)
	CGO_ENABLED=0 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

build.push/multiarch: $(addprefix build.push-,$(ARCHS))
	arch_specific_tags=()
	for arch in $(ARCHS); do \
		image="$(IMAGE):$(VERSION)-$$(echo $$arch | tr / -)" ;\
		arch_specific_tags+=( " $${image}" ) ;\
	done ;\
	echo $${arch_specific_tags[@]} ;\
        DOCKER_CLI_EXPERIMENTAL=enabled docker buildx imagetools create --tag "$(IMAGE):$(VERSION)" $${arch_specific_tags[@]} ;\

build.image/multiarch: $(addprefix build.image-,$(ARCHS))

build.image:
	$(MAKE) ARCH=$(ARCH) OUTPUT_TYPE=docker build.docker

build.image-amd64:
	$(MAKE) ARCH=amd64 build.image

build.image-arm64:
	$(MAKE) ARCH=arm64 build.image

build.image-arm/v7:
	$(MAKE) ARCH=arm/v7 build.image

build.push:
	$(MAKE) ARCH=$(ARCH) OUTPUT_TYPE=registry build.docker

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

build.setup:
	docker buildx inspect img-builder > /dev/null || docker buildx create --name img-builder --use

build.docker: build.setup build.$(ARCH)
	docker build --rm --tag "$(IMAGE):$(VERSION)" --build-arg VERSION="$(VERSION)" --build-arg ARCH="$(ARCH)" .
	image="$(IMAGE):$(VERSION)-$(subst /,-,$(ARCH))"; \
	docker buildx build \
		--pull \
		--provenance=false \
		--sbom=false \
		--output=type=$(OUTPUT_TYPE) \
		--platform linux/$(ARCH) \
		--build-arg ARCH="$(ARCH)" \
		--build-arg VERSION="$(VERSION)" \
		--tag $${image} .

build.mini:
	docker build --rm --tag "$(IMAGE):$(VERSION)-mini" --build-arg VERSION="$(VERSION)" -f Dockerfile.mini .

clean:
	@rm -rf build

 # Builds and push container images to the staging bucket.
.PHONY: release.staging

release.staging: test
	IMAGE=$(IMAGE_STAGING) $(MAKE) build.push/multiarch

release.prod: test
	$(MAKE) build.push/multiarch

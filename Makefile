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
	$(eval PKGS_DELIM := $(shell echo $(PKGS) | sed -e 's/ /,/g'))
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
VERSION       ?= $(shell git describe --tags --always --dirty)
BUILD_FLAGS   ?= -v
LDFLAGS       ?= -X sigs.k8s.io/external-dns/pkg/apis/externaldns.Version=$(VERSION) -w -s
ARCHS         = amd64 arm64v8 arm32v7
SHELL         = /bin/bash


build: build/$(BINARY)

build/$(BINARY): $(SOURCES)
	CGO_ENABLED=0 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

build.push/multiarch:
	arch_specific_tags=()
	for arch in $(ARCHS); do \
		image="$(IMAGE):$(VERSION)-$${arch}" ;\
		# pre-pull due to https://github.com/kubernetes-sigs/cluster-addons/pull/84/files ;\
		docker pull $${arch}/alpine:3.14 ;\
		docker pull golang:1.18 ;\
		DOCKER_BUILDKIT=1 docker build --rm --tag $${image} --build-arg VERSION="$(VERSION)" --build-arg ARCH="$${arch}" . ;\
		docker push $${image} ;\
		arch_specific_tags+=( "--amend $${image}" ) ;\
	done ;\
	DOCKER_CLI_EXPERIMENTAL=enabled docker manifest create "$(IMAGE):$(VERSION)" $${arch_specific_tags[@]} ;\
	for arch in $(ARCHS); do \
		DOCKER_CLI_EXPERIMENTAL=enabled docker manifest annotate --arch $${arch} "$(IMAGE):$(VERSION)" "$(IMAGE):$(VERSION)-$${arch}" ;\
	done;\
	DOCKER_CLI_EXPERIMENTAL=enabled docker manifest push "$(IMAGE):$(VERSION)" \

build.push: build.docker
	docker push "$(IMAGE):$(VERSION)"

build.arm64v8:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

build.amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

build.arm32v7:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o build/$(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)" .

build.docker:
	docker build --rm --tag "$(IMAGE):$(VERSION)" --build-arg VERSION="$(VERSION)" --build-arg ARCH="amd64" .

build.mini:
	docker build --rm --tag "$(IMAGE):$(VERSION)-mini" --build-arg VERSION="$(VERSION)" -f Dockerfile.mini .

clean:
	@rm -rf build

 # Builds and push container images to the staging bucket.
.PHONY: release.staging

release.staging:
	IMAGE=$(IMAGE_STAGING) $(MAKE) build.push/multiarch

release.prod:
	$(MAKE) build.push/multiarch

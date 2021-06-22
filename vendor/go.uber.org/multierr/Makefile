# Directory to put `go install`ed binaries in.
export GOBIN ?= $(shell pwd)/bin

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test -race ./...

.PHONY: gofmt
gofmt:
	$(eval FMT_LOG := $(shell mktemp -t gofmt.XXXXX))
	@gofmt -e -s -l $(GO_FILES) > $(FMT_LOG) || true
	@[ ! -s "$(FMT_LOG)" ] || (echo "gofmt failed:" | cat - $(FMT_LOG) && false)

.PHONY: golint
golint:
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	@cd tools && go install golang.org/x/lint/golint
	@$(GOBIN)/golint ./...

.PHONY: staticcheck
staticcheck:
	@cd tools && go install honnef.co/go/tools/cmd/staticcheck
	@$(GOBIN)/staticcheck ./...

.PHONY: lint
lint: gofmt golint staticcheck

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html

update-license:
	@cd tools && go install go.uber.org/tools/update-license
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	@go install golang.org/x/lint/golint
||||||| parent of 5ce8c7613 (update vendored files)
	@go install golang.org/x/lint/golint
=======
	@cd tools && go install golang.org/x/lint/golint
>>>>>>> 5ce8c7613 (update vendored files)
	@$(GOBIN)/golint ./...

.PHONY: staticcheck
staticcheck:
	@cd tools && go install honnef.co/go/tools/cmd/staticcheck
	@$(GOBIN)/staticcheck ./...

.PHONY: lint
lint: gofmt golint staticcheck

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html

update-license:
<<<<<<< HEAD
	@go install go.uber.org/tools/update-license
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	@go install go.uber.org/tools/update-license
=======
	@cd tools && go install go.uber.org/tools/update-license
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	@go install golang.org/x/lint/golint
||||||| parent of 6b7ce455e (update vendored files)
	@go install golang.org/x/lint/golint
=======
	@cd tools && go install golang.org/x/lint/golint
>>>>>>> 6b7ce455e (update vendored files)
	@$(GOBIN)/golint ./...

.PHONY: staticcheck
staticcheck:
	@cd tools && go install honnef.co/go/tools/cmd/staticcheck
	@$(GOBIN)/staticcheck ./...

.PHONY: lint
lint: gofmt golint staticcheck

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html

update-license:
<<<<<<< HEAD
	@go install go.uber.org/tools/update-license
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	@go install go.uber.org/tools/update-license
=======
	@cd tools && go install go.uber.org/tools/update-license
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	@go install golang.org/x/lint/golint
||||||| parent of 4d7e5ad26 (update vendored files)
	@go install golang.org/x/lint/golint
=======
	@cd tools && go install golang.org/x/lint/golint
>>>>>>> 4d7e5ad26 (update vendored files)
	@$(GOBIN)/golint ./...

.PHONY: staticcheck
staticcheck:
	@cd tools && go install honnef.co/go/tools/cmd/staticcheck
	@$(GOBIN)/staticcheck ./...

.PHONY: lint
lint: gofmt golint staticcheck

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html

update-license:
<<<<<<< HEAD
	@go install go.uber.org/tools/update-license
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	@go install go.uber.org/tools/update-license
=======
	@cd tools && go install go.uber.org/tools/update-license
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	@go install golang.org/x/lint/golint
	@$(GOBIN)/golint ./...

.PHONY: staticcheck
staticcheck:
	@go install honnef.co/go/tools/cmd/staticcheck
	@$(GOBIN)/staticcheck ./...

.PHONY: lint
lint: gofmt golint staticcheck

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html

update-license:
	@go install go.uber.org/tools/update-license
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	@$(GOBIN)/update-license $(GO_FILES)

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
# Directory to put `go install`ed binaries in.
export GOBIN ?= $(shell pwd)/bin

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

.PHONY: bench
bench:
	go test -bench=. ./...

bin/golint: tools/go.mod
	@cd tools && go install golang.org/x/lint/golint

bin/staticcheck: tools/go.mod
	@cd tools && go install honnef.co/go/tools/cmd/staticcheck

.PHONY: build
build:
	go build ./...

.PHONY: cover
cover:
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: gofmt
gofmt:
	$(eval FMT_LOG := $(shell mktemp -t gofmt.XXXXX))
	@gofmt -e -s -l $(GO_FILES) > $(FMT_LOG) || true
	@[ ! -s "$(FMT_LOG)" ] || (echo "gofmt failed:" | cat - $(FMT_LOG) && false)

.PHONY: golint
golint: bin/golint
	@$(GOBIN)/golint -set_exit_status ./...

.PHONY: lint
lint: gofmt golint staticcheck

.PHONY: staticcheck
staticcheck: bin/staticcheck
	@$(GOBIN)/staticcheck ./...

.PHONY: test
test:
	go test -race ./...
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
test:
	go test .
||||||| parent of 5ce8c7613 (update vendored files)
test:
	go test .
=======
# Directory to put `go install`ed binaries in.
export GOBIN ?= $(shell pwd)/bin

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

.PHONY: bench
bench:
	go test -bench=. ./...

bin/golint: tools/go.mod
	@cd tools && go install golang.org/x/lint/golint

bin/staticcheck: tools/go.mod
	@cd tools && go install honnef.co/go/tools/cmd/staticcheck

.PHONY: build
build:
	go build ./...

.PHONY: cover
>>>>>>> 5ce8c7613 (update vendored files)
cover:
<<<<<<< HEAD
	go test . -coverprofile=cover.out
	go tool cover -html=cover.out
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	go test . -coverprofile=cover.out
	go tool cover -html=cover.out
=======
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: gofmt
gofmt:
	$(eval FMT_LOG := $(shell mktemp -t gofmt.XXXXX))
	@gofmt -e -s -l $(GO_FILES) > $(FMT_LOG) || true
	@[ ! -s "$(FMT_LOG)" ] || (echo "gofmt failed:" | cat - $(FMT_LOG) && false)

.PHONY: golint
golint: bin/golint
	@$(GOBIN)/golint -set_exit_status ./...

.PHONY: lint
lint: gofmt golint staticcheck

.PHONY: staticcheck
staticcheck: bin/staticcheck
	@$(GOBIN)/staticcheck ./...

.PHONY: test
test:
	go test -race ./...
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
test:
	go test .
||||||| parent of 6b7ce455e (update vendored files)
test:
	go test .
=======
# Directory to put `go install`ed binaries in.
export GOBIN ?= $(shell pwd)/bin

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

.PHONY: bench
bench:
	go test -bench=. ./...

bin/golint: tools/go.mod
	@cd tools && go install golang.org/x/lint/golint

bin/staticcheck: tools/go.mod
	@cd tools && go install honnef.co/go/tools/cmd/staticcheck

.PHONY: build
build:
	go build ./...

.PHONY: cover
>>>>>>> 6b7ce455e (update vendored files)
cover:
<<<<<<< HEAD
	go test . -coverprofile=cover.out
	go tool cover -html=cover.out
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	go test . -coverprofile=cover.out
	go tool cover -html=cover.out
=======
	go test -coverprofile=cover.out -coverpkg=./... -v ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: gofmt
gofmt:
	$(eval FMT_LOG := $(shell mktemp -t gofmt.XXXXX))
	@gofmt -e -s -l $(GO_FILES) > $(FMT_LOG) || true
	@[ ! -s "$(FMT_LOG)" ] || (echo "gofmt failed:" | cat - $(FMT_LOG) && false)

.PHONY: golint
golint: bin/golint
	@$(GOBIN)/golint -set_exit_status ./...

.PHONY: lint
lint: gofmt golint staticcheck

.PHONY: staticcheck
staticcheck: bin/staticcheck
	@$(GOBIN)/staticcheck ./...

.PHONY: test
test:
	go test -race ./...
>>>>>>> 6b7ce455e (update vendored files)

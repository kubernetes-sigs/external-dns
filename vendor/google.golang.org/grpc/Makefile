all: vet test testrace

<<<<<<< HEAD
<<<<<<< HEAD
build:
	go build google.golang.org/grpc/...

clean:
	go clean -i google.golang.org/grpc/...

deps:
	GO111MODULE=on go get -d -v google.golang.org/grpc/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate google.golang.org/grpc/...

test:
	go test -cpu 1,4 -timeout 7m google.golang.org/grpc/...

testsubmodule:
	cd security/advancedtls && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/advancedtls/...
	cd security/authorization && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/authorization/...

testrace:
	go test -race -cpu 1,4 -timeout 7m google.golang.org/grpc/...

testdeps:
	GO111MODULE=on go get -d -v -t google.golang.org/grpc/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testrace \
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
build: deps
||||||| parent of 5ce8c7613 (update vendored files)
build: deps
=======
build:
>>>>>>> 5ce8c7613 (update vendored files)
	go build google.golang.org/grpc/...

clean:
	go clean -i google.golang.org/grpc/...

deps:
	GO111MODULE=on go get -d -v google.golang.org/grpc/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate google.golang.org/grpc/...

test:
	go test -cpu 1,4 -timeout 7m google.golang.org/grpc/...

testsubmodule:
	cd security/advancedtls && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/advancedtls/...
	cd security/authorization && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/authorization/...

testrace:
	go test -race -cpu 1,4 -timeout 7m google.golang.org/grpc/...

testdeps:
	GO111MODULE=on go get -d -v -t google.golang.org/grpc/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testrace \
<<<<<<< HEAD
	updatedeps \
	updatetestdeps \
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	updatedeps \
	updatetestdeps \
=======
>>>>>>> 5ce8c7613 (update vendored files)
	vet \
	vetdeps

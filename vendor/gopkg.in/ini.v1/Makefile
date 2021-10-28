.PHONY: build test bench vet coverage

build: vet bench

test:
	go test -v -cover -race

bench:
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	go test -v -cover -test.bench=. -test.benchmem
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	go test -v -cover -race -test.bench=. -test.benchmem
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	go test -v -cover -race -test.bench=. -test.benchmem
=======
	go test -v -cover -test.bench=. -test.benchmem
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	go test -v -cover -race -test.bench=. -test.benchmem
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	go test -v -cover -race -test.bench=. -test.benchmem
=======
	go test -v -cover -test.bench=. -test.benchmem
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	go test -v -cover -race -test.bench=. -test.benchmem
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	go test -v -cover -race -test.bench=. -test.benchmem
=======
	go test -v -cover -test.bench=. -test.benchmem
>>>>>>> 4d7e5ad26 (update vendored files)

vet:
	go vet

coverage:
	go test -coverprofile=c.out && go tool cover -html=c.out && rm c.out

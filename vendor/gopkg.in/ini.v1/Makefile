.PHONY: build test bench vet coverage

build: vet bench

test:
	go test -v -cover -race

bench:
<<<<<<< HEAD
	go test -v -cover -test.bench=. -test.benchmem
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	go test -v -cover -race -test.bench=. -test.benchmem
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

vet:
	go vet

coverage:
	go test -coverprofile=c.out && go tool cover -html=c.out && rm c.out

# Build Instructions

The base tag this release is branched from is `v0.7.1`

## To build the Go code for external-dns
Does a local build of external-dns.

```make build```

## To build a docker image for external-dns
Builds external-dns and creates a docker image

```
export IMAGE=<image-prefix>/external-dns/external-dns
export VERSION=v0.7.1
make build.docker
```

## To build and push a docker image for external-dns
Builds external-dns, creates a docker image and pushes it

```
export IMAGE=<image-prefix>/external-dns/external-dns
export VERSION=v0.7.1
make build.push
```

To cleanup after a build of external-dns

```make clean```

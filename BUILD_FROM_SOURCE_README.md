# Build Instructions

The base tag this release is branched from is `v0.7.1`

To build the Go code for external-dns.

```make build```

To build a docker image for external-dns.

```
export IMAGE=<image-prefix>/external-dns/external-dns
export VERSION=v0.7.1
make build.docker
```

To push a docker image for external-dns.

```
export IMAGE=<image-prefix>/external-dns/external-dns
export VERSION=v0.7.1
make push.docker
```

To cleanup after a build of external-dns

```make clean```

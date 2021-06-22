# gssapi

[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![Build][Build-Status-Image]][Build-Status-Url] [![Coverage][Coverage-Image]][Coverage-Url] [![GoDoc][GoDoc-Image]][GoDoc-URL]

The gssapi package is a Golang wrapper around [RFC 2743](https://www.ietf.org/rfc/rfc2743.txt),
the Generic Security Services Application Programming Interface. (GSSAPI)

## Uses

We use it to authenticate clients with our authentication server. Clients talk
to a Kerberos or Active Directory Domain Controller to retrieve a Kerberos
service ticket, which we verify with a keytab on our authentication server.

When a user logs into Kerberos using `kinit`, they get a Kerberos TGT. During
Kerberos authentication, that TGT is used to retrieve a Service Ticket from the
Domain Controller. GSSAPI lets us authenticate without having to know where or
in what form the TGT is stored. Because each operating system vendor might
move that, this package wraps your system GSSAPI implementation.

What do you use it for? Let us know!

## Building

This library is `go get` compatible.  However, it also requires header files
to build against the GSSAPI C library on your platform.

Golang needs to be able to find a gcc compiler (and one which is recent enough
to support gccgo).  If the system compiler isn't gcc, then use `CC` in environ
to point the Golang build tools at your gcc.  (LLVM's clang does not work and
Golang's diagnostics if it encounters clang are to spew a lot of
apparently-unrelated errors from trying to use it anyway).

On MacOS, the default headers are too old; you can use newer headers for
building but still use the normal system libraries.

* FreeBSD: `export CC=gcc48; go install`
* MacOS: `brew install homebrew/dupes/heimdal --without-x11`
* Ubuntu: see `apt-get` in `test/docker/client/Dockerfile`

## Testing

Tests in the main `gssapi` repository can be run using the built-in `go test`.

To run an integrated test against a live Heimdal Kerberos Domain Controller,
`cd test` and bring up [Docker](https://www.docker.com/), (or
[boot2docker](http://boot2docker.io/)). Then, run `./run-heimdal.sh`. This will
run some go tests using three Docker images: a client, a service, and a domain
controller. The service will receive a generated keytab file, and the client
will point to the domain controller for authentication.

**NOTE:** to run Docker tests, your `GOROOT` environment variable MUST be set.

## TODO

See our [TODO doc](TODO.md) on stuff you can do to help. We welcome
contributions!

## Verified platforms

We've tested that we can authenticate against:

- Heimdal Kerberos
- Active Directory

We suspect we can authenticate against:

- MIT Kerberos

We definitely cannot authenticate with:

- Windows clients (because Windows uses SSPI instead of GSSAPI as the library
  interface)

[License-Url]: https://opensource.org/licenses/Apache-2.0
[License-Image]: https://img.shields.io/hexpm/l/plug.svg
[Build-Status-Url]: http://travis-ci.org/apcera/gssapi
[Build-Status-Image]: https://travis-ci.org/apcera/gssapi.svg?branch=master
[Coverage-Url]: https://coveralls.io/r/apcera/gssapi?branch=master
[Coverage-image]: https://img.shields.io/coveralls/apcera/gssapi.svg?branch=master
[ReportCard-Url]: http://goreportcard.com/report/github.com/apcera/gssapi
[ReportCard-Image]: http://goreportcard.com/badge/github.com/apcera/gssapi
[Godoc-Url]: https://godoc.org/github.com/apcera/gssapi
[Godoc-Image]: https://godoc.org/github.com/apcera/gssapi?status.svg

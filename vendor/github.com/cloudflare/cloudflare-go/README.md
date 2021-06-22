# cloudflare-go

<<<<<<< HEAD
<<<<<<< HEAD
[![GoDoc](https://img.shields.io/badge/godoc-reference-5673AF.svg?style=flat-square)](https://godoc.org/github.com/cloudflare/cloudflare-go)
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
![Test](https://github.com/cloudflare/cloudflare-go/workflows/Test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudflare/cloudflare-go?style=flat-square)](https://goreportcard.com/report/github.com/cloudflare/cloudflare-go)

> **Note**: This library is under active development as we expand it to cover
> our (expanding!) API. Consider the public API of this package a little
> unstable as we work towards a v1.0.

A Go library for interacting with
[Cloudflare's API v4](https://api.cloudflare.com/). This library allows you to:

* Manage and automate changes to your DNS records within Cloudflare
* Manage and automate changes to your zones (domains) on Cloudflare, including
  adding new zones to your account
* List and modify the status of WAF (Web Application Firewall) rules for your
  zones
* Fetch Cloudflare's IP ranges for automating your firewall whitelisting

A command-line client, [flarectl](cmd/flarectl), is also available as part of
this project.

## Features

The current feature list includes:

* [x] Cache purging
* [x] Cloudflare IPs
* [x] Custom hostnames
* [x] DNS Records
* [x] Firewall (partial)
* [ ] [Keyless SSL](https://blog.cloudflare.com/keyless-ssl-the-nitty-gritty-technical-details/)
* [x] [Load Balancing](https://blog.cloudflare.com/introducing-load-balancing-intelligent-failover-with-cloudflare/)
* [x] [Logpush Jobs](https://developers.cloudflare.com/logs/logpush/)
* [ ] Organization Administration
* [x] [Origin CA](https://blog.cloudflare.com/universal-ssl-encryption-all-the-way-to-the-origin-for-free/)
* [x] [Railgun](https://www.cloudflare.com/railgun/) administration
* [x] Rate Limiting
* [x] User Administration (partial)
* [x] Virtual DNS Management
* [x] Web Application Firewall (WAF)
* [x] Zone Lockdown and User-Agent Block rules
* [x] Zones
* [x] Workers KV
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
[![Build Status](https://img.shields.io/travis/cloudflare/cloudflare-go/master.svg?style=flat-square)](https://travis-ci.org/cloudflare/cloudflare-go)
||||||| parent of 5ce8c7613 (update vendored files)
[![Build Status](https://img.shields.io/travis/cloudflare/cloudflare-go/master.svg?style=flat-square)](https://travis-ci.org/cloudflare/cloudflare-go)
=======
![Test](https://github.com/cloudflare/cloudflare-go/workflows/Test/badge.svg)
>>>>>>> 5ce8c7613 (update vendored files)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudflare/cloudflare-go?style=flat-square)](https://goreportcard.com/report/github.com/cloudflare/cloudflare-go)

> **Note**: This library is under active development as we expand it to cover
> our (expanding!) API. Consider the public API of this package a little
> unstable as we work towards a v1.0.

A Go library for interacting with
[Cloudflare's API v4](https://api.cloudflare.com/). This library allows you to:

* Manage and automate changes to your DNS records within Cloudflare
* Manage and automate changes to your zones (domains) on Cloudflare, including
  adding new zones to your account
* List and modify the status of WAF (Web Application Firewall) rules for your
  zones
* Fetch Cloudflare's IP ranges for automating your firewall whitelisting

A command-line client, [flarectl](cmd/flarectl), is also available as part of
this project.

## Features

The current feature list includes:

* [x] Cache purging
* [x] Cloudflare IPs
* [x] Custom hostnames
* [x] DNS Records
* [x] Firewall (partial)
* [ ] [Keyless SSL](https://blog.cloudflare.com/keyless-ssl-the-nitty-gritty-technical-details/)
* [x] [Load Balancing](https://blog.cloudflare.com/introducing-load-balancing-intelligent-failover-with-cloudflare/)
* [x] [Logpush Jobs](https://developers.cloudflare.com/logs/logpush/)
* [ ] Organization Administration
* [x] [Origin CA](https://blog.cloudflare.com/universal-ssl-encryption-all-the-way-to-the-origin-for-free/)
* [x] [Railgun](https://www.cloudflare.com/railgun/) administration
* [x] Rate Limiting
* [x] User Administration (partial)
* [x] Virtual DNS Management
* [x] Web Application Firewall (WAF)
* [x] Zone Lockdown and User-Agent Block rules
* [x] Zones
<<<<<<< HEAD
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
* [x] Workers KV
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
[![Build Status](https://img.shields.io/travis/cloudflare/cloudflare-go/master.svg?style=flat-square)](https://travis-ci.org/cloudflare/cloudflare-go)
||||||| parent of 6b7ce455e (update vendored files)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5673AF.svg?style=flat-square)](https://godoc.org/github.com/cloudflare/cloudflare-go)
[![Build Status](https://img.shields.io/travis/cloudflare/cloudflare-go/master.svg?style=flat-square)](https://travis-ci.org/cloudflare/cloudflare-go)
=======
[![Go Reference](https://pkg.go.dev/badge/github.com/cloudflare/cloudflare-go.svg)](https://pkg.go.dev/github.com/cloudflare/cloudflare-go)
![Test](https://github.com/cloudflare/cloudflare-go/workflows/Test/badge.svg)
>>>>>>> 6b7ce455e (update vendored files)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudflare/cloudflare-go?style=flat-square)](https://goreportcard.com/report/github.com/cloudflare/cloudflare-go)

> **Note**: This library is under active development as we expand it to cover
> our (expanding!) API. Consider the public API of this package a little
> unstable as we work towards a v1.0.

A Go library for interacting with
[Cloudflare's API v4](https://api.cloudflare.com/). This library allows you to:

* Manage and automate changes to your DNS records within Cloudflare
* Manage and automate changes to your zones (domains) on Cloudflare, including
  adding new zones to your account
* List and modify the status of WAF (Web Application Firewall) rules for your
  zones
* Fetch Cloudflare's IP ranges for automating your firewall whitelisting

A command-line client, [flarectl](cmd/flarectl), is also available as part of
this project.

## Features

The current feature list includes:

* [x] Cache purging
* [x] Cloudflare IPs
* [x] Custom hostnames
* [x] DNS Records
* [x] Firewall (partial)
* [x] [Keyless SSL](https://blog.cloudflare.com/keyless-ssl-the-nitty-gritty-technical-details/)
* [x] [Load Balancing](https://blog.cloudflare.com/introducing-load-balancing-intelligent-failover-with-cloudflare/)
* [x] [Logpush Jobs](https://developers.cloudflare.com/logs/logpush/)
* [ ] Organization Administration
* [x] [Origin CA](https://blog.cloudflare.com/universal-ssl-encryption-all-the-way-to-the-origin-for-free/)
* [x] [Railgun](https://www.cloudflare.com/railgun/) administration
* [x] Rate Limiting
* [x] User Administration (partial)
* [x] Virtual DNS Management
* [x] Web Application Firewall (WAF)
* [x] Zone Lockdown and User-Agent Block rules
* [x] Zones
<<<<<<< HEAD
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
* [x] Workers KV
* [x] Notifications
* [x] Gateway Locations
>>>>>>> 6b7ce455e (update vendored files)

Pull Requests are welcome, but please open an issue (or comment in an existing
issue) to discuss any non-trivial changes before submitting code.

## Installation

You need a working Go environment.

```
go get github.com/cloudflare/cloudflare-go
```

## Getting Started

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	// Construct a new API object
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	ctx := context.Background()

	// Fetch user details on the account
	u, err := api.UserDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Print user details
	fmt.Println(u)

	// Fetch the zone ID
	id, err := api.ZoneIDByName("example.com") // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}

	// Fetch zone details
	zone, err := api.ZoneDetails(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	// Print zone details
	fmt.Println(zone)
}
```

Also refer to the
[API documentation](https://pkg.go.dev/github.com/cloudflare/cloudflare-go) for
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
[![GoDoc](https://img.shields.io/badge/godoc-reference-5673AF.svg?style=flat-square)](https://godoc.org/github.com/cloudflare/cloudflare-go)
[![Build Status](https://img.shields.io/travis/cloudflare/cloudflare-go/master.svg?style=flat-square)](https://travis-ci.org/cloudflare/cloudflare-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudflare/cloudflare-go?style=flat-square)](https://goreportcard.com/report/github.com/cloudflare/cloudflare-go)

> **Note**: This library is under active development as we expand it to cover
> our (expanding!) API. Consider the public API of this package a little
> unstable as we work towards a v1.0.

A Go library for interacting with
[Cloudflare's API v4](https://api.cloudflare.com/). This library allows you to:

* Manage and automate changes to your DNS records within Cloudflare
* Manage and automate changes to your zones (domains) on Cloudflare, including
  adding new zones to your account
* List and modify the status of WAF (Web Application Firewall) rules for your
  zones
* Fetch Cloudflare's IP ranges for automating your firewall whitelisting

A command-line client, [flarectl](cmd/flarectl), is also available as part of
this project.

## Features

The current feature list includes:

* [x] Cache purging
* [x] Cloudflare IPs
* [x] Custom hostnames
* [x] DNS Records
* [x] Firewall (partial)
* [ ] [Keyless SSL](https://blog.cloudflare.com/keyless-ssl-the-nitty-gritty-technical-details/)
* [x] [Load Balancing](https://blog.cloudflare.com/introducing-load-balancing-intelligent-failover-with-cloudflare/)
* [x] [Logpush Jobs](https://developers.cloudflare.com/logs/logpush/)
* [ ] Organization Administration
* [x] [Origin CA](https://blog.cloudflare.com/universal-ssl-encryption-all-the-way-to-the-origin-for-free/)
* [x] [Railgun](https://www.cloudflare.com/railgun/) administration
* [x] Rate Limiting
* [x] User Administration (partial)
* [x] Virtual DNS Management
* [x] Web Application Firewall (WAF)
* [x] Zone Lockdown and User-Agent Block rules
* [x] Zones

Pull Requests are welcome, but please open an issue (or comment in an existing
issue) to discuss any non-trivial changes before submitting code.

## Installation

You need a working Go environment.

```
go get github.com/cloudflare/cloudflare-go
```

## Getting Started

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	// Construct a new API object
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
	}

	// Fetch user details on the account
	u, err := api.UserDetails()
	if err != nil {
		log.Fatal(err)
	}
	// Print user details
	fmt.Println(u)

	// Fetch the zone ID
	id, err := api.ZoneIDByName("example.com") // Assuming example.com exists in your Cloudflare account already
	if err != nil {
		log.Fatal(err)
	}

	// Fetch zone details
	zone, err := api.ZoneDetails(id)
	if err != nil {
		log.Fatal(err)
	}
	// Print zone details
	fmt.Println(zone)
}
```

Also refer to the
[API documentation](https://godoc.org/github.com/cloudflare/cloudflare-go) for
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
how to use this package in-depth.

# License

BSD licensed. See the [LICENSE](LICENSE) file for details.

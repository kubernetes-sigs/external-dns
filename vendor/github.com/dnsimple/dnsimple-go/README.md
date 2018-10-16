# DNSimple Go Client

A Go client for the [DNSimple API v2](https://developer.dnsimple.com/v2/).

[![Build Status](https://travis-ci.org/dnsimple/dnsimple-go.svg)](https://travis-ci.org/dnsimple/dnsimple-go)
[![GoDoc](https://godoc.org/github.com/dnsimple/dnsimple-go/dnsimple?status.svg)](https://godoc.org/github.com/dnsimple/dnsimple-go/dnsimple)


## Installation

```
$ go get github.com/dnsimple/dnsimple-go/dnsimple
```


## Usage

This library is a Go client you can use to interact with the [DNSimple API v2](https://developer.dnsimple.com/v2/). Here are some examples.

```go
package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/dnsimple/dnsimple-go/dnsimple"
)

func main() {
    oauthToken := "xxxxxxx"

    // new client
    client := dnsimple.NewClient(dnsimple.NewOauthTokenCredentials(oauthToken))

    // get the current authenticated account (if you don't know who you are)
    whoamiResponse, err := client.Identity.Whoami()
    if err != nil {
        fmt.Printf("Whoami() returned error: %v\n", err)
        os.Exit(1)
    }

    fmt.Println(whoamiResponse.Data.Account)
    fmt.Println(whoamiResponse.Data.User)

    // either assign the account ID or fetch it from the response
    // if you are authenticated with an account token
    accountID := strconv.Itoa(whoamiResponse.Data.Account.ID)

    // get the list of domains
    domainsResponse, err := client.Domains.ListDomains(accountID, nil)
    if err != nil {
        fmt.Printf("Domains.ListDomains() returned error: %v\n", err)
        os.Exit(1)
    }

    // iterate over all the domains in the
    // paginated response.
    for _, domain := range domainsResponse.Data {
        fmt.Println(domain)
    }

    // List methods support a variety of options to paginate, sort and filter records.
    // Here's a few example:

    // get the list of domains filtered by name and sorted by expiration
    client.Domains.ListDomains(accountID, &dnsimple.DomainListOptions{NameLike: "com", Sort: "expiration:DESC"})
}
```

For more complete documentation, see [godoc](https://godoc.org/github.com/dnsimple/dnsimple-go/dnsimple).


## Contributing

For instructions about contributing and testing, visit the [CONTRIBUTING](CONTRIBUTING.md) file.


## License

Copyright (c) 2014-2016 Aetrion LLC. This is Free Software distributed under the MIT license.

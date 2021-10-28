[![GitHub release](https://img.shields.io/github/v/release/bodgit/tsig)](https://github.com/bodgit/tsig/releases)
[![Build Status](https://img.shields.io/github/workflow/status/bodgit/tsig/build)](https://github.com/bodgit/tsig/actions?query=workflow%3Abuild)
[![Coverage Status](https://coveralls.io/repos/github/bodgit/tsig/badge.svg?branch=master)](https://coveralls.io/github/bodgit/tsig?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/bodgit/tsig)](https://goreportcard.com/report/github.com/bodgit/tsig)
[![GoDoc](https://godoc.org/github.com/bodgit/tsig?status.svg)](https://godoc.org/github.com/bodgit/tsig)
![Go version](https://img.shields.io/badge/Go-1.17-brightgreen.svg)
![Go version](https://img.shields.io/badge/Go-1.16-brightgreen.svg)

# Additional TSIG methods

The [github.com/bodgit/tsig](https://godoc.org/github.com/bodgit/tsig) package
adds support for additional TSIG methods used in DNS queries. It is designed
to be used alongside the [github.com/miekg/dns](https://github.com/miekg/dns)
package which is used to construct and parse DNS queries and responses.

This is most useful for allowing
[RFC 3645 GSS-TSIG](https://www.ietf.org/rfc/rfc3645.txt) which is necessary
for dealing with Windows DNS servers that require 'Secure only' updates or
BIND if it has been configured to use Kerberos.

> :warning: Windows DNS servers don't accept wildcard resource names in dynamic updates.

Here is an example client, it is necessary that your Kerberos or Active
Directory environment is configured and functional:

```golang
package main

import (
        "fmt"
        "time"

        "github.com/bodgit/tsig"
        "github.com/bodgit/tsig/gss"
        "github.com/miekg/dns"
)

func main() {
        dnsClient := new(dns.Client)
        dnsClient.Net = "tcp"

        gssClient, err := gss.NewClient(dnsClient)
        if err != nil {
                panic(err)
        }
        defer gssClient.Close()

        host := "ns.example.com:53"

        // Negotiate a context with the chosen server using the
        // current user. See also gssClient.NegotiateContextWithCredentials()
        // and gssClient.NegotiateContextWithKeytab() for alternatives
        keyname, _, err := gssClient.NegotiateContext(host)
        if err != nil {
                panic(err)
        }

        dnsClient.TsigProvider = gssClient

        // Use the DNS client as normal

        msg := new(dns.Msg)
        msg.SetUpdate(dns.Fqdn("example.com"))

        insert, err := dns.NewRR("test.example.com. 300 A 192.0.2.1")
        if err != nil {
                panic(err)
        }
        msg.Insert([]dns.RR{insert})

        msg.SetTsig(keyname, tsig.GSS, 300, time.Now().Unix())

        rr, _, err := dnsClient.Exchange(msg, host)
        if err != nil {
                panic(err)
        }

        if rr.Rcode != dns.RcodeSuccess {
                fmt.Printf("DNS error: %s (%d)\n", dns.RcodeToString[rr.Rcode], rr.Rcode)
        }

        // Cleanup the context
        err = gssClient.DeleteContext(keyname)
        if err != nil {
                panic(err)
        }
}
```

If you need to deal with both regular TSIG and GSS-TSIG together then this
package also exports an HMAC TSIG implementation. To use both together set
your client up something like this:

```golang
package main

import (
        "github.com/bodgit/tsig"
        "github.com/bodgit/tsig/gss"
        "github.com/miekg/dns"
)

func main() {
        dnsClient := new(dns.Client)
        dnsClient.Net = "tcp"

        // Create HMAC TSIG provider
        hmac := tsig.HMAC{"axfr.": "so6ZGir4GPAqINNh9U5c3A=="}

        // Create GSS-TSIG provider
        gssClient, err := gss.NewClient(dnsClient)
        if err != nil {
                panic(err)
        }
        defer gssClient.Close()

        // Configure DNS client with both providers
        dnsClient.TsigProvider = tsig.MultiProvider(hmac, gssClient)

        // Use the DNS client as normal
}
```

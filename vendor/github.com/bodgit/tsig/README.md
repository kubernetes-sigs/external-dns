[![Version](https://img.shields.io/github/v/tag/bodgit/tsig)](https://github.com/bodgit/tsig/tags)
[![Build Status](https://img.shields.io/github/workflow/status/bodgit/tsig/build)](https://github.com/bodgit/tsig/actions?query=workflow%3Abuild)
[![Go Report Card](https://goreportcard.com/badge/github.com/bodgit/tsig)](https://goreportcard.com/report/github.com/bodgit/tsig)
[![GoDoc](https://godoc.org/github.com/bodgit/tsig?status.svg)](https://godoc.org/github.com/bodgit/tsig)
![Go version](https://img.shields.io/badge/Go-1.15-brightgreen.svg)
![Go version](https://img.shields.io/badge/Go-1.14-brightgreen.svg)

# Additional TSIG methods

The [github.com/bodgit/tsig](https://godoc.org/github.com/bodgit/tsig) package
adds support for additional TSIG methods used in DNS queries. It is designed
to be used alongside the [github.com/miekg/dns](https://github.com/miekg/dns)
package which is used to construct and parse DNS queries and responses.

This is most useful for allowing
[RFC 3645 GSS-TSIG](https://www.ietf.org/rfc/rfc3645.txt) which is necessary
for dealing with Windows DNS servers that require 'Secure only' updates or
BIND if it has been configured to use Kerberos.

Here is an example client, it is necessary that your Kerberos or Active
Directory environment is configured and functional:

```golang
package main

import (
        "fmt"
        "net"
        "time"

        "github.com/bodgit/tsig"
        c "github.com/bodgit/tsig/client"
        "github.com/bodgit/tsig/gss"
        "github.com/miekg/dns"
)

func main() {
        host := "ns.example.com"

        g, err := gss.New()
        if err != nil {
                panic(err)
        }
        defer g.Close()

        // Negotiate a context with the chosen server using the
        // current user. See also g.NegotiateContextWithCredentials()
        // and g.NegotiateContextWithKeytab() for alternatives
        keyname, _, err := g.NegotiateContext(host)
        if err != nil {
                panic(err)
        }

        client := c.Client{}
        client.Net = "tcp"
        client.TsigAlgorithm = map[string]*c.TsigAlgorithm{
                tsig.GSS: {
                        Generate: g.GenerateGSS,
                        Verify:   g.VerifyGSS,
                },
        }
        client.TsigSecret = map[string]string{*keyname: ""}

        // Use the DNS client as normal

        msg := new(dns.Msg)
        msg.SetUpdate(dns.Fqdn("example.com"))

        insert, err := dns.NewRR("test.example.com. 300 A 192.0.2.1")
        if err != nil {
                panic(err)
        }
        msg.Insert([]dns.RR{insert})

        msg.SetTsig(*keyname, tsig.GSS, 300, time.Now().Unix())

        rr, _, err := client.Exchange(msg, net.JoinHostPort(host, "53"))
        if err != nil {
                panic(err)
        }

        if rr.Rcode != dns.RcodeSuccess {
                fmt.Printf("DNS error: %s (%d)\n", dns.RcodeToString[rr.Rcode], rr.Rcode)
        }

        // Cleanup the context
        err = g.DeleteContext(keyname)
        if err != nil {
                panic(err)
        }
}
```

Note that it is necessary for the package to ship its own DNS client rather
than use the one provided in the github/com/miekg/dns package as it needs to
permit the additional TSIG algorithms however it behaves mostly the same and
exports the same `Exchange()` method so they can be use interchangeably in
code with a suitable interface, (see `tsig.Exchanger` for an example).

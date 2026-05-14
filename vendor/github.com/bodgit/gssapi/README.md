[![GitHub release](https://img.shields.io/github/v/release/bodgit/gssapi)](https://github.com/bodgit/gssapi/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/bodgit/gssapi/build.yml?branch=main)](https://github.com/bodgit/gssapi/actions?query=workflow%3ABuild)
[![Coverage Status](https://coveralls.io/repos/github/bodgit/gssapi/badge.svg?branch=main)](https://coveralls.io/github/bodgit/gssapi?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/bodgit/gssapi)](https://goreportcard.com/report/github.com/bodgit/gssapi)
[![GoDoc](https://godoc.org/github.com/bodgit/gssapi?status.svg)](https://godoc.org/github.com/bodgit/gssapi)
![Go version](https://img.shields.io/badge/Go-1.20-brightgreen.svg)
![Go version](https://img.shields.io/badge/Go-1.19-brightgreen.svg)

# GSSAPI wrapper for gokrb5

The [github.com/bodgit/gssapi](https://godoc.org/github.com/bodgit/gssapi)
package implements a GSSAPI-like wrapper around the
[github.com/jcmturner/gokrb5](https://github.com/jcmturner/gokrb5) package.

Sample Initiator (Client):

```golang
package main

import (
	. "github.com/bodgit/gssapi"
	"github.com/jcmturner/gokrb5/v8/gssapi"
)

func main() {
	initiator, err := NewInitiator(WithRealm("EXAMPLE.COM"), WithUsername("test"), WithKeytab[Initiator]("test.keytab"))
	if err != nil {
		panic(err)
	}

	defer initiator.Close()

	output, cont, err := initiator.Initiate("host/ssh.example.com", gssapi.ContextFlagInteg|gssapi.ContextFlagMutual, nil)
	if err != nil {
		panic(err)
	}

	// transmit output to Acceptor

	signature, err := initiator.MakeSignature(message)
	if err != nil {
		panic(err)
	}

	// transmit message and signature to Acceptor
}
```

Sample Acceptor (Server):

```golang
package main

import (
	. "github.com/bodgit/gssapi"
	"github.com/jcmturner/gokrb5/v8/gssapi"
	"github.com/jcmturner/gokrb5/v8/iana/nametype"
	"github.com/jcmturner/gokrb5/v8/types"
)

func main() {
	principal := types.NewPrincipalName(nametype.KRB_NT_SRV_HST, "host/ssh.example.com")

	acceptor, err := NewAcceptor(WithServicePrincipal(&principal))
	if err != nil {
		panic(err)
	}

	defer acceptor.Close()

	// receive input from Initiator

	output, cont, err := acceptor.Accept(input)
	if err != nil {
		panic(err)
	}

	// transmit output back to Initiator

	// receive message and signature from Initiator

	if err := acceptor.VerifySignature(message, signature); err != nil {
		panic(err)
	}
}
```

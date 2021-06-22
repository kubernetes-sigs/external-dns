/*
Package gotransip implements a client for the TransIP Rest API.
This package is a complete implementation for communicating with the TransIP RestAPI.
It covers resource calls available in the TransIP RestAPI Docs and it allows your
project(s) to connect to the TransIP RestAPI easily. Using this package you can order,
update and remove products from your TransIP account.

As of version 6.0 this package is no longer compatible with TransIP SOAP API because the
library is now organized around REST. The SOAP API library versions 5.* are now deprecated
and will no longer receive future updates.


Example

The following example uses the transip demo token in order to call the api with the test repository.
For more information about authenticating with your own credentials, see the Authentication section.

	package main

	import (
		"github.com/transip/gotransip/v6"
		"github.com/transip/gotransip/v6/test"
		"log"
	)

	func main() {
		// Create a new client with the default demo client config, using the demo token
		client, err := gotransip.NewClient(gotransip.DemoClientConfiguration)
		if err != nil {
			panic(err)
		}

		testRepo := test.Repository{Client: client}
		log.Println("Executing test call to the api server")
		if err := testRepo.Test(); err != nil {
			panic(err)
		}
		log.Println("Test successful")
	}


Authentication

If you want to tinker out with the api first without setting up your authentication,
we defined a static DemoClientConfiguration.
Which can be used to create a new client:
	client, err := gotransip.NewClient(gotransip.DemoClientConfiguration)

Create a new client using a token:
	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		Token:      "this_is_where_your_token_goes",
	})

As tokens have a limited expiry time you can also request new tokens using the private key
acquired from your transip control panel:
	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:    "accountName",
		PrivateKeyPath: "/path/to/api/private.key",
	})

We also implemented a PrivateKeyReader option, for users that want to store their key elsewhere,
not on a filesystem but on X datastore:
	file, err := os.Open("/path/to/api/private.key")
	if err != nil {
		panic(err.Error())
	}
	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:      "accountName",
		PrivateKeyReader: file,
	})


TokenCache

If you would like to keep a token between multiple client instantiations,
you can provide the client with a token cache. If the file does not exist, it creates one for you
	cache, err := authenticator.NewFileTokenCache("/tmp/path/to/gotransip_token_cache")
	if err != nil {
		panic(err.Error())
	}
	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:    "accountName",
		PrivateKeyPath: "/path/to/api/private.key",
		TokenCache:     cache
	})

As long as a provided TokenCache adheres to the following interface,
the client's authenticator is able to use it. This means you can also provide
your own token cacher: for example, one that caches to etcd
	type TokenCache interface {
		// Set will save a token by name
		Set(key string, token jwt.Token) error
		// Get returns a previously acquired token by name returned as jwt.Token
		// jwt is a subpackage in the gotransip package
		Get(key string) (jwt.Token, error)
	}

Repositories

All resource calls as can be seen on https://api.transip.nl/rest/docs.html
have been grouped in the following repositories,
these are subpackages under the gotransip package:
	availabilityzone.Repository
	colocation.Repository
	domain.Repository
	haip.Repository
	invoice.Repository
	ipaddress.Repository
	mailservice.Repository
	product.Repository
	test.Repository
	traffic.Repository
	vps.BigstorageRepository
	vps.PrivateNetworkRepository
	vps.Repository

Such a repository can be initialised with a client as follows:
	domainRepo := domain.Repository{Client: client}

Each repository has a bunch methods you can use to call get/modify/update resources in
that specific subpackage. For example, here we get a list of domains from a transip account:
	domains, err := domainRepo.GetAll()
*/
package gotransip

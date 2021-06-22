# rc0go #

[![GoDoc](https://godoc.org/github.com/nic-at/rc0go?status.svg)](https://godoc.org/github.com/nic-at/rc0go) [![Test Coverage](https://coveralls.io/repos/nic-at/rc0go/badge.svg?branch=master)](https://coveralls.io/r/nic-at/rc0go?branch=master)

rc0go is a Go client library for accessing the [RcodeZero Anycast DNS API](https://www.rcodezero.at/de/home).

## Usage ##

```go
import "github.com/nic-at/rc0go"
```

Using your API token construct a new rcode0 client, then use the various services on the client to
access different parts of the rcode0 Anycast DNS API. For example:

```go
rc0client := rcode0.NewClient("myapitoken")

// List all your zone entries which are managed by rcode0
zones, _, err := rc0client.Zones.List()

// Add a new zone to rcode0
statusResponse, err := rc0client.Zones.Create("rcodezero.at", "master", []string{})

// Get a single zone
zone, err := rc0client.Zones.Get("rcodezero.at")

// Add an "A" DNS resource record to the
rrsetCreate := []*rc0go.RRSetEdit{{
    Type:   "A",
    Name: 	"www.rcodezero.at.",
    ChangeType: rc0go.ChangeTypeADD,
    Records:    []*rc0go.Record{{
        Content: "10.10.0.1",
    }},
}}

statusResponse, err = rc0client.RRSet.Create("rcodezero.at", rrsetCreate)
```

Some code snippets are provided within the https://github.com/nic-at/rc0go/tree/master/example directory.

## Services ##

As defined in [rcode0 docs](https://my.rcodezero.at/api-doc/) the API is structured in different groups. These are:

> Zone Management <br>
> Zone Statistics <br>
> Account Statistics <br>
> Message Queue <br>
> Account Settings <br>
> Reports <br>

Each of the groups is aimed to be implemented by a Go service object (f.e. `rc0go.ZoneManagementService`) which in turn
provides the corresponding methods of the group.
DNSSEC (`rc0go.DNSSECService`), however, is defined as separate service object.

Each method contains the reference to original docs to maintain a consistent content.

## Rate Limiting ##

The API is rate limited. Additional client support will be added soon.

## Status Response ##

Some endpoints (like adding a new zone to rcode0) return a `201 Created` status code with a status response.
Status response is defined in `rc0go.StatusResponse` struct and contains only two fields - status and message.

```go
statusResponse, err := rc0client.Zones.Create("rcodezero.at", "master", []string{})
if eq := strings.Compare("ok", statusResponse.Status); eq != 0 {
    log.Println("Error: " + statusResponse.Message)
}
```
	
## Pagination ##

Some requests (like listing managed zones or rrsets) support pagination. Pagination is defined in the
`rc0go.Page` struct (with original data returned within rc0go.Page.Data field). Pagination options will be supported soon.

## Contributing ##

Contributions are most welcome!

Any changes without tests are not accepted!

1. Fork it
2. Create your feature branch (git checkout -b feature-abc)
3. Commit and sign your changes (git commit -am "Add ..." -m "Fix ..." -m "Change ... a.s.o.")
4. Push to the branch (git push origin feature-abc)
5. Create new Pull Request

## License ##

rc0go released under MIT license, refer [LICENSE](LICENSE) file.

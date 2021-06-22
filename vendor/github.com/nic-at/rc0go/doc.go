// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package rc0go provides an official client for interaction with the rcode0 Anycast DNS API in Go.

This client is highly inspired by google/go-github. The main advantage of the usage is that predefined
and API-coordinated methods are already available and the further evolution of the rcode0 API is to be transparently
aligned by the client so that the end users can focus on their own products or business logic without
always having to maintain the interaction with the rcode0 API.

Usage:

	import "github.com/nic-at/rc0go"

Using your API token construct a new rcode0 client, then use the various services on the client to
access different parts of the rcode0 Anycast DNS API. For example:

	rc0client := rcode0.NewClient("myapitoken")

	// List all your zone entries which are managed by rcode0
	zones, _, err := rc0client.Zones.List()

	// Add a new zone to rcode0
	statusResponse, err := rc0client.Zones.Create("rcodezero.at", "master", []string{})

	// Get a single zone
	zone, err := rc0client.Zones.Get("rcodezero.at")

	// Add an "A" DNS resource record to the zone
	rrsetCreate := []*rc0go.RRSetChange{{
		Type: 		"A",
		Name: 		"www.rcodezero.at.",
		ChangeType: rc0go.ChangeTypeADD,
		Records:    []*rc0go.Record{{
			Content: "10.10.0.1",
		}},
	}}

	statusResponse, err = rc0client.RRSet.Create("rcodezero.at", rrsetCreate)

Some code snippets are provided within the https://github.com/nic-at/rc0go/tree/master/example directory.

Services

As defined in rcode0 docs the API is structured in different groups. These are:

    Zone Management
    Zone Statistics
    Account Statistics
    Message Queue
    Account Settings
    Reports

Each of the groups is aimed to be implemented by a Go service object (f.e. rc0go.ZoneManagementService) which in turn
provides the corresponding methods of the group.
DNSSEC, however, is defined as separate service object.

Each method contains the reference to original docs to maintain a consistent content.

Rate Limiting

The API is rate limited. Additional client support will be added soon.

Status Response

Some endpoints (like adding a new zone to rcode0) return a 201 Created status code with a status response.
Status response is defined in rc0go.StatusResponse struct and contains only two fields - status and message.

	statusResponse, err := rc0client.Zones.Create("rcodezero.at", "master", []string{})
	if eq := strings.Compare("ok", statusResponse.Status); eq != 0 {
		log.Println("Error: " + statusResponse.Message)
	}

Pagination

Some requests (like listing managed zones or rrsets) support pagination. Pagination is defined in the
rc0go.Page struct (with original data returned within rc0go.Page.Data field). Pagination options will be supported soon.

*/
package rc0go

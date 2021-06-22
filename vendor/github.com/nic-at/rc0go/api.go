// Copyright 2019 nic.at GmbH. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rc0go

const (

	/*
	 * Zone Management
	 */

	// RC0Zone is used for GET, PUT and DELETE
	// GET:    get details of a configured zone
	// PUT:    update a zone
	// DELETE: removes a zone from the RcodeZero Anycast network
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-zone-details
	RC0Zone = "/zones/{zone}"

	// RC0Zones is used for GET and POST
	// GET: returns a list of configured zones (paginated)
	// POST: adds a new zone (master or slave) to the anycast network. (see docs for additional info)
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-zones
	RC0Zones = "/zones"

	// RC0ZoneRRSets is used for GET and PATCH
	// GET:   get the RRsets for given zone. Works for master and slave zones.
	// PATCH: adds/updates or deletes rrsets.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-rrsets
	RC0ZoneRRSets = "/zones/{zone}/rrsets"

	// RC0ZoneTransfer is used for POST
	// POST: queues a zone transfer dnssecRequest for the given zone.
	// Zone will be transfered regardless of the serial on the RcodeZero Anycast Network.
	// A zone with a greater serial on the Rcode0 network will be overwritten by the newly transferred version.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-zone-transfer
	RC0ZoneTransfer = "/zones/{zone}/retrieve"

	// RC0ZoneDNSSecSign is used for POST
	// POST: starts DNSSEC signing of a zone.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-dnssec-sign-zone
	RC0ZoneDNSSecSign = "/zones/{zone}/sign"

	// RC0ZoneDNSSecUnsign is used for POST
	// POST: stops DNSSEC signing of a zone, reverting the zone to unsigned.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-dnssec-unsign-zone
	RC0ZoneDNSSecUnsign = "/zones/{zone}/unsign"

	// RC0ZoneDNSSecKeyRollover is used for POST
	// POST: starts a DNSSEC key rollover
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-dnssec-key-rollover
	RC0ZoneDNSSecKeyRollover = "/zones/{zone}/keyrollover"

	// RC0ZoneDNSSecDSUpdate is used for POST
	// POST: acknowledges a DS update
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-dnssec-acknowledge-ds-update
	RC0ZoneDNSSecDSUpdate = "/zones/{zone}/dsupdate"

	// RC0ZoneDNSSecDSSEEN is used for POST (available on test system only)
	// POST: Simulates that the DS records of all KSKs of a certain domain were seen in the parent zone.
	// This allows to test key rollovers even if the DS of the currently active KSK was not seen in the parent zone.
	// A DSSEEN event will be pushed ot the message queue.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-simulate-dnssec-event-dsseen
	RC0ZoneDNSSecDSSEEN = "/zones/{zone}/simulate/dsseen"

	// RC0ZoneDNSSecDSREMOVED is used for POST (available on test system only)
	// POST: simulates that the DS records of all KSKs of a certain domain were removed from the parent zone.
	// This allows to subsequently “unsign” a domain.
	// A DSREMOVED event will be pushed to the message queue.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-management-simulate-dnssec-event-dsremoved
	RC0ZoneDNSSecDSREMOVED = "/zones/{zone}/simulate/dsremoved"

	/*
	 * Zone Statistics
	 */

	// RC0ZoneStatsQueries is used for GET
	// GET: Get the total number of queries and the number of queries answered with NXDOMAIN for the given zone for the last 180 days (max.)
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-statistics-queries
	RC0ZoneStatsQueries = "/zones/{zone}/stats/queries"

	// RC0ZoneStatsMagnitude is used for GET
	// GET: Get the DNS magnitude for a given zone for the last 180 days.
	// The DNS magnitude reflects the popularity of a domain between 0 (low) and 10 (high).
	// The figure is based on the number of unique resolvers seen during a day.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-statistics-dns-magnitude
	RC0ZoneStatsMagnitude = "/zones/{zone}/stats/magnitude"

	// RC0ZoneStatsQNames is used for GET
	// GET: Returns yesterdays top 10 QNAMEs with QTYPE for the given domain.
	// Returns an empty array if no queries have been received for the domain
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-statistics-qnames
	RC0ZoneStatsQNames = "/zones/{zone}/stats/qnames"

	// RC0ZoneStatsNXDomains is used for GET
	// GET: Returns yesterdays top 10 labels and QTYPE answered with NXDOMAIN.
	// Returns an empty array if no (NX-)queries have been received.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-zone-statistics-nxdomains
	RC0ZoneStatsNXDomains = "/zones/{zone}/stats/nxdomains"

	/*
	 * Account Statistics
	 */

	// RC0AccStatsTopZones is used for GET
	// GET: Return the Top 1000 zones from your account with the highest number of queries in the given past period.
	// Returns an empty array if no data is available
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-account-statistics-top-zones
	RC0AccStatsTopZones = "/stats/topzones" //{?days}

	// RC0AccStatsTopQNames is used for GET
	// GET: Returns the Top 1000 QNAMEs for zones in your account with the highest number of queries in the given past period.
	// Returns an empty array if no data is available
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-account-statistics-top-qnames
	RC0AccStatsTopQNames = "/stats/topqnames" //{?days}

	// RC0AccStatsTopNXDomains is used for GET
	// GET: Returns the Top 1000 QNAMEs with QTYPE answered with NXDOMAIN for zones in your account with the highest number of queries in the given past period.
	// Returns an empty array if no data is available
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-account-statistics-top-nxdomains
	RC0AccStatsTopNXDomains = "/stats/topnxdomains" //{?days}

	// RC0AccStatsTopDNSMagnitude is used for GET
	// GET: Returns the Top 1000 zone in your account with the highest dns magnitude in the given past period.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-account-statistics-top-dns-magnitude
	RC0AccStatsTopDNSMagnitude = "/stats/topmagnitude" //{?days}

	// RC0AccStatsQueries is used for GET
	// GET: Get the total number of queries and the number of queries answered with NXDOMAIN for all zones in your account for the given past period.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-account-statistics-queries
	RC0AccStatsQueries = "/stats/queries" //{?days}

	// RC0AccStatsCountries is used for GET
	// GET: Return the number of Queries grouped by the originating country/subregion and region for the given past period.
	// Returns an empty array if no data is available.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-account-statistics-countries
	RC0AccStatsCountries = "/stats/countries" //{?days}

	/*
	 * Account Settings
	 */

	// RC0AccSettings is used for GET
	// GET: get global account settings. Value will be empty if an individual setting is not configured.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-account-settings-settings
	RC0AccSettings = "/settings"

	// RC0AccSecondaries is used for PUT and DELETE
	// PUT: configures the account setting “secondaries”.
	// Those secondaries will receive notifies and may transfer out all zones under the management of the account.
	// DELETE: removes the configured secondaries
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-account-settings-set-secondaries
	RC0AccSecondaries = "/settings/secondaries"

	// RC0AccTsigout is used for PUT and DELETE
	// PUT: configures the TSIG key used for outbound zone transfers.
	// DELETE: removes the configured TSIG key
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-account-settings-settings-tsigout
	RC0AccTsigout = "/settings/tsigout"

	/*
	 * Messages
	 */

	// RC0Messages is used for GET
	// GET: retrieves the oldest unacknowledged message from the message queue.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-message-queue-poll-message
	RC0Messages = "/messages"

	// RC0AckMessage is used for DELETE
	// DELETE: acknowlegdes (and deletes) the message with the given id
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-message-queue-ack-message
	RC0AckMessage = "/messages/{id}"

	/*
	 * Reports
	 */

	// RC0ReportsProblematiczones is used for GET
	// GET: get global account settings. Value will be empty if an individual setting is not configured.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-reports-reports-problematiczones
	RC0ReportsProblematiczones = "/reports/problematiczones"

	// RC0ReportsNXDomains is used for GET
	// GET: get all QNAMEs and QTYPE for your account which have been answered with NXDOMAIN for yesterday or today as CSV.
	// The report is updated once every hour.
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-reports-reports-nxdomains
	RC0ReportsNXDomains = "/reports/nxdomains" //{?day}

	// RC0ReportsAccounting is used for GET
	// GET: get the accounting report per day for the given month as CSV
	// Parameter: month Values: ‘YYYY-MM’
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-reports-reports-accounting
	RC0ReportsAccounting = "/reports/accounting" //{?month}

	// RC0ReportsQueryrates is used for GET
	// GET: get the number of queries per domain and day for the given month as CSV
	// Parameter: month Values: ‘YYYY-MM’
	//
	// rcode0 API docs: https://my.rcodezero.at/api-doc/#api-reports-reports-queryrates
	RC0ReportsQueryrates = "/reports/nxdomains" //{?month}
)

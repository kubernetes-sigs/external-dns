// This file provides a mapping of Cloudflare v4.5.1 API changes for reference during the refactor.
// It is not part of the build, but helps track what needs to be updated.
// Remove this file after the refactor is complete.

// Key changes in v4.5.1:
// - All major types (DNSRecord, CustomHostname, etc.) are now under the cloudflare-go/v4 package.
// - Many methods now require a ResourceContainer (zone identifier) instead of just a string zoneID.
// - ListDNSRecords, CreateDNSRecord, UpdateDNSRecord, DeleteDNSRecord, etc. all use ResourceContainer.
// - ListZonesContext returns ZonesResponse, not []Zone.
// - CustomHostnames API is now paginated and returns a slice and ResultInfo.
// - Error handling uses *cloudflare.Error for API errors.
// - Proxied is now *bool in DNSRecord and params.
// - Use cloudflare.ZoneIdentifier(zoneID) to get a ResourceContainer for a zone.
// - ResultInfo is used for pagination in most list APIs.
// - Some params structs have changed field names/types.

// See: https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v4

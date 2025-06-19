// This file is for reference only. It lists the DNS record types and methods exported by github.com/cloudflare/cloudflare-go/v4 as of v4.5.1.
// Use this to update all type references in the Cloudflare provider code.

// DNS record management types (as of v4.5.1):
// - DNSRecord (type)
// - CreateDNSRecordParams (type)
// - UpdateDNSRecordParams (type)
// - ListDNSRecordsParams (type)
// - ResultInfo (type)
// - ResourceContainer (type)
// - ZoneIdentifier (func)
// - CustomHostname (type)
// - CustomHostnameResponse (type)
// - CustomHostnameSSL (type)
// - CustomHostnameSSLSettings (type)
// - RegionalHostname (type)
// - CreateDataLocalizationRegionalHostnameParams (type)
// - UpdateDataLocalizationRegionalHostnameParams (type)
// - Error (type)
// - API (struct)

// All of these are still exported from the root package: github.com/cloudflare/cloudflare-go/v4
// If you see build errors, ensure the import path is exactly:
// import "github.com/cloudflare/cloudflare-go/v4"
// and reference types as cloudflare.DNSRecord, cloudflare.CreateDNSRecordParams, etc.

// If you still see build errors, run `go mod tidy` and ensure your go.mod and go.sum are up to date.

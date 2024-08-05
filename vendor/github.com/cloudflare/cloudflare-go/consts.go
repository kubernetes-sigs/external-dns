package cloudflare

// RouteRoot represents the name of the route namespace.
type RouteRoot string

const (
	defaultScheme   = "https"
	defaultHostname = "api.cloudflare.com"
	defaultBasePath = "/client/v4"
	userAgent       = "cloudflare-go"

	// AccountRouteRoot is the accounts route namespace.
	AccountRouteRoot RouteRoot = "accounts"

	// ZoneRouteRoot is the zones route namespace.
	ZoneRouteRoot RouteRoot = "zones"

	originCARootCertEccURL = "https://developers.cloudflare.com/ssl/static/origin_ca_ecc_root.pem"
	originCARootCertRsaURL = "https://developers.cloudflare.com/ssl/static/origin_ca_rsa_root.pem"

	// Used for testing.
	testAccountID    = "01a7362d577a6c3019a474fd6f485823"
	testZoneID       = "d56084adb405e0b7e32c52321bf07be6"
<<<<<<< HEAD
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	testUserID       = "a81be4e9b20632860d20a64c054c4150"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	testCertPackUUID = "a77f8bd7-3b47-46b4-a6f1-75cf98109948"
	testTunnelID     = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
)

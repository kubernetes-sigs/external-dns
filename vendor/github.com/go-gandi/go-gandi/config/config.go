package config

import "time"

// Config manages common config for all Gandi API types
type Config struct {
	// APIKey has been deprecated by Gandi in favor of Personal Access Token. Please refer to the Gandi authentication documentation: https://api.gandi.net/docs/authentication/ for details.
	APIKey string
	// PersonalAccessToken is a configured token from the Gandi Admin application. Please refer to the Gandi authentication documentation for its creation: https://api.gandi.net/docs/authentication/
	PersonalAccessToken string
	// SharingID is the Organization ID, available from the Organization API
	SharingID string
	// Debug enables verbose debugging of HTTP calls
	Debug bool
	// DryRun prevents the API from making changes. Only certain API calls support it.
	DryRun bool
	// APIURL is the Gandi API URL. By default, it fallbacks to
	// https://api.gandi.net.
	APIURL string
	// Timeout is the timeout for requests against the Gandi API
	Timeout time.Duration
}

const (
	// APIURL is the default Config.APIURL value
	APIURL = "https://api.gandi.net"
	// SandboxAPIURL is the URL of the Gandi Sandbox API
	SandboxAPIURL = "https://api.sandbox.gandi.net"
	// Timeout is the default timeout of 5 seconds
	Timeout = 5 * time.Second
)

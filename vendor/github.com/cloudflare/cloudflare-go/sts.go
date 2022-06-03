package cloudflare

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

var (
	ErrSTSFailure               = errors.New("failed to fetch security token")
	ErrSTSHTTPFailure           = errors.New("failed making securtiy token issuer call")
	ErrSTSHTTPResponseError     = errors.New("security token request returned a failure")
	ErrSTSMissingServiceSecret  = errors.New("service secret missing but is required")
	ErrSTSMissingServiceTag     = errors.New("service tag missing but is required")
	ErrSTSMissingIssuerHostname = errors.New("issuer hostname missing but is required")
	ErrSTSMissingServicePath    = errors.New("issuer path missing but is required")
)

// IssuerConfiguration allows the configuration of the issuance provider.
type IssuerConfiguration struct {
	Hostname string
	Path     string
}

// SecurityTokenConfiguration holds the configuration for requesting a security
// token from the service.
type SecurityTokenConfiguration struct {
	Issuer     *IssuerConfiguration
	ServiceTag string
	Secret     string
}

type securityToken struct {
	Token string `json:"json_web_token"`
}

type securityTokenResponse struct {
	Result securityToken `json:"result"`
	Response
}

// fetchSTSCredentials provides a way to authenticate with the security token
// service and issue a usable token for the system.
func fetchSTSCredentials(stsConfig *SecurityTokenConfiguration) (string, error) {
	if stsConfig.Secret == "" {
		return "", ErrSTSMissingServiceSecret
	}

	if stsConfig.ServiceTag == "" {
		return "", ErrSTSMissingServiceTag
	}

	if stsConfig.Issuer.Hostname == "" {
		return "", ErrSTSMissingIssuerHostname
	}

	if stsConfig.Issuer.Path == "" {
		return "", ErrSTSMissingServicePath
	}

	retryableClient := retryablehttp.NewClient()
	retryableClient.RetryMax = 3
	stsClient := retryableClient.StandardClient()

	uri := fmt.Sprintf("https://%s%s", stsConfig.Issuer.Hostname, stsConfig.Issuer.Path)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return "", fmt.Errorf("HTTP request creation failed: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+stsConfig.ServiceTag+stsConfig.Secret)

	resp, err := stsClient.Do(req)
	if err != nil {
		return "", ErrSTSHTTPFailure
	}

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	resp.Body.Close()

	var stsTokenResponse *securityTokenResponse
	err = json.Unmarshal(respBody, &stsTokenResponse)
	if err != nil {
		return "", fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	if !stsTokenResponse.Success {
		return "", ErrSTSHTTPResponseError
	}

	return stsTokenResponse.Result.Token, nil
}

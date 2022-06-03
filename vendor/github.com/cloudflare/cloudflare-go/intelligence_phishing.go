package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// PhishingScan represent information about a phishing scan.
type PhishingScan struct {
	URL        string  `json:"url"`
	Phishing   bool    `json:"phishing"`
	Verified   bool    `json:"verified"`
	Score      float64 `json:"score"`
	Classifier string  `json:"classifier"`
}

// PhishingScanParameters represent parameters for a phishing scan request.
type PhishingScanParameters struct {
	AccountID string `url:"-"`
	URL       string `url:"url,omitempty"`
	Skip      bool   `url:"skip,omitempty"`
}

// PhishingScanResponse represent an API response for a phishing scan.
type PhishingScanResponse struct {
	Response
	Result PhishingScan `json:"result,omitempty"`
}

// IntelligencePhishingScan scans a URL for suspected phishing
//
// API Reference: https://api.cloudflare.com/#phishing-url-scanner-scan-suspicious-url
func (api *API) IntelligencePhishingScan(ctx context.Context, params PhishingScanParameters) (PhishingScan, error) {
	if params.AccountID == "" {
		return PhishingScan{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/intel-phishing/predict", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PhishingScan{}, err
	}

	var phishingScanResponse PhishingScanResponse
	if err := json.Unmarshal(res, &phishingScanResponse); err != nil {
		return PhishingScan{}, err
	}
	return phishingScanResponse.Result, nil
}

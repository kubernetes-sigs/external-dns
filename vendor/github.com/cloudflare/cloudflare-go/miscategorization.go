package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrMissingIP is for when ipv4 or ipv6 indicator was given but ip is missing.
	ErrMissingIP = errors.New("ip is required when using 'ipv4' or 'ipv6' indicator type and is missing")
	// ErrMissingURL is for when url or domain indicator was given but url is missing.
	ErrMissingURL = errors.New("url is required when using 'domain' or 'url' indicator type and is missing")
)

// MisCategorizationParameters represents the parameters for a miscategorization request.
type MisCategorizationParameters struct {
	AccountID       string
	IndicatorType   string `json:"indicator_type,omitempty"`
	IP              string `json:"ip,omitempty"`
	URL             string `json:"url,omitempty"`
	ContentAdds     []int  `json:"content_adds,omitempty"`
	ContentRemoves  []int  `json:"content_removes,omitempty"`
	SecurityAdds    []int  `json:"security_adds,omitempty"`
	SecurityRemoves []int  `json:"security_removes,omitempty"`
}

// CreateMiscategorization creates a miscatergorization.
//
// API Reference: https://api.cloudflare.com/#miscategorization-create-miscategorization
func (api *API) CreateMiscategorization(ctx context.Context, params MisCategorizationParameters) error {
	if params.AccountID == "" {
		return ErrMissingAccountID
	}
	if (params.IndicatorType == "ipv6" || params.IndicatorType == "ipv4") && params.IP == "" {
		return ErrMissingIP
	}
	if (params.IndicatorType == "domain" || params.IndicatorType == "url") && params.URL == "" {
		return ErrMissingURL
	}

	uri := fmt.Sprintf("/accounts/%s/intel/miscategorization", params.AccountID)
	_, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return err
	}

	return nil
}

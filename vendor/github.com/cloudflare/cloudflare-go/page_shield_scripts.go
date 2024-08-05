package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// PageShieldScript represents a Page Shield script.
type PageShieldScript struct {
	AddedAt                 string   `json:"added_at"`
	DomainReportedMalicious *bool    `json:"domain_reported_malicious,omitempty"`
	FetchedAt               string   `json:"fetched_at"`
	FirstPageURL            string   `json:"first_page_url"`
	FirstSeenAt             string   `json:"first_seen_at"`
	Hash                    string   `json:"hash"`
	Host                    string   `json:"host"`
	ID                      string   `json:"id"`
	JSIntegrityScore        int      `json:"js_integrity_score"`
	LastSeenAt              string   `json:"last_seen_at"`
	PageURLs                []string `json:"page_urls"`
	URL                     string   `json:"url"`
	URLContainsCdnCgiPath   *bool    `json:"url_contains_cdn_cgi_path,omitempty"`
}

// ListPageShieldScriptsParams represents a PageShield Script request parameters.
//
// API reference: https://developers.cloudflare.com/api/operations/page-shield-list-page-shield-scripts#Query-Parameters
type ListPageShieldScriptsParams struct {
	Direction           string `url:"direction"`
	ExcludeCdnCgi       *bool  `url:"exclude_cdn_cgi,omitempty"`
	ExcludeDuplicates   *bool  `url:"exclude_duplicates,omitempty"`
	ExcludeUrls         string `url:"exclude_urls"`
	Export              string `url:"export"`
	Hosts               string `url:"hosts"`
	OrderBy             string `url:"order_by"`
	Page                string `url:"page"`
	PageURL             string `url:"page_url"`
	PerPage             int    `url:"per_page"`
	PrioritizeMalicious *bool  `url:"prioritize_malicious,omitempty"`
	Status              string `url:"status"`
	URLs                string `url:"urls"`
}

// PageShieldScriptsResponse represents the response from the PageShield Script API.
type PageShieldScriptsResponse struct {
	Results []PageShieldScript `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// PageShieldScriptResponse represents the response from the PageShield Script API.
type PageShieldScriptResponse struct {
	Result   PageShieldScript          `json:"result"`
	Versions []PageShieldScriptVersion `json:"versions"`
}

// PageShieldScriptVersion represents a Page Shield script version.
type PageShieldScriptVersion struct {
	FetchedAt        string `json:"fetched_at"`
	Hash             string `json:"hash"`
	JSIntegrityScore int    `json:"js_integrity_score"`
}

// ListPageShieldScripts returns a list of PageShield Scripts.
//
// API reference: https://developers.cloudflare.com/api/operations/page-shield-list-page-shield-scripts
func (api *API) ListPageShieldScripts(ctx context.Context, rc *ResourceContainer, params ListPageShieldScriptsParams) ([]PageShieldScript, ResultInfo, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/scripts", rc.Identifier)

	uri := buildURI(path, params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, ResultInfo{}, err
	}

	var psResponse PageShieldScriptsResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return psResponse.Results, psResponse.ResultInfo, nil
}

// GetPageShieldScript returns a PageShield Script.
//
// API reference: https://developers.cloudflare.com/api/operations/page-shield-get-a-page-shield-script
func (api *API) GetPageShieldScript(ctx context.Context, rc *ResourceContainer, scriptID string) (*PageShieldScript, []PageShieldScriptVersion, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/scripts/%s", rc.Identifier, scriptID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var psResponse PageShieldScriptResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse.Result, psResponse.Versions, nil
}

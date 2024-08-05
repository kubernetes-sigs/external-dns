package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var (
	ErrMissingProfileID = errors.New("missing required profile ID")
)

// DLPPattern represents a DLP Pattern that matches an entry.
type DLPPattern struct {
	Regex      string `json:"regex,omitempty"`
	Validation string `json:"validation,omitempty"`
}

// DLPEntry represents a DLP Entry, which can be matched in HTTP bodies or files.
type DLPEntry struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ProfileID string `json:"profile_id,omitempty"`
	Enabled   *bool  `json:"enabled,omitempty"`
	Type      string `json:"type,omitempty"`

	// The following fields are only present for custom entries.

	Pattern   *DLPPattern `json:"pattern,omitempty"`
	CreatedAt *time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty"`
}

// Content types to exclude from context analysis and return all matches.
type DLPContextAwarenessSkip struct {
	// Return all matches, regardless of context analysis result, if the data is a file.
	Files *bool `json:"files,omitempty"`
}

// Scan the context of predefined entries to only return matches surrounded by keywords.
type DLPContextAwareness struct {
	Enabled *bool                   `json:"enabled,omitempty"`
	Skip    DLPContextAwarenessSkip `json:"skip"`
}

// DLPProfile represents a DLP Profile, which contains a set
// of entries.
type DLPProfile struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Type              string `json:"type,omitempty"`
	Description       string `json:"description,omitempty"`
	AllowedMatchCount int    `json:"allowed_match_count"`
	OCREnabled        *bool  `json:"ocr_enabled,omitempty"`

	ContextAwareness *DLPContextAwareness `json:"context_awareness,omitempty"`

	// The following fields are omitted for predefined DLP
	// profiles.
	Entries   []DLPEntry `json:"entries,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// DLPProfilesCreateRequest represents a request to create a
// set of profiles.
type DLPProfilesCreateRequest struct {
	Profiles []DLPProfile `json:"profiles"`
}

// DLPProfileListResponse represents the response from the list
// dlp profiles endpoint.
type DLPProfileListResponse struct {
	Result []DLPProfile `json:"result"`
	Response
}

// DLPProfileResponse is the API response, containing a single
// access application.
type DLPProfileResponse struct {
	Success  bool       `json:"success"`
	Errors   []string   `json:"errors"`
	Messages []string   `json:"messages"`
	Result   DLPProfile `json:"result"`
}

type ListDLPProfilesParams struct{}

type CreateDLPProfilesParams struct {
	Profiles []DLPProfile `json:"profiles"`
	Type     string
}

type UpdateDLPProfileParams struct {
	ProfileID string
	Profile   DLPProfile
	Type      string
}

// ListDLPProfiles returns all DLP profiles within an account.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-list-all-profiles
func (api *API) ListDLPProfiles(ctx context.Context, rc *ResourceContainer, params ListDLPProfilesParams) ([]DLPProfile, error) {
	if rc.Identifier == "" {
		return []DLPProfile{}, ErrMissingResourceIdentifier
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles", rc.Level, rc.Identifier), nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []DLPProfile{}, err
	}

	var dlpProfilesListResponse DLPProfileListResponse
	err = json.Unmarshal(res, &dlpProfilesListResponse)
	if err != nil {
		return []DLPProfile{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpProfilesListResponse.Result, nil
}

// GetDLPProfile returns a single DLP profile (custom or predefined) based on
// the profile ID.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-get-dlp-profile
func (api *API) GetDLPProfile(ctx context.Context, rc *ResourceContainer, profileID string) (DLPProfile, error) {
	if rc.Identifier == "" {
		return DLPProfile{}, ErrMissingResourceIdentifier
	}

	if profileID == "" {
		return DLPProfile{}, ErrMissingProfileID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles/%s", rc.Level, rc.Identifier, profileID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DLPProfile{}, err
	}

	var dlpProfileResponse DLPProfileResponse
	err = json.Unmarshal(res, &dlpProfileResponse)
	if err != nil {
		return DLPProfile{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpProfileResponse.Result, nil
}

// CreateDLPProfiles creates a set of DLP Profile.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-create-custom-profiles
func (api *API) CreateDLPProfiles(ctx context.Context, rc *ResourceContainer, params CreateDLPProfilesParams) ([]DLPProfile, error) {
	if rc.Identifier == "" {
		return []DLPProfile{}, ErrMissingResourceIdentifier
	}

	if params.Type == "" || params.Type != "custom" {
		return []DLPProfile{}, fmt.Errorf("unsupported DLP profile type: %q", params.Type)
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles/%s", rc.Level, rc.Identifier, params.Type), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return []DLPProfile{}, err
	}

	var dLPCustomProfilesResponse DLPProfileListResponse
	err = json.Unmarshal(res, &dLPCustomProfilesResponse)
	if err != nil {
		return []DLPProfile{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dLPCustomProfilesResponse.Result, nil
}

// DeleteDLPProfile deletes a DLP profile. Only custom profiles can be deleted.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-delete-custom-profile
func (api *API) DeleteDLPProfile(ctx context.Context, rc *ResourceContainer, profileID string) error {
	if rc.Identifier == "" {
		return ErrMissingResourceIdentifier
	}

	if profileID == "" {
		return ErrMissingProfileID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles/custom/%s", rc.Level, rc.Identifier, profileID), nil)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	return err
}

// UpdateDLPProfile updates a DLP profile.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-update-custom-profile
// API reference: https://api.cloudflare.com/#dlp-profiles-update-predefined-profile
func (api *API) UpdateDLPProfile(ctx context.Context, rc *ResourceContainer, params UpdateDLPProfileParams) (DLPProfile, error) {
	if rc.Identifier == "" {
		return DLPProfile{}, ErrMissingResourceIdentifier
	}

	if params.Type == "" {
		params.Type = "custom"
	}

	if params.ProfileID == "" {
		return DLPProfile{}, ErrMissingProfileID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles/%s/%s", rc.Level, rc.Identifier, params.Type, params.ProfileID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.Profile)
	if err != nil {
		return DLPProfile{}, err
	}

	var dlpProfileResponse DLPProfileResponse
	err = json.Unmarshal(res, &dlpProfileResponse)
	if err != nil {
		return DLPProfile{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpProfileResponse.Result, nil
}

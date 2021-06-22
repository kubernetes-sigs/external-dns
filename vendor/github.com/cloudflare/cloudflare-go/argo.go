package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var validSettingValues = []string{"on", "off"}

// ArgoFeatureSetting is the structure of the API object for the
// argo smart routing and tiered caching settings.
type ArgoFeatureSetting struct {
	Editable   bool      `json:"editable,omitempty"`
	ID         string    `json:"id,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
	Value      string    `json:"value"`
}

// ArgoDetailsResponse is the API response for the argo smart routing
// and tiered caching response.
type ArgoDetailsResponse struct {
	Result ArgoFeatureSetting `json:"result"`
	Response
}

// ArgoSmartRouting returns the current settings for smart routing.
//
// API reference: https://api.cloudflare.com/#argo-smart-routing-get-argo-smart-routing-setting
func (api *API) ArgoSmartRouting(ctx context.Context, zoneID string) (ArgoFeatureSetting, error) {
	uri := fmt.Sprintf("/zones/%s/argo/smart_routing", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ArgoFeatureSetting{}, err
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return argoDetailsResponse.Result, nil
}

// UpdateArgoSmartRouting updates the setting for smart routing.
//
// API reference: https://api.cloudflare.com/#argo-smart-routing-patch-argo-smart-routing-setting
func (api *API) UpdateArgoSmartRouting(ctx context.Context, zoneID, settingValue string) (ArgoFeatureSetting, error) {
	if !contains(validSettingValues, settingValue) {
		return ArgoFeatureSetting{}, fmt.Errorf("invalid setting value '%s'. must be 'on' or 'off'", settingValue)
	}

	uri := fmt.Sprintf("/zones/%s/argo/smart_routing", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, ArgoFeatureSetting{Value: settingValue})
	if err != nil {
		return ArgoFeatureSetting{}, err
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return argoDetailsResponse.Result, nil
}

// ArgoTieredCaching returns the current settings for tiered caching.
//
// API reference: TBA.
func (api *API) ArgoTieredCaching(ctx context.Context, zoneID string) (ArgoFeatureSetting, error) {
	uri := fmt.Sprintf("/zones/%s/argo/tiered_caching", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ArgoFeatureSetting{}, err
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return argoDetailsResponse.Result, nil
}

// UpdateArgoTieredCaching updates the setting for tiered caching.
//
// API reference: TBA.
func (api *API) UpdateArgoTieredCaching(ctx context.Context, zoneID, settingValue string) (ArgoFeatureSetting, error) {
	if !contains(validSettingValues, settingValue) {
		return ArgoFeatureSetting{}, fmt.Errorf("invalid setting value '%s'. must be 'on' or 'off'", settingValue)
	}

	uri := fmt.Sprintf("/zones/%s/argo/tiered_caching", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, ArgoFeatureSetting{Value: settingValue})
	if err != nil {
		return ArgoFeatureSetting{}, err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var validSettingValues = []string{"on", "off"}

// ArgoFeatureSetting is the structure of the API object for the
// argo smart routing and tiered caching settings.
type ArgoFeatureSetting struct {
	Editable   bool      `json:"editable,omitempty"`
	ID         string    `json:"id,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
	Value      string    `json:"value"`
}

// ArgoDetailsResponse is the API response for the argo smart routing
// and tiered caching response.
type ArgoDetailsResponse struct {
	Result ArgoFeatureSetting `json:"result"`
	Response
}

// ArgoSmartRouting returns the current settings for smart routing.
//
// API reference: https://api.cloudflare.com/#argo-smart-routing-get-argo-smart-routing-setting
func (api *API) ArgoSmartRouting(ctx context.Context, zoneID string) (ArgoFeatureSetting, error) {
	uri := fmt.Sprintf("/zones/%s/argo/smart_routing", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ArgoFeatureSetting{}, err
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errUnmarshalError)
	}
	return argoDetailsResponse.Result, nil
}

// UpdateArgoSmartRouting updates the setting for smart routing.
//
// API reference: https://api.cloudflare.com/#argo-smart-routing-patch-argo-smart-routing-setting
func (api *API) UpdateArgoSmartRouting(ctx context.Context, zoneID, settingValue string) (ArgoFeatureSetting, error) {
	if !contains(validSettingValues, settingValue) {
		return ArgoFeatureSetting{}, errors.New(fmt.Sprintf("invalid setting value '%s'. must be 'on' or 'off'", settingValue))
	}

	uri := fmt.Sprintf("/zones/%s/argo/smart_routing", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, ArgoFeatureSetting{Value: settingValue})
	if err != nil {
		return ArgoFeatureSetting{}, err
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errUnmarshalError)
	}
	return argoDetailsResponse.Result, nil
}

// ArgoTieredCaching returns the current settings for tiered caching.
//
// API reference: TBA
func (api *API) ArgoTieredCaching(ctx context.Context, zoneID string) (ArgoFeatureSetting, error) {
	uri := fmt.Sprintf("/zones/%s/argo/tiered_caching", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ArgoFeatureSetting{}, err
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errUnmarshalError)
	}
	return argoDetailsResponse.Result, nil
}

// UpdateArgoTieredCaching updates the setting for tiered caching.
//
// API reference: TBA
func (api *API) UpdateArgoTieredCaching(ctx context.Context, zoneID, settingValue string) (ArgoFeatureSetting, error) {
	if !contains(validSettingValues, settingValue) {
		return ArgoFeatureSetting{}, errors.New(fmt.Sprintf("invalid setting value '%s'. must be 'on' or 'off'", settingValue))
	}

	uri := fmt.Sprintf("/zones/%s/argo/tiered_caching", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, ArgoFeatureSetting{Value: settingValue})
	if err != nil {
<<<<<<< HEAD
		return ArgoFeatureSetting{}, errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return ArgoFeatureSetting{}, errors.Wrap(err, errMakeRequestError)
=======
		return ArgoFeatureSetting{}, err
>>>>>>> 4d7e5ad26 (update vendored files)
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

var validSettingValues = []string{"on", "off"}

// ArgoFeatureSetting is the structure of the API object for the
// argo smart routing and tiered caching settings.
type ArgoFeatureSetting struct {
	Editable   bool      `json:"editable,omitempty"`
	ID         string    `json:"id,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
	Value      string    `json:"value"`
}

// ArgoDetailsResponse is the API response for the argo smart routing
// and tiered caching response.
type ArgoDetailsResponse struct {
	Result ArgoFeatureSetting `json:"result"`
	Response
}

// ArgoSmartRouting returns the current settings for smart routing.
//
// API reference: https://api.cloudflare.com/#argo-smart-routing-get-argo-smart-routing-setting
func (api *API) ArgoSmartRouting(zoneID string) (ArgoFeatureSetting, error) {
	uri := "/zones/" + zoneID + "/argo/smart_routing"

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errMakeRequestError)
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errUnmarshalError)
	}
	return argoDetailsResponse.Result, nil
}

// UpdateArgoSmartRouting updates the setting for smart routing.
//
// API reference: https://api.cloudflare.com/#argo-smart-routing-patch-argo-smart-routing-setting
func (api *API) UpdateArgoSmartRouting(zoneID, settingValue string) (ArgoFeatureSetting, error) {
	if !contains(validSettingValues, settingValue) {
		return ArgoFeatureSetting{}, errors.New(fmt.Sprintf("invalid setting value '%s'. must be 'on' or 'off'", settingValue))
	}

	uri := "/zones/" + zoneID + "/argo/smart_routing"

	res, err := api.makeRequest("PATCH", uri, ArgoFeatureSetting{Value: settingValue})
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errMakeRequestError)
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errUnmarshalError)
	}
	return argoDetailsResponse.Result, nil
}

// ArgoTieredCaching returns the current settings for tiered caching.
//
// API reference: TBA
func (api *API) ArgoTieredCaching(zoneID string) (ArgoFeatureSetting, error) {
	uri := "/zones/" + zoneID + "/argo/tiered_caching"

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errMakeRequestError)
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errUnmarshalError)
	}
	return argoDetailsResponse.Result, nil
}

// UpdateArgoTieredCaching updates the setting for tiered caching.
//
// API reference: TBA
func (api *API) UpdateArgoTieredCaching(zoneID, settingValue string) (ArgoFeatureSetting, error) {
	if !contains(validSettingValues, settingValue) {
		return ArgoFeatureSetting{}, errors.New(fmt.Sprintf("invalid setting value '%s'. must be 'on' or 'off'", settingValue))
	}

	uri := "/zones/" + zoneID + "/argo/tiered_caching"

	res, err := api.makeRequest("PATCH", uri, ArgoFeatureSetting{Value: settingValue})
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errMakeRequestError)
	}

	var argoDetailsResponse ArgoDetailsResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return ArgoFeatureSetting{}, errors.Wrap(err, errUnmarshalError)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	}
	return argoDetailsResponse.Result, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

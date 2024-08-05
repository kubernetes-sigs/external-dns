package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type DeviceDexTestData map[string]interface{}

type DeviceDexTest struct {
	TestID      string             `json:"test_id"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Interval    string             `json:"interval"`
	Enabled     bool               `json:"enabled"`
	Updated     time.Time          `json:"updated"`
	Created     time.Time          `json:"created"`
	Data        *DeviceDexTestData `json:"data"`
}

type DeviceDexTests struct {
	DexTests []DeviceDexTest `json:"dex_tests"`
}

type DeviceDexTestResponse struct {
	Response
	Result DeviceDexTest `json:"result"`
}

type DeviceDexTestListResponse struct {
	Response
	Result DeviceDexTests `json:"result"`
}

type ListDeviceDexTestParams struct{}

type CreateDeviceDexTestParams struct {
	TestID      string             `json:"test_id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Interval    string             `json:"interval"`
	Enabled     bool               `json:"enabled"`
	Data        *DeviceDexTestData `json:"data"`
}

type UpdateDeviceDexTestParams struct {
	TestID      string             `json:"test_id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Interval    string             `json:"interval"`
	Enabled     bool               `json:"enabled"`
	Data        *DeviceDexTestData `json:"data"`
}

// ListDexTests returns all Device Dex Tests for a given account.
//
// API reference : https://developers.cloudflare.com/api/operations/device-dex-test-details
func (api *API) ListDexTests(ctx context.Context, rc *ResourceContainer, params ListDeviceDexTestParams) (DeviceDexTests, error) {
	if rc.Level != AccountRouteLevel {
		return DeviceDexTests{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/dex_tests", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DeviceDexTests{}, err
	}

	var response DeviceDexTestListResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return DeviceDexTests{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// CreateDeviceDexTest created a new Device Dex Test
//
// API reference: https://developers.cloudflare.com/api/operations/device-dex-test-create-device-dex-test
func (api *API) CreateDeviceDexTest(ctx context.Context, rc *ResourceContainer, params CreateDeviceDexTestParams) (DeviceDexTest, error) {
	if rc.Level != AccountRouteLevel {
		return DeviceDexTest{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/dex_tests", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return DeviceDexTest{}, err
	}

	var deviceDexTestResponse DeviceDexTestResponse
	if err := json.Unmarshal(res, &deviceDexTestResponse); err != nil {
		return DeviceDexTest{}, fmt.Errorf("%s: %w\n\nres: %s", errUnmarshalError, err, string(res))
	}

	return deviceDexTestResponse.Result, err
}

// UpdateDeviceDexTest Updates a Device Dex Test.
//
// API reference: https://developers.cloudflare.com/api/operations/device-dex-test-update-device-dex-test
func (api *API) UpdateDeviceDexTest(ctx context.Context, rc *ResourceContainer, params UpdateDeviceDexTestParams) (DeviceDexTest, error) {
	if rc.Level != AccountRouteLevel {
		return DeviceDexTest{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/dex_tests/%s", rc.Level, rc.Identifier, params.TestID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return DeviceDexTest{}, err
	}

	var deviceDexTestsResponse DeviceDexTestResponse

	if err := json.Unmarshal(res, &deviceDexTestsResponse); err != nil {
		return DeviceDexTest{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return deviceDexTestsResponse.Result, err
}

// GetDeviceDexTest gets a single Device Dex Test.
//
// API reference: https://developers.cloudflare.com/api/operations/device-dex-test-get-device-dex-test
func (api *API) GetDeviceDexTest(ctx context.Context, rc *ResourceContainer, testID string) (DeviceDexTest, error) {
	if rc.Level != AccountRouteLevel {
		return DeviceDexTest{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/dex_tests/%s", rc.Level, rc.Identifier, testID)

	deviceDexTestResponse := DeviceDexTestResponse{}
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DeviceDexTest{}, err
	}

	if err := json.Unmarshal(res, &deviceDexTestResponse); err != nil {
		return DeviceDexTest{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return deviceDexTestResponse.Result, err
}

// DeleteDexTest deletes a Device Dex Test.
//
// API reference: https://developers.cloudflare.com/api/operations/device-dex-test-delete-device-dex-test
func (api *API) DeleteDexTest(ctx context.Context, rc *ResourceContainer, testID string) (DeviceDexTests, error) {
	if rc.Level != AccountRouteLevel {
		return DeviceDexTests{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/dex_tests/%s", rc.Level, rc.Identifier, testID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return DeviceDexTests{}, err
	}

	var response DeviceDexTestListResponse
	if err := json.Unmarshal(res, &response); err != nil {
		return DeviceDexTests{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, err
}

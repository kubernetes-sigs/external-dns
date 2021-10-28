package safedns

import "github.com/ukfast/sdk-go/pkg/connection"

// GetZoneSliceResponseBody represents an API response body containing []Zone data
type GetZoneSliceResponseBody struct {
	connection.APIResponseBody
	Data []Zone `json:"data"`
}

// GetZoneResponseBody represents an API response body containing Zone data
type GetZoneResponseBody struct {
	connection.APIResponseBody
	Data Zone `json:"data"`
}

// GetRecordSliceResponseBody represents an API response body containing []Record data
type GetRecordSliceResponseBody struct {
	connection.APIResponseBody
	Data []Record `json:"data"`
}

// GetRecordResponseBody represents an API response body containing Record data
type GetRecordResponseBody struct {
	connection.APIResponseBody
	Data Record `json:"data"`
}

// GetNoteSliceResponseBody represents an API response body containing []Note data
type GetNoteSliceResponseBody struct {
	connection.APIResponseBody
	Data []Note `json:"data"`
}

// GetNoteResponseBody represents an API response body containing Note data
type GetNoteResponseBody struct {
	connection.APIResponseBody
	Data Note `json:"data"`
}

// GetTemplateSliceResponseBody represents an API response body containing []Template data
type GetTemplateSliceResponseBody struct {
	connection.APIResponseBody
	Data []Template `json:"data"`
}

// GetTemplateResponseBody represents an API response body containing Template data
type GetTemplateResponseBody struct {
	connection.APIResponseBody
	Data Template `json:"data"`
}

// GetSettingsSliceResponseBody represents an API response body containing []Settings data
type GetSettingsSliceResponseBody struct {
	connection.APIResponseBody
	Data []Settings `json:"data"`
}

// GetSettingsResponseBody represents an API response body containing Settings data
type GetSettingsResponseBody struct {
	connection.APIResponseBody
	Data Settings `json:"data"`
}

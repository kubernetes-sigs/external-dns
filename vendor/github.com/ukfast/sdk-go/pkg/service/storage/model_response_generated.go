package storage

import "github.com/ukfast/sdk-go/pkg/connection"

// GetSolutionSliceResponseBody represents an API response body containing []Solution data
type GetSolutionSliceResponseBody struct {
	connection.APIResponseBody
	Data []Solution `json:"data"`
}

// GetSolutionResponseBody represents an API response body containing Solution data
type GetSolutionResponseBody struct {
	connection.APIResponseBody
	Data Solution `json:"data"`
}

// GetHostSliceResponseBody represents an API response body containing []Host data
type GetHostSliceResponseBody struct {
	connection.APIResponseBody
	Data []Host `json:"data"`
}

// GetHostResponseBody represents an API response body containing Host data
type GetHostResponseBody struct {
	connection.APIResponseBody
	Data Host `json:"data"`
}

// GetVolumeSliceResponseBody represents an API response body containing []Volume data
type GetVolumeSliceResponseBody struct {
	connection.APIResponseBody
	Data []Volume `json:"data"`
}

// GetVolumeResponseBody represents an API response body containing Volume data
type GetVolumeResponseBody struct {
	connection.APIResponseBody
	Data Volume `json:"data"`
}

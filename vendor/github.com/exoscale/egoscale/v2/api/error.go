package api

import "errors"

var (
	// ErrNotFound represents an error indicating a non-existent resource.
	ErrNotFound = errors.New("resource not found")

	// ErrTooManyFound represents an error indicating multiple results found for a single resource.
	ErrTooManyFound = errors.New("multiple resources found")

	// ErrInvalidRequest represents an error indicating that the caller's request is invalid.
	ErrInvalidRequest = errors.New("invalid request")

	// ErrAPIError represents an error indicating an API-side issue.
	ErrAPIError = errors.New("API error")
)

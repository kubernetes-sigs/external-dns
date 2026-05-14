package ovh

import (
	"fmt"
	"strings"
)

// APIError represents an error that can occurred while calling the API.
type APIError struct {
	// Error class
	Class string `json:"class,omitempty"`
	// Error message.
	Message string `json:"message"`
	// Error details
	Details map[string]string `json:"details,omitempty"`
	// HTTP code.
	Code int
	// ID of the request
	QueryID string
}

// Let's make sure that APIError always satisfies the fmt.Stringer and error interfaces
var _ fmt.Stringer = APIError{}
var _ error = APIError{}

func (err APIError) Error() string {
	var sb strings.Builder
	sb.Grow(128)

	// Base message
	fmt.Fprint(&sb, "OVHcloud API error (status code ", err.Code, "): ")

	// Append class if any
	if err.Class != "" {
		fmt.Fprint(&sb, err.Class, ": ")
	}

	// Catch missing IAM permissions, if any
	var missingIAMActionsDetails []string
	if missingAuthenticationActions, ok := err.Details["unauthorizedActionsByAuthentication"]; ok && missingAuthenticationActions != "" {
		missingIAMActionsDetails = append(missingIAMActionsDetails, missingAuthenticationActions)
	}
	if missingIAMActions, ok := err.Details["unauthorizedActionsByIAM"]; ok && missingIAMActions != "" {
		missingIAMActionsDetails = append(missingIAMActionsDetails, missingIAMActions)
	}

	message := err.Message
	if len(missingIAMActionsDetails) > 0 {
		message += fmt.Sprintf(" (missing IAM permissions: %s)", strings.Join(missingIAMActionsDetails, ", "))
	}

	// Real error message, quoted
	fmt.Fprintf(&sb, "%q", message)

	// QueryID if any
	if err.QueryID != "" {
		fmt.Fprint(&sb, " (X-OVH-Query-Id: ", err.QueryID, ")")
	}

	return sb.String()
}

func (err APIError) String() string {
	return err.Error()
}

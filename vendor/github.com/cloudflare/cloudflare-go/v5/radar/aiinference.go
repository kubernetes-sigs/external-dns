// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AIInferenceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIInferenceService] method instead.
type AIInferenceService struct {
	Options          []option.RequestOption
	Summary          *AIInferenceSummaryService
	TimeseriesGroups *AIInferenceTimeseriesGroupService
}

// NewAIInferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIInferenceService(opts ...option.RequestOption) (r *AIInferenceService) {
	r = &AIInferenceService{}
	r.Options = opts
	r.Summary = NewAIInferenceSummaryService(opts...)
	r.TimeseriesGroups = NewAIInferenceTimeseriesGroupService(opts...)
	return
}

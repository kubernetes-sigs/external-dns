// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AIService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	Options          []option.RequestOption
	ToMarkdown       *AIToMarkdownService
	Inference        *AIInferenceService
	Bots             *AIBotService
	TimeseriesGroups *AITimeseriesGroupService
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r *AIService) {
	r = &AIService{}
	r.Options = opts
	r.ToMarkdown = NewAIToMarkdownService(opts...)
	r.Inference = NewAIInferenceService(opts...)
	r.Bots = NewAIBotService(opts...)
	r.TimeseriesGroups = NewAITimeseriesGroupService(opts...)
	return
}

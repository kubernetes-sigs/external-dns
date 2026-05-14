// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// CaptionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCaptionService] method instead.
type CaptionService struct {
	Options  []option.RequestOption
	Language *CaptionLanguageService
}

// NewCaptionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCaptionService(opts ...option.RequestOption) (r *CaptionService) {
	r = &CaptionService{}
	r.Options = opts
	r.Language = NewCaptionLanguageService(opts...)
	return
}

// Lists the available captions or subtitles for a specific video.
func (r *CaptionService) Get(ctx context.Context, identifier string, query CaptionGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[Caption], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions", query.AccountID, identifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists the available captions or subtitles for a specific video.
func (r *CaptionService) GetAutoPaging(ctx context.Context, identifier string, query CaptionGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Caption] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, identifier, query, opts...))
}

type Caption struct {
	// Whether the caption was generated via AI.
	Generated bool `json:"generated"`
	// The language label displayed in the native language to users.
	Label string `json:"label"`
	// The language tag in BCP 47 format.
	Language string `json:"language"`
	// The status of a generated caption.
	Status CaptionStatus `json:"status"`
	JSON   captionJSON   `json:"-"`
}

// captionJSON contains the JSON metadata for the struct [Caption]
type captionJSON struct {
	Generated   apijson.Field
	Label       apijson.Field
	Language    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Caption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionJSON) RawJSON() string {
	return r.raw
}

// The status of a generated caption.
type CaptionStatus string

const (
	CaptionStatusReady      CaptionStatus = "ready"
	CaptionStatusInprogress CaptionStatus = "inprogress"
	CaptionStatusError      CaptionStatus = "error"
)

func (r CaptionStatus) IsKnown() bool {
	switch r {
	case CaptionStatusReady, CaptionStatusInprogress, CaptionStatusError:
		return true
	}
	return false
}

type CaptionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

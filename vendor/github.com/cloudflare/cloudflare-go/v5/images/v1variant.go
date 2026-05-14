// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package images

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// V1VariantService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1VariantService] method instead.
type V1VariantService struct {
	Options []option.RequestOption
}

// NewV1VariantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1VariantService(opts ...option.RequestOption) (r *V1VariantService) {
	r = &V1VariantService{}
	r.Options = opts
	return
}

// Specify variants that allow you to resize images for different use cases.
func (r *V1VariantService) New(ctx context.Context, params V1VariantNewParams, opts ...option.RequestOption) (res *V1VariantNewResponse, err error) {
	var env V1VariantNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing variants.
func (r *V1VariantService) List(ctx context.Context, query V1VariantListParams, opts ...option.RequestOption) (res *Variant, err error) {
	var env V1VariantListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deleting a variant purges the cache for all images associated with the variant.
func (r *V1VariantService) Delete(ctx context.Context, variantID string, body V1VariantDeleteParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env V1VariantDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%s", body.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updating a variant purges the cache for all images associated with the variant.
func (r *V1VariantService) Edit(ctx context.Context, variantID string, params V1VariantEditParams, opts ...option.RequestOption) (res *V1VariantEditResponse, err error) {
	var env V1VariantEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%s", params.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch details for a single variant.
func (r *V1VariantService) Get(ctx context.Context, variantID string, query V1VariantGetParams, opts ...option.RequestOption) (res *V1VariantGetResponse, err error) {
	var env V1VariantGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%s", query.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Variant struct {
	Variants VariantVariants `json:"variants"`
	JSON     variantJSON     `json:"-"`
}

// variantJSON contains the JSON metadata for the struct [Variant]
type variantJSON struct {
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Variant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantJSON) RawJSON() string {
	return r.raw
}

type VariantVariants struct {
	Hero VariantVariantsHero `json:"hero"`
	JSON variantVariantsJSON `json:"-"`
}

// variantVariantsJSON contains the JSON metadata for the struct [VariantVariants]
type variantVariantsJSON struct {
	Hero        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantVariantsJSON) RawJSON() string {
	return r.raw
}

type VariantVariantsHero struct {
	ID string `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options VariantVariantsHeroOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                    `json:"neverRequireSignedURLs"`
	JSON                   variantVariantsHeroJSON `json:"-"`
}

// variantVariantsHeroJSON contains the JSON metadata for the struct
// [VariantVariantsHero]
type variantVariantsHeroJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *VariantVariantsHero) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantVariantsHeroJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type VariantVariantsHeroOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit VariantVariantsHeroOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata VariantVariantsHeroOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                        `json:"width,required"`
	JSON  variantVariantsHeroOptionsJSON `json:"-"`
}

// variantVariantsHeroOptionsJSON contains the JSON metadata for the struct
// [VariantVariantsHeroOptions]
type variantVariantsHeroOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantVariantsHeroOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantVariantsHeroOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type VariantVariantsHeroOptionsFit string

const (
	VariantVariantsHeroOptionsFitScaleDown VariantVariantsHeroOptionsFit = "scale-down"
	VariantVariantsHeroOptionsFitContain   VariantVariantsHeroOptionsFit = "contain"
	VariantVariantsHeroOptionsFitCover     VariantVariantsHeroOptionsFit = "cover"
	VariantVariantsHeroOptionsFitCrop      VariantVariantsHeroOptionsFit = "crop"
	VariantVariantsHeroOptionsFitPad       VariantVariantsHeroOptionsFit = "pad"
)

func (r VariantVariantsHeroOptionsFit) IsKnown() bool {
	switch r {
	case VariantVariantsHeroOptionsFitScaleDown, VariantVariantsHeroOptionsFitContain, VariantVariantsHeroOptionsFitCover, VariantVariantsHeroOptionsFitCrop, VariantVariantsHeroOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type VariantVariantsHeroOptionsMetadata string

const (
	VariantVariantsHeroOptionsMetadataKeep      VariantVariantsHeroOptionsMetadata = "keep"
	VariantVariantsHeroOptionsMetadataCopyright VariantVariantsHeroOptionsMetadata = "copyright"
	VariantVariantsHeroOptionsMetadataNone      VariantVariantsHeroOptionsMetadata = "none"
)

func (r VariantVariantsHeroOptionsMetadata) IsKnown() bool {
	switch r {
	case VariantVariantsHeroOptionsMetadataKeep, VariantVariantsHeroOptionsMetadataCopyright, VariantVariantsHeroOptionsMetadataNone:
		return true
	}
	return false
}

type V1VariantNewResponse struct {
	Variant V1VariantNewResponseVariant `json:"variant"`
	JSON    v1VariantNewResponseJSON    `json:"-"`
}

// v1VariantNewResponseJSON contains the JSON metadata for the struct
// [V1VariantNewResponse]
type v1VariantNewResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1VariantNewResponseVariant struct {
	ID string `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options V1VariantNewResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                            `json:"neverRequireSignedURLs"`
	JSON                   v1VariantNewResponseVariantJSON `json:"-"`
}

// v1VariantNewResponseVariantJSON contains the JSON metadata for the struct
// [V1VariantNewResponseVariant]
type v1VariantNewResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1VariantNewResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseVariantJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantNewResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit V1VariantNewResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata V1VariantNewResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                `json:"width,required"`
	JSON  v1VariantNewResponseVariantOptionsJSON `json:"-"`
}

// v1VariantNewResponseVariantOptionsJSON contains the JSON metadata for the struct
// [V1VariantNewResponseVariantOptions]
type v1VariantNewResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantNewResponseVariantOptionsFit string

const (
	V1VariantNewResponseVariantOptionsFitScaleDown V1VariantNewResponseVariantOptionsFit = "scale-down"
	V1VariantNewResponseVariantOptionsFitContain   V1VariantNewResponseVariantOptionsFit = "contain"
	V1VariantNewResponseVariantOptionsFitCover     V1VariantNewResponseVariantOptionsFit = "cover"
	V1VariantNewResponseVariantOptionsFitCrop      V1VariantNewResponseVariantOptionsFit = "crop"
	V1VariantNewResponseVariantOptionsFitPad       V1VariantNewResponseVariantOptionsFit = "pad"
)

func (r V1VariantNewResponseVariantOptionsFit) IsKnown() bool {
	switch r {
	case V1VariantNewResponseVariantOptionsFitScaleDown, V1VariantNewResponseVariantOptionsFitContain, V1VariantNewResponseVariantOptionsFitCover, V1VariantNewResponseVariantOptionsFitCrop, V1VariantNewResponseVariantOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type V1VariantNewResponseVariantOptionsMetadata string

const (
	V1VariantNewResponseVariantOptionsMetadataKeep      V1VariantNewResponseVariantOptionsMetadata = "keep"
	V1VariantNewResponseVariantOptionsMetadataCopyright V1VariantNewResponseVariantOptionsMetadata = "copyright"
	V1VariantNewResponseVariantOptionsMetadataNone      V1VariantNewResponseVariantOptionsMetadata = "none"
)

func (r V1VariantNewResponseVariantOptionsMetadata) IsKnown() bool {
	switch r {
	case V1VariantNewResponseVariantOptionsMetadataKeep, V1VariantNewResponseVariantOptionsMetadataCopyright, V1VariantNewResponseVariantOptionsMetadataNone:
		return true
	}
	return false
}

type V1VariantEditResponse struct {
	Variant V1VariantEditResponseVariant `json:"variant"`
	JSON    v1VariantEditResponseJSON    `json:"-"`
}

// v1VariantEditResponseJSON contains the JSON metadata for the struct
// [V1VariantEditResponse]
type v1VariantEditResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseJSON) RawJSON() string {
	return r.raw
}

type V1VariantEditResponseVariant struct {
	ID string `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options V1VariantEditResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                             `json:"neverRequireSignedURLs"`
	JSON                   v1VariantEditResponseVariantJSON `json:"-"`
}

// v1VariantEditResponseVariantJSON contains the JSON metadata for the struct
// [V1VariantEditResponseVariant]
type v1VariantEditResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1VariantEditResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseVariantJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantEditResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit V1VariantEditResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata V1VariantEditResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                 `json:"width,required"`
	JSON  v1VariantEditResponseVariantOptionsJSON `json:"-"`
}

// v1VariantEditResponseVariantOptionsJSON contains the JSON metadata for the
// struct [V1VariantEditResponseVariantOptions]
type v1VariantEditResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantEditResponseVariantOptionsFit string

const (
	V1VariantEditResponseVariantOptionsFitScaleDown V1VariantEditResponseVariantOptionsFit = "scale-down"
	V1VariantEditResponseVariantOptionsFitContain   V1VariantEditResponseVariantOptionsFit = "contain"
	V1VariantEditResponseVariantOptionsFitCover     V1VariantEditResponseVariantOptionsFit = "cover"
	V1VariantEditResponseVariantOptionsFitCrop      V1VariantEditResponseVariantOptionsFit = "crop"
	V1VariantEditResponseVariantOptionsFitPad       V1VariantEditResponseVariantOptionsFit = "pad"
)

func (r V1VariantEditResponseVariantOptionsFit) IsKnown() bool {
	switch r {
	case V1VariantEditResponseVariantOptionsFitScaleDown, V1VariantEditResponseVariantOptionsFitContain, V1VariantEditResponseVariantOptionsFitCover, V1VariantEditResponseVariantOptionsFitCrop, V1VariantEditResponseVariantOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type V1VariantEditResponseVariantOptionsMetadata string

const (
	V1VariantEditResponseVariantOptionsMetadataKeep      V1VariantEditResponseVariantOptionsMetadata = "keep"
	V1VariantEditResponseVariantOptionsMetadataCopyright V1VariantEditResponseVariantOptionsMetadata = "copyright"
	V1VariantEditResponseVariantOptionsMetadataNone      V1VariantEditResponseVariantOptionsMetadata = "none"
)

func (r V1VariantEditResponseVariantOptionsMetadata) IsKnown() bool {
	switch r {
	case V1VariantEditResponseVariantOptionsMetadataKeep, V1VariantEditResponseVariantOptionsMetadataCopyright, V1VariantEditResponseVariantOptionsMetadataNone:
		return true
	}
	return false
}

type V1VariantGetResponse struct {
	Variant V1VariantGetResponseVariant `json:"variant"`
	JSON    v1VariantGetResponseJSON    `json:"-"`
}

// v1VariantGetResponseJSON contains the JSON metadata for the struct
// [V1VariantGetResponse]
type v1VariantGetResponseJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1VariantGetResponseVariant struct {
	ID string `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options V1VariantGetResponseVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                            `json:"neverRequireSignedURLs"`
	JSON                   v1VariantGetResponseVariantJSON `json:"-"`
}

// v1VariantGetResponseVariantJSON contains the JSON metadata for the struct
// [V1VariantGetResponseVariant]
type v1VariantGetResponseVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *V1VariantGetResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseVariantJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantGetResponseVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit V1VariantGetResponseVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata V1VariantGetResponseVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                                `json:"width,required"`
	JSON  v1VariantGetResponseVariantOptionsJSON `json:"-"`
}

// v1VariantGetResponseVariantOptionsJSON contains the JSON metadata for the struct
// [V1VariantGetResponseVariantOptions]
type v1VariantGetResponseVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponseVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantGetResponseVariantOptionsFit string

const (
	V1VariantGetResponseVariantOptionsFitScaleDown V1VariantGetResponseVariantOptionsFit = "scale-down"
	V1VariantGetResponseVariantOptionsFitContain   V1VariantGetResponseVariantOptionsFit = "contain"
	V1VariantGetResponseVariantOptionsFitCover     V1VariantGetResponseVariantOptionsFit = "cover"
	V1VariantGetResponseVariantOptionsFitCrop      V1VariantGetResponseVariantOptionsFit = "crop"
	V1VariantGetResponseVariantOptionsFitPad       V1VariantGetResponseVariantOptionsFit = "pad"
)

func (r V1VariantGetResponseVariantOptionsFit) IsKnown() bool {
	switch r {
	case V1VariantGetResponseVariantOptionsFitScaleDown, V1VariantGetResponseVariantOptionsFitContain, V1VariantGetResponseVariantOptionsFitCover, V1VariantGetResponseVariantOptionsFitCrop, V1VariantGetResponseVariantOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type V1VariantGetResponseVariantOptionsMetadata string

const (
	V1VariantGetResponseVariantOptionsMetadataKeep      V1VariantGetResponseVariantOptionsMetadata = "keep"
	V1VariantGetResponseVariantOptionsMetadataCopyright V1VariantGetResponseVariantOptionsMetadata = "copyright"
	V1VariantGetResponseVariantOptionsMetadataNone      V1VariantGetResponseVariantOptionsMetadata = "none"
)

func (r V1VariantGetResponseVariantOptionsMetadata) IsKnown() bool {
	switch r {
	case V1VariantGetResponseVariantOptionsMetadataKeep, V1VariantGetResponseVariantOptionsMetadataCopyright, V1VariantGetResponseVariantOptionsMetadataNone:
		return true
	}
	return false
}

type V1VariantNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	ID        param.Field[string] `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[V1VariantNewParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r V1VariantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantNewParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[V1VariantNewParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[V1VariantNewParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r V1VariantNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantNewParamsOptionsFit string

const (
	V1VariantNewParamsOptionsFitScaleDown V1VariantNewParamsOptionsFit = "scale-down"
	V1VariantNewParamsOptionsFitContain   V1VariantNewParamsOptionsFit = "contain"
	V1VariantNewParamsOptionsFitCover     V1VariantNewParamsOptionsFit = "cover"
	V1VariantNewParamsOptionsFitCrop      V1VariantNewParamsOptionsFit = "crop"
	V1VariantNewParamsOptionsFitPad       V1VariantNewParamsOptionsFit = "pad"
)

func (r V1VariantNewParamsOptionsFit) IsKnown() bool {
	switch r {
	case V1VariantNewParamsOptionsFitScaleDown, V1VariantNewParamsOptionsFitContain, V1VariantNewParamsOptionsFitCover, V1VariantNewParamsOptionsFitCrop, V1VariantNewParamsOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type V1VariantNewParamsOptionsMetadata string

const (
	V1VariantNewParamsOptionsMetadataKeep      V1VariantNewParamsOptionsMetadata = "keep"
	V1VariantNewParamsOptionsMetadataCopyright V1VariantNewParamsOptionsMetadata = "copyright"
	V1VariantNewParamsOptionsMetadataNone      V1VariantNewParamsOptionsMetadata = "none"
)

func (r V1VariantNewParamsOptionsMetadata) IsKnown() bool {
	switch r {
	case V1VariantNewParamsOptionsMetadataKeep, V1VariantNewParamsOptionsMetadataCopyright, V1VariantNewParamsOptionsMetadataNone:
		return true
	}
	return false
}

type V1VariantNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   V1VariantNewResponse  `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantNewResponseEnvelopeJSON    `json:"-"`
}

// v1VariantNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantNewResponseEnvelope]
type v1VariantNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantNewResponseEnvelopeSuccess bool

const (
	V1VariantNewResponseEnvelopeSuccessTrue V1VariantNewResponseEnvelopeSuccess = true
)

func (r V1VariantNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1VariantListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1VariantListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Variant               `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantListResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantListResponseEnvelopeJSON    `json:"-"`
}

// v1VariantListResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantListResponseEnvelope]
type v1VariantListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantListResponseEnvelopeSuccess bool

const (
	V1VariantListResponseEnvelopeSuccessTrue V1VariantListResponseEnvelopeSuccess = true
)

func (r V1VariantListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1VariantDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1VariantDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantDeleteResponseEnvelopeJSON    `json:"-"`
}

// v1VariantDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantDeleteResponseEnvelope]
type v1VariantDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantDeleteResponseEnvelopeSuccess bool

const (
	V1VariantDeleteResponseEnvelopeSuccessTrue V1VariantDeleteResponseEnvelopeSuccess = true
)

func (r V1VariantDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1VariantEditParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[V1VariantEditParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r V1VariantEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantEditParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[V1VariantEditParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[V1VariantEditParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r V1VariantEditParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantEditParamsOptionsFit string

const (
	V1VariantEditParamsOptionsFitScaleDown V1VariantEditParamsOptionsFit = "scale-down"
	V1VariantEditParamsOptionsFitContain   V1VariantEditParamsOptionsFit = "contain"
	V1VariantEditParamsOptionsFitCover     V1VariantEditParamsOptionsFit = "cover"
	V1VariantEditParamsOptionsFitCrop      V1VariantEditParamsOptionsFit = "crop"
	V1VariantEditParamsOptionsFitPad       V1VariantEditParamsOptionsFit = "pad"
)

func (r V1VariantEditParamsOptionsFit) IsKnown() bool {
	switch r {
	case V1VariantEditParamsOptionsFitScaleDown, V1VariantEditParamsOptionsFitContain, V1VariantEditParamsOptionsFitCover, V1VariantEditParamsOptionsFitCrop, V1VariantEditParamsOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type V1VariantEditParamsOptionsMetadata string

const (
	V1VariantEditParamsOptionsMetadataKeep      V1VariantEditParamsOptionsMetadata = "keep"
	V1VariantEditParamsOptionsMetadataCopyright V1VariantEditParamsOptionsMetadata = "copyright"
	V1VariantEditParamsOptionsMetadataNone      V1VariantEditParamsOptionsMetadata = "none"
)

func (r V1VariantEditParamsOptionsMetadata) IsKnown() bool {
	switch r {
	case V1VariantEditParamsOptionsMetadataKeep, V1VariantEditParamsOptionsMetadataCopyright, V1VariantEditParamsOptionsMetadataNone:
		return true
	}
	return false
}

type V1VariantEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   V1VariantEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantEditResponseEnvelopeJSON    `json:"-"`
}

// v1VariantEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantEditResponseEnvelope]
type v1VariantEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantEditResponseEnvelopeSuccess bool

const (
	V1VariantEditResponseEnvelopeSuccessTrue V1VariantEditResponseEnvelopeSuccess = true
)

func (r V1VariantEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1VariantGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1VariantGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   V1VariantGetResponse  `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantGetResponseEnvelopeJSON    `json:"-"`
}

// v1VariantGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantGetResponseEnvelope]
type v1VariantGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantGetResponseEnvelopeSuccess bool

const (
	V1VariantGetResponseEnvelopeSuccessTrue V1VariantGetResponseEnvelopeSuccess = true
)

func (r V1VariantGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

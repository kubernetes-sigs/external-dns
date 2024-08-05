package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type ImagesVariant struct {
	ID                     string                `json:"id,omitempty"`
	NeverRequireSignedURLs *bool                 `json:"neverRequireSignedURLs,omitempty"`
	Options                ImagesVariantsOptions `json:"options,omitempty"`
}

type ImagesVariantsOptions struct {
	Fit      string `json:"fit,omitempty"`
	Height   int    `json:"height,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type ListImageVariantsParams struct{}

type ListImagesVariantsResponse struct {
	Result ListImageVariantsResult `json:"result,omitempty"`
	Response
}

type ListImageVariantsResult struct {
	ImagesVariants map[string]ImagesVariant `json:"variants,omitempty"`
}

type CreateImagesVariantParams struct {
	ID                     string                `json:"id,omitempty"`
	NeverRequireSignedURLs *bool                 `json:"neverRequireSignedURLs,omitempty"`
	Options                ImagesVariantsOptions `json:"options,omitempty"`
}

type UpdateImagesVariantParams struct {
	ID                     string                `json:"-"`
	NeverRequireSignedURLs *bool                 `json:"neverRequireSignedURLs,omitempty"`
	Options                ImagesVariantsOptions `json:"options,omitempty"`
}

type ImagesVariantResult struct {
	Variant ImagesVariant `json:"variant,omitempty"`
}

type ImagesVariantResponse struct {
	Result ImagesVariantResult `json:"result,omitempty"`
	Response
}

// Lists existing variants.
//
// API Reference: https://developers.cloudflare.com/api/operations/cloudflare-images-variants-list-variants
func (api *API) ListImagesVariants(ctx context.Context, rc *ResourceContainer, params ListImageVariantsParams) (ListImageVariantsResult, error) {
	if rc.Identifier == "" {
		return ListImageVariantsResult{}, ErrMissingAccountID
	}

	baseURL := fmt.Sprintf("/accounts/%s/images/v1/variants", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return ListImageVariantsResult{}, err
	}

	var listImageVariantsResponse ListImagesVariantsResponse
	err = json.Unmarshal(res, &listImageVariantsResponse)
	if err != nil {
		return ListImageVariantsResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return listImageVariantsResponse.Result, nil
}

// Fetch details for a single variant.
//
// API Reference: https://developers.cloudflare.com/api/operations/cloudflare-images-variants-variant-details
func (api *API) GetImagesVariant(ctx context.Context, rc *ResourceContainer, variantID string) (ImagesVariant, error) {
	if rc.Identifier == "" {
		return ImagesVariant{}, ErrMissingAccountID
	}

	baseURL := fmt.Sprintf("/accounts/%s/images/v1/variants/%s", rc.Identifier, variantID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return ImagesVariant{}, err
	}

	var imagesVariantDetailResponse ImagesVariantResponse
	err = json.Unmarshal(res, &imagesVariantDetailResponse)
	if err != nil {
		return ImagesVariant{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return imagesVariantDetailResponse.Result.Variant, nil
}

// Specify variants that allow you to resize images for different use cases.
//
// API Reference: https://developers.cloudflare.com/api/operations/cloudflare-images-variants-create-a-variant
func (api *API) CreateImagesVariant(ctx context.Context, rc *ResourceContainer, params CreateImagesVariantParams) (ImagesVariant, error) {
	if rc.Identifier == "" {
		return ImagesVariant{}, ErrMissingAccountID
	}

	baseURL := fmt.Sprintf("/accounts/%s/images/v1/variants", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, baseURL, params)
	if err != nil {
		return ImagesVariant{}, err
	}

	var createImagesVariantResponse ImagesVariantResponse
	err = json.Unmarshal(res, &createImagesVariantResponse)
	if err != nil {
		return ImagesVariant{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return createImagesVariantResponse.Result.Variant, nil
}

// Deleting a variant purges the cache for all images associated with the variant.
//
// API Reference: https://developers.cloudflare.com/api/operations/cloudflare-images-variants-variant-details
func (api *API) DeleteImagesVariant(ctx context.Context, rc *ResourceContainer, variantID string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	baseURL := fmt.Sprintf("/accounts/%s/images/v1/variants/%s", rc.Identifier, variantID)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, baseURL, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	return nil
}

// Updating a variant purges the cache for all images associated with the variant.
//
// API Reference: https://developers.cloudflare.com/api/operations/cloudflare-images-variants-variant-details
func (api *API) UpdateImagesVariant(ctx context.Context, rc *ResourceContainer, params UpdateImagesVariantParams) (ImagesVariant, error) {
	if rc.Identifier == "" {
		return ImagesVariant{}, ErrMissingAccountID
	}

	baseURL := fmt.Sprintf("/accounts/%s/images/v1/variants/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, baseURL, params)
	if err != nil {
		return ImagesVariant{}, err
	}

	var imagesVariantDetailResponse ImagesVariantResponse
	err = json.Unmarshal(res, &imagesVariantDetailResponse)
	if err != nil {
		return ImagesVariant{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return imagesVariantDetailResponse.Result.Variant, nil
}

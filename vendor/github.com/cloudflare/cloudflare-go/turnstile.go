package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var ErrMissingSiteKey = errors.New("required site key missing")

type TurnstileWidget struct {
	SiteKey      string     `json:"sitekey,omitempty"`
	Secret       string     `json:"secret,omitempty"`
	CreatedOn    *time.Time `json:"created_on,omitempty"`
	ModifiedOn   *time.Time `json:"modified_on,omitempty"`
	Name         string     `json:"name,omitempty"`
	Domains      []string   `json:"domains,omitempty"`
	Mode         string     `json:"mode,omitempty"`
	BotFightMode bool       `json:"bot_fight_mode,omitempty"`
	Region       string     `json:"region,omitempty"`
	OffLabel     bool       `json:"offlabel,omitempty"`
}

type CreateTurnstileWidgetParams struct {
	Name         string   `json:"name,omitempty"`
	Domains      []string `json:"domains,omitempty"`
	Mode         string   `json:"mode,omitempty"`
	BotFightMode bool     `json:"bot_fight_mode,omitempty"`
	Region       string   `json:"region,omitempty"`
	OffLabel     bool     `json:"offlabel,omitempty"`
}

type UpdateTurnstileWidgetParams struct {
	SiteKey      string   `json:"-"`
	Name         string   `json:"name,omitempty"`
	Domains      []string `json:"domains,omitempty"`
	Mode         string   `json:"mode,omitempty"`
	BotFightMode bool     `json:"bot_fight_mode,omitempty"`
	Region       string   `json:"region,omitempty"`
	OffLabel     bool     `json:"offlabel,omitempty"`
}

type TurnstileWidgetResponse struct {
	Response
	Result TurnstileWidget `json:"result"`
}

type ListTurnstileWidgetParams struct {
	ResultInfo
	Direction string         `url:"direction,omitempty"`
	Order     OrderDirection `url:"order,omitempty"`
}

type ListTurnstileWidgetResponse struct {
	Response
	ResultInfo `json:"result_info"`
	Result     []TurnstileWidget `json:"result"`
}

type RotateTurnstileWidgetParams struct {
	SiteKey               string `json:"-"`
	InvalidateImmediately bool   `json:"invalidate_immediately,omitempty"`
}

// CreateTurnstileWidget creates a new challenge widgets.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-properties
func (api *API) CreateTurnstileWidget(ctx context.Context, rc *ResourceContainer, params CreateTurnstileWidgetParams) (TurnstileWidget, error) {
	if rc.Identifier == "" {
		return TurnstileWidget{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/challenges/widgets", rc.Identifier)
	res, err := api.makeRequestContext(ctx, "POST", uri, params)
	if err != nil {
		return TurnstileWidget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r TurnstileWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return TurnstileWidget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// ListTurnstileWidgets lists challenge widgets.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-list-challenge-widgets
func (api *API) ListTurnstileWidgets(ctx context.Context, rc *ResourceContainer, params ListTurnstileWidgetParams) ([]TurnstileWidget, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []TurnstileWidget{}, &ResultInfo{}, ErrMissingAccountID
	}
	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = 25
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var widgets []TurnstileWidget
	var r ListTurnstileWidgetResponse
	for {
		uri := buildURI(fmt.Sprintf("/accounts/%s/challenges/widgets", rc.Identifier), params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)

		if err != nil {
			return []TurnstileWidget{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []TurnstileWidget{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}

		widgets = append(widgets, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return widgets, &r.ResultInfo, nil
}

// GetTurnstileWidget shows a single challenge widget configuration.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-challenge-widget-details
func (api *API) GetTurnstileWidget(ctx context.Context, rc *ResourceContainer, siteKey string) (TurnstileWidget, error) {
	if rc.Identifier == "" {
		return TurnstileWidget{}, ErrMissingAccountID
	}

	if siteKey == "" {
		return TurnstileWidget{}, ErrMissingSiteKey
	}

	uri := fmt.Sprintf("/accounts/%s/challenges/widgets/%s", rc.Identifier, siteKey)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TurnstileWidget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r TurnstileWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return TurnstileWidget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// UpdateTurnstileWidget update the configuration of a widget.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-update-a-challenge-widget
func (api *API) UpdateTurnstileWidget(ctx context.Context, rc *ResourceContainer, params UpdateTurnstileWidgetParams) (TurnstileWidget, error) {
	if rc.Identifier == "" {
		return TurnstileWidget{}, ErrMissingAccountID
	}

	if params.SiteKey == "" {
		return TurnstileWidget{}, ErrMissingSiteKey
	}

	uri := fmt.Sprintf("/accounts/%s/challenges/widgets/%s", rc.Identifier, params.SiteKey)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return TurnstileWidget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r TurnstileWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return TurnstileWidget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// RotateTurnstileWidget generates a new secret key for this widget. If
// invalidate_immediately is set to false, the previous secret remains valid for
// 2 hours.
//
// Note that secrets cannot be rotated again during the grace period.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-rotate-secret-for-a-challenge-widget
func (api *API) RotateTurnstileWidget(ctx context.Context, rc *ResourceContainer, param RotateTurnstileWidgetParams) (TurnstileWidget, error) {
	if rc.Identifier == "" {
		return TurnstileWidget{}, ErrMissingAccountID
	}
	if param.SiteKey == "" {
		return TurnstileWidget{}, ErrMissingSiteKey
	}

	uri := fmt.Sprintf("/accounts/%s/challenges/widgets/%s/rotate_secret", rc.Identifier, param.SiteKey)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, param)

	if err != nil {
		return TurnstileWidget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r TurnstileWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return TurnstileWidget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// DeleteTurnstileWidget delete a challenge widget.
//
// API reference: https://api.cloudflare.com/#challenge-widgets-delete-a-challenge-widget
func (api *API) DeleteTurnstileWidget(ctx context.Context, rc *ResourceContainer, siteKey string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if siteKey == "" {
		return ErrMissingSiteKey
	}
	uri := fmt.Sprintf("/accounts/%s/challenges/widgets/%s", rc.Identifier, siteKey)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r TurnstileWidgetResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
)

// Behavior represents a single zt risk behavior config.
type Behavior struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	RiskLevel   RiskLevel `json:"risk_level"`
	Enabled     *bool     `json:"enabled"`
}

// Wrapper used to have full-fidelity repro of json structure.
type Behaviors struct {
	Behaviors map[string]Behavior `json:"behaviors"`
}

// BehaviorResponse represents the response from the zt risk scoring endpoint
// and contains risk behaviors for an account.
type BehaviorResponse struct {
	Success  bool      `json:"success"`
	Result   Behaviors `json:"result"`
	Errors   []string  `json:"errors"`
	Messages []string  `json:"messages"`
}

// Behaviors returns all zero trust risk scoring behaviors for the provided account
//
// API reference: https://developers.cloudflare.com/api/operations/dlp-zt-risk-score-get-behaviors
func (api *API) Behaviors(ctx context.Context, accountID string) (Behaviors, error) {
	uri := fmt.Sprintf("/accounts/%s/zt_risk_scoring/behaviors", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Behaviors{}, err
	}

	var r BehaviorResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Behaviors{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateBehaviors returns all zero trust risk scoring behaviors for the provided account
// NOTE: description/name updates are no-ops, risk_level [low medium high] and enabled [true/false] results in modifications
//
// API reference: https://developers.cloudflare.com/api/operations/dlp-zt-risk-score-put-behaviors
func (api *API) UpdateBehaviors(ctx context.Context, accountID string, behaviors Behaviors) (Behaviors, error) {
	uri := fmt.Sprintf("/accounts/%s/zt_risk_scoring/behaviors", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, behaviors)
	if err != nil {
		return Behaviors{}, err
	}

	var r BehaviorResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Behaviors{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

type RiskLevel int

const (
	_ RiskLevel = iota
	Low
	Medium
	High
)

func (p RiskLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p RiskLevel) String() string {
	return [...]string{"low", "medium", "high"}[p-1]
}

func (p *RiskLevel) UnmarshalJSON(data []byte) error {
	var (
		s   string
		err error
	)
	err = json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	v, err := RiskLevelFromString(s)
	if err != nil {
		return err
	}
	*p = *v
	return nil
}

func RiskLevelFromString(s string) (*RiskLevel, error) {
	s = strings.ToLower(s)
	var v RiskLevel
	switch s {
	case "low":
		v = Low
	case "medium":
		v = Medium
	case "high":
		v = High
	default:
		return nil, fmt.Errorf("unknown variant for risk level: %s", s)
	}
	return &v, nil
}

func (p RiskLevel) IntoRef() *RiskLevel {
	return &p
}

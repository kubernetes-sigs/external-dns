package rest

import (
	"fmt"

	"gopkg.in/ns1/ns1-go.v2/rest/model/alerting"
)

// NewUsageAlert creates a new usage alert with proper type and validation
func NewUsageAlert(name string, subtype string, alertAtPercent int, notifierListIds []string) (*alerting.Alert, error) {
	alert := alerting.NewUsageAlert(name, subtype, alertAtPercent, notifierListIds)

	if err := alerting.ValidateUsageAlert(alert); err != nil {
		return nil, err
	}

	return alert, nil
}

// UsageAlertPatch defines the fields that can be updated on a usage alert
// Intentionally excluding Type and Subtype which cannot be changed via PATCH
type UsageAlertPatch struct {
	Name            *string                  `json:"name,omitempty"`
	Data            *alerting.UsageAlertData `json:"data,omitempty"`
	NotifierListIds *[]string                `json:"notifier_list_ids,omitempty"`
	ZoneNames       *[]string                `json:"zone_names,omitempty"`
}

// NewUsageAlertPatchFunc is a function that modifies a UsageAlertPatch
type NewUsageAlertPatchFunc func(*UsageAlertPatch)

// NewUsageAlertPatch creates a new patch for updating a usage alert
func NewUsageAlertPatch(opts ...NewUsageAlertPatchFunc) (*UsageAlertPatch, error) {
	patch := &UsageAlertPatch{}

	for _, opt := range opts {
		opt(patch)
	}

	// Validate data if it's being updated
	if patch.Data != nil {
		if patch.Data.AlertAtPercent < 1 || patch.Data.AlertAtPercent > 100 {
			return nil, fmt.Errorf("data.alert_at_percent must be between 1 and 100")
		}
	}

	return patch, nil
}

// WithName sets the name field in the patch
func WithName(name string) NewUsageAlertPatchFunc {
	return func(p *UsageAlertPatch) {
		p.Name = &name
	}
}

// WithAlertAtPercent sets the alert_at_percent field in the patch
func WithAlertAtPercent(percent int) NewUsageAlertPatchFunc {
	return func(p *UsageAlertPatch) {
		p.Data = &alerting.UsageAlertData{AlertAtPercent: percent}
	}
}

// WithNotifierListIds sets the notifier_list_ids field in the patch
func WithNotifierListIds(ids []string) NewUsageAlertPatchFunc {
	return func(p *UsageAlertPatch) {
		p.NotifierListIds = &ids
	}
}

// WithZoneNames sets the zone_names field in the patch
func WithZoneNames(names []string) NewUsageAlertPatchFunc {
	return func(p *UsageAlertPatch) {
		p.ZoneNames = &names
	}
}

// IsUsageAlert checks if an alert is a usage alert
func IsUsageAlert(alert *alerting.Alert) bool {
	if alert == nil || alert.Type == nil || alert.Subtype == nil {
		return false
	}

	if *alert.Type != alerting.AlertTypeAccount {
		return false
	}

	if ok := alerting.AllowedUsageSubtypes[*alert.Subtype]; !ok {
		return false
	}

	return true
}

// FilterUsageAlerts filters a list of alerts to only include usage alerts
func FilterUsageAlerts(alerts []*alerting.Alert) []*alerting.Alert {
	usageAlerts := make([]*alerting.Alert, 0)
	for _, alert := range alerts {
		if IsUsageAlert(alert) {
			usageAlerts = append(usageAlerts, alert)
		}
	}
	return usageAlerts
}

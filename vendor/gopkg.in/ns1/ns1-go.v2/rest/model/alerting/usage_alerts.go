package alerting

import (
	"encoding/json"
	"fmt"
)

// ValidationError creates a formatted validation error
func ValidationError(msg string) error {
	return fmt.Errorf("validation error: %s", msg)
}

// AlertTypeAccount is the alert type for account-scoped alerts
const AlertTypeAccount = "account"

// UsageSubtypes defines all supported usage alert subtypes
var UsageSubtypes = struct {
	QueryUsage       string
	RecordUsage      string
	ChinaQueryUsage  string
	RumDecisionUsage string
	FilterChainUsage string
	MonitorUsage     string
}{
	QueryUsage:       "query_usage",
	RecordUsage:      "record_usage",
	ChinaQueryUsage:  "china_query_usage",
	RumDecisionUsage: "rum_decision_usage",
	FilterChainUsage: "filter_chain_usage",
	MonitorUsage:     "monitor_usage",
}

// AllowedUsageSubtypes is a map of valid usage alert subtypes
var AllowedUsageSubtypes = map[string]bool{
	UsageSubtypes.QueryUsage:       true,
	UsageSubtypes.RecordUsage:      true,
	UsageSubtypes.ChinaQueryUsage:  true,
	UsageSubtypes.RumDecisionUsage: true,
	UsageSubtypes.FilterChainUsage: true,
	UsageSubtypes.MonitorUsage:     true,
}

// UsageAlertData contains the threshold percentage for usage alerts
type UsageAlertData struct {
	AlertAtPercent int `json:"alert_at_percent"`
}

// NewUsageAlert creates a new account usage alert
func NewUsageAlert(name string, subtype string, alertAtPercent int, notifierListIds []string) *Alert {
	accountType := AlertTypeAccount
	return &Alert{
		Name:            &name,
		Type:            &accountType,
		Subtype:         &subtype,
		Data:            MarshalUsageAlertData(UsageAlertData{AlertAtPercent: alertAtPercent}),
		NotifierListIds: notifierListIds,
		ZoneNames:       []string{},
	}
}

// MarshalUsageAlertData converts UsageAlertData to JSON for the Alert.Data field
func MarshalUsageAlertData(data UsageAlertData) json.RawMessage {
	jsonData, _ := json.Marshal(data)
	return jsonData
}

// UnmarshalUsageAlertData extracts UsageAlertData from an Alert
func UnmarshalUsageAlertData(alert *Alert) (UsageAlertData, error) {
	if alert.Data == nil {
		return UsageAlertData{}, fmt.Errorf("alert data is nil")
	}

	var data UsageAlertData
	err := json.Unmarshal(alert.Data, &data)
	return data, err
}

// ValidateUsageAlert validates a usage alert configuration
func ValidateUsageAlert(alert *Alert) error {
	if alert.Type == nil || *alert.Type != AlertTypeAccount {
		return fmt.Errorf("type must be 'account'")
	}

	if alert.Subtype == nil {
		return fmt.Errorf("subtype is required")
	}

	if _, ok := AllowedUsageSubtypes[*alert.Subtype]; !ok {
		return fmt.Errorf("invalid subtype %q", *alert.Subtype)
	}

	data, err := UnmarshalUsageAlertData(alert)
	if err != nil {
		return fmt.Errorf("invalid data format: %v", err)
	}

	if data.AlertAtPercent < 1 || data.AlertAtPercent > 100 {
		return fmt.Errorf("data.alert_at_percent must be between 1 and 100")
	}

	if alert.Name == nil || *alert.Name == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}

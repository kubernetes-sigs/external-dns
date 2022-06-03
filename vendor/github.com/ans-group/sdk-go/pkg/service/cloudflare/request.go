package cloudflare

// CreateAccountRequest represents a request to create an account
type CreateAccountRequest struct {
	Name string `json:"name"`
}

// PatchAccountRequest represents a request to patch an account
type PatchAccountRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateAccountMemberRequest represents a request to create an account member
type CreateAccountMemberRequest struct {
	EmailAddress string `json:"email_address"`
}

// CreateOrchestrationRequest represents a request to create new orchestration
type CreateOrchestrationRequest struct {
	AccountMemberDeploymentID string `json:"account_member_deployment_id"`
	ZoneName                  string `json:"zone_name"`
	ZoneSubscriptionID        string `json:"zone_subscription_id"`
	AccountID                 string `json:"account_id"`
	AccountName               string `json:"account_name"`
	AdministratorEmailAddress string `json:"administrator_email_address"`
}

// CreateRecordRequest represents a request to create a zone
type CreateZoneRequest struct {
	AccountID      string `json:"account_id"`
	Name           string `json:"name"`
	SubscriptionID string `json:"subscription_id"`
}

// PatchZoneRequest represents a request to patch a zone
type PatchZoneRequest struct {
	SubscriptionID string `json:"subscription_id,omitempty"`
}

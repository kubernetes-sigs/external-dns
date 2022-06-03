package cloudflare

import "github.com/ans-group/sdk-go/pkg/connection"

// Account represents a Cloudflare account
type Account struct {
	ID                  string              `json:"id"`
	Status              string              `json:"status"`
	Name                string              `json:"name"`
	CloudflareAccountID string              `json:"cloudflare_account_id"`
	CreatedAt           connection.DateTime `json:"created_at"`
	UpdatedAt           connection.DateTime `json:"updated_at"`
}

// AccountMember represents a Cloudflare account member
type AccountMember struct {
	AccountID    string `json:"account_id"`
	EmailAddress string `json:"email_address"`
}

// SpendPlan represents a Cloudflare spend plan
type SpendPlan struct {
	ID        string              `json:"id"`
	Amount    float32             `json:"amount"`
	StartedAt connection.DateTime `json:"started_at"`
	EndedAt   connection.DateTime `json:"ended_at"`
	CreatedAt connection.DateTime `json:"created_at"`
	UpdatedAt connection.DateTime `json:"updated_at"`
}

// Subscription represents a Cloudflare subscription
type Subscription struct {
	ID                   string              `json:"id"`
	Name                 string              `json:"name"`
	Type                 string              `json:"type"`
	Description          string              `json:"description"`
	Price                float32             `json:"price"`
	CloudflareRatePlanID string              `json:"cloudflare_rate_plan_id"`
	CreatedAt            connection.DateTime `json:"created_at"`
	UpdatedAt            connection.DateTime `json:"updated_at"`
}

// Zone represents a Cloudflare zone
type Zone struct {
	ID               string              `json:"id"`
	AccountID        string              `json:"account_id"`
	Name             string              `json:"name"`
	SubscriptionID   string              `json:"subscription_id"`
	CloudflareZoneID string              `json:"cloudflare_zone_id"`
	CreatedAt        connection.DateTime `json:"created_at"`
	UpdatedAt        connection.DateTime `json:"updated_at"`
}

// TotalSpend represents total spend
type TotalSpend struct {
	SpendPlanAmount float32 `json:"spend_plan_amount"`
	TotalSpend      float32 `json:"total_spend"`
}

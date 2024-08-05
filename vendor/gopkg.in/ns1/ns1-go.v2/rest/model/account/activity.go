package account

// Activity wraps an NS1 /account/activity resource
type Activity struct {
	UserID       string `json:"user_id,omitempty"`
	ResourceID   string `json:"resource_id,omitempty"`
	Timestamp    int    `json:"timestamp,omitempty"`
	UserType     string `json:"user_type,omitempty"`
	Action       string `json:"action,omitempty"`
	UserName     string `json:"user_name,omitempty"`
	ID           string `json:"id,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

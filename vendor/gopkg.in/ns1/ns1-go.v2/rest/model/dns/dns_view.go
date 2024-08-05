package dns

// View wraps an NS1 views/ resource
type View struct {
	Name       string   `json:"name,omitempty"`
	CreatedAt  int      `json:"created_at,omitempty"`
	UpdatedAt  int      `json:"updated_at,omitempty"`
	ReadACLs   []string `json:"read_acls"`
	UpdateACLs []string `json:"update_acls"`
	Zones      []string `json:"zones"`
	Networks   []int    `json:"networks"`
	Preference int      `json:"preference,omitempty"`
}

// NewView takes a viewName and creates a *DNSView
func NewView(viewName string) *View {
	return &View{
		Name: viewName,
	}
}

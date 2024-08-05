package dns

type Version struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Active      bool   `json:"active,omitempty"`
	ActivatedAt int    `json:"activated_at,omitempty"`
	CreatedAt   int    `json:"created_at,omitempty"`
}

package cloudflare

type Policy struct {
	ID               string            `json:"id"`
	PermissionGroups []PermissionGroup `json:"permission_groups"`
	ResourceGroups   []ResourceGroup   `json:"resource_groups"`
	Access           string            `json:"access"`
}

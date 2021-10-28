package govultr

// ListOptions are the available query params
type ListOptions struct {
	// These query params are used for all list calls that support pagination
	PerPage int    `url:"per_page,omitempty"`
	Cursor  string `url:"cursor,omitempty"`

	// These three query params are currently used for the list instance call
	// These may be extended to other list calls
	// https://www.vultr.com/api/#operation/list-instances
	MainIP string `url:"main_ip,omitempty"`
	Label  string `url:"label,omitempty"`
	Tag    string `url:"tag,omitempty"`
	Region string `url:"region,omitempty"`

	// Query params that can be used on the list snapshots call
	// https://www.vultr.com/api/#operation/list-snapshots
	Description string `url:"description,omitempty"`
}

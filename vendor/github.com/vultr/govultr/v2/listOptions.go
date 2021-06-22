package govultr

// ListOptions are the available fields that can be used with pagination
type ListOptions struct {
	PerPage int    `url:"per_page,omitempty"`
	Cursor  string `url:"cursor,omitempty"`
}

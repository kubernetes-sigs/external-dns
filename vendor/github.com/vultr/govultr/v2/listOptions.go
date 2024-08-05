package govultr

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
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
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// ListOptions are the available fields that can be used with pagination
||||||| parent of 4d7e5ad26 (update vendored files)
// ListOptions are the available fields that can be used with pagination
=======
// ListOptions are the available query params
>>>>>>> 4d7e5ad26 (update vendored files)
type ListOptions struct {
	// These query params are used for all list calls that support pagination
	PerPage int    `url:"per_page,omitempty"`
	Cursor  string `url:"cursor,omitempty"`
<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======

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
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// ListOptions are the available fields that can be used with pagination
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// ListOptions are the available fields that can be used with pagination
=======
// ListOptions are the available query params
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
type ListOptions struct {
	// These query params are used for all list calls that support pagination
	PerPage int    `url:"per_page,omitempty"`
	Cursor  string `url:"cursor,omitempty"`
<<<<<<< HEAD
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======

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
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}

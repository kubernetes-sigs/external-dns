package l7policies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// L7Policy is a collection of L7 rules associated with a Listener, and which
// may also have an association to a back-end pool.
type L7Policy struct {
	// The unique ID for the L7 policy.
	ID string `json:"id"`

	// Name of the L7 policy.
	Name string `json:"name"`

	// The ID of the listener.
	ListenerID string `json:"listener_id"`

	// The L7 policy action. One of REDIRECT_TO_POOL, REDIRECT_TO_URL, or REJECT.
	Action string `json:"action"`

	// The position of this policy on the listener.
	Position int32 `json:"position"`

	// A human-readable description for the resource.
	Description string `json:"description"`

	// TenantID is the UUID of the project who owns the L7 policy in neutron-lbaas.
	// Only administrative users can specify a project UUID other than their own.
	TenantID string `json:"tenant_id"`

	// Requests matching this policy will be redirected to the pool with this ID.
	// Only valid if action is REDIRECT_TO_POOL.
	RedirectPoolID string `json:"redirect_pool_id"`

	// Requests matching this policy will be redirected to this URL.
	// Only valid if action is REDIRECT_TO_URL.
	RedirectURL string `json:"redirect_url"`

	// The administrative state of the L7 policy, which is up (true) or down (false).
	AdminStateUp bool `json:"admin_state_up"`

	// Rules are List of associated L7 rule IDs.
	Rules []Rule `json:"rules"`
}

// Rule represents layer 7 load balancing rule.
type Rule struct {
	// The unique ID for the L7 rule.
	ID string `json:"id"`

	// The L7 rule type. One of COOKIE, FILE_TYPE, HEADER, HOST_NAME, or PATH.
	RuleType string `json:"type"`

	// The comparison type for the L7 rule. One of CONTAINS, ENDS_WITH, EQUAL_TO, REGEX, or STARTS_WITH.
	CompareType string `json:"compare_type"`

	// The value to use for the comparison. For example, the file type to compare.
	Value string `json:"value"`

	// TenantID is the UUID of the project who owns the rule in neutron-lbaas.
	// Only administrative users can specify a project UUID other than their own.
	TenantID string `json:"tenant_id"`

	// The key to use for the comparison. For example, the name of the cookie to evaluate.
	Key string `json:"key"`

	// When true the logic of the rule is inverted. For example, with invert true,
	// equal to would become not equal to. Default is false.
	Invert bool `json:"invert"`

	// The administrative state of the L7 rule, which is up (true) or down (false).
	AdminStateUp bool `json:"admin_state_up"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a l7policy.
func (r commonResult) Extract() (*L7Policy, error) {
	var s struct {
		L7Policy *L7Policy `json:"l7policy"`
	}
	err := r.ExtractInto(&s)
	return s.L7Policy, err
}

// CreateResult represents the result of a Create operation. Call its Extract
// method to interpret the result as a L7Policy.
type CreateResult struct {
	commonResult
}

// L7PolicyPage is the page returned by a pager when traversing over a
// collection of l7policies.
type L7PolicyPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of l7policies has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r L7PolicyPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"l7policies_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a L7PolicyPage struct is empty.
func (r L7PolicyPage) IsEmpty() (bool, error) {
	is, err := ExtractL7Policies(r)
	return len(is) == 0, err
}

// ExtractL7Policies accepts a Page struct, specifically a L7PolicyPage struct,
// and extracts the elements into a slice of L7Policy structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractL7Policies(r pagination.Page) ([]L7Policy, error) {
	var s struct {
		L7Policies []L7Policy `json:"l7policies"`
	}
	err := (r.(L7PolicyPage)).ExtractInto(&s)
	return s.L7Policies, err
}

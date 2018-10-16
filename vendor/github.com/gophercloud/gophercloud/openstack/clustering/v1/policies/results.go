package policies

import (
	"encoding/json"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Policy represents a clustering policy in the Openstack cloud
type Policy struct {
	CreatedAt time.Time              `json:"-"`
	Data      map[string]interface{} `json:"data"`
	Domain    string                 `json:"domain"`
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Project   string                 `json:"project"`
	Spec      map[string]interface{} `json:"spec"`
	Type      string                 `json:"type"`
	UpdatedAt time.Time              `json:"-"`
	User      string                 `json:"user"`
}

// ExtractPolicies interprets a page of results as a slice of Policy.
func ExtractPolicies(r pagination.Page) ([]Policy, error) {
	var s struct {
		Policies []Policy `json:"policies"`
	}
	err := (r.(PolicyPage)).ExtractInto(&s)
	return s.Policies, err
}

// PolicyPage contains a list page of all policies from a List call.
type PolicyPage struct {
	pagination.MarkerPageBase
}

// IsEmpty determines if a PolicyPage contains any results.
func (page PolicyPage) IsEmpty() (bool, error) {
	policies, err := ExtractPolicies(page)
	return len(policies) == 0, err
}

// LastMarker returns the last policy ID in a ListResult.
func (r PolicyPage) LastMarker() (string, error) {
	policies, err := ExtractPolicies(r)
	if err != nil {
		return "", err
	}
	if len(policies) == 0 {
		return "", nil
	}
	return policies[len(policies)-1].ID, nil
}

func (r *Policy) UnmarshalJSON(b []byte) error {
	type tmp Policy
	var s struct {
		tmp
		CreatedAt gophercloud.JSONRFC3339MilliNoZ `json:"created_at,omitempty"`
		UpdatedAt gophercloud.JSONRFC3339MilliNoZ `json:"updated_at,omitempty"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Policy(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return nil
}

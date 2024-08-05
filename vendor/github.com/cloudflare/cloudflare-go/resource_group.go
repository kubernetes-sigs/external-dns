package cloudflare

import "fmt"

type ResourceGroup struct {
	ID    string            `json:"id"`
	Name  string            `json:"name"`
	Meta  map[string]string `json:"meta"`
	Scope Scope             `json:"scope"`
}

type Scope struct {
	Key          string        `json:"key"`
	ScopeObjects []ScopeObject `json:"objects"`
}

type ScopeObject struct {
	Key string `json:"key"`
}

// NewResourceGroupForZone takes an existing zone and provides a resource group
// to be used within a Policy that allows access to that zone.
func NewResourceGroupForZone(zone Zone) ResourceGroup {
	return NewResourceGroup(fmt.Sprintf("com.cloudflare.api.account.zone.%s", zone.ID))
}

// NewResourceGroupForAccount takes an existing zone and provides a resource group
// to be used within a Policy that allows access to that account.
func NewResourceGroupForAccount(account Account) ResourceGroup {
	return NewResourceGroup(fmt.Sprintf("com.cloudflare.api.account.%s", account.ID))
}

// NewResourceGroup takes a Cloudflare-formatted key (e.g. 'com.cloudflare.api.%s') and
// returns a resource group to be used within a Policy to allow access to that resource.
func NewResourceGroup(key string) ResourceGroup {
	scope := Scope{
		Key: key,
		ScopeObjects: []ScopeObject{
			{
				Key: "*",
			},
		},
	}
	resourceGroup := ResourceGroup{
		ID:   "",
		Name: key,
		Meta: map[string]string{
			"editable": "false",
		},
		Scope: scope,
	}
	return resourceGroup
}

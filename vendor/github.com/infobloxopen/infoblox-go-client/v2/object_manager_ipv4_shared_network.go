package ibclient

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func (d *SharedNetwork) MarshalJSON() ([]byte, error) {
	type Alias SharedNetwork
	aux := &struct {
		Networks []interface{} `json:"networks"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	// Regular expression to match CIDR format
	isCIDR := regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\/\d{1,2}$`)

	for _, network := range d.Networks {
		if network != nil && network.Ref != "" {
			if isCIDR.MatchString(network.Ref) {
				networkView := network.NetworkView
				if networkView == "" {
					networkView = "default"
				}
				aux.Networks = append(aux.Networks, map[string]interface{}{
					"_ref": map[string]string{
						"network":      network.Ref,
						"network_view": networkView,
					},
				})
			} else {
				aux.Networks = append(aux.Networks, map[string]string{"_ref": network.Ref})
			}
		}
	}
	return json.Marshal(aux)
}

func (d *SharedNetwork) UnmarshalJSON(data []byte) error {
	type Alias SharedNetwork
	aux := &struct {
		*Alias
		Networks []map[string]interface{} `json:"networks"`
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.Networks = make([]*Ipv4Network, len(aux.Networks))
	for i, network := range aux.Networks {
		if ref, ok := network["_ref"].(string); ok {
			d.Networks[i] = &Ipv4Network{Ref: ref}
		} else {
			return fmt.Errorf("invalid network reference format")
		}
	}
	return nil
}

func (objMgr *ObjectManager) CreateIpv4SharedNetwork(name string, networks []string, networkView string, eas EA, comment string, disable bool, useOptions bool, options []*Dhcpoption) (*SharedNetwork, error) {
	if name == "" && len(networks) == 0 {
		return nil, fmt.Errorf("name and networks are required to create a shared network")
	}
	var ipv4Networks []*Ipv4Network
	for _, nw := range networks {
		if nw == "" {
			return nil, fmt.Errorf("networks cannot be empty")
		}
		ipv4Networks = append(ipv4Networks, &Ipv4Network{Ref: nw, NetworkView: networkView})
	}
	if networkView == "" {
		networkView = "default"
	}

	sharedNetwork := NewIpv4SharedNetwork("", name, ipv4Networks, eas, comment, disable, useOptions, options)
	sharedNetwork.NetworkView = networkView
	ref, err := objMgr.connector.CreateObject(sharedNetwork)
	if err != nil {
		return nil, err
	}
	sharedNetwork.Ref = ref
	return sharedNetwork, nil
}

func NewEmptyIpv4SharedNetwork() *SharedNetwork {
	sharedNetwork := &SharedNetwork{}
	sharedNetwork.SetReturnFields(append(sharedNetwork.ReturnFields(), "extattrs", "disable", "use_options", "options"))
	return sharedNetwork
}

func (objMgr *ObjectManager) GetIpv4SharedNetworkByRef(ref string) (*SharedNetwork, error) {
	sharedNetwork := NewEmptyIpv4SharedNetwork()
	err := objMgr.connector.GetObject(
		sharedNetwork, ref, NewQueryParams(false, nil), &sharedNetwork)

	return sharedNetwork, err
}

func (objMgr *ObjectManager) GetAllIpv4SharedNetwork(queryParams *QueryParams) ([]SharedNetwork, error) {
	var res []SharedNetwork
	sharedNetwork := NewEmptyIpv4SharedNetwork()
	err := objMgr.connector.GetObject(
		sharedNetwork, "", queryParams, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (objMgr *ObjectManager) UpdateIpv4SharedNetwork(ref string, name string, networks []string, networkView string, comment string, eas EA, disable bool, useOptions bool, options []*Dhcpoption) (*SharedNetwork, error) {
	if name == "" && len(networks) == 0 {
		return nil, fmt.Errorf("name and networks are required for a shared network")
	}
	var ipv4Networks []*Ipv4Network
	for _, nw := range networks {
		if nw == "" {
			return nil, fmt.Errorf("networks cannot be empty")
		}
		ipv4Networks = append(ipv4Networks, &Ipv4Network{Ref: nw, NetworkView: networkView})
	}
	sharedNetwork := NewIpv4SharedNetwork(ref, name, ipv4Networks, eas, comment, disable, useOptions, options)
	updatedRef, err := objMgr.connector.UpdateObject(sharedNetwork, ref)
	if err != nil {
		return nil, err
	}
	sharedNetwork.Ref = updatedRef
	return sharedNetwork, nil
}

func (objMgr *ObjectManager) DeleteIpv4SharedNetwork(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func NewIpv4SharedNetwork(ref string, name string, networks []*Ipv4Network, eas EA, comment string, disable bool, useOptions bool, options []*Dhcpoption) *SharedNetwork {
	sharedNetwork := NewEmptyIpv4SharedNetwork()
	sharedNetwork.Ref = ref
	sharedNetwork.Name = &name
	sharedNetwork.Networks = networks
	sharedNetwork.Ea = eas
	sharedNetwork.Comment = &comment
	sharedNetwork.Disable = &disable
	sharedNetwork.UseOptions = &useOptions
	if options != nil {
		sharedNetwork.Options = options
	}
	return sharedNetwork
}

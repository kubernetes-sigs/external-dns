package ibclient

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type AuthZonesLink struct {
	Fqdn    string
	DnsView string
}

func (d *DtcLbdn) MarshalJSON() ([]byte, error) {
	type Alias DtcLbdn
	aux := &struct {
		AuthZones []string       `json:"auth_zones"`
		Pools     []*DtcPoolLink `json:"pools"`
		Patterns  []string       `json:"patterns"`
		Topology  *string        `json:"topology"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	// Convert AuthZones to a slice of strings

	for _, zone := range d.AuthZones {
		if zone != nil && zone.Ref != "" {
			aux.AuthZones = append(aux.AuthZones, zone.Ref)
		}
	}

	// Convert Pools to a slice of DtcPoolLink
	for _, pool := range d.Pools {
		if pool != nil {
			aux.Pools = append(aux.Pools, pool)
		}
	}

	// Convert Patterns to a slice of strings
	for _, pattern := range d.Patterns {
		aux.Patterns = append(aux.Patterns, pattern)
	}

	// Ensure AuthZones, Pools, and Types are set to empty slices if nil
	if aux.AuthZones == nil {
		aux.AuthZones = make([]string, 0)
	}
	if aux.Pools == nil {
		aux.Pools = make([]*DtcPoolLink, 0)
	}
	if aux.Patterns == nil {
		aux.Patterns = make([]string, 0)
	}
	if d.Topology != nil && *d.Topology == "" {
		aux.Topology = nil
	} else {
		aux.Topology = d.Topology
	}

	return json.Marshal(aux)
}

func (d *DtcLbdn) UnmarshalJSON(data []byte) error {
	type Alias DtcLbdn
	aux := &struct {
		*Alias
		AuthZones []string `json:"auth_zones"`
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.AuthZones = make([]*ZoneAuth, len(aux.AuthZones))
	for i, ref := range aux.AuthZones {
		d.AuthZones[i] = &ZoneAuth{Ref: ref}
	}
	return nil
}

func (objMgr *ObjectManager) CreateDtcLbdn(name string, authZones []AuthZonesLink, comment string, disable bool, autoConsolidatedMonitors bool, ea EA,
	lbMethod string, patterns []string, persistence uint32, pools []*DtcPoolLink, priority uint32, topology *string, types []string, ttl uint32, usettl bool) (*DtcLbdn, error) {

	if name == "" || lbMethod == "" {
		return nil, fmt.Errorf("name and load balancing method fields are required to create a Dtc Lbdn object")
	}
	// get ref id of authzones and replace
	var zones []*ZoneAuth
	var err error
	if len(authZones) > 0 {
		zones, err = getAuthZones(authZones, objMgr)
		if err != nil {
			return nil, err
		}
	}

	// get ref id of pools and replace
	var dtcPoolLink []*DtcPoolLink
	if len(pools) > 0 {
		dtcPoolLink, err = getPools(pools, objMgr)
		if err != nil {
			return nil, err
		}
	}

	if lbMethod == "TOPOLOGY" && topology == nil {
		return nil, fmt.Errorf("topology field is required when load balancing method is TOPOLOGY")
	}
	//get ref id of topology and replace
	var topologyRef string
	if topology != nil {
		if *topology == "" && lbMethod != "TOPOLOGY" {
			topologyRef = ""
		} else {
			topologyRef, err = getTopology(*topology, objMgr)
			if err != nil {
				return nil, err
			}
		}
	}

	dtcLbdn := NewDtcLbdn("", name, zones, comment, disable, autoConsolidatedMonitors, ea,
		lbMethod, patterns, persistence, dtcPoolLink, priority, &topologyRef, types, ttl, usettl)
	ref, err := objMgr.connector.CreateObject(dtcLbdn)
	if err != nil {
		return nil, fmt.Errorf("error creating Dtc Lbdn object %s, err: %s", name, err)
	}
	dtcLbdn.Ref = ref
	return dtcLbdn, nil
}

func getTopology(topology string, objMgr *ObjectManager) (string, error) {
	var dtcTopology []DtcTopology
	var topologyRef string
	if topology == "" {
		return "", fmt.Errorf("topology field is required to retreive a unique Dtc Topology record")
	}
	isRef := regexp.MustCompile("^dtc:topology:*")
	if !isRef.MatchString(topology) {
		sf := map[string]string{
			"name": topology,
		}
		err := objMgr.connector.GetObject(&DtcTopology{}, "", NewQueryParams(false, sf), &dtcTopology)
		if err != nil {
			return "", fmt.Errorf("error getting Dtc Topology object %s, err: %s", topology, err)
		}

		if len(dtcTopology) > 0 {
			topologyRef = dtcTopology[0].Ref
		}
	}
	return topologyRef, nil
}

func getPools(pools []*DtcPoolLink, objMgr *ObjectManager) ([]*DtcPoolLink, error) {
	var dtcPoolLink []*DtcPoolLink

	for _, pool := range pools {
		sf := map[string]string{"name": pool.Pool}
		var dtcPools []DtcPool

		isRef := regexp.MustCompile("^dtc:pool:*")
		if !isRef.MatchString(pool.Pool) {
			err := objMgr.connector.GetObject(&DtcPool{}, "", NewQueryParams(false, sf), &dtcPools)
			if err != nil {
				return nil, fmt.Errorf("error getting Dtc Pool object %s, err: %s", pool.Pool, err)
			}
			if len(dtcPools) > 0 {
				dtcPoolLink = append(dtcPoolLink, &DtcPoolLink{Pool: dtcPools[0].Ref, Ratio: pool.Ratio})
			}
		}

	}
	return dtcPoolLink, nil
}

func getAuthZones(authZones []AuthZonesLink, objMgr *ObjectManager) ([]*ZoneAuth, error) {
	var zones []*ZoneAuth
	for _, authZone := range authZones {
		sf := map[string]string{
			"fqdn": authZone.Fqdn,
			"view": authZone.DnsView,
		}
		var zoneAuth []ZoneAuth
		err := objMgr.connector.GetObject(&ZoneAuth{}, "", NewQueryParams(false, sf), &zoneAuth)
		if err != nil {
			return nil, fmt.Errorf("error getting ZoneAuth object %s in %s DNS view, err: %s", authZone.Fqdn, authZone.DnsView, err)
		}
		if len(zoneAuth) > 0 {
			zones = append(zones, &zoneAuth[0])
		}
	}
	return zones, nil
}

func NewDtcLbdn(ref string, name string, authZones []*ZoneAuth, comment string, disable bool, autoConsolidatedMonitors bool, ea EA,
	lbMethod string, patterns []string, persistence uint32, pools []*DtcPoolLink, priority uint32, topology *string, types []string, ttl uint32, usettl bool) *DtcLbdn {

	lbdn := NewEmptyDtcLbdn()
	lbdn.Name = &name
	lbdn.Ref = ref
	lbdn.AuthZones = authZones
	lbdn.Comment = &comment
	lbdn.Disable = &disable
	lbdn.AutoConsolidatedMonitors = &autoConsolidatedMonitors
	lbdn.Ea = ea
	lbdn.LbMethod = lbMethod
	lbdn.Patterns = patterns
	lbdn.Persistence = &persistence
	lbdn.Pools = pools
	lbdn.Topology = topology
	lbdn.Priority = &priority

	lbdn.Types = types
	lbdn.Ttl = &ttl
	lbdn.UseTtl = &usettl
	return lbdn
}

func NewEmptyDtcLbdn() *DtcLbdn {
	dtcLbdn := &DtcLbdn{}
	dtcLbdn.SetReturnFields(append(dtcLbdn.ReturnFields(), "extattrs", "disable", "auth_zones", "auto_consolidated_monitors", "lb_method", "patterns", "persistence", "pools", "priority", "topology", "types", "health", "ttl", "use_ttl"))
	return dtcLbdn
}

func (objMgr *ObjectManager) DeleteDtcLbdn(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func (objMgr *ObjectManager) GetAllDtcLbdn(queryParams *QueryParams) ([]DtcLbdn, error) {
	var res []DtcLbdn
	lbdn := NewEmptyDtcLbdn()
	err := objMgr.connector.GetObject(lbdn, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("error getting Dtc Lbdn object, err: %s", err)
	}
	return res, nil
}

func (objMgr *ObjectManager) GetDtcLbdn(name string) (*DtcLbdn, error) {
	dtcLbdn := NewEmptyDtcLbdn()
	var res []DtcLbdn
	if name == "" {
		return nil, fmt.Errorf("name of the record is required to retrieve a unique Dtc Lbdn record")
	}
	sf := map[string]string{
		"name": name,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(dtcLbdn, "", queryParams, &res)
	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf("Dtc Lbdn record with name '%s' not found", name))
	}
	return &res[0], nil
}

func (objMgr *ObjectManager) GetDtcLbdnByRef(ref string) (*DtcLbdn, error) {
	dtcLbdn := NewEmptyDtcLbdn()
	err := objMgr.connector.GetObject(dtcLbdn, ref, NewQueryParams(false, nil), &dtcLbdn)
	if err != nil {
		return nil, err
	}
	return dtcLbdn, nil
}

func (objMgr *ObjectManager) UpdateDtcLbdn(ref string, name string, authZones []AuthZonesLink, comment string, disable bool, autoConsolidatedMonitors bool, ea EA,
	lbMethod string, patterns []string, persistence uint32, pools []*DtcPoolLink, priority uint32, topology *string, types []string, ttl uint32, usettl bool) (*DtcLbdn, error) {

	if lbMethod == "TOPOLOGY" && topology == nil {
		return nil, fmt.Errorf("topology field is required when load balancing method is TOPOLOGY")
	}
	// get ref id of authzones and replace
	var zones []*ZoneAuth
	var err error
	if len(authZones) > 0 {
		zones, err = getAuthZones(authZones, objMgr)
		if err != nil {
			return nil, err
		}
	}

	// get ref id of pools and replace
	var dtcPoolLink []*DtcPoolLink
	if len(pools) > 0 {
		dtcPoolLink, err = getPools(pools, objMgr)
		if err != nil {
			return nil, err
		}
	}

	//get ref id of topology and replace
	var topologyRef string
	if topology != nil {
		if *topology == "" && lbMethod != "TOPOLOGY" {
			topologyRef = ""
		} else {
			topologyRef, err = getTopology(*topology, objMgr)
			if err != nil {
				return nil, err
			}
		}
	}

	dtcLbdn := NewDtcLbdn(ref, name, zones, comment, disable, autoConsolidatedMonitors, ea,
		lbMethod, patterns, persistence, dtcPoolLink, priority, &topologyRef, types, ttl, usettl)
	newRef, err := objMgr.connector.UpdateObject(dtcLbdn, ref)
	if err != nil {
		return nil, fmt.Errorf("error updating Dtc Lbdn object %s, err: %s", name, err)
	}
	dtcLbdn.Ref = newRef
	dtcLbdn, err = objMgr.GetDtcLbdnByRef(newRef)
	if err != nil {
		return nil, fmt.Errorf("error getting updated Dtc Lbdn object %s, err: %s", name, err)
	}
	return dtcLbdn, nil
}

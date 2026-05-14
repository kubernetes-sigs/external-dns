package ibclient

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Monitor struct {
	Name string
	Type string
}

// Updating Servers in DtcServerLink with reference
func updateServerReferences(servers []*DtcServerLink, objMgr *ObjectManager) error {
	for _, link := range servers {
		sf := map[string]string{"name": link.Server}
		queryParams := NewQueryParams(false, sf)
		var serverResult []DtcServer
		err := objMgr.connector.GetObject(&DtcServer{}, "dtc:server", queryParams, &serverResult)
		if err != nil {
			return fmt.Errorf("error getting Dtc Server %s, err: %s", link.Server, err)
		}
		if len(serverResult) > 0 {
			link.Server = serverResult[0].Ref
		}
	}
	return nil
}

// get the monitor reference
func getMonitorReference(monitorName string, monitorType string, objMgr *ObjectManager) (string, error) {
	if monitorType == "" {
		return "", nil
	}
	fields := map[string]string{"name": monitorName}
	queryParams := NewQueryParams(false, fields)
	var monitorResult []DtcMonitorHttp

	monitorTypeKey := fmt.Sprintf("dtc:monitor:%s", monitorType)
	err := objMgr.connector.GetObject(&DtcMonitorHttp{}, monitorTypeKey, queryParams, &monitorResult)
	if err != nil {
		return "", fmt.Errorf("error getting Dtc Monitor object %s, err: %s", monitorName, err)
	}
	if len(monitorResult) > 0 {
		return monitorResult[0].Ref, nil
	}
	return "", fmt.Errorf("Dtc Monitor with name %s not found", monitorName)
}

func (cm ConsolidatedMonitorsWrapper) MarshalJSON() ([]byte, error) {
	if !cm.IsNull {
		if reflect.DeepEqual(cm.ConsolidatedMonitors, []*DtcPoolConsolidatedMonitorHealth{}) {
			return []byte("[]"), nil
		}
	}
	return json.Marshal(cm.ConsolidatedMonitors)
}

func (cm *ConsolidatedMonitorsWrapper) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		cm.IsNull = true
		cm.ConsolidatedMonitors = nil
		return nil
	}
	cm.IsNull = false
	return json.Unmarshal(data, &cm.ConsolidatedMonitors)
}

type ConsolidatedMonitorsWrapper struct {
	ConsolidatedMonitors []*DtcPoolConsolidatedMonitorHealth
	IsNull               bool
}

func (d *DtcPool) MarshalJSON() ([]byte, error) {
	type Alias DtcPool
	aux := &struct {
		Monitors             []string                     `json:"monitors"`
		Servers              []*DtcServerLink             `json:"servers"`
		ConsolidatedMonitors *ConsolidatedMonitorsWrapper `json:"consolidated_monitors,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	// Convert Monitors to a slice of strings
	if len(d.Monitors) == 0 {
		aux.Monitors = []string{}
	} else {
		for _, monitor := range d.Monitors {
			if monitor != nil {
				aux.Monitors = append(aux.Monitors, monitor.Ref)
			}
		}
	}

	// Unsetting Servers if they are empty
	if len(d.Servers) == 0 {
		aux.Servers = []*DtcServerLink{}
	} else {
		aux.Servers = d.Servers
	}

	// Conditionally handle ConsolidatedMonitors
	if d.AutoConsolidatedMonitors != nil {
		if *d.AutoConsolidatedMonitors {
			// auto_consolidated_monitors = true
			if d.ConsolidatedMonitors != nil {
				// consolidated_monitors is empty, omit it
				aux.ConsolidatedMonitors = nil
			}
		} else {
			// auto_consolidated_monitors = false
			if len(d.ConsolidatedMonitors) == 0 {
				// consolidated_monitors is empty, marshal as "[]"
				aux.ConsolidatedMonitors = &ConsolidatedMonitorsWrapper{IsNull: false, ConsolidatedMonitors: []*DtcPoolConsolidatedMonitorHealth{}}
			} else {
				// consolidated_monitors is non-empty, marshal as is
				aux.ConsolidatedMonitors = &ConsolidatedMonitorsWrapper{IsNull: false, ConsolidatedMonitors: d.ConsolidatedMonitors}
			}
		}
	}

	return json.Marshal(aux)
}
func (d *DtcPool) UnmarshalJSON(data []byte) error {
	type Alias DtcPool
	aux := &struct {
		Monitors []string `json:"monitors,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// Convert Monitors from []string to []*DtcMonitorHttp
	for _, ref := range aux.Monitors {
		d.Monitors = append(d.Monitors, &DtcMonitorHttp{Ref: ref})
	}

	return nil
}

func NewEmptyDtcPool() *DtcPool {
	poolDtc := &DtcPool{}
	poolDtc.SetReturnFields(append(poolDtc.ReturnFields(), "lb_preferred_method", "servers", "lb_dynamic_ratio_preferred", "monitors", "auto_consolidated_monitors", "consolidated_monitors", "disable",
		"extattrs", "health", "lb_alternate_method", "lb_alternate_topology", "lb_dynamic_ratio_alternate", "lb_preferred_topology", "quorum", "ttl", "use_ttl", "availability"))

	return poolDtc
}

func NewDtcPool(comment string,
	name string,
	lbPreferredMethod string,
	lbDynamicRatioPreferred *SettingDynamicratio,
	servers []*DtcServerLink,
	monitors []*DtcMonitorHttp,
	lbPreferredTopology *string,
	lbAlternateMethod string,
	lbAlternateTopology *string,
	lbDynamicRatioAlternate *SettingDynamicratio,
	eas EA,
	autoConsolidatedMonitors bool,
	availability string,
	consolidatedMonitors []*DtcPoolConsolidatedMonitorHealth,
	ttl uint32,
	useTTL bool,
	disable bool,
	quorum uint32) *DtcPool {
	DtcPool := NewEmptyDtcPool()
	DtcPool.Comment = &comment
	DtcPool.Name = &name
	DtcPool.LbPreferredMethod = lbPreferredMethod
	DtcPool.Servers = servers
	DtcPool.LbDynamicRatioPreferred = lbDynamicRatioPreferred
	DtcPool.Monitors = monitors
	DtcPool.LbPreferredTopology = lbPreferredTopology
	DtcPool.LbAlternateMethod = lbAlternateMethod
	DtcPool.LbAlternateTopology = lbAlternateTopology
	DtcPool.LbDynamicRatioAlternate = lbDynamicRatioAlternate
	DtcPool.Ea = eas
	DtcPool.AutoConsolidatedMonitors = &autoConsolidatedMonitors
	DtcPool.Availability = availability
	DtcPool.ConsolidatedMonitors = consolidatedMonitors
	DtcPool.Ttl = &ttl
	DtcPool.UseTtl = &useTTL
	DtcPool.Disable = &disable
	DtcPool.Quorum = &quorum
	return DtcPool
}

func (objMgr *ObjectManager) CreateDtcPool(
	comment string,
	name string,
	lbPreferredMethod string,
	lbDynamicRatioPreferred map[string]interface{},
	servers []*DtcServerLink,
	monitors []Monitor,
	lbPreferredTopology *string,
	lbAlternateMethod string,
	lbAlternateTopology *string,
	lbDynamicRatioAlternate map[string]interface{},
	eas EA,
	autoConsolidatedMonitors bool,
	userMonitors []map[string]interface{},
	availability string,
	ttl uint32,
	useTTL bool,
	disable bool,
	quorum uint32) (*DtcPool, error) {
	if name == "" || lbPreferredMethod == "" {
		return nil, fmt.Errorf("name and preferred load balancing method must be provided to create a pool")
	}
	if lbPreferredMethod == "DYNAMIC_RATIO" && lbDynamicRatioPreferred == nil {
		return nil, fmt.Errorf("LbDynamicRatioPreferred cannot be nil when the preferred load balancing method is set to DYNAMIC_RATIO")
	}
	if lbPreferredMethod == "TOPOLOGY" && lbPreferredTopology == nil {
		return nil, fmt.Errorf("preferred topology cannot be nil when preferred load balancing method is set to TOPOLOGY")
	}
	//update servers with server references
	err := updateServerReferences(servers, objMgr)
	if err != nil {
		return nil, err
	}
	// update the monitor in LbDynamicRatioPreferred with reference
	var lbDynamicRatioPreferredMethod *SettingDynamicratio
	if lbDynamicRatioPreferred != nil {
		monitor, _ := lbDynamicRatioPreferred["monitor"].(Monitor)
		method, _ := lbDynamicRatioPreferred["method"].(string)
		monitorMetric, _ := lbDynamicRatioPreferred["monitor_metric"].(string)
		monitorWeighing, _ := lbDynamicRatioPreferred["monitor_weighing"].(string)
		invertMonitorMetric, _ := lbDynamicRatioPreferred["invert_monitor_metric"].(bool)

		monitorRef, err := getMonitorReference(monitor.Name, monitor.Type, objMgr)
		if err != nil {
			return nil, err
		}
		lbDynamicRatioPreferredMethod = &SettingDynamicratio{
			Method:              method,
			Monitor:             monitorRef,
			MonitorMetric:       monitorMetric,
			MonitorWeighing:     monitorWeighing,
			InvertMonitorMetric: invertMonitorMetric,
		}
	} else {
		lbDynamicRatioPreferredMethod = nil
	}

	// Convert monitor names to monitor references
	var monitorResults []*DtcMonitorHttp
	for _, monitor := range monitors {
		monitorRef, err := getMonitorReference(monitor.Name, monitor.Type, objMgr)
		if err != nil {
			return nil, err
		}
		monitorResults = append(monitorResults, &DtcMonitorHttp{Ref: monitorRef})
	}
	//Update the topology name with the topology reference
	if lbPreferredTopology != nil {
		topology, err := getTopology(*lbPreferredTopology, objMgr)
		if err != nil {
			return nil, err
		}
		lbPreferredTopology = &topology
	}

	//Update the topology name with the topology reference
	if lbAlternateTopology != nil {
		topologyAlternate, err := getTopology(*lbAlternateTopology, objMgr)
		if err != nil {
			return nil, err
		}
		lbAlternateTopology = &topologyAlternate
	}
	//update the monitor in LbDynamicRatioPreferred with reference
	var lbDynamicRatioAlternateMethod *SettingDynamicratio
	if lbDynamicRatioAlternate != nil {
		monitorAlternate, _ := lbDynamicRatioAlternate["monitor"].(Monitor)
		methodAlternate, _ := lbDynamicRatioAlternate["method"].(string)
		monitorMetricAlternate, _ := lbDynamicRatioAlternate["monitor_metric"].(string)
		monitorWeighingAlternate, _ := lbDynamicRatioAlternate["monitor_weighing"].(string)
		interferometricAlternate, _ := lbDynamicRatioAlternate["invert_monitor_metric"].(bool)

		monitorRefAlternate, err := getMonitorReference(monitorAlternate.Name, monitorAlternate.Type, objMgr)
		if err != nil {
			return nil, err
		}
		lbDynamicRatioAlternateMethod = &SettingDynamicratio{
			Method:              methodAlternate,
			Monitor:             monitorRefAlternate,
			MonitorMetric:       monitorMetricAlternate,
			MonitorWeighing:     monitorWeighingAlternate,
			InvertMonitorMetric: interferometricAlternate,
		}
	} else {
		lbDynamicRatioAlternateMethod = nil
	}

	var consolidatedMonitors []*DtcPoolConsolidatedMonitorHealth
	if userMonitors != nil {
		if len(userMonitors) == 0 {
			consolidatedMonitors = []*DtcPoolConsolidatedMonitorHealth{}
		} else {
			for _, userMonitor := range userMonitors {
				monitor, okMonitor := userMonitor["monitor"].(Monitor)
				monitorAvailability, okAvail := userMonitor["availability"].(string)
				fullHealthComm, _ := userMonitor["full_health_communication"].(bool)
				members, okMember := userMonitor["members"].([]string)
				if !okMonitor {
					return nil, fmt.Errorf("required field missing: monitor")
				}

				if !okAvail {
					return nil, fmt.Errorf("required field missing: availability")
				}

				if !okMember {
					return nil, fmt.Errorf("required field missing: members")
				}
				monitorRef, err := getMonitorReference(monitor.Name, monitor.Type, objMgr)
				if err != nil {
					return nil, err
				}

				consolidatedMonitor := &DtcPoolConsolidatedMonitorHealth{
					Members:                 members,
					Monitor:                 monitorRef,
					Availability:            monitorAvailability,
					FullHealthCommunication: fullHealthComm,
				}
				consolidatedMonitors = append(consolidatedMonitors, consolidatedMonitor)
			}
		}
	}
	// Create the DtcPool
	poolDtc := NewDtcPool(comment, name, lbPreferredMethod, lbDynamicRatioPreferredMethod, servers, monitorResults, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology, lbDynamicRatioAlternateMethod, eas, autoConsolidatedMonitors, availability, consolidatedMonitors, ttl, useTTL, disable, quorum)
	ref, err := objMgr.connector.CreateObject(poolDtc)
	if err != nil {
		return nil, err
	}
	poolDtc.Ref = ref
	return poolDtc, nil
}

func (objMgr *ObjectManager) GetAllDtcPool(queryParams *QueryParams) ([]DtcPool, error) {
	var res []DtcPool
	pool := NewEmptyDtcPool()
	err := objMgr.connector.GetObject(pool, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("error getting Dtc Pool object, err: %s", err)
	}
	return res, nil
}

func (objMgr *ObjectManager) UpdateDtcPool(
	ref string,
	comment string,
	name string,
	lbPreferredMethod string,
	lbDynamicRatioPreferred map[string]interface{},
	servers []*DtcServerLink,
	monitors []Monitor,
	lbPreferredTopology *string,
	lbAlternateMethod string,
	lbAlternateTopology *string,
	lbDynamicRatioAlternate map[string]interface{},
	eas EA,
	autoConsolidatedMonitors bool,
	availability string,
	userMonitors []map[string]interface{},
	ttl uint32,
	useTTL bool,
	disable bool,
	quorum uint32) (*DtcPool, error) {
	if lbPreferredMethod == "DYNAMIC_RATIO" && lbDynamicRatioPreferred == nil {
		return nil, fmt.Errorf("LbDynamicRatioPreferred cannot be nil when the preferred load balancing method is set to DYNAMIC_RATIO")
	}
	if lbPreferredMethod == "TOPOLOGY" && lbPreferredTopology == nil {
		return nil, fmt.Errorf("preferred topology cannot be nil when preferred load balancing method is set to TOPOLOGY")
	}
	if autoConsolidatedMonitors && len(userMonitors) > 0 {
		return nil, fmt.Errorf("either AutoConsolidatedMonitors or ConsolidatedMonitors should be set.")
	}
	//update servers with server references
	err := updateServerReferences(servers, objMgr)
	if err != nil {
		return nil, err
	}
	// Convert LbDynamicRatioPreferred to use monitor reference
	var lbDynamicRatioPreferredMethod *SettingDynamicratio
	if lbDynamicRatioPreferred != nil {
		monitor, _ := lbDynamicRatioPreferred["monitor"].(Monitor)
		method, _ := lbDynamicRatioPreferred["method"].(string)
		monitorMetric, _ := lbDynamicRatioPreferred["monitor_metric"].(string)
		monitorWeighing, _ := lbDynamicRatioPreferred["monitor_weighing"].(string)
		invertMonitorMetric, _ := lbDynamicRatioPreferred["invert_monitor_metric"].(bool)

		monitorRef, err := getMonitorReference(monitor.Name, monitor.Type, objMgr)
		if err != nil {
			return nil, err
		}
		lbDynamicRatioPreferredMethod = &SettingDynamicratio{
			Method:              method,
			Monitor:             monitorRef,
			MonitorMetric:       monitorMetric,
			MonitorWeighing:     monitorWeighing,
			InvertMonitorMetric: invertMonitorMetric,
		}
	} else {
		lbDynamicRatioPreferredMethod = nil
	}
	// Convert monitor names to monitor references
	var monitorResults []*DtcMonitorHttp
	for _, monitor := range monitors {
		monitorRef, err := getMonitorReference(monitor.Name, monitor.Type, objMgr)
		if err != nil {
			return nil, err
		}
		monitorResults = append(monitorResults, &DtcMonitorHttp{Ref: monitorRef})
	}
	//Update the topology name with the topology reference
	if lbPreferredTopology != nil {
		topology, err := getTopology(*lbPreferredTopology, objMgr)
		if err != nil {
			return nil, err
		}
		lbPreferredTopology = &topology
	}
	//Update the topology name with the topology reference
	if lbAlternateTopology != nil {
		topologyAlternate, err := getTopology(*lbAlternateTopology, objMgr)
		if err != nil {
			return nil, err
		}
		lbAlternateTopology = &topologyAlternate
	}
	//Convert LbDynamicRatioAlternate to use monitor reference
	var lbDynamicRatioAlternateMethod *SettingDynamicratio
	if lbDynamicRatioAlternate != nil {
		monitorAlternate, _ := lbDynamicRatioAlternate["monitor"].(Monitor)
		methodAlternate, _ := lbDynamicRatioAlternate["method"].(string)
		monitorMetricAlternate, _ := lbDynamicRatioAlternate["monitor_metric"].(string)
		monitorWeighingAlternate, _ := lbDynamicRatioAlternate["monitor_weighing"].(string)
		invertMonitorMetricAlternate, _ := lbDynamicRatioAlternate["invert_monitor_metric"].(bool)

		monitorRefAlternate, err := getMonitorReference(monitorAlternate.Name, monitorAlternate.Type, objMgr)
		if err != nil {
			return nil, err
		}
		lbDynamicRatioAlternateMethod = &SettingDynamicratio{
			Method:              methodAlternate,
			Monitor:             monitorRefAlternate,
			MonitorMetric:       monitorMetricAlternate,
			MonitorWeighing:     monitorWeighingAlternate,
			InvertMonitorMetric: invertMonitorMetricAlternate,
		}
	} else {
		lbDynamicRatioAlternateMethod = nil
	}
	//processing user input to retrieve monitor references and creating a slice of *DtcPoolConsolidatedMonitorHealth structs with updated monitor references.
	var consolidatedMonitors []*DtcPoolConsolidatedMonitorHealth
	if userMonitors != nil {
		if len(userMonitors) == 0 {
			consolidatedMonitors = []*DtcPoolConsolidatedMonitorHealth{}
		} else {
			for _, userMonitor := range userMonitors {
				monitor, okMonitor := userMonitor["monitor"].(Monitor)
				monitorAvailability, okAvail := userMonitor["availability"].(string)
				fullHealthComm, _ := userMonitor["full_health_communication"].(bool)
				members, okMember := userMonitor["members"].([]string)
				if !okMonitor {
					return nil, fmt.Errorf("required field missing: monitor")
				}

				if !okAvail {
					return nil, fmt.Errorf("required field missing: availability")
				}

				if !okMember {
					return nil, fmt.Errorf("required field missing: members")
				}
				monitorRef, err := getMonitorReference(monitor.Name, monitor.Type, objMgr)
				if err != nil {
					return nil, err
				}

				consolidatedMonitor := &DtcPoolConsolidatedMonitorHealth{
					Members:                 members,
					Monitor:                 monitorRef,
					Availability:            monitorAvailability,
					FullHealthCommunication: fullHealthComm,
				}
				consolidatedMonitors = append(consolidatedMonitors, consolidatedMonitor)
			}
		}
	}

	poolDtc := NewDtcPool(comment, name, lbPreferredMethod, lbDynamicRatioPreferredMethod, servers, monitorResults, lbPreferredTopology, lbAlternateMethod, lbAlternateTopology, lbDynamicRatioAlternateMethod, eas, autoConsolidatedMonitors, availability, consolidatedMonitors, ttl, useTTL, disable, quorum)
	poolDtc.Ref = ref
	reference, err := objMgr.connector.UpdateObject(poolDtc, ref)
	if err != nil {
		return nil, err
	}
	poolDtc.Ref = reference

	poolDtc, err = objMgr.GetDtcPoolByRef(reference)
	if err != nil {
		return nil, err
	}

	return poolDtc, nil

}

func (objMgr *ObjectManager) GetDtcPoolByRef(ref string) (*DtcPool, error) {
	poolDtc := NewEmptyDtcPool()
	err := objMgr.connector.GetObject(
		poolDtc, ref, NewQueryParams(false, nil), &poolDtc)
	return poolDtc, err
}

func (objMgr *ObjectManager) GetDtcPool(name string) (*DtcPool, error) {
	dtcPool := NewEmptyDtcPool()
	var res []DtcPool
	if name == "" {
		return nil, fmt.Errorf("name of the record is required to retrieve a unique Dtc Pool record")
	}
	sf := map[string]string{
		"name": name,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(dtcPool, "", queryParams, &res)
	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf("Dtc Pool record with name '%s' not found", name))
	}
	return &res[0], nil
}

func (objMgr *ObjectManager) DeleteDtcPool(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

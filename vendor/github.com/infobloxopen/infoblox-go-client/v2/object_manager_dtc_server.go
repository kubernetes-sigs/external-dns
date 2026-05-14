package ibclient

import (
	"encoding/json"
	"fmt"
)

func (d *DtcServer) MarshalJSON() ([]byte, error) {
	type Alias DtcServer
	aux := &struct {
		Monitors []*DtcServerMonitor `json:"monitors"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	if len(d.Monitors) == 0 {
		aux.Monitors = []*DtcServerMonitor{}
	} else {
		aux.Monitors = d.Monitors
	}

	return json.Marshal(aux)
}

func NewEmptyDtcServer() *DtcServer {
	dtcServer := &DtcServer{}
	dtcServer.SetReturnFields(append(dtcServer.ReturnFields(), "extattrs", "auto_create_host_record", "disable", "health", "monitors", "sni_hostname", "use_sni_hostname"))
	return dtcServer
}

func NewDtcServer(comment string,
	name string,
	host string,
	autoCreateHostRecord bool,
	disable bool,
	ea EA,
	monitors []*DtcServerMonitor,
	sniHostname string,
	useSniHostname bool,
) *DtcServer {
	DtcServer := NewEmptyDtcServer()
	DtcServer.Comment = &comment
	DtcServer.Name = &name
	DtcServer.Host = &host
	DtcServer.AutoCreateHostRecord = &autoCreateHostRecord
	DtcServer.Disable = &disable
	DtcServer.Ea = ea
	DtcServer.Monitors = monitors
	DtcServer.SniHostname = &sniHostname
	DtcServer.UseSniHostname = &useSniHostname
	return DtcServer
}

func (objMgr *ObjectManager) CreateDtcServer(
	comment string,
	name string,
	host string,
	autoCreateHostRecord bool,
	disable bool,
	ea EA,
	monitors []map[string]interface{},
	sniHostname string,
	useSniHostname bool,
) (*DtcServer, error) {

	if name == "" || host == "" {
		return nil, fmt.Errorf("name and host fields are required to create a Dtc Server object")
	}
	if (useSniHostname && sniHostname == "") || (!useSniHostname && sniHostname != "") {
		return nil, fmt.Errorf("'sni_hostname' must be provided when 'use_sni_hostname' is enabled, " +
			"and 'use_sni_hostname' must be enabled if 'sni_hostname' is provided")
	}
	var serverMonitors []*DtcServerMonitor
	for _, userMonitor := range monitors {
		monitor, okMonitor := userMonitor["monitor"].(Monitor)
		monitorHost, _ := userMonitor["host"].(string)
		if !okMonitor {
			return nil, fmt.Errorf("required field missing: monitor")
		}
		monitorRef, err := getMonitorReference(monitor.Name, monitor.Type, objMgr)
		if err != nil {
			return nil, err
		}

		serverMonitor := &DtcServerMonitor{
			Monitor: monitorRef,
			Host:    monitorHost,
		}

		serverMonitors = append(serverMonitors, serverMonitor)
	}
	dtcServer := NewDtcServer(comment, name, host, autoCreateHostRecord, disable, ea, serverMonitors, sniHostname, useSniHostname)
	ref, err := objMgr.connector.CreateObject(dtcServer)
	if err != nil {
		return nil, err
	}
	dtcServer.Ref = ref
	return dtcServer, nil
}

func (objMgr *ObjectManager) GetAllDtcServer(queryParams *QueryParams) ([]DtcServer, error) {
	var res []DtcServer
	server := NewEmptyDtcServer()
	err := objMgr.connector.GetObject(server, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("error getting Dtc Server object, err: %s", err)
	}
	return res, nil
}

func (objMgr *ObjectManager) GetDtcServer(name string, host string) (*DtcServer, error) {
	var res []DtcServer
	server := NewEmptyDtcServer()
	if name == "" || host == "" {
		return nil, fmt.Errorf("name and host of the server are required to retreive a unique dtc server")
	}
	sf := map[string]string{
		"name": name,
		"host": host,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(server, "", queryParams, &res)
	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf("Dtc server with name '%s' and host '%s' not found", name, host))
	}
	return &res[0], nil
}

func (objMgr *ObjectManager) UpdateDtcServer(
	ref string,
	comment string,
	name string,
	host string,
	autoCreateHostRecord bool,
	disable bool,
	ea EA,
	monitors []map[string]interface{},
	sniHostname string,
	useSniHostname bool) (*DtcServer, error) {
	if (useSniHostname && sniHostname == "") || (!useSniHostname && sniHostname != "") {
		return nil, fmt.Errorf("'sni_hostname' must be provided when 'use_sni_hostname' is enabled, " +
			"and 'use_sni_hostname' must be enabled if 'sni_hostname' is provided")
	}
	var serverMonitors []*DtcServerMonitor
	for _, userMonitor := range monitors {
		monitor, okMonitor := userMonitor["monitor"].(Monitor)
		monitorHost, _ := userMonitor["host"].(string)
		if !okMonitor {
			return nil, fmt.Errorf("required field missing: monitor")
		}
		monitorRef, err := getMonitorReference(monitor.Name, monitor.Type, objMgr)
		if err != nil {
			return nil, err
		}

		serverMonitor := &DtcServerMonitor{
			Monitor: monitorRef,
			Host:    monitorHost,
		}

		serverMonitors = append(serverMonitors, serverMonitor)
	}
	dtcServer := NewDtcServer(comment, name, host, autoCreateHostRecord, disable, ea, serverMonitors, sniHostname, useSniHostname)
	dtcServer.Ref = ref
	ref, err := objMgr.connector.UpdateObject(dtcServer, ref)
	if err != nil {
		return nil, err
	}
	dtcServer.Ref = ref
	return dtcServer, nil
}

func (objMgr *ObjectManager) GetDtcServerByRef(ref string) (*DtcServer, error) {
	serverDtc := NewEmptyDtcServer()
	err := objMgr.connector.GetObject(
		serverDtc, ref, NewQueryParams(false, nil), &serverDtc)
	return serverDtc, err
}

func (objMgr *ObjectManager) DeleteDtcServer(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

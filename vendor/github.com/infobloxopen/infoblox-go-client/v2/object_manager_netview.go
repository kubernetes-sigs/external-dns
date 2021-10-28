package ibclient

import (
	"fmt"
	"strings"
)

func (objMgr *ObjectManager) CreateNetworkView(name string, comment string, setEas EA) (*NetworkView, error) {
	networkView := NewNetworkView(name, comment, setEas, "")

	ref, err := objMgr.connector.CreateObject(networkView)
	networkView.Ref = ref

	return networkView, err
}

func (objMgr *ObjectManager) makeNetworkView(netviewName string) (netviewRef string, err error) {
	var netviewObj *NetworkView
	if netviewObj, err = objMgr.GetNetworkView(netviewName); err != nil {
		return
	}
	if netviewObj == nil {
		if netviewObj, err = objMgr.CreateNetworkView(netviewName, "", nil); err != nil {
			return
		}
	}

	netviewRef = netviewObj.Ref

	return
}

func (objMgr *ObjectManager) CreateDefaultNetviews(globalNetview string, localNetview string) (globalNetviewRef string, localNetviewRef string, err error) {
	if globalNetviewRef, err = objMgr.makeNetworkView(globalNetview); err != nil {
		return
	}

	if localNetviewRef, err = objMgr.makeNetworkView(localNetview); err != nil {
		return
	}

	return
}

func (objMgr *ObjectManager) GetNetworkView(name string) (*NetworkView, error) {
	var res []NetworkView

	netview := NewEmptyNetworkView()
	sf := map[string]string{
		"name": name,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(netview, "", queryParams, &res)

	if err != nil {
		return nil, err
	}
	if res == nil || len(res) == 0 {
		return nil, fmt.Errorf("network view '%s' not found", name)
	}

	return &res[0], nil
}

func (objMgr *ObjectManager) GetNetworkViewByRef(ref string) (*NetworkView, error) {
	res := NewEmptyNetworkView()
	queryParams := NewQueryParams(false, nil)
	if err := objMgr.connector.GetObject(res, ref, queryParams, &res); err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("network view not found")
	}

	return res, nil
}

func (objMgr *ObjectManager) UpdateNetworkView(ref string, name string, comment string, setEas EA) (*NetworkView, error) {

	nv := NewEmptyNetworkView()

	err := objMgr.connector.GetObject(
		nv, ref, NewQueryParams(false, nil), nv)
	if err != nil {
		return nil, err
	}
	cleanName := strings.TrimSpace(name)
	if cleanName != "" {
		nv.Name = cleanName
	}
	nv.Comment = comment
	nv.Ea = setEas

	updatedRef, err := objMgr.connector.UpdateObject(nv, ref)
	nv.Ref = updatedRef

	return nv, err
}

func (objMgr *ObjectManager) DeleteNetworkView(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

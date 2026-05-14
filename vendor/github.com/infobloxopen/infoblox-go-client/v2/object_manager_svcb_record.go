package ibclient

import "fmt"

func (objMgr *ObjectManager) CreateSVCBRecord(name string, priority uint32, targetName string, comment string,
	creator string, ddnsPrincipal string, ddnsProtected bool, disable bool, ea EA, forbidReclamation bool,
	svcParams []SVCParams, ttl uint32, useTtl bool, view string) (*RecordSVCB, error) {
	if name == "" || targetName == "" {
		return nil, fmt.Errorf("name and target name fields are required to create a SVCB Record")
	}
	if priority > 65535 {
		return nil, fmt.Errorf("priority must be between 0 and 65535")
	}
	recordSVCB := NewSVCBRecord("", name, priority, targetName, comment, creator, ddnsPrincipal, ddnsProtected, disable, ea,
		forbidReclamation, svcParams, ttl, useTtl)
	recordSVCB.View = view
	ref, err := objMgr.connector.CreateObject(recordSVCB)
	if err != nil {
		return nil, fmt.Errorf("error creating SVCB Record %s, err: %s", name, err)
	}
	recordSVCB.Ref = ref
	return recordSVCB, nil
}

func (objMgr *ObjectManager) DeleteSVCBRecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func (objMgr *ObjectManager) GetAllSVCBRecords(queryParams *QueryParams) ([]RecordSVCB, error) {
	var res []RecordSVCB
	recordSVCB := NewEmptyRecordSVCB()
	err := objMgr.connector.GetObject(recordSVCB, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("failed getting SVCB Record: %s", err)
	}
	return res, nil
}

func (objMgr *ObjectManager) GetSVCBRecordByRef(ref string) (*RecordSVCB, error) {
	recordSVCB := NewEmptyRecordSVCB()
	err := objMgr.connector.GetObject(recordSVCB, ref, NewQueryParams(false, nil), &recordSVCB)
	if err != nil {
		return nil, err
	}
	return recordSVCB, nil
}

func (objMgr *ObjectManager) UpdateSVCBRecord(ref string, name string, priority uint32, targetName string, comment string,
	creator string, ddnsPrincipal string, ddnsProtected bool, disable bool, ea EA, forbidReclamation bool,
	svcParams []SVCParams, ttl uint32, useTtl bool) (*RecordSVCB, error) {
	if name == "" || targetName == "" {
		return nil, fmt.Errorf("name and target name fields are required for SVCB Record")
	}
	if priority > 65535 {
		return nil, fmt.Errorf("priority must be between 0 and 65535")
	}
	recordSVCB := NewSVCBRecord(ref, name, priority, targetName, comment, creator, ddnsPrincipal, ddnsProtected, disable, ea,
		forbidReclamation, svcParams, ttl, useTtl)
	newRef, err := objMgr.connector.UpdateObject(recordSVCB, ref)
	if err != nil {
		return nil, fmt.Errorf("error updating SVCB Record %s, err: %s", name, err)
	}
	recordSVCB.Ref = newRef
	recordSVCB, err = objMgr.GetSVCBRecordByRef(newRef)
	if err != nil {
		return nil, fmt.Errorf("error getting updated SVCB Record %s, err: %s", name, err)
	}
	return recordSVCB, nil
}

func NewEmptyRecordSVCB() *RecordSVCB {
	recordSVCB := RecordSVCB{}
	recordSVCB.SetReturnFields(append(recordSVCB.ReturnFields(), "aws_rte53_record_info", "cloud_info", "comment", "creation_time", "creator", "ddns_principal", "ddns_protected", "disable", "extattrs", "forbid_reclamation", "last_queried", "reclaimable", "svc_parameters", "ttl", "use_ttl", "zone"))
	return &recordSVCB
}

func NewSVCBRecord(ref string, name string, priority uint32, targetName string, comment string,
	creator string, ddnsPrincipal string, ddnsProtected bool, disable bool, ea EA, forbidReclamation bool,
	svcParams []SVCParams, ttl uint32, useTtl bool) *RecordSVCB {
	recordSVCB := NewEmptyRecordSVCB()
	recordSVCB.Ref = ref
	recordSVCB.Name = name
	recordSVCB.Comment = comment
	recordSVCB.Disable = disable
	recordSVCB.Ea = ea
	recordSVCB.Priority = priority
	recordSVCB.SvcParameters = svcParams
	recordSVCB.TargetName = targetName
	recordSVCB.Creator = creator
	recordSVCB.DdnsProtected = ddnsProtected
	recordSVCB.DdnsPrincipal = ddnsPrincipal
	recordSVCB.UseTtl = useTtl
	recordSVCB.Ttl = ttl
	recordSVCB.ForbidReclamation = forbidReclamation
	return recordSVCB
}

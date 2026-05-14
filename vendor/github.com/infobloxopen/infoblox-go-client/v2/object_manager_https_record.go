package ibclient

import "fmt"

func NewEmptyHttpsRecord() *RecordHttps {
	newRecordHttps := &RecordHttps{}
	newRecordHttps.SetReturnFields(append(newRecordHttps.ReturnFields(),"zone", "comment", "svc_parameters", "disable", "extattrs", "forbid_reclamation", "ttl", "use_ttl", "creator", "ddns_principal", "ddns_protected","reclaimable","last_queried","creation_time","aws_rte53_record_info","cloud_info"))
	return newRecordHttps
}
func NewHttpsRecord(
	name string,
	priority uint32,
	targetName string,
	comment string,
	creator string,
	ddnsPrincipal string,
	ddnsProtected bool,
	svcParameters []SVCParams,
	disable bool,
	extAttrs EA,
	forbidReclamation bool,
	ttl uint32,
	useTtl bool,
	view string,
	ref string) *RecordHttps {

	res := NewEmptyHttpsRecord()

	res.Name = name
	res.Priority = priority
	res.TargetName = targetName
	res.Comment = comment
	res.Creator = creator
	res.DdnsPrincipal = ddnsPrincipal
	res.DdnsProtected = ddnsProtected
	res.SvcParameters = svcParameters
	res.Disable = disable
	res.Ea = extAttrs
	res.ForbidReclamation = forbidReclamation
	res.Ttl = ttl
	res.UseTtl = useTtl
	res.View = view
	res.Ref = ref
	return res
}

func (obj *ObjectManager) CreateHTTPSRecord(name string, priority uint32, targetName string, comment string, creator string, ddnsPrincipal string, ddnsProtected bool , disable bool, ea EA, forbidReclamation bool, svcParams []SVCParams, ttl uint32, useTtl bool, view string) (*RecordHttps, error) {

	if priority > 65535 {
		return nil, fmt.Errorf("priority must be between 0 and 65535")
	}

	if name == "" || targetName == "" {
		return nil, fmt.Errorf("name and targetName are required to create HTTPS Record")
	}

	recordHttps := NewHttpsRecord(name, priority, targetName, comment, creator, ddnsPrincipal, ddnsProtected,svcParams, disable, ea , forbidReclamation, ttl, useTtl, view, "")
	ref, err := obj.connector.CreateObject(recordHttps)
	if err != nil {
		return nil, err
	}
	recordHttps.Ref = ref
	return recordHttps, nil
}

func (objMgr *ObjectManager) GetHTTPSRecordByRef(ref string) (*RecordHttps, error) {
	recordHTTPS := NewEmptyHttpsRecord()
	err := objMgr.connector.GetObject(recordHTTPS, ref, NewQueryParams(false, nil), &recordHTTPS)
	if err != nil {
		return nil, err
	}
	return recordHTTPS, nil
}

func (objMgr *ObjectManager) GetAllHTTPSRecord(queryParams *QueryParams) ([]RecordHttps, error) {
	var res []RecordHttps
	recordHttps := NewEmptyHttpsRecord()
	err := objMgr.connector.GetObject(recordHttps, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("failed getting HTTPS Record: %s", err)
	}
	return res, nil
}

func (objMgr *ObjectManager) UpdateHTTPSRecord(ref string, name string, priority uint32, targetName string, comment string, creator string, ddnsPrincipal string, ddnsProtected bool, disable bool, ea EA, forbidReclamation bool, svcParams []SVCParams, ttl uint32, useTtl bool) (*RecordHttps, error) {
	if priority > 65535 {
		return nil, fmt.Errorf("priority must be between 0 and 65535")
	}
	if name == "" || targetName == "" {
		return nil, fmt.Errorf("name and targetName cannot be empty")
	}
	httpsRecord := NewHttpsRecord(name, priority, targetName, comment, creator, ddnsPrincipal, ddnsProtected, svcParams, disable, ea, forbidReclamation, ttl, useTtl, "", ref)
	updatedRef, err := objMgr.connector.UpdateObject(httpsRecord, ref)
	if err != nil {
		return nil, err
	}
	httpsRecord.Ref = updatedRef
	httpsRecord, err = objMgr.GetHTTPSRecordByRef(updatedRef)
	if err != nil {
		return nil, fmt.Errorf("error getting updated HTTPS Record %s, err: %s", name, err)
	}
	return httpsRecord, nil
}

func (objMgr *ObjectManager) DeleteHTTPSRecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

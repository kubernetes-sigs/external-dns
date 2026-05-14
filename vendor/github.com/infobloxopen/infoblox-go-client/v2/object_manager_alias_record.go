package ibclient

import "fmt"

func NewEmptyAliasRecord() *RecordAlias {
	aliasRecord := &RecordAlias{}
	aliasRecord.SetReturnFields(append(aliasRecord.ReturnFields(), "extattrs", "cloud_info", "comment", "disable", "dns_name", "dns_target_name", "creator", "ttl", "use_ttl", "view", "zone"))
	return aliasRecord
}

func (objMgr *ObjectManager) CreateAliasRecord(name string, dnsView string, targetName string, targetType string, comment string, disable bool, ea EA, ttl uint32, useTtl bool) (*RecordAlias, error) {
	if name == "" || targetName == "" || targetType == "" {
		return nil, fmt.Errorf("name, targetName and targetType are required to create an Alias Record")
	}
	if dnsView == "" {
		dnsView = "default"
	}
	aliasRecord := NewAliasRecord(name, dnsView, targetName, targetType, comment, disable, ea, ttl, useTtl)
	ref, err := objMgr.connector.CreateObject(aliasRecord)
	if err != nil {
		return nil, err
	}

	aliasRecord.Ref = ref
	return aliasRecord, nil
}

func (objMgr *ObjectManager) GetAliasRecordByRef(ref string) (*RecordAlias, error) {
	aliasRecord := NewEmptyAliasRecord()
	err := objMgr.connector.GetObject(aliasRecord, ref, NewQueryParams(false, nil), &aliasRecord)
	if err != nil {
		return nil, err
	}
	return aliasRecord, nil
}

func (objMgr *ObjectManager) GetAllAliasRecord(queryParams *QueryParams) ([]RecordAlias, error) {
	var res []RecordAlias
	aliasRecord := NewEmptyAliasRecord()
	err := objMgr.connector.GetObject(aliasRecord, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("failed getting Alias Record: %s", err)
	}
	return res, nil
}

func (objMgr *ObjectManager) DeleteAliasRecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func (objMgr *ObjectManager) UpdateAliasRecord(ref string, name string, dnsView string, targetName string, targetType string, comment string, disable bool, ea EA, ttl uint32, useTtl bool) (*RecordAlias, error) {
	if name == "" || targetName == "" || targetType == "" {
		return nil, fmt.Errorf("name, targetName and targetType are required to create an Alias Record")
	}
	if dnsView == "" {
		dnsView = "default"
	}
	aliasRecord := NewAliasRecord(name, dnsView, targetName, targetType, comment, disable, ea, ttl, useTtl)
	updatedRef, err := objMgr.connector.UpdateObject(aliasRecord, ref)
	if err != nil {
		return nil, err
	}
	aliasRecord.Ref = updatedRef
	return aliasRecord, nil
}

func NewAliasRecord(name string, dnsView string, targetName string, targetType string, comment string, disable bool, ea EA, ttl uint32, useTtl bool) *RecordAlias {
	recordAlias := NewEmptyAliasRecord()
	recordAlias.Name = &name
	recordAlias.View = &dnsView
	recordAlias.TargetName = &targetName
	recordAlias.TargetType = targetType
	recordAlias.Comment = &comment
	recordAlias.Disable = &disable
	recordAlias.Ea = ea
	recordAlias.Ttl = &ttl
	recordAlias.UseTtl = &useTtl
	return recordAlias
}

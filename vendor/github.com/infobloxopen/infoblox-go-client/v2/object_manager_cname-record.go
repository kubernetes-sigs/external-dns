package ibclient

import "fmt"

func (objMgr *ObjectManager) CreateCNAMERecord(
	dnsview string,
	canonical string,
	recordname string,
	useTtl bool,
	ttl uint32,
	comment string,
	eas EA) (*RecordCNAME, error) {

	if canonical == "" || recordname == "" {
		return nil, fmt.Errorf("canonical name and record name fields are required to create a CNAME record")
	}
	recordCNAME := NewRecordCNAME(dnsview, canonical, recordname, useTtl, ttl, comment, eas, "")

	ref, err := objMgr.connector.CreateObject(recordCNAME)
	if err != nil {
		return nil, err
	}
	recordCNAME, err = objMgr.GetCNAMERecordByRef(ref)
	if err != nil {
		return nil, err
	}
	return recordCNAME, err
}

func (objMgr *ObjectManager) GetCNAMERecord(dnsview string, canonical string, recordName string) (*RecordCNAME, error) {
	var res []RecordCNAME
	recordCNAME := NewEmptyRecordCNAME()
	if dnsview == "" || canonical == "" || recordName == "" {
		return nil, fmt.Errorf("DNS view, canonical name and record name of the record are required to retreive a unique CNAME record")
	}
	sf := map[string]string{
		"view":      dnsview,
		"canonical": canonical,
		"name":      recordName,
	}

	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(recordCNAME, "", queryParams, &res)

	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf(
				"CNAME record with name '%s' and canonical name '%s' in DNS view '%s' is not found",
				recordName, canonical, dnsview))
	}
	return &res[0], nil
}

func (objMgr *ObjectManager) GetCNAMERecordByRef(ref string) (*RecordCNAME, error) {
	recordCNAME := NewEmptyRecordCNAME()
	err := objMgr.connector.GetObject(
		recordCNAME, ref, NewQueryParams(false, nil), &recordCNAME)
	return recordCNAME, err
}

func (objMgr *ObjectManager) DeleteCNAMERecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func (objMgr *ObjectManager) UpdateCNAMERecord(
	ref string,
	canonical string,
	recordName string,
	useTtl bool,
	ttl uint32,
	comment string,
	setEas EA) (*RecordCNAME, error) {

	recordCNAME := NewRecordCNAME("", canonical, recordName, useTtl, ttl, comment, setEas, ref)
	updatedRef, err := objMgr.connector.UpdateObject(recordCNAME, ref)
	if err != nil {
		return nil, err
	}

	recordCNAME, err = objMgr.GetCNAMERecordByRef(updatedRef)
	return recordCNAME, err
}

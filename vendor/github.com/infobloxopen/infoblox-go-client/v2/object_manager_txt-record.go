package ibclient

import "fmt"

// Creates TXT Record. Use TTL of 0 to inherit TTL from the Zone
<<<<<<< HEAD
func (objMgr *ObjectManager) CreateTXTRecord(recordname string, text string, ttl uint, dnsview string) (*RecordTXT, error) {

	recordTXT := NewRecordTXT(RecordTXT{
		View: dnsview,
		Name: recordname,
		Text: text,
		Ttl:  ttl,
	})

	ref, err := objMgr.connector.CreateObject(recordTXT)
	recordTXT.Ref = ref
	return recordTXT, err
}

func (objMgr *ObjectManager) GetTXTRecordByRef(ref string) (*RecordTXT, error) {
	recordTXT := NewRecordTXT(RecordTXT{})
	err := objMgr.connector.GetObject(
		recordTXT, ref, NewQueryParams(false, nil), &recordTXT)
	return recordTXT, err
}

func (objMgr *ObjectManager) GetTXTRecord(name string) (*RecordTXT, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}
	var res []RecordTXT

	recordTXT := NewRecordTXT(RecordTXT{})

	sf := map[string]string{
		"name": name,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(recordTXT, "", queryParams, &res)

	if err != nil || res == nil || len(res) == 0 {
		return nil, err
	}

	return &res[0], nil
}

func (objMgr *ObjectManager) UpdateTXTRecord(recordname string, text string) (*RecordTXT, error) {
	var res []RecordTXT

	recordTXT := NewRecordTXT(RecordTXT{Name: recordname})

	sf := map[string]string{
		"name": recordname,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(recordTXT, "", queryParams, &res)

	if len(res) == 0 {
		return nil, nil
	}

	res[0].Text = text

	res[0].Zone = "" //  set the Zone value to "" as its a non writable field

	_, err = objMgr.connector.UpdateObject(&res[0], res[0].Ref)

	if err != nil || res == nil || len(res) == 0 {
		return nil, err
	}

	return &res[0], nil
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
func (objMgr *ObjectManager) CreateTXTRecord(
	dnsView string,
	recordName string,
	text string,
	ttl uint32,
	useTtl bool,
	comment string,
	eas EA) (*RecordTXT, error) {

	recordTXT := NewRecordTXT(dnsView, "", recordName, text, ttl, useTtl, comment, eas)

	ref, err := objMgr.connector.CreateObject(recordTXT)
	if err != nil {
		return nil, err
	}
	recordTXT, err = objMgr.GetTXTRecordByRef(ref)
	return recordTXT, err
}

func (objMgr *ObjectManager) GetTXTRecordByRef(ref string) (*RecordTXT, error) {
	recordTXT := NewEmptyRecordTXT()
	err := objMgr.connector.GetObject(
		recordTXT, ref, NewQueryParams(false, nil), &recordTXT)
	return recordTXT, err
}

func (objMgr *ObjectManager) GetTXTRecord(dnsview string, name string) (*RecordTXT, error) {
	if dnsview == "" || name == "" {
		return nil, fmt.Errorf("DNS view and name are required to retrieve a unique txt record")
	}
	var res []RecordTXT

	recordTXT := NewEmptyRecordTXT()

	sf := map[string]string{
		"view": dnsview,
		"name": name,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(recordTXT, "", queryParams, &res)

	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(
			fmt.Sprintf(
				"TXT record with name '%s' in DNS view '%s' is not found",
				name, dnsview))
	}
	return &res[0], nil
}

func (objMgr *ObjectManager) UpdateTXTRecord(
	ref string,
	recordName string,
	text string,
	ttl uint32,
	useTtl bool,
	comment string,
	eas EA) (*RecordTXT, error) {

	recordTXT := NewRecordTXT("", "", recordName, text, ttl, useTtl, comment, eas)
	recordTXT.Ref = ref

	reference, err := objMgr.connector.UpdateObject(recordTXT, ref)
	if err != nil {
		return nil, err
	}

	recordTXT, err = objMgr.GetTXTRecordByRef(reference)
	return recordTXT, err
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}

func (objMgr *ObjectManager) DeleteTXTRecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

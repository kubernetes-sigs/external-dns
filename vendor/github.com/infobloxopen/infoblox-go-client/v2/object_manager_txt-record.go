package ibclient

import "fmt"

// Creates TXT Record. Use TTL of 0 to inherit TTL from the Zone
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
}

func (objMgr *ObjectManager) DeleteTXTRecord(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

package ibclient

func (objMgr *ObjectManager) CreateEADefinition(eadef EADefinition) (*EADefinition, error) {
	newEadef := NewEADefinition(eadef)

	ref, err := objMgr.connector.CreateObject(newEadef)
	newEadef.Ref = ref

	return newEadef, err
}

func (objMgr *ObjectManager) GetEADefinition(name string) (*EADefinition, error) {
	var res []EADefinition

	eadef := NewEADefinition(EADefinition{Name: name})

	sf := map[string]string{
		"name": name,
	}
	queryParams := NewQueryParams(false, sf)
	err := objMgr.connector.GetObject(eadef, "", queryParams, &res)

	if err != nil || res == nil || len(res) == 0 {
		return nil, err
	}

	return &res[0], nil
}

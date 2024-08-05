package ibclient

import "fmt"

func (objMgr *ObjectManager) GetDNSView(name string) (*View, error) {
	var res []View
	if name == "" {
		return nil, fmt.Errorf(
			"DNS view's name is required to retreive DNS view object")
	}
	queryParams := NewQueryParams(false, map[string]string{"name": name})
	err := objMgr.connector.GetObject(NewEmptyDNSView(), "", queryParams, &res)
	if err != nil {
		return nil, err
	} else if res == nil || len(res) == 0 {
		return nil, NewNotFoundError(fmt.Sprintf("DNS view with name '%s' not found", name))
	}

	return &res[0], nil
}

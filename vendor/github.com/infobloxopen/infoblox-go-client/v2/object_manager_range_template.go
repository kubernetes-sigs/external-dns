package ibclient

import (
	"encoding/json"
	"fmt"
)

func (d Rangetemplate) MarshalJSON() ([]byte, error) {
	type Alias Rangetemplate
	aux := &struct {
		Member *Dhcpmember `json:"member"`
		*Alias
	}{
		Member: d.Member,
		Alias:  (*Alias)(&d),
	}
	return json.Marshal(aux)
}

func (d *Rangetemplate) UnmarshalJSON(data []byte) error {
	type Alias Rangetemplate
	aux := &struct {
		Member *Dhcpmember `json:"member"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.Member = aux.Member
	return nil
}

func (ms Msdhcpserver) MarshalJSON() ([]byte, error) {
	if ms.Ipv4Addr == "" {
		return []byte("null"), nil
	}
	return json.Marshal(map[string]string{
		"ipv4addr": ms.Ipv4Addr,
	})
}

func (objMgr *ObjectManager) CreateRangeTemplate(name string, numberOfAdresses uint32, offset uint32, comment string, ea EA,
	options []*Dhcpoption, useOption bool, serverAssociationType string, failOverAssociation string, member *Dhcpmember, cloudApiCompatible bool, msServer string) (*Rangetemplate, error) {
	if name == "" {
		return nil, fmt.Errorf("name field is required to create a Range Template object")
	}
	rangeTemplate := NewRangeTemplate("", name, numberOfAdresses, offset, comment, ea, options,
		useOption, serverAssociationType, failOverAssociation, member, cloudApiCompatible, msServer)
	ref, err := objMgr.connector.CreateObject(rangeTemplate)
	if err != nil {
		return nil, fmt.Errorf("error creating Range Template object %s, err: %s", name, err)
	}
	rangeTemplate.Ref = ref
	return rangeTemplate, nil
}

func (objMgr *ObjectManager) DeleteRangeTemplate(ref string) (string, error) {
	return objMgr.connector.DeleteObject(ref)
}

func (objMgr *ObjectManager) GetAllRangeTemplate(queryParams *QueryParams) ([]Rangetemplate, error) {
	var res []Rangetemplate
	rangeTemplate := NewEmptyRangeTemplate()
	err := objMgr.connector.GetObject(rangeTemplate, "", queryParams, &res)
	if err != nil {
		return nil, fmt.Errorf("failed getting Range Template Record: %s", err)
	}
	return res, nil
}

func (objMgr *ObjectManager) GetRangeTemplateByRef(ref string) (*Rangetemplate, error) {
	rangeTemplate := NewEmptyRangeTemplate()
	err := objMgr.connector.GetObject(rangeTemplate, ref, NewQueryParams(false, nil), &rangeTemplate)
	if err != nil {
		return nil, err
	}
	return rangeTemplate, nil
}

func (objMgr *ObjectManager) UpdateRangeTemplate(ref string, name string, numberOfAddresses uint32, offset uint32, comment string, ea EA,
	options []*Dhcpoption, useOption bool, serverAssociationType string, failOverAssociation string, member *Dhcpmember, cloudApiCompatible bool, msServer string) (*Rangetemplate, error) {
	if name == "" {
		return nil, fmt.Errorf("name field is required to update a Range Template object")
	}
	rangeTemplate := NewRangeTemplate(ref, name, numberOfAddresses, offset, comment, ea, options, useOption,
		serverAssociationType, failOverAssociation, member, cloudApiCompatible, msServer)
	newRef, err := objMgr.connector.UpdateObject(rangeTemplate, ref)
	if err != nil {
		return nil, fmt.Errorf("error updating Range Template object %s, err: %s", name, err)
	}
	rangeTemplate.Ref = newRef
	rangeTemplate, err = objMgr.GetRangeTemplateByRef(newRef)
	if err != nil {
		return nil, fmt.Errorf("error getting updated Range Template object %s, err: %s", name, err)
	}
	return rangeTemplate, nil
}

func NewRangeTemplate(ref string, name string, numberOfAddresses uint32, offset uint32, comment string, ea EA,
	options []*Dhcpoption, useOption bool, serverAssociationType string, failOverAssociation string, member *Dhcpmember, cloudApiCompatible bool, msServer string) *Rangetemplate {
	rangeTemplate := NewEmptyRangeTemplate()
	rangeTemplate.Ref = ref
	rangeTemplate.Name = &name
	rangeTemplate.NumberOfAddresses = &numberOfAddresses
	rangeTemplate.Offset = &offset
	rangeTemplate.Comment = &comment
	rangeTemplate.Ea = ea
	rangeTemplate.Options = options
	rangeTemplate.UseOptions = &useOption
	rangeTemplate.ServerAssociationType = serverAssociationType
	rangeTemplate.FailoverAssociation = &failOverAssociation
	rangeTemplate.Member = member
	rangeTemplate.CloudApiCompatible = &cloudApiCompatible
	rangeTemplate.MsServer = &Msdhcpserver{Ipv4Addr: msServer}
	return rangeTemplate
}

func NewEmptyRangeTemplate() *Rangetemplate {
	rangeTemplate := &Rangetemplate{}
	rangeTemplate.SetReturnFields(append(rangeTemplate.ReturnFields(), "extattrs", "options", "use_options",
		"server_association_type", "failover_association", "member", "cloud_api_compatible", "ms_server"))
	return rangeTemplate
}

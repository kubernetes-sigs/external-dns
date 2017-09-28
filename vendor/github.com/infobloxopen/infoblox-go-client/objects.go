package ibclient

import (
	"bytes"
	"encoding/json"
	"reflect"
)

const MACADDR_ZERO = "00:00:00:00:00:00"

type Bool bool

type EA map[string]interface{}

type EASearch map[string]interface{}

type EADefListValue string

type IBBase struct {
	objectType   string   `json:"-"`
	returnFields []string `json:"-"`
	eaSearch     EASearch `json:"-"`
}

type IBObject interface {
	ObjectType() string
	ReturnFields() []string
	EaSearch() EASearch
}

func (obj *IBBase) ObjectType() string {
	return obj.objectType
}

func (obj *IBBase) ReturnFields() []string {
	return obj.returnFields
}

func (obj *IBBase) EaSearch() EASearch {
	return obj.eaSearch
}

type NetworkView struct {
	IBBase `json:"-"`
	Ref    string `json:"_ref,omitempty"`
	Name   string `json:"name,omitempty"`
	Ea     EA     `json:"extattrs,omitempty"`
}

func NewNetworkView(nv NetworkView) *NetworkView {
	res := nv
	res.objectType = "networkview"
	res.returnFields = []string{"extattrs", "name"}

	return &res
}

type Network struct {
	IBBase
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	Ea          EA     `json:"extattrs,omitempty"`
}

func NewNetwork(nw Network) *Network {
	res := nw
	res.objectType = "network"
	res.returnFields = []string{"extattrs", "network", "network_view"}

	return &res
}

type NetworkContainer struct {
	IBBase      `json:"-"`
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	Ea          EA     `json:"extattrs,omitempty"`
}

func NewNetworkContainer(nc NetworkContainer) *NetworkContainer {
	res := nc
	res.objectType = "networkcontainer"
	res.returnFields = []string{"extattrs", "network", "network_view"}

	return &res
}

type FixedAddress struct {
	IBBase      `json:"-"`
	Ref         string `json:"_ref,omitempty"`
	NetviewName string `json:"network_view,omitempty"`
	Cidr        string `json:"network,omitempty"`
	IPAddress   string `json:"ipv4addr,omitempty"`
	Mac         string `json:"mac,omitempty"`
	Ea          EA     `json:"extattrs,omitempty"`
}

func NewFixedAddress(fixedAddr FixedAddress) *FixedAddress {
	res := fixedAddr
	res.objectType = "fixedaddress"
	res.returnFields = []string{"extattrs", "ipv4addr", "mac", "network", "network_view"}

	return &res
}

type EADefinition struct {
	IBBase             `json:"-"`
	Ref                string           `json:"_ref,omitempty"`
	Comment            string           `json:"comment,omitempty"`
	Flags              string           `json:"flags,omitempty"`
	ListValues         []EADefListValue `json:"list_values,omitempty"`
	Name               string           `json:"name,omitempty"`
	Type               string           `json:"type,omitempty"`
	AllowedObjectTypes []string         `json:"allowed_object_types,omitempty"`
}

func NewEADefinition(eadef EADefinition) *EADefinition {
	res := eadef
	res.objectType = "extensibleattributedef"
	res.returnFields = []string{"allowed_object_types", "comment", "flags", "list_values", "name", "type"}

	return &res
}

type UserProfile struct {
	IBBase `json:"-"`
	Ref    string `json:"_ref,omitempty"`
	Name   string `json:"name,omitempty"`
}

func NewUserProfile(userprofile UserProfile) *UserProfile {
	res := userprofile
	res.objectType = "userprofile"
	res.returnFields = []string{"name"}

	return &res
}

type RecordA struct {
	IBBase   `json:"-"`
	Ref      string `json:"_ref,omitempty"`
	Ipv4Addr string `json:"ipv4addr,omitempty"`
	Name     string `json:"name,omitempty"`
	View     string `json:"view,omitempty"`
	Zone     string `json:"zone,omitempty"`
	Ea       EA     `json:"extattrs,omitempty"`
}

func NewRecordA(ra RecordA) *RecordA {
	res := ra
	res.objectType = "record:a"
	res.returnFields = []string{"extattrs", "ipv4addr", "name", "view", "zone"}

	return &res
}

type RecordCNAME struct {
	IBBase    `json:"-"`
	Ref       string `json:"_ref,omitempty"`
	Canonical string `json:"canonical,omitempty"`
	Name      string `json:"name,omitempty"`
	View      string `json:"view,omitempty"`
	Zone      string `json:"zone,omitempty"`
	Ea        EA     `json:"extattrs,omitempty"`
}

func NewRecordCNAME(rc RecordCNAME) *RecordCNAME {
	res := rc
	res.objectType = "record:cname"
	res.returnFields = []string{"extattrs", "canonical", "name", "view", "zone"}

	return &res
}

type RecordTXT struct {
	IBBase `json:"-"`
	Ref    string `json:"_ref,omitempty"`
	Name   string `json:"name,omitempty"`
	Text   string `json:"text,omitempty"`
	View   string `json:"view,omitempty"`
	Zone   string `json:"zone,omitempty"`
	Ea     EA     `json:"extattrs,omitempty"`
}

func NewRecordTXT(rt RecordTXT) *RecordTXT {
	res := rt
	res.objectType = "record:txt"
	res.returnFields = []string{"extattrs", "name", "text", "view", "zone"}

	return &res
}

type ZoneAuth struct {
	IBBase `json:"-"`
	Ref    string `json:"_ref,omitempty"`
	Fqdn   string `json:"fqdn,omitempty"`
	View   string `json:"view,omitempty"`
	Ea     EA     `json:"extattrs,omitempty"`
}

func NewZoneAuth(za ZoneAuth) *ZoneAuth {
	res := za
	res.objectType = "zone_auth"
	res.returnFields = []string{"extattrs", "fqdn", "view"}

	return &res
}

func (ea EA) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for k, v := range ea {
		value := make(map[string]interface{})
		value["value"] = v
		m[k] = value
	}

	return json.Marshal(m)
}

func (eas EASearch) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for k, v := range eas {
		m["*"+k] = v
	}

	return json.Marshal(m)
}

func (val EADefListValue) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m["value"] = string(val)

	return json.Marshal(m)
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b {
		return json.Marshal("True")
	}

	return json.Marshal("False")
}

func (ea *EA) UnmarshalJSON(b []byte) (err error) {
	var m map[string]map[string]interface{}

	decoder := json.NewDecoder(bytes.NewBuffer(b))
	decoder.UseNumber()
	err = decoder.Decode(&m)
	if err != nil {
		return
	}

	*ea = make(EA)
	for k, v := range m {
		val := v["value"]
		if reflect.TypeOf(val).String() == "json.Number" {
			var i64 int64
			i64, err = val.(json.Number).Int64()
			val = int(i64)
		} else if val.(string) == "True" {
			val = Bool(true)
		} else if val.(string) == "False" {
			val = Bool(false)
		}

		(*ea)[k] = val
	}

	return
}

func (v *EADefListValue) UnmarshalJSON(b []byte) (err error) {
	var m map[string]string
	err = json.Unmarshal(b, &m)
	if err != nil {
		return
	}

	*v = EADefListValue(m["value"])
	return
}

type RequestBody struct {
	Data               map[string]interface{} `json:"data,omitempty"`
	Args               map[string]string      `json:"args,omitempty"`
	Method             string                 `json:"method"`
	Object             string                 `json:"object,omitempty"`
	EnableSubstitution bool                   `json:"enable_substitution,omitempty"`
	AssignState        map[string]string      `json:"assign_state,omitempty"`
	Discard            bool                   `json:"discard,omitempty"`
}

type SingleRequest struct {
	IBBase `json:"-"`
	Body   *RequestBody
}

type MultiRequest struct {
	IBBase `json:"-"`
	Body   []*RequestBody
}

func (r *MultiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Body)
}

func NewMultiRequest(body []*RequestBody) *MultiRequest {
	req := &MultiRequest{Body: body}
	req.objectType = "request"
	return req
}

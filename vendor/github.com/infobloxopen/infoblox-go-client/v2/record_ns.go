package ibclient

type RecordNS struct {
	IBBase     `json:"-"`
	Ref        string           `json:"_ref,omitempty"`
	Addresses  []ZoneNameServer `json:"addresses,omitempty"`
	Name       string           `json:"name,omitempty"`
	Nameserver string           `json:"nameserver,omitempty"`
	View       string           `json:"view,omitempty"`
	Zone       string           `json:"zone,omitempty"`
}

func NewRecordNS(rc RecordNS) *RecordNS {
	res := rc
	res.objectType = "record:ns"
	res.returnFields = []string{"addresses", "name", "nameserver", "view", "zone"}

	return &res
}

type ZoneNameServer struct {
	Address string `json:"address,omitempty"`
}

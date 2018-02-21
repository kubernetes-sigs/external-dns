package dynect

// DSFSResponse is used for holding the data returned by a call to
// "https://api.dynect.net/REST/DSF/" with 'detail: Y'.
type AllDSFDetailedResponse struct {
	ResponseBlock
	Data []DSFService `json:"data"`
}

// DSFResponse is used for holding the data returned by a call to
// "https://api.dynect.net/REST/DSF/SERVICE_ID".
type DSFResponse struct {
	ResponseBlock
	Data DSFService `json:"data"`
}

// Type DSFService is used as a nested struct, which holds the data for a
// DSF Service returned by a call to "https://api.dynect.net/REST/DSF/SERVICE_ID".
type DSFService struct {
	ID            string       `json:"service_id"`
	Label         string       `json:"label"`
	Active        string       `json:"active"`
	TTL           string       `json:"ttl"`
	PendingChange string       `json:"pending_change"`
	Notifiers     []Notifier   `json:"notifiers"`
	Nodes         []DSFNode    `json:"nodes"`
	Rulesets      []DSFRuleset `json:"rulesets"`
}

type DSFRuleset struct {
	ID            string            `json:"dsf_ruleset_id`
	Label         string            `json:"label"`
	CriteriaType  string            `json:"criteria_type"`
	Criteria      interface{}       `json:"criteria"`
	Ordering      string            `json:"ordering"`
	Eligible      string            `json:"eligible"`
	PendingChange string            `json:"pending_change"`
	ResponsePools []DSFResponsePool `json:"response_pools"`
}

type DSFResponsePool struct {
	ID            string              `json:"dsf_response_pool_id"`
	Label         string              `json:"label"`
	Automation    string              `json:"automation"`
	CoreSetCount  string              `json:"core_set_count"`
	Eligible      string              `json:"eligible"`
	PendingChange string              `json:"pending_change"`
	RsChains      []DSFRecordSetChain `json:"rs_chains"`
	Rulesets      []DSFRuleset        `json:"rulesets"`
	Status        string              `json:"status"`
	LastMonitored string              `json:"last_monitored"`
	Notifier      string              `json:"notifier"`
}

type DSFRecordSetChain struct {
	ID                string         `json:"dsf_record_set_failover_chain_id"`
	Status            string         `json:"status"`
	Core              string         `json:"core"`
	Label             string         `json:"label"`
	DSFResponsePoolID string         `json:"dsf_response_pool_id"`
	DSFServiceID      string         `json:"service_id"`
	PendingChange     string         `json:"pending_change"`
	DSFRecordSets     []DSFRecordSet `json:"record_sets"`
}

type DSFRecordSet struct {
	Status        string      `json:"status"`
	Eligible      string      `json:"eligible"`
	ID            string      `json:"dsf_record_set_id"`
	MonitorID     string      `json:"dsf_monitor_id"`
	Label         string      `json:"label"`
	TroubleCount  string      `json:"trouble_count"`
	Records       []DSFRecord `json:"records"`
	FailCount     string      `json:"fail_count"`
	TorpidityMax  string      `json:"torpidity_max"`
	TTLDerived    string      `json:"ttl_derived"`
	LastMonitored string      `json:"last_monitored"`
	TTL           string      `json:"ttl"`
	ServiceID     string      `json:"service_id"`
	ServeCount    string      `json:"serve_count"`
	Automation    string      `json:"automation"`
	PendingChange string      `json:"pending_change"`
}

type DSFRecord struct {
	Status         string   `json:"status"`
	Endpoints      []string `json:"endpoints"`
	RDataClass     string   `json:"rdata_class"`
	Weight         int      `json:"weight"`
	Eligible       string   `json:"eligible"`
	ID             string   `json:"dsf_record_id"`
	DSFRecordSetID string   `json:"dsf_record_set_id"`
	//RData           interface{} `json:"rdata"`
	EndpointUpCount int    `json:"endpoint_up_count"`
	Label           string `json:"label"`
	MasterLine      string `json:"master_line"`
	Torpidity       int    `json:"torpidity"`
	LastMonitored   int    `json:"last_monitored"`
	TTL             string `json:"ttl"`
	DSFServiceID    string `json:"service_id"`
	PendingChange   string `json:"pending_change"`
	Automation      string `json:"automation"`
	ReponseTime     int    `json:"response_time"`
	Publish         string `json:"publish",omit_empty`
}

type DSFNode struct {
	Zone string `json:"zone"`
	FQDN string `json:"fqdn"`
}

type Notifier struct {
	ID         int    `json:"notifier_id"`
	Label      string `json:"label"`
	Recipients string `json:"recipients"`
	Active     string `json:"active"`
}

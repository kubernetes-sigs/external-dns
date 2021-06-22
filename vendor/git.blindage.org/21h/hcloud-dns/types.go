package hclouddns

// RecordType supported by Hetzner
type RecordType string

const (
	A     RecordType = "A"
	AAAA  RecordType = "AAAA"
	CNAME RecordType = "CNAME"
	MX    RecordType = "MX"
	NS    RecordType = "NS"
	TXT   RecordType = "TXT"
	RP    RecordType = "RP"
	SOA   RecordType = "SOA"
	HINFO RecordType = "HINFO"
	SRV   RecordType = "SRV"
	DANE  RecordType = "DANE"
	TLSA  RecordType = "TLSA"
	DS    RecordType = "DS"
	CAA   RecordType = "CAA"
)

type HCloudClientAdapter interface {
	GetZone(ID string) (HCloudAnswerGetZone, error)
	GetZones(params HCloudGetZonesParams) (HCloudAnswerGetZones, error)
	UpdateZone(zone HCloudZone) (HCloudAnswerGetZone, error)
	DeleteZone(ID string) (HCloudAnswerDeleteZone, error)
	CreateZone(zone HCloudZone) (HCloudAnswerGetZone, error)
	ImportZoneString(zoneID string, zonePlainText string) (HCloudAnswerGetZone, error)
	ExportZoneToString(zoneID string) (HCloudAnswerGetZonePlainText, error)
	ValidateZoneString(zonePlainText string) (HCloudAnswerZoneValidate, error)
	GetRecord(ID string) (HCloudAnswerGetRecord, error)
	GetRecords(params HCloudGetRecordsParams) (HCloudAnswerGetRecords, error)
	UpdateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error)
	DeleteRecord(ID string) (HCloudAnswerDeleteRecord, error)
	CreateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error)
	CreateRecordBulk(record []HCloudRecord) (HCloudAnswerCreateRecords, error)
	UpdateRecordBulk(record []HCloudRecord) (HCloudAnswerUpdateRecords, error)
}

type HCloudClient struct {
	Token string `yaml:"token"`
}

// Hetzner errors roundabout. Fuck you Hetzner.
type HCloudError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// sometime can be returned HCloudError
type HCloudAnswerError struct {
	Error HCloudError `json:"error,omitempty"`
}

// or plain string
type HCloudAnswerErrorString struct {
	Error string `json:"error,omitempty"`
}

type HCloudRecord struct {
	RecordType RecordType `json:"type"`
	ID         string     `json:"id"`
	Created    string     `json:"created"`
	Modified   string     `json:"modified"`
	ZoneID     string     `json:"zone_id"`
	Name       string     `json:"name"`
	Value      string     `json:"value"`
	TTL        int        `json:"ttl"`
}

type HCloudMeta struct {
	Pagination struct {
		Page         int `json:"page"`
		PerPage      int `json:"per_page"`
		LastPage     int `json:"last_page"`
		TotalEntries int `json:"total_entries"`
	} `json:"pagination,omitempty"`
}

type HCloudZone struct {
	ID              string   `json:"id,omitempty"`
	Created         string   `json:"created,omitempty"`
	Modified        string   `json:"modified,omitempty"`
	LegacyDNSHost   string   `json:"legacy_dns_host,omitempty"`
	LegacyNS        []string `json:"legacy_ns,omitempty"`
	Name            string   `json:"name,omitempty"`
	NS              []string `json:"ns,omitempty"`
	Owner           string   `json:"owner,omitempty"`
	Paused          bool     `json:"paused,omitempty"`
	Permission      string   `json:"permission,omitempty"`
	Project         string   `json:"project,omitempty"`
	Registrar       string   `json:"registrar,omitempty"`
	Status          string   `json:"status,omitempty"`
	TTL             int      `json:"ttl,omitempty"`
	Verified        string   `json:"verified,omitempty"`
	RecordsCount    int      `json:"records_count,omitempty"`
	IsSecondaryDNS  bool     `json:"is_secondary_dns,omitempty"`
	TXTverification struct {
		Name  string `json:"name,omitempty"`
		Token string `json:"token,omitempty"`
	} `json:"txt_verification,omitempty"`
}

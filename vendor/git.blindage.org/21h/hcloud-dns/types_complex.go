package hclouddns

// Records answers

type HCloudAnswerGetRecord struct {
	Record HCloudRecord `json:"record,omitempty"`
	Error  HCloudError
}

type HCloudAnswerGetRecords struct {
	Records []HCloudRecord `json:"records,omitempty"`
	Meta    HCloudMeta     `json:"meta,omitempty"`
	Error   HCloudError
}

type HCloudAnswerDeleteRecord struct {
	Error HCloudError
}

type HCloudAnswerCreateRecords struct {
	Records        []HCloudRecord `json:"records,omitempty"`
	ValidRecords   []HCloudRecord `json:"valid_records,omitempty"`
	InvalidRecords []HCloudRecord `json:"invalid_records,omitempty"`
	Error          HCloudError
}

type HCloudAnswerUpdateRecords struct {
	Records        []HCloudRecord `json:"records,omitempty"`
	InvalidRecords []HCloudRecord `json:"failed_records,omitempty"`
	Error          HCloudError
}

// Zones answers

type HCloudAnswerGetZone struct {
	Zone  HCloudZone `json:"zone,omitempty"`
	Error HCloudError
}

type HCloudAnswerGetZonePlainText struct {
	ZonePlainText string `json:"zone,omitempty"`
	Error         HCloudError
}

type HCloudAnswerZoneValidate struct {
	ParsedRecords int          `json:"parsed_records,omitempty"`
	ValidRecords  []HCloudZone `json:"valid_records,omitempty"`
	Error         HCloudError
}

type HCloudAnswerGetZones struct {
	Zones []HCloudZone `json:"zones,omitempty"`
	Meta  HCloudMeta   `json:"meta,omitempty"`
	Error HCloudError
}

type HCloudAnswerDeleteZone struct {
	Error HCloudError
}

// Params

type HCloudGetZonesParams struct {
	Name       string
	SearchName string
	Page       string
	PerPage    string
}

type HCloudGetRecordsParams struct {
	ZoneID  string
	Page    string
	PerPage string
}

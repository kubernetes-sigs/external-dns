# hclouddns
--
    import "git.blindage.org/21h/hcloud-dns"


## Usage

#### type HCloudAnswerCreateRecords

```go
type HCloudAnswerCreateRecords struct {
	Records        []HCloudRecord `json:"records,omitempty"`
	ValidRecords   []HCloudRecord `json:"valid_records,omitempty"`
	InvalidRecords []HCloudRecord `json:"invalid_records,omitempty"`
	Error          HCloudError
}
```


#### type HCloudAnswerDeleteRecord

```go
type HCloudAnswerDeleteRecord struct {
	Error HCloudError
}
```


#### type HCloudAnswerDeleteZone

```go
type HCloudAnswerDeleteZone struct {
	Error HCloudError
}
```


#### type HCloudAnswerError

```go
type HCloudAnswerError struct {
	Error HCloudError `json:"error,omitempty"`
}
```

sometime can be returned HCloudError

#### type HCloudAnswerErrorString

```go
type HCloudAnswerErrorString struct {
	Error string `json:"error,omitempty"`
}
```

or plain string

#### type HCloudAnswerGetRecord

```go
type HCloudAnswerGetRecord struct {
	Record HCloudRecord `json:"record,omitempty"`
	Error  HCloudError
}
```


#### type HCloudAnswerGetRecords

```go
type HCloudAnswerGetRecords struct {
	Records []HCloudRecord `json:"records,omitempty"`
	Meta    HCloudMeta     `json:"meta,omitempty"`
	Error   HCloudError
}
```


#### type HCloudAnswerGetZone

```go
type HCloudAnswerGetZone struct {
	Zone  HCloudZone `json:"zone,omitempty"`
	Error HCloudError
}
```


#### type HCloudAnswerGetZonePlainText

```go
type HCloudAnswerGetZonePlainText struct {
	ZonePlainText string `json:"zone,omitempty"`
	Error         HCloudError
}
```


#### type HCloudAnswerGetZones

```go
type HCloudAnswerGetZones struct {
	Zones []HCloudZone `json:"zones,omitempty"`
	Meta  HCloudMeta   `json:"meta,omitempty"`
	Error HCloudError
}
```


#### type HCloudAnswerUpdateRecords

```go
type HCloudAnswerUpdateRecords struct {
	Records        []HCloudRecord `json:"records,omitempty"`
	InvalidRecords []HCloudRecord `json:"failed_records,omitempty"`
	Error          HCloudError
}
```


#### type HCloudAnswerZoneValidate

```go
type HCloudAnswerZoneValidate struct {
	ParsedRecords int          `json:"parsed_records,omitempty"`
	ValidRecords  []HCloudZone `json:"valid_records,omitempty"`
	Error         HCloudError
}
```


#### type HCloudDNS

```go
type HCloudDNS struct {
}
```

Base types

#### func  New

```go
func New(t string) *HCloudDNS
```
New instance

#### func (*HCloudDNS) CreateRecord

```go
func (d *HCloudDNS) CreateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error)
```
CreateRecord creates new single record. Accepts HCloudRecord with record to
create, of cource no ID. Returns HCloudAnswerGetRecord with HCloudRecord and
error.

#### func (*HCloudDNS) CreateRecordBulk

```go
func (d *HCloudDNS) CreateRecordBulk(record []HCloudRecord) (HCloudAnswerCreateRecords, error)
```
CreateRecordBulk creates many records at once. Accepts array of HCloudRecord,
converts it to json and makes POST to Hetzner. Returns HCloudAnswerCreateRecords
with arrays of HCloudRecord with whole list, valid and invalid, error.

#### func (*HCloudDNS) CreateZone

```go
func (d *HCloudDNS) CreateZone(zone HCloudZone) (HCloudAnswerGetZone, error)
```
CreateZone creates new single zone. Accepts HCloudZone with record to create, of
cource no ID. Returns HCloudAnswerGetZone with HCloudZone and error.

#### func (*HCloudDNS) DeleteRecord

```go
func (d *HCloudDNS) DeleteRecord(ID string) (HCloudAnswerDeleteRecord, error)
```
DeleteRecord remove record by ID. Accepts single ID string. Returns
HCloudAnswerDeleteRecord and error.

#### func (*HCloudDNS) DeleteZone

```go
func (d *HCloudDNS) DeleteZone(ID string) (HCloudAnswerDeleteZone, error)
```
DeleteZone remove zone by ID. Accepts single ID string. Returns
HCloudAnswerDeleteZone with error.

#### func (*HCloudDNS) ExportZoneToString

```go
func (d *HCloudDNS) ExportZoneToString(zoneID string) (HCloudAnswerGetZonePlainText, error)
```
ExportZoneToString exports single zone from imported text. Accepts ID and
zonePlainText strings. Returns HCloudAnswerGetZonePlainText with HCloudZone and
error.

#### func (*HCloudDNS) GetRecord

```go
func (d *HCloudDNS) GetRecord(ID string) (HCloudAnswerGetRecord, error)
```
GetRecord retrieve one single record by ID. Accepts single ID of record. Returns
HCloudAnswerGetRecord with HCloudRecord and error.

#### func (*HCloudDNS) GetRecords

```go
func (d *HCloudDNS) GetRecords(params HCloudGetRecordsParams) (HCloudAnswerGetRecords, error)
```
GetRecords retrieve all records of user. Accepts HCloudGetRecordsParams struct
Returns HCloudAnswerGetRecords with array of HCloudRecord, Meta and error.

#### func (*HCloudDNS) GetZone

```go
func (d *HCloudDNS) GetZone(ID string) (HCloudAnswerGetZone, error)
```
GetZone retrieve one single zone by ID. Accepts zone ID string. Returns
HCloudAnswerGetZone with HCloudZone and error

#### func (*HCloudDNS) GetZones

```go
func (d *HCloudDNS) GetZones(params HCloudGetZonesParams) (HCloudAnswerGetZones, error)
```
GetZones retrieve all zones of user. Accepts exact name as string, search name
with partial name. Returns HCloudAnswerGetZones with array of HCloudZone, Meta
and error.

#### func (*HCloudDNS) ImportZoneString

```go
func (d *HCloudDNS) ImportZoneString(zoneID string, zonePlainText string) (HCloudAnswerGetZone, error)
```
ImportZoneString imports single zone from imported text. Accepts ID and
zonePlainText strings. Returns HCloudAnswerGetZone with HCloudZone and error.

#### func (*HCloudDNS) UpdateRecord

```go
func (d *HCloudDNS) UpdateRecord(record HCloudRecord) (HCloudAnswerGetRecord, error)
```
UpdateRecord makes update of single record by ID. Accepts HCloudRecord with
fullfilled fields. Returns HCloudAnswerGetRecord with HCloudRecord and error.

#### func (*HCloudDNS) UpdateRecordBulk

```go
func (d *HCloudDNS) UpdateRecordBulk(record []HCloudRecord) (HCloudAnswerUpdateRecords, error)
```
UpdateRecordBulk updates many records at once. Accepts array of HCloudRecord,
converting to json and makes PUT to Hetzner. Returns HCloudAnswerUpdateRecords
with arrays of HCloudRecord with updated and failed, error.

#### func (*HCloudDNS) UpdateZone

```go
func (d *HCloudDNS) UpdateZone(zone HCloudZone) (HCloudAnswerGetZone, error)
```
UpdateZone makes update of single zone by ID. Accepts HCloudZone with fullfilled
fields. Returns HCloudAnswerGetZone with HCloudZone and error.

#### func (*HCloudDNS) ValidateZoneString

```go
func (d *HCloudDNS) ValidateZoneString(zonePlainText string) (HCloudAnswerZoneValidate, error)
```
ValidateZoneString validate single zone from imported text. Accepts ID and
zonePlainText strings. Returns HCloudAnswerZoneValidate with HCloudZone and
error.

#### type HCloudError

```go
type HCloudError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
```

Hetzner errors roundabout. Fuck you Hetzner.

#### type HCloudGetRecordsParams

```go
type HCloudGetRecordsParams struct {
	ZoneID  string
	Page    string
	PerPage string
}
```


#### type HCloudGetZonesParams

```go
type HCloudGetZonesParams struct {
	Name       string
	SearchName string
	Page       string
	PerPage    string
}
```


#### type HCloudMeta

```go
type HCloudMeta struct {
	Pagination struct {
		Page         int `json:"page"`
		PerPage      int `json:"per_page"`
		LastPage     int `json:"last_page"`
		TotalEntries int `json:"total_entries"`
	} `json:"pagination,omitempty"`
}
```


#### type HCloudRecord

```go
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
```


#### type HCloudZone

```go
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
```


#### type RecordType

```go
type RecordType string
```

RecordType supported by Hetzner

```go
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
```

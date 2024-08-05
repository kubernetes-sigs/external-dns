package dataset

import (
	"bytes"
	"encoding/json"
	"time"
)

// ExportType is a string enum
type ExportType string

const (
	ExportTypeCSV  = ExportType("csv")
	ExportTypeJSON = ExportType("json")
	ExportTypeXLSX = ExportType("xlsx")
)

// Dataset wraps an NS1 /datasets resource
type Dataset struct {
	ID              string        `json:"id,omitempty"`
	Name            string        `json:"name,omitempty"`
	Datatype        *Datatype     `json:"datatype,omitempty"`
	Repeat          *Repeat       `json:"repeat,omitempty"`
	Timeframe       *Timeframe    `json:"timeframe,omitempty"`
	ExportType      ExportType    `json:"export_type,omitempty"`
	Reports         []*Report     `json:"reports,omitempty"`
	RecipientEmails []string      `json:"recipient_emails,omitempty"`
	CreatedAt       UnixTimestamp `json:"created_at,omitempty"`
	UpdatedAt       UnixTimestamp `json:"updated_at,omitempty"`
}

type DatatypeType string

const (
	DatatypeTypeNumQueries   = DatatypeType("num_queries")
	DatatypeTypeEBOTResponse = DatatypeType("num_ebot_response")
	DatatypeTypeNXDResponse  = DatatypeType("num_nxd_response")
	DatatypeTypeZeroQueries  = DatatypeType("zero_queries")
)

type DatatypeScope string

const (
	DatatypeScopeAccount       = DatatypeScope("account")
	DatatypeScopeNetworkSingle = DatatypeScope("network_single")
	DatatypeScopeRecordSingle  = DatatypeScope("record_single")
	DatatypeScopeZoneSingle    = DatatypeScope("zone_single")
	DatatypeScopeNetworkEach   = DatatypeScope("network_each")
	DatatypeScopeRecordEach    = DatatypeScope("record_each")
	DatatypeScopeZoneEach      = DatatypeScope("zone_each")
	DatatypeScopeTopNZones     = DatatypeScope("top_n_zones")
	DatatypeScopeTopNRecords   = DatatypeScope("top_n_records")
)

// Datatype wraps Dataset's "Datatype" attribute
type Datatype struct {
	Type  DatatypeType      `json:"type,omitempty"`
	Scope DatatypeScope     `json:"scope,omitempty"`
	Data  map[string]string `json:"data,omitempty"`
}

// RepeatsEvery is a string enum
type RepeatsEvery string

const (
	RepeatsEveryWeek  = RepeatsEvery("week")
	RepeatsEveryMonth = RepeatsEvery("month")
	RepeatsEveryYear  = RepeatsEvery("year")
)

// Repeat wraps Dataset's "Repeat" attribute
type Repeat struct {
	Start        UnixTimestamp `json:"start,omitempty"`
	RepeatsEvery RepeatsEvery  `json:"repeats_every,omitempty"`
	EndAfterN    int32         `json:"end_after_n,omitempty"`
}

// TimeframeAggregation is a string enum
type TimeframeAggregation string

const (
	TimeframeAggregationDaily         = TimeframeAggregation("daily")
	TimeframeAggregationMontly        = TimeframeAggregation("monthly")
	TimeframeAggregationBillingPeriod = TimeframeAggregation("billing_period")
)

// Timeframe wraps Dataset's "Timeframe" attribute
type Timeframe struct {
	Aggregation TimeframeAggregation `json:"aggregation,omitempty"`
	Cycles      *int32               `json:"cycles,omitempty"`
	From        *UnixTimestamp       `json:"from,omitempty"`
	To          *UnixTimestamp       `json:"to,omitempty"`
}

// ReportStatus is a string enum
type ReportStatus string

const (
	ReportStatusQueued     = ReportStatus("queued")
	ReportStatusGenerating = ReportStatus("generating")
	ReportStatusAvailable  = ReportStatus("available")
	ReportStatusFailed     = ReportStatus("failed")
)

// Report wraps Dataset's "Report" attribute
type Report struct {
	ID        string        `json:"id,omitempty"`
	Status    ReportStatus  `json:"status,omitempty"`
	Start     UnixTimestamp `json:"start,omitempty"`
	End       UnixTimestamp `json:"end,omitempty"`
	CreatedAt UnixTimestamp `json:"created_at,omitempty"`
}

// NewDataset takes the properties for a Dataset and creates a new instance
func NewDataset(
	id string,
	name string,
	datatype *Datatype,
	repeat *Repeat,
	timeframe *Timeframe,
	exportType ExportType,
	reports []*Report,
	recipientEmails []string,
	createdAt UnixTimestamp,
	updatedAt UnixTimestamp,
) *Dataset {
	return &Dataset{
		ID:              id,
		Name:            name,
		Datatype:        datatype,
		Repeat:          repeat,
		Timeframe:       timeframe,
		ExportType:      exportType,
		Reports:         reports,
		RecipientEmails: recipientEmails,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}

// NewDatatype takes the properties for a Datatype and creates a new instance
func NewDatatype(
	dtype DatatypeType,
	scope DatatypeScope,
	data map[string]string,
) *Datatype {
	return &Datatype{
		Type:  dtype,
		Scope: scope,
		Data:  data,
	}
}

// NewRepeat takes the properties for a Repeat and creates a new instance
func NewRepeat(
	start UnixTimestamp,
	repeatsEvery RepeatsEvery,
	endAfterN int32,
) *Repeat {
	return &Repeat{
		Start:        start,
		RepeatsEvery: repeatsEvery,
		EndAfterN:    endAfterN,
	}
}

// NewTimeframe takes the properties for a Timeframe and creates a new instance
func NewTimeframe(
	aggregation TimeframeAggregation,
	cycles *int32,
	from *UnixTimestamp,
	to *UnixTimestamp,
) *Timeframe {
	return &Timeframe{
		Aggregation: aggregation,
		Cycles:      cycles,
		From:        from,
		To:          to,
	}
}

// NewReport takes the properties for a Report and creates a new instance
func NewReport(
	id string,
	status ReportStatus,
	start UnixTimestamp,
	end UnixTimestamp,
	createdAt UnixTimestamp,
) *Report {
	return &Report{
		ID:        id,
		Status:    status,
		Start:     start,
		End:       end,
		CreatedAt: createdAt,
	}
}

// UnixTimestamp represents a timestamp field that comes as a string-based unix timestamp
type UnixTimestamp time.Time

func (ut *UnixTimestamp) UnmarshalJSON(data []byte) error {
	var unix int64
	data = bytes.Replace(data, []byte(`"`), []byte(""), -1)
	if err := json.Unmarshal(data, &unix); err != nil {
		return err
	}
	*ut = UnixTimestamp(time.Unix(unix, 0))
	return nil
}

func (ut *UnixTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*ut).Unix())
}

package pulsar

// DecisionsQueryParams contains common query parameters for decisions endpoints
type DecisionsQueryParams struct {
	Start      int64    `url:"start,omitempty"`
	End        int64    `url:"end,omitempty"`
	Period     string   `url:"period,omitempty"`
	Area       string   `url:"area,omitempty"`
	ASN        string   `url:"asn,omitempty"`
	Job        string   `url:"job,omitempty"`
	Jobs       []string `url:"jobs,omitempty"`
	Record     string   `url:"record,omitempty"`
	Result     string   `url:"result,omitempty"`
	Agg        string   `url:"agg,omitempty"`
	Geo        string   `url:"geo,omitempty"`
	ZoneID     string   `url:"zone_id,omitempty"`
	CustomerID int64    `url:"customer_id,omitempty"`
}

// DecisionsResponse wraps the response for GetDecisions
type DecisionsResponse struct {
	Graphs []*DecisionsGraph `json:"graphs,omitempty"`
	Total  int64             `json:"total,omitempty"`
}

// DecisionsGraph represents a decision graph with tags
type DecisionsGraph struct {
	Count     int64                 `json:"count,omitempty"`
	GraphData []*DecisionsGraphData `json:"graph_data,omitempty"`
	Avg       int64                 `json:"avg,omitempty"`
	Tags      *Tags                 `json:"tags,omitempty"`
}

// DecisionsGraphData represents a single graph data point
type DecisionsGraphData struct {
	Timestamp int64 `json:"timestamp,omitempty"`
	Count     int64 `json:"count,omitempty"`
}

// Tags represents tag parameters for decisions
type Tags struct {
	JobID string `json:"job_id,omitempty"`
}

// DecisionsGraphRegionResponse wraps the response for GetDecisionsGraphRegion
type DecisionsGraphRegionResponse struct {
	Data []*DecisionsGraphRegionData `json:"data,omitempty"`
	Unit *string                     `json:"unit,omitempty"`
}

// DecisionsGraphRegionData represents regional graph data
type DecisionsGraphRegionData struct {
	Region string         `json:"region,omitempty"`
	Counts []*JobIDCounts `json:"counts,omitempty"`
}

// DecisionsGraphTimeResponse wraps the response for GetDecisionsGraphTime
type DecisionsGraphTimeResponse struct {
	Data []*DecisionsGraphTimeData `json:"data,omitempty"`
	Unit *string                   `json:"unit,omitempty"`
}

// DecisionsGraphTimeData represents time-based graph data
type DecisionsGraphTimeData struct {
	Counts    []*JobIDCounts `json:"counts,omitempty"`
	Timestamp int64          `json:"timestamp,omitempty"`
}

// DecisionsAreaResponse wraps the response for GetDecisionsArea
type DecisionsAreaResponse struct {
	Areas []*AreaData `json:"areas,omitempty"`
}

// AreaData represents area-specific decision data
type AreaData struct {
	AreaName  string         `json:"area_name,omitempty"`
	Parent    *string        `json:"parent,omitempty"`
	Count     int64          `json:"count,omitempty"`
	JobCounts []*JobIDCounts `json:"job_counts,omitempty"`
}

// JobIDCounts represents job-specific counts
type JobIDCounts struct {
	JobID string `json:"job_id,omitempty"`
	Count int64  `json:"count,omitempty"`
}

// DecisionsASNResponse wraps the response for GetDecisionsAsn
type DecisionsASNResponse struct {
	Data []*ASNData `json:"data,omitempty"`
	Unit *string    `json:"unit,omitempty"`
}

// ASNData represents ASN-specific decision data
type ASNData struct {
	ASN                 int64   `json:"asn,omitempty"`
	Count               int64   `json:"count,omitempty"`
	TrafficDistribution float64 `json:"traffic_distribution,omitempty"`
	PreviousDay         float64 `json:"previousDay,omitempty"`
	PreviousWeek        float64 `json:"previousWeek,omitempty"`
}

// DecisionsResultsTimeResponse wraps the response for GetDecisionsResultsTime
type DecisionsResultsTimeResponse struct {
	Data []*ResultsTimeData `json:"data,omitempty"`
}

// ResultsTimeData represents time-based results data
type ResultsTimeData struct {
	Timestamp int64          `json:"timestamp,omitempty"`
	Results   []*ResultCount `json:"results,omitempty"`
}

// ResultCount represents a result with its count
type ResultCount struct {
	Result string `json:"result,omitempty"`
	Count  int64  `json:"count,omitempty"`
}

// DecisionsResultsAreaResponse wraps the response for GetDecisionsResultsArea
type DecisionsResultsAreaResponse struct {
	Area []*DecisionsResultsArea `json:"area,omitempty"`
}

// DecisionsResultsArea represents area-based results data
type DecisionsResultsArea struct {
	Area          string         `json:"area,omitempty"`
	Parent        string         `json:"parent,omitempty"`
	DecisionCount int64          `json:"decision_count,omitempty"`
	Results       []*ResultCount `json:"results,omitempty"`
}

// FiltersTimeResponse wraps the response for GetFiltersTime
type FiltersTimeResponse struct {
	Filters []*FilterTimeData `json:"filters,omitempty"`
}

// FilterTimeData represents time-based filter data
type FilterTimeData struct {
	Timestamp int64            `json:"timestamp,omitempty"`
	Filters   map[string]int64 `json:"filters,omitempty"`
}

// DecisionCustomerResponse wraps the response for customer decision endpoints
type DecisionCustomerResponse struct {
	Data []*CustomerDecisionData `json:"data,omitempty"`
}

// CustomerDecisionData represents customer-specific decision data
type CustomerDecisionData struct {
	Timestamp int64            `json:"timestamp,omitempty"`
	Total     int64            `json:"total,omitempty"`
	JobCounts map[string]int64 `json:"job_counts,omitempty"`
}

// DecisionRecordResponse wraps the response for record decision endpoints
type DecisionRecordResponse struct {
	Data []*RecordDecisionData `json:"data,omitempty"`
}

// RecordDecisionData represents record-specific decision data
type RecordDecisionData struct {
	Timestamp int64            `json:"timestamp,omitempty"`
	Domain    string           `json:"domain,omitempty"`
	RecType   string           `json:"rec_type,omitempty"`
	Total     int64            `json:"total,omitempty"`
	JobCounts map[string]int64 `json:"job_counts,omitempty"`
}

// DecisionTotalResponse wraps the response for GetDecisionTotal
type DecisionTotalResponse struct {
	Total int64 `json:"total,omitempty"`
}

// DecisionsRecordsResponse wraps the response for GetPulsarDecisionsRecords
type DecisionsRecordsResponse struct {
	Total   int64              `json:"total,omitempty"`
	Records map[string]*Record `json:"records,omitempty"`
}

// Record represents a record with count and percentage
type Record struct {
	Count             int64   `json:"count,omitempty"`
	PercentageOfTotal float64 `json:"percentage_of_total,omitempty"`
}

// DecisionsResultsRecordResponse wraps the response for GetPulsarDecisionsResultsRecord
type DecisionsResultsRecordResponse struct {
	Record map[string]*Results `json:"record,omitempty"`
}

// Results represents results with decision count and result map
type Results struct {
	DecisionCount int64            `json:"decision_count,omitempty"`
	Results       map[string]int64 `json:"results,omitempty"`
}

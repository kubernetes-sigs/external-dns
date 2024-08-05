package pulsar

// Job wraps an NS1 pulsar/apps/{appid}/jobs/{jobid} resource
type Job struct {
	Customer  int        `json:"customer,omitempty"`
	TypeID    string     `json:"typeid"`
	Name      string     `json:"name"`
	Community bool       `json:"community,omitempty"`
	JobID     string     `json:"jobid,omitempty"`
	AppID     string     `json:"appid"`
	Active    bool       `json:"active"`
	Shared    bool       `json:"shared"`
	Config    *JobConfig `json:"config,omitempty"`
}

// JobConfig config parameter struct
type JobConfig struct {
	Host                 *string             `json:"host"`
	URLPath              *string             `json:"url_path"`
	HTTP                 *bool               `json:"http,omitempty"`
	HTTPS                *bool               `json:"https,omitempty"`
	RequestTimeoutMillis *int                `json:"request_timeout_millis,omitempty"`
	JobTimeoutMillis     *int                `json:"job_timeout_millis,omitempty"`
	UseXHR               *bool               `json:"use_xhr,omitempty"`
	StaticValues         *bool               `json:"static_values,omitempty"`
	BlendMetricWeights   *BlendMetricWeights `json:"blend_metric_weights,omitempty"`
}

// BlendMetricWeights parameter struct
type BlendMetricWeights struct {
	Timestamp int        `json:"timestamp"`
	Weights   []*Weights `json:"weights"`
}

// Weights parameter struct
type Weights struct {
	Name         string  `json:"name,omitempty"`
	Weight       int     `json:"weight"`
	DefaultValue float64 `json:"default_value"`
	Maximize     bool    `json:"maximize"`
}

// NewJSPulsarJob takes a name, appid, host and urlPath and creates a JavaScript Pulsar job (type *Job)
func NewJSPulsarJob(name string, appid string, host string, urlPath string) *Job {
	return &Job{
		Name:   name,
		TypeID: "latency",
		AppID:  appid,
		Config: &JobConfig{
			Host:    &host,
			URLPath: &urlPath,
		},
	}
}

// NewBBPulsarJob takes a name and appid and creates a Bulk Beacon Pulsar job (type *PulsarJob)
func NewBBPulsarJob(name string, appid string) *Job {
	return &Job{
		Name:   name,
		TypeID: "custom",
		AppID:  appid,
	}
}

package pulsar

// Application wraps an NS1 /pulsar/apps/{appid} resource
type Application struct {
	ID                 string        `json:"appid,omitempty"`
	Name               string        `json:"name"`
	Active             bool          `json:"active"`
	BrowserWaitMillis  int           `json:"browser_wait_millis"`
	JobsPerTransaction int           `json:"jobs_per_transaction"`
	DefaultConfig      DefaultConfig `json:"default_config"`
}

// DefaultConfig contains configuration parameters for application
type DefaultConfig struct {
	HTTP                 bool `json:"http"`
	HTTPS                bool `json:"https"`
	RequestTimeoutMillis int  `json:"request_timeout_millis"`
	JobTimeoutMillis     int  `json:"job_timeout_millis"`
	UseXhr               bool `json:"use_xhr"`
	StaticValues         bool `json:"static_values"`
}

// NewApplication takes a application name and creates a *Application
func NewApplication(name string) *Application {
	return &Application{Name: name}
}

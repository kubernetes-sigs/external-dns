package redirect

// Certificate represents an NS1 redirect configuration object
type Configuration struct {
	ID              *string         `json:"id,omitempty"`
	CertificateID   *string         `json:"certificate_id,omitempty"`
	Domain          string          `json:"domain,omitempty"`
	Path            string          `json:"path,omitempty"`
	Target          string          `json:"target,omitempty"`
	Tags            []string        `json:"tags"`
	ForwardingMode  *ForwardingMode `json:"forwarding_mode,omitempty"`
	ForwardingType  *ForwardingType `json:"forwarding_type,omitempty"`
	HttpsEnabled    *bool           `json:"https_enabled,omitempty"`
	HttpsForced     *bool           `json:"https_forced,omitempty"`
	QueryForwarding *bool           `json:"query_forwarding,omitempty"`
	LastUpdated     *int64          `json:"last_updated,omitempty"`
}

// ConfigurationList represents an NS1 redirect configuration list object
type ConfigurationList struct {
	After   *string          `json:"after,omitempty"`
	Count   int64            `json:"count,omitempty"`
	Limit   *int64           `json:"limit,omitempty"`
	Results []*Configuration `json:"results"`
	Total   int64            `json:"total,omitempty"`
}

// NewConfiguration creates a new configuration with the given parameters
func NewConfiguration(
	domain string,
	path string,
	target string,
	tags []string,
	fwMode *ForwardingMode,
	fwType *ForwardingType,
	httpsEnabled *bool,
	httpsForced *bool,
	queryFwd *bool,
) *Configuration {
	cfg := Configuration{
		Domain:          domain,
		Path:            path,
		Target:          target,
		Tags:            tags,
		ForwardingMode:  fwMode,
		ForwardingType:  fwType,
		HttpsEnabled:    httpsEnabled,
		HttpsForced:     httpsForced,
		QueryForwarding: queryFwd,
	}
	return &cfg
}

// NewConfiguration creates a new configuration with the given parameters
func NewConfigurationMinimal(domain string, path string, target string) *Configuration {
	cfg := Configuration{
		Domain: domain,
		Path:   path,
		Target: target,
	}
	return &cfg
}

// ForwardingMode is a string enum
type ForwardingMode string

const (
	All     ForwardingMode = "all"
	Capture ForwardingMode = "capture"
	None    ForwardingMode = "none"
)

func (s ForwardingMode) String() string {
	switch s {
	case All:
		return "all"
	case Capture:
		return "capture"
	case None:
		return "none"
	}
	return "unknown"
}

func ParseForwardingMode(str string) (ForwardingMode, bool) {
	switch str {
	case "all":
		return All, true
	case "capture":
		return Capture, true
	case "none":
		return None, true
	}
	return "unknown", false
}

// ForwardingType is a string enum
type ForwardingType string

const (
	Masking   ForwardingType = "masking"
	Permanent ForwardingType = "permanent"
	Temporary ForwardingType = "temporary"
)

func (s ForwardingType) String() string {
	switch s {
	case Masking:
		return "masking"
	case Permanent:
		return "permanent"
	case Temporary:
		return "temporary"
	}
	return "unknown"
}

func ParseForwardingType(str string) (ForwardingType, bool) {
	switch str {
	case "masking":
		return Masking, true
	case "permanent":
		return Permanent, true
	case "temporary":
		return Temporary, true
	}
	return "unknown", false
}

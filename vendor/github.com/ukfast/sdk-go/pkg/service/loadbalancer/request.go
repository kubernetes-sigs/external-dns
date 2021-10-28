package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// PatchClusterRequest represents a request to patch a cluster
type PatchClusterRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateTargetRequest represents a request to create a target
type CreateTargetRequest struct {
	Name          string               `json:"name,omitempty"`
	IP            connection.IPAddress `json:"ip"`
	Port          int                  `json:"port,omitempty"`
	Weight        int                  `json:"weight,omitempty"`
	Backup        bool                 `json:"backup"`
	CheckInterval int                  `json:"check_interval,omitempty"`
	CheckSSL      bool                 `json:"check_ssl"`
	CheckRise     int                  `json:"check_rise,omitempty"`
	CheckFall     int                  `json:"check_fall,omitempty"`
	DisableHTTP2  bool                 `json:"disable_http2"`
	HTTP2Only     bool                 `json:"http2_only"`
	Active        bool                 `json:"active"`
}

// PatchTargetRequest represents a request to patch a target
type PatchTargetRequest struct {
	Name          string               `json:"name,omitempty"`
	IP            connection.IPAddress `json:"ip,omitempty"`
	Port          int                  `json:"port,omitempty"`
	Weight        int                  `json:"weight,omitempty"`
	Backup        *bool                `json:"backup,omitempty"`
	CheckInterval int                  `json:"check_interval,omitempty"`
	CheckSSL      *bool                `json:"check_ssl,omitempty"`
	CheckRise     int                  `json:"check_rise,omitempty"`
	CheckFall     int                  `json:"check_fall,omitempty"`
	DisableHTTP2  *bool                `json:"disable_http2,omitempty"`
	HTTP2Only     *bool                `json:"http2_only,omitempty"`
	Active        *bool                `json:"active,omitempty"`
}

// CreateTargetGroupRequest represents a request to create a target group
type CreateTargetGroupRequest struct {
	ClusterID            int                      `json:"cluster_id"`
	Name                 string                   `json:"name"`
	Balance              TargetGroupBalance       `json:"balance"`
	Mode                 Mode                     `json:"mode"`
	Close                bool                     `json:"close"`
	Sticky               bool                     `json:"sticky"`
	CookieOpts           string                   `json:"cookie_opts,omitempty"`
	Source               string                   `json:"source,omitempty"`
	TimeoutsConnect      int                      `json:"timeouts_connect,omitempty"`
	TimeoutsServer       int                      `json:"timeouts_server,omitempty"`
	CustomOptions        string                   `json:"custom_options,omitempty"`
	MonitorURL           string                   `json:"monitor_url,omitempty"`
	MonitorMethod        TargetGroupMonitorMethod `json:"monitor_method,omitempty"`
	MonitorHost          string                   `json:"monitor_host,omitempty"`
	MonitorHTTPVersion   string                   `json:"monitor_http_version,omitempty"`
	MonitorExpect        string                   `json:"monitor_expect,omitempty"`
	MonitorTCPMonitoring bool                     `json:"monitor_tcp_monitoring"`
	CheckPort            int                      `json:"check_port,omitempty"`
	SendProxy            bool                     `json:"send_proxy"`
	SendProxyV2          bool                     `json:"send_proxy_v2"`
	SSL                  bool                     `json:"ssl"`
	SSLVerify            bool                     `json:"ssl_verify"`
	SNI                  bool                     `json:"sni"`
}

// PatchTargetGroupRequest represents a request to patch a target group
type PatchTargetGroupRequest struct {
	Name                 string                   `json:"name,omitempty"`
	Balance              TargetGroupBalance       `json:"balance,omitempty"`
	Mode                 Mode                     `json:"mode,omitempty"`
	Close                *bool                    `json:"close,omitempty"`
	Sticky               *bool                    `json:"sticky,omitempty"`
	CookieOpts           string                   `json:"cookie_opts,omitempty"`
	Source               string                   `json:"source,omitempty"`
	TimeoutsConnect      int                      `json:"timeouts_connect,omitempty"`
	TimeoutsServer       int                      `json:"timeouts_server,omitempty"`
	CustomOptions        string                   `json:"custom_options,omitempty"`
	MonitorURL           string                   `json:"monitor_url,omitempty"`
	MonitorMethod        TargetGroupMonitorMethod `json:"monitor_method,omitempty"`
	MonitorHost          string                   `json:"monitor_host,omitempty"`
	MonitorHTTPVersion   string                   `json:"monitor_http_version,omitempty"`
	MonitorExpect        string                   `json:"monitor_expect,omitempty"`
	MonitorTCPMonitoring *bool                    `json:"monitor_tcp_monitoring,omitempty"`
	CheckPort            int                      `json:"check_port,omitempty"`
	SendProxy            *bool                    `json:"send_proxy,omitempty"`
	SendProxyV2          *bool                    `json:"send_proxy_v2,omitempty"`
	SSL                  *bool                    `json:"ssl,omitempty"`
	SSLVerify            *bool                    `json:"ssl_verify,omitempty"`
	SNI                  *bool                    `json:"sni,omitempty"`
}

// CreateVIPRequest represents a request to create a target group
type CreateVIPRequest struct {
	ClusterID int    `json:"cluster_id"`
	Type      string `json:"type"`
	CIDR      string `json:"cidr"`
}

// PatchVIPRequest represents a request to patch a target group
type PatchVIPRequest struct {
	Type string `json:"type,omitempty"`
	CIDR string `json:"cidr,omitempty"`
}

// CreateListenerRequest represents a request to create a listener
type CreateListenerRequest struct {
	Name                 string `json:"name"`
	ClusterID            int    `json:"cluster_id"`
	HSTSEnabled          bool   `json:"hsts_enabled"`
	Mode                 Mode   `json:"mode"`
	HSTSMaxAge           int    `json:"hsts_maxage"`
	Close                bool   `json:"close"`
	RedirectHTTPS        bool   `json:"redirect_https"`
	DefaultTargetGroupID int    `json:"default_target_group_id"`
	AccessIsAllowList    bool   `json:"access_is_allow_list"`
	AllowTLSV1           bool   `json:"allow_tlsv1"`
	AllowTLSV11          bool   `json:"allow_tlsv11"`
	DisableTLSV12        bool   `json:"disable_tlsv12"`
	DisableHTTP2         bool   `json:"disable_http2"`
	HTTP2Only            bool   `json:"http2_only"`
	CustomCiphers        string `json:"custom_ciphers"`
}

// PatchListenerRequest represents a request to patch a listener
type PatchListenerRequest struct {
	Name                 string `json:"name,omitempty"`
	HSTSEnabled          *bool  `json:"hsts_enabled,omitempty"`
	Mode                 Mode   `json:"mode,omitempty"`
	HSTSMaxAge           int    `json:"hsts_maxage,omitempty"`
	Close                *bool  `json:"close,omitempty"`
	RedirectHTTPS        *bool  `json:"redirect_https,omitempty"`
	DefaultTargetGroupID int    `json:"default_target_group_id,omitempty"`
	AccessIsAllowList    *bool  `json:"access_is_allow_list,omitempty"`
	AllowTLSV1           *bool  `json:"allow_tlsv1,omitempty"`
	AllowTLSV11          *bool  `json:"allow_tlsv11,omitempty"`
	DisableTLSV12        *bool  `json:"disable_tlsv12,omitempty"`
	DisableHTTP2         *bool  `json:"disable_http2,omitempty"`
	HTTP2Only            *bool  `json:"http2_only,omitempty"`
	CustomCiphers        string `json:"custom_ciphers,omitempty"`
}

// CreateAccessIPRequest represents a request to create an access IP
type CreateAccessIPRequest struct {
	IP connection.IPAddress `json:"ip"`
}

// PatchAccessIPRequest represents a request to patch an access IP
type PatchAccessIPRequest struct {
	IP connection.IPAddress `json:"ip,omitempty"`
}

// CreateBindRequest represents a request to create a bind
type CreateBindRequest struct {
	VIPID int `json:"vip_id"`
	Port  int `json:"port"`
}

// PatchBindRequest represents a request to patch a bind
type PatchBindRequest struct {
	VIPID int `json:"vip_id,omitempty"`
	Port  int `json:"port,omitempty"`
}

// CreateCertificateRequest represents a request to create a certificate
type CreateCertificateRequest struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	Certificate string `json:"certificate"`
	CABundle    string `json:"ca_bundle"`
}

// PatchListenerCertificateRequest represents a request to patch a certificate
type PatchCertificateRequest struct {
	Name        string `json:"name,omitempty"`
	Key         string `json:"key,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	CABundle    string `json:"ca_bundle,omitempty"`
}

// CreateACLRequest represents a request to create a ACL
type CreateACLRequest struct {
	Name          string         `json:"name"`
	ListenerID    int            `json:"listener_id,omitempty"`
	TargetGroupID int            `json:"target_group_id,omitempty"`
	Conditions    []ACLCondition `json:"conditions,omitempty"`
	Actions       []ACLAction    `json:"actions"`
}

// PatchListenerACLRequest represents a request to patch a ACL
type PatchACLRequest struct {
	Name       string         `json:"name,omitempty"`
	Conditions []ACLCondition `json:"conditions,omitempty"`
	Actions    []ACLAction    `json:"actions,omitempty"`
}

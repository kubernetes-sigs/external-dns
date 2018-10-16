package linodego

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

type NodeBalancerConfig struct {
	ID             int
	Port           int
	Protocol       ConfigProtocol
	Algorithm      ConfigAlgorithm
	Stickiness     ConfigStickiness
	Check          ConfigCheck
	CheckInterval  int                     `json:"check_interval"`
	CheckAttempts  int                     `json:"check_attempts"`
	CheckPath      string                  `json:"check_path"`
	CheckBody      string                  `json:"check_body"`
	CheckPassive   bool                    `json:"check_passive"`
	CheckTimeout   int                     `json:"check_timeout"`
	CipherSuite    ConfigCipher            `json:"cipher_suite"`
	NodeBalancerID int                     `json:"nodebalancer_id"`
	SSLCommonName  string                  `json:"ssl_commonname"`
	SSLFingerprint string                  `json:"ssl_fingerprint"`
	SSLCert        string                  `json:"ssl_cert"`
	SSLKey         string                  `json:"ssl_key"`
	NodesStatus    *NodeBalancerNodeStatus `json:"nodes_status"`
}

type ConfigAlgorithm string

var (
	AlgorithmRoundRobin ConfigAlgorithm = "roundrobin"
	AlgorithmLeastConn  ConfigAlgorithm = "leastconn"
	AlgorithmSource     ConfigAlgorithm = "source"
)

type ConfigStickiness string

var (
	StickinessNone       ConfigStickiness = "none"
	StickinessTable      ConfigStickiness = "table"
	StickinessHTTPCookie ConfigStickiness = "http_cookie"
)

type ConfigCheck string

var (
	CheckNone       ConfigCheck = "none"
	CheckConnection ConfigCheck = "connection"
	CheckHTTP       ConfigCheck = "http"
	CheckHTTPBody   ConfigCheck = "http_body"
)

type ConfigProtocol string

var (
	ProtocolHTTP  ConfigProtocol = "http"
	ProtocolHTTPS ConfigProtocol = "https"
	ProtocolTCP   ConfigProtocol = "tcp"
)

type ConfigCipher string

var (
	CipherRecommended ConfigCipher = "recommended"
	CipherLegacy      ConfigCipher = "legacy"
)

type NodeBalancerNodeStatus struct {
	Up   int
	Down int
}

// NodeBalancerConfigUpdateOptions are permitted by CreateNodeBalancerConfig
type NodeBalancerConfigCreateOptions struct {
	Port          int              `json:"port"`
	Protocol      ConfigProtocol   `json:"protocol,omitempty"`
	Algorithm     ConfigAlgorithm  `json:"algorithm,omitempty"`
	Stickiness    ConfigStickiness `json:"stickiness,omitempty"`
	Check         ConfigCheck      `json:"check,omitempty"`
	CheckInterval int              `json:"check_interval,omitempty"`
	CheckAttempts int              `json:"check_attempts,omitempty"`
	CheckPath     string           `json:"check_path,omitempty"`
	CheckBody     string           `json:"check_body,omitempty"`
	CheckPassive  *bool            `json:"check_passive,omitempty"`
	CheckTimeout  int              `json:"check_timeout,omitempty"`
	CipherSuite   ConfigCipher     `json:"cipher_suite,omitempty"`
	SSLCert       string           `json:"ssl_cert,omitempty"`
	SSLKey        string           `json:"ssl_key,omitempty"`
}

// NodeBalancerConfigUpdateOptions are permitted by UpdateNodeBalancerConfig
type NodeBalancerConfigUpdateOptions NodeBalancerConfigCreateOptions

func (i NodeBalancerConfig) GetCreateOptions() NodeBalancerConfigCreateOptions {
	return NodeBalancerConfigCreateOptions{
		Port:          i.Port,
		Protocol:      i.Protocol,
		Algorithm:     i.Algorithm,
		Stickiness:    i.Stickiness,
		Check:         i.Check,
		CheckInterval: i.CheckInterval,
		CheckAttempts: i.CheckAttempts,
		CheckTimeout:  i.CheckTimeout,
		CheckPath:     i.CheckPath,
		CheckBody:     i.CheckBody,
		CheckPassive:  &i.CheckPassive,
		CipherSuite:   i.CipherSuite,
		SSLCert:       i.SSLCert,
		SSLKey:        i.SSLKey,
	}
}

func (i NodeBalancerConfig) GetUpdateOptions() NodeBalancerConfigUpdateOptions {
	return NodeBalancerConfigUpdateOptions{
		Port:          i.Port,
		Protocol:      i.Protocol,
		Algorithm:     i.Algorithm,
		Stickiness:    i.Stickiness,
		Check:         i.Check,
		CheckInterval: i.CheckInterval,
		CheckAttempts: i.CheckAttempts,
		CheckPath:     i.CheckPath,
		CheckBody:     i.CheckBody,
		CheckPassive:  &i.CheckPassive,
		CheckTimeout:  i.CheckTimeout,
		CipherSuite:   i.CipherSuite,
		SSLCert:       i.SSLCert,
		SSLKey:        i.SSLKey,
	}
}

// NodeBalancerConfigsPagedResponse represents a paginated NodeBalancerConfig API response
type NodeBalancerConfigsPagedResponse struct {
	*PageOptions
	Data []*NodeBalancerConfig
}

// endpointWithID gets the endpoint URL for NodeBalancerConfig
func (NodeBalancerConfigsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.NodeBalancerConfigs.endpointWithID(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends NodeBalancerConfigs when processing paginated NodeBalancerConfig responses
func (resp *NodeBalancerConfigsPagedResponse) appendData(r *NodeBalancerConfigsPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of NodeBalancerConfig
func (NodeBalancerConfigsPagedResponse) setResult(r *resty.Request) {
	r.SetResult(NodeBalancerConfigsPagedResponse{})
}

// ListNodeBalancerConfigs lists NodeBalancerConfigs
func (c *Client) ListNodeBalancerConfigs(ctx context.Context, nodebalancerID int, opts *ListOptions) ([]*NodeBalancerConfig, error) {
	response := NodeBalancerConfigsPagedResponse{}
	err := c.listHelperWithID(ctx, &response, nodebalancerID, opts)
	for _, el := range response.Data {
		el.fixDates()
	}
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *NodeBalancerConfig) fixDates() *NodeBalancerConfig {
	return v
}

// GetNodeBalancerConfig gets the template with the provided ID
func (c *Client) GetNodeBalancerConfig(ctx context.Context, nodebalancerID int, configID int) (*NodeBalancerConfig, error) {
	e, err := c.NodeBalancerConfigs.endpointWithID(nodebalancerID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, configID)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&NodeBalancerConfig{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancerConfig).fixDates(), nil
}

// CreateNodeBalancerConfig creates a NodeBalancerConfig
func (c *Client) CreateNodeBalancerConfig(ctx context.Context, nodebalancerID int, nodebalancerConfig NodeBalancerConfigCreateOptions) (*NodeBalancerConfig, error) {
	var body string
	e, err := c.NodeBalancerConfigs.endpointWithID(nodebalancerID)

	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&NodeBalancerConfig{})

	if bodyData, err := json.Marshal(nodebalancerConfig); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(e))

	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancerConfig).fixDates(), nil
}

// UpdateNodeBalancerConfig updates the NodeBalancerConfig with the specified id
func (c *Client) UpdateNodeBalancerConfig(ctx context.Context, nodebalancerID int, configID int, updateOpts NodeBalancerConfigUpdateOptions) (*NodeBalancerConfig, error) {
	var body string
	e, err := c.NodeBalancerConfigs.endpointWithID(nodebalancerID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, configID)

	req := c.R(ctx).SetResult(&NodeBalancerConfig{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))

	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancerConfig).fixDates(), nil
}

// DeleteNodeBalancerConfig deletes the NodeBalancerConfig with the specified id
func (c *Client) DeleteNodeBalancerConfig(ctx context.Context, nodebalancerID int, configID int) error {
	e, err := c.NodeBalancerConfigs.endpointWithID(nodebalancerID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, configID)

	if _, err := coupleAPIErrors(c.R(ctx).Delete(e)); err != nil {
		return err
	}

	return nil
}

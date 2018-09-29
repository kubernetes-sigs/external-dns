package linodego

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

type NodeBalancerNode struct {
	ID             int
	Address        string
	Label          string
	Status         string
	Weight         int
	Mode           NodeMode
	ConfigID       int `json:"config_id"`
	NodeBalancerID int `json:"nodebalancer_id"`
}

// NodeMode is the mode a NodeBalancer should use when sending traffic to a NodeBalancer Node
type NodeMode string

var (
	// ModeAccept is the NodeMode indicating a NodeBalancer Node is accepting traffic
	ModeAccept NodeMode = "accept"

	// ModeReject is the NodeMode indicating a NodeBalancer Node is not receiving traffic
	ModeReject NodeMode = "reject"

	// ModeDrain is the NodeMode indicating a NodeBalancer Node is not receiving new traffic, but may continue receiving traffic from pinned connections
	ModeDrain NodeMode = "drain"
)

type NodeBalancerNodeCreateOptions struct {
	Address string   `json:"address"`
	Label   string   `json:"label"`
	Weight  int      `json:"weight,omitempty"`
	Mode    NodeMode `json:"mode,omitempty"`
}

type NodeBalancerNodeUpdateOptions struct {
	Address string   `json:"address,omitempty"`
	Label   string   `json:"label,omitempty"`
	Weight  int      `json:"weight,omitempty"`
	Mode    NodeMode `json:"mode,omitempty"`
}

func (i NodeBalancerNode) GetCreateOptions() NodeBalancerNodeCreateOptions {
	return NodeBalancerNodeCreateOptions{
		Address: i.Address,
		Label:   i.Label,
		Weight:  i.Weight,
		Mode:    i.Mode,
	}
}

func (i NodeBalancerNode) GetUpdateOptions() NodeBalancerNodeUpdateOptions {
	return NodeBalancerNodeUpdateOptions{
		Address: i.Address,
		Label:   i.Label,
		Weight:  i.Weight,
		Mode:    i.Mode,
	}
}

// NodeBalancerNodesPagedResponse represents a paginated NodeBalancerNode API response
type NodeBalancerNodesPagedResponse struct {
	*PageOptions
	Data []*NodeBalancerNode
}

// endpoint gets the endpoint URL for NodeBalancerNode
func (NodeBalancerNodesPagedResponse) endpointWithTwoIDs(c *Client, nodebalancerID int, configID int) string {
	endpoint, err := c.NodeBalancerNodes.endpointWithID(nodebalancerID, configID)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends NodeBalancerNodes when processing paginated NodeBalancerNode responses
func (resp *NodeBalancerNodesPagedResponse) appendData(r *NodeBalancerNodesPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of NodeBalancerNode
func (NodeBalancerNodesPagedResponse) setResult(r *resty.Request) {
	r.SetResult(NodeBalancerNodesPagedResponse{})
}

// ListNodeBalancerNodes lists NodeBalancerNodes
func (c *Client) ListNodeBalancerNodes(ctx context.Context, nodebalancerID int, configID int, opts *ListOptions) ([]*NodeBalancerNode, error) {
	response := NodeBalancerNodesPagedResponse{}
	err := c.listHelperWithTwoIDs(ctx, &response, nodebalancerID, configID, opts)
	for _, el := range response.Data {
		el.fixDates()
	}
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *NodeBalancerNode) fixDates() *NodeBalancerNode {
	return v
}

// GetNodeBalancerNode gets the template with the provided ID
func (c *Client) GetNodeBalancerNode(ctx context.Context, nodebalancerID int, configID int, nodeID int) (*NodeBalancerNode, error) {
	e, err := c.NodeBalancerNodes.endpointWithID(nodebalancerID, configID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, nodeID)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&NodeBalancerNode{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancerNode).fixDates(), nil
}

// CreateNodeBalancerNode creates a NodeBalancerNode
func (c *Client) CreateNodeBalancerNode(ctx context.Context, nodebalancerID int, configID int, createOpts NodeBalancerNodeCreateOptions) (*NodeBalancerNode, error) {
	var body string
	e, err := c.NodeBalancerNodes.endpointWithID(nodebalancerID, configID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&NodeBalancerNode{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	if err != nil {
		return nil, err
	}
	return r.Result().(*NodeBalancerNode).fixDates(), nil
}

// UpdateNodeBalancerNode updates the NodeBalancerNode with the specified id
func (c *Client) UpdateNodeBalancerNode(ctx context.Context, nodebalancerID int, configID int, nodeID int, updateOpts NodeBalancerNodeUpdateOptions) (*NodeBalancerNode, error) {
	var body string
	e, err := c.NodeBalancerNodes.endpointWithID(nodebalancerID, configID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, nodeID)

	req := c.R(ctx).SetResult(&NodeBalancerNode{})

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
	return r.Result().(*NodeBalancerNode).fixDates(), nil
}

// DeleteNodeBalancerNode deletes the NodeBalancerNode with the specified id
func (c *Client) DeleteNodeBalancerNode(ctx context.Context, nodebalancerID int, configID int, nodeID int) error {
	e, err := c.NodeBalancerNodes.endpointWithID(nodebalancerID, configID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, nodeID)

	if _, err := coupleAPIErrors(c.R(ctx).Delete(e)); err != nil {
		return err
	}

	return nil
}

package linodego

import (
	"context"
	"fmt"
)

// NodeBalancerStats represents a nodebalancer stats object
type NodeBalancerStats struct {
	Title string                `json:"title"`
	Data  NodeBalancerStatsData `json:"data"`
}

// NodeBalancerStatsData represents a nodebalancer stats data object
type NodeBalancerStatsData struct {
	Connections [][]float64  `json:"connections"`
	Traffic     StatsTraffic `json:"traffic"`
}

// StatsTraffic represents a Traffic stats object
type StatsTraffic struct {
	In  [][]float64 `json:"in"`
	Out [][]float64 `json:"out"`
}

// GetNodeBalancerStats gets the template with the provided ID
<<<<<<< HEAD
func (c *Client) GetNodeBalancerStats(ctx context.Context, linodeID int) (*NodeBalancerStats, error) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	e, err := c.NodeBalancerStats.endpointWithParams(linodeID)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	e, err := c.NodeBalancerStats.endpointWithID(linodeID)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	e, err := c.NodeBalancerStats.endpointWithID(linodeID)
=======
	e, err := c.NodeBalancerStats.endpointWithParams(linodeID)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	e, err := c.NodeBalancerStats.endpointWithID(linodeID)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	e, err := c.NodeBalancerStats.endpointWithID(linodeID)
=======
	e, err := c.NodeBalancerStats.endpointWithParams(linodeID)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	e, err := c.NodeBalancerStats.endpointWithID(linodeID)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	e, err := c.NodeBalancerStats.endpointWithID(linodeID)
=======
	e, err := c.NodeBalancerStats.endpointWithParams(linodeID)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	e, err := c.NodeBalancerStats.endpointWithID(linodeID)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	if err != nil {
		return nil, err
	}
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&NodeBalancerStats{}).Get(e))
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func (c *Client) GetNodeBalancerStats(ctx context.Context, linodeID int) (*NodeBalancerStats, error) {
	e, err := c.NodeBalancerStats.endpointWithID(linodeID)
	if err != nil {
		return nil, err
	}
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&NodeBalancerStats{}).Get(e))
=======
func (c *Client) GetNodeBalancerStats(ctx context.Context, nodebalancerID int) (*NodeBalancerStats, error) {
	e := fmt.Sprintf("nodebalancers/%d/stats", nodebalancerID)
	req := c.R(ctx).SetResult(&NodeBalancerStats{})
	r, err := coupleAPIErrors(req.Get(e))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return nil, err
	}

	return r.Result().(*NodeBalancerStats), nil
}

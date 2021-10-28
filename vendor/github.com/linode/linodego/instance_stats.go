package linodego

import (
	"context"
	"fmt"
)

// StatsNet represents a network stats object
type StatsNet struct {
	In         [][]float64 `json:"in"`
	Out        [][]float64 `json:"out"`
	PrivateIn  [][]float64 `json:"private_in"`
	PrivateOut [][]float64 `json:"private_out"`
}

// StatsIO represents an IO stats object
type StatsIO struct {
	IO   [][]float64 `json:"io"`
	Swap [][]float64 `json:"swap"`
}

// InstanceStatsData represents an instance stats data object
type InstanceStatsData struct {
	CPU   [][]float64 `json:"cpu"`
	IO    StatsIO     `json:"io"`
	NetV4 StatsNet    `json:"netv4"`
	NetV6 StatsNet    `json:"netv6"`
}

// InstanceStats represents an instance stats object
type InstanceStats struct {
	Title string            `json:"title"`
	Data  InstanceStatsData `json:"data"`
}

// endpointWithIDAndDate gets the endpoint URL for InstanceStats of a given Instance and Year/Month
func endpointWithIDAndDate(c *Client, id int, year int, month int) string {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	endpoint, err := c.InstanceStats.endpointWithParams(id)
	if err != nil {
		panic(err)
	}

	endpoint = fmt.Sprintf("%s/%d/%d", endpoint, year, month)
	return endpoint
}

// GetInstanceStats gets the template with the provided ID
func (c *Client) GetInstanceStats(ctx context.Context, linodeID int) (*InstanceStats, error) {
	e, err := c.InstanceStats.endpointWithParams(linodeID)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceStats.endpointWithID(id)
||||||| parent of 5ce8c7613 (update vendored files)
	endpoint, err := c.InstanceStats.endpointWithID(id)
=======
	endpoint, err := c.InstanceStats.endpointWithParams(id)
>>>>>>> 5ce8c7613 (update vendored files)
	if err != nil {
		panic(err)
	}

	endpoint = fmt.Sprintf("%s/%d/%d", endpoint, year, month)
	return endpoint
}

// GetInstanceStats gets the template with the provided ID
func (c *Client) GetInstanceStats(ctx context.Context, linodeID int) (*InstanceStats, error) {
<<<<<<< HEAD
	e, err := c.InstanceStats.endpointWithID(linodeID)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	e, err := c.InstanceStats.endpointWithID(linodeID)
=======
	e, err := c.InstanceStats.endpointWithParams(linodeID)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceStats.endpointWithID(id)
||||||| parent of 6b7ce455e (update vendored files)
	endpoint, err := c.InstanceStats.endpointWithID(id)
=======
	endpoint, err := c.InstanceStats.endpointWithParams(id)
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		panic(err)
	}

	endpoint = fmt.Sprintf("%s/%d/%d", endpoint, year, month)
	return endpoint
}

// GetInstanceStats gets the template with the provided ID
func (c *Client) GetInstanceStats(ctx context.Context, linodeID int) (*InstanceStats, error) {
<<<<<<< HEAD
	e, err := c.InstanceStats.endpointWithID(linodeID)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	e, err := c.InstanceStats.endpointWithID(linodeID)
=======
	e, err := c.InstanceStats.endpointWithParams(linodeID)
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		return nil, err
	}
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceStats{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceStats), nil
}

// GetInstanceStatsByDate gets the template with the provided ID, year, and month
func (c *Client) GetInstanceStatsByDate(ctx context.Context, linodeID int, year int, month int) (*InstanceStats, error) {
	e := endpointWithIDAndDate(c, linodeID, year, month)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceStats{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceStats), nil
}

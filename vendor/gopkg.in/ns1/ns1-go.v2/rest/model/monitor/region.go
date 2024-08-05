package monitor

// Region wraps an NS1 /monitoring/regions resource.
type Region struct {
	Code    string   `json:"code"`
	Name    string   `json:"name"`
	Subnets []string `json:"subnets"`
}

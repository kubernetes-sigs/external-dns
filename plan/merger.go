package plan

import "sigs.k8s.io/external-dns/endpoint"

const (
	mergerKey    string = "merge-strategy"
	mergeTargets string = "merge-targets"
)

// EndpointMerger defines how fields from Candidate and Current endpoints
// are merged into the resulting endpoint. For example, targets belonging
// to different owners can be merged into a single list of targets.
// The merger implementation can be selected using the mergerKey
// label on the DNSEndpoint resource.
type EndpointMerger interface {
	Merge(candidate *endpoint.Endpoint, current *endpoint.Endpoint) *endpoint.Endpoint
}

// resolveMerger returns an EndpointMerger based on the labels of the candidate endpoint.
// if merger is unknown or not set, it returns a default merger.
func resolveMerger(candidate *endpoint.Endpoint) EndpointMerger {
	switch candidate.Labels[mergerKey] {
	case mergeTargets:
		return &TargetMerger{}
	}
	return &DefaulMerger{}
}

type DefaulMerger struct {
}

type TargetMerger struct {
}

// Merge default returns the candidate endpoint as is.
func (d DefaulMerger) Merge(candidate *endpoint.Endpoint, _ *endpoint.Endpoint) *endpoint.Endpoint {
	return candidate
}

// Merge combines targets from both candidate and current endpoints into the candidate endpoint.
// Function doesn't preserve order
func (t TargetMerger) Merge(candidate *endpoint.Endpoint, current *endpoint.Endpoint) *endpoint.Endpoint {
	m := map[string]bool{}
	for _, v := range current.Targets {
		m[v] = true
	}
	for _, v := range candidate.Targets {
		m[v] = true
	}
	candidate.Targets = endpoint.Targets{}
	for k, _ := range m {
		candidate.Targets = append(candidate.Targets, k)
	}
	return candidate
}

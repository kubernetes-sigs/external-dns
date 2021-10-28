package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedTarget represents a paginated collection of Target
type PaginatedTarget struct {
	*connection.PaginatedBase
	Items []Target
}

// NewPaginatedTarget returns a pointer to an initialized PaginatedTarget struct
func NewPaginatedTarget(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Target) *PaginatedTarget {
	return &PaginatedTarget{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedTargetGroup represents a paginated collection of TargetGroup
type PaginatedTargetGroup struct {
	*connection.PaginatedBase
	Items []TargetGroup
}

// NewPaginatedTargetGroup returns a pointer to an initialized PaginatedTargetGroup struct
func NewPaginatedTargetGroup(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []TargetGroup) *PaginatedTargetGroup {
	return &PaginatedTargetGroup{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedCluster represents a paginated collection of Cluster
type PaginatedCluster struct {
	*connection.PaginatedBase
	Items []Cluster
}

// NewPaginatedCluster returns a pointer to an initialized PaginatedCluster struct
func NewPaginatedCluster(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Cluster) *PaginatedCluster {
	return &PaginatedCluster{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedVIP represents a paginated collection of VIP
type PaginatedVIP struct {
	*connection.PaginatedBase
	Items []VIP
}

// NewPaginatedVIP returns a pointer to an initialized PaginatedVIP struct
func NewPaginatedVIP(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []VIP) *PaginatedVIP {
	return &PaginatedVIP{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedListener represents a paginated collection of Listener
type PaginatedListener struct {
	*connection.PaginatedBase
	Items []Listener
}

// NewPaginatedListener returns a pointer to an initialized PaginatedListener struct
func NewPaginatedListener(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Listener) *PaginatedListener {
	return &PaginatedListener{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedAccessIP represents a paginated collection of AccessIP
type PaginatedAccessIP struct {
	*connection.PaginatedBase
	Items []AccessIP
}

// NewPaginatedAccessIP returns a pointer to an initialized PaginatedAccessIP struct
func NewPaginatedAccessIP(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []AccessIP) *PaginatedAccessIP {
	return &PaginatedAccessIP{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedBind represents a paginated collection of Bind
type PaginatedBind struct {
	*connection.PaginatedBase
	Items []Bind
}

// NewPaginatedBind returns a pointer to an initialized PaginatedBind struct
func NewPaginatedBind(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Bind) *PaginatedBind {
	return &PaginatedBind{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedCertificate represents a paginated collection of Certificate
type PaginatedCertificate struct {
	*connection.PaginatedBase
	Items []Certificate
}

// NewPaginatedCertificate returns a pointer to an initialized PaginatedCertificate struct
func NewPaginatedCertificate(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Certificate) *PaginatedCertificate {
	return &PaginatedCertificate{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedHeader represents a paginated collection of Header
type PaginatedHeader struct {
	*connection.PaginatedBase
	Items []Header
}

// NewPaginatedHeader returns a pointer to an initialized PaginatedHeader struct
func NewPaginatedHeader(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Header) *PaginatedHeader {
	return &PaginatedHeader{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedACL represents a paginated collection of ACL
type PaginatedACL struct {
	*connection.PaginatedBase
	Items []ACL
}

// NewPaginatedACL returns a pointer to an initialized PaginatedACL struct
func NewPaginatedACL(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ACL) *PaginatedACL {
	return &PaginatedACL{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedACLArgument represents a paginated collection of ACLArgument
type PaginatedACLArgument struct {
	*connection.PaginatedBase
	Items []ACLArgument
}

// NewPaginatedACLArgument returns a pointer to an initialized PaginatedACLArgument struct
func NewPaginatedACLArgument(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ACLArgument) *PaginatedACLArgument {
	return &PaginatedACLArgument{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedACLCondition represents a paginated collection of ACLCondition
type PaginatedACLCondition struct {
	*connection.PaginatedBase
	Items []ACLCondition
}

// NewPaginatedACLCondition returns a pointer to an initialized PaginatedACLCondition struct
func NewPaginatedACLCondition(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ACLCondition) *PaginatedACLCondition {
	return &PaginatedACLCondition{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedACLAction represents a paginated collection of ACLAction
type PaginatedACLAction struct {
	*connection.PaginatedBase
	Items []ACLAction
}

// NewPaginatedACLAction returns a pointer to an initialized PaginatedACLAction struct
func NewPaginatedACLAction(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ACLAction) *PaginatedACLAction {
	return &PaginatedACLAction{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

package loadbalancer

import "github.com/ukfast/sdk-go/pkg/connection"

// GetTargetSliceResponseBody represents an API response body containing []Target data
type GetTargetSliceResponseBody struct {
	connection.APIResponseBody
	Data []Target `json:"data"`
}

// GetTargetResponseBody represents an API response body containing Target data
type GetTargetResponseBody struct {
	connection.APIResponseBody
	Data Target `json:"data"`
}

// GetTargetGroupSliceResponseBody represents an API response body containing []TargetGroup data
type GetTargetGroupSliceResponseBody struct {
	connection.APIResponseBody
	Data []TargetGroup `json:"data"`
}

// GetTargetGroupResponseBody represents an API response body containing TargetGroup data
type GetTargetGroupResponseBody struct {
	connection.APIResponseBody
	Data TargetGroup `json:"data"`
}

// GetClusterSliceResponseBody represents an API response body containing []Cluster data
type GetClusterSliceResponseBody struct {
	connection.APIResponseBody
	Data []Cluster `json:"data"`
}

// GetClusterResponseBody represents an API response body containing Cluster data
type GetClusterResponseBody struct {
	connection.APIResponseBody
	Data Cluster `json:"data"`
}

// GetVIPSliceResponseBody represents an API response body containing []VIP data
type GetVIPSliceResponseBody struct {
	connection.APIResponseBody
	Data []VIP `json:"data"`
}

// GetVIPResponseBody represents an API response body containing VIP data
type GetVIPResponseBody struct {
	connection.APIResponseBody
	Data VIP `json:"data"`
}

// GetListenerSliceResponseBody represents an API response body containing []Listener data
type GetListenerSliceResponseBody struct {
	connection.APIResponseBody
	Data []Listener `json:"data"`
}

// GetListenerResponseBody represents an API response body containing Listener data
type GetListenerResponseBody struct {
	connection.APIResponseBody
	Data Listener `json:"data"`
}

// GetAccessIPSliceResponseBody represents an API response body containing []AccessIP data
type GetAccessIPSliceResponseBody struct {
	connection.APIResponseBody
	Data []AccessIP `json:"data"`
}

// GetAccessIPResponseBody represents an API response body containing AccessIP data
type GetAccessIPResponseBody struct {
	connection.APIResponseBody
	Data AccessIP `json:"data"`
}

// GetBindSliceResponseBody represents an API response body containing []Bind data
type GetBindSliceResponseBody struct {
	connection.APIResponseBody
	Data []Bind `json:"data"`
}

// GetBindResponseBody represents an API response body containing Bind data
type GetBindResponseBody struct {
	connection.APIResponseBody
	Data Bind `json:"data"`
}

// GetCertificateSliceResponseBody represents an API response body containing []Certificate data
type GetCertificateSliceResponseBody struct {
	connection.APIResponseBody
	Data []Certificate `json:"data"`
}

// GetCertificateResponseBody represents an API response body containing Certificate data
type GetCertificateResponseBody struct {
	connection.APIResponseBody
	Data Certificate `json:"data"`
}

// GetHeaderSliceResponseBody represents an API response body containing []Header data
type GetHeaderSliceResponseBody struct {
	connection.APIResponseBody
	Data []Header `json:"data"`
}

// GetHeaderResponseBody represents an API response body containing Header data
type GetHeaderResponseBody struct {
	connection.APIResponseBody
	Data Header `json:"data"`
}

// GetACLSliceResponseBody represents an API response body containing []ACL data
type GetACLSliceResponseBody struct {
	connection.APIResponseBody
	Data []ACL `json:"data"`
}

// GetACLResponseBody represents an API response body containing ACL data
type GetACLResponseBody struct {
	connection.APIResponseBody
	Data ACL `json:"data"`
}

// GetACLArgumentSliceResponseBody represents an API response body containing []ACLArgument data
type GetACLArgumentSliceResponseBody struct {
	connection.APIResponseBody
	Data []ACLArgument `json:"data"`
}

// GetACLArgumentResponseBody represents an API response body containing ACLArgument data
type GetACLArgumentResponseBody struct {
	connection.APIResponseBody
	Data ACLArgument `json:"data"`
}

// GetACLConditionSliceResponseBody represents an API response body containing []ACLCondition data
type GetACLConditionSliceResponseBody struct {
	connection.APIResponseBody
	Data []ACLCondition `json:"data"`
}

// GetACLConditionResponseBody represents an API response body containing ACLCondition data
type GetACLConditionResponseBody struct {
	connection.APIResponseBody
	Data ACLCondition `json:"data"`
}

// GetACLActionSliceResponseBody represents an API response body containing []ACLAction data
type GetACLActionSliceResponseBody struct {
	connection.APIResponseBody
	Data []ACLAction `json:"data"`
}

// GetACLActionResponseBody represents an API response body containing ACLAction data
type GetACLActionResponseBody struct {
	connection.APIResponseBody
	Data ACLAction `json:"data"`
}

// GetACLTemplatesSliceResponseBody represents an API response body containing []ACLTemplates data
type GetACLTemplatesSliceResponseBody struct {
	connection.APIResponseBody
	Data []ACLTemplates `json:"data"`
}

// GetACLTemplatesResponseBody represents an API response body containing ACLTemplates data
type GetACLTemplatesResponseBody struct {
	connection.APIResponseBody
	Data ACLTemplates `json:"data"`
}

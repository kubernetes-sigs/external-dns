// Package cloudresourcemanager provides access to the Google Cloud Resource Manager API.
//
// See https://cloud.google.com/resource-manager
//
// Usage example:
//
//   import "google.golang.org/api/cloudresourcemanager/v1beta1"
//   ...
//   cloudresourcemanagerService, err := cloudresourcemanager.New(oauthHttpClient)
package cloudresourcemanager // import "google.golang.org/api/cloudresourcemanager/v1beta1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "cloudresourcemanager:v1beta1"
const apiName = "cloudresourcemanager"
const apiVersion = "v1beta1"
const basePath = "https://cloudresourcemanager.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Organizations = NewOrganizationsService(s)
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Organizations *OrganizationsService

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewOrganizationsService(s *Service) *OrganizationsService {
	rs := &OrganizationsService{s: s}
	return rs
}

type OrganizationsService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	return rs
}

type ProjectsService struct {
	s *Service
}

// Ancestor: Identifying information for a single ancestor of a project.
type Ancestor struct {
	// ResourceId: Resource id of the ancestor.
	ResourceId *ResourceId `json:"resourceId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResourceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResourceId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Ancestor) MarshalJSON() ([]byte, error) {
	type noMethod Ancestor
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AuditConfig: Specifies the audit configuration for a service.
// The configuration determines which permission types are logged, and
// what
// identities, if any, are exempted from logging.
// An AuditConifg must have one or more AuditLogConfigs.
//
// If there are AuditConfigs for both `allServices` and a specific
// service,
// the union of the two AuditConfigs is used for that service: the
// log_types
// specified in each AuditConfig are enabled, and the exempted_members
// in each
// AuditConfig are exempted.
// Example Policy with multiple AuditConfigs:
// {
//   "audit_configs": [
//     {
//       "service": "allServices"
//       "audit_log_configs": [
//         {
//           "log_type": "DATA_READ",
//           "exempted_members": [
//             "user:foo@gmail.com"
//           ]
//         },
//         {
//           "log_type": "DATA_WRITE",
//         },
//         {
//           "log_type": "ADMIN_READ",
//         }
//       ]
//     },
//     {
//       "service": "fooservice@googleapis.com"
//       "audit_log_configs": [
//         {
//           "log_type": "DATA_READ",
//         },
//         {
//           "log_type": "DATA_WRITE",
//           "exempted_members": [
//             "user:bar@gmail.com"
//           ]
//         }
//       ]
//     }
//   ]
// }
// For fooservice, this policy enables DATA_READ, DATA_WRITE and
// ADMIN_READ
// logging. It also exempts foo@gmail.com from DATA_READ logging,
// and
// bar@gmail.com from DATA_WRITE logging.
type AuditConfig struct {
	// AuditLogConfigs: The configuration for logging of each type of
	// permission.
	// Next ID: 4
	AuditLogConfigs []*AuditLogConfig `json:"auditLogConfigs,omitempty"`

	// Service: Specifies a service that will be enabled for audit
	// logging.
	// For example, `storage.googleapis.com`,
	// `cloudsql.googleapis.com`.
	// `allServices` is a special value that covers all services.
	Service string `json:"service,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuditLogConfigs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditLogConfigs") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AuditConfig) MarshalJSON() ([]byte, error) {
	type noMethod AuditConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AuditLogConfig: Provides the configuration for logging a type of
// permissions.
// Example:
//
//     {
//       "audit_log_configs": [
//         {
//           "log_type": "DATA_READ",
//           "exempted_members": [
//             "user:foo@gmail.com"
//           ]
//         },
//         {
//           "log_type": "DATA_WRITE",
//         }
//       ]
//     }
//
// This enables 'DATA_READ' and 'DATA_WRITE' logging, while
// exempting
// foo@gmail.com from DATA_READ logging.
type AuditLogConfig struct {
	// ExemptedMembers: Specifies the identities that do not cause logging
	// for this type of
	// permission.
	// Follows the same format of Binding.members.
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// LogType: The log type that this config enables.
	//
	// Possible values:
	//   "LOG_TYPE_UNSPECIFIED" - Default case. Should never be this.
	//   "ADMIN_READ" - Admin reads. Example: CloudIAM getIamPolicy
	//   "DATA_WRITE" - Data writes. Example: CloudSQL Users create
	//   "DATA_READ" - Data reads. Example: CloudSQL Users list
	LogType string `json:"logType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExemptedMembers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExemptedMembers") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AuditLogConfig) MarshalJSON() ([]byte, error) {
	type noMethod AuditLogConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Binding: Associates `members` with a `role`.
type Binding struct {
	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents
	// anyone
	//    who is authenticated with a Google account or a service
	// account.
	//
	// * `user:{emailid}`: An email address that represents a specific
	// Google
	//    account. For example, `alice@gmail.com` or `joe@example.com`.
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a
	// service
	//    account. For example,
	// `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google
	// group.
	//    For example, `admins@example.com`.
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all
	// the
	//    users of that domain. For example, `google.com` or
	// `example.com`.
	//
	//
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or
	// `roles/owner`.
	// Required
	Role string `json:"role,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Members") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Members") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Binding) MarshalJSON() ([]byte, error) {
	type noMethod Binding
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BooleanConstraint: A `Constraint` that is either enforced or
// not.
//
// For example a constraint
// `constraints/compute.disableSerialPortAccess`.
// If it is enforced on a VM instance, serial port connections will not
// be
// opened to that instance.
type BooleanConstraint struct {
}

// BooleanPolicy: Used in `policy_type` to specify how `boolean_policy`
// will behave at this
// resource.
type BooleanPolicy struct {
	// Enforced: If `true`, then the `Policy` is enforced. If `false`, then
	// any
	// configuration is acceptable.
	//
	// Suppose you have a `Constraint`
	// `constraints/compute.disableSerialPortAccess`
	// with `constraint_default` set to `ALLOW`. A `Policy` for
	// that
	// `Constraint` exhibits the following behavior:
	//   - If the `Policy` at this resource has enforced set to `false`,
	// serial
	//     port connection attempts will be allowed.
	//   - If the `Policy` at this resource has enforced set to `true`,
	// serial
	//     port connection attempts will be refused.
	//   - If the `Policy` at this resource is `RestoreDefault`, serial
	// port
	//     connection attempts will be allowed.
	//   - If no `Policy` is set at this resource or anywhere higher in the
	//     resource hierarchy, serial port connection attempts will be
	// allowed.
	//   - If no `Policy` is set at this resource, but one exists higher in
	// the
	//     resource hierarchy, the behavior is as if the`Policy` were set
	// at
	//     this resource.
	//
	// The following examples demonstrate the different possible
	// layerings:
	//
	// Example 1 (nearest `Constraint` wins):
	//   `organizations/foo` has a `Policy` with:
	//     {enforced: false}
	//   `projects/bar` has no `Policy` set.
	// The constraint at `projects/bar` and `organizations/foo` will not
	// be
	// enforced.
	//
	// Example 2 (enforcement gets replaced):
	//   `organizations/foo` has a `Policy` with:
	//     {enforced: false}
	//   `projects/bar` has a `Policy` with:
	//     {enforced: true}
	// The constraint at `organizations/foo` is not enforced.
	// The constraint at `projects/bar` is enforced.
	//
	// Example 3 (RestoreDefault):
	//   `organizations/foo` has a `Policy` with:
	//     {enforced: true}
	//   `projects/bar` has a `Policy` with:
	//     {RestoreDefault: {}}
	// The constraint at `organizations/foo` is enforced.
	// The constraint at `projects/bar` is not enforced,
	// because
	// `constraint_default` for the `Constraint` is `ALLOW`.
	Enforced bool `json:"enforced,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Enforced") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Enforced") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BooleanPolicy) MarshalJSON() ([]byte, error) {
	type noMethod BooleanPolicy
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ClearOrgPolicyRequest: The request sent to the ClearOrgPolicy method.
type ClearOrgPolicyRequest struct {
	// Constraint: Name of the `Constraint` of the `Policy` to clear.
	Constraint string `json:"constraint,omitempty"`

	// Etag: The current version, for concurrency control. Not sending an
	// `etag`
	// will cause the `Policy` to be cleared blindly.
	Etag string `json:"etag,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Constraint") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Constraint") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ClearOrgPolicyRequest) MarshalJSON() ([]byte, error) {
	type noMethod ClearOrgPolicyRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Constraint: A `Constraint` describes a way in which a resource's
// configuration can be
// restricted. For example, it controls which cloud services can be
// activated
// across an organization, or whether a Compute Engine instance can
// have
// serial port connections established. `Constraints` can be configured
// by the
// organization's policy adminstrator to fit the needs of the
// organzation by
// setting Policies for `Constraints` at different locations in
// the
// organization's resource hierarchy. Policies are inherited down the
// resource
// hierarchy from higher levels, but can also be overridden. For details
// about
// the inheritance rules please read about
// Policies.
//
// `Constraints` have a default behavior determined by the
// `constraint_default`
// field, which is the enforcement behavior that is used in the absence
// of a
// `Policy` being defined or inherited for the resource in question.
type Constraint struct {
	// BooleanConstraint: Defines this constraint as being a
	// BooleanConstraint.
	BooleanConstraint *BooleanConstraint `json:"booleanConstraint,omitempty"`

	// ConstraintDefault: The evaluation behavior of this constraint in the
	// absense of 'Policy'.
	//
	// Possible values:
	//   "CONSTRAINT_DEFAULT_UNSPECIFIED" - This is only used for
	// distinguishing unset values and should never be
	// used.
	//   "ALLOW" - Indicate that all values are allowed for list
	// constraints.
	// Indicate that enforcement is off for boolean constraints.
	//   "DENY" - Indicate that all values are denied for list
	// constraints.
	// Indicate that enforcement is on for boolean constraints.
	ConstraintDefault string `json:"constraintDefault,omitempty"`

	// Description: Detailed description of what this `Constraint` controls
	// as well as how and
	// where it is enforced.
	//
	// Mutable.
	Description string `json:"description,omitempty"`

	// DisplayName: The human readable name.
	//
	// Mutable.
	DisplayName string `json:"displayName,omitempty"`

	// ListConstraint: Defines this constraint as being a ListConstraint.
	ListConstraint *ListConstraint `json:"listConstraint,omitempty"`

	// Name: Immutable value, required to globally be unique. For
	// example,
	// `constraints/serviceuser.services`
	Name string `json:"name,omitempty"`

	// Version: Version of the `Constraint`. Default version is 0;
	Version int64 `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BooleanConstraint")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BooleanConstraint") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Constraint) MarshalJSON() ([]byte, error) {
	type noMethod Constraint
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// FolderOperation: Metadata describing a long running folder operation
type FolderOperation struct {
	// DestinationParent: The resource name of the folder or organization we
	// are either creating
	// the folder under or moving the folder to.
	DestinationParent string `json:"destinationParent,omitempty"`

	// DisplayName: The display name of the folder.
	DisplayName string `json:"displayName,omitempty"`

	// OperationType: The type of this operation.
	//
	// Possible values:
	//   "OPERATION_TYPE_UNSPECIFIED" - Operation type not specified.
	//   "CREATE" - A create folder operation.
	//   "MOVE" - A move folder operation.
	OperationType string `json:"operationType,omitempty"`

	// SourceParent: The resource name of the folder's parent.
	// Only applicable when the operation_type is MOVE.
	SourceParent string `json:"sourceParent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DestinationParent")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DestinationParent") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FolderOperation) MarshalJSON() ([]byte, error) {
	type noMethod FolderOperation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FolderOperationError: A classification of the Folder Operation error.
type FolderOperationError struct {
	// ErrorMessageId: The type of operation error experienced.
	//
	// Possible values:
	//   "ERROR_TYPE_UNSPECIFIED" - The error type was unrecognized or
	// unspecified.
	//   "FOLDER_HEIGHT_VIOLATION" - The attempted action would violate the
	// max folder depth constraint.
	//   "MAX_CHILD_FOLDERS_VIOLATION" - The attempted action would violate
	// the max child folders constraint.
	//   "FOLDER_NAME_UNIQUENESS_VIOLATION" - The attempted action would
	// violate the locally-unique folder
	// display_name constraint.
	//   "RESOURCE_DELETED" - The resource being moved has been deleted.
	//   "PARENT_DELETED" - The resource a folder was being added to has
	// been deleted.
	//   "CYCLE_INTRODUCED_ERROR" - The attempted action would introduce
	// cycle in resource path.
	//   "FOLDER_BEING_MOVED" - The attempted action would move a folder
	// that is already being moved.
	//   "FOLDER_TO_DELETE_NON_EMPTY" - The folder the caller is trying to
	// delete contains active resources.
	ErrorMessageId string `json:"errorMessageId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ErrorMessageId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ErrorMessageId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FolderOperationError) MarshalJSON() ([]byte, error) {
	type noMethod FolderOperationError
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetAncestryRequest: The request sent to the
// GetAncestry
// method.
type GetAncestryRequest struct {
}

// GetAncestryResponse: Response from the GetAncestry method.
type GetAncestryResponse struct {
	// Ancestor: Ancestors are ordered from bottom to top of the resource
	// hierarchy. The
	// first ancestor is the project itself, followed by the project's
	// parent,
	// etc.
	Ancestor []*Ancestor `json:"ancestor,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Ancestor") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Ancestor") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetAncestryResponse) MarshalJSON() ([]byte, error) {
	type noMethod GetAncestryResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetEffectiveOrgPolicyRequest: The request sent to the
// GetEffectiveOrgPolicy method.
type GetEffectiveOrgPolicyRequest struct {
	// Constraint: The name of the `Constraint` to compute the effective
	// `Policy`.
	Constraint string `json:"constraint,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Constraint") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Constraint") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetEffectiveOrgPolicyRequest) MarshalJSON() ([]byte, error) {
	type noMethod GetEffectiveOrgPolicyRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetIamPolicyRequest: Request message for `GetIamPolicy` method.
type GetIamPolicyRequest struct {
}

// GetOrgPolicyRequest: The request sent to the GetOrgPolicy method.
type GetOrgPolicyRequest struct {
	// Constraint: Name of the `Constraint` to get the `Policy`.
	Constraint string `json:"constraint,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Constraint") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Constraint") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetOrgPolicyRequest) MarshalJSON() ([]byte, error) {
	type noMethod GetOrgPolicyRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListAvailableOrgPolicyConstraintsRequest: The request sent to the
// [ListAvailableOrgPolicyConstraints]
// google.cloud.OrgPolicy.v1.ListAvai
// lableOrgPolicyConstraints] method.
type ListAvailableOrgPolicyConstraintsRequest struct {
	// PageSize: Size of the pages to be returned. This is currently
	// unsupported and will
	// be ignored. The server may at any point start using this field to
	// limit
	// page size.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: Page token used to retrieve the next page. This is
	// currently unsupported
	// and will be ignored. The server may at any point start using this
	// field.
	PageToken string `json:"pageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PageSize") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PageSize") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListAvailableOrgPolicyConstraintsRequest) MarshalJSON() ([]byte, error) {
	type noMethod ListAvailableOrgPolicyConstraintsRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListAvailableOrgPolicyConstraintsResponse: The response returned from
// the ListAvailableOrgPolicyConstraints method.
// Returns all `Constraints` that could be set at this level of the
// hierarchy
// (contrast with the response from `ListPolicies`, which returns all
// policies
// which are set).
type ListAvailableOrgPolicyConstraintsResponse struct {
	// Constraints: The collection of constraints that are settable on the
	// request resource.
	Constraints []*Constraint `json:"constraints,omitempty"`

	// NextPageToken: Page token used to retrieve the next page. This is
	// currently not used.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Constraints") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Constraints") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListAvailableOrgPolicyConstraintsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListAvailableOrgPolicyConstraintsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListConstraint: A `Constraint` that allows or disallows a list of
// string values, which are
// configured by an Organization's policy administrator with a `Policy`.
type ListConstraint struct {
	// SuggestedValue: Optional. The Google Cloud Console will try to
	// default to a configuration
	// that matches the value specified in this `Constraint`.
	SuggestedValue string `json:"suggestedValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SuggestedValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SuggestedValue") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListConstraint) MarshalJSON() ([]byte, error) {
	type noMethod ListConstraint
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListOrgPoliciesRequest: The request sent to the ListOrgPolicies
// method.
type ListOrgPoliciesRequest struct {
	// PageSize: Size of the pages to be returned. This is currently
	// unsupported and will
	// be ignored. The server may at any point start using this field to
	// limit
	// page size.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: Page token used to retrieve the next page. This is
	// currently unsupported
	// and will be ignored. The server may at any point start using this
	// field.
	PageToken string `json:"pageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PageSize") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PageSize") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListOrgPoliciesRequest) MarshalJSON() ([]byte, error) {
	type noMethod ListOrgPoliciesRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListOrgPoliciesResponse: The response returned from the
// ListOrgPolicies method. It will be empty
// if no `Policies` are set on the resource.
type ListOrgPoliciesResponse struct {
	// NextPageToken: Page token used to retrieve the next page. This is
	// currently not used, but
	// the server may at any point start supplying a valid token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Policies: The `Policies` that are set on the resource. It will be
	// empty if no
	// `Policies` are set.
	Policies []*OrgPolicy `json:"policies,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListOrgPoliciesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListOrgPoliciesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListOrganizationsResponse: The response returned from the
// `ListOrganizations` method.
type ListOrganizationsResponse struct {
	// NextPageToken: A pagination token to be used to retrieve the next
	// page of results. If the
	// result is too large to fit within the page size specified in the
	// request,
	// this field will be set with a token that can be used to fetch the
	// next page
	// of results. If this field is empty, it indicates that this
	// response
	// contains the last page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Organizations: The list of Organizations that matched the list query,
	// possibly paginated.
	Organizations []*Organization `json:"organizations,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListOrganizationsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListOrganizationsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListPolicy: Used in `policy_type` to specify how `list_policy`
// behaves at this
// resource.
//
// A `ListPolicy` can define specific values that are allowed or denied
// by
// setting either the `allowed_values` or `denied_values` fields. It can
// also
// be used to allow or deny all values, by setting the `all_values`
// field. If
// `all_values` is `ALL_VALUES_UNSPECIFIED`, exactly one of
// `allowed_values`
// or `denied_values` must be set (attempting to set both or neither
// will
// result in a failed request). If `all_values` is set to either `ALLOW`
// or
// `DENY`, `allowed_values` and `denied_values` must be unset.
type ListPolicy struct {
	// AllValues: The policy all_values state.
	//
	// Possible values:
	//   "ALL_VALUES_UNSPECIFIED" - Indicates that either allowed_values or
	// denied_values must be set.
	//   "ALLOW" - A policy with this set allows all values.
	//   "DENY" - A policy with this set denies all values.
	AllValues string `json:"allValues,omitempty"`

	// AllowedValues: List of values allowed  at this resource. an only be
	// set if no values are
	// set for `denied_values` and `all_values` is set
	// to
	// `ALL_VALUES_UNSPECIFIED`.
	AllowedValues []string `json:"allowedValues,omitempty"`

	// DeniedValues: List of values denied at this resource. Can only be set
	// if no values are
	// set for `allowed_values` and `all_values` is set
	// to
	// `ALL_VALUES_UNSPECIFIED`.
	DeniedValues []string `json:"deniedValues,omitempty"`

	// InheritFromParent: Determines the inheritance behavior for this
	// `Policy`.
	//
	// By default, a `ListPolicy` set at a resource supercedes any `Policy`
	// set
	// anywhere up the resource hierarchy. However, if `inherit_from_parent`
	// is
	// set to `true`, then the values from the effective `Policy` of the
	// parent
	// resource are inherited, meaning the values set in this `Policy`
	// are
	// added to the values inherited up the hierarchy.
	//
	// Setting `Policy` hierarchies that inherit both allowed values and
	// denied
	// values isn't recommended in most circumstances to keep the
	// configuration
	// simple and understandable. However, it is possible to set a `Policy`
	// with
	// `allowed_values` set that inherits a `Policy` with `denied_values`
	// set.
	// In this case, the values that are allowed must be in `allowed_values`
	// and
	// not present in `denied_values`.
	//
	// For example, suppose you have a
	// `Constraint`
	// `constraints/serviceuser.services`, which has a `constraint_type`
	// of
	// `list_constraint`, and with `constraint_default` set to
	// `ALLOW`.
	// Suppose that at the Organization level, a `Policy` is applied
	// that
	// restricts the allowed API activations to {`E1`, `E2`}. Then, if
	// a
	// `Policy` is applied to a project below the Organization that
	// has
	// `inherit_from_parent` set to `false` and field all_values set to
	// DENY,
	// then an attempt to activate any API will be denied.
	//
	// The following examples demonstrate different possible
	// layerings:
	//
	// Example 1 (no inherited values):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: “E1” allowed_values:”E2”}
	//   ``projects/bar`` has `inherit_from_parent` `false` and values:
	//     {allowed_values: "E3" allowed_values: "E4"}
	// The accepted values at `organizations/foo` are `E1`, `E2`.
	// The accepted values at `projects/bar` are `E3`, and `E4`.
	//
	// Example 2 (inherited values):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: “E1” allowed_values:”E2”}
	//   `projects/bar` has a `Policy` with values:
	//     {value: “E3” value: ”E4” inherit_from_parent: true}
	// The accepted values at `organizations/foo` are `E1`, `E2`.
	// The accepted values at `projects/bar` are `E1`, `E2`, `E3`, and
	// `E4`.
	//
	// Example 3 (inheriting both allowed and denied values):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: "E1" allowed_values: "E2"}
	//   `projects/bar` has a `Policy` with:
	//     {denied_values: "E1"}
	// The accepted values at `organizations/foo` are `E1`, `E2`.
	// The value accepted at `projects/bar` is `E2`.
	//
	// Example 4 (RestoreDefault):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: “E1” allowed_values:”E2”}
	//   `projects/bar` has a `Policy` with values:
	//     {RestoreDefault: {}}
	// The accepted values at `organizations/foo` are `E1`, `E2`.
	// The accepted values at `projects/bar` are either all or none
	// depending on
	// the value of `constraint_default` (if `ALLOW`, all; if
	// `DENY`, none).
	//
	// Example 5 (no policy inherits parent policy):
	//   `organizations/foo` has no `Policy` set.
	//   `projects/bar` has no `Policy` set.
	// The accepted values at both levels are either all or none depending
	// on
	// the value of `constraint_default` (if `ALLOW`, all; if
	// `DENY`, none).
	//
	// Example 6 (ListConstraint allowing all):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: “E1” allowed_values: ”E2”}
	//   `projects/bar` has a `Policy` with:
	//     {all: ALLOW}
	// The accepted values at `organizations/foo` are `E1`, E2`.
	// Any value is accepted at `projects/bar`.
	//
	// Example 7 (ListConstraint allowing none):
	//   `organizations/foo` has a `Policy` with values:
	//     {allowed_values: “E1” allowed_values: ”E2”}
	//   `projects/bar` has a `Policy` with:
	//     {all: DENY}
	// The accepted values at `organizations/foo` are `E1`, E2`.
	// No value is accepted at `projects/bar`.
	InheritFromParent bool `json:"inheritFromParent,omitempty"`

	// SuggestedValue: Optional. The Google Cloud Console will try to
	// default to a configuration
	// that matches the value specified in this `Policy`. If
	// `suggested_value`
	// is not set, it will inherit the value specified higher in the
	// hierarchy,
	// unless `inherit_from_parent` is `false`.
	SuggestedValue string `json:"suggestedValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllValues") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AllValues") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListPolicy) MarshalJSON() ([]byte, error) {
	type noMethod ListPolicy
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListProjectsResponse: A page of the response received from
// the
// ListProjects
// method.
//
// A paginated response where more pages are available
// has
// `next_page_token` set. This token can be used in a subsequent request
// to
// retrieve the next request page.
type ListProjectsResponse struct {
	// NextPageToken: Pagination token.
	//
	// If the result set is too large to fit in a single response, this
	// token
	// is returned. It encodes the position of the current result
	// cursor.
	// Feeding this value into a new list request with the `page_token`
	// parameter
	// gives the next page of the results.
	//
	// When `next_page_token` is not filled in, there is no next page
	// and
	// the list returned is the last page in the result set.
	//
	// Pagination tokens have a limited lifetime.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Projects: The list of Projects that matched the list filter. This
	// list can
	// be paginated.
	Projects []*Project `json:"projects,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListProjectsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListProjectsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OrgPolicy: Defines a Cloud Organization `Policy` which is used to
// specify `Constraints`
// for configurations of Cloud Platform resources.
type OrgPolicy struct {
	// BooleanPolicy: For boolean `Constraints`, whether to enforce the
	// `Constraint` or not.
	BooleanPolicy *BooleanPolicy `json:"booleanPolicy,omitempty"`

	// Constraint: The name of the `Constraint` the `Policy` is configuring,
	// for example,
	// `constraints/serviceuser.services`.
	//
	// Immutable after creation.
	Constraint string `json:"constraint,omitempty"`

	// Etag: An opaque tag indicating the current version of the `Policy`,
	// used for
	// concurrency control.
	//
	// When the `Policy` is returned from either a `GetPolicy` or
	// a
	// `ListOrgPolicy` request, this `etag` indicates the version of the
	// current
	// `Policy` to use when executing a read-modify-write loop.
	//
	// When the `Policy` is returned from a `GetEffectivePolicy` request,
	// the
	// `etag` will be unset.
	//
	// When the `Policy` is used in a `SetOrgPolicy` method, use the `etag`
	// value
	// that was returned from a `GetOrgPolicy` request as part of
	// a
	// read-modify-write loop for concurrency control. Not setting the
	// `etag`in a
	// `SetOrgPolicy` request will result in an unconditional write of
	// the
	// `Policy`.
	Etag string `json:"etag,omitempty"`

	// ListPolicy: List of values either allowed or disallowed.
	ListPolicy *ListPolicy `json:"listPolicy,omitempty"`

	// RestoreDefault: Restores the default behavior of the constraint;
	// independent of
	// `Constraint` type.
	RestoreDefault *RestoreDefault `json:"restoreDefault,omitempty"`

	// UpdateTime: The time stamp the `Policy` was previously updated. This
	// is set by the
	// server, not specified by the caller, and represents the last time a
	// call to
	// `SetOrgPolicy` was made for that `Policy`. Any value set by the
	// client will
	// be ignored.
	UpdateTime string `json:"updateTime,omitempty"`

	// Version: Version of the `Policy`. Default version is 0;
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BooleanPolicy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BooleanPolicy") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OrgPolicy) MarshalJSON() ([]byte, error) {
	type noMethod OrgPolicy
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Organization: The root node in the resource hierarchy to which a
// particular entity's
// (e.g., company) resources belong.
type Organization struct {
	// CreationTime: Timestamp when the Organization was created. Assigned
	// by the server.
	// @OutputOnly
	CreationTime string `json:"creationTime,omitempty"`

	// DisplayName: A friendly string to be used to refer to the
	// Organization in the UI.
	// Assigned by the server, set to the primary domain of the G
	// Suite
	// customer that owns the organization.
	// @OutputOnly
	DisplayName string `json:"displayName,omitempty"`

	// LifecycleState: The organization's current lifecycle state. Assigned
	// by the server.
	// @OutputOnly
	//
	// Possible values:
	//   "LIFECYCLE_STATE_UNSPECIFIED" - Unspecified state.  This is only
	// useful for distinguishing unset values.
	//   "ACTIVE" - The normal and active state.
	//   "DELETE_REQUESTED" - The organization has been marked for deletion
	// by the user.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Name: Output Only. The resource name of the organization. This is
	// the
	// organization's relative path in the API. Its format
	// is
	// "organizations/[organization_id]". For example, "organizations/1234".
	Name string `json:"name,omitempty"`

	// OrganizationId: An immutable id for the Organization that is assigned
	// on creation. This
	// should be omitted when creating a new Organization.
	// This field is read-only.
	// This field is deprecated and will be removed in v1. Use name instead.
	OrganizationId string `json:"organizationId,omitempty"`

	// Owner: The owner of this Organization. The owner should be specified
	// on
	// creation. Once set, it cannot be changed.
	// This field is required.
	Owner *OrganizationOwner `json:"owner,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreationTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreationTime") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Organization) MarshalJSON() ([]byte, error) {
	type noMethod Organization
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OrganizationOwner: The entity that owns an Organization. The lifetime
// of the Organization and
// all of its descendants are bound to the `OrganizationOwner`. If
// the
// `OrganizationOwner` is deleted, the Organization and all its
// descendants will
// be deleted.
type OrganizationOwner struct {
	// DirectoryCustomerId: The Google for Work customer id used in the
	// Directory API.
	DirectoryCustomerId string `json:"directoryCustomerId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DirectoryCustomerId")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DirectoryCustomerId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *OrganizationOwner) MarshalJSON() ([]byte, error) {
	type noMethod OrganizationOwner
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Policy: Defines an Identity and Access Management (IAM) policy. It is
// used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `Binding` binds a list
// of
// `members` to a `role`, where the members can be user accounts, Google
// groups,
// Google domains, and service accounts. A `role` is a named list of
// permissions
// defined by IAM.
//
// **Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//
// "serviceAccount:my-other-app@appspot.gserviceaccount.com",
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam).
type Policy struct {
	// AuditConfigs: Specifies cloud audit logging configuration for this
	// policy.
	AuditConfigs []*AuditConfig `json:"auditConfigs,omitempty"`

	// Bindings: Associates a list of `members` to a `role`.
	// Multiple `bindings` must not be specified for the same
	// `role`.
	// `bindings` with no members will result in an error.
	Bindings []*Binding `json:"bindings,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a policy from overwriting each
	// other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform policy updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `getIamPolicy`,
	// and
	// systems are expected to put that etag in the request to
	// `setIamPolicy` to
	// ensure that their change will be applied to the same version of the
	// policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the
	// existing
	// policy is overwritten blindly.
	Etag string `json:"etag,omitempty"`

	// Version: Version of the `Policy`. The default version is 0.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AuditConfigs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditConfigs") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Policy) MarshalJSON() ([]byte, error) {
	type noMethod Policy
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Project: A Project is a high-level Google Cloud Platform entity.  It
// is a
// container for ACLs, APIs, App Engine Apps, VMs, and other
// Google Cloud Platform resources.
type Project struct {
	// CreateTime: Creation time.
	//
	// Read-only.
	CreateTime string `json:"createTime,omitempty"`

	// Labels: The labels associated with this Project.
	//
	// Label keys must be between 1 and 63 characters long and must
	// conform
	// to the following regular expression:
	// \[a-z\](\[-a-z0-9\]*\[a-z0-9\])?.
	//
	// Label values must be between 0 and 63 characters long and must
	// conform
	// to the regular expression (\[a-z\](\[-a-z0-9\]*\[a-z0-9\])?)?.
	//
	// No more than 256 labels can be associated with a given
	// resource.
	//
	// Clients should store labels in a representation such as JSON that
	// does not
	// depend on specific characters being disallowed.
	//
	// Example: <code>"environment" : "dev"</code>
	// Read-write.
	Labels map[string]string `json:"labels,omitempty"`

	// LifecycleState: The Project lifecycle state.
	//
	// Read-only.
	//
	// Possible values:
	//   "LIFECYCLE_STATE_UNSPECIFIED" - Unspecified state.  This is only
	// used/useful for distinguishing
	// unset values.
	//   "ACTIVE" - The normal and active state.
	//   "DELETE_REQUESTED" - The project has been marked for deletion by
	// the user
	// (by invoking DeleteProject)
	// or by the system (Google Cloud Platform).
	// This can generally be reversed by invoking UndeleteProject.
	//   "DELETE_IN_PROGRESS" - This lifecycle state is no longer used and
	// is not returned by the API.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Name: The user-assigned display name of the Project.
	// It must be 4 to 30 characters.
	// Allowed characters are: lowercase and uppercase letters,
	// numbers,
	// hyphen, single-quote, double-quote, space, and exclamation
	// point.
	//
	// Example: <code>My Project</code>
	// Read-write.
	Name string `json:"name,omitempty"`

	// Parent: An optional reference to a parent Resource.
	//
	// The only supported parent type is "organization". Once set, the
	// parent
	// cannot be modified. The `parent` can be set on creation or using
	// the
	// `UpdateProject` method; the end user must have
	// the
	// `resourcemanager.projects.create` permission on the
	// parent.
	//
	// Read-write.
	Parent *ResourceId `json:"parent,omitempty"`

	// ProjectId: The unique, user-assigned ID of the Project.
	// It must be 6 to 30 lowercase letters, digits, or hyphens.
	// It must start with a letter.
	// Trailing hyphens are prohibited.
	//
	// Example: <code>tokyo-rain-123</code>
	// Read-only after creation.
	ProjectId string `json:"projectId,omitempty"`

	// ProjectNumber: The number uniquely identifying the project.
	//
	// Example: <code>415104041262</code>
	// Read-only.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Project) MarshalJSON() ([]byte, error) {
	type noMethod Project
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ProjectCreationStatus: A status object which is used as the
// `metadata` field for the Operation
// returned by CreateProject. It provides insight for when significant
// phases of
// Project creation have completed.
type ProjectCreationStatus struct {
	// CreateTime: Creation time of the project creation workflow.
	CreateTime string `json:"createTime,omitempty"`

	// Gettable: True if the project can be retrieved using GetProject. No
	// other operations
	// on the project are guaranteed to work until the project creation
	// is
	// complete.
	Gettable bool `json:"gettable,omitempty"`

	// Ready: True if the project creation process is complete.
	Ready bool `json:"ready,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ProjectCreationStatus) MarshalJSON() ([]byte, error) {
	type noMethod ProjectCreationStatus
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResourceId: A container to reference an id for any resource type. A
// `resource` in Google
// Cloud Platform is a generic term for something you (a developer) may
// want to
// interact with through one of our API's. Some examples are an App
// Engine app,
// a Compute Engine instance, a Cloud SQL database, and so on.
type ResourceId struct {
	// Id: Required field for the type-specific id. This should correspond
	// to the id
	// used in the type-specific API's.
	Id string `json:"id,omitempty"`

	// Type: Required field representing the resource type this id is
	// for.
	// At present, the valid types are "project" and "organization".
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceId) MarshalJSON() ([]byte, error) {
	type noMethod ResourceId
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RestoreDefault: Ignores policies set above this resource and restores
// the
// `constraint_default` enforcement behavior of the specific
// `Constraint` at
// this resource.
//
// Suppose that `constraint_default` is set to `ALLOW` for
// the
// `Constraint` `constraints/serviceuser.services`. Suppose that
// organization
// foo.com sets a `Policy` at their Organization resource node that
// restricts
// the allowed service activations to deny all service activations.
// They
// could then set a `Policy` with the `policy_type` `restore_default`
// on
// several experimental projects, restoring the
// `constraint_default`
// enforcement of the `Constraint` for only those projects, allowing
// those
// projects to have all services activated.
type RestoreDefault struct {
}

// SetIamPolicyRequest: Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the
	// `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as
	// Projects)
	// might reject them.
	Policy *Policy `json:"policy,omitempty"`

	// UpdateMask: OPTIONAL: A FieldMask specifying which fields of the
	// policy to modify. Only
	// the fields in the mask will be modified. If no mask is provided,
	// the
	// following default mask is used:
	// paths: "bindings, etag"
	// This field is only used by Cloud IAM.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Policy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Policy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SetIamPolicyRequest) MarshalJSON() ([]byte, error) {
	type noMethod SetIamPolicyRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SetOrgPolicyRequest: The request sent to the SetOrgPolicyRequest
// method.
type SetOrgPolicyRequest struct {
	// Policy: `Policy` to set on the resource.
	Policy *OrgPolicy `json:"policy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Policy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Policy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SetOrgPolicyRequest) MarshalJSON() ([]byte, error) {
	type noMethod SetOrgPolicyRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestIamPermissionsRequest: Request message for `TestIamPermissions`
// method.
type TestIamPermissionsRequest struct {
	// Permissions: The set of permissions to check for the `resource`.
	// Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For
	// more
	// information see
	// [IAM
	// Overview](https://cloud.google.com/iam/docs/overview#permissions).
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestIamPermissionsRequest) MarshalJSON() ([]byte, error) {
	type noMethod TestIamPermissionsRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestIamPermissionsResponse: Response message for `TestIamPermissions`
// method.
type TestIamPermissionsResponse struct {
	// Permissions: A subset of `TestPermissionsRequest.permissions` that
	// the caller is
	// allowed.
	Permissions []string `json:"permissions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestIamPermissionsResponse) MarshalJSON() ([]byte, error) {
	type noMethod TestIamPermissionsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UndeleteProjectRequest: The request sent to the
// UndeleteProject
// method.
type UndeleteProjectRequest struct {
}

// method id "cloudresourcemanager.organizations.clearOrgPolicy":

type OrganizationsClearOrgPolicyCall struct {
	s                     *Service
	resource              string
	clearorgpolicyrequest *ClearOrgPolicyRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// ClearOrgPolicy: Clears a `Policy` from a resource.
func (r *OrganizationsService) ClearOrgPolicy(resource string, clearorgpolicyrequest *ClearOrgPolicyRequest) *OrganizationsClearOrgPolicyCall {
	c := &OrganizationsClearOrgPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.clearorgpolicyrequest = clearorgpolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsClearOrgPolicyCall) Fields(s ...googleapi.Field) *OrganizationsClearOrgPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsClearOrgPolicyCall) Context(ctx context.Context) *OrganizationsClearOrgPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsClearOrgPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsClearOrgPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.clearorgpolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:clearOrgPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.clearOrgPolicy" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OrganizationsClearOrgPolicyCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Clears a `Policy` from a resource.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:clearOrgPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.clearOrgPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "Name of the resource for the `Policy` to clear.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:clearOrgPolicy",
	//   "request": {
	//     "$ref": "ClearOrgPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.clearOrgPolicyV1":

type OrganizationsClearOrgPolicyV1Call struct {
	s                     *Service
	resource              string
	clearorgpolicyrequest *ClearOrgPolicyRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
	header_               http.Header
}

// ClearOrgPolicyV1: Clears a `Policy` from a resource.
func (r *OrganizationsService) ClearOrgPolicyV1(resource string, clearorgpolicyrequest *ClearOrgPolicyRequest) *OrganizationsClearOrgPolicyV1Call {
	c := &OrganizationsClearOrgPolicyV1Call{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.clearorgpolicyrequest = clearorgpolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsClearOrgPolicyV1Call) Fields(s ...googleapi.Field) *OrganizationsClearOrgPolicyV1Call {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsClearOrgPolicyV1Call) Context(ctx context.Context) *OrganizationsClearOrgPolicyV1Call {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsClearOrgPolicyV1Call) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsClearOrgPolicyV1Call) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.clearorgpolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:clearOrgPolicyV1")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.clearOrgPolicyV1" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OrganizationsClearOrgPolicyV1Call) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Clears a `Policy` from a resource.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:clearOrgPolicyV1",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.clearOrgPolicyV1",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "Name of the resource for the `Policy` to clear.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:clearOrgPolicyV1",
	//   "request": {
	//     "$ref": "ClearOrgPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.get":

type OrganizationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Fetches an Organization resource identified by the specified
// resource name.
func (r *OrganizationsService) Get(name string) *OrganizationsGetCall {
	c := &OrganizationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// OrganizationId sets the optional parameter "organizationId": The id
// of the Organization resource to fetch.
// This field is deprecated and will be removed in v1. Use name instead.
func (c *OrganizationsGetCall) OrganizationId(organizationId string) *OrganizationsGetCall {
	c.urlParams_.Set("organizationId", organizationId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsGetCall) Fields(s ...googleapi.Field) *OrganizationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OrganizationsGetCall) IfNoneMatch(entityTag string) *OrganizationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsGetCall) Context(ctx context.Context) *OrganizationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.get" call.
// Exactly one of *Organization or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Organization.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsGetCall) Do(opts ...googleapi.CallOption) (*Organization, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Organization{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Fetches an Organization resource identified by the specified resource name.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.organizations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the Organization to fetch, e.g. \"organizations/1234\".",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "organizationId": {
	//       "description": "The id of the Organization resource to fetch.\nThis field is deprecated and will be removed in v1. Use name instead.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Organization"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.getEffectiveOrgPolicy":

type OrganizationsGetEffectiveOrgPolicyCall struct {
	s                            *Service
	resource                     string
	geteffectiveorgpolicyrequest *GetEffectiveOrgPolicyRequest
	urlParams_                   gensupport.URLParams
	ctx_                         context.Context
	header_                      http.Header
}

// GetEffectiveOrgPolicy: Gets the effective `Policy` on a resource.
// This is the result of merging
// `Policies` in the resource hierarchy. The returned `Policy` will not
// have
// an `etag`set because it is a computed `Policy` across multiple
// resources.
func (r *OrganizationsService) GetEffectiveOrgPolicy(resource string, geteffectiveorgpolicyrequest *GetEffectiveOrgPolicyRequest) *OrganizationsGetEffectiveOrgPolicyCall {
	c := &OrganizationsGetEffectiveOrgPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.geteffectiveorgpolicyrequest = geteffectiveorgpolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsGetEffectiveOrgPolicyCall) Fields(s ...googleapi.Field) *OrganizationsGetEffectiveOrgPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsGetEffectiveOrgPolicyCall) Context(ctx context.Context) *OrganizationsGetEffectiveOrgPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsGetEffectiveOrgPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsGetEffectiveOrgPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.geteffectiveorgpolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:getEffectiveOrgPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.getEffectiveOrgPolicy" call.
// Exactly one of *OrgPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *OrgPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsGetEffectiveOrgPolicyCall) Do(opts ...googleapi.CallOption) (*OrgPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &OrgPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the effective `Policy` on a resource. This is the result of merging\n`Policies` in the resource hierarchy. The returned `Policy` will not have\nan `etag`set because it is a computed `Policy` across multiple resources.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:getEffectiveOrgPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.getEffectiveOrgPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "The name of the resource to start computing the effective `Policy`.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:getEffectiveOrgPolicy",
	//   "request": {
	//     "$ref": "GetEffectiveOrgPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "OrgPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.getEffectiveOrgPolicyV1":

type OrganizationsGetEffectiveOrgPolicyV1Call struct {
	s                            *Service
	resource                     string
	geteffectiveorgpolicyrequest *GetEffectiveOrgPolicyRequest
	urlParams_                   gensupport.URLParams
	ctx_                         context.Context
	header_                      http.Header
}

// GetEffectiveOrgPolicyV1: Gets the effective `Policy` on a resource.
// This is the result of merging
// `Policies` in the resource hierarchy. The returned `Policy` will not
// have
// an `etag`set because it is a computed `Policy` across multiple
// resources.
func (r *OrganizationsService) GetEffectiveOrgPolicyV1(resource string, geteffectiveorgpolicyrequest *GetEffectiveOrgPolicyRequest) *OrganizationsGetEffectiveOrgPolicyV1Call {
	c := &OrganizationsGetEffectiveOrgPolicyV1Call{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.geteffectiveorgpolicyrequest = geteffectiveorgpolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsGetEffectiveOrgPolicyV1Call) Fields(s ...googleapi.Field) *OrganizationsGetEffectiveOrgPolicyV1Call {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsGetEffectiveOrgPolicyV1Call) Context(ctx context.Context) *OrganizationsGetEffectiveOrgPolicyV1Call {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsGetEffectiveOrgPolicyV1Call) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsGetEffectiveOrgPolicyV1Call) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.geteffectiveorgpolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:getEffectiveOrgPolicyV1")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.getEffectiveOrgPolicyV1" call.
// Exactly one of *OrgPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *OrgPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsGetEffectiveOrgPolicyV1Call) Do(opts ...googleapi.CallOption) (*OrgPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &OrgPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the effective `Policy` on a resource. This is the result of merging\n`Policies` in the resource hierarchy. The returned `Policy` will not have\nan `etag`set because it is a computed `Policy` across multiple resources.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:getEffectiveOrgPolicyV1",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.getEffectiveOrgPolicyV1",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "The name of the resource to start computing the effective `Policy`.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:getEffectiveOrgPolicyV1",
	//   "request": {
	//     "$ref": "GetEffectiveOrgPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "OrgPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.getIamPolicy":

type OrganizationsGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetIamPolicy: Gets the access control policy for an Organization
// resource. May be empty
// if no such policy or resource exists. The `resource` field should be
// the
// organization's resource name, e.g. "organizations/123".
func (r *OrganizationsService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *OrganizationsGetIamPolicyCall {
	c := &OrganizationsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsGetIamPolicyCall) Fields(s ...googleapi.Field) *OrganizationsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsGetIamPolicyCall) Context(ctx context.Context) *OrganizationsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OrganizationsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for an Organization resource. May be empty\nif no such policy or resource exists. The `resource` field should be the\norganization's resource name, e.g. \"organizations/123\".",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.getOrgPolicy":

type OrganizationsGetOrgPolicyCall struct {
	s                   *Service
	resource            string
	getorgpolicyrequest *GetOrgPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetOrgPolicy: Gets a `Policy` on a resource.
//
// If no `Policy` is set on the resource, a `Policy` is returned with
// default
// values including `POLICY_TYPE_NOT_SET` for the `policy_type oneof`.
// The
// `etag` value can be used with `SetOrgPolicy()` to create or update
// a
// `Policy` during read-modify-write.
func (r *OrganizationsService) GetOrgPolicy(resource string, getorgpolicyrequest *GetOrgPolicyRequest) *OrganizationsGetOrgPolicyCall {
	c := &OrganizationsGetOrgPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getorgpolicyrequest = getorgpolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsGetOrgPolicyCall) Fields(s ...googleapi.Field) *OrganizationsGetOrgPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsGetOrgPolicyCall) Context(ctx context.Context) *OrganizationsGetOrgPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsGetOrgPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsGetOrgPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getorgpolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:getOrgPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.getOrgPolicy" call.
// Exactly one of *OrgPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *OrgPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsGetOrgPolicyCall) Do(opts ...googleapi.CallOption) (*OrgPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &OrgPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a `Policy` on a resource.\n\nIf no `Policy` is set on the resource, a `Policy` is returned with default\nvalues including `POLICY_TYPE_NOT_SET` for the `policy_type oneof`. The\n`etag` value can be used with `SetOrgPolicy()` to create or update a\n`Policy` during read-modify-write.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:getOrgPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.getOrgPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "Name of the resource the `Policy` is set on.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:getOrgPolicy",
	//   "request": {
	//     "$ref": "GetOrgPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "OrgPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.getOrgPolicyV1":

type OrganizationsGetOrgPolicyV1Call struct {
	s                   *Service
	resource            string
	getorgpolicyrequest *GetOrgPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetOrgPolicyV1: Gets a `Policy` on a resource.
//
// If no `Policy` is set on the resource, a `Policy` is returned with
// default
// values including `POLICY_TYPE_NOT_SET` for the `policy_type oneof`.
// The
// `etag` value can be used with `SetOrgPolicy()` to create or update
// a
// `Policy` during read-modify-write.
func (r *OrganizationsService) GetOrgPolicyV1(resource string, getorgpolicyrequest *GetOrgPolicyRequest) *OrganizationsGetOrgPolicyV1Call {
	c := &OrganizationsGetOrgPolicyV1Call{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getorgpolicyrequest = getorgpolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsGetOrgPolicyV1Call) Fields(s ...googleapi.Field) *OrganizationsGetOrgPolicyV1Call {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsGetOrgPolicyV1Call) Context(ctx context.Context) *OrganizationsGetOrgPolicyV1Call {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsGetOrgPolicyV1Call) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsGetOrgPolicyV1Call) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getorgpolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:getOrgPolicyV1")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.getOrgPolicyV1" call.
// Exactly one of *OrgPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *OrgPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsGetOrgPolicyV1Call) Do(opts ...googleapi.CallOption) (*OrgPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &OrgPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a `Policy` on a resource.\n\nIf no `Policy` is set on the resource, a `Policy` is returned with default\nvalues including `POLICY_TYPE_NOT_SET` for the `policy_type oneof`. The\n`etag` value can be used with `SetOrgPolicy()` to create or update a\n`Policy` during read-modify-write.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:getOrgPolicyV1",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.getOrgPolicyV1",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "Name of the resource the `Policy` is set on.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:getOrgPolicyV1",
	//   "request": {
	//     "$ref": "GetOrgPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "OrgPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.list":

type OrganizationsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists Organization resources that are visible to the user and
// satisfy
// the specified filter. This method returns Organizations in an
// unspecified
// order. New Organizations do not necessarily appear at the end of the
// list.
func (r *OrganizationsService) List() *OrganizationsListCall {
	c := &OrganizationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Filter sets the optional parameter "filter": An optional query string
// used to filter the Organizations to return in
// the response. Filter rules are case-insensitive.
//
//
// Organizations may be filtered by `owner.directoryCustomerId` or
// by
// `domain`, where the domain is a Google for Work domain, for
// example:
//
// |Filter|Description|
// |------|-----------|
// |owner.directorycu
// stomerid:123456789|Organizations with `owner.directory_customer_id`
// equal to `123456789`.|
// |domain:google.com|Organizations corresponding to the domain
// `google.com`.|
//
// This field is optional.
func (c *OrganizationsListCall) Filter(filter string) *OrganizationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of Organizations to return in the response.
// This field is optional.
func (c *OrganizationsListCall) PageSize(pageSize int64) *OrganizationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A pagination token
// returned from a previous call to `ListOrganizations`
// that indicates from where listing should continue.
// This field is optional.
func (c *OrganizationsListCall) PageToken(pageToken string) *OrganizationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsListCall) Fields(s ...googleapi.Field) *OrganizationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OrganizationsListCall) IfNoneMatch(entityTag string) *OrganizationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsListCall) Context(ctx context.Context) *OrganizationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/organizations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.list" call.
// Exactly one of *ListOrganizationsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListOrganizationsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OrganizationsListCall) Do(opts ...googleapi.CallOption) (*ListOrganizationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListOrganizationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists Organization resources that are visible to the user and satisfy\nthe specified filter. This method returns Organizations in an unspecified\norder. New Organizations do not necessarily appear at the end of the list.",
	//   "flatPath": "v1beta1/organizations",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.organizations.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "filter": {
	//       "description": "An optional query string used to filter the Organizations to return in\nthe response. Filter rules are case-insensitive.\n\n\nOrganizations may be filtered by `owner.directoryCustomerId` or by\n`domain`, where the domain is a Google for Work domain, for example:\n\n|Filter|Description|\n|------|-----------|\n|owner.directorycustomerid:123456789|Organizations with `owner.directory_customer_id` equal to `123456789`.|\n|domain:google.com|Organizations corresponding to the domain `google.com`.|\n\nThis field is optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of Organizations to return in the response.\nThis field is optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A pagination token returned from a previous call to `ListOrganizations`\nthat indicates from where listing should continue.\nThis field is optional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/organizations",
	//   "response": {
	//     "$ref": "ListOrganizationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OrganizationsListCall) Pages(ctx context.Context, f func(*ListOrganizationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudresourcemanager.organizations.listAvailableOrgPolicyConstraints":

type OrganizationsListAvailableOrgPolicyConstraintsCall struct {
	s                                        *Service
	resource                                 string
	listavailableorgpolicyconstraintsrequest *ListAvailableOrgPolicyConstraintsRequest
	urlParams_                               gensupport.URLParams
	ctx_                                     context.Context
	header_                                  http.Header
}

// ListAvailableOrgPolicyConstraints: Lists `Constraints` that could be
// applied on the specified resource.
func (r *OrganizationsService) ListAvailableOrgPolicyConstraints(resource string, listavailableorgpolicyconstraintsrequest *ListAvailableOrgPolicyConstraintsRequest) *OrganizationsListAvailableOrgPolicyConstraintsCall {
	c := &OrganizationsListAvailableOrgPolicyConstraintsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.listavailableorgpolicyconstraintsrequest = listavailableorgpolicyconstraintsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsListAvailableOrgPolicyConstraintsCall) Fields(s ...googleapi.Field) *OrganizationsListAvailableOrgPolicyConstraintsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsListAvailableOrgPolicyConstraintsCall) Context(ctx context.Context) *OrganizationsListAvailableOrgPolicyConstraintsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsListAvailableOrgPolicyConstraintsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsListAvailableOrgPolicyConstraintsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.listavailableorgpolicyconstraintsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:listAvailableOrgPolicyConstraints")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.listAvailableOrgPolicyConstraints" call.
// Exactly one of *ListAvailableOrgPolicyConstraintsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *ListAvailableOrgPolicyConstraintsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *OrganizationsListAvailableOrgPolicyConstraintsCall) Do(opts ...googleapi.CallOption) (*ListAvailableOrgPolicyConstraintsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListAvailableOrgPolicyConstraintsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists `Constraints` that could be applied on the specified resource.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:listAvailableOrgPolicyConstraints",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.listAvailableOrgPolicyConstraints",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "Name of the resource to list `Constraints` for.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:listAvailableOrgPolicyConstraints",
	//   "request": {
	//     "$ref": "ListAvailableOrgPolicyConstraintsRequest"
	//   },
	//   "response": {
	//     "$ref": "ListAvailableOrgPolicyConstraintsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OrganizationsListAvailableOrgPolicyConstraintsCall) Pages(ctx context.Context, f func(*ListAvailableOrgPolicyConstraintsResponse) error) error {
	c.ctx_ = ctx
	defer func(pt string) { c.listavailableorgpolicyconstraintsrequest.PageToken = pt }(c.listavailableorgpolicyconstraintsrequest.PageToken) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.listavailableorgpolicyconstraintsrequest.PageToken = x.NextPageToken
	}
}

// method id "cloudresourcemanager.organizations.listOrgPolicies":

type OrganizationsListOrgPoliciesCall struct {
	s                      *Service
	resource               string
	listorgpoliciesrequest *ListOrgPoliciesRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// ListOrgPolicies: Lists all the `Policies` set for a particular
// resource.
func (r *OrganizationsService) ListOrgPolicies(resource string, listorgpoliciesrequest *ListOrgPoliciesRequest) *OrganizationsListOrgPoliciesCall {
	c := &OrganizationsListOrgPoliciesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.listorgpoliciesrequest = listorgpoliciesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsListOrgPoliciesCall) Fields(s ...googleapi.Field) *OrganizationsListOrgPoliciesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsListOrgPoliciesCall) Context(ctx context.Context) *OrganizationsListOrgPoliciesCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsListOrgPoliciesCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsListOrgPoliciesCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.listorgpoliciesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:listOrgPolicies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.listOrgPolicies" call.
// Exactly one of *ListOrgPoliciesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOrgPoliciesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OrganizationsListOrgPoliciesCall) Do(opts ...googleapi.CallOption) (*ListOrgPoliciesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListOrgPoliciesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the `Policies` set for a particular resource.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:listOrgPolicies",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.listOrgPolicies",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "Name of the resource to list Policies for.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:listOrgPolicies",
	//   "request": {
	//     "$ref": "ListOrgPoliciesRequest"
	//   },
	//   "response": {
	//     "$ref": "ListOrgPoliciesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OrganizationsListOrgPoliciesCall) Pages(ctx context.Context, f func(*ListOrgPoliciesResponse) error) error {
	c.ctx_ = ctx
	defer func(pt string) { c.listorgpoliciesrequest.PageToken = pt }(c.listorgpoliciesrequest.PageToken) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.listorgpoliciesrequest.PageToken = x.NextPageToken
	}
}

// method id "cloudresourcemanager.organizations.setIamPolicy":

type OrganizationsSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetIamPolicy: Sets the access control policy on an Organization
// resource. Replaces any
// existing policy. The `resource` field should be the organization's
// resource
// name, e.g. "organizations/123".
func (r *OrganizationsService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *OrganizationsSetIamPolicyCall {
	c := &OrganizationsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsSetIamPolicyCall) Fields(s ...googleapi.Field) *OrganizationsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsSetIamPolicyCall) Context(ctx context.Context) *OrganizationsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OrganizationsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on an Organization resource. Replaces any\nexisting policy. The `resource` field should be the organization's resource\nname, e.g. \"organizations/123\".",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.setOrgPolicy":

type OrganizationsSetOrgPolicyCall struct {
	s                   *Service
	resource            string
	setorgpolicyrequest *SetOrgPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetOrgPolicy: Updates the specified `Policy` on the resource. Creates
// a new `Policy` for
// that `Constraint` on the resource if one does not exist.
//
// Not supplying an `etag` on the request `Policy` results in an
// unconditional
// write of the `Policy`.
func (r *OrganizationsService) SetOrgPolicy(resource string, setorgpolicyrequest *SetOrgPolicyRequest) *OrganizationsSetOrgPolicyCall {
	c := &OrganizationsSetOrgPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setorgpolicyrequest = setorgpolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsSetOrgPolicyCall) Fields(s ...googleapi.Field) *OrganizationsSetOrgPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsSetOrgPolicyCall) Context(ctx context.Context) *OrganizationsSetOrgPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsSetOrgPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsSetOrgPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setorgpolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:setOrgPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.setOrgPolicy" call.
// Exactly one of *OrgPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *OrgPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsSetOrgPolicyCall) Do(opts ...googleapi.CallOption) (*OrgPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &OrgPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified `Policy` on the resource. Creates a new `Policy` for\nthat `Constraint` on the resource if one does not exist.\n\nNot supplying an `etag` on the request `Policy` results in an unconditional\nwrite of the `Policy`.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:setOrgPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.setOrgPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "Resource name of the resource to attach the `Policy`.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:setOrgPolicy",
	//   "request": {
	//     "$ref": "SetOrgPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "OrgPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.setOrgPolicyV1":

type OrganizationsSetOrgPolicyV1Call struct {
	s                   *Service
	resource            string
	setorgpolicyrequest *SetOrgPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetOrgPolicyV1: Updates the specified `Policy` on the resource.
// Creates a new `Policy` for
// that `Constraint` on the resource if one does not exist.
//
// Not supplying an `etag` on the request `Policy` results in an
// unconditional
// write of the `Policy`.
func (r *OrganizationsService) SetOrgPolicyV1(resource string, setorgpolicyrequest *SetOrgPolicyRequest) *OrganizationsSetOrgPolicyV1Call {
	c := &OrganizationsSetOrgPolicyV1Call{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setorgpolicyrequest = setorgpolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsSetOrgPolicyV1Call) Fields(s ...googleapi.Field) *OrganizationsSetOrgPolicyV1Call {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsSetOrgPolicyV1Call) Context(ctx context.Context) *OrganizationsSetOrgPolicyV1Call {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsSetOrgPolicyV1Call) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsSetOrgPolicyV1Call) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setorgpolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:setOrgPolicyV1")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.setOrgPolicyV1" call.
// Exactly one of *OrgPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *OrgPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsSetOrgPolicyV1Call) Do(opts ...googleapi.CallOption) (*OrgPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &OrgPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified `Policy` on the resource. Creates a new `Policy` for\nthat `Constraint` on the resource if one does not exist.\n\nNot supplying an `etag` on the request `Policy` results in an unconditional\nwrite of the `Policy`.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:setOrgPolicyV1",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.setOrgPolicyV1",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "Resource name of the resource to attach the `Policy`.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:setOrgPolicyV1",
	//   "request": {
	//     "$ref": "SetOrgPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "OrgPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.testIamPermissions":

type OrganizationsTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified Organization.
// The `resource` field should be the organization's resource name,
// e.g. "organizations/123".
func (r *OrganizationsService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *OrganizationsTestIamPermissionsCall {
	c := &OrganizationsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsTestIamPermissionsCall) Fields(s ...googleapi.Field) *OrganizationsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsTestIamPermissionsCall) Context(ctx context.Context) *OrganizationsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OrganizationsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified Organization.\nThe `resource` field should be the organization's resource name,\ne.g. \"organizations/123\".",
	//   "flatPath": "v1beta1/organizations/{organizationsId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.update":

type OrganizationsUpdateCall struct {
	s            *Service
	name         string
	organization *Organization
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Update: Updates an Organization resource identified by the specified
// resource name.
func (r *OrganizationsService) Update(name string, organization *Organization) *OrganizationsUpdateCall {
	c := &OrganizationsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.organization = organization
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsUpdateCall) Fields(s ...googleapi.Field) *OrganizationsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsUpdateCall) Context(ctx context.Context) *OrganizationsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OrganizationsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OrganizationsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.organization)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.organizations.update" call.
// Exactly one of *Organization or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Organization.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OrganizationsUpdateCall) Do(opts ...googleapi.CallOption) (*Organization, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Organization{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an Organization resource identified by the specified resource name.",
	//   "flatPath": "v1beta1/organizations/{organizationsId}",
	//   "httpMethod": "PUT",
	//   "id": "cloudresourcemanager.organizations.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Output Only. The resource name of the organization. This is the\norganization's relative path in the API. Its format is\n\"organizations/[organization_id]\". For example, \"organizations/1234\".",
	//       "location": "path",
	//       "pattern": "^organizations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "request": {
	//     "$ref": "Organization"
	//   },
	//   "response": {
	//     "$ref": "Organization"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.create":

type ProjectsCreateCall struct {
	s          *Service
	project    *Project
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a Project resource.
//
// Initially, the Project resource is owned by its creator
// exclusively.
// The creator can later grant permission to others to read or update
// the
// Project.
//
// Several APIs are activated automatically for the Project,
// including
// Google Cloud Storage.
func (r *ProjectsService) Create(project *Project) *ProjectsCreateCall {
	c := &ProjectsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	return c
}

// UseLegacyStack sets the optional parameter "useLegacyStack": A safety
// hatch to opt out of the new reliable project creation process.
func (c *ProjectsCreateCall) UseLegacyStack(useLegacyStack bool) *ProjectsCreateCall {
	c.urlParams_.Set("useLegacyStack", fmt.Sprint(useLegacyStack))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsCreateCall) Fields(s ...googleapi.Field) *ProjectsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsCreateCall) Context(ctx context.Context) *ProjectsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.project)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.create" call.
// Exactly one of *Project or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Project.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsCreateCall) Do(opts ...googleapi.CallOption) (*Project, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Project{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a Project resource.\n\nInitially, the Project resource is owned by its creator exclusively.\nThe creator can later grant permission to others to read or update the\nProject.\n\nSeveral APIs are activated automatically for the Project, including\nGoogle Cloud Storage.",
	//   "flatPath": "v1beta1/projects",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.create",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "useLegacyStack": {
	//       "description": "A safety hatch to opt out of the new reliable project creation process.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "v1beta1/projects",
	//   "request": {
	//     "$ref": "Project"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.delete":

type ProjectsDeleteCall struct {
	s          *Service
	projectId  string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Marks the Project identified by the specified
// `project_id` (for example, `my-project-123`) for deletion.
// This method will only affect the Project if the following criteria
// are met:
//
// + The Project does not have a billing account associated with it.
// + The Project has a lifecycle state of
// ACTIVE.
//
// This method changes the Project's lifecycle state from
// ACTIVE
// to DELETE_REQUESTED.
// The deletion starts at an unspecified time, at which point the
// project is
// no longer accessible.
//
// Until the deletion completes, you can check the lifecycle
// state
// checked by retrieving the Project with GetProject,
// and the Project remains visible to ListProjects.
// However, you cannot update the project.
//
// After the deletion completes, the Project is not retrievable by
// the  GetProject and
// ListProjects methods.
//
// The caller must have modify permissions for this Project.
func (r *ProjectsService) Delete(projectId string) *ProjectsDeleteCall {
	c := &ProjectsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeleteCall) Fields(s ...googleapi.Field) *ProjectsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDeleteCall) Context(ctx context.Context) *ProjectsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Marks the Project identified by the specified\n`project_id` (for example, `my-project-123`) for deletion.\nThis method will only affect the Project if the following criteria are met:\n\n+ The Project does not have a billing account associated with it.\n+ The Project has a lifecycle state of\nACTIVE.\n\nThis method changes the Project's lifecycle state from\nACTIVE\nto DELETE_REQUESTED.\nThe deletion starts at an unspecified time, at which point the project is\nno longer accessible.\n\nUntil the deletion completes, you can check the lifecycle state\nchecked by retrieving the Project with GetProject,\nand the Project remains visible to ListProjects.\nHowever, you cannot update the project.\n\nAfter the deletion completes, the Project is not retrievable by\nthe  GetProject and\nListProjects methods.\n\nThe caller must have modify permissions for this Project.",
	//   "flatPath": "v1beta1/projects/{projectId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudresourcemanager.projects.delete",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Project ID (for example, `foo-bar-123`).\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.get":

type ProjectsGetCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Retrieves the Project identified by the specified
// `project_id` (for example, `my-project-123`).
//
// The caller must have read permissions for this Project.
func (r *ProjectsService) Get(projectId string) *ProjectsGetCall {
	c := &ProjectsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetCall) Fields(s ...googleapi.Field) *ProjectsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetCall) IfNoneMatch(entityTag string) *ProjectsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetCall) Context(ctx context.Context) *ProjectsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.get" call.
// Exactly one of *Project or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Project.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGetCall) Do(opts ...googleapi.CallOption) (*Project, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Project{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the Project identified by the specified\n`project_id` (for example, `my-project-123`).\n\nThe caller must have read permissions for this Project.",
	//   "flatPath": "v1beta1/projects/{projectId}",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.projects.get",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Project ID (for example, `my-project-123`).\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}",
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.getAncestry":

type ProjectsGetAncestryCall struct {
	s                  *Service
	projectId          string
	getancestryrequest *GetAncestryRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// GetAncestry: Gets a list of ancestors in the resource hierarchy for
// the Project
// identified by the specified `project_id` (for example,
// `my-project-123`).
//
// The caller must have read permissions for this Project.
func (r *ProjectsService) GetAncestry(projectId string, getancestryrequest *GetAncestryRequest) *ProjectsGetAncestryCall {
	c := &ProjectsGetAncestryCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.getancestryrequest = getancestryrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetAncestryCall) Fields(s ...googleapi.Field) *ProjectsGetAncestryCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetAncestryCall) Context(ctx context.Context) *ProjectsGetAncestryCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetAncestryCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetAncestryCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getancestryrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}:getAncestry")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.getAncestry" call.
// Exactly one of *GetAncestryResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GetAncestryResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsGetAncestryCall) Do(opts ...googleapi.CallOption) (*GetAncestryResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GetAncestryResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a list of ancestors in the resource hierarchy for the Project\nidentified by the specified `project_id` (for example, `my-project-123`).\n\nThe caller must have read permissions for this Project.",
	//   "flatPath": "v1beta1/projects/{projectId}:getAncestry",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.getAncestry",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Project ID (for example, `my-project-123`).\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}:getAncestry",
	//   "request": {
	//     "$ref": "GetAncestryRequest"
	//   },
	//   "response": {
	//     "$ref": "GetAncestryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.getIamPolicy":

type ProjectsGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetIamPolicy: Returns the IAM access control policy for the specified
// Project.
// Permission is denied if the policy or the resource does not exist.
func (r *ProjectsService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsGetIamPolicyCall {
	c := &ProjectsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetIamPolicyCall) Context(ctx context.Context) *ProjectsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the IAM access control policy for the specified Project.\nPermission is denied if the policy or the resource does not exist.",
	//   "flatPath": "v1beta1/projects/{resource}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.list":

type ProjectsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists Projects that are visible to the user and satisfy
// the
// specified filter. This method returns Projects in an unspecified
// order.
// New Projects do not necessarily appear at the end of the list.
func (r *ProjectsService) List() *ProjectsListCall {
	c := &ProjectsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Filter sets the optional parameter "filter": An expression for
// filtering the results of the request.  Filter rules are
// case insensitive. The fields eligible for filtering are:
//
// + `name`
// + `id`
// + <code>labels.<em>key</em></code> where *key* is the name of a
// label
//
// Some examples of using labels as
// filters:
//
// |Filter|Description|
// |------|-----------|
// |name:*|The project has a name.|
// |name:Howl|The project's name is `Howl` or
// `howl`.|
// |name:HOWL|Equivalent to above.|
// |NAME:howl|Equivalent to above.|
// |labels.color:*|The project has the label
// `color`.|
// |labels.color:red|The project's label `color` has the value
// `red`.|
// |labels.color:red&nbsp;labels.size:big|The project's label `color`
// has the value `red` and its label `size` has the value `big`.
func (c *ProjectsListCall) Filter(filter string) *ProjectsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of Projects to return in the response.
// The server can return fewer Projects than requested.
// If unspecified, server picks an appropriate default.
func (c *ProjectsListCall) PageSize(pageSize int64) *ProjectsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A pagination token
// returned from a previous call to ListProjects
// that indicates from where listing should continue.
func (c *ProjectsListCall) PageToken(pageToken string) *ProjectsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsListCall) Fields(s ...googleapi.Field) *ProjectsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsListCall) IfNoneMatch(entityTag string) *ProjectsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsListCall) Context(ctx context.Context) *ProjectsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.list" call.
// Exactly one of *ListProjectsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListProjectsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsListCall) Do(opts ...googleapi.CallOption) (*ListProjectsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListProjectsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists Projects that are visible to the user and satisfy the\nspecified filter. This method returns Projects in an unspecified order.\nNew Projects do not necessarily appear at the end of the list.",
	//   "flatPath": "v1beta1/projects",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.projects.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "filter": {
	//       "description": "An expression for filtering the results of the request.  Filter rules are\ncase insensitive. The fields eligible for filtering are:\n\n+ `name`\n+ `id`\n+ \u003ccode\u003elabels.\u003cem\u003ekey\u003c/em\u003e\u003c/code\u003e where *key* is the name of a label\n\nSome examples of using labels as filters:\n\n|Filter|Description|\n|------|-----------|\n|name:*|The project has a name.|\n|name:Howl|The project's name is `Howl` or `howl`.|\n|name:HOWL|Equivalent to above.|\n|NAME:howl|Equivalent to above.|\n|labels.color:*|The project has the label `color`.|\n|labels.color:red|The project's label `color` has the value `red`.|\n|labels.color:red\u0026nbsp;labels.size:big|The project's label `color` has the value `red` and its label `size` has the value `big`.\n\nOptional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of Projects to return in the response.\nThe server can return fewer Projects than requested.\nIf unspecified, server picks an appropriate default.\n\nOptional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A pagination token returned from a previous call to ListProjects\nthat indicates from where listing should continue.\n\nOptional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects",
	//   "response": {
	//     "$ref": "ListProjectsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsListCall) Pages(ctx context.Context, f func(*ListProjectsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudresourcemanager.projects.setIamPolicy":

type ProjectsSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetIamPolicy: Sets the IAM access control policy for the specified
// Project. Replaces
// any existing policy.
//
// The following constraints apply when using `setIamPolicy()`:
//
// + Project does not support `allUsers` and `allAuthenticatedUsers`
// as
// `members` in a `Binding` of a `Policy`.
//
// + The owner role can be granted only to `user` and
// `serviceAccount`.
//
// + Service accounts can be made owners of a project directly
// without any restrictions. However, to be added as an owner, a user
// must be
// invited via Cloud Platform console and must accept the invitation.
//
// + A user cannot be granted the owner role using `setIamPolicy()`. The
// user
// must be granted the owner role using the Cloud Platform Console and
// must
// explicitly accept the invitation.
//
// + Invitations to grant the owner role cannot be sent
// using
// `setIamPolicy()`; they must be sent only using the Cloud Platform
// Console.
//
// + Membership changes that leave the project without any owners that
// have
// accepted the Terms of Service (ToS) will be rejected.
//
// + There must be at least one owner who has accepted the Terms
// of
// Service (ToS) agreement in the policy. Calling `setIamPolicy()`
// to
// remove the last ToS-accepted owner from the policy will fail.
// This
// restriction also applies to legacy projects that no longer have
// owners
// who have accepted the ToS. Edits to IAM policies will be rejected
// until
// the lack of a ToS-accepting owner is rectified.
//
// + Calling this method requires enabling the App Engine Admin
// API.
//
// Note: Removing service accounts from policies or changing their
// roles
// can render services completely inoperable. It is important to
// understand
// how the service account is being used before removing or updating
// its
// roles.
func (r *ProjectsService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsSetIamPolicyCall {
	c := &ProjectsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsSetIamPolicyCall) Context(ctx context.Context) *ProjectsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the IAM access control policy for the specified Project. Replaces\nany existing policy.\n\nThe following constraints apply when using `setIamPolicy()`:\n\n+ Project does not support `allUsers` and `allAuthenticatedUsers` as\n`members` in a `Binding` of a `Policy`.\n\n+ The owner role can be granted only to `user` and `serviceAccount`.\n\n+ Service accounts can be made owners of a project directly\nwithout any restrictions. However, to be added as an owner, a user must be\ninvited via Cloud Platform console and must accept the invitation.\n\n+ A user cannot be granted the owner role using `setIamPolicy()`. The user\nmust be granted the owner role using the Cloud Platform Console and must\nexplicitly accept the invitation.\n\n+ Invitations to grant the owner role cannot be sent using\n`setIamPolicy()`; they must be sent only using the Cloud Platform Console.\n\n+ Membership changes that leave the project without any owners that have\naccepted the Terms of Service (ToS) will be rejected.\n\n+ There must be at least one owner who has accepted the Terms of\nService (ToS) agreement in the policy. Calling `setIamPolicy()` to\nremove the last ToS-accepted owner from the policy will fail. This\nrestriction also applies to legacy projects that no longer have owners\nwho have accepted the ToS. Edits to IAM policies will be rejected until\nthe lack of a ToS-accepting owner is rectified.\n\n+ Calling this method requires enabling the App Engine Admin API.\n\nNote: Removing service accounts from policies or changing their roles\ncan render services completely inoperable. It is important to understand\nhow the service account is being used before removing or updating its\nroles.",
	//   "flatPath": "v1beta1/projects/{resource}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.testIamPermissions":

type ProjectsTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified Project.
func (r *ProjectsService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsTestIamPermissionsCall {
	c := &ProjectsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTestIamPermissionsCall) Context(ctx context.Context) *ProjectsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified Project.",
	//   "flatPath": "v1beta1/projects/{resource}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.undelete":

type ProjectsUndeleteCall struct {
	s                      *Service
	projectId              string
	undeleteprojectrequest *UndeleteProjectRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Undelete: Restores the Project identified by the
// specified
// `project_id` (for example, `my-project-123`).
// You can only use this method for a Project that has a lifecycle state
// of
// DELETE_REQUESTED.
// After deletion starts, the Project cannot be restored.
//
// The caller must have modify permissions for this Project.
func (r *ProjectsService) Undelete(projectId string, undeleteprojectrequest *UndeleteProjectRequest) *ProjectsUndeleteCall {
	c := &ProjectsUndeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.undeleteprojectrequest = undeleteprojectrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUndeleteCall) Fields(s ...googleapi.Field) *ProjectsUndeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsUndeleteCall) Context(ctx context.Context) *ProjectsUndeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsUndeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsUndeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.undeleteprojectrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}:undelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.undelete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsUndeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Restores the Project identified by the specified\n`project_id` (for example, `my-project-123`).\nYou can only use this method for a Project that has a lifecycle state of\nDELETE_REQUESTED.\nAfter deletion starts, the Project cannot be restored.\n\nThe caller must have modify permissions for this Project.",
	//   "flatPath": "v1beta1/projects/{projectId}:undelete",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.undelete",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `foo-bar-123`).\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}:undelete",
	//   "request": {
	//     "$ref": "UndeleteProjectRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.update":

type ProjectsUpdateCall struct {
	s          *Service
	projectId  string
	project    *Project
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates the attributes of the Project identified by the
// specified
// `project_id` (for example, `my-project-123`).
//
// The caller must have modify permissions for this Project.
func (r *ProjectsService) Update(projectId string, project *Project) *ProjectsUpdateCall {
	c := &ProjectsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.project = project
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUpdateCall) Fields(s ...googleapi.Field) *ProjectsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsUpdateCall) Context(ctx context.Context) *ProjectsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.project)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudresourcemanager.projects.update" call.
// Exactly one of *Project or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Project.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsUpdateCall) Do(opts ...googleapi.CallOption) (*Project, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Project{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the attributes of the Project identified by the specified\n`project_id` (for example, `my-project-123`).\n\nThe caller must have modify permissions for this Project.",
	//   "flatPath": "v1beta1/projects/{projectId}",
	//   "httpMethod": "PUT",
	//   "id": "cloudresourcemanager.projects.update",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `my-project-123`).\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}",
	//   "request": {
	//     "$ref": "Project"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

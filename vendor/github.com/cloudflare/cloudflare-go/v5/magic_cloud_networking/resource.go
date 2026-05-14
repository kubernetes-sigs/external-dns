// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_cloud_networking

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/tidwall/gjson"
)

// ResourceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceService] method instead.
type ResourceService struct {
	Options []option.RequestOption
}

// NewResourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceService(opts ...option.RequestOption) (r *ResourceService) {
	r = &ResourceService{}
	r.Options = opts
	return
}

// List resources in the Resource Catalog (Closed Beta).
func (r *ResourceService) List(ctx context.Context, params ResourceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ResourceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/resources", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List resources in the Resource Catalog (Closed Beta).
func (r *ResourceService) ListAutoPaging(ctx context.Context, params ResourceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ResourceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Export resources in the Resource Catalog as a JSON file (Closed Beta).
func (r *ResourceService) Export(ctx context.Context, params ResourceExportParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/resources/export", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Read an resource from the Resource Catalog (Closed Beta).
func (r *ResourceService) Get(ctx context.Context, resourceID string, params ResourceGetParams, opts ...option.RequestOption) (res *ResourceGetResponse, err error) {
	var env ResourceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/resources/%s", params.AccountID, resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Preview Rego query result against the latest resource catalog (Closed Beta).
func (r *ResourceService) PolicyPreview(ctx context.Context, params ResourcePolicyPreviewParams, opts ...option.RequestOption) (res *string, err error) {
	var env ResourcePolicyPreviewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/resources/policy-preview", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ResourceListResponse struct {
	ID                  string                                     `json:"id,required" format:"uuid"`
	AccountID           string                                     `json:"account_id,required"`
	CloudType           ResourceListResponseCloudType              `json:"cloud_type,required"`
	Config              map[string]interface{}                     `json:"config,required"`
	DeploymentProvider  string                                     `json:"deployment_provider,required" format:"uuid"`
	Managed             bool                                       `json:"managed,required"`
	MonthlyCostEstimate ResourceListResponseMonthlyCostEstimate    `json:"monthly_cost_estimate,required"`
	Name                string                                     `json:"name,required"`
	NativeID            string                                     `json:"native_id,required"`
	Observations        map[string]ResourceListResponseObservation `json:"observations,required"`
	ProviderIDs         []string                                   `json:"provider_ids,required" format:"uuid"`
	ProviderNamesByID   map[string]string                          `json:"provider_names_by_id,required"`
	Region              string                                     `json:"region,required"`
	ResourceGroup       string                                     `json:"resource_group,required"`
	ResourceType        ResourceListResponseResourceType           `json:"resource_type,required"`
	Sections            []ResourceListResponseSection              `json:"sections,required"`
	State               map[string]interface{}                     `json:"state,required"`
	Tags                map[string]string                          `json:"tags,required"`
	UpdatedAt           string                                     `json:"updated_at,required"`
	URL                 string                                     `json:"url,required"`
	ManagedBy           []ResourceListResponseManagedBy            `json:"managed_by"`
	JSON                resourceListResponseJSON                   `json:"-"`
}

// resourceListResponseJSON contains the JSON metadata for the struct
// [ResourceListResponse]
type resourceListResponseJSON struct {
	ID                  apijson.Field
	AccountID           apijson.Field
	CloudType           apijson.Field
	Config              apijson.Field
	DeploymentProvider  apijson.Field
	Managed             apijson.Field
	MonthlyCostEstimate apijson.Field
	Name                apijson.Field
	NativeID            apijson.Field
	Observations        apijson.Field
	ProviderIDs         apijson.Field
	ProviderNamesByID   apijson.Field
	Region              apijson.Field
	ResourceGroup       apijson.Field
	ResourceType        apijson.Field
	Sections            apijson.Field
	State               apijson.Field
	Tags                apijson.Field
	UpdatedAt           apijson.Field
	URL                 apijson.Field
	ManagedBy           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ResourceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseCloudType string

const (
	ResourceListResponseCloudTypeAws        ResourceListResponseCloudType = "AWS"
	ResourceListResponseCloudTypeAzure      ResourceListResponseCloudType = "AZURE"
	ResourceListResponseCloudTypeGoogle     ResourceListResponseCloudType = "GOOGLE"
	ResourceListResponseCloudTypeCloudflare ResourceListResponseCloudType = "CLOUDFLARE"
)

func (r ResourceListResponseCloudType) IsKnown() bool {
	switch r {
	case ResourceListResponseCloudTypeAws, ResourceListResponseCloudTypeAzure, ResourceListResponseCloudTypeGoogle, ResourceListResponseCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceListResponseMonthlyCostEstimate struct {
	Currency    string                                      `json:"currency,required"`
	MonthlyCost float64                                     `json:"monthly_cost,required"`
	JSON        resourceListResponseMonthlyCostEstimateJSON `json:"-"`
}

// resourceListResponseMonthlyCostEstimateJSON contains the JSON metadata for the
// struct [ResourceListResponseMonthlyCostEstimate]
type resourceListResponseMonthlyCostEstimateJSON struct {
	Currency    apijson.Field
	MonthlyCost apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseMonthlyCostEstimate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseMonthlyCostEstimateJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseObservation struct {
	FirstObservedAt string                              `json:"first_observed_at,required"`
	LastObservedAt  string                              `json:"last_observed_at,required"`
	ProviderID      string                              `json:"provider_id,required" format:"uuid"`
	ResourceID      string                              `json:"resource_id,required" format:"uuid"`
	JSON            resourceListResponseObservationJSON `json:"-"`
}

// resourceListResponseObservationJSON contains the JSON metadata for the struct
// [ResourceListResponseObservation]
type resourceListResponseObservationJSON struct {
	FirstObservedAt apijson.Field
	LastObservedAt  apijson.Field
	ProviderID      apijson.Field
	ResourceID      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceListResponseObservation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseObservationJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseResourceType string

const (
	ResourceListResponseResourceTypeAwsCustomerGateway                                         ResourceListResponseResourceType = "aws_customer_gateway"
	ResourceListResponseResourceTypeAwsEgressOnlyInternetGateway                               ResourceListResponseResourceType = "aws_egress_only_internet_gateway"
	ResourceListResponseResourceTypeAwsInternetGateway                                         ResourceListResponseResourceType = "aws_internet_gateway"
	ResourceListResponseResourceTypeAwsInstance                                                ResourceListResponseResourceType = "aws_instance"
	ResourceListResponseResourceTypeAwsNetworkInterface                                        ResourceListResponseResourceType = "aws_network_interface"
	ResourceListResponseResourceTypeAwsRoute                                                   ResourceListResponseResourceType = "aws_route"
	ResourceListResponseResourceTypeAwsRouteTable                                              ResourceListResponseResourceType = "aws_route_table"
	ResourceListResponseResourceTypeAwsRouteTableAssociation                                   ResourceListResponseResourceType = "aws_route_table_association"
	ResourceListResponseResourceTypeAwsSubnet                                                  ResourceListResponseResourceType = "aws_subnet"
	ResourceListResponseResourceTypeAwsVPC                                                     ResourceListResponseResourceType = "aws_vpc"
	ResourceListResponseResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceListResponseResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceListResponseResourceTypeAwsVpnConnection                                           ResourceListResponseResourceType = "aws_vpn_connection"
	ResourceListResponseResourceTypeAwsVpnConnectionRoute                                      ResourceListResponseResourceType = "aws_vpn_connection_route"
	ResourceListResponseResourceTypeAwsVpnGateway                                              ResourceListResponseResourceType = "aws_vpn_gateway"
	ResourceListResponseResourceTypeAwsSecurityGroup                                           ResourceListResponseResourceType = "aws_security_group"
	ResourceListResponseResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceListResponseResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceListResponseResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceListResponseResourceType = "aws_vpc_security_group_egress_rule"
	ResourceListResponseResourceTypeAwsEc2ManagedPrefixList                                    ResourceListResponseResourceType = "aws_ec2_managed_prefix_list"
	ResourceListResponseResourceTypeAwsEc2TransitGateway                                       ResourceListResponseResourceType = "aws_ec2_transit_gateway"
	ResourceListResponseResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceListResponseResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceListResponseResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceListResponseResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceListResponseResourceTypeAzurermApplicationSecurityGroup                            ResourceListResponseResourceType = "azurerm_application_security_group"
	ResourceListResponseResourceTypeAzurermLB                                                  ResourceListResponseResourceType = "azurerm_lb"
	ResourceListResponseResourceTypeAzurermLBBackendAddressPool                                ResourceListResponseResourceType = "azurerm_lb_backend_address_pool"
	ResourceListResponseResourceTypeAzurermLBNatPool                                           ResourceListResponseResourceType = "azurerm_lb_nat_pool"
	ResourceListResponseResourceTypeAzurermLBNatRule                                           ResourceListResponseResourceType = "azurerm_lb_nat_rule"
	ResourceListResponseResourceTypeAzurermLBRule                                              ResourceListResponseResourceType = "azurerm_lb_rule"
	ResourceListResponseResourceTypeAzurermLocalNetworkGateway                                 ResourceListResponseResourceType = "azurerm_local_network_gateway"
	ResourceListResponseResourceTypeAzurermNetworkInterface                                    ResourceListResponseResourceType = "azurerm_network_interface"
	ResourceListResponseResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceListResponseResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceListResponseResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceListResponseResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceListResponseResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceListResponseResourceType = "azurerm_network_interface_security_group_association"
	ResourceListResponseResourceTypeAzurermNetworkSecurityGroup                                ResourceListResponseResourceType = "azurerm_network_security_group"
	ResourceListResponseResourceTypeAzurermPublicIP                                            ResourceListResponseResourceType = "azurerm_public_ip"
	ResourceListResponseResourceTypeAzurermRoute                                               ResourceListResponseResourceType = "azurerm_route"
	ResourceListResponseResourceTypeAzurermRouteTable                                          ResourceListResponseResourceType = "azurerm_route_table"
	ResourceListResponseResourceTypeAzurermSubnet                                              ResourceListResponseResourceType = "azurerm_subnet"
	ResourceListResponseResourceTypeAzurermSubnetRouteTableAssociation                         ResourceListResponseResourceType = "azurerm_subnet_route_table_association"
	ResourceListResponseResourceTypeAzurermVirtualMachine                                      ResourceListResponseResourceType = "azurerm_virtual_machine"
	ResourceListResponseResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceListResponseResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceListResponseResourceTypeAzurermVirtualNetwork                                      ResourceListResponseResourceType = "azurerm_virtual_network"
	ResourceListResponseResourceTypeAzurermVirtualNetworkGateway                               ResourceListResponseResourceType = "azurerm_virtual_network_gateway"
	ResourceListResponseResourceTypeGoogleComputeNetwork                                       ResourceListResponseResourceType = "google_compute_network"
	ResourceListResponseResourceTypeGoogleComputeSubnetwork                                    ResourceListResponseResourceType = "google_compute_subnetwork"
	ResourceListResponseResourceTypeGoogleComputeVpnGateway                                    ResourceListResponseResourceType = "google_compute_vpn_gateway"
	ResourceListResponseResourceTypeGoogleComputeVpnTunnel                                     ResourceListResponseResourceType = "google_compute_vpn_tunnel"
	ResourceListResponseResourceTypeGoogleComputeRoute                                         ResourceListResponseResourceType = "google_compute_route"
	ResourceListResponseResourceTypeGoogleComputeAddress                                       ResourceListResponseResourceType = "google_compute_address"
	ResourceListResponseResourceTypeGoogleComputeGlobalAddress                                 ResourceListResponseResourceType = "google_compute_global_address"
	ResourceListResponseResourceTypeGoogleComputeRouter                                        ResourceListResponseResourceType = "google_compute_router"
	ResourceListResponseResourceTypeGoogleComputeInterconnectAttachment                        ResourceListResponseResourceType = "google_compute_interconnect_attachment"
	ResourceListResponseResourceTypeGoogleComputeHaVpnGateway                                  ResourceListResponseResourceType = "google_compute_ha_vpn_gateway"
	ResourceListResponseResourceTypeGoogleComputeForwardingRule                                ResourceListResponseResourceType = "google_compute_forwarding_rule"
	ResourceListResponseResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceListResponseResourceType = "google_compute_network_firewall_policy"
	ResourceListResponseResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceListResponseResourceType = "google_compute_network_firewall_policy_rule"
	ResourceListResponseResourceTypeCloudflareStaticRoute                                      ResourceListResponseResourceType = "cloudflare_static_route"
	ResourceListResponseResourceTypeCloudflareIPSECTunnel                                      ResourceListResponseResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceListResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceListResponseResourceTypeAwsCustomerGateway, ResourceListResponseResourceTypeAwsEgressOnlyInternetGateway, ResourceListResponseResourceTypeAwsInternetGateway, ResourceListResponseResourceTypeAwsInstance, ResourceListResponseResourceTypeAwsNetworkInterface, ResourceListResponseResourceTypeAwsRoute, ResourceListResponseResourceTypeAwsRouteTable, ResourceListResponseResourceTypeAwsRouteTableAssociation, ResourceListResponseResourceTypeAwsSubnet, ResourceListResponseResourceTypeAwsVPC, ResourceListResponseResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceListResponseResourceTypeAwsVpnConnection, ResourceListResponseResourceTypeAwsVpnConnectionRoute, ResourceListResponseResourceTypeAwsVpnGateway, ResourceListResponseResourceTypeAwsSecurityGroup, ResourceListResponseResourceTypeAwsVPCSecurityGroupIngressRule, ResourceListResponseResourceTypeAwsVPCSecurityGroupEgressRule, ResourceListResponseResourceTypeAwsEc2ManagedPrefixList, ResourceListResponseResourceTypeAwsEc2TransitGateway, ResourceListResponseResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceListResponseResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceListResponseResourceTypeAzurermApplicationSecurityGroup, ResourceListResponseResourceTypeAzurermLB, ResourceListResponseResourceTypeAzurermLBBackendAddressPool, ResourceListResponseResourceTypeAzurermLBNatPool, ResourceListResponseResourceTypeAzurermLBNatRule, ResourceListResponseResourceTypeAzurermLBRule, ResourceListResponseResourceTypeAzurermLocalNetworkGateway, ResourceListResponseResourceTypeAzurermNetworkInterface, ResourceListResponseResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceListResponseResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceListResponseResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceListResponseResourceTypeAzurermNetworkSecurityGroup, ResourceListResponseResourceTypeAzurermPublicIP, ResourceListResponseResourceTypeAzurermRoute, ResourceListResponseResourceTypeAzurermRouteTable, ResourceListResponseResourceTypeAzurermSubnet, ResourceListResponseResourceTypeAzurermSubnetRouteTableAssociation, ResourceListResponseResourceTypeAzurermVirtualMachine, ResourceListResponseResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceListResponseResourceTypeAzurermVirtualNetwork, ResourceListResponseResourceTypeAzurermVirtualNetworkGateway, ResourceListResponseResourceTypeGoogleComputeNetwork, ResourceListResponseResourceTypeGoogleComputeSubnetwork, ResourceListResponseResourceTypeGoogleComputeVpnGateway, ResourceListResponseResourceTypeGoogleComputeVpnTunnel, ResourceListResponseResourceTypeGoogleComputeRoute, ResourceListResponseResourceTypeGoogleComputeAddress, ResourceListResponseResourceTypeGoogleComputeGlobalAddress, ResourceListResponseResourceTypeGoogleComputeRouter, ResourceListResponseResourceTypeGoogleComputeInterconnectAttachment, ResourceListResponseResourceTypeGoogleComputeHaVpnGateway, ResourceListResponseResourceTypeGoogleComputeForwardingRule, ResourceListResponseResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceListResponseResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceListResponseResourceTypeCloudflareStaticRoute, ResourceListResponseResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceListResponseSection struct {
	HiddenItems  []ResourceListResponseSectionsHiddenItem  `json:"hidden_items,required"`
	Name         string                                    `json:"name,required"`
	VisibleItems []ResourceListResponseSectionsVisibleItem `json:"visible_items,required"`
	HelpText     string                                    `json:"help_text"`
	JSON         resourceListResponseSectionJSON           `json:"-"`
}

// resourceListResponseSectionJSON contains the JSON metadata for the struct
// [ResourceListResponseSection]
type resourceListResponseSectionJSON struct {
	HiddenItems  apijson.Field
	Name         apijson.Field
	VisibleItems apijson.Field
	HelpText     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceListResponseSection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseSectionsHiddenItem struct {
	HelpText string                                       `json:"helpText"`
	Name     string                                       `json:"name"`
	Value    ResourceListResponseSectionsHiddenItemsValue `json:"value"`
	JSON     resourceListResponseSectionsHiddenItemJSON   `json:"-"`
}

// resourceListResponseSectionsHiddenItemJSON contains the JSON metadata for the
// struct [ResourceListResponseSectionsHiddenItem]
type resourceListResponseSectionsHiddenItemJSON struct {
	HelpText    apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseSectionsHiddenItemsValue struct {
	ItemType string `json:"item_type,required"`
	// This field can have the runtime type of
	// [[]ResourceListResponseSectionsHiddenItemsValueMcnListItemList].
	List interface{} `json:"list"`
	// This field can have the runtime type of
	// [ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview].
	ResourcePreview interface{} `json:"resource_preview"`
	String          string      `json:"string"`
	Yaml            string      `json:"yaml"`
	// This field can have the runtime type of
	// [ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff].
	YamlDiff interface{}                                      `json:"yaml_diff"`
	JSON     resourceListResponseSectionsHiddenItemsValueJSON `json:"-"`
	union    ResourceListResponseSectionsHiddenItemsValueUnion
}

// resourceListResponseSectionsHiddenItemsValueJSON contains the JSON metadata for
// the struct [ResourceListResponseSectionsHiddenItemsValue]
type resourceListResponseSectionsHiddenItemsValueJSON struct {
	ItemType        apijson.Field
	List            apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	Yaml            apijson.Field
	YamlDiff        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r resourceListResponseSectionsHiddenItemsValueJSON) RawJSON() string {
	return r.raw
}

func (r *ResourceListResponseSectionsHiddenItemsValue) UnmarshalJSON(data []byte) (err error) {
	*r = ResourceListResponseSectionsHiddenItemsValue{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ResourceListResponseSectionsHiddenItemsValueUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ResourceListResponseSectionsHiddenItemsValueMcnStringItem],
// [ResourceListResponseSectionsHiddenItemsValueMcnYamlItem],
// [ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItem],
// [ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItem],
// [ResourceListResponseSectionsHiddenItemsValueMcnListItem].
func (r ResourceListResponseSectionsHiddenItemsValue) AsUnion() ResourceListResponseSectionsHiddenItemsValueUnion {
	return r.union
}

// Union satisfied by [ResourceListResponseSectionsHiddenItemsValueMcnStringItem],
// [ResourceListResponseSectionsHiddenItemsValueMcnYamlItem],
// [ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItem],
// [ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItem] or
// [ResourceListResponseSectionsHiddenItemsValueMcnListItem].
type ResourceListResponseSectionsHiddenItemsValueUnion interface {
	implementsResourceListResponseSectionsHiddenItemsValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResourceListResponseSectionsHiddenItemsValueUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsHiddenItemsValueMcnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsHiddenItemsValueMcnYamlItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsHiddenItemsValueMcnListItem{}),
		},
	)
}

type ResourceListResponseSectionsHiddenItemsValueMcnStringItem struct {
	ItemType string                                                        `json:"item_type,required"`
	String   string                                                        `json:"string,required"`
	JSON     resourceListResponseSectionsHiddenItemsValueMcnStringItemJSON `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnStringItemJSON contains the JSON
// metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnStringItem]
type resourceListResponseSectionsHiddenItemsValueMcnStringItemJSON struct {
	ItemType    apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnStringItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnStringItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsHiddenItemsValueMcnStringItem) implementsResourceListResponseSectionsHiddenItemsValue() {
}

type ResourceListResponseSectionsHiddenItemsValueMcnYamlItem struct {
	ItemType string                                                      `json:"item_type,required"`
	Yaml     string                                                      `json:"yaml,required"`
	JSON     resourceListResponseSectionsHiddenItemsValueMcnYamlItemJSON `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnYamlItemJSON contains the JSON
// metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnYamlItem]
type resourceListResponseSectionsHiddenItemsValueMcnYamlItemJSON struct {
	ItemType    apijson.Field
	Yaml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnYamlItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnYamlItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsHiddenItemsValueMcnYamlItem) implementsResourceListResponseSectionsHiddenItemsValue() {
}

type ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItem struct {
	ItemType string                                                              `json:"item_type,required"`
	YamlDiff ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff `json:"yaml_diff,required"`
	JSON     resourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemJSON     `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemJSON contains the
// JSON metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItem]
type resourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemJSON struct {
	ItemType    apijson.Field
	YamlDiff    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItem) implementsResourceListResponseSectionsHiddenItemsValue() {
}

type ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff struct {
	Diff             string                                                                  `json:"diff,required"`
	LeftDescription  string                                                                  `json:"left_description,required"`
	LeftYaml         string                                                                  `json:"left_yaml,required"`
	RightDescription string                                                                  `json:"right_description,required"`
	RightYaml        string                                                                  `json:"right_yaml,required"`
	JSON             resourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiffJSON `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiffJSON contains
// the JSON metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff]
type resourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiffJSON struct {
	Diff             apijson.Field
	LeftDescription  apijson.Field
	LeftYaml         apijson.Field
	RightDescription apijson.Field
	RightYaml        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiffJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItem struct {
	ItemType        string                                                                            `json:"item_type,required"`
	ResourcePreview ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview `json:"resource_preview,required"`
	JSON            resourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemJSON            `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemJSON contains
// the JSON metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItem]
type resourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItem) implementsResourceListResponseSectionsHiddenItemsValue() {
}

type ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview struct {
	ID           string                                                                                        `json:"id,required" format:"uuid"`
	CloudType    ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType    `json:"cloud_type,required"`
	Detail       string                                                                                        `json:"detail,required"`
	Name         string                                                                                        `json:"name,required"`
	ResourceType ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType `json:"resource_type,required"`
	Title        string                                                                                        `json:"title,required"`
	JSON         resourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewJSON         `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewJSON
// contains the JSON metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview]
type resourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewJSON struct {
	ID           apijson.Field
	CloudType    apijson.Field
	Detail       apijson.Field
	Name         apijson.Field
	ResourceType apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType string

const (
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAws        ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "AWS"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAzure      ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "AZURE"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeGoogle     ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "GOOGLE"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "CLOUDFLARE"
)

func (r ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType) IsKnown() bool {
	switch r {
	case ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAws, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAzure, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeGoogle, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType string

const (
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway                                         ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_customer_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway                               ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_egress_only_internet_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway                                         ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_internet_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance                                                ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_instance"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface                                        ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_network_interface"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute                                                   ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable                                              ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation                                   ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table_association"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet                                                  ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_subnet"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC                                                     ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection                                           ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute                                      ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection_route"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway                                              ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup                                           ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_security_group"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_egress_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList                                    ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_managed_prefix_list"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway                                       ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup                            ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_application_security_group"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB                                                  ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool                                ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_backend_address_pool"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool                                           ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_pool"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule                                           ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule                                              ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway                                 ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_local_network_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface                                    ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_security_group_association"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup                                ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_security_group"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP                                            ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_public_ip"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute                                               ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable                                          ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route_table"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet                                              ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation                         ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet_route_table_association"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine                                      ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_machine"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork                                      ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway                               ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork                                       ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork                                    ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_subnetwork"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway                                    ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel                                     ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_tunnel"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute                                         ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_route"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress                                       ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_address"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress                                 ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_global_address"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter                                        ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_router"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment                        ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_interconnect_attachment"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway                                  ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_ha_vpn_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule                                ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_forwarding_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute                                      ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_static_route"
	ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel                                      ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType) IsKnown() bool {
	switch r {
	case ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute, ResourceListResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceListResponseSectionsHiddenItemsValueMcnListItem struct {
	ItemType string                                                        `json:"item_type,required"`
	List     []ResourceListResponseSectionsHiddenItemsValueMcnListItemList `json:"list,required"`
	JSON     resourceListResponseSectionsHiddenItemsValueMcnListItemJSON   `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnListItemJSON contains the JSON
// metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnListItem]
type resourceListResponseSectionsHiddenItemsValueMcnListItemJSON struct {
	ItemType    apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnListItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsHiddenItemsValueMcnListItem) implementsResourceListResponseSectionsHiddenItemsValue() {
}

type ResourceListResponseSectionsHiddenItemsValueMcnListItemList struct {
	ItemType string `json:"item_type,required"`
	// This field can have the runtime type of
	// [ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview].
	ResourcePreview interface{}                                                     `json:"resource_preview"`
	String          string                                                          `json:"string"`
	JSON            resourceListResponseSectionsHiddenItemsValueMcnListItemListJSON `json:"-"`
	union           ResourceListResponseSectionsHiddenItemsValueMcnListItemListUnion
}

// resourceListResponseSectionsHiddenItemsValueMcnListItemListJSON contains the
// JSON metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnListItemList]
type resourceListResponseSectionsHiddenItemsValueMcnListItemListJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r resourceListResponseSectionsHiddenItemsValueMcnListItemListJSON) RawJSON() string {
	return r.raw
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnListItemList) UnmarshalJSON(data []byte) (err error) {
	*r = ResourceListResponseSectionsHiddenItemsValueMcnListItemList{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ResourceListResponseSectionsHiddenItemsValueMcnListItemListUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem],
// [ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem].
func (r ResourceListResponseSectionsHiddenItemsValueMcnListItemList) AsUnion() ResourceListResponseSectionsHiddenItemsValueMcnListItemListUnion {
	return r.union
}

// Union satisfied by
// [ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem] or
// [ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem].
type ResourceListResponseSectionsHiddenItemsValueMcnListItemListUnion interface {
	implementsResourceListResponseSectionsHiddenItemsValueMcnListItemList()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResourceListResponseSectionsHiddenItemsValueMcnListItemListUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem{}),
		},
	)
}

type ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem struct {
	ItemType string                                                                       `json:"item_type,required"`
	String   string                                                                       `json:"string,required"`
	JSON     resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItemJSON `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItemJSON
// contains the JSON metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem]
type resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItemJSON struct {
	ItemType    apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem) implementsResourceListResponseSectionsHiddenItemsValueMcnListItemList() {
}

type ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem struct {
	ItemType        string                                                                                           `json:"item_type,required"`
	ResourcePreview ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview `json:"resource_preview,required"`
	JSON            resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemJSON            `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemJSON
// contains the JSON metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem]
type resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem) implementsResourceListResponseSectionsHiddenItemsValueMcnListItemList() {
}

type ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview struct {
	ID           string                                                                                                       `json:"id,required" format:"uuid"`
	CloudType    ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType    `json:"cloud_type,required"`
	Detail       string                                                                                                       `json:"detail,required"`
	Name         string                                                                                                       `json:"name,required"`
	ResourceType ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType `json:"resource_type,required"`
	Title        string                                                                                                       `json:"title,required"`
	JSON         resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON         `json:"-"`
}

// resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON
// contains the JSON metadata for the struct
// [ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview]
type resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON struct {
	ID           apijson.Field
	CloudType    apijson.Field
	Detail       apijson.Field
	Name         apijson.Field
	ResourceType apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType string

const (
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAws        ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "AWS"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAzure      ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "AZURE"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeGoogle     ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "GOOGLE"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "CLOUDFLARE"
)

func (r ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType) IsKnown() bool {
	switch r {
	case ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAws, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAzure, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeGoogle, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType string

const (
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway                                         ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_customer_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway                               ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_egress_only_internet_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway                                         ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_internet_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance                                                ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_instance"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface                                        ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_network_interface"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute                                                   ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable                                              ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation                                   ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table_association"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet                                                  ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_subnet"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC                                                     ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection                                           ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute                                      ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection_route"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway                                              ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup                                           ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_security_group"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_egress_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList                                    ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_managed_prefix_list"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway                                       ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup                            ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_application_security_group"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB                                                  ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool                                ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_backend_address_pool"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool                                           ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_pool"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule                                           ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule                                              ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway                                 ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_local_network_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface                                    ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_security_group_association"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup                                ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_security_group"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP                                            ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_public_ip"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute                                               ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable                                          ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route_table"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet                                              ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation                         ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet_route_table_association"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine                                      ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_machine"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork                                      ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway                               ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork                                       ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork                                    ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_subnetwork"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway                                    ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel                                     ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_tunnel"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute                                         ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_route"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress                                       ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_address"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress                                 ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_global_address"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter                                        ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_router"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment                        ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_interconnect_attachment"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway                                  ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_ha_vpn_gateway"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule                                ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_forwarding_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy_rule"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute                                      ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_static_route"
	ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel                                      ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType) IsKnown() bool {
	switch r {
	case ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute, ResourceListResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceListResponseSectionsVisibleItem struct {
	HelpText string                                        `json:"helpText"`
	Name     string                                        `json:"name"`
	Value    ResourceListResponseSectionsVisibleItemsValue `json:"value"`
	JSON     resourceListResponseSectionsVisibleItemJSON   `json:"-"`
}

// resourceListResponseSectionsVisibleItemJSON contains the JSON metadata for the
// struct [ResourceListResponseSectionsVisibleItem]
type resourceListResponseSectionsVisibleItemJSON struct {
	HelpText    apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseSectionsVisibleItemsValue struct {
	ItemType string `json:"item_type,required"`
	// This field can have the runtime type of
	// [[]ResourceListResponseSectionsVisibleItemsValueMcnListItemList].
	List interface{} `json:"list"`
	// This field can have the runtime type of
	// [ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview].
	ResourcePreview interface{} `json:"resource_preview"`
	String          string      `json:"string"`
	Yaml            string      `json:"yaml"`
	// This field can have the runtime type of
	// [ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff].
	YamlDiff interface{}                                       `json:"yaml_diff"`
	JSON     resourceListResponseSectionsVisibleItemsValueJSON `json:"-"`
	union    ResourceListResponseSectionsVisibleItemsValueUnion
}

// resourceListResponseSectionsVisibleItemsValueJSON contains the JSON metadata for
// the struct [ResourceListResponseSectionsVisibleItemsValue]
type resourceListResponseSectionsVisibleItemsValueJSON struct {
	ItemType        apijson.Field
	List            apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	Yaml            apijson.Field
	YamlDiff        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r resourceListResponseSectionsVisibleItemsValueJSON) RawJSON() string {
	return r.raw
}

func (r *ResourceListResponseSectionsVisibleItemsValue) UnmarshalJSON(data []byte) (err error) {
	*r = ResourceListResponseSectionsVisibleItemsValue{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ResourceListResponseSectionsVisibleItemsValueUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ResourceListResponseSectionsVisibleItemsValueMcnStringItem],
// [ResourceListResponseSectionsVisibleItemsValueMcnYamlItem],
// [ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItem],
// [ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItem],
// [ResourceListResponseSectionsVisibleItemsValueMcnListItem].
func (r ResourceListResponseSectionsVisibleItemsValue) AsUnion() ResourceListResponseSectionsVisibleItemsValueUnion {
	return r.union
}

// Union satisfied by [ResourceListResponseSectionsVisibleItemsValueMcnStringItem],
// [ResourceListResponseSectionsVisibleItemsValueMcnYamlItem],
// [ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItem],
// [ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItem] or
// [ResourceListResponseSectionsVisibleItemsValueMcnListItem].
type ResourceListResponseSectionsVisibleItemsValueUnion interface {
	implementsResourceListResponseSectionsVisibleItemsValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResourceListResponseSectionsVisibleItemsValueUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsVisibleItemsValueMcnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsVisibleItemsValueMcnYamlItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsVisibleItemsValueMcnListItem{}),
		},
	)
}

type ResourceListResponseSectionsVisibleItemsValueMcnStringItem struct {
	ItemType string                                                         `json:"item_type,required"`
	String   string                                                         `json:"string,required"`
	JSON     resourceListResponseSectionsVisibleItemsValueMcnStringItemJSON `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnStringItemJSON contains the JSON
// metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnStringItem]
type resourceListResponseSectionsVisibleItemsValueMcnStringItemJSON struct {
	ItemType    apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnStringItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnStringItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsVisibleItemsValueMcnStringItem) implementsResourceListResponseSectionsVisibleItemsValue() {
}

type ResourceListResponseSectionsVisibleItemsValueMcnYamlItem struct {
	ItemType string                                                       `json:"item_type,required"`
	Yaml     string                                                       `json:"yaml,required"`
	JSON     resourceListResponseSectionsVisibleItemsValueMcnYamlItemJSON `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnYamlItemJSON contains the JSON
// metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnYamlItem]
type resourceListResponseSectionsVisibleItemsValueMcnYamlItemJSON struct {
	ItemType    apijson.Field
	Yaml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnYamlItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnYamlItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsVisibleItemsValueMcnYamlItem) implementsResourceListResponseSectionsVisibleItemsValue() {
}

type ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItem struct {
	ItemType string                                                               `json:"item_type,required"`
	YamlDiff ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff `json:"yaml_diff,required"`
	JSON     resourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemJSON     `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemJSON contains the
// JSON metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItem]
type resourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemJSON struct {
	ItemType    apijson.Field
	YamlDiff    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItem) implementsResourceListResponseSectionsVisibleItemsValue() {
}

type ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff struct {
	Diff             string                                                                   `json:"diff,required"`
	LeftDescription  string                                                                   `json:"left_description,required"`
	LeftYaml         string                                                                   `json:"left_yaml,required"`
	RightDescription string                                                                   `json:"right_description,required"`
	RightYaml        string                                                                   `json:"right_yaml,required"`
	JSON             resourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiffJSON `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiffJSON
// contains the JSON metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff]
type resourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiffJSON struct {
	Diff             apijson.Field
	LeftDescription  apijson.Field
	LeftYaml         apijson.Field
	RightDescription apijson.Field
	RightYaml        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiffJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItem struct {
	ItemType        string                                                                             `json:"item_type,required"`
	ResourcePreview ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview `json:"resource_preview,required"`
	JSON            resourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemJSON            `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemJSON contains
// the JSON metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItem]
type resourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItem) implementsResourceListResponseSectionsVisibleItemsValue() {
}

type ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview struct {
	ID           string                                                                                         `json:"id,required" format:"uuid"`
	CloudType    ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType    `json:"cloud_type,required"`
	Detail       string                                                                                         `json:"detail,required"`
	Name         string                                                                                         `json:"name,required"`
	ResourceType ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType `json:"resource_type,required"`
	Title        string                                                                                         `json:"title,required"`
	JSON         resourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewJSON         `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewJSON
// contains the JSON metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview]
type resourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewJSON struct {
	ID           apijson.Field
	CloudType    apijson.Field
	Detail       apijson.Field
	Name         apijson.Field
	ResourceType apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType string

const (
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAws        ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "AWS"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAzure      ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "AZURE"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeGoogle     ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "GOOGLE"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "CLOUDFLARE"
)

func (r ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType) IsKnown() bool {
	switch r {
	case ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAws, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAzure, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeGoogle, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType string

const (
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway                                         ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_customer_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway                               ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_egress_only_internet_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway                                         ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_internet_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance                                                ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_instance"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface                                        ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_network_interface"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute                                                   ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable                                              ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation                                   ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table_association"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet                                                  ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_subnet"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC                                                     ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection                                           ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute                                      ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection_route"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway                                              ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup                                           ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_security_group"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_egress_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList                                    ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_managed_prefix_list"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway                                       ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup                            ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_application_security_group"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB                                                  ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool                                ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_backend_address_pool"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool                                           ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_pool"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule                                           ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule                                              ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway                                 ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_local_network_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface                                    ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_security_group_association"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup                                ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_security_group"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP                                            ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_public_ip"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute                                               ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable                                          ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route_table"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet                                              ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation                         ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet_route_table_association"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine                                      ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_machine"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork                                      ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway                               ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork                                       ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork                                    ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_subnetwork"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway                                    ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel                                     ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_tunnel"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute                                         ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_route"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress                                       ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_address"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress                                 ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_global_address"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter                                        ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_router"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment                        ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_interconnect_attachment"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway                                  ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_ha_vpn_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule                                ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_forwarding_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute                                      ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_static_route"
	ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel                                      ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType) IsKnown() bool {
	switch r {
	case ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute, ResourceListResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceListResponseSectionsVisibleItemsValueMcnListItem struct {
	ItemType string                                                         `json:"item_type,required"`
	List     []ResourceListResponseSectionsVisibleItemsValueMcnListItemList `json:"list,required"`
	JSON     resourceListResponseSectionsVisibleItemsValueMcnListItemJSON   `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnListItemJSON contains the JSON
// metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnListItem]
type resourceListResponseSectionsVisibleItemsValueMcnListItemJSON struct {
	ItemType    apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnListItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsVisibleItemsValueMcnListItem) implementsResourceListResponseSectionsVisibleItemsValue() {
}

type ResourceListResponseSectionsVisibleItemsValueMcnListItemList struct {
	ItemType string `json:"item_type,required"`
	// This field can have the runtime type of
	// [ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview].
	ResourcePreview interface{}                                                      `json:"resource_preview"`
	String          string                                                           `json:"string"`
	JSON            resourceListResponseSectionsVisibleItemsValueMcnListItemListJSON `json:"-"`
	union           ResourceListResponseSectionsVisibleItemsValueMcnListItemListUnion
}

// resourceListResponseSectionsVisibleItemsValueMcnListItemListJSON contains the
// JSON metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnListItemList]
type resourceListResponseSectionsVisibleItemsValueMcnListItemListJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r resourceListResponseSectionsVisibleItemsValueMcnListItemListJSON) RawJSON() string {
	return r.raw
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnListItemList) UnmarshalJSON(data []byte) (err error) {
	*r = ResourceListResponseSectionsVisibleItemsValueMcnListItemList{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ResourceListResponseSectionsVisibleItemsValueMcnListItemListUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem],
// [ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem].
func (r ResourceListResponseSectionsVisibleItemsValueMcnListItemList) AsUnion() ResourceListResponseSectionsVisibleItemsValueMcnListItemListUnion {
	return r.union
}

// Union satisfied by
// [ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem] or
// [ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem].
type ResourceListResponseSectionsVisibleItemsValueMcnListItemListUnion interface {
	implementsResourceListResponseSectionsVisibleItemsValueMcnListItemList()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResourceListResponseSectionsVisibleItemsValueMcnListItemListUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem{}),
		},
	)
}

type ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem struct {
	ItemType string                                                                        `json:"item_type,required"`
	String   string                                                                        `json:"string,required"`
	JSON     resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItemJSON `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItemJSON
// contains the JSON metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem]
type resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItemJSON struct {
	ItemType    apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem) implementsResourceListResponseSectionsVisibleItemsValueMcnListItemList() {
}

type ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem struct {
	ItemType        string                                                                                            `json:"item_type,required"`
	ResourcePreview ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview `json:"resource_preview,required"`
	JSON            resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemJSON            `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemJSON
// contains the JSON metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem]
type resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem) implementsResourceListResponseSectionsVisibleItemsValueMcnListItemList() {
}

type ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview struct {
	ID           string                                                                                                        `json:"id,required" format:"uuid"`
	CloudType    ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType    `json:"cloud_type,required"`
	Detail       string                                                                                                        `json:"detail,required"`
	Name         string                                                                                                        `json:"name,required"`
	ResourceType ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType `json:"resource_type,required"`
	Title        string                                                                                                        `json:"title,required"`
	JSON         resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON         `json:"-"`
}

// resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON
// contains the JSON metadata for the struct
// [ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview]
type resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON struct {
	ID           apijson.Field
	CloudType    apijson.Field
	Detail       apijson.Field
	Name         apijson.Field
	ResourceType apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType string

const (
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAws        ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "AWS"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAzure      ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "AZURE"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeGoogle     ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "GOOGLE"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "CLOUDFLARE"
)

func (r ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType) IsKnown() bool {
	switch r {
	case ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAws, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAzure, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeGoogle, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType string

const (
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway                                         ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_customer_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway                               ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_egress_only_internet_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway                                         ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_internet_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance                                                ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_instance"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface                                        ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_network_interface"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute                                                   ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable                                              ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation                                   ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table_association"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet                                                  ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_subnet"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC                                                     ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection                                           ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute                                      ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection_route"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway                                              ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup                                           ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_security_group"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_egress_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList                                    ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_managed_prefix_list"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway                                       ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup                            ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_application_security_group"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB                                                  ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool                                ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_backend_address_pool"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool                                           ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_pool"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule                                           ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule                                              ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway                                 ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_local_network_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface                                    ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_security_group_association"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup                                ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_security_group"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP                                            ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_public_ip"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute                                               ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable                                          ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route_table"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet                                              ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation                         ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet_route_table_association"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine                                      ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_machine"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork                                      ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway                               ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork                                       ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork                                    ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_subnetwork"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway                                    ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel                                     ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_tunnel"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute                                         ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_route"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress                                       ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_address"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress                                 ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_global_address"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter                                        ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_router"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment                        ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_interconnect_attachment"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway                                  ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_ha_vpn_gateway"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule                                ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_forwarding_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy_rule"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute                                      ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_static_route"
	ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel                                      ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType) IsKnown() bool {
	switch r {
	case ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute, ResourceListResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceListResponseManagedBy struct {
	ID         string                                  `json:"id,required" format:"uuid"`
	ClientType ResourceListResponseManagedByClientType `json:"client_type,required"`
	Name       string                                  `json:"name,required"`
	JSON       resourceListResponseManagedByJSON       `json:"-"`
}

// resourceListResponseManagedByJSON contains the JSON metadata for the struct
// [ResourceListResponseManagedBy]
type resourceListResponseManagedByJSON struct {
	ID          apijson.Field
	ClientType  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceListResponseManagedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceListResponseManagedByJSON) RawJSON() string {
	return r.raw
}

type ResourceListResponseManagedByClientType string

const (
	ResourceListResponseManagedByClientTypeMagicWANCloudOnramp ResourceListResponseManagedByClientType = "MAGIC_WAN_CLOUD_ONRAMP"
)

func (r ResourceListResponseManagedByClientType) IsKnown() bool {
	switch r {
	case ResourceListResponseManagedByClientTypeMagicWANCloudOnramp:
		return true
	}
	return false
}

type ResourceGetResponse struct {
	ID                  string                                    `json:"id,required" format:"uuid"`
	AccountID           string                                    `json:"account_id,required"`
	CloudType           ResourceGetResponseCloudType              `json:"cloud_type,required"`
	Config              map[string]interface{}                    `json:"config,required"`
	DeploymentProvider  string                                    `json:"deployment_provider,required" format:"uuid"`
	Managed             bool                                      `json:"managed,required"`
	MonthlyCostEstimate ResourceGetResponseMonthlyCostEstimate    `json:"monthly_cost_estimate,required"`
	Name                string                                    `json:"name,required"`
	NativeID            string                                    `json:"native_id,required"`
	Observations        map[string]ResourceGetResponseObservation `json:"observations,required"`
	ProviderIDs         []string                                  `json:"provider_ids,required" format:"uuid"`
	ProviderNamesByID   map[string]string                         `json:"provider_names_by_id,required"`
	Region              string                                    `json:"region,required"`
	ResourceGroup       string                                    `json:"resource_group,required"`
	ResourceType        ResourceGetResponseResourceType           `json:"resource_type,required"`
	Sections            []ResourceGetResponseSection              `json:"sections,required"`
	State               map[string]interface{}                    `json:"state,required"`
	Tags                map[string]string                         `json:"tags,required"`
	UpdatedAt           string                                    `json:"updated_at,required"`
	URL                 string                                    `json:"url,required"`
	ManagedBy           []ResourceGetResponseManagedBy            `json:"managed_by"`
	JSON                resourceGetResponseJSON                   `json:"-"`
}

// resourceGetResponseJSON contains the JSON metadata for the struct
// [ResourceGetResponse]
type resourceGetResponseJSON struct {
	ID                  apijson.Field
	AccountID           apijson.Field
	CloudType           apijson.Field
	Config              apijson.Field
	DeploymentProvider  apijson.Field
	Managed             apijson.Field
	MonthlyCostEstimate apijson.Field
	Name                apijson.Field
	NativeID            apijson.Field
	Observations        apijson.Field
	ProviderIDs         apijson.Field
	ProviderNamesByID   apijson.Field
	Region              apijson.Field
	ResourceGroup       apijson.Field
	ResourceType        apijson.Field
	Sections            apijson.Field
	State               apijson.Field
	Tags                apijson.Field
	UpdatedAt           apijson.Field
	URL                 apijson.Field
	ManagedBy           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ResourceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseCloudType string

const (
	ResourceGetResponseCloudTypeAws        ResourceGetResponseCloudType = "AWS"
	ResourceGetResponseCloudTypeAzure      ResourceGetResponseCloudType = "AZURE"
	ResourceGetResponseCloudTypeGoogle     ResourceGetResponseCloudType = "GOOGLE"
	ResourceGetResponseCloudTypeCloudflare ResourceGetResponseCloudType = "CLOUDFLARE"
)

func (r ResourceGetResponseCloudType) IsKnown() bool {
	switch r {
	case ResourceGetResponseCloudTypeAws, ResourceGetResponseCloudTypeAzure, ResourceGetResponseCloudTypeGoogle, ResourceGetResponseCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceGetResponseMonthlyCostEstimate struct {
	Currency    string                                     `json:"currency,required"`
	MonthlyCost float64                                    `json:"monthly_cost,required"`
	JSON        resourceGetResponseMonthlyCostEstimateJSON `json:"-"`
}

// resourceGetResponseMonthlyCostEstimateJSON contains the JSON metadata for the
// struct [ResourceGetResponseMonthlyCostEstimate]
type resourceGetResponseMonthlyCostEstimateJSON struct {
	Currency    apijson.Field
	MonthlyCost apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseMonthlyCostEstimate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseMonthlyCostEstimateJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseObservation struct {
	FirstObservedAt string                             `json:"first_observed_at,required"`
	LastObservedAt  string                             `json:"last_observed_at,required"`
	ProviderID      string                             `json:"provider_id,required" format:"uuid"`
	ResourceID      string                             `json:"resource_id,required" format:"uuid"`
	JSON            resourceGetResponseObservationJSON `json:"-"`
}

// resourceGetResponseObservationJSON contains the JSON metadata for the struct
// [ResourceGetResponseObservation]
type resourceGetResponseObservationJSON struct {
	FirstObservedAt apijson.Field
	LastObservedAt  apijson.Field
	ProviderID      apijson.Field
	ResourceID      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceGetResponseObservation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseObservationJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseResourceType string

const (
	ResourceGetResponseResourceTypeAwsCustomerGateway                                         ResourceGetResponseResourceType = "aws_customer_gateway"
	ResourceGetResponseResourceTypeAwsEgressOnlyInternetGateway                               ResourceGetResponseResourceType = "aws_egress_only_internet_gateway"
	ResourceGetResponseResourceTypeAwsInternetGateway                                         ResourceGetResponseResourceType = "aws_internet_gateway"
	ResourceGetResponseResourceTypeAwsInstance                                                ResourceGetResponseResourceType = "aws_instance"
	ResourceGetResponseResourceTypeAwsNetworkInterface                                        ResourceGetResponseResourceType = "aws_network_interface"
	ResourceGetResponseResourceTypeAwsRoute                                                   ResourceGetResponseResourceType = "aws_route"
	ResourceGetResponseResourceTypeAwsRouteTable                                              ResourceGetResponseResourceType = "aws_route_table"
	ResourceGetResponseResourceTypeAwsRouteTableAssociation                                   ResourceGetResponseResourceType = "aws_route_table_association"
	ResourceGetResponseResourceTypeAwsSubnet                                                  ResourceGetResponseResourceType = "aws_subnet"
	ResourceGetResponseResourceTypeAwsVPC                                                     ResourceGetResponseResourceType = "aws_vpc"
	ResourceGetResponseResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceGetResponseResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceGetResponseResourceTypeAwsVpnConnection                                           ResourceGetResponseResourceType = "aws_vpn_connection"
	ResourceGetResponseResourceTypeAwsVpnConnectionRoute                                      ResourceGetResponseResourceType = "aws_vpn_connection_route"
	ResourceGetResponseResourceTypeAwsVpnGateway                                              ResourceGetResponseResourceType = "aws_vpn_gateway"
	ResourceGetResponseResourceTypeAwsSecurityGroup                                           ResourceGetResponseResourceType = "aws_security_group"
	ResourceGetResponseResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceGetResponseResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceGetResponseResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceGetResponseResourceType = "aws_vpc_security_group_egress_rule"
	ResourceGetResponseResourceTypeAwsEc2ManagedPrefixList                                    ResourceGetResponseResourceType = "aws_ec2_managed_prefix_list"
	ResourceGetResponseResourceTypeAwsEc2TransitGateway                                       ResourceGetResponseResourceType = "aws_ec2_transit_gateway"
	ResourceGetResponseResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceGetResponseResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceGetResponseResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceGetResponseResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceGetResponseResourceTypeAzurermApplicationSecurityGroup                            ResourceGetResponseResourceType = "azurerm_application_security_group"
	ResourceGetResponseResourceTypeAzurermLB                                                  ResourceGetResponseResourceType = "azurerm_lb"
	ResourceGetResponseResourceTypeAzurermLBBackendAddressPool                                ResourceGetResponseResourceType = "azurerm_lb_backend_address_pool"
	ResourceGetResponseResourceTypeAzurermLBNatPool                                           ResourceGetResponseResourceType = "azurerm_lb_nat_pool"
	ResourceGetResponseResourceTypeAzurermLBNatRule                                           ResourceGetResponseResourceType = "azurerm_lb_nat_rule"
	ResourceGetResponseResourceTypeAzurermLBRule                                              ResourceGetResponseResourceType = "azurerm_lb_rule"
	ResourceGetResponseResourceTypeAzurermLocalNetworkGateway                                 ResourceGetResponseResourceType = "azurerm_local_network_gateway"
	ResourceGetResponseResourceTypeAzurermNetworkInterface                                    ResourceGetResponseResourceType = "azurerm_network_interface"
	ResourceGetResponseResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceGetResponseResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceGetResponseResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceGetResponseResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceGetResponseResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceGetResponseResourceType = "azurerm_network_interface_security_group_association"
	ResourceGetResponseResourceTypeAzurermNetworkSecurityGroup                                ResourceGetResponseResourceType = "azurerm_network_security_group"
	ResourceGetResponseResourceTypeAzurermPublicIP                                            ResourceGetResponseResourceType = "azurerm_public_ip"
	ResourceGetResponseResourceTypeAzurermRoute                                               ResourceGetResponseResourceType = "azurerm_route"
	ResourceGetResponseResourceTypeAzurermRouteTable                                          ResourceGetResponseResourceType = "azurerm_route_table"
	ResourceGetResponseResourceTypeAzurermSubnet                                              ResourceGetResponseResourceType = "azurerm_subnet"
	ResourceGetResponseResourceTypeAzurermSubnetRouteTableAssociation                         ResourceGetResponseResourceType = "azurerm_subnet_route_table_association"
	ResourceGetResponseResourceTypeAzurermVirtualMachine                                      ResourceGetResponseResourceType = "azurerm_virtual_machine"
	ResourceGetResponseResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceGetResponseResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceGetResponseResourceTypeAzurermVirtualNetwork                                      ResourceGetResponseResourceType = "azurerm_virtual_network"
	ResourceGetResponseResourceTypeAzurermVirtualNetworkGateway                               ResourceGetResponseResourceType = "azurerm_virtual_network_gateway"
	ResourceGetResponseResourceTypeGoogleComputeNetwork                                       ResourceGetResponseResourceType = "google_compute_network"
	ResourceGetResponseResourceTypeGoogleComputeSubnetwork                                    ResourceGetResponseResourceType = "google_compute_subnetwork"
	ResourceGetResponseResourceTypeGoogleComputeVpnGateway                                    ResourceGetResponseResourceType = "google_compute_vpn_gateway"
	ResourceGetResponseResourceTypeGoogleComputeVpnTunnel                                     ResourceGetResponseResourceType = "google_compute_vpn_tunnel"
	ResourceGetResponseResourceTypeGoogleComputeRoute                                         ResourceGetResponseResourceType = "google_compute_route"
	ResourceGetResponseResourceTypeGoogleComputeAddress                                       ResourceGetResponseResourceType = "google_compute_address"
	ResourceGetResponseResourceTypeGoogleComputeGlobalAddress                                 ResourceGetResponseResourceType = "google_compute_global_address"
	ResourceGetResponseResourceTypeGoogleComputeRouter                                        ResourceGetResponseResourceType = "google_compute_router"
	ResourceGetResponseResourceTypeGoogleComputeInterconnectAttachment                        ResourceGetResponseResourceType = "google_compute_interconnect_attachment"
	ResourceGetResponseResourceTypeGoogleComputeHaVpnGateway                                  ResourceGetResponseResourceType = "google_compute_ha_vpn_gateway"
	ResourceGetResponseResourceTypeGoogleComputeForwardingRule                                ResourceGetResponseResourceType = "google_compute_forwarding_rule"
	ResourceGetResponseResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceGetResponseResourceType = "google_compute_network_firewall_policy"
	ResourceGetResponseResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceGetResponseResourceType = "google_compute_network_firewall_policy_rule"
	ResourceGetResponseResourceTypeCloudflareStaticRoute                                      ResourceGetResponseResourceType = "cloudflare_static_route"
	ResourceGetResponseResourceTypeCloudflareIPSECTunnel                                      ResourceGetResponseResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceGetResponseResourceType) IsKnown() bool {
	switch r {
	case ResourceGetResponseResourceTypeAwsCustomerGateway, ResourceGetResponseResourceTypeAwsEgressOnlyInternetGateway, ResourceGetResponseResourceTypeAwsInternetGateway, ResourceGetResponseResourceTypeAwsInstance, ResourceGetResponseResourceTypeAwsNetworkInterface, ResourceGetResponseResourceTypeAwsRoute, ResourceGetResponseResourceTypeAwsRouteTable, ResourceGetResponseResourceTypeAwsRouteTableAssociation, ResourceGetResponseResourceTypeAwsSubnet, ResourceGetResponseResourceTypeAwsVPC, ResourceGetResponseResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceGetResponseResourceTypeAwsVpnConnection, ResourceGetResponseResourceTypeAwsVpnConnectionRoute, ResourceGetResponseResourceTypeAwsVpnGateway, ResourceGetResponseResourceTypeAwsSecurityGroup, ResourceGetResponseResourceTypeAwsVPCSecurityGroupIngressRule, ResourceGetResponseResourceTypeAwsVPCSecurityGroupEgressRule, ResourceGetResponseResourceTypeAwsEc2ManagedPrefixList, ResourceGetResponseResourceTypeAwsEc2TransitGateway, ResourceGetResponseResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceGetResponseResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceGetResponseResourceTypeAzurermApplicationSecurityGroup, ResourceGetResponseResourceTypeAzurermLB, ResourceGetResponseResourceTypeAzurermLBBackendAddressPool, ResourceGetResponseResourceTypeAzurermLBNatPool, ResourceGetResponseResourceTypeAzurermLBNatRule, ResourceGetResponseResourceTypeAzurermLBRule, ResourceGetResponseResourceTypeAzurermLocalNetworkGateway, ResourceGetResponseResourceTypeAzurermNetworkInterface, ResourceGetResponseResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceGetResponseResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceGetResponseResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceGetResponseResourceTypeAzurermNetworkSecurityGroup, ResourceGetResponseResourceTypeAzurermPublicIP, ResourceGetResponseResourceTypeAzurermRoute, ResourceGetResponseResourceTypeAzurermRouteTable, ResourceGetResponseResourceTypeAzurermSubnet, ResourceGetResponseResourceTypeAzurermSubnetRouteTableAssociation, ResourceGetResponseResourceTypeAzurermVirtualMachine, ResourceGetResponseResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceGetResponseResourceTypeAzurermVirtualNetwork, ResourceGetResponseResourceTypeAzurermVirtualNetworkGateway, ResourceGetResponseResourceTypeGoogleComputeNetwork, ResourceGetResponseResourceTypeGoogleComputeSubnetwork, ResourceGetResponseResourceTypeGoogleComputeVpnGateway, ResourceGetResponseResourceTypeGoogleComputeVpnTunnel, ResourceGetResponseResourceTypeGoogleComputeRoute, ResourceGetResponseResourceTypeGoogleComputeAddress, ResourceGetResponseResourceTypeGoogleComputeGlobalAddress, ResourceGetResponseResourceTypeGoogleComputeRouter, ResourceGetResponseResourceTypeGoogleComputeInterconnectAttachment, ResourceGetResponseResourceTypeGoogleComputeHaVpnGateway, ResourceGetResponseResourceTypeGoogleComputeForwardingRule, ResourceGetResponseResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceGetResponseResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceGetResponseResourceTypeCloudflareStaticRoute, ResourceGetResponseResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceGetResponseSection struct {
	HiddenItems  []ResourceGetResponseSectionsHiddenItem  `json:"hidden_items,required"`
	Name         string                                   `json:"name,required"`
	VisibleItems []ResourceGetResponseSectionsVisibleItem `json:"visible_items,required"`
	HelpText     string                                   `json:"help_text"`
	JSON         resourceGetResponseSectionJSON           `json:"-"`
}

// resourceGetResponseSectionJSON contains the JSON metadata for the struct
// [ResourceGetResponseSection]
type resourceGetResponseSectionJSON struct {
	HiddenItems  apijson.Field
	Name         apijson.Field
	VisibleItems apijson.Field
	HelpText     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceGetResponseSection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseSectionsHiddenItem struct {
	HelpText string                                      `json:"helpText"`
	Name     string                                      `json:"name"`
	Value    ResourceGetResponseSectionsHiddenItemsValue `json:"value"`
	JSON     resourceGetResponseSectionsHiddenItemJSON   `json:"-"`
}

// resourceGetResponseSectionsHiddenItemJSON contains the JSON metadata for the
// struct [ResourceGetResponseSectionsHiddenItem]
type resourceGetResponseSectionsHiddenItemJSON struct {
	HelpText    apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseSectionsHiddenItemsValue struct {
	ItemType string `json:"item_type,required"`
	// This field can have the runtime type of
	// [[]ResourceGetResponseSectionsHiddenItemsValueMcnListItemList].
	List interface{} `json:"list"`
	// This field can have the runtime type of
	// [ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview].
	ResourcePreview interface{} `json:"resource_preview"`
	String          string      `json:"string"`
	Yaml            string      `json:"yaml"`
	// This field can have the runtime type of
	// [ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff].
	YamlDiff interface{}                                     `json:"yaml_diff"`
	JSON     resourceGetResponseSectionsHiddenItemsValueJSON `json:"-"`
	union    ResourceGetResponseSectionsHiddenItemsValueUnion
}

// resourceGetResponseSectionsHiddenItemsValueJSON contains the JSON metadata for
// the struct [ResourceGetResponseSectionsHiddenItemsValue]
type resourceGetResponseSectionsHiddenItemsValueJSON struct {
	ItemType        apijson.Field
	List            apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	Yaml            apijson.Field
	YamlDiff        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r resourceGetResponseSectionsHiddenItemsValueJSON) RawJSON() string {
	return r.raw
}

func (r *ResourceGetResponseSectionsHiddenItemsValue) UnmarshalJSON(data []byte) (err error) {
	*r = ResourceGetResponseSectionsHiddenItemsValue{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ResourceGetResponseSectionsHiddenItemsValueUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ResourceGetResponseSectionsHiddenItemsValueMcnStringItem],
// [ResourceGetResponseSectionsHiddenItemsValueMcnYamlItem],
// [ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItem],
// [ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItem],
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItem].
func (r ResourceGetResponseSectionsHiddenItemsValue) AsUnion() ResourceGetResponseSectionsHiddenItemsValueUnion {
	return r.union
}

// Union satisfied by [ResourceGetResponseSectionsHiddenItemsValueMcnStringItem],
// [ResourceGetResponseSectionsHiddenItemsValueMcnYamlItem],
// [ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItem],
// [ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItem] or
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItem].
type ResourceGetResponseSectionsHiddenItemsValueUnion interface {
	implementsResourceGetResponseSectionsHiddenItemsValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResourceGetResponseSectionsHiddenItemsValueUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsHiddenItemsValueMcnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsHiddenItemsValueMcnYamlItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsHiddenItemsValueMcnListItem{}),
		},
	)
}

type ResourceGetResponseSectionsHiddenItemsValueMcnStringItem struct {
	ItemType string                                                       `json:"item_type,required"`
	String   string                                                       `json:"string,required"`
	JSON     resourceGetResponseSectionsHiddenItemsValueMcnStringItemJSON `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnStringItemJSON contains the JSON
// metadata for the struct
// [ResourceGetResponseSectionsHiddenItemsValueMcnStringItem]
type resourceGetResponseSectionsHiddenItemsValueMcnStringItemJSON struct {
	ItemType    apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnStringItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnStringItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsHiddenItemsValueMcnStringItem) implementsResourceGetResponseSectionsHiddenItemsValue() {
}

type ResourceGetResponseSectionsHiddenItemsValueMcnYamlItem struct {
	ItemType string                                                     `json:"item_type,required"`
	Yaml     string                                                     `json:"yaml,required"`
	JSON     resourceGetResponseSectionsHiddenItemsValueMcnYamlItemJSON `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnYamlItemJSON contains the JSON
// metadata for the struct [ResourceGetResponseSectionsHiddenItemsValueMcnYamlItem]
type resourceGetResponseSectionsHiddenItemsValueMcnYamlItemJSON struct {
	ItemType    apijson.Field
	Yaml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnYamlItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnYamlItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsHiddenItemsValueMcnYamlItem) implementsResourceGetResponseSectionsHiddenItemsValue() {
}

type ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItem struct {
	ItemType string                                                             `json:"item_type,required"`
	YamlDiff ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff `json:"yaml_diff,required"`
	JSON     resourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemJSON     `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemJSON contains the JSON
// metadata for the struct
// [ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItem]
type resourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemJSON struct {
	ItemType    apijson.Field
	YamlDiff    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItem) implementsResourceGetResponseSectionsHiddenItemsValue() {
}

type ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff struct {
	Diff             string                                                                 `json:"diff,required"`
	LeftDescription  string                                                                 `json:"left_description,required"`
	LeftYaml         string                                                                 `json:"left_yaml,required"`
	RightDescription string                                                                 `json:"right_description,required"`
	RightYaml        string                                                                 `json:"right_yaml,required"`
	JSON             resourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiffJSON `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiffJSON contains
// the JSON metadata for the struct
// [ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff]
type resourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiffJSON struct {
	Diff             apijson.Field
	LeftDescription  apijson.Field
	LeftYaml         apijson.Field
	RightDescription apijson.Field
	RightYaml        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnYamlDiffItemYamlDiffJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItem struct {
	ItemType        string                                                                           `json:"item_type,required"`
	ResourcePreview ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview `json:"resource_preview,required"`
	JSON            resourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemJSON            `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemJSON contains
// the JSON metadata for the struct
// [ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItem]
type resourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItem) implementsResourceGetResponseSectionsHiddenItemsValue() {
}

type ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview struct {
	ID           string                                                                                       `json:"id,required" format:"uuid"`
	CloudType    ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType    `json:"cloud_type,required"`
	Detail       string                                                                                       `json:"detail,required"`
	Name         string                                                                                       `json:"name,required"`
	ResourceType ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType `json:"resource_type,required"`
	Title        string                                                                                       `json:"title,required"`
	JSON         resourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewJSON         `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewJSON
// contains the JSON metadata for the struct
// [ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview]
type resourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewJSON struct {
	ID           apijson.Field
	CloudType    apijson.Field
	Detail       apijson.Field
	Name         apijson.Field
	ResourceType apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType string

const (
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAws        ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "AWS"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAzure      ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "AZURE"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeGoogle     ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "GOOGLE"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "CLOUDFLARE"
)

func (r ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudType) IsKnown() bool {
	switch r {
	case ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAws, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAzure, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeGoogle, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType string

const (
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway                                         ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_customer_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway                               ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_egress_only_internet_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway                                         ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_internet_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance                                                ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_instance"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface                                        ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_network_interface"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute                                                   ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable                                              ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation                                   ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet                                                  ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_subnet"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC                                                     ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection                                           ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute                                      ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway                                              ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup                                           ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_security_group"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_egress_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList                                    ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_managed_prefix_list"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway                                       ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup                            ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_application_security_group"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB                                                  ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool                                ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_backend_address_pool"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool                                           ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_pool"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule                                           ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule                                              ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway                                 ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_local_network_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface                                    ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_security_group_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup                                ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_security_group"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP                                            ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_public_ip"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute                                               ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable                                          ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route_table"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet                                              ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation                         ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet_route_table_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine                                      ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_machine"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork                                      ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway                               ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork                                       ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork                                    ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_subnetwork"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway                                    ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel                                     ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_tunnel"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute                                         ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress                                       ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_address"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress                                 ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_global_address"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter                                        ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_router"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment                        ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_interconnect_attachment"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway                                  ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_ha_vpn_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule                                ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_forwarding_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute                                      ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_static_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel                                      ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceType) IsKnown() bool {
	switch r {
	case ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute, ResourceGetResponseSectionsHiddenItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceGetResponseSectionsHiddenItemsValueMcnListItem struct {
	ItemType string                                                       `json:"item_type,required"`
	List     []ResourceGetResponseSectionsHiddenItemsValueMcnListItemList `json:"list,required"`
	JSON     resourceGetResponseSectionsHiddenItemsValueMcnListItemJSON   `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnListItemJSON contains the JSON
// metadata for the struct [ResourceGetResponseSectionsHiddenItemsValueMcnListItem]
type resourceGetResponseSectionsHiddenItemsValueMcnListItemJSON struct {
	ItemType    apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnListItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsHiddenItemsValueMcnListItem) implementsResourceGetResponseSectionsHiddenItemsValue() {
}

type ResourceGetResponseSectionsHiddenItemsValueMcnListItemList struct {
	ItemType string `json:"item_type,required"`
	// This field can have the runtime type of
	// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview].
	ResourcePreview interface{}                                                    `json:"resource_preview"`
	String          string                                                         `json:"string"`
	JSON            resourceGetResponseSectionsHiddenItemsValueMcnListItemListJSON `json:"-"`
	union           ResourceGetResponseSectionsHiddenItemsValueMcnListItemListUnion
}

// resourceGetResponseSectionsHiddenItemsValueMcnListItemListJSON contains the JSON
// metadata for the struct
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemList]
type resourceGetResponseSectionsHiddenItemsValueMcnListItemListJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnListItemListJSON) RawJSON() string {
	return r.raw
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnListItemList) UnmarshalJSON(data []byte) (err error) {
	*r = ResourceGetResponseSectionsHiddenItemsValueMcnListItemList{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemListUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem],
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem].
func (r ResourceGetResponseSectionsHiddenItemsValueMcnListItemList) AsUnion() ResourceGetResponseSectionsHiddenItemsValueMcnListItemListUnion {
	return r.union
}

// Union satisfied by
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem] or
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem].
type ResourceGetResponseSectionsHiddenItemsValueMcnListItemListUnion interface {
	implementsResourceGetResponseSectionsHiddenItemsValueMcnListItemList()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResourceGetResponseSectionsHiddenItemsValueMcnListItemListUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem{}),
		},
	)
}

type ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem struct {
	ItemType string                                                                      `json:"item_type,required"`
	String   string                                                                      `json:"string,required"`
	JSON     resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItemJSON `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItemJSON
// contains the JSON metadata for the struct
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem]
type resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItemJSON struct {
	ItemType    apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnStringItem) implementsResourceGetResponseSectionsHiddenItemsValueMcnListItemList() {
}

type ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem struct {
	ItemType        string                                                                                          `json:"item_type,required"`
	ResourcePreview ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview `json:"resource_preview,required"`
	JSON            resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemJSON            `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemJSON
// contains the JSON metadata for the struct
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem]
type resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItem) implementsResourceGetResponseSectionsHiddenItemsValueMcnListItemList() {
}

type ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview struct {
	ID           string                                                                                                      `json:"id,required" format:"uuid"`
	CloudType    ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType    `json:"cloud_type,required"`
	Detail       string                                                                                                      `json:"detail,required"`
	Name         string                                                                                                      `json:"name,required"`
	ResourceType ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType `json:"resource_type,required"`
	Title        string                                                                                                      `json:"title,required"`
	JSON         resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON         `json:"-"`
}

// resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON
// contains the JSON metadata for the struct
// [ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview]
type resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON struct {
	ID           apijson.Field
	CloudType    apijson.Field
	Detail       apijson.Field
	Name         apijson.Field
	ResourceType apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType string

const (
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAws        ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "AWS"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAzure      ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "AZURE"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeGoogle     ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "GOOGLE"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "CLOUDFLARE"
)

func (r ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType) IsKnown() bool {
	switch r {
	case ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAws, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAzure, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeGoogle, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType string

const (
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway                                         ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_customer_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway                               ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_egress_only_internet_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway                                         ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_internet_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance                                                ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_instance"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface                                        ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_network_interface"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute                                                   ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable                                              ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation                                   ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet                                                  ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_subnet"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC                                                     ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection                                           ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute                                      ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway                                              ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup                                           ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_security_group"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_egress_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList                                    ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_managed_prefix_list"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway                                       ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup                            ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_application_security_group"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB                                                  ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool                                ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_backend_address_pool"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool                                           ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_pool"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule                                           ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule                                              ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway                                 ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_local_network_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface                                    ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_security_group_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup                                ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_security_group"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP                                            ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_public_ip"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute                                               ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable                                          ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route_table"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet                                              ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation                         ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet_route_table_association"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine                                      ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_machine"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork                                      ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway                               ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork                                       ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork                                    ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_subnetwork"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway                                    ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel                                     ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_tunnel"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute                                         ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress                                       ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_address"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress                                 ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_global_address"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter                                        ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_router"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment                        ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_interconnect_attachment"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway                                  ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_ha_vpn_gateway"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule                                ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_forwarding_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy_rule"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute                                      ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_static_route"
	ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel                                      ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType) IsKnown() bool {
	switch r {
	case ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute, ResourceGetResponseSectionsHiddenItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceGetResponseSectionsVisibleItem struct {
	HelpText string                                       `json:"helpText"`
	Name     string                                       `json:"name"`
	Value    ResourceGetResponseSectionsVisibleItemsValue `json:"value"`
	JSON     resourceGetResponseSectionsVisibleItemJSON   `json:"-"`
}

// resourceGetResponseSectionsVisibleItemJSON contains the JSON metadata for the
// struct [ResourceGetResponseSectionsVisibleItem]
type resourceGetResponseSectionsVisibleItemJSON struct {
	HelpText    apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseSectionsVisibleItemsValue struct {
	ItemType string `json:"item_type,required"`
	// This field can have the runtime type of
	// [[]ResourceGetResponseSectionsVisibleItemsValueMcnListItemList].
	List interface{} `json:"list"`
	// This field can have the runtime type of
	// [ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview].
	ResourcePreview interface{} `json:"resource_preview"`
	String          string      `json:"string"`
	Yaml            string      `json:"yaml"`
	// This field can have the runtime type of
	// [ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff].
	YamlDiff interface{}                                      `json:"yaml_diff"`
	JSON     resourceGetResponseSectionsVisibleItemsValueJSON `json:"-"`
	union    ResourceGetResponseSectionsVisibleItemsValueUnion
}

// resourceGetResponseSectionsVisibleItemsValueJSON contains the JSON metadata for
// the struct [ResourceGetResponseSectionsVisibleItemsValue]
type resourceGetResponseSectionsVisibleItemsValueJSON struct {
	ItemType        apijson.Field
	List            apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	Yaml            apijson.Field
	YamlDiff        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r resourceGetResponseSectionsVisibleItemsValueJSON) RawJSON() string {
	return r.raw
}

func (r *ResourceGetResponseSectionsVisibleItemsValue) UnmarshalJSON(data []byte) (err error) {
	*r = ResourceGetResponseSectionsVisibleItemsValue{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ResourceGetResponseSectionsVisibleItemsValueUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ResourceGetResponseSectionsVisibleItemsValueMcnStringItem],
// [ResourceGetResponseSectionsVisibleItemsValueMcnYamlItem],
// [ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItem],
// [ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItem],
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItem].
func (r ResourceGetResponseSectionsVisibleItemsValue) AsUnion() ResourceGetResponseSectionsVisibleItemsValueUnion {
	return r.union
}

// Union satisfied by [ResourceGetResponseSectionsVisibleItemsValueMcnStringItem],
// [ResourceGetResponseSectionsVisibleItemsValueMcnYamlItem],
// [ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItem],
// [ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItem] or
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItem].
type ResourceGetResponseSectionsVisibleItemsValueUnion interface {
	implementsResourceGetResponseSectionsVisibleItemsValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResourceGetResponseSectionsVisibleItemsValueUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsVisibleItemsValueMcnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsVisibleItemsValueMcnYamlItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsVisibleItemsValueMcnListItem{}),
		},
	)
}

type ResourceGetResponseSectionsVisibleItemsValueMcnStringItem struct {
	ItemType string                                                        `json:"item_type,required"`
	String   string                                                        `json:"string,required"`
	JSON     resourceGetResponseSectionsVisibleItemsValueMcnStringItemJSON `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnStringItemJSON contains the JSON
// metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnStringItem]
type resourceGetResponseSectionsVisibleItemsValueMcnStringItemJSON struct {
	ItemType    apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnStringItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnStringItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsVisibleItemsValueMcnStringItem) implementsResourceGetResponseSectionsVisibleItemsValue() {
}

type ResourceGetResponseSectionsVisibleItemsValueMcnYamlItem struct {
	ItemType string                                                      `json:"item_type,required"`
	Yaml     string                                                      `json:"yaml,required"`
	JSON     resourceGetResponseSectionsVisibleItemsValueMcnYamlItemJSON `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnYamlItemJSON contains the JSON
// metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnYamlItem]
type resourceGetResponseSectionsVisibleItemsValueMcnYamlItemJSON struct {
	ItemType    apijson.Field
	Yaml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnYamlItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnYamlItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsVisibleItemsValueMcnYamlItem) implementsResourceGetResponseSectionsVisibleItemsValue() {
}

type ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItem struct {
	ItemType string                                                              `json:"item_type,required"`
	YamlDiff ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff `json:"yaml_diff,required"`
	JSON     resourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemJSON     `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemJSON contains the
// JSON metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItem]
type resourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemJSON struct {
	ItemType    apijson.Field
	YamlDiff    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItem) implementsResourceGetResponseSectionsVisibleItemsValue() {
}

type ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff struct {
	Diff             string                                                                  `json:"diff,required"`
	LeftDescription  string                                                                  `json:"left_description,required"`
	LeftYaml         string                                                                  `json:"left_yaml,required"`
	RightDescription string                                                                  `json:"right_description,required"`
	RightYaml        string                                                                  `json:"right_yaml,required"`
	JSON             resourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiffJSON `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiffJSON contains
// the JSON metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff]
type resourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiffJSON struct {
	Diff             apijson.Field
	LeftDescription  apijson.Field
	LeftYaml         apijson.Field
	RightDescription apijson.Field
	RightYaml        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiff) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnYamlDiffItemYamlDiffJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItem struct {
	ItemType        string                                                                            `json:"item_type,required"`
	ResourcePreview ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview `json:"resource_preview,required"`
	JSON            resourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemJSON            `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemJSON contains
// the JSON metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItem]
type resourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItem) implementsResourceGetResponseSectionsVisibleItemsValue() {
}

type ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview struct {
	ID           string                                                                                        `json:"id,required" format:"uuid"`
	CloudType    ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType    `json:"cloud_type,required"`
	Detail       string                                                                                        `json:"detail,required"`
	Name         string                                                                                        `json:"name,required"`
	ResourceType ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType `json:"resource_type,required"`
	Title        string                                                                                        `json:"title,required"`
	JSON         resourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewJSON         `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewJSON
// contains the JSON metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview]
type resourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewJSON struct {
	ID           apijson.Field
	CloudType    apijson.Field
	Detail       apijson.Field
	Name         apijson.Field
	ResourceType apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType string

const (
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAws        ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "AWS"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAzure      ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "AZURE"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeGoogle     ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "GOOGLE"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType = "CLOUDFLARE"
)

func (r ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudType) IsKnown() bool {
	switch r {
	case ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAws, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeAzure, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeGoogle, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType string

const (
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway                                         ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_customer_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway                               ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_egress_only_internet_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway                                         ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_internet_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance                                                ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_instance"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface                                        ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_network_interface"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute                                                   ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable                                              ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation                                   ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet                                                  ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_subnet"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC                                                     ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection                                           ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute                                      ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway                                              ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup                                           ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_security_group"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_egress_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList                                    ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_managed_prefix_list"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway                                       ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup                            ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_application_security_group"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB                                                  ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool                                ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_backend_address_pool"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool                                           ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_pool"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule                                           ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule                                              ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway                                 ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_local_network_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface                                    ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_security_group_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup                                ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_security_group"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP                                            ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_public_ip"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute                                               ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable                                          ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route_table"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet                                              ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation                         ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet_route_table_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine                                      ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_machine"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork                                      ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway                               ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork                                       ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork                                    ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_subnetwork"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway                                    ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel                                     ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_tunnel"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute                                         ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress                                       ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_address"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress                                 ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_global_address"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter                                        ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_router"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment                        ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_interconnect_attachment"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway                                  ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_ha_vpn_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule                                ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_forwarding_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute                                      ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_static_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel                                      ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceType) IsKnown() bool {
	switch r {
	case ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute, ResourceGetResponseSectionsVisibleItemsValueMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceGetResponseSectionsVisibleItemsValueMcnListItem struct {
	ItemType string                                                        `json:"item_type,required"`
	List     []ResourceGetResponseSectionsVisibleItemsValueMcnListItemList `json:"list,required"`
	JSON     resourceGetResponseSectionsVisibleItemsValueMcnListItemJSON   `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnListItemJSON contains the JSON
// metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItem]
type resourceGetResponseSectionsVisibleItemsValueMcnListItemJSON struct {
	ItemType    apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnListItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsVisibleItemsValueMcnListItem) implementsResourceGetResponseSectionsVisibleItemsValue() {
}

type ResourceGetResponseSectionsVisibleItemsValueMcnListItemList struct {
	ItemType string `json:"item_type,required"`
	// This field can have the runtime type of
	// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview].
	ResourcePreview interface{}                                                     `json:"resource_preview"`
	String          string                                                          `json:"string"`
	JSON            resourceGetResponseSectionsVisibleItemsValueMcnListItemListJSON `json:"-"`
	union           ResourceGetResponseSectionsVisibleItemsValueMcnListItemListUnion
}

// resourceGetResponseSectionsVisibleItemsValueMcnListItemListJSON contains the
// JSON metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemList]
type resourceGetResponseSectionsVisibleItemsValueMcnListItemListJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnListItemListJSON) RawJSON() string {
	return r.raw
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnListItemList) UnmarshalJSON(data []byte) (err error) {
	*r = ResourceGetResponseSectionsVisibleItemsValueMcnListItemList{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemListUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem],
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem].
func (r ResourceGetResponseSectionsVisibleItemsValueMcnListItemList) AsUnion() ResourceGetResponseSectionsVisibleItemsValueMcnListItemListUnion {
	return r.union
}

// Union satisfied by
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem] or
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem].
type ResourceGetResponseSectionsVisibleItemsValueMcnListItemListUnion interface {
	implementsResourceGetResponseSectionsVisibleItemsValueMcnListItemList()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ResourceGetResponseSectionsVisibleItemsValueMcnListItemListUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem{}),
		},
	)
}

type ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem struct {
	ItemType string                                                                       `json:"item_type,required"`
	String   string                                                                       `json:"string,required"`
	JSON     resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItemJSON `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItemJSON
// contains the JSON metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem]
type resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItemJSON struct {
	ItemType    apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnStringItem) implementsResourceGetResponseSectionsVisibleItemsValueMcnListItemList() {
}

type ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem struct {
	ItemType        string                                                                                           `json:"item_type,required"`
	ResourcePreview ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview `json:"resource_preview,required"`
	JSON            resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemJSON            `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemJSON
// contains the JSON metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem]
type resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemJSON) RawJSON() string {
	return r.raw
}

func (r ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItem) implementsResourceGetResponseSectionsVisibleItemsValueMcnListItemList() {
}

type ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview struct {
	ID           string                                                                                                       `json:"id,required" format:"uuid"`
	CloudType    ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType    `json:"cloud_type,required"`
	Detail       string                                                                                                       `json:"detail,required"`
	Name         string                                                                                                       `json:"name,required"`
	ResourceType ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType `json:"resource_type,required"`
	Title        string                                                                                                       `json:"title,required"`
	JSON         resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON         `json:"-"`
}

// resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON
// contains the JSON metadata for the struct
// [ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview]
type resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON struct {
	ID           apijson.Field
	CloudType    apijson.Field
	Detail       apijson.Field
	Name         apijson.Field
	ResourceType apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType string

const (
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAws        ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "AWS"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAzure      ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "AZURE"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeGoogle     ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "GOOGLE"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType = "CLOUDFLARE"
)

func (r ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudType) IsKnown() bool {
	switch r {
	case ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAws, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeAzure, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeGoogle, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewCloudTypeCloudflare:
		return true
	}
	return false
}

type ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType string

const (
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway                                         ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_customer_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway                               ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_egress_only_internet_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway                                         ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_internet_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance                                                ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_instance"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface                                        ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_network_interface"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute                                                   ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable                                              ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation                                   ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_route_table_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet                                                  ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_subnet"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC                                                     ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection                                           ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute                                      ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_connection_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway                                              ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpn_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup                                           ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_security_group"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_vpc_security_group_egress_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList                                    ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_managed_prefix_list"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway                                       ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup                            ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_application_security_group"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB                                                  ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool                                ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_backend_address_pool"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool                                           ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_pool"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule                                           ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_nat_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule                                              ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_lb_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway                                 ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_local_network_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface                                    ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_interface_security_group_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup                                ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_network_security_group"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP                                            ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_public_ip"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute                                               ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable                                          ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_route_table"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet                                              ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation                         ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_subnet_route_table_association"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine                                      ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_machine"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork                                      ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway                               ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "azurerm_virtual_network_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork                                       ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork                                    ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_subnetwork"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway                                    ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel                                     ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_vpn_tunnel"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute                                         ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress                                       ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_address"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress                                 ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_global_address"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter                                        ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_router"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment                        ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_interconnect_attachment"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway                                  ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_ha_vpn_gateway"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule                                ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_forwarding_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "google_compute_network_firewall_policy_rule"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute                                      ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_static_route"
	ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel                                      ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceType) IsKnown() bool {
	switch r {
	case ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsCustomerGateway, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEgressOnlyInternetGateway, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInternetGateway, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsInstance, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsNetworkInterface, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRoute, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTable, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsRouteTableAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSubnet, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPC, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnection, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnConnectionRoute, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVpnGateway, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsSecurityGroup, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupIngressRule, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsVPCSecurityGroupEgressRule, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2ManagedPrefixList, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGateway, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermApplicationSecurityGroup, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLB, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBBackendAddressPool, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatPool, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBNatRule, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLBRule, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermLocalNetworkGateway, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterface, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermNetworkSecurityGroup, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermPublicIP, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRoute, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermRouteTable, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnet, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermSubnetRouteTableAssociation, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualMachine, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetwork, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeAzurermVirtualNetworkGateway, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetwork, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeSubnetwork, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnGateway, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeVpnTunnel, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRoute, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeAddress, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeGlobalAddress, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeRouter, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeInterconnectAttachment, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeHaVpnGateway, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeForwardingRule, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareStaticRoute, ResourceGetResponseSectionsVisibleItemsValueMcnListItemListMcnResourcePreviewItemResourcePreviewResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceGetResponseManagedBy struct {
	ID         string                                 `json:"id,required" format:"uuid"`
	ClientType ResourceGetResponseManagedByClientType `json:"client_type,required"`
	Name       string                                 `json:"name,required"`
	JSON       resourceGetResponseManagedByJSON       `json:"-"`
}

// resourceGetResponseManagedByJSON contains the JSON metadata for the struct
// [ResourceGetResponseManagedBy]
type resourceGetResponseManagedByJSON struct {
	ID          apijson.Field
	ClientType  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseManagedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseManagedByJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseManagedByClientType string

const (
	ResourceGetResponseManagedByClientTypeMagicWANCloudOnramp ResourceGetResponseManagedByClientType = "MAGIC_WAN_CLOUD_ONRAMP"
)

func (r ResourceGetResponseManagedByClientType) IsKnown() bool {
	switch r {
	case ResourceGetResponseManagedByClientTypeMagicWANCloudOnramp:
		return true
	}
	return false
}

type ResourceListParams struct {
	AccountID  param.Field[string] `path:"account_id,required"`
	Cloudflare param.Field[bool]   `query:"cloudflare"`
	Desc       param.Field[bool]   `query:"desc"`
	Managed    param.Field[bool]   `query:"managed"`
	// One of ["id", "resource_type", "region"].
	OrderBy       param.Field[string]                           `query:"order_by"`
	Page          param.Field[int64]                            `query:"page"`
	PerPage       param.Field[int64]                            `query:"per_page"`
	ProviderID    param.Field[string]                           `query:"provider_id"`
	Region        param.Field[string]                           `query:"region"`
	ResourceGroup param.Field[string]                           `query:"resource_group"`
	ResourceID    param.Field[[]string]                         `query:"resource_id" format:"uuid"`
	ResourceType  param.Field[[]ResourceListParamsResourceType] `query:"resource_type"`
	Search        param.Field[[]string]                         `query:"search"`
	V2            param.Field[bool]                             `query:"v2"`
}

// URLQuery serializes [ResourceListParams]'s query parameters as `url.Values`.
func (r ResourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ResourceListParamsResourceType string

const (
	ResourceListParamsResourceTypeAwsCustomerGateway                                         ResourceListParamsResourceType = "aws_customer_gateway"
	ResourceListParamsResourceTypeAwsEgressOnlyInternetGateway                               ResourceListParamsResourceType = "aws_egress_only_internet_gateway"
	ResourceListParamsResourceTypeAwsInternetGateway                                         ResourceListParamsResourceType = "aws_internet_gateway"
	ResourceListParamsResourceTypeAwsInstance                                                ResourceListParamsResourceType = "aws_instance"
	ResourceListParamsResourceTypeAwsNetworkInterface                                        ResourceListParamsResourceType = "aws_network_interface"
	ResourceListParamsResourceTypeAwsRoute                                                   ResourceListParamsResourceType = "aws_route"
	ResourceListParamsResourceTypeAwsRouteTable                                              ResourceListParamsResourceType = "aws_route_table"
	ResourceListParamsResourceTypeAwsRouteTableAssociation                                   ResourceListParamsResourceType = "aws_route_table_association"
	ResourceListParamsResourceTypeAwsSubnet                                                  ResourceListParamsResourceType = "aws_subnet"
	ResourceListParamsResourceTypeAwsVPC                                                     ResourceListParamsResourceType = "aws_vpc"
	ResourceListParamsResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceListParamsResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceListParamsResourceTypeAwsVpnConnection                                           ResourceListParamsResourceType = "aws_vpn_connection"
	ResourceListParamsResourceTypeAwsVpnConnectionRoute                                      ResourceListParamsResourceType = "aws_vpn_connection_route"
	ResourceListParamsResourceTypeAwsVpnGateway                                              ResourceListParamsResourceType = "aws_vpn_gateway"
	ResourceListParamsResourceTypeAwsSecurityGroup                                           ResourceListParamsResourceType = "aws_security_group"
	ResourceListParamsResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceListParamsResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceListParamsResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceListParamsResourceType = "aws_vpc_security_group_egress_rule"
	ResourceListParamsResourceTypeAwsEc2ManagedPrefixList                                    ResourceListParamsResourceType = "aws_ec2_managed_prefix_list"
	ResourceListParamsResourceTypeAwsEc2TransitGateway                                       ResourceListParamsResourceType = "aws_ec2_transit_gateway"
	ResourceListParamsResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceListParamsResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceListParamsResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceListParamsResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceListParamsResourceTypeAzurermApplicationSecurityGroup                            ResourceListParamsResourceType = "azurerm_application_security_group"
	ResourceListParamsResourceTypeAzurermLB                                                  ResourceListParamsResourceType = "azurerm_lb"
	ResourceListParamsResourceTypeAzurermLBBackendAddressPool                                ResourceListParamsResourceType = "azurerm_lb_backend_address_pool"
	ResourceListParamsResourceTypeAzurermLBNatPool                                           ResourceListParamsResourceType = "azurerm_lb_nat_pool"
	ResourceListParamsResourceTypeAzurermLBNatRule                                           ResourceListParamsResourceType = "azurerm_lb_nat_rule"
	ResourceListParamsResourceTypeAzurermLBRule                                              ResourceListParamsResourceType = "azurerm_lb_rule"
	ResourceListParamsResourceTypeAzurermLocalNetworkGateway                                 ResourceListParamsResourceType = "azurerm_local_network_gateway"
	ResourceListParamsResourceTypeAzurermNetworkInterface                                    ResourceListParamsResourceType = "azurerm_network_interface"
	ResourceListParamsResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceListParamsResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceListParamsResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceListParamsResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceListParamsResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceListParamsResourceType = "azurerm_network_interface_security_group_association"
	ResourceListParamsResourceTypeAzurermNetworkSecurityGroup                                ResourceListParamsResourceType = "azurerm_network_security_group"
	ResourceListParamsResourceTypeAzurermPublicIP                                            ResourceListParamsResourceType = "azurerm_public_ip"
	ResourceListParamsResourceTypeAzurermRoute                                               ResourceListParamsResourceType = "azurerm_route"
	ResourceListParamsResourceTypeAzurermRouteTable                                          ResourceListParamsResourceType = "azurerm_route_table"
	ResourceListParamsResourceTypeAzurermSubnet                                              ResourceListParamsResourceType = "azurerm_subnet"
	ResourceListParamsResourceTypeAzurermSubnetRouteTableAssociation                         ResourceListParamsResourceType = "azurerm_subnet_route_table_association"
	ResourceListParamsResourceTypeAzurermVirtualMachine                                      ResourceListParamsResourceType = "azurerm_virtual_machine"
	ResourceListParamsResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceListParamsResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceListParamsResourceTypeAzurermVirtualNetwork                                      ResourceListParamsResourceType = "azurerm_virtual_network"
	ResourceListParamsResourceTypeAzurermVirtualNetworkGateway                               ResourceListParamsResourceType = "azurerm_virtual_network_gateway"
	ResourceListParamsResourceTypeGoogleComputeNetwork                                       ResourceListParamsResourceType = "google_compute_network"
	ResourceListParamsResourceTypeGoogleComputeSubnetwork                                    ResourceListParamsResourceType = "google_compute_subnetwork"
	ResourceListParamsResourceTypeGoogleComputeVpnGateway                                    ResourceListParamsResourceType = "google_compute_vpn_gateway"
	ResourceListParamsResourceTypeGoogleComputeVpnTunnel                                     ResourceListParamsResourceType = "google_compute_vpn_tunnel"
	ResourceListParamsResourceTypeGoogleComputeRoute                                         ResourceListParamsResourceType = "google_compute_route"
	ResourceListParamsResourceTypeGoogleComputeAddress                                       ResourceListParamsResourceType = "google_compute_address"
	ResourceListParamsResourceTypeGoogleComputeGlobalAddress                                 ResourceListParamsResourceType = "google_compute_global_address"
	ResourceListParamsResourceTypeGoogleComputeRouter                                        ResourceListParamsResourceType = "google_compute_router"
	ResourceListParamsResourceTypeGoogleComputeInterconnectAttachment                        ResourceListParamsResourceType = "google_compute_interconnect_attachment"
	ResourceListParamsResourceTypeGoogleComputeHaVpnGateway                                  ResourceListParamsResourceType = "google_compute_ha_vpn_gateway"
	ResourceListParamsResourceTypeGoogleComputeForwardingRule                                ResourceListParamsResourceType = "google_compute_forwarding_rule"
	ResourceListParamsResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceListParamsResourceType = "google_compute_network_firewall_policy"
	ResourceListParamsResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceListParamsResourceType = "google_compute_network_firewall_policy_rule"
	ResourceListParamsResourceTypeCloudflareStaticRoute                                      ResourceListParamsResourceType = "cloudflare_static_route"
	ResourceListParamsResourceTypeCloudflareIPSECTunnel                                      ResourceListParamsResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceListParamsResourceType) IsKnown() bool {
	switch r {
	case ResourceListParamsResourceTypeAwsCustomerGateway, ResourceListParamsResourceTypeAwsEgressOnlyInternetGateway, ResourceListParamsResourceTypeAwsInternetGateway, ResourceListParamsResourceTypeAwsInstance, ResourceListParamsResourceTypeAwsNetworkInterface, ResourceListParamsResourceTypeAwsRoute, ResourceListParamsResourceTypeAwsRouteTable, ResourceListParamsResourceTypeAwsRouteTableAssociation, ResourceListParamsResourceTypeAwsSubnet, ResourceListParamsResourceTypeAwsVPC, ResourceListParamsResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceListParamsResourceTypeAwsVpnConnection, ResourceListParamsResourceTypeAwsVpnConnectionRoute, ResourceListParamsResourceTypeAwsVpnGateway, ResourceListParamsResourceTypeAwsSecurityGroup, ResourceListParamsResourceTypeAwsVPCSecurityGroupIngressRule, ResourceListParamsResourceTypeAwsVPCSecurityGroupEgressRule, ResourceListParamsResourceTypeAwsEc2ManagedPrefixList, ResourceListParamsResourceTypeAwsEc2TransitGateway, ResourceListParamsResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceListParamsResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceListParamsResourceTypeAzurermApplicationSecurityGroup, ResourceListParamsResourceTypeAzurermLB, ResourceListParamsResourceTypeAzurermLBBackendAddressPool, ResourceListParamsResourceTypeAzurermLBNatPool, ResourceListParamsResourceTypeAzurermLBNatRule, ResourceListParamsResourceTypeAzurermLBRule, ResourceListParamsResourceTypeAzurermLocalNetworkGateway, ResourceListParamsResourceTypeAzurermNetworkInterface, ResourceListParamsResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceListParamsResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceListParamsResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceListParamsResourceTypeAzurermNetworkSecurityGroup, ResourceListParamsResourceTypeAzurermPublicIP, ResourceListParamsResourceTypeAzurermRoute, ResourceListParamsResourceTypeAzurermRouteTable, ResourceListParamsResourceTypeAzurermSubnet, ResourceListParamsResourceTypeAzurermSubnetRouteTableAssociation, ResourceListParamsResourceTypeAzurermVirtualMachine, ResourceListParamsResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceListParamsResourceTypeAzurermVirtualNetwork, ResourceListParamsResourceTypeAzurermVirtualNetworkGateway, ResourceListParamsResourceTypeGoogleComputeNetwork, ResourceListParamsResourceTypeGoogleComputeSubnetwork, ResourceListParamsResourceTypeGoogleComputeVpnGateway, ResourceListParamsResourceTypeGoogleComputeVpnTunnel, ResourceListParamsResourceTypeGoogleComputeRoute, ResourceListParamsResourceTypeGoogleComputeAddress, ResourceListParamsResourceTypeGoogleComputeGlobalAddress, ResourceListParamsResourceTypeGoogleComputeRouter, ResourceListParamsResourceTypeGoogleComputeInterconnectAttachment, ResourceListParamsResourceTypeGoogleComputeHaVpnGateway, ResourceListParamsResourceTypeGoogleComputeForwardingRule, ResourceListParamsResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceListParamsResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceListParamsResourceTypeCloudflareStaticRoute, ResourceListParamsResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceExportParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Desc      param.Field[bool]   `query:"desc"`
	// One of ["id", "resource_type", "region"].
	OrderBy       param.Field[string]                             `query:"order_by"`
	ProviderID    param.Field[string]                             `query:"provider_id"`
	Region        param.Field[string]                             `query:"region"`
	ResourceGroup param.Field[string]                             `query:"resource_group"`
	ResourceID    param.Field[[]string]                           `query:"resource_id" format:"uuid"`
	ResourceType  param.Field[[]ResourceExportParamsResourceType] `query:"resource_type"`
	Search        param.Field[[]string]                           `query:"search"`
	V2            param.Field[bool]                               `query:"v2"`
}

// URLQuery serializes [ResourceExportParams]'s query parameters as `url.Values`.
func (r ResourceExportParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ResourceExportParamsResourceType string

const (
	ResourceExportParamsResourceTypeAwsCustomerGateway                                         ResourceExportParamsResourceType = "aws_customer_gateway"
	ResourceExportParamsResourceTypeAwsEgressOnlyInternetGateway                               ResourceExportParamsResourceType = "aws_egress_only_internet_gateway"
	ResourceExportParamsResourceTypeAwsInternetGateway                                         ResourceExportParamsResourceType = "aws_internet_gateway"
	ResourceExportParamsResourceTypeAwsInstance                                                ResourceExportParamsResourceType = "aws_instance"
	ResourceExportParamsResourceTypeAwsNetworkInterface                                        ResourceExportParamsResourceType = "aws_network_interface"
	ResourceExportParamsResourceTypeAwsRoute                                                   ResourceExportParamsResourceType = "aws_route"
	ResourceExportParamsResourceTypeAwsRouteTable                                              ResourceExportParamsResourceType = "aws_route_table"
	ResourceExportParamsResourceTypeAwsRouteTableAssociation                                   ResourceExportParamsResourceType = "aws_route_table_association"
	ResourceExportParamsResourceTypeAwsSubnet                                                  ResourceExportParamsResourceType = "aws_subnet"
	ResourceExportParamsResourceTypeAwsVPC                                                     ResourceExportParamsResourceType = "aws_vpc"
	ResourceExportParamsResourceTypeAwsVPCIPV4CIDRBlockAssociation                             ResourceExportParamsResourceType = "aws_vpc_ipv4_cidr_block_association"
	ResourceExportParamsResourceTypeAwsVpnConnection                                           ResourceExportParamsResourceType = "aws_vpn_connection"
	ResourceExportParamsResourceTypeAwsVpnConnectionRoute                                      ResourceExportParamsResourceType = "aws_vpn_connection_route"
	ResourceExportParamsResourceTypeAwsVpnGateway                                              ResourceExportParamsResourceType = "aws_vpn_gateway"
	ResourceExportParamsResourceTypeAwsSecurityGroup                                           ResourceExportParamsResourceType = "aws_security_group"
	ResourceExportParamsResourceTypeAwsVPCSecurityGroupIngressRule                             ResourceExportParamsResourceType = "aws_vpc_security_group_ingress_rule"
	ResourceExportParamsResourceTypeAwsVPCSecurityGroupEgressRule                              ResourceExportParamsResourceType = "aws_vpc_security_group_egress_rule"
	ResourceExportParamsResourceTypeAwsEc2ManagedPrefixList                                    ResourceExportParamsResourceType = "aws_ec2_managed_prefix_list"
	ResourceExportParamsResourceTypeAwsEc2TransitGateway                                       ResourceExportParamsResourceType = "aws_ec2_transit_gateway"
	ResourceExportParamsResourceTypeAwsEc2TransitGatewayPrefixListReference                    ResourceExportParamsResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	ResourceExportParamsResourceTypeAwsEc2TransitGatewayVPCAttachment                          ResourceExportParamsResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	ResourceExportParamsResourceTypeAzurermApplicationSecurityGroup                            ResourceExportParamsResourceType = "azurerm_application_security_group"
	ResourceExportParamsResourceTypeAzurermLB                                                  ResourceExportParamsResourceType = "azurerm_lb"
	ResourceExportParamsResourceTypeAzurermLBBackendAddressPool                                ResourceExportParamsResourceType = "azurerm_lb_backend_address_pool"
	ResourceExportParamsResourceTypeAzurermLBNatPool                                           ResourceExportParamsResourceType = "azurerm_lb_nat_pool"
	ResourceExportParamsResourceTypeAzurermLBNatRule                                           ResourceExportParamsResourceType = "azurerm_lb_nat_rule"
	ResourceExportParamsResourceTypeAzurermLBRule                                              ResourceExportParamsResourceType = "azurerm_lb_rule"
	ResourceExportParamsResourceTypeAzurermLocalNetworkGateway                                 ResourceExportParamsResourceType = "azurerm_local_network_gateway"
	ResourceExportParamsResourceTypeAzurermNetworkInterface                                    ResourceExportParamsResourceType = "azurerm_network_interface"
	ResourceExportParamsResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation ResourceExportParamsResourceType = "azurerm_network_interface_application_security_group_association"
	ResourceExportParamsResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       ResourceExportParamsResourceType = "azurerm_network_interface_backend_address_pool_association"
	ResourceExportParamsResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            ResourceExportParamsResourceType = "azurerm_network_interface_security_group_association"
	ResourceExportParamsResourceTypeAzurermNetworkSecurityGroup                                ResourceExportParamsResourceType = "azurerm_network_security_group"
	ResourceExportParamsResourceTypeAzurermPublicIP                                            ResourceExportParamsResourceType = "azurerm_public_ip"
	ResourceExportParamsResourceTypeAzurermRoute                                               ResourceExportParamsResourceType = "azurerm_route"
	ResourceExportParamsResourceTypeAzurermRouteTable                                          ResourceExportParamsResourceType = "azurerm_route_table"
	ResourceExportParamsResourceTypeAzurermSubnet                                              ResourceExportParamsResourceType = "azurerm_subnet"
	ResourceExportParamsResourceTypeAzurermSubnetRouteTableAssociation                         ResourceExportParamsResourceType = "azurerm_subnet_route_table_association"
	ResourceExportParamsResourceTypeAzurermVirtualMachine                                      ResourceExportParamsResourceType = "azurerm_virtual_machine"
	ResourceExportParamsResourceTypeAzurermVirtualNetworkGatewayConnection                     ResourceExportParamsResourceType = "azurerm_virtual_network_gateway_connection"
	ResourceExportParamsResourceTypeAzurermVirtualNetwork                                      ResourceExportParamsResourceType = "azurerm_virtual_network"
	ResourceExportParamsResourceTypeAzurermVirtualNetworkGateway                               ResourceExportParamsResourceType = "azurerm_virtual_network_gateway"
	ResourceExportParamsResourceTypeGoogleComputeNetwork                                       ResourceExportParamsResourceType = "google_compute_network"
	ResourceExportParamsResourceTypeGoogleComputeSubnetwork                                    ResourceExportParamsResourceType = "google_compute_subnetwork"
	ResourceExportParamsResourceTypeGoogleComputeVpnGateway                                    ResourceExportParamsResourceType = "google_compute_vpn_gateway"
	ResourceExportParamsResourceTypeGoogleComputeVpnTunnel                                     ResourceExportParamsResourceType = "google_compute_vpn_tunnel"
	ResourceExportParamsResourceTypeGoogleComputeRoute                                         ResourceExportParamsResourceType = "google_compute_route"
	ResourceExportParamsResourceTypeGoogleComputeAddress                                       ResourceExportParamsResourceType = "google_compute_address"
	ResourceExportParamsResourceTypeGoogleComputeGlobalAddress                                 ResourceExportParamsResourceType = "google_compute_global_address"
	ResourceExportParamsResourceTypeGoogleComputeRouter                                        ResourceExportParamsResourceType = "google_compute_router"
	ResourceExportParamsResourceTypeGoogleComputeInterconnectAttachment                        ResourceExportParamsResourceType = "google_compute_interconnect_attachment"
	ResourceExportParamsResourceTypeGoogleComputeHaVpnGateway                                  ResourceExportParamsResourceType = "google_compute_ha_vpn_gateway"
	ResourceExportParamsResourceTypeGoogleComputeForwardingRule                                ResourceExportParamsResourceType = "google_compute_forwarding_rule"
	ResourceExportParamsResourceTypeGoogleComputeNetworkFirewallPolicy                         ResourceExportParamsResourceType = "google_compute_network_firewall_policy"
	ResourceExportParamsResourceTypeGoogleComputeNetworkFirewallPolicyRule                     ResourceExportParamsResourceType = "google_compute_network_firewall_policy_rule"
	ResourceExportParamsResourceTypeCloudflareStaticRoute                                      ResourceExportParamsResourceType = "cloudflare_static_route"
	ResourceExportParamsResourceTypeCloudflareIPSECTunnel                                      ResourceExportParamsResourceType = "cloudflare_ipsec_tunnel"
)

func (r ResourceExportParamsResourceType) IsKnown() bool {
	switch r {
	case ResourceExportParamsResourceTypeAwsCustomerGateway, ResourceExportParamsResourceTypeAwsEgressOnlyInternetGateway, ResourceExportParamsResourceTypeAwsInternetGateway, ResourceExportParamsResourceTypeAwsInstance, ResourceExportParamsResourceTypeAwsNetworkInterface, ResourceExportParamsResourceTypeAwsRoute, ResourceExportParamsResourceTypeAwsRouteTable, ResourceExportParamsResourceTypeAwsRouteTableAssociation, ResourceExportParamsResourceTypeAwsSubnet, ResourceExportParamsResourceTypeAwsVPC, ResourceExportParamsResourceTypeAwsVPCIPV4CIDRBlockAssociation, ResourceExportParamsResourceTypeAwsVpnConnection, ResourceExportParamsResourceTypeAwsVpnConnectionRoute, ResourceExportParamsResourceTypeAwsVpnGateway, ResourceExportParamsResourceTypeAwsSecurityGroup, ResourceExportParamsResourceTypeAwsVPCSecurityGroupIngressRule, ResourceExportParamsResourceTypeAwsVPCSecurityGroupEgressRule, ResourceExportParamsResourceTypeAwsEc2ManagedPrefixList, ResourceExportParamsResourceTypeAwsEc2TransitGateway, ResourceExportParamsResourceTypeAwsEc2TransitGatewayPrefixListReference, ResourceExportParamsResourceTypeAwsEc2TransitGatewayVPCAttachment, ResourceExportParamsResourceTypeAzurermApplicationSecurityGroup, ResourceExportParamsResourceTypeAzurermLB, ResourceExportParamsResourceTypeAzurermLBBackendAddressPool, ResourceExportParamsResourceTypeAzurermLBNatPool, ResourceExportParamsResourceTypeAzurermLBNatRule, ResourceExportParamsResourceTypeAzurermLBRule, ResourceExportParamsResourceTypeAzurermLocalNetworkGateway, ResourceExportParamsResourceTypeAzurermNetworkInterface, ResourceExportParamsResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, ResourceExportParamsResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, ResourceExportParamsResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, ResourceExportParamsResourceTypeAzurermNetworkSecurityGroup, ResourceExportParamsResourceTypeAzurermPublicIP, ResourceExportParamsResourceTypeAzurermRoute, ResourceExportParamsResourceTypeAzurermRouteTable, ResourceExportParamsResourceTypeAzurermSubnet, ResourceExportParamsResourceTypeAzurermSubnetRouteTableAssociation, ResourceExportParamsResourceTypeAzurermVirtualMachine, ResourceExportParamsResourceTypeAzurermVirtualNetworkGatewayConnection, ResourceExportParamsResourceTypeAzurermVirtualNetwork, ResourceExportParamsResourceTypeAzurermVirtualNetworkGateway, ResourceExportParamsResourceTypeGoogleComputeNetwork, ResourceExportParamsResourceTypeGoogleComputeSubnetwork, ResourceExportParamsResourceTypeGoogleComputeVpnGateway, ResourceExportParamsResourceTypeGoogleComputeVpnTunnel, ResourceExportParamsResourceTypeGoogleComputeRoute, ResourceExportParamsResourceTypeGoogleComputeAddress, ResourceExportParamsResourceTypeGoogleComputeGlobalAddress, ResourceExportParamsResourceTypeGoogleComputeRouter, ResourceExportParamsResourceTypeGoogleComputeInterconnectAttachment, ResourceExportParamsResourceTypeGoogleComputeHaVpnGateway, ResourceExportParamsResourceTypeGoogleComputeForwardingRule, ResourceExportParamsResourceTypeGoogleComputeNetworkFirewallPolicy, ResourceExportParamsResourceTypeGoogleComputeNetworkFirewallPolicyRule, ResourceExportParamsResourceTypeCloudflareStaticRoute, ResourceExportParamsResourceTypeCloudflareIPSECTunnel:
		return true
	}
	return false
}

type ResourceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	V2        param.Field[bool]   `query:"v2"`
}

// URLQuery serializes [ResourceGetParams]'s query parameters as `url.Values`.
func (r ResourceGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ResourceGetResponseEnvelope struct {
	Errors   []ResourceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ResourceGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ResourceGetResponse                   `json:"result,required"`
	Success  bool                                  `json:"success,required"`
	JSON     resourceGetResponseEnvelopeJSON       `json:"-"`
}

// resourceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResourceGetResponseEnvelope]
type resourceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseEnvelopeErrors struct {
	Code             ResourceGetResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Meta             ResourceGetResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           ResourceGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ResourceGetResponseEnvelopeErrors]
type resourceGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseEnvelopeErrorsCode int64

const (
	ResourceGetResponseEnvelopeErrorsCode1001   ResourceGetResponseEnvelopeErrorsCode = 1001
	ResourceGetResponseEnvelopeErrorsCode1002   ResourceGetResponseEnvelopeErrorsCode = 1002
	ResourceGetResponseEnvelopeErrorsCode1003   ResourceGetResponseEnvelopeErrorsCode = 1003
	ResourceGetResponseEnvelopeErrorsCode1004   ResourceGetResponseEnvelopeErrorsCode = 1004
	ResourceGetResponseEnvelopeErrorsCode1005   ResourceGetResponseEnvelopeErrorsCode = 1005
	ResourceGetResponseEnvelopeErrorsCode1006   ResourceGetResponseEnvelopeErrorsCode = 1006
	ResourceGetResponseEnvelopeErrorsCode1007   ResourceGetResponseEnvelopeErrorsCode = 1007
	ResourceGetResponseEnvelopeErrorsCode1008   ResourceGetResponseEnvelopeErrorsCode = 1008
	ResourceGetResponseEnvelopeErrorsCode1009   ResourceGetResponseEnvelopeErrorsCode = 1009
	ResourceGetResponseEnvelopeErrorsCode1010   ResourceGetResponseEnvelopeErrorsCode = 1010
	ResourceGetResponseEnvelopeErrorsCode1011   ResourceGetResponseEnvelopeErrorsCode = 1011
	ResourceGetResponseEnvelopeErrorsCode1012   ResourceGetResponseEnvelopeErrorsCode = 1012
	ResourceGetResponseEnvelopeErrorsCode1013   ResourceGetResponseEnvelopeErrorsCode = 1013
	ResourceGetResponseEnvelopeErrorsCode1014   ResourceGetResponseEnvelopeErrorsCode = 1014
	ResourceGetResponseEnvelopeErrorsCode1015   ResourceGetResponseEnvelopeErrorsCode = 1015
	ResourceGetResponseEnvelopeErrorsCode1016   ResourceGetResponseEnvelopeErrorsCode = 1016
	ResourceGetResponseEnvelopeErrorsCode1017   ResourceGetResponseEnvelopeErrorsCode = 1017
	ResourceGetResponseEnvelopeErrorsCode2001   ResourceGetResponseEnvelopeErrorsCode = 2001
	ResourceGetResponseEnvelopeErrorsCode2002   ResourceGetResponseEnvelopeErrorsCode = 2002
	ResourceGetResponseEnvelopeErrorsCode2003   ResourceGetResponseEnvelopeErrorsCode = 2003
	ResourceGetResponseEnvelopeErrorsCode2004   ResourceGetResponseEnvelopeErrorsCode = 2004
	ResourceGetResponseEnvelopeErrorsCode2005   ResourceGetResponseEnvelopeErrorsCode = 2005
	ResourceGetResponseEnvelopeErrorsCode2006   ResourceGetResponseEnvelopeErrorsCode = 2006
	ResourceGetResponseEnvelopeErrorsCode2007   ResourceGetResponseEnvelopeErrorsCode = 2007
	ResourceGetResponseEnvelopeErrorsCode2008   ResourceGetResponseEnvelopeErrorsCode = 2008
	ResourceGetResponseEnvelopeErrorsCode2009   ResourceGetResponseEnvelopeErrorsCode = 2009
	ResourceGetResponseEnvelopeErrorsCode2010   ResourceGetResponseEnvelopeErrorsCode = 2010
	ResourceGetResponseEnvelopeErrorsCode2011   ResourceGetResponseEnvelopeErrorsCode = 2011
	ResourceGetResponseEnvelopeErrorsCode2012   ResourceGetResponseEnvelopeErrorsCode = 2012
	ResourceGetResponseEnvelopeErrorsCode2013   ResourceGetResponseEnvelopeErrorsCode = 2013
	ResourceGetResponseEnvelopeErrorsCode2014   ResourceGetResponseEnvelopeErrorsCode = 2014
	ResourceGetResponseEnvelopeErrorsCode2015   ResourceGetResponseEnvelopeErrorsCode = 2015
	ResourceGetResponseEnvelopeErrorsCode2016   ResourceGetResponseEnvelopeErrorsCode = 2016
	ResourceGetResponseEnvelopeErrorsCode2017   ResourceGetResponseEnvelopeErrorsCode = 2017
	ResourceGetResponseEnvelopeErrorsCode2018   ResourceGetResponseEnvelopeErrorsCode = 2018
	ResourceGetResponseEnvelopeErrorsCode2019   ResourceGetResponseEnvelopeErrorsCode = 2019
	ResourceGetResponseEnvelopeErrorsCode2020   ResourceGetResponseEnvelopeErrorsCode = 2020
	ResourceGetResponseEnvelopeErrorsCode2021   ResourceGetResponseEnvelopeErrorsCode = 2021
	ResourceGetResponseEnvelopeErrorsCode2022   ResourceGetResponseEnvelopeErrorsCode = 2022
	ResourceGetResponseEnvelopeErrorsCode3001   ResourceGetResponseEnvelopeErrorsCode = 3001
	ResourceGetResponseEnvelopeErrorsCode3002   ResourceGetResponseEnvelopeErrorsCode = 3002
	ResourceGetResponseEnvelopeErrorsCode3003   ResourceGetResponseEnvelopeErrorsCode = 3003
	ResourceGetResponseEnvelopeErrorsCode3004   ResourceGetResponseEnvelopeErrorsCode = 3004
	ResourceGetResponseEnvelopeErrorsCode3005   ResourceGetResponseEnvelopeErrorsCode = 3005
	ResourceGetResponseEnvelopeErrorsCode3006   ResourceGetResponseEnvelopeErrorsCode = 3006
	ResourceGetResponseEnvelopeErrorsCode3007   ResourceGetResponseEnvelopeErrorsCode = 3007
	ResourceGetResponseEnvelopeErrorsCode4001   ResourceGetResponseEnvelopeErrorsCode = 4001
	ResourceGetResponseEnvelopeErrorsCode4002   ResourceGetResponseEnvelopeErrorsCode = 4002
	ResourceGetResponseEnvelopeErrorsCode4003   ResourceGetResponseEnvelopeErrorsCode = 4003
	ResourceGetResponseEnvelopeErrorsCode4004   ResourceGetResponseEnvelopeErrorsCode = 4004
	ResourceGetResponseEnvelopeErrorsCode4005   ResourceGetResponseEnvelopeErrorsCode = 4005
	ResourceGetResponseEnvelopeErrorsCode4006   ResourceGetResponseEnvelopeErrorsCode = 4006
	ResourceGetResponseEnvelopeErrorsCode4007   ResourceGetResponseEnvelopeErrorsCode = 4007
	ResourceGetResponseEnvelopeErrorsCode4008   ResourceGetResponseEnvelopeErrorsCode = 4008
	ResourceGetResponseEnvelopeErrorsCode4009   ResourceGetResponseEnvelopeErrorsCode = 4009
	ResourceGetResponseEnvelopeErrorsCode4010   ResourceGetResponseEnvelopeErrorsCode = 4010
	ResourceGetResponseEnvelopeErrorsCode4011   ResourceGetResponseEnvelopeErrorsCode = 4011
	ResourceGetResponseEnvelopeErrorsCode4012   ResourceGetResponseEnvelopeErrorsCode = 4012
	ResourceGetResponseEnvelopeErrorsCode4013   ResourceGetResponseEnvelopeErrorsCode = 4013
	ResourceGetResponseEnvelopeErrorsCode4014   ResourceGetResponseEnvelopeErrorsCode = 4014
	ResourceGetResponseEnvelopeErrorsCode4015   ResourceGetResponseEnvelopeErrorsCode = 4015
	ResourceGetResponseEnvelopeErrorsCode4016   ResourceGetResponseEnvelopeErrorsCode = 4016
	ResourceGetResponseEnvelopeErrorsCode4017   ResourceGetResponseEnvelopeErrorsCode = 4017
	ResourceGetResponseEnvelopeErrorsCode4018   ResourceGetResponseEnvelopeErrorsCode = 4018
	ResourceGetResponseEnvelopeErrorsCode4019   ResourceGetResponseEnvelopeErrorsCode = 4019
	ResourceGetResponseEnvelopeErrorsCode4020   ResourceGetResponseEnvelopeErrorsCode = 4020
	ResourceGetResponseEnvelopeErrorsCode4021   ResourceGetResponseEnvelopeErrorsCode = 4021
	ResourceGetResponseEnvelopeErrorsCode4022   ResourceGetResponseEnvelopeErrorsCode = 4022
	ResourceGetResponseEnvelopeErrorsCode4023   ResourceGetResponseEnvelopeErrorsCode = 4023
	ResourceGetResponseEnvelopeErrorsCode5001   ResourceGetResponseEnvelopeErrorsCode = 5001
	ResourceGetResponseEnvelopeErrorsCode5002   ResourceGetResponseEnvelopeErrorsCode = 5002
	ResourceGetResponseEnvelopeErrorsCode5003   ResourceGetResponseEnvelopeErrorsCode = 5003
	ResourceGetResponseEnvelopeErrorsCode5004   ResourceGetResponseEnvelopeErrorsCode = 5004
	ResourceGetResponseEnvelopeErrorsCode102000 ResourceGetResponseEnvelopeErrorsCode = 102000
	ResourceGetResponseEnvelopeErrorsCode102001 ResourceGetResponseEnvelopeErrorsCode = 102001
	ResourceGetResponseEnvelopeErrorsCode102002 ResourceGetResponseEnvelopeErrorsCode = 102002
	ResourceGetResponseEnvelopeErrorsCode102003 ResourceGetResponseEnvelopeErrorsCode = 102003
	ResourceGetResponseEnvelopeErrorsCode102004 ResourceGetResponseEnvelopeErrorsCode = 102004
	ResourceGetResponseEnvelopeErrorsCode102005 ResourceGetResponseEnvelopeErrorsCode = 102005
	ResourceGetResponseEnvelopeErrorsCode102006 ResourceGetResponseEnvelopeErrorsCode = 102006
	ResourceGetResponseEnvelopeErrorsCode102007 ResourceGetResponseEnvelopeErrorsCode = 102007
	ResourceGetResponseEnvelopeErrorsCode102008 ResourceGetResponseEnvelopeErrorsCode = 102008
	ResourceGetResponseEnvelopeErrorsCode102009 ResourceGetResponseEnvelopeErrorsCode = 102009
	ResourceGetResponseEnvelopeErrorsCode102010 ResourceGetResponseEnvelopeErrorsCode = 102010
	ResourceGetResponseEnvelopeErrorsCode102011 ResourceGetResponseEnvelopeErrorsCode = 102011
	ResourceGetResponseEnvelopeErrorsCode102012 ResourceGetResponseEnvelopeErrorsCode = 102012
	ResourceGetResponseEnvelopeErrorsCode102013 ResourceGetResponseEnvelopeErrorsCode = 102013
	ResourceGetResponseEnvelopeErrorsCode102014 ResourceGetResponseEnvelopeErrorsCode = 102014
	ResourceGetResponseEnvelopeErrorsCode102015 ResourceGetResponseEnvelopeErrorsCode = 102015
	ResourceGetResponseEnvelopeErrorsCode102016 ResourceGetResponseEnvelopeErrorsCode = 102016
	ResourceGetResponseEnvelopeErrorsCode102017 ResourceGetResponseEnvelopeErrorsCode = 102017
	ResourceGetResponseEnvelopeErrorsCode102018 ResourceGetResponseEnvelopeErrorsCode = 102018
	ResourceGetResponseEnvelopeErrorsCode102019 ResourceGetResponseEnvelopeErrorsCode = 102019
	ResourceGetResponseEnvelopeErrorsCode102020 ResourceGetResponseEnvelopeErrorsCode = 102020
	ResourceGetResponseEnvelopeErrorsCode102021 ResourceGetResponseEnvelopeErrorsCode = 102021
	ResourceGetResponseEnvelopeErrorsCode102022 ResourceGetResponseEnvelopeErrorsCode = 102022
	ResourceGetResponseEnvelopeErrorsCode102023 ResourceGetResponseEnvelopeErrorsCode = 102023
	ResourceGetResponseEnvelopeErrorsCode102024 ResourceGetResponseEnvelopeErrorsCode = 102024
	ResourceGetResponseEnvelopeErrorsCode102025 ResourceGetResponseEnvelopeErrorsCode = 102025
	ResourceGetResponseEnvelopeErrorsCode102026 ResourceGetResponseEnvelopeErrorsCode = 102026
	ResourceGetResponseEnvelopeErrorsCode102027 ResourceGetResponseEnvelopeErrorsCode = 102027
	ResourceGetResponseEnvelopeErrorsCode102028 ResourceGetResponseEnvelopeErrorsCode = 102028
	ResourceGetResponseEnvelopeErrorsCode102029 ResourceGetResponseEnvelopeErrorsCode = 102029
	ResourceGetResponseEnvelopeErrorsCode102030 ResourceGetResponseEnvelopeErrorsCode = 102030
	ResourceGetResponseEnvelopeErrorsCode102031 ResourceGetResponseEnvelopeErrorsCode = 102031
	ResourceGetResponseEnvelopeErrorsCode102032 ResourceGetResponseEnvelopeErrorsCode = 102032
	ResourceGetResponseEnvelopeErrorsCode102033 ResourceGetResponseEnvelopeErrorsCode = 102033
	ResourceGetResponseEnvelopeErrorsCode102034 ResourceGetResponseEnvelopeErrorsCode = 102034
	ResourceGetResponseEnvelopeErrorsCode102035 ResourceGetResponseEnvelopeErrorsCode = 102035
	ResourceGetResponseEnvelopeErrorsCode102036 ResourceGetResponseEnvelopeErrorsCode = 102036
	ResourceGetResponseEnvelopeErrorsCode102037 ResourceGetResponseEnvelopeErrorsCode = 102037
	ResourceGetResponseEnvelopeErrorsCode102038 ResourceGetResponseEnvelopeErrorsCode = 102038
	ResourceGetResponseEnvelopeErrorsCode102039 ResourceGetResponseEnvelopeErrorsCode = 102039
	ResourceGetResponseEnvelopeErrorsCode102040 ResourceGetResponseEnvelopeErrorsCode = 102040
	ResourceGetResponseEnvelopeErrorsCode102041 ResourceGetResponseEnvelopeErrorsCode = 102041
	ResourceGetResponseEnvelopeErrorsCode102042 ResourceGetResponseEnvelopeErrorsCode = 102042
	ResourceGetResponseEnvelopeErrorsCode102043 ResourceGetResponseEnvelopeErrorsCode = 102043
	ResourceGetResponseEnvelopeErrorsCode102044 ResourceGetResponseEnvelopeErrorsCode = 102044
	ResourceGetResponseEnvelopeErrorsCode102045 ResourceGetResponseEnvelopeErrorsCode = 102045
	ResourceGetResponseEnvelopeErrorsCode102046 ResourceGetResponseEnvelopeErrorsCode = 102046
	ResourceGetResponseEnvelopeErrorsCode102047 ResourceGetResponseEnvelopeErrorsCode = 102047
	ResourceGetResponseEnvelopeErrorsCode102048 ResourceGetResponseEnvelopeErrorsCode = 102048
	ResourceGetResponseEnvelopeErrorsCode102049 ResourceGetResponseEnvelopeErrorsCode = 102049
	ResourceGetResponseEnvelopeErrorsCode102050 ResourceGetResponseEnvelopeErrorsCode = 102050
	ResourceGetResponseEnvelopeErrorsCode102051 ResourceGetResponseEnvelopeErrorsCode = 102051
	ResourceGetResponseEnvelopeErrorsCode102052 ResourceGetResponseEnvelopeErrorsCode = 102052
	ResourceGetResponseEnvelopeErrorsCode102053 ResourceGetResponseEnvelopeErrorsCode = 102053
	ResourceGetResponseEnvelopeErrorsCode102054 ResourceGetResponseEnvelopeErrorsCode = 102054
	ResourceGetResponseEnvelopeErrorsCode102055 ResourceGetResponseEnvelopeErrorsCode = 102055
	ResourceGetResponseEnvelopeErrorsCode102056 ResourceGetResponseEnvelopeErrorsCode = 102056
	ResourceGetResponseEnvelopeErrorsCode102057 ResourceGetResponseEnvelopeErrorsCode = 102057
	ResourceGetResponseEnvelopeErrorsCode102058 ResourceGetResponseEnvelopeErrorsCode = 102058
	ResourceGetResponseEnvelopeErrorsCode102059 ResourceGetResponseEnvelopeErrorsCode = 102059
	ResourceGetResponseEnvelopeErrorsCode102060 ResourceGetResponseEnvelopeErrorsCode = 102060
	ResourceGetResponseEnvelopeErrorsCode102061 ResourceGetResponseEnvelopeErrorsCode = 102061
	ResourceGetResponseEnvelopeErrorsCode102062 ResourceGetResponseEnvelopeErrorsCode = 102062
	ResourceGetResponseEnvelopeErrorsCode102063 ResourceGetResponseEnvelopeErrorsCode = 102063
	ResourceGetResponseEnvelopeErrorsCode102064 ResourceGetResponseEnvelopeErrorsCode = 102064
	ResourceGetResponseEnvelopeErrorsCode102065 ResourceGetResponseEnvelopeErrorsCode = 102065
	ResourceGetResponseEnvelopeErrorsCode102066 ResourceGetResponseEnvelopeErrorsCode = 102066
	ResourceGetResponseEnvelopeErrorsCode103001 ResourceGetResponseEnvelopeErrorsCode = 103001
	ResourceGetResponseEnvelopeErrorsCode103002 ResourceGetResponseEnvelopeErrorsCode = 103002
	ResourceGetResponseEnvelopeErrorsCode103003 ResourceGetResponseEnvelopeErrorsCode = 103003
	ResourceGetResponseEnvelopeErrorsCode103004 ResourceGetResponseEnvelopeErrorsCode = 103004
	ResourceGetResponseEnvelopeErrorsCode103005 ResourceGetResponseEnvelopeErrorsCode = 103005
	ResourceGetResponseEnvelopeErrorsCode103006 ResourceGetResponseEnvelopeErrorsCode = 103006
	ResourceGetResponseEnvelopeErrorsCode103007 ResourceGetResponseEnvelopeErrorsCode = 103007
	ResourceGetResponseEnvelopeErrorsCode103008 ResourceGetResponseEnvelopeErrorsCode = 103008
)

func (r ResourceGetResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case ResourceGetResponseEnvelopeErrorsCode1001, ResourceGetResponseEnvelopeErrorsCode1002, ResourceGetResponseEnvelopeErrorsCode1003, ResourceGetResponseEnvelopeErrorsCode1004, ResourceGetResponseEnvelopeErrorsCode1005, ResourceGetResponseEnvelopeErrorsCode1006, ResourceGetResponseEnvelopeErrorsCode1007, ResourceGetResponseEnvelopeErrorsCode1008, ResourceGetResponseEnvelopeErrorsCode1009, ResourceGetResponseEnvelopeErrorsCode1010, ResourceGetResponseEnvelopeErrorsCode1011, ResourceGetResponseEnvelopeErrorsCode1012, ResourceGetResponseEnvelopeErrorsCode1013, ResourceGetResponseEnvelopeErrorsCode1014, ResourceGetResponseEnvelopeErrorsCode1015, ResourceGetResponseEnvelopeErrorsCode1016, ResourceGetResponseEnvelopeErrorsCode1017, ResourceGetResponseEnvelopeErrorsCode2001, ResourceGetResponseEnvelopeErrorsCode2002, ResourceGetResponseEnvelopeErrorsCode2003, ResourceGetResponseEnvelopeErrorsCode2004, ResourceGetResponseEnvelopeErrorsCode2005, ResourceGetResponseEnvelopeErrorsCode2006, ResourceGetResponseEnvelopeErrorsCode2007, ResourceGetResponseEnvelopeErrorsCode2008, ResourceGetResponseEnvelopeErrorsCode2009, ResourceGetResponseEnvelopeErrorsCode2010, ResourceGetResponseEnvelopeErrorsCode2011, ResourceGetResponseEnvelopeErrorsCode2012, ResourceGetResponseEnvelopeErrorsCode2013, ResourceGetResponseEnvelopeErrorsCode2014, ResourceGetResponseEnvelopeErrorsCode2015, ResourceGetResponseEnvelopeErrorsCode2016, ResourceGetResponseEnvelopeErrorsCode2017, ResourceGetResponseEnvelopeErrorsCode2018, ResourceGetResponseEnvelopeErrorsCode2019, ResourceGetResponseEnvelopeErrorsCode2020, ResourceGetResponseEnvelopeErrorsCode2021, ResourceGetResponseEnvelopeErrorsCode2022, ResourceGetResponseEnvelopeErrorsCode3001, ResourceGetResponseEnvelopeErrorsCode3002, ResourceGetResponseEnvelopeErrorsCode3003, ResourceGetResponseEnvelopeErrorsCode3004, ResourceGetResponseEnvelopeErrorsCode3005, ResourceGetResponseEnvelopeErrorsCode3006, ResourceGetResponseEnvelopeErrorsCode3007, ResourceGetResponseEnvelopeErrorsCode4001, ResourceGetResponseEnvelopeErrorsCode4002, ResourceGetResponseEnvelopeErrorsCode4003, ResourceGetResponseEnvelopeErrorsCode4004, ResourceGetResponseEnvelopeErrorsCode4005, ResourceGetResponseEnvelopeErrorsCode4006, ResourceGetResponseEnvelopeErrorsCode4007, ResourceGetResponseEnvelopeErrorsCode4008, ResourceGetResponseEnvelopeErrorsCode4009, ResourceGetResponseEnvelopeErrorsCode4010, ResourceGetResponseEnvelopeErrorsCode4011, ResourceGetResponseEnvelopeErrorsCode4012, ResourceGetResponseEnvelopeErrorsCode4013, ResourceGetResponseEnvelopeErrorsCode4014, ResourceGetResponseEnvelopeErrorsCode4015, ResourceGetResponseEnvelopeErrorsCode4016, ResourceGetResponseEnvelopeErrorsCode4017, ResourceGetResponseEnvelopeErrorsCode4018, ResourceGetResponseEnvelopeErrorsCode4019, ResourceGetResponseEnvelopeErrorsCode4020, ResourceGetResponseEnvelopeErrorsCode4021, ResourceGetResponseEnvelopeErrorsCode4022, ResourceGetResponseEnvelopeErrorsCode4023, ResourceGetResponseEnvelopeErrorsCode5001, ResourceGetResponseEnvelopeErrorsCode5002, ResourceGetResponseEnvelopeErrorsCode5003, ResourceGetResponseEnvelopeErrorsCode5004, ResourceGetResponseEnvelopeErrorsCode102000, ResourceGetResponseEnvelopeErrorsCode102001, ResourceGetResponseEnvelopeErrorsCode102002, ResourceGetResponseEnvelopeErrorsCode102003, ResourceGetResponseEnvelopeErrorsCode102004, ResourceGetResponseEnvelopeErrorsCode102005, ResourceGetResponseEnvelopeErrorsCode102006, ResourceGetResponseEnvelopeErrorsCode102007, ResourceGetResponseEnvelopeErrorsCode102008, ResourceGetResponseEnvelopeErrorsCode102009, ResourceGetResponseEnvelopeErrorsCode102010, ResourceGetResponseEnvelopeErrorsCode102011, ResourceGetResponseEnvelopeErrorsCode102012, ResourceGetResponseEnvelopeErrorsCode102013, ResourceGetResponseEnvelopeErrorsCode102014, ResourceGetResponseEnvelopeErrorsCode102015, ResourceGetResponseEnvelopeErrorsCode102016, ResourceGetResponseEnvelopeErrorsCode102017, ResourceGetResponseEnvelopeErrorsCode102018, ResourceGetResponseEnvelopeErrorsCode102019, ResourceGetResponseEnvelopeErrorsCode102020, ResourceGetResponseEnvelopeErrorsCode102021, ResourceGetResponseEnvelopeErrorsCode102022, ResourceGetResponseEnvelopeErrorsCode102023, ResourceGetResponseEnvelopeErrorsCode102024, ResourceGetResponseEnvelopeErrorsCode102025, ResourceGetResponseEnvelopeErrorsCode102026, ResourceGetResponseEnvelopeErrorsCode102027, ResourceGetResponseEnvelopeErrorsCode102028, ResourceGetResponseEnvelopeErrorsCode102029, ResourceGetResponseEnvelopeErrorsCode102030, ResourceGetResponseEnvelopeErrorsCode102031, ResourceGetResponseEnvelopeErrorsCode102032, ResourceGetResponseEnvelopeErrorsCode102033, ResourceGetResponseEnvelopeErrorsCode102034, ResourceGetResponseEnvelopeErrorsCode102035, ResourceGetResponseEnvelopeErrorsCode102036, ResourceGetResponseEnvelopeErrorsCode102037, ResourceGetResponseEnvelopeErrorsCode102038, ResourceGetResponseEnvelopeErrorsCode102039, ResourceGetResponseEnvelopeErrorsCode102040, ResourceGetResponseEnvelopeErrorsCode102041, ResourceGetResponseEnvelopeErrorsCode102042, ResourceGetResponseEnvelopeErrorsCode102043, ResourceGetResponseEnvelopeErrorsCode102044, ResourceGetResponseEnvelopeErrorsCode102045, ResourceGetResponseEnvelopeErrorsCode102046, ResourceGetResponseEnvelopeErrorsCode102047, ResourceGetResponseEnvelopeErrorsCode102048, ResourceGetResponseEnvelopeErrorsCode102049, ResourceGetResponseEnvelopeErrorsCode102050, ResourceGetResponseEnvelopeErrorsCode102051, ResourceGetResponseEnvelopeErrorsCode102052, ResourceGetResponseEnvelopeErrorsCode102053, ResourceGetResponseEnvelopeErrorsCode102054, ResourceGetResponseEnvelopeErrorsCode102055, ResourceGetResponseEnvelopeErrorsCode102056, ResourceGetResponseEnvelopeErrorsCode102057, ResourceGetResponseEnvelopeErrorsCode102058, ResourceGetResponseEnvelopeErrorsCode102059, ResourceGetResponseEnvelopeErrorsCode102060, ResourceGetResponseEnvelopeErrorsCode102061, ResourceGetResponseEnvelopeErrorsCode102062, ResourceGetResponseEnvelopeErrorsCode102063, ResourceGetResponseEnvelopeErrorsCode102064, ResourceGetResponseEnvelopeErrorsCode102065, ResourceGetResponseEnvelopeErrorsCode102066, ResourceGetResponseEnvelopeErrorsCode103001, ResourceGetResponseEnvelopeErrorsCode103002, ResourceGetResponseEnvelopeErrorsCode103003, ResourceGetResponseEnvelopeErrorsCode103004, ResourceGetResponseEnvelopeErrorsCode103005, ResourceGetResponseEnvelopeErrorsCode103006, ResourceGetResponseEnvelopeErrorsCode103007, ResourceGetResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type ResourceGetResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                    `json:"l10n_key"`
	LoggableError string                                    `json:"loggable_error"`
	TemplateData  interface{}                               `json:"template_data"`
	TraceID       string                                    `json:"trace_id"`
	JSON          resourceGetResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// resourceGetResponseEnvelopeErrorsMetaJSON contains the JSON metadata for the
// struct [ResourceGetResponseEnvelopeErrorsMeta]
type resourceGetResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceGetResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseEnvelopeErrorsSource struct {
	Parameter           string                                      `json:"parameter"`
	ParameterValueIndex int64                                       `json:"parameter_value_index"`
	Pointer             string                                      `json:"pointer"`
	JSON                resourceGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ResourceGetResponseEnvelopeErrorsSource]
type resourceGetResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ResourceGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseEnvelopeMessages struct {
	Code             ResourceGetResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Meta             ResourceGetResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           ResourceGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ResourceGetResponseEnvelopeMessages]
type resourceGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseEnvelopeMessagesCode int64

const (
	ResourceGetResponseEnvelopeMessagesCode1001   ResourceGetResponseEnvelopeMessagesCode = 1001
	ResourceGetResponseEnvelopeMessagesCode1002   ResourceGetResponseEnvelopeMessagesCode = 1002
	ResourceGetResponseEnvelopeMessagesCode1003   ResourceGetResponseEnvelopeMessagesCode = 1003
	ResourceGetResponseEnvelopeMessagesCode1004   ResourceGetResponseEnvelopeMessagesCode = 1004
	ResourceGetResponseEnvelopeMessagesCode1005   ResourceGetResponseEnvelopeMessagesCode = 1005
	ResourceGetResponseEnvelopeMessagesCode1006   ResourceGetResponseEnvelopeMessagesCode = 1006
	ResourceGetResponseEnvelopeMessagesCode1007   ResourceGetResponseEnvelopeMessagesCode = 1007
	ResourceGetResponseEnvelopeMessagesCode1008   ResourceGetResponseEnvelopeMessagesCode = 1008
	ResourceGetResponseEnvelopeMessagesCode1009   ResourceGetResponseEnvelopeMessagesCode = 1009
	ResourceGetResponseEnvelopeMessagesCode1010   ResourceGetResponseEnvelopeMessagesCode = 1010
	ResourceGetResponseEnvelopeMessagesCode1011   ResourceGetResponseEnvelopeMessagesCode = 1011
	ResourceGetResponseEnvelopeMessagesCode1012   ResourceGetResponseEnvelopeMessagesCode = 1012
	ResourceGetResponseEnvelopeMessagesCode1013   ResourceGetResponseEnvelopeMessagesCode = 1013
	ResourceGetResponseEnvelopeMessagesCode1014   ResourceGetResponseEnvelopeMessagesCode = 1014
	ResourceGetResponseEnvelopeMessagesCode1015   ResourceGetResponseEnvelopeMessagesCode = 1015
	ResourceGetResponseEnvelopeMessagesCode1016   ResourceGetResponseEnvelopeMessagesCode = 1016
	ResourceGetResponseEnvelopeMessagesCode1017   ResourceGetResponseEnvelopeMessagesCode = 1017
	ResourceGetResponseEnvelopeMessagesCode2001   ResourceGetResponseEnvelopeMessagesCode = 2001
	ResourceGetResponseEnvelopeMessagesCode2002   ResourceGetResponseEnvelopeMessagesCode = 2002
	ResourceGetResponseEnvelopeMessagesCode2003   ResourceGetResponseEnvelopeMessagesCode = 2003
	ResourceGetResponseEnvelopeMessagesCode2004   ResourceGetResponseEnvelopeMessagesCode = 2004
	ResourceGetResponseEnvelopeMessagesCode2005   ResourceGetResponseEnvelopeMessagesCode = 2005
	ResourceGetResponseEnvelopeMessagesCode2006   ResourceGetResponseEnvelopeMessagesCode = 2006
	ResourceGetResponseEnvelopeMessagesCode2007   ResourceGetResponseEnvelopeMessagesCode = 2007
	ResourceGetResponseEnvelopeMessagesCode2008   ResourceGetResponseEnvelopeMessagesCode = 2008
	ResourceGetResponseEnvelopeMessagesCode2009   ResourceGetResponseEnvelopeMessagesCode = 2009
	ResourceGetResponseEnvelopeMessagesCode2010   ResourceGetResponseEnvelopeMessagesCode = 2010
	ResourceGetResponseEnvelopeMessagesCode2011   ResourceGetResponseEnvelopeMessagesCode = 2011
	ResourceGetResponseEnvelopeMessagesCode2012   ResourceGetResponseEnvelopeMessagesCode = 2012
	ResourceGetResponseEnvelopeMessagesCode2013   ResourceGetResponseEnvelopeMessagesCode = 2013
	ResourceGetResponseEnvelopeMessagesCode2014   ResourceGetResponseEnvelopeMessagesCode = 2014
	ResourceGetResponseEnvelopeMessagesCode2015   ResourceGetResponseEnvelopeMessagesCode = 2015
	ResourceGetResponseEnvelopeMessagesCode2016   ResourceGetResponseEnvelopeMessagesCode = 2016
	ResourceGetResponseEnvelopeMessagesCode2017   ResourceGetResponseEnvelopeMessagesCode = 2017
	ResourceGetResponseEnvelopeMessagesCode2018   ResourceGetResponseEnvelopeMessagesCode = 2018
	ResourceGetResponseEnvelopeMessagesCode2019   ResourceGetResponseEnvelopeMessagesCode = 2019
	ResourceGetResponseEnvelopeMessagesCode2020   ResourceGetResponseEnvelopeMessagesCode = 2020
	ResourceGetResponseEnvelopeMessagesCode2021   ResourceGetResponseEnvelopeMessagesCode = 2021
	ResourceGetResponseEnvelopeMessagesCode2022   ResourceGetResponseEnvelopeMessagesCode = 2022
	ResourceGetResponseEnvelopeMessagesCode3001   ResourceGetResponseEnvelopeMessagesCode = 3001
	ResourceGetResponseEnvelopeMessagesCode3002   ResourceGetResponseEnvelopeMessagesCode = 3002
	ResourceGetResponseEnvelopeMessagesCode3003   ResourceGetResponseEnvelopeMessagesCode = 3003
	ResourceGetResponseEnvelopeMessagesCode3004   ResourceGetResponseEnvelopeMessagesCode = 3004
	ResourceGetResponseEnvelopeMessagesCode3005   ResourceGetResponseEnvelopeMessagesCode = 3005
	ResourceGetResponseEnvelopeMessagesCode3006   ResourceGetResponseEnvelopeMessagesCode = 3006
	ResourceGetResponseEnvelopeMessagesCode3007   ResourceGetResponseEnvelopeMessagesCode = 3007
	ResourceGetResponseEnvelopeMessagesCode4001   ResourceGetResponseEnvelopeMessagesCode = 4001
	ResourceGetResponseEnvelopeMessagesCode4002   ResourceGetResponseEnvelopeMessagesCode = 4002
	ResourceGetResponseEnvelopeMessagesCode4003   ResourceGetResponseEnvelopeMessagesCode = 4003
	ResourceGetResponseEnvelopeMessagesCode4004   ResourceGetResponseEnvelopeMessagesCode = 4004
	ResourceGetResponseEnvelopeMessagesCode4005   ResourceGetResponseEnvelopeMessagesCode = 4005
	ResourceGetResponseEnvelopeMessagesCode4006   ResourceGetResponseEnvelopeMessagesCode = 4006
	ResourceGetResponseEnvelopeMessagesCode4007   ResourceGetResponseEnvelopeMessagesCode = 4007
	ResourceGetResponseEnvelopeMessagesCode4008   ResourceGetResponseEnvelopeMessagesCode = 4008
	ResourceGetResponseEnvelopeMessagesCode4009   ResourceGetResponseEnvelopeMessagesCode = 4009
	ResourceGetResponseEnvelopeMessagesCode4010   ResourceGetResponseEnvelopeMessagesCode = 4010
	ResourceGetResponseEnvelopeMessagesCode4011   ResourceGetResponseEnvelopeMessagesCode = 4011
	ResourceGetResponseEnvelopeMessagesCode4012   ResourceGetResponseEnvelopeMessagesCode = 4012
	ResourceGetResponseEnvelopeMessagesCode4013   ResourceGetResponseEnvelopeMessagesCode = 4013
	ResourceGetResponseEnvelopeMessagesCode4014   ResourceGetResponseEnvelopeMessagesCode = 4014
	ResourceGetResponseEnvelopeMessagesCode4015   ResourceGetResponseEnvelopeMessagesCode = 4015
	ResourceGetResponseEnvelopeMessagesCode4016   ResourceGetResponseEnvelopeMessagesCode = 4016
	ResourceGetResponseEnvelopeMessagesCode4017   ResourceGetResponseEnvelopeMessagesCode = 4017
	ResourceGetResponseEnvelopeMessagesCode4018   ResourceGetResponseEnvelopeMessagesCode = 4018
	ResourceGetResponseEnvelopeMessagesCode4019   ResourceGetResponseEnvelopeMessagesCode = 4019
	ResourceGetResponseEnvelopeMessagesCode4020   ResourceGetResponseEnvelopeMessagesCode = 4020
	ResourceGetResponseEnvelopeMessagesCode4021   ResourceGetResponseEnvelopeMessagesCode = 4021
	ResourceGetResponseEnvelopeMessagesCode4022   ResourceGetResponseEnvelopeMessagesCode = 4022
	ResourceGetResponseEnvelopeMessagesCode4023   ResourceGetResponseEnvelopeMessagesCode = 4023
	ResourceGetResponseEnvelopeMessagesCode5001   ResourceGetResponseEnvelopeMessagesCode = 5001
	ResourceGetResponseEnvelopeMessagesCode5002   ResourceGetResponseEnvelopeMessagesCode = 5002
	ResourceGetResponseEnvelopeMessagesCode5003   ResourceGetResponseEnvelopeMessagesCode = 5003
	ResourceGetResponseEnvelopeMessagesCode5004   ResourceGetResponseEnvelopeMessagesCode = 5004
	ResourceGetResponseEnvelopeMessagesCode102000 ResourceGetResponseEnvelopeMessagesCode = 102000
	ResourceGetResponseEnvelopeMessagesCode102001 ResourceGetResponseEnvelopeMessagesCode = 102001
	ResourceGetResponseEnvelopeMessagesCode102002 ResourceGetResponseEnvelopeMessagesCode = 102002
	ResourceGetResponseEnvelopeMessagesCode102003 ResourceGetResponseEnvelopeMessagesCode = 102003
	ResourceGetResponseEnvelopeMessagesCode102004 ResourceGetResponseEnvelopeMessagesCode = 102004
	ResourceGetResponseEnvelopeMessagesCode102005 ResourceGetResponseEnvelopeMessagesCode = 102005
	ResourceGetResponseEnvelopeMessagesCode102006 ResourceGetResponseEnvelopeMessagesCode = 102006
	ResourceGetResponseEnvelopeMessagesCode102007 ResourceGetResponseEnvelopeMessagesCode = 102007
	ResourceGetResponseEnvelopeMessagesCode102008 ResourceGetResponseEnvelopeMessagesCode = 102008
	ResourceGetResponseEnvelopeMessagesCode102009 ResourceGetResponseEnvelopeMessagesCode = 102009
	ResourceGetResponseEnvelopeMessagesCode102010 ResourceGetResponseEnvelopeMessagesCode = 102010
	ResourceGetResponseEnvelopeMessagesCode102011 ResourceGetResponseEnvelopeMessagesCode = 102011
	ResourceGetResponseEnvelopeMessagesCode102012 ResourceGetResponseEnvelopeMessagesCode = 102012
	ResourceGetResponseEnvelopeMessagesCode102013 ResourceGetResponseEnvelopeMessagesCode = 102013
	ResourceGetResponseEnvelopeMessagesCode102014 ResourceGetResponseEnvelopeMessagesCode = 102014
	ResourceGetResponseEnvelopeMessagesCode102015 ResourceGetResponseEnvelopeMessagesCode = 102015
	ResourceGetResponseEnvelopeMessagesCode102016 ResourceGetResponseEnvelopeMessagesCode = 102016
	ResourceGetResponseEnvelopeMessagesCode102017 ResourceGetResponseEnvelopeMessagesCode = 102017
	ResourceGetResponseEnvelopeMessagesCode102018 ResourceGetResponseEnvelopeMessagesCode = 102018
	ResourceGetResponseEnvelopeMessagesCode102019 ResourceGetResponseEnvelopeMessagesCode = 102019
	ResourceGetResponseEnvelopeMessagesCode102020 ResourceGetResponseEnvelopeMessagesCode = 102020
	ResourceGetResponseEnvelopeMessagesCode102021 ResourceGetResponseEnvelopeMessagesCode = 102021
	ResourceGetResponseEnvelopeMessagesCode102022 ResourceGetResponseEnvelopeMessagesCode = 102022
	ResourceGetResponseEnvelopeMessagesCode102023 ResourceGetResponseEnvelopeMessagesCode = 102023
	ResourceGetResponseEnvelopeMessagesCode102024 ResourceGetResponseEnvelopeMessagesCode = 102024
	ResourceGetResponseEnvelopeMessagesCode102025 ResourceGetResponseEnvelopeMessagesCode = 102025
	ResourceGetResponseEnvelopeMessagesCode102026 ResourceGetResponseEnvelopeMessagesCode = 102026
	ResourceGetResponseEnvelopeMessagesCode102027 ResourceGetResponseEnvelopeMessagesCode = 102027
	ResourceGetResponseEnvelopeMessagesCode102028 ResourceGetResponseEnvelopeMessagesCode = 102028
	ResourceGetResponseEnvelopeMessagesCode102029 ResourceGetResponseEnvelopeMessagesCode = 102029
	ResourceGetResponseEnvelopeMessagesCode102030 ResourceGetResponseEnvelopeMessagesCode = 102030
	ResourceGetResponseEnvelopeMessagesCode102031 ResourceGetResponseEnvelopeMessagesCode = 102031
	ResourceGetResponseEnvelopeMessagesCode102032 ResourceGetResponseEnvelopeMessagesCode = 102032
	ResourceGetResponseEnvelopeMessagesCode102033 ResourceGetResponseEnvelopeMessagesCode = 102033
	ResourceGetResponseEnvelopeMessagesCode102034 ResourceGetResponseEnvelopeMessagesCode = 102034
	ResourceGetResponseEnvelopeMessagesCode102035 ResourceGetResponseEnvelopeMessagesCode = 102035
	ResourceGetResponseEnvelopeMessagesCode102036 ResourceGetResponseEnvelopeMessagesCode = 102036
	ResourceGetResponseEnvelopeMessagesCode102037 ResourceGetResponseEnvelopeMessagesCode = 102037
	ResourceGetResponseEnvelopeMessagesCode102038 ResourceGetResponseEnvelopeMessagesCode = 102038
	ResourceGetResponseEnvelopeMessagesCode102039 ResourceGetResponseEnvelopeMessagesCode = 102039
	ResourceGetResponseEnvelopeMessagesCode102040 ResourceGetResponseEnvelopeMessagesCode = 102040
	ResourceGetResponseEnvelopeMessagesCode102041 ResourceGetResponseEnvelopeMessagesCode = 102041
	ResourceGetResponseEnvelopeMessagesCode102042 ResourceGetResponseEnvelopeMessagesCode = 102042
	ResourceGetResponseEnvelopeMessagesCode102043 ResourceGetResponseEnvelopeMessagesCode = 102043
	ResourceGetResponseEnvelopeMessagesCode102044 ResourceGetResponseEnvelopeMessagesCode = 102044
	ResourceGetResponseEnvelopeMessagesCode102045 ResourceGetResponseEnvelopeMessagesCode = 102045
	ResourceGetResponseEnvelopeMessagesCode102046 ResourceGetResponseEnvelopeMessagesCode = 102046
	ResourceGetResponseEnvelopeMessagesCode102047 ResourceGetResponseEnvelopeMessagesCode = 102047
	ResourceGetResponseEnvelopeMessagesCode102048 ResourceGetResponseEnvelopeMessagesCode = 102048
	ResourceGetResponseEnvelopeMessagesCode102049 ResourceGetResponseEnvelopeMessagesCode = 102049
	ResourceGetResponseEnvelopeMessagesCode102050 ResourceGetResponseEnvelopeMessagesCode = 102050
	ResourceGetResponseEnvelopeMessagesCode102051 ResourceGetResponseEnvelopeMessagesCode = 102051
	ResourceGetResponseEnvelopeMessagesCode102052 ResourceGetResponseEnvelopeMessagesCode = 102052
	ResourceGetResponseEnvelopeMessagesCode102053 ResourceGetResponseEnvelopeMessagesCode = 102053
	ResourceGetResponseEnvelopeMessagesCode102054 ResourceGetResponseEnvelopeMessagesCode = 102054
	ResourceGetResponseEnvelopeMessagesCode102055 ResourceGetResponseEnvelopeMessagesCode = 102055
	ResourceGetResponseEnvelopeMessagesCode102056 ResourceGetResponseEnvelopeMessagesCode = 102056
	ResourceGetResponseEnvelopeMessagesCode102057 ResourceGetResponseEnvelopeMessagesCode = 102057
	ResourceGetResponseEnvelopeMessagesCode102058 ResourceGetResponseEnvelopeMessagesCode = 102058
	ResourceGetResponseEnvelopeMessagesCode102059 ResourceGetResponseEnvelopeMessagesCode = 102059
	ResourceGetResponseEnvelopeMessagesCode102060 ResourceGetResponseEnvelopeMessagesCode = 102060
	ResourceGetResponseEnvelopeMessagesCode102061 ResourceGetResponseEnvelopeMessagesCode = 102061
	ResourceGetResponseEnvelopeMessagesCode102062 ResourceGetResponseEnvelopeMessagesCode = 102062
	ResourceGetResponseEnvelopeMessagesCode102063 ResourceGetResponseEnvelopeMessagesCode = 102063
	ResourceGetResponseEnvelopeMessagesCode102064 ResourceGetResponseEnvelopeMessagesCode = 102064
	ResourceGetResponseEnvelopeMessagesCode102065 ResourceGetResponseEnvelopeMessagesCode = 102065
	ResourceGetResponseEnvelopeMessagesCode102066 ResourceGetResponseEnvelopeMessagesCode = 102066
	ResourceGetResponseEnvelopeMessagesCode103001 ResourceGetResponseEnvelopeMessagesCode = 103001
	ResourceGetResponseEnvelopeMessagesCode103002 ResourceGetResponseEnvelopeMessagesCode = 103002
	ResourceGetResponseEnvelopeMessagesCode103003 ResourceGetResponseEnvelopeMessagesCode = 103003
	ResourceGetResponseEnvelopeMessagesCode103004 ResourceGetResponseEnvelopeMessagesCode = 103004
	ResourceGetResponseEnvelopeMessagesCode103005 ResourceGetResponseEnvelopeMessagesCode = 103005
	ResourceGetResponseEnvelopeMessagesCode103006 ResourceGetResponseEnvelopeMessagesCode = 103006
	ResourceGetResponseEnvelopeMessagesCode103007 ResourceGetResponseEnvelopeMessagesCode = 103007
	ResourceGetResponseEnvelopeMessagesCode103008 ResourceGetResponseEnvelopeMessagesCode = 103008
)

func (r ResourceGetResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case ResourceGetResponseEnvelopeMessagesCode1001, ResourceGetResponseEnvelopeMessagesCode1002, ResourceGetResponseEnvelopeMessagesCode1003, ResourceGetResponseEnvelopeMessagesCode1004, ResourceGetResponseEnvelopeMessagesCode1005, ResourceGetResponseEnvelopeMessagesCode1006, ResourceGetResponseEnvelopeMessagesCode1007, ResourceGetResponseEnvelopeMessagesCode1008, ResourceGetResponseEnvelopeMessagesCode1009, ResourceGetResponseEnvelopeMessagesCode1010, ResourceGetResponseEnvelopeMessagesCode1011, ResourceGetResponseEnvelopeMessagesCode1012, ResourceGetResponseEnvelopeMessagesCode1013, ResourceGetResponseEnvelopeMessagesCode1014, ResourceGetResponseEnvelopeMessagesCode1015, ResourceGetResponseEnvelopeMessagesCode1016, ResourceGetResponseEnvelopeMessagesCode1017, ResourceGetResponseEnvelopeMessagesCode2001, ResourceGetResponseEnvelopeMessagesCode2002, ResourceGetResponseEnvelopeMessagesCode2003, ResourceGetResponseEnvelopeMessagesCode2004, ResourceGetResponseEnvelopeMessagesCode2005, ResourceGetResponseEnvelopeMessagesCode2006, ResourceGetResponseEnvelopeMessagesCode2007, ResourceGetResponseEnvelopeMessagesCode2008, ResourceGetResponseEnvelopeMessagesCode2009, ResourceGetResponseEnvelopeMessagesCode2010, ResourceGetResponseEnvelopeMessagesCode2011, ResourceGetResponseEnvelopeMessagesCode2012, ResourceGetResponseEnvelopeMessagesCode2013, ResourceGetResponseEnvelopeMessagesCode2014, ResourceGetResponseEnvelopeMessagesCode2015, ResourceGetResponseEnvelopeMessagesCode2016, ResourceGetResponseEnvelopeMessagesCode2017, ResourceGetResponseEnvelopeMessagesCode2018, ResourceGetResponseEnvelopeMessagesCode2019, ResourceGetResponseEnvelopeMessagesCode2020, ResourceGetResponseEnvelopeMessagesCode2021, ResourceGetResponseEnvelopeMessagesCode2022, ResourceGetResponseEnvelopeMessagesCode3001, ResourceGetResponseEnvelopeMessagesCode3002, ResourceGetResponseEnvelopeMessagesCode3003, ResourceGetResponseEnvelopeMessagesCode3004, ResourceGetResponseEnvelopeMessagesCode3005, ResourceGetResponseEnvelopeMessagesCode3006, ResourceGetResponseEnvelopeMessagesCode3007, ResourceGetResponseEnvelopeMessagesCode4001, ResourceGetResponseEnvelopeMessagesCode4002, ResourceGetResponseEnvelopeMessagesCode4003, ResourceGetResponseEnvelopeMessagesCode4004, ResourceGetResponseEnvelopeMessagesCode4005, ResourceGetResponseEnvelopeMessagesCode4006, ResourceGetResponseEnvelopeMessagesCode4007, ResourceGetResponseEnvelopeMessagesCode4008, ResourceGetResponseEnvelopeMessagesCode4009, ResourceGetResponseEnvelopeMessagesCode4010, ResourceGetResponseEnvelopeMessagesCode4011, ResourceGetResponseEnvelopeMessagesCode4012, ResourceGetResponseEnvelopeMessagesCode4013, ResourceGetResponseEnvelopeMessagesCode4014, ResourceGetResponseEnvelopeMessagesCode4015, ResourceGetResponseEnvelopeMessagesCode4016, ResourceGetResponseEnvelopeMessagesCode4017, ResourceGetResponseEnvelopeMessagesCode4018, ResourceGetResponseEnvelopeMessagesCode4019, ResourceGetResponseEnvelopeMessagesCode4020, ResourceGetResponseEnvelopeMessagesCode4021, ResourceGetResponseEnvelopeMessagesCode4022, ResourceGetResponseEnvelopeMessagesCode4023, ResourceGetResponseEnvelopeMessagesCode5001, ResourceGetResponseEnvelopeMessagesCode5002, ResourceGetResponseEnvelopeMessagesCode5003, ResourceGetResponseEnvelopeMessagesCode5004, ResourceGetResponseEnvelopeMessagesCode102000, ResourceGetResponseEnvelopeMessagesCode102001, ResourceGetResponseEnvelopeMessagesCode102002, ResourceGetResponseEnvelopeMessagesCode102003, ResourceGetResponseEnvelopeMessagesCode102004, ResourceGetResponseEnvelopeMessagesCode102005, ResourceGetResponseEnvelopeMessagesCode102006, ResourceGetResponseEnvelopeMessagesCode102007, ResourceGetResponseEnvelopeMessagesCode102008, ResourceGetResponseEnvelopeMessagesCode102009, ResourceGetResponseEnvelopeMessagesCode102010, ResourceGetResponseEnvelopeMessagesCode102011, ResourceGetResponseEnvelopeMessagesCode102012, ResourceGetResponseEnvelopeMessagesCode102013, ResourceGetResponseEnvelopeMessagesCode102014, ResourceGetResponseEnvelopeMessagesCode102015, ResourceGetResponseEnvelopeMessagesCode102016, ResourceGetResponseEnvelopeMessagesCode102017, ResourceGetResponseEnvelopeMessagesCode102018, ResourceGetResponseEnvelopeMessagesCode102019, ResourceGetResponseEnvelopeMessagesCode102020, ResourceGetResponseEnvelopeMessagesCode102021, ResourceGetResponseEnvelopeMessagesCode102022, ResourceGetResponseEnvelopeMessagesCode102023, ResourceGetResponseEnvelopeMessagesCode102024, ResourceGetResponseEnvelopeMessagesCode102025, ResourceGetResponseEnvelopeMessagesCode102026, ResourceGetResponseEnvelopeMessagesCode102027, ResourceGetResponseEnvelopeMessagesCode102028, ResourceGetResponseEnvelopeMessagesCode102029, ResourceGetResponseEnvelopeMessagesCode102030, ResourceGetResponseEnvelopeMessagesCode102031, ResourceGetResponseEnvelopeMessagesCode102032, ResourceGetResponseEnvelopeMessagesCode102033, ResourceGetResponseEnvelopeMessagesCode102034, ResourceGetResponseEnvelopeMessagesCode102035, ResourceGetResponseEnvelopeMessagesCode102036, ResourceGetResponseEnvelopeMessagesCode102037, ResourceGetResponseEnvelopeMessagesCode102038, ResourceGetResponseEnvelopeMessagesCode102039, ResourceGetResponseEnvelopeMessagesCode102040, ResourceGetResponseEnvelopeMessagesCode102041, ResourceGetResponseEnvelopeMessagesCode102042, ResourceGetResponseEnvelopeMessagesCode102043, ResourceGetResponseEnvelopeMessagesCode102044, ResourceGetResponseEnvelopeMessagesCode102045, ResourceGetResponseEnvelopeMessagesCode102046, ResourceGetResponseEnvelopeMessagesCode102047, ResourceGetResponseEnvelopeMessagesCode102048, ResourceGetResponseEnvelopeMessagesCode102049, ResourceGetResponseEnvelopeMessagesCode102050, ResourceGetResponseEnvelopeMessagesCode102051, ResourceGetResponseEnvelopeMessagesCode102052, ResourceGetResponseEnvelopeMessagesCode102053, ResourceGetResponseEnvelopeMessagesCode102054, ResourceGetResponseEnvelopeMessagesCode102055, ResourceGetResponseEnvelopeMessagesCode102056, ResourceGetResponseEnvelopeMessagesCode102057, ResourceGetResponseEnvelopeMessagesCode102058, ResourceGetResponseEnvelopeMessagesCode102059, ResourceGetResponseEnvelopeMessagesCode102060, ResourceGetResponseEnvelopeMessagesCode102061, ResourceGetResponseEnvelopeMessagesCode102062, ResourceGetResponseEnvelopeMessagesCode102063, ResourceGetResponseEnvelopeMessagesCode102064, ResourceGetResponseEnvelopeMessagesCode102065, ResourceGetResponseEnvelopeMessagesCode102066, ResourceGetResponseEnvelopeMessagesCode103001, ResourceGetResponseEnvelopeMessagesCode103002, ResourceGetResponseEnvelopeMessagesCode103003, ResourceGetResponseEnvelopeMessagesCode103004, ResourceGetResponseEnvelopeMessagesCode103005, ResourceGetResponseEnvelopeMessagesCode103006, ResourceGetResponseEnvelopeMessagesCode103007, ResourceGetResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type ResourceGetResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                      `json:"l10n_key"`
	LoggableError string                                      `json:"loggable_error"`
	TemplateData  interface{}                                 `json:"template_data"`
	TraceID       string                                      `json:"trace_id"`
	JSON          resourceGetResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// resourceGetResponseEnvelopeMessagesMetaJSON contains the JSON metadata for the
// struct [ResourceGetResponseEnvelopeMessagesMeta]
type resourceGetResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourceGetResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type ResourceGetResponseEnvelopeMessagesSource struct {
	Parameter           string                                        `json:"parameter"`
	ParameterValueIndex int64                                         `json:"parameter_value_index"`
	Pointer             string                                        `json:"pointer"`
	JSON                resourceGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [ResourceGetResponseEnvelopeMessagesSource]
type resourceGetResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ResourceGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ResourcePolicyPreviewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Policy    param.Field[string] `json:"policy,required"`
}

func (r ResourcePolicyPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourcePolicyPreviewResponseEnvelope struct {
	Errors   []ResourcePolicyPreviewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ResourcePolicyPreviewResponseEnvelopeMessages `json:"messages,required"`
	Result   string                                          `json:"result,required"`
	Success  bool                                            `json:"success,required"`
	JSON     resourcePolicyPreviewResponseEnvelopeJSON       `json:"-"`
}

// resourcePolicyPreviewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ResourcePolicyPreviewResponseEnvelope]
type resourcePolicyPreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourcePolicyPreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourcePolicyPreviewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourcePolicyPreviewResponseEnvelopeErrors struct {
	Code             ResourcePolicyPreviewResponseEnvelopeErrorsCode   `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Meta             ResourcePolicyPreviewResponseEnvelopeErrorsMeta   `json:"meta"`
	Source           ResourcePolicyPreviewResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourcePolicyPreviewResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourcePolicyPreviewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ResourcePolicyPreviewResponseEnvelopeErrors]
type resourcePolicyPreviewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourcePolicyPreviewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourcePolicyPreviewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourcePolicyPreviewResponseEnvelopeErrorsCode int64

const (
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1001   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1001
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1002   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1002
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1003   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1003
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1004   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1004
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1005   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1005
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1006   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1006
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1007   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1007
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1008   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1008
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1009   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1009
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1010   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1010
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1011   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1011
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1012   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1012
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1013   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1013
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1014   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1014
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1015   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1015
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1016   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1016
	ResourcePolicyPreviewResponseEnvelopeErrorsCode1017   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 1017
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2001   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2001
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2002   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2002
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2003   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2003
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2004   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2004
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2005   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2005
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2006   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2006
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2007   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2007
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2008   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2008
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2009   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2009
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2010   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2010
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2011   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2011
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2012   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2012
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2013   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2013
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2014   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2014
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2015   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2015
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2016   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2016
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2017   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2017
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2018   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2018
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2019   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2019
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2020   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2020
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2021   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2021
	ResourcePolicyPreviewResponseEnvelopeErrorsCode2022   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 2022
	ResourcePolicyPreviewResponseEnvelopeErrorsCode3001   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 3001
	ResourcePolicyPreviewResponseEnvelopeErrorsCode3002   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 3002
	ResourcePolicyPreviewResponseEnvelopeErrorsCode3003   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 3003
	ResourcePolicyPreviewResponseEnvelopeErrorsCode3004   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 3004
	ResourcePolicyPreviewResponseEnvelopeErrorsCode3005   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 3005
	ResourcePolicyPreviewResponseEnvelopeErrorsCode3006   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 3006
	ResourcePolicyPreviewResponseEnvelopeErrorsCode3007   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 3007
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4001   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4001
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4002   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4002
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4003   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4003
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4004   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4004
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4005   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4005
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4006   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4006
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4007   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4007
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4008   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4008
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4009   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4009
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4010   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4010
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4011   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4011
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4012   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4012
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4013   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4013
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4014   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4014
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4015   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4015
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4016   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4016
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4017   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4017
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4018   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4018
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4019   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4019
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4020   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4020
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4021   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4021
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4022   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4022
	ResourcePolicyPreviewResponseEnvelopeErrorsCode4023   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 4023
	ResourcePolicyPreviewResponseEnvelopeErrorsCode5001   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 5001
	ResourcePolicyPreviewResponseEnvelopeErrorsCode5002   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 5002
	ResourcePolicyPreviewResponseEnvelopeErrorsCode5003   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 5003
	ResourcePolicyPreviewResponseEnvelopeErrorsCode5004   ResourcePolicyPreviewResponseEnvelopeErrorsCode = 5004
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102000 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102000
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102001 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102001
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102002 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102002
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102003 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102003
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102004 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102004
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102005 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102005
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102006 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102006
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102007 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102007
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102008 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102008
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102009 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102009
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102010 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102010
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102011 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102011
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102012 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102012
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102013 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102013
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102014 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102014
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102015 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102015
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102016 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102016
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102017 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102017
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102018 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102018
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102019 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102019
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102020 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102020
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102021 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102021
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102022 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102022
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102023 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102023
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102024 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102024
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102025 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102025
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102026 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102026
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102027 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102027
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102028 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102028
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102029 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102029
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102030 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102030
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102031 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102031
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102032 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102032
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102033 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102033
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102034 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102034
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102035 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102035
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102036 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102036
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102037 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102037
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102038 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102038
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102039 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102039
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102040 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102040
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102041 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102041
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102042 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102042
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102043 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102043
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102044 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102044
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102045 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102045
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102046 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102046
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102047 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102047
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102048 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102048
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102049 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102049
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102050 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102050
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102051 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102051
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102052 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102052
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102053 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102053
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102054 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102054
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102055 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102055
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102056 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102056
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102057 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102057
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102058 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102058
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102059 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102059
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102060 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102060
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102061 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102061
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102062 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102062
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102063 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102063
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102064 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102064
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102065 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102065
	ResourcePolicyPreviewResponseEnvelopeErrorsCode102066 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 102066
	ResourcePolicyPreviewResponseEnvelopeErrorsCode103001 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 103001
	ResourcePolicyPreviewResponseEnvelopeErrorsCode103002 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 103002
	ResourcePolicyPreviewResponseEnvelopeErrorsCode103003 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 103003
	ResourcePolicyPreviewResponseEnvelopeErrorsCode103004 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 103004
	ResourcePolicyPreviewResponseEnvelopeErrorsCode103005 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 103005
	ResourcePolicyPreviewResponseEnvelopeErrorsCode103006 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 103006
	ResourcePolicyPreviewResponseEnvelopeErrorsCode103007 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 103007
	ResourcePolicyPreviewResponseEnvelopeErrorsCode103008 ResourcePolicyPreviewResponseEnvelopeErrorsCode = 103008
)

func (r ResourcePolicyPreviewResponseEnvelopeErrorsCode) IsKnown() bool {
	switch r {
	case ResourcePolicyPreviewResponseEnvelopeErrorsCode1001, ResourcePolicyPreviewResponseEnvelopeErrorsCode1002, ResourcePolicyPreviewResponseEnvelopeErrorsCode1003, ResourcePolicyPreviewResponseEnvelopeErrorsCode1004, ResourcePolicyPreviewResponseEnvelopeErrorsCode1005, ResourcePolicyPreviewResponseEnvelopeErrorsCode1006, ResourcePolicyPreviewResponseEnvelopeErrorsCode1007, ResourcePolicyPreviewResponseEnvelopeErrorsCode1008, ResourcePolicyPreviewResponseEnvelopeErrorsCode1009, ResourcePolicyPreviewResponseEnvelopeErrorsCode1010, ResourcePolicyPreviewResponseEnvelopeErrorsCode1011, ResourcePolicyPreviewResponseEnvelopeErrorsCode1012, ResourcePolicyPreviewResponseEnvelopeErrorsCode1013, ResourcePolicyPreviewResponseEnvelopeErrorsCode1014, ResourcePolicyPreviewResponseEnvelopeErrorsCode1015, ResourcePolicyPreviewResponseEnvelopeErrorsCode1016, ResourcePolicyPreviewResponseEnvelopeErrorsCode1017, ResourcePolicyPreviewResponseEnvelopeErrorsCode2001, ResourcePolicyPreviewResponseEnvelopeErrorsCode2002, ResourcePolicyPreviewResponseEnvelopeErrorsCode2003, ResourcePolicyPreviewResponseEnvelopeErrorsCode2004, ResourcePolicyPreviewResponseEnvelopeErrorsCode2005, ResourcePolicyPreviewResponseEnvelopeErrorsCode2006, ResourcePolicyPreviewResponseEnvelopeErrorsCode2007, ResourcePolicyPreviewResponseEnvelopeErrorsCode2008, ResourcePolicyPreviewResponseEnvelopeErrorsCode2009, ResourcePolicyPreviewResponseEnvelopeErrorsCode2010, ResourcePolicyPreviewResponseEnvelopeErrorsCode2011, ResourcePolicyPreviewResponseEnvelopeErrorsCode2012, ResourcePolicyPreviewResponseEnvelopeErrorsCode2013, ResourcePolicyPreviewResponseEnvelopeErrorsCode2014, ResourcePolicyPreviewResponseEnvelopeErrorsCode2015, ResourcePolicyPreviewResponseEnvelopeErrorsCode2016, ResourcePolicyPreviewResponseEnvelopeErrorsCode2017, ResourcePolicyPreviewResponseEnvelopeErrorsCode2018, ResourcePolicyPreviewResponseEnvelopeErrorsCode2019, ResourcePolicyPreviewResponseEnvelopeErrorsCode2020, ResourcePolicyPreviewResponseEnvelopeErrorsCode2021, ResourcePolicyPreviewResponseEnvelopeErrorsCode2022, ResourcePolicyPreviewResponseEnvelopeErrorsCode3001, ResourcePolicyPreviewResponseEnvelopeErrorsCode3002, ResourcePolicyPreviewResponseEnvelopeErrorsCode3003, ResourcePolicyPreviewResponseEnvelopeErrorsCode3004, ResourcePolicyPreviewResponseEnvelopeErrorsCode3005, ResourcePolicyPreviewResponseEnvelopeErrorsCode3006, ResourcePolicyPreviewResponseEnvelopeErrorsCode3007, ResourcePolicyPreviewResponseEnvelopeErrorsCode4001, ResourcePolicyPreviewResponseEnvelopeErrorsCode4002, ResourcePolicyPreviewResponseEnvelopeErrorsCode4003, ResourcePolicyPreviewResponseEnvelopeErrorsCode4004, ResourcePolicyPreviewResponseEnvelopeErrorsCode4005, ResourcePolicyPreviewResponseEnvelopeErrorsCode4006, ResourcePolicyPreviewResponseEnvelopeErrorsCode4007, ResourcePolicyPreviewResponseEnvelopeErrorsCode4008, ResourcePolicyPreviewResponseEnvelopeErrorsCode4009, ResourcePolicyPreviewResponseEnvelopeErrorsCode4010, ResourcePolicyPreviewResponseEnvelopeErrorsCode4011, ResourcePolicyPreviewResponseEnvelopeErrorsCode4012, ResourcePolicyPreviewResponseEnvelopeErrorsCode4013, ResourcePolicyPreviewResponseEnvelopeErrorsCode4014, ResourcePolicyPreviewResponseEnvelopeErrorsCode4015, ResourcePolicyPreviewResponseEnvelopeErrorsCode4016, ResourcePolicyPreviewResponseEnvelopeErrorsCode4017, ResourcePolicyPreviewResponseEnvelopeErrorsCode4018, ResourcePolicyPreviewResponseEnvelopeErrorsCode4019, ResourcePolicyPreviewResponseEnvelopeErrorsCode4020, ResourcePolicyPreviewResponseEnvelopeErrorsCode4021, ResourcePolicyPreviewResponseEnvelopeErrorsCode4022, ResourcePolicyPreviewResponseEnvelopeErrorsCode4023, ResourcePolicyPreviewResponseEnvelopeErrorsCode5001, ResourcePolicyPreviewResponseEnvelopeErrorsCode5002, ResourcePolicyPreviewResponseEnvelopeErrorsCode5003, ResourcePolicyPreviewResponseEnvelopeErrorsCode5004, ResourcePolicyPreviewResponseEnvelopeErrorsCode102000, ResourcePolicyPreviewResponseEnvelopeErrorsCode102001, ResourcePolicyPreviewResponseEnvelopeErrorsCode102002, ResourcePolicyPreviewResponseEnvelopeErrorsCode102003, ResourcePolicyPreviewResponseEnvelopeErrorsCode102004, ResourcePolicyPreviewResponseEnvelopeErrorsCode102005, ResourcePolicyPreviewResponseEnvelopeErrorsCode102006, ResourcePolicyPreviewResponseEnvelopeErrorsCode102007, ResourcePolicyPreviewResponseEnvelopeErrorsCode102008, ResourcePolicyPreviewResponseEnvelopeErrorsCode102009, ResourcePolicyPreviewResponseEnvelopeErrorsCode102010, ResourcePolicyPreviewResponseEnvelopeErrorsCode102011, ResourcePolicyPreviewResponseEnvelopeErrorsCode102012, ResourcePolicyPreviewResponseEnvelopeErrorsCode102013, ResourcePolicyPreviewResponseEnvelopeErrorsCode102014, ResourcePolicyPreviewResponseEnvelopeErrorsCode102015, ResourcePolicyPreviewResponseEnvelopeErrorsCode102016, ResourcePolicyPreviewResponseEnvelopeErrorsCode102017, ResourcePolicyPreviewResponseEnvelopeErrorsCode102018, ResourcePolicyPreviewResponseEnvelopeErrorsCode102019, ResourcePolicyPreviewResponseEnvelopeErrorsCode102020, ResourcePolicyPreviewResponseEnvelopeErrorsCode102021, ResourcePolicyPreviewResponseEnvelopeErrorsCode102022, ResourcePolicyPreviewResponseEnvelopeErrorsCode102023, ResourcePolicyPreviewResponseEnvelopeErrorsCode102024, ResourcePolicyPreviewResponseEnvelopeErrorsCode102025, ResourcePolicyPreviewResponseEnvelopeErrorsCode102026, ResourcePolicyPreviewResponseEnvelopeErrorsCode102027, ResourcePolicyPreviewResponseEnvelopeErrorsCode102028, ResourcePolicyPreviewResponseEnvelopeErrorsCode102029, ResourcePolicyPreviewResponseEnvelopeErrorsCode102030, ResourcePolicyPreviewResponseEnvelopeErrorsCode102031, ResourcePolicyPreviewResponseEnvelopeErrorsCode102032, ResourcePolicyPreviewResponseEnvelopeErrorsCode102033, ResourcePolicyPreviewResponseEnvelopeErrorsCode102034, ResourcePolicyPreviewResponseEnvelopeErrorsCode102035, ResourcePolicyPreviewResponseEnvelopeErrorsCode102036, ResourcePolicyPreviewResponseEnvelopeErrorsCode102037, ResourcePolicyPreviewResponseEnvelopeErrorsCode102038, ResourcePolicyPreviewResponseEnvelopeErrorsCode102039, ResourcePolicyPreviewResponseEnvelopeErrorsCode102040, ResourcePolicyPreviewResponseEnvelopeErrorsCode102041, ResourcePolicyPreviewResponseEnvelopeErrorsCode102042, ResourcePolicyPreviewResponseEnvelopeErrorsCode102043, ResourcePolicyPreviewResponseEnvelopeErrorsCode102044, ResourcePolicyPreviewResponseEnvelopeErrorsCode102045, ResourcePolicyPreviewResponseEnvelopeErrorsCode102046, ResourcePolicyPreviewResponseEnvelopeErrorsCode102047, ResourcePolicyPreviewResponseEnvelopeErrorsCode102048, ResourcePolicyPreviewResponseEnvelopeErrorsCode102049, ResourcePolicyPreviewResponseEnvelopeErrorsCode102050, ResourcePolicyPreviewResponseEnvelopeErrorsCode102051, ResourcePolicyPreviewResponseEnvelopeErrorsCode102052, ResourcePolicyPreviewResponseEnvelopeErrorsCode102053, ResourcePolicyPreviewResponseEnvelopeErrorsCode102054, ResourcePolicyPreviewResponseEnvelopeErrorsCode102055, ResourcePolicyPreviewResponseEnvelopeErrorsCode102056, ResourcePolicyPreviewResponseEnvelopeErrorsCode102057, ResourcePolicyPreviewResponseEnvelopeErrorsCode102058, ResourcePolicyPreviewResponseEnvelopeErrorsCode102059, ResourcePolicyPreviewResponseEnvelopeErrorsCode102060, ResourcePolicyPreviewResponseEnvelopeErrorsCode102061, ResourcePolicyPreviewResponseEnvelopeErrorsCode102062, ResourcePolicyPreviewResponseEnvelopeErrorsCode102063, ResourcePolicyPreviewResponseEnvelopeErrorsCode102064, ResourcePolicyPreviewResponseEnvelopeErrorsCode102065, ResourcePolicyPreviewResponseEnvelopeErrorsCode102066, ResourcePolicyPreviewResponseEnvelopeErrorsCode103001, ResourcePolicyPreviewResponseEnvelopeErrorsCode103002, ResourcePolicyPreviewResponseEnvelopeErrorsCode103003, ResourcePolicyPreviewResponseEnvelopeErrorsCode103004, ResourcePolicyPreviewResponseEnvelopeErrorsCode103005, ResourcePolicyPreviewResponseEnvelopeErrorsCode103006, ResourcePolicyPreviewResponseEnvelopeErrorsCode103007, ResourcePolicyPreviewResponseEnvelopeErrorsCode103008:
		return true
	}
	return false
}

type ResourcePolicyPreviewResponseEnvelopeErrorsMeta struct {
	L10nKey       string                                              `json:"l10n_key"`
	LoggableError string                                              `json:"loggable_error"`
	TemplateData  interface{}                                         `json:"template_data"`
	TraceID       string                                              `json:"trace_id"`
	JSON          resourcePolicyPreviewResponseEnvelopeErrorsMetaJSON `json:"-"`
}

// resourcePolicyPreviewResponseEnvelopeErrorsMetaJSON contains the JSON metadata
// for the struct [ResourcePolicyPreviewResponseEnvelopeErrorsMeta]
type resourcePolicyPreviewResponseEnvelopeErrorsMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourcePolicyPreviewResponseEnvelopeErrorsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourcePolicyPreviewResponseEnvelopeErrorsMetaJSON) RawJSON() string {
	return r.raw
}

type ResourcePolicyPreviewResponseEnvelopeErrorsSource struct {
	Parameter           string                                                `json:"parameter"`
	ParameterValueIndex int64                                                 `json:"parameter_value_index"`
	Pointer             string                                                `json:"pointer"`
	JSON                resourcePolicyPreviewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourcePolicyPreviewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ResourcePolicyPreviewResponseEnvelopeErrorsSource]
type resourcePolicyPreviewResponseEnvelopeErrorsSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ResourcePolicyPreviewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourcePolicyPreviewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourcePolicyPreviewResponseEnvelopeMessages struct {
	Code             ResourcePolicyPreviewResponseEnvelopeMessagesCode   `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Meta             ResourcePolicyPreviewResponseEnvelopeMessagesMeta   `json:"meta"`
	Source           ResourcePolicyPreviewResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourcePolicyPreviewResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourcePolicyPreviewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ResourcePolicyPreviewResponseEnvelopeMessages]
type resourcePolicyPreviewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Meta             apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourcePolicyPreviewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourcePolicyPreviewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourcePolicyPreviewResponseEnvelopeMessagesCode int64

const (
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1001   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1001
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1002   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1002
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1003   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1003
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1004   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1004
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1005   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1005
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1006   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1006
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1007   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1007
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1008   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1008
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1009   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1009
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1010   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1010
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1011   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1011
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1012   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1012
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1013   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1013
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1014   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1014
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1015   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1015
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1016   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1016
	ResourcePolicyPreviewResponseEnvelopeMessagesCode1017   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 1017
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2001   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2001
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2002   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2002
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2003   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2003
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2004   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2004
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2005   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2005
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2006   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2006
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2007   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2007
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2008   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2008
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2009   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2009
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2010   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2010
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2011   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2011
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2012   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2012
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2013   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2013
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2014   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2014
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2015   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2015
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2016   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2016
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2017   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2017
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2018   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2018
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2019   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2019
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2020   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2020
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2021   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2021
	ResourcePolicyPreviewResponseEnvelopeMessagesCode2022   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 2022
	ResourcePolicyPreviewResponseEnvelopeMessagesCode3001   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 3001
	ResourcePolicyPreviewResponseEnvelopeMessagesCode3002   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 3002
	ResourcePolicyPreviewResponseEnvelopeMessagesCode3003   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 3003
	ResourcePolicyPreviewResponseEnvelopeMessagesCode3004   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 3004
	ResourcePolicyPreviewResponseEnvelopeMessagesCode3005   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 3005
	ResourcePolicyPreviewResponseEnvelopeMessagesCode3006   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 3006
	ResourcePolicyPreviewResponseEnvelopeMessagesCode3007   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 3007
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4001   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4001
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4002   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4002
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4003   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4003
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4004   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4004
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4005   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4005
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4006   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4006
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4007   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4007
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4008   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4008
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4009   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4009
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4010   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4010
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4011   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4011
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4012   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4012
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4013   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4013
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4014   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4014
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4015   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4015
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4016   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4016
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4017   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4017
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4018   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4018
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4019   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4019
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4020   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4020
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4021   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4021
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4022   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4022
	ResourcePolicyPreviewResponseEnvelopeMessagesCode4023   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 4023
	ResourcePolicyPreviewResponseEnvelopeMessagesCode5001   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 5001
	ResourcePolicyPreviewResponseEnvelopeMessagesCode5002   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 5002
	ResourcePolicyPreviewResponseEnvelopeMessagesCode5003   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 5003
	ResourcePolicyPreviewResponseEnvelopeMessagesCode5004   ResourcePolicyPreviewResponseEnvelopeMessagesCode = 5004
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102000 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102000
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102001 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102001
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102002 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102002
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102003 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102003
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102004 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102004
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102005 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102005
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102006 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102006
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102007 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102007
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102008 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102008
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102009 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102009
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102010 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102010
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102011 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102011
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102012 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102012
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102013 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102013
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102014 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102014
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102015 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102015
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102016 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102016
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102017 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102017
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102018 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102018
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102019 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102019
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102020 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102020
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102021 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102021
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102022 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102022
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102023 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102023
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102024 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102024
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102025 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102025
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102026 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102026
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102027 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102027
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102028 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102028
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102029 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102029
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102030 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102030
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102031 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102031
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102032 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102032
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102033 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102033
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102034 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102034
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102035 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102035
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102036 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102036
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102037 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102037
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102038 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102038
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102039 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102039
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102040 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102040
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102041 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102041
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102042 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102042
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102043 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102043
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102044 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102044
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102045 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102045
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102046 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102046
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102047 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102047
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102048 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102048
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102049 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102049
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102050 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102050
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102051 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102051
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102052 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102052
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102053 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102053
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102054 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102054
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102055 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102055
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102056 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102056
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102057 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102057
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102058 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102058
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102059 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102059
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102060 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102060
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102061 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102061
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102062 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102062
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102063 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102063
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102064 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102064
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102065 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102065
	ResourcePolicyPreviewResponseEnvelopeMessagesCode102066 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 102066
	ResourcePolicyPreviewResponseEnvelopeMessagesCode103001 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 103001
	ResourcePolicyPreviewResponseEnvelopeMessagesCode103002 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 103002
	ResourcePolicyPreviewResponseEnvelopeMessagesCode103003 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 103003
	ResourcePolicyPreviewResponseEnvelopeMessagesCode103004 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 103004
	ResourcePolicyPreviewResponseEnvelopeMessagesCode103005 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 103005
	ResourcePolicyPreviewResponseEnvelopeMessagesCode103006 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 103006
	ResourcePolicyPreviewResponseEnvelopeMessagesCode103007 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 103007
	ResourcePolicyPreviewResponseEnvelopeMessagesCode103008 ResourcePolicyPreviewResponseEnvelopeMessagesCode = 103008
)

func (r ResourcePolicyPreviewResponseEnvelopeMessagesCode) IsKnown() bool {
	switch r {
	case ResourcePolicyPreviewResponseEnvelopeMessagesCode1001, ResourcePolicyPreviewResponseEnvelopeMessagesCode1002, ResourcePolicyPreviewResponseEnvelopeMessagesCode1003, ResourcePolicyPreviewResponseEnvelopeMessagesCode1004, ResourcePolicyPreviewResponseEnvelopeMessagesCode1005, ResourcePolicyPreviewResponseEnvelopeMessagesCode1006, ResourcePolicyPreviewResponseEnvelopeMessagesCode1007, ResourcePolicyPreviewResponseEnvelopeMessagesCode1008, ResourcePolicyPreviewResponseEnvelopeMessagesCode1009, ResourcePolicyPreviewResponseEnvelopeMessagesCode1010, ResourcePolicyPreviewResponseEnvelopeMessagesCode1011, ResourcePolicyPreviewResponseEnvelopeMessagesCode1012, ResourcePolicyPreviewResponseEnvelopeMessagesCode1013, ResourcePolicyPreviewResponseEnvelopeMessagesCode1014, ResourcePolicyPreviewResponseEnvelopeMessagesCode1015, ResourcePolicyPreviewResponseEnvelopeMessagesCode1016, ResourcePolicyPreviewResponseEnvelopeMessagesCode1017, ResourcePolicyPreviewResponseEnvelopeMessagesCode2001, ResourcePolicyPreviewResponseEnvelopeMessagesCode2002, ResourcePolicyPreviewResponseEnvelopeMessagesCode2003, ResourcePolicyPreviewResponseEnvelopeMessagesCode2004, ResourcePolicyPreviewResponseEnvelopeMessagesCode2005, ResourcePolicyPreviewResponseEnvelopeMessagesCode2006, ResourcePolicyPreviewResponseEnvelopeMessagesCode2007, ResourcePolicyPreviewResponseEnvelopeMessagesCode2008, ResourcePolicyPreviewResponseEnvelopeMessagesCode2009, ResourcePolicyPreviewResponseEnvelopeMessagesCode2010, ResourcePolicyPreviewResponseEnvelopeMessagesCode2011, ResourcePolicyPreviewResponseEnvelopeMessagesCode2012, ResourcePolicyPreviewResponseEnvelopeMessagesCode2013, ResourcePolicyPreviewResponseEnvelopeMessagesCode2014, ResourcePolicyPreviewResponseEnvelopeMessagesCode2015, ResourcePolicyPreviewResponseEnvelopeMessagesCode2016, ResourcePolicyPreviewResponseEnvelopeMessagesCode2017, ResourcePolicyPreviewResponseEnvelopeMessagesCode2018, ResourcePolicyPreviewResponseEnvelopeMessagesCode2019, ResourcePolicyPreviewResponseEnvelopeMessagesCode2020, ResourcePolicyPreviewResponseEnvelopeMessagesCode2021, ResourcePolicyPreviewResponseEnvelopeMessagesCode2022, ResourcePolicyPreviewResponseEnvelopeMessagesCode3001, ResourcePolicyPreviewResponseEnvelopeMessagesCode3002, ResourcePolicyPreviewResponseEnvelopeMessagesCode3003, ResourcePolicyPreviewResponseEnvelopeMessagesCode3004, ResourcePolicyPreviewResponseEnvelopeMessagesCode3005, ResourcePolicyPreviewResponseEnvelopeMessagesCode3006, ResourcePolicyPreviewResponseEnvelopeMessagesCode3007, ResourcePolicyPreviewResponseEnvelopeMessagesCode4001, ResourcePolicyPreviewResponseEnvelopeMessagesCode4002, ResourcePolicyPreviewResponseEnvelopeMessagesCode4003, ResourcePolicyPreviewResponseEnvelopeMessagesCode4004, ResourcePolicyPreviewResponseEnvelopeMessagesCode4005, ResourcePolicyPreviewResponseEnvelopeMessagesCode4006, ResourcePolicyPreviewResponseEnvelopeMessagesCode4007, ResourcePolicyPreviewResponseEnvelopeMessagesCode4008, ResourcePolicyPreviewResponseEnvelopeMessagesCode4009, ResourcePolicyPreviewResponseEnvelopeMessagesCode4010, ResourcePolicyPreviewResponseEnvelopeMessagesCode4011, ResourcePolicyPreviewResponseEnvelopeMessagesCode4012, ResourcePolicyPreviewResponseEnvelopeMessagesCode4013, ResourcePolicyPreviewResponseEnvelopeMessagesCode4014, ResourcePolicyPreviewResponseEnvelopeMessagesCode4015, ResourcePolicyPreviewResponseEnvelopeMessagesCode4016, ResourcePolicyPreviewResponseEnvelopeMessagesCode4017, ResourcePolicyPreviewResponseEnvelopeMessagesCode4018, ResourcePolicyPreviewResponseEnvelopeMessagesCode4019, ResourcePolicyPreviewResponseEnvelopeMessagesCode4020, ResourcePolicyPreviewResponseEnvelopeMessagesCode4021, ResourcePolicyPreviewResponseEnvelopeMessagesCode4022, ResourcePolicyPreviewResponseEnvelopeMessagesCode4023, ResourcePolicyPreviewResponseEnvelopeMessagesCode5001, ResourcePolicyPreviewResponseEnvelopeMessagesCode5002, ResourcePolicyPreviewResponseEnvelopeMessagesCode5003, ResourcePolicyPreviewResponseEnvelopeMessagesCode5004, ResourcePolicyPreviewResponseEnvelopeMessagesCode102000, ResourcePolicyPreviewResponseEnvelopeMessagesCode102001, ResourcePolicyPreviewResponseEnvelopeMessagesCode102002, ResourcePolicyPreviewResponseEnvelopeMessagesCode102003, ResourcePolicyPreviewResponseEnvelopeMessagesCode102004, ResourcePolicyPreviewResponseEnvelopeMessagesCode102005, ResourcePolicyPreviewResponseEnvelopeMessagesCode102006, ResourcePolicyPreviewResponseEnvelopeMessagesCode102007, ResourcePolicyPreviewResponseEnvelopeMessagesCode102008, ResourcePolicyPreviewResponseEnvelopeMessagesCode102009, ResourcePolicyPreviewResponseEnvelopeMessagesCode102010, ResourcePolicyPreviewResponseEnvelopeMessagesCode102011, ResourcePolicyPreviewResponseEnvelopeMessagesCode102012, ResourcePolicyPreviewResponseEnvelopeMessagesCode102013, ResourcePolicyPreviewResponseEnvelopeMessagesCode102014, ResourcePolicyPreviewResponseEnvelopeMessagesCode102015, ResourcePolicyPreviewResponseEnvelopeMessagesCode102016, ResourcePolicyPreviewResponseEnvelopeMessagesCode102017, ResourcePolicyPreviewResponseEnvelopeMessagesCode102018, ResourcePolicyPreviewResponseEnvelopeMessagesCode102019, ResourcePolicyPreviewResponseEnvelopeMessagesCode102020, ResourcePolicyPreviewResponseEnvelopeMessagesCode102021, ResourcePolicyPreviewResponseEnvelopeMessagesCode102022, ResourcePolicyPreviewResponseEnvelopeMessagesCode102023, ResourcePolicyPreviewResponseEnvelopeMessagesCode102024, ResourcePolicyPreviewResponseEnvelopeMessagesCode102025, ResourcePolicyPreviewResponseEnvelopeMessagesCode102026, ResourcePolicyPreviewResponseEnvelopeMessagesCode102027, ResourcePolicyPreviewResponseEnvelopeMessagesCode102028, ResourcePolicyPreviewResponseEnvelopeMessagesCode102029, ResourcePolicyPreviewResponseEnvelopeMessagesCode102030, ResourcePolicyPreviewResponseEnvelopeMessagesCode102031, ResourcePolicyPreviewResponseEnvelopeMessagesCode102032, ResourcePolicyPreviewResponseEnvelopeMessagesCode102033, ResourcePolicyPreviewResponseEnvelopeMessagesCode102034, ResourcePolicyPreviewResponseEnvelopeMessagesCode102035, ResourcePolicyPreviewResponseEnvelopeMessagesCode102036, ResourcePolicyPreviewResponseEnvelopeMessagesCode102037, ResourcePolicyPreviewResponseEnvelopeMessagesCode102038, ResourcePolicyPreviewResponseEnvelopeMessagesCode102039, ResourcePolicyPreviewResponseEnvelopeMessagesCode102040, ResourcePolicyPreviewResponseEnvelopeMessagesCode102041, ResourcePolicyPreviewResponseEnvelopeMessagesCode102042, ResourcePolicyPreviewResponseEnvelopeMessagesCode102043, ResourcePolicyPreviewResponseEnvelopeMessagesCode102044, ResourcePolicyPreviewResponseEnvelopeMessagesCode102045, ResourcePolicyPreviewResponseEnvelopeMessagesCode102046, ResourcePolicyPreviewResponseEnvelopeMessagesCode102047, ResourcePolicyPreviewResponseEnvelopeMessagesCode102048, ResourcePolicyPreviewResponseEnvelopeMessagesCode102049, ResourcePolicyPreviewResponseEnvelopeMessagesCode102050, ResourcePolicyPreviewResponseEnvelopeMessagesCode102051, ResourcePolicyPreviewResponseEnvelopeMessagesCode102052, ResourcePolicyPreviewResponseEnvelopeMessagesCode102053, ResourcePolicyPreviewResponseEnvelopeMessagesCode102054, ResourcePolicyPreviewResponseEnvelopeMessagesCode102055, ResourcePolicyPreviewResponseEnvelopeMessagesCode102056, ResourcePolicyPreviewResponseEnvelopeMessagesCode102057, ResourcePolicyPreviewResponseEnvelopeMessagesCode102058, ResourcePolicyPreviewResponseEnvelopeMessagesCode102059, ResourcePolicyPreviewResponseEnvelopeMessagesCode102060, ResourcePolicyPreviewResponseEnvelopeMessagesCode102061, ResourcePolicyPreviewResponseEnvelopeMessagesCode102062, ResourcePolicyPreviewResponseEnvelopeMessagesCode102063, ResourcePolicyPreviewResponseEnvelopeMessagesCode102064, ResourcePolicyPreviewResponseEnvelopeMessagesCode102065, ResourcePolicyPreviewResponseEnvelopeMessagesCode102066, ResourcePolicyPreviewResponseEnvelopeMessagesCode103001, ResourcePolicyPreviewResponseEnvelopeMessagesCode103002, ResourcePolicyPreviewResponseEnvelopeMessagesCode103003, ResourcePolicyPreviewResponseEnvelopeMessagesCode103004, ResourcePolicyPreviewResponseEnvelopeMessagesCode103005, ResourcePolicyPreviewResponseEnvelopeMessagesCode103006, ResourcePolicyPreviewResponseEnvelopeMessagesCode103007, ResourcePolicyPreviewResponseEnvelopeMessagesCode103008:
		return true
	}
	return false
}

type ResourcePolicyPreviewResponseEnvelopeMessagesMeta struct {
	L10nKey       string                                                `json:"l10n_key"`
	LoggableError string                                                `json:"loggable_error"`
	TemplateData  interface{}                                           `json:"template_data"`
	TraceID       string                                                `json:"trace_id"`
	JSON          resourcePolicyPreviewResponseEnvelopeMessagesMetaJSON `json:"-"`
}

// resourcePolicyPreviewResponseEnvelopeMessagesMetaJSON contains the JSON metadata
// for the struct [ResourcePolicyPreviewResponseEnvelopeMessagesMeta]
type resourcePolicyPreviewResponseEnvelopeMessagesMetaJSON struct {
	L10nKey       apijson.Field
	LoggableError apijson.Field
	TemplateData  apijson.Field
	TraceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResourcePolicyPreviewResponseEnvelopeMessagesMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourcePolicyPreviewResponseEnvelopeMessagesMetaJSON) RawJSON() string {
	return r.raw
}

type ResourcePolicyPreviewResponseEnvelopeMessagesSource struct {
	Parameter           string                                                  `json:"parameter"`
	ParameterValueIndex int64                                                   `json:"parameter_value_index"`
	Pointer             string                                                  `json:"pointer"`
	JSON                resourcePolicyPreviewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourcePolicyPreviewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ResourcePolicyPreviewResponseEnvelopeMessagesSource]
type resourcePolicyPreviewResponseEnvelopeMessagesSourceJSON struct {
	Parameter           apijson.Field
	ParameterValueIndex apijson.Field
	Pointer             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ResourcePolicyPreviewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourcePolicyPreviewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

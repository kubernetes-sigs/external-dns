// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DnsClient a client for Dns
type DnsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDnsClientWithConfigurationProvider Creates a new default Dns client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDnsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DnsClient, err error) {
	if enabled := common.CheckForEnabledServices("dns"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDnsClientFromBaseClient(baseClient, provider)
}

// NewDnsClientWithOboToken Creates a new default Dns client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDnsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DnsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDnsClientFromBaseClient(baseClient, configProvider)
}

func newDnsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DnsClient, err error) {
	// Dns service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Dns"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DnsClient{BaseClient: baseClient}
	client.BasePath = "20180115"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DnsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("dns", "https://dns.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DnsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *DnsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ChangeResolverCompartment Moves a resolver into a different compartment along with its protected default view and any endpoints.
// Zones in the default view are not moved. VCN-dedicated resolvers are initially created in the same compartment
// as their corresponding VCN, but can then be moved to a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ChangeResolverCompartment.go.html to see an example of how to use ChangeResolverCompartment API.
// A default retry strategy applies to this operation ChangeResolverCompartment()
func (client DnsClient) ChangeResolverCompartment(ctx context.Context, request ChangeResolverCompartmentRequest) (response ChangeResolverCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeResolverCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeResolverCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeResolverCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeResolverCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeResolverCompartmentResponse")
	}
	return
}

// changeResolverCompartment implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) changeResolverCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/resolvers/{resolverId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeResolverCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Resolver/ChangeResolverCompartment"
		err = common.PostProcessServiceError(err, "Dns", "ChangeResolverCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeSteeringPolicyCompartment Moves a steering policy into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ChangeSteeringPolicyCompartment.go.html to see an example of how to use ChangeSteeringPolicyCompartment API.
// A default retry strategy applies to this operation ChangeSteeringPolicyCompartment()
func (client DnsClient) ChangeSteeringPolicyCompartment(ctx context.Context, request ChangeSteeringPolicyCompartmentRequest) (response ChangeSteeringPolicyCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeSteeringPolicyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeSteeringPolicyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeSteeringPolicyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeSteeringPolicyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeSteeringPolicyCompartmentResponse")
	}
	return
}

// changeSteeringPolicyCompartment implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) changeSteeringPolicyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/steeringPolicies/{steeringPolicyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeSteeringPolicyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/SteeringPolicy/ChangeSteeringPolicyCompartment"
		err = common.PostProcessServiceError(err, "Dns", "ChangeSteeringPolicyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeTsigKeyCompartment Moves a TSIG key into a different compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ChangeTsigKeyCompartment.go.html to see an example of how to use ChangeTsigKeyCompartment API.
// A default retry strategy applies to this operation ChangeTsigKeyCompartment()
func (client DnsClient) ChangeTsigKeyCompartment(ctx context.Context, request ChangeTsigKeyCompartmentRequest) (response ChangeTsigKeyCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeTsigKeyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeTsigKeyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeTsigKeyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeTsigKeyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeTsigKeyCompartmentResponse")
	}
	return
}

// changeTsigKeyCompartment implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) changeTsigKeyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tsigKeys/{tsigKeyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeTsigKeyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/TsigKey/ChangeTsigKeyCompartment"
		err = common.PostProcessServiceError(err, "Dns", "ChangeTsigKeyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeViewCompartment Moves a view into a different compartment.
// To change the compartment of a protected view, change the compartment of its corresponding resolver.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ChangeViewCompartment.go.html to see an example of how to use ChangeViewCompartment API.
// A default retry strategy applies to this operation ChangeViewCompartment()
func (client DnsClient) ChangeViewCompartment(ctx context.Context, request ChangeViewCompartmentRequest) (response ChangeViewCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeViewCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeViewCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeViewCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeViewCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeViewCompartmentResponse")
	}
	return
}

// changeViewCompartment implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) changeViewCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/views/{viewId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeViewCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/View/ChangeViewCompartment"
		err = common.PostProcessServiceError(err, "Dns", "ChangeViewCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeZoneCompartment Moves a zone into a different compartment.
// Protected zones cannot have their compartment changed. When the zone name is provided as a path
// parameter and `PRIVATE` is used for the scope query parameter then the viewId query parameter is
// required.
// **Note:** All SteeringPolicyAttachment objects associated with this zone will also be moved into
// the provided compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ChangeZoneCompartment.go.html to see an example of how to use ChangeZoneCompartment API.
// A default retry strategy applies to this operation ChangeZoneCompartment()
func (client DnsClient) ChangeZoneCompartment(ctx context.Context, request ChangeZoneCompartmentRequest) (response ChangeZoneCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeZoneCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeZoneCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeZoneCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeZoneCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeZoneCompartmentResponse")
	}
	return
}

// changeZoneCompartment implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) changeZoneCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/zones/{zoneId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeZoneCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Zone/ChangeZoneCompartment"
		err = common.PostProcessServiceError(err, "Dns", "ChangeZoneCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateResolverEndpoint Creates a new resolver endpoint in the same compartment as the resolver.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateResolverEndpoint.go.html to see an example of how to use CreateResolverEndpoint API.
// A default retry strategy applies to this operation CreateResolverEndpoint()
func (client DnsClient) CreateResolverEndpoint(ctx context.Context, request CreateResolverEndpointRequest) (response CreateResolverEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createResolverEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateResolverEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateResolverEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateResolverEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateResolverEndpointResponse")
	}
	return
}

// createResolverEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) createResolverEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/resolvers/{resolverId}/endpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateResolverEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "CreateResolverEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &resolverendpoint{})
	return response, err
}

// CreateSteeringPolicy Creates a new steering policy in the specified compartment. For more information on
// creating policies with templates, see Traffic Management API Guide (https://docs.cloud.oracle.com/iaas/Content/TrafficManagement/Concepts/trafficmanagementapi.htm).
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateSteeringPolicy.go.html to see an example of how to use CreateSteeringPolicy API.
// A default retry strategy applies to this operation CreateSteeringPolicy()
func (client DnsClient) CreateSteeringPolicy(ctx context.Context, request CreateSteeringPolicyRequest) (response CreateSteeringPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSteeringPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSteeringPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSteeringPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSteeringPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSteeringPolicyResponse")
	}
	return
}

// createSteeringPolicy implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) createSteeringPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/steeringPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSteeringPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "CreateSteeringPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSteeringPolicyAttachment Creates a new attachment between a steering policy and a domain, giving the
// policy permission to answer queries for the specified domain. A steering policy must
// be attached to a domain for the policy to answer DNS queries for that domain.
// For the purposes of access control, the attachment is automatically placed
// into the same compartment as the domain's zone.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateSteeringPolicyAttachment.go.html to see an example of how to use CreateSteeringPolicyAttachment API.
// A default retry strategy applies to this operation CreateSteeringPolicyAttachment()
func (client DnsClient) CreateSteeringPolicyAttachment(ctx context.Context, request CreateSteeringPolicyAttachmentRequest) (response CreateSteeringPolicyAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSteeringPolicyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSteeringPolicyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSteeringPolicyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSteeringPolicyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSteeringPolicyAttachmentResponse")
	}
	return
}

// createSteeringPolicyAttachment implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) createSteeringPolicyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/steeringPolicyAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSteeringPolicyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "CreateSteeringPolicyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTsigKey Creates a new TSIG key in the specified compartment. There is no
// `opc-retry-token` header since TSIG key names must be globally unique.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateTsigKey.go.html to see an example of how to use CreateTsigKey API.
// A default retry strategy applies to this operation CreateTsigKey()
func (client DnsClient) CreateTsigKey(ctx context.Context, request CreateTsigKeyRequest) (response CreateTsigKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createTsigKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTsigKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTsigKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTsigKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTsigKeyResponse")
	}
	return
}

// createTsigKey implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) createTsigKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/tsigKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTsigKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "CreateTsigKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateView Creates a new view in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateView.go.html to see an example of how to use CreateView API.
// A default retry strategy applies to this operation CreateView()
func (client DnsClient) CreateView(ctx context.Context, request CreateViewRequest) (response CreateViewResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createView, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateViewResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateViewResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateViewResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateViewResponse")
	}
	return
}

// createView implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) createView(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/views", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateViewResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "CreateView", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateZone Creates a new zone in the specified compartment.
// Private zones must have a zone type of `PRIMARY`. Creating a private zone at or under `oraclevcn.com`
// within the default protected view of a VCN-dedicated resolver is not permitted.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateZone.go.html to see an example of how to use CreateZone API.
func (client DnsClient) CreateZone(ctx context.Context, request CreateZoneRequest) (response CreateZoneResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createZone, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateZoneResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateZoneResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateZoneResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateZoneResponse")
	}
	return
}

// createZone implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) createZone(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/zones", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateZoneResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "CreateZone", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateZoneFromZoneFile Creates a new zone from a zone file in the specified compartment. Not supported for private zones.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateZoneFromZoneFile.go.html to see an example of how to use CreateZoneFromZoneFile API.
func (client DnsClient) CreateZoneFromZoneFile(ctx context.Context, request CreateZoneFromZoneFileRequest) (response CreateZoneFromZoneFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createZoneFromZoneFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateZoneFromZoneFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateZoneFromZoneFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateZoneFromZoneFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateZoneFromZoneFileResponse")
	}
	return
}

// createZoneFromZoneFile implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) createZoneFromZoneFile(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/actions/createZoneFromZoneFile", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateZoneFromZoneFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Zone/CreateZoneFromZoneFile"
		err = common.PostProcessServiceError(err, "Dns", "CreateZoneFromZoneFile", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDomainRecords Deletes all records at the specified zone and domain.
// When the zone name is provided as a path parameter and `PRIVATE` is used for the scope query parameter
// then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteDomainRecords.go.html to see an example of how to use DeleteDomainRecords API.
// A default retry strategy applies to this operation DeleteDomainRecords()
func (client DnsClient) DeleteDomainRecords(ctx context.Context, request DeleteDomainRecordsRequest) (response DeleteDomainRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDomainRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDomainRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDomainRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDomainRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDomainRecordsResponse")
	}
	return
}

// deleteDomainRecords implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) deleteDomainRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/zones/{zoneNameOrId}/records/{domain}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDomainRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Records/DeleteDomainRecords"
		err = common.PostProcessServiceError(err, "Dns", "DeleteDomainRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteRRSet Deletes all records in the specified RRSet.
// When the zone name is provided as a path parameter and `PRIVATE` is used for the scope
// query parameter then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteRRSet.go.html to see an example of how to use DeleteRRSet API.
// A default retry strategy applies to this operation DeleteRRSet()
func (client DnsClient) DeleteRRSet(ctx context.Context, request DeleteRRSetRequest) (response DeleteRRSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteRRSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteRRSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteRRSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteRRSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteRRSetResponse")
	}
	return
}

// deleteRRSet implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) deleteRRSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/zones/{zoneNameOrId}/records/{domain}/{rtype}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteRRSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "DeleteRRSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteResolverEndpoint Deletes the specified resolver endpoint.
// Note that attempting to delete a resolver endpoint in the DELETED lifecycle state will result in
// a `404` response to be consistent with other operations of the API. Resolver endpoints may not
// be deleted if they are referenced by a resolver rule.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteResolverEndpoint.go.html to see an example of how to use DeleteResolverEndpoint API.
func (client DnsClient) DeleteResolverEndpoint(ctx context.Context, request DeleteResolverEndpointRequest) (response DeleteResolverEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteResolverEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteResolverEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteResolverEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteResolverEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteResolverEndpointResponse")
	}
	return
}

// deleteResolverEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) deleteResolverEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/resolvers/{resolverId}/endpoints/{resolverEndpointName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteResolverEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/ResolverEndpoint/DeleteResolverEndpoint"
		err = common.PostProcessServiceError(err, "Dns", "DeleteResolverEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSteeringPolicy Deletes the specified steering policy.
// A `204` response indicates that the delete has been successful.
// Deletion will fail if the policy is attached to any zones. To detach a
// policy from a zone, see `DeleteSteeringPolicyAttachment`.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteSteeringPolicy.go.html to see an example of how to use DeleteSteeringPolicy API.
func (client DnsClient) DeleteSteeringPolicy(ctx context.Context, request DeleteSteeringPolicyRequest) (response DeleteSteeringPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSteeringPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSteeringPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSteeringPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSteeringPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSteeringPolicyResponse")
	}
	return
}

// deleteSteeringPolicy implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) deleteSteeringPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/steeringPolicies/{steeringPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSteeringPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "DeleteSteeringPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSteeringPolicyAttachment Deletes the specified steering policy attachment.
// A `204` response indicates that the delete has been successful.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteSteeringPolicyAttachment.go.html to see an example of how to use DeleteSteeringPolicyAttachment API.
func (client DnsClient) DeleteSteeringPolicyAttachment(ctx context.Context, request DeleteSteeringPolicyAttachmentRequest) (response DeleteSteeringPolicyAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSteeringPolicyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSteeringPolicyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSteeringPolicyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSteeringPolicyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSteeringPolicyAttachmentResponse")
	}
	return
}

// deleteSteeringPolicyAttachment implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) deleteSteeringPolicyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/steeringPolicyAttachments/{steeringPolicyAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSteeringPolicyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "DeleteSteeringPolicyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTsigKey Deletes the specified TSIG key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteTsigKey.go.html to see an example of how to use DeleteTsigKey API.
func (client DnsClient) DeleteTsigKey(ctx context.Context, request DeleteTsigKeyRequest) (response DeleteTsigKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTsigKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTsigKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTsigKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTsigKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTsigKeyResponse")
	}
	return
}

// deleteTsigKey implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) deleteTsigKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/tsigKeys/{tsigKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTsigKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "DeleteTsigKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteView Deletes the specified view.
// Note that attempting to delete a view in the DELETED lifecycleState will result in a `404`
// response to be consistent with other operations of the API. Views cannot be
// deleted if they are referenced by non-deleted zones or resolvers.
// Protected views cannot be deleted.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteView.go.html to see an example of how to use DeleteView API.
func (client DnsClient) DeleteView(ctx context.Context, request DeleteViewRequest) (response DeleteViewResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteView, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteViewResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteViewResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteViewResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteViewResponse")
	}
	return
}

// deleteView implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) deleteView(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/views/{viewId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteViewResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/View/DeleteView"
		err = common.PostProcessServiceError(err, "Dns", "DeleteView", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteZone Deletes the specified zone and all its steering policy attachments.
// A `204` response indicates that the zone has been successfully deleted. Protected zones cannot be deleted.
// When the zone name is provided as a path parameter and `PRIVATE` is used for the scope query parameter
// then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteZone.go.html to see an example of how to use DeleteZone API.
func (client DnsClient) DeleteZone(ctx context.Context, request DeleteZoneRequest) (response DeleteZoneResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteZone, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteZoneResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteZoneResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteZoneResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteZoneResponse")
	}
	return
}

// deleteZone implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) deleteZone(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/zones/{zoneNameOrId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteZoneResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "Dns", "DeleteZone", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDomainRecords Gets a list of all records at the specified zone and domain.
// The results are sorted by `rtype` in alphabetical order by default. You can optionally filter and/or sort
// the results using the listed parameters. When the zone name is provided as a path parameter and `PRIVATE`
// is used for the scope query parameter then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetDomainRecords.go.html to see an example of how to use GetDomainRecords API.
// A default retry strategy applies to this operation GetDomainRecords()
func (client DnsClient) GetDomainRecords(ctx context.Context, request GetDomainRecordsRequest) (response GetDomainRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDomainRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDomainRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDomainRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDomainRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDomainRecordsResponse")
	}
	return
}

// getDomainRecords implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getDomainRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zones/{zoneNameOrId}/records/{domain}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDomainRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Records/GetDomainRecords"
		err = common.PostProcessServiceError(err, "Dns", "GetDomainRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRRSet Gets a list of all records in the specified RRSet.
// The results are sorted by `recordHash` by default. When the zone name is provided as a path parameter
// and `PRIVATE` is used for the scope query parameter then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetRRSet.go.html to see an example of how to use GetRRSet API.
// A default retry strategy applies to this operation GetRRSet()
func (client DnsClient) GetRRSet(ctx context.Context, request GetRRSetRequest) (response GetRRSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRRSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRRSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRRSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRRSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRRSetResponse")
	}
	return
}

// getRRSet implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getRRSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zones/{zoneNameOrId}/records/{domain}/{rtype}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRRSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/RRSet/GetRRSet"
		err = common.PostProcessServiceError(err, "Dns", "GetRRSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResolver Gets information about a specific resolver.
// Note that attempting to get a resolver in the DELETED lifecycleState will result in a `404`
// response to be consistent with other operations of the API.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetResolver.go.html to see an example of how to use GetResolver API.
// A default retry strategy applies to this operation GetResolver()
func (client DnsClient) GetResolver(ctx context.Context, request GetResolverRequest) (response GetResolverResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResolver, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResolverResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResolverResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResolverResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResolverResponse")
	}
	return
}

// getResolver implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getResolver(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resolvers/{resolverId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResolverResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Resolver/GetResolver"
		err = common.PostProcessServiceError(err, "Dns", "GetResolver", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetResolverEndpoint Gets information about a specific resolver endpoint.
// Note that attempting to get a resolver endpoint in the DELETED lifecycle state will result
// in a `404` response to be consistent with other operations of the API.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetResolverEndpoint.go.html to see an example of how to use GetResolverEndpoint API.
// A default retry strategy applies to this operation GetResolverEndpoint()
func (client DnsClient) GetResolverEndpoint(ctx context.Context, request GetResolverEndpointRequest) (response GetResolverEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getResolverEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetResolverEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetResolverEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetResolverEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetResolverEndpointResponse")
	}
	return
}

// getResolverEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getResolverEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resolvers/{resolverId}/endpoints/{resolverEndpointName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetResolverEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/ResolverEndpoint/GetResolverEndpoint"
		err = common.PostProcessServiceError(err, "Dns", "GetResolverEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &resolverendpoint{})
	return response, err
}

// GetSteeringPolicy Gets information about the specified steering policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetSteeringPolicy.go.html to see an example of how to use GetSteeringPolicy API.
// A default retry strategy applies to this operation GetSteeringPolicy()
func (client DnsClient) GetSteeringPolicy(ctx context.Context, request GetSteeringPolicyRequest) (response GetSteeringPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSteeringPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSteeringPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSteeringPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSteeringPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSteeringPolicyResponse")
	}
	return
}

// getSteeringPolicy implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getSteeringPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/steeringPolicies/{steeringPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSteeringPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/SteeringPolicy/GetSteeringPolicy"
		err = common.PostProcessServiceError(err, "Dns", "GetSteeringPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSteeringPolicyAttachment Gets information about the specified steering policy attachment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetSteeringPolicyAttachment.go.html to see an example of how to use GetSteeringPolicyAttachment API.
// A default retry strategy applies to this operation GetSteeringPolicyAttachment()
func (client DnsClient) GetSteeringPolicyAttachment(ctx context.Context, request GetSteeringPolicyAttachmentRequest) (response GetSteeringPolicyAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSteeringPolicyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSteeringPolicyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSteeringPolicyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSteeringPolicyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSteeringPolicyAttachmentResponse")
	}
	return
}

// getSteeringPolicyAttachment implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getSteeringPolicyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/steeringPolicyAttachments/{steeringPolicyAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSteeringPolicyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/SteeringPolicyAttachment/GetSteeringPolicyAttachment"
		err = common.PostProcessServiceError(err, "Dns", "GetSteeringPolicyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTsigKey Gets information about the specified TSIG key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetTsigKey.go.html to see an example of how to use GetTsigKey API.
// A default retry strategy applies to this operation GetTsigKey()
func (client DnsClient) GetTsigKey(ctx context.Context, request GetTsigKeyRequest) (response GetTsigKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTsigKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTsigKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTsigKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTsigKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTsigKeyResponse")
	}
	return
}

// getTsigKey implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getTsigKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tsigKeys/{tsigKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTsigKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/TsigKey/GetTsigKey"
		err = common.PostProcessServiceError(err, "Dns", "GetTsigKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetView Gets information about a specific view.
// Note that attempting to get a
// view in the DELETED lifecycleState will result in a `404` response to be
// consistent with other operations of the API.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetView.go.html to see an example of how to use GetView API.
// A default retry strategy applies to this operation GetView()
func (client DnsClient) GetView(ctx context.Context, request GetViewRequest) (response GetViewResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getView, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetViewResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetViewResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetViewResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetViewResponse")
	}
	return
}

// getView implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getView(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/views/{viewId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetViewResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/View/GetView"
		err = common.PostProcessServiceError(err, "Dns", "GetView", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetZone Gets information about the specified zone, including its creation date, zone type, and serial.
// When the zone name is provided as a path parameter and `PRIVATE` is used for the scope query
// parameter then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetZone.go.html to see an example of how to use GetZone API.
// A default retry strategy applies to this operation GetZone()
func (client DnsClient) GetZone(ctx context.Context, request GetZoneRequest) (response GetZoneResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getZone, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetZoneResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetZoneResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetZoneResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetZoneResponse")
	}
	return
}

// getZone implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getZone(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zones/{zoneNameOrId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetZoneResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Zone/GetZone"
		err = common.PostProcessServiceError(err, "Dns", "GetZone", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetZoneContent Gets the requested zone's zone file.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetZoneContent.go.html to see an example of how to use GetZoneContent API.
// A default retry strategy applies to this operation GetZoneContent()
func (client DnsClient) GetZoneContent(ctx context.Context, request GetZoneContentRequest) (response GetZoneContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getZoneContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetZoneContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetZoneContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetZoneContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetZoneContentResponse")
	}
	return
}

// getZoneContent implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getZoneContent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zones/{zoneNameOrId}/content", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetZoneContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Zone/GetZoneContent"
		err = common.PostProcessServiceError(err, "Dns", "GetZoneContent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetZoneRecords Gets all records in the specified zone.
// The results are sorted by `domain` in alphabetical order by default. For more information about records,
// see Resource Record (RR) TYPEs (https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4).
// When the zone name is provided as a path parameter and `PRIVATE` is used for the scope query parameter
// then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetZoneRecords.go.html to see an example of how to use GetZoneRecords API.
// A default retry strategy applies to this operation GetZoneRecords()
func (client DnsClient) GetZoneRecords(ctx context.Context, request GetZoneRecordsRequest) (response GetZoneRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getZoneRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetZoneRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetZoneRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetZoneRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetZoneRecordsResponse")
	}
	return
}

// getZoneRecords implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) getZoneRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zones/{zoneNameOrId}/records", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetZoneRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Records/GetZoneRecords"
		err = common.PostProcessServiceError(err, "Dns", "GetZoneRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// listresolverendpointsummary allows to unmarshal list of polymorphic ResolverEndpointSummary
type listresolverendpointsummary []resolverendpointsummary

// UnmarshalPolymorphicJSON unmarshals polymorphic json list of items
func (m *listresolverendpointsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	res := make([]ResolverEndpointSummary, len(*m))
	for i, v := range *m {
		nn, err := v.UnmarshalPolymorphicJSON(v.JsonData)
		if err != nil {
			return nil, err
		}
		res[i] = nn.(ResolverEndpointSummary)
	}
	return res, nil
}

// ListResolverEndpoints Gets a list of all endpoints within a resolver. The collection can be filtered by name or lifecycle state.
// It can be sorted on creation time or name both in ASC or DESC order. Note that when no lifecycleState
// query parameter is provided, the collection does not include resolver endpoints in the DELETED
// lifecycle state to be consistent with other operations of the API.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListResolverEndpoints.go.html to see an example of how to use ListResolverEndpoints API.
// A default retry strategy applies to this operation ListResolverEndpoints()
func (client DnsClient) ListResolverEndpoints(ctx context.Context, request ListResolverEndpointsRequest) (response ListResolverEndpointsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResolverEndpoints, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResolverEndpointsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResolverEndpointsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResolverEndpointsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResolverEndpointsResponse")
	}
	return
}

// listResolverEndpoints implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) listResolverEndpoints(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resolvers/{resolverId}/endpoints", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResolverEndpointsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/ResolverEndpoint/ListResolverEndpoints"
		err = common.PostProcessServiceError(err, "Dns", "ListResolverEndpoints", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &listresolverendpointsummary{})
	return response, err
}

// ListResolvers Gets a list of all resolvers within a compartment.
// The collection can be filtered by display name, id, or lifecycle state. It can be sorted
// on creation time or displayName both in ASC or DESC order. Note that when no lifecycleState
// query parameter is provided, the collection does not include resolvers in the DELETED
// lifecycleState to be consistent with other operations of the API.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListResolvers.go.html to see an example of how to use ListResolvers API.
// A default retry strategy applies to this operation ListResolvers()
func (client DnsClient) ListResolvers(ctx context.Context, request ListResolversRequest) (response ListResolversResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listResolvers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListResolversResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListResolversResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListResolversResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListResolversResponse")
	}
	return
}

// listResolvers implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) listResolvers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/resolvers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListResolversResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Resolver/ListResolvers"
		err = common.PostProcessServiceError(err, "Dns", "ListResolvers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSteeringPolicies Gets a list of all steering policies in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListSteeringPolicies.go.html to see an example of how to use ListSteeringPolicies API.
// A default retry strategy applies to this operation ListSteeringPolicies()
func (client DnsClient) ListSteeringPolicies(ctx context.Context, request ListSteeringPoliciesRequest) (response ListSteeringPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSteeringPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSteeringPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSteeringPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSteeringPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSteeringPoliciesResponse")
	}
	return
}

// listSteeringPolicies implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) listSteeringPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/steeringPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSteeringPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/SteeringPolicy/ListSteeringPolicies"
		err = common.PostProcessServiceError(err, "Dns", "ListSteeringPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSteeringPolicyAttachments Lists the steering policy attachments in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListSteeringPolicyAttachments.go.html to see an example of how to use ListSteeringPolicyAttachments API.
// A default retry strategy applies to this operation ListSteeringPolicyAttachments()
func (client DnsClient) ListSteeringPolicyAttachments(ctx context.Context, request ListSteeringPolicyAttachmentsRequest) (response ListSteeringPolicyAttachmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSteeringPolicyAttachments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSteeringPolicyAttachmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSteeringPolicyAttachmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSteeringPolicyAttachmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSteeringPolicyAttachmentsResponse")
	}
	return
}

// listSteeringPolicyAttachments implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) listSteeringPolicyAttachments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/steeringPolicyAttachments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSteeringPolicyAttachmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/SteeringPolicyAttachment/ListSteeringPolicyAttachments"
		err = common.PostProcessServiceError(err, "Dns", "ListSteeringPolicyAttachments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTsigKeys Gets a list of all TSIG keys in the specified compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListTsigKeys.go.html to see an example of how to use ListTsigKeys API.
// A default retry strategy applies to this operation ListTsigKeys()
func (client DnsClient) ListTsigKeys(ctx context.Context, request ListTsigKeysRequest) (response ListTsigKeysResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTsigKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTsigKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTsigKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTsigKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTsigKeysResponse")
	}
	return
}

// listTsigKeys implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) listTsigKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/tsigKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTsigKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/TsigKey/ListTsigKeys"
		err = common.PostProcessServiceError(err, "Dns", "ListTsigKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListViews Gets a list of all views within a compartment.
// The collection can be filtered by display name, id, or lifecycle state. It can be sorted
// on creation time or displayName both in ASC or DESC order. Note that when no lifecycleState
// query parameter is provided, the collection does not include views in the DELETED
// lifecycleState to be consistent with other operations of the API.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListViews.go.html to see an example of how to use ListViews API.
// A default retry strategy applies to this operation ListViews()
func (client DnsClient) ListViews(ctx context.Context, request ListViewsRequest) (response ListViewsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listViews, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListViewsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListViewsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListViewsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListViewsResponse")
	}
	return
}

// listViews implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) listViews(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/views", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListViewsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/View/ListViews"
		err = common.PostProcessServiceError(err, "Dns", "ListViews", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListZoneTransferServers Gets a list of IP addresses of OCI nameservers for inbound and outbound transfer of zones in the specified
// compartment (which must be the root compartment of a tenancy) that transfer zone data with external master or
// downstream nameservers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListZoneTransferServers.go.html to see an example of how to use ListZoneTransferServers API.
// A default retry strategy applies to this operation ListZoneTransferServers()
func (client DnsClient) ListZoneTransferServers(ctx context.Context, request ListZoneTransferServersRequest) (response ListZoneTransferServersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listZoneTransferServers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListZoneTransferServersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListZoneTransferServersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListZoneTransferServersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListZoneTransferServersResponse")
	}
	return
}

// listZoneTransferServers implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) listZoneTransferServers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zoneTransferServers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListZoneTransferServersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/ZoneTransferServer/ListZoneTransferServers"
		err = common.PostProcessServiceError(err, "Dns", "ListZoneTransferServers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListZones Gets a list of all zones in the specified compartment.
// The collection can be filtered by name, time created, scope, associated view, and zone type.
// Filtering by view is only supported for private zones.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListZones.go.html to see an example of how to use ListZones API.
// A default retry strategy applies to this operation ListZones()
func (client DnsClient) ListZones(ctx context.Context, request ListZonesRequest) (response ListZonesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listZones, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListZonesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListZonesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListZonesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListZonesResponse")
	}
	return
}

// listZones implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) listZones(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/zones", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListZonesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Zone/ListZones"
		err = common.PostProcessServiceError(err, "Dns", "ListZones", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchDomainRecords Updates records in the specified zone at a domain.
// You can update one record or all records for the specified zone depending on the changes provided in the
// request body. You can also add or remove records using this function. When the zone name is provided as
// a path parameter and `PRIVATE` is used for the scope query parameter then the viewId query parameter is
// required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/PatchDomainRecords.go.html to see an example of how to use PatchDomainRecords API.
// A default retry strategy applies to this operation PatchDomainRecords()
func (client DnsClient) PatchDomainRecords(ctx context.Context, request PatchDomainRecordsRequest) (response PatchDomainRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchDomainRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchDomainRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchDomainRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchDomainRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchDomainRecordsResponse")
	}
	return
}

// patchDomainRecords implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) patchDomainRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/zones/{zoneNameOrId}/records/{domain}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchDomainRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Records/PatchDomainRecords"
		err = common.PostProcessServiceError(err, "Dns", "PatchDomainRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchRRSet Updates records in the specified RRSet.
// When the zone name is provided as a path parameter and `PRIVATE` is used for the scope query
// parameter then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/PatchRRSet.go.html to see an example of how to use PatchRRSet API.
// A default retry strategy applies to this operation PatchRRSet()
func (client DnsClient) PatchRRSet(ctx context.Context, request PatchRRSetRequest) (response PatchRRSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchRRSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchRRSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchRRSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchRRSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchRRSetResponse")
	}
	return
}

// patchRRSet implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) patchRRSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/zones/{zoneNameOrId}/records/{domain}/{rtype}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchRRSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/RRSet/PatchRRSet"
		err = common.PostProcessServiceError(err, "Dns", "PatchRRSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PatchZoneRecords Updates a collection of records in the specified zone.
// You can update one record or all records for the specified zone depending on the changes provided in the
// request body. You can also add or remove records using this function. When the zone name is provided as
// a path parameter and `PRIVATE` is used for the scope query parameter then the viewId query parameter is
// required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/PatchZoneRecords.go.html to see an example of how to use PatchZoneRecords API.
// A default retry strategy applies to this operation PatchZoneRecords()
func (client DnsClient) PatchZoneRecords(ctx context.Context, request PatchZoneRecordsRequest) (response PatchZoneRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.patchZoneRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PatchZoneRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PatchZoneRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PatchZoneRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PatchZoneRecordsResponse")
	}
	return
}

// patchZoneRecords implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) patchZoneRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPatch, "/zones/{zoneNameOrId}/records", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PatchZoneRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Records/PatchZoneRecords"
		err = common.PostProcessServiceError(err, "Dns", "PatchZoneRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDomainRecords Replaces records in the specified zone at a domain with the records specified in the request body.
// If a specified record does not exist, it will be created. If the record exists, then it will be updated to
// represent the record in the body of the request. If a record in the zone does not exist in the request body,
// the record will be removed from the zone. When the zone name is provided as a path parameter and `PRIVATE`
// is used for the scope query parameter then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateDomainRecords.go.html to see an example of how to use UpdateDomainRecords API.
// A default retry strategy applies to this operation UpdateDomainRecords()
func (client DnsClient) UpdateDomainRecords(ctx context.Context, request UpdateDomainRecordsRequest) (response UpdateDomainRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDomainRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDomainRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDomainRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDomainRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDomainRecordsResponse")
	}
	return
}

// updateDomainRecords implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateDomainRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/zones/{zoneNameOrId}/records/{domain}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDomainRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Records/UpdateDomainRecords"
		err = common.PostProcessServiceError(err, "Dns", "UpdateDomainRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRRSet Replaces records in the specified RRSet.
// When the zone name is provided as a path parameter and `PRIVATE` is used for the scope
// query parameter then the viewId query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateRRSet.go.html to see an example of how to use UpdateRRSet API.
// A default retry strategy applies to this operation UpdateRRSet()
func (client DnsClient) UpdateRRSet(ctx context.Context, request UpdateRRSetRequest) (response UpdateRRSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRRSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRRSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRRSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRRSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRRSetResponse")
	}
	return
}

// updateRRSet implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateRRSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/zones/{zoneNameOrId}/records/{domain}/{rtype}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRRSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/RRSet/UpdateRRSet"
		err = common.PostProcessServiceError(err, "Dns", "UpdateRRSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateResolver Updates the specified resolver with your new information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateResolver.go.html to see an example of how to use UpdateResolver API.
// A default retry strategy applies to this operation UpdateResolver()
func (client DnsClient) UpdateResolver(ctx context.Context, request UpdateResolverRequest) (response UpdateResolverResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResolver, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResolverResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResolverResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResolverResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResolverResponse")
	}
	return
}

// updateResolver implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateResolver(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/resolvers/{resolverId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateResolverResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Resolver/UpdateResolver"
		err = common.PostProcessServiceError(err, "Dns", "UpdateResolver", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateResolverEndpoint Updates the specified resolver endpoint with your new information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateResolverEndpoint.go.html to see an example of how to use UpdateResolverEndpoint API.
// A default retry strategy applies to this operation UpdateResolverEndpoint()
func (client DnsClient) UpdateResolverEndpoint(ctx context.Context, request UpdateResolverEndpointRequest) (response UpdateResolverEndpointResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateResolverEndpoint, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateResolverEndpointResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateResolverEndpointResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateResolverEndpointResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateResolverEndpointResponse")
	}
	return
}

// updateResolverEndpoint implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateResolverEndpoint(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/resolvers/{resolverId}/endpoints/{resolverEndpointName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateResolverEndpointResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/ResolverEndpoint/UpdateResolverEndpoint"
		err = common.PostProcessServiceError(err, "Dns", "UpdateResolverEndpoint", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &resolverendpoint{})
	return response, err
}

// UpdateSteeringPolicy Updates the configuration of the specified steering policy.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateSteeringPolicy.go.html to see an example of how to use UpdateSteeringPolicy API.
// A default retry strategy applies to this operation UpdateSteeringPolicy()
func (client DnsClient) UpdateSteeringPolicy(ctx context.Context, request UpdateSteeringPolicyRequest) (response UpdateSteeringPolicyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSteeringPolicy, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSteeringPolicyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSteeringPolicyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSteeringPolicyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSteeringPolicyResponse")
	}
	return
}

// updateSteeringPolicy implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateSteeringPolicy(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/steeringPolicies/{steeringPolicyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSteeringPolicyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/SteeringPolicy/UpdateSteeringPolicy"
		err = common.PostProcessServiceError(err, "Dns", "UpdateSteeringPolicy", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSteeringPolicyAttachment Updates the specified steering policy attachment with your new information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateSteeringPolicyAttachment.go.html to see an example of how to use UpdateSteeringPolicyAttachment API.
// A default retry strategy applies to this operation UpdateSteeringPolicyAttachment()
func (client DnsClient) UpdateSteeringPolicyAttachment(ctx context.Context, request UpdateSteeringPolicyAttachmentRequest) (response UpdateSteeringPolicyAttachmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSteeringPolicyAttachment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSteeringPolicyAttachmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSteeringPolicyAttachmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSteeringPolicyAttachmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSteeringPolicyAttachmentResponse")
	}
	return
}

// updateSteeringPolicyAttachment implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateSteeringPolicyAttachment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/steeringPolicyAttachments/{steeringPolicyAttachmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSteeringPolicyAttachmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/SteeringPolicyAttachment/UpdateSteeringPolicyAttachment"
		err = common.PostProcessServiceError(err, "Dns", "UpdateSteeringPolicyAttachment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTsigKey Updates the specified TSIG key.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateTsigKey.go.html to see an example of how to use UpdateTsigKey API.
// A default retry strategy applies to this operation UpdateTsigKey()
func (client DnsClient) UpdateTsigKey(ctx context.Context, request UpdateTsigKeyRequest) (response UpdateTsigKeyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTsigKey, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTsigKeyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTsigKeyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTsigKeyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTsigKeyResponse")
	}
	return
}

// updateTsigKey implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateTsigKey(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/tsigKeys/{tsigKeyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTsigKeyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/TsigKey/UpdateTsigKey"
		err = common.PostProcessServiceError(err, "Dns", "UpdateTsigKey", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateView Updates the specified view with your new information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateView.go.html to see an example of how to use UpdateView API.
// A default retry strategy applies to this operation UpdateView()
func (client DnsClient) UpdateView(ctx context.Context, request UpdateViewRequest) (response UpdateViewResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateView, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateViewResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateViewResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateViewResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateViewResponse")
	}
	return
}

// updateView implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateView(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/views/{viewId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateViewResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/View/UpdateView"
		err = common.PostProcessServiceError(err, "Dns", "UpdateView", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateZone Updates the zone with the specified information.
// Global secondary zones may have their external masters updated. For more information about secondary
// zones, see Manage DNS Service Zone (https://docs.cloud.oracle.com/iaas/Content/DNS/Tasks/managingdnszones.htm). When the zone name
// is provided as a path parameter and `PRIVATE` is used for the scope query parameter then the viewId
// query parameter is required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateZone.go.html to see an example of how to use UpdateZone API.
// A default retry strategy applies to this operation UpdateZone()
func (client DnsClient) UpdateZone(ctx context.Context, request UpdateZoneRequest) (response UpdateZoneResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateZone, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateZoneResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateZoneResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateZoneResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateZoneResponse")
	}
	return
}

// updateZone implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateZone(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/zones/{zoneNameOrId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateZoneResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Zone/UpdateZone"
		err = common.PostProcessServiceError(err, "Dns", "UpdateZone", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateZoneRecords Replaces records in the specified zone with the records specified in the request body.
// If a specified record does not exist, it will be created. If the record exists, then it will be updated
// to represent the record in the body of the request. If a record in the zone does not exist in the
// request body, the record will be removed from the zone. When the zone name is provided as a path
// parameter and `PRIVATE` is used for the scope query parameter then the viewId query parameter is
// required.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateZoneRecords.go.html to see an example of how to use UpdateZoneRecords API.
// A default retry strategy applies to this operation UpdateZoneRecords()
func (client DnsClient) UpdateZoneRecords(ctx context.Context, request UpdateZoneRecordsRequest) (response UpdateZoneRecordsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateZoneRecords, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateZoneRecordsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateZoneRecordsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateZoneRecordsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateZoneRecordsResponse")
	}
	return
}

// updateZoneRecords implements the OCIOperation interface (enables retrying operations)
func (client DnsClient) updateZoneRecords(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/zones/{zoneNameOrId}/records", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateZoneRecordsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/dns/20180115/Records/UpdateZoneRecords"
		err = common.PostProcessServiceError(err, "Dns", "UpdateZoneRecords", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

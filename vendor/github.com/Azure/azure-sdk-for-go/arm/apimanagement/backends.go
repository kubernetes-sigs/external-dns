package apimanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// BackendsClient is the composite Swagger for ApiManagement Client
type BackendsClient struct {
	ManagementClient
}

// NewBackendsClient creates an instance of the BackendsClient client.
func NewBackendsClient(subscriptionID string) BackendsClient {
	return NewBackendsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackendsClientWithBaseURI creates an instance of the BackendsClient
// client.
func NewBackendsClientWithBaseURI(baseURI string, subscriptionID string) BackendsClient {
	return BackendsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or Updates a backend.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. backendid is identifier of the Backend
// entity. Must be unique in the current API Management service instance.
// parameters is create parameters.
func (client BackendsClient) CreateOrUpdate(resourceGroupName string, serviceName string, backendid string, parameters BackendContract) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: backendid,
			Constraints: []validation.Constraint{{Target: "backendid", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "backendid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "backendid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.BackendsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, serviceName, backendid, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BackendsClient) CreateOrUpdatePreparer(resourceGroupName string, serviceName string, backendid string, parameters BackendContract) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendid":         autorest.Encode("path", backendid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BackendsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BackendsClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes the specified backend.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. backendid is identifier of the Backend
// entity. Must be unique in the current API Management service instance.
// ifMatch is the entity state (Etag) version of the backend to delete. A value
// of "*" can be used for If-Match to unconditionally apply the operation.
func (client BackendsClient) Delete(resourceGroupName string, serviceName string, backendid string, ifMatch string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: backendid,
			Constraints: []validation.Constraint{{Target: "backendid", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "backendid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "backendid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.BackendsClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, serviceName, backendid, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BackendsClient) DeletePreparer(resourceGroupName string, serviceName string, backendid string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendid":         autorest.Encode("path", backendid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendid}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackendsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackendsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the backend specified by its identifier.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. backendid is identifier of the Backend
// entity. Must be unique in the current API Management service instance.
func (client BackendsClient) Get(resourceGroupName string, serviceName string, backendid string) (result BackendResponse, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: backendid,
			Constraints: []validation.Constraint{{Target: "backendid", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "backendid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "backendid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.BackendsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, serviceName, backendid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackendsClient) GetPreparer(resourceGroupName string, serviceName string, backendid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendid":         autorest.Encode("path", backendid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackendsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackendsClient) GetResponder(resp *http.Response) (result BackendResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByService lists a collection of backends in the specified service
// instance.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. filter is | Field | Supported operators    |
// Supported functions                         |
// |-------|------------------------|---------------------------------------------|
// | id    | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | host  | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith | top is number of records to return. skip is number of records to
// skip.
func (client BackendsClient) ListByService(resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result BackendCollection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.BackendsClient", "ListByService")
	}

	req, err := client.ListByServicePreparer(resourceGroupName, serviceName, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "ListByService", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client BackendsClient) ListByServicePreparer(resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client BackendsClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client BackendsClient) ListByServiceResponder(resp *http.Response) (result BackendCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServiceNextResults retrieves the next set of results, if any.
func (client BackendsClient) ListByServiceNextResults(lastResults BackendCollection) (result BackendCollection, err error) {
	req, err := lastResults.BackendCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "ListByService", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "ListByService", resp, "Failure sending next results request")
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "ListByService", resp, "Failure responding to next results request")
	}

	return
}

// Update updates an existing backend.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. backendid is identifier of the Backend
// entity. Must be unique in the current API Management service instance.
// parameters is update parameters. ifMatch is the entity state (Etag) version
// of the backend to update. A value of "*" can be used for If-Match to
// unconditionally apply the operation.
func (client BackendsClient) Update(resourceGroupName string, serviceName string, backendid string, parameters BackendUpdateParameters, ifMatch string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: backendid,
			Constraints: []validation.Constraint{{Target: "backendid", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "backendid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "backendid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.BackendsClient", "Update")
	}

	req, err := client.UpdatePreparer(resourceGroupName, serviceName, backendid, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.BackendsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client BackendsClient) UpdatePreparer(resourceGroupName string, serviceName string, backendid string, parameters BackendUpdateParameters, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendid":         autorest.Encode("path", backendid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client BackendsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client BackendsClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

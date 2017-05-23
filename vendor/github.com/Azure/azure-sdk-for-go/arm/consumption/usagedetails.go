package consumption

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

// UsageDetailsClient is the consumption management client provides access to
// consumption resources for Azure Web-Direct subscriptions. Other subscription
// types which were not purchased directly through the Azure web portal are not
// supported through this preview API.
type UsageDetailsClient struct {
	ManagementClient
}

// NewUsageDetailsClient creates an instance of the UsageDetailsClient client.
func NewUsageDetailsClient(subscriptionID string) UsageDetailsClient {
	return NewUsageDetailsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageDetailsClientWithBaseURI creates an instance of the
// UsageDetailsClient client.
func NewUsageDetailsClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsClient {
	return UsageDetailsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the usage details for a scope in reverse chronological order by
// billing period. Usage details are available via this API only for January 1,
// 2017 or later.
//
// scope is the scope of the usage details. The scope can be
// '/subscriptions/{subscriptionId}/' for a subscription, or
// '/subscriptions/{subscriptionId}/providers/Microsoft.Billing/invoices/{invoiceName}'
// for an invoice or
// '/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'
// for a billing perdiod. expand is may be used to expand the
// additionalProperties or meterDetails property within a list of usage
// details. By default, these fields are not included when listing usage
// details. filter is may be used to filter usageDetails by usageEnd (Utc
// time). The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'. It does
// not currently support 'ne', 'or', or 'not'. skiptoken is skiptoken is only
// used if a previous operation returned a partial result. If a previous
// response contains a nextLink element, the value of the nextLink element will
// include a skiptoken parameter that specifies a starting point to use for
// subsequent calls. top is may be used to limit the number of results to the
// most recent N usageDetails.
func (client UsageDetailsClient) List(scope string, expand string, filter string, skiptoken string, top *int32) (result UsageDetailsListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "consumption.UsageDetailsClient", "List")
	}

	req, err := client.ListPreparer(scope, expand, filter, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client UsageDetailsClient) ListPreparer(scope string, expand string, filter string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": autorest.Encode("path", scope),
	}

	const APIVersion = "2017-04-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Consumption/usageDetails", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UsageDetailsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UsageDetailsClient) ListResponder(resp *http.Response) (result UsageDetailsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client UsageDetailsClient) ListNextResults(lastResults UsageDetailsListResult) (result UsageDetailsListResult, err error) {
	req, err := lastResults.UsageDetailsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.UsageDetailsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

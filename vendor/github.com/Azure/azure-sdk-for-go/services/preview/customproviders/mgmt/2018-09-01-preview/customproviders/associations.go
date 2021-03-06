package customproviders

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AssociationsClient is the allows extension of ARM control plane with custom resource providers.
type AssociationsClient struct {
	BaseClient
}

// NewAssociationsClient creates an instance of the AssociationsClient client.
func NewAssociationsClient(subscriptionID string) AssociationsClient {
	return NewAssociationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAssociationsClientWithBaseURI creates an instance of the AssociationsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAssociationsClientWithBaseURI(baseURI string, subscriptionID string) AssociationsClient {
	return AssociationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update an association.
// Parameters:
// scope - the scope of the association. The scope can be any valid REST resource instance. For example, use
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/Microsoft.Compute/virtualMachines/{vm-name}'
// for a virtual machine resource.
// associationName - the name of the association.
// association - the parameters required to create or update an association.
func (client AssociationsClient) CreateOrUpdate(ctx context.Context, scope string, associationName string, association Association) (result AssociationsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssociationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, scope, associationName, association)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AssociationsClient) CreateOrUpdatePreparer(ctx context.Context, scope string, associationName string, association Association) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"associationName": autorest.Encode("path", associationName),
		"scope":           scope,
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	association.ID = nil
	association.Name = nil
	association.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}", pathParameters),
		autorest.WithJSON(association),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AssociationsClient) CreateOrUpdateSender(req *http.Request) (future AssociationsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AssociationsClient) CreateOrUpdateResponder(resp *http.Response) (result Association, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an association.
// Parameters:
// scope - the scope of the association.
// associationName - the name of the association.
func (client AssociationsClient) Delete(ctx context.Context, scope string, associationName string) (result AssociationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssociationsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, scope, associationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AssociationsClient) DeletePreparer(ctx context.Context, scope string, associationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"associationName": autorest.Encode("path", associationName),
		"scope":           scope,
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AssociationsClient) DeleteSender(req *http.Request) (future AssociationsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AssociationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an association.
// Parameters:
// scope - the scope of the association.
// associationName - the name of the association.
func (client AssociationsClient) Get(ctx context.Context, scope string, associationName string) (result Association, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssociationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, associationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssociationsClient) GetPreparer(ctx context.Context, scope string, associationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"associationName": autorest.Encode("path", associationName),
		"scope":           scope,
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssociationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssociationsClient) GetResponder(resp *http.Response) (result Association, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAll gets all association for the given scope.
// Parameters:
// scope - the scope of the association.
func (client AssociationsClient) ListAll(ctx context.Context, scope string) (result AssociationsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssociationsClient.ListAll")
		defer func() {
			sc := -1
			if result.al.Response.Response != nil {
				sc = result.al.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.al, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "ListAll", resp, "Failure responding to request")
		return
	}
	if result.al.hasNextLink() && result.al.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client AssociationsClient) ListAllPreparer(ctx context.Context, scope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.CustomProviders/associations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client AssociationsClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client AssociationsClient) ListAllResponder(resp *http.Response) (result AssociationsList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client AssociationsClient) listAllNextResults(ctx context.Context, lastResults AssociationsList) (result AssociationsList, err error) {
	req, err := lastResults.associationsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customproviders.AssociationsClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client AssociationsClient) ListAllComplete(ctx context.Context, scope string) (result AssociationsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssociationsClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx, scope)
	return
}

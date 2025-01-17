//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AuthorizationPoliciesClient contains the methods for the AuthorizationPolicies group.
// Don't use this type directly, use NewAuthorizationPoliciesClient() instead.
type AuthorizationPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAuthorizationPoliciesClient creates a new instance of AuthorizationPoliciesClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAuthorizationPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AuthorizationPoliciesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AuthorizationPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates an authorization policy or updates an existing authorization policy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// authorizationPolicyName - The name of the policy.
// parameters - Parameters supplied to the CreateOrUpdate authorization policy operation.
// options - AuthorizationPoliciesClientCreateOrUpdateOptions contains the optional parameters for the AuthorizationPoliciesClient.CreateOrUpdate
// method.
func (client *AuthorizationPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, parameters AuthorizationPolicyResourceFormat, options *AuthorizationPoliciesClientCreateOrUpdateOptions) (AuthorizationPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, authorizationPolicyName, parameters, options)
	if err != nil {
		return AuthorizationPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AuthorizationPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AuthorizationPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AuthorizationPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, parameters AuthorizationPolicyResourceFormat, options *AuthorizationPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if authorizationPolicyName == "" {
		return nil, errors.New("parameter authorizationPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationPolicyName}", url.PathEscape(authorizationPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AuthorizationPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (AuthorizationPoliciesClientCreateOrUpdateResponse, error) {
	result := AuthorizationPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationPolicyResourceFormat); err != nil {
		return AuthorizationPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets an authorization policy in the hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// authorizationPolicyName - The name of the policy.
// options - AuthorizationPoliciesClientGetOptions contains the optional parameters for the AuthorizationPoliciesClient.Get
// method.
func (client *AuthorizationPoliciesClient) Get(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, options *AuthorizationPoliciesClientGetOptions) (AuthorizationPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, authorizationPolicyName, options)
	if err != nil {
		return AuthorizationPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AuthorizationPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AuthorizationPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AuthorizationPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, options *AuthorizationPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if authorizationPolicyName == "" {
		return nil, errors.New("parameter authorizationPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationPolicyName}", url.PathEscape(authorizationPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AuthorizationPoliciesClient) getHandleResponse(resp *http.Response) (AuthorizationPoliciesClientGetResponse, error) {
	result := AuthorizationPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationPolicyResourceFormat); err != nil {
		return AuthorizationPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHubPager - Gets all the authorization policies in a specified hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// options - AuthorizationPoliciesClientListByHubOptions contains the optional parameters for the AuthorizationPoliciesClient.ListByHub
// method.
func (client *AuthorizationPoliciesClient) NewListByHubPager(resourceGroupName string, hubName string, options *AuthorizationPoliciesClientListByHubOptions) *runtime.Pager[AuthorizationPoliciesClientListByHubResponse] {
	return runtime.NewPager(runtime.PagingHandler[AuthorizationPoliciesClientListByHubResponse]{
		More: func(page AuthorizationPoliciesClientListByHubResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AuthorizationPoliciesClientListByHubResponse) (AuthorizationPoliciesClientListByHubResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHubCreateRequest(ctx, resourceGroupName, hubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AuthorizationPoliciesClientListByHubResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AuthorizationPoliciesClientListByHubResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AuthorizationPoliciesClientListByHubResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHubHandleResponse(resp)
		},
	})
}

// listByHubCreateRequest creates the ListByHub request.
func (client *AuthorizationPoliciesClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *AuthorizationPoliciesClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *AuthorizationPoliciesClient) listByHubHandleResponse(resp *http.Response) (AuthorizationPoliciesClientListByHubResponse, error) {
	result := AuthorizationPoliciesClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationPolicyListResult); err != nil {
		return AuthorizationPoliciesClientListByHubResponse{}, err
	}
	return result, nil
}

// RegeneratePrimaryKey - Regenerates the primary policy key of the specified authorization policy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// authorizationPolicyName - The name of the policy.
// options - AuthorizationPoliciesClientRegeneratePrimaryKeyOptions contains the optional parameters for the AuthorizationPoliciesClient.RegeneratePrimaryKey
// method.
func (client *AuthorizationPoliciesClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, options *AuthorizationPoliciesClientRegeneratePrimaryKeyOptions) (AuthorizationPoliciesClientRegeneratePrimaryKeyResponse, error) {
	req, err := client.regeneratePrimaryKeyCreateRequest(ctx, resourceGroupName, hubName, authorizationPolicyName, options)
	if err != nil {
		return AuthorizationPoliciesClientRegeneratePrimaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AuthorizationPoliciesClientRegeneratePrimaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AuthorizationPoliciesClientRegeneratePrimaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regeneratePrimaryKeyHandleResponse(resp)
}

// regeneratePrimaryKeyCreateRequest creates the RegeneratePrimaryKey request.
func (client *AuthorizationPoliciesClient) regeneratePrimaryKeyCreateRequest(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, options *AuthorizationPoliciesClientRegeneratePrimaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}/regeneratePrimaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if authorizationPolicyName == "" {
		return nil, errors.New("parameter authorizationPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationPolicyName}", url.PathEscape(authorizationPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// regeneratePrimaryKeyHandleResponse handles the RegeneratePrimaryKey response.
func (client *AuthorizationPoliciesClient) regeneratePrimaryKeyHandleResponse(resp *http.Response) (AuthorizationPoliciesClientRegeneratePrimaryKeyResponse, error) {
	result := AuthorizationPoliciesClientRegeneratePrimaryKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationPolicy); err != nil {
		return AuthorizationPoliciesClientRegeneratePrimaryKeyResponse{}, err
	}
	return result, nil
}

// RegenerateSecondaryKey - Regenerates the secondary policy key of the specified authorization policy.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// authorizationPolicyName - The name of the policy.
// options - AuthorizationPoliciesClientRegenerateSecondaryKeyOptions contains the optional parameters for the AuthorizationPoliciesClient.RegenerateSecondaryKey
// method.
func (client *AuthorizationPoliciesClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, options *AuthorizationPoliciesClientRegenerateSecondaryKeyOptions) (AuthorizationPoliciesClientRegenerateSecondaryKeyResponse, error) {
	req, err := client.regenerateSecondaryKeyCreateRequest(ctx, resourceGroupName, hubName, authorizationPolicyName, options)
	if err != nil {
		return AuthorizationPoliciesClientRegenerateSecondaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AuthorizationPoliciesClientRegenerateSecondaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AuthorizationPoliciesClientRegenerateSecondaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateSecondaryKeyHandleResponse(resp)
}

// regenerateSecondaryKeyCreateRequest creates the RegenerateSecondaryKey request.
func (client *AuthorizationPoliciesClient) regenerateSecondaryKeyCreateRequest(ctx context.Context, resourceGroupName string, hubName string, authorizationPolicyName string, options *AuthorizationPoliciesClientRegenerateSecondaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}/regenerateSecondaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if authorizationPolicyName == "" {
		return nil, errors.New("parameter authorizationPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationPolicyName}", url.PathEscape(authorizationPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// regenerateSecondaryKeyHandleResponse handles the RegenerateSecondaryKey response.
func (client *AuthorizationPoliciesClient) regenerateSecondaryKeyHandleResponse(resp *http.Response) (AuthorizationPoliciesClientRegenerateSecondaryKeyResponse, error) {
	result := AuthorizationPoliciesClientRegenerateSecondaryKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationPolicy); err != nil {
		return AuthorizationPoliciesClientRegenerateSecondaryKeyResponse{}, err
	}
	return result, nil
}

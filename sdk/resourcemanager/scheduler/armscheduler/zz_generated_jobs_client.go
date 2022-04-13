//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

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
	"strconv"
	"strings"
)

// JobsClient contains the methods for the Jobs group.
// Don't use this type directly, use NewJobsClient() instead.
type JobsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewJobsClient creates a new instance of JobsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &JobsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Provisions a new job or updates an existing job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// jobCollectionName - The job collection name.
// jobName - The job name.
// job - The job definition.
// options - JobsClientCreateOrUpdateOptions contains the optional parameters for the JobsClient.CreateOrUpdate method.
func (client *JobsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition, options *JobsClientCreateOrUpdateOptions) (JobsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, job, options)
	if err != nil {
		return JobsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return JobsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition, options *JobsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, job)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *JobsClient) createOrUpdateHandleResponse(resp *http.Response) (JobsClientCreateOrUpdateResponse, error) {
	result := JobsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// jobCollectionName - The job collection name.
// jobName - The job name.
// options - JobsClientDeleteOptions contains the optional parameters for the JobsClient.Delete method.
func (client *JobsClient) Delete(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsClientDeleteOptions) (JobsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, options)
	if err != nil {
		return JobsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return JobsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// jobCollectionName - The job collection name.
// jobName - The job name.
// options - JobsClientGetOptions contains the optional parameters for the JobsClient.Get method.
func (client *JobsClient) Get(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsClientGetOptions) (JobsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, options)
	if err != nil {
		return JobsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobsClient) getHandleResponse(resp *http.Response) (JobsClientGetResponse, error) {
	result := JobsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all jobs under the specified job collection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// jobCollectionName - The job collection name.
// options - JobsClientListOptions contains the optional parameters for the JobsClient.List method.
func (client *JobsClient) List(resourceGroupName string, jobCollectionName string, options *JobsClientListOptions) *runtime.Pager[JobsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[JobsClientListResponse]{
		More: func(page JobsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobsClientListResponse) (JobsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *JobsClient) listCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *JobsClient) listHandleResponse(resp *http.Response) (JobsClientListResponse, error) {
	result := JobsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobListResult); err != nil {
		return JobsClientListResponse{}, err
	}
	return result, nil
}

// ListJobHistory - Lists job history.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// jobCollectionName - The job collection name.
// jobName - The job name.
// options - JobsClientListJobHistoryOptions contains the optional parameters for the JobsClient.ListJobHistory method.
func (client *JobsClient) ListJobHistory(resourceGroupName string, jobCollectionName string, jobName string, options *JobsClientListJobHistoryOptions) *runtime.Pager[JobsClientListJobHistoryResponse] {
	return runtime.NewPager(runtime.PageProcessor[JobsClientListJobHistoryResponse]{
		More: func(page JobsClientListJobHistoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobsClientListJobHistoryResponse) (JobsClientListJobHistoryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listJobHistoryCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobsClientListJobHistoryResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return JobsClientListJobHistoryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobsClientListJobHistoryResponse{}, runtime.NewResponseError(resp)
			}
			return client.listJobHistoryHandleResponse(resp)
		},
	})
}

// listJobHistoryCreateRequest creates the ListJobHistory request.
func (client *JobsClient) listJobHistoryCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsClientListJobHistoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/history"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listJobHistoryHandleResponse handles the ListJobHistory response.
func (client *JobsClient) listJobHistoryHandleResponse(resp *http.Response) (JobsClientListJobHistoryResponse, error) {
	result := JobsClientListJobHistoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobHistoryListResult); err != nil {
		return JobsClientListJobHistoryResponse{}, err
	}
	return result, nil
}

// Patch - Patches an existing job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// jobCollectionName - The job collection name.
// jobName - The job name.
// job - The job definition.
// options - JobsClientPatchOptions contains the optional parameters for the JobsClient.Patch method.
func (client *JobsClient) Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition, options *JobsClientPatchOptions) (JobsClientPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, job, options)
	if err != nil {
		return JobsClientPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsClientPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsClientPatchResponse{}, runtime.NewResponseError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *JobsClient) patchCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job JobDefinition, options *JobsClientPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, job)
}

// patchHandleResponse handles the Patch response.
func (client *JobsClient) patchHandleResponse(resp *http.Response) (JobsClientPatchResponse, error) {
	result := JobsClientPatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobDefinition); err != nil {
		return JobsClientPatchResponse{}, err
	}
	return result, nil
}

// Run - Runs a job.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// jobCollectionName - The job collection name.
// jobName - The job name.
// options - JobsClientRunOptions contains the optional parameters for the JobsClient.Run method.
func (client *JobsClient) Run(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsClientRunOptions) (JobsClientRunResponse, error) {
	req, err := client.runCreateRequest(ctx, resourceGroupName, jobCollectionName, jobName, options)
	if err != nil {
		return JobsClientRunResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobsClientRunResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobsClientRunResponse{}, runtime.NewResponseError(resp)
	}
	return JobsClientRunResponse{}, nil
}

// runCreateRequest creates the Run request.
func (client *JobsClient) runCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, options *JobsClientRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/run"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ExtensionsListByFarmBeatsPager provides operations for iterating over paged responses.
type ExtensionsListByFarmBeatsPager struct {
	client    *ExtensionsClient
	current   ExtensionsListByFarmBeatsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ExtensionsListByFarmBeatsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ExtensionsListByFarmBeatsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ExtensionsListByFarmBeatsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ExtensionListResponse.NextLink == nil || len(*p.current.ExtensionListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByFarmBeatsHandleError(resp)
		return false
	}
	result, err := p.client.listByFarmBeatsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ExtensionsListByFarmBeatsResponse page.
func (p *ExtensionsListByFarmBeatsPager) PageResponse() ExtensionsListByFarmBeatsResponse {
	return p.current
}

// FarmBeatsExtensionsListPager provides operations for iterating over paged responses.
type FarmBeatsExtensionsListPager struct {
	client    *FarmBeatsExtensionsClient
	current   FarmBeatsExtensionsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FarmBeatsExtensionsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FarmBeatsExtensionsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FarmBeatsExtensionsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FarmBeatsExtensionListResponse.NextLink == nil || len(*p.current.FarmBeatsExtensionListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FarmBeatsExtensionsListResponse page.
func (p *FarmBeatsExtensionsListPager) PageResponse() FarmBeatsExtensionsListResponse {
	return p.current
}

// FarmBeatsModelsListByResourceGroupPager provides operations for iterating over paged responses.
type FarmBeatsModelsListByResourceGroupPager struct {
	client    *FarmBeatsModelsClient
	current   FarmBeatsModelsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FarmBeatsModelsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FarmBeatsModelsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FarmBeatsModelsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FarmBeatsListResponse.NextLink == nil || len(*p.current.FarmBeatsListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FarmBeatsModelsListByResourceGroupResponse page.
func (p *FarmBeatsModelsListByResourceGroupPager) PageResponse() FarmBeatsModelsListByResourceGroupResponse {
	return p.current
}

// FarmBeatsModelsListBySubscriptionPager provides operations for iterating over paged responses.
type FarmBeatsModelsListBySubscriptionPager struct {
	client    *FarmBeatsModelsClient
	current   FarmBeatsModelsListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, FarmBeatsModelsListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *FarmBeatsModelsListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *FarmBeatsModelsListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.FarmBeatsListResponse.NextLink == nil || len(*p.current.FarmBeatsListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current FarmBeatsModelsListBySubscriptionResponse page.
func (p *FarmBeatsModelsListBySubscriptionPager) PageResponse() FarmBeatsModelsListBySubscriptionResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}
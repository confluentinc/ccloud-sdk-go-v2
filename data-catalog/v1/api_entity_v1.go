// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Confluent Data Catalog

REST API for the Data Catalog

API version: 1.0.0
Contact: data-governance@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type EntityV1Api interface {

	/*
		CreateTags Bulk Create Tags

		[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

	Bulk API to create multiple tags.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiCreateTagsRequest
	*/
	CreateTags(ctx _context.Context) ApiCreateTagsRequest

	// CreateTagsExecute executes the request
	//  @return []TagResponse
	CreateTagsExecute(r ApiCreateTagsRequest) ([]TagResponse, *_nethttp.Response, error)

	/*
		DeleteTag Delete a Tag for an Entity

		[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

	Delete a tag for an entity.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param typeName The type of the entity
		 @param qualifiedName The qualified name of the entity
		 @param tagName The name of the tag
		 @return ApiDeleteTagRequest
	*/
	DeleteTag(ctx _context.Context, typeName string, qualifiedName string, tagName string) ApiDeleteTagRequest

	// DeleteTagExecute executes the request
	DeleteTagExecute(r ApiDeleteTagRequest) (*_nethttp.Response, error)

	/*
		GetByUniqueAttributes Read an Entity

		[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

	Fetch complete definition of an entity given its type and unique attribute.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param typeName The type of the entity
		 @param qualifiedName The qualified name of the entity
		 @return ApiGetByUniqueAttributesRequest
	*/
	GetByUniqueAttributes(ctx _context.Context, typeName string, qualifiedName string) ApiGetByUniqueAttributesRequest

	// GetByUniqueAttributesExecute executes the request
	//  @return EntityWithExtInfo
	GetByUniqueAttributesExecute(r ApiGetByUniqueAttributesRequest) (EntityWithExtInfo, *_nethttp.Response, error)

	/*
		GetTags Read Tags for an Entity

		[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

	Gets the list of tags for a given entity represented by a qualified name.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param typeName The type of the entity
		 @param qualifiedName The qualified name of the entity
		 @return ApiGetTagsRequest
	*/
	GetTags(ctx _context.Context, typeName string, qualifiedName string) ApiGetTagsRequest

	// GetTagsExecute executes the request
	//  @return []TagResponse
	GetTagsExecute(r ApiGetTagsRequest) ([]TagResponse, *_nethttp.Response, error)

	/*
		UpdateTags Bulk Update Tags

		[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

	Bulk API to update multiple tags.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiUpdateTagsRequest
	*/
	UpdateTags(ctx _context.Context) ApiUpdateTagsRequest

	// UpdateTagsExecute executes the request
	//  @return []TagResponse
	UpdateTagsExecute(r ApiUpdateTagsRequest) ([]TagResponse, *_nethttp.Response, error)
}

// EntityV1ApiService EntityV1Api service
type EntityV1ApiService service

type ApiCreateTagsRequest struct {
	ctx        _context.Context
	ApiService EntityV1Api
	tag        *[]Tag
}

// The tags
func (r ApiCreateTagsRequest) Tag(tag []Tag) ApiCreateTagsRequest {
	r.tag = &tag
	return r
}

func (r ApiCreateTagsRequest) Execute() ([]TagResponse, *_nethttp.Response, error) {
	return r.ApiService.CreateTagsExecute(r)
}

/*
CreateTags Bulk Create Tags

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Bulk API to create multiple tags.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTagsRequest
*/
func (a *EntityV1ApiService) CreateTags(ctx _context.Context) ApiCreateTagsRequest {
	return ApiCreateTagsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TagResponse
func (a *EntityV1ApiService) CreateTagsExecute(r ApiCreateTagsRequest) ([]TagResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TagResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntityV1ApiService.CreateTags")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/catalog/v1/entity/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tag
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteTagRequest struct {
	ctx           _context.Context
	ApiService    EntityV1Api
	typeName      string
	qualifiedName string
	tagName       string
}

func (r ApiDeleteTagRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteTagExecute(r)
}

/*
DeleteTag Delete a Tag for an Entity

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Delete a tag for an entity.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param typeName The type of the entity
 @param qualifiedName The qualified name of the entity
 @param tagName The name of the tag
 @return ApiDeleteTagRequest
*/
func (a *EntityV1ApiService) DeleteTag(ctx _context.Context, typeName string, qualifiedName string, tagName string) ApiDeleteTagRequest {
	return ApiDeleteTagRequest{
		ApiService:    a,
		ctx:           ctx,
		typeName:      typeName,
		qualifiedName: qualifiedName,
		tagName:       tagName,
	}
}

// Execute executes the request
func (a *EntityV1ApiService) DeleteTagExecute(r ApiDeleteTagRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntityV1ApiService.DeleteTag")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags/{tagName}"
	localVarPath = strings.Replace(localVarPath, "{"+"typeName"+"}", _neturl.PathEscape(parameterToString(r.typeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"qualifiedName"+"}", _neturl.PathEscape(parameterToString(r.qualifiedName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tagName"+"}", _neturl.PathEscape(parameterToString(r.tagName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetByUniqueAttributesRequest struct {
	ctx                 _context.Context
	ApiService          EntityV1Api
	typeName            string
	qualifiedName       string
	minExtInfo          *bool
	ignoreRelationships *bool
}

// Whether to populate on header and schema attributes
func (r ApiGetByUniqueAttributesRequest) MinExtInfo(minExtInfo bool) ApiGetByUniqueAttributesRequest {
	r.minExtInfo = &minExtInfo
	return r
}

// Whether to ignore relationships
func (r ApiGetByUniqueAttributesRequest) IgnoreRelationships(ignoreRelationships bool) ApiGetByUniqueAttributesRequest {
	r.ignoreRelationships = &ignoreRelationships
	return r
}

func (r ApiGetByUniqueAttributesRequest) Execute() (EntityWithExtInfo, *_nethttp.Response, error) {
	return r.ApiService.GetByUniqueAttributesExecute(r)
}

/*
GetByUniqueAttributes Read an Entity

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Fetch complete definition of an entity given its type and unique attribute.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param typeName The type of the entity
 @param qualifiedName The qualified name of the entity
 @return ApiGetByUniqueAttributesRequest
*/
func (a *EntityV1ApiService) GetByUniqueAttributes(ctx _context.Context, typeName string, qualifiedName string) ApiGetByUniqueAttributesRequest {
	return ApiGetByUniqueAttributesRequest{
		ApiService:    a,
		ctx:           ctx,
		typeName:      typeName,
		qualifiedName: qualifiedName,
	}
}

// Execute executes the request
//  @return EntityWithExtInfo
func (a *EntityV1ApiService) GetByUniqueAttributesExecute(r ApiGetByUniqueAttributesRequest) (EntityWithExtInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EntityWithExtInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntityV1ApiService.GetByUniqueAttributes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/catalog/v1/entity/type/{typeName}/name/{qualifiedName}"
	localVarPath = strings.Replace(localVarPath, "{"+"typeName"+"}", _neturl.PathEscape(parameterToString(r.typeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"qualifiedName"+"}", _neturl.PathEscape(parameterToString(r.qualifiedName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.minExtInfo != nil {
		localVarQueryParams.Add("minExtInfo", parameterToString(*r.minExtInfo, ""))
	}
	if r.ignoreRelationships != nil {
		localVarQueryParams.Add("ignoreRelationships", parameterToString(*r.ignoreRelationships, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTagsRequest struct {
	ctx           _context.Context
	ApiService    EntityV1Api
	typeName      string
	qualifiedName string
}

func (r ApiGetTagsRequest) Execute() ([]TagResponse, *_nethttp.Response, error) {
	return r.ApiService.GetTagsExecute(r)
}

/*
GetTags Read Tags for an Entity

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Gets the list of tags for a given entity represented by a qualified name.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param typeName The type of the entity
 @param qualifiedName The qualified name of the entity
 @return ApiGetTagsRequest
*/
func (a *EntityV1ApiService) GetTags(ctx _context.Context, typeName string, qualifiedName string) ApiGetTagsRequest {
	return ApiGetTagsRequest{
		ApiService:    a,
		ctx:           ctx,
		typeName:      typeName,
		qualifiedName: qualifiedName,
	}
}

// Execute executes the request
//  @return []TagResponse
func (a *EntityV1ApiService) GetTagsExecute(r ApiGetTagsRequest) ([]TagResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TagResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntityV1ApiService.GetTags")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"typeName"+"}", _neturl.PathEscape(parameterToString(r.typeName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"qualifiedName"+"}", _neturl.PathEscape(parameterToString(r.qualifiedName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateTagsRequest struct {
	ctx        _context.Context
	ApiService EntityV1Api
	tag        *[]Tag
}

// The tags
func (r ApiUpdateTagsRequest) Tag(tag []Tag) ApiUpdateTagsRequest {
	r.tag = &tag
	return r
}

func (r ApiUpdateTagsRequest) Execute() ([]TagResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateTagsExecute(r)
}

/*
UpdateTags Bulk Update Tags

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Bulk API to update multiple tags.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateTagsRequest
*/
func (a *EntityV1ApiService) UpdateTags(ctx _context.Context) ApiUpdateTagsRequest {
	return ApiUpdateTagsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TagResponse
func (a *EntityV1ApiService) UpdateTagsExecute(r ApiUpdateTagsRequest) ([]TagResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TagResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntityV1ApiService.UpdateTags")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/catalog/v1/entity/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tag
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
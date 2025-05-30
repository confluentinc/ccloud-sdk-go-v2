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
Custom Connect Plugin Management API

This is Custom Connect Plugin Management API.

API version: 0.1.0
Contact: connect-team@confluent.io
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

type CustomConnectPluginVersionsCcpmV1Api interface {

	/*
			CreateCcpmV1CustomConnectPluginVersion Create a Custom Connect Plugin Version

			[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to create a custom connect plugin version.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param pluginId The Plugin
			 @return ApiCreateCcpmV1CustomConnectPluginVersionRequest
	*/
	CreateCcpmV1CustomConnectPluginVersion(ctx _context.Context, pluginId string) ApiCreateCcpmV1CustomConnectPluginVersionRequest

	// CreateCcpmV1CustomConnectPluginVersionExecute executes the request
	//  @return CcpmV1CustomConnectPluginVersion
	CreateCcpmV1CustomConnectPluginVersionExecute(r ApiCreateCcpmV1CustomConnectPluginVersionRequest) (CcpmV1CustomConnectPluginVersion, *_nethttp.Response, error)

	/*
			DeleteCcpmV1CustomConnectPluginVersion Delete a Custom Connect Plugin Version

			[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to delete a custom connect plugin version.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param pluginId The Plugin
			 @param id The unique identifier for the custom connect plugin version.
			 @return ApiDeleteCcpmV1CustomConnectPluginVersionRequest
	*/
	DeleteCcpmV1CustomConnectPluginVersion(ctx _context.Context, pluginId string, id string) ApiDeleteCcpmV1CustomConnectPluginVersionRequest

	// DeleteCcpmV1CustomConnectPluginVersionExecute executes the request
	DeleteCcpmV1CustomConnectPluginVersionExecute(r ApiDeleteCcpmV1CustomConnectPluginVersionRequest) (*_nethttp.Response, error)

	/*
			GetCcpmV1CustomConnectPluginVersion Read a Custom Connect Plugin Version

			[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to read a custom connect plugin version.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param pluginId The Plugin
			 @param id The unique identifier for the custom connect plugin version.
			 @return ApiGetCcpmV1CustomConnectPluginVersionRequest
	*/
	GetCcpmV1CustomConnectPluginVersion(ctx _context.Context, pluginId string, id string) ApiGetCcpmV1CustomConnectPluginVersionRequest

	// GetCcpmV1CustomConnectPluginVersionExecute executes the request
	//  @return CcpmV1CustomConnectPluginVersion
	GetCcpmV1CustomConnectPluginVersionExecute(r ApiGetCcpmV1CustomConnectPluginVersionRequest) (CcpmV1CustomConnectPluginVersion, *_nethttp.Response, error)

	/*
			ListCcpmV1CustomConnectPluginVersions List of Custom Connect Plugin Versions

			[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Retrieve a sorted, filtered, paginated list of all custom connect plugin versions.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param pluginId The Plugin
			 @return ApiListCcpmV1CustomConnectPluginVersionsRequest
	*/
	ListCcpmV1CustomConnectPluginVersions(ctx _context.Context, pluginId string) ApiListCcpmV1CustomConnectPluginVersionsRequest

	// ListCcpmV1CustomConnectPluginVersionsExecute executes the request
	//  @return CcpmV1CustomConnectPluginVersionList
	ListCcpmV1CustomConnectPluginVersionsExecute(r ApiListCcpmV1CustomConnectPluginVersionsRequest) (CcpmV1CustomConnectPluginVersionList, *_nethttp.Response, error)
}

// CustomConnectPluginVersionsCcpmV1ApiService CustomConnectPluginVersionsCcpmV1Api service
type CustomConnectPluginVersionsCcpmV1ApiService service

type ApiCreateCcpmV1CustomConnectPluginVersionRequest struct {
	ctx                              _context.Context
	ApiService                       CustomConnectPluginVersionsCcpmV1Api
	pluginId                         string
	ccpmV1CustomConnectPluginVersion *CcpmV1CustomConnectPluginVersion
}

func (r ApiCreateCcpmV1CustomConnectPluginVersionRequest) CcpmV1CustomConnectPluginVersion(ccpmV1CustomConnectPluginVersion CcpmV1CustomConnectPluginVersion) ApiCreateCcpmV1CustomConnectPluginVersionRequest {
	r.ccpmV1CustomConnectPluginVersion = &ccpmV1CustomConnectPluginVersion
	return r
}

func (r ApiCreateCcpmV1CustomConnectPluginVersionRequest) Execute() (CcpmV1CustomConnectPluginVersion, *_nethttp.Response, error) {
	return r.ApiService.CreateCcpmV1CustomConnectPluginVersionExecute(r)
}

/*
CreateCcpmV1CustomConnectPluginVersion Create a Custom Connect Plugin Version

[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to create a custom connect plugin version.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pluginId The Plugin
	@return ApiCreateCcpmV1CustomConnectPluginVersionRequest
*/
func (a *CustomConnectPluginVersionsCcpmV1ApiService) CreateCcpmV1CustomConnectPluginVersion(ctx _context.Context, pluginId string) ApiCreateCcpmV1CustomConnectPluginVersionRequest {
	return ApiCreateCcpmV1CustomConnectPluginVersionRequest{
		ApiService: a,
		ctx:        ctx,
		pluginId:   pluginId,
	}
}

// Execute executes the request
//
//	@return CcpmV1CustomConnectPluginVersion
func (a *CustomConnectPluginVersionsCcpmV1ApiService) CreateCcpmV1CustomConnectPluginVersionExecute(r ApiCreateCcpmV1CustomConnectPluginVersionRequest) (CcpmV1CustomConnectPluginVersion, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CcpmV1CustomConnectPluginVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomConnectPluginVersionsCcpmV1ApiService.CreateCcpmV1CustomConnectPluginVersion")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ccpm/v1/plugins/{plugin_id}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"plugin_id"+"}", _neturl.PathEscape(parameterToString(r.pluginId, "")), -1)

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
	localVarPostBody = r.ccpmV1CustomConnectPluginVersion
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiDeleteCcpmV1CustomConnectPluginVersionRequest struct {
	ctx         _context.Context
	ApiService  CustomConnectPluginVersionsCcpmV1Api
	environment *string
	pluginId    string
	id          string
}

// Scope the operation to the given environment.
func (r ApiDeleteCcpmV1CustomConnectPluginVersionRequest) Environment(environment string) ApiDeleteCcpmV1CustomConnectPluginVersionRequest {
	r.environment = &environment
	return r
}

func (r ApiDeleteCcpmV1CustomConnectPluginVersionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteCcpmV1CustomConnectPluginVersionExecute(r)
}

/*
DeleteCcpmV1CustomConnectPluginVersion Delete a Custom Connect Plugin Version

[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to delete a custom connect plugin version.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pluginId The Plugin
	@param id The unique identifier for the custom connect plugin version.
	@return ApiDeleteCcpmV1CustomConnectPluginVersionRequest
*/
func (a *CustomConnectPluginVersionsCcpmV1ApiService) DeleteCcpmV1CustomConnectPluginVersion(ctx _context.Context, pluginId string, id string) ApiDeleteCcpmV1CustomConnectPluginVersionRequest {
	return ApiDeleteCcpmV1CustomConnectPluginVersionRequest{
		ApiService: a,
		ctx:        ctx,
		pluginId:   pluginId,
		id:         id,
	}
}

// Execute executes the request
func (a *CustomConnectPluginVersionsCcpmV1ApiService) DeleteCcpmV1CustomConnectPluginVersionExecute(r ApiDeleteCcpmV1CustomConnectPluginVersionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomConnectPluginVersionsCcpmV1ApiService.DeleteCcpmV1CustomConnectPluginVersion")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ccpm/v1/plugins/{plugin_id}/versions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"plugin_id"+"}", _neturl.PathEscape(parameterToString(r.pluginId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environment == nil {
		return nil, reportError("environment is required and must be specified")
	}

	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCcpmV1CustomConnectPluginVersionRequest struct {
	ctx         _context.Context
	ApiService  CustomConnectPluginVersionsCcpmV1Api
	environment *string
	pluginId    string
	id          string
}

// Scope the operation to the given environment.
func (r ApiGetCcpmV1CustomConnectPluginVersionRequest) Environment(environment string) ApiGetCcpmV1CustomConnectPluginVersionRequest {
	r.environment = &environment
	return r
}

func (r ApiGetCcpmV1CustomConnectPluginVersionRequest) Execute() (CcpmV1CustomConnectPluginVersion, *_nethttp.Response, error) {
	return r.ApiService.GetCcpmV1CustomConnectPluginVersionExecute(r)
}

/*
GetCcpmV1CustomConnectPluginVersion Read a Custom Connect Plugin Version

[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to read a custom connect plugin version.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pluginId The Plugin
	@param id The unique identifier for the custom connect plugin version.
	@return ApiGetCcpmV1CustomConnectPluginVersionRequest
*/
func (a *CustomConnectPluginVersionsCcpmV1ApiService) GetCcpmV1CustomConnectPluginVersion(ctx _context.Context, pluginId string, id string) ApiGetCcpmV1CustomConnectPluginVersionRequest {
	return ApiGetCcpmV1CustomConnectPluginVersionRequest{
		ApiService: a,
		ctx:        ctx,
		pluginId:   pluginId,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CcpmV1CustomConnectPluginVersion
func (a *CustomConnectPluginVersionsCcpmV1ApiService) GetCcpmV1CustomConnectPluginVersionExecute(r ApiGetCcpmV1CustomConnectPluginVersionRequest) (CcpmV1CustomConnectPluginVersion, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CcpmV1CustomConnectPluginVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomConnectPluginVersionsCcpmV1ApiService.GetCcpmV1CustomConnectPluginVersion")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ccpm/v1/plugins/{plugin_id}/versions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"plugin_id"+"}", _neturl.PathEscape(parameterToString(r.pluginId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environment == nil {
		return localVarReturnValue, nil, reportError("environment is required and must be specified")
	}

	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiListCcpmV1CustomConnectPluginVersionsRequest struct {
	ctx         _context.Context
	ApiService  CustomConnectPluginVersionsCcpmV1Api
	environment *string
	pluginId    string
}

// Filter the results by exact match for environment.
func (r ApiListCcpmV1CustomConnectPluginVersionsRequest) Environment(environment string) ApiListCcpmV1CustomConnectPluginVersionsRequest {
	r.environment = &environment
	return r
}

func (r ApiListCcpmV1CustomConnectPluginVersionsRequest) Execute() (CcpmV1CustomConnectPluginVersionList, *_nethttp.Response, error) {
	return r.ApiService.ListCcpmV1CustomConnectPluginVersionsExecute(r)
}

/*
ListCcpmV1CustomConnectPluginVersions List of Custom Connect Plugin Versions

[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Retrieve a sorted, filtered, paginated list of all custom connect plugin versions.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param pluginId The Plugin
	@return ApiListCcpmV1CustomConnectPluginVersionsRequest
*/
func (a *CustomConnectPluginVersionsCcpmV1ApiService) ListCcpmV1CustomConnectPluginVersions(ctx _context.Context, pluginId string) ApiListCcpmV1CustomConnectPluginVersionsRequest {
	return ApiListCcpmV1CustomConnectPluginVersionsRequest{
		ApiService: a,
		ctx:        ctx,
		pluginId:   pluginId,
	}
}

// Execute executes the request
//
//	@return CcpmV1CustomConnectPluginVersionList
func (a *CustomConnectPluginVersionsCcpmV1ApiService) ListCcpmV1CustomConnectPluginVersionsExecute(r ApiListCcpmV1CustomConnectPluginVersionsRequest) (CcpmV1CustomConnectPluginVersionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CcpmV1CustomConnectPluginVersionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomConnectPluginVersionsCcpmV1ApiService.ListCcpmV1CustomConnectPluginVersions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ccpm/v1/plugins/{plugin_id}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"plugin_id"+"}", _neturl.PathEscape(parameterToString(r.pluginId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environment == nil {
		return localVarReturnValue, nil, reportError("environment is required and must be specified")
	}

	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

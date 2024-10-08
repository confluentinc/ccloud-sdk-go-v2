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
SQL API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
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

type ConnectionsSqlV1Api interface {

	/*
		CreateSqlv1Connection Create a Connection

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Make a request to create a Connection.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @return ApiCreateSqlv1ConnectionRequest
	*/
	CreateSqlv1Connection(ctx _context.Context, organizationId string, environmentId string) ApiCreateSqlv1ConnectionRequest

	// CreateSqlv1ConnectionExecute executes the request
	//  @return SqlV1Connection
	CreateSqlv1ConnectionExecute(r ApiCreateSqlv1ConnectionRequest) (SqlV1Connection, *_nethttp.Response, error)

	/*
		DeleteSqlv1Connection Delete a Connection

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Make a request to delete a statement.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @param connectionName The unique identifier for the connection.
		 @return ApiDeleteSqlv1ConnectionRequest
	*/
	DeleteSqlv1Connection(ctx _context.Context, organizationId string, environmentId string, connectionName string) ApiDeleteSqlv1ConnectionRequest

	// DeleteSqlv1ConnectionExecute executes the request
	DeleteSqlv1ConnectionExecute(r ApiDeleteSqlv1ConnectionRequest) (*_nethttp.Response, error)

	/*
		GetSqlv1Connection Read a Connection

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Make a request to read a Connection.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @param connectionName The user provided name of the Connection. Unique within a region within an org and env.
		 @return ApiGetSqlv1ConnectionRequest
	*/
	GetSqlv1Connection(ctx _context.Context, organizationId string, environmentId string, connectionName string) ApiGetSqlv1ConnectionRequest

	// GetSqlv1ConnectionExecute executes the request
	//  @return SqlV1Connection
	GetSqlv1ConnectionExecute(r ApiGetSqlv1ConnectionRequest) (SqlV1Connection, *_nethttp.Response, error)

	/*
		ListSqlv1Connections List of Connections

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Retrieve a sorted, filtered and paginated list of all Connections.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @return ApiListSqlv1ConnectionsRequest
	*/
	ListSqlv1Connections(ctx _context.Context, organizationId string, environmentId string) ApiListSqlv1ConnectionsRequest

	// ListSqlv1ConnectionsExecute executes the request
	//  @return SqlV1ConnectionList
	ListSqlv1ConnectionsExecute(r ApiListSqlv1ConnectionsRequest) (SqlV1ConnectionList, *_nethttp.Response, error)

	/*
		UpdateSqlv1Connection Update a Connection

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Make a request to update a connection.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @param connectionName The unique identifier for the connection.
		 @return ApiUpdateSqlv1ConnectionRequest
	*/
	UpdateSqlv1Connection(ctx _context.Context, organizationId string, environmentId string, connectionName string) ApiUpdateSqlv1ConnectionRequest

	// UpdateSqlv1ConnectionExecute executes the request
	UpdateSqlv1ConnectionExecute(r ApiUpdateSqlv1ConnectionRequest) (*_nethttp.Response, error)
}

// ConnectionsSqlV1ApiService ConnectionsSqlV1Api service
type ConnectionsSqlV1ApiService service

type ApiCreateSqlv1ConnectionRequest struct {
	ctx             _context.Context
	ApiService      ConnectionsSqlV1Api
	organizationId  string
	environmentId   string
	sqlV1Connection *SqlV1Connection
}

func (r ApiCreateSqlv1ConnectionRequest) SqlV1Connection(sqlV1Connection SqlV1Connection) ApiCreateSqlv1ConnectionRequest {
	r.sqlV1Connection = &sqlV1Connection
	return r
}

func (r ApiCreateSqlv1ConnectionRequest) Execute() (SqlV1Connection, *_nethttp.Response, error) {
	return r.ApiService.CreateSqlv1ConnectionExecute(r)
}

/*
CreateSqlv1Connection Create a Connection

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to create a Connection.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@return ApiCreateSqlv1ConnectionRequest
*/
func (a *ConnectionsSqlV1ApiService) CreateSqlv1Connection(ctx _context.Context, organizationId string, environmentId string) ApiCreateSqlv1ConnectionRequest {
	return ApiCreateSqlv1ConnectionRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
	}
}

// Execute executes the request
//
//	@return SqlV1Connection
func (a *ConnectionsSqlV1ApiService) CreateSqlv1ConnectionExecute(r ApiCreateSqlv1ConnectionRequest) (SqlV1Connection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1Connection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsSqlV1ApiService.CreateSqlv1Connection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/connections"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

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
	localVarPostBody = r.sqlV1Connection
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

type ApiDeleteSqlv1ConnectionRequest struct {
	ctx            _context.Context
	ApiService     ConnectionsSqlV1Api
	organizationId string
	environmentId  string
	connectionName string
}

func (r ApiDeleteSqlv1ConnectionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteSqlv1ConnectionExecute(r)
}

/*
DeleteSqlv1Connection Delete a Connection

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to delete a statement.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@param connectionName The unique identifier for the connection.
	@return ApiDeleteSqlv1ConnectionRequest
*/
func (a *ConnectionsSqlV1ApiService) DeleteSqlv1Connection(ctx _context.Context, organizationId string, environmentId string, connectionName string) ApiDeleteSqlv1ConnectionRequest {
	return ApiDeleteSqlv1ConnectionRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
		connectionName: connectionName,
	}
}

// Execute executes the request
func (a *ConnectionsSqlV1ApiService) DeleteSqlv1ConnectionExecute(r ApiDeleteSqlv1ConnectionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsSqlV1ApiService.DeleteSqlv1Connection")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/connections/{connection_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connection_name"+"}", _neturl.PathEscape(parameterToString(r.connectionName, "")), -1)

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

type ApiGetSqlv1ConnectionRequest struct {
	ctx            _context.Context
	ApiService     ConnectionsSqlV1Api
	organizationId string
	environmentId  string
	connectionName string
}

func (r ApiGetSqlv1ConnectionRequest) Execute() (SqlV1Connection, *_nethttp.Response, error) {
	return r.ApiService.GetSqlv1ConnectionExecute(r)
}

/*
GetSqlv1Connection Read a Connection

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to read a Connection.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@param connectionName The user provided name of the Connection. Unique within a region within an org and env.
	@return ApiGetSqlv1ConnectionRequest
*/
func (a *ConnectionsSqlV1ApiService) GetSqlv1Connection(ctx _context.Context, organizationId string, environmentId string, connectionName string) ApiGetSqlv1ConnectionRequest {
	return ApiGetSqlv1ConnectionRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
		connectionName: connectionName,
	}
}

// Execute executes the request
//
//	@return SqlV1Connection
func (a *ConnectionsSqlV1ApiService) GetSqlv1ConnectionExecute(r ApiGetSqlv1ConnectionRequest) (SqlV1Connection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1Connection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsSqlV1ApiService.GetSqlv1Connection")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/connections/{connection_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connection_name"+"}", _neturl.PathEscape(parameterToString(r.connectionName, "")), -1)

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

type ApiListSqlv1ConnectionsRequest struct {
	ctx                _context.Context
	ApiService         ConnectionsSqlV1Api
	organizationId     string
	environmentId      string
	specConnectionType *string
	pageSize           *int32
	pageToken          *string
}

// Filter the results by exact match for spec.connection_type
func (r ApiListSqlv1ConnectionsRequest) SpecConnectionType(specConnectionType string) ApiListSqlv1ConnectionsRequest {
	r.specConnectionType = &specConnectionType
	return r
}

// A pagination size for collection requests.
func (r ApiListSqlv1ConnectionsRequest) PageSize(pageSize int32) ApiListSqlv1ConnectionsRequest {
	r.pageSize = &pageSize
	return r
}

// An opaque pagination token for collection requests.
func (r ApiListSqlv1ConnectionsRequest) PageToken(pageToken string) ApiListSqlv1ConnectionsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListSqlv1ConnectionsRequest) Execute() (SqlV1ConnectionList, *_nethttp.Response, error) {
	return r.ApiService.ListSqlv1ConnectionsExecute(r)
}

/*
ListSqlv1Connections List of Connections

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Retrieve a sorted, filtered and paginated list of all Connections.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@return ApiListSqlv1ConnectionsRequest
*/
func (a *ConnectionsSqlV1ApiService) ListSqlv1Connections(ctx _context.Context, organizationId string, environmentId string) ApiListSqlv1ConnectionsRequest {
	return ApiListSqlv1ConnectionsRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
	}
}

// Execute executes the request
//
//	@return SqlV1ConnectionList
func (a *ConnectionsSqlV1ApiService) ListSqlv1ConnectionsExecute(r ApiListSqlv1ConnectionsRequest) (SqlV1ConnectionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1ConnectionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsSqlV1ApiService.ListSqlv1Connections")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/connections"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.specConnectionType != nil {
		localVarQueryParams.Add("spec.connection_type", parameterToString(*r.specConnectionType, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
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

type ApiUpdateSqlv1ConnectionRequest struct {
	ctx             _context.Context
	ApiService      ConnectionsSqlV1Api
	organizationId  string
	environmentId   string
	connectionName  string
	sqlV1Connection *SqlV1Connection
}

func (r ApiUpdateSqlv1ConnectionRequest) SqlV1Connection(sqlV1Connection SqlV1Connection) ApiUpdateSqlv1ConnectionRequest {
	r.sqlV1Connection = &sqlV1Connection
	return r
}

func (r ApiUpdateSqlv1ConnectionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateSqlv1ConnectionExecute(r)
}

/*
UpdateSqlv1Connection Update a Connection

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to update a connection.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@param connectionName The unique identifier for the connection.
	@return ApiUpdateSqlv1ConnectionRequest
*/
func (a *ConnectionsSqlV1ApiService) UpdateSqlv1Connection(ctx _context.Context, organizationId string, environmentId string, connectionName string) ApiUpdateSqlv1ConnectionRequest {
	return ApiUpdateSqlv1ConnectionRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
		connectionName: connectionName,
	}
}

// Execute executes the request
func (a *ConnectionsSqlV1ApiService) UpdateSqlv1ConnectionExecute(r ApiUpdateSqlv1ConnectionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectionsSqlV1ApiService.UpdateSqlv1Connection")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/connections/{connection_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"connection_name"+"}", _neturl.PathEscape(parameterToString(r.connectionName, "")), -1)

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
	localVarPostBody = r.sqlV1Connection
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

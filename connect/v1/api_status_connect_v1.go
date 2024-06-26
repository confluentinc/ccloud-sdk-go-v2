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
Kafka Connect APIs

REST API for managing connectors

API version: 1.0
Contact: connect@confluent.io
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

type StatusConnectV1Api interface {

	/*
		ListConnectv1ConnectorTasks List of Connector Tasks

		[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

	Get a list of tasks currently running for the connector.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param connectorName The unique name of the connector.
		 @param environmentId The unique identifier of the environment this resource belongs to.
		 @param kafkaClusterId The unique identifier for the Kafka cluster.
		 @return ApiListConnectv1ConnectorTasksRequest
	*/
	ListConnectv1ConnectorTasks(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiListConnectv1ConnectorTasksRequest

	// ListConnectv1ConnectorTasksExecute executes the request
	//  @return []map[string]interface{}
	ListConnectv1ConnectorTasksExecute(r ApiListConnectv1ConnectorTasksRequest) ([]map[string]interface{}, *_nethttp.Response, error)

	/*
		ReadConnectv1ConnectorStatus Read a Connector Status

		Get current status of the connector. This includes whether it is running, failed, or paused. Also includes which worker it is assigned to, error information if it has failed, and the state of all its tasks.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param connectorName The unique name of the connector.
		 @param environmentId The unique identifier of the environment this resource belongs to.
		 @param kafkaClusterId The unique identifier for the Kafka cluster.
		 @return ApiReadConnectv1ConnectorStatusRequest
	*/
	ReadConnectv1ConnectorStatus(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiReadConnectv1ConnectorStatusRequest

	// ReadConnectv1ConnectorStatusExecute executes the request
	//  @return InlineResponse2001
	ReadConnectv1ConnectorStatusExecute(r ApiReadConnectv1ConnectorStatusRequest) (InlineResponse2001, *_nethttp.Response, error)
}

// StatusConnectV1ApiService StatusConnectV1Api service
type StatusConnectV1ApiService service

type ApiListConnectv1ConnectorTasksRequest struct {
	ctx            _context.Context
	ApiService     StatusConnectV1Api
	connectorName  string
	environmentId  string
	kafkaClusterId string
}

func (r ApiListConnectv1ConnectorTasksRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ListConnectv1ConnectorTasksExecute(r)
}

/*
ListConnectv1ConnectorTasks List of Connector Tasks

[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Get a list of tasks currently running for the connector.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connectorName The unique name of the connector.
	@param environmentId The unique identifier of the environment this resource belongs to.
	@param kafkaClusterId The unique identifier for the Kafka cluster.
	@return ApiListConnectv1ConnectorTasksRequest
*/
func (a *StatusConnectV1ApiService) ListConnectv1ConnectorTasks(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiListConnectv1ConnectorTasksRequest {
	return ApiListConnectv1ConnectorTasksRequest{
		ApiService:     a,
		ctx:            ctx,
		connectorName:  connectorName,
		environmentId:  environmentId,
		kafkaClusterId: kafkaClusterId,
	}
}

// Execute executes the request
//
//	@return []map[string]interface{}
func (a *StatusConnectV1ApiService) ListConnectv1ConnectorTasksExecute(r ApiListConnectv1ConnectorTasksRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusConnectV1ApiService.ListConnectv1ConnectorTasks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/tasks"
	localVarPath = strings.Replace(localVarPath, "{"+"connector_name"+"}", _neturl.PathEscape(parameterToString(r.connectorName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kafka_cluster_id"+"}", _neturl.PathEscape(parameterToString(r.kafkaClusterId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ConnectV1ConnectorError
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

type ApiReadConnectv1ConnectorStatusRequest struct {
	ctx            _context.Context
	ApiService     StatusConnectV1Api
	connectorName  string
	environmentId  string
	kafkaClusterId string
}

func (r ApiReadConnectv1ConnectorStatusRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.ReadConnectv1ConnectorStatusExecute(r)
}

/*
ReadConnectv1ConnectorStatus Read a Connector Status

Get current status of the connector. This includes whether it is running, failed, or paused. Also includes which worker it is assigned to, error information if it has failed, and the state of all its tasks.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connectorName The unique name of the connector.
	@param environmentId The unique identifier of the environment this resource belongs to.
	@param kafkaClusterId The unique identifier for the Kafka cluster.
	@return ApiReadConnectv1ConnectorStatusRequest
*/
func (a *StatusConnectV1ApiService) ReadConnectv1ConnectorStatus(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiReadConnectv1ConnectorStatusRequest {
	return ApiReadConnectv1ConnectorStatusRequest{
		ApiService:     a,
		ctx:            ctx,
		connectorName:  connectorName,
		environmentId:  environmentId,
		kafkaClusterId: kafkaClusterId,
	}
}

// Execute executes the request
//
//	@return InlineResponse2001
func (a *StatusConnectV1ApiService) ReadConnectv1ConnectorStatusExecute(r ApiReadConnectv1ConnectorStatusRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusConnectV1ApiService.ReadConnectv1ConnectorStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"connector_name"+"}", _neturl.PathEscape(parameterToString(r.connectorName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kafka_cluster_id"+"}", _neturl.PathEscape(parameterToString(r.kafkaClusterId, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ConnectV1ConnectorError
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

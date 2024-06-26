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
Custom Connector Plugin Management API

This is Custom Connector Plugin Management API.

API version: 1.0.0
Contact: compute-platform-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type PresignedUrlsConnectV1Api interface {

	/*
		PresignedUploadUrlConnectV1PresignedUrl Request a presigned upload URL for a new Custom Connector Plugin.

		[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

	Request a presigned upload URL to upload a Custom Connector Plugin archive.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiPresignedUploadUrlConnectV1PresignedUrlRequest
	*/
	PresignedUploadUrlConnectV1PresignedUrl(ctx _context.Context) ApiPresignedUploadUrlConnectV1PresignedUrlRequest

	// PresignedUploadUrlConnectV1PresignedUrlExecute executes the request
	//  @return ConnectV1PresignedUrl
	PresignedUploadUrlConnectV1PresignedUrlExecute(r ApiPresignedUploadUrlConnectV1PresignedUrlRequest) (ConnectV1PresignedUrl, *_nethttp.Response, error)
}

// PresignedUrlsConnectV1ApiService PresignedUrlsConnectV1Api service
type PresignedUrlsConnectV1ApiService service

type ApiPresignedUploadUrlConnectV1PresignedUrlRequest struct {
	ctx                          _context.Context
	ApiService                   PresignedUrlsConnectV1Api
	connectV1PresignedUrlRequest *ConnectV1PresignedUrlRequest
}

func (r ApiPresignedUploadUrlConnectV1PresignedUrlRequest) ConnectV1PresignedUrlRequest(connectV1PresignedUrlRequest ConnectV1PresignedUrlRequest) ApiPresignedUploadUrlConnectV1PresignedUrlRequest {
	r.connectV1PresignedUrlRequest = &connectV1PresignedUrlRequest
	return r
}

func (r ApiPresignedUploadUrlConnectV1PresignedUrlRequest) Execute() (ConnectV1PresignedUrl, *_nethttp.Response, error) {
	return r.ApiService.PresignedUploadUrlConnectV1PresignedUrlExecute(r)
}

/*
PresignedUploadUrlConnectV1PresignedUrl Request a presigned upload URL for a new Custom Connector Plugin.

[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Request a presigned upload URL to upload a Custom Connector Plugin archive.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPresignedUploadUrlConnectV1PresignedUrlRequest
*/
func (a *PresignedUrlsConnectV1ApiService) PresignedUploadUrlConnectV1PresignedUrl(ctx _context.Context) ApiPresignedUploadUrlConnectV1PresignedUrlRequest {
	return ApiPresignedUploadUrlConnectV1PresignedUrlRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ConnectV1PresignedUrl
func (a *PresignedUrlsConnectV1ApiService) PresignedUploadUrlConnectV1PresignedUrlExecute(r ApiPresignedUploadUrlConnectV1PresignedUrlRequest) (ConnectV1PresignedUrl, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectV1PresignedUrl
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PresignedUrlsConnectV1ApiService.PresignedUploadUrlConnectV1PresignedUrl")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connect/v1/presigned-upload-url"

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
	localVarPostBody = r.connectV1PresignedUrlRequest
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

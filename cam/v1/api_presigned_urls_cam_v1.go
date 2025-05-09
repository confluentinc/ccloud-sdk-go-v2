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
Connect Artifact Management API

This is the Connect Management API.

API version: 0.0.1
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
)

// Linger please
var (
	_ _context.Context
)

type PresignedUrlsCamV1Api interface {

	/*
			PresignedUploadUrlCamV1PresignedUrl Request a presigned upload URL for a new Connect Artifact.

			[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Request a presigned upload URL to upload a Connect Artifact archive.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiPresignedUploadUrlCamV1PresignedUrlRequest
	*/
	PresignedUploadUrlCamV1PresignedUrl(ctx _context.Context) ApiPresignedUploadUrlCamV1PresignedUrlRequest

	// PresignedUploadUrlCamV1PresignedUrlExecute executes the request
	//  @return CamV1PresignedUrl
	PresignedUploadUrlCamV1PresignedUrlExecute(r ApiPresignedUploadUrlCamV1PresignedUrlRequest) (CamV1PresignedUrl, *_nethttp.Response, error)
}

// PresignedUrlsCamV1ApiService PresignedUrlsCamV1Api service
type PresignedUrlsCamV1ApiService service

type ApiPresignedUploadUrlCamV1PresignedUrlRequest struct {
	ctx                      _context.Context
	ApiService               PresignedUrlsCamV1Api
	camV1PresignedUrlRequest *CamV1PresignedUrlRequest
}

func (r ApiPresignedUploadUrlCamV1PresignedUrlRequest) CamV1PresignedUrlRequest(camV1PresignedUrlRequest CamV1PresignedUrlRequest) ApiPresignedUploadUrlCamV1PresignedUrlRequest {
	r.camV1PresignedUrlRequest = &camV1PresignedUrlRequest
	return r
}

func (r ApiPresignedUploadUrlCamV1PresignedUrlRequest) Execute() (CamV1PresignedUrl, *_nethttp.Response, error) {
	return r.ApiService.PresignedUploadUrlCamV1PresignedUrlExecute(r)
}

/*
PresignedUploadUrlCamV1PresignedUrl Request a presigned upload URL for a new Connect Artifact.

[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Request a presigned upload URL to upload a Connect Artifact archive.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPresignedUploadUrlCamV1PresignedUrlRequest
*/
func (a *PresignedUrlsCamV1ApiService) PresignedUploadUrlCamV1PresignedUrl(ctx _context.Context) ApiPresignedUploadUrlCamV1PresignedUrlRequest {
	return ApiPresignedUploadUrlCamV1PresignedUrlRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CamV1PresignedUrl
func (a *PresignedUrlsCamV1ApiService) PresignedUploadUrlCamV1PresignedUrlExecute(r ApiPresignedUploadUrlCamV1PresignedUrlRequest) (CamV1PresignedUrl, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CamV1PresignedUrl
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PresignedUrlsCamV1ApiService.PresignedUploadUrlCamV1PresignedUrl")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cam/v1/presigned-upload-url"

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
	localVarPostBody = r.camV1PresignedUrlRequest
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

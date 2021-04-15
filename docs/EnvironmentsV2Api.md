# \EnvironmentsV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateV2Environment**](EnvironmentsV2Api.md#CreateV2Environment) | **Post** /v2/environments | Create an Environment
[**DeleteV2Environment**](EnvironmentsV2Api.md#DeleteV2Environment) | **Delete** /v2/environments/{id} | Delete an Environment
[**GetV2Environment**](EnvironmentsV2Api.md#GetV2Environment) | **Get** /v2/environments/{id} | Read an Environment
[**ListV2Environments**](EnvironmentsV2Api.md#ListV2Environments) | **Get** /v2/environments | List of Environments
[**UpdateV2Environment**](EnvironmentsV2Api.md#UpdateV2Environment) | **Patch** /v2/environments/{id} | Update an Environment



## CreateV2Environment

> V2Environment CreateV2Environment(ctx).V2Environment(v2Environment).Execute()

Create an Environment



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateV2EnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2Environment** | [**V2Environment**](V2Environment.md) |  | 

### Return type

[**V2Environment**](v2.Environment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteV2Environment

> DeleteV2Environment(ctx, id).Execute()

Delete an Environment



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV2EnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV2Environment

> V2Environment GetV2Environment(ctx, id).Execute()

Read an Environment



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2EnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2Environment**](v2.Environment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListV2Environments

> V2EnvironmentList ListV2Environments(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Environments



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListV2EnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**V2EnvironmentList**](v2.EnvironmentList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateV2Environment

> V2Environment UpdateV2Environment(ctx, id).V2Environment(v2Environment).Execute()

Update an Environment



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateV2EnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2Environment** | [**V2Environment**](V2Environment.md) |  | 

### Return type

[**V2Environment**](v2.Environment.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


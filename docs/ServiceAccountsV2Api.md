# \ServiceAccountsV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateV2ServiceAccount**](ServiceAccountsV2Api.md#CreateV2ServiceAccount) | **Post** /v2/service-accounts | Create a Service Account
[**DeleteV2ServiceAccount**](ServiceAccountsV2Api.md#DeleteV2ServiceAccount) | **Delete** /v2/service-accounts/{id} | Delete a Service Account
[**GetV2ServiceAccount**](ServiceAccountsV2Api.md#GetV2ServiceAccount) | **Get** /v2/service-accounts/{id} | Read a Service Account
[**ListV2ServiceAccounts**](ServiceAccountsV2Api.md#ListV2ServiceAccounts) | **Get** /v2/service-accounts | List of Service Accounts
[**UpdateV2ServiceAccount**](ServiceAccountsV2Api.md#UpdateV2ServiceAccount) | **Patch** /v2/service-accounts/{id} | Update a Service Account



## CreateV2ServiceAccount

> V2ServiceAccount CreateV2ServiceAccount(ctx).V2ServiceAccount(v2ServiceAccount).Execute()

Create a Service Account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateV2ServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2ServiceAccount** | [**V2ServiceAccount**](V2ServiceAccount.md) |  | 

### Return type

[**V2ServiceAccount**](v2.ServiceAccount.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteV2ServiceAccount

> DeleteV2ServiceAccount(ctx, id).Execute()

Delete a Service Account



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV2ServiceAccountRequest struct via the builder pattern


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


## GetV2ServiceAccount

> V2ServiceAccount GetV2ServiceAccount(ctx, id).Execute()

Read a Service Account



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2ServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2ServiceAccount**](v2.ServiceAccount.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListV2ServiceAccounts

> V2ServiceAccountList ListV2ServiceAccounts(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Service Accounts



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListV2ServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**V2ServiceAccountList**](v2.ServiceAccountList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateV2ServiceAccount

> V2ServiceAccount UpdateV2ServiceAccount(ctx, id).V2ServiceAccount(v2ServiceAccount).Execute()

Update a Service Account



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateV2ServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2ServiceAccount** | [**V2ServiceAccount**](V2ServiceAccount.md) |  | 

### Return type

[**V2ServiceAccount**](v2.ServiceAccount.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


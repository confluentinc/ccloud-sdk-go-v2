# \UsersV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV2User**](UsersV2Api.md#DeleteV2User) | **Delete** /v2/users/{id} | Delete a User
[**GetV2User**](UsersV2Api.md#GetV2User) | **Get** /v2/users/{id} | Read a User
[**ListV2Users**](UsersV2Api.md#ListV2Users) | **Get** /v2/users | List of Users
[**UpdateV2User**](UsersV2Api.md#UpdateV2User) | **Patch** /v2/users/{id} | Update a User



## DeleteV2User

> DeleteV2User(ctx, id).Execute()

Delete a User



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV2UserRequest struct via the builder pattern


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


## GetV2User

> V2User GetV2User(ctx, id).Execute()

Read a User



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2UserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2User**](v2.User.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListV2Users

> V2UserList ListV2Users(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Users



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListV2UsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**V2UserList**](v2.UserList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateV2User

> V2User UpdateV2User(ctx, id).V2User(v2User).Execute()

Update a User



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateV2UserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2User** | [**V2User**](V2User.md) |  | 

### Return type

[**V2User**](v2.User.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


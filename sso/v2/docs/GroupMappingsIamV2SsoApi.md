# \GroupMappingsIamV2SsoApi

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2SsoGroupMapping**](GroupMappingsIamV2SsoApi.md#CreateIamV2SsoGroupMapping) | **Post** /iam/v2/sso/group-mappings | Create a Group Mapping
[**DeleteIamV2SsoGroupMapping**](GroupMappingsIamV2SsoApi.md#DeleteIamV2SsoGroupMapping) | **Delete** /iam/v2/sso/group-mappings/{id} | Delete a Group Mapping
[**GetIamV2SsoGroupMapping**](GroupMappingsIamV2SsoApi.md#GetIamV2SsoGroupMapping) | **Get** /iam/v2/sso/group-mappings/{id} | Read a Group Mapping
[**ListIamV2SsoGroupMappings**](GroupMappingsIamV2SsoApi.md#ListIamV2SsoGroupMappings) | **Get** /iam/v2/sso/group-mappings | List of Group Mappings
[**UpdateIamV2SsoGroupMapping**](GroupMappingsIamV2SsoApi.md#UpdateIamV2SsoGroupMapping) | **Patch** /iam/v2/sso/group-mappings/{id} | Update a Group Mapping



## CreateIamV2SsoGroupMapping

> IamV2SsoGroupMapping CreateIamV2SsoGroupMapping(ctx).IamV2SsoGroupMapping(iamV2SsoGroupMapping).Execute()

Create a Group Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    iamV2SsoGroupMapping := *openapiclient.NewIamV2SsoGroupMapping() // IamV2SsoGroupMapping |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupMappingsIamV2SsoApi.CreateIamV2SsoGroupMapping(context.Background()).IamV2SsoGroupMapping(iamV2SsoGroupMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupMappingsIamV2SsoApi.CreateIamV2SsoGroupMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2SsoGroupMapping`: IamV2SsoGroupMapping
    fmt.Fprintf(os.Stdout, "Response from `GroupMappingsIamV2SsoApi.CreateIamV2SsoGroupMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2SsoGroupMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamV2SsoGroupMapping** | [**IamV2SsoGroupMapping**](IamV2SsoGroupMapping.md) |  | 

### Return type

[**IamV2SsoGroupMapping**](iam.v2.sso.GroupMapping.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2SsoGroupMapping

> DeleteIamV2SsoGroupMapping(ctx, id).Execute()

Delete a Group Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The unique identifier for the group mapping.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupMappingsIamV2SsoApi.DeleteIamV2SsoGroupMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupMappingsIamV2SsoApi.DeleteIamV2SsoGroupMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the group mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2SsoGroupMappingRequest struct via the builder pattern


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


## GetIamV2SsoGroupMapping

> IamV2SsoGroupMapping GetIamV2SsoGroupMapping(ctx, id).Execute()

Read a Group Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The unique identifier for the group mapping.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupMappingsIamV2SsoApi.GetIamV2SsoGroupMapping(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupMappingsIamV2SsoApi.GetIamV2SsoGroupMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2SsoGroupMapping`: IamV2SsoGroupMapping
    fmt.Fprintf(os.Stdout, "Response from `GroupMappingsIamV2SsoApi.GetIamV2SsoGroupMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the group mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2SsoGroupMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamV2SsoGroupMapping**](iam.v2.sso.GroupMapping.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2SsoGroupMappings

> IamV2SsoGroupMappingList ListIamV2SsoGroupMappings(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Group Mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupMappingsIamV2SsoApi.ListIamV2SsoGroupMappings(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupMappingsIamV2SsoApi.ListIamV2SsoGroupMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2SsoGroupMappings`: IamV2SsoGroupMappingList
    fmt.Fprintf(os.Stdout, "Response from `GroupMappingsIamV2SsoApi.ListIamV2SsoGroupMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2SsoGroupMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**IamV2SsoGroupMappingList**](iam.v2.sso.GroupMappingList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIamV2SsoGroupMapping

> IamV2SsoGroupMapping UpdateIamV2SsoGroupMapping(ctx, id).IamV2SsoGroupMapping(iamV2SsoGroupMapping).Execute()

Update a Group Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The unique identifier for the group mapping.
    iamV2SsoGroupMapping := *openapiclient.NewIamV2SsoGroupMapping() // IamV2SsoGroupMapping |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupMappingsIamV2SsoApi.UpdateIamV2SsoGroupMapping(context.Background(), id).IamV2SsoGroupMapping(iamV2SsoGroupMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupMappingsIamV2SsoApi.UpdateIamV2SsoGroupMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIamV2SsoGroupMapping`: IamV2SsoGroupMapping
    fmt.Fprintf(os.Stdout, "Response from `GroupMappingsIamV2SsoApi.UpdateIamV2SsoGroupMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the group mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIamV2SsoGroupMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamV2SsoGroupMapping** | [**IamV2SsoGroupMapping**](IamV2SsoGroupMapping.md) |  | 

### Return type

[**IamV2SsoGroupMapping**](iam.v2.sso.GroupMapping.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


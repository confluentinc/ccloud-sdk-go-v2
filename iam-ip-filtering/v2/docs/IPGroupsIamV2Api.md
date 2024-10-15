# \IPGroupsIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2IpGroup**](IPGroupsIamV2Api.md#CreateIamV2IpGroup) | **Post** /iam/v2/ip-groups | Create an IP Group
[**DeleteIamV2IpGroup**](IPGroupsIamV2Api.md#DeleteIamV2IpGroup) | **Delete** /iam/v2/ip-groups/{id} | Delete an IP Group
[**GetIamV2IpGroup**](IPGroupsIamV2Api.md#GetIamV2IpGroup) | **Get** /iam/v2/ip-groups/{id} | Read an IP Group
[**ListIamV2IpGroups**](IPGroupsIamV2Api.md#ListIamV2IpGroups) | **Get** /iam/v2/ip-groups | List of IP Groups
[**UpdateIamV2IpGroup**](IPGroupsIamV2Api.md#UpdateIamV2IpGroup) | **Patch** /iam/v2/ip-groups/{id} | Update an IP Group



## CreateIamV2IpGroup

> IamV2IpGroup CreateIamV2IpGroup(ctx).IamV2IpGroup(iamV2IpGroup).Execute()

Create an IP Group



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
    iamV2IpGroup := *openapiclient.NewIamV2IpGroup() // IamV2IpGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPGroupsIamV2Api.CreateIamV2IpGroup(context.Background()).IamV2IpGroup(iamV2IpGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPGroupsIamV2Api.CreateIamV2IpGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2IpGroup`: IamV2IpGroup
    fmt.Fprintf(os.Stdout, "Response from `IPGroupsIamV2Api.CreateIamV2IpGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2IpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamV2IpGroup** | [**IamV2IpGroup**](IamV2IpGroup.md) |  | 

### Return type

[**IamV2IpGroup**](iam.v2.IpGroup.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2IpGroup

> DeleteIamV2IpGroup(ctx, id).Execute()

Delete an IP Group



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
    id := "id_example" // string | The unique identifier for the IP group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPGroupsIamV2Api.DeleteIamV2IpGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPGroupsIamV2Api.DeleteIamV2IpGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the IP group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2IpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamV2IpGroup

> IamV2IpGroup GetIamV2IpGroup(ctx, id).Execute()

Read an IP Group



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
    id := "id_example" // string | The unique identifier for the IP group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPGroupsIamV2Api.GetIamV2IpGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPGroupsIamV2Api.GetIamV2IpGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2IpGroup`: IamV2IpGroup
    fmt.Fprintf(os.Stdout, "Response from `IPGroupsIamV2Api.GetIamV2IpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the IP group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2IpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamV2IpGroup**](iam.v2.IpGroup.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2IpGroups

> IamV2IpGroupList ListIamV2IpGroups(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of IP Groups



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 25)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPGroupsIamV2Api.ListIamV2IpGroups(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPGroupsIamV2Api.ListIamV2IpGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2IpGroups`: IamV2IpGroupList
    fmt.Fprintf(os.Stdout, "Response from `IPGroupsIamV2Api.ListIamV2IpGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2IpGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 25]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**IamV2IpGroupList**](iam.v2.IpGroupList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIamV2IpGroup

> IamV2IpGroup UpdateIamV2IpGroup(ctx, id).IamV2IpGroup(iamV2IpGroup).Execute()

Update an IP Group



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
    id := "id_example" // string | The unique identifier for the IP group.
    iamV2IpGroup := *openapiclient.NewIamV2IpGroup() // IamV2IpGroup |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IPGroupsIamV2Api.UpdateIamV2IpGroup(context.Background(), id).IamV2IpGroup(iamV2IpGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPGroupsIamV2Api.UpdateIamV2IpGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIamV2IpGroup`: IamV2IpGroup
    fmt.Fprintf(os.Stdout, "Response from `IPGroupsIamV2Api.UpdateIamV2IpGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the IP group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIamV2IpGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamV2IpGroup** | [**IamV2IpGroup**](IamV2IpGroup.md) |  | 

### Return type

[**IamV2IpGroup**](iam.v2.IpGroup.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


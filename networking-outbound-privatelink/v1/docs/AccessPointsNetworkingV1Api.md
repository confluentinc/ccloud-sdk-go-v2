# \AccessPointsNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1AccessPoint**](AccessPointsNetworkingV1Api.md#CreateNetworkingV1AccessPoint) | **Post** /networking/v1/access-points | Create an Access Point
[**DeleteNetworkingV1AccessPoint**](AccessPointsNetworkingV1Api.md#DeleteNetworkingV1AccessPoint) | **Delete** /networking/v1/access-points/{id} | Delete an Access Point
[**GetNetworkingV1AccessPoint**](AccessPointsNetworkingV1Api.md#GetNetworkingV1AccessPoint) | **Get** /networking/v1/access-points/{id} | Read an Access Point
[**ListNetworkingV1AccessPoints**](AccessPointsNetworkingV1Api.md#ListNetworkingV1AccessPoints) | **Get** /networking/v1/access-points | List of Access Points
[**UpdateNetworkingV1AccessPoint**](AccessPointsNetworkingV1Api.md#UpdateNetworkingV1AccessPoint) | **Patch** /networking/v1/access-points/{id} | Update an Access Point



## CreateNetworkingV1AccessPoint

> NetworkingV1AccessPoint CreateNetworkingV1AccessPoint(ctx).NetworkingV1AccessPoint(networkingV1AccessPoint).Execute()

Create an Access Point



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
    networkingV1AccessPoint := *openapiclient.NewNetworkingV1AccessPoint() // NetworkingV1AccessPoint |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessPointsNetworkingV1Api.CreateNetworkingV1AccessPoint(context.Background()).NetworkingV1AccessPoint(networkingV1AccessPoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPointsNetworkingV1Api.CreateNetworkingV1AccessPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1AccessPoint`: NetworkingV1AccessPoint
    fmt.Fprintf(os.Stdout, "Response from `AccessPointsNetworkingV1Api.CreateNetworkingV1AccessPoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1AccessPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1AccessPoint** | [**NetworkingV1AccessPoint**](NetworkingV1AccessPoint.md) |  | 

### Return type

[**NetworkingV1AccessPoint**](networking.v1.AccessPoint.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1AccessPoint

> DeleteNetworkingV1AccessPoint(ctx, id).Environment(environment).Execute()

Delete an Access Point



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the access point.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessPointsNetworkingV1Api.DeleteNetworkingV1AccessPoint(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPointsNetworkingV1Api.DeleteNetworkingV1AccessPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the access point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1AccessPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


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


## GetNetworkingV1AccessPoint

> NetworkingV1AccessPoint GetNetworkingV1AccessPoint(ctx, id).Environment(environment).Execute()

Read an Access Point



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
    environment := "env-00000" // string | Scope the operation to the given environment.
    id := "id_example" // string | The unique identifier for the access point.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessPointsNetworkingV1Api.GetNetworkingV1AccessPoint(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPointsNetworkingV1Api.GetNetworkingV1AccessPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1AccessPoint`: NetworkingV1AccessPoint
    fmt.Fprintf(os.Stdout, "Response from `AccessPointsNetworkingV1Api.GetNetworkingV1AccessPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the access point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1AccessPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1AccessPoint**](networking.v1.AccessPoint.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1AccessPoints

> NetworkingV1AccessPointList ListNetworkingV1AccessPoints(ctx).Environment(environment).SpecDisplayName(specDisplayName).PageSize(pageSize).PageToken(pageToken).Execute()

List of Access Points



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
    environment := "env-00000" // string | Filter the results by exact match for environment.
    specDisplayName := []string{"Inner_example"} // []string | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessPointsNetworkingV1Api.ListNetworkingV1AccessPoints(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPointsNetworkingV1Api.ListNetworkingV1AccessPoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1AccessPoints`: NetworkingV1AccessPointList
    fmt.Fprintf(os.Stdout, "Response from `AccessPointsNetworkingV1Api.ListNetworkingV1AccessPoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1AccessPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | **[]string** | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1AccessPointList**](networking.v1.AccessPointList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1AccessPoint

> NetworkingV1AccessPoint UpdateNetworkingV1AccessPoint(ctx, id).NetworkingV1AccessPointUpdate(networkingV1AccessPointUpdate).Execute()

Update an Access Point



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
    id := "id_example" // string | The unique identifier for the access point.
    networkingV1AccessPointUpdate := *openapiclient.NewNetworkingV1AccessPointUpdate() // NetworkingV1AccessPointUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessPointsNetworkingV1Api.UpdateNetworkingV1AccessPoint(context.Background(), id).NetworkingV1AccessPointUpdate(networkingV1AccessPointUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPointsNetworkingV1Api.UpdateNetworkingV1AccessPoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1AccessPoint`: NetworkingV1AccessPoint
    fmt.Fprintf(os.Stdout, "Response from `AccessPointsNetworkingV1Api.UpdateNetworkingV1AccessPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the access point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1AccessPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1AccessPointUpdate** | [**NetworkingV1AccessPointUpdate**](NetworkingV1AccessPointUpdate.md) |  | 

### Return type

[**NetworkingV1AccessPoint**](networking.v1.AccessPoint.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


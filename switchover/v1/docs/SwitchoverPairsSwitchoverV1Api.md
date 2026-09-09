# \SwitchoverPairsSwitchoverV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSwitchoverV1SwitchoverPair**](SwitchoverPairsSwitchoverV1Api.md#CreateSwitchoverV1SwitchoverPair) | **Post** /switchover/v1/switchover-pairs | Create a Switchover Pair
[**DeleteSwitchoverV1SwitchoverPair**](SwitchoverPairsSwitchoverV1Api.md#DeleteSwitchoverV1SwitchoverPair) | **Delete** /switchover/v1/switchover-pairs/{id} | Delete a Switchover Pair
[**FailoverSwitchoverV1SwitchoverPair**](SwitchoverPairsSwitchoverV1Api.md#FailoverSwitchoverV1SwitchoverPair) | **Post** /switchover/v1/switchover-pairs/{id}:failover | Trigger a failover on a switchover pair
[**GetSwitchoverV1SwitchoverPair**](SwitchoverPairsSwitchoverV1Api.md#GetSwitchoverV1SwitchoverPair) | **Get** /switchover/v1/switchover-pairs/{id} | Read a Switchover Pair
[**ListSwitchoverV1SwitchoverPairs**](SwitchoverPairsSwitchoverV1Api.md#ListSwitchoverV1SwitchoverPairs) | **Get** /switchover/v1/switchover-pairs | List of Switchover Pairs
[**UpdateSwitchoverV1SwitchoverPair**](SwitchoverPairsSwitchoverV1Api.md#UpdateSwitchoverV1SwitchoverPair) | **Put** /switchover/v1/switchover-pairs/{id} | Update a Switchover Pair



## CreateSwitchoverV1SwitchoverPair

> SwitchoverV1SwitchoverPair CreateSwitchoverV1SwitchoverPair(ctx).SwitchoverV1SwitchoverPair(switchoverV1SwitchoverPair).Execute()

Create a Switchover Pair



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
    switchoverV1SwitchoverPair := *openapiclient.NewSwitchoverV1SwitchoverPair() // SwitchoverV1SwitchoverPair |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverPairsSwitchoverV1Api.CreateSwitchoverV1SwitchoverPair(context.Background()).SwitchoverV1SwitchoverPair(switchoverV1SwitchoverPair).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverPairsSwitchoverV1Api.CreateSwitchoverV1SwitchoverPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSwitchoverV1SwitchoverPair`: SwitchoverV1SwitchoverPair
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverPairsSwitchoverV1Api.CreateSwitchoverV1SwitchoverPair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSwitchoverV1SwitchoverPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **switchoverV1SwitchoverPair** | [**SwitchoverV1SwitchoverPair**](SwitchoverV1SwitchoverPair.md) |  | 

### Return type

[**SwitchoverV1SwitchoverPair**](switchover.v1.SwitchoverPair.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSwitchoverV1SwitchoverPair

> SwitchoverV1SwitchoverPair DeleteSwitchoverV1SwitchoverPair(ctx, id).Environment(environment).Execute()

Delete a Switchover Pair



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
    environment := "env-00000" // string | Scope the operation to the environment with this ID. The bare-ID form of `environment_crn`; it is the standard Confluent collection filter, passed as the `?environment=` query parameter. 
    id := "id_example" // string | The unique identifier for the switchover pair.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverPairsSwitchoverV1Api.DeleteSwitchoverV1SwitchoverPair(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverPairsSwitchoverV1Api.DeleteSwitchoverV1SwitchoverPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSwitchoverV1SwitchoverPair`: SwitchoverV1SwitchoverPair
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverPairsSwitchoverV1Api.DeleteSwitchoverV1SwitchoverPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the switchover pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSwitchoverV1SwitchoverPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the environment with this ID. The bare-ID form of &#x60;environment_crn&#x60;; it is the standard Confluent collection filter, passed as the &#x60;?environment&#x3D;&#x60; query parameter.  | 


### Return type

[**SwitchoverV1SwitchoverPair**](SwitchoverV1SwitchoverPair.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FailoverSwitchoverV1SwitchoverPair

> SwitchoverV1SwitchoverPair FailoverSwitchoverV1SwitchoverPair(ctx, id).SwitchoverV1SwitchoverPairFailoverRequest(switchoverV1SwitchoverPairFailoverRequest).Execute()

Trigger a failover on a switchover pair



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
    id := "id_example" // string | The unique identifier for the switchover pair.
    switchoverV1SwitchoverPairFailoverRequest := *openapiclient.NewSwitchoverV1SwitchoverPairFailoverRequest(*openapiclient.NewSwitchoverV1SwitchoverPairFailoverRequestSpec("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-00000", "PLANNED")) // SwitchoverV1SwitchoverPairFailoverRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverPairsSwitchoverV1Api.FailoverSwitchoverV1SwitchoverPair(context.Background(), id).SwitchoverV1SwitchoverPairFailoverRequest(switchoverV1SwitchoverPairFailoverRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverPairsSwitchoverV1Api.FailoverSwitchoverV1SwitchoverPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FailoverSwitchoverV1SwitchoverPair`: SwitchoverV1SwitchoverPair
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverPairsSwitchoverV1Api.FailoverSwitchoverV1SwitchoverPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the switchover pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFailoverSwitchoverV1SwitchoverPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **switchoverV1SwitchoverPairFailoverRequest** | [**SwitchoverV1SwitchoverPairFailoverRequest**](SwitchoverV1SwitchoverPairFailoverRequest.md) |  | 

### Return type

[**SwitchoverV1SwitchoverPair**](SwitchoverV1SwitchoverPair.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSwitchoverV1SwitchoverPair

> SwitchoverV1SwitchoverPair GetSwitchoverV1SwitchoverPair(ctx, id).Environment(environment).Execute()

Read a Switchover Pair



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
    environment := "env-00000" // string | Scope the operation to the environment with this ID. The bare-ID form of `environment_crn`; it is the standard Confluent collection filter, passed as the `?environment=` query parameter. 
    id := "id_example" // string | The unique identifier for the switchover pair.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverPairsSwitchoverV1Api.GetSwitchoverV1SwitchoverPair(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverPairsSwitchoverV1Api.GetSwitchoverV1SwitchoverPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSwitchoverV1SwitchoverPair`: SwitchoverV1SwitchoverPair
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverPairsSwitchoverV1Api.GetSwitchoverV1SwitchoverPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the switchover pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchoverV1SwitchoverPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the environment with this ID. The bare-ID form of &#x60;environment_crn&#x60;; it is the standard Confluent collection filter, passed as the &#x60;?environment&#x3D;&#x60; query parameter.  | 


### Return type

[**SwitchoverV1SwitchoverPair**](switchover.v1.SwitchoverPair.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSwitchoverV1SwitchoverPairs

> SwitchoverV1SwitchoverPairList ListSwitchoverV1SwitchoverPairs(ctx).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()

List of Switchover Pairs



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
    environment := "env-00000" // string | Scope the operation to the environment with this ID. The bare-ID form of `environment_crn`; it is the standard Confluent collection filter, passed as the `?environment=` query parameter. 
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverPairsSwitchoverV1Api.ListSwitchoverV1SwitchoverPairs(context.Background()).Environment(environment).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverPairsSwitchoverV1Api.ListSwitchoverV1SwitchoverPairs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSwitchoverV1SwitchoverPairs`: SwitchoverV1SwitchoverPairList
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverPairsSwitchoverV1Api.ListSwitchoverV1SwitchoverPairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSwitchoverV1SwitchoverPairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the environment with this ID. The bare-ID form of &#x60;environment_crn&#x60;; it is the standard Confluent collection filter, passed as the &#x60;?environment&#x3D;&#x60; query parameter.  | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**SwitchoverV1SwitchoverPairList**](switchover.v1.SwitchoverPairList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSwitchoverV1SwitchoverPair

> SwitchoverV1SwitchoverPair UpdateSwitchoverV1SwitchoverPair(ctx, id).Environment(environment).SwitchoverV1SwitchoverPairUpdateRequest(switchoverV1SwitchoverPairUpdateRequest).Execute()

Update a Switchover Pair



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
    environment := "env-00000" // string | Scope the operation to the environment with this ID. The bare-ID form of `environment_crn`; it is the standard Confluent collection filter, passed as the `?environment=` query parameter. 
    id := "id_example" // string | The unique identifier for the switchover pair.
    switchoverV1SwitchoverPairUpdateRequest := *openapiclient.NewSwitchoverV1SwitchoverPairUpdateRequest(*openapiclient.NewSwitchoverV1SwitchoverPairUpdateRequestSpec()) // SwitchoverV1SwitchoverPairUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SwitchoverPairsSwitchoverV1Api.UpdateSwitchoverV1SwitchoverPair(context.Background(), id).Environment(environment).SwitchoverV1SwitchoverPairUpdateRequest(switchoverV1SwitchoverPairUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SwitchoverPairsSwitchoverV1Api.UpdateSwitchoverV1SwitchoverPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSwitchoverV1SwitchoverPair`: SwitchoverV1SwitchoverPair
    fmt.Fprintf(os.Stdout, "Response from `SwitchoverPairsSwitchoverV1Api.UpdateSwitchoverV1SwitchoverPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the switchover pair. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSwitchoverV1SwitchoverPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the environment with this ID. The bare-ID form of &#x60;environment_crn&#x60;; it is the standard Confluent collection filter, passed as the &#x60;?environment&#x3D;&#x60; query parameter.  | 

 **switchoverV1SwitchoverPairUpdateRequest** | [**SwitchoverV1SwitchoverPairUpdateRequest**](SwitchoverV1SwitchoverPairUpdateRequest.md) |  | 

### Return type

[**SwitchoverV1SwitchoverPair**](switchover.v1.SwitchoverPair.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [global-api-key](../README.md#global-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


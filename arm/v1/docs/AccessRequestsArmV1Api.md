# \AccessRequestsArmV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArmV1AccessRequest**](AccessRequestsArmV1Api.md#CreateArmV1AccessRequest) | **Post** /arm/v1/access-requests | Create an Access Request
[**ListArmV1AccessRequests**](AccessRequestsArmV1Api.md#ListArmV1AccessRequests) | **Get** /arm/v1/access-requests | List of Access Requests
[**UpdateArmV1AccessRequest**](AccessRequestsArmV1Api.md#UpdateArmV1AccessRequest) | **Patch** /arm/v1/access-requests | Act on access requests



## CreateArmV1AccessRequest

> ArmV1AccessRequest CreateArmV1AccessRequest(ctx).ArmV1AccessRequest(armV1AccessRequest).Execute()

Create an Access Request



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
    armV1AccessRequest := *openapiclient.NewArmV1AccessRequest() // ArmV1AccessRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessRequestsArmV1Api.CreateArmV1AccessRequest(context.Background()).ArmV1AccessRequest(armV1AccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsArmV1Api.CreateArmV1AccessRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArmV1AccessRequest`: ArmV1AccessRequest
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsArmV1Api.CreateArmV1AccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArmV1AccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **armV1AccessRequest** | [**ArmV1AccessRequest**](ArmV1AccessRequest.md) |  | 

### Return type

[**ArmV1AccessRequest**](arm.v1.AccessRequest.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArmV1AccessRequests

> ArmV1AccessRequestList ListArmV1AccessRequests(ctx).Status(status).Resource(resource).SelfInitiatedOnly(selfInitiatedOnly).CanActOnOnly(canActOnOnly).Execute()

List of Access Requests



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
    status := "Pending" // string | Filter the results by exact match for status. (optional)
    resource := "crn://confluent.cloud/organization=92fe5f20-f7c2-4a1b-8f62-b1aab73c915c/environment=env-mx05q/cloud-cluster=lkc-111aaa/kafka=lkc-111aaa/topic=myTopic" // string | Filter the results by exact match for resource. (optional)
    selfInitiatedOnly := true // bool | Include requests initiated by current user only (optional)
    canActOnOnly := true // bool | Include requests that current user can operate on only (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessRequestsArmV1Api.ListArmV1AccessRequests(context.Background()).Status(status).Resource(resource).SelfInitiatedOnly(selfInitiatedOnly).CanActOnOnly(canActOnOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsArmV1Api.ListArmV1AccessRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArmV1AccessRequests`: ArmV1AccessRequestList
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsArmV1Api.ListArmV1AccessRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListArmV1AccessRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter the results by exact match for status. | 
 **resource** | **string** | Filter the results by exact match for resource. | 
 **selfInitiatedOnly** | **bool** | Include requests initiated by current user only | 
 **canActOnOnly** | **bool** | Include requests that current user can operate on only | 

### Return type

[**ArmV1AccessRequestList**](arm.v1.AccessRequestList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArmV1AccessRequest

> InlineResponse200 UpdateArmV1AccessRequest(ctx).ArmV1UpdateAccessRequestRequest(armV1UpdateAccessRequestRequest).Execute()

Act on access requests



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
    armV1UpdateAccessRequestRequest := *openapiclient.NewArmV1UpdateAccessRequestRequest() // ArmV1UpdateAccessRequestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessRequestsArmV1Api.UpdateArmV1AccessRequest(context.Background()).ArmV1UpdateAccessRequestRequest(armV1UpdateAccessRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsArmV1Api.UpdateArmV1AccessRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateArmV1AccessRequest`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsArmV1Api.UpdateArmV1AccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArmV1AccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **armV1UpdateAccessRequestRequest** | [**ArmV1UpdateAccessRequestRequest**](ArmV1UpdateAccessRequestRequest.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \PoliciesConfigurationcontrolV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigurationcontrolV1Policy**](PoliciesConfigurationcontrolV1Api.md#CreateConfigurationcontrolV1Policy) | **Post** /configurationcontrol/v1/policies | Create a Policy
[**DeleteConfigurationcontrolV1Policy**](PoliciesConfigurationcontrolV1Api.md#DeleteConfigurationcontrolV1Policy) | **Delete** /configurationcontrol/v1/policies/{name} | Delete a Policy
[**GetConfigurationcontrolV1Policy**](PoliciesConfigurationcontrolV1Api.md#GetConfigurationcontrolV1Policy) | **Get** /configurationcontrol/v1/policies/{name} | Read a Policy
[**ListConfigurationcontrolV1Policies**](PoliciesConfigurationcontrolV1Api.md#ListConfigurationcontrolV1Policies) | **Get** /configurationcontrol/v1/policies | List of Policies
[**UpdateConfigurationcontrolV1Policy**](PoliciesConfigurationcontrolV1Api.md#UpdateConfigurationcontrolV1Policy) | **Patch** /configurationcontrol/v1/policies/{name} | Update a Policy



## CreateConfigurationcontrolV1Policy

> ConfigurationcontrolV1Policy CreateConfigurationcontrolV1Policy(ctx).ConfigurationcontrolV1Policy(configurationcontrolV1Policy).Execute()

Create a Policy



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
    configurationcontrolV1Policy := *openapiclient.NewConfigurationcontrolV1Policy() // ConfigurationcontrolV1Policy |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesConfigurationcontrolV1Api.CreateConfigurationcontrolV1Policy(context.Background()).ConfigurationcontrolV1Policy(configurationcontrolV1Policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConfigurationcontrolV1Api.CreateConfigurationcontrolV1Policy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConfigurationcontrolV1Policy`: ConfigurationcontrolV1Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesConfigurationcontrolV1Api.CreateConfigurationcontrolV1Policy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationcontrolV1PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationcontrolV1Policy** | [**ConfigurationcontrolV1Policy**](ConfigurationcontrolV1Policy.md) |  | 

### Return type

[**ConfigurationcontrolV1Policy**](configurationcontrol.v1.Policy.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfigurationcontrolV1Policy

> DeleteConfigurationcontrolV1Policy(ctx, name).Execute()

Delete a Policy



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
    name := "name_example" // string | Client-supplied stable identifier for the policy. Used as the URL identifier and must be unique within the caller's organization. Immutable. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesConfigurationcontrolV1Api.DeleteConfigurationcontrolV1Policy(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConfigurationcontrolV1Api.DeleteConfigurationcontrolV1Policy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Client-supplied stable identifier for the policy. Used as the URL identifier and must be unique within the caller&#39;s organization. Immutable.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationcontrolV1PolicyRequest struct via the builder pattern


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


## GetConfigurationcontrolV1Policy

> ConfigurationcontrolV1Policy GetConfigurationcontrolV1Policy(ctx, name).Execute()

Read a Policy



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
    name := "name_example" // string | Client-supplied stable identifier for the policy. Used as the URL identifier and must be unique within the caller's organization. Immutable. 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesConfigurationcontrolV1Api.GetConfigurationcontrolV1Policy(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConfigurationcontrolV1Api.GetConfigurationcontrolV1Policy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurationcontrolV1Policy`: ConfigurationcontrolV1Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesConfigurationcontrolV1Api.GetConfigurationcontrolV1Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Client-supplied stable identifier for the policy. Used as the URL identifier and must be unique within the caller&#39;s organization. Immutable.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationcontrolV1PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationcontrolV1Policy**](configurationcontrol.v1.Policy.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurationcontrolV1Policies

> ConfigurationcontrolV1PolicyList ListConfigurationcontrolV1Policies(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Policies



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
    resp, r, err := api_client.PoliciesConfigurationcontrolV1Api.ListConfigurationcontrolV1Policies(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConfigurationcontrolV1Api.ListConfigurationcontrolV1Policies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfigurationcontrolV1Policies`: ConfigurationcontrolV1PolicyList
    fmt.Fprintf(os.Stdout, "Response from `PoliciesConfigurationcontrolV1Api.ListConfigurationcontrolV1Policies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConfigurationcontrolV1PoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**ConfigurationcontrolV1PolicyList**](configurationcontrol.v1.PolicyList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationcontrolV1Policy

> ConfigurationcontrolV1Policy UpdateConfigurationcontrolV1Policy(ctx, name).ConfigurationcontrolV1PolicyUpdate(configurationcontrolV1PolicyUpdate).Execute()

Update a Policy



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
    name := "name_example" // string | Client-supplied stable identifier for the policy. Used as the URL identifier and must be unique within the caller's organization. Immutable. 
    configurationcontrolV1PolicyUpdate := *openapiclient.NewConfigurationcontrolV1PolicyUpdate() // ConfigurationcontrolV1PolicyUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PoliciesConfigurationcontrolV1Api.UpdateConfigurationcontrolV1Policy(context.Background(), name).ConfigurationcontrolV1PolicyUpdate(configurationcontrolV1PolicyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PoliciesConfigurationcontrolV1Api.UpdateConfigurationcontrolV1Policy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigurationcontrolV1Policy`: ConfigurationcontrolV1Policy
    fmt.Fprintf(os.Stdout, "Response from `PoliciesConfigurationcontrolV1Api.UpdateConfigurationcontrolV1Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Client-supplied stable identifier for the policy. Used as the URL identifier and must be unique within the caller&#39;s organization. Immutable.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationcontrolV1PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationcontrolV1PolicyUpdate** | [**ConfigurationcontrolV1PolicyUpdate**](ConfigurationcontrolV1PolicyUpdate.md) |  | 

### Return type

[**ConfigurationcontrolV1Policy**](configurationcontrol.v1.Policy.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


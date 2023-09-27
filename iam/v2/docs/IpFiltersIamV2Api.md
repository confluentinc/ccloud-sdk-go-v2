# \IpFiltersIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2IpFilter**](IpFiltersIamV2Api.md#CreateIamV2IpFilter) | **Post** /iam/v2/ip-filters | Create an Ip Filter
[**DeleteIamV2IpFilter**](IpFiltersIamV2Api.md#DeleteIamV2IpFilter) | **Delete** /iam/v2/ip-filters/{id} | Delete an Ip Filter
[**GetIamV2IpFilter**](IpFiltersIamV2Api.md#GetIamV2IpFilter) | **Get** /iam/v2/ip-filters/{id} | Read an Ip Filter
[**ListIamV2IpFilters**](IpFiltersIamV2Api.md#ListIamV2IpFilters) | **Get** /iam/v2/ip-filters | List of Ip Filters



## CreateIamV2IpFilter

> IamV2IpFilter CreateIamV2IpFilter(ctx).IamV2IpFilter(iamV2IpFilter).Execute()

Create an Ip Filter



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
    iamV2IpFilter := *openapiclient.NewIamV2IpFilter() // IamV2IpFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IpFiltersIamV2Api.CreateIamV2IpFilter(context.Background()).IamV2IpFilter(iamV2IpFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpFiltersIamV2Api.CreateIamV2IpFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2IpFilter`: IamV2IpFilter
    fmt.Fprintf(os.Stdout, "Response from `IpFiltersIamV2Api.CreateIamV2IpFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2IpFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamV2IpFilter** | [**IamV2IpFilter**](IamV2IpFilter.md) |  | 

### Return type

[**IamV2IpFilter**](iam.v2.IpFilter.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2IpFilter

> DeleteIamV2IpFilter(ctx, id).Execute()

Delete an Ip Filter



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
    id := "id_example" // string | The unique identifier for the ip filter.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IpFiltersIamV2Api.DeleteIamV2IpFilter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpFiltersIamV2Api.DeleteIamV2IpFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the ip filter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2IpFilterRequest struct via the builder pattern


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


## GetIamV2IpFilter

> IamV2IpFilter GetIamV2IpFilter(ctx, id).Execute()

Read an Ip Filter



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
    id := "id_example" // string | The unique identifier for the ip filter.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IpFiltersIamV2Api.GetIamV2IpFilter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpFiltersIamV2Api.GetIamV2IpFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2IpFilter`: IamV2IpFilter
    fmt.Fprintf(os.Stdout, "Response from `IpFiltersIamV2Api.GetIamV2IpFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the ip filter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2IpFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamV2IpFilter**](iam.v2.IpFilter.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2IpFilters

> IamV2IpFilterList ListIamV2IpFilters(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Ip Filters



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
    resp, r, err := api_client.IpFiltersIamV2Api.ListIamV2IpFilters(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IpFiltersIamV2Api.ListIamV2IpFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2IpFilters`: IamV2IpFilterList
    fmt.Fprintf(os.Stdout, "Response from `IpFiltersIamV2Api.ListIamV2IpFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2IpFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 25]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**IamV2IpFilterList**](iam.v2.IpFilterList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


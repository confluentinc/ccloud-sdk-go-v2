# \OptInsCdxV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCdxV1OptIn**](OptInsCdxV1Api.md#GetCdxV1OptIn) | **Get** /cdx/v1/opt-in | Read the organization&#39;s stream sharing opt-in settings
[**UpdateCdxV1OptIn**](OptInsCdxV1Api.md#UpdateCdxV1OptIn) | **Patch** /cdx/v1/opt-in | Set the organization&#39;s stream sharing opt-in settings



## GetCdxV1OptIn

> CdxV1OptIn GetCdxV1OptIn(ctx).Execute()

Read the organization's stream sharing opt-in settings



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OptInsCdxV1Api.GetCdxV1OptIn(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptInsCdxV1Api.GetCdxV1OptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCdxV1OptIn`: CdxV1OptIn
    fmt.Fprintf(os.Stdout, "Response from `OptInsCdxV1Api.GetCdxV1OptIn`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCdxV1OptInRequest struct via the builder pattern


### Return type

[**CdxV1OptIn**](cdx.v1.OptIn.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCdxV1OptIn

> CdxV1OptIn UpdateCdxV1OptIn(ctx).CdxV1OptIn(cdxV1OptIn).Execute()

Set the organization's stream sharing opt-in settings



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
    cdxV1OptIn := *openapiclient.NewCdxV1OptIn() // CdxV1OptIn |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OptInsCdxV1Api.UpdateCdxV1OptIn(context.Background()).CdxV1OptIn(cdxV1OptIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptInsCdxV1Api.UpdateCdxV1OptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCdxV1OptIn`: CdxV1OptIn
    fmt.Fprintf(os.Stdout, "Response from `OptInsCdxV1Api.UpdateCdxV1OptIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCdxV1OptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cdxV1OptIn** | [**CdxV1OptIn**](CdxV1OptIn.md) |  | 

### Return type

[**CdxV1OptIn**](cdx.v1.OptIn.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


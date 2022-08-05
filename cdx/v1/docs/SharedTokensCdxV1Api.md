# \SharedTokensCdxV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RedeemCdxV1SharedToken**](SharedTokensCdxV1Api.md#RedeemCdxV1SharedToken) | **Post** /cdx/v1/shared-tokens:redeem | Redeem a Shared Token
[**ResourcesCdxV1SharedToken**](SharedTokensCdxV1Api.md#ResourcesCdxV1SharedToken) | **Post** /cdx/v1/shared-tokens:resources | Resources a Shared Token



## RedeemCdxV1SharedToken

> CdxV1RedeemShare RedeemCdxV1SharedToken(ctx).CdxV1SharedToken(cdxV1SharedToken).Execute()

Redeem a Shared Token



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
    cdxV1SharedToken := *openapiclient.NewCdxV1SharedToken() // CdxV1SharedToken |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharedTokensCdxV1Api.RedeemCdxV1SharedToken(context.Background()).CdxV1SharedToken(cdxV1SharedToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTokensCdxV1Api.RedeemCdxV1SharedToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedeemCdxV1SharedToken`: CdxV1RedeemShare
    fmt.Fprintf(os.Stdout, "Response from `SharedTokensCdxV1Api.RedeemCdxV1SharedToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRedeemCdxV1SharedTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cdxV1SharedToken** | [**CdxV1SharedToken**](CdxV1SharedToken.md) |  | 

### Return type

[**CdxV1RedeemShare**](CdxV1RedeemShare.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourcesCdxV1SharedToken

> InlineResponse200 ResourcesCdxV1SharedToken(ctx).CdxV1SharedToken(cdxV1SharedToken).Execute()

Resources a Shared Token



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
    cdxV1SharedToken := *openapiclient.NewCdxV1SharedToken() // CdxV1SharedToken |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SharedTokensCdxV1Api.ResourcesCdxV1SharedToken(context.Background()).CdxV1SharedToken(cdxV1SharedToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharedTokensCdxV1Api.ResourcesCdxV1SharedToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResourcesCdxV1SharedToken`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `SharedTokensCdxV1Api.ResourcesCdxV1SharedToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourcesCdxV1SharedTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cdxV1SharedToken** | [**CdxV1SharedToken**](CdxV1SharedToken.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \CertificateAuthoritiesIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2CertificateAuthority**](CertificateAuthoritiesIamV2Api.md#CreateIamV2CertificateAuthority) | **Post** /iam/v2/certificate-authorities | Create a Certificate Authority
[**DeleteIamV2CertificateAuthority**](CertificateAuthoritiesIamV2Api.md#DeleteIamV2CertificateAuthority) | **Delete** /iam/v2/certificate-authorities/{id} | Delete a Certificate Authority
[**GetIamV2CertificateAuthority**](CertificateAuthoritiesIamV2Api.md#GetIamV2CertificateAuthority) | **Get** /iam/v2/certificate-authorities/{id} | Read a Certificate Authority
[**ListIamV2CertificateAuthorities**](CertificateAuthoritiesIamV2Api.md#ListIamV2CertificateAuthorities) | **Get** /iam/v2/certificate-authorities | List of Certificate Authorities
[**UpdateIamV2CertificateAuthority**](CertificateAuthoritiesIamV2Api.md#UpdateIamV2CertificateAuthority) | **Put** /iam/v2/certificate-authorities/{id} | Update a Certificate Authority



## CreateIamV2CertificateAuthority

> IamV2CertificateAuthority CreateIamV2CertificateAuthority(ctx).IamV2CreateCertRequest(iamV2CreateCertRequest).Execute()

Create a Certificate Authority



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
    iamV2CreateCertRequest := *openapiclient.NewIamV2CreateCertRequest() // IamV2CreateCertRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateAuthoritiesIamV2Api.CreateIamV2CertificateAuthority(context.Background()).IamV2CreateCertRequest(iamV2CreateCertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthoritiesIamV2Api.CreateIamV2CertificateAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2CertificateAuthority`: IamV2CertificateAuthority
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthoritiesIamV2Api.CreateIamV2CertificateAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2CertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamV2CreateCertRequest** | [**IamV2CreateCertRequest**](IamV2CreateCertRequest.md) |  | 

### Return type

[**IamV2CertificateAuthority**](iam.v2.CertificateAuthority.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2CertificateAuthority

> IamV2CertificateAuthority DeleteIamV2CertificateAuthority(ctx, id).Execute()

Delete a Certificate Authority



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
    id := "id_example" // string | The unique identifier for the certificate authority.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateAuthoritiesIamV2Api.DeleteIamV2CertificateAuthority(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthoritiesIamV2Api.DeleteIamV2CertificateAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIamV2CertificateAuthority`: IamV2CertificateAuthority
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthoritiesIamV2Api.DeleteIamV2CertificateAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the certificate authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2CertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamV2CertificateAuthority**](iam.v2.CertificateAuthority.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamV2CertificateAuthority

> IamV2CertificateAuthority GetIamV2CertificateAuthority(ctx, id).Execute()

Read a Certificate Authority



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
    id := "id_example" // string | The unique identifier for the certificate authority.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateAuthoritiesIamV2Api.GetIamV2CertificateAuthority(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthoritiesIamV2Api.GetIamV2CertificateAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2CertificateAuthority`: IamV2CertificateAuthority
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthoritiesIamV2Api.GetIamV2CertificateAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the certificate authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2CertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamV2CertificateAuthority**](iam.v2.CertificateAuthority.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2CertificateAuthorities

> IamV2CertificateAuthorityList ListIamV2CertificateAuthorities(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Certificate Authorities



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
    resp, r, err := api_client.CertificateAuthoritiesIamV2Api.ListIamV2CertificateAuthorities(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthoritiesIamV2Api.ListIamV2CertificateAuthorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2CertificateAuthorities`: IamV2CertificateAuthorityList
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthoritiesIamV2Api.ListIamV2CertificateAuthorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2CertificateAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**IamV2CertificateAuthorityList**](iam.v2.CertificateAuthorityList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIamV2CertificateAuthority

> IamV2CertificateAuthority UpdateIamV2CertificateAuthority(ctx, id).IamV2UpdateCertRequest(iamV2UpdateCertRequest).Execute()

Update a Certificate Authority



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
    id := "id_example" // string | The unique identifier for the certificate authority.
    iamV2UpdateCertRequest := *openapiclient.NewIamV2UpdateCertRequest() // IamV2UpdateCertRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateAuthoritiesIamV2Api.UpdateIamV2CertificateAuthority(context.Background(), id).IamV2UpdateCertRequest(iamV2UpdateCertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthoritiesIamV2Api.UpdateIamV2CertificateAuthority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIamV2CertificateAuthority`: IamV2CertificateAuthority
    fmt.Fprintf(os.Stdout, "Response from `CertificateAuthoritiesIamV2Api.UpdateIamV2CertificateAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the certificate authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIamV2CertificateAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamV2UpdateCertRequest** | [**IamV2UpdateCertRequest**](IamV2UpdateCertRequest.md) |  | 

### Return type

[**IamV2CertificateAuthority**](iam.v2.CertificateAuthority.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


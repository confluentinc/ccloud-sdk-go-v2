# \CertificateIdentityPoolsIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2CertificateIdentityPool**](CertificateIdentityPoolsIamV2Api.md#CreateIamV2CertificateIdentityPool) | **Post** /iam/v2/certificate-authorities/{certificate_authority_id}/identity-pools | Create a Certificate Identity Pool
[**DeleteIamV2CertificateIdentityPool**](CertificateIdentityPoolsIamV2Api.md#DeleteIamV2CertificateIdentityPool) | **Delete** /iam/v2/certificate-authorities/{certificate_authority_id}/identity-pools/{id} | Delete a Certificate Identity Pool
[**GetIamV2CertificateIdentityPool**](CertificateIdentityPoolsIamV2Api.md#GetIamV2CertificateIdentityPool) | **Get** /iam/v2/certificate-authorities/{certificate_authority_id}/identity-pools/{id} | Read a Certificate Identity Pool
[**ListIamV2CertificateIdentityPools**](CertificateIdentityPoolsIamV2Api.md#ListIamV2CertificateIdentityPools) | **Get** /iam/v2/certificate-authorities/{certificate_authority_id}/identity-pools | List of Certificate Identity Pools
[**UpdateIamV2CertificateIdentityPool**](CertificateIdentityPoolsIamV2Api.md#UpdateIamV2CertificateIdentityPool) | **Put** /iam/v2/certificate-authorities/{certificate_authority_id}/identity-pools/{id} | Update a Certificate Identity Pool



## CreateIamV2CertificateIdentityPool

> IamV2CertificateIdentityPool CreateIamV2CertificateIdentityPool(ctx, certificateAuthorityId).AssignedResourceOwner(assignedResourceOwner).IamV2CertificateIdentityPool(iamV2CertificateIdentityPool).Execute()

Create a Certificate Identity Pool



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
    certificateAuthorityId := "certificateAuthorityId_example" // string | The Certificate Authority
    assignedResourceOwner := "u-a83k9b" // string | The resource_id of the principal who will be assigned resource owner on the created certificate identity pool. Principal can be group-mapping (group-xxx),  user (u-xxx), service-account (sa-xxx) or identity-pool (pool-xxx). (optional)
    iamV2CertificateIdentityPool := *openapiclient.NewIamV2CertificateIdentityPool() // IamV2CertificateIdentityPool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateIdentityPoolsIamV2Api.CreateIamV2CertificateIdentityPool(context.Background(), certificateAuthorityId).AssignedResourceOwner(assignedResourceOwner).IamV2CertificateIdentityPool(iamV2CertificateIdentityPool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateIdentityPoolsIamV2Api.CreateIamV2CertificateIdentityPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2CertificateIdentityPool`: IamV2CertificateIdentityPool
    fmt.Fprintf(os.Stdout, "Response from `CertificateIdentityPoolsIamV2Api.CreateIamV2CertificateIdentityPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateAuthorityId** | **string** | The Certificate Authority | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2CertificateIdentityPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignedResourceOwner** | **string** | The resource_id of the principal who will be assigned resource owner on the created certificate identity pool. Principal can be group-mapping (group-xxx),  user (u-xxx), service-account (sa-xxx) or identity-pool (pool-xxx). | 
 **iamV2CertificateIdentityPool** | [**IamV2CertificateIdentityPool**](IamV2CertificateIdentityPool.md) |  | 

### Return type

[**IamV2CertificateIdentityPool**](iam.v2.CertificateIdentityPool.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2CertificateIdentityPool

> IamV2CertificateIdentityPool DeleteIamV2CertificateIdentityPool(ctx, certificateAuthorityId, id).Execute()

Delete a Certificate Identity Pool



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
    certificateAuthorityId := "certificateAuthorityId_example" // string | The Certificate Authority
    id := "id_example" // string | The unique identifier for the certificate identity pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateIdentityPoolsIamV2Api.DeleteIamV2CertificateIdentityPool(context.Background(), certificateAuthorityId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateIdentityPoolsIamV2Api.DeleteIamV2CertificateIdentityPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIamV2CertificateIdentityPool`: IamV2CertificateIdentityPool
    fmt.Fprintf(os.Stdout, "Response from `CertificateIdentityPoolsIamV2Api.DeleteIamV2CertificateIdentityPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateAuthorityId** | **string** | The Certificate Authority | 
**id** | **string** | The unique identifier for the certificate identity pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2CertificateIdentityPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamV2CertificateIdentityPool**](iam.v2.CertificateIdentityPool.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIamV2CertificateIdentityPool

> IamV2CertificateIdentityPool GetIamV2CertificateIdentityPool(ctx, certificateAuthorityId, id).Execute()

Read a Certificate Identity Pool



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
    certificateAuthorityId := "certificateAuthorityId_example" // string | The Certificate Authority
    id := "id_example" // string | The unique identifier for the certificate identity pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateIdentityPoolsIamV2Api.GetIamV2CertificateIdentityPool(context.Background(), certificateAuthorityId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateIdentityPoolsIamV2Api.GetIamV2CertificateIdentityPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2CertificateIdentityPool`: IamV2CertificateIdentityPool
    fmt.Fprintf(os.Stdout, "Response from `CertificateIdentityPoolsIamV2Api.GetIamV2CertificateIdentityPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateAuthorityId** | **string** | The Certificate Authority | 
**id** | **string** | The unique identifier for the certificate identity pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2CertificateIdentityPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IamV2CertificateIdentityPool**](iam.v2.CertificateIdentityPool.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2CertificateIdentityPools

> IamV2CertificateIdentityPoolList ListIamV2CertificateIdentityPools(ctx, certificateAuthorityId).PageSize(pageSize).PageToken(pageToken).Execute()

List of Certificate Identity Pools



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
    certificateAuthorityId := "certificateAuthorityId_example" // string | The Certificate Authority
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateIdentityPoolsIamV2Api.ListIamV2CertificateIdentityPools(context.Background(), certificateAuthorityId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateIdentityPoolsIamV2Api.ListIamV2CertificateIdentityPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2CertificateIdentityPools`: IamV2CertificateIdentityPoolList
    fmt.Fprintf(os.Stdout, "Response from `CertificateIdentityPoolsIamV2Api.ListIamV2CertificateIdentityPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateAuthorityId** | **string** | The Certificate Authority | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2CertificateIdentityPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**IamV2CertificateIdentityPoolList**](iam.v2.CertificateIdentityPoolList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIamV2CertificateIdentityPool

> IamV2CertificateIdentityPool UpdateIamV2CertificateIdentityPool(ctx, certificateAuthorityId, id).IamV2CertificateIdentityPool(iamV2CertificateIdentityPool).Execute()

Update a Certificate Identity Pool



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
    certificateAuthorityId := "certificateAuthorityId_example" // string | The Certificate Authority
    id := "id_example" // string | The unique identifier for the certificate identity pool.
    iamV2CertificateIdentityPool := *openapiclient.NewIamV2CertificateIdentityPool() // IamV2CertificateIdentityPool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificateIdentityPoolsIamV2Api.UpdateIamV2CertificateIdentityPool(context.Background(), certificateAuthorityId, id).IamV2CertificateIdentityPool(iamV2CertificateIdentityPool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateIdentityPoolsIamV2Api.UpdateIamV2CertificateIdentityPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIamV2CertificateIdentityPool`: IamV2CertificateIdentityPool
    fmt.Fprintf(os.Stdout, "Response from `CertificateIdentityPoolsIamV2Api.UpdateIamV2CertificateIdentityPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateAuthorityId** | **string** | The Certificate Authority | 
**id** | **string** | The unique identifier for the certificate identity pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIamV2CertificateIdentityPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iamV2CertificateIdentityPool** | [**IamV2CertificateIdentityPool**](IamV2CertificateIdentityPool.md) |  | 

### Return type

[**IamV2CertificateIdentityPool**](iam.v2.CertificateIdentityPool.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


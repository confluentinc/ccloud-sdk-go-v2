# \DNSRecordsNetworkingV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkingV1DnsRecord**](DNSRecordsNetworkingV1Api.md#CreateNetworkingV1DnsRecord) | **Post** /networking/v1/dns-records | Create a DNS Record
[**DeleteNetworkingV1DnsRecord**](DNSRecordsNetworkingV1Api.md#DeleteNetworkingV1DnsRecord) | **Delete** /networking/v1/dns-records/{id} | Delete a DNS Record
[**GetNetworkingV1DnsRecord**](DNSRecordsNetworkingV1Api.md#GetNetworkingV1DnsRecord) | **Get** /networking/v1/dns-records/{id} | Read a DNS Record
[**ListNetworkingV1DnsRecords**](DNSRecordsNetworkingV1Api.md#ListNetworkingV1DnsRecords) | **Get** /networking/v1/dns-records | List of DNS Records
[**UpdateNetworkingV1DnsRecord**](DNSRecordsNetworkingV1Api.md#UpdateNetworkingV1DnsRecord) | **Patch** /networking/v1/dns-records/{id} | Update a DNS Record



## CreateNetworkingV1DnsRecord

> NetworkingV1DnsRecord CreateNetworkingV1DnsRecord(ctx).NetworkingV1DnsRecord(networkingV1DnsRecord).Execute()

Create a DNS Record



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
    networkingV1DnsRecord := *openapiclient.NewNetworkingV1DnsRecord() // NetworkingV1DnsRecord |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSRecordsNetworkingV1Api.CreateNetworkingV1DnsRecord(context.Background()).NetworkingV1DnsRecord(networkingV1DnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsNetworkingV1Api.CreateNetworkingV1DnsRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkingV1DnsRecord`: NetworkingV1DnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DNSRecordsNetworkingV1Api.CreateNetworkingV1DnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkingV1DnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkingV1DnsRecord** | [**NetworkingV1DnsRecord**](NetworkingV1DnsRecord.md) |  | 

### Return type

[**NetworkingV1DnsRecord**](networking.v1.DnsRecord.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkingV1DnsRecord

> DeleteNetworkingV1DnsRecord(ctx, id).Environment(environment).Execute()

Delete a DNS Record



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
    id := "id_example" // string | The unique identifier for the DNS record.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSRecordsNetworkingV1Api.DeleteNetworkingV1DnsRecord(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsNetworkingV1Api.DeleteNetworkingV1DnsRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the DNS record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkingV1DnsRecordRequest struct via the builder pattern


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


## GetNetworkingV1DnsRecord

> NetworkingV1DnsRecord GetNetworkingV1DnsRecord(ctx, id).Environment(environment).Execute()

Read a DNS Record



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
    id := "id_example" // string | The unique identifier for the DNS record.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSRecordsNetworkingV1Api.GetNetworkingV1DnsRecord(context.Background(), id).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsNetworkingV1Api.GetNetworkingV1DnsRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkingV1DnsRecord`: NetworkingV1DnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DNSRecordsNetworkingV1Api.GetNetworkingV1DnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the DNS record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkingV1DnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Scope the operation to the given environment. | 


### Return type

[**NetworkingV1DnsRecord**](networking.v1.DnsRecord.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkingV1DnsRecords

> NetworkingV1DnsRecordList ListNetworkingV1DnsRecords(ctx).Environment(environment).SpecDisplayName(specDisplayName).SpecDomain(specDomain).SpecGateway(specGateway).ResourceId(resourceId).PageSize(pageSize).PageToken(pageToken).Execute()

List of DNS Records



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
    specDomain := []string{"Inner_example"} // []string | Filter the results by exact match for spec.domain. Pass multiple times to see results matching any of the values. (optional)
    specGateway := []string{"Inner_example"} // []string | Filter the results by exact match for spec.gateway. Pass multiple times to see results matching any of the values. (optional)
    resourceId := []string{"Inner_example"} // []string | Filter the results by exact match for resource_id. Pass multiple times to see results matching any of the values. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSRecordsNetworkingV1Api.ListNetworkingV1DnsRecords(context.Background()).Environment(environment).SpecDisplayName(specDisplayName).SpecDomain(specDomain).SpecGateway(specGateway).ResourceId(resourceId).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsNetworkingV1Api.ListNetworkingV1DnsRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkingV1DnsRecords`: NetworkingV1DnsRecordList
    fmt.Fprintf(os.Stdout, "Response from `DNSRecordsNetworkingV1Api.ListNetworkingV1DnsRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkingV1DnsRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** | Filter the results by exact match for environment. | 
 **specDisplayName** | **[]string** | Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values. | 
 **specDomain** | **[]string** | Filter the results by exact match for spec.domain. Pass multiple times to see results matching any of the values. | 
 **specGateway** | **[]string** | Filter the results by exact match for spec.gateway. Pass multiple times to see results matching any of the values. | 
 **resourceId** | **[]string** | Filter the results by exact match for resource_id. Pass multiple times to see results matching any of the values. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**NetworkingV1DnsRecordList**](networking.v1.DnsRecordList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkingV1DnsRecord

> NetworkingV1DnsRecord UpdateNetworkingV1DnsRecord(ctx, id).NetworkingV1DnsRecordUpdate(networkingV1DnsRecordUpdate).Execute()

Update a DNS Record



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
    id := "id_example" // string | The unique identifier for the DNS record.
    networkingV1DnsRecordUpdate := *openapiclient.NewNetworkingV1DnsRecordUpdate() // NetworkingV1DnsRecordUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DNSRecordsNetworkingV1Api.UpdateNetworkingV1DnsRecord(context.Background(), id).NetworkingV1DnsRecordUpdate(networkingV1DnsRecordUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSRecordsNetworkingV1Api.UpdateNetworkingV1DnsRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkingV1DnsRecord`: NetworkingV1DnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DNSRecordsNetworkingV1Api.UpdateNetworkingV1DnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the DNS record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkingV1DnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkingV1DnsRecordUpdate** | [**NetworkingV1DnsRecordUpdate**](NetworkingV1DnsRecordUpdate.md) |  | 

### Return type

[**NetworkingV1DnsRecord**](networking.v1.DnsRecord.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


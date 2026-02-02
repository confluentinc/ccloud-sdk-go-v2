# \EndpointsEndpointV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEndpointV1Endpoints**](EndpointsEndpointV1Api.md#ListEndpointV1Endpoints) | **Get** /endpoint/v1/endpoints | List of Endpoints



## ListEndpointV1Endpoints

> EndpointV1EndpointList ListEndpointV1Endpoints(ctx).Service(service).Environment(environment).Cloud(cloud).Region(region).IsPrivate(isPrivate).Resource(resource).PageSize(pageSize).PageToken(pageToken).Execute()

List of Endpoints



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
    service := "KAFKA" // string | Filter the results by exact match for service.
    environment := "env-00000" // string | Filter the results by exact match for environment.
    cloud := "AWS" // string | Filter the results by exact match for cloud. (optional)
    region := "us-west-2" // string | Filter the results by exact match for region. (optional)
    isPrivate := true // bool | Filter the results by whether the endpoint is private (true) or public (false). If not specified, returns both private and public endpoints.  (optional)
    resource := "lkc-abc123" // string | Filter the results by exact match for resource. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 100)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EndpointsEndpointV1Api.ListEndpointV1Endpoints(context.Background()).Service(service).Environment(environment).Cloud(cloud).Region(region).IsPrivate(isPrivate).Resource(resource).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndpointsEndpointV1Api.ListEndpointV1Endpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpointV1Endpoints`: EndpointV1EndpointList
    fmt.Fprintf(os.Stdout, "Response from `EndpointsEndpointV1Api.ListEndpointV1Endpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointV1EndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Filter the results by exact match for service. | 
 **environment** | **string** | Filter the results by exact match for environment. | 
 **cloud** | **string** | Filter the results by exact match for cloud. | 
 **region** | **string** | Filter the results by exact match for region. | 
 **isPrivate** | **bool** | Filter the results by whether the endpoint is private (true) or public (false). If not specified, returns both private and public endpoints.  | 
 **resource** | **string** | Filter the results by exact match for resource. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 100]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**EndpointV1EndpointList**](endpoint.v1.EndpointList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


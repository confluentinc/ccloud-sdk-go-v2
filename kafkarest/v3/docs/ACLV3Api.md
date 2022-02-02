# \ACLV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaV3Acls**](ACLV3Api.md#CreateKafkaV3Acls) | **Post** /kafka/v3/clusters/{cluster_id}/acls | Create ACLs
[**DeleteKafkaV3Acls**](ACLV3Api.md#DeleteKafkaV3Acls) | **Delete** /kafka/v3/clusters/{cluster_id}/acls | Delete ACLs
[**GetKafkaV3Acls**](ACLV3Api.md#GetKafkaV3Acls) | **Get** /kafka/v3/clusters/{cluster_id}/acls | Search ACLs



## CreateKafkaV3Acls

> CreateKafkaV3Acls(ctx, clusterId).CreateAclRequestData(createAclRequestData).Execute()

Create ACLs



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
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    createAclRequestData := *openapiclient.NewCreateAclRequestData(openapiclient.AclResourceType("UNKNOWN"), "ResourceName_example", "PatternType_example", "Principal_example", "Host_example", "Operation_example", "Permission_example") // CreateAclRequestData | The ACL creation request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ACLV3Api.CreateKafkaV3Acls(context.Background(), clusterId).CreateAclRequestData(createAclRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLV3Api.CreateKafkaV3Acls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKafkaV3AclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAclRequestData** | [**CreateAclRequestData**](CreateAclRequestData.md) | The ACL creation request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaV3Acls

> InlineResponse200 DeleteKafkaV3Acls(ctx, clusterId).ResourceType(resourceType).ResourceName(resourceName).PatternType(patternType).Principal(principal).Host(host).Operation(operation).Permission(permission).Execute()

Delete ACLs



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
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    resourceType := openapiclient.AclResourceType("UNKNOWN") // AclResourceType | The ACL resource type. (optional)
    resourceName := "resourceName_example" // string | The ACL resource name. (optional)
    patternType := "patternType_example" // string | The ACL pattern type. (optional)
    principal := "principal_example" // string | The ACL principal. (optional)
    host := "host_example" // string | The ACL host. (optional)
    operation := "operation_example" // string | The ACL operation. (optional)
    permission := "permission_example" // string | The ACL permission. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ACLV3Api.DeleteKafkaV3Acls(context.Background(), clusterId).ResourceType(resourceType).ResourceName(resourceName).PatternType(patternType).Principal(principal).Host(host).Operation(operation).Permission(permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLV3Api.DeleteKafkaV3Acls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteKafkaV3Acls`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `ACLV3Api.DeleteKafkaV3Acls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKafkaV3AclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceType** | [**AclResourceType**](AclResourceType.md) | The ACL resource type. | 
 **resourceName** | **string** | The ACL resource name. | 
 **patternType** | **string** | The ACL pattern type. | 
 **principal** | **string** | The ACL principal. | 
 **host** | **string** | The ACL host. | 
 **operation** | **string** | The ACL operation. | 
 **permission** | **string** | The ACL permission. | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3Acls

> AclDataList GetKafkaV3Acls(ctx, clusterId).ResourceType(resourceType).ResourceName(resourceName).PatternType(patternType).Principal(principal).Host(host).Operation(operation).Permission(permission).Execute()

Search ACLs



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
    clusterId := "cluster-1" // string | The Kafka cluster ID.
    resourceType := openapiclient.AclResourceType("UNKNOWN") // AclResourceType | The ACL resource type. (optional)
    resourceName := "resourceName_example" // string | The ACL resource name. (optional)
    patternType := "patternType_example" // string | The ACL pattern type. (optional)
    principal := "principal_example" // string | The ACL principal. (optional)
    host := "host_example" // string | The ACL host. (optional)
    operation := "operation_example" // string | The ACL operation. (optional)
    permission := "permission_example" // string | The ACL permission. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ACLV3Api.GetKafkaV3Acls(context.Background(), clusterId).ResourceType(resourceType).ResourceName(resourceName).PatternType(patternType).Principal(principal).Host(host).Operation(operation).Permission(permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACLV3Api.GetKafkaV3Acls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3Acls`: AclDataList
    fmt.Fprintf(os.Stdout, "Response from `ACLV3Api.GetKafkaV3Acls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3AclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceType** | [**AclResourceType**](AclResourceType.md) | The ACL resource type. | 
 **resourceName** | **string** | The ACL resource name. | 
 **patternType** | **string** | The ACL pattern type. | 
 **principal** | **string** | The ACL principal. | 
 **host** | **string** | The ACL host. | 
 **operation** | **string** | The ACL operation. | 
 **permission** | **string** | The ACL permission. | 

### Return type

[**AclDataList**](AclDataList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


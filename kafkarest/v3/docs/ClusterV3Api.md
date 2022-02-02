# \ClusterV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKafkaV3Cluster**](ClusterV3Api.md#GetKafkaV3Cluster) | **Get** /kafka/v3/clusters/{cluster_id} | Get Cluster



## GetKafkaV3Cluster

> ClusterData GetKafkaV3Cluster(ctx, clusterId).Execute()

Get Cluster



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV3Api.GetKafkaV3Cluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV3Api.GetKafkaV3Cluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKafkaV3Cluster`: ClusterData
    fmt.Fprintf(os.Stdout, "Response from `ClusterV3Api.GetKafkaV3Cluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKafkaV3ClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterData**](ClusterData.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


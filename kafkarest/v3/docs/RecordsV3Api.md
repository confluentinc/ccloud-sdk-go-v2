# \RecordsV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProduceRecords**](RecordsV3Api.md#ProduceRecords) | **Post** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/records | Produce records to the given topic.



## ProduceRecords

> ProduceResponse ProduceRecords(ctx, clusterId, topicName).ProduceRequest(produceRequest).Execute()

Produce records to the given topic.



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
    topicName := "topic-1" // string | The topic name.
    produceRequest := *openapiclient.NewProduceRequest() // ProduceRequest | A single record to be produced to Kafka. To produce multiple records on the same connection, simply concatenate all the records, e.g.: {\"partition_id\":1}{\"partition_id\":2}. Delivery reports will be concatenated on the same order as the records are sent. See examples for the options available. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordsV3Api.ProduceRecords(context.Background(), clusterId, topicName).ProduceRequest(produceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsV3Api.ProduceRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProduceRecords`: ProduceResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsV3Api.ProduceRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProduceRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **produceRequest** | [**ProduceRequest**](ProduceRequest.md) | A single record to be produced to Kafka. To produce multiple records on the same connection, simply concatenate all the records, e.g.: {\&quot;partition_id\&quot;:1}{\&quot;partition_id\&quot;:2}. Delivery reports will be concatenated on the same order as the records are sent. See examples for the options available. | 

### Return type

[**ProduceResponse**](ProduceResponse.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


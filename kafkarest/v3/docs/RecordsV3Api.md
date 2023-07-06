# \RecordsV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProduceRecord**](RecordsV3Api.md#ProduceRecord) | **Post** /kafka/v3/clusters/{cluster_id}/topics/{topic_name}/records | Produce Records



## ProduceRecord

> ProduceResponse ProduceRecord(ctx, clusterId, topicName).ProduceRequest(produceRequest).Execute()

Produce Records



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
    produceRequest := *openapiclient.NewProduceRequest() // ProduceRequest | A single record to be produced to Kafka. To produce multiple records in the same request, simply concatenate the records. The delivery reports are concatenated in the same order as the records are sent. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecordsV3Api.ProduceRecord(context.Background(), clusterId, topicName).ProduceRequest(produceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsV3Api.ProduceRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProduceRecord`: ProduceResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsV3Api.ProduceRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Kafka cluster ID. | 
**topicName** | **string** | The topic name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProduceRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **produceRequest** | [**ProduceRequest**](ProduceRequest.md) | A single record to be produced to Kafka. To produce multiple records in the same request, simply concatenate the records. The delivery reports are concatenated in the same order as the records are sent. | 

### Return type

[**ProduceResponse**](ProduceResponse.md)

### Authorization

[api-key](../README.md#api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


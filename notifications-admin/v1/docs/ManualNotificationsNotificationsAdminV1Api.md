# \ManualNotificationsNotificationsAdminV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendNotificationsAdminV1ManualNotification**](ManualNotificationsNotificationsAdminV1Api.md#SendNotificationsAdminV1ManualNotification) | **Post** /notifications-admin/v1/notifications/send | Send manual notification.



## SendNotificationsAdminV1ManualNotification

> InlineResponse207 SendNotificationsAdminV1ManualNotification(ctx).InlineObject(inlineObject).Execute()

Send manual notification.



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
    inlineObject := *openapiclient.NewInlineObject("There is a problem with Kafka cluster lkc-23ds8.", []int32{int32(123)}, false) // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManualNotificationsNotificationsAdminV1Api.SendNotificationsAdminV1ManualNotification(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualNotificationsNotificationsAdminV1Api.SendNotificationsAdminV1ManualNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendNotificationsAdminV1ManualNotification`: InlineResponse207
    fmt.Fprintf(os.Stdout, "Response from `ManualNotificationsNotificationsAdminV1Api.SendNotificationsAdminV1ManualNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendNotificationsAdminV1ManualNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**InlineResponse207**](InlineResponse207.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


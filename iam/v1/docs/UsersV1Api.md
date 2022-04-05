# \UsersV1Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListV1Users**](UsersV1Api.md#ListV1Users) | **Get** /users | List of Users



## ListV1Users

> V1ServiceAccountList ListV1Users(ctx).Execute()

List of Users



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersV1Api.ListV1Users(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersV1Api.ListV1Users``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListV1Users`: V1ServiceAccountList
    fmt.Fprintf(os.Stdout, "Response from `UsersV1Api.ListV1Users`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListV1UsersRequest struct via the builder pattern


### Return type

[**V1ServiceAccountList**](v1.ServiceAccountList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


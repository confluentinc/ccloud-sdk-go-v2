# \UserInvitationsV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateV2UserInvitation**](UserInvitationsV2Api.md#CreateV2UserInvitation) | **Post** /v2/user-invitations | Create a User Invitation



## CreateV2UserInvitation

> V2UserInvitation CreateV2UserInvitation(ctx).V2UserInvitation(v2UserInvitation).Execute()

Create a User Invitation



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
    v2UserInvitation := *openapiclient.NewV2UserInvitation() // V2UserInvitation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserInvitationsV2Api.CreateV2UserInvitation(context.Background()).V2UserInvitation(v2UserInvitation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInvitationsV2Api.CreateV2UserInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateV2UserInvitation`: V2UserInvitation
    fmt.Fprintf(os.Stdout, "Response from `UserInvitationsV2Api.CreateV2UserInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateV2UserInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2UserInvitation** | [**V2UserInvitation**](V2UserInvitation.md) |  | 

### Return type

[**V2UserInvitation**](v2.UserInvitation.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


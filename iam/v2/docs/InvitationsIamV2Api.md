# \InvitationsIamV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIamV2Invitation**](InvitationsIamV2Api.md#CreateIamV2Invitation) | **Post** /iam/v2/invitations | Create an Invitation
[**DeleteIamV2Invitation**](InvitationsIamV2Api.md#DeleteIamV2Invitation) | **Delete** /iam/v2/invitations/{id} | Delete an Invitation
[**GetIamV2Invitation**](InvitationsIamV2Api.md#GetIamV2Invitation) | **Get** /iam/v2/invitations/{id} | Read an Invitation
[**ListIamV2Invitations**](InvitationsIamV2Api.md#ListIamV2Invitations) | **Get** /iam/v2/invitations | List of Invitations



## CreateIamV2Invitation

> IamV2Invitation CreateIamV2Invitation(ctx).IamV2Invitation(iamV2Invitation).Execute()

Create an Invitation



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
    iamV2Invitation := *openapiclient.NewIamV2Invitation() // IamV2Invitation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsIamV2Api.CreateIamV2Invitation(context.Background()).IamV2Invitation(iamV2Invitation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsIamV2Api.CreateIamV2Invitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIamV2Invitation`: IamV2Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsIamV2Api.CreateIamV2Invitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIamV2InvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamV2Invitation** | [**IamV2Invitation**](IamV2Invitation.md) |  | 

### Return type

[**IamV2Invitation**](iam.v2.Invitation.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIamV2Invitation

> DeleteIamV2Invitation(ctx, id).Execute()

Delete an Invitation



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
    id := "id_example" // string | The unique identifier for the invitation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsIamV2Api.DeleteIamV2Invitation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsIamV2Api.DeleteIamV2Invitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIamV2InvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetIamV2Invitation

> IamV2Invitation GetIamV2Invitation(ctx, id).Execute()

Read an Invitation



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
    id := "id_example" // string | The unique identifier for the invitation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsIamV2Api.GetIamV2Invitation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsIamV2Api.GetIamV2Invitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIamV2Invitation`: IamV2Invitation
    fmt.Fprintf(os.Stdout, "Response from `InvitationsIamV2Api.GetIamV2Invitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIamV2InvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamV2Invitation**](iam.v2.Invitation.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIamV2Invitations

> IamV2InvitationList ListIamV2Invitations(ctx).Email(email).Status(status).User(user).Creator(creator).PageSize(pageSize).PageToken(pageToken).Execute()

List of Invitations



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
    email := "johndoe@confluent.io" // string | Filter the results by exact match for email. (optional)
    status := "INVITE_STATUS_SENT" // string | Filter the results by exact match for status. (optional)
    user := "u-j93dy8" // string | Filter the results by exact match for user. (optional)
    creator := "u-m2r9o7" // string | Filter the results by exact match for creator. (optional)
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvitationsIamV2Api.ListIamV2Invitations(context.Background()).Email(email).Status(status).User(user).Creator(creator).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvitationsIamV2Api.ListIamV2Invitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIamV2Invitations`: IamV2InvitationList
    fmt.Fprintf(os.Stdout, "Response from `InvitationsIamV2Api.ListIamV2Invitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIamV2InvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Filter the results by exact match for email. | 
 **status** | **string** | Filter the results by exact match for status. | 
 **user** | **string** | Filter the results by exact match for user. | 
 **creator** | **string** | Filter the results by exact match for creator. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**IamV2InvitationList**](iam.v2.InvitationList.md)

### Authorization

[cloud-api-key](../README.md#cloud-api-key), [confluent-sts-access-token](../README.md#confluent-sts-access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


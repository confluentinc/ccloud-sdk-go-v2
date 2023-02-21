# \OrganizationsOrgV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgV2Organization**](OrganizationsOrgV2Api.md#GetOrgV2Organization) | **Get** /org/v2/organizations/{id} | Read an Organization
[**ListOrgV2Organizations**](OrganizationsOrgV2Api.md#ListOrgV2Organizations) | **Get** /org/v2/organizations | List of Organizations
[**UpdateOrgV2Organization**](OrganizationsOrgV2Api.md#UpdateOrgV2Organization) | **Patch** /org/v2/organizations/{id} | Update an Organization



## GetOrgV2Organization

> OrgV2Organization GetOrgV2Organization(ctx, id).Execute()

Read an Organization



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
    id := "id_example" // string | The unique identifier for the organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsOrgV2Api.GetOrgV2Organization(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsOrgV2Api.GetOrgV2Organization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgV2Organization`: OrgV2Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsOrgV2Api.GetOrgV2Organization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgV2OrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgV2Organization**](org.v2.Organization.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgV2Organizations

> OrgV2OrganizationList ListOrgV2Organizations(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Organizations



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
    pageSize := int32(56) // int32 | A pagination size for collection requests. (optional) (default to 10)
    pageToken := "pageToken_example" // string | An opaque pagination token for collection requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsOrgV2Api.ListOrgV2Organizations(context.Background()).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsOrgV2Api.ListOrgV2Organizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgV2Organizations`: OrgV2OrganizationList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsOrgV2Api.ListOrgV2Organizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrgV2OrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**OrgV2OrganizationList**](org.v2.OrganizationList.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgV2Organization

> OrgV2Organization UpdateOrgV2Organization(ctx, id).OrgV2Organization(orgV2Organization).Execute()

Update an Organization



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
    id := "id_example" // string | The unique identifier for the organization.
    orgV2Organization := *openapiclient.NewOrgV2Organization() // OrgV2Organization |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsOrgV2Api.UpdateOrgV2Organization(context.Background(), id).OrgV2Organization(orgV2Organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsOrgV2Api.UpdateOrgV2Organization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgV2Organization`: OrgV2Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsOrgV2Api.UpdateOrgV2Organization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgV2OrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgV2Organization** | [**OrgV2Organization**](OrgV2Organization.md) |  | 

### Return type

[**OrgV2Organization**](org.v2.Organization.md)

### Authorization

[api-key](../README.md#api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


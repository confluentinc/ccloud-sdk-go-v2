# \OrganizationsPartnerV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPartnerV2Organization**](OrganizationsPartnerV2Api.md#GetPartnerV2Organization) | **Get** /partner/v2/organizations/{id} | Read an Organization
[**ListPartnerV2Organizations**](OrganizationsPartnerV2Api.md#ListPartnerV2Organizations) | **Get** /partner/v2/organizations | List of Organizations



## GetPartnerV2Organization

> PartnerV2Organization GetPartnerV2Organization(ctx, id).Execute()

Read an Organization



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerV2OrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PartnerV2Organization**](partner.v2.Organization.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartnerV2Organizations

> PartnerV2OrganizationList ListPartnerV2Organizations(ctx).PageSize(pageSize).PageToken(pageToken).Execute()

List of Organizations



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPartnerV2OrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**PartnerV2OrganizationList**](partner.v2.OrganizationList.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


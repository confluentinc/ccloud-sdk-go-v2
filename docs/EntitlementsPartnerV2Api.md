# \EntitlementsPartnerV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePartnerV2Entitlement**](EntitlementsPartnerV2Api.md#CreatePartnerV2Entitlement) | **Post** /partner/v2/entitlements | Create an Entitlement
[**GetPartnerV2Entitlement**](EntitlementsPartnerV2Api.md#GetPartnerV2Entitlement) | **Get** /partner/v2/entitlements/{id} | Read an Entitlement
[**ListPartnerV2Entitlements**](EntitlementsPartnerV2Api.md#ListPartnerV2Entitlements) | **Get** /partner/v2/entitlements | List of Entitlements



## CreatePartnerV2Entitlement

> PartnerV2Entitlement CreatePartnerV2Entitlement(ctx).PartnerV2Entitlement(partnerV2Entitlement).Execute()

Create an Entitlement



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartnerV2EntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partnerV2Entitlement** | [**PartnerV2Entitlement**](PartnerV2Entitlement.md) |  | 

### Return type

[**PartnerV2Entitlement**](partner.v2.Entitlement.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartnerV2Entitlement

> PartnerV2Entitlement GetPartnerV2Entitlement(ctx, id).OrganizationId(organizationId).Execute()

Read an Entitlement



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier for the entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerV2EntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **string** | Scope the operation to the given organization.id. | 

### Return type

[**PartnerV2Entitlement**](partner.v2.Entitlement.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartnerV2Entitlements

> PartnerV2EntitlementList ListPartnerV2Entitlements(ctx).OrganizationId(organizationId).PageSize(pageSize).PageToken(pageToken).Execute()

List of Entitlements



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPartnerV2EntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | Filter the results by exact match for organization.id. | 
 **pageSize** | **int32** | A pagination size for collection requests. | [default to 10]
 **pageToken** | **string** | An opaque pagination token for collection requests. | 

### Return type

[**PartnerV2EntitlementList**](partner.v2.EntitlementList.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


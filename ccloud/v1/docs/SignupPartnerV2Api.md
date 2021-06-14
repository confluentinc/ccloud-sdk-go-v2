# \SignupPartnerV2Api

All URIs are relative to *https://api.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSignup**](SignupPartnerV2Api.md#ActivateSignup) | **Post** /partner/v2/signup/activate | Activate an Incomplete Signup
[**Signup**](SignupPartnerV2Api.md#Signup) | **Post** /partner/v2/signup | Signup an Organization on behalf of a Customer



## ActivateSignup

> PartnerSignupResponse ActivateSignup(ctx).ActivatePartnerSignupRequest(activatePartnerSignupRequest).Execute()

Activate an Incomplete Signup



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activatePartnerSignupRequest** | [**ActivatePartnerSignupRequest**](ActivatePartnerSignupRequest.md) | A JSON object containing signup information | 

### Return type

[**PartnerSignupResponse**](PartnerSignupResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Signup

> PartnerSignupResponse Signup(ctx).PartnerSignupRequest(partnerSignupRequest).Execute()

Signup an Organization on behalf of a Customer



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partnerSignupRequest** | [**PartnerSignupRequest**](PartnerSignupRequest.md) | A JSON object containing signup information | 

### Return type

[**PartnerSignupResponse**](PartnerSignupResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


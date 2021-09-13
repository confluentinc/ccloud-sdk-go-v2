# \ACLV3Api

All URIs are relative to *https://pkc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKafkaV3Acls**](ACLV3Api.md#CreateKafkaV3Acls) | **Post** /kafka/v3/clusters/{cluster_id}/acls | Create ACLs
[**DeleteKafkaV3Acls**](ACLV3Api.md#DeleteKafkaV3Acls) | **Delete** /kafka/v3/clusters/{cluster_id}/acls | Delete ACLs
[**GetKafkaV3Acls**](ACLV3Api.md#GetKafkaV3Acls) | **Get** /kafka/v3/clusters/{cluster_id}/acls | Search ACLs



## CreateKafkaV3Acls

> CreateKafkaV3Acls(ctx, clusterId, optional)

Create ACLs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Creates an ACL.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***CreateKafkaV3AclsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKafkaV3AclsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAclRequestData** | [**optional.Interface of CreateAclRequestData**](CreateAclRequestData.md)| The ACL creation request. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKafkaV3Acls

> InlineResponse200 DeleteKafkaV3Acls(ctx, clusterId, optional)

Delete ACLs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Deletes the list of ACLs that matches the search criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***DeleteKafkaV3AclsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteKafkaV3AclsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceType** | [**optional.Interface of AclResourceType**](.md)| The ACL resource type. | 
 **resourceName** | **optional.String**| The ACL resource name. | 
 **patternType** | **optional.String**| The ACL pattern type. | 
 **principal** | **optional.String**| The ACL principal. | 
 **host** | **optional.String**| The ACL host. | 
 **operation** | **optional.String**| The ACL operation. | 
 **permission** | **optional.String**| The ACL permission. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKafkaV3Acls

> AclDataList GetKafkaV3Acls(ctx, clusterId, optional)

Search ACLs

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)  Returns a list of ACLs that match the search criteria.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string**| The Kafka cluster ID. | 
 **optional** | ***GetKafkaV3AclsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetKafkaV3AclsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceType** | [**optional.Interface of AclResourceType**](.md)| The ACL resource type. | 
 **resourceName** | **optional.String**| The ACL resource name. | 
 **patternType** | **optional.String**| The ACL pattern type. | 
 **principal** | **optional.String**| The ACL principal. | 
 **host** | **optional.String**| The ACL host. | 
 **operation** | **optional.String**| The ACL operation. | 
 **permission** | **optional.String**| The ACL permission. | 

### Return type

[**AclDataList**](AclDataList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


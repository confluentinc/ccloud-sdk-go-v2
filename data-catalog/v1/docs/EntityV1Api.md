# \EntityV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBusinessMetadata**](EntityV1Api.md#CreateBusinessMetadata) | **Post** /catalog/v1/entity/businessmetadata | Bulk Create Business Metadata
[**CreateTags**](EntityV1Api.md#CreateTags) | **Post** /catalog/v1/entity/tags | Bulk Create Tags
[**DeleteBusinessMetadata**](EntityV1Api.md#DeleteBusinessMetadata) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/businessmetadata/{bmName} | Delete a Business Metadata for an Entity
[**DeleteTag**](EntityV1Api.md#DeleteTag) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags/{tagName} | Delete a Tag for an Entity
[**GetBusinessMetadata**](EntityV1Api.md#GetBusinessMetadata) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/businessmetadata | Read Business Metadata for an Entity
[**GetByUniqueAttributes**](EntityV1Api.md#GetByUniqueAttributes) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName} | Read an Entity
[**GetTags**](EntityV1Api.md#GetTags) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags | Read Tags for an Entity
[**PartialEntityUpdate**](EntityV1Api.md#PartialEntityUpdate) | **Put** /catalog/v1/entity | Update an Entity Attribute
[**RepublishAllEntities**](EntityV1Api.md#RepublishAllEntities) | **Post** /catalog/v1/entity-notifications-snapshot | Republish all entities events.
[**UpdateBusinessMetadata**](EntityV1Api.md#UpdateBusinessMetadata) | **Put** /catalog/v1/entity/businessmetadata | Bulk Update Business Metadata
[**UpdateTags**](EntityV1Api.md#UpdateTags) | **Put** /catalog/v1/entity/tags | Bulk Update Tags



## CreateBusinessMetadata

> []BusinessMetadataResponse CreateBusinessMetadata(ctx).BusinessMetadata(businessMetadata).Execute()

Bulk Create Business Metadata



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
    businessMetadata := []openapiclient.BusinessMetadata{*openapiclient.NewBusinessMetadata()} // []BusinessMetadata | The business metadata (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.CreateBusinessMetadata(context.Background()).BusinessMetadata(businessMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.CreateBusinessMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBusinessMetadata`: []BusinessMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.CreateBusinessMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBusinessMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessMetadata** | [**[]BusinessMetadata**](BusinessMetadata.md) | The business metadata | 

### Return type

[**[]BusinessMetadataResponse**](BusinessMetadataResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTags

> []TagResponse CreateTags(ctx).Tag(tag).Execute()

Bulk Create Tags



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
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | The tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.CreateTags(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.CreateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTags`: []TagResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.CreateTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**[]Tag**](Tag.md) | The tags | 

### Return type

[**[]TagResponse**](TagResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBusinessMetadata

> DeleteBusinessMetadata(ctx, typeName, qualifiedName, bmName).Execute()

Delete a Business Metadata for an Entity



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
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity
    bmName := "bmName_example" // string | The name of the business metadata

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.DeleteBusinessMetadata(context.Background(), typeName, qualifiedName, bmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.DeleteBusinessMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 
**bmName** | **string** | The name of the business metadata | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBusinessMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(ctx, typeName, qualifiedName, tagName).Execute()

Delete a Tag for an Entity



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
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity
    tagName := "tagName_example" // string | The name of the tag

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.DeleteTag(context.Background(), typeName, qualifiedName, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 
**tagName** | **string** | The name of the tag | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessMetadata

> []BusinessMetadataResponse GetBusinessMetadata(ctx, typeName, qualifiedName).Execute()

Read Business Metadata for an Entity



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
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.GetBusinessMetadata(context.Background(), typeName, qualifiedName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.GetBusinessMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessMetadata`: []BusinessMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.GetBusinessMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]BusinessMetadataResponse**](BusinessMetadataResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByUniqueAttributes

> EntityWithExtInfo GetByUniqueAttributes(ctx, typeName, qualifiedName).MinExtInfo(minExtInfo).IgnoreRelationships(ignoreRelationships).Execute()

Read an Entity



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
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity
    minExtInfo := true // bool | Whether to populate on header and schema attributes (optional) (default to false)
    ignoreRelationships := true // bool | Whether to ignore relationships (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.GetByUniqueAttributes(context.Background(), typeName, qualifiedName).MinExtInfo(minExtInfo).IgnoreRelationships(ignoreRelationships).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.GetByUniqueAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByUniqueAttributes`: EntityWithExtInfo
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.GetByUniqueAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByUniqueAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **minExtInfo** | **bool** | Whether to populate on header and schema attributes | [default to false]
 **ignoreRelationships** | **bool** | Whether to ignore relationships | [default to false]

### Return type

[**EntityWithExtInfo**](EntityWithExtInfo.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> []TagResponse GetTags(ctx, typeName, qualifiedName).Execute()

Read Tags for an Entity



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
    typeName := "typeName_example" // string | The type of the entity
    qualifiedName := "qualifiedName_example" // string | The qualified name of the entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.GetTags(context.Background(), typeName, qualifiedName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: []TagResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeName** | **string** | The type of the entity | 
**qualifiedName** | **string** | The qualified name of the entity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TagResponse**](TagResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialEntityUpdate

> EntityPartialUpdateResponse PartialEntityUpdate(ctx).EntityWithExtInfo(entityWithExtInfo).Execute()

Update an Entity Attribute



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
    entityWithExtInfo := *openapiclient.NewEntityWithExtInfo() // EntityWithExtInfo | The entity to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.PartialEntityUpdate(context.Background()).EntityWithExtInfo(entityWithExtInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.PartialEntityUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialEntityUpdate`: EntityPartialUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.PartialEntityUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPartialEntityUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityWithExtInfo** | [**EntityWithExtInfo**](EntityWithExtInfo.md) | The entity to update | 

### Return type

[**EntityPartialUpdateResponse**](EntityPartialUpdateResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepublishAllEntities

> RepublishAllEntities(ctx).Execute()

Republish all entities events.



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
    resp, r, err := api_client.EntityV1Api.RepublishAllEntities(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.RepublishAllEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRepublishAllEntitiesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBusinessMetadata

> []BusinessMetadataResponse UpdateBusinessMetadata(ctx).BusinessMetadata(businessMetadata).Execute()

Bulk Update Business Metadata



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
    businessMetadata := []openapiclient.BusinessMetadata{*openapiclient.NewBusinessMetadata()} // []BusinessMetadata | The business metadata (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.UpdateBusinessMetadata(context.Background()).BusinessMetadata(businessMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.UpdateBusinessMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBusinessMetadata`: []BusinessMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.UpdateBusinessMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBusinessMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessMetadata** | [**[]BusinessMetadata**](BusinessMetadata.md) | The business metadata | 

### Return type

[**[]BusinessMetadataResponse**](BusinessMetadataResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTags

> []TagResponse UpdateTags(ctx).Tag(tag).Execute()

Bulk Update Tags



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
    tag := []openapiclient.Tag{*openapiclient.NewTag()} // []Tag | The tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntityV1Api.UpdateTags(context.Background()).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntityV1Api.UpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTags`: []TagResponse
    fmt.Fprintf(os.Stdout, "Response from `EntityV1Api.UpdateTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**[]Tag**](Tag.md) | The tags | 

### Return type

[**[]TagResponse**](TagResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


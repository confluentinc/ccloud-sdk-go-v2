# \TypesV1Api

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBusinessMetadataDefs**](TypesV1Api.md#CreateBusinessMetadataDefs) | **Post** /catalog/v1/types/businessmetadatadefs | Bulk Create Business Metadata Definitions
[**CreateTagDefs**](TypesV1Api.md#CreateTagDefs) | **Post** /catalog/v1/types/tagdefs | Bulk Create Tag Definitions
[**DeleteBusinessMetadataDef**](TypesV1Api.md#DeleteBusinessMetadataDef) | **Delete** /catalog/v1/types/businessmetadatadefs/{bmName} | Delete Business Metadata Definition
[**DeleteTagDef**](TypesV1Api.md#DeleteTagDef) | **Delete** /catalog/v1/types/tagdefs/{tagName} | Delete Tag Definition
[**GetAllBusinessMetadataDefs**](TypesV1Api.md#GetAllBusinessMetadataDefs) | **Get** /catalog/v1/types/businessmetadatadefs | Bulk Read Business Metadata Definitions
[**GetAllTagDefs**](TypesV1Api.md#GetAllTagDefs) | **Get** /catalog/v1/types/tagdefs | Bulk Read Tag Definitions
[**GetBusinessMetadataDefByName**](TypesV1Api.md#GetBusinessMetadataDefByName) | **Get** /catalog/v1/types/businessmetadatadefs/{bmName} | Read Business Metadata Definition
[**GetTagDefByName**](TypesV1Api.md#GetTagDefByName) | **Get** /catalog/v1/types/tagdefs/{tagName} | Read Tag Definition
[**UpdateBusinessMetadataDefs**](TypesV1Api.md#UpdateBusinessMetadataDefs) | **Put** /catalog/v1/types/businessmetadatadefs | Bulk Update Business Metadata Definitions
[**UpdateTagDefs**](TypesV1Api.md#UpdateTagDefs) | **Put** /catalog/v1/types/tagdefs | Bulk Update Tag Definitions



## CreateBusinessMetadataDefs

> []BusinessMetadataDefResponse CreateBusinessMetadataDefs(ctx).BusinessMetadataDef(businessMetadataDef).Execute()

Bulk Create Business Metadata Definitions



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
    businessMetadataDef := []openapiclient.BusinessMetadataDef{*openapiclient.NewBusinessMetadataDef()} // []BusinessMetadataDef | The business metadata definitions to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.CreateBusinessMetadataDefs(context.Background()).BusinessMetadataDef(businessMetadataDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.CreateBusinessMetadataDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBusinessMetadataDefs`: []BusinessMetadataDefResponse
    fmt.Fprintf(os.Stdout, "Response from `TypesV1Api.CreateBusinessMetadataDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBusinessMetadataDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessMetadataDef** | [**[]BusinessMetadataDef**](BusinessMetadataDef.md) | The business metadata definitions to create | 

### Return type

[**[]BusinessMetadataDefResponse**](BusinessMetadataDefResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagDefs

> []TagDefResponse CreateTagDefs(ctx).TagDef(tagDef).Execute()

Bulk Create Tag Definitions



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
    tagDef := []openapiclient.TagDef{*openapiclient.NewTagDef()} // []TagDef | The tag definitions to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.CreateTagDefs(context.Background()).TagDef(tagDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.CreateTagDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagDefs`: []TagDefResponse
    fmt.Fprintf(os.Stdout, "Response from `TypesV1Api.CreateTagDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagDef** | [**[]TagDef**](TagDef.md) | The tag definitions to create | 

### Return type

[**[]TagDefResponse**](TagDefResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBusinessMetadataDef

> DeleteBusinessMetadataDef(ctx, bmName).Execute()

Delete Business Metadata Definition



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
    bmName := "bmName_example" // string | The name of the business metadata definition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.DeleteBusinessMetadataDef(context.Background(), bmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.DeleteBusinessMetadataDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bmName** | **string** | The name of the business metadata definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBusinessMetadataDefRequest struct via the builder pattern


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


## DeleteTagDef

> DeleteTagDef(ctx, tagName).Execute()

Delete Tag Definition



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
    tagName := "tagName_example" // string | The name of the tag definition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.DeleteTagDef(context.Background(), tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.DeleteTagDef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string** | The name of the tag definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagDefRequest struct via the builder pattern


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


## GetAllBusinessMetadataDefs

> []BusinessMetadataDefResponse GetAllBusinessMetadataDefs(ctx).Prefix(prefix).Execute()

Bulk Read Business Metadata Definitions



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
    prefix := "prefix_example" // string | The prefix of a business metadata definition name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.GetAllBusinessMetadataDefs(context.Background()).Prefix(prefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.GetAllBusinessMetadataDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllBusinessMetadataDefs`: []BusinessMetadataDefResponse
    fmt.Fprintf(os.Stdout, "Response from `TypesV1Api.GetAllBusinessMetadataDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBusinessMetadataDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | The prefix of a business metadata definition name | 

### Return type

[**[]BusinessMetadataDefResponse**](BusinessMetadataDefResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTagDefs

> []TagDefResponse GetAllTagDefs(ctx).Prefix(prefix).Execute()

Bulk Read Tag Definitions



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
    prefix := "prefix_example" // string | The prefix of a tag definition name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.GetAllTagDefs(context.Background()).Prefix(prefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.GetAllTagDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTagDefs`: []TagDefResponse
    fmt.Fprintf(os.Stdout, "Response from `TypesV1Api.GetAllTagDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTagDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | The prefix of a tag definition name | 

### Return type

[**[]TagDefResponse**](TagDefResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessMetadataDefByName

> BusinessMetadataDef GetBusinessMetadataDefByName(ctx, bmName).Execute()

Read Business Metadata Definition



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
    bmName := "bmName_example" // string | The name of the business metadata definition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.GetBusinessMetadataDefByName(context.Background(), bmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.GetBusinessMetadataDefByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessMetadataDefByName`: BusinessMetadataDef
    fmt.Fprintf(os.Stdout, "Response from `TypesV1Api.GetBusinessMetadataDefByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bmName** | **string** | The name of the business metadata definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessMetadataDefByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BusinessMetadataDef**](BusinessMetadataDef.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagDefByName

> TagDef GetTagDefByName(ctx, tagName).Execute()

Read Tag Definition



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
    tagName := "tagName_example" // string | The name of the tag definiton

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.GetTagDefByName(context.Background(), tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.GetTagDefByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagDefByName`: TagDef
    fmt.Fprintf(os.Stdout, "Response from `TypesV1Api.GetTagDefByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string** | The name of the tag definiton | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagDefByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagDef**](TagDef.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBusinessMetadataDefs

> []BusinessMetadataDefResponse UpdateBusinessMetadataDefs(ctx).BusinessMetadataDef(businessMetadataDef).Execute()

Bulk Update Business Metadata Definitions



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
    businessMetadataDef := []openapiclient.BusinessMetadataDef{*openapiclient.NewBusinessMetadataDef()} // []BusinessMetadataDef | The business metadata definitions to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.UpdateBusinessMetadataDefs(context.Background()).BusinessMetadataDef(businessMetadataDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.UpdateBusinessMetadataDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBusinessMetadataDefs`: []BusinessMetadataDefResponse
    fmt.Fprintf(os.Stdout, "Response from `TypesV1Api.UpdateBusinessMetadataDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBusinessMetadataDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessMetadataDef** | [**[]BusinessMetadataDef**](BusinessMetadataDef.md) | The business metadata definitions to update | 

### Return type

[**[]BusinessMetadataDefResponse**](BusinessMetadataDefResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagDefs

> []TagDefResponse UpdateTagDefs(ctx).TagDef(tagDef).Execute()

Bulk Update Tag Definitions



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
    tagDef := []openapiclient.TagDef{*openapiclient.NewTagDef()} // []TagDef | The tag definitions to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TypesV1Api.UpdateTagDefs(context.Background()).TagDef(tagDef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TypesV1Api.UpdateTagDefs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTagDefs`: []TagDefResponse
    fmt.Fprintf(os.Stdout, "Response from `TypesV1Api.UpdateTagDefs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagDefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagDef** | [**[]TagDef**](TagDef.md) | The tag definitions to update | 

### Return type

[**[]TagDefResponse**](TagDefResponse.md)

### Authorization

[external-access-token](../README.md#external-access-token), [resource-api-key](../README.md#resource-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


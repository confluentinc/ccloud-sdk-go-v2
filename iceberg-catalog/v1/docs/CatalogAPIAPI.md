# \CatalogAPIAPI

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNamespaces**](CatalogAPIAPI.md#ListNamespaces) | **Get** /v1/{prefix}/namespaces | List namespaces, optionally providing a parent namespace to list underneath
[**ListTables**](CatalogAPIAPI.md#ListTables) | **Get** /v1/{prefix}/namespaces/{namespace}/tables | List all table identifiers underneath a given namespace
[**LoadNamespaceMetadata**](CatalogAPIAPI.md#LoadNamespaceMetadata) | **Get** /v1/{prefix}/namespaces/{namespace} | Load the metadata properties for a namespace
[**LoadTable**](CatalogAPIAPI.md#LoadTable) | **Get** /v1/{prefix}/namespaces/{namespace}/tables/{table} | Load a table from the catalog
[**NamespaceExists**](CatalogAPIAPI.md#NamespaceExists) | **Head** /v1/{prefix}/namespaces/{namespace} | Check if a namespace exists
[**TableExists**](CatalogAPIAPI.md#TableExists) | **Head** /v1/{prefix}/namespaces/{namespace}/tables/{table} | Check if a table exists



## ListNamespaces

> ListNamespacesResponse ListNamespaces(ctx, prefix).Parent(parent).Execute()

List namespaces, optionally providing a parent namespace to list underneath



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/confluentinc/ccloud-sdk-go-v2"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	parent := "accounting%1Ftax" // string | An optional namespace, underneath which to list namespaces. If not provided or empty, all top-level namespaces should be listed. If parent is a multipart namespace, the parts must be separated by the unit separator (`0x1F`) byte. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.ListNamespaces(context.Background(), prefix).Parent(parent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.ListNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespaces`: ListNamespacesResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.ListNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parent** | **string** | An optional namespace, underneath which to list namespaces. If not provided or empty, all top-level namespaces should be listed. If parent is a multipart namespace, the parts must be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Return type

[**ListNamespacesResponse**](ListNamespacesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTables

> ListTablesResponse ListTables(ctx, prefix, namespace).Execute()

List all table identifiers underneath a given namespace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/confluentinc/ccloud-sdk-go-v2"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.ListTables(context.Background(), prefix, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.ListTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTables`: ListTablesResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.ListTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListTablesResponse**](ListTablesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadNamespaceMetadata

> GetNamespaceResponse LoadNamespaceMetadata(ctx, prefix, namespace).Execute()

Load the metadata properties for a namespace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/confluentinc/ccloud-sdk-go-v2"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.LoadNamespaceMetadata(context.Background(), prefix, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.LoadNamespaceMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadNamespaceMetadata`: GetNamespaceResponse
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.LoadNamespaceMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadNamespaceMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNamespaceResponse**](GetNamespaceResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadTable

> LoadTableResult LoadTable(ctx, prefix, namespace, table).XIcebergAccessDelegation(xIcebergAccessDelegation).Snapshots(snapshots).Execute()

Load a table from the catalog



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/confluentinc/ccloud-sdk-go-v2"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name
	xIcebergAccessDelegation := "vended-credentials,remote-signing" // string | Optional signal to the server that the client supports delegated access via a comma-separated list of access mechanisms.  The server may choose to supply access via any or none of the requested mechanisms.  Specific properties and handling for `vended-credentials` is documented in the `LoadTableResult` schema section of this spec document.  The protocol and specification for `remote-signing` is documented in the `s3-signer-open-api.yaml` OpenApi spec in the `aws` module.  (optional)
	snapshots := "snapshots_example" // string | The snapshots to return in the body of the metadata. Setting the value to `all` would return the full set of snapshots currently valid for the table. Setting the value to `refs` would load all snapshots referenced by branches or tags. Default if no param is provided is `all`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogAPIAPI.LoadTable(context.Background(), prefix, namespace, table).XIcebergAccessDelegation(xIcebergAccessDelegation).Snapshots(snapshots).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.LoadTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadTable`: LoadTableResult
	fmt.Fprintf(os.Stdout, "Response from `CatalogAPIAPI.LoadTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xIcebergAccessDelegation** | **string** | Optional signal to the server that the client supports delegated access via a comma-separated list of access mechanisms.  The server may choose to supply access via any or none of the requested mechanisms.  Specific properties and handling for &#x60;vended-credentials&#x60; is documented in the &#x60;LoadTableResult&#x60; schema section of this spec document.  The protocol and specification for &#x60;remote-signing&#x60; is documented in the &#x60;s3-signer-open-api.yaml&#x60; OpenApi spec in the &#x60;aws&#x60; module.  | 
 **snapshots** | **string** | The snapshots to return in the body of the metadata. Setting the value to &#x60;all&#x60; would return the full set of snapshots currently valid for the table. Setting the value to &#x60;refs&#x60; would load all snapshots referenced by branches or tags. Default if no param is provided is &#x60;all&#x60;. | 

### Return type

[**LoadTableResult**](LoadTableResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NamespaceExists

> NamespaceExists(ctx, prefix, namespace).Execute()

Check if a namespace exists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/confluentinc/ccloud-sdk-go-v2"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.NamespaceExists(context.Background(), prefix, namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.NamespaceExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNamespaceExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TableExists

> TableExists(ctx, prefix, namespace, table).Execute()

Check if a table exists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/confluentinc/ccloud-sdk-go-v2"
)

func main() {
	prefix := "prefix_example" // string | An optional prefix in the path
	namespace := "accounting" // string | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (`0x1F`) byte.
	table := "sales" // string | A table name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CatalogAPIAPI.TableExists(context.Background(), prefix, namespace, table).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogAPIAPI.TableExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | An optional prefix in the path | 
**namespace** | **string** | A namespace identifier as a single string. Multipart namespace parts should be separated by the unit separator (&#x60;0x1F&#x60;) byte. | 
**table** | **string** | A table name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTableExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


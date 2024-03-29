# Go API client for v1

REST API for the Data Catalog

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://confluent.slack.com/app_redirect?channel=data-governance-eng](https://confluent.slack.com/app_redirect?channel=data-governance-eng)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./v1"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://psrc-00000.region.provider.confluent.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EntityV1Api* | [**CreateBusinessMetadata**](docs/EntityV1Api.md#createbusinessmetadata) | **Post** /catalog/v1/entity/businessmetadata | Bulk Create Business Metadata
*EntityV1Api* | [**CreateTags**](docs/EntityV1Api.md#createtags) | **Post** /catalog/v1/entity/tags | Bulk Create Tags
*EntityV1Api* | [**DeleteBusinessMetadata**](docs/EntityV1Api.md#deletebusinessmetadata) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/businessmetadata/{bmName} | Delete a Business Metadata for an Entity
*EntityV1Api* | [**DeleteTag**](docs/EntityV1Api.md#deletetag) | **Delete** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags/{tagName} | Delete a Tag for an Entity
*EntityV1Api* | [**GetBusinessMetadata**](docs/EntityV1Api.md#getbusinessmetadata) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/businessmetadata | Read Business Metadata for an Entity
*EntityV1Api* | [**GetByUniqueAttributes**](docs/EntityV1Api.md#getbyuniqueattributes) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName} | Read an Entity
*EntityV1Api* | [**GetTags**](docs/EntityV1Api.md#gettags) | **Get** /catalog/v1/entity/type/{typeName}/name/{qualifiedName}/tags | Read Tags for an Entity
*EntityV1Api* | [**PartialEntityUpdate**](docs/EntityV1Api.md#partialentityupdate) | **Put** /catalog/v1/entity | Update an Entity Attribute
*EntityV1Api* | [**UpdateBusinessMetadata**](docs/EntityV1Api.md#updatebusinessmetadata) | **Put** /catalog/v1/entity/businessmetadata | Bulk Update Business Metadata
*EntityV1Api* | [**UpdateTags**](docs/EntityV1Api.md#updatetags) | **Put** /catalog/v1/entity/tags | Bulk Update Tags
*SearchV1Api* | [**SearchUsingAttribute**](docs/SearchV1Api.md#searchusingattribute) | **Get** /catalog/v1/search/attribute | Search by Attribute
*SearchV1Api* | [**SearchUsingBasic**](docs/SearchV1Api.md#searchusingbasic) | **Get** /catalog/v1/search/basic | Search by Fulltext Query
*TypesV1Api* | [**CreateBusinessMetadataDefs**](docs/TypesV1Api.md#createbusinessmetadatadefs) | **Post** /catalog/v1/types/businessmetadatadefs | Bulk Create Business Metadata Definitions
*TypesV1Api* | [**CreateTagDefs**](docs/TypesV1Api.md#createtagdefs) | **Post** /catalog/v1/types/tagdefs | Bulk Create Tag Definitions
*TypesV1Api* | [**DeleteBusinessMetadataDef**](docs/TypesV1Api.md#deletebusinessmetadatadef) | **Delete** /catalog/v1/types/businessmetadatadefs/{bmName} | Delete Business Metadata Definition
*TypesV1Api* | [**DeleteTagDef**](docs/TypesV1Api.md#deletetagdef) | **Delete** /catalog/v1/types/tagdefs/{tagName} | Delete Tag Definition
*TypesV1Api* | [**GetAllBusinessMetadataDefs**](docs/TypesV1Api.md#getallbusinessmetadatadefs) | **Get** /catalog/v1/types/businessmetadatadefs | Bulk Read Business Metadata Definitions
*TypesV1Api* | [**GetAllTagDefs**](docs/TypesV1Api.md#getalltagdefs) | **Get** /catalog/v1/types/tagdefs | Bulk Read Tag Definitions
*TypesV1Api* | [**GetBusinessMetadataDefByName**](docs/TypesV1Api.md#getbusinessmetadatadefbyname) | **Get** /catalog/v1/types/businessmetadatadefs/{bmName} | Read Business Metadata Definition
*TypesV1Api* | [**GetTagDefByName**](docs/TypesV1Api.md#gettagdefbyname) | **Get** /catalog/v1/types/tagdefs/{tagName} | Read Tag Definition
*TypesV1Api* | [**UpdateBusinessMetadataDefs**](docs/TypesV1Api.md#updatebusinessmetadatadefs) | **Put** /catalog/v1/types/businessmetadatadefs | Bulk Update Business Metadata Definitions
*TypesV1Api* | [**UpdateTagDefs**](docs/TypesV1Api.md#updatetagdefs) | **Put** /catalog/v1/types/tagdefs | Bulk Update Tag Definitions


## Documentation For Models

 - [AttributeDef](docs/AttributeDef.md)
 - [BusinessMetadata](docs/BusinessMetadata.md)
 - [BusinessMetadataDef](docs/BusinessMetadataDef.md)
 - [BusinessMetadataDefResponse](docs/BusinessMetadataDefResponse.md)
 - [BusinessMetadataResponse](docs/BusinessMetadataResponse.md)
 - [Classification](docs/Classification.md)
 - [ClassificationHeader](docs/ClassificationHeader.md)
 - [ConstraintDef](docs/ConstraintDef.md)
 - [Entity](docs/Entity.md)
 - [EntityHeader](docs/EntityHeader.md)
 - [EntityPartialUpdate](docs/EntityPartialUpdate.md)
 - [EntityPartialUpdateResponse](docs/EntityPartialUpdateResponse.md)
 - [EntityWithExtInfo](docs/EntityWithExtInfo.md)
 - [ErrorMessage](docs/ErrorMessage.md)
 - [PartialUpdateParams](docs/PartialUpdateParams.md)
 - [SearchParams](docs/SearchParams.md)
 - [SearchResult](docs/SearchResult.md)
 - [Tag](docs/Tag.md)
 - [TagDef](docs/TagDef.md)
 - [TagDefResponse](docs/TagDefResponse.md)
 - [TagResponse](docs/TagResponse.md)
 - [TermAssignmentHeader](docs/TermAssignmentHeader.md)
 - [TimeBoundary](docs/TimeBoundary.md)


## Documentation For Authorization



### external-access-token


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### resource-api-key

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

data-governance@confluent.io


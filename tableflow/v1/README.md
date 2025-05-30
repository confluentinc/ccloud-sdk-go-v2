# Go API client for v1

Tableflow Management API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://confluent.slack.com/app_redirect?channel=cts-eng](https://confluent.slack.com/app_redirect?channel=cts-eng)

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

All URIs are relative to *https://api.confluent.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CatalogIntegrationsTableflowV1Api* | [**CreateTableflowV1CatalogIntegration**](docs/CatalogIntegrationsTableflowV1Api.md#createtableflowv1catalogintegration) | **Post** /tableflow/v1/catalog-integrations | Create a Catalog Integration
*CatalogIntegrationsTableflowV1Api* | [**DeleteTableflowV1CatalogIntegration**](docs/CatalogIntegrationsTableflowV1Api.md#deletetableflowv1catalogintegration) | **Delete** /tableflow/v1/catalog-integrations/{id} | Delete a Catalog Integration
*CatalogIntegrationsTableflowV1Api* | [**GetTableflowV1CatalogIntegration**](docs/CatalogIntegrationsTableflowV1Api.md#gettableflowv1catalogintegration) | **Get** /tableflow/v1/catalog-integrations/{id} | Read a Catalog Integration
*CatalogIntegrationsTableflowV1Api* | [**ListTableflowV1CatalogIntegrations**](docs/CatalogIntegrationsTableflowV1Api.md#listtableflowv1catalogintegrations) | **Get** /tableflow/v1/catalog-integrations | List of Catalog Integrations
*CatalogIntegrationsTableflowV1Api* | [**UpdateTableflowV1CatalogIntegration**](docs/CatalogIntegrationsTableflowV1Api.md#updatetableflowv1catalogintegration) | **Patch** /tableflow/v1/catalog-integrations/{id} | Update a Catalog Integration
*RegionsTableflowV1Api* | [**ListTableflowV1Regions**](docs/RegionsTableflowV1Api.md#listtableflowv1regions) | **Get** /tableflow/v1/regions | List of Regions
*TableflowTopicsTableflowV1Api* | [**CreateTableflowV1TableflowTopic**](docs/TableflowTopicsTableflowV1Api.md#createtableflowv1tableflowtopic) | **Post** /tableflow/v1/tableflow-topics | Create a Tableflow Topic
*TableflowTopicsTableflowV1Api* | [**DeleteTableflowV1TableflowTopic**](docs/TableflowTopicsTableflowV1Api.md#deletetableflowv1tableflowtopic) | **Delete** /tableflow/v1/tableflow-topics/{display_name} | Delete a Tableflow Topic
*TableflowTopicsTableflowV1Api* | [**GetTableflowV1TableflowTopic**](docs/TableflowTopicsTableflowV1Api.md#gettableflowv1tableflowtopic) | **Get** /tableflow/v1/tableflow-topics/{display_name} | Read a Tableflow Topic
*TableflowTopicsTableflowV1Api* | [**ListTableflowV1TableflowTopics**](docs/TableflowTopicsTableflowV1Api.md#listtableflowv1tableflowtopics) | **Get** /tableflow/v1/tableflow-topics | List of Tableflow Topics
*TableflowTopicsTableflowV1Api* | [**UpdateTableflowV1TableflowTopic**](docs/TableflowTopicsTableflowV1Api.md#updatetableflowv1tableflowtopic) | **Patch** /tableflow/v1/tableflow-topics/{display_name} | Update a Tableflow Topic


## Documentation For Models

 - [EnvScopedObjectReference](docs/EnvScopedObjectReference.md)
 - [Error](docs/Error.md)
 - [ErrorSource](docs/ErrorSource.md)
 - [Failure](docs/Failure.md)
 - [GlobalObjectReference](docs/GlobalObjectReference.md)
 - [ListMeta](docs/ListMeta.md)
 - [ObjectMeta](docs/ObjectMeta.md)
 - [ObjectReference](docs/ObjectReference.md)
 - [TableflowV1ByobAwsSpec](docs/TableflowV1ByobAwsSpec.md)
 - [TableflowV1CatalogIntegration](docs/TableflowV1CatalogIntegration.md)
 - [TableflowV1CatalogIntegrationAwsGlueSpec](docs/TableflowV1CatalogIntegrationAwsGlueSpec.md)
 - [TableflowV1CatalogIntegrationAwsGlueUpdateSpec](docs/TableflowV1CatalogIntegrationAwsGlueUpdateSpec.md)
 - [TableflowV1CatalogIntegrationList](docs/TableflowV1CatalogIntegrationList.md)
 - [TableflowV1CatalogIntegrationSnowflakeSpec](docs/TableflowV1CatalogIntegrationSnowflakeSpec.md)
 - [TableflowV1CatalogIntegrationSnowflakeUpdateSpec](docs/TableflowV1CatalogIntegrationSnowflakeUpdateSpec.md)
 - [TableflowV1CatalogIntegrationSpec](docs/TableflowV1CatalogIntegrationSpec.md)
 - [TableflowV1CatalogIntegrationSpecConfigOneOf](docs/TableflowV1CatalogIntegrationSpecConfigOneOf.md)
 - [TableflowV1CatalogIntegrationStatus](docs/TableflowV1CatalogIntegrationStatus.md)
 - [TableflowV1CatalogIntegrationUpdateRequest](docs/TableflowV1CatalogIntegrationUpdateRequest.md)
 - [TableflowV1CatalogIntegrationUpdateSpec](docs/TableflowV1CatalogIntegrationUpdateSpec.md)
 - [TableflowV1CatalogIntegrationUpdateSpecConfigOneOf](docs/TableflowV1CatalogIntegrationUpdateSpecConfigOneOf.md)
 - [TableflowV1ManagedStorageSpec](docs/TableflowV1ManagedStorageSpec.md)
 - [TableflowV1Region](docs/TableflowV1Region.md)
 - [TableflowV1RegionList](docs/TableflowV1RegionList.md)
 - [TableflowV1TableFlowTopicConfigsSpec](docs/TableflowV1TableFlowTopicConfigsSpec.md)
 - [TableflowV1TableflowTopic](docs/TableflowV1TableflowTopic.md)
 - [TableflowV1TableflowTopicList](docs/TableflowV1TableflowTopicList.md)
 - [TableflowV1TableflowTopicSpec](docs/TableflowV1TableflowTopicSpec.md)
 - [TableflowV1TableflowTopicSpecStorageOneOf](docs/TableflowV1TableflowTopicSpecStorageOneOf.md)
 - [TableflowV1TableflowTopicSpecUpdate](docs/TableflowV1TableflowTopicSpecUpdate.md)
 - [TableflowV1TableflowTopicStatus](docs/TableflowV1TableflowTopicStatus.md)
 - [TableflowV1TableflowTopicUpdate](docs/TableflowV1TableflowTopicUpdate.md)


## Documentation For Authorization



### confluent-sts-access-token


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

cts-team@confluent.io


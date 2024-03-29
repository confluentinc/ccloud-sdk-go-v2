# Go API client for v1

API to manage various Kafka Quotas.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1-alpha1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://confluent.slack.com/app_redirect?channel=kafka-cloud-fundamentals-eng](https://confluent.slack.com/app_redirect?channel=kafka-cloud-fundamentals-eng)

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
*ClientQuotasKafkaQuotasV1Api* | [**CreateKafkaQuotasV1ClientQuota**](docs/ClientQuotasKafkaQuotasV1Api.md#createkafkaquotasv1clientquota) | **Post** /kafka-quotas/v1/client-quotas | Create a Client Quota
*ClientQuotasKafkaQuotasV1Api* | [**DeleteKafkaQuotasV1ClientQuota**](docs/ClientQuotasKafkaQuotasV1Api.md#deletekafkaquotasv1clientquota) | **Delete** /kafka-quotas/v1/client-quotas/{id} | Delete a Client Quota
*ClientQuotasKafkaQuotasV1Api* | [**GetKafkaQuotasV1ClientQuota**](docs/ClientQuotasKafkaQuotasV1Api.md#getkafkaquotasv1clientquota) | **Get** /kafka-quotas/v1/client-quotas/{id} | Read a Client Quota
*ClientQuotasKafkaQuotasV1Api* | [**ListKafkaQuotasV1ClientQuotas**](docs/ClientQuotasKafkaQuotasV1Api.md#listkafkaquotasv1clientquotas) | **Get** /kafka-quotas/v1/client-quotas | List of Client Quotas
*ClientQuotasKafkaQuotasV1Api* | [**UpdateKafkaQuotasV1ClientQuota**](docs/ClientQuotasKafkaQuotasV1Api.md#updatekafkaquotasv1clientquota) | **Patch** /kafka-quotas/v1/client-quotas/{id} | Update a Client Quota


## Documentation For Models

 - [EnvScopedObjectReference](docs/EnvScopedObjectReference.md)
 - [Error](docs/Error.md)
 - [ErrorSource](docs/ErrorSource.md)
 - [Failure](docs/Failure.md)
 - [GlobalObjectReference](docs/GlobalObjectReference.md)
 - [KafkaQuotasV1ClientQuota](docs/KafkaQuotasV1ClientQuota.md)
 - [KafkaQuotasV1ClientQuotaList](docs/KafkaQuotasV1ClientQuotaList.md)
 - [KafkaQuotasV1ClientQuotaSpec](docs/KafkaQuotasV1ClientQuotaSpec.md)
 - [KafkaQuotasV1ClientQuotaSpecUpdate](docs/KafkaQuotasV1ClientQuotaSpecUpdate.md)
 - [KafkaQuotasV1ClientQuotaUpdate](docs/KafkaQuotasV1ClientQuotaUpdate.md)
 - [KafkaQuotasV1Throughput](docs/KafkaQuotasV1Throughput.md)
 - [ListMeta](docs/ListMeta.md)
 - [ObjectMeta](docs/ObjectMeta.md)
 - [ObjectReference](docs/ObjectReference.md)


## Documentation For Authorization



### api-key

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

kafka-cloud-fundament-aaaacmo35d4fp7t7p45tvuw6uq@confluent.slack.com


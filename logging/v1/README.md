# Go API client for v1

# Introduction

The Logs API is served by `cc-logging-service` and returns customer-facing operational log
documents for one Confluent Cloud resource, filtered by time window and by the search fields
declared for that resource type. Its consumers today are the
Confluent Cloud Console log viewer and the Confluent CLI (`pkg/ccloudv2/logging.go`). This spec
documents the API exactly as the service serves it on `master` as of 2026-10-15; it is the
source for the generated Go client.

# Authentication

Requests carry a Confluent Cloud data-plane bearer token obtained from a user session
(`Authorization: Bearer <token>`). Cloud API keys are not accepted. The organization in the
request's `crn` must match the organization of the token's principal, otherwise the request is
rejected with `422`.

Access is gated per organization by the `logging.external.api.enable` LaunchDarkly flag. An
organization without the flag receives `403` with the detail `Feature is not enabled for this
organization`.

# Resource types

The service resolves the trailing element of `crn` against its resource catalog, which is
deployment configuration (`resourceCatalog` in `.halyard/cc-logging-service.yaml`). The catalog
decides which OpenSearch indices are queried, which `search` keys are accepted, which fields a
document carries in the response, and which RBAC operation is checked. As of 2026-10-15 the
production catalog serves these resource types:

| CRN shape (trailing element) | Accepted `search` keys | Allowed values | Returned fields | RBAC operation |
|------------------------------|------------------------|----------------|-----------------|----------------|
| `organization/environment/cloud-cluster/connector` | `search_text`, `level`, `task_id`, `id` | `level`: `INFO`, `WARN`, `ERROR`, `FATAL` | `timestamp`, `level`, `task_id`, `message`, `exception.stacktrace`, `id`, `name` | `ReadLogs` |
| `organization/environment/flink-region/statement` | `search_text`, `level`, `source`, `id` | `level`: `INFO`, `WARN`, `ERROR` | `id`, `timestamp`, `level`, `message`, `crn`, `data`, `exception`, `function_name`, `function_crn`, `source` | `Describe` |
| `organization/environment/flink-region/cloud-cluster/materialized-table` | `search_text`, `level`, `source`, `id` | `level`: `INFO`, `WARN`, `ERROR`; `source`: `MaterializedTable`, `Function` | `id`, `timestamp`, `level`, `message`, `crn`, `data`, `source`, `materialized_table_name`, `materialized_table_version` | `Describe` |
| `organization/environment/cloud-cluster/kafka` | `search_text`, `event_type`, `level`, `id` | `level`: `INFO`, `WARN`, `ERROR`; `event_type`: `SERVICE_MAINTENANCE`, `HEALTH_ALERT`, `CONFLUENT_STATUS_UPDATED`, `SERVICE_ISSUE_DETECTED` | `id`, `timestamp`, `level`, `message`, `crn`, `source`, `event_type`, `event_id`, `status`, `cloud_cluster_id` | `ReadLogs` |

A `crn` that cannot be parsed, or that ends in one of these resource types but does not have every
element of its shape, is rejected with `422 Invalid CRN`, and an element that does not match its
position in the shape with `422 Invalid CRN element {element}`; see Choosing a log source for the
order of the checks. For connectors, the service first
resolves the connector's logical cluster ID through the metadata service; a connector it cannot
resolve is reported as `403`, never `404`, so that the response does not reveal whether the
resource exists.

The lookback for every resource type above is 3 days. An `end_time` older than that returns an
empty `200` without querying OpenSearch; a `start_time` older than that is silently clamped.

Audit logs (`audit_logs/v1`) are served by the same host under a separate path and are not part
of this spec.

# Discovery

`GET /logs/v1/sources` and `GET /logs/v1/sources/{name}` let a client learn which log source
types this deployment's catalog serves, and what each one supports (crn shape, accepted
`search` keys and their allowed values, returned fields, and lookback).

These endpoints make no OpenSearch calls themselves and are
not subject to the per-resource rate limit described below. They're instead subject to the
per-IP limit plus their own (user, organization) limit.

# Choosing a log source

`source` is optional. The service infers the log source from the `crn`'s trailing element, so a
client only sets it to pick one of several log sources for the same resource type.

The service checks the `crn` and `source` in this order and returns the first failure as a `422`:

1. The `crn` cannot be parsed: `Invalid CRN`.
2. There is no `source`, and the `crn`'s trailing element is a resource type with more than one
   log source: `source is required for resource type {type}; one of [{sources}]`.
3. The log source cannot be resolved. With no `source` and a trailing element that is not a
   catalog resource type: `Invalid resource type {type}`. With a `source` that does not exist:
   `Unknown source {name}`. With a `source` that does not serve the `crn`'s trailing resource
   type: `source {name} does not support resource type {type}`.
4. The `crn` does not have every element of the resolved resource type's shape: `Invalid CRN`.
   An element that does not match its position: `Invalid CRN element {element}`.

Every query targets one resource, so the `crn` must be complete. A `crn` truncated to an
organization or an environment ends in `organization` or `environment`, which is not a resource
type, so it fails at step 3: `Invalid resource type environment` without a `source`, and
`source cluster-health does not support resource type environment` with one. It never reaches
step 4. A `crn` that ends in a real resource type but is missing elements, for example one that
stops after `connector=my-connector` with no environment before it, fails at step 4 with
`Invalid CRN`.

# Pagination

`POST /logs/v1/search` paginates with an opaque `page_token`. When a page is full, the response
carries the URL of the next page in `metadata.next`; `metadata` is absent when the page was not
full or was empty. The default page size is 100 documents; the maximum is 1000.

# Rate limiting

The rate-limit sidecar enforces, per (resource CRN, user, organization, client type), 1 request
per second for UI clients (User-Agent containing `mozilla`) and 5 requests per minute for every
other client, plus 5 requests per second per client IP before authentication. The per-resource
limit applies only to `POST /logs/v1/search`, which is the only endpoint with a resource CRN to
key it on; the per-IP limit applies to every endpoint in this spec, including the sources
endpoints. The sources endpoints additionally enforce their own limit of 60 requests per minute
per (user, organization) pair, independent of the per-IP check — discovery has no resource to
key a per-resource limit on, and no UI-vs-other-client distinction the way search does. Every
response that passed the per-IP limiter carries `X-RateLimit-Limit`, `X-RateLimit-Remaining`,
and `X-RateLimit-Reset` headers, including non-429 responses; a rate-limited request receives
`429`. `Retry-After` is not sent. These numbers come from `ratelimit-config.yaml` and
`discovery-ratelimit-config.yaml` in the service's Halyard spec as of 2026-09-15 and are
deployment configuration, not part of the API contract.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://confluent.slack.com/app_redirect?channel=logging-platform](https://confluent.slack.com/app_redirect?channel=logging-platform)

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

All URIs are relative to *https://api.logging.confluent.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*LogsLoggingV1Api* | [**DescribeLoggingV1Source**](docs/LogsLoggingV1Api.md#describeloggingv1source) | **Get** /logs/v1/sources/{name} | Describe one log source
*LogsLoggingV1Api* | [**ListLoggingV1Sources**](docs/LogsLoggingV1Api.md#listloggingv1sources) | **Get** /logs/v1/sources | List available log sources
*LogsLoggingV1Api* | [**SearchLoggingV1Logs**](docs/LogsLoggingV1Api.md#searchloggingv1logs) | **Post** /logs/v1/search | Search logs


## Documentation For Models

 - [ListMeta](docs/ListMeta.md)
 - [LoggingV1Aggregation](docs/LoggingV1Aggregation.md)
 - [LoggingV1AggregationResponse](docs/LoggingV1AggregationResponse.md)
 - [LoggingV1AggregationResponseTimestampHistogram](docs/LoggingV1AggregationResponseTimestampHistogram.md)
 - [LoggingV1AggregationResponseTimestampHistogramAggregatedTerms](docs/LoggingV1AggregationResponseTimestampHistogramAggregatedTerms.md)
 - [LoggingV1AggregationResponseTimestampHistogramAggregatedTermsBuckets](docs/LoggingV1AggregationResponseTimestampHistogramAggregatedTermsBuckets.md)
 - [LoggingV1AggregationResponseTimestampHistogramBuckets](docs/LoggingV1AggregationResponseTimestampHistogramBuckets.md)
 - [LoggingV1ErrorResponse](docs/LoggingV1ErrorResponse.md)
 - [LoggingV1ErrorResponseErrors](docs/LoggingV1ErrorResponseErrors.md)
 - [LoggingV1ReturnedFieldDescriptor](docs/LoggingV1ReturnedFieldDescriptor.md)
 - [LoggingV1SearchFieldDescriptor](docs/LoggingV1SearchFieldDescriptor.md)
 - [LoggingV1SearchRequest](docs/LoggingV1SearchRequest.md)
 - [LoggingV1SearchResponse](docs/LoggingV1SearchResponse.md)
 - [LoggingV1SearchResponseMetadata](docs/LoggingV1SearchResponseMetadata.md)
 - [LoggingV1Source](docs/LoggingV1Source.md)
 - [LoggingV1SourceList](docs/LoggingV1SourceList.md)


## Documentation For Authorization



### confluent_auth


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://confluent.cloud/oauth
- **Scopes**: 
 - **Complex**: RBAC permissions required for this endpoint depend on the resource type of the queried CRN.
 - **None**: No authorization check is performed on this endpoint.

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


### data-plane-bearer-token

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
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

logging-platform@confluent.io


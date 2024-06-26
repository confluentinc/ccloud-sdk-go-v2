# Go API client for v1

API for interacting with AI models from within Confluent Cloud.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://confluent.slack.com/app_redirect?channel=api-eng](https://confluent.slack.com/app_redirect?channel=api-eng)

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
*AvailabilitiesAiV1Api* | [**GetAiV1Availability**](docs/AvailabilitiesAiV1Api.md#getaiv1availability) | **Get** /ai/v1/availability | Read the organization&#39;s ai-assistant setting.
*ChatCompletionsAiV1Api* | [**GetAiV1ChatCompletion**](docs/ChatCompletionsAiV1Api.md#getaiv1chatcompletion) | **Options** /ai/v1/chat-completions | Read a Chat Completion
*ChatCompletionsAiV1Api* | [**QueryAiV1ChatCompletion**](docs/ChatCompletionsAiV1Api.md#queryaiv1chatcompletion) | **Post** /ai/v1/chat-completions | Query a Chat Completion
*DocCompletionsDocsAiV1Api* | [**QueryDocsAiV1DocCompletion**](docs/DocCompletionsDocsAiV1Api.md#querydocsaiv1doccompletion) | **Post** /docs-ai/v1/doc-completions | Query a Doc Completion
*FeedbacksAiV1Api* | [**CreateAiV1ChatCompletionFeedback**](docs/FeedbacksAiV1Api.md#createaiv1chatcompletionfeedback) | **Post** /ai/v1/chat-completions/{chat_completion_id}/feedback | Create a Feedback
*FeedbacksAiV1Api* | [**CreateAiV1DocCompletionFeedback**](docs/FeedbacksAiV1Api.md#createaiv1doccompletionfeedback) | **Post** /docs-ai/v1/doc-completions/{doc_completion_id}/feedback | Create a Feedback
*OrgPreferencesAiV1Api* | [**GetAiV1OrgPreference**](docs/OrgPreferencesAiV1Api.md#getaiv1orgpreference) | **Get** /ai/v1/org-preferences | Read the organization&#39;s ai-assistant setting in org-preferences.
*OrgPreferencesAiV1Api* | [**UpdateAiV1OrgPreference**](docs/OrgPreferencesAiV1Api.md#updateaiv1orgpreference) | **Patch** /ai/v1/org-preferences | Set the organization&#39;s ai-assistant setting in org-preferences.


## Documentation For Models

 - [AiV1Availability](docs/AiV1Availability.md)
 - [AiV1ChatCompletionsHistory](docs/AiV1ChatCompletionsHistory.md)
 - [AiV1ChatCompletionsReply](docs/AiV1ChatCompletionsReply.md)
 - [AiV1ChatCompletionsRequest](docs/AiV1ChatCompletionsRequest.md)
 - [AiV1Feedback](docs/AiV1Feedback.md)
 - [AiV1OrgPreferences](docs/AiV1OrgPreferences.md)
 - [Error](docs/Error.md)
 - [ErrorSource](docs/ErrorSource.md)
 - [Failure](docs/Failure.md)
 - [ListMeta](docs/ListMeta.md)
 - [ObjectMeta](docs/ObjectMeta.md)
 - [ObjectReference](docs/ObjectReference.md)


## Documentation For Authorization



### cloud-api-key

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

## Author of this repo

api-team@confluent.io


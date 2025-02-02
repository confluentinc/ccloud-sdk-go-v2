# Go API client for v2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1-alpha0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://confluent.slack.com/app_redirect?channel=paas-eng](https://confluent.slack.com/app_redirect?channel=paas-eng)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./v2"
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
*InvitationsIamV2Api* | [**CreateIamV2Invitation**](docs/InvitationsIamV2Api.md#createiamv2invitation) | **Post** /iam/v2/invitations | Create an Invitation
*InvitationsIamV2Api* | [**DeleteIamV2Invitation**](docs/InvitationsIamV2Api.md#deleteiamv2invitation) | **Delete** /iam/v2/invitations/{id} | Delete an Invitation
*InvitationsIamV2Api* | [**GetIamV2Invitation**](docs/InvitationsIamV2Api.md#getiamv2invitation) | **Get** /iam/v2/invitations/{id} | Read an Invitation
*InvitationsIamV2Api* | [**ListIamV2Invitations**](docs/InvitationsIamV2Api.md#listiamv2invitations) | **Get** /iam/v2/invitations | List of Invitations
*ServiceAccountsIamV2Api* | [**CreateIamV2ServiceAccount**](docs/ServiceAccountsIamV2Api.md#createiamv2serviceaccount) | **Post** /iam/v2/service-accounts | Create a Service Account
*ServiceAccountsIamV2Api* | [**DeleteIamV2ServiceAccount**](docs/ServiceAccountsIamV2Api.md#deleteiamv2serviceaccount) | **Delete** /iam/v2/service-accounts/{id} | Delete a Service Account
*ServiceAccountsIamV2Api* | [**GetIamV2ServiceAccount**](docs/ServiceAccountsIamV2Api.md#getiamv2serviceaccount) | **Get** /iam/v2/service-accounts/{id} | Read a Service Account
*ServiceAccountsIamV2Api* | [**ListIamV2ServiceAccounts**](docs/ServiceAccountsIamV2Api.md#listiamv2serviceaccounts) | **Get** /iam/v2/service-accounts | List of Service Accounts
*ServiceAccountsIamV2Api* | [**UpdateIamV2ServiceAccount**](docs/ServiceAccountsIamV2Api.md#updateiamv2serviceaccount) | **Patch** /iam/v2/service-accounts/{id} | Update a Service Account
*UsersIamV2Api* | [**DeleteIamV2User**](docs/UsersIamV2Api.md#deleteiamv2user) | **Delete** /iam/v2/users/{id} | Delete a User
*UsersIamV2Api* | [**GetIamV2User**](docs/UsersIamV2Api.md#getiamv2user) | **Get** /iam/v2/users/{id} | Read a User
*UsersIamV2Api* | [**ListIamV2Users**](docs/UsersIamV2Api.md#listiamv2users) | **Get** /iam/v2/users | List of Users
*UsersIamV2Api* | [**UpdateIamV2User**](docs/UsersIamV2Api.md#updateiamv2user) | **Patch** /iam/v2/users/{id} | Update a User


## Documentation For Models

 - [Error](docs/Error.md)
 - [ErrorSource](docs/ErrorSource.md)
 - [Failure](docs/Failure.md)
 - [GlobalObjectReference](docs/GlobalObjectReference.md)
 - [IamV2Invitation](docs/IamV2Invitation.md)
 - [IamV2InvitationList](docs/IamV2InvitationList.md)
 - [IamV2ServiceAccount](docs/IamV2ServiceAccount.md)
 - [IamV2ServiceAccountList](docs/IamV2ServiceAccountList.md)
 - [IamV2User](docs/IamV2User.md)
 - [IamV2UserList](docs/IamV2UserList.md)
 - [IamV2UserUpdate](docs/IamV2UserUpdate.md)
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

paas-team@confluent.io


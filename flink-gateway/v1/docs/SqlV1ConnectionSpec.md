# SqlV1ConnectionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to **string** | The type of this connection. | [optional] 
**Endpoint** | Pointer to **string** | The endpoint that is used to run model inferencing. | [optional] 
**AuthData** | Pointer to [**SqlV1ConnectionSpecAuthDataOneOf**](SqlV1ConnectionSpecAuthDataOneOf.md) | The vendor specific authentication token details  The contents are stored as opaque bytes given in plaintext by an EnvAdmin. In future, we would support more secure methods for distributing authentication tokens.  | [optional] 

## Methods

### NewSqlV1ConnectionSpec

`func NewSqlV1ConnectionSpec() *SqlV1ConnectionSpec`

NewSqlV1ConnectionSpec instantiates a new SqlV1ConnectionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ConnectionSpecWithDefaults

`func NewSqlV1ConnectionSpecWithDefaults() *SqlV1ConnectionSpec`

NewSqlV1ConnectionSpecWithDefaults instantiates a new SqlV1ConnectionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *SqlV1ConnectionSpec) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SqlV1ConnectionSpec) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SqlV1ConnectionSpec) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *SqlV1ConnectionSpec) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetEndpoint

`func (o *SqlV1ConnectionSpec) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SqlV1ConnectionSpec) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SqlV1ConnectionSpec) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *SqlV1ConnectionSpec) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAuthData

`func (o *SqlV1ConnectionSpec) GetAuthData() SqlV1ConnectionSpecAuthDataOneOf`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *SqlV1ConnectionSpec) GetAuthDataOk() (*SqlV1ConnectionSpecAuthDataOneOf, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *SqlV1ConnectionSpec) SetAuthData(v SqlV1ConnectionSpecAuthDataOneOf)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *SqlV1ConnectionSpec) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



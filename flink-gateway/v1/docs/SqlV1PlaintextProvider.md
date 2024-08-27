# SqlV1PlaintextProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Plaintext Provider Kind Type  | [optional] 
**Data** | Pointer to **string** | Base64 encoded opaque piece of sensitive information.  Scoped to an endpoint of a &#x60;Connection&#x60; resource.  | [optional] 

## Methods

### NewSqlV1PlaintextProvider

`func NewSqlV1PlaintextProvider() *SqlV1PlaintextProvider`

NewSqlV1PlaintextProvider instantiates a new SqlV1PlaintextProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1PlaintextProviderWithDefaults

`func NewSqlV1PlaintextProviderWithDefaults() *SqlV1PlaintextProvider`

NewSqlV1PlaintextProviderWithDefaults instantiates a new SqlV1PlaintextProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SqlV1PlaintextProvider) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1PlaintextProvider) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1PlaintextProvider) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SqlV1PlaintextProvider) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetData

`func (o *SqlV1PlaintextProvider) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SqlV1PlaintextProvider) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SqlV1PlaintextProvider) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SqlV1PlaintextProvider) HasData() bool`

HasData returns a boolean if a field has been set.


### AsSqlV1ConnectionSpecAuthDataOneOf

`func (s *SqlV1PlaintextProvider) AsSqlV1ConnectionSpecAuthDataOneOf() SqlV1ConnectionSpecAuthDataOneOf`

Convenience method to wrap this instance of SqlV1PlaintextProvider in SqlV1ConnectionSpecAuthDataOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



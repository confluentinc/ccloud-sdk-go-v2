# SqlV1beta1StatementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | Pointer to **string** | The raw SQL text statement. | [optional] 
**Properties** | Pointer to **map[string]string** | A map (key-value pairs) of statement properties. | [optional] 
**ComputePoolId** | Pointer to **string** | The id associated with the compute pool in context. | [optional] 
**Principal** | Pointer to **string** | The id of a principal this statement runs as. | [optional] 
**Stopped** | Pointer to **bool** | Indicates whether the statement should be stopped. | [optional] 

## Methods

### NewSqlV1beta1StatementSpec

`func NewSqlV1beta1StatementSpec() *SqlV1beta1StatementSpec`

NewSqlV1beta1StatementSpec instantiates a new SqlV1beta1StatementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1beta1StatementSpecWithDefaults

`func NewSqlV1beta1StatementSpecWithDefaults() *SqlV1beta1StatementSpec`

NewSqlV1beta1StatementSpecWithDefaults instantiates a new SqlV1beta1StatementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *SqlV1beta1StatementSpec) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *SqlV1beta1StatementSpec) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *SqlV1beta1StatementSpec) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *SqlV1beta1StatementSpec) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetProperties

`func (o *SqlV1beta1StatementSpec) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SqlV1beta1StatementSpec) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SqlV1beta1StatementSpec) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SqlV1beta1StatementSpec) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetComputePoolId

`func (o *SqlV1beta1StatementSpec) GetComputePoolId() string`

GetComputePoolId returns the ComputePoolId field if non-nil, zero value otherwise.

### GetComputePoolIdOk

`func (o *SqlV1beta1StatementSpec) GetComputePoolIdOk() (*string, bool)`

GetComputePoolIdOk returns a tuple with the ComputePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePoolId

`func (o *SqlV1beta1StatementSpec) SetComputePoolId(v string)`

SetComputePoolId sets ComputePoolId field to given value.

### HasComputePoolId

`func (o *SqlV1beta1StatementSpec) HasComputePoolId() bool`

HasComputePoolId returns a boolean if a field has been set.

### GetPrincipal

`func (o *SqlV1beta1StatementSpec) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SqlV1beta1StatementSpec) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SqlV1beta1StatementSpec) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *SqlV1beta1StatementSpec) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetStopped

`func (o *SqlV1beta1StatementSpec) GetStopped() bool`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *SqlV1beta1StatementSpec) GetStoppedOk() (*bool, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *SqlV1beta1StatementSpec) SetStopped(v bool)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *SqlV1beta1StatementSpec) HasStopped() bool`

HasStopped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



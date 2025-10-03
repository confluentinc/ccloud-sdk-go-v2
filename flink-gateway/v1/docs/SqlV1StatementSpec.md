# SqlV1StatementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | Pointer to **string** | The raw SQL text statement. | [optional] 
**Properties** | Pointer to **map[string]string** | A map (key-value pairs) of statement properties. | [optional] 
**ComputePoolId** | Pointer to **string** | The id associated with the compute pool in context.  If not specified, the statement will use the default compute pool. The default pool is automatically determined by the system. | [optional]
**Principal** | Pointer to **string** | The id of a principal this statement runs as. | [optional] 
**AuthorizedPrincipals** | Pointer to **[]string** | The list of ids of the principals granting permissions to run this statement. | [optional] 
**Stopped** | Pointer to **bool** | Indicates whether the statement should be stopped. | [optional] 

## Methods

### NewSqlV1StatementSpec

`func NewSqlV1StatementSpec() *SqlV1StatementSpec`

NewSqlV1StatementSpec instantiates a new SqlV1StatementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1StatementSpecWithDefaults

`func NewSqlV1StatementSpecWithDefaults() *SqlV1StatementSpec`

NewSqlV1StatementSpecWithDefaults instantiates a new SqlV1StatementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *SqlV1StatementSpec) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *SqlV1StatementSpec) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *SqlV1StatementSpec) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *SqlV1StatementSpec) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetProperties

`func (o *SqlV1StatementSpec) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SqlV1StatementSpec) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SqlV1StatementSpec) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SqlV1StatementSpec) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetComputePoolId

`func (o *SqlV1StatementSpec) GetComputePoolId() string`

GetComputePoolId returns the ComputePoolId field if non-nil, zero value otherwise.

### GetComputePoolIdOk

`func (o *SqlV1StatementSpec) GetComputePoolIdOk() (*string, bool)`

GetComputePoolIdOk returns a tuple with the ComputePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePoolId

`func (o *SqlV1StatementSpec) SetComputePoolId(v string)`

SetComputePoolId sets ComputePoolId field to given value.

### HasComputePoolId

`func (o *SqlV1StatementSpec) HasComputePoolId() bool`

HasComputePoolId returns a boolean if a field has been set.

### GetPrincipal

`func (o *SqlV1StatementSpec) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SqlV1StatementSpec) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SqlV1StatementSpec) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *SqlV1StatementSpec) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetAuthorizedPrincipals

`func (o *SqlV1StatementSpec) GetAuthorizedPrincipals() []string`

GetAuthorizedPrincipals returns the AuthorizedPrincipals field if non-nil, zero value otherwise.

### GetAuthorizedPrincipalsOk

`func (o *SqlV1StatementSpec) GetAuthorizedPrincipalsOk() (*[]string, bool)`

GetAuthorizedPrincipalsOk returns a tuple with the AuthorizedPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedPrincipals

`func (o *SqlV1StatementSpec) SetAuthorizedPrincipals(v []string)`

SetAuthorizedPrincipals sets AuthorizedPrincipals field to given value.

### HasAuthorizedPrincipals

`func (o *SqlV1StatementSpec) HasAuthorizedPrincipals() bool`

HasAuthorizedPrincipals returns a boolean if a field has been set.

### GetStopped

`func (o *SqlV1StatementSpec) GetStopped() bool`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *SqlV1StatementSpec) GetStoppedOk() (*bool, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *SqlV1StatementSpec) SetStopped(v bool)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *SqlV1StatementSpec) HasStopped() bool`

HasStopped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



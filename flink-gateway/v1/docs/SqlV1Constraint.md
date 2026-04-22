# SqlV1Constraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of constraint. | [optional] 
**Columns** | Pointer to **[]string** |  | [optional] 
**Enforced** | Pointer to **bool** | Whether the constraint is enforced. | [optional] [default to false]

## Methods

### NewSqlV1Constraint

`func NewSqlV1Constraint() *SqlV1Constraint`

NewSqlV1Constraint instantiates a new SqlV1Constraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ConstraintWithDefaults

`func NewSqlV1ConstraintWithDefaults() *SqlV1Constraint`

NewSqlV1ConstraintWithDefaults instantiates a new SqlV1Constraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlV1Constraint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1Constraint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1Constraint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlV1Constraint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SqlV1Constraint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SqlV1Constraint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SqlV1Constraint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SqlV1Constraint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetColumns

`func (o *SqlV1Constraint) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *SqlV1Constraint) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *SqlV1Constraint) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *SqlV1Constraint) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetEnforced

`func (o *SqlV1Constraint) GetEnforced() bool`

GetEnforced returns the Enforced field if non-nil, zero value otherwise.

### GetEnforcedOk

`func (o *SqlV1Constraint) GetEnforcedOk() (*bool, bool)`

GetEnforcedOk returns a tuple with the Enforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforced

`func (o *SqlV1Constraint) SetEnforced(v bool)`

SetEnforced sets Enforced field to given value.

### HasEnforced

`func (o *SqlV1Constraint) HasEnforced() bool`

HasEnforced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



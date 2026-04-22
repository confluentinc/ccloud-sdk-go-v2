# SqlV1ComputedColumnAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **interface{}** |  | [optional] 
**Kind** | **string** | The kind of column. | 
**Expression** | **string** | The SQL expression used to compute the column value. | 
**Virtual** | Pointer to **bool** | Indicates if the computed column is virtual. | [optional] [default to false]

## Methods

### NewSqlV1ComputedColumnAllOf

`func NewSqlV1ComputedColumnAllOf(kind string, expression string, ) *SqlV1ComputedColumnAllOf`

NewSqlV1ComputedColumnAllOf instantiates a new SqlV1ComputedColumnAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ComputedColumnAllOfWithDefaults

`func NewSqlV1ComputedColumnAllOfWithDefaults() *SqlV1ComputedColumnAllOf`

NewSqlV1ComputedColumnAllOfWithDefaults instantiates a new SqlV1ComputedColumnAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlV1ComputedColumnAllOf) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1ComputedColumnAllOf) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1ComputedColumnAllOf) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *SqlV1ComputedColumnAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SqlV1ComputedColumnAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SqlV1ComputedColumnAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *SqlV1ComputedColumnAllOf) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SqlV1ComputedColumnAllOf) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SqlV1ComputedColumnAllOf) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *SqlV1ComputedColumnAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SqlV1ComputedColumnAllOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SqlV1ComputedColumnAllOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetKind

`func (o *SqlV1ComputedColumnAllOf) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1ComputedColumnAllOf) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1ComputedColumnAllOf) SetKind(v string)`

SetKind sets Kind field to given value.


### GetExpression

`func (o *SqlV1ComputedColumnAllOf) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SqlV1ComputedColumnAllOf) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SqlV1ComputedColumnAllOf) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetVirtual

`func (o *SqlV1ComputedColumnAllOf) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *SqlV1ComputedColumnAllOf) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *SqlV1ComputedColumnAllOf) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *SqlV1ComputedColumnAllOf) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



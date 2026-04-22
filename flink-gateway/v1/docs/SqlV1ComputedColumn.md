# SqlV1ComputedColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** |  | 
**Type** | **interface{}** |  | 
**Comment** | Pointer to **string** | A comment or description for the column. | [optional] 
**Kind** | **string** | The kind of column. | 
**Expression** | **string** | The SQL expression used to compute the column value. | 
**Virtual** | Pointer to **bool** | Indicates if the computed column is virtual. | [optional] [default to false]

## Methods

### NewSqlV1ComputedColumn

`func NewSqlV1ComputedColumn(name interface{}, type_ interface{}, kind string, expression string, ) *SqlV1ComputedColumn`

NewSqlV1ComputedColumn instantiates a new SqlV1ComputedColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ComputedColumnWithDefaults

`func NewSqlV1ComputedColumnWithDefaults() *SqlV1ComputedColumn`

NewSqlV1ComputedColumnWithDefaults instantiates a new SqlV1ComputedColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlV1ComputedColumn) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1ComputedColumn) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1ComputedColumn) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *SqlV1ComputedColumn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SqlV1ComputedColumn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *SqlV1ComputedColumn) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SqlV1ComputedColumn) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SqlV1ComputedColumn) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SqlV1ComputedColumn) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SqlV1ComputedColumn) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetComment

`func (o *SqlV1ComputedColumn) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SqlV1ComputedColumn) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SqlV1ComputedColumn) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SqlV1ComputedColumn) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetKind

`func (o *SqlV1ComputedColumn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1ComputedColumn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1ComputedColumn) SetKind(v string)`

SetKind sets Kind field to given value.


### GetExpression

`func (o *SqlV1ComputedColumn) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SqlV1ComputedColumn) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SqlV1ComputedColumn) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetVirtual

`func (o *SqlV1ComputedColumn) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *SqlV1ComputedColumn) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *SqlV1ComputedColumn) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *SqlV1ComputedColumn) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.


### AsSqlV1ColumnDetails

`func (s *SqlV1ComputedColumn) AsSqlV1ColumnDetails() SqlV1ColumnDetails`

Convenience method to wrap this instance of SqlV1ComputedColumn in SqlV1ColumnDetails

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



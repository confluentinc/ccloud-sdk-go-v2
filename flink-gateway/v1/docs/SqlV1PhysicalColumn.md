# SqlV1PhysicalColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** |  | 
**Type** | **interface{}** |  | 
**Comment** | Pointer to **string** | A comment or description for the column. | [optional] 
**Kind** | **string** | The kind of column. | 

## Methods

### NewSqlV1PhysicalColumn

`func NewSqlV1PhysicalColumn(name interface{}, type_ interface{}, kind string, ) *SqlV1PhysicalColumn`

NewSqlV1PhysicalColumn instantiates a new SqlV1PhysicalColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1PhysicalColumnWithDefaults

`func NewSqlV1PhysicalColumnWithDefaults() *SqlV1PhysicalColumn`

NewSqlV1PhysicalColumnWithDefaults instantiates a new SqlV1PhysicalColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlV1PhysicalColumn) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1PhysicalColumn) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1PhysicalColumn) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *SqlV1PhysicalColumn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SqlV1PhysicalColumn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *SqlV1PhysicalColumn) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SqlV1PhysicalColumn) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SqlV1PhysicalColumn) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SqlV1PhysicalColumn) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SqlV1PhysicalColumn) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetComment

`func (o *SqlV1PhysicalColumn) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SqlV1PhysicalColumn) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SqlV1PhysicalColumn) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SqlV1PhysicalColumn) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetKind

`func (o *SqlV1PhysicalColumn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1PhysicalColumn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1PhysicalColumn) SetKind(v string)`

SetKind sets Kind field to given value.



### AsSqlV1ColumnDetails

`func (s *SqlV1PhysicalColumn) AsSqlV1ColumnDetails() SqlV1ColumnDetails`

Convenience method to wrap this instance of SqlV1PhysicalColumn in SqlV1ColumnDetails

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



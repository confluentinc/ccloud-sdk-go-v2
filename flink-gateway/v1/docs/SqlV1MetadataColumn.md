# SqlV1MetadataColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** |  | 
**Type** | **interface{}** |  | 
**Comment** | Pointer to **string** | A comment or description for the column. | [optional] 
**Kind** | **string** | The kind of column. | 
**MetadataKey** | **string** | The system metadata key to reference. | 
**Virtual** | Pointer to **bool** | Indicates if the metadata column is virtual. | [optional] [default to false]

## Methods

### NewSqlV1MetadataColumn

`func NewSqlV1MetadataColumn(name interface{}, type_ interface{}, kind string, metadataKey string, ) *SqlV1MetadataColumn`

NewSqlV1MetadataColumn instantiates a new SqlV1MetadataColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1MetadataColumnWithDefaults

`func NewSqlV1MetadataColumnWithDefaults() *SqlV1MetadataColumn`

NewSqlV1MetadataColumnWithDefaults instantiates a new SqlV1MetadataColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlV1MetadataColumn) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1MetadataColumn) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1MetadataColumn) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *SqlV1MetadataColumn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SqlV1MetadataColumn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *SqlV1MetadataColumn) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SqlV1MetadataColumn) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SqlV1MetadataColumn) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SqlV1MetadataColumn) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SqlV1MetadataColumn) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetComment

`func (o *SqlV1MetadataColumn) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SqlV1MetadataColumn) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SqlV1MetadataColumn) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SqlV1MetadataColumn) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetKind

`func (o *SqlV1MetadataColumn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1MetadataColumn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1MetadataColumn) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadataKey

`func (o *SqlV1MetadataColumn) GetMetadataKey() string`

GetMetadataKey returns the MetadataKey field if non-nil, zero value otherwise.

### GetMetadataKeyOk

`func (o *SqlV1MetadataColumn) GetMetadataKeyOk() (*string, bool)`

GetMetadataKeyOk returns a tuple with the MetadataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataKey

`func (o *SqlV1MetadataColumn) SetMetadataKey(v string)`

SetMetadataKey sets MetadataKey field to given value.


### GetVirtual

`func (o *SqlV1MetadataColumn) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *SqlV1MetadataColumn) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *SqlV1MetadataColumn) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *SqlV1MetadataColumn) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.


### AsSqlV1ColumnDetails

`func (s *SqlV1MetadataColumn) AsSqlV1ColumnDetails() SqlV1ColumnDetails`

Convenience method to wrap this instance of SqlV1MetadataColumn in SqlV1ColumnDetails

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



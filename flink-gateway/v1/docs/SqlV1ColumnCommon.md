# SqlV1ColumnCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the column. | 
**Type** | [**DataType**](DataType.md) |  | 
**Comment** | Pointer to **string** | A comment or description for the column. | [optional] 

## Methods

### NewSqlV1ColumnCommon

`func NewSqlV1ColumnCommon(name string, type_ DataType, ) *SqlV1ColumnCommon`

NewSqlV1ColumnCommon instantiates a new SqlV1ColumnCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ColumnCommonWithDefaults

`func NewSqlV1ColumnCommonWithDefaults() *SqlV1ColumnCommon`

NewSqlV1ColumnCommonWithDefaults instantiates a new SqlV1ColumnCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlV1ColumnCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1ColumnCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1ColumnCommon) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SqlV1ColumnCommon) GetType() DataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SqlV1ColumnCommon) GetTypeOk() (*DataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SqlV1ColumnCommon) SetType(v DataType)`

SetType sets Type field to given value.


### GetComment

`func (o *SqlV1ColumnCommon) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SqlV1ColumnCommon) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SqlV1ColumnCommon) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SqlV1ColumnCommon) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



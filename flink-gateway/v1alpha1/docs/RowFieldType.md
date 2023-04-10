# RowFieldType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the field. | 
**Type** | [**DataType**](DataType.md) |  | 

## Methods

### NewRowFieldType

`func NewRowFieldType(name string, type_ DataType, ) *RowFieldType`

NewRowFieldType instantiates a new RowFieldType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRowFieldTypeWithDefaults

`func NewRowFieldTypeWithDefaults() *RowFieldType`

NewRowFieldTypeWithDefaults instantiates a new RowFieldType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RowFieldType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RowFieldType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RowFieldType) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RowFieldType) GetType() DataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RowFieldType) GetTypeOk() (*DataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RowFieldType) SetType(v DataType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



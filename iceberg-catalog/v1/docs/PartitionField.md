# PartitionField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **int32** |  | [optional] 
**SourceId** | **int32** |  | 
**Name** | **string** |  | 
**Transform** | **string** |  | 

## Methods

### NewPartitionField

`func NewPartitionField(sourceId int32, name string, transform string, ) *PartitionField`

NewPartitionField instantiates a new PartitionField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionFieldWithDefaults

`func NewPartitionFieldWithDefaults() *PartitionField`

NewPartitionFieldWithDefaults instantiates a new PartitionField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *PartitionField) GetFieldId() int32`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *PartitionField) GetFieldIdOk() (*int32, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *PartitionField) SetFieldId(v int32)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *PartitionField) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetSourceId

`func (o *PartitionField) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PartitionField) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PartitionField) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.


### GetName

`func (o *PartitionField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartitionField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartitionField) SetName(v string)`

SetName sets Name field to given value.


### GetTransform

`func (o *PartitionField) GetTransform() string`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *PartitionField) GetTransformOk() (*string, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *PartitionField) SetTransform(v string)`

SetTransform sets Transform field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



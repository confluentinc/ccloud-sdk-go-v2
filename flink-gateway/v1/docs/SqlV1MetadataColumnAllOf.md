# SqlV1MetadataColumnAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **interface{}** |  | [optional] 
**Kind** | **string** | The kind of column. | 
**MetadataKey** | **string** | The system metadata key to reference. | 
**Virtual** | Pointer to **bool** | Indicates if the metadata column is virtual. | [optional] [default to false]

## Methods

### NewSqlV1MetadataColumnAllOf

`func NewSqlV1MetadataColumnAllOf(kind string, metadataKey string, ) *SqlV1MetadataColumnAllOf`

NewSqlV1MetadataColumnAllOf instantiates a new SqlV1MetadataColumnAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1MetadataColumnAllOfWithDefaults

`func NewSqlV1MetadataColumnAllOfWithDefaults() *SqlV1MetadataColumnAllOf`

NewSqlV1MetadataColumnAllOfWithDefaults instantiates a new SqlV1MetadataColumnAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SqlV1MetadataColumnAllOf) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1MetadataColumnAllOf) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1MetadataColumnAllOf) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *SqlV1MetadataColumnAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SqlV1MetadataColumnAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SqlV1MetadataColumnAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *SqlV1MetadataColumnAllOf) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SqlV1MetadataColumnAllOf) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SqlV1MetadataColumnAllOf) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *SqlV1MetadataColumnAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SqlV1MetadataColumnAllOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SqlV1MetadataColumnAllOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetKind

`func (o *SqlV1MetadataColumnAllOf) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1MetadataColumnAllOf) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1MetadataColumnAllOf) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadataKey

`func (o *SqlV1MetadataColumnAllOf) GetMetadataKey() string`

GetMetadataKey returns the MetadataKey field if non-nil, zero value otherwise.

### GetMetadataKeyOk

`func (o *SqlV1MetadataColumnAllOf) GetMetadataKeyOk() (*string, bool)`

GetMetadataKeyOk returns a tuple with the MetadataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataKey

`func (o *SqlV1MetadataColumnAllOf) SetMetadataKey(v string)`

SetMetadataKey sets MetadataKey field to given value.


### GetVirtual

`func (o *SqlV1MetadataColumnAllOf) GetVirtual() bool`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### GetVirtualOk

`func (o *SqlV1MetadataColumnAllOf) GetVirtualOk() (*bool, bool)`

GetVirtualOk returns a tuple with the Virtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtual

`func (o *SqlV1MetadataColumnAllOf) SetVirtual(v bool)`

SetVirtual sets Virtual field to given value.

### HasVirtual

`func (o *SqlV1MetadataColumnAllOf) HasVirtual() bool`

HasVirtual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



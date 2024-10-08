# ObjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **string** | Self is a Uniform Resource Locator (URL) at which an object can be addressed. This URL encodes the service location, API version, and other particulars necessary to locate the resource at a point in time | 
**CreatedAt** | Pointer to **time.Time** | The date and time at which this object was created. It is represented in RFC3339 format and is in UTC. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time at which this object was last updated. It is represented in RFC3339 format and is in UTC. | [optional] 
**Uid** | Pointer to **string** | A system generated globally unique identifier for this resource. | [optional] 
**ResourceVersion** | Pointer to **string** | A system generated string that uniquely identifies the version of this resource. | [optional] 
**Labels** | Pointer to **map[string]string** | A map of key-value pairs that describe the resource. | [optional] 

## Methods

### NewObjectMeta

`func NewObjectMeta(self string, ) *ObjectMeta`

NewObjectMeta instantiates a new ObjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMetaWithDefaults

`func NewObjectMetaWithDefaults() *ObjectMeta`

NewObjectMetaWithDefaults instantiates a new ObjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ObjectMeta) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ObjectMeta) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ObjectMeta) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetCreatedAt

`func (o *ObjectMeta) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ObjectMeta) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ObjectMeta) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ObjectMeta) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ObjectMeta) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ObjectMeta) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ObjectMeta) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ObjectMeta) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUid

`func (o *ObjectMeta) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ObjectMeta) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ObjectMeta) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ObjectMeta) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetResourceVersion

`func (o *ObjectMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ObjectMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ObjectMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ObjectMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetLabels

`func (o *ObjectMeta) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ObjectMeta) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ObjectMeta) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ObjectMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StatementObjectMeta

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

### NewStatementObjectMeta

`func NewStatementObjectMeta(self string, ) *StatementObjectMeta`

NewStatementObjectMeta instantiates a new StatementObjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementObjectMetaWithDefaults

`func NewStatementObjectMetaWithDefaults() *StatementObjectMeta`

NewStatementObjectMetaWithDefaults instantiates a new StatementObjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *StatementObjectMeta) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *StatementObjectMeta) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *StatementObjectMeta) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetCreatedAt

`func (o *StatementObjectMeta) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StatementObjectMeta) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StatementObjectMeta) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StatementObjectMeta) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StatementObjectMeta) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StatementObjectMeta) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StatementObjectMeta) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StatementObjectMeta) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUid

`func (o *StatementObjectMeta) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *StatementObjectMeta) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *StatementObjectMeta) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *StatementObjectMeta) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetResourceVersion

`func (o *StatementObjectMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *StatementObjectMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *StatementObjectMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *StatementObjectMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetLabels

`func (o *StatementObjectMeta) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StatementObjectMeta) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StatementObjectMeta) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *StatementObjectMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



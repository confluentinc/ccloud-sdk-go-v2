# ObjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **string** | Self is a Uniform Resource Locator (URL) at which an object can be addressed. This URL encodes the service location, API version, and other particulars necessary to locate the resource at a point in time | [readonly] 
**ResourceName** | Pointer to **string** | Resource Name is a Uniform Resource Identifier (URI) that is globally unique across space and time. It is represented as a Confluent Resource Name | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which this object was created. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time at which this object was last updated. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**DeletedAt** | Pointer to **time.Time** | The date and time at which this object was (or will be) deleted. It is represented in RFC3339 format and is in UTC. | [optional] [readonly] 
**Name** | **string** | Client-supplied stable identifier for the policy. Used as the URL identifier and must be unique within the caller&#39;s organization. Immutable.  | 
**Uid** | Pointer to **string** | System-generated unique identifier for this policy. Distinct from &#x60;name&#x60;: while &#x60;name&#x60; is the client-supplied URL identifier, &#x60;uid&#x60; remains unique across time even if a policy with the same &#x60;name&#x60; is deleted and recreated.  | [optional] [readonly] 
**ResourceVersion** | Pointer to **string** | Opaque value identifying the current version of the policy. | [optional] [readonly] 

## Methods

### NewObjectMeta

`func NewObjectMeta(self string, name string, ) *ObjectMeta`

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


### GetResourceName

`func (o *ObjectMeta) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ObjectMeta) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ObjectMeta) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ObjectMeta) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

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

### GetDeletedAt

`func (o *ObjectMeta) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ObjectMeta) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ObjectMeta) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ObjectMeta) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetName

`func (o *ObjectMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectMeta) SetName(v string)`

SetName sets Name field to given value.


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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



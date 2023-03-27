# FrpmV2ResourcePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**FrpmV2ResourcePoolSpec**](FrpmV2ResourcePoolSpec.md) |  | [optional] 
**Status** | Pointer to [**FrpmV2ResourcePoolStatus**](FrpmV2ResourcePoolStatus.md) |  | [optional] 

## Methods

### NewFrpmV2ResourcePool

`func NewFrpmV2ResourcePool() *FrpmV2ResourcePool`

NewFrpmV2ResourcePool instantiates a new FrpmV2ResourcePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrpmV2ResourcePoolWithDefaults

`func NewFrpmV2ResourcePoolWithDefaults() *FrpmV2ResourcePool`

NewFrpmV2ResourcePoolWithDefaults instantiates a new FrpmV2ResourcePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FrpmV2ResourcePool) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FrpmV2ResourcePool) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FrpmV2ResourcePool) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FrpmV2ResourcePool) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *FrpmV2ResourcePool) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FrpmV2ResourcePool) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FrpmV2ResourcePool) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FrpmV2ResourcePool) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *FrpmV2ResourcePool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FrpmV2ResourcePool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FrpmV2ResourcePool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FrpmV2ResourcePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FrpmV2ResourcePool) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FrpmV2ResourcePool) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FrpmV2ResourcePool) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FrpmV2ResourcePool) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *FrpmV2ResourcePool) GetSpec() FrpmV2ResourcePoolSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *FrpmV2ResourcePool) GetSpecOk() (*FrpmV2ResourcePoolSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *FrpmV2ResourcePool) SetSpec(v FrpmV2ResourcePoolSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *FrpmV2ResourcePool) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *FrpmV2ResourcePool) GetStatus() FrpmV2ResourcePoolStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FrpmV2ResourcePool) GetStatusOk() (*FrpmV2ResourcePoolStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FrpmV2ResourcePool) SetStatus(v FrpmV2ResourcePoolStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FrpmV2ResourcePool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



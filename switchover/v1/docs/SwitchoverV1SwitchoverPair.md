# SwitchoverV1SwitchoverPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**SwitchoverV1SwitchoverPairSpec**](SwitchoverV1SwitchoverPairSpec.md) |  | [optional] 
**Status** | Pointer to [**SwitchoverV1SwitchoverPairStatus**](SwitchoverV1SwitchoverPairStatus.md) |  | [optional] 

## Methods

### NewSwitchoverV1SwitchoverPair

`func NewSwitchoverV1SwitchoverPair() *SwitchoverV1SwitchoverPair`

NewSwitchoverV1SwitchoverPair instantiates a new SwitchoverV1SwitchoverPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverPairWithDefaults

`func NewSwitchoverV1SwitchoverPairWithDefaults() *SwitchoverV1SwitchoverPair`

NewSwitchoverV1SwitchoverPairWithDefaults instantiates a new SwitchoverV1SwitchoverPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SwitchoverV1SwitchoverPair) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SwitchoverV1SwitchoverPair) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SwitchoverV1SwitchoverPair) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SwitchoverV1SwitchoverPair) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SwitchoverV1SwitchoverPair) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SwitchoverV1SwitchoverPair) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SwitchoverV1SwitchoverPair) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SwitchoverV1SwitchoverPair) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *SwitchoverV1SwitchoverPair) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SwitchoverV1SwitchoverPair) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SwitchoverV1SwitchoverPair) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SwitchoverV1SwitchoverPair) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *SwitchoverV1SwitchoverPair) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SwitchoverV1SwitchoverPair) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SwitchoverV1SwitchoverPair) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SwitchoverV1SwitchoverPair) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *SwitchoverV1SwitchoverPair) GetSpec() SwitchoverV1SwitchoverPairSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SwitchoverV1SwitchoverPair) GetSpecOk() (*SwitchoverV1SwitchoverPairSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SwitchoverV1SwitchoverPair) SetSpec(v SwitchoverV1SwitchoverPairSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SwitchoverV1SwitchoverPair) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SwitchoverV1SwitchoverPair) GetStatus() SwitchoverV1SwitchoverPairStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwitchoverV1SwitchoverPair) GetStatusOk() (*SwitchoverV1SwitchoverPairStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwitchoverV1SwitchoverPair) SetStatus(v SwitchoverV1SwitchoverPairStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SwitchoverV1SwitchoverPair) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



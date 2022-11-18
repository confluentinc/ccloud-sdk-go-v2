# SrcmV2ClusterUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**SrcmV2ClusterSpecUpdate**](SrcmV2ClusterSpecUpdate.md) |  | [optional] 
**Status** | Pointer to [**SrcmV2ClusterStatus**](SrcmV2ClusterStatus.md) |  | [optional] 

## Methods

### NewSrcmV2ClusterUpdate

`func NewSrcmV2ClusterUpdate() *SrcmV2ClusterUpdate`

NewSrcmV2ClusterUpdate instantiates a new SrcmV2ClusterUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrcmV2ClusterUpdateWithDefaults

`func NewSrcmV2ClusterUpdateWithDefaults() *SrcmV2ClusterUpdate`

NewSrcmV2ClusterUpdateWithDefaults instantiates a new SrcmV2ClusterUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SrcmV2ClusterUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SrcmV2ClusterUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SrcmV2ClusterUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SrcmV2ClusterUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SrcmV2ClusterUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SrcmV2ClusterUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SrcmV2ClusterUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SrcmV2ClusterUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *SrcmV2ClusterUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SrcmV2ClusterUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SrcmV2ClusterUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SrcmV2ClusterUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *SrcmV2ClusterUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SrcmV2ClusterUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SrcmV2ClusterUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SrcmV2ClusterUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *SrcmV2ClusterUpdate) GetSpec() SrcmV2ClusterSpecUpdate`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SrcmV2ClusterUpdate) GetSpecOk() (*SrcmV2ClusterSpecUpdate, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SrcmV2ClusterUpdate) SetSpec(v SrcmV2ClusterSpecUpdate)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SrcmV2ClusterUpdate) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SrcmV2ClusterUpdate) GetStatus() SrcmV2ClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SrcmV2ClusterUpdate) GetStatusOk() (*SrcmV2ClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SrcmV2ClusterUpdate) SetStatus(v SrcmV2ClusterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SrcmV2ClusterUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



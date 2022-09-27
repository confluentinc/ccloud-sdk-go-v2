# SdV1PipelineUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**SdV1PipelineSpecUpdate**](SdV1PipelineSpecUpdate.md) |  | [optional] 
**Status** | Pointer to [**SdV1PipelineStatus**](SdV1PipelineStatus.md) |  | [optional] 

## Methods

### NewSdV1PipelineUpdate

`func NewSdV1PipelineUpdate() *SdV1PipelineUpdate`

NewSdV1PipelineUpdate instantiates a new SdV1PipelineUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdV1PipelineUpdateWithDefaults

`func NewSdV1PipelineUpdateWithDefaults() *SdV1PipelineUpdate`

NewSdV1PipelineUpdateWithDefaults instantiates a new SdV1PipelineUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SdV1PipelineUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SdV1PipelineUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SdV1PipelineUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SdV1PipelineUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SdV1PipelineUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SdV1PipelineUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SdV1PipelineUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SdV1PipelineUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *SdV1PipelineUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SdV1PipelineUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SdV1PipelineUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SdV1PipelineUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *SdV1PipelineUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SdV1PipelineUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SdV1PipelineUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SdV1PipelineUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *SdV1PipelineUpdate) GetSpec() SdV1PipelineSpecUpdate`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SdV1PipelineUpdate) GetSpecOk() (*SdV1PipelineSpecUpdate, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SdV1PipelineUpdate) SetSpec(v SdV1PipelineSpecUpdate)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SdV1PipelineUpdate) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SdV1PipelineUpdate) GetStatus() SdV1PipelineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SdV1PipelineUpdate) GetStatusOk() (*SdV1PipelineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SdV1PipelineUpdate) SetStatus(v SdV1PipelineStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SdV1PipelineUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



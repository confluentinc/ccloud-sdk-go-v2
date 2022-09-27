# SdV1Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**SdV1PipelineSpec**](SdV1PipelineSpec.md) |  | [optional] 
**Status** | Pointer to [**SdV1PipelineStatus**](SdV1PipelineStatus.md) |  | [optional] 

## Methods

### NewSdV1Pipeline

`func NewSdV1Pipeline() *SdV1Pipeline`

NewSdV1Pipeline instantiates a new SdV1Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdV1PipelineWithDefaults

`func NewSdV1PipelineWithDefaults() *SdV1Pipeline`

NewSdV1PipelineWithDefaults instantiates a new SdV1Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SdV1Pipeline) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SdV1Pipeline) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SdV1Pipeline) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SdV1Pipeline) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SdV1Pipeline) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SdV1Pipeline) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SdV1Pipeline) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SdV1Pipeline) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *SdV1Pipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SdV1Pipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SdV1Pipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SdV1Pipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *SdV1Pipeline) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SdV1Pipeline) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SdV1Pipeline) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SdV1Pipeline) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *SdV1Pipeline) GetSpec() SdV1PipelineSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SdV1Pipeline) GetSpecOk() (*SdV1PipelineSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SdV1Pipeline) SetSpec(v SdV1PipelineSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SdV1Pipeline) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SdV1Pipeline) GetStatus() SdV1PipelineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SdV1Pipeline) GetStatusOk() (*SdV1PipelineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SdV1Pipeline) SetStatus(v SdV1PipelineStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SdV1Pipeline) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RtceV1RtceTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**RtceV1RtceTopicSpec**](RtceV1RtceTopicSpec.md) |  | [optional] 
**Status** | Pointer to [**RtceV1RtceTopicStatus**](RtceV1RtceTopicStatus.md) |  | [optional] 

## Methods

### NewRtceV1RtceTopic

`func NewRtceV1RtceTopic() *RtceV1RtceTopic`

NewRtceV1RtceTopic instantiates a new RtceV1RtceTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRtceV1RtceTopicWithDefaults

`func NewRtceV1RtceTopicWithDefaults() *RtceV1RtceTopic`

NewRtceV1RtceTopicWithDefaults instantiates a new RtceV1RtceTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *RtceV1RtceTopic) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RtceV1RtceTopic) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RtceV1RtceTopic) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RtceV1RtceTopic) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *RtceV1RtceTopic) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RtceV1RtceTopic) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RtceV1RtceTopic) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RtceV1RtceTopic) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *RtceV1RtceTopic) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RtceV1RtceTopic) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RtceV1RtceTopic) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RtceV1RtceTopic) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *RtceV1RtceTopic) GetSpec() RtceV1RtceTopicSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RtceV1RtceTopic) GetSpecOk() (*RtceV1RtceTopicSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RtceV1RtceTopic) SetSpec(v RtceV1RtceTopicSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RtceV1RtceTopic) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *RtceV1RtceTopic) GetStatus() RtceV1RtceTopicStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RtceV1RtceTopic) GetStatusOk() (*RtceV1RtceTopicStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RtceV1RtceTopic) SetStatus(v RtceV1RtceTopicStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RtceV1RtceTopic) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



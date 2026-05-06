# RtceV1RtceTopicUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**RtceV1RtceTopicSpecUpdate**](RtceV1RtceTopicSpecUpdate.md) |  | [optional] 
**Status** | Pointer to [**RtceV1RtceTopicStatus**](RtceV1RtceTopicStatus.md) |  | [optional] 

## Methods

### NewRtceV1RtceTopicUpdate

`func NewRtceV1RtceTopicUpdate() *RtceV1RtceTopicUpdate`

NewRtceV1RtceTopicUpdate instantiates a new RtceV1RtceTopicUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRtceV1RtceTopicUpdateWithDefaults

`func NewRtceV1RtceTopicUpdateWithDefaults() *RtceV1RtceTopicUpdate`

NewRtceV1RtceTopicUpdateWithDefaults instantiates a new RtceV1RtceTopicUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *RtceV1RtceTopicUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RtceV1RtceTopicUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RtceV1RtceTopicUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RtceV1RtceTopicUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *RtceV1RtceTopicUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RtceV1RtceTopicUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RtceV1RtceTopicUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RtceV1RtceTopicUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *RtceV1RtceTopicUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RtceV1RtceTopicUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RtceV1RtceTopicUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RtceV1RtceTopicUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *RtceV1RtceTopicUpdate) GetSpec() RtceV1RtceTopicSpecUpdate`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RtceV1RtceTopicUpdate) GetSpecOk() (*RtceV1RtceTopicSpecUpdate, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RtceV1RtceTopicUpdate) SetSpec(v RtceV1RtceTopicSpecUpdate)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RtceV1RtceTopicUpdate) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *RtceV1RtceTopicUpdate) GetStatus() RtceV1RtceTopicStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RtceV1RtceTopicUpdate) GetStatusOk() (*RtceV1RtceTopicStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RtceV1RtceTopicUpdate) SetStatus(v RtceV1RtceTopicStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RtceV1RtceTopicUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



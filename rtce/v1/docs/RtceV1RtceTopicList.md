# RtceV1RtceTopicList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]RtceV1RtceTopic**](RtceV1RtceTopic.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewRtceV1RtceTopicList

`func NewRtceV1RtceTopicList(apiVersion string, kind string, metadata ListMeta, data []RtceV1RtceTopic, ) *RtceV1RtceTopicList`

NewRtceV1RtceTopicList instantiates a new RtceV1RtceTopicList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRtceV1RtceTopicListWithDefaults

`func NewRtceV1RtceTopicListWithDefaults() *RtceV1RtceTopicList`

NewRtceV1RtceTopicListWithDefaults instantiates a new RtceV1RtceTopicList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *RtceV1RtceTopicList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RtceV1RtceTopicList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RtceV1RtceTopicList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *RtceV1RtceTopicList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RtceV1RtceTopicList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RtceV1RtceTopicList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *RtceV1RtceTopicList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RtceV1RtceTopicList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RtceV1RtceTopicList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *RtceV1RtceTopicList) GetData() []RtceV1RtceTopic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RtceV1RtceTopicList) GetDataOk() (*[]RtceV1RtceTopic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RtceV1RtceTopicList) SetData(v []RtceV1RtceTopic)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListMirrorTopicsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**LinkName** | **string** |  | 
**MirrorTopicName** | **string** |  | 
**SourceTopicName** | **string** |  | 
**NumPartitions** | **int32** |  | 
**MirrorLags** | [**MirrorLags**](MirrorLags.md) |  | 
**MirrorStatus** | [**MirrorTopicStatus**](MirrorTopicStatus.md) |  | 
**StateTimeMs** | **int64** |  | 
**MirrorStateTransitionErrors** | Pointer to [**[]LinkTaskError**](LinkTaskError.md) |  | [optional] 

## Methods

### NewListMirrorTopicsResponseData

`func NewListMirrorTopicsResponseData(kind string, metadata ResourceMetadata, linkName string, mirrorTopicName string, sourceTopicName string, numPartitions int32, mirrorLags MirrorLags, mirrorStatus MirrorTopicStatus, stateTimeMs int64, ) *ListMirrorTopicsResponseData`

NewListMirrorTopicsResponseData instantiates a new ListMirrorTopicsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMirrorTopicsResponseDataWithDefaults

`func NewListMirrorTopicsResponseDataWithDefaults() *ListMirrorTopicsResponseData`

NewListMirrorTopicsResponseDataWithDefaults instantiates a new ListMirrorTopicsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ListMirrorTopicsResponseData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListMirrorTopicsResponseData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListMirrorTopicsResponseData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListMirrorTopicsResponseData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListMirrorTopicsResponseData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListMirrorTopicsResponseData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetLinkName

`func (o *ListMirrorTopicsResponseData) GetLinkName() string`

GetLinkName returns the LinkName field if non-nil, zero value otherwise.

### GetLinkNameOk

`func (o *ListMirrorTopicsResponseData) GetLinkNameOk() (*string, bool)`

GetLinkNameOk returns a tuple with the LinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkName

`func (o *ListMirrorTopicsResponseData) SetLinkName(v string)`

SetLinkName sets LinkName field to given value.


### GetMirrorTopicName

`func (o *ListMirrorTopicsResponseData) GetMirrorTopicName() string`

GetMirrorTopicName returns the MirrorTopicName field if non-nil, zero value otherwise.

### GetMirrorTopicNameOk

`func (o *ListMirrorTopicsResponseData) GetMirrorTopicNameOk() (*string, bool)`

GetMirrorTopicNameOk returns a tuple with the MirrorTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorTopicName

`func (o *ListMirrorTopicsResponseData) SetMirrorTopicName(v string)`

SetMirrorTopicName sets MirrorTopicName field to given value.


### GetSourceTopicName

`func (o *ListMirrorTopicsResponseData) GetSourceTopicName() string`

GetSourceTopicName returns the SourceTopicName field if non-nil, zero value otherwise.

### GetSourceTopicNameOk

`func (o *ListMirrorTopicsResponseData) GetSourceTopicNameOk() (*string, bool)`

GetSourceTopicNameOk returns a tuple with the SourceTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTopicName

`func (o *ListMirrorTopicsResponseData) SetSourceTopicName(v string)`

SetSourceTopicName sets SourceTopicName field to given value.


### GetNumPartitions

`func (o *ListMirrorTopicsResponseData) GetNumPartitions() int32`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *ListMirrorTopicsResponseData) GetNumPartitionsOk() (*int32, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *ListMirrorTopicsResponseData) SetNumPartitions(v int32)`

SetNumPartitions sets NumPartitions field to given value.


### GetMirrorLags

`func (o *ListMirrorTopicsResponseData) GetMirrorLags() MirrorLags`

GetMirrorLags returns the MirrorLags field if non-nil, zero value otherwise.

### GetMirrorLagsOk

`func (o *ListMirrorTopicsResponseData) GetMirrorLagsOk() (*MirrorLags, bool)`

GetMirrorLagsOk returns a tuple with the MirrorLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorLags

`func (o *ListMirrorTopicsResponseData) SetMirrorLags(v MirrorLags)`

SetMirrorLags sets MirrorLags field to given value.


### GetMirrorStatus

`func (o *ListMirrorTopicsResponseData) GetMirrorStatus() MirrorTopicStatus`

GetMirrorStatus returns the MirrorStatus field if non-nil, zero value otherwise.

### GetMirrorStatusOk

`func (o *ListMirrorTopicsResponseData) GetMirrorStatusOk() (*MirrorTopicStatus, bool)`

GetMirrorStatusOk returns a tuple with the MirrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorStatus

`func (o *ListMirrorTopicsResponseData) SetMirrorStatus(v MirrorTopicStatus)`

SetMirrorStatus sets MirrorStatus field to given value.


### GetStateTimeMs

`func (o *ListMirrorTopicsResponseData) GetStateTimeMs() int64`

GetStateTimeMs returns the StateTimeMs field if non-nil, zero value otherwise.

### GetStateTimeMsOk

`func (o *ListMirrorTopicsResponseData) GetStateTimeMsOk() (*int64, bool)`

GetStateTimeMsOk returns a tuple with the StateTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTimeMs

`func (o *ListMirrorTopicsResponseData) SetStateTimeMs(v int64)`

SetStateTimeMs sets StateTimeMs field to given value.


### GetMirrorStateTransitionErrors

`func (o *ListMirrorTopicsResponseData) GetMirrorStateTransitionErrors() []LinkTaskError`

GetMirrorStateTransitionErrors returns the MirrorStateTransitionErrors field if non-nil, zero value otherwise.

### GetMirrorStateTransitionErrorsOk

`func (o *ListMirrorTopicsResponseData) GetMirrorStateTransitionErrorsOk() (*[]LinkTaskError, bool)`

GetMirrorStateTransitionErrorsOk returns a tuple with the MirrorStateTransitionErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorStateTransitionErrors

`func (o *ListMirrorTopicsResponseData) SetMirrorStateTransitionErrors(v []LinkTaskError)`

SetMirrorStateTransitionErrors sets MirrorStateTransitionErrors field to given value.

### HasMirrorStateTransitionErrors

`func (o *ListMirrorTopicsResponseData) HasMirrorStateTransitionErrors() bool`

HasMirrorStateTransitionErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListMirrorTopicsResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkName** | **string** |  | 
**MirrorTopicName** | **string** |  | 
**SourceTopicName** | **string** |  | 
**NumPartitions** | **int32** |  | 
**MirrorLags** | [**MirrorLags**](MirrorLags.md) |  | 
**MirrorStatus** | [**MirrorTopicStatus**](MirrorTopicStatus.md) |  | 
**StateTimeMs** | **int64** |  | 

## Methods

### NewListMirrorTopicsResponseDataAllOf

`func NewListMirrorTopicsResponseDataAllOf(linkName string, mirrorTopicName string, sourceTopicName string, numPartitions int32, mirrorLags MirrorLags, mirrorStatus MirrorTopicStatus, stateTimeMs int64, ) *ListMirrorTopicsResponseDataAllOf`

NewListMirrorTopicsResponseDataAllOf instantiates a new ListMirrorTopicsResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMirrorTopicsResponseDataAllOfWithDefaults

`func NewListMirrorTopicsResponseDataAllOfWithDefaults() *ListMirrorTopicsResponseDataAllOf`

NewListMirrorTopicsResponseDataAllOfWithDefaults instantiates a new ListMirrorTopicsResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkName

`func (o *ListMirrorTopicsResponseDataAllOf) GetLinkName() string`

GetLinkName returns the LinkName field if non-nil, zero value otherwise.

### GetLinkNameOk

`func (o *ListMirrorTopicsResponseDataAllOf) GetLinkNameOk() (*string, bool)`

GetLinkNameOk returns a tuple with the LinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkName

`func (o *ListMirrorTopicsResponseDataAllOf) SetLinkName(v string)`

SetLinkName sets LinkName field to given value.


### GetMirrorTopicName

`func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorTopicName() string`

GetMirrorTopicName returns the MirrorTopicName field if non-nil, zero value otherwise.

### GetMirrorTopicNameOk

`func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorTopicNameOk() (*string, bool)`

GetMirrorTopicNameOk returns a tuple with the MirrorTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorTopicName

`func (o *ListMirrorTopicsResponseDataAllOf) SetMirrorTopicName(v string)`

SetMirrorTopicName sets MirrorTopicName field to given value.


### GetSourceTopicName

`func (o *ListMirrorTopicsResponseDataAllOf) GetSourceTopicName() string`

GetSourceTopicName returns the SourceTopicName field if non-nil, zero value otherwise.

### GetSourceTopicNameOk

`func (o *ListMirrorTopicsResponseDataAllOf) GetSourceTopicNameOk() (*string, bool)`

GetSourceTopicNameOk returns a tuple with the SourceTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTopicName

`func (o *ListMirrorTopicsResponseDataAllOf) SetSourceTopicName(v string)`

SetSourceTopicName sets SourceTopicName field to given value.


### GetNumPartitions

`func (o *ListMirrorTopicsResponseDataAllOf) GetNumPartitions() int32`

GetNumPartitions returns the NumPartitions field if non-nil, zero value otherwise.

### GetNumPartitionsOk

`func (o *ListMirrorTopicsResponseDataAllOf) GetNumPartitionsOk() (*int32, bool)`

GetNumPartitionsOk returns a tuple with the NumPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitions

`func (o *ListMirrorTopicsResponseDataAllOf) SetNumPartitions(v int32)`

SetNumPartitions sets NumPartitions field to given value.


### GetMirrorLags

`func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorLags() MirrorLags`

GetMirrorLags returns the MirrorLags field if non-nil, zero value otherwise.

### GetMirrorLagsOk

`func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorLagsOk() (*MirrorLags, bool)`

GetMirrorLagsOk returns a tuple with the MirrorLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorLags

`func (o *ListMirrorTopicsResponseDataAllOf) SetMirrorLags(v MirrorLags)`

SetMirrorLags sets MirrorLags field to given value.


### GetMirrorStatus

`func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorStatus() MirrorTopicStatus`

GetMirrorStatus returns the MirrorStatus field if non-nil, zero value otherwise.

### GetMirrorStatusOk

`func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorStatusOk() (*MirrorTopicStatus, bool)`

GetMirrorStatusOk returns a tuple with the MirrorStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorStatus

`func (o *ListMirrorTopicsResponseDataAllOf) SetMirrorStatus(v MirrorTopicStatus)`

SetMirrorStatus sets MirrorStatus field to given value.


### GetStateTimeMs

`func (o *ListMirrorTopicsResponseDataAllOf) GetStateTimeMs() int64`

GetStateTimeMs returns the StateTimeMs field if non-nil, zero value otherwise.

### GetStateTimeMsOk

`func (o *ListMirrorTopicsResponseDataAllOf) GetStateTimeMsOk() (*int64, bool)`

GetStateTimeMsOk returns a tuple with the StateTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTimeMs

`func (o *ListMirrorTopicsResponseDataAllOf) SetStateTimeMs(v int64)`

SetStateTimeMs sets StateTimeMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



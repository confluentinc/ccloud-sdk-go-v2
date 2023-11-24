# CreateMirrorTopicRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceTopicName** | **string** |  | 
**MirrorTopicName** | Pointer to **string** |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**Configs** | Pointer to [**[]ConfigData**](ConfigData.md) |  | [optional] 

## Methods

### NewCreateMirrorTopicRequestData

`func NewCreateMirrorTopicRequestData(sourceTopicName string, ) *CreateMirrorTopicRequestData`

NewCreateMirrorTopicRequestData instantiates a new CreateMirrorTopicRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMirrorTopicRequestDataWithDefaults

`func NewCreateMirrorTopicRequestDataWithDefaults() *CreateMirrorTopicRequestData`

NewCreateMirrorTopicRequestDataWithDefaults instantiates a new CreateMirrorTopicRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceTopicName

`func (o *CreateMirrorTopicRequestData) GetSourceTopicName() string`

GetSourceTopicName returns the SourceTopicName field if non-nil, zero value otherwise.

### GetSourceTopicNameOk

`func (o *CreateMirrorTopicRequestData) GetSourceTopicNameOk() (*string, bool)`

GetSourceTopicNameOk returns a tuple with the SourceTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTopicName

`func (o *CreateMirrorTopicRequestData) SetSourceTopicName(v string)`

SetSourceTopicName sets SourceTopicName field to given value.


### GetMirrorTopicName

`func (o *CreateMirrorTopicRequestData) GetMirrorTopicName() string`

GetMirrorTopicName returns the MirrorTopicName field if non-nil, zero value otherwise.

### GetMirrorTopicNameOk

`func (o *CreateMirrorTopicRequestData) GetMirrorTopicNameOk() (*string, bool)`

GetMirrorTopicNameOk returns a tuple with the MirrorTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorTopicName

`func (o *CreateMirrorTopicRequestData) SetMirrorTopicName(v string)`

SetMirrorTopicName sets MirrorTopicName field to given value.

### HasMirrorTopicName

`func (o *CreateMirrorTopicRequestData) HasMirrorTopicName() bool`

HasMirrorTopicName returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *CreateMirrorTopicRequestData) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *CreateMirrorTopicRequestData) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *CreateMirrorTopicRequestData) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *CreateMirrorTopicRequestData) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetConfigs

`func (o *CreateMirrorTopicRequestData) GetConfigs() []ConfigData`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *CreateMirrorTopicRequestData) GetConfigsOk() (*[]ConfigData, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *CreateMirrorTopicRequestData) SetConfigs(v []ConfigData)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *CreateMirrorTopicRequestData) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



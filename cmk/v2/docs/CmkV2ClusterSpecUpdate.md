# CmkV2ClusterSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the cluster. | [optional] 
**Config** | Pointer to [**CmkV2ClusterSpecUpdateConfigOneOf**](CmkV2ClusterSpecUpdateConfigOneOf.md) | The configuration of the Kafka cluster.  Note: Clusters can be upgraded from Basic to Standard, but cannot be downgraded from Standard to Basic.  | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewCmkV2ClusterSpecUpdate

`func NewCmkV2ClusterSpecUpdate() *CmkV2ClusterSpecUpdate`

NewCmkV2ClusterSpecUpdate instantiates a new CmkV2ClusterSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2ClusterSpecUpdateWithDefaults

`func NewCmkV2ClusterSpecUpdateWithDefaults() *CmkV2ClusterSpecUpdate`

NewCmkV2ClusterSpecUpdateWithDefaults instantiates a new CmkV2ClusterSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CmkV2ClusterSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CmkV2ClusterSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CmkV2ClusterSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CmkV2ClusterSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetConfig

`func (o *CmkV2ClusterSpecUpdate) GetConfig() CmkV2ClusterSpecUpdateConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CmkV2ClusterSpecUpdate) GetConfigOk() (*CmkV2ClusterSpecUpdateConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CmkV2ClusterSpecUpdate) SetConfig(v CmkV2ClusterSpecUpdateConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CmkV2ClusterSpecUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *CmkV2ClusterSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CmkV2ClusterSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CmkV2ClusterSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CmkV2ClusterSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



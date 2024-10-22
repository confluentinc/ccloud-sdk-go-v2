# NetworkingV1AccessPointSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the access point. | [optional] 
**Config** | Pointer to [**NetworkingV1AccessPointSpecUpdateConfigOneOf**](NetworkingV1AccessPointSpecUpdateConfigOneOf.md) | The specific details of the different access point configurations. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1AccessPointSpecUpdate

`func NewNetworkingV1AccessPointSpecUpdate() *NetworkingV1AccessPointSpecUpdate`

NewNetworkingV1AccessPointSpecUpdate instantiates a new NetworkingV1AccessPointSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1AccessPointSpecUpdateWithDefaults

`func NewNetworkingV1AccessPointSpecUpdateWithDefaults() *NetworkingV1AccessPointSpecUpdate`

NewNetworkingV1AccessPointSpecUpdateWithDefaults instantiates a new NetworkingV1AccessPointSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1AccessPointSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1AccessPointSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1AccessPointSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1AccessPointSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkingV1AccessPointSpecUpdate) GetConfig() NetworkingV1AccessPointSpecUpdateConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkingV1AccessPointSpecUpdate) GetConfigOk() (*NetworkingV1AccessPointSpecUpdateConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkingV1AccessPointSpecUpdate) SetConfig(v NetworkingV1AccessPointSpecUpdateConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkingV1AccessPointSpecUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1AccessPointSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1AccessPointSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1AccessPointSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1AccessPointSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



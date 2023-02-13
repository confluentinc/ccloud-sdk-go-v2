# NetworkingV1NetworkLinkEndpointSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network link endpoint | [optional] 
**Description** | Pointer to **string** | The description of the network link endpoint | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1NetworkLinkEndpointSpecUpdate

`func NewNetworkingV1NetworkLinkEndpointSpecUpdate() *NetworkingV1NetworkLinkEndpointSpecUpdate`

NewNetworkingV1NetworkLinkEndpointSpecUpdate instantiates a new NetworkingV1NetworkLinkEndpointSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkEndpointSpecUpdateWithDefaults

`func NewNetworkingV1NetworkLinkEndpointSpecUpdateWithDefaults() *NetworkingV1NetworkLinkEndpointSpecUpdate`

NewNetworkingV1NetworkLinkEndpointSpecUpdateWithDefaults instantiates a new NetworkingV1NetworkLinkEndpointSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkLinkEndpointSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



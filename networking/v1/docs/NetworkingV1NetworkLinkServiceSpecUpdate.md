# NetworkingV1NetworkLinkServiceSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network link service | [optional] 
**Description** | Pointer to **string** | The description of the network link service | [optional] 
**Accept** | Pointer to [**NetworkingV1NetworkLinkServiceAcceptPolicy**](networking.v1.NetworkLinkServiceAcceptPolicy.md) | Network Link Service Accept policy | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1NetworkLinkServiceSpecUpdate

`func NewNetworkingV1NetworkLinkServiceSpecUpdate() *NetworkingV1NetworkLinkServiceSpecUpdate`

NewNetworkingV1NetworkLinkServiceSpecUpdate instantiates a new NetworkingV1NetworkLinkServiceSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkServiceSpecUpdateWithDefaults

`func NewNetworkingV1NetworkLinkServiceSpecUpdateWithDefaults() *NetworkingV1NetworkLinkServiceSpecUpdate`

NewNetworkingV1NetworkLinkServiceSpecUpdateWithDefaults instantiates a new NetworkingV1NetworkLinkServiceSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccept

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetAccept() NetworkingV1NetworkLinkServiceAcceptPolicy`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetAcceptOk() (*NetworkingV1NetworkLinkServiceAcceptPolicy, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) SetAccept(v NetworkingV1NetworkLinkServiceAcceptPolicy)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) HasAccept() bool`

HasAccept returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkLinkServiceSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetworkingV1NetworkSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the network | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**Gateway** | Pointer to [**NullableTypedEnvScopedObjectReference**](TypedEnvScopedObjectReference.md) | The gateway associated with this object. The gateway can be one of networking.v1.Gateway. May be &#x60;null&#x60; or omitted if not associated with a gateway. | [optional] 

## Methods

### NewNetworkingV1NetworkSpecUpdate

`func NewNetworkingV1NetworkSpecUpdate() *NetworkingV1NetworkSpecUpdate`

NewNetworkingV1NetworkSpecUpdate instantiates a new NetworkingV1NetworkSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkSpecUpdateWithDefaults

`func NewNetworkingV1NetworkSpecUpdateWithDefaults() *NetworkingV1NetworkSpecUpdate`

NewNetworkingV1NetworkSpecUpdateWithDefaults instantiates a new NetworkingV1NetworkSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1NetworkSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1NetworkSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1NetworkSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1NetworkSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1NetworkSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1NetworkSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1NetworkSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1NetworkSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetGateway

`func (o *NetworkingV1NetworkSpecUpdate) GetGateway() TypedEnvScopedObjectReference`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkingV1NetworkSpecUpdate) GetGatewayOk() (*TypedEnvScopedObjectReference, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkingV1NetworkSpecUpdate) SetGateway(v TypedEnvScopedObjectReference)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkingV1NetworkSpecUpdate) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *NetworkingV1NetworkSpecUpdate) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *NetworkingV1NetworkSpecUpdate) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SwitchoverV1SwitchoverEndpointSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentCrn** | Pointer to **string** | The CRN of the environment this switchover endpoint belongs to — always its parent pair&#39;s environment, derived from &#x60;parent_resource_crn&#x60; on create (ORC-9794 decision log). Read-only: accepting a second copy of the same scope would only create the possibility of the two disagreeing. It is also the value whose ID is passed as the &#x60;?environment&#x3D;&#x60; query parameter on the other operations.  | [optional] [readonly] 
**ParentResourceCrn** | Pointer to **string** | The CRN of the switchover pair this endpoint is bound to. Must name the same environment as &#x60;environment_crn&#x60;: the pair is looked up in that scope, so a mismatch would report a confusing not-found rather than the inconsistency the caller introduced.  | [optional] 
**DisplayName** | Pointer to **string** | A human-readable name for the switchover endpoint. | [optional] 
**Target** | Pointer to **string** | The name of the endpoint that should be active. For stateful pairs the control plane owns this value (it follows the pair&#39;s active member) and customers cannot set it directly; on create it may be provided as an initial value. Must match one of &#x60;endpoints[].name&#x60;.  | [optional] 
**Endpoints** | Pointer to [**[]SwitchoverV1EndpointConfig**](SwitchoverV1EndpointConfig.md) | The endpoint definitions, one per side (e.g. west/east). Must contain exactly 2 entries.  | [optional] 

## Methods

### NewSwitchoverV1SwitchoverEndpointSpec

`func NewSwitchoverV1SwitchoverEndpointSpec() *SwitchoverV1SwitchoverEndpointSpec`

NewSwitchoverV1SwitchoverEndpointSpec instantiates a new SwitchoverV1SwitchoverEndpointSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1SwitchoverEndpointSpecWithDefaults

`func NewSwitchoverV1SwitchoverEndpointSpecWithDefaults() *SwitchoverV1SwitchoverEndpointSpec`

NewSwitchoverV1SwitchoverEndpointSpecWithDefaults instantiates a new SwitchoverV1SwitchoverEndpointSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentCrn

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetEnvironmentCrn() string`

GetEnvironmentCrn returns the EnvironmentCrn field if non-nil, zero value otherwise.

### GetEnvironmentCrnOk

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetEnvironmentCrnOk() (*string, bool)`

GetEnvironmentCrnOk returns a tuple with the EnvironmentCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCrn

`func (o *SwitchoverV1SwitchoverEndpointSpec) SetEnvironmentCrn(v string)`

SetEnvironmentCrn sets EnvironmentCrn field to given value.

### HasEnvironmentCrn

`func (o *SwitchoverV1SwitchoverEndpointSpec) HasEnvironmentCrn() bool`

HasEnvironmentCrn returns a boolean if a field has been set.

### GetParentResourceCrn

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetParentResourceCrn() string`

GetParentResourceCrn returns the ParentResourceCrn field if non-nil, zero value otherwise.

### GetParentResourceCrnOk

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetParentResourceCrnOk() (*string, bool)`

GetParentResourceCrnOk returns a tuple with the ParentResourceCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceCrn

`func (o *SwitchoverV1SwitchoverEndpointSpec) SetParentResourceCrn(v string)`

SetParentResourceCrn sets ParentResourceCrn field to given value.

### HasParentResourceCrn

`func (o *SwitchoverV1SwitchoverEndpointSpec) HasParentResourceCrn() bool`

HasParentResourceCrn returns a boolean if a field has been set.

### GetDisplayName

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SwitchoverV1SwitchoverEndpointSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SwitchoverV1SwitchoverEndpointSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTarget

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SwitchoverV1SwitchoverEndpointSpec) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *SwitchoverV1SwitchoverEndpointSpec) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetEndpoints

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetEndpoints() []SwitchoverV1EndpointConfig`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *SwitchoverV1SwitchoverEndpointSpec) GetEndpointsOk() (*[]SwitchoverV1EndpointConfig, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *SwitchoverV1SwitchoverEndpointSpec) SetEndpoints(v []SwitchoverV1EndpointConfig)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *SwitchoverV1SwitchoverEndpointSpec) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



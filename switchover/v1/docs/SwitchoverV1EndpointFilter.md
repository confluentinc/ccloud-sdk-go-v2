# SwitchoverV1EndpointFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Whether the endpoint is private or public. | 
**NetworkCrn** | Pointer to **string** | CRN of the network to resolve this endpoint against. A private endpoint sets exactly one of this or &#x60;access_point_crn&#x60; — each independently resolves a unique endpoint, so a second selector is never needed and could only disagree with the first. A public endpoint sets neither.  | [optional] 
**AccessPointCrn** | Pointer to **string** | CRN of the network access point, for access-point (PNI) endpoints — the alternative to &#x60;network_crn&#x60;; a private endpoint sets exactly one of the two. The &#x60;gateway&#x60; segment is part of the registered canonical form and is required.  | [optional] 

## Methods

### NewSwitchoverV1EndpointFilter

`func NewSwitchoverV1EndpointFilter(type_ string, ) *SwitchoverV1EndpointFilter`

NewSwitchoverV1EndpointFilter instantiates a new SwitchoverV1EndpointFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchoverV1EndpointFilterWithDefaults

`func NewSwitchoverV1EndpointFilterWithDefaults() *SwitchoverV1EndpointFilter`

NewSwitchoverV1EndpointFilterWithDefaults instantiates a new SwitchoverV1EndpointFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SwitchoverV1EndpointFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SwitchoverV1EndpointFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SwitchoverV1EndpointFilter) SetType(v string)`

SetType sets Type field to given value.


### GetNetworkCrn

`func (o *SwitchoverV1EndpointFilter) GetNetworkCrn() string`

GetNetworkCrn returns the NetworkCrn field if non-nil, zero value otherwise.

### GetNetworkCrnOk

`func (o *SwitchoverV1EndpointFilter) GetNetworkCrnOk() (*string, bool)`

GetNetworkCrnOk returns a tuple with the NetworkCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCrn

`func (o *SwitchoverV1EndpointFilter) SetNetworkCrn(v string)`

SetNetworkCrn sets NetworkCrn field to given value.

### HasNetworkCrn

`func (o *SwitchoverV1EndpointFilter) HasNetworkCrn() bool`

HasNetworkCrn returns a boolean if a field has been set.

### GetAccessPointCrn

`func (o *SwitchoverV1EndpointFilter) GetAccessPointCrn() string`

GetAccessPointCrn returns the AccessPointCrn field if non-nil, zero value otherwise.

### GetAccessPointCrnOk

`func (o *SwitchoverV1EndpointFilter) GetAccessPointCrnOk() (*string, bool)`

GetAccessPointCrnOk returns a tuple with the AccessPointCrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPointCrn

`func (o *SwitchoverV1EndpointFilter) SetAccessPointCrn(v string)`

SetAccessPointCrn sets AccessPointCrn field to given value.

### HasAccessPointCrn

`func (o *SwitchoverV1EndpointFilter) HasAccessPointCrn() bool`

HasAccessPointCrn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



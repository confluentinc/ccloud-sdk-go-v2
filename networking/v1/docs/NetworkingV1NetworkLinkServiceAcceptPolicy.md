# NetworkingV1NetworkLinkServiceAcceptPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | Pointer to **[]string** | List of environments from which connections can be accepted. All networks win the list of environment will be allowed.  | [optional] 
**Networks** | Pointer to **[]string** | List of networks from which connections can be accepted.  | [optional] 

## Methods

### NewNetworkingV1NetworkLinkServiceAcceptPolicy

`func NewNetworkingV1NetworkLinkServiceAcceptPolicy() *NetworkingV1NetworkLinkServiceAcceptPolicy`

NewNetworkingV1NetworkLinkServiceAcceptPolicy instantiates a new NetworkingV1NetworkLinkServiceAcceptPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkServiceAcceptPolicyWithDefaults

`func NewNetworkingV1NetworkLinkServiceAcceptPolicyWithDefaults() *NetworkingV1NetworkLinkServiceAcceptPolicy`

NewNetworkingV1NetworkLinkServiceAcceptPolicyWithDefaults instantiates a new NetworkingV1NetworkLinkServiceAcceptPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *NetworkingV1NetworkLinkServiceAcceptPolicy) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *NetworkingV1NetworkLinkServiceAcceptPolicy) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *NetworkingV1NetworkLinkServiceAcceptPolicy) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *NetworkingV1NetworkLinkServiceAcceptPolicy) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetNetworks

`func (o *NetworkingV1NetworkLinkServiceAcceptPolicy) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NetworkingV1NetworkLinkServiceAcceptPolicy) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NetworkingV1NetworkLinkServiceAcceptPolicy) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NetworkingV1NetworkLinkServiceAcceptPolicy) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



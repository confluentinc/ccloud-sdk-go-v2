# NetworkingV1DnsForwarderSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the DNS forwarder | [optional] 
**Domains** | Pointer to **[]string** | List of domains for the DNS forwarder to use | [optional] 
**Config** | Pointer to [**NetworkingV1DnsForwarderSpecUpdateConfigOneOf**](NetworkingV1DnsForwarderSpecUpdateConfigOneOf.md) | The specific details of different kinds of configuration for DNS Forwarder. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1DnsForwarderSpecUpdate

`func NewNetworkingV1DnsForwarderSpecUpdate() *NetworkingV1DnsForwarderSpecUpdate`

NewNetworkingV1DnsForwarderSpecUpdate instantiates a new NetworkingV1DnsForwarderSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1DnsForwarderSpecUpdateWithDefaults

`func NewNetworkingV1DnsForwarderSpecUpdateWithDefaults() *NetworkingV1DnsForwarderSpecUpdate`

NewNetworkingV1DnsForwarderSpecUpdateWithDefaults instantiates a new NetworkingV1DnsForwarderSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1DnsForwarderSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1DnsForwarderSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1DnsForwarderSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1DnsForwarderSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDomains

`func (o *NetworkingV1DnsForwarderSpecUpdate) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *NetworkingV1DnsForwarderSpecUpdate) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *NetworkingV1DnsForwarderSpecUpdate) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *NetworkingV1DnsForwarderSpecUpdate) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkingV1DnsForwarderSpecUpdate) GetConfig() NetworkingV1DnsForwarderSpecUpdateConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkingV1DnsForwarderSpecUpdate) GetConfigOk() (*NetworkingV1DnsForwarderSpecUpdateConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkingV1DnsForwarderSpecUpdate) SetConfig(v NetworkingV1DnsForwarderSpecUpdateConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkingV1DnsForwarderSpecUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1DnsForwarderSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1DnsForwarderSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1DnsForwarderSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1DnsForwarderSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



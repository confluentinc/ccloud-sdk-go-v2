# NetworkingV1DnsForwarderSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | List of domains for the DNS forwarder to use | [optional] 
**Config** | Pointer to [**NetworkingV1DnsForwarderSpecConfigOneOf**](NetworkingV1DnsForwarderSpecConfigOneOf.md) | The specific details of different kinds of configuration for DNS Forwarder. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**Gateway** | Pointer to [**ObjectReference**](ObjectReference.md) | The gateway to which this belongs. | [optional] 

## Methods

### NewNetworkingV1DnsForwarderSpec

`func NewNetworkingV1DnsForwarderSpec() *NetworkingV1DnsForwarderSpec`

NewNetworkingV1DnsForwarderSpec instantiates a new NetworkingV1DnsForwarderSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1DnsForwarderSpecWithDefaults

`func NewNetworkingV1DnsForwarderSpecWithDefaults() *NetworkingV1DnsForwarderSpec`

NewNetworkingV1DnsForwarderSpecWithDefaults instantiates a new NetworkingV1DnsForwarderSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *NetworkingV1DnsForwarderSpec) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *NetworkingV1DnsForwarderSpec) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *NetworkingV1DnsForwarderSpec) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *NetworkingV1DnsForwarderSpec) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkingV1DnsForwarderSpec) GetConfig() NetworkingV1DnsForwarderSpecConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkingV1DnsForwarderSpec) GetConfigOk() (*NetworkingV1DnsForwarderSpecConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkingV1DnsForwarderSpec) SetConfig(v NetworkingV1DnsForwarderSpecConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkingV1DnsForwarderSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1DnsForwarderSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1DnsForwarderSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1DnsForwarderSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1DnsForwarderSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetGateway

`func (o *NetworkingV1DnsForwarderSpec) GetGateway() ObjectReference`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkingV1DnsForwarderSpec) GetGatewayOk() (*ObjectReference, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkingV1DnsForwarderSpec) SetGateway(v ObjectReference)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkingV1DnsForwarderSpec) HasGateway() bool`

HasGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetworkingV1ForwardViaIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | DNS Forwarder Configured via DNS Server IPs kind type. | 
**DnsServerIps** | **[]string** | List of IP addresses of the DNS server | 

## Methods

### NewNetworkingV1ForwardViaIp

`func NewNetworkingV1ForwardViaIp(kind string, dnsServerIps []string, ) *NetworkingV1ForwardViaIp`

NewNetworkingV1ForwardViaIp instantiates a new NetworkingV1ForwardViaIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1ForwardViaIpWithDefaults

`func NewNetworkingV1ForwardViaIpWithDefaults() *NetworkingV1ForwardViaIp`

NewNetworkingV1ForwardViaIpWithDefaults instantiates a new NetworkingV1ForwardViaIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1ForwardViaIp) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1ForwardViaIp) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1ForwardViaIp) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDnsServerIps

`func (o *NetworkingV1ForwardViaIp) GetDnsServerIps() []string`

GetDnsServerIps returns the DnsServerIps field if non-nil, zero value otherwise.

### GetDnsServerIpsOk

`func (o *NetworkingV1ForwardViaIp) GetDnsServerIpsOk() (*[]string, bool)`

GetDnsServerIpsOk returns a tuple with the DnsServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerIps

`func (o *NetworkingV1ForwardViaIp) SetDnsServerIps(v []string)`

SetDnsServerIps sets DnsServerIps field to given value.



### AsNetworkingV1DnsForwarderSpecConfigOneOf

`func (s *NetworkingV1ForwardViaIp) AsNetworkingV1DnsForwarderSpecConfigOneOf() NetworkingV1DnsForwarderSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1ForwardViaIp in NetworkingV1DnsForwarderSpecConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



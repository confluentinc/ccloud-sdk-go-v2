# NetworkingV1ForwardViaGcpDnsZones

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | DNS Forwarder Configured via GCP DNS Zones kind type. | 
**DomainMappings** | [**map[string]NetworkingV1ForwardViaGcpDnsZonesDomainMappings**](NetworkingV1ForwardViaGcpDnsZonesDomainMappings.md) | Mapping of domain names to GCP DNS Zones and Project ID. | 

## Methods

### NewNetworkingV1ForwardViaGcpDnsZones

`func NewNetworkingV1ForwardViaGcpDnsZones(kind string, domainMappings map[string]NetworkingV1ForwardViaGcpDnsZonesDomainMappings, ) *NetworkingV1ForwardViaGcpDnsZones`

NewNetworkingV1ForwardViaGcpDnsZones instantiates a new NetworkingV1ForwardViaGcpDnsZones object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1ForwardViaGcpDnsZonesWithDefaults

`func NewNetworkingV1ForwardViaGcpDnsZonesWithDefaults() *NetworkingV1ForwardViaGcpDnsZones`

NewNetworkingV1ForwardViaGcpDnsZonesWithDefaults instantiates a new NetworkingV1ForwardViaGcpDnsZones object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1ForwardViaGcpDnsZones) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1ForwardViaGcpDnsZones) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1ForwardViaGcpDnsZones) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDomainMappings

`func (o *NetworkingV1ForwardViaGcpDnsZones) GetDomainMappings() map[string]NetworkingV1ForwardViaGcpDnsZonesDomainMappings`

GetDomainMappings returns the DomainMappings field if non-nil, zero value otherwise.

### GetDomainMappingsOk

`func (o *NetworkingV1ForwardViaGcpDnsZones) GetDomainMappingsOk() (*map[string]NetworkingV1ForwardViaGcpDnsZonesDomainMappings, bool)`

GetDomainMappingsOk returns a tuple with the DomainMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainMappings

`func (o *NetworkingV1ForwardViaGcpDnsZones) SetDomainMappings(v map[string]NetworkingV1ForwardViaGcpDnsZonesDomainMappings)`

SetDomainMappings sets DomainMappings field to given value.



### AsNetworkingV1DnsForwarderSpecConfigOneOf

`func (s *NetworkingV1ForwardViaGcpDnsZones) AsNetworkingV1DnsForwarderSpecConfigOneOf() NetworkingV1DnsForwarderSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1ForwardViaGcpDnsZones in NetworkingV1DnsForwarderSpecConfigOneOf

### AsNetworkingV1DnsForwarderSpecUpdateConfigOneOf

`func (s *NetworkingV1ForwardViaGcpDnsZones) AsNetworkingV1DnsForwarderSpecUpdateConfigOneOf() NetworkingV1DnsForwarderSpecUpdateConfigOneOf`

Convenience method to wrap this instance of NetworkingV1ForwardViaGcpDnsZones in NetworkingV1DnsForwarderSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



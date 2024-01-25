# NetworkingV1PrivateLinkAccessPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | PrivateLinkAccessPoint kind. | 
**ResourceId** | **string** | Id of the target resource. | 

## Methods

### NewNetworkingV1PrivateLinkAccessPoint

`func NewNetworkingV1PrivateLinkAccessPoint(kind string, resourceId string, ) *NetworkingV1PrivateLinkAccessPoint`

NewNetworkingV1PrivateLinkAccessPoint instantiates a new NetworkingV1PrivateLinkAccessPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1PrivateLinkAccessPointWithDefaults

`func NewNetworkingV1PrivateLinkAccessPointWithDefaults() *NetworkingV1PrivateLinkAccessPoint`

NewNetworkingV1PrivateLinkAccessPointWithDefaults instantiates a new NetworkingV1PrivateLinkAccessPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1PrivateLinkAccessPoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1PrivateLinkAccessPoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1PrivateLinkAccessPoint) SetKind(v string)`

SetKind sets Kind field to given value.


### GetResourceId

`func (o *NetworkingV1PrivateLinkAccessPoint) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *NetworkingV1PrivateLinkAccessPoint) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *NetworkingV1PrivateLinkAccessPoint) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



### AsNetworkingV1DnsRecordSpecConfigOneOf

`func (s *NetworkingV1PrivateLinkAccessPoint) AsNetworkingV1DnsRecordSpecConfigOneOf() NetworkingV1DnsRecordSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1PrivateLinkAccessPoint in NetworkingV1DnsRecordSpecConfigOneOf

### AsNetworkingV1DnsRecordSpecUpdateConfigOneOf

`func (s *NetworkingV1PrivateLinkAccessPoint) AsNetworkingV1DnsRecordSpecUpdateConfigOneOf() NetworkingV1DnsRecordSpecUpdateConfigOneOf`

Convenience method to wrap this instance of NetworkingV1PrivateLinkAccessPoint in NetworkingV1DnsRecordSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



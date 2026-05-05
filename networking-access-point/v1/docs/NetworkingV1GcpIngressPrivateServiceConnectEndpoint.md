# NetworkingV1GcpIngressPrivateServiceConnectEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | GcpIngressPrivateServiceConnectEndpoint kind. | 
**PrivateServiceConnectConnectionId** | **string** | The ID of the Private Service Connect connection. | 

## Methods

### NewNetworkingV1GcpIngressPrivateServiceConnectEndpoint

`func NewNetworkingV1GcpIngressPrivateServiceConnectEndpoint(kind string, privateServiceConnectConnectionId string, ) *NetworkingV1GcpIngressPrivateServiceConnectEndpoint`

NewNetworkingV1GcpIngressPrivateServiceConnectEndpoint instantiates a new NetworkingV1GcpIngressPrivateServiceConnectEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpIngressPrivateServiceConnectEndpointWithDefaults

`func NewNetworkingV1GcpIngressPrivateServiceConnectEndpointWithDefaults() *NetworkingV1GcpIngressPrivateServiceConnectEndpoint`

NewNetworkingV1GcpIngressPrivateServiceConnectEndpointWithDefaults instantiates a new NetworkingV1GcpIngressPrivateServiceConnectEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpoint) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateServiceConnectConnectionId

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpoint) GetPrivateServiceConnectConnectionId() string`

GetPrivateServiceConnectConnectionId returns the PrivateServiceConnectConnectionId field if non-nil, zero value otherwise.

### GetPrivateServiceConnectConnectionIdOk

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpoint) GetPrivateServiceConnectConnectionIdOk() (*string, bool)`

GetPrivateServiceConnectConnectionIdOk returns a tuple with the PrivateServiceConnectConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectConnectionId

`func (o *NetworkingV1GcpIngressPrivateServiceConnectEndpoint) SetPrivateServiceConnectConnectionId(v string)`

SetPrivateServiceConnectConnectionId sets PrivateServiceConnectConnectionId field to given value.



### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1GcpIngressPrivateServiceConnectEndpoint) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1GcpIngressPrivateServiceConnectEndpoint in NetworkingV1AccessPointSpecConfigOneOf

### AsNetworkingV1AccessPointSpecUpdateConfigOneOf

`func (s *NetworkingV1GcpIngressPrivateServiceConnectEndpoint) AsNetworkingV1AccessPointSpecUpdateConfigOneOf() NetworkingV1AccessPointSpecUpdateConfigOneOf`

Convenience method to wrap this instance of NetworkingV1GcpIngressPrivateServiceConnectEndpoint in NetworkingV1AccessPointSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



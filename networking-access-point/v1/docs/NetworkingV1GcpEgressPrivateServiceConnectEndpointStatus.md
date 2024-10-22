# NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | GcpEgressPrivateServiceConnectEndpointStatus kind. | 
**PrivateServiceConnectEndpointConnectionId** | **string** | Connection ID of the Private Service Connect Endpoint (if any) that is connected to the endpoint target. | [readonly] 
**PrivateServiceConnectEndpointName** | **string** | Name of the Private Service Connect Endpoint (if any) that is connected to the endpoint target. | [readonly] 
**PrivateServiceConnectEndpointIpAddress** | **string** | IP address of the Private Service Connect Endpoint (if any) that is connected to the endpoint target. | [readonly] 

## Methods

### NewNetworkingV1GcpEgressPrivateServiceConnectEndpointStatus

`func NewNetworkingV1GcpEgressPrivateServiceConnectEndpointStatus(kind string, privateServiceConnectEndpointConnectionId string, privateServiceConnectEndpointName string, privateServiceConnectEndpointIpAddress string, ) *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus`

NewNetworkingV1GcpEgressPrivateServiceConnectEndpointStatus instantiates a new NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpEgressPrivateServiceConnectEndpointStatusWithDefaults

`func NewNetworkingV1GcpEgressPrivateServiceConnectEndpointStatusWithDefaults() *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus`

NewNetworkingV1GcpEgressPrivateServiceConnectEndpointStatusWithDefaults instantiates a new NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateServiceConnectEndpointConnectionId

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectEndpointConnectionId() string`

GetPrivateServiceConnectEndpointConnectionId returns the PrivateServiceConnectEndpointConnectionId field if non-nil, zero value otherwise.

### GetPrivateServiceConnectEndpointConnectionIdOk

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectEndpointConnectionIdOk() (*string, bool)`

GetPrivateServiceConnectEndpointConnectionIdOk returns a tuple with the PrivateServiceConnectEndpointConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectEndpointConnectionId

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) SetPrivateServiceConnectEndpointConnectionId(v string)`

SetPrivateServiceConnectEndpointConnectionId sets PrivateServiceConnectEndpointConnectionId field to given value.


### GetPrivateServiceConnectEndpointName

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectEndpointName() string`

GetPrivateServiceConnectEndpointName returns the PrivateServiceConnectEndpointName field if non-nil, zero value otherwise.

### GetPrivateServiceConnectEndpointNameOk

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectEndpointNameOk() (*string, bool)`

GetPrivateServiceConnectEndpointNameOk returns a tuple with the PrivateServiceConnectEndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectEndpointName

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) SetPrivateServiceConnectEndpointName(v string)`

SetPrivateServiceConnectEndpointName sets PrivateServiceConnectEndpointName field to given value.


### GetPrivateServiceConnectEndpointIpAddress

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectEndpointIpAddress() string`

GetPrivateServiceConnectEndpointIpAddress returns the PrivateServiceConnectEndpointIpAddress field if non-nil, zero value otherwise.

### GetPrivateServiceConnectEndpointIpAddressOk

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) GetPrivateServiceConnectEndpointIpAddressOk() (*string, bool)`

GetPrivateServiceConnectEndpointIpAddressOk returns a tuple with the PrivateServiceConnectEndpointIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectEndpointIpAddress

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) SetPrivateServiceConnectEndpointIpAddress(v string)`

SetPrivateServiceConnectEndpointIpAddress sets PrivateServiceConnectEndpointIpAddress field to given value.



### AsNetworkingV1AccessPointStatusConfigOneOf

`func (s *NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus) AsNetworkingV1AccessPointStatusConfigOneOf() NetworkingV1AccessPointStatusConfigOneOf`

Convenience method to wrap this instance of NetworkingV1GcpEgressPrivateServiceConnectEndpointStatus in NetworkingV1AccessPointStatusConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetworkingV1GcpEgressPrivateServiceConnectEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | GcpEgressPrivateServiceConnectEndpoint kind. | 
**PrivateServiceConnectEndpointTarget** | **string** | URI of the service attachment for the published service that the Private Service Connect Endpoint connects to or \&quot;ALL_GOOGLE_APIS\&quot; for global Google APIs. | 
**TargetSystem** | Pointer to **string** | [Used by the Confluent Cloud Console] The target system or service that the PrivateLink Endpoint connects to (e.g. \&quot;GCS\&quot; or \&quot;SNOWFLAKE\&quot;). | [optional] 

## Methods

### NewNetworkingV1GcpEgressPrivateServiceConnectEndpoint

`func NewNetworkingV1GcpEgressPrivateServiceConnectEndpoint(kind string, privateServiceConnectEndpointTarget string, ) *NetworkingV1GcpEgressPrivateServiceConnectEndpoint`

NewNetworkingV1GcpEgressPrivateServiceConnectEndpoint instantiates a new NetworkingV1GcpEgressPrivateServiceConnectEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpEgressPrivateServiceConnectEndpointWithDefaults

`func NewNetworkingV1GcpEgressPrivateServiceConnectEndpointWithDefaults() *NetworkingV1GcpEgressPrivateServiceConnectEndpoint`

NewNetworkingV1GcpEgressPrivateServiceConnectEndpointWithDefaults instantiates a new NetworkingV1GcpEgressPrivateServiceConnectEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPrivateServiceConnectEndpointTarget

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) GetPrivateServiceConnectEndpointTarget() string`

GetPrivateServiceConnectEndpointTarget returns the PrivateServiceConnectEndpointTarget field if non-nil, zero value otherwise.

### GetPrivateServiceConnectEndpointTargetOk

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) GetPrivateServiceConnectEndpointTargetOk() (*string, bool)`

GetPrivateServiceConnectEndpointTargetOk returns a tuple with the PrivateServiceConnectEndpointTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateServiceConnectEndpointTarget

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) SetPrivateServiceConnectEndpointTarget(v string)`

SetPrivateServiceConnectEndpointTarget sets PrivateServiceConnectEndpointTarget field to given value.


### GetTargetSystem

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) GetTargetSystem() string`

GetTargetSystem returns the TargetSystem field if non-nil, zero value otherwise.

### GetTargetSystemOk

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) GetTargetSystemOk() (*string, bool)`

GetTargetSystemOk returns a tuple with the TargetSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystem

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) SetTargetSystem(v string)`

SetTargetSystem sets TargetSystem field to given value.

### HasTargetSystem

`func (o *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) HasTargetSystem() bool`

HasTargetSystem returns a boolean if a field has been set.


### AsNetworkingV1AccessPointSpecConfigOneOf

`func (s *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) AsNetworkingV1AccessPointSpecConfigOneOf() NetworkingV1AccessPointSpecConfigOneOf`

Convenience method to wrap this instance of NetworkingV1GcpEgressPrivateServiceConnectEndpoint in NetworkingV1AccessPointSpecConfigOneOf

### AsNetworkingV1AccessPointSpecUpdateConfigOneOf

`func (s *NetworkingV1GcpEgressPrivateServiceConnectEndpoint) AsNetworkingV1AccessPointSpecUpdateConfigOneOf() NetworkingV1AccessPointSpecUpdateConfigOneOf`

Convenience method to wrap this instance of NetworkingV1GcpEgressPrivateServiceConnectEndpoint in NetworkingV1AccessPointSpecUpdateConfigOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



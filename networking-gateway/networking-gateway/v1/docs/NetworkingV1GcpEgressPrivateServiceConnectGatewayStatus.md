# NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | GCP Private Service Connect Gateway Status kind type. | 
**Project** | Pointer to **string** | The GCP project used by the GCP Private Service Connect Gateway. | [optional] [readonly] 

## Methods

### NewNetworkingV1GcpEgressPrivateServiceConnectGatewayStatus

`func NewNetworkingV1GcpEgressPrivateServiceConnectGatewayStatus(kind string, ) *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus`

NewNetworkingV1GcpEgressPrivateServiceConnectGatewayStatus instantiates a new NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpEgressPrivateServiceConnectGatewayStatusWithDefaults

`func NewNetworkingV1GcpEgressPrivateServiceConnectGatewayStatusWithDefaults() *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus`

NewNetworkingV1GcpEgressPrivateServiceConnectGatewayStatusWithDefaults instantiates a new NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus) SetKind(v string)`

SetKind sets Kind field to given value.


### GetProject

`func (o *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus) HasProject() bool`

HasProject returns a boolean if a field has been set.


### AsNetworkingV1GatewayStatusCloudGatewayOneOf

`func (s *NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus) AsNetworkingV1GatewayStatusCloudGatewayOneOf() NetworkingV1GatewayStatusCloudGatewayOneOf`

Convenience method to wrap this instance of NetworkingV1GcpEgressPrivateServiceConnectGatewayStatus in NetworkingV1GatewayStatusCloudGatewayOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



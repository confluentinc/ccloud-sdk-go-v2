# NetworkingV1GcpNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Project** | **string** | The GCP project. | [readonly] 
**VpcNetwork** | **string** | The GCP VPC network name. | [readonly] 

## Methods

### NewNetworkingV1GcpNetwork

`func NewNetworkingV1GcpNetwork(kind string, project string, vpcNetwork string, ) *NetworkingV1GcpNetwork`

NewNetworkingV1GcpNetwork instantiates a new NetworkingV1GcpNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpNetworkWithDefaults

`func NewNetworkingV1GcpNetworkWithDefaults() *NetworkingV1GcpNetwork`

NewNetworkingV1GcpNetworkWithDefaults instantiates a new NetworkingV1GcpNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpNetwork) SetKind(v string)`

SetKind sets Kind field to given value.


### GetProject

`func (o *NetworkingV1GcpNetwork) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingV1GcpNetwork) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingV1GcpNetwork) SetProject(v string)`

SetProject sets Project field to given value.


### GetVpcNetwork

`func (o *NetworkingV1GcpNetwork) GetVpcNetwork() string`

GetVpcNetwork returns the VpcNetwork field if non-nil, zero value otherwise.

### GetVpcNetworkOk

`func (o *NetworkingV1GcpNetwork) GetVpcNetworkOk() (*string, bool)`

GetVpcNetworkOk returns a tuple with the VpcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNetwork

`func (o *NetworkingV1GcpNetwork) SetVpcNetwork(v string)`

SetVpcNetwork sets VpcNetwork field to given value.



### AsNetworkingV1NetworkStatusCloudOneOf

`func (s *NetworkingV1GcpNetwork) AsNetworkingV1NetworkStatusCloudOneOf() NetworkingV1NetworkStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingV1GcpNetwork in NetworkingV1NetworkStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



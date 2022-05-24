# NetworkingAdminV1GcpNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Network kind type. | 
**Project** | **string** | The GCP project. | [readonly] 
**VpcNetwork** | **string** | The GCP VPC network name. | [readonly] 

## Methods

### NewNetworkingAdminV1GcpNetwork

`func NewNetworkingAdminV1GcpNetwork(kind string, project string, vpcNetwork string, ) *NetworkingAdminV1GcpNetwork`

NewNetworkingAdminV1GcpNetwork instantiates a new NetworkingAdminV1GcpNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1GcpNetworkWithDefaults

`func NewNetworkingAdminV1GcpNetworkWithDefaults() *NetworkingAdminV1GcpNetwork`

NewNetworkingAdminV1GcpNetworkWithDefaults instantiates a new NetworkingAdminV1GcpNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingAdminV1GcpNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingAdminV1GcpNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingAdminV1GcpNetwork) SetKind(v string)`

SetKind sets Kind field to given value.


### GetProject

`func (o *NetworkingAdminV1GcpNetwork) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingAdminV1GcpNetwork) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingAdminV1GcpNetwork) SetProject(v string)`

SetProject sets Project field to given value.


### GetVpcNetwork

`func (o *NetworkingAdminV1GcpNetwork) GetVpcNetwork() string`

GetVpcNetwork returns the VpcNetwork field if non-nil, zero value otherwise.

### GetVpcNetworkOk

`func (o *NetworkingAdminV1GcpNetwork) GetVpcNetworkOk() (*string, bool)`

GetVpcNetworkOk returns a tuple with the VpcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNetwork

`func (o *NetworkingAdminV1GcpNetwork) SetVpcNetwork(v string)`

SetVpcNetwork sets VpcNetwork field to given value.



### AsNetworkingAdminV1NetworkStatusCloudOneOf

`func (s *NetworkingAdminV1GcpNetwork) AsNetworkingAdminV1NetworkStatusCloudOneOf() NetworkingAdminV1NetworkStatusCloudOneOf`

Convenience method to wrap this instance of NetworkingAdminV1GcpNetwork in NetworkingAdminV1NetworkStatusCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



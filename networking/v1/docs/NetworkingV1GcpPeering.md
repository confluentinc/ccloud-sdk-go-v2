# NetworkingV1GcpPeering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Peering kind type. | 
**Project** | **string** | The name of the GCP project. | 
**VpcNetwork** | **string** | The name of the GCP VPC network to peer with. | 
**ImportCustomRoutes** | Pointer to **bool** | Enable customer route import. | [optional] [default to false]

## Methods

### NewNetworkingV1GcpPeering

`func NewNetworkingV1GcpPeering(kind string, project string, vpcNetwork string, ) *NetworkingV1GcpPeering`

NewNetworkingV1GcpPeering instantiates a new NetworkingV1GcpPeering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GcpPeeringWithDefaults

`func NewNetworkingV1GcpPeeringWithDefaults() *NetworkingV1GcpPeering`

NewNetworkingV1GcpPeeringWithDefaults instantiates a new NetworkingV1GcpPeering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *NetworkingV1GcpPeering) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkingV1GcpPeering) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkingV1GcpPeering) SetKind(v string)`

SetKind sets Kind field to given value.


### GetProject

`func (o *NetworkingV1GcpPeering) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkingV1GcpPeering) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkingV1GcpPeering) SetProject(v string)`

SetProject sets Project field to given value.


### GetVpcNetwork

`func (o *NetworkingV1GcpPeering) GetVpcNetwork() string`

GetVpcNetwork returns the VpcNetwork field if non-nil, zero value otherwise.

### GetVpcNetworkOk

`func (o *NetworkingV1GcpPeering) GetVpcNetworkOk() (*string, bool)`

GetVpcNetworkOk returns a tuple with the VpcNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNetwork

`func (o *NetworkingV1GcpPeering) SetVpcNetwork(v string)`

SetVpcNetwork sets VpcNetwork field to given value.


### GetImportCustomRoutes

`func (o *NetworkingV1GcpPeering) GetImportCustomRoutes() bool`

GetImportCustomRoutes returns the ImportCustomRoutes field if non-nil, zero value otherwise.

### GetImportCustomRoutesOk

`func (o *NetworkingV1GcpPeering) GetImportCustomRoutesOk() (*bool, bool)`

GetImportCustomRoutesOk returns a tuple with the ImportCustomRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportCustomRoutes

`func (o *NetworkingV1GcpPeering) SetImportCustomRoutes(v bool)`

SetImportCustomRoutes sets ImportCustomRoutes field to given value.

### HasImportCustomRoutes

`func (o *NetworkingV1GcpPeering) HasImportCustomRoutes() bool`

HasImportCustomRoutes returns a boolean if a field has been set.


### AsNetworkingV1PeeringSpecCloudOneOf

`func (s *NetworkingV1GcpPeering) AsNetworkingV1PeeringSpecCloudOneOf() NetworkingV1PeeringSpecCloudOneOf`

Convenience method to wrap this instance of NetworkingV1GcpPeering in NetworkingV1PeeringSpecCloudOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



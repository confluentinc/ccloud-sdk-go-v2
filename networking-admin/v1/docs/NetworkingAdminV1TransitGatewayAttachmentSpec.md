# NetworkingAdminV1TransitGatewayAttachmentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the TGW attachment | [optional] 
**RamShareArn** | Pointer to **string** | The full AWS Resource Name (ARN) for the AWS Resource Access Manager (RAM) Share of the Transit Gateways that you want Confluent Cloud attached to | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**Network** | Pointer to [**ObjectReference**](ObjectReference.md) | The network to which this belongs. | [optional] 

## Methods

### NewNetworkingAdminV1TransitGatewayAttachmentSpec

`func NewNetworkingAdminV1TransitGatewayAttachmentSpec() *NetworkingAdminV1TransitGatewayAttachmentSpec`

NewNetworkingAdminV1TransitGatewayAttachmentSpec instantiates a new NetworkingAdminV1TransitGatewayAttachmentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1TransitGatewayAttachmentSpecWithDefaults

`func NewNetworkingAdminV1TransitGatewayAttachmentSpecWithDefaults() *NetworkingAdminV1TransitGatewayAttachmentSpec`

NewNetworkingAdminV1TransitGatewayAttachmentSpecWithDefaults instantiates a new NetworkingAdminV1TransitGatewayAttachmentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRamShareArn

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetRamShareArn() string`

GetRamShareArn returns the RamShareArn field if non-nil, zero value otherwise.

### GetRamShareArnOk

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetRamShareArnOk() (*string, bool)`

GetRamShareArnOk returns a tuple with the RamShareArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamShareArn

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) SetRamShareArn(v string)`

SetRamShareArn sets RamShareArn field to given value.

### HasRamShareArn

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) HasRamShareArn() bool`

HasRamShareArn returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetNetwork() ObjectReference`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetNetworkOk() (*ObjectReference, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) SetNetwork(v ObjectReference)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



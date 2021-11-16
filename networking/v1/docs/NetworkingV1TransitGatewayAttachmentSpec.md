# NetworkingV1TransitGatewayAttachmentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the TGW attachment | [optional] 
**RamShareArn** | Pointer to **string** | The full AWS Resource Name (ARN) for the AWS Resource Access Manager (RAM) Share of the Transit Gateways that you want Confluent Cloud attached to | [optional] 
**Network** | Pointer to **string** | The network to use for the Transit Gateway | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1TransitGatewayAttachmentSpec

`func NewNetworkingV1TransitGatewayAttachmentSpec() *NetworkingV1TransitGatewayAttachmentSpec`

NewNetworkingV1TransitGatewayAttachmentSpec instantiates a new NetworkingV1TransitGatewayAttachmentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1TransitGatewayAttachmentSpecWithDefaults

`func NewNetworkingV1TransitGatewayAttachmentSpecWithDefaults() *NetworkingV1TransitGatewayAttachmentSpec`

NewNetworkingV1TransitGatewayAttachmentSpecWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1TransitGatewayAttachmentSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1TransitGatewayAttachmentSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRamShareArn

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetRamShareArn() string`

GetRamShareArn returns the RamShareArn field if non-nil, zero value otherwise.

### GetRamShareArnOk

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetRamShareArnOk() (*string, bool)`

GetRamShareArnOk returns a tuple with the RamShareArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamShareArn

`func (o *NetworkingV1TransitGatewayAttachmentSpec) SetRamShareArn(v string)`

SetRamShareArn sets RamShareArn field to given value.

### HasRamShareArn

`func (o *NetworkingV1TransitGatewayAttachmentSpec) HasRamShareArn() bool`

HasRamShareArn returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingV1TransitGatewayAttachmentSpec) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingV1TransitGatewayAttachmentSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1TransitGatewayAttachmentSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1TransitGatewayAttachmentSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1TransitGatewayAttachmentSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



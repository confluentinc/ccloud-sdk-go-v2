# NetworkingAdminV1TransitGatewayAttachmentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the TGW attachment | [optional] 
**Cloud** | Pointer to [**NetworkingAdminV1TransitGatewayAttachmentSpecCloudOneOf**](NetworkingAdminV1TransitGatewayAttachmentSpecCloudOneOf.md) | The cloud-specific Transit Gateway details. | [optional] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 
**Network** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The network to which this belongs. | [optional] 

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

### GetCloud

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetCloud() NetworkingAdminV1TransitGatewayAttachmentSpecCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetCloudOk() (*NetworkingAdminV1TransitGatewayAttachmentSpecCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) SetCloud(v NetworkingAdminV1TransitGatewayAttachmentSpecCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNetwork

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetNetwork() EnvScopedObjectReference`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) GetNetworkOk() (*EnvScopedObjectReference, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) SetNetwork(v EnvScopedObjectReference)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkingAdminV1TransitGatewayAttachmentSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



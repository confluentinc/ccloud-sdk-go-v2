# NetworkingV1TransitGatewayAttachmentSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the TGW attachment | [optional] 
**Cloud** | Pointer to [**NetworkingV1TransitGatewayAttachmentSpecUpdateCloudOneOf**](NetworkingV1TransitGatewayAttachmentSpecUpdateCloudOneOf.md) | The cloud-specific Transit Gateway details. | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewNetworkingV1TransitGatewayAttachmentSpecUpdate

`func NewNetworkingV1TransitGatewayAttachmentSpecUpdate() *NetworkingV1TransitGatewayAttachmentSpecUpdate`

NewNetworkingV1TransitGatewayAttachmentSpecUpdate instantiates a new NetworkingV1TransitGatewayAttachmentSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1TransitGatewayAttachmentSpecUpdateWithDefaults

`func NewNetworkingV1TransitGatewayAttachmentSpecUpdateWithDefaults() *NetworkingV1TransitGatewayAttachmentSpecUpdate`

NewNetworkingV1TransitGatewayAttachmentSpecUpdateWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetCloud() NetworkingV1TransitGatewayAttachmentSpecUpdateCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetCloudOk() (*NetworkingV1TransitGatewayAttachmentSpecUpdateCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) SetCloud(v NetworkingV1TransitGatewayAttachmentSpecUpdateCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



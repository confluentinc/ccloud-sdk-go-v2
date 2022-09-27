# NetworkingV1TransitGatewayAttachmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the TGW attachment:   PROVISIONING: attachment provisioning is in progress;   PENDING_ACCEPT: attachment request is pending acceptance by the customer;   READY:  attachment is ready;   FAILED: attachment is in a failed state;   DEPROVISIONING: attachment deprovisioning is in progress;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if TGW attachment is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if TGW attachment is in a failed state | [optional] [readonly] 
**Cloud** | Pointer to [**NetworkingV1TransitGatewayAttachmentStatusCloudOneOf**](NetworkingV1TransitGatewayAttachmentStatusCloudOneOf.md) | The cloud-specific TGW attachment details. | [optional] [readonly] 

## Methods

### NewNetworkingV1TransitGatewayAttachmentStatus

`func NewNetworkingV1TransitGatewayAttachmentStatus(phase string, ) *NetworkingV1TransitGatewayAttachmentStatus`

NewNetworkingV1TransitGatewayAttachmentStatus instantiates a new NetworkingV1TransitGatewayAttachmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1TransitGatewayAttachmentStatusWithDefaults

`func NewNetworkingV1TransitGatewayAttachmentStatusWithDefaults() *NetworkingV1TransitGatewayAttachmentStatus`

NewNetworkingV1TransitGatewayAttachmentStatusWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1TransitGatewayAttachmentStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1TransitGatewayAttachmentStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1TransitGatewayAttachmentStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingV1TransitGatewayAttachmentStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1TransitGatewayAttachmentStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1TransitGatewayAttachmentStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1TransitGatewayAttachmentStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1TransitGatewayAttachmentStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1TransitGatewayAttachmentStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1TransitGatewayAttachmentStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1TransitGatewayAttachmentStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCloud

`func (o *NetworkingV1TransitGatewayAttachmentStatus) GetCloud() NetworkingV1TransitGatewayAttachmentStatusCloudOneOf`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *NetworkingV1TransitGatewayAttachmentStatus) GetCloudOk() (*NetworkingV1TransitGatewayAttachmentStatusCloudOneOf, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *NetworkingV1TransitGatewayAttachmentStatus) SetCloud(v NetworkingV1TransitGatewayAttachmentStatusCloudOneOf)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *NetworkingV1TransitGatewayAttachmentStatus) HasCloud() bool`

HasCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



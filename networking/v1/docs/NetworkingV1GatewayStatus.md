# NetworkingV1GatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the gateway:    PROVISIONING: gateway provisioning is in progress;    READY:  gateway is ready;    FAILED: gateway is in a failed state;    DEPROVISIONING: gateway deprovisioning is in progress;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if gateway is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if gateway is in a failed state | [optional] [readonly] 
**CloudGateway** | Pointer to [**NetworkingV1GatewayStatusCloudGatewayOneOf**](NetworkingV1GatewayStatusCloudGatewayOneOf.md) | Gateway type specific status | [optional] [readonly] 

## Methods

### NewNetworkingV1GatewayStatus

`func NewNetworkingV1GatewayStatus(phase string, ) *NetworkingV1GatewayStatus`

NewNetworkingV1GatewayStatus instantiates a new NetworkingV1GatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1GatewayStatusWithDefaults

`func NewNetworkingV1GatewayStatusWithDefaults() *NetworkingV1GatewayStatus`

NewNetworkingV1GatewayStatusWithDefaults instantiates a new NetworkingV1GatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1GatewayStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1GatewayStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1GatewayStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingV1GatewayStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1GatewayStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1GatewayStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1GatewayStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1GatewayStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1GatewayStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1GatewayStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1GatewayStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCloudGateway

`func (o *NetworkingV1GatewayStatus) GetCloudGateway() NetworkingV1GatewayStatusCloudGatewayOneOf`

GetCloudGateway returns the CloudGateway field if non-nil, zero value otherwise.

### GetCloudGatewayOk

`func (o *NetworkingV1GatewayStatus) GetCloudGatewayOk() (*NetworkingV1GatewayStatusCloudGatewayOneOf, bool)`

GetCloudGatewayOk returns a tuple with the CloudGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudGateway

`func (o *NetworkingV1GatewayStatus) SetCloudGateway(v NetworkingV1GatewayStatusCloudGatewayOneOf)`

SetCloudGateway sets CloudGateway field to given value.

### HasCloudGateway

`func (o *NetworkingV1GatewayStatus) HasCloudGateway() bool`

HasCloudGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



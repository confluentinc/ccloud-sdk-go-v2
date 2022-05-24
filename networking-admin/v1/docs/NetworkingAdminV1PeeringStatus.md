# NetworkingAdminV1PeeringStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the peering:   PROVISIONING: peering provisioning is in progress;   PENDING_ACCEPT: peering connection request is pending acceptance by the customer;   READY:  peering is ready;   FAILED: peering is in a failed state;   DEPROVISIONING: peering deprovisioning is in progress;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if peering is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if peering is in a failed state | [optional] [readonly] 

## Methods

### NewNetworkingAdminV1PeeringStatus

`func NewNetworkingAdminV1PeeringStatus(phase string, ) *NetworkingAdminV1PeeringStatus`

NewNetworkingAdminV1PeeringStatus instantiates a new NetworkingAdminV1PeeringStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingAdminV1PeeringStatusWithDefaults

`func NewNetworkingAdminV1PeeringStatusWithDefaults() *NetworkingAdminV1PeeringStatus`

NewNetworkingAdminV1PeeringStatusWithDefaults instantiates a new NetworkingAdminV1PeeringStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingAdminV1PeeringStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingAdminV1PeeringStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingAdminV1PeeringStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingAdminV1PeeringStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingAdminV1PeeringStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingAdminV1PeeringStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingAdminV1PeeringStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingAdminV1PeeringStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingAdminV1PeeringStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingAdminV1PeeringStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingAdminV1PeeringStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



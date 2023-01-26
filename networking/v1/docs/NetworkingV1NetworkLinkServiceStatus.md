# NetworkingV1NetworkLinkServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the network link service:  READY:  network link service is ready;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if network link service is in a failed state. May be used for programmatic error checking.  | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if network link service is in a failed state | [optional] [readonly] 

## Methods

### NewNetworkingV1NetworkLinkServiceStatus

`func NewNetworkingV1NetworkLinkServiceStatus(phase string, ) *NetworkingV1NetworkLinkServiceStatus`

NewNetworkingV1NetworkLinkServiceStatus instantiates a new NetworkingV1NetworkLinkServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1NetworkLinkServiceStatusWithDefaults

`func NewNetworkingV1NetworkLinkServiceStatusWithDefaults() *NetworkingV1NetworkLinkServiceStatus`

NewNetworkingV1NetworkLinkServiceStatusWithDefaults instantiates a new NetworkingV1NetworkLinkServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1NetworkLinkServiceStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1NetworkLinkServiceStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1NetworkLinkServiceStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingV1NetworkLinkServiceStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1NetworkLinkServiceStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1NetworkLinkServiceStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1NetworkLinkServiceStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1NetworkLinkServiceStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1NetworkLinkServiceStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1NetworkLinkServiceStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1NetworkLinkServiceStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



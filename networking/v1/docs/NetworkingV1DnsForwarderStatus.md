# NetworkingV1DnsForwarderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The lifecycle phase of the dns forwarder:    PROVISIONING: dns forwarder provisioning is in progress;    READY: dns forwarder is ready;    FAILED: dns forwarder is in a failed state;    DEPROVISIONING: dns forwarder deprovisioning is in progress;  | [readonly] 
**ErrorCode** | Pointer to **string** | Error code if dns forwarder is in a failed state. May be used for programmatic error checking. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Displayable error message if dns forwarder is in a failed state | [optional] [readonly] 

## Methods

### NewNetworkingV1DnsForwarderStatus

`func NewNetworkingV1DnsForwarderStatus(phase string, ) *NetworkingV1DnsForwarderStatus`

NewNetworkingV1DnsForwarderStatus instantiates a new NetworkingV1DnsForwarderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkingV1DnsForwarderStatusWithDefaults

`func NewNetworkingV1DnsForwarderStatusWithDefaults() *NetworkingV1DnsForwarderStatus`

NewNetworkingV1DnsForwarderStatusWithDefaults instantiates a new NetworkingV1DnsForwarderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *NetworkingV1DnsForwarderStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NetworkingV1DnsForwarderStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NetworkingV1DnsForwarderStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetErrorCode

`func (o *NetworkingV1DnsForwarderStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NetworkingV1DnsForwarderStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NetworkingV1DnsForwarderStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NetworkingV1DnsForwarderStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *NetworkingV1DnsForwarderStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NetworkingV1DnsForwarderStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NetworkingV1DnsForwarderStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NetworkingV1DnsForwarderStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



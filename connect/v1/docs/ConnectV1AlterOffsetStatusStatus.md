# ConnectV1AlterOffsetStatusStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | **string** | The phase of the alter offset operation.   PENDING: The offset alter operation is in progress.  APPLIED: The offset alter operation has been applied to the connector.  FAILED:  The offset alter operation has failed to be applied to the connector. | 
**Message** | Pointer to **string** | An info message from the alter offset operation. | [optional] 

## Methods

### NewConnectV1AlterOffsetStatusStatus

`func NewConnectV1AlterOffsetStatusStatus(phase string, ) *ConnectV1AlterOffsetStatusStatus`

NewConnectV1AlterOffsetStatusStatus instantiates a new ConnectV1AlterOffsetStatusStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1AlterOffsetStatusStatusWithDefaults

`func NewConnectV1AlterOffsetStatusStatusWithDefaults() *ConnectV1AlterOffsetStatusStatus`

NewConnectV1AlterOffsetStatusStatusWithDefaults instantiates a new ConnectV1AlterOffsetStatusStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *ConnectV1AlterOffsetStatusStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ConnectV1AlterOffsetStatusStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ConnectV1AlterOffsetStatusStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetMessage

`func (o *ConnectV1AlterOffsetStatusStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConnectV1AlterOffsetStatusStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConnectV1AlterOffsetStatusStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConnectV1AlterOffsetStatusStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



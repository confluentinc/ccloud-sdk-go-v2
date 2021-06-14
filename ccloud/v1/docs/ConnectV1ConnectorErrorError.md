# ConnectV1ConnectorErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | Error code for the type of error | [optional] 
**Message** | Pointer to **string** | Human readable error message | [optional] 

## Methods

### NewConnectV1ConnectorErrorError

`func NewConnectV1ConnectorErrorError() *ConnectV1ConnectorErrorError`

NewConnectV1ConnectorErrorError instantiates a new ConnectV1ConnectorErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorErrorErrorWithDefaults

`func NewConnectV1ConnectorErrorErrorWithDefaults() *ConnectV1ConnectorErrorError`

NewConnectV1ConnectorErrorErrorWithDefaults instantiates a new ConnectV1ConnectorErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ConnectV1ConnectorErrorError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ConnectV1ConnectorErrorError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ConnectV1ConnectorErrorError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ConnectV1ConnectorErrorError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ConnectV1ConnectorErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConnectV1ConnectorErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConnectV1ConnectorErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConnectV1ConnectorErrorError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



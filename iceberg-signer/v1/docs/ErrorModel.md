# ErrorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Human-readable error message | 
**Type** | **string** | Internal type definition of the error | 
**Code** | **int32** | HTTP response code | 
**Stack** | Pointer to **[]string** |  | [optional] 

## Methods

### NewErrorModel

`func NewErrorModel(message string, type_ string, code int32, ) *ErrorModel`

NewErrorModel instantiates a new ErrorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorModelWithDefaults

`func NewErrorModelWithDefaults() *ErrorModel`

NewErrorModelWithDefaults instantiates a new ErrorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *ErrorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorModel) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *ErrorModel) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorModel) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorModel) SetCode(v int32)`

SetCode sets Code field to given value.


### GetStack

`func (o *ErrorModel) GetStack() []string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *ErrorModel) GetStackOk() (*[]string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *ErrorModel) SetStack(v []string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *ErrorModel) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



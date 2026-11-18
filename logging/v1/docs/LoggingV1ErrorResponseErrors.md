# LoggingV1ErrorResponseErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for this occurrence of the problem. Not returned by the service today. | [optional] 
**Status** | **string** | HTTP status as a string. | 
**Code** | Pointer to **string** | An application-specific error code. Not returned by the service today. | [optional] 
**Detail** | **string** | A human-readable explanation specific to this occurrence of the problem. | 

## Methods

### NewLoggingV1ErrorResponseErrors

`func NewLoggingV1ErrorResponseErrors(status string, detail string, ) *LoggingV1ErrorResponseErrors`

NewLoggingV1ErrorResponseErrors instantiates a new LoggingV1ErrorResponseErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1ErrorResponseErrorsWithDefaults

`func NewLoggingV1ErrorResponseErrorsWithDefaults() *LoggingV1ErrorResponseErrors`

NewLoggingV1ErrorResponseErrorsWithDefaults instantiates a new LoggingV1ErrorResponseErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoggingV1ErrorResponseErrors) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoggingV1ErrorResponseErrors) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoggingV1ErrorResponseErrors) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoggingV1ErrorResponseErrors) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *LoggingV1ErrorResponseErrors) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoggingV1ErrorResponseErrors) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoggingV1ErrorResponseErrors) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCode

`func (o *LoggingV1ErrorResponseErrors) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LoggingV1ErrorResponseErrors) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LoggingV1ErrorResponseErrors) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LoggingV1ErrorResponseErrors) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *LoggingV1ErrorResponseErrors) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *LoggingV1ErrorResponseErrors) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *LoggingV1ErrorResponseErrors) SetDetail(v string)`

SetDetail sets Detail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



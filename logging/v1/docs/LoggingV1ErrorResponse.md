# LoggingV1ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]LoggingV1ErrorResponseErrors**](LoggingV1ErrorResponseErrors.md) | One or more problems encountered while processing the request. Bean-validation failures produce one entry per violated constraint. | 

## Methods

### NewLoggingV1ErrorResponse

`func NewLoggingV1ErrorResponse(errors []LoggingV1ErrorResponseErrors, ) *LoggingV1ErrorResponse`

NewLoggingV1ErrorResponse instantiates a new LoggingV1ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1ErrorResponseWithDefaults

`func NewLoggingV1ErrorResponseWithDefaults() *LoggingV1ErrorResponse`

NewLoggingV1ErrorResponseWithDefaults instantiates a new LoggingV1ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *LoggingV1ErrorResponse) GetErrors() []LoggingV1ErrorResponseErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LoggingV1ErrorResponse) GetErrorsOk() (*[]LoggingV1ErrorResponseErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LoggingV1ErrorResponse) SetErrors(v []LoggingV1ErrorResponseErrors)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



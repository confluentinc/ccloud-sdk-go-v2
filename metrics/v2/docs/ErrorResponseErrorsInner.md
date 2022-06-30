# ErrorResponseErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for this particular occurrence of the problem. | [optional] 
**Status** | Pointer to **string** | The HTTP status code applicable to this problem, expressed as a string value. | [optional] 
**Code** | Pointer to **string** | An application-specific error code, expressed as a string value. | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewErrorResponseErrorsInner

`func NewErrorResponseErrorsInner() *ErrorResponseErrorsInner`

NewErrorResponseErrorsInner instantiates a new ErrorResponseErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseErrorsInnerWithDefaults

`func NewErrorResponseErrorsInnerWithDefaults() *ErrorResponseErrorsInner`

NewErrorResponseErrorsInnerWithDefaults instantiates a new ErrorResponseErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorResponseErrorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorResponseErrorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorResponseErrorsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorResponseErrorsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ErrorResponseErrorsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponseErrorsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponseErrorsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorResponseErrorsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *ErrorResponseErrorsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseErrorsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseErrorsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorResponseErrorsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetail

`func (o *ErrorResponseErrorsInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorResponseErrorsInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorResponseErrorsInner) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorResponseErrorsInner) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetMeta

`func (o *ErrorResponseErrorsInner) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorResponseErrorsInner) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorResponseErrorsInner) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorResponseErrorsInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *ErrorResponseErrorsInner) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ErrorResponseErrorsInner) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ErrorResponseErrorsInner) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ErrorResponseErrorsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



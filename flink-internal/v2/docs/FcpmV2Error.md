# FcpmV2Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for this particular occurrence of the problem. | [optional] 
**Status** | Pointer to **string** | The HTTP status code applicable to this problem, expressed as a string value. | [optional] 
**Code** | Pointer to **string** | An application-specific error code, expressed as a string value. | [optional] 
**Title** | Pointer to **string** | A short, human-readable summary of the problem. It **SHOULD NOT** change from occurrence to occurrence of the problem, except for purposes of localization. | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Source** | Pointer to [**FcpmV2ErrorSource**](FcpmV2ErrorSource.md) |  | [optional] 

## Methods

### NewFcpmV2Error

`func NewFcpmV2Error() *FcpmV2Error`

NewFcpmV2Error instantiates a new FcpmV2Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2ErrorWithDefaults

`func NewFcpmV2ErrorWithDefaults() *FcpmV2Error`

NewFcpmV2ErrorWithDefaults instantiates a new FcpmV2Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FcpmV2Error) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FcpmV2Error) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FcpmV2Error) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FcpmV2Error) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *FcpmV2Error) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FcpmV2Error) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FcpmV2Error) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FcpmV2Error) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *FcpmV2Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FcpmV2Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FcpmV2Error) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FcpmV2Error) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTitle

`func (o *FcpmV2Error) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FcpmV2Error) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FcpmV2Error) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FcpmV2Error) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *FcpmV2Error) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *FcpmV2Error) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *FcpmV2Error) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *FcpmV2Error) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetSource

`func (o *FcpmV2Error) GetSource() FcpmV2ErrorSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FcpmV2Error) GetSourceOk() (*FcpmV2ErrorSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FcpmV2Error) SetSource(v FcpmV2ErrorSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *FcpmV2Error) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



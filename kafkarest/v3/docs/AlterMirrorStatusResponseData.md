# AlterMirrorStatusResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**MirrorTopicName** | **string** |  | 
**ErrorMessage** | **NullableString** |  | 
**ErrorCode** | **NullableInt32** |  | 
**MirrorLags** | [**MirrorLags**](MirrorLags.md) |  | 

## Methods

### NewAlterMirrorStatusResponseData

`func NewAlterMirrorStatusResponseData(kind string, metadata ResourceMetadata, mirrorTopicName string, errorMessage NullableString, errorCode NullableInt32, mirrorLags MirrorLags, ) *AlterMirrorStatusResponseData`

NewAlterMirrorStatusResponseData instantiates a new AlterMirrorStatusResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlterMirrorStatusResponseDataWithDefaults

`func NewAlterMirrorStatusResponseDataWithDefaults() *AlterMirrorStatusResponseData`

NewAlterMirrorStatusResponseDataWithDefaults instantiates a new AlterMirrorStatusResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AlterMirrorStatusResponseData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AlterMirrorStatusResponseData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AlterMirrorStatusResponseData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *AlterMirrorStatusResponseData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlterMirrorStatusResponseData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlterMirrorStatusResponseData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetMirrorTopicName

`func (o *AlterMirrorStatusResponseData) GetMirrorTopicName() string`

GetMirrorTopicName returns the MirrorTopicName field if non-nil, zero value otherwise.

### GetMirrorTopicNameOk

`func (o *AlterMirrorStatusResponseData) GetMirrorTopicNameOk() (*string, bool)`

GetMirrorTopicNameOk returns a tuple with the MirrorTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorTopicName

`func (o *AlterMirrorStatusResponseData) SetMirrorTopicName(v string)`

SetMirrorTopicName sets MirrorTopicName field to given value.


### GetErrorMessage

`func (o *AlterMirrorStatusResponseData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AlterMirrorStatusResponseData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AlterMirrorStatusResponseData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *AlterMirrorStatusResponseData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AlterMirrorStatusResponseData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorCode

`func (o *AlterMirrorStatusResponseData) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AlterMirrorStatusResponseData) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AlterMirrorStatusResponseData) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.


### SetErrorCodeNil

`func (o *AlterMirrorStatusResponseData) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AlterMirrorStatusResponseData) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetMirrorLags

`func (o *AlterMirrorStatusResponseData) GetMirrorLags() MirrorLags`

GetMirrorLags returns the MirrorLags field if non-nil, zero value otherwise.

### GetMirrorLagsOk

`func (o *AlterMirrorStatusResponseData) GetMirrorLagsOk() (*MirrorLags, bool)`

GetMirrorLagsOk returns a tuple with the MirrorLags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorLags

`func (o *AlterMirrorStatusResponseData) SetMirrorLags(v MirrorLags)`

SetMirrorLags sets MirrorLags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



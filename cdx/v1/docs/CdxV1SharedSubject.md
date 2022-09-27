# CdxV1SharedSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The shared resource kind | 
**Subject** | **string** | The subject name | 

## Methods

### NewCdxV1SharedSubject

`func NewCdxV1SharedSubject(kind string, subject string, ) *CdxV1SharedSubject`

NewCdxV1SharedSubject instantiates a new CdxV1SharedSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1SharedSubjectWithDefaults

`func NewCdxV1SharedSubjectWithDefaults() *CdxV1SharedSubject`

NewCdxV1SharedSubjectWithDefaults instantiates a new CdxV1SharedSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CdxV1SharedSubject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1SharedSubject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1SharedSubject) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSubject

`func (o *CdxV1SharedSubject) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CdxV1SharedSubject) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CdxV1SharedSubject) SetSubject(v string)`

SetSubject sets Subject field to given value.



### AsCdxV1RedeemTokenResponseResourcesOneOf

`func (s *CdxV1SharedSubject) AsCdxV1RedeemTokenResponseResourcesOneOf() CdxV1RedeemTokenResponseResourcesOneOf`

Convenience method to wrap this instance of CdxV1SharedSubject in CdxV1RedeemTokenResponseResourcesOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



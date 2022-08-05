# CdxV1EmailConsumerRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The resource kind | 
**Email** | **string** | Email based matching for the consumers | 

## Methods

### NewCdxV1EmailConsumerRestriction

`func NewCdxV1EmailConsumerRestriction(kind string, email string, ) *CdxV1EmailConsumerRestriction`

NewCdxV1EmailConsumerRestriction instantiates a new CdxV1EmailConsumerRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1EmailConsumerRestrictionWithDefaults

`func NewCdxV1EmailConsumerRestrictionWithDefaults() *CdxV1EmailConsumerRestriction`

NewCdxV1EmailConsumerRestrictionWithDefaults instantiates a new CdxV1EmailConsumerRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CdxV1EmailConsumerRestriction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1EmailConsumerRestriction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1EmailConsumerRestriction) SetKind(v string)`

SetKind sets Kind field to given value.


### GetEmail

`func (o *CdxV1EmailConsumerRestriction) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CdxV1EmailConsumerRestriction) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CdxV1EmailConsumerRestriction) SetEmail(v string)`

SetEmail sets Email field to given value.



### AsCdxV1CreateShareRequestConsumerRestrictionOneOf

`func (s *CdxV1EmailConsumerRestriction) AsCdxV1CreateShareRequestConsumerRestrictionOneOf() CdxV1CreateShareRequestConsumerRestrictionOneOf`

Convenience method to wrap this instance of CdxV1EmailConsumerRestriction in CdxV1CreateShareRequestConsumerRestrictionOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



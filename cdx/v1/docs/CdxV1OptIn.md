# CdxV1OptIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**StreamShareEnabled** | Pointer to **bool** | Enable stream sharing for the organization | [optional] 

## Methods

### NewCdxV1OptIn

`func NewCdxV1OptIn() *CdxV1OptIn`

NewCdxV1OptIn instantiates a new CdxV1OptIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1OptInWithDefaults

`func NewCdxV1OptInWithDefaults() *CdxV1OptIn`

NewCdxV1OptInWithDefaults instantiates a new CdxV1OptIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1OptIn) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1OptIn) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1OptIn) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1OptIn) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1OptIn) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1OptIn) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1OptIn) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1OptIn) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetStreamShareEnabled

`func (o *CdxV1OptIn) GetStreamShareEnabled() bool`

GetStreamShareEnabled returns the StreamShareEnabled field if non-nil, zero value otherwise.

### GetStreamShareEnabledOk

`func (o *CdxV1OptIn) GetStreamShareEnabledOk() (*bool, bool)`

GetStreamShareEnabledOk returns a tuple with the StreamShareEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamShareEnabled

`func (o *CdxV1OptIn) SetStreamShareEnabled(v bool)`

SetStreamShareEnabled sets StreamShareEnabled field to given value.

### HasStreamShareEnabled

`func (o *CdxV1OptIn) HasStreamShareEnabled() bool`

HasStreamShareEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



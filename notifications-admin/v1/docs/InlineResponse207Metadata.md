# InlineResponse207Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **int32** | The number of notifications sent successfully. | [optional] 
**Failure** | Pointer to **int32** | The number of notifications failed to send. | [optional] 
**Total** | Pointer to **int32** | Total number of notifications attempted to send. | [optional] 

## Methods

### NewInlineResponse207Metadata

`func NewInlineResponse207Metadata() *InlineResponse207Metadata`

NewInlineResponse207Metadata instantiates a new InlineResponse207Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse207MetadataWithDefaults

`func NewInlineResponse207MetadataWithDefaults() *InlineResponse207Metadata`

NewInlineResponse207MetadataWithDefaults instantiates a new InlineResponse207Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *InlineResponse207Metadata) GetSuccess() int32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InlineResponse207Metadata) GetSuccessOk() (*int32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InlineResponse207Metadata) SetSuccess(v int32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *InlineResponse207Metadata) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailure

`func (o *InlineResponse207Metadata) GetFailure() int32`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *InlineResponse207Metadata) GetFailureOk() (*int32, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *InlineResponse207Metadata) SetFailure(v int32)`

SetFailure sets Failure field to given value.

### HasFailure

`func (o *InlineResponse207Metadata) HasFailure() bool`

HasFailure returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse207Metadata) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse207Metadata) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse207Metadata) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse207Metadata) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



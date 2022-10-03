# CompatibilityCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCompatible** | Pointer to **bool** | Whether the compared schemas are compatible | [optional] 
**Messages** | Pointer to **[]string** | Error messages | [optional] 

## Methods

### NewCompatibilityCheckResponse

`func NewCompatibilityCheckResponse() *CompatibilityCheckResponse`

NewCompatibilityCheckResponse instantiates a new CompatibilityCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompatibilityCheckResponseWithDefaults

`func NewCompatibilityCheckResponseWithDefaults() *CompatibilityCheckResponse`

NewCompatibilityCheckResponseWithDefaults instantiates a new CompatibilityCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCompatible

`func (o *CompatibilityCheckResponse) GetIsCompatible() bool`

GetIsCompatible returns the IsCompatible field if non-nil, zero value otherwise.

### GetIsCompatibleOk

`func (o *CompatibilityCheckResponse) GetIsCompatibleOk() (*bool, bool)`

GetIsCompatibleOk returns a tuple with the IsCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompatible

`func (o *CompatibilityCheckResponse) SetIsCompatible(v bool)`

SetIsCompatible sets IsCompatible field to given value.

### HasIsCompatible

`func (o *CompatibilityCheckResponse) HasIsCompatible() bool`

HasIsCompatible returns a boolean if a field has been set.

### GetMessages

`func (o *CompatibilityCheckResponse) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *CompatibilityCheckResponse) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *CompatibilityCheckResponse) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *CompatibilityCheckResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# JsonPatchRequestRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | A JSON Pointer path. | 
**Op** | **string** | The operation to perform. | 

## Methods

### NewJsonPatchRequestRemove

`func NewJsonPatchRequestRemove(path string, op string, ) *JsonPatchRequestRemove`

NewJsonPatchRequestRemove instantiates a new JsonPatchRequestRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPatchRequestRemoveWithDefaults

`func NewJsonPatchRequestRemoveWithDefaults() *JsonPatchRequestRemove`

NewJsonPatchRequestRemoveWithDefaults instantiates a new JsonPatchRequestRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *JsonPatchRequestRemove) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JsonPatchRequestRemove) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JsonPatchRequestRemove) SetPath(v string)`

SetPath sets Path field to given value.


### GetOp

`func (o *JsonPatchRequestRemove) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JsonPatchRequestRemove) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JsonPatchRequestRemove) SetOp(v string)`

SetOp sets Op field to given value.



### AsPatchRequestOneOf

`func (s *JsonPatchRequestRemove) AsPatchRequestOneOf() PatchRequestOneOf`

Convenience method to wrap this instance of JsonPatchRequestRemove in PatchRequestOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



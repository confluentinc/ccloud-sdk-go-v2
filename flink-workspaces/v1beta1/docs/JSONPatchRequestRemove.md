# JSONPatchRequestRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | A JSON Pointer path. | 
**Op** | **string** | The operation to perform. | 

## Methods

### NewJSONPatchRequestRemove

`func NewJSONPatchRequestRemove(path string, op string, ) *JSONPatchRequestRemove`

NewJSONPatchRequestRemove instantiates a new JSONPatchRequestRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONPatchRequestRemoveWithDefaults

`func NewJSONPatchRequestRemoveWithDefaults() *JSONPatchRequestRemove`

NewJSONPatchRequestRemoveWithDefaults instantiates a new JSONPatchRequestRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *JSONPatchRequestRemove) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JSONPatchRequestRemove) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JSONPatchRequestRemove) SetPath(v string)`

SetPath sets Path field to given value.


### GetOp

`func (o *JSONPatchRequestRemove) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JSONPatchRequestRemove) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JSONPatchRequestRemove) SetOp(v string)`

SetOp sets Op field to given value.



### AsPatchRequestOneOf

`func (s *JSONPatchRequestRemove) AsPatchRequestOneOf() PatchRequestOneOf`

Convenience method to wrap this instance of JSONPatchRequestRemove in PatchRequestOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



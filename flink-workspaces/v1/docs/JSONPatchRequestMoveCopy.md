# JSONPatchRequestMoveCopy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | A JSON Pointer path. | 
**Op** | **string** | The operation to perform. | 
**From** | **string** | A JSON Pointer path. | 

## Methods

### NewJSONPatchRequestMoveCopy

`func NewJSONPatchRequestMoveCopy(path string, op string, from string, ) *JSONPatchRequestMoveCopy`

NewJSONPatchRequestMoveCopy instantiates a new JSONPatchRequestMoveCopy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSONPatchRequestMoveCopyWithDefaults

`func NewJSONPatchRequestMoveCopyWithDefaults() *JSONPatchRequestMoveCopy`

NewJSONPatchRequestMoveCopyWithDefaults instantiates a new JSONPatchRequestMoveCopy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *JSONPatchRequestMoveCopy) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JSONPatchRequestMoveCopy) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JSONPatchRequestMoveCopy) SetPath(v string)`

SetPath sets Path field to given value.


### GetOp

`func (o *JSONPatchRequestMoveCopy) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JSONPatchRequestMoveCopy) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JSONPatchRequestMoveCopy) SetOp(v string)`

SetOp sets Op field to given value.


### GetFrom

`func (o *JSONPatchRequestMoveCopy) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *JSONPatchRequestMoveCopy) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *JSONPatchRequestMoveCopy) SetFrom(v string)`

SetFrom sets From field to given value.



### AsPatchRequestOneOf

`func (s *JSONPatchRequestMoveCopy) AsPatchRequestOneOf() PatchRequestOneOf`

Convenience method to wrap this instance of JSONPatchRequestMoveCopy in PatchRequestOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



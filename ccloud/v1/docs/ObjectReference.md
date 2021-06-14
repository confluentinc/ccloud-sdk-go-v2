# ObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Related** | Pointer to **string** | URL for accessing or modifying the referred object | [optional] 

## Methods

### NewObjectReference

`func NewObjectReference() *ObjectReference`

NewObjectReference instantiates a new ObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectReferenceWithDefaults

`func NewObjectReferenceWithDefaults() *ObjectReference`

NewObjectReferenceWithDefaults instantiates a new ObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelated

`func (o *ObjectReference) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *ObjectReference) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *ObjectReference) SetRelated(v string)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *ObjectReference) HasRelated() bool`

HasRelated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



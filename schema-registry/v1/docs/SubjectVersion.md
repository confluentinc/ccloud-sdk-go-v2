# SubjectVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Name of the subject | [optional] 
**Version** | Pointer to **int32** | Version number | [optional] 

## Methods

### NewSubjectVersion

`func NewSubjectVersion() *SubjectVersion`

NewSubjectVersion instantiates a new SubjectVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectVersionWithDefaults

`func NewSubjectVersionWithDefaults() *SubjectVersion`

NewSubjectVersionWithDefaults instantiates a new SubjectVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *SubjectVersion) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SubjectVersion) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SubjectVersion) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SubjectVersion) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVersion

`func (o *SubjectVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SubjectVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SubjectVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SubjectVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



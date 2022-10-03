# SchemaReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Reference name | [optional] 
**Subject** | Pointer to **string** | Name of the referenced subject | [optional] 
**Version** | Pointer to **int32** | Version number of the referenced subject | [optional] 

## Methods

### NewSchemaReference

`func NewSchemaReference() *SchemaReference`

NewSchemaReference instantiates a new SchemaReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaReferenceWithDefaults

`func NewSchemaReferenceWithDefaults() *SchemaReference`

NewSchemaReferenceWithDefaults instantiates a new SchemaReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SchemaReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemaReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubject

`func (o *SchemaReference) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SchemaReference) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SchemaReference) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SchemaReference) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVersion

`func (o *SchemaReference) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaReference) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaReference) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SchemaReference) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



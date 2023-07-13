# RegisterExporterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**References** | Pointer to [**[]ExporterReference**](ExporterReference.md) | References to other schemas | [optional] 

## Methods

### NewRegisterExporterRequest

`func NewRegisterExporterRequest() *RegisterExporterRequest`

NewRegisterExporterRequest instantiates a new RegisterExporterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterExporterRequestWithDefaults

`func NewRegisterExporterRequestWithDefaults() *RegisterExporterRequest`

NewRegisterExporterRequestWithDefaults instantiates a new RegisterExporterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferences

`func (o *RegisterExporterRequest) GetReferences() []ExporterReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *RegisterExporterRequest) GetReferencesOk() (*[]ExporterReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *RegisterExporterRequest) SetReferences(v []ExporterReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *RegisterExporterRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



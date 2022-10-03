# RegisterSchemaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** | Version number | [optional] 
**Id** | Pointer to **int32** | Globally unique identifier of the schema | [optional] 
**SchemaType** | Pointer to **string** | Schema type | [optional] 
**References** | Pointer to [**[]SchemaReference**](SchemaReference.md) | References to other schemas | [optional] 
**Schema** | Pointer to **string** | Schema definition string | [optional] 

## Methods

### NewRegisterSchemaRequest

`func NewRegisterSchemaRequest() *RegisterSchemaRequest`

NewRegisterSchemaRequest instantiates a new RegisterSchemaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterSchemaRequestWithDefaults

`func NewRegisterSchemaRequestWithDefaults() *RegisterSchemaRequest`

NewRegisterSchemaRequestWithDefaults instantiates a new RegisterSchemaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *RegisterSchemaRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RegisterSchemaRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RegisterSchemaRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RegisterSchemaRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *RegisterSchemaRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterSchemaRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterSchemaRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegisterSchemaRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchemaType

`func (o *RegisterSchemaRequest) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *RegisterSchemaRequest) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *RegisterSchemaRequest) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *RegisterSchemaRequest) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetReferences

`func (o *RegisterSchemaRequest) GetReferences() []SchemaReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *RegisterSchemaRequest) GetReferencesOk() (*[]SchemaReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *RegisterSchemaRequest) SetReferences(v []SchemaReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *RegisterSchemaRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSchema

`func (o *RegisterSchemaRequest) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *RegisterSchemaRequest) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *RegisterSchemaRequest) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *RegisterSchemaRequest) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



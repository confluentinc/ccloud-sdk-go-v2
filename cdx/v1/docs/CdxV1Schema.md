# CdxV1Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Name of the subject | [optional] 
**Version** | Pointer to **int32** | Version number | [optional] 
**Id** | Pointer to **int32** | Globally unique identifier of the schema | [optional] 
**SchemaType** | Pointer to **string** | Schema type | [optional] 
**Schema** | Pointer to **string** | Schema definition string | [optional] 

## Methods

### NewCdxV1Schema

`func NewCdxV1Schema() *CdxV1Schema`

NewCdxV1Schema instantiates a new CdxV1Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1SchemaWithDefaults

`func NewCdxV1SchemaWithDefaults() *CdxV1Schema`

NewCdxV1SchemaWithDefaults instantiates a new CdxV1Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *CdxV1Schema) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CdxV1Schema) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CdxV1Schema) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CdxV1Schema) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVersion

`func (o *CdxV1Schema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CdxV1Schema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CdxV1Schema) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CdxV1Schema) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *CdxV1Schema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1Schema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1Schema) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1Schema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSchemaType

`func (o *CdxV1Schema) GetSchemaType() string`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *CdxV1Schema) GetSchemaTypeOk() (*string, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *CdxV1Schema) SetSchemaType(v string)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *CdxV1Schema) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetSchema

`func (o *CdxV1Schema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CdxV1Schema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CdxV1Schema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *CdxV1Schema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



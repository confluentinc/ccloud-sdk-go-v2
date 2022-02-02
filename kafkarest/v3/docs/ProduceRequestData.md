# ProduceRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**SubjectNameStrategy** | Pointer to **NullableString** |  | [optional] 
**SchemaId** | Pointer to **NullableInt32** |  | [optional] 
**SchemaVersion** | Pointer to **NullableInt32** |  | [optional] 
**RawSchema** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewProduceRequestData

`func NewProduceRequestData() *ProduceRequestData`

NewProduceRequestData instantiates a new ProduceRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProduceRequestDataWithDefaults

`func NewProduceRequestDataWithDefaults() *ProduceRequestData`

NewProduceRequestDataWithDefaults instantiates a new ProduceRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProduceRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProduceRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProduceRequestData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProduceRequestData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubject

`func (o *ProduceRequestData) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ProduceRequestData) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ProduceRequestData) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ProduceRequestData) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *ProduceRequestData) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *ProduceRequestData) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetSubjectNameStrategy

`func (o *ProduceRequestData) GetSubjectNameStrategy() string`

GetSubjectNameStrategy returns the SubjectNameStrategy field if non-nil, zero value otherwise.

### GetSubjectNameStrategyOk

`func (o *ProduceRequestData) GetSubjectNameStrategyOk() (*string, bool)`

GetSubjectNameStrategyOk returns a tuple with the SubjectNameStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameStrategy

`func (o *ProduceRequestData) SetSubjectNameStrategy(v string)`

SetSubjectNameStrategy sets SubjectNameStrategy field to given value.

### HasSubjectNameStrategy

`func (o *ProduceRequestData) HasSubjectNameStrategy() bool`

HasSubjectNameStrategy returns a boolean if a field has been set.

### SetSubjectNameStrategyNil

`func (o *ProduceRequestData) SetSubjectNameStrategyNil(b bool)`

 SetSubjectNameStrategyNil sets the value for SubjectNameStrategy to be an explicit nil

### UnsetSubjectNameStrategy
`func (o *ProduceRequestData) UnsetSubjectNameStrategy()`

UnsetSubjectNameStrategy ensures that no value is present for SubjectNameStrategy, not even an explicit nil
### GetSchemaId

`func (o *ProduceRequestData) GetSchemaId() int32`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *ProduceRequestData) GetSchemaIdOk() (*int32, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *ProduceRequestData) SetSchemaId(v int32)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *ProduceRequestData) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### SetSchemaIdNil

`func (o *ProduceRequestData) SetSchemaIdNil(b bool)`

 SetSchemaIdNil sets the value for SchemaId to be an explicit nil

### UnsetSchemaId
`func (o *ProduceRequestData) UnsetSchemaId()`

UnsetSchemaId ensures that no value is present for SchemaId, not even an explicit nil
### GetSchemaVersion

`func (o *ProduceRequestData) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ProduceRequestData) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ProduceRequestData) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ProduceRequestData) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### SetSchemaVersionNil

`func (o *ProduceRequestData) SetSchemaVersionNil(b bool)`

 SetSchemaVersionNil sets the value for SchemaVersion to be an explicit nil

### UnsetSchemaVersion
`func (o *ProduceRequestData) UnsetSchemaVersion()`

UnsetSchemaVersion ensures that no value is present for SchemaVersion, not even an explicit nil
### GetRawSchema

`func (o *ProduceRequestData) GetRawSchema() string`

GetRawSchema returns the RawSchema field if non-nil, zero value otherwise.

### GetRawSchemaOk

`func (o *ProduceRequestData) GetRawSchemaOk() (*string, bool)`

GetRawSchemaOk returns a tuple with the RawSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSchema

`func (o *ProduceRequestData) SetRawSchema(v string)`

SetRawSchema sets RawSchema field to given value.

### HasRawSchema

`func (o *ProduceRequestData) HasRawSchema() bool`

HasRawSchema returns a boolean if a field has been set.

### SetRawSchemaNil

`func (o *ProduceRequestData) SetRawSchemaNil(b bool)`

 SetRawSchemaNil sets the value for RawSchema to be an explicit nil

### UnsetRawSchema
`func (o *ProduceRequestData) UnsetRawSchema()`

UnsetRawSchema ensures that no value is present for RawSchema, not even an explicit nil
### GetData

`func (o *ProduceRequestData) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProduceRequestData) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProduceRequestData) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ProduceRequestData) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ProduceRequestData) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ProduceRequestData) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StructField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Type** | [**Type**](Type.md) |  | 
**Required** | **bool** |  | 
**Doc** | Pointer to **string** |  | [optional] 

## Methods

### NewStructField

`func NewStructField(id int32, name string, type_ Type, required bool, ) *StructField`

NewStructField instantiates a new StructField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructFieldWithDefaults

`func NewStructFieldWithDefaults() *StructField`

NewStructFieldWithDefaults instantiates a new StructField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StructField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StructField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StructField) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StructField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StructField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StructField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *StructField) GetType() Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StructField) GetTypeOk() (*Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StructField) SetType(v Type)`

SetType sets Type field to given value.


### GetRequired

`func (o *StructField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *StructField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *StructField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDoc

`func (o *StructField) GetDoc() string`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *StructField) GetDocOk() (*string, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *StructField) SetDoc(v string)`

SetDoc sets Doc field to given value.

### HasDoc

`func (o *StructField) HasDoc() bool`

HasDoc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



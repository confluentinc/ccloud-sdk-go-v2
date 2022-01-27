# V1ServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ServiceAccount

`func NewV1ServiceAccount() *V1ServiceAccount`

NewV1ServiceAccount instantiates a new V1ServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ServiceAccountWithDefaults

`func NewV1ServiceAccountWithDefaults() *V1ServiceAccount`

NewV1ServiceAccountWithDefaults instantiates a new V1ServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1ServiceAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ServiceAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ServiceAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V1ServiceAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceId

`func (o *V1ServiceAccount) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *V1ServiceAccount) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *V1ServiceAccount) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *V1ServiceAccount) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



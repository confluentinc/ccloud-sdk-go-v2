# ByokV1KeySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to [**ByokV1KeySpecKeyOneOf**](ByokV1KeySpecKeyOneOf.md) | The cloud-specific key details.  For AWS please provide the corresponding &#x60;key_arn&#x60;. For Azure please provide the corresponding &#x60;key_id&#x60;.  | [optional] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewByokV1KeySpec

`func NewByokV1KeySpec() *ByokV1KeySpec`

NewByokV1KeySpec instantiates a new ByokV1KeySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewByokV1KeySpecWithDefaults

`func NewByokV1KeySpecWithDefaults() *ByokV1KeySpec`

NewByokV1KeySpecWithDefaults instantiates a new ByokV1KeySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ByokV1KeySpec) GetKey() ByokV1KeySpecKeyOneOf`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ByokV1KeySpec) GetKeyOk() (*ByokV1KeySpecKeyOneOf, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ByokV1KeySpec) SetKey(v ByokV1KeySpecKeyOneOf)`

SetKey sets Key field to given value.

### HasKey

`func (o *ByokV1KeySpec) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetEnvironment

`func (o *ByokV1KeySpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ByokV1KeySpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ByokV1KeySpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ByokV1KeySpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



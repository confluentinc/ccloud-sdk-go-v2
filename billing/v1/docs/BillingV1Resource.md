# BillingV1Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the resource. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the resource. | [optional] 
**Environment** | Pointer to [**NullableBillingV1Environment**](billing.v1.Environment.md) | The environment associated with this resource | [optional] 

## Methods

### NewBillingV1Resource

`func NewBillingV1Resource() *BillingV1Resource`

NewBillingV1Resource instantiates a new BillingV1Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1ResourceWithDefaults

`func NewBillingV1ResourceWithDefaults() *BillingV1Resource`

NewBillingV1ResourceWithDefaults instantiates a new BillingV1Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingV1Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingV1Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingV1Resource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingV1Resource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *BillingV1Resource) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BillingV1Resource) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BillingV1Resource) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BillingV1Resource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnvironment

`func (o *BillingV1Resource) GetEnvironment() BillingV1Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BillingV1Resource) GetEnvironmentOk() (*BillingV1Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BillingV1Resource) SetEnvironment(v BillingV1Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BillingV1Resource) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *BillingV1Resource) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *BillingV1Resource) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# BillingV1alpha1Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the resource. | [optional] 
**Name** | Pointer to **string** | Display name of the resource. | [optional] 
**Environment** | Pointer to [**NullableBillingV1alpha1Environment**](billing.v1alpha1.Environment.md) | The environment associated with this resource | [optional] 

## Methods

### NewBillingV1alpha1Resource

`func NewBillingV1alpha1Resource() *BillingV1alpha1Resource`

NewBillingV1alpha1Resource instantiates a new BillingV1alpha1Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1alpha1ResourceWithDefaults

`func NewBillingV1alpha1ResourceWithDefaults() *BillingV1alpha1Resource`

NewBillingV1alpha1ResourceWithDefaults instantiates a new BillingV1alpha1Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingV1alpha1Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingV1alpha1Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingV1alpha1Resource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingV1alpha1Resource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BillingV1alpha1Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingV1alpha1Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingV1alpha1Resource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingV1alpha1Resource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnvironment

`func (o *BillingV1alpha1Resource) GetEnvironment() BillingV1alpha1Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BillingV1alpha1Resource) GetEnvironmentOk() (*BillingV1alpha1Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BillingV1alpha1Resource) SetEnvironment(v BillingV1alpha1Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BillingV1alpha1Resource) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *BillingV1alpha1Resource) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *BillingV1alpha1Resource) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



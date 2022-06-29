# ResourceDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A named type for a resource (e.g. &#x60;kafka&#x60;, &#x60;connector&#x60;). | 
**Description** | **string** | The description of the resource. | 
**Labels** | [**[]LabelDescriptor**](LabelDescriptor.md) | Labels for the resource. Resource labels can be used for filtering and aggregating metrics associated with a resource.  | 

## Methods

### NewResourceDescriptor

`func NewResourceDescriptor(type_ string, description string, labels []LabelDescriptor, ) *ResourceDescriptor`

NewResourceDescriptor instantiates a new ResourceDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDescriptorWithDefaults

`func NewResourceDescriptorWithDefaults() *ResourceDescriptor`

NewResourceDescriptorWithDefaults instantiates a new ResourceDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceDescriptor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceDescriptor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceDescriptor) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *ResourceDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLabels

`func (o *ResourceDescriptor) GetLabels() []LabelDescriptor`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ResourceDescriptor) GetLabelsOk() (*[]LabelDescriptor, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ResourceDescriptor) SetLabels(v []LabelDescriptor)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



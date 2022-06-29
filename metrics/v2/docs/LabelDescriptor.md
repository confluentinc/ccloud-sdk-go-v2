# LabelDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the metric. | 
**Key** | **string** | The key of the label. | 
**Exportable** | Pointer to **bool** | Is this label included in the &#x60;/export&#x60; endpoint response? | [optional] 

## Methods

### NewLabelDescriptor

`func NewLabelDescriptor(description string, key string, ) *LabelDescriptor`

NewLabelDescriptor instantiates a new LabelDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelDescriptorWithDefaults

`func NewLabelDescriptorWithDefaults() *LabelDescriptor`

NewLabelDescriptorWithDefaults instantiates a new LabelDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LabelDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LabelDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LabelDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetKey

`func (o *LabelDescriptor) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LabelDescriptor) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LabelDescriptor) SetKey(v string)`

SetKey sets Key field to given value.


### GetExportable

`func (o *LabelDescriptor) GetExportable() bool`

GetExportable returns the Exportable field if non-nil, zero value otherwise.

### GetExportableOk

`func (o *LabelDescriptor) GetExportableOk() (*bool, bool)`

GetExportableOk returns a tuple with the Exportable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportable

`func (o *LabelDescriptor) SetExportable(v bool)`

SetExportable sets Exportable field to given value.

### HasExportable

`func (o *LabelDescriptor) HasExportable() bool`

HasExportable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



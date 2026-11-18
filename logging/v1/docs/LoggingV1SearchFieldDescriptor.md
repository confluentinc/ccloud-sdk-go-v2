# LoggingV1SearchFieldDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The &#x60;search&#x60; key. | 
**Type** | **string** | &#x60;enum&#x60; when the catalog restricts this field to &#x60;allowed_values&#x60;; &#x60;string&#x60; for any free-form phrase or array match. | 
**AllowedValues** | Pointer to **[]string** | The values this field accepts. Present only when &#x60;type&#x60; is &#x60;enum&#x60;. | [optional] 

## Methods

### NewLoggingV1SearchFieldDescriptor

`func NewLoggingV1SearchFieldDescriptor(name string, type_ string, ) *LoggingV1SearchFieldDescriptor`

NewLoggingV1SearchFieldDescriptor instantiates a new LoggingV1SearchFieldDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1SearchFieldDescriptorWithDefaults

`func NewLoggingV1SearchFieldDescriptorWithDefaults() *LoggingV1SearchFieldDescriptor`

NewLoggingV1SearchFieldDescriptorWithDefaults instantiates a new LoggingV1SearchFieldDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingV1SearchFieldDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingV1SearchFieldDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingV1SearchFieldDescriptor) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LoggingV1SearchFieldDescriptor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoggingV1SearchFieldDescriptor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoggingV1SearchFieldDescriptor) SetType(v string)`

SetType sets Type field to given value.


### GetAllowedValues

`func (o *LoggingV1SearchFieldDescriptor) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *LoggingV1SearchFieldDescriptor) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *LoggingV1SearchFieldDescriptor) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *LoggingV1SearchFieldDescriptor) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



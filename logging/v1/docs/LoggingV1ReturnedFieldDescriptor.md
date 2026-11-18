# LoggingV1ReturnedFieldDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The field name, as it appears in a search response document. | 
**Type** | **string** | A rendering hint for this field&#39;s value, e.g. &#x60;timestamp&#x60;, &#x60;enum&#x60;, or &#x60;string&#x60;. Not a closed set. | 
**DefaultColumn** | **bool** | Whether this field is shown by default, versus only when explicitly requested. | 
**Description** | Pointer to **string** | Human-readable explanation of this field&#39;s meaning. Omitted when not configured. | [optional] 

## Methods

### NewLoggingV1ReturnedFieldDescriptor

`func NewLoggingV1ReturnedFieldDescriptor(name string, type_ string, defaultColumn bool, ) *LoggingV1ReturnedFieldDescriptor`

NewLoggingV1ReturnedFieldDescriptor instantiates a new LoggingV1ReturnedFieldDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingV1ReturnedFieldDescriptorWithDefaults

`func NewLoggingV1ReturnedFieldDescriptorWithDefaults() *LoggingV1ReturnedFieldDescriptor`

NewLoggingV1ReturnedFieldDescriptorWithDefaults instantiates a new LoggingV1ReturnedFieldDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingV1ReturnedFieldDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingV1ReturnedFieldDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingV1ReturnedFieldDescriptor) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LoggingV1ReturnedFieldDescriptor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoggingV1ReturnedFieldDescriptor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoggingV1ReturnedFieldDescriptor) SetType(v string)`

SetType sets Type field to given value.


### GetDefaultColumn

`func (o *LoggingV1ReturnedFieldDescriptor) GetDefaultColumn() bool`

GetDefaultColumn returns the DefaultColumn field if non-nil, zero value otherwise.

### GetDefaultColumnOk

`func (o *LoggingV1ReturnedFieldDescriptor) GetDefaultColumnOk() (*bool, bool)`

GetDefaultColumnOk returns a tuple with the DefaultColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColumn

`func (o *LoggingV1ReturnedFieldDescriptor) SetDefaultColumn(v bool)`

SetDefaultColumn sets DefaultColumn field to given value.


### GetDescription

`func (o *LoggingV1ReturnedFieldDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoggingV1ReturnedFieldDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoggingV1ReturnedFieldDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoggingV1ReturnedFieldDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



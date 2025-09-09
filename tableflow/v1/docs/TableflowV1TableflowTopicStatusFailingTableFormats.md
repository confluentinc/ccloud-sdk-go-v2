# TableflowV1TableflowTopicStatusFailingTableFormats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | The name of the table format (e.g., DELTA, ICEBERG). | 
**ErrorMessage** | **string** | The error message for the failing table format. | 

## Methods

### NewTableflowV1TableflowTopicStatusFailingTableFormats

`func NewTableflowV1TableflowTopicStatusFailingTableFormats(format string, errorMessage string, ) *TableflowV1TableflowTopicStatusFailingTableFormats`

NewTableflowV1TableflowTopicStatusFailingTableFormats instantiates a new TableflowV1TableflowTopicStatusFailingTableFormats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1TableflowTopicStatusFailingTableFormatsWithDefaults

`func NewTableflowV1TableflowTopicStatusFailingTableFormatsWithDefaults() *TableflowV1TableflowTopicStatusFailingTableFormats`

NewTableflowV1TableflowTopicStatusFailingTableFormatsWithDefaults instantiates a new TableflowV1TableflowTopicStatusFailingTableFormats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *TableflowV1TableflowTopicStatusFailingTableFormats) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TableflowV1TableflowTopicStatusFailingTableFormats) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TableflowV1TableflowTopicStatusFailingTableFormats) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetErrorMessage

`func (o *TableflowV1TableflowTopicStatusFailingTableFormats) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TableflowV1TableflowTopicStatusFailingTableFormats) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TableflowV1TableflowTopicStatusFailingTableFormats) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



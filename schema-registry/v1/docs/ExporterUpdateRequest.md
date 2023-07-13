# ExporterUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextType** | Pointer to **string** | Context type of the exporter. One of CUSTOM, NONE or AUTO (default) | [optional] 
**Context** | Pointer to **string** | Context of the exporter if contextType equals CUSTOM | [optional] 

## Methods

### NewExporterUpdateRequest

`func NewExporterUpdateRequest() *ExporterUpdateRequest`

NewExporterUpdateRequest instantiates a new ExporterUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExporterUpdateRequestWithDefaults

`func NewExporterUpdateRequestWithDefaults() *ExporterUpdateRequest`

NewExporterUpdateRequestWithDefaults instantiates a new ExporterUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextType

`func (o *ExporterUpdateRequest) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *ExporterUpdateRequest) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *ExporterUpdateRequest) SetContextType(v string)`

SetContextType sets ContextType field to given value.

### HasContextType

`func (o *ExporterUpdateRequest) HasContextType() bool`

HasContextType returns a boolean if a field has been set.

### GetContext

`func (o *ExporterUpdateRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ExporterUpdateRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ExporterUpdateRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ExporterUpdateRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



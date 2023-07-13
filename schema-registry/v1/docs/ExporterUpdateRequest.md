# ExporterUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextType** | Pointer to **string** | Context type of the exporter. One of CUSTOM, NONE or AUTO (default) | [optional] 
**Context** | Pointer to **string** | Customized context of the exporter if contextType equals CUSTOM. | [optional] 
**Subjects** | Pointer to **[]string** | Name of each exporter subject | [optional] 
**SubjectRenameFormat** | Pointer to **string** | Format string for the subject name in the destination cluster, which may contain ${subject} as a placeholder for the originating subject name. For example, dc_${subject} for the subject orders will map to the destination subject name dc_orders. | [optional] 
**Config** | Pointer to **map[string]string** | The map containing exporter’s configurations | [optional] 

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

### GetSubjects

`func (o *ExporterUpdateRequest) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *ExporterUpdateRequest) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *ExporterUpdateRequest) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *ExporterUpdateRequest) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetSubjectRenameFormat

`func (o *ExporterUpdateRequest) GetSubjectRenameFormat() string`

GetSubjectRenameFormat returns the SubjectRenameFormat field if non-nil, zero value otherwise.

### GetSubjectRenameFormatOk

`func (o *ExporterUpdateRequest) GetSubjectRenameFormatOk() (*string, bool)`

GetSubjectRenameFormatOk returns a tuple with the SubjectRenameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRenameFormat

`func (o *ExporterUpdateRequest) SetSubjectRenameFormat(v string)`

SetSubjectRenameFormat sets SubjectRenameFormat field to given value.

### HasSubjectRenameFormat

`func (o *ExporterUpdateRequest) HasSubjectRenameFormat() bool`

HasSubjectRenameFormat returns a boolean if a field has been set.

### GetConfig

`func (o *ExporterUpdateRequest) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExporterUpdateRequest) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExporterUpdateRequest) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExporterUpdateRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



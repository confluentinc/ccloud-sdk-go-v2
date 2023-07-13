# ExporterReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the exporter | [optional] 
**ContextType** | Pointer to **string** | Context type of the exporter. One of CUSTOM, NONE or AUTO (default) | [optional] 
**Context** | Pointer to **string** | Customized context of the exporter if contextType equals CUSTOM. | [optional] 
**Subjects** | Pointer to **[]string** | Name of each exporter subject | [optional] 
**SubjectRenameFormat** | Pointer to **string** | Format string for the subject name in the destination cluster, which may contain ${subject} as a placeholder for the originating subject name. For example, dc_${subject} for the subject orders will map to the destination subject name dc_orders. | [optional] 
**Config** | Pointer to **map[string]string** | The map containing exporter’s configurations | [optional] 

## Methods

### NewExporterReference

`func NewExporterReference() *ExporterReference`

NewExporterReference instantiates a new ExporterReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExporterReferenceWithDefaults

`func NewExporterReferenceWithDefaults() *ExporterReference`

NewExporterReferenceWithDefaults instantiates a new ExporterReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExporterReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExporterReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExporterReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExporterReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContextType

`func (o *ExporterReference) GetContextType() string`

GetContextType returns the ContextType field if non-nil, zero value otherwise.

### GetContextTypeOk

`func (o *ExporterReference) GetContextTypeOk() (*string, bool)`

GetContextTypeOk returns a tuple with the ContextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextType

`func (o *ExporterReference) SetContextType(v string)`

SetContextType sets ContextType field to given value.

### HasContextType

`func (o *ExporterReference) HasContextType() bool`

HasContextType returns a boolean if a field has been set.

### GetContext

`func (o *ExporterReference) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ExporterReference) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ExporterReference) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ExporterReference) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetSubjects

`func (o *ExporterReference) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *ExporterReference) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *ExporterReference) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *ExporterReference) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetSubjectRenameFormat

`func (o *ExporterReference) GetSubjectRenameFormat() string`

GetSubjectRenameFormat returns the SubjectRenameFormat field if non-nil, zero value otherwise.

### GetSubjectRenameFormatOk

`func (o *ExporterReference) GetSubjectRenameFormatOk() (*string, bool)`

GetSubjectRenameFormatOk returns a tuple with the SubjectRenameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectRenameFormat

`func (o *ExporterReference) SetSubjectRenameFormat(v string)`

SetSubjectRenameFormat sets SubjectRenameFormat field to given value.

### HasSubjectRenameFormat

`func (o *ExporterReference) HasSubjectRenameFormat() bool`

HasSubjectRenameFormat returns a boolean if a field has been set.

### GetConfig

`func (o *ExporterReference) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExporterReference) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExporterReference) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExporterReference) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



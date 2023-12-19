# TemplateV1WidgetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]TemplateV1Widget**](TemplateV1Widget.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewTemplateV1WidgetList

`func NewTemplateV1WidgetList(apiVersion string, kind string, metadata ListMeta, data []TemplateV1Widget, ) *TemplateV1WidgetList`

NewTemplateV1WidgetList instantiates a new TemplateV1WidgetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateV1WidgetListWithDefaults

`func NewTemplateV1WidgetListWithDefaults() *TemplateV1WidgetList`

NewTemplateV1WidgetListWithDefaults instantiates a new TemplateV1WidgetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *TemplateV1WidgetList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TemplateV1WidgetList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TemplateV1WidgetList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *TemplateV1WidgetList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TemplateV1WidgetList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TemplateV1WidgetList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *TemplateV1WidgetList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TemplateV1WidgetList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TemplateV1WidgetList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *TemplateV1WidgetList) GetData() []TemplateV1Widget`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TemplateV1WidgetList) GetDataOk() (*[]TemplateV1Widget, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TemplateV1WidgetList) SetData(v []TemplateV1Widget)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



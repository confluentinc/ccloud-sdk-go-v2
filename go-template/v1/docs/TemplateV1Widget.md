# TemplateV1Widget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**TemplateV1WidgetSpec**](TemplateV1WidgetSpec.md) |  | [optional] 
**Status** | Pointer to [**TemplateV1WidgetStatus**](TemplateV1WidgetStatus.md) |  | [optional] 

## Methods

### NewTemplateV1Widget

`func NewTemplateV1Widget() *TemplateV1Widget`

NewTemplateV1Widget instantiates a new TemplateV1Widget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateV1WidgetWithDefaults

`func NewTemplateV1WidgetWithDefaults() *TemplateV1Widget`

NewTemplateV1WidgetWithDefaults instantiates a new TemplateV1Widget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *TemplateV1Widget) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TemplateV1Widget) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TemplateV1Widget) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TemplateV1Widget) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *TemplateV1Widget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TemplateV1Widget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TemplateV1Widget) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TemplateV1Widget) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *TemplateV1Widget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateV1Widget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateV1Widget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateV1Widget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *TemplateV1Widget) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TemplateV1Widget) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TemplateV1Widget) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TemplateV1Widget) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *TemplateV1Widget) GetSpec() TemplateV1WidgetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *TemplateV1Widget) GetSpecOk() (*TemplateV1WidgetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *TemplateV1Widget) SetSpec(v TemplateV1WidgetSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *TemplateV1Widget) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *TemplateV1Widget) GetStatus() TemplateV1WidgetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateV1Widget) GetStatusOk() (*TemplateV1WidgetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateV1Widget) SetStatus(v TemplateV1WidgetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TemplateV1Widget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



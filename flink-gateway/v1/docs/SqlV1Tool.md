# SqlV1Tool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Name** | Pointer to **string** | The user provided name of the tool, unique within this environment. | [optional] 
**Spec** | Pointer to [**SqlV1ToolSpec**](SqlV1ToolSpec.md) |  | [optional] 
**Status** | Pointer to [**SqlV1ToolStatus**](SqlV1ToolStatus.md) |  | [optional] 

## Methods

### NewSqlV1Tool

`func NewSqlV1Tool() *SqlV1Tool`

NewSqlV1Tool instantiates a new SqlV1Tool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ToolWithDefaults

`func NewSqlV1ToolWithDefaults() *SqlV1Tool`

NewSqlV1ToolWithDefaults instantiates a new SqlV1Tool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1Tool) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1Tool) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1Tool) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SqlV1Tool) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SqlV1Tool) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1Tool) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1Tool) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SqlV1Tool) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *SqlV1Tool) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1Tool) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1Tool) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SqlV1Tool) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SqlV1Tool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1Tool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1Tool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlV1Tool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpec

`func (o *SqlV1Tool) GetSpec() SqlV1ToolSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SqlV1Tool) GetSpecOk() (*SqlV1ToolSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SqlV1Tool) SetSpec(v SqlV1ToolSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SqlV1Tool) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SqlV1Tool) GetStatus() SqlV1ToolStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SqlV1Tool) GetStatusOk() (*SqlV1ToolStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SqlV1Tool) SetStatus(v SqlV1ToolStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SqlV1Tool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



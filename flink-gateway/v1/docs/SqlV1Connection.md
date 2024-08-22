# SqlV1Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Name** | Pointer to **string** | The user provided name of the resource, unique within this environment. | [optional] 
**Spec** | Pointer to [**SqlV1ConnectionSpec**](SqlV1ConnectionSpec.md) |  | [optional] 
**Status** | Pointer to [**SqlV1ConnectionStatus**](SqlV1ConnectionStatus.md) |  | [optional] 

## Methods

### NewSqlV1Connection

`func NewSqlV1Connection() *SqlV1Connection`

NewSqlV1Connection instantiates a new SqlV1Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ConnectionWithDefaults

`func NewSqlV1ConnectionWithDefaults() *SqlV1Connection`

NewSqlV1ConnectionWithDefaults instantiates a new SqlV1Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SqlV1Connection) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SqlV1Connection) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SqlV1Connection) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SqlV1Connection) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SqlV1Connection) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1Connection) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1Connection) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SqlV1Connection) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *SqlV1Connection) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SqlV1Connection) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SqlV1Connection) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SqlV1Connection) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SqlV1Connection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1Connection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1Connection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlV1Connection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpec

`func (o *SqlV1Connection) GetSpec() SqlV1ConnectionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SqlV1Connection) GetSpecOk() (*SqlV1ConnectionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SqlV1Connection) SetSpec(v SqlV1ConnectionSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SqlV1Connection) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SqlV1Connection) GetStatus() SqlV1ConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SqlV1Connection) GetStatusOk() (*SqlV1ConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SqlV1Connection) SetStatus(v SqlV1ConnectionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SqlV1Connection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



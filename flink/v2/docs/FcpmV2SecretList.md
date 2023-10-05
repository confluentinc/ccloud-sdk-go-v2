# FcpmV2SecretList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]FcpmV2Secret**](FcpmV2Secret.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewFcpmV2SecretList

`func NewFcpmV2SecretList(apiVersion string, kind string, metadata ListMeta, data []FcpmV2Secret, ) *FcpmV2SecretList`

NewFcpmV2SecretList instantiates a new FcpmV2SecretList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcpmV2SecretListWithDefaults

`func NewFcpmV2SecretListWithDefaults() *FcpmV2SecretList`

NewFcpmV2SecretListWithDefaults instantiates a new FcpmV2SecretList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FcpmV2SecretList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FcpmV2SecretList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FcpmV2SecretList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *FcpmV2SecretList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FcpmV2SecretList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FcpmV2SecretList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *FcpmV2SecretList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FcpmV2SecretList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FcpmV2SecretList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *FcpmV2SecretList) GetData() []FcpmV2Secret`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FcpmV2SecretList) GetDataOk() (*[]FcpmV2Secret, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FcpmV2SecretList) SetData(v []FcpmV2Secret)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



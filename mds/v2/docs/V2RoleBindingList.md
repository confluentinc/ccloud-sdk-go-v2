# V2RoleBindingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]V2RoleBinding**](V2RoleBinding.md) |  | 

## Methods

### NewV2RoleBindingList

`func NewV2RoleBindingList(apiVersion string, kind string, metadata ListMeta, data []V2RoleBinding, ) *V2RoleBindingList`

NewV2RoleBindingList instantiates a new V2RoleBindingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RoleBindingListWithDefaults

`func NewV2RoleBindingListWithDefaults() *V2RoleBindingList`

NewV2RoleBindingListWithDefaults instantiates a new V2RoleBindingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V2RoleBindingList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V2RoleBindingList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V2RoleBindingList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *V2RoleBindingList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V2RoleBindingList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V2RoleBindingList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *V2RoleBindingList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V2RoleBindingList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V2RoleBindingList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *V2RoleBindingList) GetData() []V2RoleBinding`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2RoleBindingList) GetDataOk() (*[]V2RoleBinding, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2RoleBindingList) SetData(v []V2RoleBinding)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



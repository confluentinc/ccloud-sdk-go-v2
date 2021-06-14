# V2UserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | Pointer to [**ListMeta**](ListMeta.md) |  | 
**Data** | Pointer to [**[]V2User**](v2.User.md) |  | 

## Methods

### NewV2UserList

`func NewV2UserList(apiVersion string, kind string, metadata ListMeta, data []V2User, ) *V2UserList`

NewV2UserList instantiates a new V2UserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2UserListWithDefaults

`func NewV2UserListWithDefaults() *V2UserList`

NewV2UserListWithDefaults instantiates a new V2UserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V2UserList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V2UserList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V2UserList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *V2UserList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V2UserList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V2UserList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *V2UserList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V2UserList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V2UserList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *V2UserList) GetData() []V2User`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2UserList) GetDataOk() (*[]V2User, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2UserList) SetData(v []V2User)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



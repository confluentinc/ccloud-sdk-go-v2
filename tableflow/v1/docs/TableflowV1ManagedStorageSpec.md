# TableflowV1ManagedStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The storage type.  | 
**TablePath** | Pointer to **string** | The current storage path where the data and metadata is stored for this table | [optional] [readonly] 

## Methods

### NewTableflowV1ManagedStorageSpec

`func NewTableflowV1ManagedStorageSpec(kind string, ) *TableflowV1ManagedStorageSpec`

NewTableflowV1ManagedStorageSpec instantiates a new TableflowV1ManagedStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1ManagedStorageSpecWithDefaults

`func NewTableflowV1ManagedStorageSpecWithDefaults() *TableflowV1ManagedStorageSpec`

NewTableflowV1ManagedStorageSpecWithDefaults instantiates a new TableflowV1ManagedStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1ManagedStorageSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1ManagedStorageSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1ManagedStorageSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTablePath

`func (o *TableflowV1ManagedStorageSpec) GetTablePath() string`

GetTablePath returns the TablePath field if non-nil, zero value otherwise.

### GetTablePathOk

`func (o *TableflowV1ManagedStorageSpec) GetTablePathOk() (*string, bool)`

GetTablePathOk returns a tuple with the TablePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablePath

`func (o *TableflowV1ManagedStorageSpec) SetTablePath(v string)`

SetTablePath sets TablePath field to given value.

### HasTablePath

`func (o *TableflowV1ManagedStorageSpec) HasTablePath() bool`

HasTablePath returns a boolean if a field has been set.


### AsTableflowV1TableflowTopicSpecStorageOneOf

`func (s *TableflowV1ManagedStorageSpec) AsTableflowV1TableflowTopicSpecStorageOneOf() TableflowV1TableflowTopicSpecStorageOneOf`

Convenience method to wrap this instance of TableflowV1ManagedStorageSpec in TableflowV1TableflowTopicSpecStorageOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



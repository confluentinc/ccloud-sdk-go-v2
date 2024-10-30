# ListLinksResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**SourceClusterId** | Pointer to **NullableString** |  | [optional] 
**DestinationClusterId** | Pointer to **NullableString** |  | [optional] 
**RemoteClusterId** | Pointer to **NullableString** |  | [optional] 
**LinkName** | **string** |  | 
**LinkId** | Pointer to **string** |  | [optional] 
**ClusterLinkId** | **string** |  | 
**TopicNames** | **[]string** |  | 
**LinkError** | Pointer to **string** |  | [optional] 
**LinkErrorMessage** | Pointer to **NullableString** |  | [optional] 
**LinkState** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to [**[]LinkTask**](LinkTask.md) |  | [optional] 
**CategoryCounts** | Pointer to [**[]LinkCategory**](LinkCategory.md) |  | [optional] 

## Methods

### NewListLinksResponseData

`func NewListLinksResponseData(kind string, metadata ResourceMetadata, linkName string, clusterLinkId string, topicNames []string, ) *ListLinksResponseData`

NewListLinksResponseData instantiates a new ListLinksResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLinksResponseDataWithDefaults

`func NewListLinksResponseDataWithDefaults() *ListLinksResponseData`

NewListLinksResponseDataWithDefaults instantiates a new ListLinksResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ListLinksResponseData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListLinksResponseData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListLinksResponseData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ListLinksResponseData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListLinksResponseData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListLinksResponseData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetSourceClusterId

`func (o *ListLinksResponseData) GetSourceClusterId() string`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### GetSourceClusterIdOk

`func (o *ListLinksResponseData) GetSourceClusterIdOk() (*string, bool)`

GetSourceClusterIdOk returns a tuple with the SourceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterId

`func (o *ListLinksResponseData) SetSourceClusterId(v string)`

SetSourceClusterId sets SourceClusterId field to given value.

### HasSourceClusterId

`func (o *ListLinksResponseData) HasSourceClusterId() bool`

HasSourceClusterId returns a boolean if a field has been set.

### SetSourceClusterIdNil

`func (o *ListLinksResponseData) SetSourceClusterIdNil(b bool)`

 SetSourceClusterIdNil sets the value for SourceClusterId to be an explicit nil

### UnsetSourceClusterId
`func (o *ListLinksResponseData) UnsetSourceClusterId()`

UnsetSourceClusterId ensures that no value is present for SourceClusterId, not even an explicit nil
### GetDestinationClusterId

`func (o *ListLinksResponseData) GetDestinationClusterId() string`

GetDestinationClusterId returns the DestinationClusterId field if non-nil, zero value otherwise.

### GetDestinationClusterIdOk

`func (o *ListLinksResponseData) GetDestinationClusterIdOk() (*string, bool)`

GetDestinationClusterIdOk returns a tuple with the DestinationClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationClusterId

`func (o *ListLinksResponseData) SetDestinationClusterId(v string)`

SetDestinationClusterId sets DestinationClusterId field to given value.

### HasDestinationClusterId

`func (o *ListLinksResponseData) HasDestinationClusterId() bool`

HasDestinationClusterId returns a boolean if a field has been set.

### SetDestinationClusterIdNil

`func (o *ListLinksResponseData) SetDestinationClusterIdNil(b bool)`

 SetDestinationClusterIdNil sets the value for DestinationClusterId to be an explicit nil

### UnsetDestinationClusterId
`func (o *ListLinksResponseData) UnsetDestinationClusterId()`

UnsetDestinationClusterId ensures that no value is present for DestinationClusterId, not even an explicit nil
### GetRemoteClusterId

`func (o *ListLinksResponseData) GetRemoteClusterId() string`

GetRemoteClusterId returns the RemoteClusterId field if non-nil, zero value otherwise.

### GetRemoteClusterIdOk

`func (o *ListLinksResponseData) GetRemoteClusterIdOk() (*string, bool)`

GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterId

`func (o *ListLinksResponseData) SetRemoteClusterId(v string)`

SetRemoteClusterId sets RemoteClusterId field to given value.

### HasRemoteClusterId

`func (o *ListLinksResponseData) HasRemoteClusterId() bool`

HasRemoteClusterId returns a boolean if a field has been set.

### SetRemoteClusterIdNil

`func (o *ListLinksResponseData) SetRemoteClusterIdNil(b bool)`

 SetRemoteClusterIdNil sets the value for RemoteClusterId to be an explicit nil

### UnsetRemoteClusterId
`func (o *ListLinksResponseData) UnsetRemoteClusterId()`

UnsetRemoteClusterId ensures that no value is present for RemoteClusterId, not even an explicit nil
### GetLinkName

`func (o *ListLinksResponseData) GetLinkName() string`

GetLinkName returns the LinkName field if non-nil, zero value otherwise.

### GetLinkNameOk

`func (o *ListLinksResponseData) GetLinkNameOk() (*string, bool)`

GetLinkNameOk returns a tuple with the LinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkName

`func (o *ListLinksResponseData) SetLinkName(v string)`

SetLinkName sets LinkName field to given value.


### GetLinkId

`func (o *ListLinksResponseData) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *ListLinksResponseData) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *ListLinksResponseData) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *ListLinksResponseData) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetClusterLinkId

`func (o *ListLinksResponseData) GetClusterLinkId() string`

GetClusterLinkId returns the ClusterLinkId field if non-nil, zero value otherwise.

### GetClusterLinkIdOk

`func (o *ListLinksResponseData) GetClusterLinkIdOk() (*string, bool)`

GetClusterLinkIdOk returns a tuple with the ClusterLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterLinkId

`func (o *ListLinksResponseData) SetClusterLinkId(v string)`

SetClusterLinkId sets ClusterLinkId field to given value.


### GetTopicNames

`func (o *ListLinksResponseData) GetTopicNames() []string`

GetTopicNames returns the TopicNames field if non-nil, zero value otherwise.

### GetTopicNamesOk

`func (o *ListLinksResponseData) GetTopicNamesOk() (*[]string, bool)`

GetTopicNamesOk returns a tuple with the TopicNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicNames

`func (o *ListLinksResponseData) SetTopicNames(v []string)`

SetTopicNames sets TopicNames field to given value.


### GetLinkError

`func (o *ListLinksResponseData) GetLinkError() string`

GetLinkError returns the LinkError field if non-nil, zero value otherwise.

### GetLinkErrorOk

`func (o *ListLinksResponseData) GetLinkErrorOk() (*string, bool)`

GetLinkErrorOk returns a tuple with the LinkError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkError

`func (o *ListLinksResponseData) SetLinkError(v string)`

SetLinkError sets LinkError field to given value.

### HasLinkError

`func (o *ListLinksResponseData) HasLinkError() bool`

HasLinkError returns a boolean if a field has been set.

### GetLinkErrorMessage

`func (o *ListLinksResponseData) GetLinkErrorMessage() string`

GetLinkErrorMessage returns the LinkErrorMessage field if non-nil, zero value otherwise.

### GetLinkErrorMessageOk

`func (o *ListLinksResponseData) GetLinkErrorMessageOk() (*string, bool)`

GetLinkErrorMessageOk returns a tuple with the LinkErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkErrorMessage

`func (o *ListLinksResponseData) SetLinkErrorMessage(v string)`

SetLinkErrorMessage sets LinkErrorMessage field to given value.

### HasLinkErrorMessage

`func (o *ListLinksResponseData) HasLinkErrorMessage() bool`

HasLinkErrorMessage returns a boolean if a field has been set.

### SetLinkErrorMessageNil

`func (o *ListLinksResponseData) SetLinkErrorMessageNil(b bool)`

 SetLinkErrorMessageNil sets the value for LinkErrorMessage to be an explicit nil

### UnsetLinkErrorMessage
`func (o *ListLinksResponseData) UnsetLinkErrorMessage()`

UnsetLinkErrorMessage ensures that no value is present for LinkErrorMessage, not even an explicit nil
### GetLinkState

`func (o *ListLinksResponseData) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *ListLinksResponseData) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *ListLinksResponseData) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *ListLinksResponseData) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetTasks

`func (o *ListLinksResponseData) GetTasks() []LinkTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ListLinksResponseData) GetTasksOk() (*[]LinkTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ListLinksResponseData) SetTasks(v []LinkTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ListLinksResponseData) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *ListLinksResponseData) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *ListLinksResponseData) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil
### GetCategoryCounts

`func (o *ListLinksResponseData) GetCategoryCounts() []LinkCategory`

GetCategoryCounts returns the CategoryCounts field if non-nil, zero value otherwise.

### GetCategoryCountsOk

`func (o *ListLinksResponseData) GetCategoryCountsOk() (*[]LinkCategory, bool)`

GetCategoryCountsOk returns a tuple with the CategoryCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryCounts

`func (o *ListLinksResponseData) SetCategoryCounts(v []LinkCategory)`

SetCategoryCounts sets CategoryCounts field to given value.

### HasCategoryCounts

`func (o *ListLinksResponseData) HasCategoryCounts() bool`

HasCategoryCounts returns a boolean if a field has been set.

### SetCategoryCountsNil

`func (o *ListLinksResponseData) SetCategoryCountsNil(b bool)`

 SetCategoryCountsNil sets the value for CategoryCounts to be an explicit nil

### UnsetCategoryCounts
`func (o *ListLinksResponseData) UnsetCategoryCounts()`

UnsetCategoryCounts ensures that no value is present for CategoryCounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



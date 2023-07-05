# ListLinksResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewListLinksResponseDataAllOf

`func NewListLinksResponseDataAllOf(linkName string, clusterLinkId string, topicNames []string, ) *ListLinksResponseDataAllOf`

NewListLinksResponseDataAllOf instantiates a new ListLinksResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLinksResponseDataAllOfWithDefaults

`func NewListLinksResponseDataAllOfWithDefaults() *ListLinksResponseDataAllOf`

NewListLinksResponseDataAllOfWithDefaults instantiates a new ListLinksResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceClusterId

`func (o *ListLinksResponseDataAllOf) GetSourceClusterId() string`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### GetSourceClusterIdOk

`func (o *ListLinksResponseDataAllOf) GetSourceClusterIdOk() (*string, bool)`

GetSourceClusterIdOk returns a tuple with the SourceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterId

`func (o *ListLinksResponseDataAllOf) SetSourceClusterId(v string)`

SetSourceClusterId sets SourceClusterId field to given value.

### HasSourceClusterId

`func (o *ListLinksResponseDataAllOf) HasSourceClusterId() bool`

HasSourceClusterId returns a boolean if a field has been set.

### SetSourceClusterIdNil

`func (o *ListLinksResponseDataAllOf) SetSourceClusterIdNil(b bool)`

 SetSourceClusterIdNil sets the value for SourceClusterId to be an explicit nil

### UnsetSourceClusterId
`func (o *ListLinksResponseDataAllOf) UnsetSourceClusterId()`

UnsetSourceClusterId ensures that no value is present for SourceClusterId, not even an explicit nil
### GetDestinationClusterId

`func (o *ListLinksResponseDataAllOf) GetDestinationClusterId() string`

GetDestinationClusterId returns the DestinationClusterId field if non-nil, zero value otherwise.

### GetDestinationClusterIdOk

`func (o *ListLinksResponseDataAllOf) GetDestinationClusterIdOk() (*string, bool)`

GetDestinationClusterIdOk returns a tuple with the DestinationClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationClusterId

`func (o *ListLinksResponseDataAllOf) SetDestinationClusterId(v string)`

SetDestinationClusterId sets DestinationClusterId field to given value.

### HasDestinationClusterId

`func (o *ListLinksResponseDataAllOf) HasDestinationClusterId() bool`

HasDestinationClusterId returns a boolean if a field has been set.

### SetDestinationClusterIdNil

`func (o *ListLinksResponseDataAllOf) SetDestinationClusterIdNil(b bool)`

 SetDestinationClusterIdNil sets the value for DestinationClusterId to be an explicit nil

### UnsetDestinationClusterId
`func (o *ListLinksResponseDataAllOf) UnsetDestinationClusterId()`

UnsetDestinationClusterId ensures that no value is present for DestinationClusterId, not even an explicit nil
### GetRemoteClusterId

`func (o *ListLinksResponseDataAllOf) GetRemoteClusterId() string`

GetRemoteClusterId returns the RemoteClusterId field if non-nil, zero value otherwise.

### GetRemoteClusterIdOk

`func (o *ListLinksResponseDataAllOf) GetRemoteClusterIdOk() (*string, bool)`

GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterId

`func (o *ListLinksResponseDataAllOf) SetRemoteClusterId(v string)`

SetRemoteClusterId sets RemoteClusterId field to given value.

### HasRemoteClusterId

`func (o *ListLinksResponseDataAllOf) HasRemoteClusterId() bool`

HasRemoteClusterId returns a boolean if a field has been set.

### SetRemoteClusterIdNil

`func (o *ListLinksResponseDataAllOf) SetRemoteClusterIdNil(b bool)`

 SetRemoteClusterIdNil sets the value for RemoteClusterId to be an explicit nil

### UnsetRemoteClusterId
`func (o *ListLinksResponseDataAllOf) UnsetRemoteClusterId()`

UnsetRemoteClusterId ensures that no value is present for RemoteClusterId, not even an explicit nil
### GetLinkName

`func (o *ListLinksResponseDataAllOf) GetLinkName() string`

GetLinkName returns the LinkName field if non-nil, zero value otherwise.

### GetLinkNameOk

`func (o *ListLinksResponseDataAllOf) GetLinkNameOk() (*string, bool)`

GetLinkNameOk returns a tuple with the LinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkName

`func (o *ListLinksResponseDataAllOf) SetLinkName(v string)`

SetLinkName sets LinkName field to given value.


### GetLinkId

`func (o *ListLinksResponseDataAllOf) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *ListLinksResponseDataAllOf) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *ListLinksResponseDataAllOf) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *ListLinksResponseDataAllOf) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetClusterLinkId

`func (o *ListLinksResponseDataAllOf) GetClusterLinkId() string`

GetClusterLinkId returns the ClusterLinkId field if non-nil, zero value otherwise.

### GetClusterLinkIdOk

`func (o *ListLinksResponseDataAllOf) GetClusterLinkIdOk() (*string, bool)`

GetClusterLinkIdOk returns a tuple with the ClusterLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterLinkId

`func (o *ListLinksResponseDataAllOf) SetClusterLinkId(v string)`

SetClusterLinkId sets ClusterLinkId field to given value.


### GetTopicNames

`func (o *ListLinksResponseDataAllOf) GetTopicNames() []string`

GetTopicNames returns the TopicNames field if non-nil, zero value otherwise.

### GetTopicNamesOk

`func (o *ListLinksResponseDataAllOf) GetTopicNamesOk() (*[]string, bool)`

GetTopicNamesOk returns a tuple with the TopicNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicNames

`func (o *ListLinksResponseDataAllOf) SetTopicNames(v []string)`

SetTopicNames sets TopicNames field to given value.


### GetLinkError

`func (o *ListLinksResponseDataAllOf) GetLinkError() string`

GetLinkError returns the LinkError field if non-nil, zero value otherwise.

### GetLinkErrorOk

`func (o *ListLinksResponseDataAllOf) GetLinkErrorOk() (*string, bool)`

GetLinkErrorOk returns a tuple with the LinkError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkError

`func (o *ListLinksResponseDataAllOf) SetLinkError(v string)`

SetLinkError sets LinkError field to given value.

### HasLinkError

`func (o *ListLinksResponseDataAllOf) HasLinkError() bool`

HasLinkError returns a boolean if a field has been set.

### GetLinkErrorMessage

`func (o *ListLinksResponseDataAllOf) GetLinkErrorMessage() string`

GetLinkErrorMessage returns the LinkErrorMessage field if non-nil, zero value otherwise.

### GetLinkErrorMessageOk

`func (o *ListLinksResponseDataAllOf) GetLinkErrorMessageOk() (*string, bool)`

GetLinkErrorMessageOk returns a tuple with the LinkErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkErrorMessage

`func (o *ListLinksResponseDataAllOf) SetLinkErrorMessage(v string)`

SetLinkErrorMessage sets LinkErrorMessage field to given value.

### HasLinkErrorMessage

`func (o *ListLinksResponseDataAllOf) HasLinkErrorMessage() bool`

HasLinkErrorMessage returns a boolean if a field has been set.

### SetLinkErrorMessageNil

`func (o *ListLinksResponseDataAllOf) SetLinkErrorMessageNil(b bool)`

 SetLinkErrorMessageNil sets the value for LinkErrorMessage to be an explicit nil

### UnsetLinkErrorMessage
`func (o *ListLinksResponseDataAllOf) UnsetLinkErrorMessage()`

UnsetLinkErrorMessage ensures that no value is present for LinkErrorMessage, not even an explicit nil
### GetLinkState

`func (o *ListLinksResponseDataAllOf) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *ListLinksResponseDataAllOf) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *ListLinksResponseDataAllOf) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *ListLinksResponseDataAllOf) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListLinksResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**SourceClusterId** | Pointer to **NullableString** |  | [optional] 
**DestinationClusterId** | Pointer to **NullableString** |  | [optional] 
**LinkName** | **string** |  | 
**LinkId** | **string** |  | 
**TopicsNames** | Pointer to **[]string** |  | [optional] 
**LinkError** | **string** |  | 
**LinkErrorMessage** | **string** |  | 
**LinkState** | **string** |  | 

## Methods

### NewListLinksResponseData

`func NewListLinksResponseData(kind string, metadata ResourceMetadata, linkName string, linkId string, linkError string, linkErrorMessage string, linkState string, ) *ListLinksResponseData`

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


### GetTopicsNames

`func (o *ListLinksResponseData) GetTopicsNames() []string`

GetTopicsNames returns the TopicsNames field if non-nil, zero value otherwise.

### GetTopicsNamesOk

`func (o *ListLinksResponseData) GetTopicsNamesOk() (*[]string, bool)`

GetTopicsNamesOk returns a tuple with the TopicsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicsNames

`func (o *ListLinksResponseData) SetTopicsNames(v []string)`

SetTopicsNames sets TopicsNames field to given value.

### HasTopicsNames

`func (o *ListLinksResponseData) HasTopicsNames() bool`

HasTopicsNames returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



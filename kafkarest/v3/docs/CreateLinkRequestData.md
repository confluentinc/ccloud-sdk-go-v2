# CreateLinkRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceClusterId** | Pointer to **string** |  | [optional] 
**DestinationClusterId** | Pointer to **string** |  | [optional] 
**RemoteClusterId** | Pointer to **string** | The expected remote cluster ID. | [optional] 
**ClusterLinkId** | Pointer to **string** | The expected cluster link ID. Can be provided when creating the second side of a bidirectional link for validating the link ID is as expected. If it&#39;s not provided, it&#39;s inferred from the remote cluster. | [optional] 
**Configs** | Pointer to [**[]ConfigData**](ConfigData.md) |  | [optional] 

## Methods

### NewCreateLinkRequestData

`func NewCreateLinkRequestData() *CreateLinkRequestData`

NewCreateLinkRequestData instantiates a new CreateLinkRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLinkRequestDataWithDefaults

`func NewCreateLinkRequestDataWithDefaults() *CreateLinkRequestData`

NewCreateLinkRequestDataWithDefaults instantiates a new CreateLinkRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceClusterId

`func (o *CreateLinkRequestData) GetSourceClusterId() string`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### GetSourceClusterIdOk

`func (o *CreateLinkRequestData) GetSourceClusterIdOk() (*string, bool)`

GetSourceClusterIdOk returns a tuple with the SourceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterId

`func (o *CreateLinkRequestData) SetSourceClusterId(v string)`

SetSourceClusterId sets SourceClusterId field to given value.

### HasSourceClusterId

`func (o *CreateLinkRequestData) HasSourceClusterId() bool`

HasSourceClusterId returns a boolean if a field has been set.

### GetDestinationClusterId

`func (o *CreateLinkRequestData) GetDestinationClusterId() string`

GetDestinationClusterId returns the DestinationClusterId field if non-nil, zero value otherwise.

### GetDestinationClusterIdOk

`func (o *CreateLinkRequestData) GetDestinationClusterIdOk() (*string, bool)`

GetDestinationClusterIdOk returns a tuple with the DestinationClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationClusterId

`func (o *CreateLinkRequestData) SetDestinationClusterId(v string)`

SetDestinationClusterId sets DestinationClusterId field to given value.

### HasDestinationClusterId

`func (o *CreateLinkRequestData) HasDestinationClusterId() bool`

HasDestinationClusterId returns a boolean if a field has been set.

### GetRemoteClusterId

`func (o *CreateLinkRequestData) GetRemoteClusterId() string`

GetRemoteClusterId returns the RemoteClusterId field if non-nil, zero value otherwise.

### GetRemoteClusterIdOk

`func (o *CreateLinkRequestData) GetRemoteClusterIdOk() (*string, bool)`

GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClusterId

`func (o *CreateLinkRequestData) SetRemoteClusterId(v string)`

SetRemoteClusterId sets RemoteClusterId field to given value.

### HasRemoteClusterId

`func (o *CreateLinkRequestData) HasRemoteClusterId() bool`

HasRemoteClusterId returns a boolean if a field has been set.

### GetClusterLinkId

`func (o *CreateLinkRequestData) GetClusterLinkId() string`

GetClusterLinkId returns the ClusterLinkId field if non-nil, zero value otherwise.

### GetClusterLinkIdOk

`func (o *CreateLinkRequestData) GetClusterLinkIdOk() (*string, bool)`

GetClusterLinkIdOk returns a tuple with the ClusterLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterLinkId

`func (o *CreateLinkRequestData) SetClusterLinkId(v string)`

SetClusterLinkId sets ClusterLinkId field to given value.

### HasClusterLinkId

`func (o *CreateLinkRequestData) HasClusterLinkId() bool`

HasClusterLinkId returns a boolean if a field has been set.

### GetConfigs

`func (o *CreateLinkRequestData) GetConfigs() []ConfigData`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *CreateLinkRequestData) GetConfigsOk() (*[]ConfigData, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *CreateLinkRequestData) SetConfigs(v []ConfigData)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *CreateLinkRequestData) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



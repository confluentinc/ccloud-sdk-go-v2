# ShareGroupConsumerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**GroupId** | **string** |  | 
**ConsumerId** | **string** |  | 
**ClientId** | **string** |  | 
**Assignments** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewShareGroupConsumerData

`func NewShareGroupConsumerData(kind string, metadata ResourceMetadata, clusterId string, groupId string, consumerId string, clientId string, assignments Relationship, ) *ShareGroupConsumerData`

NewShareGroupConsumerData instantiates a new ShareGroupConsumerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareGroupConsumerDataWithDefaults

`func NewShareGroupConsumerDataWithDefaults() *ShareGroupConsumerData`

NewShareGroupConsumerDataWithDefaults instantiates a new ShareGroupConsumerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ShareGroupConsumerData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ShareGroupConsumerData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ShareGroupConsumerData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ShareGroupConsumerData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShareGroupConsumerData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShareGroupConsumerData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *ShareGroupConsumerData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ShareGroupConsumerData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ShareGroupConsumerData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetGroupId

`func (o *ShareGroupConsumerData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ShareGroupConsumerData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ShareGroupConsumerData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetConsumerId

`func (o *ShareGroupConsumerData) GetConsumerId() string`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *ShareGroupConsumerData) GetConsumerIdOk() (*string, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *ShareGroupConsumerData) SetConsumerId(v string)`

SetConsumerId sets ConsumerId field to given value.


### GetClientId

`func (o *ShareGroupConsumerData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ShareGroupConsumerData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ShareGroupConsumerData) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetAssignments

`func (o *ShareGroupConsumerData) GetAssignments() Relationship`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *ShareGroupConsumerData) GetAssignmentsOk() (*Relationship, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *ShareGroupConsumerData) SetAssignments(v Relationship)`

SetAssignments sets Assignments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



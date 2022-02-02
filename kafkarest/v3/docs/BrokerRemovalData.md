# BrokerRemovalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**BrokerTask** | [**Relationship**](Relationship.md) |  | 
**Broker** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewBrokerRemovalData

`func NewBrokerRemovalData(kind string, metadata ResourceMetadata, clusterId string, brokerId int32, brokerTask Relationship, broker Relationship, ) *BrokerRemovalData`

NewBrokerRemovalData instantiates a new BrokerRemovalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerRemovalDataWithDefaults

`func NewBrokerRemovalDataWithDefaults() *BrokerRemovalData`

NewBrokerRemovalDataWithDefaults instantiates a new BrokerRemovalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BrokerRemovalData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrokerRemovalData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrokerRemovalData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *BrokerRemovalData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BrokerRemovalData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BrokerRemovalData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *BrokerRemovalData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BrokerRemovalData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BrokerRemovalData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *BrokerRemovalData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerRemovalData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerRemovalData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetBrokerTask

`func (o *BrokerRemovalData) GetBrokerTask() Relationship`

GetBrokerTask returns the BrokerTask field if non-nil, zero value otherwise.

### GetBrokerTaskOk

`func (o *BrokerRemovalData) GetBrokerTaskOk() (*Relationship, bool)`

GetBrokerTaskOk returns a tuple with the BrokerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerTask

`func (o *BrokerRemovalData) SetBrokerTask(v Relationship)`

SetBrokerTask sets BrokerTask field to given value.


### GetBroker

`func (o *BrokerRemovalData) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *BrokerRemovalData) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *BrokerRemovalData) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



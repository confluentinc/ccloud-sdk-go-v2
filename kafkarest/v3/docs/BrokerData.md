# BrokerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | 
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**Host** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Rack** | Pointer to **NullableString** |  | [optional] 
**Configs** | [**Relationship**](Relationship.md) |  | 
**PartitionReplicas** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewBrokerData

`func NewBrokerData(kind string, metadata ResourceMetadata, clusterId string, brokerId int32, configs Relationship, partitionReplicas Relationship, ) *BrokerData`

NewBrokerData instantiates a new BrokerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerDataWithDefaults

`func NewBrokerDataWithDefaults() *BrokerData`

NewBrokerDataWithDefaults instantiates a new BrokerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *BrokerData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrokerData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrokerData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *BrokerData) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BrokerData) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BrokerData) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetClusterId

`func (o *BrokerData) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BrokerData) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BrokerData) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *BrokerData) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerData) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerData) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetHost

`func (o *BrokerData) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *BrokerData) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *BrokerData) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *BrokerData) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *BrokerData) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *BrokerData) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *BrokerData) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BrokerData) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BrokerData) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *BrokerData) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *BrokerData) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *BrokerData) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetRack

`func (o *BrokerData) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *BrokerData) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *BrokerData) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *BrokerData) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *BrokerData) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *BrokerData) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetConfigs

`func (o *BrokerData) GetConfigs() Relationship`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *BrokerData) GetConfigsOk() (*Relationship, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *BrokerData) SetConfigs(v Relationship)`

SetConfigs sets Configs field to given value.


### GetPartitionReplicas

`func (o *BrokerData) GetPartitionReplicas() Relationship`

GetPartitionReplicas returns the PartitionReplicas field if non-nil, zero value otherwise.

### GetPartitionReplicasOk

`func (o *BrokerData) GetPartitionReplicasOk() (*Relationship, bool)`

GetPartitionReplicasOk returns a tuple with the PartitionReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionReplicas

`func (o *BrokerData) SetPartitionReplicas(v Relationship)`

SetPartitionReplicas sets PartitionReplicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



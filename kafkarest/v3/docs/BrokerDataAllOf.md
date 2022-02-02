# BrokerDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**Host** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Rack** | Pointer to **NullableString** |  | [optional] 
**Configs** | [**Relationship**](Relationship.md) |  | 
**PartitionReplicas** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewBrokerDataAllOf

`func NewBrokerDataAllOf(clusterId string, brokerId int32, configs Relationship, partitionReplicas Relationship, ) *BrokerDataAllOf`

NewBrokerDataAllOf instantiates a new BrokerDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerDataAllOfWithDefaults

`func NewBrokerDataAllOfWithDefaults() *BrokerDataAllOf`

NewBrokerDataAllOfWithDefaults instantiates a new BrokerDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *BrokerDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BrokerDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BrokerDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *BrokerDataAllOf) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerDataAllOf) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerDataAllOf) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetHost

`func (o *BrokerDataAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *BrokerDataAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *BrokerDataAllOf) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *BrokerDataAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *BrokerDataAllOf) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *BrokerDataAllOf) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *BrokerDataAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BrokerDataAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BrokerDataAllOf) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *BrokerDataAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *BrokerDataAllOf) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *BrokerDataAllOf) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetRack

`func (o *BrokerDataAllOf) GetRack() string`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *BrokerDataAllOf) GetRackOk() (*string, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *BrokerDataAllOf) SetRack(v string)`

SetRack sets Rack field to given value.

### HasRack

`func (o *BrokerDataAllOf) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *BrokerDataAllOf) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *BrokerDataAllOf) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetConfigs

`func (o *BrokerDataAllOf) GetConfigs() Relationship`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *BrokerDataAllOf) GetConfigsOk() (*Relationship, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *BrokerDataAllOf) SetConfigs(v Relationship)`

SetConfigs sets Configs field to given value.


### GetPartitionReplicas

`func (o *BrokerDataAllOf) GetPartitionReplicas() Relationship`

GetPartitionReplicas returns the PartitionReplicas field if non-nil, zero value otherwise.

### GetPartitionReplicasOk

`func (o *BrokerDataAllOf) GetPartitionReplicasOk() (*Relationship, bool)`

GetPartitionReplicasOk returns a tuple with the PartitionReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionReplicas

`func (o *BrokerDataAllOf) SetPartitionReplicas(v Relationship)`

SetPartitionReplicas sets PartitionReplicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



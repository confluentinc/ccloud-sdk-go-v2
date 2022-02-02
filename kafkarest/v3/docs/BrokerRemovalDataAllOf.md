# BrokerRemovalDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**BrokerId** | **int32** |  | 
**BrokerTask** | [**Relationship**](Relationship.md) |  | 
**Broker** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewBrokerRemovalDataAllOf

`func NewBrokerRemovalDataAllOf(clusterId string, brokerId int32, brokerTask Relationship, broker Relationship, ) *BrokerRemovalDataAllOf`

NewBrokerRemovalDataAllOf instantiates a new BrokerRemovalDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerRemovalDataAllOfWithDefaults

`func NewBrokerRemovalDataAllOfWithDefaults() *BrokerRemovalDataAllOf`

NewBrokerRemovalDataAllOfWithDefaults instantiates a new BrokerRemovalDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *BrokerRemovalDataAllOf) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BrokerRemovalDataAllOf) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BrokerRemovalDataAllOf) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetBrokerId

`func (o *BrokerRemovalDataAllOf) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerRemovalDataAllOf) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerRemovalDataAllOf) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetBrokerTask

`func (o *BrokerRemovalDataAllOf) GetBrokerTask() Relationship`

GetBrokerTask returns the BrokerTask field if non-nil, zero value otherwise.

### GetBrokerTaskOk

`func (o *BrokerRemovalDataAllOf) GetBrokerTaskOk() (*Relationship, bool)`

GetBrokerTaskOk returns a tuple with the BrokerTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerTask

`func (o *BrokerRemovalDataAllOf) SetBrokerTask(v Relationship)`

SetBrokerTask sets BrokerTask field to given value.


### GetBroker

`func (o *BrokerRemovalDataAllOf) GetBroker() Relationship`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *BrokerRemovalDataAllOf) GetBrokerOk() (*Relationship, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *BrokerRemovalDataAllOf) SetBroker(v Relationship)`

SetBroker sets Broker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



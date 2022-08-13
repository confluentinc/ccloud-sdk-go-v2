# KsqldbcmV2ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpEndpoint** | Pointer to **string** | The dataplane endpoint of the ksqlDB cluster. | [optional] [readonly] 
**Phase** | **string** | Status of the ksqlDB cluster. | [readonly] 
**IsPaused** | **bool** | Tells you if the cluster has been paused | [readonly] 
**Storage** | **int32** | Amount of storage (in GB) provisioned to this cluster | [readonly] 
**TopicPrefix** | Pointer to **string** | Topic name prefix used by this ksqlDB cluster. Used to assign ACLs for this ksqlDB cluster to use. | [optional] [readonly] 

## Methods

### NewKsqldbcmV2ClusterStatus

`func NewKsqldbcmV2ClusterStatus(phase string, isPaused bool, storage int32, ) *KsqldbcmV2ClusterStatus`

NewKsqldbcmV2ClusterStatus instantiates a new KsqldbcmV2ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKsqldbcmV2ClusterStatusWithDefaults

`func NewKsqldbcmV2ClusterStatusWithDefaults() *KsqldbcmV2ClusterStatus`

NewKsqldbcmV2ClusterStatusWithDefaults instantiates a new KsqldbcmV2ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpEndpoint

`func (o *KsqldbcmV2ClusterStatus) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *KsqldbcmV2ClusterStatus) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *KsqldbcmV2ClusterStatus) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *KsqldbcmV2ClusterStatus) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetPhase

`func (o *KsqldbcmV2ClusterStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KsqldbcmV2ClusterStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KsqldbcmV2ClusterStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetIsPaused

`func (o *KsqldbcmV2ClusterStatus) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *KsqldbcmV2ClusterStatus) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *KsqldbcmV2ClusterStatus) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.


### GetStorage

`func (o *KsqldbcmV2ClusterStatus) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *KsqldbcmV2ClusterStatus) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *KsqldbcmV2ClusterStatus) SetStorage(v int32)`

SetStorage sets Storage field to given value.


### GetTopicPrefix

`func (o *KsqldbcmV2ClusterStatus) GetTopicPrefix() string`

GetTopicPrefix returns the TopicPrefix field if non-nil, zero value otherwise.

### GetTopicPrefixOk

`func (o *KsqldbcmV2ClusterStatus) GetTopicPrefixOk() (*string, bool)`

GetTopicPrefixOk returns a tuple with the TopicPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicPrefix

`func (o *KsqldbcmV2ClusterStatus) SetTopicPrefix(v string)`

SetTopicPrefix sets TopicPrefix field to given value.

### HasTopicPrefix

`func (o *KsqldbcmV2ClusterStatus) HasTopicPrefix() bool`

HasTopicPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



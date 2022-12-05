# ClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSchemas** | Pointer to **int32** | Maximum number of registered schemas allowed | [optional] 
**MaxRequestsPerSec** | Pointer to **int32** | Maximum number of allowed requests per second | [optional] 

## Methods

### NewClusterConfig

`func NewClusterConfig() *ClusterConfig`

NewClusterConfig instantiates a new ClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigWithDefaults

`func NewClusterConfigWithDefaults() *ClusterConfig`

NewClusterConfigWithDefaults instantiates a new ClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSchemas

`func (o *ClusterConfig) GetMaxSchemas() int32`

GetMaxSchemas returns the MaxSchemas field if non-nil, zero value otherwise.

### GetMaxSchemasOk

`func (o *ClusterConfig) GetMaxSchemasOk() (*int32, bool)`

GetMaxSchemasOk returns a tuple with the MaxSchemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSchemas

`func (o *ClusterConfig) SetMaxSchemas(v int32)`

SetMaxSchemas sets MaxSchemas field to given value.

### HasMaxSchemas

`func (o *ClusterConfig) HasMaxSchemas() bool`

HasMaxSchemas returns a boolean if a field has been set.

### GetMaxRequestsPerSec

`func (o *ClusterConfig) GetMaxRequestsPerSec() int32`

GetMaxRequestsPerSec returns the MaxRequestsPerSec field if non-nil, zero value otherwise.

### GetMaxRequestsPerSecOk

`func (o *ClusterConfig) GetMaxRequestsPerSecOk() (*int32, bool)`

GetMaxRequestsPerSecOk returns a tuple with the MaxRequestsPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestsPerSec

`func (o *ClusterConfig) SetMaxRequestsPerSec(v int32)`

SetMaxRequestsPerSec sets MaxRequestsPerSec field to given value.

### HasMaxRequestsPerSec

`func (o *ClusterConfig) HasMaxRequestsPerSec() bool`

HasMaxRequestsPerSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



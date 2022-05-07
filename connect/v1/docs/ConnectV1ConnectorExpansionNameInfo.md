# ConnectV1ConnectorExpansionNameInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the connector. | [optional] 
**Config** | Pointer to **map[string]string** | Configuration parameters for the connector. These configurations are the minimum set of key-value pairs (KVP) which are used to define how the connector connects Kafka to the external system. Some of these KVPs are common to all the connectors, such as connection parameters to Kafka, connector metadata, etc. The list of common connector configurations is as follows    - cloud.environment   - cloud.provider   - connector.class   - kafka.api.key   - kafka.api.secret   - kafka.endpoint   - kafka.region   - name  For example, a connector like &#x60;GcsSink&#x60; would have additional parameters such as &#x60;gcs.bucket.name&#x60;, &#x60;flush.size&#x60;, etc. | [optional] 

## Methods

### NewConnectV1ConnectorExpansionNameInfo

`func NewConnectV1ConnectorExpansionNameInfo() *ConnectV1ConnectorExpansionNameInfo`

NewConnectV1ConnectorExpansionNameInfo instantiates a new ConnectV1ConnectorExpansionNameInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectV1ConnectorExpansionNameInfoWithDefaults

`func NewConnectV1ConnectorExpansionNameInfoWithDefaults() *ConnectV1ConnectorExpansionNameInfo`

NewConnectV1ConnectorExpansionNameInfoWithDefaults instantiates a new ConnectV1ConnectorExpansionNameInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectV1ConnectorExpansionNameInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectV1ConnectorExpansionNameInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectV1ConnectorExpansionNameInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectV1ConnectorExpansionNameInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *ConnectV1ConnectorExpansionNameInfo) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ConnectV1ConnectorExpansionNameInfo) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ConnectV1ConnectorExpansionNameInfo) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ConnectV1ConnectorExpansionNameInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



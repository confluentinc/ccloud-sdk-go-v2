# CclV1KafkaDestinationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The destination where Custom Code Logging is sent. | 
**ClusterId** | **string** | The kafka cluster id where Custom Code Logging is sent. | 
**Topic** | **string** | The kafka topic where Custom Code Logging is sent. | 
**LogLevel** | Pointer to **string** | Minimum log level for Custom Code Logging. | [optional] [default to "INFO"]

## Methods

### NewCclV1KafkaDestinationSettings

`func NewCclV1KafkaDestinationSettings(kind string, clusterId string, topic string, ) *CclV1KafkaDestinationSettings`

NewCclV1KafkaDestinationSettings instantiates a new CclV1KafkaDestinationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCclV1KafkaDestinationSettingsWithDefaults

`func NewCclV1KafkaDestinationSettingsWithDefaults() *CclV1KafkaDestinationSettings`

NewCclV1KafkaDestinationSettingsWithDefaults instantiates a new CclV1KafkaDestinationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CclV1KafkaDestinationSettings) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CclV1KafkaDestinationSettings) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CclV1KafkaDestinationSettings) SetKind(v string)`

SetKind sets Kind field to given value.


### GetClusterId

`func (o *CclV1KafkaDestinationSettings) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CclV1KafkaDestinationSettings) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CclV1KafkaDestinationSettings) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetTopic

`func (o *CclV1KafkaDestinationSettings) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *CclV1KafkaDestinationSettings) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *CclV1KafkaDestinationSettings) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetLogLevel

`func (o *CclV1KafkaDestinationSettings) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *CclV1KafkaDestinationSettings) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *CclV1KafkaDestinationSettings) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *CclV1KafkaDestinationSettings) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.


### AsCclV1CustomCodeLoggingDestinationSettingsOneOf

`func (s *CclV1KafkaDestinationSettings) AsCclV1CustomCodeLoggingDestinationSettingsOneOf() CclV1CustomCodeLoggingDestinationSettingsOneOf`

Convenience method to wrap this instance of CclV1KafkaDestinationSettings in CclV1CustomCodeLoggingDestinationSettingsOneOf

### AsCclV1CustomCodeLoggingUpdateDestinationSettingsOneOf

`func (s *CclV1KafkaDestinationSettings) AsCclV1CustomCodeLoggingUpdateDestinationSettingsOneOf() CclV1CustomCodeLoggingUpdateDestinationSettingsOneOf`

Convenience method to wrap this instance of CclV1KafkaDestinationSettings in CclV1CustomCodeLoggingUpdateDestinationSettingsOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



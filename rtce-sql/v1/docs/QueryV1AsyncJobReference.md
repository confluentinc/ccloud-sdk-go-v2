# QueryV1AsyncJobReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatementId** | **string** | Unique identifier generated to trace the decoupled tracking process. | 
**EnvironmentId** | **string** | Confluent Cloud environment ID (echoed back from &#x60;catalog_name&#x60;). | 
**KafkaClusterId** | **string** | Kafka cluster ID (echoed back from &#x60;database_name&#x60;). | 
**TopicName** | **string** | Name of the table/topic the statement queried. | 

## Methods

### NewQueryV1AsyncJobReference

`func NewQueryV1AsyncJobReference(statementId string, environmentId string, kafkaClusterId string, topicName string, ) *QueryV1AsyncJobReference`

NewQueryV1AsyncJobReference instantiates a new QueryV1AsyncJobReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1AsyncJobReferenceWithDefaults

`func NewQueryV1AsyncJobReferenceWithDefaults() *QueryV1AsyncJobReference`

NewQueryV1AsyncJobReferenceWithDefaults instantiates a new QueryV1AsyncJobReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatementId

`func (o *QueryV1AsyncJobReference) GetStatementId() string`

GetStatementId returns the StatementId field if non-nil, zero value otherwise.

### GetStatementIdOk

`func (o *QueryV1AsyncJobReference) GetStatementIdOk() (*string, bool)`

GetStatementIdOk returns a tuple with the StatementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementId

`func (o *QueryV1AsyncJobReference) SetStatementId(v string)`

SetStatementId sets StatementId field to given value.


### GetEnvironmentId

`func (o *QueryV1AsyncJobReference) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *QueryV1AsyncJobReference) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *QueryV1AsyncJobReference) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetKafkaClusterId

`func (o *QueryV1AsyncJobReference) GetKafkaClusterId() string`

GetKafkaClusterId returns the KafkaClusterId field if non-nil, zero value otherwise.

### GetKafkaClusterIdOk

`func (o *QueryV1AsyncJobReference) GetKafkaClusterIdOk() (*string, bool)`

GetKafkaClusterIdOk returns a tuple with the KafkaClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusterId

`func (o *QueryV1AsyncJobReference) SetKafkaClusterId(v string)`

SetKafkaClusterId sets KafkaClusterId field to given value.


### GetTopicName

`func (o *QueryV1AsyncJobReference) GetTopicName() string`

GetTopicName returns the TopicName field if non-nil, zero value otherwise.

### GetTopicNameOk

`func (o *QueryV1AsyncJobReference) GetTopicNameOk() (*string, bool)`

GetTopicNameOk returns a tuple with the TopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicName

`func (o *QueryV1AsyncJobReference) SetTopicName(v string)`

SetTopicName sets TopicName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CmkV2Endpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaBootstrapEndpoint** | **string** | The bootstrap endpoint used by Kafka clients to connect to the cluster.  | 
**HttpEndpoint** | **string** | The REST endpoint for the Kafka cluster.  | 
**ConnectionType** | **string** | The type of connection used for the endpoint.  | 

## Methods

### NewCmkV2Endpoints

`func NewCmkV2Endpoints(kafkaBootstrapEndpoint string, httpEndpoint string, connectionType string, ) *CmkV2Endpoints`

NewCmkV2Endpoints instantiates a new CmkV2Endpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2EndpointsWithDefaults

`func NewCmkV2EndpointsWithDefaults() *CmkV2Endpoints`

NewCmkV2EndpointsWithDefaults instantiates a new CmkV2Endpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKafkaBootstrapEndpoint

`func (o *CmkV2Endpoints) GetKafkaBootstrapEndpoint() string`

GetKafkaBootstrapEndpoint returns the KafkaBootstrapEndpoint field if non-nil, zero value otherwise.

### GetKafkaBootstrapEndpointOk

`func (o *CmkV2Endpoints) GetKafkaBootstrapEndpointOk() (*string, bool)`

GetKafkaBootstrapEndpointOk returns a tuple with the KafkaBootstrapEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaBootstrapEndpoint

`func (o *CmkV2Endpoints) SetKafkaBootstrapEndpoint(v string)`

SetKafkaBootstrapEndpoint sets KafkaBootstrapEndpoint field to given value.


### GetHttpEndpoint

`func (o *CmkV2Endpoints) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *CmkV2Endpoints) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *CmkV2Endpoints) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.


### GetConnectionType

`func (o *CmkV2Endpoints) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CmkV2Endpoints) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CmkV2Endpoints) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



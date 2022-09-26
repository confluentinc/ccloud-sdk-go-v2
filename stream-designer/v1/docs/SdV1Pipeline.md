# SdV1Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**OrgId** | Pointer to **string** | The organization that owns this resource. | [optional] 
**Name** | Pointer to **string** | The name of this pipeline. | [optional] 
**Description** | Pointer to **string** | The description of this pipeline. | [optional] 
**EnvironmentId** | Pointer to **string** | The unique identifier of the environment this pipeline belongs to. | [optional] 
**KafkaClusterId** | Pointer to **string** | The unique identifier for the Kafka cluster this pipeline uses. | [optional] 
**KsqlId** | Pointer to **string** | The unique identifier of the ksqlDB application this pipeline uses. | [optional] 
**SchemaRegistryId** | Pointer to **string** | The unique identifier of the Schema Registry this pipeline uses. | [optional] 
**State** | Pointer to **string** | The current state of the pipeline. | [optional] 
**Activated** | Pointer to **bool** | The desired state of the pipeline. | [optional] 
**KafkaClusterEndpoint** | Pointer to **string** | The endpoint URL of the kafka cluster this pipeline uses. | [optional] 
**KsqlEndpoint** | Pointer to **string** | The endpoint URL of the ksqlDB application this pipeline uses. | [optional] 
**ConnectEndpoint** | Pointer to **string** | The endpoint URL of the CCloud Connect service this pipeline uses. | [optional] 
**SchemaRegistryEndpoint** | Pointer to **string** | The endpoint URL of the Schema Registry this pipeline uses. | [optional] 

## Methods

### NewSdV1Pipeline

`func NewSdV1Pipeline() *SdV1Pipeline`

NewSdV1Pipeline instantiates a new SdV1Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdV1PipelineWithDefaults

`func NewSdV1PipelineWithDefaults() *SdV1Pipeline`

NewSdV1PipelineWithDefaults instantiates a new SdV1Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SdV1Pipeline) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SdV1Pipeline) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SdV1Pipeline) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SdV1Pipeline) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SdV1Pipeline) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SdV1Pipeline) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SdV1Pipeline) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SdV1Pipeline) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *SdV1Pipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SdV1Pipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SdV1Pipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SdV1Pipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *SdV1Pipeline) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SdV1Pipeline) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SdV1Pipeline) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SdV1Pipeline) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrgId

`func (o *SdV1Pipeline) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *SdV1Pipeline) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *SdV1Pipeline) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *SdV1Pipeline) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *SdV1Pipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdV1Pipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdV1Pipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdV1Pipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SdV1Pipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdV1Pipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdV1Pipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdV1Pipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *SdV1Pipeline) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SdV1Pipeline) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SdV1Pipeline) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *SdV1Pipeline) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetKafkaClusterId

`func (o *SdV1Pipeline) GetKafkaClusterId() string`

GetKafkaClusterId returns the KafkaClusterId field if non-nil, zero value otherwise.

### GetKafkaClusterIdOk

`func (o *SdV1Pipeline) GetKafkaClusterIdOk() (*string, bool)`

GetKafkaClusterIdOk returns a tuple with the KafkaClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusterId

`func (o *SdV1Pipeline) SetKafkaClusterId(v string)`

SetKafkaClusterId sets KafkaClusterId field to given value.

### HasKafkaClusterId

`func (o *SdV1Pipeline) HasKafkaClusterId() bool`

HasKafkaClusterId returns a boolean if a field has been set.

### GetKsqlId

`func (o *SdV1Pipeline) GetKsqlId() string`

GetKsqlId returns the KsqlId field if non-nil, zero value otherwise.

### GetKsqlIdOk

`func (o *SdV1Pipeline) GetKsqlIdOk() (*string, bool)`

GetKsqlIdOk returns a tuple with the KsqlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKsqlId

`func (o *SdV1Pipeline) SetKsqlId(v string)`

SetKsqlId sets KsqlId field to given value.

### HasKsqlId

`func (o *SdV1Pipeline) HasKsqlId() bool`

HasKsqlId returns a boolean if a field has been set.

### GetSchemaRegistryId

`func (o *SdV1Pipeline) GetSchemaRegistryId() string`

GetSchemaRegistryId returns the SchemaRegistryId field if non-nil, zero value otherwise.

### GetSchemaRegistryIdOk

`func (o *SdV1Pipeline) GetSchemaRegistryIdOk() (*string, bool)`

GetSchemaRegistryIdOk returns a tuple with the SchemaRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistryId

`func (o *SdV1Pipeline) SetSchemaRegistryId(v string)`

SetSchemaRegistryId sets SchemaRegistryId field to given value.

### HasSchemaRegistryId

`func (o *SdV1Pipeline) HasSchemaRegistryId() bool`

HasSchemaRegistryId returns a boolean if a field has been set.

### GetState

`func (o *SdV1Pipeline) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SdV1Pipeline) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SdV1Pipeline) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SdV1Pipeline) HasState() bool`

HasState returns a boolean if a field has been set.

### GetActivated

`func (o *SdV1Pipeline) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *SdV1Pipeline) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *SdV1Pipeline) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *SdV1Pipeline) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetKafkaClusterEndpoint

`func (o *SdV1Pipeline) GetKafkaClusterEndpoint() string`

GetKafkaClusterEndpoint returns the KafkaClusterEndpoint field if non-nil, zero value otherwise.

### GetKafkaClusterEndpointOk

`func (o *SdV1Pipeline) GetKafkaClusterEndpointOk() (*string, bool)`

GetKafkaClusterEndpointOk returns a tuple with the KafkaClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusterEndpoint

`func (o *SdV1Pipeline) SetKafkaClusterEndpoint(v string)`

SetKafkaClusterEndpoint sets KafkaClusterEndpoint field to given value.

### HasKafkaClusterEndpoint

`func (o *SdV1Pipeline) HasKafkaClusterEndpoint() bool`

HasKafkaClusterEndpoint returns a boolean if a field has been set.

### GetKsqlEndpoint

`func (o *SdV1Pipeline) GetKsqlEndpoint() string`

GetKsqlEndpoint returns the KsqlEndpoint field if non-nil, zero value otherwise.

### GetKsqlEndpointOk

`func (o *SdV1Pipeline) GetKsqlEndpointOk() (*string, bool)`

GetKsqlEndpointOk returns a tuple with the KsqlEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKsqlEndpoint

`func (o *SdV1Pipeline) SetKsqlEndpoint(v string)`

SetKsqlEndpoint sets KsqlEndpoint field to given value.

### HasKsqlEndpoint

`func (o *SdV1Pipeline) HasKsqlEndpoint() bool`

HasKsqlEndpoint returns a boolean if a field has been set.

### GetConnectEndpoint

`func (o *SdV1Pipeline) GetConnectEndpoint() string`

GetConnectEndpoint returns the ConnectEndpoint field if non-nil, zero value otherwise.

### GetConnectEndpointOk

`func (o *SdV1Pipeline) GetConnectEndpointOk() (*string, bool)`

GetConnectEndpointOk returns a tuple with the ConnectEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectEndpoint

`func (o *SdV1Pipeline) SetConnectEndpoint(v string)`

SetConnectEndpoint sets ConnectEndpoint field to given value.

### HasConnectEndpoint

`func (o *SdV1Pipeline) HasConnectEndpoint() bool`

HasConnectEndpoint returns a boolean if a field has been set.

### GetSchemaRegistryEndpoint

`func (o *SdV1Pipeline) GetSchemaRegistryEndpoint() string`

GetSchemaRegistryEndpoint returns the SchemaRegistryEndpoint field if non-nil, zero value otherwise.

### GetSchemaRegistryEndpointOk

`func (o *SdV1Pipeline) GetSchemaRegistryEndpointOk() (*string, bool)`

GetSchemaRegistryEndpointOk returns a tuple with the SchemaRegistryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistryEndpoint

`func (o *SdV1Pipeline) SetSchemaRegistryEndpoint(v string)`

SetSchemaRegistryEndpoint sets SchemaRegistryEndpoint field to given value.

### HasSchemaRegistryEndpoint

`func (o *SdV1Pipeline) HasSchemaRegistryEndpoint() bool`

HasSchemaRegistryEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SdV1PipelineUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Name** | Pointer to **string** | The name of this pipeline. | [optional] 
**Description** | Pointer to **string** | The description of this pipeline. | [optional] 
**KsqlId** | Pointer to **string** | The unique identifier of the ksqlDB application this pipeline uses. | [optional] 
**SchemaRegistryId** | Pointer to **string** | The unique identifier of the Schema Registry this pipeline uses. | [optional] 
**KafkaClusterEndpoint** | Pointer to **string** | The endpoint URL of the kafka cluster this pipeline uses. | [optional] 
**KsqlEndpoint** | Pointer to **string** | The endpoint URL of the ksqlDB application this pipeline uses. | [optional] 
**ConnectEndpoint** | Pointer to **string** | The endpoint URL of the CCloud Connect service this pipeline uses. | [optional] 
**SchemaRegistryEndpoint** | Pointer to **string** | The endpoint URL of the Schema Registry this pipeline uses. | [optional] 

## Methods

### NewSdV1PipelineUpdate

`func NewSdV1PipelineUpdate() *SdV1PipelineUpdate`

NewSdV1PipelineUpdate instantiates a new SdV1PipelineUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdV1PipelineUpdateWithDefaults

`func NewSdV1PipelineUpdateWithDefaults() *SdV1PipelineUpdate`

NewSdV1PipelineUpdateWithDefaults instantiates a new SdV1PipelineUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SdV1PipelineUpdate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SdV1PipelineUpdate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SdV1PipelineUpdate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SdV1PipelineUpdate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SdV1PipelineUpdate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SdV1PipelineUpdate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SdV1PipelineUpdate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SdV1PipelineUpdate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *SdV1PipelineUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SdV1PipelineUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SdV1PipelineUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SdV1PipelineUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *SdV1PipelineUpdate) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SdV1PipelineUpdate) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SdV1PipelineUpdate) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SdV1PipelineUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SdV1PipelineUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdV1PipelineUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdV1PipelineUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdV1PipelineUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SdV1PipelineUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SdV1PipelineUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SdV1PipelineUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SdV1PipelineUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKsqlId

`func (o *SdV1PipelineUpdate) GetKsqlId() string`

GetKsqlId returns the KsqlId field if non-nil, zero value otherwise.

### GetKsqlIdOk

`func (o *SdV1PipelineUpdate) GetKsqlIdOk() (*string, bool)`

GetKsqlIdOk returns a tuple with the KsqlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKsqlId

`func (o *SdV1PipelineUpdate) SetKsqlId(v string)`

SetKsqlId sets KsqlId field to given value.

### HasKsqlId

`func (o *SdV1PipelineUpdate) HasKsqlId() bool`

HasKsqlId returns a boolean if a field has been set.

### GetSchemaRegistryId

`func (o *SdV1PipelineUpdate) GetSchemaRegistryId() string`

GetSchemaRegistryId returns the SchemaRegistryId field if non-nil, zero value otherwise.

### GetSchemaRegistryIdOk

`func (o *SdV1PipelineUpdate) GetSchemaRegistryIdOk() (*string, bool)`

GetSchemaRegistryIdOk returns a tuple with the SchemaRegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistryId

`func (o *SdV1PipelineUpdate) SetSchemaRegistryId(v string)`

SetSchemaRegistryId sets SchemaRegistryId field to given value.

### HasSchemaRegistryId

`func (o *SdV1PipelineUpdate) HasSchemaRegistryId() bool`

HasSchemaRegistryId returns a boolean if a field has been set.

### GetKafkaClusterEndpoint

`func (o *SdV1PipelineUpdate) GetKafkaClusterEndpoint() string`

GetKafkaClusterEndpoint returns the KafkaClusterEndpoint field if non-nil, zero value otherwise.

### GetKafkaClusterEndpointOk

`func (o *SdV1PipelineUpdate) GetKafkaClusterEndpointOk() (*string, bool)`

GetKafkaClusterEndpointOk returns a tuple with the KafkaClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaClusterEndpoint

`func (o *SdV1PipelineUpdate) SetKafkaClusterEndpoint(v string)`

SetKafkaClusterEndpoint sets KafkaClusterEndpoint field to given value.

### HasKafkaClusterEndpoint

`func (o *SdV1PipelineUpdate) HasKafkaClusterEndpoint() bool`

HasKafkaClusterEndpoint returns a boolean if a field has been set.

### GetKsqlEndpoint

`func (o *SdV1PipelineUpdate) GetKsqlEndpoint() string`

GetKsqlEndpoint returns the KsqlEndpoint field if non-nil, zero value otherwise.

### GetKsqlEndpointOk

`func (o *SdV1PipelineUpdate) GetKsqlEndpointOk() (*string, bool)`

GetKsqlEndpointOk returns a tuple with the KsqlEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKsqlEndpoint

`func (o *SdV1PipelineUpdate) SetKsqlEndpoint(v string)`

SetKsqlEndpoint sets KsqlEndpoint field to given value.

### HasKsqlEndpoint

`func (o *SdV1PipelineUpdate) HasKsqlEndpoint() bool`

HasKsqlEndpoint returns a boolean if a field has been set.

### GetConnectEndpoint

`func (o *SdV1PipelineUpdate) GetConnectEndpoint() string`

GetConnectEndpoint returns the ConnectEndpoint field if non-nil, zero value otherwise.

### GetConnectEndpointOk

`func (o *SdV1PipelineUpdate) GetConnectEndpointOk() (*string, bool)`

GetConnectEndpointOk returns a tuple with the ConnectEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectEndpoint

`func (o *SdV1PipelineUpdate) SetConnectEndpoint(v string)`

SetConnectEndpoint sets ConnectEndpoint field to given value.

### HasConnectEndpoint

`func (o *SdV1PipelineUpdate) HasConnectEndpoint() bool`

HasConnectEndpoint returns a boolean if a field has been set.

### GetSchemaRegistryEndpoint

`func (o *SdV1PipelineUpdate) GetSchemaRegistryEndpoint() string`

GetSchemaRegistryEndpoint returns the SchemaRegistryEndpoint field if non-nil, zero value otherwise.

### GetSchemaRegistryEndpointOk

`func (o *SdV1PipelineUpdate) GetSchemaRegistryEndpointOk() (*string, bool)`

GetSchemaRegistryEndpointOk returns a tuple with the SchemaRegistryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistryEndpoint

`func (o *SdV1PipelineUpdate) SetSchemaRegistryEndpoint(v string)`

SetSchemaRegistryEndpoint sets SchemaRegistryEndpoint field to given value.

### HasSchemaRegistryEndpoint

`func (o *SdV1PipelineUpdate) HasSchemaRegistryEndpoint() bool`

HasSchemaRegistryEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



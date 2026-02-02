# EndpointV1Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Cloud** | Pointer to **string** | The cloud service provider. | [optional] [readonly] 
**Region** | Pointer to **string** | The cloud service provider region in which the resource is located. | [optional] [readonly] 
**Service** | Pointer to **string** | The Confluent Cloud service. | [optional] [readonly] 
**IsPrivate** | Pointer to **bool** | Whether the endpoint is private (true) or public (false). | [optional] [readonly] 
**ConnectionType** | Pointer to **string** | The network connection type. | [optional] [readonly] 
**Endpoint** | Pointer to **string** | The endpoint URL or address. | [optional] [readonly] 
**EndpointType** | Pointer to **string** | The endpoint type enum values: * &#x60;REST&#x60; - REST API endpoint for HTTP/HTTPS access, used by Kafka/Flink/Schema Registry services. * &#x60;BOOTSTRAP&#x60; - Kafka native protocol bootstrap servers for direct client connections, used by Kafka only. * &#x60;LANGUAGE_SERVICE&#x60; - Flink language service endpoint for SQL/Table API, used by Flink only.  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 
**Resource** | Pointer to [**TypedEnvScopedObjectReference**](TypedEnvScopedObjectReference.md) | The resource associated with the endpoint. The resource can be one of Kafka Cluster ID (example: lkc-12345), or Schema Registry Cluster ID (example: lsrc-12345). May be null or omitted if not associated with a resource.  | [optional] 
**Gateway** | Pointer to [**ObjectReference**](ObjectReference.md) | The gateway to which this belongs. | [optional] 
**AccessPoint** | Pointer to [**ObjectReference**](ObjectReference.md) | The access_point to which this belongs. | [optional] 

## Methods

### NewEndpointV1Endpoint

`func NewEndpointV1Endpoint() *EndpointV1Endpoint`

NewEndpointV1Endpoint instantiates a new EndpointV1Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointV1EndpointWithDefaults

`func NewEndpointV1EndpointWithDefaults() *EndpointV1Endpoint`

NewEndpointV1EndpointWithDefaults instantiates a new EndpointV1Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *EndpointV1Endpoint) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EndpointV1Endpoint) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EndpointV1Endpoint) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *EndpointV1Endpoint) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *EndpointV1Endpoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EndpointV1Endpoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EndpointV1Endpoint) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EndpointV1Endpoint) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *EndpointV1Endpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointV1Endpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointV1Endpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointV1Endpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *EndpointV1Endpoint) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EndpointV1Endpoint) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EndpointV1Endpoint) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EndpointV1Endpoint) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCloud

`func (o *EndpointV1Endpoint) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *EndpointV1Endpoint) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *EndpointV1Endpoint) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *EndpointV1Endpoint) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *EndpointV1Endpoint) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EndpointV1Endpoint) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EndpointV1Endpoint) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EndpointV1Endpoint) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetService

`func (o *EndpointV1Endpoint) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *EndpointV1Endpoint) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *EndpointV1Endpoint) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *EndpointV1Endpoint) HasService() bool`

HasService returns a boolean if a field has been set.

### GetIsPrivate

`func (o *EndpointV1Endpoint) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *EndpointV1Endpoint) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *EndpointV1Endpoint) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *EndpointV1Endpoint) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetConnectionType

`func (o *EndpointV1Endpoint) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *EndpointV1Endpoint) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *EndpointV1Endpoint) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *EndpointV1Endpoint) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetEndpoint

`func (o *EndpointV1Endpoint) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EndpointV1Endpoint) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EndpointV1Endpoint) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *EndpointV1Endpoint) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEndpointType

`func (o *EndpointV1Endpoint) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *EndpointV1Endpoint) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *EndpointV1Endpoint) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *EndpointV1Endpoint) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetEnvironment

`func (o *EndpointV1Endpoint) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EndpointV1Endpoint) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EndpointV1Endpoint) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EndpointV1Endpoint) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetResource

`func (o *EndpointV1Endpoint) GetResource() TypedEnvScopedObjectReference`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EndpointV1Endpoint) GetResourceOk() (*TypedEnvScopedObjectReference, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EndpointV1Endpoint) SetResource(v TypedEnvScopedObjectReference)`

SetResource sets Resource field to given value.

### HasResource

`func (o *EndpointV1Endpoint) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetGateway

`func (o *EndpointV1Endpoint) GetGateway() ObjectReference`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *EndpointV1Endpoint) GetGatewayOk() (*ObjectReference, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *EndpointV1Endpoint) SetGateway(v ObjectReference)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *EndpointV1Endpoint) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetAccessPoint

`func (o *EndpointV1Endpoint) GetAccessPoint() ObjectReference`

GetAccessPoint returns the AccessPoint field if non-nil, zero value otherwise.

### GetAccessPointOk

`func (o *EndpointV1Endpoint) GetAccessPointOk() (*ObjectReference, bool)`

GetAccessPointOk returns a tuple with the AccessPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPoint

`func (o *EndpointV1Endpoint) SetAccessPoint(v ObjectReference)`

SetAccessPoint sets AccessPoint field to given value.

### HasAccessPoint

`func (o *EndpointV1Endpoint) HasAccessPoint() bool`

HasAccessPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



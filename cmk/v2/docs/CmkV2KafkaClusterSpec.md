# CmkV2KafkaClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The name of the cluster. | [optional] 
**Availability** | Pointer to **string** | The availability zone configuration of the cluster. | [optional] [default to "SINGLE_ZONE"]
**Placement** | Pointer to [**CmkV2KafkaClusterSpecPlacementOneOf**](CmkV2KafkaClusterSpecPlacementOneOf.md) | The geographical location where to place the Kafka cluster. | [optional] 
**ClusterType** | Pointer to **string** | The type of cluster.  Note: Clusters can only be upgraded from BASIC to STANDARD but cannot be downgraded from STANDARD to BASIC.  | [optional] [default to "BASIC"]
**KafkaBootstrapEndpoint** | Pointer to **string** | The bootstrap endpoint used by Kafka clients to connect to the cluster. | [optional] [readonly] 
**HttpEndpoint** | Pointer to **string** | The cluster HTTP request URL. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewCmkV2KafkaClusterSpec

`func NewCmkV2KafkaClusterSpec() *CmkV2KafkaClusterSpec`

NewCmkV2KafkaClusterSpec instantiates a new CmkV2KafkaClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCmkV2KafkaClusterSpecWithDefaults

`func NewCmkV2KafkaClusterSpecWithDefaults() *CmkV2KafkaClusterSpec`

NewCmkV2KafkaClusterSpecWithDefaults instantiates a new CmkV2KafkaClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CmkV2KafkaClusterSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CmkV2KafkaClusterSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CmkV2KafkaClusterSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CmkV2KafkaClusterSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAvailability

`func (o *CmkV2KafkaClusterSpec) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *CmkV2KafkaClusterSpec) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *CmkV2KafkaClusterSpec) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *CmkV2KafkaClusterSpec) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetPlacement

`func (o *CmkV2KafkaClusterSpec) GetPlacement() CmkV2KafkaClusterSpecPlacementOneOf`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *CmkV2KafkaClusterSpec) GetPlacementOk() (*CmkV2KafkaClusterSpecPlacementOneOf, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *CmkV2KafkaClusterSpec) SetPlacement(v CmkV2KafkaClusterSpecPlacementOneOf)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *CmkV2KafkaClusterSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetClusterType

`func (o *CmkV2KafkaClusterSpec) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *CmkV2KafkaClusterSpec) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *CmkV2KafkaClusterSpec) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *CmkV2KafkaClusterSpec) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetKafkaBootstrapEndpoint

`func (o *CmkV2KafkaClusterSpec) GetKafkaBootstrapEndpoint() string`

GetKafkaBootstrapEndpoint returns the KafkaBootstrapEndpoint field if non-nil, zero value otherwise.

### GetKafkaBootstrapEndpointOk

`func (o *CmkV2KafkaClusterSpec) GetKafkaBootstrapEndpointOk() (*string, bool)`

GetKafkaBootstrapEndpointOk returns a tuple with the KafkaBootstrapEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaBootstrapEndpoint

`func (o *CmkV2KafkaClusterSpec) SetKafkaBootstrapEndpoint(v string)`

SetKafkaBootstrapEndpoint sets KafkaBootstrapEndpoint field to given value.

### HasKafkaBootstrapEndpoint

`func (o *CmkV2KafkaClusterSpec) HasKafkaBootstrapEndpoint() bool`

HasKafkaBootstrapEndpoint returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *CmkV2KafkaClusterSpec) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *CmkV2KafkaClusterSpec) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *CmkV2KafkaClusterSpec) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *CmkV2KafkaClusterSpec) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.

### GetEnvironment

`func (o *CmkV2KafkaClusterSpec) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CmkV2KafkaClusterSpec) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CmkV2KafkaClusterSpec) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CmkV2KafkaClusterSpec) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



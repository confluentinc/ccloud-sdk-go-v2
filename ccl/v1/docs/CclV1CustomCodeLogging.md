# CclV1CustomCodeLogging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Cloud** | Pointer to **string** | Cloud provider where the Custom Code Logging is sent. | [optional] 
**Region** | Pointer to **string** | The Cloud provider region the Custom Code Logging is sent. | [optional] 
**DestinationSettings** | Pointer to [**CclV1CustomCodeLoggingDestinationSettingsOneOf**](CclV1CustomCodeLoggingDestinationSettingsOneOf.md) | Destination Settings of the Custom Code Logging. | [optional] 
**Environment** | Pointer to [**EnvScopedObjectReference**](EnvScopedObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewCclV1CustomCodeLogging

`func NewCclV1CustomCodeLogging() *CclV1CustomCodeLogging`

NewCclV1CustomCodeLogging instantiates a new CclV1CustomCodeLogging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCclV1CustomCodeLoggingWithDefaults

`func NewCclV1CustomCodeLoggingWithDefaults() *CclV1CustomCodeLogging`

NewCclV1CustomCodeLoggingWithDefaults instantiates a new CclV1CustomCodeLogging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CclV1CustomCodeLogging) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CclV1CustomCodeLogging) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CclV1CustomCodeLogging) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CclV1CustomCodeLogging) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CclV1CustomCodeLogging) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CclV1CustomCodeLogging) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CclV1CustomCodeLogging) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CclV1CustomCodeLogging) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CclV1CustomCodeLogging) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CclV1CustomCodeLogging) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CclV1CustomCodeLogging) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CclV1CustomCodeLogging) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CclV1CustomCodeLogging) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CclV1CustomCodeLogging) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CclV1CustomCodeLogging) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CclV1CustomCodeLogging) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCloud

`func (o *CclV1CustomCodeLogging) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CclV1CustomCodeLogging) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CclV1CustomCodeLogging) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CclV1CustomCodeLogging) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *CclV1CustomCodeLogging) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CclV1CustomCodeLogging) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CclV1CustomCodeLogging) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CclV1CustomCodeLogging) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetDestinationSettings

`func (o *CclV1CustomCodeLogging) GetDestinationSettings() CclV1CustomCodeLoggingDestinationSettingsOneOf`

GetDestinationSettings returns the DestinationSettings field if non-nil, zero value otherwise.

### GetDestinationSettingsOk

`func (o *CclV1CustomCodeLogging) GetDestinationSettingsOk() (*CclV1CustomCodeLoggingDestinationSettingsOneOf, bool)`

GetDestinationSettingsOk returns a tuple with the DestinationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationSettings

`func (o *CclV1CustomCodeLogging) SetDestinationSettings(v CclV1CustomCodeLoggingDestinationSettingsOneOf)`

SetDestinationSettings sets DestinationSettings field to given value.

### HasDestinationSettings

`func (o *CclV1CustomCodeLogging) HasDestinationSettings() bool`

HasDestinationSettings returns a boolean if a field has been set.

### GetEnvironment

`func (o *CclV1CustomCodeLogging) GetEnvironment() EnvScopedObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CclV1CustomCodeLogging) GetEnvironmentOk() (*EnvScopedObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CclV1CustomCodeLogging) SetEnvironment(v EnvScopedObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CclV1CustomCodeLogging) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



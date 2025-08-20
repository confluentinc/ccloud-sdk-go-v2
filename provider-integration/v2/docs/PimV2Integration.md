# PimV2Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**DisplayName** | Pointer to **string** | Display name of Provider Integration. | [optional] 
**Provider** | Pointer to **string** | Cloud provider to which access is provided through provider integration. | [optional] [default to "GCP"]
**Config** | Pointer to [**PimV2IntegrationConfigOneOf**](PimV2IntegrationConfigOneOf.md) | Cloud provider specific configuration for the provider integration. Required only when updating integrations with &#x60;DRAFT&#x60; status. Not required during creation.  | [optional] 
**Usages** | Pointer to **[]string** | List of resource crns where this integration is being used. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the provider integration. - &#x60;DRAFT&#x60;: Integration exists but is not associated with customer configuration - &#x60;CREATED&#x60;: Integration has been associated with customer configuration - &#x60;ACTIVE&#x60;: Integration is in use by Confluent resources  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectReference**](ObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewPimV2Integration

`func NewPimV2Integration() *PimV2Integration`

NewPimV2Integration instantiates a new PimV2Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimV2IntegrationWithDefaults

`func NewPimV2IntegrationWithDefaults() *PimV2Integration`

NewPimV2IntegrationWithDefaults instantiates a new PimV2Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PimV2Integration) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PimV2Integration) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PimV2Integration) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PimV2Integration) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PimV2Integration) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PimV2Integration) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PimV2Integration) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PimV2Integration) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *PimV2Integration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PimV2Integration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PimV2Integration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PimV2Integration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *PimV2Integration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PimV2Integration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PimV2Integration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PimV2Integration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProvider

`func (o *PimV2Integration) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PimV2Integration) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PimV2Integration) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PimV2Integration) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetConfig

`func (o *PimV2Integration) GetConfig() PimV2IntegrationConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PimV2Integration) GetConfigOk() (*PimV2IntegrationConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PimV2Integration) SetConfig(v PimV2IntegrationConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PimV2Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetUsages

`func (o *PimV2Integration) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *PimV2Integration) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *PimV2Integration) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *PimV2Integration) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetStatus

`func (o *PimV2Integration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PimV2Integration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PimV2Integration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PimV2Integration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnvironment

`func (o *PimV2Integration) GetEnvironment() ObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PimV2Integration) GetEnvironmentOk() (*ObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PimV2Integration) SetEnvironment(v ObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PimV2Integration) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



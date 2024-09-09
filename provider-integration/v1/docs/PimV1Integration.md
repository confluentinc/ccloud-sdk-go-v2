# PimV1Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**DisplayName** | Pointer to **string** | Display name of Provider Integration. | [optional] 
**Provider** | Pointer to **string** | Cloud provider to which access is provided through provider integration. | [optional] [default to "AWS"]
**Config** | Pointer to [**PimV1IntegrationConfigOneOf**](PimV1IntegrationConfigOneOf.md) | Cloud provider specific configs for provider integration | [optional] 
**Usages** | Pointer to **[]string** | List of resource crns where this integration is being used. | [optional] [readonly] 
**Environment** | Pointer to [**GlobalObjectReference**](GlobalObjectReference.md) | The environment to which this belongs. | [optional] 

## Methods

### NewPimV1Integration

`func NewPimV1Integration() *PimV1Integration`

NewPimV1Integration instantiates a new PimV1Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimV1IntegrationWithDefaults

`func NewPimV1IntegrationWithDefaults() *PimV1Integration`

NewPimV1IntegrationWithDefaults instantiates a new PimV1Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PimV1Integration) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PimV1Integration) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PimV1Integration) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PimV1Integration) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PimV1Integration) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PimV1Integration) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PimV1Integration) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PimV1Integration) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *PimV1Integration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PimV1Integration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PimV1Integration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PimV1Integration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *PimV1Integration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PimV1Integration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PimV1Integration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PimV1Integration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProvider

`func (o *PimV1Integration) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PimV1Integration) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PimV1Integration) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PimV1Integration) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetConfig

`func (o *PimV1Integration) GetConfig() PimV1IntegrationConfigOneOf`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PimV1Integration) GetConfigOk() (*PimV1IntegrationConfigOneOf, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PimV1Integration) SetConfig(v PimV1IntegrationConfigOneOf)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PimV1Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetUsages

`func (o *PimV1Integration) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *PimV1Integration) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *PimV1Integration) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *PimV1Integration) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetEnvironment

`func (o *PimV1Integration) GetEnvironment() GlobalObjectReference`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PimV1Integration) GetEnvironmentOk() (*GlobalObjectReference, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PimV1Integration) SetEnvironment(v GlobalObjectReference)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PimV1Integration) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



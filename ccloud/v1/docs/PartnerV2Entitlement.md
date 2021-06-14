# PartnerV2Entitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**ExternalId** | Pointer to **string** | The unique external ID of the entitlement (this should be unique to customer) | [optional] 
**Name** | Pointer to **string** | The name of the entitlement | [optional] 
**PlanId** | Pointer to **string** | The plan ID the entitlement | [optional] 
**ProductId** | Pointer to **string** | The product ID of the entitlement | [optional] 
**UsageReportingId** | Pointer to **string** | The usage reporting ID of the entitlement (if usage reporting uses a different ID, otherwise, same as external_id)  | [optional] 
**ResourceId** | Pointer to **string** | The resource ID of the entitlement | [optional] 
**Organization** | Pointer to [**ObjectReference**](ObjectReference.md) | The organization associated with this object. | [optional] 

## Methods

### NewPartnerV2Entitlement

`func NewPartnerV2Entitlement() *PartnerV2Entitlement`

NewPartnerV2Entitlement instantiates a new PartnerV2Entitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerV2EntitlementWithDefaults

`func NewPartnerV2EntitlementWithDefaults() *PartnerV2Entitlement`

NewPartnerV2EntitlementWithDefaults instantiates a new PartnerV2Entitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PartnerV2Entitlement) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PartnerV2Entitlement) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PartnerV2Entitlement) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PartnerV2Entitlement) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PartnerV2Entitlement) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PartnerV2Entitlement) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PartnerV2Entitlement) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PartnerV2Entitlement) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *PartnerV2Entitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerV2Entitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerV2Entitlement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PartnerV2Entitlement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *PartnerV2Entitlement) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartnerV2Entitlement) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartnerV2Entitlement) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PartnerV2Entitlement) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExternalId

`func (o *PartnerV2Entitlement) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PartnerV2Entitlement) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PartnerV2Entitlement) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PartnerV2Entitlement) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetName

`func (o *PartnerV2Entitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerV2Entitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerV2Entitlement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartnerV2Entitlement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanId

`func (o *PartnerV2Entitlement) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PartnerV2Entitlement) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PartnerV2Entitlement) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PartnerV2Entitlement) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProductId

`func (o *PartnerV2Entitlement) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PartnerV2Entitlement) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PartnerV2Entitlement) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *PartnerV2Entitlement) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUsageReportingId

`func (o *PartnerV2Entitlement) GetUsageReportingId() string`

GetUsageReportingId returns the UsageReportingId field if non-nil, zero value otherwise.

### GetUsageReportingIdOk

`func (o *PartnerV2Entitlement) GetUsageReportingIdOk() (*string, bool)`

GetUsageReportingIdOk returns a tuple with the UsageReportingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageReportingId

`func (o *PartnerV2Entitlement) SetUsageReportingId(v string)`

SetUsageReportingId sets UsageReportingId field to given value.

### HasUsageReportingId

`func (o *PartnerV2Entitlement) HasUsageReportingId() bool`

HasUsageReportingId returns a boolean if a field has been set.

### GetResourceId

`func (o *PartnerV2Entitlement) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PartnerV2Entitlement) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PartnerV2Entitlement) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *PartnerV2Entitlement) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetOrganization

`func (o *PartnerV2Entitlement) GetOrganization() ObjectReference`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PartnerV2Entitlement) GetOrganizationOk() (*ObjectReference, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PartnerV2Entitlement) SetOrganization(v ObjectReference)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PartnerV2Entitlement) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



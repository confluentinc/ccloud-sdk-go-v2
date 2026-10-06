# BudgetV1InferenceBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**BudgetV1InferenceBudgetSpec**](BudgetV1InferenceBudgetSpec.md) |  | [optional] 
**Status** | Pointer to [**BudgetV1InferenceBudgetStatus**](BudgetV1InferenceBudgetStatus.md) |  | [optional] 

## Methods

### NewBudgetV1InferenceBudget

`func NewBudgetV1InferenceBudget() *BudgetV1InferenceBudget`

NewBudgetV1InferenceBudget instantiates a new BudgetV1InferenceBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBudgetV1InferenceBudgetWithDefaults

`func NewBudgetV1InferenceBudgetWithDefaults() *BudgetV1InferenceBudget`

NewBudgetV1InferenceBudgetWithDefaults instantiates a new BudgetV1InferenceBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BudgetV1InferenceBudget) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BudgetV1InferenceBudget) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BudgetV1InferenceBudget) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BudgetV1InferenceBudget) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *BudgetV1InferenceBudget) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BudgetV1InferenceBudget) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BudgetV1InferenceBudget) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BudgetV1InferenceBudget) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *BudgetV1InferenceBudget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BudgetV1InferenceBudget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BudgetV1InferenceBudget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BudgetV1InferenceBudget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *BudgetV1InferenceBudget) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BudgetV1InferenceBudget) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BudgetV1InferenceBudget) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BudgetV1InferenceBudget) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *BudgetV1InferenceBudget) GetSpec() BudgetV1InferenceBudgetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *BudgetV1InferenceBudget) GetSpecOk() (*BudgetV1InferenceBudgetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *BudgetV1InferenceBudget) SetSpec(v BudgetV1InferenceBudgetSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *BudgetV1InferenceBudget) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *BudgetV1InferenceBudget) GetStatus() BudgetV1InferenceBudgetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BudgetV1InferenceBudget) GetStatusOk() (*BudgetV1InferenceBudgetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BudgetV1InferenceBudget) SetStatus(v BudgetV1InferenceBudgetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BudgetV1InferenceBudget) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



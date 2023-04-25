# BillingV1alpha1Cost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**StartDate** | Pointer to **string** | Start date of time period (inclusive) to retrieve billing costs. It is represented in RFC3339 format and is in UTC. | [optional] 
**EndDate** | Pointer to **string** | End date of time period (exclusive) to retrieve billing costs. It is represented in RFC3339 format and is in UTC. | [optional] 
**Granularity** | Pointer to **string** | Granularity at which each line item is aggregated. | [optional] [default to "DAILY"]
**NetworkAccessType** | Pointer to **string** | Network access type for the cluster. | [optional] 
**Product** | Pointer to **string** | Product name. | [optional] 
**LineType** | Pointer to **string** | Type of the line item. | [optional] 
**Price** | Pointer to **float64** | Price for the line item in dollars. | [optional] 
**Unit** | Pointer to **string** | Unit of the line item. | [optional] 
**Quantity** | Pointer to **float64** | Quantity of the line item. | [optional] 
**OriginalAmount** | Pointer to **float64** | Original amount accrued for this line item. | [optional] 
**DiscountAmount** | Pointer to **float64** | Amount discounted from the original amount in dollars. | [optional] 
**Amount** | Pointer to **float64** | Final amount after deducting discounts. | [optional] 
**Resource** | Pointer to [**BillingV1alpha1Resource**](billing.v1alpha1.Resource.md) | The resource for a given object | [optional] 

## Methods

### NewBillingV1alpha1Cost

`func NewBillingV1alpha1Cost() *BillingV1alpha1Cost`

NewBillingV1alpha1Cost instantiates a new BillingV1alpha1Cost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1alpha1CostWithDefaults

`func NewBillingV1alpha1CostWithDefaults() *BillingV1alpha1Cost`

NewBillingV1alpha1CostWithDefaults instantiates a new BillingV1alpha1Cost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BillingV1alpha1Cost) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BillingV1alpha1Cost) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BillingV1alpha1Cost) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BillingV1alpha1Cost) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *BillingV1alpha1Cost) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BillingV1alpha1Cost) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BillingV1alpha1Cost) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BillingV1alpha1Cost) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *BillingV1alpha1Cost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingV1alpha1Cost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingV1alpha1Cost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingV1alpha1Cost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *BillingV1alpha1Cost) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BillingV1alpha1Cost) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BillingV1alpha1Cost) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BillingV1alpha1Cost) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStartDate

`func (o *BillingV1alpha1Cost) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingV1alpha1Cost) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingV1alpha1Cost) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingV1alpha1Cost) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingV1alpha1Cost) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingV1alpha1Cost) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingV1alpha1Cost) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingV1alpha1Cost) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGranularity

`func (o *BillingV1alpha1Cost) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *BillingV1alpha1Cost) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *BillingV1alpha1Cost) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *BillingV1alpha1Cost) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetNetworkAccessType

`func (o *BillingV1alpha1Cost) GetNetworkAccessType() string`

GetNetworkAccessType returns the NetworkAccessType field if non-nil, zero value otherwise.

### GetNetworkAccessTypeOk

`func (o *BillingV1alpha1Cost) GetNetworkAccessTypeOk() (*string, bool)`

GetNetworkAccessTypeOk returns a tuple with the NetworkAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAccessType

`func (o *BillingV1alpha1Cost) SetNetworkAccessType(v string)`

SetNetworkAccessType sets NetworkAccessType field to given value.

### HasNetworkAccessType

`func (o *BillingV1alpha1Cost) HasNetworkAccessType() bool`

HasNetworkAccessType returns a boolean if a field has been set.

### GetProduct

`func (o *BillingV1alpha1Cost) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *BillingV1alpha1Cost) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *BillingV1alpha1Cost) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *BillingV1alpha1Cost) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetLineType

`func (o *BillingV1alpha1Cost) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *BillingV1alpha1Cost) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *BillingV1alpha1Cost) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *BillingV1alpha1Cost) HasLineType() bool`

HasLineType returns a boolean if a field has been set.

### GetPrice

`func (o *BillingV1alpha1Cost) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingV1alpha1Cost) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingV1alpha1Cost) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingV1alpha1Cost) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetUnit

`func (o *BillingV1alpha1Cost) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BillingV1alpha1Cost) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BillingV1alpha1Cost) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BillingV1alpha1Cost) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetQuantity

`func (o *BillingV1alpha1Cost) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingV1alpha1Cost) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingV1alpha1Cost) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingV1alpha1Cost) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *BillingV1alpha1Cost) GetOriginalAmount() float64`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *BillingV1alpha1Cost) GetOriginalAmountOk() (*float64, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *BillingV1alpha1Cost) SetOriginalAmount(v float64)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *BillingV1alpha1Cost) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *BillingV1alpha1Cost) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *BillingV1alpha1Cost) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *BillingV1alpha1Cost) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *BillingV1alpha1Cost) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetAmount

`func (o *BillingV1alpha1Cost) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingV1alpha1Cost) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingV1alpha1Cost) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingV1alpha1Cost) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetResource

`func (o *BillingV1alpha1Cost) GetResource() BillingV1alpha1Resource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BillingV1alpha1Cost) GetResourceOk() (*BillingV1alpha1Resource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BillingV1alpha1Cost) SetResource(v BillingV1alpha1Resource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *BillingV1alpha1Cost) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



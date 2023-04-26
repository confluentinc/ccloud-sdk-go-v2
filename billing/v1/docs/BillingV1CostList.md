# BillingV1CostList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]BillingV1Cost**](BillingV1Cost.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewBillingV1CostList

`func NewBillingV1CostList(apiVersion string, kind string, metadata ListMeta, data []BillingV1Cost, ) *BillingV1CostList`

NewBillingV1CostList instantiates a new BillingV1CostList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1CostListWithDefaults

`func NewBillingV1CostListWithDefaults() *BillingV1CostList`

NewBillingV1CostListWithDefaults instantiates a new BillingV1CostList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BillingV1CostList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BillingV1CostList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BillingV1CostList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *BillingV1CostList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BillingV1CostList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BillingV1CostList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *BillingV1CostList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BillingV1CostList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BillingV1CostList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *BillingV1CostList) GetData() []BillingV1Cost`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillingV1CostList) GetDataOk() (*[]BillingV1Cost, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillingV1CostList) SetData(v []BillingV1Cost)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



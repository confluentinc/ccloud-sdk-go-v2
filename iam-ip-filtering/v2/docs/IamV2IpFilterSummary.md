# IamV2IpFilterSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Scope** | Pointer to **string** | The scope associated with this object. | [optional] 
**Categories** | Pointer to [**[]IamV2IpFilterSummaryCategories**](IamV2IpFilterSummaryCategories.md) | Summary of the operation groups and IP filters created in those operation groups.  | [optional] 

## Methods

### NewIamV2IpFilterSummary

`func NewIamV2IpFilterSummary() *IamV2IpFilterSummary`

NewIamV2IpFilterSummary instantiates a new IamV2IpFilterSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IpFilterSummaryWithDefaults

`func NewIamV2IpFilterSummaryWithDefaults() *IamV2IpFilterSummary`

NewIamV2IpFilterSummaryWithDefaults instantiates a new IamV2IpFilterSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IamV2IpFilterSummary) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IamV2IpFilterSummary) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IamV2IpFilterSummary) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IamV2IpFilterSummary) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IamV2IpFilterSummary) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IamV2IpFilterSummary) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IamV2IpFilterSummary) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IamV2IpFilterSummary) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetScope

`func (o *IamV2IpFilterSummary) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamV2IpFilterSummary) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamV2IpFilterSummary) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamV2IpFilterSummary) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetCategories

`func (o *IamV2IpFilterSummary) GetCategories() []IamV2IpFilterSummaryCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *IamV2IpFilterSummary) GetCategoriesOk() (*[]IamV2IpFilterSummaryCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *IamV2IpFilterSummary) SetCategories(v []IamV2IpFilterSummaryCategories)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *IamV2IpFilterSummary) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



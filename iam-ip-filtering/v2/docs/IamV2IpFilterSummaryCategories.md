# IamV2IpFilterSummaryCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the category. | [optional] 
**Status** | Pointer to **string** | Open, limited, or mixed. | [optional] 
**OperationGroups** | Pointer to [**[]IamV2IpFilterSummaryOperationGroups**](IamV2IpFilterSummaryOperationGroups.md) | Operation groups part of this category. | [optional] 

## Methods

### NewIamV2IpFilterSummaryCategories

`func NewIamV2IpFilterSummaryCategories() *IamV2IpFilterSummaryCategories`

NewIamV2IpFilterSummaryCategories instantiates a new IamV2IpFilterSummaryCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamV2IpFilterSummaryCategoriesWithDefaults

`func NewIamV2IpFilterSummaryCategoriesWithDefaults() *IamV2IpFilterSummaryCategories`

NewIamV2IpFilterSummaryCategoriesWithDefaults instantiates a new IamV2IpFilterSummaryCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IamV2IpFilterSummaryCategories) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamV2IpFilterSummaryCategories) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamV2IpFilterSummaryCategories) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamV2IpFilterSummaryCategories) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *IamV2IpFilterSummaryCategories) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamV2IpFilterSummaryCategories) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamV2IpFilterSummaryCategories) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamV2IpFilterSummaryCategories) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOperationGroups

`func (o *IamV2IpFilterSummaryCategories) GetOperationGroups() []IamV2IpFilterSummaryOperationGroups`

GetOperationGroups returns the OperationGroups field if non-nil, zero value otherwise.

### GetOperationGroupsOk

`func (o *IamV2IpFilterSummaryCategories) GetOperationGroupsOk() (*[]IamV2IpFilterSummaryOperationGroups, bool)`

GetOperationGroupsOk returns a tuple with the OperationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationGroups

`func (o *IamV2IpFilterSummaryCategories) SetOperationGroups(v []IamV2IpFilterSummaryOperationGroups)`

SetOperationGroups sets OperationGroups field to given value.

### HasOperationGroups

`func (o *IamV2IpFilterSummaryCategories) HasOperationGroups() bool`

HasOperationGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



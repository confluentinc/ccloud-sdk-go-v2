# V1ServiceAccountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]V1ServiceAccount**](V1ServiceAccount.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewV1ServiceAccountList

`func NewV1ServiceAccountList(users []V1ServiceAccount, ) *V1ServiceAccountList`

NewV1ServiceAccountList instantiates a new V1ServiceAccountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ServiceAccountListWithDefaults

`func NewV1ServiceAccountListWithDefaults() *V1ServiceAccountList`

NewV1ServiceAccountListWithDefaults instantiates a new V1ServiceAccountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *V1ServiceAccountList) GetUsers() []V1ServiceAccount`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V1ServiceAccountList) GetUsersOk() (*[]V1ServiceAccount, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V1ServiceAccountList) SetUsers(v []V1ServiceAccount)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



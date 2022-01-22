# V1UserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]V1User**](V1User.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewV1UserList

`func NewV1UserList(users []V1User, ) *V1UserList`

NewV1UserList instantiates a new V1UserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UserListWithDefaults

`func NewV1UserListWithDefaults() *V1UserList`

NewV1UserListWithDefaults instantiates a new V1UserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *V1UserList) GetUsers() []V1User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V1UserList) GetUsersOk() (*[]V1User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V1UserList) SetUsers(v []V1User)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateAclRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | [**AclResourceType**](AclResourceType.md) |  | 
**ResourceName** | **string** |  | 
**PatternType** | **string** |  | 
**Principal** | **string** |  | 
**Host** | **string** |  | 
**Operation** | **string** |  | 
**Permission** | **string** |  | 

## Methods

### NewCreateAclRequestData

`func NewCreateAclRequestData(resourceType AclResourceType, resourceName string, patternType string, principal string, host string, operation string, permission string, ) *CreateAclRequestData`

NewCreateAclRequestData instantiates a new CreateAclRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAclRequestDataWithDefaults

`func NewCreateAclRequestDataWithDefaults() *CreateAclRequestData`

NewCreateAclRequestDataWithDefaults instantiates a new CreateAclRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *CreateAclRequestData) GetResourceType() AclResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateAclRequestData) GetResourceTypeOk() (*AclResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateAclRequestData) SetResourceType(v AclResourceType)`

SetResourceType sets ResourceType field to given value.


### GetResourceName

`func (o *CreateAclRequestData) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *CreateAclRequestData) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *CreateAclRequestData) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetPatternType

`func (o *CreateAclRequestData) GetPatternType() string`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *CreateAclRequestData) GetPatternTypeOk() (*string, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *CreateAclRequestData) SetPatternType(v string)`

SetPatternType sets PatternType field to given value.


### GetPrincipal

`func (o *CreateAclRequestData) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *CreateAclRequestData) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *CreateAclRequestData) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetHost

`func (o *CreateAclRequestData) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateAclRequestData) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateAclRequestData) SetHost(v string)`

SetHost sets Host field to given value.


### GetOperation

`func (o *CreateAclRequestData) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CreateAclRequestData) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CreateAclRequestData) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetPermission

`func (o *CreateAclRequestData) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CreateAclRequestData) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CreateAclRequestData) SetPermission(v string)`

SetPermission sets Permission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



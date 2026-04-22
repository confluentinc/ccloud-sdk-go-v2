# SqlV1StatementStatusAffectedResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind of resource that was created. | 
**EnvironmentId** | **string** | The unique identifier for the environment containing the resource. | 
**DatabaseId** | Pointer to **string** | The unique identifier for the database containing the resource. Only present for resource kinds that are scoped to a database.  | [optional] 
**ResourceName** | **string** | The name of the created resource, unique within its scope. | 

## Methods

### NewSqlV1StatementStatusAffectedResource

`func NewSqlV1StatementStatusAffectedResource(kind string, environmentId string, resourceName string, ) *SqlV1StatementStatusAffectedResource`

NewSqlV1StatementStatusAffectedResource instantiates a new SqlV1StatementStatusAffectedResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1StatementStatusAffectedResourceWithDefaults

`func NewSqlV1StatementStatusAffectedResourceWithDefaults() *SqlV1StatementStatusAffectedResource`

NewSqlV1StatementStatusAffectedResourceWithDefaults instantiates a new SqlV1StatementStatusAffectedResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SqlV1StatementStatusAffectedResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1StatementStatusAffectedResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1StatementStatusAffectedResource) SetKind(v string)`

SetKind sets Kind field to given value.


### GetEnvironmentId

`func (o *SqlV1StatementStatusAffectedResource) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SqlV1StatementStatusAffectedResource) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SqlV1StatementStatusAffectedResource) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetDatabaseId

`func (o *SqlV1StatementStatusAffectedResource) GetDatabaseId() string`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *SqlV1StatementStatusAffectedResource) GetDatabaseIdOk() (*string, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *SqlV1StatementStatusAffectedResource) SetDatabaseId(v string)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *SqlV1StatementStatusAffectedResource) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### GetResourceName

`func (o *SqlV1StatementStatusAffectedResource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *SqlV1StatementStatusAffectedResource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *SqlV1StatementStatusAffectedResource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



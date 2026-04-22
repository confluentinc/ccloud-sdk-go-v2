# SqlV1ResourceChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The type of resource that was changed. | 
**EnvironmentLocator** | **string** | The environment containing the resource. Can be either the environment name or ID, depending on how it is referenced in the SQL statement text. | 
**DatabaseLocator** | **string** | The database containing the resource. Can be either the database name or ID, depending on how it is referenced in the SQL statement text. | 
**Name** | **string** | The name of the resource, unique within its scope (environment and database). | 
**Details** | **[]string** | Human-readable descriptions of the changes made to this resource. | 

## Methods

### NewSqlV1ResourceChange

`func NewSqlV1ResourceChange(kind string, environmentLocator string, databaseLocator string, name string, details []string, ) *SqlV1ResourceChange`

NewSqlV1ResourceChange instantiates a new SqlV1ResourceChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlV1ResourceChangeWithDefaults

`func NewSqlV1ResourceChangeWithDefaults() *SqlV1ResourceChange`

NewSqlV1ResourceChangeWithDefaults instantiates a new SqlV1ResourceChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SqlV1ResourceChange) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SqlV1ResourceChange) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SqlV1ResourceChange) SetKind(v string)`

SetKind sets Kind field to given value.


### GetEnvironmentLocator

`func (o *SqlV1ResourceChange) GetEnvironmentLocator() string`

GetEnvironmentLocator returns the EnvironmentLocator field if non-nil, zero value otherwise.

### GetEnvironmentLocatorOk

`func (o *SqlV1ResourceChange) GetEnvironmentLocatorOk() (*string, bool)`

GetEnvironmentLocatorOk returns a tuple with the EnvironmentLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentLocator

`func (o *SqlV1ResourceChange) SetEnvironmentLocator(v string)`

SetEnvironmentLocator sets EnvironmentLocator field to given value.


### GetDatabaseLocator

`func (o *SqlV1ResourceChange) GetDatabaseLocator() string`

GetDatabaseLocator returns the DatabaseLocator field if non-nil, zero value otherwise.

### GetDatabaseLocatorOk

`func (o *SqlV1ResourceChange) GetDatabaseLocatorOk() (*string, bool)`

GetDatabaseLocatorOk returns a tuple with the DatabaseLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseLocator

`func (o *SqlV1ResourceChange) SetDatabaseLocator(v string)`

SetDatabaseLocator sets DatabaseLocator field to given value.


### GetName

`func (o *SqlV1ResourceChange) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlV1ResourceChange) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlV1ResourceChange) SetName(v string)`

SetName sets Name field to given value.


### GetDetails

`func (o *SqlV1ResourceChange) GetDetails() []string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SqlV1ResourceChange) GetDetailsOk() (*[]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SqlV1ResourceChange) SetDetails(v []string)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



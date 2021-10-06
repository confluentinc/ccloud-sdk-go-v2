# V2ResourceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | What type of resource these operations are allowed on (\&quot;Cluster\&quot;, \&quot;Topic\&quot;, \&quot;Group\&quot;, \&quot;Subject\&quot;, etc.) | [optional] 
**Operations** | Pointer to **[]string** | List of allowed operations on this resourceType | [optional] 

## Methods

### NewV2ResourceOperation

`func NewV2ResourceOperation() *V2ResourceOperation`

NewV2ResourceOperation instantiates a new V2ResourceOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ResourceOperationWithDefaults

`func NewV2ResourceOperationWithDefaults() *V2ResourceOperation`

NewV2ResourceOperationWithDefaults instantiates a new V2ResourceOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *V2ResourceOperation) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *V2ResourceOperation) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *V2ResourceOperation) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *V2ResourceOperation) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetOperations

`func (o *V2ResourceOperation) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *V2ResourceOperation) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *V2ResourceOperation) SetOperations(v []string)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *V2ResourceOperation) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



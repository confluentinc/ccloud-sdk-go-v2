# TypedEnvScopedObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Environment** | Pointer to **string** | Environment of the referred resource, if env-scoped | [optional] 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 
**ApiVersion** | Pointer to **string** | API group and version of the referred resource | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the referred resource | [optional] [readonly] 

## Methods

### NewTypedEnvScopedObjectReference

`func NewTypedEnvScopedObjectReference(id string, related string, resourceName string, ) *TypedEnvScopedObjectReference`

NewTypedEnvScopedObjectReference instantiates a new TypedEnvScopedObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedEnvScopedObjectReferenceWithDefaults

`func NewTypedEnvScopedObjectReferenceWithDefaults() *TypedEnvScopedObjectReference`

NewTypedEnvScopedObjectReferenceWithDefaults instantiates a new TypedEnvScopedObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypedEnvScopedObjectReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypedEnvScopedObjectReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypedEnvScopedObjectReference) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *TypedEnvScopedObjectReference) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TypedEnvScopedObjectReference) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TypedEnvScopedObjectReference) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *TypedEnvScopedObjectReference) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetRelated

`func (o *TypedEnvScopedObjectReference) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *TypedEnvScopedObjectReference) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *TypedEnvScopedObjectReference) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *TypedEnvScopedObjectReference) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *TypedEnvScopedObjectReference) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *TypedEnvScopedObjectReference) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetApiVersion

`func (o *TypedEnvScopedObjectReference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TypedEnvScopedObjectReference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TypedEnvScopedObjectReference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TypedEnvScopedObjectReference) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *TypedEnvScopedObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TypedEnvScopedObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TypedEnvScopedObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TypedEnvScopedObjectReference) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



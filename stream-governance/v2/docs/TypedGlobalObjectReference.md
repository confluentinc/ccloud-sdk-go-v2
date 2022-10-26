# TypedGlobalObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the referred resource | 
**Related** | **string** | API URL for accessing or modifying the referred object | [readonly] 
**ResourceName** | **string** | CRN reference to the referred resource | [readonly] 
**ApiVersion** | Pointer to **string** | API group and version of the referred resource | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind of the referred resource | [optional] [readonly] 

## Methods

### NewTypedGlobalObjectReference

`func NewTypedGlobalObjectReference(id string, related string, resourceName string, ) *TypedGlobalObjectReference`

NewTypedGlobalObjectReference instantiates a new TypedGlobalObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypedGlobalObjectReferenceWithDefaults

`func NewTypedGlobalObjectReferenceWithDefaults() *TypedGlobalObjectReference`

NewTypedGlobalObjectReferenceWithDefaults instantiates a new TypedGlobalObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TypedGlobalObjectReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypedGlobalObjectReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypedGlobalObjectReference) SetId(v string)`

SetId sets Id field to given value.


### GetRelated

`func (o *TypedGlobalObjectReference) GetRelated() string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *TypedGlobalObjectReference) GetRelatedOk() (*string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *TypedGlobalObjectReference) SetRelated(v string)`

SetRelated sets Related field to given value.


### GetResourceName

`func (o *TypedGlobalObjectReference) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *TypedGlobalObjectReference) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *TypedGlobalObjectReference) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetApiVersion

`func (o *TypedGlobalObjectReference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TypedGlobalObjectReference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TypedGlobalObjectReference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TypedGlobalObjectReference) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *TypedGlobalObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TypedGlobalObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TypedGlobalObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TypedGlobalObjectReference) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



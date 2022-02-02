# ResourceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **string** |  | 
**ResourceName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResourceMetadata

`func NewResourceMetadata(self string, ) *ResourceMetadata`

NewResourceMetadata instantiates a new ResourceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceMetadataWithDefaults

`func NewResourceMetadataWithDefaults() *ResourceMetadata`

NewResourceMetadataWithDefaults instantiates a new ResourceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ResourceMetadata) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ResourceMetadata) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ResourceMetadata) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetResourceName

`func (o *ResourceMetadata) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceMetadata) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceMetadata) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ResourceMetadata) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *ResourceMetadata) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *ResourceMetadata) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



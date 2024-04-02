# StreamCatalogV2Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the domain. | 
**Name** | Pointer to **string** | The name of the domain. | [optional] 
**ResourceType** | Pointer to **string** | The type of the resource | [optional] 
**Community** | [**StreamCatalogV2DomainCommunity**](StreamCatalogV2DomainCommunity.md) |  | 
**Description** | Pointer to **string** | A brief description of the domain. | [optional] 

## Methods

### NewStreamCatalogV2Domain

`func NewStreamCatalogV2Domain(id string, community StreamCatalogV2DomainCommunity, ) *StreamCatalogV2Domain`

NewStreamCatalogV2Domain instantiates a new StreamCatalogV2Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamCatalogV2DomainWithDefaults

`func NewStreamCatalogV2DomainWithDefaults() *StreamCatalogV2Domain`

NewStreamCatalogV2DomainWithDefaults instantiates a new StreamCatalogV2Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamCatalogV2Domain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamCatalogV2Domain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamCatalogV2Domain) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StreamCatalogV2Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamCatalogV2Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamCatalogV2Domain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamCatalogV2Domain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResourceType

`func (o *StreamCatalogV2Domain) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *StreamCatalogV2Domain) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *StreamCatalogV2Domain) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *StreamCatalogV2Domain) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetCommunity

`func (o *StreamCatalogV2Domain) GetCommunity() StreamCatalogV2DomainCommunity`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *StreamCatalogV2Domain) GetCommunityOk() (*StreamCatalogV2DomainCommunity, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *StreamCatalogV2Domain) SetCommunity(v StreamCatalogV2DomainCommunity)`

SetCommunity sets Community field to given value.


### GetDescription

`func (o *StreamCatalogV2Domain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamCatalogV2Domain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamCatalogV2Domain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StreamCatalogV2Domain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



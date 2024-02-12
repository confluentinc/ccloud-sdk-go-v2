# CatalogConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | **map[string]string** | Properties that should be used to override client configuration; applied after defaults and client configuration. | 
**Defaults** | **map[string]string** | Properties that should be used as default configuration; applied before client configuration. | 

## Methods

### NewCatalogConfig

`func NewCatalogConfig(overrides map[string]string, defaults map[string]string, ) *CatalogConfig`

NewCatalogConfig instantiates a new CatalogConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogConfigWithDefaults

`func NewCatalogConfigWithDefaults() *CatalogConfig`

NewCatalogConfigWithDefaults instantiates a new CatalogConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *CatalogConfig) GetOverrides() map[string]string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *CatalogConfig) GetOverridesOk() (*map[string]string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *CatalogConfig) SetOverrides(v map[string]string)`

SetOverrides sets Overrides field to given value.


### GetDefaults

`func (o *CatalogConfig) GetDefaults() map[string]string`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *CatalogConfig) GetDefaultsOk() (*map[string]string, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *CatalogConfig) SetDefaults(v map[string]string)`

SetDefaults sets Defaults field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ExporterConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaRegistryUrl** | Pointer to **string** | Config SR URL | [optional] 
**BasicAuthCredentialsSource** | Pointer to **string** | Config SR Auth | [optional] 
**BasicAuthUserInfo** | Pointer to **string** | Config SR User Info | [optional] 

## Methods

### NewExporterConfigResponse

`func NewExporterConfigResponse() *ExporterConfigResponse`

NewExporterConfigResponse instantiates a new ExporterConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExporterConfigResponseWithDefaults

`func NewExporterConfigResponseWithDefaults() *ExporterConfigResponse`

NewExporterConfigResponseWithDefaults instantiates a new ExporterConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaRegistryUrl

`func (o *ExporterConfigResponse) GetSchemaRegistryUrl() string`

GetSchemaRegistryUrl returns the SchemaRegistryUrl field if non-nil, zero value otherwise.

### GetSchemaRegistryUrlOk

`func (o *ExporterConfigResponse) GetSchemaRegistryUrlOk() (*string, bool)`

GetSchemaRegistryUrlOk returns a tuple with the SchemaRegistryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistryUrl

`func (o *ExporterConfigResponse) SetSchemaRegistryUrl(v string)`

SetSchemaRegistryUrl sets SchemaRegistryUrl field to given value.

### HasSchemaRegistryUrl

`func (o *ExporterConfigResponse) HasSchemaRegistryUrl() bool`

HasSchemaRegistryUrl returns a boolean if a field has been set.

### GetBasicAuthCredentialsSource

`func (o *ExporterConfigResponse) GetBasicAuthCredentialsSource() string`

GetBasicAuthCredentialsSource returns the BasicAuthCredentialsSource field if non-nil, zero value otherwise.

### GetBasicAuthCredentialsSourceOk

`func (o *ExporterConfigResponse) GetBasicAuthCredentialsSourceOk() (*string, bool)`

GetBasicAuthCredentialsSourceOk returns a tuple with the BasicAuthCredentialsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthCredentialsSource

`func (o *ExporterConfigResponse) SetBasicAuthCredentialsSource(v string)`

SetBasicAuthCredentialsSource sets BasicAuthCredentialsSource field to given value.

### HasBasicAuthCredentialsSource

`func (o *ExporterConfigResponse) HasBasicAuthCredentialsSource() bool`

HasBasicAuthCredentialsSource returns a boolean if a field has been set.

### GetBasicAuthUserInfo

`func (o *ExporterConfigResponse) GetBasicAuthUserInfo() string`

GetBasicAuthUserInfo returns the BasicAuthUserInfo field if non-nil, zero value otherwise.

### GetBasicAuthUserInfoOk

`func (o *ExporterConfigResponse) GetBasicAuthUserInfoOk() (*string, bool)`

GetBasicAuthUserInfoOk returns a tuple with the BasicAuthUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUserInfo

`func (o *ExporterConfigResponse) SetBasicAuthUserInfo(v string)`

SetBasicAuthUserInfo sets BasicAuthUserInfo field to given value.

### HasBasicAuthUserInfo

`func (o *ExporterConfigResponse) HasBasicAuthUserInfo() bool`

HasBasicAuthUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



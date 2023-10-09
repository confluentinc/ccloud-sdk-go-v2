# WsV1beta1Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of block. | [optional] 
**Properties** | Pointer to **map[string]string** | A map (key-value pairs) of cell properties. All key-value pairs are optional, and clients choose whether and how to use these key-value pairs. | [optional] 
**CodeOptions** | Pointer to [**WsV1beta1CodeOptions**](WsV1beta1CodeOptions.md) |  | [optional] 

## Methods

### NewWsV1beta1Block

`func NewWsV1beta1Block() *WsV1beta1Block`

NewWsV1beta1Block instantiates a new WsV1beta1Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWsV1beta1BlockWithDefaults

`func NewWsV1beta1BlockWithDefaults() *WsV1beta1Block`

NewWsV1beta1BlockWithDefaults instantiates a new WsV1beta1Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WsV1beta1Block) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WsV1beta1Block) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WsV1beta1Block) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WsV1beta1Block) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProperties

`func (o *WsV1beta1Block) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WsV1beta1Block) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WsV1beta1Block) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WsV1beta1Block) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCodeOptions

`func (o *WsV1beta1Block) GetCodeOptions() WsV1beta1CodeOptions`

GetCodeOptions returns the CodeOptions field if non-nil, zero value otherwise.

### GetCodeOptionsOk

`func (o *WsV1beta1Block) GetCodeOptionsOk() (*WsV1beta1CodeOptions, bool)`

GetCodeOptionsOk returns a tuple with the CodeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeOptions

`func (o *WsV1beta1Block) SetCodeOptions(v WsV1beta1CodeOptions)`

SetCodeOptions sets CodeOptions field to given value.

### HasCodeOptions

`func (o *WsV1beta1Block) HasCodeOptions() bool`

HasCodeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



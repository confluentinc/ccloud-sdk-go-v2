# AttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **[]map[string]string** | The enumerated labels, as an array of key/value pairs. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewAttributesResponse

`func NewAttributesResponse(data []map[string]string, meta Meta, ) *AttributesResponse`

NewAttributesResponse instantiates a new AttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributesResponseWithDefaults

`func NewAttributesResponseWithDefaults() *AttributesResponse`

NewAttributesResponseWithDefaults instantiates a new AttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AttributesResponse) GetData() []map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AttributesResponse) GetDataOk() (*[]map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AttributesResponse) SetData(v []map[string]string)`

SetData sets Data field to given value.


### GetMeta

`func (o *AttributesResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AttributesResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AttributesResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *AttributesResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AttributesResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AttributesResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AttributesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



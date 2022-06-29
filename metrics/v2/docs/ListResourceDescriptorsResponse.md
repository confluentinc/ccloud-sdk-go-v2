# ListResourceDescriptorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ResourceDescriptor**](ResourceDescriptor.md) | The resource descriptors for the specified dataset. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewListResourceDescriptorsResponse

`func NewListResourceDescriptorsResponse(data []ResourceDescriptor, meta Meta, ) *ListResourceDescriptorsResponse`

NewListResourceDescriptorsResponse instantiates a new ListResourceDescriptorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourceDescriptorsResponseWithDefaults

`func NewListResourceDescriptorsResponseWithDefaults() *ListResourceDescriptorsResponse`

NewListResourceDescriptorsResponseWithDefaults instantiates a new ListResourceDescriptorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListResourceDescriptorsResponse) GetData() []ResourceDescriptor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListResourceDescriptorsResponse) GetDataOk() (*[]ResourceDescriptor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListResourceDescriptorsResponse) SetData(v []ResourceDescriptor)`

SetData sets Data field to given value.


### GetMeta

`func (o *ListResourceDescriptorsResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListResourceDescriptorsResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListResourceDescriptorsResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *ListResourceDescriptorsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListResourceDescriptorsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListResourceDescriptorsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListResourceDescriptorsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



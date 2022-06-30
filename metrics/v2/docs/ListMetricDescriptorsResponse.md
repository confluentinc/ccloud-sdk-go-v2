# ListMetricDescriptorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]MetricDescriptor**](MetricDescriptor.md) | The metric descriptors for the specified dataset. | 
**Meta** | [**Meta**](Meta.md) |  | 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewListMetricDescriptorsResponse

`func NewListMetricDescriptorsResponse(data []MetricDescriptor, meta Meta, ) *ListMetricDescriptorsResponse`

NewListMetricDescriptorsResponse instantiates a new ListMetricDescriptorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricDescriptorsResponseWithDefaults

`func NewListMetricDescriptorsResponseWithDefaults() *ListMetricDescriptorsResponse`

NewListMetricDescriptorsResponseWithDefaults instantiates a new ListMetricDescriptorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListMetricDescriptorsResponse) GetData() []MetricDescriptor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListMetricDescriptorsResponse) GetDataOk() (*[]MetricDescriptor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListMetricDescriptorsResponse) SetData(v []MetricDescriptor)`

SetData sets Data field to given value.


### GetMeta

`func (o *ListMetricDescriptorsResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMetricDescriptorsResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMetricDescriptorsResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.


### GetLinks

`func (o *ListMetricDescriptorsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListMetricDescriptorsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListMetricDescriptorsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListMetricDescriptorsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FlatQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Point**](Point.md) | An array of results for this query. Each item includes &#x60;timestamp&#x60;, &#x60;value&#x60;, and an attribute for each label specified in the request&#39;s &#x60;group_by&#x60;. | 

## Methods

### NewFlatQueryResponse

`func NewFlatQueryResponse(data []Point, ) *FlatQueryResponse`

NewFlatQueryResponse instantiates a new FlatQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatQueryResponseWithDefaults

`func NewFlatQueryResponseWithDefaults() *FlatQueryResponse`

NewFlatQueryResponseWithDefaults instantiates a new FlatQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FlatQueryResponse) GetData() []Point`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FlatQueryResponse) GetDataOk() (*[]Point, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FlatQueryResponse) SetData(v []Point)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



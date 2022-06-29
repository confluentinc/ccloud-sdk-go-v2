# GroupedQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GroupedQueryResponseDataInner**](GroupedQueryResponseDataInner.md) | An array of results for this query. Each item represents a group bucket having a distinct set of label values for the request&#39;s &#x60;group_by&#x60;.  The groups are ordered lexicographically by the label values for that group. | 

## Methods

### NewGroupedQueryResponse

`func NewGroupedQueryResponse(data []GroupedQueryResponseDataInner, ) *GroupedQueryResponse`

NewGroupedQueryResponse instantiates a new GroupedQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupedQueryResponseWithDefaults

`func NewGroupedQueryResponseWithDefaults() *GroupedQueryResponse`

NewGroupedQueryResponseWithDefaults instantiates a new GroupedQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GroupedQueryResponse) GetData() []GroupedQueryResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupedQueryResponse) GetDataOk() (*[]GroupedQueryResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupedQueryResponse) SetData(v []GroupedQueryResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



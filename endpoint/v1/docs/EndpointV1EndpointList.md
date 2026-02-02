# EndpointV1EndpointList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the schema version of this representation of a resource. | [readonly] 
**Kind** | **string** | Kind defines the object this REST resource represents. | [readonly] 
**Metadata** | [**ListMeta**](ListMeta.md) |  | 
**Data** | [**[]EndpointV1Endpoint**](EndpointV1Endpoint.md) | A data property that contains an array of resource items. Each entry in the array is a separate resource. | 

## Methods

### NewEndpointV1EndpointList

`func NewEndpointV1EndpointList(apiVersion string, kind string, metadata ListMeta, data []EndpointV1Endpoint, ) *EndpointV1EndpointList`

NewEndpointV1EndpointList instantiates a new EndpointV1EndpointList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointV1EndpointListWithDefaults

`func NewEndpointV1EndpointListWithDefaults() *EndpointV1EndpointList`

NewEndpointV1EndpointListWithDefaults instantiates a new EndpointV1EndpointList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *EndpointV1EndpointList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EndpointV1EndpointList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EndpointV1EndpointList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *EndpointV1EndpointList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EndpointV1EndpointList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EndpointV1EndpointList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *EndpointV1EndpointList) GetMetadata() ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EndpointV1EndpointList) GetMetadataOk() (*ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EndpointV1EndpointList) SetMetadata(v ListMeta)`

SetMetadata sets Metadata field to given value.


### GetData

`func (o *EndpointV1EndpointList) GetData() []EndpointV1Endpoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EndpointV1EndpointList) GetDataOk() (*[]EndpointV1Endpoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EndpointV1EndpointList) SetData(v []EndpointV1Endpoint)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



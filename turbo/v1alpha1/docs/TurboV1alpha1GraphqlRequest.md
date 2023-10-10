# TurboV1alpha1GraphqlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Query** | Pointer to **string** | The GraphQL query string. | [optional] 

## Methods

### NewTurboV1alpha1GraphqlRequest

`func NewTurboV1alpha1GraphqlRequest() *TurboV1alpha1GraphqlRequest`

NewTurboV1alpha1GraphqlRequest instantiates a new TurboV1alpha1GraphqlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTurboV1alpha1GraphqlRequestWithDefaults

`func NewTurboV1alpha1GraphqlRequestWithDefaults() *TurboV1alpha1GraphqlRequest`

NewTurboV1alpha1GraphqlRequestWithDefaults instantiates a new TurboV1alpha1GraphqlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *TurboV1alpha1GraphqlRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TurboV1alpha1GraphqlRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TurboV1alpha1GraphqlRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TurboV1alpha1GraphqlRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *TurboV1alpha1GraphqlRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TurboV1alpha1GraphqlRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TurboV1alpha1GraphqlRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TurboV1alpha1GraphqlRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *TurboV1alpha1GraphqlRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TurboV1alpha1GraphqlRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TurboV1alpha1GraphqlRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TurboV1alpha1GraphqlRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *TurboV1alpha1GraphqlRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TurboV1alpha1GraphqlRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TurboV1alpha1GraphqlRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TurboV1alpha1GraphqlRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQuery

`func (o *TurboV1alpha1GraphqlRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TurboV1alpha1GraphqlRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TurboV1alpha1GraphqlRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *TurboV1alpha1GraphqlRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CdxV1RedeemTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Token** | Pointer to **string** | The encrypted token | [optional] 
**AwsAccount** | Pointer to **string** | The AWS account id for the consumer network. | [optional] 
**AzureSubscription** | Pointer to **string** | The Azure subscription for the consumer network. | [optional] 

## Methods

### NewCdxV1RedeemTokenRequest

`func NewCdxV1RedeemTokenRequest() *CdxV1RedeemTokenRequest`

NewCdxV1RedeemTokenRequest instantiates a new CdxV1RedeemTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdxV1RedeemTokenRequestWithDefaults

`func NewCdxV1RedeemTokenRequestWithDefaults() *CdxV1RedeemTokenRequest`

NewCdxV1RedeemTokenRequestWithDefaults instantiates a new CdxV1RedeemTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CdxV1RedeemTokenRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CdxV1RedeemTokenRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CdxV1RedeemTokenRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CdxV1RedeemTokenRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CdxV1RedeemTokenRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CdxV1RedeemTokenRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CdxV1RedeemTokenRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CdxV1RedeemTokenRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CdxV1RedeemTokenRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdxV1RedeemTokenRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdxV1RedeemTokenRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdxV1RedeemTokenRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CdxV1RedeemTokenRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CdxV1RedeemTokenRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CdxV1RedeemTokenRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CdxV1RedeemTokenRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetToken

`func (o *CdxV1RedeemTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CdxV1RedeemTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CdxV1RedeemTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CdxV1RedeemTokenRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAwsAccount

`func (o *CdxV1RedeemTokenRequest) GetAwsAccount() string`

GetAwsAccount returns the AwsAccount field if non-nil, zero value otherwise.

### GetAwsAccountOk

`func (o *CdxV1RedeemTokenRequest) GetAwsAccountOk() (*string, bool)`

GetAwsAccountOk returns a tuple with the AwsAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccount

`func (o *CdxV1RedeemTokenRequest) SetAwsAccount(v string)`

SetAwsAccount sets AwsAccount field to given value.

### HasAwsAccount

`func (o *CdxV1RedeemTokenRequest) HasAwsAccount() bool`

HasAwsAccount returns a boolean if a field has been set.

### GetAzureSubscription

`func (o *CdxV1RedeemTokenRequest) GetAzureSubscription() string`

GetAzureSubscription returns the AzureSubscription field if non-nil, zero value otherwise.

### GetAzureSubscriptionOk

`func (o *CdxV1RedeemTokenRequest) GetAzureSubscriptionOk() (*string, bool)`

GetAzureSubscriptionOk returns a tuple with the AzureSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscription

`func (o *CdxV1RedeemTokenRequest) SetAzureSubscription(v string)`

SetAzureSubscription sets AzureSubscription field to given value.

### HasAzureSubscription

`func (o *CdxV1RedeemTokenRequest) HasAzureSubscription() bool`

HasAzureSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



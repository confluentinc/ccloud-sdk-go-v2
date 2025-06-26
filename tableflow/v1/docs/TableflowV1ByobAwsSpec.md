# TableflowV1ByobAwsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The storage type  | 
**BucketName** | **string** | Bucket name | 
**BucketRegion** | Pointer to **string** | Bucket region | [optional] [readonly] 
**ProviderIntegrationId** | **string** | The provider integration id | 
**TablePath** | Pointer to **string** | The current storage path where the data and metadata is stored for this table | [optional] [readonly] 

## Methods

### NewTableflowV1ByobAwsSpec

`func NewTableflowV1ByobAwsSpec(kind string, bucketName string, providerIntegrationId string, ) *TableflowV1ByobAwsSpec`

NewTableflowV1ByobAwsSpec instantiates a new TableflowV1ByobAwsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1ByobAwsSpecWithDefaults

`func NewTableflowV1ByobAwsSpecWithDefaults() *TableflowV1ByobAwsSpec`

NewTableflowV1ByobAwsSpecWithDefaults instantiates a new TableflowV1ByobAwsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1ByobAwsSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1ByobAwsSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1ByobAwsSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetBucketName

`func (o *TableflowV1ByobAwsSpec) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *TableflowV1ByobAwsSpec) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *TableflowV1ByobAwsSpec) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetBucketRegion

`func (o *TableflowV1ByobAwsSpec) GetBucketRegion() string`

GetBucketRegion returns the BucketRegion field if non-nil, zero value otherwise.

### GetBucketRegionOk

`func (o *TableflowV1ByobAwsSpec) GetBucketRegionOk() (*string, bool)`

GetBucketRegionOk returns a tuple with the BucketRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRegion

`func (o *TableflowV1ByobAwsSpec) SetBucketRegion(v string)`

SetBucketRegion sets BucketRegion field to given value.

### HasBucketRegion

`func (o *TableflowV1ByobAwsSpec) HasBucketRegion() bool`

HasBucketRegion returns a boolean if a field has been set.

### GetProviderIntegrationId

`func (o *TableflowV1ByobAwsSpec) GetProviderIntegrationId() string`

GetProviderIntegrationId returns the ProviderIntegrationId field if non-nil, zero value otherwise.

### GetProviderIntegrationIdOk

`func (o *TableflowV1ByobAwsSpec) GetProviderIntegrationIdOk() (*string, bool)`

GetProviderIntegrationIdOk returns a tuple with the ProviderIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIntegrationId

`func (o *TableflowV1ByobAwsSpec) SetProviderIntegrationId(v string)`

SetProviderIntegrationId sets ProviderIntegrationId field to given value.


### GetTablePath

`func (o *TableflowV1ByobAwsSpec) GetTablePath() string`

GetTablePath returns the TablePath field if non-nil, zero value otherwise.

### GetTablePathOk

`func (o *TableflowV1ByobAwsSpec) GetTablePathOk() (*string, bool)`

GetTablePathOk returns a tuple with the TablePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablePath

`func (o *TableflowV1ByobAwsSpec) SetTablePath(v string)`

SetTablePath sets TablePath field to given value.

### HasTablePath

`func (o *TableflowV1ByobAwsSpec) HasTablePath() bool`

HasTablePath returns a boolean if a field has been set.


### AsTableflowV1TableflowTopicSpecStorageOneOf

`func (s *TableflowV1ByobAwsSpec) AsTableflowV1TableflowTopicSpecStorageOneOf() TableflowV1TableflowTopicSpecStorageOneOf`

Convenience method to wrap this instance of TableflowV1ByobAwsSpec in TableflowV1TableflowTopicSpecStorageOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



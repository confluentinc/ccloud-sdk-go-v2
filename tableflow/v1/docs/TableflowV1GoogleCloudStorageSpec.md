# TableflowV1GoogleCloudStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The storage type.  | 
**BucketName** | **string** | Bucket name | 
**BucketRegion** | Pointer to **string** | Bucket region | [optional] [readonly] 
**ProviderIntegrationId** | **string** | The provider integration id | 
**TablePath** | Pointer to **string** | The current storage path where the data and metadata is stored for this table | [optional] [readonly] 

## Methods

### NewTableflowV1GoogleCloudStorageSpec

`func NewTableflowV1GoogleCloudStorageSpec(kind string, bucketName string, providerIntegrationId string, ) *TableflowV1GoogleCloudStorageSpec`

NewTableflowV1GoogleCloudStorageSpec instantiates a new TableflowV1GoogleCloudStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableflowV1GoogleCloudStorageSpecWithDefaults

`func NewTableflowV1GoogleCloudStorageSpecWithDefaults() *TableflowV1GoogleCloudStorageSpec`

NewTableflowV1GoogleCloudStorageSpecWithDefaults instantiates a new TableflowV1GoogleCloudStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TableflowV1GoogleCloudStorageSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TableflowV1GoogleCloudStorageSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TableflowV1GoogleCloudStorageSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetBucketName

`func (o *TableflowV1GoogleCloudStorageSpec) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *TableflowV1GoogleCloudStorageSpec) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *TableflowV1GoogleCloudStorageSpec) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetBucketRegion

`func (o *TableflowV1GoogleCloudStorageSpec) GetBucketRegion() string`

GetBucketRegion returns the BucketRegion field if non-nil, zero value otherwise.

### GetBucketRegionOk

`func (o *TableflowV1GoogleCloudStorageSpec) GetBucketRegionOk() (*string, bool)`

GetBucketRegionOk returns a tuple with the BucketRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRegion

`func (o *TableflowV1GoogleCloudStorageSpec) SetBucketRegion(v string)`

SetBucketRegion sets BucketRegion field to given value.

### HasBucketRegion

`func (o *TableflowV1GoogleCloudStorageSpec) HasBucketRegion() bool`

HasBucketRegion returns a boolean if a field has been set.

### GetProviderIntegrationId

`func (o *TableflowV1GoogleCloudStorageSpec) GetProviderIntegrationId() string`

GetProviderIntegrationId returns the ProviderIntegrationId field if non-nil, zero value otherwise.

### GetProviderIntegrationIdOk

`func (o *TableflowV1GoogleCloudStorageSpec) GetProviderIntegrationIdOk() (*string, bool)`

GetProviderIntegrationIdOk returns a tuple with the ProviderIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIntegrationId

`func (o *TableflowV1GoogleCloudStorageSpec) SetProviderIntegrationId(v string)`

SetProviderIntegrationId sets ProviderIntegrationId field to given value.


### GetTablePath

`func (o *TableflowV1GoogleCloudStorageSpec) GetTablePath() string`

GetTablePath returns the TablePath field if non-nil, zero value otherwise.

### GetTablePathOk

`func (o *TableflowV1GoogleCloudStorageSpec) GetTablePathOk() (*string, bool)`

GetTablePathOk returns a tuple with the TablePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablePath

`func (o *TableflowV1GoogleCloudStorageSpec) SetTablePath(v string)`

SetTablePath sets TablePath field to given value.

### HasTablePath

`func (o *TableflowV1GoogleCloudStorageSpec) HasTablePath() bool`

HasTablePath returns a boolean if a field has been set.


### AsTableflowV1TableflowTopicSpecStorageOneOf

`func (s *TableflowV1GoogleCloudStorageSpec) AsTableflowV1TableflowTopicSpecStorageOneOf() TableflowV1TableflowTopicSpecStorageOneOf`

Convenience method to wrap this instance of TableflowV1GoogleCloudStorageSpec in TableflowV1TableflowTopicSpecStorageOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



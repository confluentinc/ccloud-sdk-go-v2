# QueryV1QueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogName** | **string** | Confluent Cloud environment ID (e.g., env-xyz123) | 
**DatabaseName** | **string** | Kafka cluster ID (e.g., lkc-xyz123) | 
**Options** | Pointer to [**QueryV1QueryOptions**](QueryV1QueryOptions.md) | Options for this query. | [optional] 
**Query** | **string** | SQL query to execute | 
**ClientInfo** | Pointer to **map[string]string** | Free-form client metadata for telemetry and debugging (e.g., client_name, client_version, driver_version). Values are strings; not interpreted by the server.  | [optional] 

## Methods

### NewQueryV1QueryRequest

`func NewQueryV1QueryRequest(catalogName string, databaseName string, query string, ) *QueryV1QueryRequest`

NewQueryV1QueryRequest instantiates a new QueryV1QueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryV1QueryRequestWithDefaults

`func NewQueryV1QueryRequestWithDefaults() *QueryV1QueryRequest`

NewQueryV1QueryRequestWithDefaults instantiates a new QueryV1QueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogName

`func (o *QueryV1QueryRequest) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *QueryV1QueryRequest) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *QueryV1QueryRequest) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### GetDatabaseName

`func (o *QueryV1QueryRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *QueryV1QueryRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *QueryV1QueryRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetOptions

`func (o *QueryV1QueryRequest) GetOptions() QueryV1QueryOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *QueryV1QueryRequest) GetOptionsOk() (*QueryV1QueryOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *QueryV1QueryRequest) SetOptions(v QueryV1QueryOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *QueryV1QueryRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetQuery

`func (o *QueryV1QueryRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryV1QueryRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryV1QueryRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetClientInfo

`func (o *QueryV1QueryRequest) GetClientInfo() map[string]string`

GetClientInfo returns the ClientInfo field if non-nil, zero value otherwise.

### GetClientInfoOk

`func (o *QueryV1QueryRequest) GetClientInfoOk() (*map[string]string, bool)`

GetClientInfoOk returns a tuple with the ClientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientInfo

`func (o *QueryV1QueryRequest) SetClientInfo(v map[string]string)`

SetClientInfo sets ClientInfo field to given value.

### HasClientInfo

`func (o *QueryV1QueryRequest) HasClientInfo() bool`

HasClientInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



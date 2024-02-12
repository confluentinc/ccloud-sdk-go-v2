# ListTablesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifiers** | Pointer to [**[]TableIdentifier**](TableIdentifier.md) |  | [optional] 

## Methods

### NewListTablesResponse

`func NewListTablesResponse() *ListTablesResponse`

NewListTablesResponse instantiates a new ListTablesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTablesResponseWithDefaults

`func NewListTablesResponseWithDefaults() *ListTablesResponse`

NewListTablesResponseWithDefaults instantiates a new ListTablesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifiers

`func (o *ListTablesResponse) GetIdentifiers() []TableIdentifier`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ListTablesResponse) GetIdentifiersOk() (*[]TableIdentifier, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ListTablesResponse) SetIdentifiers(v []TableIdentifier)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *ListTablesResponse) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



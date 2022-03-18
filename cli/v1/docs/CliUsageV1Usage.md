# CliUsageV1Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Os** | Pointer to **string** | Operating system of the CLI binary used to run the command | [optional] 
**Arch** | Pointer to **string** | Architecture of the CLI binary used to run the command | [optional] 
**Version** | Pointer to **string** | Version of the CLI binary used to run the command | [optional] 
**Command** | Pointer to **string** | CLI command that was run | [optional] 
**Flags** | Pointer to **[]string** | Names of the flags passed with the CLI command | [optional] 
**Error** | Pointer to **bool** | If an error occurred while running the CLI command | [optional] 

## Methods

### NewCliUsageV1Usage

`func NewCliUsageV1Usage() *CliUsageV1Usage`

NewCliUsageV1Usage instantiates a new CliUsageV1Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCliUsageV1UsageWithDefaults

`func NewCliUsageV1UsageWithDefaults() *CliUsageV1Usage`

NewCliUsageV1UsageWithDefaults instantiates a new CliUsageV1Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CliUsageV1Usage) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CliUsageV1Usage) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CliUsageV1Usage) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CliUsageV1Usage) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *CliUsageV1Usage) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CliUsageV1Usage) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CliUsageV1Usage) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CliUsageV1Usage) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *CliUsageV1Usage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CliUsageV1Usage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CliUsageV1Usage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CliUsageV1Usage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *CliUsageV1Usage) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CliUsageV1Usage) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CliUsageV1Usage) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CliUsageV1Usage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOs

`func (o *CliUsageV1Usage) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *CliUsageV1Usage) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *CliUsageV1Usage) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *CliUsageV1Usage) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetArch

`func (o *CliUsageV1Usage) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CliUsageV1Usage) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CliUsageV1Usage) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CliUsageV1Usage) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetVersion

`func (o *CliUsageV1Usage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CliUsageV1Usage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CliUsageV1Usage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CliUsageV1Usage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCommand

`func (o *CliUsageV1Usage) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CliUsageV1Usage) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CliUsageV1Usage) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CliUsageV1Usage) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetFlags

`func (o *CliUsageV1Usage) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *CliUsageV1Usage) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *CliUsageV1Usage) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *CliUsageV1Usage) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetError

`func (o *CliUsageV1Usage) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CliUsageV1Usage) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CliUsageV1Usage) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *CliUsageV1Usage) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



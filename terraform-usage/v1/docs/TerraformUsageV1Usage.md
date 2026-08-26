# TerraformUsageV1Usage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | Pointer to **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [optional] [readonly] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**RunId** | Pointer to **string** | Identifier generated once per provider process and shared by every operation that process reports.  This identifier is process-scoped, not workflow-scoped. Terraform Core launches a separate provider process for the plan phase and for the apply phase, and a separate process for each aliased &#x60;provider&#x60; block, so a &#x60;run_id&#x60; does not correlate a plan with its paired apply, nor one aliased provider instance with another. Consumers must not treat it as a plan, apply, workflow, or session identifier.  The value is randomly generated per process. It must not be derived from the Terraform workspace, module path, or any other user-authored name.  | [optional] 
**Sequence** | Pointer to **int32** | Position of this operation in the order the provider started operations within its &#x60;run_id&#x60;, counting from 1.  Values are unique within a &#x60;run_id&#x60; and increase in the order operations were started, but the set of events the API receives is neither ordered nor complete. Terraform walks the resource graph concurrently and each event is reported only after its operation finishes, so a lower sequence number can arrive after a higher one. Reporting is best-effort and the provider stops reporting after a fixed number of events in a run, so numbers can be missing and the highest value received can undercount the operations the run actually performed.  | [optional] 
**StartedAt** | Pointer to **time.Time** | The date and time at which the operation began. It is represented in RFC3339 format and is in UTC. | [optional] 
**DurationMs** | Pointer to **int32** | How long the operation took, in milliseconds, measured by the provider from the start of the operation until just before it returned | [optional] 
**Os** | Pointer to **string** | Operating system of the provider binary that ran the operation, as reported by the Go runtime&#39;s &#x60;GOOS&#x60; | [optional] 
**Arch** | Pointer to **string** | Architecture of the provider binary that ran the operation, as reported by the Go runtime&#39;s &#x60;GOARCH&#x60; | [optional] 
**ProviderVersion** | Pointer to **string** | Version of the Confluent Terraform Provider that ran the operation. One provider process loads exactly one provider binary, so this value is constant for a given &#x60;run_id&#x60; | [optional] 
**TerraformVersion** | Pointer to **string** | Version of Terraform that launched the provider process, as reported to the provider by Terraform Core | [optional] 
**ResourceType** | Pointer to **string** | Terraform resource type that the operation was performed on. This is the resource type the provider registers, never a Terraform resource address, whose second component is a user-authored name | [optional] 
**Operation** | Pointer to **string** | Operation that the provider performed on the resource. Values outside the listed set are accepted so that a newer provider can report an operation this version of the API does not yet name | [optional] 
**ChangedAttributes** | Pointer to **[]string** | Names of the resource&#39;s attributes that the operation changed. Empty for operations that change no attributes, such as &#x60;READ&#x60; and &#x60;IMPORT&#x60;.  Only top-level attribute names declared in the provider&#39;s resource schema are reported. Attribute values are never reported, and a map-typed attribute contributes only its own declared name, never the map&#39;s runtime keys or values. Elements must be valid Terraform SDK field names, which rejects dotted paths such as &#x60;config.cleanup.policy&#x60; and anything containing whitespace or punctuation.  | [optional] 
**Error** | Pointer to **bool** | If an error occurred while running the operation. Also set when the operation panicked, in which case &#x60;stack_frames&#x60; is populated | [optional] 
**StackFrames** | Pointer to **[]string** | Source locations from the stack trace of a panic, reported only when the operation panicked.  Each entry is a Go source location of the form &#x60;&lt;path&gt;:&lt;line&gt;&#x60;, taken from the provider&#39;s own stack trace and filtered to Go source lines. Panic messages, function arguments, and resource state are not included. Paths are recorded when the provider binary is compiled, not when it runs, so they do not describe the machine running Terraform, unless the binary was built locally rather than from an official release, in which case they are paths on the machine that built it. Official release builds trim the module prefix but not the Go runtime prefix, so runtime frames still carry an absolute path from the release build machine.  | [optional] 

## Methods

### NewTerraformUsageV1Usage

`func NewTerraformUsageV1Usage() *TerraformUsageV1Usage`

NewTerraformUsageV1Usage instantiates a new TerraformUsageV1Usage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformUsageV1UsageWithDefaults

`func NewTerraformUsageV1UsageWithDefaults() *TerraformUsageV1Usage`

NewTerraformUsageV1UsageWithDefaults instantiates a new TerraformUsageV1Usage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *TerraformUsageV1Usage) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TerraformUsageV1Usage) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TerraformUsageV1Usage) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TerraformUsageV1Usage) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *TerraformUsageV1Usage) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TerraformUsageV1Usage) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TerraformUsageV1Usage) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TerraformUsageV1Usage) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *TerraformUsageV1Usage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerraformUsageV1Usage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerraformUsageV1Usage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TerraformUsageV1Usage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *TerraformUsageV1Usage) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TerraformUsageV1Usage) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TerraformUsageV1Usage) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TerraformUsageV1Usage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRunId

`func (o *TerraformUsageV1Usage) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *TerraformUsageV1Usage) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *TerraformUsageV1Usage) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *TerraformUsageV1Usage) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetSequence

`func (o *TerraformUsageV1Usage) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *TerraformUsageV1Usage) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *TerraformUsageV1Usage) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *TerraformUsageV1Usage) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetStartedAt

`func (o *TerraformUsageV1Usage) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TerraformUsageV1Usage) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TerraformUsageV1Usage) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TerraformUsageV1Usage) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetDurationMs

`func (o *TerraformUsageV1Usage) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *TerraformUsageV1Usage) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *TerraformUsageV1Usage) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *TerraformUsageV1Usage) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetOs

`func (o *TerraformUsageV1Usage) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *TerraformUsageV1Usage) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *TerraformUsageV1Usage) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *TerraformUsageV1Usage) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetArch

`func (o *TerraformUsageV1Usage) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *TerraformUsageV1Usage) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *TerraformUsageV1Usage) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *TerraformUsageV1Usage) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetProviderVersion

`func (o *TerraformUsageV1Usage) GetProviderVersion() string`

GetProviderVersion returns the ProviderVersion field if non-nil, zero value otherwise.

### GetProviderVersionOk

`func (o *TerraformUsageV1Usage) GetProviderVersionOk() (*string, bool)`

GetProviderVersionOk returns a tuple with the ProviderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVersion

`func (o *TerraformUsageV1Usage) SetProviderVersion(v string)`

SetProviderVersion sets ProviderVersion field to given value.

### HasProviderVersion

`func (o *TerraformUsageV1Usage) HasProviderVersion() bool`

HasProviderVersion returns a boolean if a field has been set.

### GetTerraformVersion

`func (o *TerraformUsageV1Usage) GetTerraformVersion() string`

GetTerraformVersion returns the TerraformVersion field if non-nil, zero value otherwise.

### GetTerraformVersionOk

`func (o *TerraformUsageV1Usage) GetTerraformVersionOk() (*string, bool)`

GetTerraformVersionOk returns a tuple with the TerraformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformVersion

`func (o *TerraformUsageV1Usage) SetTerraformVersion(v string)`

SetTerraformVersion sets TerraformVersion field to given value.

### HasTerraformVersion

`func (o *TerraformUsageV1Usage) HasTerraformVersion() bool`

HasTerraformVersion returns a boolean if a field has been set.

### GetResourceType

`func (o *TerraformUsageV1Usage) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TerraformUsageV1Usage) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TerraformUsageV1Usage) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *TerraformUsageV1Usage) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetOperation

`func (o *TerraformUsageV1Usage) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TerraformUsageV1Usage) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TerraformUsageV1Usage) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *TerraformUsageV1Usage) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetChangedAttributes

`func (o *TerraformUsageV1Usage) GetChangedAttributes() []string`

GetChangedAttributes returns the ChangedAttributes field if non-nil, zero value otherwise.

### GetChangedAttributesOk

`func (o *TerraformUsageV1Usage) GetChangedAttributesOk() (*[]string, bool)`

GetChangedAttributesOk returns a tuple with the ChangedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAttributes

`func (o *TerraformUsageV1Usage) SetChangedAttributes(v []string)`

SetChangedAttributes sets ChangedAttributes field to given value.

### HasChangedAttributes

`func (o *TerraformUsageV1Usage) HasChangedAttributes() bool`

HasChangedAttributes returns a boolean if a field has been set.

### GetError

`func (o *TerraformUsageV1Usage) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TerraformUsageV1Usage) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TerraformUsageV1Usage) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *TerraformUsageV1Usage) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStackFrames

`func (o *TerraformUsageV1Usage) GetStackFrames() []string`

GetStackFrames returns the StackFrames field if non-nil, zero value otherwise.

### GetStackFramesOk

`func (o *TerraformUsageV1Usage) GetStackFramesOk() (*[]string, bool)`

GetStackFramesOk returns a tuple with the StackFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackFrames

`func (o *TerraformUsageV1Usage) SetStackFrames(v []string)`

SetStackFrames sets StackFrames field to given value.

### HasStackFrames

`func (o *TerraformUsageV1Usage) HasStackFrames() bool`

HasStackFrames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



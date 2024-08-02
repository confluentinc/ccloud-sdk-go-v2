# ArtifactV1FlinkArtifactVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version id of the Flink Artifact. | 
**ReleaseNotes** | Pointer to **string** | Release Notes of the Flink Artifact version. | [optional] 
**IsBeta** | Pointer to **bool** | Flag to specify stability of the version | [optional] 
**ArtifactId** | [**ArtifactV1FlinkArtifact**](artifact.v1.FlinkArtifact.md) | The Flink Artifact this version belongs to. | 
**UploadSource** | [**ArtifactV1FlinkArtifactVersionUploadSourceOneOf**](ArtifactV1FlinkArtifactVersionUploadSourceOneOf.md) | Upload source of the Flink Artifact Version. | 

## Methods

### NewArtifactV1FlinkArtifactVersion

`func NewArtifactV1FlinkArtifactVersion(version string, artifactId ArtifactV1FlinkArtifact, uploadSource ArtifactV1FlinkArtifactVersionUploadSourceOneOf, ) *ArtifactV1FlinkArtifactVersion`

NewArtifactV1FlinkArtifactVersion instantiates a new ArtifactV1FlinkArtifactVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactV1FlinkArtifactVersionWithDefaults

`func NewArtifactV1FlinkArtifactVersionWithDefaults() *ArtifactV1FlinkArtifactVersion`

NewArtifactV1FlinkArtifactVersionWithDefaults instantiates a new ArtifactV1FlinkArtifactVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ArtifactV1FlinkArtifactVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactV1FlinkArtifactVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactV1FlinkArtifactVersion) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetReleaseNotes

`func (o *ArtifactV1FlinkArtifactVersion) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *ArtifactV1FlinkArtifactVersion) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *ArtifactV1FlinkArtifactVersion) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *ArtifactV1FlinkArtifactVersion) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### GetIsBeta

`func (o *ArtifactV1FlinkArtifactVersion) GetIsBeta() bool`

GetIsBeta returns the IsBeta field if non-nil, zero value otherwise.

### GetIsBetaOk

`func (o *ArtifactV1FlinkArtifactVersion) GetIsBetaOk() (*bool, bool)`

GetIsBetaOk returns a tuple with the IsBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBeta

`func (o *ArtifactV1FlinkArtifactVersion) SetIsBeta(v bool)`

SetIsBeta sets IsBeta field to given value.

### HasIsBeta

`func (o *ArtifactV1FlinkArtifactVersion) HasIsBeta() bool`

HasIsBeta returns a boolean if a field has been set.

### GetArtifactId

`func (o *ArtifactV1FlinkArtifactVersion) GetArtifactId() ArtifactV1FlinkArtifact`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *ArtifactV1FlinkArtifactVersion) GetArtifactIdOk() (*ArtifactV1FlinkArtifact, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *ArtifactV1FlinkArtifactVersion) SetArtifactId(v ArtifactV1FlinkArtifact)`

SetArtifactId sets ArtifactId field to given value.


### GetUploadSource

`func (o *ArtifactV1FlinkArtifactVersion) GetUploadSource() ArtifactV1FlinkArtifactVersionUploadSourceOneOf`

GetUploadSource returns the UploadSource field if non-nil, zero value otherwise.

### GetUploadSourceOk

`func (o *ArtifactV1FlinkArtifactVersion) GetUploadSourceOk() (*ArtifactV1FlinkArtifactVersionUploadSourceOneOf, bool)`

GetUploadSourceOk returns a tuple with the UploadSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSource

`func (o *ArtifactV1FlinkArtifactVersion) SetUploadSource(v ArtifactV1FlinkArtifactVersionUploadSourceOneOf)`

SetUploadSource sets UploadSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



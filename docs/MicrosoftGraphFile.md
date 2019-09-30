# MicrosoftGraphFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hashes** | Pointer to [**AnyOfmicrosoftGraphHashes**](anyOf&lt;microsoft.graph.hashes&gt;.md) |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**ProcessingMetadata** | Pointer to **bool** |  | [optional] 

## Methods

### GetHashes

`func (o *MicrosoftGraphFile) GetHashes() AnyOfmicrosoftGraphHashes`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *MicrosoftGraphFile) GetHashesOk() (AnyOfmicrosoftGraphHashes, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHashes

`func (o *MicrosoftGraphFile) HasHashes() bool`

HasHashes returns a boolean if a field has been set.

### SetHashes

`func (o *MicrosoftGraphFile) SetHashes(v AnyOfmicrosoftGraphHashes)`

SetHashes gets a reference to the given AnyOfmicrosoftGraphHashes and assigns it to the Hashes field.

### SetHashesExplicitNull

`func (o *MicrosoftGraphFile) SetHashesExplicitNull(b bool)`

SetHashesExplicitNull (un)sets Hashes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Hashes value is set to nil even if false is passed
### GetMimeType

`func (o *MicrosoftGraphFile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *MicrosoftGraphFile) GetMimeTypeOk() (string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMimeType

`func (o *MicrosoftGraphFile) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeType

`func (o *MicrosoftGraphFile) SetMimeType(v string)`

SetMimeType gets a reference to the given string and assigns it to the MimeType field.

### SetMimeTypeExplicitNull

`func (o *MicrosoftGraphFile) SetMimeTypeExplicitNull(b bool)`

SetMimeTypeExplicitNull (un)sets MimeType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MimeType value is set to nil even if false is passed
### GetProcessingMetadata

`func (o *MicrosoftGraphFile) GetProcessingMetadata() bool`

GetProcessingMetadata returns the ProcessingMetadata field if non-nil, zero value otherwise.

### GetProcessingMetadataOk

`func (o *MicrosoftGraphFile) GetProcessingMetadataOk() (bool, bool)`

GetProcessingMetadataOk returns a tuple with the ProcessingMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcessingMetadata

`func (o *MicrosoftGraphFile) HasProcessingMetadata() bool`

HasProcessingMetadata returns a boolean if a field has been set.

### SetProcessingMetadata

`func (o *MicrosoftGraphFile) SetProcessingMetadata(v bool)`

SetProcessingMetadata gets a reference to the given bool and assigns it to the ProcessingMetadata field.

### SetProcessingMetadataExplicitNull

`func (o *MicrosoftGraphFile) SetProcessingMetadataExplicitNull(b bool)`

SetProcessingMetadataExplicitNull (un)sets ProcessingMetadata to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProcessingMetadata value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



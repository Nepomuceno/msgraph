# MicrosoftGraphFileEncryptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionKey** | Pointer to **string** | The key used to encrypt the file content. | [optional] 
**InitializationVector** | Pointer to **string** | The initialization vector used for the encryption algorithm. | [optional] 
**Mac** | Pointer to **string** | The hash of the encrypted file content + IV (content hash). | [optional] 
**MacKey** | Pointer to **string** | The key used to get mac. | [optional] 
**ProfileIdentifier** | Pointer to **string** | The the profile identifier. | [optional] 
**FileDigest** | Pointer to **string** | The file digest prior to encryption. | [optional] 
**FileDigestAlgorithm** | Pointer to **string** | The file digest algorithm. | [optional] 

## Methods

### GetEncryptionKey

`func (o *MicrosoftGraphFileEncryptionInfo) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetEncryptionKeyOk() (string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEncryptionKey

`func (o *MicrosoftGraphFileEncryptionInfo) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKey

`func (o *MicrosoftGraphFileEncryptionInfo) SetEncryptionKey(v string)`

SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.

### SetEncryptionKeyExplicitNull

`func (o *MicrosoftGraphFileEncryptionInfo) SetEncryptionKeyExplicitNull(b bool)`

SetEncryptionKeyExplicitNull (un)sets EncryptionKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EncryptionKey value is set to nil even if false is passed
### GetInitializationVector

`func (o *MicrosoftGraphFileEncryptionInfo) GetInitializationVector() string`

GetInitializationVector returns the InitializationVector field if non-nil, zero value otherwise.

### GetInitializationVectorOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetInitializationVectorOk() (string, bool)`

GetInitializationVectorOk returns a tuple with the InitializationVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInitializationVector

`func (o *MicrosoftGraphFileEncryptionInfo) HasInitializationVector() bool`

HasInitializationVector returns a boolean if a field has been set.

### SetInitializationVector

`func (o *MicrosoftGraphFileEncryptionInfo) SetInitializationVector(v string)`

SetInitializationVector gets a reference to the given string and assigns it to the InitializationVector field.

### SetInitializationVectorExplicitNull

`func (o *MicrosoftGraphFileEncryptionInfo) SetInitializationVectorExplicitNull(b bool)`

SetInitializationVectorExplicitNull (un)sets InitializationVector to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InitializationVector value is set to nil even if false is passed
### GetMac

`func (o *MicrosoftGraphFileEncryptionInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetMacOk() (string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMac

`func (o *MicrosoftGraphFileEncryptionInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### SetMac

`func (o *MicrosoftGraphFileEncryptionInfo) SetMac(v string)`

SetMac gets a reference to the given string and assigns it to the Mac field.

### SetMacExplicitNull

`func (o *MicrosoftGraphFileEncryptionInfo) SetMacExplicitNull(b bool)`

SetMacExplicitNull (un)sets Mac to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Mac value is set to nil even if false is passed
### GetMacKey

`func (o *MicrosoftGraphFileEncryptionInfo) GetMacKey() string`

GetMacKey returns the MacKey field if non-nil, zero value otherwise.

### GetMacKeyOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetMacKeyOk() (string, bool)`

GetMacKeyOk returns a tuple with the MacKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMacKey

`func (o *MicrosoftGraphFileEncryptionInfo) HasMacKey() bool`

HasMacKey returns a boolean if a field has been set.

### SetMacKey

`func (o *MicrosoftGraphFileEncryptionInfo) SetMacKey(v string)`

SetMacKey gets a reference to the given string and assigns it to the MacKey field.

### SetMacKeyExplicitNull

`func (o *MicrosoftGraphFileEncryptionInfo) SetMacKeyExplicitNull(b bool)`

SetMacKeyExplicitNull (un)sets MacKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MacKey value is set to nil even if false is passed
### GetProfileIdentifier

`func (o *MicrosoftGraphFileEncryptionInfo) GetProfileIdentifier() string`

GetProfileIdentifier returns the ProfileIdentifier field if non-nil, zero value otherwise.

### GetProfileIdentifierOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetProfileIdentifierOk() (string, bool)`

GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileIdentifier

`func (o *MicrosoftGraphFileEncryptionInfo) HasProfileIdentifier() bool`

HasProfileIdentifier returns a boolean if a field has been set.

### SetProfileIdentifier

`func (o *MicrosoftGraphFileEncryptionInfo) SetProfileIdentifier(v string)`

SetProfileIdentifier gets a reference to the given string and assigns it to the ProfileIdentifier field.

### SetProfileIdentifierExplicitNull

`func (o *MicrosoftGraphFileEncryptionInfo) SetProfileIdentifierExplicitNull(b bool)`

SetProfileIdentifierExplicitNull (un)sets ProfileIdentifier to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ProfileIdentifier value is set to nil even if false is passed
### GetFileDigest

`func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigest() string`

GetFileDigest returns the FileDigest field if non-nil, zero value otherwise.

### GetFileDigestOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigestOk() (string, bool)`

GetFileDigestOk returns a tuple with the FileDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileDigest

`func (o *MicrosoftGraphFileEncryptionInfo) HasFileDigest() bool`

HasFileDigest returns a boolean if a field has been set.

### SetFileDigest

`func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigest(v string)`

SetFileDigest gets a reference to the given string and assigns it to the FileDigest field.

### SetFileDigestExplicitNull

`func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigestExplicitNull(b bool)`

SetFileDigestExplicitNull (un)sets FileDigest to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileDigest value is set to nil even if false is passed
### GetFileDigestAlgorithm

`func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigestAlgorithm() string`

GetFileDigestAlgorithm returns the FileDigestAlgorithm field if non-nil, zero value otherwise.

### GetFileDigestAlgorithmOk

`func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigestAlgorithmOk() (string, bool)`

GetFileDigestAlgorithmOk returns a tuple with the FileDigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileDigestAlgorithm

`func (o *MicrosoftGraphFileEncryptionInfo) HasFileDigestAlgorithm() bool`

HasFileDigestAlgorithm returns a boolean if a field has been set.

### SetFileDigestAlgorithm

`func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigestAlgorithm(v string)`

SetFileDigestAlgorithm gets a reference to the given string and assigns it to the FileDigestAlgorithm field.

### SetFileDigestAlgorithmExplicitNull

`func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigestAlgorithmExplicitNull(b bool)`

SetFileDigestAlgorithmExplicitNull (un)sets FileDigestAlgorithm to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileDigestAlgorithm value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



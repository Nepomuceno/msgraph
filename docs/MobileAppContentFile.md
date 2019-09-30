# MobileAppContentFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureStorageUri** | Pointer to **string** | The Azure Storage URI. | [optional] 
**IsCommitted** | Pointer to **bool** | A value indicating whether the file is committed. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | The time the file was created. | [optional] 
**Name** | Pointer to **string** | the file name. | [optional] 
**Size** | Pointer to **int64** | The size of the file prior to encryption. | [optional] 
**SizeEncrypted** | Pointer to **int64** | The size of the file after encryption. | [optional] 
**AzureStorageUriExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) | The time the Azure storage Uri expires. | [optional] 
**Manifest** | Pointer to **string** | The manifest information. | [optional] 
**UploadState** | Pointer to [**AnyOfmicrosoftGraphMobileAppContentFileUploadState**](anyOf&lt;microsoft.graph.mobileAppContentFileUploadState&gt;.md) | The state of the current upload request. | [optional] 

## Methods

### GetAzureStorageUri

`func (o *MobileAppContentFile) GetAzureStorageUri() string`

GetAzureStorageUri returns the AzureStorageUri field if non-nil, zero value otherwise.

### GetAzureStorageUriOk

`func (o *MobileAppContentFile) GetAzureStorageUriOk() (string, bool)`

GetAzureStorageUriOk returns a tuple with the AzureStorageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureStorageUri

`func (o *MobileAppContentFile) HasAzureStorageUri() bool`

HasAzureStorageUri returns a boolean if a field has been set.

### SetAzureStorageUri

`func (o *MobileAppContentFile) SetAzureStorageUri(v string)`

SetAzureStorageUri gets a reference to the given string and assigns it to the AzureStorageUri field.

### SetAzureStorageUriExplicitNull

`func (o *MobileAppContentFile) SetAzureStorageUriExplicitNull(b bool)`

SetAzureStorageUriExplicitNull (un)sets AzureStorageUri to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureStorageUri value is set to nil even if false is passed
### GetIsCommitted

`func (o *MobileAppContentFile) GetIsCommitted() bool`

GetIsCommitted returns the IsCommitted field if non-nil, zero value otherwise.

### GetIsCommittedOk

`func (o *MobileAppContentFile) GetIsCommittedOk() (bool, bool)`

GetIsCommittedOk returns a tuple with the IsCommitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsCommitted

`func (o *MobileAppContentFile) HasIsCommitted() bool`

HasIsCommitted returns a boolean if a field has been set.

### SetIsCommitted

`func (o *MobileAppContentFile) SetIsCommitted(v bool)`

SetIsCommitted gets a reference to the given bool and assigns it to the IsCommitted field.

### GetCreatedDateTime

`func (o *MobileAppContentFile) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MobileAppContentFile) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MobileAppContentFile) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MobileAppContentFile) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetName

`func (o *MobileAppContentFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileAppContentFile) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MobileAppContentFile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MobileAppContentFile) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MobileAppContentFile) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetSize

`func (o *MobileAppContentFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MobileAppContentFile) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *MobileAppContentFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *MobileAppContentFile) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### GetSizeEncrypted

`func (o *MobileAppContentFile) GetSizeEncrypted() int64`

GetSizeEncrypted returns the SizeEncrypted field if non-nil, zero value otherwise.

### GetSizeEncryptedOk

`func (o *MobileAppContentFile) GetSizeEncryptedOk() (int64, bool)`

GetSizeEncryptedOk returns a tuple with the SizeEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSizeEncrypted

`func (o *MobileAppContentFile) HasSizeEncrypted() bool`

HasSizeEncrypted returns a boolean if a field has been set.

### SetSizeEncrypted

`func (o *MobileAppContentFile) SetSizeEncrypted(v int64)`

SetSizeEncrypted gets a reference to the given int64 and assigns it to the SizeEncrypted field.

### GetAzureStorageUriExpirationDateTime

`func (o *MobileAppContentFile) GetAzureStorageUriExpirationDateTime() time.Time`

GetAzureStorageUriExpirationDateTime returns the AzureStorageUriExpirationDateTime field if non-nil, zero value otherwise.

### GetAzureStorageUriExpirationDateTimeOk

`func (o *MobileAppContentFile) GetAzureStorageUriExpirationDateTimeOk() (time.Time, bool)`

GetAzureStorageUriExpirationDateTimeOk returns a tuple with the AzureStorageUriExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureStorageUriExpirationDateTime

`func (o *MobileAppContentFile) HasAzureStorageUriExpirationDateTime() bool`

HasAzureStorageUriExpirationDateTime returns a boolean if a field has been set.

### SetAzureStorageUriExpirationDateTime

`func (o *MobileAppContentFile) SetAzureStorageUriExpirationDateTime(v time.Time)`

SetAzureStorageUriExpirationDateTime gets a reference to the given time.Time and assigns it to the AzureStorageUriExpirationDateTime field.

### SetAzureStorageUriExpirationDateTimeExplicitNull

`func (o *MobileAppContentFile) SetAzureStorageUriExpirationDateTimeExplicitNull(b bool)`

SetAzureStorageUriExpirationDateTimeExplicitNull (un)sets AzureStorageUriExpirationDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AzureStorageUriExpirationDateTime value is set to nil even if false is passed
### GetManifest

`func (o *MobileAppContentFile) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *MobileAppContentFile) GetManifestOk() (string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManifest

`func (o *MobileAppContentFile) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### SetManifest

`func (o *MobileAppContentFile) SetManifest(v string)`

SetManifest gets a reference to the given string and assigns it to the Manifest field.

### SetManifestExplicitNull

`func (o *MobileAppContentFile) SetManifestExplicitNull(b bool)`

SetManifestExplicitNull (un)sets Manifest to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Manifest value is set to nil even if false is passed
### GetUploadState

`func (o *MobileAppContentFile) GetUploadState() AnyOfmicrosoftGraphMobileAppContentFileUploadState`

GetUploadState returns the UploadState field if non-nil, zero value otherwise.

### GetUploadStateOk

`func (o *MobileAppContentFile) GetUploadStateOk() (AnyOfmicrosoftGraphMobileAppContentFileUploadState, bool)`

GetUploadStateOk returns a tuple with the UploadState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUploadState

`func (o *MobileAppContentFile) HasUploadState() bool`

HasUploadState returns a boolean if a field has been set.

### SetUploadState

`func (o *MobileAppContentFile) SetUploadState(v AnyOfmicrosoftGraphMobileAppContentFileUploadState)`

SetUploadState gets a reference to the given AnyOfmicrosoftGraphMobileAppContentFileUploadState and assigns it to the UploadState field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



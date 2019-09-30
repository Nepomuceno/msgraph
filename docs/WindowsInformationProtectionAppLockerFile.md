# WindowsInformationProtectionAppLockerFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The friendly name | [optional] 
**FileHash** | Pointer to **string** | SHA256 hash of the file | [optional] 
**File** | Pointer to **string** | File as a byte array | [optional] 
**Version** | Pointer to **string** | Version of the entity. | [optional] 

## Methods

### GetDisplayName

`func (o *WindowsInformationProtectionAppLockerFile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WindowsInformationProtectionAppLockerFile) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *WindowsInformationProtectionAppLockerFile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *WindowsInformationProtectionAppLockerFile) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *WindowsInformationProtectionAppLockerFile) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetFileHash

`func (o *WindowsInformationProtectionAppLockerFile) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *WindowsInformationProtectionAppLockerFile) GetFileHashOk() (string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileHash

`func (o *WindowsInformationProtectionAppLockerFile) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### SetFileHash

`func (o *WindowsInformationProtectionAppLockerFile) SetFileHash(v string)`

SetFileHash gets a reference to the given string and assigns it to the FileHash field.

### SetFileHashExplicitNull

`func (o *WindowsInformationProtectionAppLockerFile) SetFileHashExplicitNull(b bool)`

SetFileHashExplicitNull (un)sets FileHash to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FileHash value is set to nil even if false is passed
### GetFile

`func (o *WindowsInformationProtectionAppLockerFile) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *WindowsInformationProtectionAppLockerFile) GetFileOk() (string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFile

`func (o *WindowsInformationProtectionAppLockerFile) HasFile() bool`

HasFile returns a boolean if a field has been set.

### SetFile

`func (o *WindowsInformationProtectionAppLockerFile) SetFile(v string)`

SetFile gets a reference to the given string and assigns it to the File field.

### SetFileExplicitNull

`func (o *WindowsInformationProtectionAppLockerFile) SetFileExplicitNull(b bool)`

SetFileExplicitNull (un)sets File to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The File value is set to nil even if false is passed
### GetVersion

`func (o *WindowsInformationProtectionAppLockerFile) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WindowsInformationProtectionAppLockerFile) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *WindowsInformationProtectionAppLockerFile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *WindowsInformationProtectionAppLockerFile) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *WindowsInformationProtectionAppLockerFile) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



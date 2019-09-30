# MicrosoftGraphEBookInstallSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**InstalledDeviceCount** | Pointer to **int32** | Number of Devices that have successfully installed this book. | [optional] 
**FailedDeviceCount** | Pointer to **int32** | Number of Devices that have failed to install this book. | [optional] 
**NotInstalledDeviceCount** | Pointer to **int32** | Number of Devices that does not have this book installed. | [optional] 
**InstalledUserCount** | Pointer to **int32** | Number of Users whose devices have all succeeded to install this book. | [optional] 
**FailedUserCount** | Pointer to **int32** | Number of Users that have 1 or more device that failed to install this book. | [optional] 
**NotInstalledUserCount** | Pointer to **int32** | Number of Users that did not install this book. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphEBookInstallSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEBookInstallSummary) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphEBookInstallSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphEBookInstallSummary) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) GetInstalledDeviceCount() int32`

GetInstalledDeviceCount returns the InstalledDeviceCount field if non-nil, zero value otherwise.

### GetInstalledDeviceCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetInstalledDeviceCountOk() (int32, bool)`

GetInstalledDeviceCountOk returns a tuple with the InstalledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) HasInstalledDeviceCount() bool`

HasInstalledDeviceCount returns a boolean if a field has been set.

### SetInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) SetInstalledDeviceCount(v int32)`

SetInstalledDeviceCount gets a reference to the given int32 and assigns it to the InstalledDeviceCount field.

### GetFailedDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) GetFailedDeviceCount() int32`

GetFailedDeviceCount returns the FailedDeviceCount field if non-nil, zero value otherwise.

### GetFailedDeviceCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetFailedDeviceCountOk() (int32, bool)`

GetFailedDeviceCountOk returns a tuple with the FailedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailedDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) HasFailedDeviceCount() bool`

HasFailedDeviceCount returns a boolean if a field has been set.

### SetFailedDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) SetFailedDeviceCount(v int32)`

SetFailedDeviceCount gets a reference to the given int32 and assigns it to the FailedDeviceCount field.

### GetNotInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) GetNotInstalledDeviceCount() int32`

GetNotInstalledDeviceCount returns the NotInstalledDeviceCount field if non-nil, zero value otherwise.

### GetNotInstalledDeviceCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetNotInstalledDeviceCountOk() (int32, bool)`

GetNotInstalledDeviceCountOk returns a tuple with the NotInstalledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) HasNotInstalledDeviceCount() bool`

HasNotInstalledDeviceCount returns a boolean if a field has been set.

### SetNotInstalledDeviceCount

`func (o *MicrosoftGraphEBookInstallSummary) SetNotInstalledDeviceCount(v int32)`

SetNotInstalledDeviceCount gets a reference to the given int32 and assigns it to the NotInstalledDeviceCount field.

### GetInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) GetInstalledUserCount() int32`

GetInstalledUserCount returns the InstalledUserCount field if non-nil, zero value otherwise.

### GetInstalledUserCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetInstalledUserCountOk() (int32, bool)`

GetInstalledUserCountOk returns a tuple with the InstalledUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) HasInstalledUserCount() bool`

HasInstalledUserCount returns a boolean if a field has been set.

### SetInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) SetInstalledUserCount(v int32)`

SetInstalledUserCount gets a reference to the given int32 and assigns it to the InstalledUserCount field.

### GetFailedUserCount

`func (o *MicrosoftGraphEBookInstallSummary) GetFailedUserCount() int32`

GetFailedUserCount returns the FailedUserCount field if non-nil, zero value otherwise.

### GetFailedUserCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetFailedUserCountOk() (int32, bool)`

GetFailedUserCountOk returns a tuple with the FailedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailedUserCount

`func (o *MicrosoftGraphEBookInstallSummary) HasFailedUserCount() bool`

HasFailedUserCount returns a boolean if a field has been set.

### SetFailedUserCount

`func (o *MicrosoftGraphEBookInstallSummary) SetFailedUserCount(v int32)`

SetFailedUserCount gets a reference to the given int32 and assigns it to the FailedUserCount field.

### GetNotInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) GetNotInstalledUserCount() int32`

GetNotInstalledUserCount returns the NotInstalledUserCount field if non-nil, zero value otherwise.

### GetNotInstalledUserCountOk

`func (o *MicrosoftGraphEBookInstallSummary) GetNotInstalledUserCountOk() (int32, bool)`

GetNotInstalledUserCountOk returns a tuple with the NotInstalledUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) HasNotInstalledUserCount() bool`

HasNotInstalledUserCount returns a boolean if a field has been set.

### SetNotInstalledUserCount

`func (o *MicrosoftGraphEBookInstallSummary) SetNotInstalledUserCount(v int32)`

SetNotInstalledUserCount gets a reference to the given int32 and assigns it to the NotInstalledUserCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



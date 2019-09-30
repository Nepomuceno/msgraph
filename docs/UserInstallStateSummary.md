# UserInstallStateSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** | User name. | [optional] 
**InstalledDeviceCount** | Pointer to **int32** | Installed Device Count. | [optional] 
**FailedDeviceCount** | Pointer to **int32** | Failed Device Count. | [optional] 
**NotInstalledDeviceCount** | Pointer to **int32** | Not installed device count. | [optional] 
**DeviceStates** | Pointer to [**[]MicrosoftGraphDeviceInstallState**](microsoft.graph.deviceInstallState.md) |  | [optional] 

## Methods

### GetUserName

`func (o *UserInstallStateSummary) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserInstallStateSummary) GetUserNameOk() (string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserName

`func (o *UserInstallStateSummary) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserName

`func (o *UserInstallStateSummary) SetUserName(v string)`

SetUserName gets a reference to the given string and assigns it to the UserName field.

### SetUserNameExplicitNull

`func (o *UserInstallStateSummary) SetUserNameExplicitNull(b bool)`

SetUserNameExplicitNull (un)sets UserName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserName value is set to nil even if false is passed
### GetInstalledDeviceCount

`func (o *UserInstallStateSummary) GetInstalledDeviceCount() int32`

GetInstalledDeviceCount returns the InstalledDeviceCount field if non-nil, zero value otherwise.

### GetInstalledDeviceCountOk

`func (o *UserInstallStateSummary) GetInstalledDeviceCountOk() (int32, bool)`

GetInstalledDeviceCountOk returns a tuple with the InstalledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstalledDeviceCount

`func (o *UserInstallStateSummary) HasInstalledDeviceCount() bool`

HasInstalledDeviceCount returns a boolean if a field has been set.

### SetInstalledDeviceCount

`func (o *UserInstallStateSummary) SetInstalledDeviceCount(v int32)`

SetInstalledDeviceCount gets a reference to the given int32 and assigns it to the InstalledDeviceCount field.

### GetFailedDeviceCount

`func (o *UserInstallStateSummary) GetFailedDeviceCount() int32`

GetFailedDeviceCount returns the FailedDeviceCount field if non-nil, zero value otherwise.

### GetFailedDeviceCountOk

`func (o *UserInstallStateSummary) GetFailedDeviceCountOk() (int32, bool)`

GetFailedDeviceCountOk returns a tuple with the FailedDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailedDeviceCount

`func (o *UserInstallStateSummary) HasFailedDeviceCount() bool`

HasFailedDeviceCount returns a boolean if a field has been set.

### SetFailedDeviceCount

`func (o *UserInstallStateSummary) SetFailedDeviceCount(v int32)`

SetFailedDeviceCount gets a reference to the given int32 and assigns it to the FailedDeviceCount field.

### GetNotInstalledDeviceCount

`func (o *UserInstallStateSummary) GetNotInstalledDeviceCount() int32`

GetNotInstalledDeviceCount returns the NotInstalledDeviceCount field if non-nil, zero value otherwise.

### GetNotInstalledDeviceCountOk

`func (o *UserInstallStateSummary) GetNotInstalledDeviceCountOk() (int32, bool)`

GetNotInstalledDeviceCountOk returns a tuple with the NotInstalledDeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotInstalledDeviceCount

`func (o *UserInstallStateSummary) HasNotInstalledDeviceCount() bool`

HasNotInstalledDeviceCount returns a boolean if a field has been set.

### SetNotInstalledDeviceCount

`func (o *UserInstallStateSummary) SetNotInstalledDeviceCount(v int32)`

SetNotInstalledDeviceCount gets a reference to the given int32 and assigns it to the NotInstalledDeviceCount field.

### GetDeviceStates

`func (o *UserInstallStateSummary) GetDeviceStates() []MicrosoftGraphDeviceInstallState`

GetDeviceStates returns the DeviceStates field if non-nil, zero value otherwise.

### GetDeviceStatesOk

`func (o *UserInstallStateSummary) GetDeviceStatesOk() ([]MicrosoftGraphDeviceInstallState, bool)`

GetDeviceStatesOk returns a tuple with the DeviceStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStates

`func (o *UserInstallStateSummary) HasDeviceStates() bool`

HasDeviceStates returns a boolean if a field has been set.

### SetDeviceStates

`func (o *UserInstallStateSummary) SetDeviceStates(v []MicrosoftGraphDeviceInstallState)`

SetDeviceStates gets a reference to the given []MicrosoftGraphDeviceInstallState and assigns it to the DeviceStates field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



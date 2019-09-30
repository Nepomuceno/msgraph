# DeviceInstallState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** | Device name. | [optional] 
**DeviceId** | Pointer to **string** | Device Id. | [optional] 
**LastSyncDateTime** | Pointer to [**time.Time**](time.Time.md) | Last sync date and time. | [optional] 
**InstallState** | Pointer to [**AnyOfmicrosoftGraphInstallState**](anyOf&lt;microsoft.graph.installState&gt;.md) | The install state of the eBook. | [optional] 
**ErrorCode** | Pointer to **string** | The error code for install failures. | [optional] 
**OsVersion** | Pointer to **string** | OS Version. | [optional] 
**OsDescription** | Pointer to **string** | OS Description. | [optional] 
**UserName** | Pointer to **string** | Device User Name. | [optional] 

## Methods

### GetDeviceName

`func (o *DeviceInstallState) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceInstallState) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *DeviceInstallState) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *DeviceInstallState) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### SetDeviceNameExplicitNull

`func (o *DeviceInstallState) SetDeviceNameExplicitNull(b bool)`

SetDeviceNameExplicitNull (un)sets DeviceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceName value is set to nil even if false is passed
### GetDeviceId

`func (o *DeviceInstallState) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceInstallState) GetDeviceIdOk() (string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceId

`func (o *DeviceInstallState) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceId

`func (o *DeviceInstallState) SetDeviceId(v string)`

SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.

### SetDeviceIdExplicitNull

`func (o *DeviceInstallState) SetDeviceIdExplicitNull(b bool)`

SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceId value is set to nil even if false is passed
### GetLastSyncDateTime

`func (o *DeviceInstallState) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *DeviceInstallState) GetLastSyncDateTimeOk() (time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastSyncDateTime

`func (o *DeviceInstallState) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### SetLastSyncDateTime

`func (o *DeviceInstallState) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.

### GetInstallState

`func (o *DeviceInstallState) GetInstallState() AnyOfmicrosoftGraphInstallState`

GetInstallState returns the InstallState field if non-nil, zero value otherwise.

### GetInstallStateOk

`func (o *DeviceInstallState) GetInstallStateOk() (AnyOfmicrosoftGraphInstallState, bool)`

GetInstallStateOk returns a tuple with the InstallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstallState

`func (o *DeviceInstallState) HasInstallState() bool`

HasInstallState returns a boolean if a field has been set.

### SetInstallState

`func (o *DeviceInstallState) SetInstallState(v AnyOfmicrosoftGraphInstallState)`

SetInstallState gets a reference to the given AnyOfmicrosoftGraphInstallState and assigns it to the InstallState field.

### GetErrorCode

`func (o *DeviceInstallState) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *DeviceInstallState) GetErrorCodeOk() (string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorCode

`func (o *DeviceInstallState) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCode

`func (o *DeviceInstallState) SetErrorCode(v string)`

SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.

### SetErrorCodeExplicitNull

`func (o *DeviceInstallState) SetErrorCodeExplicitNull(b bool)`

SetErrorCodeExplicitNull (un)sets ErrorCode to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ErrorCode value is set to nil even if false is passed
### GetOsVersion

`func (o *DeviceInstallState) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceInstallState) GetOsVersionOk() (string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsVersion

`func (o *DeviceInstallState) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersion

`func (o *DeviceInstallState) SetOsVersion(v string)`

SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.

### SetOsVersionExplicitNull

`func (o *DeviceInstallState) SetOsVersionExplicitNull(b bool)`

SetOsVersionExplicitNull (un)sets OsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsVersion value is set to nil even if false is passed
### GetOsDescription

`func (o *DeviceInstallState) GetOsDescription() string`

GetOsDescription returns the OsDescription field if non-nil, zero value otherwise.

### GetOsDescriptionOk

`func (o *DeviceInstallState) GetOsDescriptionOk() (string, bool)`

GetOsDescriptionOk returns a tuple with the OsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsDescription

`func (o *DeviceInstallState) HasOsDescription() bool`

HasOsDescription returns a boolean if a field has been set.

### SetOsDescription

`func (o *DeviceInstallState) SetOsDescription(v string)`

SetOsDescription gets a reference to the given string and assigns it to the OsDescription field.

### SetOsDescriptionExplicitNull

`func (o *DeviceInstallState) SetOsDescriptionExplicitNull(b bool)`

SetOsDescriptionExplicitNull (un)sets OsDescription to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OsDescription value is set to nil even if false is passed
### GetUserName

`func (o *DeviceInstallState) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DeviceInstallState) GetUserNameOk() (string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserName

`func (o *DeviceInstallState) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserName

`func (o *DeviceInstallState) SetUserName(v string)`

SetUserName gets a reference to the given string and assigns it to the UserName field.

### SetUserNameExplicitNull

`func (o *DeviceInstallState) SetUserNameExplicitNull(b bool)`

SetUserNameExplicitNull (un)sets UserName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserName value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



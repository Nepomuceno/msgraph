# DeviceComplianceSettingState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Setting** | Pointer to **string** | The setting class name and property name. | [optional] 
**SettingName** | Pointer to **string** | The Setting Name that is being reported | [optional] 
**DeviceId** | Pointer to **string** | The Device Id that is being reported | [optional] 
**DeviceName** | Pointer to **string** | The Device Name that is being reported | [optional] 
**UserId** | Pointer to **string** | The user Id that is being reported | [optional] 
**UserEmail** | Pointer to **string** | The User email address that is being reported | [optional] 
**UserName** | Pointer to **string** | The User Name that is being reported | [optional] 
**UserPrincipalName** | Pointer to **string** | The User PrincipalName that is being reported | [optional] 
**DeviceModel** | Pointer to **string** | The device model that is being reported | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the setting | [optional] 
**ComplianceGracePeriodExpirationDateTime** | Pointer to [**time.Time**](time.Time.md) | The DateTime when device compliance grace period expires | [optional] 

## Methods

### GetSetting

`func (o *DeviceComplianceSettingState) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *DeviceComplianceSettingState) GetSettingOk() (string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSetting

`func (o *DeviceComplianceSettingState) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSetting

`func (o *DeviceComplianceSettingState) SetSetting(v string)`

SetSetting gets a reference to the given string and assigns it to the Setting field.

### SetSettingExplicitNull

`func (o *DeviceComplianceSettingState) SetSettingExplicitNull(b bool)`

SetSettingExplicitNull (un)sets Setting to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Setting value is set to nil even if false is passed
### GetSettingName

`func (o *DeviceComplianceSettingState) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *DeviceComplianceSettingState) GetSettingNameOk() (string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingName

`func (o *DeviceComplianceSettingState) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingName

`func (o *DeviceComplianceSettingState) SetSettingName(v string)`

SetSettingName gets a reference to the given string and assigns it to the SettingName field.

### SetSettingNameExplicitNull

`func (o *DeviceComplianceSettingState) SetSettingNameExplicitNull(b bool)`

SetSettingNameExplicitNull (un)sets SettingName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingName value is set to nil even if false is passed
### GetDeviceId

`func (o *DeviceComplianceSettingState) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceComplianceSettingState) GetDeviceIdOk() (string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceId

`func (o *DeviceComplianceSettingState) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceId

`func (o *DeviceComplianceSettingState) SetDeviceId(v string)`

SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.

### SetDeviceIdExplicitNull

`func (o *DeviceComplianceSettingState) SetDeviceIdExplicitNull(b bool)`

SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceId value is set to nil even if false is passed
### GetDeviceName

`func (o *DeviceComplianceSettingState) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceComplianceSettingState) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *DeviceComplianceSettingState) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *DeviceComplianceSettingState) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### SetDeviceNameExplicitNull

`func (o *DeviceComplianceSettingState) SetDeviceNameExplicitNull(b bool)`

SetDeviceNameExplicitNull (un)sets DeviceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceName value is set to nil even if false is passed
### GetUserId

`func (o *DeviceComplianceSettingState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DeviceComplianceSettingState) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *DeviceComplianceSettingState) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *DeviceComplianceSettingState) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *DeviceComplianceSettingState) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetUserEmail

`func (o *DeviceComplianceSettingState) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *DeviceComplianceSettingState) GetUserEmailOk() (string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserEmail

`func (o *DeviceComplianceSettingState) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmail

`func (o *DeviceComplianceSettingState) SetUserEmail(v string)`

SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.

### SetUserEmailExplicitNull

`func (o *DeviceComplianceSettingState) SetUserEmailExplicitNull(b bool)`

SetUserEmailExplicitNull (un)sets UserEmail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserEmail value is set to nil even if false is passed
### GetUserName

`func (o *DeviceComplianceSettingState) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DeviceComplianceSettingState) GetUserNameOk() (string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserName

`func (o *DeviceComplianceSettingState) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserName

`func (o *DeviceComplianceSettingState) SetUserName(v string)`

SetUserName gets a reference to the given string and assigns it to the UserName field.

### SetUserNameExplicitNull

`func (o *DeviceComplianceSettingState) SetUserNameExplicitNull(b bool)`

SetUserNameExplicitNull (un)sets UserName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserName value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *DeviceComplianceSettingState) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *DeviceComplianceSettingState) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *DeviceComplianceSettingState) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *DeviceComplianceSettingState) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *DeviceComplianceSettingState) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetDeviceModel

`func (o *DeviceComplianceSettingState) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *DeviceComplianceSettingState) GetDeviceModelOk() (string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceModel

`func (o *DeviceComplianceSettingState) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModel

`func (o *DeviceComplianceSettingState) SetDeviceModel(v string)`

SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.

### SetDeviceModelExplicitNull

`func (o *DeviceComplianceSettingState) SetDeviceModelExplicitNull(b bool)`

SetDeviceModelExplicitNull (un)sets DeviceModel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceModel value is set to nil even if false is passed
### GetState

`func (o *DeviceComplianceSettingState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeviceComplianceSettingState) GetStateOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *DeviceComplianceSettingState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *DeviceComplianceSettingState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.

### GetComplianceGracePeriodExpirationDateTime

`func (o *DeviceComplianceSettingState) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *DeviceComplianceSettingState) GetComplianceGracePeriodExpirationDateTimeOk() (time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceGracePeriodExpirationDateTime

`func (o *DeviceComplianceSettingState) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *DeviceComplianceSettingState) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceGracePeriodExpirationDateTime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



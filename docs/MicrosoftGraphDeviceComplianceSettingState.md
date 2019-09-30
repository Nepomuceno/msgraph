# MicrosoftGraphDeviceComplianceSettingState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSetting

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetSettingOk() (string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSetting

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSetting

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetSetting(v string)`

SetSetting gets a reference to the given string and assigns it to the Setting field.

### SetSettingExplicitNull

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetSettingExplicitNull(b bool)`

SetSettingExplicitNull (un)sets Setting to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Setting value is set to nil even if false is passed
### GetSettingName

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetSettingNameOk() (string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingName

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingName

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetSettingName(v string)`

SetSettingName gets a reference to the given string and assigns it to the SettingName field.

### SetSettingNameExplicitNull

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetSettingNameExplicitNull(b bool)`

SetSettingNameExplicitNull (un)sets SettingName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingName value is set to nil even if false is passed
### GetDeviceId

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceIdOk() (string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceId

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceId

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceId(v string)`

SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.

### SetDeviceIdExplicitNull

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceIdExplicitNull(b bool)`

SetDeviceIdExplicitNull (un)sets DeviceId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceId value is set to nil even if false is passed
### GetDeviceName

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### SetDeviceNameExplicitNull

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceNameExplicitNull(b bool)`

SetDeviceNameExplicitNull (un)sets DeviceName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceName value is set to nil even if false is passed
### GetUserId

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetUserEmail

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserEmailOk() (string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserEmail

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmail

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserEmail(v string)`

SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.

### SetUserEmailExplicitNull

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserEmailExplicitNull(b bool)`

SetUserEmailExplicitNull (un)sets UserEmail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserEmail value is set to nil even if false is passed
### GetUserName

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserNameOk() (string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserName

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserName

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserName(v string)`

SetUserName gets a reference to the given string and assigns it to the UserName field.

### SetUserNameExplicitNull

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserNameExplicitNull(b bool)`

SetUserNameExplicitNull (un)sets UserName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserName value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetDeviceModel

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetDeviceModelOk() (string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceModel

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModel

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceModel(v string)`

SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.

### SetDeviceModelExplicitNull

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetDeviceModelExplicitNull(b bool)`

SetDeviceModelExplicitNull (un)sets DeviceModel to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceModel value is set to nil even if false is passed
### GetState

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetStateOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.

### GetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetComplianceGracePeriodExpirationDateTime() time.Time`

GetComplianceGracePeriodExpirationDateTime returns the ComplianceGracePeriodExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceGracePeriodExpirationDateTimeOk

`func (o *MicrosoftGraphDeviceComplianceSettingState) GetComplianceGracePeriodExpirationDateTimeOk() (time.Time, bool)`

GetComplianceGracePeriodExpirationDateTimeOk returns a tuple with the ComplianceGracePeriodExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphDeviceComplianceSettingState) HasComplianceGracePeriodExpirationDateTime() bool`

HasComplianceGracePeriodExpirationDateTime returns a boolean if a field has been set.

### SetComplianceGracePeriodExpirationDateTime

`func (o *MicrosoftGraphDeviceComplianceSettingState) SetComplianceGracePeriodExpirationDateTime(v time.Time)`

SetComplianceGracePeriodExpirationDateTime gets a reference to the given time.Time and assigns it to the ComplianceGracePeriodExpirationDateTime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



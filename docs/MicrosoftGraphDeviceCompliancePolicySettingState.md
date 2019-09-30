# MicrosoftGraphDeviceCompliancePolicySettingState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Setting** | Pointer to **string** | The setting that is being reported | [optional] 
**SettingName** | Pointer to **string** | Localized/user friendly setting name that is being reported | [optional] 
**InstanceDisplayName** | Pointer to **string** | Name of setting instance that is being reported. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphComplianceStatus**](anyOf&lt;microsoft.graph.complianceStatus&gt;.md) | The compliance state of the setting | [optional] 
**ErrorCode** | Pointer to **int64** | Error code for the setting | [optional] 
**ErrorDescription** | Pointer to **string** | Error description | [optional] 
**UserId** | Pointer to **string** | UserId | [optional] 
**UserName** | Pointer to **string** | UserName | [optional] 
**UserEmail** | Pointer to **string** | UserEmail | [optional] 
**UserPrincipalName** | Pointer to **string** | UserPrincipalName. | [optional] 
**Sources** | Pointer to [**[]AnyOfmicrosoftGraphSettingSource**](anyOf&lt;microsoft.graph.settingSource&gt;.md) | Contributing policies | [optional] 
**CurrentValue** | Pointer to **string** | Current value of setting on device | [optional] 

## Methods

### GetSetting

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSettingOk() (string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSetting

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### SetSetting

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSetting(v string)`

SetSetting gets a reference to the given string and assigns it to the Setting field.

### SetSettingExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSettingExplicitNull(b bool)`

SetSettingExplicitNull (un)sets Setting to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Setting value is set to nil even if false is passed
### GetSettingName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSettingNameOk() (string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettingName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSettingName(v string)`

SetSettingName gets a reference to the given string and assigns it to the SettingName field.

### SetSettingNameExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSettingNameExplicitNull(b bool)`

SetSettingNameExplicitNull (un)sets SettingName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SettingName value is set to nil even if false is passed
### GetInstanceDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetInstanceDisplayName() string`

GetInstanceDisplayName returns the InstanceDisplayName field if non-nil, zero value otherwise.

### GetInstanceDisplayNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetInstanceDisplayNameOk() (string, bool)`

GetInstanceDisplayNameOk returns a tuple with the InstanceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstanceDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasInstanceDisplayName() bool`

HasInstanceDisplayName returns a boolean if a field has been set.

### SetInstanceDisplayName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetInstanceDisplayName(v string)`

SetInstanceDisplayName gets a reference to the given string and assigns it to the InstanceDisplayName field.

### SetInstanceDisplayNameExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetInstanceDisplayNameExplicitNull(b bool)`

SetInstanceDisplayNameExplicitNull (un)sets InstanceDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InstanceDisplayName value is set to nil even if false is passed
### GetState

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetState() AnyOfmicrosoftGraphComplianceStatus`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetStateOk() (AnyOfmicrosoftGraphComplianceStatus, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetState(v AnyOfmicrosoftGraphComplianceStatus)`

SetState gets a reference to the given AnyOfmicrosoftGraphComplianceStatus and assigns it to the State field.

### GetErrorCode

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetErrorCode() int64`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetErrorCodeOk() (int64, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorCode

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCode

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetErrorCode(v int64)`

SetErrorCode gets a reference to the given int64 and assigns it to the ErrorCode field.

### GetErrorDescription

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetErrorDescriptionOk() (string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrorDescription

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescription

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetErrorDescription(v string)`

SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.

### SetErrorDescriptionExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetErrorDescriptionExplicitNull(b bool)`

SetErrorDescriptionExplicitNull (un)sets ErrorDescription to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ErrorDescription value is set to nil even if false is passed
### GetUserId

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserIdOk() (string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserId

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserId

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserId(v string)`

SetUserId gets a reference to the given string and assigns it to the UserId field.

### SetUserIdExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserIdExplicitNull(b bool)`

SetUserIdExplicitNull (un)sets UserId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserId value is set to nil even if false is passed
### GetUserName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserNameOk() (string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserName(v string)`

SetUserName gets a reference to the given string and assigns it to the UserName field.

### SetUserNameExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserNameExplicitNull(b bool)`

SetUserNameExplicitNull (un)sets UserName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserName value is set to nil even if false is passed
### GetUserEmail

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserEmailOk() (string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserEmail

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmail

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserEmail(v string)`

SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.

### SetUserEmailExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserEmailExplicitNull(b bool)`

SetUserEmailExplicitNull (un)sets UserEmail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserEmail value is set to nil even if false is passed
### GetUserPrincipalName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetUserPrincipalNameOk() (string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserPrincipalName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserPrincipalName(v string)`

SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.

### SetUserPrincipalNameExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetUserPrincipalNameExplicitNull(b bool)`

SetUserPrincipalNameExplicitNull (un)sets UserPrincipalName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserPrincipalName value is set to nil even if false is passed
### GetSources

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSources() []AnyOfmicrosoftGraphSettingSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetSourcesOk() ([]AnyOfmicrosoftGraphSettingSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSources

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSources

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetSources(v []AnyOfmicrosoftGraphSettingSource)`

SetSources gets a reference to the given []AnyOfmicrosoftGraphSettingSource and assigns it to the Sources field.

### GetCurrentValue

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) GetCurrentValueOk() (string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentValue

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### SetCurrentValue

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetCurrentValue(v string)`

SetCurrentValue gets a reference to the given string and assigns it to the CurrentValue field.

### SetCurrentValueExplicitNull

`func (o *MicrosoftGraphDeviceCompliancePolicySettingState) SetCurrentValueExplicitNull(b bool)`

SetCurrentValueExplicitNull (un)sets CurrentValue to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CurrentValue value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



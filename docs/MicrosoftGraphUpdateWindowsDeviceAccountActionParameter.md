# MicrosoftGraphUpdateWindowsDeviceAccountActionParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceAccount** | Pointer to [**AnyOfmicrosoftGraphWindowsDeviceAccount**](anyOf&lt;microsoft.graph.windowsDeviceAccount&gt;.md) |  | [optional] 
**PasswordRotationEnabled** | Pointer to **bool** |  | [optional] 
**CalendarSyncEnabled** | Pointer to **bool** |  | [optional] 
**DeviceAccountEmail** | Pointer to **string** |  | [optional] 
**ExchangeServer** | Pointer to **string** |  | [optional] 
**SessionInitiationProtocalAddress** | Pointer to **string** |  | [optional] 

## Methods

### GetDeviceAccount

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccount() AnyOfmicrosoftGraphWindowsDeviceAccount`

GetDeviceAccount returns the DeviceAccount field if non-nil, zero value otherwise.

### GetDeviceAccountOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccountOk() (AnyOfmicrosoftGraphWindowsDeviceAccount, bool)`

GetDeviceAccountOk returns a tuple with the DeviceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceAccount

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasDeviceAccount() bool`

HasDeviceAccount returns a boolean if a field has been set.

### SetDeviceAccount

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccount(v AnyOfmicrosoftGraphWindowsDeviceAccount)`

SetDeviceAccount gets a reference to the given AnyOfmicrosoftGraphWindowsDeviceAccount and assigns it to the DeviceAccount field.

### SetDeviceAccountExplicitNull

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccountExplicitNull(b bool)`

SetDeviceAccountExplicitNull (un)sets DeviceAccount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceAccount value is set to nil even if false is passed
### GetPasswordRotationEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabled() bool`

GetPasswordRotationEnabled returns the PasswordRotationEnabled field if non-nil, zero value otherwise.

### GetPasswordRotationEnabledOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabledOk() (bool, bool)`

GetPasswordRotationEnabledOk returns a tuple with the PasswordRotationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRotationEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasPasswordRotationEnabled() bool`

HasPasswordRotationEnabled returns a boolean if a field has been set.

### SetPasswordRotationEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabled(v bool)`

SetPasswordRotationEnabled gets a reference to the given bool and assigns it to the PasswordRotationEnabled field.

### SetPasswordRotationEnabledExplicitNull

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabledExplicitNull(b bool)`

SetPasswordRotationEnabledExplicitNull (un)sets PasswordRotationEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordRotationEnabled value is set to nil even if false is passed
### GetCalendarSyncEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabled() bool`

GetCalendarSyncEnabled returns the CalendarSyncEnabled field if non-nil, zero value otherwise.

### GetCalendarSyncEnabledOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabledOk() (bool, bool)`

GetCalendarSyncEnabledOk returns a tuple with the CalendarSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendarSyncEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasCalendarSyncEnabled() bool`

HasCalendarSyncEnabled returns a boolean if a field has been set.

### SetCalendarSyncEnabled

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabled(v bool)`

SetCalendarSyncEnabled gets a reference to the given bool and assigns it to the CalendarSyncEnabled field.

### SetCalendarSyncEnabledExplicitNull

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabledExplicitNull(b bool)`

SetCalendarSyncEnabledExplicitNull (un)sets CalendarSyncEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CalendarSyncEnabled value is set to nil even if false is passed
### GetDeviceAccountEmail

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmail() string`

GetDeviceAccountEmail returns the DeviceAccountEmail field if non-nil, zero value otherwise.

### GetDeviceAccountEmailOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmailOk() (string, bool)`

GetDeviceAccountEmailOk returns a tuple with the DeviceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceAccountEmail

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasDeviceAccountEmail() bool`

HasDeviceAccountEmail returns a boolean if a field has been set.

### SetDeviceAccountEmail

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmail(v string)`

SetDeviceAccountEmail gets a reference to the given string and assigns it to the DeviceAccountEmail field.

### SetDeviceAccountEmailExplicitNull

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmailExplicitNull(b bool)`

SetDeviceAccountEmailExplicitNull (un)sets DeviceAccountEmail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceAccountEmail value is set to nil even if false is passed
### GetExchangeServer

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetExchangeServer() string`

GetExchangeServer returns the ExchangeServer field if non-nil, zero value otherwise.

### GetExchangeServerOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetExchangeServerOk() (string, bool)`

GetExchangeServerOk returns a tuple with the ExchangeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeServer

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasExchangeServer() bool`

HasExchangeServer returns a boolean if a field has been set.

### SetExchangeServer

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetExchangeServer(v string)`

SetExchangeServer gets a reference to the given string and assigns it to the ExchangeServer field.

### SetExchangeServerExplicitNull

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetExchangeServerExplicitNull(b bool)`

SetExchangeServerExplicitNull (un)sets ExchangeServer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExchangeServer value is set to nil even if false is passed
### GetSessionInitiationProtocalAddress

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddress() string`

GetSessionInitiationProtocalAddress returns the SessionInitiationProtocalAddress field if non-nil, zero value otherwise.

### GetSessionInitiationProtocalAddressOk

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddressOk() (string, bool)`

GetSessionInitiationProtocalAddressOk returns a tuple with the SessionInitiationProtocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionInitiationProtocalAddress

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) HasSessionInitiationProtocalAddress() bool`

HasSessionInitiationProtocalAddress returns a boolean if a field has been set.

### SetSessionInitiationProtocalAddress

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddress(v string)`

SetSessionInitiationProtocalAddress gets a reference to the given string and assigns it to the SessionInitiationProtocalAddress field.

### SetSessionInitiationProtocalAddressExplicitNull

`func (o *MicrosoftGraphUpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddressExplicitNull(b bool)`

SetSessionInitiationProtocalAddressExplicitNull (un)sets SessionInitiationProtocalAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SessionInitiationProtocalAddress value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



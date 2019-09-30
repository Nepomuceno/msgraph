# TelecomExpenseManagementPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the TEM partner. | [optional] 
**Url** | Pointer to **string** | URL of the TEM partner&#39;s administrative control panel, where an administrator can configure their TEM service. | [optional] 
**AppAuthorized** | Pointer to **bool** | Whether the partner&#39;s AAD app has been authorized to access Intune. | [optional] 
**Enabled** | Pointer to **bool** | Whether Intune&#39;s connection to the TEM service is currently enabled or disabled. | [optional] 
**LastConnectionDateTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the last request sent to Intune by the TEM partner. | [optional] 

## Methods

### GetDisplayName

`func (o *TelecomExpenseManagementPartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TelecomExpenseManagementPartner) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *TelecomExpenseManagementPartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *TelecomExpenseManagementPartner) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *TelecomExpenseManagementPartner) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetUrl

`func (o *TelecomExpenseManagementPartner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TelecomExpenseManagementPartner) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *TelecomExpenseManagementPartner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *TelecomExpenseManagementPartner) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.

### SetUrlExplicitNull

`func (o *TelecomExpenseManagementPartner) SetUrlExplicitNull(b bool)`

SetUrlExplicitNull (un)sets Url to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Url value is set to nil even if false is passed
### GetAppAuthorized

`func (o *TelecomExpenseManagementPartner) GetAppAuthorized() bool`

GetAppAuthorized returns the AppAuthorized field if non-nil, zero value otherwise.

### GetAppAuthorizedOk

`func (o *TelecomExpenseManagementPartner) GetAppAuthorizedOk() (bool, bool)`

GetAppAuthorizedOk returns a tuple with the AppAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppAuthorized

`func (o *TelecomExpenseManagementPartner) HasAppAuthorized() bool`

HasAppAuthorized returns a boolean if a field has been set.

### SetAppAuthorized

`func (o *TelecomExpenseManagementPartner) SetAppAuthorized(v bool)`

SetAppAuthorized gets a reference to the given bool and assigns it to the AppAuthorized field.

### GetEnabled

`func (o *TelecomExpenseManagementPartner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TelecomExpenseManagementPartner) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *TelecomExpenseManagementPartner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *TelecomExpenseManagementPartner) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetLastConnectionDateTime

`func (o *TelecomExpenseManagementPartner) GetLastConnectionDateTime() time.Time`

GetLastConnectionDateTime returns the LastConnectionDateTime field if non-nil, zero value otherwise.

### GetLastConnectionDateTimeOk

`func (o *TelecomExpenseManagementPartner) GetLastConnectionDateTimeOk() (time.Time, bool)`

GetLastConnectionDateTimeOk returns a tuple with the LastConnectionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastConnectionDateTime

`func (o *TelecomExpenseManagementPartner) HasLastConnectionDateTime() bool`

HasLastConnectionDateTime returns a boolean if a field has been set.

### SetLastConnectionDateTime

`func (o *TelecomExpenseManagementPartner) SetLastConnectionDateTime(v time.Time)`

SetLastConnectionDateTime gets a reference to the given time.Time and assigns it to the LastConnectionDateTime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



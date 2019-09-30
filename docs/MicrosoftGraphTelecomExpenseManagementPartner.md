# MicrosoftGraphTelecomExpenseManagementPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Display name of the TEM partner. | [optional] 
**Url** | Pointer to **string** | URL of the TEM partner&#39;s administrative control panel, where an administrator can configure their TEM service. | [optional] 
**AppAuthorized** | Pointer to **bool** | Whether the partner&#39;s AAD app has been authorized to access Intune. | [optional] 
**Enabled** | Pointer to **bool** | Whether Intune&#39;s connection to the TEM service is currently enabled or disabled. | [optional] 
**LastConnectionDateTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the last request sent to Intune by the TEM partner. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetUrl

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.

### SetUrlExplicitNull

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetUrlExplicitNull(b bool)`

SetUrlExplicitNull (un)sets Url to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Url value is set to nil even if false is passed
### GetAppAuthorized

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetAppAuthorized() bool`

GetAppAuthorized returns the AppAuthorized field if non-nil, zero value otherwise.

### GetAppAuthorizedOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetAppAuthorizedOk() (bool, bool)`

GetAppAuthorizedOk returns a tuple with the AppAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppAuthorized

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasAppAuthorized() bool`

HasAppAuthorized returns a boolean if a field has been set.

### SetAppAuthorized

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetAppAuthorized(v bool)`

SetAppAuthorized gets a reference to the given bool and assigns it to the AppAuthorized field.

### GetEnabled

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetLastConnectionDateTime

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetLastConnectionDateTime() time.Time`

GetLastConnectionDateTime returns the LastConnectionDateTime field if non-nil, zero value otherwise.

### GetLastConnectionDateTimeOk

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) GetLastConnectionDateTimeOk() (time.Time, bool)`

GetLastConnectionDateTimeOk returns a tuple with the LastConnectionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastConnectionDateTime

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) HasLastConnectionDateTime() bool`

HasLastConnectionDateTime returns a boolean if a field has been set.

### SetLastConnectionDateTime

`func (o *MicrosoftGraphTelecomExpenseManagementPartner) SetLastConnectionDateTime(v time.Time)`

SetLastConnectionDateTime gets a reference to the given time.Time and assigns it to the LastConnectionDateTime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



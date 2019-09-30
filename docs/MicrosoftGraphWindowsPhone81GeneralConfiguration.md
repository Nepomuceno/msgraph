# MicrosoftGraphWindowsPhone81GeneralConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was created. | [optional] 
**Description** | Pointer to **string** | Admin provided description of the Device Configuration. | [optional] 
**DisplayName** | Pointer to **string** | Admin provided name of the device configuration. | [optional] 
**Version** | Pointer to **int32** | Version of the device configuration. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphDeviceConfigurationAssignment**](microsoft.graph.deviceConfigurationAssignment.md) |  | [optional] 
**DeviceStatuses** | Pointer to [**[]MicrosoftGraphDeviceConfigurationDeviceStatus**](microsoft.graph.deviceConfigurationDeviceStatus.md) |  | [optional] 
**UserStatuses** | Pointer to [**[]MicrosoftGraphDeviceConfigurationUserStatus**](microsoft.graph.deviceConfigurationUserStatus.md) |  | [optional] 
**DeviceStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview**](anyOf&lt;microsoft.graph.deviceConfigurationDeviceOverview&gt;.md) |  | [optional] 
**UserStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationUserOverview**](anyOf&lt;microsoft.graph.deviceConfigurationUserOverview&gt;.md) |  | [optional] 
**DeviceSettingStateSummaries** | Pointer to [**[]MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md) |  | [optional] 
**ApplyOnlyToWindowsPhone81** | Pointer to **bool** | Value indicating whether this policy only applies to Windows Phone 8.1. This property is read-only. | [optional] 
**AppsBlockCopyPaste** | Pointer to **bool** | Indicates whether or not to block copy paste. | [optional] 
**BluetoothBlocked** | Pointer to **bool** | Indicates whether or not to block bluetooth. | [optional] 
**CameraBlocked** | Pointer to **bool** | Indicates whether or not to block camera. | [optional] 
**CellularBlockWifiTethering** | Pointer to **bool** | Indicates whether or not to block Wi-Fi tethering. Has no impact if Wi-Fi is blocked. | [optional] 
**CompliantAppsList** | Pointer to [**[]AnyOfmicrosoftGraphAppListItem**](anyOf&lt;microsoft.graph.appListItem&gt;.md) | List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements. | [optional] 
**CompliantAppListType** | Pointer to [**AnyOfmicrosoftGraphAppListType**](anyOf&lt;microsoft.graph.appListType&gt;.md) | List that is in the AppComplianceList. | [optional] 
**DiagnosticDataBlockSubmission** | Pointer to **bool** | Indicates whether or not to block diagnostic data submission. | [optional] 
**EmailBlockAddingAccounts** | Pointer to **bool** | Indicates whether or not to block custom email accounts. | [optional] 
**LocationServicesBlocked** | Pointer to **bool** | Indicates whether or not to block location services. | [optional] 
**MicrosoftAccountBlocked** | Pointer to **bool** | Indicates whether or not to block using a Microsoft Account. | [optional] 
**NfcBlocked** | Pointer to **bool** | Indicates whether or not to block Near-Field Communication. | [optional] 
**PasswordBlockSimple** | Pointer to **bool** | Indicates whether or not to block syncing the calendar. | [optional] 
**PasswordExpirationDays** | Pointer to **int32** | Number of days before the password expires. | [optional] 
**PasswordMinimumLength** | Pointer to **int32** | Minimum length of passwords. | [optional] 
**PasswordMinutesOfInactivityBeforeScreenTimeout** | Pointer to **int32** | Minutes of inactivity before screen timeout. | [optional] 
**PasswordMinimumCharacterSetCount** | Pointer to **int32** | Number of character sets a password must contain. | [optional] 
**PasswordPreviousPasswordBlockCount** | Pointer to **int32** | Number of previous passwords to block. Valid values 0 to 24 | [optional] 
**PasswordSignInFailureCountBeforeFactoryReset** | Pointer to **int32** | Number of sign in failures allowed before factory reset. | [optional] 
**PasswordRequiredType** | Pointer to [**AnyOfmicrosoftGraphRequiredPasswordType**](anyOf&lt;microsoft.graph.requiredPasswordType&gt;.md) | Password type that is required. | [optional] 
**PasswordRequired** | Pointer to **bool** | Indicates whether or not to require a password. | [optional] 
**ScreenCaptureBlocked** | Pointer to **bool** | Indicates whether or not to block screenshots. | [optional] 
**StorageBlockRemovableStorage** | Pointer to **bool** | Indicates whether or not to block removable storage. | [optional] 
**StorageRequireEncryption** | Pointer to **bool** | Indicates whether or not to require encryption. | [optional] 
**WebBrowserBlocked** | Pointer to **bool** | Indicates whether or not to block the web browser. | [optional] 
**WifiBlocked** | Pointer to **bool** | Indicates whether or not to block Wi-Fi. | [optional] 
**WifiBlockAutomaticConnectHotspots** | Pointer to **bool** | Indicates whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked. | [optional] 
**WifiBlockHotspotReporting** | Pointer to **bool** | Indicates whether or not to block Wi-Fi hotspot reporting. Has no impact if Wi-Fi is blocked. | [optional] 
**WindowsStoreBlocked** | Pointer to **bool** | Indicates whether or not to block the Windows Store. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetApplyOnlyToWindowsPhone81

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetApplyOnlyToWindowsPhone81() bool`

GetApplyOnlyToWindowsPhone81 returns the ApplyOnlyToWindowsPhone81 field if non-nil, zero value otherwise.

### GetApplyOnlyToWindowsPhone81Ok

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetApplyOnlyToWindowsPhone81Ok() (bool, bool)`

GetApplyOnlyToWindowsPhone81Ok returns a tuple with the ApplyOnlyToWindowsPhone81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplyOnlyToWindowsPhone81

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasApplyOnlyToWindowsPhone81() bool`

HasApplyOnlyToWindowsPhone81 returns a boolean if a field has been set.

### SetApplyOnlyToWindowsPhone81

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetApplyOnlyToWindowsPhone81(v bool)`

SetApplyOnlyToWindowsPhone81 gets a reference to the given bool and assigns it to the ApplyOnlyToWindowsPhone81 field.

### GetAppsBlockCopyPaste

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetAppsBlockCopyPaste() bool`

GetAppsBlockCopyPaste returns the AppsBlockCopyPaste field if non-nil, zero value otherwise.

### GetAppsBlockCopyPasteOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetAppsBlockCopyPasteOk() (bool, bool)`

GetAppsBlockCopyPasteOk returns a tuple with the AppsBlockCopyPaste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockCopyPaste

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasAppsBlockCopyPaste() bool`

HasAppsBlockCopyPaste returns a boolean if a field has been set.

### SetAppsBlockCopyPaste

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetAppsBlockCopyPaste(v bool)`

SetAppsBlockCopyPaste gets a reference to the given bool and assigns it to the AppsBlockCopyPaste field.

### GetBluetoothBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetBluetoothBlocked() bool`

GetBluetoothBlocked returns the BluetoothBlocked field if non-nil, zero value otherwise.

### GetBluetoothBlockedOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetBluetoothBlockedOk() (bool, bool)`

GetBluetoothBlockedOk returns a tuple with the BluetoothBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasBluetoothBlocked() bool`

HasBluetoothBlocked returns a boolean if a field has been set.

### SetBluetoothBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetBluetoothBlocked(v bool)`

SetBluetoothBlocked gets a reference to the given bool and assigns it to the BluetoothBlocked field.

### GetCameraBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCameraBlocked() bool`

GetCameraBlocked returns the CameraBlocked field if non-nil, zero value otherwise.

### GetCameraBlockedOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCameraBlockedOk() (bool, bool)`

GetCameraBlockedOk returns a tuple with the CameraBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasCameraBlocked() bool`

HasCameraBlocked returns a boolean if a field has been set.

### SetCameraBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetCameraBlocked(v bool)`

SetCameraBlocked gets a reference to the given bool and assigns it to the CameraBlocked field.

### GetCellularBlockWifiTethering

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCellularBlockWifiTethering() bool`

GetCellularBlockWifiTethering returns the CellularBlockWifiTethering field if non-nil, zero value otherwise.

### GetCellularBlockWifiTetheringOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCellularBlockWifiTetheringOk() (bool, bool)`

GetCellularBlockWifiTetheringOk returns a tuple with the CellularBlockWifiTethering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockWifiTethering

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasCellularBlockWifiTethering() bool`

HasCellularBlockWifiTethering returns a boolean if a field has been set.

### SetCellularBlockWifiTethering

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetCellularBlockWifiTethering(v bool)`

SetCellularBlockWifiTethering gets a reference to the given bool and assigns it to the CellularBlockWifiTethering field.

### GetCompliantAppsList

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCompliantAppsList() []AnyOfmicrosoftGraphAppListItem`

GetCompliantAppsList returns the CompliantAppsList field if non-nil, zero value otherwise.

### GetCompliantAppsListOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCompliantAppsListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetCompliantAppsListOk returns a tuple with the CompliantAppsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppsList

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasCompliantAppsList() bool`

HasCompliantAppsList returns a boolean if a field has been set.

### SetCompliantAppsList

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetCompliantAppsList(v []AnyOfmicrosoftGraphAppListItem)`

SetCompliantAppsList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the CompliantAppsList field.

### GetCompliantAppListType

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCompliantAppListType() AnyOfmicrosoftGraphAppListType`

GetCompliantAppListType returns the CompliantAppListType field if non-nil, zero value otherwise.

### GetCompliantAppListTypeOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetCompliantAppListTypeOk() (AnyOfmicrosoftGraphAppListType, bool)`

GetCompliantAppListTypeOk returns a tuple with the CompliantAppListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppListType

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasCompliantAppListType() bool`

HasCompliantAppListType returns a boolean if a field has been set.

### SetCompliantAppListType

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetCompliantAppListType(v AnyOfmicrosoftGraphAppListType)`

SetCompliantAppListType gets a reference to the given AnyOfmicrosoftGraphAppListType and assigns it to the CompliantAppListType field.

### GetDiagnosticDataBlockSubmission

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDiagnosticDataBlockSubmission() bool`

GetDiagnosticDataBlockSubmission returns the DiagnosticDataBlockSubmission field if non-nil, zero value otherwise.

### GetDiagnosticDataBlockSubmissionOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetDiagnosticDataBlockSubmissionOk() (bool, bool)`

GetDiagnosticDataBlockSubmissionOk returns a tuple with the DiagnosticDataBlockSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticDataBlockSubmission

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasDiagnosticDataBlockSubmission() bool`

HasDiagnosticDataBlockSubmission returns a boolean if a field has been set.

### SetDiagnosticDataBlockSubmission

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetDiagnosticDataBlockSubmission(v bool)`

SetDiagnosticDataBlockSubmission gets a reference to the given bool and assigns it to the DiagnosticDataBlockSubmission field.

### GetEmailBlockAddingAccounts

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetEmailBlockAddingAccounts() bool`

GetEmailBlockAddingAccounts returns the EmailBlockAddingAccounts field if non-nil, zero value otherwise.

### GetEmailBlockAddingAccountsOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetEmailBlockAddingAccountsOk() (bool, bool)`

GetEmailBlockAddingAccountsOk returns a tuple with the EmailBlockAddingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailBlockAddingAccounts

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasEmailBlockAddingAccounts() bool`

HasEmailBlockAddingAccounts returns a boolean if a field has been set.

### SetEmailBlockAddingAccounts

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetEmailBlockAddingAccounts(v bool)`

SetEmailBlockAddingAccounts gets a reference to the given bool and assigns it to the EmailBlockAddingAccounts field.

### GetLocationServicesBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetLocationServicesBlocked() bool`

GetLocationServicesBlocked returns the LocationServicesBlocked field if non-nil, zero value otherwise.

### GetLocationServicesBlockedOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetLocationServicesBlockedOk() (bool, bool)`

GetLocationServicesBlockedOk returns a tuple with the LocationServicesBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationServicesBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasLocationServicesBlocked() bool`

HasLocationServicesBlocked returns a boolean if a field has been set.

### SetLocationServicesBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetLocationServicesBlocked(v bool)`

SetLocationServicesBlocked gets a reference to the given bool and assigns it to the LocationServicesBlocked field.

### GetMicrosoftAccountBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetMicrosoftAccountBlocked() bool`

GetMicrosoftAccountBlocked returns the MicrosoftAccountBlocked field if non-nil, zero value otherwise.

### GetMicrosoftAccountBlockedOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetMicrosoftAccountBlockedOk() (bool, bool)`

GetMicrosoftAccountBlockedOk returns a tuple with the MicrosoftAccountBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftAccountBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasMicrosoftAccountBlocked() bool`

HasMicrosoftAccountBlocked returns a boolean if a field has been set.

### SetMicrosoftAccountBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetMicrosoftAccountBlocked(v bool)`

SetMicrosoftAccountBlocked gets a reference to the given bool and assigns it to the MicrosoftAccountBlocked field.

### GetNfcBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetNfcBlocked() bool`

GetNfcBlocked returns the NfcBlocked field if non-nil, zero value otherwise.

### GetNfcBlockedOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetNfcBlockedOk() (bool, bool)`

GetNfcBlockedOk returns a tuple with the NfcBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNfcBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasNfcBlocked() bool`

HasNfcBlocked returns a boolean if a field has been set.

### SetNfcBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetNfcBlocked(v bool)`

SetNfcBlocked gets a reference to the given bool and assigns it to the NfcBlocked field.

### GetPasswordBlockSimple

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordExpirationDays

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordRequired

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetScreenCaptureBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetStorageBlockRemovableStorage

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetStorageBlockRemovableStorage() bool`

GetStorageBlockRemovableStorage returns the StorageBlockRemovableStorage field if non-nil, zero value otherwise.

### GetStorageBlockRemovableStorageOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetStorageBlockRemovableStorageOk() (bool, bool)`

GetStorageBlockRemovableStorageOk returns a tuple with the StorageBlockRemovableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageBlockRemovableStorage

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasStorageBlockRemovableStorage() bool`

HasStorageBlockRemovableStorage returns a boolean if a field has been set.

### SetStorageBlockRemovableStorage

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetStorageBlockRemovableStorage(v bool)`

SetStorageBlockRemovableStorage gets a reference to the given bool and assigns it to the StorageBlockRemovableStorage field.

### GetStorageRequireEncryption

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.

### GetWebBrowserBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWebBrowserBlocked() bool`

GetWebBrowserBlocked returns the WebBrowserBlocked field if non-nil, zero value otherwise.

### GetWebBrowserBlockedOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWebBrowserBlockedOk() (bool, bool)`

GetWebBrowserBlockedOk returns a tuple with the WebBrowserBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasWebBrowserBlocked() bool`

HasWebBrowserBlocked returns a boolean if a field has been set.

### SetWebBrowserBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetWebBrowserBlocked(v bool)`

SetWebBrowserBlocked gets a reference to the given bool and assigns it to the WebBrowserBlocked field.

### GetWifiBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWifiBlocked() bool`

GetWifiBlocked returns the WifiBlocked field if non-nil, zero value otherwise.

### GetWifiBlockedOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWifiBlockedOk() (bool, bool)`

GetWifiBlockedOk returns a tuple with the WifiBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWifiBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasWifiBlocked() bool`

HasWifiBlocked returns a boolean if a field has been set.

### SetWifiBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetWifiBlocked(v bool)`

SetWifiBlocked gets a reference to the given bool and assigns it to the WifiBlocked field.

### GetWifiBlockAutomaticConnectHotspots

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWifiBlockAutomaticConnectHotspots() bool`

GetWifiBlockAutomaticConnectHotspots returns the WifiBlockAutomaticConnectHotspots field if non-nil, zero value otherwise.

### GetWifiBlockAutomaticConnectHotspotsOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWifiBlockAutomaticConnectHotspotsOk() (bool, bool)`

GetWifiBlockAutomaticConnectHotspotsOk returns a tuple with the WifiBlockAutomaticConnectHotspots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWifiBlockAutomaticConnectHotspots

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasWifiBlockAutomaticConnectHotspots() bool`

HasWifiBlockAutomaticConnectHotspots returns a boolean if a field has been set.

### SetWifiBlockAutomaticConnectHotspots

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetWifiBlockAutomaticConnectHotspots(v bool)`

SetWifiBlockAutomaticConnectHotspots gets a reference to the given bool and assigns it to the WifiBlockAutomaticConnectHotspots field.

### GetWifiBlockHotspotReporting

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWifiBlockHotspotReporting() bool`

GetWifiBlockHotspotReporting returns the WifiBlockHotspotReporting field if non-nil, zero value otherwise.

### GetWifiBlockHotspotReportingOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWifiBlockHotspotReportingOk() (bool, bool)`

GetWifiBlockHotspotReportingOk returns a tuple with the WifiBlockHotspotReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWifiBlockHotspotReporting

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasWifiBlockHotspotReporting() bool`

HasWifiBlockHotspotReporting returns a boolean if a field has been set.

### SetWifiBlockHotspotReporting

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetWifiBlockHotspotReporting(v bool)`

SetWifiBlockHotspotReporting gets a reference to the given bool and assigns it to the WifiBlockHotspotReporting field.

### GetWindowsStoreBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWindowsStoreBlocked() bool`

GetWindowsStoreBlocked returns the WindowsStoreBlocked field if non-nil, zero value otherwise.

### GetWindowsStoreBlockedOk

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) GetWindowsStoreBlockedOk() (bool, bool)`

GetWindowsStoreBlockedOk returns a tuple with the WindowsStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsStoreBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) HasWindowsStoreBlocked() bool`

HasWindowsStoreBlocked returns a boolean if a field has been set.

### SetWindowsStoreBlocked

`func (o *MicrosoftGraphWindowsPhone81GeneralConfiguration) SetWindowsStoreBlocked(v bool)`

SetWindowsStoreBlocked gets a reference to the given bool and assigns it to the WindowsStoreBlocked field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



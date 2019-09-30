# WindowsPhone81GeneralConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetApplyOnlyToWindowsPhone81

`func (o *WindowsPhone81GeneralConfiguration) GetApplyOnlyToWindowsPhone81() bool`

GetApplyOnlyToWindowsPhone81 returns the ApplyOnlyToWindowsPhone81 field if non-nil, zero value otherwise.

### GetApplyOnlyToWindowsPhone81Ok

`func (o *WindowsPhone81GeneralConfiguration) GetApplyOnlyToWindowsPhone81Ok() (bool, bool)`

GetApplyOnlyToWindowsPhone81Ok returns a tuple with the ApplyOnlyToWindowsPhone81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplyOnlyToWindowsPhone81

`func (o *WindowsPhone81GeneralConfiguration) HasApplyOnlyToWindowsPhone81() bool`

HasApplyOnlyToWindowsPhone81 returns a boolean if a field has been set.

### SetApplyOnlyToWindowsPhone81

`func (o *WindowsPhone81GeneralConfiguration) SetApplyOnlyToWindowsPhone81(v bool)`

SetApplyOnlyToWindowsPhone81 gets a reference to the given bool and assigns it to the ApplyOnlyToWindowsPhone81 field.

### GetAppsBlockCopyPaste

`func (o *WindowsPhone81GeneralConfiguration) GetAppsBlockCopyPaste() bool`

GetAppsBlockCopyPaste returns the AppsBlockCopyPaste field if non-nil, zero value otherwise.

### GetAppsBlockCopyPasteOk

`func (o *WindowsPhone81GeneralConfiguration) GetAppsBlockCopyPasteOk() (bool, bool)`

GetAppsBlockCopyPasteOk returns a tuple with the AppsBlockCopyPaste field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppsBlockCopyPaste

`func (o *WindowsPhone81GeneralConfiguration) HasAppsBlockCopyPaste() bool`

HasAppsBlockCopyPaste returns a boolean if a field has been set.

### SetAppsBlockCopyPaste

`func (o *WindowsPhone81GeneralConfiguration) SetAppsBlockCopyPaste(v bool)`

SetAppsBlockCopyPaste gets a reference to the given bool and assigns it to the AppsBlockCopyPaste field.

### GetBluetoothBlocked

`func (o *WindowsPhone81GeneralConfiguration) GetBluetoothBlocked() bool`

GetBluetoothBlocked returns the BluetoothBlocked field if non-nil, zero value otherwise.

### GetBluetoothBlockedOk

`func (o *WindowsPhone81GeneralConfiguration) GetBluetoothBlockedOk() (bool, bool)`

GetBluetoothBlockedOk returns a tuple with the BluetoothBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBluetoothBlocked

`func (o *WindowsPhone81GeneralConfiguration) HasBluetoothBlocked() bool`

HasBluetoothBlocked returns a boolean if a field has been set.

### SetBluetoothBlocked

`func (o *WindowsPhone81GeneralConfiguration) SetBluetoothBlocked(v bool)`

SetBluetoothBlocked gets a reference to the given bool and assigns it to the BluetoothBlocked field.

### GetCameraBlocked

`func (o *WindowsPhone81GeneralConfiguration) GetCameraBlocked() bool`

GetCameraBlocked returns the CameraBlocked field if non-nil, zero value otherwise.

### GetCameraBlockedOk

`func (o *WindowsPhone81GeneralConfiguration) GetCameraBlockedOk() (bool, bool)`

GetCameraBlockedOk returns a tuple with the CameraBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCameraBlocked

`func (o *WindowsPhone81GeneralConfiguration) HasCameraBlocked() bool`

HasCameraBlocked returns a boolean if a field has been set.

### SetCameraBlocked

`func (o *WindowsPhone81GeneralConfiguration) SetCameraBlocked(v bool)`

SetCameraBlocked gets a reference to the given bool and assigns it to the CameraBlocked field.

### GetCellularBlockWifiTethering

`func (o *WindowsPhone81GeneralConfiguration) GetCellularBlockWifiTethering() bool`

GetCellularBlockWifiTethering returns the CellularBlockWifiTethering field if non-nil, zero value otherwise.

### GetCellularBlockWifiTetheringOk

`func (o *WindowsPhone81GeneralConfiguration) GetCellularBlockWifiTetheringOk() (bool, bool)`

GetCellularBlockWifiTetheringOk returns a tuple with the CellularBlockWifiTethering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellularBlockWifiTethering

`func (o *WindowsPhone81GeneralConfiguration) HasCellularBlockWifiTethering() bool`

HasCellularBlockWifiTethering returns a boolean if a field has been set.

### SetCellularBlockWifiTethering

`func (o *WindowsPhone81GeneralConfiguration) SetCellularBlockWifiTethering(v bool)`

SetCellularBlockWifiTethering gets a reference to the given bool and assigns it to the CellularBlockWifiTethering field.

### GetCompliantAppsList

`func (o *WindowsPhone81GeneralConfiguration) GetCompliantAppsList() []AnyOfmicrosoftGraphAppListItem`

GetCompliantAppsList returns the CompliantAppsList field if non-nil, zero value otherwise.

### GetCompliantAppsListOk

`func (o *WindowsPhone81GeneralConfiguration) GetCompliantAppsListOk() ([]AnyOfmicrosoftGraphAppListItem, bool)`

GetCompliantAppsListOk returns a tuple with the CompliantAppsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppsList

`func (o *WindowsPhone81GeneralConfiguration) HasCompliantAppsList() bool`

HasCompliantAppsList returns a boolean if a field has been set.

### SetCompliantAppsList

`func (o *WindowsPhone81GeneralConfiguration) SetCompliantAppsList(v []AnyOfmicrosoftGraphAppListItem)`

SetCompliantAppsList gets a reference to the given []AnyOfmicrosoftGraphAppListItem and assigns it to the CompliantAppsList field.

### GetCompliantAppListType

`func (o *WindowsPhone81GeneralConfiguration) GetCompliantAppListType() AnyOfmicrosoftGraphAppListType`

GetCompliantAppListType returns the CompliantAppListType field if non-nil, zero value otherwise.

### GetCompliantAppListTypeOk

`func (o *WindowsPhone81GeneralConfiguration) GetCompliantAppListTypeOk() (AnyOfmicrosoftGraphAppListType, bool)`

GetCompliantAppListTypeOk returns a tuple with the CompliantAppListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompliantAppListType

`func (o *WindowsPhone81GeneralConfiguration) HasCompliantAppListType() bool`

HasCompliantAppListType returns a boolean if a field has been set.

### SetCompliantAppListType

`func (o *WindowsPhone81GeneralConfiguration) SetCompliantAppListType(v AnyOfmicrosoftGraphAppListType)`

SetCompliantAppListType gets a reference to the given AnyOfmicrosoftGraphAppListType and assigns it to the CompliantAppListType field.

### GetDiagnosticDataBlockSubmission

`func (o *WindowsPhone81GeneralConfiguration) GetDiagnosticDataBlockSubmission() bool`

GetDiagnosticDataBlockSubmission returns the DiagnosticDataBlockSubmission field if non-nil, zero value otherwise.

### GetDiagnosticDataBlockSubmissionOk

`func (o *WindowsPhone81GeneralConfiguration) GetDiagnosticDataBlockSubmissionOk() (bool, bool)`

GetDiagnosticDataBlockSubmissionOk returns a tuple with the DiagnosticDataBlockSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiagnosticDataBlockSubmission

`func (o *WindowsPhone81GeneralConfiguration) HasDiagnosticDataBlockSubmission() bool`

HasDiagnosticDataBlockSubmission returns a boolean if a field has been set.

### SetDiagnosticDataBlockSubmission

`func (o *WindowsPhone81GeneralConfiguration) SetDiagnosticDataBlockSubmission(v bool)`

SetDiagnosticDataBlockSubmission gets a reference to the given bool and assigns it to the DiagnosticDataBlockSubmission field.

### GetEmailBlockAddingAccounts

`func (o *WindowsPhone81GeneralConfiguration) GetEmailBlockAddingAccounts() bool`

GetEmailBlockAddingAccounts returns the EmailBlockAddingAccounts field if non-nil, zero value otherwise.

### GetEmailBlockAddingAccountsOk

`func (o *WindowsPhone81GeneralConfiguration) GetEmailBlockAddingAccountsOk() (bool, bool)`

GetEmailBlockAddingAccountsOk returns a tuple with the EmailBlockAddingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailBlockAddingAccounts

`func (o *WindowsPhone81GeneralConfiguration) HasEmailBlockAddingAccounts() bool`

HasEmailBlockAddingAccounts returns a boolean if a field has been set.

### SetEmailBlockAddingAccounts

`func (o *WindowsPhone81GeneralConfiguration) SetEmailBlockAddingAccounts(v bool)`

SetEmailBlockAddingAccounts gets a reference to the given bool and assigns it to the EmailBlockAddingAccounts field.

### GetLocationServicesBlocked

`func (o *WindowsPhone81GeneralConfiguration) GetLocationServicesBlocked() bool`

GetLocationServicesBlocked returns the LocationServicesBlocked field if non-nil, zero value otherwise.

### GetLocationServicesBlockedOk

`func (o *WindowsPhone81GeneralConfiguration) GetLocationServicesBlockedOk() (bool, bool)`

GetLocationServicesBlockedOk returns a tuple with the LocationServicesBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocationServicesBlocked

`func (o *WindowsPhone81GeneralConfiguration) HasLocationServicesBlocked() bool`

HasLocationServicesBlocked returns a boolean if a field has been set.

### SetLocationServicesBlocked

`func (o *WindowsPhone81GeneralConfiguration) SetLocationServicesBlocked(v bool)`

SetLocationServicesBlocked gets a reference to the given bool and assigns it to the LocationServicesBlocked field.

### GetMicrosoftAccountBlocked

`func (o *WindowsPhone81GeneralConfiguration) GetMicrosoftAccountBlocked() bool`

GetMicrosoftAccountBlocked returns the MicrosoftAccountBlocked field if non-nil, zero value otherwise.

### GetMicrosoftAccountBlockedOk

`func (o *WindowsPhone81GeneralConfiguration) GetMicrosoftAccountBlockedOk() (bool, bool)`

GetMicrosoftAccountBlockedOk returns a tuple with the MicrosoftAccountBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftAccountBlocked

`func (o *WindowsPhone81GeneralConfiguration) HasMicrosoftAccountBlocked() bool`

HasMicrosoftAccountBlocked returns a boolean if a field has been set.

### SetMicrosoftAccountBlocked

`func (o *WindowsPhone81GeneralConfiguration) SetMicrosoftAccountBlocked(v bool)`

SetMicrosoftAccountBlocked gets a reference to the given bool and assigns it to the MicrosoftAccountBlocked field.

### GetNfcBlocked

`func (o *WindowsPhone81GeneralConfiguration) GetNfcBlocked() bool`

GetNfcBlocked returns the NfcBlocked field if non-nil, zero value otherwise.

### GetNfcBlockedOk

`func (o *WindowsPhone81GeneralConfiguration) GetNfcBlockedOk() (bool, bool)`

GetNfcBlockedOk returns a tuple with the NfcBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNfcBlocked

`func (o *WindowsPhone81GeneralConfiguration) HasNfcBlocked() bool`

HasNfcBlocked returns a boolean if a field has been set.

### SetNfcBlocked

`func (o *WindowsPhone81GeneralConfiguration) SetNfcBlocked(v bool)`

SetNfcBlocked gets a reference to the given bool and assigns it to the NfcBlocked field.

### GetPasswordBlockSimple

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordBlockSimple() bool`

GetPasswordBlockSimple returns the PasswordBlockSimple field if non-nil, zero value otherwise.

### GetPasswordBlockSimpleOk

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordBlockSimpleOk() (bool, bool)`

GetPasswordBlockSimpleOk returns a tuple with the PasswordBlockSimple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordBlockSimple

`func (o *WindowsPhone81GeneralConfiguration) HasPasswordBlockSimple() bool`

HasPasswordBlockSimple returns a boolean if a field has been set.

### SetPasswordBlockSimple

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordBlockSimple(v bool)`

SetPasswordBlockSimple gets a reference to the given bool and assigns it to the PasswordBlockSimple field.

### GetPasswordExpirationDays

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordExpirationDays() int32`

GetPasswordExpirationDays returns the PasswordExpirationDays field if non-nil, zero value otherwise.

### GetPasswordExpirationDaysOk

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordExpirationDaysOk() (int32, bool)`

GetPasswordExpirationDaysOk returns a tuple with the PasswordExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordExpirationDays

`func (o *WindowsPhone81GeneralConfiguration) HasPasswordExpirationDays() bool`

HasPasswordExpirationDays returns a boolean if a field has been set.

### SetPasswordExpirationDays

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordExpirationDays(v int32)`

SetPasswordExpirationDays gets a reference to the given int32 and assigns it to the PasswordExpirationDays field.

### SetPasswordExpirationDaysExplicitNull

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordExpirationDaysExplicitNull(b bool)`

SetPasswordExpirationDaysExplicitNull (un)sets PasswordExpirationDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordExpirationDays value is set to nil even if false is passed
### GetPasswordMinimumLength

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordMinimumLength() int32`

GetPasswordMinimumLength returns the PasswordMinimumLength field if non-nil, zero value otherwise.

### GetPasswordMinimumLengthOk

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordMinimumLengthOk() (int32, bool)`

GetPasswordMinimumLengthOk returns a tuple with the PasswordMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumLength

`func (o *WindowsPhone81GeneralConfiguration) HasPasswordMinimumLength() bool`

HasPasswordMinimumLength returns a boolean if a field has been set.

### SetPasswordMinimumLength

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordMinimumLength(v int32)`

SetPasswordMinimumLength gets a reference to the given int32 and assigns it to the PasswordMinimumLength field.

### SetPasswordMinimumLengthExplicitNull

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordMinimumLengthExplicitNull(b bool)`

SetPasswordMinimumLengthExplicitNull (un)sets PasswordMinimumLength to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumLength value is set to nil even if false is passed
### GetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout() int32`

GetPasswordMinutesOfInactivityBeforeScreenTimeout returns the PasswordMinutesOfInactivityBeforeScreenTimeout field if non-nil, zero value otherwise.

### GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk() (int32, bool)`

GetPasswordMinutesOfInactivityBeforeScreenTimeoutOk returns a tuple with the PasswordMinutesOfInactivityBeforeScreenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *WindowsPhone81GeneralConfiguration) HasPasswordMinutesOfInactivityBeforeScreenTimeout() bool`

HasPasswordMinutesOfInactivityBeforeScreenTimeout returns a boolean if a field has been set.

### SetPasswordMinutesOfInactivityBeforeScreenTimeout

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(v int32)`

SetPasswordMinutesOfInactivityBeforeScreenTimeout gets a reference to the given int32 and assigns it to the PasswordMinutesOfInactivityBeforeScreenTimeout field.

### SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull(b bool)`

SetPasswordMinutesOfInactivityBeforeScreenTimeoutExplicitNull (un)sets PasswordMinutesOfInactivityBeforeScreenTimeout to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinutesOfInactivityBeforeScreenTimeout value is set to nil even if false is passed
### GetPasswordMinimumCharacterSetCount

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordMinimumCharacterSetCount() int32`

GetPasswordMinimumCharacterSetCount returns the PasswordMinimumCharacterSetCount field if non-nil, zero value otherwise.

### GetPasswordMinimumCharacterSetCountOk

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordMinimumCharacterSetCountOk() (int32, bool)`

GetPasswordMinimumCharacterSetCountOk returns a tuple with the PasswordMinimumCharacterSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMinimumCharacterSetCount

`func (o *WindowsPhone81GeneralConfiguration) HasPasswordMinimumCharacterSetCount() bool`

HasPasswordMinimumCharacterSetCount returns a boolean if a field has been set.

### SetPasswordMinimumCharacterSetCount

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordMinimumCharacterSetCount(v int32)`

SetPasswordMinimumCharacterSetCount gets a reference to the given int32 and assigns it to the PasswordMinimumCharacterSetCount field.

### SetPasswordMinimumCharacterSetCountExplicitNull

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordMinimumCharacterSetCountExplicitNull(b bool)`

SetPasswordMinimumCharacterSetCountExplicitNull (un)sets PasswordMinimumCharacterSetCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordMinimumCharacterSetCount value is set to nil even if false is passed
### GetPasswordPreviousPasswordBlockCount

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordPreviousPasswordBlockCount() int32`

GetPasswordPreviousPasswordBlockCount returns the PasswordPreviousPasswordBlockCount field if non-nil, zero value otherwise.

### GetPasswordPreviousPasswordBlockCountOk

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordPreviousPasswordBlockCountOk() (int32, bool)`

GetPasswordPreviousPasswordBlockCountOk returns a tuple with the PasswordPreviousPasswordBlockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordPreviousPasswordBlockCount

`func (o *WindowsPhone81GeneralConfiguration) HasPasswordPreviousPasswordBlockCount() bool`

HasPasswordPreviousPasswordBlockCount returns a boolean if a field has been set.

### SetPasswordPreviousPasswordBlockCount

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordPreviousPasswordBlockCount(v int32)`

SetPasswordPreviousPasswordBlockCount gets a reference to the given int32 and assigns it to the PasswordPreviousPasswordBlockCount field.

### SetPasswordPreviousPasswordBlockCountExplicitNull

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordPreviousPasswordBlockCountExplicitNull(b bool)`

SetPasswordPreviousPasswordBlockCountExplicitNull (un)sets PasswordPreviousPasswordBlockCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordPreviousPasswordBlockCount value is set to nil even if false is passed
### GetPasswordSignInFailureCountBeforeFactoryReset

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset() int32`

GetPasswordSignInFailureCountBeforeFactoryReset returns the PasswordSignInFailureCountBeforeFactoryReset field if non-nil, zero value otherwise.

### GetPasswordSignInFailureCountBeforeFactoryResetOk

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryResetOk() (int32, bool)`

GetPasswordSignInFailureCountBeforeFactoryResetOk returns a tuple with the PasswordSignInFailureCountBeforeFactoryReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordSignInFailureCountBeforeFactoryReset

`func (o *WindowsPhone81GeneralConfiguration) HasPasswordSignInFailureCountBeforeFactoryReset() bool`

HasPasswordSignInFailureCountBeforeFactoryReset returns a boolean if a field has been set.

### SetPasswordSignInFailureCountBeforeFactoryReset

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(v int32)`

SetPasswordSignInFailureCountBeforeFactoryReset gets a reference to the given int32 and assigns it to the PasswordSignInFailureCountBeforeFactoryReset field.

### SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull(b bool)`

SetPasswordSignInFailureCountBeforeFactoryResetExplicitNull (un)sets PasswordSignInFailureCountBeforeFactoryReset to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PasswordSignInFailureCountBeforeFactoryReset value is set to nil even if false is passed
### GetPasswordRequiredType

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordRequiredType() AnyOfmicrosoftGraphRequiredPasswordType`

GetPasswordRequiredType returns the PasswordRequiredType field if non-nil, zero value otherwise.

### GetPasswordRequiredTypeOk

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordRequiredTypeOk() (AnyOfmicrosoftGraphRequiredPasswordType, bool)`

GetPasswordRequiredTypeOk returns a tuple with the PasswordRequiredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequiredType

`func (o *WindowsPhone81GeneralConfiguration) HasPasswordRequiredType() bool`

HasPasswordRequiredType returns a boolean if a field has been set.

### SetPasswordRequiredType

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordRequiredType(v AnyOfmicrosoftGraphRequiredPasswordType)`

SetPasswordRequiredType gets a reference to the given AnyOfmicrosoftGraphRequiredPasswordType and assigns it to the PasswordRequiredType field.

### GetPasswordRequired

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *WindowsPhone81GeneralConfiguration) GetPasswordRequiredOk() (bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordRequired

`func (o *WindowsPhone81GeneralConfiguration) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### SetPasswordRequired

`func (o *WindowsPhone81GeneralConfiguration) SetPasswordRequired(v bool)`

SetPasswordRequired gets a reference to the given bool and assigns it to the PasswordRequired field.

### GetScreenCaptureBlocked

`func (o *WindowsPhone81GeneralConfiguration) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *WindowsPhone81GeneralConfiguration) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *WindowsPhone81GeneralConfiguration) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *WindowsPhone81GeneralConfiguration) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetStorageBlockRemovableStorage

`func (o *WindowsPhone81GeneralConfiguration) GetStorageBlockRemovableStorage() bool`

GetStorageBlockRemovableStorage returns the StorageBlockRemovableStorage field if non-nil, zero value otherwise.

### GetStorageBlockRemovableStorageOk

`func (o *WindowsPhone81GeneralConfiguration) GetStorageBlockRemovableStorageOk() (bool, bool)`

GetStorageBlockRemovableStorageOk returns a tuple with the StorageBlockRemovableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageBlockRemovableStorage

`func (o *WindowsPhone81GeneralConfiguration) HasStorageBlockRemovableStorage() bool`

HasStorageBlockRemovableStorage returns a boolean if a field has been set.

### SetStorageBlockRemovableStorage

`func (o *WindowsPhone81GeneralConfiguration) SetStorageBlockRemovableStorage(v bool)`

SetStorageBlockRemovableStorage gets a reference to the given bool and assigns it to the StorageBlockRemovableStorage field.

### GetStorageRequireEncryption

`func (o *WindowsPhone81GeneralConfiguration) GetStorageRequireEncryption() bool`

GetStorageRequireEncryption returns the StorageRequireEncryption field if non-nil, zero value otherwise.

### GetStorageRequireEncryptionOk

`func (o *WindowsPhone81GeneralConfiguration) GetStorageRequireEncryptionOk() (bool, bool)`

GetStorageRequireEncryptionOk returns a tuple with the StorageRequireEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStorageRequireEncryption

`func (o *WindowsPhone81GeneralConfiguration) HasStorageRequireEncryption() bool`

HasStorageRequireEncryption returns a boolean if a field has been set.

### SetStorageRequireEncryption

`func (o *WindowsPhone81GeneralConfiguration) SetStorageRequireEncryption(v bool)`

SetStorageRequireEncryption gets a reference to the given bool and assigns it to the StorageRequireEncryption field.

### GetWebBrowserBlocked

`func (o *WindowsPhone81GeneralConfiguration) GetWebBrowserBlocked() bool`

GetWebBrowserBlocked returns the WebBrowserBlocked field if non-nil, zero value otherwise.

### GetWebBrowserBlockedOk

`func (o *WindowsPhone81GeneralConfiguration) GetWebBrowserBlockedOk() (bool, bool)`

GetWebBrowserBlockedOk returns a tuple with the WebBrowserBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebBrowserBlocked

`func (o *WindowsPhone81GeneralConfiguration) HasWebBrowserBlocked() bool`

HasWebBrowserBlocked returns a boolean if a field has been set.

### SetWebBrowserBlocked

`func (o *WindowsPhone81GeneralConfiguration) SetWebBrowserBlocked(v bool)`

SetWebBrowserBlocked gets a reference to the given bool and assigns it to the WebBrowserBlocked field.

### GetWifiBlocked

`func (o *WindowsPhone81GeneralConfiguration) GetWifiBlocked() bool`

GetWifiBlocked returns the WifiBlocked field if non-nil, zero value otherwise.

### GetWifiBlockedOk

`func (o *WindowsPhone81GeneralConfiguration) GetWifiBlockedOk() (bool, bool)`

GetWifiBlockedOk returns a tuple with the WifiBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWifiBlocked

`func (o *WindowsPhone81GeneralConfiguration) HasWifiBlocked() bool`

HasWifiBlocked returns a boolean if a field has been set.

### SetWifiBlocked

`func (o *WindowsPhone81GeneralConfiguration) SetWifiBlocked(v bool)`

SetWifiBlocked gets a reference to the given bool and assigns it to the WifiBlocked field.

### GetWifiBlockAutomaticConnectHotspots

`func (o *WindowsPhone81GeneralConfiguration) GetWifiBlockAutomaticConnectHotspots() bool`

GetWifiBlockAutomaticConnectHotspots returns the WifiBlockAutomaticConnectHotspots field if non-nil, zero value otherwise.

### GetWifiBlockAutomaticConnectHotspotsOk

`func (o *WindowsPhone81GeneralConfiguration) GetWifiBlockAutomaticConnectHotspotsOk() (bool, bool)`

GetWifiBlockAutomaticConnectHotspotsOk returns a tuple with the WifiBlockAutomaticConnectHotspots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWifiBlockAutomaticConnectHotspots

`func (o *WindowsPhone81GeneralConfiguration) HasWifiBlockAutomaticConnectHotspots() bool`

HasWifiBlockAutomaticConnectHotspots returns a boolean if a field has been set.

### SetWifiBlockAutomaticConnectHotspots

`func (o *WindowsPhone81GeneralConfiguration) SetWifiBlockAutomaticConnectHotspots(v bool)`

SetWifiBlockAutomaticConnectHotspots gets a reference to the given bool and assigns it to the WifiBlockAutomaticConnectHotspots field.

### GetWifiBlockHotspotReporting

`func (o *WindowsPhone81GeneralConfiguration) GetWifiBlockHotspotReporting() bool`

GetWifiBlockHotspotReporting returns the WifiBlockHotspotReporting field if non-nil, zero value otherwise.

### GetWifiBlockHotspotReportingOk

`func (o *WindowsPhone81GeneralConfiguration) GetWifiBlockHotspotReportingOk() (bool, bool)`

GetWifiBlockHotspotReportingOk returns a tuple with the WifiBlockHotspotReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWifiBlockHotspotReporting

`func (o *WindowsPhone81GeneralConfiguration) HasWifiBlockHotspotReporting() bool`

HasWifiBlockHotspotReporting returns a boolean if a field has been set.

### SetWifiBlockHotspotReporting

`func (o *WindowsPhone81GeneralConfiguration) SetWifiBlockHotspotReporting(v bool)`

SetWifiBlockHotspotReporting gets a reference to the given bool and assigns it to the WifiBlockHotspotReporting field.

### GetWindowsStoreBlocked

`func (o *WindowsPhone81GeneralConfiguration) GetWindowsStoreBlocked() bool`

GetWindowsStoreBlocked returns the WindowsStoreBlocked field if non-nil, zero value otherwise.

### GetWindowsStoreBlockedOk

`func (o *WindowsPhone81GeneralConfiguration) GetWindowsStoreBlockedOk() (bool, bool)`

GetWindowsStoreBlockedOk returns a tuple with the WindowsStoreBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsStoreBlocked

`func (o *WindowsPhone81GeneralConfiguration) HasWindowsStoreBlocked() bool`

HasWindowsStoreBlocked returns a boolean if a field has been set.

### SetWindowsStoreBlocked

`func (o *WindowsPhone81GeneralConfiguration) SetWindowsStoreBlocked(v bool)`

SetWindowsStoreBlocked gets a reference to the given bool and assigns it to the WindowsStoreBlocked field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



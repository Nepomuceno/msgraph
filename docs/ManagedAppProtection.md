# ManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodOfflineBeforeAccessCheck** | Pointer to **string** | The period after which access is checked when the device is not connected to the internet. | [optional] 
**PeriodOnlineBeforeAccessCheck** | Pointer to **string** | The period after which access is checked when the device is connected to the internet. | [optional] 
**AllowedInboundDataTransferSources** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataTransferLevel**](anyOf&lt;microsoft.graph.managedAppDataTransferLevel&gt;.md) | Sources from which data is allowed to be transferred. | [optional] 
**AllowedOutboundDataTransferDestinations** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataTransferLevel**](anyOf&lt;microsoft.graph.managedAppDataTransferLevel&gt;.md) | Destinations to which data is allowed to be transferred. | [optional] 
**OrganizationalCredentialsRequired** | Pointer to **bool** | Indicates whether organizational credentials are required for app use. | [optional] 
**AllowedOutboundClipboardSharingLevel** | Pointer to [**AnyOfmicrosoftGraphManagedAppClipboardSharingLevel**](anyOf&lt;microsoft.graph.managedAppClipboardSharingLevel&gt;.md) | The level to which the clipboard may be shared between apps on the managed device. | [optional] 
**DataBackupBlocked** | Pointer to **bool** | Indicates whether the backup of a managed app&#39;s data is blocked. | [optional] 
**DeviceComplianceRequired** | Pointer to **bool** | Indicates whether device compliance is required. | [optional] 
**ManagedBrowserToOpenLinksRequired** | Pointer to **bool** | Indicates whether internet links should be opened in the managed browser app. | [optional] 
**SaveAsBlocked** | Pointer to **bool** | Indicates whether users may use the \&quot;Save As\&quot; menu item to save a copy of protected files. | [optional] 
**PeriodOfflineBeforeWipeIsEnforced** | Pointer to **string** | The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped. | [optional] 
**PinRequired** | Pointer to **bool** | Indicates whether an app-level pin is required. | [optional] 
**MaximumPinRetries** | Pointer to **int32** | Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped. | [optional] 
**SimplePinBlocked** | Pointer to **bool** | Indicates whether simplePin is blocked. | [optional] 
**MinimumPinLength** | Pointer to **int32** | Minimum pin length required for an app-level pin if PinRequired is set to True | [optional] 
**PinCharacterSet** | Pointer to [**AnyOfmicrosoftGraphManagedAppPinCharacterSet**](anyOf&lt;microsoft.graph.managedAppPinCharacterSet&gt;.md) | Character set which may be used for an app-level pin if PinRequired is set to True. | [optional] 
**PeriodBeforePinReset** | Pointer to **string** | TimePeriod before the all-level pin must be reset if PinRequired is set to True. | [optional] 
**AllowedDataStorageLocations** | Pointer to [**[]AnyOfmicrosoftGraphManagedAppDataStorageLocation**](anyOf&lt;microsoft.graph.managedAppDataStorageLocation&gt;.md) | Data storage locations where a user may store managed data. | [optional] 
**ContactSyncBlocked** | Pointer to **bool** | Indicates whether contacts can be synced to the user&#39;s device. | [optional] 
**PrintBlocked** | Pointer to **bool** | Indicates whether printing is allowed from managed apps. | [optional] 
**FingerprintBlocked** | Pointer to **bool** | Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True. | [optional] 
**DisableAppPinIfDevicePinIsSet** | Pointer to **bool** | Indicates whether use of the app pin is required if the device pin is set. | [optional] 
**MinimumRequiredOsVersion** | Pointer to **string** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**MinimumWarningOsVersion** | Pointer to **string** | Versions less than the specified version will result in warning message on the managed app from accessing company data. | [optional] 
**MinimumRequiredAppVersion** | Pointer to **string** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**MinimumWarningAppVersion** | Pointer to **string** | Versions less than the specified version will result in warning message on the managed app. | [optional] 

## Methods

### GetPeriodOfflineBeforeAccessCheck

`func (o *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeAccessCheck

`func (o *ManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *ManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOfflineBeforeAccessCheck field.

### GetPeriodOnlineBeforeAccessCheck

`func (o *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOnlineBeforeAccessCheck

`func (o *ManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *ManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOnlineBeforeAccessCheck field.

### GetAllowedInboundDataTransferSources

`func (o *ManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *ManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedInboundDataTransferSources

`func (o *ManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSources

`func (o *ManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedInboundDataTransferSources field.

### GetAllowedOutboundDataTransferDestinations

`func (o *ManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *ManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundDataTransferDestinations

`func (o *ManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *ManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedOutboundDataTransferDestinations field.

### GetOrganizationalCredentialsRequired

`func (o *ManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *ManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationalCredentialsRequired

`func (o *ManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### SetOrganizationalCredentialsRequired

`func (o *ManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired gets a reference to the given bool and assigns it to the OrganizationalCredentialsRequired field.

### GetAllowedOutboundClipboardSharingLevel

`func (o *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundClipboardSharingLevel

`func (o *ManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *ManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel gets a reference to the given AnyOfmicrosoftGraphManagedAppClipboardSharingLevel and assigns it to the AllowedOutboundClipboardSharingLevel field.

### GetDataBackupBlocked

`func (o *ManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *ManagedAppProtection) GetDataBackupBlockedOk() (bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataBackupBlocked

`func (o *ManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### SetDataBackupBlocked

`func (o *ManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked gets a reference to the given bool and assigns it to the DataBackupBlocked field.

### GetDeviceComplianceRequired

`func (o *ManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *ManagedAppProtection) GetDeviceComplianceRequiredOk() (bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceComplianceRequired

`func (o *ManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### SetDeviceComplianceRequired

`func (o *ManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired gets a reference to the given bool and assigns it to the DeviceComplianceRequired field.

### GetManagedBrowserToOpenLinksRequired

`func (o *ManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *ManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedBrowserToOpenLinksRequired

`func (o *ManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *ManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired gets a reference to the given bool and assigns it to the ManagedBrowserToOpenLinksRequired field.

### GetSaveAsBlocked

`func (o *ManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *ManagedAppProtection) GetSaveAsBlockedOk() (bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSaveAsBlocked

`func (o *ManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### SetSaveAsBlocked

`func (o *ManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked gets a reference to the given bool and assigns it to the SaveAsBlocked field.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *ManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *ManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced gets a reference to the given string and assigns it to the PeriodOfflineBeforeWipeIsEnforced field.

### GetPinRequired

`func (o *ManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *ManagedAppProtection) GetPinRequiredOk() (bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinRequired

`func (o *ManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### SetPinRequired

`func (o *ManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired gets a reference to the given bool and assigns it to the PinRequired field.

### GetMaximumPinRetries

`func (o *ManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *ManagedAppProtection) GetMaximumPinRetriesOk() (int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaximumPinRetries

`func (o *ManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### SetMaximumPinRetries

`func (o *ManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries gets a reference to the given int32 and assigns it to the MaximumPinRetries field.

### GetSimplePinBlocked

`func (o *ManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *ManagedAppProtection) GetSimplePinBlockedOk() (bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSimplePinBlocked

`func (o *ManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.

### SetSimplePinBlocked

`func (o *ManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked gets a reference to the given bool and assigns it to the SimplePinBlocked field.

### GetMinimumPinLength

`func (o *ManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *ManagedAppProtection) GetMinimumPinLengthOk() (int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumPinLength

`func (o *ManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### SetMinimumPinLength

`func (o *ManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength gets a reference to the given int32 and assigns it to the MinimumPinLength field.

### GetPinCharacterSet

`func (o *ManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *ManagedAppProtection) GetPinCharacterSetOk() (AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinCharacterSet

`func (o *ManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSet

`func (o *ManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet gets a reference to the given AnyOfmicrosoftGraphManagedAppPinCharacterSet and assigns it to the PinCharacterSet field.

### GetPeriodBeforePinReset

`func (o *ManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *ManagedAppProtection) GetPeriodBeforePinResetOk() (string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodBeforePinReset

`func (o *ManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### SetPeriodBeforePinReset

`func (o *ManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset gets a reference to the given string and assigns it to the PeriodBeforePinReset field.

### GetAllowedDataStorageLocations

`func (o *ManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *ManagedAppProtection) GetAllowedDataStorageLocationsOk() ([]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedDataStorageLocations

`func (o *ManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### SetAllowedDataStorageLocations

`func (o *ManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations gets a reference to the given []AnyOfmicrosoftGraphManagedAppDataStorageLocation and assigns it to the AllowedDataStorageLocations field.

### GetContactSyncBlocked

`func (o *ManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *ManagedAppProtection) GetContactSyncBlockedOk() (bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactSyncBlocked

`func (o *ManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### SetContactSyncBlocked

`func (o *ManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked gets a reference to the given bool and assigns it to the ContactSyncBlocked field.

### GetPrintBlocked

`func (o *ManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *ManagedAppProtection) GetPrintBlockedOk() (bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrintBlocked

`func (o *ManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### SetPrintBlocked

`func (o *ManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked gets a reference to the given bool and assigns it to the PrintBlocked field.

### GetFingerprintBlocked

`func (o *ManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *ManagedAppProtection) GetFingerprintBlockedOk() (bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFingerprintBlocked

`func (o *ManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### SetFingerprintBlocked

`func (o *ManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked gets a reference to the given bool and assigns it to the FingerprintBlocked field.

### GetDisableAppPinIfDevicePinIsSet

`func (o *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAppPinIfDevicePinIsSet

`func (o *ManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *ManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet gets a reference to the given bool and assigns it to the DisableAppPinIfDevicePinIsSet field.

### GetMinimumRequiredOsVersion

`func (o *ManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *ManagedAppProtection) GetMinimumRequiredOsVersionOk() (string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredOsVersion

`func (o *ManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersion

`func (o *ManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion gets a reference to the given string and assigns it to the MinimumRequiredOsVersion field.

### SetMinimumRequiredOsVersionExplicitNull

`func (o *ManagedAppProtection) SetMinimumRequiredOsVersionExplicitNull(b bool)`

SetMinimumRequiredOsVersionExplicitNull (un)sets MinimumRequiredOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredOsVersion value is set to nil even if false is passed
### GetMinimumWarningOsVersion

`func (o *ManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *ManagedAppProtection) GetMinimumWarningOsVersionOk() (string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningOsVersion

`func (o *ManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersion

`func (o *ManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion gets a reference to the given string and assigns it to the MinimumWarningOsVersion field.

### SetMinimumWarningOsVersionExplicitNull

`func (o *ManagedAppProtection) SetMinimumWarningOsVersionExplicitNull(b bool)`

SetMinimumWarningOsVersionExplicitNull (un)sets MinimumWarningOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningOsVersion value is set to nil even if false is passed
### GetMinimumRequiredAppVersion

`func (o *ManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *ManagedAppProtection) GetMinimumRequiredAppVersionOk() (string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredAppVersion

`func (o *ManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersion

`func (o *ManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion gets a reference to the given string and assigns it to the MinimumRequiredAppVersion field.

### SetMinimumRequiredAppVersionExplicitNull

`func (o *ManagedAppProtection) SetMinimumRequiredAppVersionExplicitNull(b bool)`

SetMinimumRequiredAppVersionExplicitNull (un)sets MinimumRequiredAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredAppVersion value is set to nil even if false is passed
### GetMinimumWarningAppVersion

`func (o *ManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *ManagedAppProtection) GetMinimumWarningAppVersionOk() (string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningAppVersion

`func (o *ManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersion

`func (o *ManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion gets a reference to the given string and assigns it to the MinimumWarningAppVersion field.

### SetMinimumWarningAppVersionExplicitNull

`func (o *ManagedAppProtection) SetMinimumWarningAppVersionExplicitNull(b bool)`

SetMinimumWarningAppVersionExplicitNull (un)sets MinimumWarningAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningAppVersion value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



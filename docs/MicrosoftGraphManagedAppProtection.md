# MicrosoftGraphManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Policy display name. | [optional] 
**Description** | Pointer to **string** | The policy&#39;s description. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time the policy was created. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last time the policy was modified. | [optional] 
**Version** | Pointer to **string** | Version of the entity. | [optional] 
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

### GetId

`func (o *MicrosoftGraphManagedAppProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedAppProtection) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedAppProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedAppProtection) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphManagedAppProtection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedAppProtection) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphManagedAppProtection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedAppProtection) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphManagedAppProtection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphManagedAppProtection) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphManagedAppProtection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphManagedAppProtection) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphManagedAppProtection) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphManagedAppProtection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphManagedAppProtection) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphManagedAppProtection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphManagedAppProtection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphManagedAppProtection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphManagedAppProtection) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphManagedAppProtection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphManagedAppProtection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphManagedAppProtection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedAppProtection) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphManagedAppProtection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphManagedAppProtection) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphManagedAppProtection) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *MicrosoftGraphManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOfflineBeforeAccessCheck field.

### GetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *MicrosoftGraphManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOnlineBeforeAccessCheck field.

### GetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *MicrosoftGraphManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedInboundDataTransferSources

`func (o *MicrosoftGraphManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedInboundDataTransferSources field.

### GetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *MicrosoftGraphManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedOutboundDataTransferDestinations field.

### GetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *MicrosoftGraphManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationalCredentialsRequired

`func (o *MicrosoftGraphManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### SetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired gets a reference to the given bool and assigns it to the OrganizationalCredentialsRequired field.

### GetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *MicrosoftGraphManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel gets a reference to the given AnyOfmicrosoftGraphManagedAppClipboardSharingLevel and assigns it to the AllowedOutboundClipboardSharingLevel field.

### GetDataBackupBlocked

`func (o *MicrosoftGraphManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *MicrosoftGraphManagedAppProtection) GetDataBackupBlockedOk() (bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataBackupBlocked

`func (o *MicrosoftGraphManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### SetDataBackupBlocked

`func (o *MicrosoftGraphManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked gets a reference to the given bool and assigns it to the DataBackupBlocked field.

### GetDeviceComplianceRequired

`func (o *MicrosoftGraphManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *MicrosoftGraphManagedAppProtection) GetDeviceComplianceRequiredOk() (bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceComplianceRequired

`func (o *MicrosoftGraphManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### SetDeviceComplianceRequired

`func (o *MicrosoftGraphManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired gets a reference to the given bool and assigns it to the DeviceComplianceRequired field.

### GetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *MicrosoftGraphManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired gets a reference to the given bool and assigns it to the ManagedBrowserToOpenLinksRequired field.

### GetSaveAsBlocked

`func (o *MicrosoftGraphManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *MicrosoftGraphManagedAppProtection) GetSaveAsBlockedOk() (bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSaveAsBlocked

`func (o *MicrosoftGraphManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### SetSaveAsBlocked

`func (o *MicrosoftGraphManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked gets a reference to the given bool and assigns it to the SaveAsBlocked field.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *MicrosoftGraphManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced gets a reference to the given string and assigns it to the PeriodOfflineBeforeWipeIsEnforced field.

### GetPinRequired

`func (o *MicrosoftGraphManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *MicrosoftGraphManagedAppProtection) GetPinRequiredOk() (bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinRequired

`func (o *MicrosoftGraphManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### SetPinRequired

`func (o *MicrosoftGraphManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired gets a reference to the given bool and assigns it to the PinRequired field.

### GetMaximumPinRetries

`func (o *MicrosoftGraphManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *MicrosoftGraphManagedAppProtection) GetMaximumPinRetriesOk() (int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaximumPinRetries

`func (o *MicrosoftGraphManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### SetMaximumPinRetries

`func (o *MicrosoftGraphManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries gets a reference to the given int32 and assigns it to the MaximumPinRetries field.

### GetSimplePinBlocked

`func (o *MicrosoftGraphManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *MicrosoftGraphManagedAppProtection) GetSimplePinBlockedOk() (bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSimplePinBlocked

`func (o *MicrosoftGraphManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.

### SetSimplePinBlocked

`func (o *MicrosoftGraphManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked gets a reference to the given bool and assigns it to the SimplePinBlocked field.

### GetMinimumPinLength

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumPinLengthOk() (int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumPinLength

`func (o *MicrosoftGraphManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### SetMinimumPinLength

`func (o *MicrosoftGraphManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength gets a reference to the given int32 and assigns it to the MinimumPinLength field.

### GetPinCharacterSet

`func (o *MicrosoftGraphManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *MicrosoftGraphManagedAppProtection) GetPinCharacterSetOk() (AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinCharacterSet

`func (o *MicrosoftGraphManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSet

`func (o *MicrosoftGraphManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet gets a reference to the given AnyOfmicrosoftGraphManagedAppPinCharacterSet and assigns it to the PinCharacterSet field.

### GetPeriodBeforePinReset

`func (o *MicrosoftGraphManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *MicrosoftGraphManagedAppProtection) GetPeriodBeforePinResetOk() (string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodBeforePinReset

`func (o *MicrosoftGraphManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### SetPeriodBeforePinReset

`func (o *MicrosoftGraphManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset gets a reference to the given string and assigns it to the PeriodBeforePinReset field.

### GetAllowedDataStorageLocations

`func (o *MicrosoftGraphManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *MicrosoftGraphManagedAppProtection) GetAllowedDataStorageLocationsOk() ([]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedDataStorageLocations

`func (o *MicrosoftGraphManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### SetAllowedDataStorageLocations

`func (o *MicrosoftGraphManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations gets a reference to the given []AnyOfmicrosoftGraphManagedAppDataStorageLocation and assigns it to the AllowedDataStorageLocations field.

### GetContactSyncBlocked

`func (o *MicrosoftGraphManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *MicrosoftGraphManagedAppProtection) GetContactSyncBlockedOk() (bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactSyncBlocked

`func (o *MicrosoftGraphManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### SetContactSyncBlocked

`func (o *MicrosoftGraphManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked gets a reference to the given bool and assigns it to the ContactSyncBlocked field.

### GetPrintBlocked

`func (o *MicrosoftGraphManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *MicrosoftGraphManagedAppProtection) GetPrintBlockedOk() (bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrintBlocked

`func (o *MicrosoftGraphManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### SetPrintBlocked

`func (o *MicrosoftGraphManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked gets a reference to the given bool and assigns it to the PrintBlocked field.

### GetFingerprintBlocked

`func (o *MicrosoftGraphManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *MicrosoftGraphManagedAppProtection) GetFingerprintBlockedOk() (bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFingerprintBlocked

`func (o *MicrosoftGraphManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### SetFingerprintBlocked

`func (o *MicrosoftGraphManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked gets a reference to the given bool and assigns it to the FingerprintBlocked field.

### GetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *MicrosoftGraphManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet gets a reference to the given bool and assigns it to the DisableAppPinIfDevicePinIsSet field.

### GetMinimumRequiredOsVersion

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumRequiredOsVersionOk() (string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredOsVersion

`func (o *MicrosoftGraphManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersion

`func (o *MicrosoftGraphManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion gets a reference to the given string and assigns it to the MinimumRequiredOsVersion field.

### SetMinimumRequiredOsVersionExplicitNull

`func (o *MicrosoftGraphManagedAppProtection) SetMinimumRequiredOsVersionExplicitNull(b bool)`

SetMinimumRequiredOsVersionExplicitNull (un)sets MinimumRequiredOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredOsVersion value is set to nil even if false is passed
### GetMinimumWarningOsVersion

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumWarningOsVersionOk() (string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningOsVersion

`func (o *MicrosoftGraphManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersion

`func (o *MicrosoftGraphManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion gets a reference to the given string and assigns it to the MinimumWarningOsVersion field.

### SetMinimumWarningOsVersionExplicitNull

`func (o *MicrosoftGraphManagedAppProtection) SetMinimumWarningOsVersionExplicitNull(b bool)`

SetMinimumWarningOsVersionExplicitNull (un)sets MinimumWarningOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningOsVersion value is set to nil even if false is passed
### GetMinimumRequiredAppVersion

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumRequiredAppVersionOk() (string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredAppVersion

`func (o *MicrosoftGraphManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersion

`func (o *MicrosoftGraphManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion gets a reference to the given string and assigns it to the MinimumRequiredAppVersion field.

### SetMinimumRequiredAppVersionExplicitNull

`func (o *MicrosoftGraphManagedAppProtection) SetMinimumRequiredAppVersionExplicitNull(b bool)`

SetMinimumRequiredAppVersionExplicitNull (un)sets MinimumRequiredAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredAppVersion value is set to nil even if false is passed
### GetMinimumWarningAppVersion

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *MicrosoftGraphManagedAppProtection) GetMinimumWarningAppVersionOk() (string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningAppVersion

`func (o *MicrosoftGraphManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersion

`func (o *MicrosoftGraphManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion gets a reference to the given string and assigns it to the MinimumWarningAppVersion field.

### SetMinimumWarningAppVersionExplicitNull

`func (o *MicrosoftGraphManagedAppProtection) SetMinimumWarningAppVersionExplicitNull(b bool)`

SetMinimumWarningAppVersionExplicitNull (un)sets MinimumWarningAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningAppVersion value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



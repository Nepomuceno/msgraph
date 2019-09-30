# MicrosoftGraphAndroidManagedAppProtection

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
**IsAssigned** | Pointer to **bool** | Indicates if the policy is deployed to any inclusion groups or not. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphTargetedManagedAppPolicyAssignment**](microsoft.graph.targetedManagedAppPolicyAssignment.md) |  | [optional] 
**ScreenCaptureBlocked** | Pointer to **bool** | Indicates whether a managed user can take screen captures of managed apps | [optional] 
**DisableAppEncryptionIfDeviceEncryptionIsEnabled** | Pointer to **bool** | When this setting is enabled, app level encryption is disabled if device level encryption is enabled | [optional] 
**EncryptAppData** | Pointer to **bool** | Indicates whether application data for managed apps should be encrypted | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**MinimumRequiredPatchVersion** | Pointer to **string** | Define the oldest required Android security patch level a user can have to gain secure access to the app. | [optional] 
**MinimumWarningPatchVersion** | Pointer to **string** | Define the oldest recommended Android security patch level a user can have for secure access to the app. | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOfflineBeforeAccessCheck field.

### GetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOnlineBeforeAccessCheck field.

### GetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedInboundDataTransferSources

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedInboundDataTransferSources field.

### GetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedOutboundDataTransferDestinations field.

### GetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationalCredentialsRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### SetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired gets a reference to the given bool and assigns it to the OrganizationalCredentialsRequired field.

### GetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel gets a reference to the given AnyOfmicrosoftGraphManagedAppClipboardSharingLevel and assigns it to the AllowedOutboundClipboardSharingLevel field.

### GetDataBackupBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDataBackupBlockedOk() (bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataBackupBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### SetDataBackupBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked gets a reference to the given bool and assigns it to the DataBackupBlocked field.

### GetDeviceComplianceRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeviceComplianceRequiredOk() (bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceComplianceRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### SetDeviceComplianceRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired gets a reference to the given bool and assigns it to the DeviceComplianceRequired field.

### GetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired gets a reference to the given bool and assigns it to the ManagedBrowserToOpenLinksRequired field.

### GetSaveAsBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetSaveAsBlockedOk() (bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSaveAsBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### SetSaveAsBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked gets a reference to the given bool and assigns it to the SaveAsBlocked field.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced gets a reference to the given string and assigns it to the PeriodOfflineBeforeWipeIsEnforced field.

### GetPinRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPinRequiredOk() (bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### SetPinRequired

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired gets a reference to the given bool and assigns it to the PinRequired field.

### GetMaximumPinRetries

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMaximumPinRetriesOk() (int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaximumPinRetries

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### SetMaximumPinRetries

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries gets a reference to the given int32 and assigns it to the MaximumPinRetries field.

### GetSimplePinBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetSimplePinBlockedOk() (bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSimplePinBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.

### SetSimplePinBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked gets a reference to the given bool and assigns it to the SimplePinBlocked field.

### GetMinimumPinLength

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumPinLengthOk() (int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumPinLength

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### SetMinimumPinLength

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength gets a reference to the given int32 and assigns it to the MinimumPinLength field.

### GetPinCharacterSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPinCharacterSetOk() (AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinCharacterSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet gets a reference to the given AnyOfmicrosoftGraphManagedAppPinCharacterSet and assigns it to the PinCharacterSet field.

### GetPeriodBeforePinReset

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPeriodBeforePinResetOk() (string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodBeforePinReset

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### SetPeriodBeforePinReset

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset gets a reference to the given string and assigns it to the PeriodBeforePinReset field.

### GetAllowedDataStorageLocations

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAllowedDataStorageLocationsOk() ([]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedDataStorageLocations

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### SetAllowedDataStorageLocations

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations gets a reference to the given []AnyOfmicrosoftGraphManagedAppDataStorageLocation and assigns it to the AllowedDataStorageLocations field.

### GetContactSyncBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetContactSyncBlockedOk() (bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactSyncBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### SetContactSyncBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked gets a reference to the given bool and assigns it to the ContactSyncBlocked field.

### GetPrintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetPrintBlockedOk() (bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### SetPrintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked gets a reference to the given bool and assigns it to the PrintBlocked field.

### GetFingerprintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetFingerprintBlockedOk() (bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFingerprintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### SetFingerprintBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked gets a reference to the given bool and assigns it to the FingerprintBlocked field.

### GetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet gets a reference to the given bool and assigns it to the DisableAppPinIfDevicePinIsSet field.

### GetMinimumRequiredOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredOsVersionOk() (string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion gets a reference to the given string and assigns it to the MinimumRequiredOsVersion field.

### SetMinimumRequiredOsVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredOsVersionExplicitNull(b bool)`

SetMinimumRequiredOsVersionExplicitNull (un)sets MinimumRequiredOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredOsVersion value is set to nil even if false is passed
### GetMinimumWarningOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningOsVersionOk() (string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion gets a reference to the given string and assigns it to the MinimumWarningOsVersion field.

### SetMinimumWarningOsVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningOsVersionExplicitNull(b bool)`

SetMinimumWarningOsVersionExplicitNull (un)sets MinimumWarningOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningOsVersion value is set to nil even if false is passed
### GetMinimumRequiredAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredAppVersionOk() (string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion gets a reference to the given string and assigns it to the MinimumRequiredAppVersion field.

### SetMinimumRequiredAppVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredAppVersionExplicitNull(b bool)`

SetMinimumRequiredAppVersionExplicitNull (un)sets MinimumRequiredAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredAppVersion value is set to nil even if false is passed
### GetMinimumWarningAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningAppVersionOk() (string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion gets a reference to the given string and assigns it to the MinimumWarningAppVersion field.

### SetMinimumWarningAppVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningAppVersionExplicitNull(b bool)`

SetMinimumWarningAppVersionExplicitNull (un)sets MinimumWarningAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningAppVersion value is set to nil even if false is passed
### GetIsAssigned

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetIsAssignedOk() (bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAssigned

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### SetIsAssigned

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetIsAssigned(v bool)`

SetIsAssigned gets a reference to the given bool and assigns it to the IsAssigned field.

### GetAssignments

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAssignmentsOk() ([]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.

### GetScreenCaptureBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabled returns the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if non-nil, zero value otherwise.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk() (bool, bool)`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk returns a tuple with the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

HasDisableAppEncryptionIfDeviceEncryptionIsEnabled returns a boolean if a field has been set.

### SetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(v bool)`

SetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets a reference to the given bool and assigns it to the DisableAppEncryptionIfDeviceEncryptionIsEnabled field.

### GetEncryptAppData

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetEncryptAppData() bool`

GetEncryptAppData returns the EncryptAppData field if non-nil, zero value otherwise.

### GetEncryptAppDataOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetEncryptAppDataOk() (bool, bool)`

GetEncryptAppDataOk returns a tuple with the EncryptAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEncryptAppData

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasEncryptAppData() bool`

HasEncryptAppData returns a boolean if a field has been set.

### SetEncryptAppData

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetEncryptAppData(v bool)`

SetEncryptAppData gets a reference to the given bool and assigns it to the EncryptAppData field.

### GetDeployedAppCount

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeployedAppCountOk() (int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeployedAppCount

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### SetDeployedAppCount

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount gets a reference to the given int32 and assigns it to the DeployedAppCount field.

### GetMinimumRequiredPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredPatchVersion() string`

GetMinimumRequiredPatchVersion returns the MinimumRequiredPatchVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredPatchVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumRequiredPatchVersionOk() (string, bool)`

GetMinimumRequiredPatchVersionOk returns a tuple with the MinimumRequiredPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumRequiredPatchVersion() bool`

HasMinimumRequiredPatchVersion returns a boolean if a field has been set.

### SetMinimumRequiredPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredPatchVersion(v string)`

SetMinimumRequiredPatchVersion gets a reference to the given string and assigns it to the MinimumRequiredPatchVersion field.

### SetMinimumRequiredPatchVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumRequiredPatchVersionExplicitNull(b bool)`

SetMinimumRequiredPatchVersionExplicitNull (un)sets MinimumRequiredPatchVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredPatchVersion value is set to nil even if false is passed
### GetMinimumWarningPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningPatchVersion() string`

GetMinimumWarningPatchVersion returns the MinimumWarningPatchVersion field if non-nil, zero value otherwise.

### GetMinimumWarningPatchVersionOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetMinimumWarningPatchVersionOk() (string, bool)`

GetMinimumWarningPatchVersionOk returns a tuple with the MinimumWarningPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasMinimumWarningPatchVersion() bool`

HasMinimumWarningPatchVersion returns a boolean if a field has been set.

### SetMinimumWarningPatchVersion

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningPatchVersion(v string)`

SetMinimumWarningPatchVersion gets a reference to the given string and assigns it to the MinimumWarningPatchVersion field.

### SetMinimumWarningPatchVersionExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetMinimumWarningPatchVersionExplicitNull(b bool)`

SetMinimumWarningPatchVersionExplicitNull (un)sets MinimumWarningPatchVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningPatchVersion value is set to nil even if false is passed
### GetApps

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetAppsOk() ([]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApps

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetApps

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps gets a reference to the given []MicrosoftGraphManagedMobileApp and assigns it to the Apps field.

### GetDeploymentSummary

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *MicrosoftGraphAndroidManagedAppProtection) GetDeploymentSummaryOk() (AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeploymentSummary

`func (o *MicrosoftGraphAndroidManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummary

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary gets a reference to the given AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary and assigns it to the DeploymentSummary field.

### SetDeploymentSummaryExplicitNull

`func (o *MicrosoftGraphAndroidManagedAppProtection) SetDeploymentSummaryExplicitNull(b bool)`

SetDeploymentSummaryExplicitNull (un)sets DeploymentSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeploymentSummary value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



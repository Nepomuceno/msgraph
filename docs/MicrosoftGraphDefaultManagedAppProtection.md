# MicrosoftGraphDefaultManagedAppProtection

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
**AppDataEncryptionType** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataEncryptionType**](anyOf&lt;microsoft.graph.managedAppDataEncryptionType&gt;.md) | Type of encryption which should be used for data in a managed app. (iOS Only) | [optional] 
**ScreenCaptureBlocked** | Pointer to **bool** | Indicates whether screen capture is blocked. (Android only) | [optional] 
**EncryptAppData** | Pointer to **bool** | Indicates whether managed-app data should be encrypted. (Android only) | [optional] 
**DisableAppEncryptionIfDeviceEncryptionIsEnabled** | Pointer to **bool** | When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only) | [optional] 
**MinimumRequiredSdkVersion** | Pointer to **string** | Versions less than the specified version will block the managed app from accessing company data. (iOS Only) | [optional] 
**CustomSettings** | Pointer to [**[]MicrosoftGraphKeyValuePair**](microsoft.graph.keyValuePair.md) | A set of string key and string value pairs to be sent to the affected users, unalterned by this service | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**MinimumRequiredPatchVersion** | Pointer to **string** | Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only) | [optional] 
**MinimumWarningPatchVersion** | Pointer to **string** | Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only) | [optional] 
**FaceIdBlocked** | Pointer to **bool** | Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only) | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOfflineBeforeAccessCheck field.

### GetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOnlineBeforeAccessCheck field.

### GetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedInboundDataTransferSources

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedInboundDataTransferSources field.

### GetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedOutboundDataTransferDestinations field.

### GetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationalCredentialsRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### SetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired gets a reference to the given bool and assigns it to the OrganizationalCredentialsRequired field.

### GetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel gets a reference to the given AnyOfmicrosoftGraphManagedAppClipboardSharingLevel and assigns it to the AllowedOutboundClipboardSharingLevel field.

### GetDataBackupBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDataBackupBlockedOk() (bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataBackupBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### SetDataBackupBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked gets a reference to the given bool and assigns it to the DataBackupBlocked field.

### GetDeviceComplianceRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeviceComplianceRequiredOk() (bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceComplianceRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### SetDeviceComplianceRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired gets a reference to the given bool and assigns it to the DeviceComplianceRequired field.

### GetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired gets a reference to the given bool and assigns it to the ManagedBrowserToOpenLinksRequired field.

### GetSaveAsBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetSaveAsBlockedOk() (bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSaveAsBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### SetSaveAsBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked gets a reference to the given bool and assigns it to the SaveAsBlocked field.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced gets a reference to the given string and assigns it to the PeriodOfflineBeforeWipeIsEnforced field.

### GetPinRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPinRequiredOk() (bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### SetPinRequired

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired gets a reference to the given bool and assigns it to the PinRequired field.

### GetMaximumPinRetries

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMaximumPinRetriesOk() (int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaximumPinRetries

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### SetMaximumPinRetries

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries gets a reference to the given int32 and assigns it to the MaximumPinRetries field.

### GetSimplePinBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetSimplePinBlockedOk() (bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSimplePinBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.

### SetSimplePinBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked gets a reference to the given bool and assigns it to the SimplePinBlocked field.

### GetMinimumPinLength

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumPinLengthOk() (int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumPinLength

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### SetMinimumPinLength

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength gets a reference to the given int32 and assigns it to the MinimumPinLength field.

### GetPinCharacterSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPinCharacterSetOk() (AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinCharacterSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet gets a reference to the given AnyOfmicrosoftGraphManagedAppPinCharacterSet and assigns it to the PinCharacterSet field.

### GetPeriodBeforePinReset

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPeriodBeforePinResetOk() (string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodBeforePinReset

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### SetPeriodBeforePinReset

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset gets a reference to the given string and assigns it to the PeriodBeforePinReset field.

### GetAllowedDataStorageLocations

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAllowedDataStorageLocationsOk() ([]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedDataStorageLocations

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### SetAllowedDataStorageLocations

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations gets a reference to the given []AnyOfmicrosoftGraphManagedAppDataStorageLocation and assigns it to the AllowedDataStorageLocations field.

### GetContactSyncBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetContactSyncBlockedOk() (bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactSyncBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### SetContactSyncBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked gets a reference to the given bool and assigns it to the ContactSyncBlocked field.

### GetPrintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetPrintBlockedOk() (bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### SetPrintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked gets a reference to the given bool and assigns it to the PrintBlocked field.

### GetFingerprintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetFingerprintBlockedOk() (bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFingerprintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### SetFingerprintBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked gets a reference to the given bool and assigns it to the FingerprintBlocked field.

### GetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet gets a reference to the given bool and assigns it to the DisableAppPinIfDevicePinIsSet field.

### GetMinimumRequiredOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredOsVersionOk() (string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion gets a reference to the given string and assigns it to the MinimumRequiredOsVersion field.

### SetMinimumRequiredOsVersionExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredOsVersionExplicitNull(b bool)`

SetMinimumRequiredOsVersionExplicitNull (un)sets MinimumRequiredOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredOsVersion value is set to nil even if false is passed
### GetMinimumWarningOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningOsVersionOk() (string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion gets a reference to the given string and assigns it to the MinimumWarningOsVersion field.

### SetMinimumWarningOsVersionExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningOsVersionExplicitNull(b bool)`

SetMinimumWarningOsVersionExplicitNull (un)sets MinimumWarningOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningOsVersion value is set to nil even if false is passed
### GetMinimumRequiredAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredAppVersionOk() (string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion gets a reference to the given string and assigns it to the MinimumRequiredAppVersion field.

### SetMinimumRequiredAppVersionExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredAppVersionExplicitNull(b bool)`

SetMinimumRequiredAppVersionExplicitNull (un)sets MinimumRequiredAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredAppVersion value is set to nil even if false is passed
### GetMinimumWarningAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningAppVersionOk() (string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion gets a reference to the given string and assigns it to the MinimumWarningAppVersion field.

### SetMinimumWarningAppVersionExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningAppVersionExplicitNull(b bool)`

SetMinimumWarningAppVersionExplicitNull (un)sets MinimumWarningAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningAppVersion value is set to nil even if false is passed
### GetAppDataEncryptionType

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAppDataEncryptionType() AnyOfmicrosoftGraphManagedAppDataEncryptionType`

GetAppDataEncryptionType returns the AppDataEncryptionType field if non-nil, zero value otherwise.

### GetAppDataEncryptionTypeOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAppDataEncryptionTypeOk() (AnyOfmicrosoftGraphManagedAppDataEncryptionType, bool)`

GetAppDataEncryptionTypeOk returns a tuple with the AppDataEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDataEncryptionType

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasAppDataEncryptionType() bool`

HasAppDataEncryptionType returns a boolean if a field has been set.

### SetAppDataEncryptionType

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetAppDataEncryptionType(v AnyOfmicrosoftGraphManagedAppDataEncryptionType)`

SetAppDataEncryptionType gets a reference to the given AnyOfmicrosoftGraphManagedAppDataEncryptionType and assigns it to the AppDataEncryptionType field.

### GetScreenCaptureBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetEncryptAppData

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetEncryptAppData() bool`

GetEncryptAppData returns the EncryptAppData field if non-nil, zero value otherwise.

### GetEncryptAppDataOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetEncryptAppDataOk() (bool, bool)`

GetEncryptAppDataOk returns a tuple with the EncryptAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEncryptAppData

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasEncryptAppData() bool`

HasEncryptAppData returns a boolean if a field has been set.

### SetEncryptAppData

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetEncryptAppData(v bool)`

SetEncryptAppData gets a reference to the given bool and assigns it to the EncryptAppData field.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabled returns the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if non-nil, zero value otherwise.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk() (bool, bool)`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk returns a tuple with the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

HasDisableAppEncryptionIfDeviceEncryptionIsEnabled returns a boolean if a field has been set.

### SetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(v bool)`

SetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets a reference to the given bool and assigns it to the DisableAppEncryptionIfDeviceEncryptionIsEnabled field.

### GetMinimumRequiredSdkVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredSdkVersion() string`

GetMinimumRequiredSdkVersion returns the MinimumRequiredSdkVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredSdkVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredSdkVersionOk() (string, bool)`

GetMinimumRequiredSdkVersionOk returns a tuple with the MinimumRequiredSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredSdkVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumRequiredSdkVersion() bool`

HasMinimumRequiredSdkVersion returns a boolean if a field has been set.

### SetMinimumRequiredSdkVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredSdkVersion(v string)`

SetMinimumRequiredSdkVersion gets a reference to the given string and assigns it to the MinimumRequiredSdkVersion field.

### SetMinimumRequiredSdkVersionExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredSdkVersionExplicitNull(b bool)`

SetMinimumRequiredSdkVersionExplicitNull (un)sets MinimumRequiredSdkVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredSdkVersion value is set to nil even if false is passed
### GetCustomSettings

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetCustomSettings() []MicrosoftGraphKeyValuePair`

GetCustomSettings returns the CustomSettings field if non-nil, zero value otherwise.

### GetCustomSettingsOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetCustomSettingsOk() ([]MicrosoftGraphKeyValuePair, bool)`

GetCustomSettingsOk returns a tuple with the CustomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomSettings

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasCustomSettings() bool`

HasCustomSettings returns a boolean if a field has been set.

### SetCustomSettings

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetCustomSettings(v []MicrosoftGraphKeyValuePair)`

SetCustomSettings gets a reference to the given []MicrosoftGraphKeyValuePair and assigns it to the CustomSettings field.

### GetDeployedAppCount

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeployedAppCountOk() (int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeployedAppCount

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### SetDeployedAppCount

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount gets a reference to the given int32 and assigns it to the DeployedAppCount field.

### GetMinimumRequiredPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredPatchVersion() string`

GetMinimumRequiredPatchVersion returns the MinimumRequiredPatchVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredPatchVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumRequiredPatchVersionOk() (string, bool)`

GetMinimumRequiredPatchVersionOk returns a tuple with the MinimumRequiredPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumRequiredPatchVersion() bool`

HasMinimumRequiredPatchVersion returns a boolean if a field has been set.

### SetMinimumRequiredPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredPatchVersion(v string)`

SetMinimumRequiredPatchVersion gets a reference to the given string and assigns it to the MinimumRequiredPatchVersion field.

### SetMinimumRequiredPatchVersionExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumRequiredPatchVersionExplicitNull(b bool)`

SetMinimumRequiredPatchVersionExplicitNull (un)sets MinimumRequiredPatchVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredPatchVersion value is set to nil even if false is passed
### GetMinimumWarningPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningPatchVersion() string`

GetMinimumWarningPatchVersion returns the MinimumWarningPatchVersion field if non-nil, zero value otherwise.

### GetMinimumWarningPatchVersionOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetMinimumWarningPatchVersionOk() (string, bool)`

GetMinimumWarningPatchVersionOk returns a tuple with the MinimumWarningPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasMinimumWarningPatchVersion() bool`

HasMinimumWarningPatchVersion returns a boolean if a field has been set.

### SetMinimumWarningPatchVersion

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningPatchVersion(v string)`

SetMinimumWarningPatchVersion gets a reference to the given string and assigns it to the MinimumWarningPatchVersion field.

### SetMinimumWarningPatchVersionExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetMinimumWarningPatchVersionExplicitNull(b bool)`

SetMinimumWarningPatchVersionExplicitNull (un)sets MinimumWarningPatchVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningPatchVersion value is set to nil even if false is passed
### GetFaceIdBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetFaceIdBlocked() bool`

GetFaceIdBlocked returns the FaceIdBlocked field if non-nil, zero value otherwise.

### GetFaceIdBlockedOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetFaceIdBlockedOk() (bool, bool)`

GetFaceIdBlockedOk returns a tuple with the FaceIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFaceIdBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasFaceIdBlocked() bool`

HasFaceIdBlocked returns a boolean if a field has been set.

### SetFaceIdBlocked

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetFaceIdBlocked(v bool)`

SetFaceIdBlocked gets a reference to the given bool and assigns it to the FaceIdBlocked field.

### GetApps

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetAppsOk() ([]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApps

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetApps

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps gets a reference to the given []MicrosoftGraphManagedMobileApp and assigns it to the Apps field.

### GetDeploymentSummary

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *MicrosoftGraphDefaultManagedAppProtection) GetDeploymentSummaryOk() (AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeploymentSummary

`func (o *MicrosoftGraphDefaultManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummary

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary gets a reference to the given AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary and assigns it to the DeploymentSummary field.

### SetDeploymentSummaryExplicitNull

`func (o *MicrosoftGraphDefaultManagedAppProtection) SetDeploymentSummaryExplicitNull(b bool)`

SetDeploymentSummaryExplicitNull (un)sets DeploymentSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeploymentSummary value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



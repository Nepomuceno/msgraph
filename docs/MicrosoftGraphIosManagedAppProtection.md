# MicrosoftGraphIosManagedAppProtection

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
**AppDataEncryptionType** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataEncryptionType**](anyOf&lt;microsoft.graph.managedAppDataEncryptionType&gt;.md) | Type of encryption which should be used for data in a managed app. | [optional] 
**MinimumRequiredSdkVersion** | Pointer to **string** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**FaceIdBlocked** | Pointer to **bool** | Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphIosManagedAppProtection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphIosManagedAppProtection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphIosManagedAppProtection) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphIosManagedAppProtection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosManagedAppProtection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosManagedAppProtection) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphIosManagedAppProtection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphIosManagedAppProtection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphIosManagedAppProtection) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphIosManagedAppProtection) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphIosManagedAppProtection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphIosManagedAppProtection) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOfflineBeforeAccessCheck() string`

GetPeriodOfflineBeforeAccessCheck returns the PeriodOfflineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeAccessCheckOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOfflineBeforeAccessCheckOk() (string, bool)`

GetPeriodOfflineBeforeAccessCheckOk returns a tuple with the PeriodOfflineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) HasPeriodOfflineBeforeAccessCheck() bool`

HasPeriodOfflineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOfflineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(v string)`

SetPeriodOfflineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOfflineBeforeAccessCheck field.

### GetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOnlineBeforeAccessCheck() string`

GetPeriodOnlineBeforeAccessCheck returns the PeriodOnlineBeforeAccessCheck field if non-nil, zero value otherwise.

### GetPeriodOnlineBeforeAccessCheckOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOnlineBeforeAccessCheckOk() (string, bool)`

GetPeriodOnlineBeforeAccessCheckOk returns a tuple with the PeriodOnlineBeforeAccessCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) HasPeriodOnlineBeforeAccessCheck() bool`

HasPeriodOnlineBeforeAccessCheck returns a boolean if a field has been set.

### SetPeriodOnlineBeforeAccessCheck

`func (o *MicrosoftGraphIosManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(v string)`

SetPeriodOnlineBeforeAccessCheck gets a reference to the given string and assigns it to the PeriodOnlineBeforeAccessCheck field.

### GetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedInboundDataTransferSources() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedInboundDataTransferSources returns the AllowedInboundDataTransferSources field if non-nil, zero value otherwise.

### GetAllowedInboundDataTransferSourcesOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedInboundDataTransferSourcesOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedInboundDataTransferSourcesOk returns a tuple with the AllowedInboundDataTransferSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedInboundDataTransferSources

`func (o *MicrosoftGraphIosManagedAppProtection) HasAllowedInboundDataTransferSources() bool`

HasAllowedInboundDataTransferSources returns a boolean if a field has been set.

### SetAllowedInboundDataTransferSources

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedInboundDataTransferSources(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedInboundDataTransferSources gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedInboundDataTransferSources field.

### GetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedOutboundDataTransferDestinations() AnyOfmicrosoftGraphManagedAppDataTransferLevel`

GetAllowedOutboundDataTransferDestinations returns the AllowedOutboundDataTransferDestinations field if non-nil, zero value otherwise.

### GetAllowedOutboundDataTransferDestinationsOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedOutboundDataTransferDestinationsOk() (AnyOfmicrosoftGraphManagedAppDataTransferLevel, bool)`

GetAllowedOutboundDataTransferDestinationsOk returns a tuple with the AllowedOutboundDataTransferDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphIosManagedAppProtection) HasAllowedOutboundDataTransferDestinations() bool`

HasAllowedOutboundDataTransferDestinations returns a boolean if a field has been set.

### SetAllowedOutboundDataTransferDestinations

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedOutboundDataTransferDestinations(v AnyOfmicrosoftGraphManagedAppDataTransferLevel)`

SetAllowedOutboundDataTransferDestinations gets a reference to the given AnyOfmicrosoftGraphManagedAppDataTransferLevel and assigns it to the AllowedOutboundDataTransferDestinations field.

### GetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphIosManagedAppProtection) GetOrganizationalCredentialsRequired() bool`

GetOrganizationalCredentialsRequired returns the OrganizationalCredentialsRequired field if non-nil, zero value otherwise.

### GetOrganizationalCredentialsRequiredOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetOrganizationalCredentialsRequiredOk() (bool, bool)`

GetOrganizationalCredentialsRequiredOk returns a tuple with the OrganizationalCredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizationalCredentialsRequired

`func (o *MicrosoftGraphIosManagedAppProtection) HasOrganizationalCredentialsRequired() bool`

HasOrganizationalCredentialsRequired returns a boolean if a field has been set.

### SetOrganizationalCredentialsRequired

`func (o *MicrosoftGraphIosManagedAppProtection) SetOrganizationalCredentialsRequired(v bool)`

SetOrganizationalCredentialsRequired gets a reference to the given bool and assigns it to the OrganizationalCredentialsRequired field.

### GetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedOutboundClipboardSharingLevel() AnyOfmicrosoftGraphManagedAppClipboardSharingLevel`

GetAllowedOutboundClipboardSharingLevel returns the AllowedOutboundClipboardSharingLevel field if non-nil, zero value otherwise.

### GetAllowedOutboundClipboardSharingLevelOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedOutboundClipboardSharingLevelOk() (AnyOfmicrosoftGraphManagedAppClipboardSharingLevel, bool)`

GetAllowedOutboundClipboardSharingLevelOk returns a tuple with the AllowedOutboundClipboardSharingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphIosManagedAppProtection) HasAllowedOutboundClipboardSharingLevel() bool`

HasAllowedOutboundClipboardSharingLevel returns a boolean if a field has been set.

### SetAllowedOutboundClipboardSharingLevel

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(v AnyOfmicrosoftGraphManagedAppClipboardSharingLevel)`

SetAllowedOutboundClipboardSharingLevel gets a reference to the given AnyOfmicrosoftGraphManagedAppClipboardSharingLevel and assigns it to the AllowedOutboundClipboardSharingLevel field.

### GetDataBackupBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetDataBackupBlocked() bool`

GetDataBackupBlocked returns the DataBackupBlocked field if non-nil, zero value otherwise.

### GetDataBackupBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDataBackupBlockedOk() (bool, bool)`

GetDataBackupBlockedOk returns a tuple with the DataBackupBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataBackupBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasDataBackupBlocked() bool`

HasDataBackupBlocked returns a boolean if a field has been set.

### SetDataBackupBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetDataBackupBlocked(v bool)`

SetDataBackupBlocked gets a reference to the given bool and assigns it to the DataBackupBlocked field.

### GetDeviceComplianceRequired

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeviceComplianceRequired() bool`

GetDeviceComplianceRequired returns the DeviceComplianceRequired field if non-nil, zero value otherwise.

### GetDeviceComplianceRequiredOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeviceComplianceRequiredOk() (bool, bool)`

GetDeviceComplianceRequiredOk returns a tuple with the DeviceComplianceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceComplianceRequired

`func (o *MicrosoftGraphIosManagedAppProtection) HasDeviceComplianceRequired() bool`

HasDeviceComplianceRequired returns a boolean if a field has been set.

### SetDeviceComplianceRequired

`func (o *MicrosoftGraphIosManagedAppProtection) SetDeviceComplianceRequired(v bool)`

SetDeviceComplianceRequired gets a reference to the given bool and assigns it to the DeviceComplianceRequired field.

### GetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphIosManagedAppProtection) GetManagedBrowserToOpenLinksRequired() bool`

GetManagedBrowserToOpenLinksRequired returns the ManagedBrowserToOpenLinksRequired field if non-nil, zero value otherwise.

### GetManagedBrowserToOpenLinksRequiredOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetManagedBrowserToOpenLinksRequiredOk() (bool, bool)`

GetManagedBrowserToOpenLinksRequiredOk returns a tuple with the ManagedBrowserToOpenLinksRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphIosManagedAppProtection) HasManagedBrowserToOpenLinksRequired() bool`

HasManagedBrowserToOpenLinksRequired returns a boolean if a field has been set.

### SetManagedBrowserToOpenLinksRequired

`func (o *MicrosoftGraphIosManagedAppProtection) SetManagedBrowserToOpenLinksRequired(v bool)`

SetManagedBrowserToOpenLinksRequired gets a reference to the given bool and assigns it to the ManagedBrowserToOpenLinksRequired field.

### GetSaveAsBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetSaveAsBlocked() bool`

GetSaveAsBlocked returns the SaveAsBlocked field if non-nil, zero value otherwise.

### GetSaveAsBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetSaveAsBlockedOk() (bool, bool)`

GetSaveAsBlockedOk returns a tuple with the SaveAsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSaveAsBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasSaveAsBlocked() bool`

HasSaveAsBlocked returns a boolean if a field has been set.

### SetSaveAsBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetSaveAsBlocked(v bool)`

SetSaveAsBlocked gets a reference to the given bool and assigns it to the SaveAsBlocked field.

### GetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced() string`

GetPeriodOfflineBeforeWipeIsEnforced returns the PeriodOfflineBeforeWipeIsEnforced field if non-nil, zero value otherwise.

### GetPeriodOfflineBeforeWipeIsEnforcedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforcedOk() (string, bool)`

GetPeriodOfflineBeforeWipeIsEnforcedOk returns a tuple with the PeriodOfflineBeforeWipeIsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphIosManagedAppProtection) HasPeriodOfflineBeforeWipeIsEnforced() bool`

HasPeriodOfflineBeforeWipeIsEnforced returns a boolean if a field has been set.

### SetPeriodOfflineBeforeWipeIsEnforced

`func (o *MicrosoftGraphIosManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(v string)`

SetPeriodOfflineBeforeWipeIsEnforced gets a reference to the given string and assigns it to the PeriodOfflineBeforeWipeIsEnforced field.

### GetPinRequired

`func (o *MicrosoftGraphIosManagedAppProtection) GetPinRequired() bool`

GetPinRequired returns the PinRequired field if non-nil, zero value otherwise.

### GetPinRequiredOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPinRequiredOk() (bool, bool)`

GetPinRequiredOk returns a tuple with the PinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinRequired

`func (o *MicrosoftGraphIosManagedAppProtection) HasPinRequired() bool`

HasPinRequired returns a boolean if a field has been set.

### SetPinRequired

`func (o *MicrosoftGraphIosManagedAppProtection) SetPinRequired(v bool)`

SetPinRequired gets a reference to the given bool and assigns it to the PinRequired field.

### GetMaximumPinRetries

`func (o *MicrosoftGraphIosManagedAppProtection) GetMaximumPinRetries() int32`

GetMaximumPinRetries returns the MaximumPinRetries field if non-nil, zero value otherwise.

### GetMaximumPinRetriesOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMaximumPinRetriesOk() (int32, bool)`

GetMaximumPinRetriesOk returns a tuple with the MaximumPinRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaximumPinRetries

`func (o *MicrosoftGraphIosManagedAppProtection) HasMaximumPinRetries() bool`

HasMaximumPinRetries returns a boolean if a field has been set.

### SetMaximumPinRetries

`func (o *MicrosoftGraphIosManagedAppProtection) SetMaximumPinRetries(v int32)`

SetMaximumPinRetries gets a reference to the given int32 and assigns it to the MaximumPinRetries field.

### GetSimplePinBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetSimplePinBlocked() bool`

GetSimplePinBlocked returns the SimplePinBlocked field if non-nil, zero value otherwise.

### GetSimplePinBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetSimplePinBlockedOk() (bool, bool)`

GetSimplePinBlockedOk returns a tuple with the SimplePinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSimplePinBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasSimplePinBlocked() bool`

HasSimplePinBlocked returns a boolean if a field has been set.

### SetSimplePinBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetSimplePinBlocked(v bool)`

SetSimplePinBlocked gets a reference to the given bool and assigns it to the SimplePinBlocked field.

### GetMinimumPinLength

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumPinLength() int32`

GetMinimumPinLength returns the MinimumPinLength field if non-nil, zero value otherwise.

### GetMinimumPinLengthOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumPinLengthOk() (int32, bool)`

GetMinimumPinLengthOk returns a tuple with the MinimumPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumPinLength

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumPinLength() bool`

HasMinimumPinLength returns a boolean if a field has been set.

### SetMinimumPinLength

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumPinLength(v int32)`

SetMinimumPinLength gets a reference to the given int32 and assigns it to the MinimumPinLength field.

### GetPinCharacterSet

`func (o *MicrosoftGraphIosManagedAppProtection) GetPinCharacterSet() AnyOfmicrosoftGraphManagedAppPinCharacterSet`

GetPinCharacterSet returns the PinCharacterSet field if non-nil, zero value otherwise.

### GetPinCharacterSetOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPinCharacterSetOk() (AnyOfmicrosoftGraphManagedAppPinCharacterSet, bool)`

GetPinCharacterSetOk returns a tuple with the PinCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinCharacterSet

`func (o *MicrosoftGraphIosManagedAppProtection) HasPinCharacterSet() bool`

HasPinCharacterSet returns a boolean if a field has been set.

### SetPinCharacterSet

`func (o *MicrosoftGraphIosManagedAppProtection) SetPinCharacterSet(v AnyOfmicrosoftGraphManagedAppPinCharacterSet)`

SetPinCharacterSet gets a reference to the given AnyOfmicrosoftGraphManagedAppPinCharacterSet and assigns it to the PinCharacterSet field.

### GetPeriodBeforePinReset

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodBeforePinReset() string`

GetPeriodBeforePinReset returns the PeriodBeforePinReset field if non-nil, zero value otherwise.

### GetPeriodBeforePinResetOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPeriodBeforePinResetOk() (string, bool)`

GetPeriodBeforePinResetOk returns a tuple with the PeriodBeforePinReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriodBeforePinReset

`func (o *MicrosoftGraphIosManagedAppProtection) HasPeriodBeforePinReset() bool`

HasPeriodBeforePinReset returns a boolean if a field has been set.

### SetPeriodBeforePinReset

`func (o *MicrosoftGraphIosManagedAppProtection) SetPeriodBeforePinReset(v string)`

SetPeriodBeforePinReset gets a reference to the given string and assigns it to the PeriodBeforePinReset field.

### GetAllowedDataStorageLocations

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedDataStorageLocations() []AnyOfmicrosoftGraphManagedAppDataStorageLocation`

GetAllowedDataStorageLocations returns the AllowedDataStorageLocations field if non-nil, zero value otherwise.

### GetAllowedDataStorageLocationsOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAllowedDataStorageLocationsOk() ([]AnyOfmicrosoftGraphManagedAppDataStorageLocation, bool)`

GetAllowedDataStorageLocationsOk returns a tuple with the AllowedDataStorageLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedDataStorageLocations

`func (o *MicrosoftGraphIosManagedAppProtection) HasAllowedDataStorageLocations() bool`

HasAllowedDataStorageLocations returns a boolean if a field has been set.

### SetAllowedDataStorageLocations

`func (o *MicrosoftGraphIosManagedAppProtection) SetAllowedDataStorageLocations(v []AnyOfmicrosoftGraphManagedAppDataStorageLocation)`

SetAllowedDataStorageLocations gets a reference to the given []AnyOfmicrosoftGraphManagedAppDataStorageLocation and assigns it to the AllowedDataStorageLocations field.

### GetContactSyncBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetContactSyncBlocked() bool`

GetContactSyncBlocked returns the ContactSyncBlocked field if non-nil, zero value otherwise.

### GetContactSyncBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetContactSyncBlockedOk() (bool, bool)`

GetContactSyncBlockedOk returns a tuple with the ContactSyncBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContactSyncBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasContactSyncBlocked() bool`

HasContactSyncBlocked returns a boolean if a field has been set.

### SetContactSyncBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetContactSyncBlocked(v bool)`

SetContactSyncBlocked gets a reference to the given bool and assigns it to the ContactSyncBlocked field.

### GetPrintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetPrintBlocked() bool`

GetPrintBlocked returns the PrintBlocked field if non-nil, zero value otherwise.

### GetPrintBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetPrintBlockedOk() (bool, bool)`

GetPrintBlockedOk returns a tuple with the PrintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasPrintBlocked() bool`

HasPrintBlocked returns a boolean if a field has been set.

### SetPrintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetPrintBlocked(v bool)`

SetPrintBlocked gets a reference to the given bool and assigns it to the PrintBlocked field.

### GetFingerprintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetFingerprintBlocked() bool`

GetFingerprintBlocked returns the FingerprintBlocked field if non-nil, zero value otherwise.

### GetFingerprintBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetFingerprintBlockedOk() (bool, bool)`

GetFingerprintBlockedOk returns a tuple with the FingerprintBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFingerprintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasFingerprintBlocked() bool`

HasFingerprintBlocked returns a boolean if a field has been set.

### SetFingerprintBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetFingerprintBlocked(v bool)`

SetFingerprintBlocked gets a reference to the given bool and assigns it to the FingerprintBlocked field.

### GetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphIosManagedAppProtection) GetDisableAppPinIfDevicePinIsSet() bool`

GetDisableAppPinIfDevicePinIsSet returns the DisableAppPinIfDevicePinIsSet field if non-nil, zero value otherwise.

### GetDisableAppPinIfDevicePinIsSetOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDisableAppPinIfDevicePinIsSetOk() (bool, bool)`

GetDisableAppPinIfDevicePinIsSetOk returns a tuple with the DisableAppPinIfDevicePinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphIosManagedAppProtection) HasDisableAppPinIfDevicePinIsSet() bool`

HasDisableAppPinIfDevicePinIsSet returns a boolean if a field has been set.

### SetDisableAppPinIfDevicePinIsSet

`func (o *MicrosoftGraphIosManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(v bool)`

SetDisableAppPinIfDevicePinIsSet gets a reference to the given bool and assigns it to the DisableAppPinIfDevicePinIsSet field.

### GetMinimumRequiredOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredOsVersion() string`

GetMinimumRequiredOsVersion returns the MinimumRequiredOsVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredOsVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredOsVersionOk() (string, bool)`

GetMinimumRequiredOsVersionOk returns a tuple with the MinimumRequiredOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumRequiredOsVersion() bool`

HasMinimumRequiredOsVersion returns a boolean if a field has been set.

### SetMinimumRequiredOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredOsVersion(v string)`

SetMinimumRequiredOsVersion gets a reference to the given string and assigns it to the MinimumRequiredOsVersion field.

### SetMinimumRequiredOsVersionExplicitNull

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredOsVersionExplicitNull(b bool)`

SetMinimumRequiredOsVersionExplicitNull (un)sets MinimumRequiredOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredOsVersion value is set to nil even if false is passed
### GetMinimumWarningOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumWarningOsVersion() string`

GetMinimumWarningOsVersion returns the MinimumWarningOsVersion field if non-nil, zero value otherwise.

### GetMinimumWarningOsVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumWarningOsVersionOk() (string, bool)`

GetMinimumWarningOsVersionOk returns a tuple with the MinimumWarningOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumWarningOsVersion() bool`

HasMinimumWarningOsVersion returns a boolean if a field has been set.

### SetMinimumWarningOsVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumWarningOsVersion(v string)`

SetMinimumWarningOsVersion gets a reference to the given string and assigns it to the MinimumWarningOsVersion field.

### SetMinimumWarningOsVersionExplicitNull

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumWarningOsVersionExplicitNull(b bool)`

SetMinimumWarningOsVersionExplicitNull (un)sets MinimumWarningOsVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningOsVersion value is set to nil even if false is passed
### GetMinimumRequiredAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredAppVersion() string`

GetMinimumRequiredAppVersion returns the MinimumRequiredAppVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredAppVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredAppVersionOk() (string, bool)`

GetMinimumRequiredAppVersionOk returns a tuple with the MinimumRequiredAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumRequiredAppVersion() bool`

HasMinimumRequiredAppVersion returns a boolean if a field has been set.

### SetMinimumRequiredAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredAppVersion(v string)`

SetMinimumRequiredAppVersion gets a reference to the given string and assigns it to the MinimumRequiredAppVersion field.

### SetMinimumRequiredAppVersionExplicitNull

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredAppVersionExplicitNull(b bool)`

SetMinimumRequiredAppVersionExplicitNull (un)sets MinimumRequiredAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredAppVersion value is set to nil even if false is passed
### GetMinimumWarningAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumWarningAppVersion() string`

GetMinimumWarningAppVersion returns the MinimumWarningAppVersion field if non-nil, zero value otherwise.

### GetMinimumWarningAppVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumWarningAppVersionOk() (string, bool)`

GetMinimumWarningAppVersionOk returns a tuple with the MinimumWarningAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumWarningAppVersion() bool`

HasMinimumWarningAppVersion returns a boolean if a field has been set.

### SetMinimumWarningAppVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumWarningAppVersion(v string)`

SetMinimumWarningAppVersion gets a reference to the given string and assigns it to the MinimumWarningAppVersion field.

### SetMinimumWarningAppVersionExplicitNull

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumWarningAppVersionExplicitNull(b bool)`

SetMinimumWarningAppVersionExplicitNull (un)sets MinimumWarningAppVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningAppVersion value is set to nil even if false is passed
### GetIsAssigned

`func (o *MicrosoftGraphIosManagedAppProtection) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetIsAssignedOk() (bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAssigned

`func (o *MicrosoftGraphIosManagedAppProtection) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### SetIsAssigned

`func (o *MicrosoftGraphIosManagedAppProtection) SetIsAssigned(v bool)`

SetIsAssigned gets a reference to the given bool and assigns it to the IsAssigned field.

### GetAssignments

`func (o *MicrosoftGraphIosManagedAppProtection) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAssignmentsOk() ([]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphIosManagedAppProtection) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphIosManagedAppProtection) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.

### GetAppDataEncryptionType

`func (o *MicrosoftGraphIosManagedAppProtection) GetAppDataEncryptionType() AnyOfmicrosoftGraphManagedAppDataEncryptionType`

GetAppDataEncryptionType returns the AppDataEncryptionType field if non-nil, zero value otherwise.

### GetAppDataEncryptionTypeOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAppDataEncryptionTypeOk() (AnyOfmicrosoftGraphManagedAppDataEncryptionType, bool)`

GetAppDataEncryptionTypeOk returns a tuple with the AppDataEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDataEncryptionType

`func (o *MicrosoftGraphIosManagedAppProtection) HasAppDataEncryptionType() bool`

HasAppDataEncryptionType returns a boolean if a field has been set.

### SetAppDataEncryptionType

`func (o *MicrosoftGraphIosManagedAppProtection) SetAppDataEncryptionType(v AnyOfmicrosoftGraphManagedAppDataEncryptionType)`

SetAppDataEncryptionType gets a reference to the given AnyOfmicrosoftGraphManagedAppDataEncryptionType and assigns it to the AppDataEncryptionType field.

### GetMinimumRequiredSdkVersion

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredSdkVersion() string`

GetMinimumRequiredSdkVersion returns the MinimumRequiredSdkVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredSdkVersionOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetMinimumRequiredSdkVersionOk() (string, bool)`

GetMinimumRequiredSdkVersionOk returns a tuple with the MinimumRequiredSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredSdkVersion

`func (o *MicrosoftGraphIosManagedAppProtection) HasMinimumRequiredSdkVersion() bool`

HasMinimumRequiredSdkVersion returns a boolean if a field has been set.

### SetMinimumRequiredSdkVersion

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredSdkVersion(v string)`

SetMinimumRequiredSdkVersion gets a reference to the given string and assigns it to the MinimumRequiredSdkVersion field.

### SetMinimumRequiredSdkVersionExplicitNull

`func (o *MicrosoftGraphIosManagedAppProtection) SetMinimumRequiredSdkVersionExplicitNull(b bool)`

SetMinimumRequiredSdkVersionExplicitNull (un)sets MinimumRequiredSdkVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredSdkVersion value is set to nil even if false is passed
### GetDeployedAppCount

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeployedAppCountOk() (int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeployedAppCount

`func (o *MicrosoftGraphIosManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### SetDeployedAppCount

`func (o *MicrosoftGraphIosManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount gets a reference to the given int32 and assigns it to the DeployedAppCount field.

### GetFaceIdBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) GetFaceIdBlocked() bool`

GetFaceIdBlocked returns the FaceIdBlocked field if non-nil, zero value otherwise.

### GetFaceIdBlockedOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetFaceIdBlockedOk() (bool, bool)`

GetFaceIdBlockedOk returns a tuple with the FaceIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFaceIdBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) HasFaceIdBlocked() bool`

HasFaceIdBlocked returns a boolean if a field has been set.

### SetFaceIdBlocked

`func (o *MicrosoftGraphIosManagedAppProtection) SetFaceIdBlocked(v bool)`

SetFaceIdBlocked gets a reference to the given bool and assigns it to the FaceIdBlocked field.

### GetApps

`func (o *MicrosoftGraphIosManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetAppsOk() ([]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApps

`func (o *MicrosoftGraphIosManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetApps

`func (o *MicrosoftGraphIosManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps gets a reference to the given []MicrosoftGraphManagedMobileApp and assigns it to the Apps field.

### GetDeploymentSummary

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *MicrosoftGraphIosManagedAppProtection) GetDeploymentSummaryOk() (AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeploymentSummary

`func (o *MicrosoftGraphIosManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummary

`func (o *MicrosoftGraphIosManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary gets a reference to the given AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary and assigns it to the DeploymentSummary field.

### SetDeploymentSummaryExplicitNull

`func (o *MicrosoftGraphIosManagedAppProtection) SetDeploymentSummaryExplicitNull(b bool)`

SetDeploymentSummaryExplicitNull (un)sets DeploymentSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeploymentSummary value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



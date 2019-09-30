# MicrosoftGraphWindows10EndpointProtectionConfiguration

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
**FirewallBlockStatefulFTP** | Pointer to **bool** | Blocks stateful FTP connections to the device | [optional] 
**FirewallIdleTimeoutForSecurityAssociationInSeconds** | Pointer to **int32** | Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600 | [optional] 
**FirewallPreSharedKeyEncodingMethod** | Pointer to [**AnyOfmicrosoftGraphFirewallPreSharedKeyEncodingMethodType**](anyOf&lt;microsoft.graph.firewallPreSharedKeyEncodingMethodType&gt;.md) | Select the preshared key encoding to be used | [optional] 
**FirewallIPSecExemptionsAllowNeighborDiscovery** | Pointer to **bool** | Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes | [optional] 
**FirewallIPSecExemptionsAllowICMP** | Pointer to **bool** | Configures IPSec exemptions to allow ICMP | [optional] 
**FirewallIPSecExemptionsAllowRouterDiscovery** | Pointer to **bool** | Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes | [optional] 
**FirewallIPSecExemptionsAllowDHCP** | Pointer to **bool** | Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic | [optional] 
**FirewallCertificateRevocationListCheckMethod** | Pointer to [**AnyOfmicrosoftGraphFirewallCertificateRevocationListCheckMethodType**](anyOf&lt;microsoft.graph.firewallCertificateRevocationListCheckMethodType&gt;.md) | Specify how the certificate revocation list is to be enforced | [optional] 
**FirewallMergeKeyingModuleSettings** | Pointer to **bool** | If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set | [optional] 
**FirewallPacketQueueingMethod** | Pointer to [**AnyOfmicrosoftGraphFirewallPacketQueueingMethodType**](anyOf&lt;microsoft.graph.firewallPacketQueueingMethodType&gt;.md) | Configures how packet queueing should be applied in the tunnel gateway scenario | [optional] 
**FirewallProfileDomain** | Pointer to [**AnyOfmicrosoftGraphWindowsFirewallNetworkProfile**](anyOf&lt;microsoft.graph.windowsFirewallNetworkProfile&gt;.md) | Configures the firewall profile settings for domain networks | [optional] 
**FirewallProfilePublic** | Pointer to [**AnyOfmicrosoftGraphWindowsFirewallNetworkProfile**](anyOf&lt;microsoft.graph.windowsFirewallNetworkProfile&gt;.md) | Configures the firewall profile settings for public networks | [optional] 
**FirewallProfilePrivate** | Pointer to [**AnyOfmicrosoftGraphWindowsFirewallNetworkProfile**](anyOf&lt;microsoft.graph.windowsFirewallNetworkProfile&gt;.md) | Configures the firewall profile settings for private networks | [optional] 
**DefenderAttackSurfaceReductionExcludedPaths** | Pointer to **[]string** | List of exe files and folders to be excluded from attack surface reduction rules | [optional] 
**DefenderGuardedFoldersAllowedAppPaths** | Pointer to **[]string** | List of paths to exe that are allowed to access protected folders | [optional] 
**DefenderAdditionalGuardedFolders** | Pointer to **[]string** | List of folder paths to be added to the list of protected folders | [optional] 
**DefenderExploitProtectionXml** | Pointer to **string** | Xml content containing information regarding exploit protection details. | [optional] 
**DefenderExploitProtectionXmlFileName** | Pointer to **string** | Name of the file from which DefenderExploitProtectionXml was obtained. | [optional] 
**DefenderSecurityCenterBlockExploitProtectionOverride** | Pointer to **bool** | Indicates whether or not to block user from overriding Exploit Protection settings. | [optional] 
**AppLockerApplicationControl** | Pointer to [**AnyOfmicrosoftGraphAppLockerApplicationControlType**](anyOf&lt;microsoft.graph.appLockerApplicationControlType&gt;.md) | Enables the Admin to choose what types of app to allow on devices. | [optional] 
**SmartScreenEnableInShell** | Pointer to **bool** | Allows IT Admins to configure SmartScreen for Windows. | [optional] 
**SmartScreenBlockOverrideForFiles** | Pointer to **bool** | Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files. | [optional] 
**ApplicationGuardEnabled** | Pointer to **bool** | Enable Windows Defender Application Guard | [optional] 
**ApplicationGuardBlockFileTransfer** | Pointer to [**AnyOfmicrosoftGraphApplicationGuardBlockFileTransferType**](anyOf&lt;microsoft.graph.applicationGuardBlockFileTransferType&gt;.md) | Block clipboard to transfer image file, text file or neither of them | [optional] 
**ApplicationGuardBlockNonEnterpriseContent** | Pointer to **bool** | Block enterprise sites to load non-enterprise content, such as third party plug-ins | [optional] 
**ApplicationGuardAllowPersistence** | Pointer to **bool** | Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.) | [optional] 
**ApplicationGuardForceAuditing** | Pointer to **bool** | Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.) | [optional] 
**ApplicationGuardBlockClipboardSharing** | Pointer to [**AnyOfmicrosoftGraphApplicationGuardBlockClipboardSharingType**](anyOf&lt;microsoft.graph.applicationGuardBlockClipboardSharingType&gt;.md) | Block clipboard to share data from Host to Container, or from Container to Host, or both ways, or neither ways. | [optional] 
**ApplicationGuardAllowPrintToPDF** | Pointer to **bool** | Allow printing to PDF from Container | [optional] 
**ApplicationGuardAllowPrintToXPS** | Pointer to **bool** | Allow printing to XPS from Container | [optional] 
**ApplicationGuardAllowPrintToLocalPrinters** | Pointer to **bool** | Allow printing to Local Printers from Container | [optional] 
**ApplicationGuardAllowPrintToNetworkPrinters** | Pointer to **bool** | Allow printing to Network Printers from Container | [optional] 
**BitLockerDisableWarningForOtherDiskEncryption** | Pointer to **bool** | Allows the Admin to disable the warning prompt for other disk encryption on the user machines. | [optional] 
**BitLockerEnableStorageCardEncryptionOnMobile** | Pointer to **bool** | Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU. | [optional] 
**BitLockerEncryptDevice** | Pointer to **bool** | Allows the admin to require encryption to be turned on using BitLocker. | [optional] 
**BitLockerRemovableDrivePolicy** | Pointer to [**AnyOfmicrosoftGraphBitLockerRemovableDrivePolicy**](anyOf&lt;microsoft.graph.bitLockerRemovableDrivePolicy&gt;.md) | BitLocker Removable Drive Policy. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetFirewallBlockStatefulFTP

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallBlockStatefulFTP() bool`

GetFirewallBlockStatefulFTP returns the FirewallBlockStatefulFTP field if non-nil, zero value otherwise.

### GetFirewallBlockStatefulFTPOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallBlockStatefulFTPOk() (bool, bool)`

GetFirewallBlockStatefulFTPOk returns a tuple with the FirewallBlockStatefulFTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallBlockStatefulFTP

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallBlockStatefulFTP() bool`

HasFirewallBlockStatefulFTP returns a boolean if a field has been set.

### SetFirewallBlockStatefulFTP

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallBlockStatefulFTP(v bool)`

SetFirewallBlockStatefulFTP gets a reference to the given bool and assigns it to the FirewallBlockStatefulFTP field.

### GetFirewallIdleTimeoutForSecurityAssociationInSeconds

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIdleTimeoutForSecurityAssociationInSeconds() int32`

GetFirewallIdleTimeoutForSecurityAssociationInSeconds returns the FirewallIdleTimeoutForSecurityAssociationInSeconds field if non-nil, zero value otherwise.

### GetFirewallIdleTimeoutForSecurityAssociationInSecondsOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIdleTimeoutForSecurityAssociationInSecondsOk() (int32, bool)`

GetFirewallIdleTimeoutForSecurityAssociationInSecondsOk returns a tuple with the FirewallIdleTimeoutForSecurityAssociationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallIdleTimeoutForSecurityAssociationInSeconds

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallIdleTimeoutForSecurityAssociationInSeconds() bool`

HasFirewallIdleTimeoutForSecurityAssociationInSeconds returns a boolean if a field has been set.

### SetFirewallIdleTimeoutForSecurityAssociationInSeconds

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallIdleTimeoutForSecurityAssociationInSeconds(v int32)`

SetFirewallIdleTimeoutForSecurityAssociationInSeconds gets a reference to the given int32 and assigns it to the FirewallIdleTimeoutForSecurityAssociationInSeconds field.

### SetFirewallIdleTimeoutForSecurityAssociationInSecondsExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallIdleTimeoutForSecurityAssociationInSecondsExplicitNull(b bool)`

SetFirewallIdleTimeoutForSecurityAssociationInSecondsExplicitNull (un)sets FirewallIdleTimeoutForSecurityAssociationInSeconds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FirewallIdleTimeoutForSecurityAssociationInSeconds value is set to nil even if false is passed
### GetFirewallPreSharedKeyEncodingMethod

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallPreSharedKeyEncodingMethod() AnyOfmicrosoftGraphFirewallPreSharedKeyEncodingMethodType`

GetFirewallPreSharedKeyEncodingMethod returns the FirewallPreSharedKeyEncodingMethod field if non-nil, zero value otherwise.

### GetFirewallPreSharedKeyEncodingMethodOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallPreSharedKeyEncodingMethodOk() (AnyOfmicrosoftGraphFirewallPreSharedKeyEncodingMethodType, bool)`

GetFirewallPreSharedKeyEncodingMethodOk returns a tuple with the FirewallPreSharedKeyEncodingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallPreSharedKeyEncodingMethod

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallPreSharedKeyEncodingMethod() bool`

HasFirewallPreSharedKeyEncodingMethod returns a boolean if a field has been set.

### SetFirewallPreSharedKeyEncodingMethod

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallPreSharedKeyEncodingMethod(v AnyOfmicrosoftGraphFirewallPreSharedKeyEncodingMethodType)`

SetFirewallPreSharedKeyEncodingMethod gets a reference to the given AnyOfmicrosoftGraphFirewallPreSharedKeyEncodingMethodType and assigns it to the FirewallPreSharedKeyEncodingMethod field.

### GetFirewallIPSecExemptionsAllowNeighborDiscovery

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowNeighborDiscovery() bool`

GetFirewallIPSecExemptionsAllowNeighborDiscovery returns the FirewallIPSecExemptionsAllowNeighborDiscovery field if non-nil, zero value otherwise.

### GetFirewallIPSecExemptionsAllowNeighborDiscoveryOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowNeighborDiscoveryOk() (bool, bool)`

GetFirewallIPSecExemptionsAllowNeighborDiscoveryOk returns a tuple with the FirewallIPSecExemptionsAllowNeighborDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallIPSecExemptionsAllowNeighborDiscovery

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallIPSecExemptionsAllowNeighborDiscovery() bool`

HasFirewallIPSecExemptionsAllowNeighborDiscovery returns a boolean if a field has been set.

### SetFirewallIPSecExemptionsAllowNeighborDiscovery

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowNeighborDiscovery(v bool)`

SetFirewallIPSecExemptionsAllowNeighborDiscovery gets a reference to the given bool and assigns it to the FirewallIPSecExemptionsAllowNeighborDiscovery field.

### GetFirewallIPSecExemptionsAllowICMP

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowICMP() bool`

GetFirewallIPSecExemptionsAllowICMP returns the FirewallIPSecExemptionsAllowICMP field if non-nil, zero value otherwise.

### GetFirewallIPSecExemptionsAllowICMPOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowICMPOk() (bool, bool)`

GetFirewallIPSecExemptionsAllowICMPOk returns a tuple with the FirewallIPSecExemptionsAllowICMP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallIPSecExemptionsAllowICMP

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallIPSecExemptionsAllowICMP() bool`

HasFirewallIPSecExemptionsAllowICMP returns a boolean if a field has been set.

### SetFirewallIPSecExemptionsAllowICMP

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowICMP(v bool)`

SetFirewallIPSecExemptionsAllowICMP gets a reference to the given bool and assigns it to the FirewallIPSecExemptionsAllowICMP field.

### GetFirewallIPSecExemptionsAllowRouterDiscovery

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowRouterDiscovery() bool`

GetFirewallIPSecExemptionsAllowRouterDiscovery returns the FirewallIPSecExemptionsAllowRouterDiscovery field if non-nil, zero value otherwise.

### GetFirewallIPSecExemptionsAllowRouterDiscoveryOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowRouterDiscoveryOk() (bool, bool)`

GetFirewallIPSecExemptionsAllowRouterDiscoveryOk returns a tuple with the FirewallIPSecExemptionsAllowRouterDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallIPSecExemptionsAllowRouterDiscovery

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallIPSecExemptionsAllowRouterDiscovery() bool`

HasFirewallIPSecExemptionsAllowRouterDiscovery returns a boolean if a field has been set.

### SetFirewallIPSecExemptionsAllowRouterDiscovery

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowRouterDiscovery(v bool)`

SetFirewallIPSecExemptionsAllowRouterDiscovery gets a reference to the given bool and assigns it to the FirewallIPSecExemptionsAllowRouterDiscovery field.

### GetFirewallIPSecExemptionsAllowDHCP

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowDHCP() bool`

GetFirewallIPSecExemptionsAllowDHCP returns the FirewallIPSecExemptionsAllowDHCP field if non-nil, zero value otherwise.

### GetFirewallIPSecExemptionsAllowDHCPOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowDHCPOk() (bool, bool)`

GetFirewallIPSecExemptionsAllowDHCPOk returns a tuple with the FirewallIPSecExemptionsAllowDHCP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallIPSecExemptionsAllowDHCP

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallIPSecExemptionsAllowDHCP() bool`

HasFirewallIPSecExemptionsAllowDHCP returns a boolean if a field has been set.

### SetFirewallIPSecExemptionsAllowDHCP

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowDHCP(v bool)`

SetFirewallIPSecExemptionsAllowDHCP gets a reference to the given bool and assigns it to the FirewallIPSecExemptionsAllowDHCP field.

### GetFirewallCertificateRevocationListCheckMethod

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallCertificateRevocationListCheckMethod() AnyOfmicrosoftGraphFirewallCertificateRevocationListCheckMethodType`

GetFirewallCertificateRevocationListCheckMethod returns the FirewallCertificateRevocationListCheckMethod field if non-nil, zero value otherwise.

### GetFirewallCertificateRevocationListCheckMethodOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallCertificateRevocationListCheckMethodOk() (AnyOfmicrosoftGraphFirewallCertificateRevocationListCheckMethodType, bool)`

GetFirewallCertificateRevocationListCheckMethodOk returns a tuple with the FirewallCertificateRevocationListCheckMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallCertificateRevocationListCheckMethod

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallCertificateRevocationListCheckMethod() bool`

HasFirewallCertificateRevocationListCheckMethod returns a boolean if a field has been set.

### SetFirewallCertificateRevocationListCheckMethod

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallCertificateRevocationListCheckMethod(v AnyOfmicrosoftGraphFirewallCertificateRevocationListCheckMethodType)`

SetFirewallCertificateRevocationListCheckMethod gets a reference to the given AnyOfmicrosoftGraphFirewallCertificateRevocationListCheckMethodType and assigns it to the FirewallCertificateRevocationListCheckMethod field.

### GetFirewallMergeKeyingModuleSettings

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallMergeKeyingModuleSettings() bool`

GetFirewallMergeKeyingModuleSettings returns the FirewallMergeKeyingModuleSettings field if non-nil, zero value otherwise.

### GetFirewallMergeKeyingModuleSettingsOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallMergeKeyingModuleSettingsOk() (bool, bool)`

GetFirewallMergeKeyingModuleSettingsOk returns a tuple with the FirewallMergeKeyingModuleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallMergeKeyingModuleSettings

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallMergeKeyingModuleSettings() bool`

HasFirewallMergeKeyingModuleSettings returns a boolean if a field has been set.

### SetFirewallMergeKeyingModuleSettings

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallMergeKeyingModuleSettings(v bool)`

SetFirewallMergeKeyingModuleSettings gets a reference to the given bool and assigns it to the FirewallMergeKeyingModuleSettings field.

### GetFirewallPacketQueueingMethod

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallPacketQueueingMethod() AnyOfmicrosoftGraphFirewallPacketQueueingMethodType`

GetFirewallPacketQueueingMethod returns the FirewallPacketQueueingMethod field if non-nil, zero value otherwise.

### GetFirewallPacketQueueingMethodOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallPacketQueueingMethodOk() (AnyOfmicrosoftGraphFirewallPacketQueueingMethodType, bool)`

GetFirewallPacketQueueingMethodOk returns a tuple with the FirewallPacketQueueingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallPacketQueueingMethod

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallPacketQueueingMethod() bool`

HasFirewallPacketQueueingMethod returns a boolean if a field has been set.

### SetFirewallPacketQueueingMethod

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallPacketQueueingMethod(v AnyOfmicrosoftGraphFirewallPacketQueueingMethodType)`

SetFirewallPacketQueueingMethod gets a reference to the given AnyOfmicrosoftGraphFirewallPacketQueueingMethodType and assigns it to the FirewallPacketQueueingMethod field.

### GetFirewallProfileDomain

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallProfileDomain() AnyOfmicrosoftGraphWindowsFirewallNetworkProfile`

GetFirewallProfileDomain returns the FirewallProfileDomain field if non-nil, zero value otherwise.

### GetFirewallProfileDomainOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallProfileDomainOk() (AnyOfmicrosoftGraphWindowsFirewallNetworkProfile, bool)`

GetFirewallProfileDomainOk returns a tuple with the FirewallProfileDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallProfileDomain

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallProfileDomain() bool`

HasFirewallProfileDomain returns a boolean if a field has been set.

### SetFirewallProfileDomain

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallProfileDomain(v AnyOfmicrosoftGraphWindowsFirewallNetworkProfile)`

SetFirewallProfileDomain gets a reference to the given AnyOfmicrosoftGraphWindowsFirewallNetworkProfile and assigns it to the FirewallProfileDomain field.

### SetFirewallProfileDomainExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallProfileDomainExplicitNull(b bool)`

SetFirewallProfileDomainExplicitNull (un)sets FirewallProfileDomain to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FirewallProfileDomain value is set to nil even if false is passed
### GetFirewallProfilePublic

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallProfilePublic() AnyOfmicrosoftGraphWindowsFirewallNetworkProfile`

GetFirewallProfilePublic returns the FirewallProfilePublic field if non-nil, zero value otherwise.

### GetFirewallProfilePublicOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallProfilePublicOk() (AnyOfmicrosoftGraphWindowsFirewallNetworkProfile, bool)`

GetFirewallProfilePublicOk returns a tuple with the FirewallProfilePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallProfilePublic

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallProfilePublic() bool`

HasFirewallProfilePublic returns a boolean if a field has been set.

### SetFirewallProfilePublic

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallProfilePublic(v AnyOfmicrosoftGraphWindowsFirewallNetworkProfile)`

SetFirewallProfilePublic gets a reference to the given AnyOfmicrosoftGraphWindowsFirewallNetworkProfile and assigns it to the FirewallProfilePublic field.

### SetFirewallProfilePublicExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallProfilePublicExplicitNull(b bool)`

SetFirewallProfilePublicExplicitNull (un)sets FirewallProfilePublic to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FirewallProfilePublic value is set to nil even if false is passed
### GetFirewallProfilePrivate

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallProfilePrivate() AnyOfmicrosoftGraphWindowsFirewallNetworkProfile`

GetFirewallProfilePrivate returns the FirewallProfilePrivate field if non-nil, zero value otherwise.

### GetFirewallProfilePrivateOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetFirewallProfilePrivateOk() (AnyOfmicrosoftGraphWindowsFirewallNetworkProfile, bool)`

GetFirewallProfilePrivateOk returns a tuple with the FirewallProfilePrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirewallProfilePrivate

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasFirewallProfilePrivate() bool`

HasFirewallProfilePrivate returns a boolean if a field has been set.

### SetFirewallProfilePrivate

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallProfilePrivate(v AnyOfmicrosoftGraphWindowsFirewallNetworkProfile)`

SetFirewallProfilePrivate gets a reference to the given AnyOfmicrosoftGraphWindowsFirewallNetworkProfile and assigns it to the FirewallProfilePrivate field.

### SetFirewallProfilePrivateExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetFirewallProfilePrivateExplicitNull(b bool)`

SetFirewallProfilePrivateExplicitNull (un)sets FirewallProfilePrivate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FirewallProfilePrivate value is set to nil even if false is passed
### GetDefenderAttackSurfaceReductionExcludedPaths

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderAttackSurfaceReductionExcludedPaths() []string`

GetDefenderAttackSurfaceReductionExcludedPaths returns the DefenderAttackSurfaceReductionExcludedPaths field if non-nil, zero value otherwise.

### GetDefenderAttackSurfaceReductionExcludedPathsOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderAttackSurfaceReductionExcludedPathsOk() ([]string, bool)`

GetDefenderAttackSurfaceReductionExcludedPathsOk returns a tuple with the DefenderAttackSurfaceReductionExcludedPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderAttackSurfaceReductionExcludedPaths

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDefenderAttackSurfaceReductionExcludedPaths() bool`

HasDefenderAttackSurfaceReductionExcludedPaths returns a boolean if a field has been set.

### SetDefenderAttackSurfaceReductionExcludedPaths

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDefenderAttackSurfaceReductionExcludedPaths(v []string)`

SetDefenderAttackSurfaceReductionExcludedPaths gets a reference to the given []string and assigns it to the DefenderAttackSurfaceReductionExcludedPaths field.

### GetDefenderGuardedFoldersAllowedAppPaths

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderGuardedFoldersAllowedAppPaths() []string`

GetDefenderGuardedFoldersAllowedAppPaths returns the DefenderGuardedFoldersAllowedAppPaths field if non-nil, zero value otherwise.

### GetDefenderGuardedFoldersAllowedAppPathsOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderGuardedFoldersAllowedAppPathsOk() ([]string, bool)`

GetDefenderGuardedFoldersAllowedAppPathsOk returns a tuple with the DefenderGuardedFoldersAllowedAppPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderGuardedFoldersAllowedAppPaths

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDefenderGuardedFoldersAllowedAppPaths() bool`

HasDefenderGuardedFoldersAllowedAppPaths returns a boolean if a field has been set.

### SetDefenderGuardedFoldersAllowedAppPaths

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDefenderGuardedFoldersAllowedAppPaths(v []string)`

SetDefenderGuardedFoldersAllowedAppPaths gets a reference to the given []string and assigns it to the DefenderGuardedFoldersAllowedAppPaths field.

### GetDefenderAdditionalGuardedFolders

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderAdditionalGuardedFolders() []string`

GetDefenderAdditionalGuardedFolders returns the DefenderAdditionalGuardedFolders field if non-nil, zero value otherwise.

### GetDefenderAdditionalGuardedFoldersOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderAdditionalGuardedFoldersOk() ([]string, bool)`

GetDefenderAdditionalGuardedFoldersOk returns a tuple with the DefenderAdditionalGuardedFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderAdditionalGuardedFolders

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDefenderAdditionalGuardedFolders() bool`

HasDefenderAdditionalGuardedFolders returns a boolean if a field has been set.

### SetDefenderAdditionalGuardedFolders

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDefenderAdditionalGuardedFolders(v []string)`

SetDefenderAdditionalGuardedFolders gets a reference to the given []string and assigns it to the DefenderAdditionalGuardedFolders field.

### GetDefenderExploitProtectionXml

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXml() string`

GetDefenderExploitProtectionXml returns the DefenderExploitProtectionXml field if non-nil, zero value otherwise.

### GetDefenderExploitProtectionXmlOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXmlOk() (string, bool)`

GetDefenderExploitProtectionXmlOk returns a tuple with the DefenderExploitProtectionXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderExploitProtectionXml

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDefenderExploitProtectionXml() bool`

HasDefenderExploitProtectionXml returns a boolean if a field has been set.

### SetDefenderExploitProtectionXml

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXml(v string)`

SetDefenderExploitProtectionXml gets a reference to the given string and assigns it to the DefenderExploitProtectionXml field.

### SetDefenderExploitProtectionXmlExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXmlExplicitNull(b bool)`

SetDefenderExploitProtectionXmlExplicitNull (un)sets DefenderExploitProtectionXml to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderExploitProtectionXml value is set to nil even if false is passed
### GetDefenderExploitProtectionXmlFileName

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXmlFileName() string`

GetDefenderExploitProtectionXmlFileName returns the DefenderExploitProtectionXmlFileName field if non-nil, zero value otherwise.

### GetDefenderExploitProtectionXmlFileNameOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXmlFileNameOk() (string, bool)`

GetDefenderExploitProtectionXmlFileNameOk returns a tuple with the DefenderExploitProtectionXmlFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderExploitProtectionXmlFileName

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDefenderExploitProtectionXmlFileName() bool`

HasDefenderExploitProtectionXmlFileName returns a boolean if a field has been set.

### SetDefenderExploitProtectionXmlFileName

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXmlFileName(v string)`

SetDefenderExploitProtectionXmlFileName gets a reference to the given string and assigns it to the DefenderExploitProtectionXmlFileName field.

### SetDefenderExploitProtectionXmlFileNameExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXmlFileNameExplicitNull(b bool)`

SetDefenderExploitProtectionXmlFileNameExplicitNull (un)sets DefenderExploitProtectionXmlFileName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DefenderExploitProtectionXmlFileName value is set to nil even if false is passed
### GetDefenderSecurityCenterBlockExploitProtectionOverride

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderSecurityCenterBlockExploitProtectionOverride() bool`

GetDefenderSecurityCenterBlockExploitProtectionOverride returns the DefenderSecurityCenterBlockExploitProtectionOverride field if non-nil, zero value otherwise.

### GetDefenderSecurityCenterBlockExploitProtectionOverrideOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetDefenderSecurityCenterBlockExploitProtectionOverrideOk() (bool, bool)`

GetDefenderSecurityCenterBlockExploitProtectionOverrideOk returns a tuple with the DefenderSecurityCenterBlockExploitProtectionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefenderSecurityCenterBlockExploitProtectionOverride

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasDefenderSecurityCenterBlockExploitProtectionOverride() bool`

HasDefenderSecurityCenterBlockExploitProtectionOverride returns a boolean if a field has been set.

### SetDefenderSecurityCenterBlockExploitProtectionOverride

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetDefenderSecurityCenterBlockExploitProtectionOverride(v bool)`

SetDefenderSecurityCenterBlockExploitProtectionOverride gets a reference to the given bool and assigns it to the DefenderSecurityCenterBlockExploitProtectionOverride field.

### GetAppLockerApplicationControl

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetAppLockerApplicationControl() AnyOfmicrosoftGraphAppLockerApplicationControlType`

GetAppLockerApplicationControl returns the AppLockerApplicationControl field if non-nil, zero value otherwise.

### GetAppLockerApplicationControlOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetAppLockerApplicationControlOk() (AnyOfmicrosoftGraphAppLockerApplicationControlType, bool)`

GetAppLockerApplicationControlOk returns a tuple with the AppLockerApplicationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppLockerApplicationControl

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasAppLockerApplicationControl() bool`

HasAppLockerApplicationControl returns a boolean if a field has been set.

### SetAppLockerApplicationControl

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetAppLockerApplicationControl(v AnyOfmicrosoftGraphAppLockerApplicationControlType)`

SetAppLockerApplicationControl gets a reference to the given AnyOfmicrosoftGraphAppLockerApplicationControlType and assigns it to the AppLockerApplicationControl field.

### GetSmartScreenEnableInShell

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetSmartScreenEnableInShell() bool`

GetSmartScreenEnableInShell returns the SmartScreenEnableInShell field if non-nil, zero value otherwise.

### GetSmartScreenEnableInShellOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetSmartScreenEnableInShellOk() (bool, bool)`

GetSmartScreenEnableInShellOk returns a tuple with the SmartScreenEnableInShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmartScreenEnableInShell

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasSmartScreenEnableInShell() bool`

HasSmartScreenEnableInShell returns a boolean if a field has been set.

### SetSmartScreenEnableInShell

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetSmartScreenEnableInShell(v bool)`

SetSmartScreenEnableInShell gets a reference to the given bool and assigns it to the SmartScreenEnableInShell field.

### GetSmartScreenBlockOverrideForFiles

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetSmartScreenBlockOverrideForFiles() bool`

GetSmartScreenBlockOverrideForFiles returns the SmartScreenBlockOverrideForFiles field if non-nil, zero value otherwise.

### GetSmartScreenBlockOverrideForFilesOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetSmartScreenBlockOverrideForFilesOk() (bool, bool)`

GetSmartScreenBlockOverrideForFilesOk returns a tuple with the SmartScreenBlockOverrideForFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmartScreenBlockOverrideForFiles

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasSmartScreenBlockOverrideForFiles() bool`

HasSmartScreenBlockOverrideForFiles returns a boolean if a field has been set.

### SetSmartScreenBlockOverrideForFiles

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetSmartScreenBlockOverrideForFiles(v bool)`

SetSmartScreenBlockOverrideForFiles gets a reference to the given bool and assigns it to the SmartScreenBlockOverrideForFiles field.

### GetApplicationGuardEnabled

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardEnabled() bool`

GetApplicationGuardEnabled returns the ApplicationGuardEnabled field if non-nil, zero value otherwise.

### GetApplicationGuardEnabledOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardEnabledOk() (bool, bool)`

GetApplicationGuardEnabledOk returns a tuple with the ApplicationGuardEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardEnabled

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardEnabled() bool`

HasApplicationGuardEnabled returns a boolean if a field has been set.

### SetApplicationGuardEnabled

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardEnabled(v bool)`

SetApplicationGuardEnabled gets a reference to the given bool and assigns it to the ApplicationGuardEnabled field.

### GetApplicationGuardBlockFileTransfer

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardBlockFileTransfer() AnyOfmicrosoftGraphApplicationGuardBlockFileTransferType`

GetApplicationGuardBlockFileTransfer returns the ApplicationGuardBlockFileTransfer field if non-nil, zero value otherwise.

### GetApplicationGuardBlockFileTransferOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardBlockFileTransferOk() (AnyOfmicrosoftGraphApplicationGuardBlockFileTransferType, bool)`

GetApplicationGuardBlockFileTransferOk returns a tuple with the ApplicationGuardBlockFileTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardBlockFileTransfer

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardBlockFileTransfer() bool`

HasApplicationGuardBlockFileTransfer returns a boolean if a field has been set.

### SetApplicationGuardBlockFileTransfer

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardBlockFileTransfer(v AnyOfmicrosoftGraphApplicationGuardBlockFileTransferType)`

SetApplicationGuardBlockFileTransfer gets a reference to the given AnyOfmicrosoftGraphApplicationGuardBlockFileTransferType and assigns it to the ApplicationGuardBlockFileTransfer field.

### GetApplicationGuardBlockNonEnterpriseContent

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardBlockNonEnterpriseContent() bool`

GetApplicationGuardBlockNonEnterpriseContent returns the ApplicationGuardBlockNonEnterpriseContent field if non-nil, zero value otherwise.

### GetApplicationGuardBlockNonEnterpriseContentOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardBlockNonEnterpriseContentOk() (bool, bool)`

GetApplicationGuardBlockNonEnterpriseContentOk returns a tuple with the ApplicationGuardBlockNonEnterpriseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardBlockNonEnterpriseContent

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardBlockNonEnterpriseContent() bool`

HasApplicationGuardBlockNonEnterpriseContent returns a boolean if a field has been set.

### SetApplicationGuardBlockNonEnterpriseContent

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardBlockNonEnterpriseContent(v bool)`

SetApplicationGuardBlockNonEnterpriseContent gets a reference to the given bool and assigns it to the ApplicationGuardBlockNonEnterpriseContent field.

### GetApplicationGuardAllowPersistence

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPersistence() bool`

GetApplicationGuardAllowPersistence returns the ApplicationGuardAllowPersistence field if non-nil, zero value otherwise.

### GetApplicationGuardAllowPersistenceOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPersistenceOk() (bool, bool)`

GetApplicationGuardAllowPersistenceOk returns a tuple with the ApplicationGuardAllowPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardAllowPersistence

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardAllowPersistence() bool`

HasApplicationGuardAllowPersistence returns a boolean if a field has been set.

### SetApplicationGuardAllowPersistence

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardAllowPersistence(v bool)`

SetApplicationGuardAllowPersistence gets a reference to the given bool and assigns it to the ApplicationGuardAllowPersistence field.

### GetApplicationGuardForceAuditing

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardForceAuditing() bool`

GetApplicationGuardForceAuditing returns the ApplicationGuardForceAuditing field if non-nil, zero value otherwise.

### GetApplicationGuardForceAuditingOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardForceAuditingOk() (bool, bool)`

GetApplicationGuardForceAuditingOk returns a tuple with the ApplicationGuardForceAuditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardForceAuditing

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardForceAuditing() bool`

HasApplicationGuardForceAuditing returns a boolean if a field has been set.

### SetApplicationGuardForceAuditing

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardForceAuditing(v bool)`

SetApplicationGuardForceAuditing gets a reference to the given bool and assigns it to the ApplicationGuardForceAuditing field.

### GetApplicationGuardBlockClipboardSharing

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardBlockClipboardSharing() AnyOfmicrosoftGraphApplicationGuardBlockClipboardSharingType`

GetApplicationGuardBlockClipboardSharing returns the ApplicationGuardBlockClipboardSharing field if non-nil, zero value otherwise.

### GetApplicationGuardBlockClipboardSharingOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardBlockClipboardSharingOk() (AnyOfmicrosoftGraphApplicationGuardBlockClipboardSharingType, bool)`

GetApplicationGuardBlockClipboardSharingOk returns a tuple with the ApplicationGuardBlockClipboardSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardBlockClipboardSharing

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardBlockClipboardSharing() bool`

HasApplicationGuardBlockClipboardSharing returns a boolean if a field has been set.

### SetApplicationGuardBlockClipboardSharing

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardBlockClipboardSharing(v AnyOfmicrosoftGraphApplicationGuardBlockClipboardSharingType)`

SetApplicationGuardBlockClipboardSharing gets a reference to the given AnyOfmicrosoftGraphApplicationGuardBlockClipboardSharingType and assigns it to the ApplicationGuardBlockClipboardSharing field.

### GetApplicationGuardAllowPrintToPDF

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToPDF() bool`

GetApplicationGuardAllowPrintToPDF returns the ApplicationGuardAllowPrintToPDF field if non-nil, zero value otherwise.

### GetApplicationGuardAllowPrintToPDFOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToPDFOk() (bool, bool)`

GetApplicationGuardAllowPrintToPDFOk returns a tuple with the ApplicationGuardAllowPrintToPDF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardAllowPrintToPDF

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardAllowPrintToPDF() bool`

HasApplicationGuardAllowPrintToPDF returns a boolean if a field has been set.

### SetApplicationGuardAllowPrintToPDF

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToPDF(v bool)`

SetApplicationGuardAllowPrintToPDF gets a reference to the given bool and assigns it to the ApplicationGuardAllowPrintToPDF field.

### GetApplicationGuardAllowPrintToXPS

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToXPS() bool`

GetApplicationGuardAllowPrintToXPS returns the ApplicationGuardAllowPrintToXPS field if non-nil, zero value otherwise.

### GetApplicationGuardAllowPrintToXPSOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToXPSOk() (bool, bool)`

GetApplicationGuardAllowPrintToXPSOk returns a tuple with the ApplicationGuardAllowPrintToXPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardAllowPrintToXPS

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardAllowPrintToXPS() bool`

HasApplicationGuardAllowPrintToXPS returns a boolean if a field has been set.

### SetApplicationGuardAllowPrintToXPS

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToXPS(v bool)`

SetApplicationGuardAllowPrintToXPS gets a reference to the given bool and assigns it to the ApplicationGuardAllowPrintToXPS field.

### GetApplicationGuardAllowPrintToLocalPrinters

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToLocalPrinters() bool`

GetApplicationGuardAllowPrintToLocalPrinters returns the ApplicationGuardAllowPrintToLocalPrinters field if non-nil, zero value otherwise.

### GetApplicationGuardAllowPrintToLocalPrintersOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToLocalPrintersOk() (bool, bool)`

GetApplicationGuardAllowPrintToLocalPrintersOk returns a tuple with the ApplicationGuardAllowPrintToLocalPrinters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardAllowPrintToLocalPrinters

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardAllowPrintToLocalPrinters() bool`

HasApplicationGuardAllowPrintToLocalPrinters returns a boolean if a field has been set.

### SetApplicationGuardAllowPrintToLocalPrinters

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToLocalPrinters(v bool)`

SetApplicationGuardAllowPrintToLocalPrinters gets a reference to the given bool and assigns it to the ApplicationGuardAllowPrintToLocalPrinters field.

### GetApplicationGuardAllowPrintToNetworkPrinters

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToNetworkPrinters() bool`

GetApplicationGuardAllowPrintToNetworkPrinters returns the ApplicationGuardAllowPrintToNetworkPrinters field if non-nil, zero value otherwise.

### GetApplicationGuardAllowPrintToNetworkPrintersOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToNetworkPrintersOk() (bool, bool)`

GetApplicationGuardAllowPrintToNetworkPrintersOk returns a tuple with the ApplicationGuardAllowPrintToNetworkPrinters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationGuardAllowPrintToNetworkPrinters

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasApplicationGuardAllowPrintToNetworkPrinters() bool`

HasApplicationGuardAllowPrintToNetworkPrinters returns a boolean if a field has been set.

### SetApplicationGuardAllowPrintToNetworkPrinters

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToNetworkPrinters(v bool)`

SetApplicationGuardAllowPrintToNetworkPrinters gets a reference to the given bool and assigns it to the ApplicationGuardAllowPrintToNetworkPrinters field.

### GetBitLockerDisableWarningForOtherDiskEncryption

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetBitLockerDisableWarningForOtherDiskEncryption() bool`

GetBitLockerDisableWarningForOtherDiskEncryption returns the BitLockerDisableWarningForOtherDiskEncryption field if non-nil, zero value otherwise.

### GetBitLockerDisableWarningForOtherDiskEncryptionOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetBitLockerDisableWarningForOtherDiskEncryptionOk() (bool, bool)`

GetBitLockerDisableWarningForOtherDiskEncryptionOk returns a tuple with the BitLockerDisableWarningForOtherDiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitLockerDisableWarningForOtherDiskEncryption

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasBitLockerDisableWarningForOtherDiskEncryption() bool`

HasBitLockerDisableWarningForOtherDiskEncryption returns a boolean if a field has been set.

### SetBitLockerDisableWarningForOtherDiskEncryption

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetBitLockerDisableWarningForOtherDiskEncryption(v bool)`

SetBitLockerDisableWarningForOtherDiskEncryption gets a reference to the given bool and assigns it to the BitLockerDisableWarningForOtherDiskEncryption field.

### GetBitLockerEnableStorageCardEncryptionOnMobile

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetBitLockerEnableStorageCardEncryptionOnMobile() bool`

GetBitLockerEnableStorageCardEncryptionOnMobile returns the BitLockerEnableStorageCardEncryptionOnMobile field if non-nil, zero value otherwise.

### GetBitLockerEnableStorageCardEncryptionOnMobileOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetBitLockerEnableStorageCardEncryptionOnMobileOk() (bool, bool)`

GetBitLockerEnableStorageCardEncryptionOnMobileOk returns a tuple with the BitLockerEnableStorageCardEncryptionOnMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitLockerEnableStorageCardEncryptionOnMobile

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasBitLockerEnableStorageCardEncryptionOnMobile() bool`

HasBitLockerEnableStorageCardEncryptionOnMobile returns a boolean if a field has been set.

### SetBitLockerEnableStorageCardEncryptionOnMobile

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetBitLockerEnableStorageCardEncryptionOnMobile(v bool)`

SetBitLockerEnableStorageCardEncryptionOnMobile gets a reference to the given bool and assigns it to the BitLockerEnableStorageCardEncryptionOnMobile field.

### GetBitLockerEncryptDevice

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetBitLockerEncryptDevice() bool`

GetBitLockerEncryptDevice returns the BitLockerEncryptDevice field if non-nil, zero value otherwise.

### GetBitLockerEncryptDeviceOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetBitLockerEncryptDeviceOk() (bool, bool)`

GetBitLockerEncryptDeviceOk returns a tuple with the BitLockerEncryptDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitLockerEncryptDevice

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasBitLockerEncryptDevice() bool`

HasBitLockerEncryptDevice returns a boolean if a field has been set.

### SetBitLockerEncryptDevice

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetBitLockerEncryptDevice(v bool)`

SetBitLockerEncryptDevice gets a reference to the given bool and assigns it to the BitLockerEncryptDevice field.

### GetBitLockerRemovableDrivePolicy

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetBitLockerRemovableDrivePolicy() AnyOfmicrosoftGraphBitLockerRemovableDrivePolicy`

GetBitLockerRemovableDrivePolicy returns the BitLockerRemovableDrivePolicy field if non-nil, zero value otherwise.

### GetBitLockerRemovableDrivePolicyOk

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) GetBitLockerRemovableDrivePolicyOk() (AnyOfmicrosoftGraphBitLockerRemovableDrivePolicy, bool)`

GetBitLockerRemovableDrivePolicyOk returns a tuple with the BitLockerRemovableDrivePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBitLockerRemovableDrivePolicy

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) HasBitLockerRemovableDrivePolicy() bool`

HasBitLockerRemovableDrivePolicy returns a boolean if a field has been set.

### SetBitLockerRemovableDrivePolicy

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetBitLockerRemovableDrivePolicy(v AnyOfmicrosoftGraphBitLockerRemovableDrivePolicy)`

SetBitLockerRemovableDrivePolicy gets a reference to the given AnyOfmicrosoftGraphBitLockerRemovableDrivePolicy and assigns it to the BitLockerRemovableDrivePolicy field.

### SetBitLockerRemovableDrivePolicyExplicitNull

`func (o *MicrosoftGraphWindows10EndpointProtectionConfiguration) SetBitLockerRemovableDrivePolicyExplicitNull(b bool)`

SetBitLockerRemovableDrivePolicyExplicitNull (un)sets BitLockerRemovableDrivePolicy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BitLockerRemovableDrivePolicy value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



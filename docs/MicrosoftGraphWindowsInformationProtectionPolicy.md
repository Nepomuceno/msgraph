# MicrosoftGraphWindowsInformationProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Policy display name. | [optional] 
**Description** | Pointer to **string** | The policy&#39;s description. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time the policy was created. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last time the policy was modified. | [optional] 
**Version** | Pointer to **string** | Version of the entity. | [optional] 
**EnforcementLevel** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel**](anyOf&lt;microsoft.graph.windowsInformationProtectionEnforcementLevel&gt;.md) | WIP enforcement level.See the Enum definition for supported values | [optional] 
**EnterpriseDomain** | Pointer to **string** | Primary enterprise domain | [optional] 
**EnterpriseProtectedDomainNames** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](anyOf&lt;microsoft.graph.windowsInformationProtectionResourceCollection&gt;.md) | List of enterprise domains to be protected | [optional] 
**ProtectionUnderLockConfigRequired** | Pointer to **bool** | Specifies whether the protection under lock feature (also known as encrypt under pin) should be configured | [optional] 
**DataRecoveryCertificate** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate**](anyOf&lt;microsoft.graph.windowsInformationProtectionDataRecoveryCertificate&gt;.md) | Specifies a recovery certificate that can be used for data recovery of encrypted files. This is the same as the data recovery agent(DRA) certificate for encrypting file system(EFS) | [optional] 
**RevokeOnUnenrollDisabled** | Pointer to **bool** | This policy controls whether to revoke the WIP keys when a device unenrolls from the management service. If set to 1 (Don&#39;t revoke keys), the keys will not be revoked and the user will continue to have access to protected files after unenrollment. If the keys are not revoked, there will be no revoked file cleanup subsequently. | [optional] 
**RightsManagementServicesTemplateId** | Pointer to **string** | TemplateID GUID to use for RMS encryption. The RMS template allows the IT admin to configure the details about who has access to RMS-protected file and how long they have access | [optional] 
**AzureRightsManagementServicesAllowed** | Pointer to **bool** | Specifies whether to allow Azure RMS encryption for WIP | [optional] 
**IconsVisible** | Pointer to **bool** | Determines whether overlays are added to icons for WIP protected files in Explorer and enterprise only app tiles in the Start menu. Starting in Windows 10, version 1703 this setting also configures the visibility of the WIP icon in the title bar of a WIP-protected app | [optional] 
**ProtectedApps** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionApp**](anyOf&lt;microsoft.graph.windowsInformationProtectionApp&gt;.md) | Protected applications can access enterprise data and the data handled by those applications are protected with encryption | [optional] 
**ExemptApps** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionApp**](anyOf&lt;microsoft.graph.windowsInformationProtectionApp&gt;.md) | Exempt applications can also access enterprise data, but the data handled by those applications are not protected. This is because some critical enterprise applications may have compatibility problems with encrypted data. | [optional] 
**EnterpriseNetworkDomainNames** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](anyOf&lt;microsoft.graph.windowsInformationProtectionResourceCollection&gt;.md) | This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is sent to a device will be considered enterprise data and protected These locations will be considered a safe destination for enterprise data to be shared to | [optional] 
**EnterpriseProxiedDomains** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection**](anyOf&lt;microsoft.graph.windowsInformationProtectionProxiedDomainCollection&gt;.md) | Contains a list of Enterprise resource domains hosted in the cloud that need to be protected. Connections to these resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this purpose must also be configured using the EnterpriseInternalProxyServers policy | [optional] 
**EnterpriseIPRanges** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection**](anyOf&lt;microsoft.graph.windowsInformationProtectionIPRangeCollection&gt;.md) | Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those computers will be considered part of the enterprise and protected. These locations will be considered a safe destination for enterprise data to be shared to | [optional] 
**EnterpriseIPRangesAreAuthoritative** | Pointer to **bool** | Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other subnets. Default is false | [optional] 
**EnterpriseProxyServers** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](anyOf&lt;microsoft.graph.windowsInformationProtectionResourceCollection&gt;.md) | This is a list of proxy servers. Any server not on this list is considered non-enterprise | [optional] 
**EnterpriseInternalProxyServers** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](anyOf&lt;microsoft.graph.windowsInformationProtectionResourceCollection&gt;.md) | This is the comma-separated list of internal proxy servers. For example, \&quot;157.54.14.28, 157.54.11.118, 10.202.14.167, 157.53.14.163, 157.69.210.59\&quot;. These proxies have been configured by the admin to connect to specific resources on the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring the EnterpriseProxiedDomains policy to force traffic to the matched domains through these proxies | [optional] 
**EnterpriseProxyServersAreAuthoritative** | Pointer to **bool** | Boolean value that tells the client to accept the configured list of proxies and not try to detect other work proxies. Default is false | [optional] 
**NeutralDomainResources** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](anyOf&lt;microsoft.graph.windowsInformationProtectionResourceCollection&gt;.md) | List of domain names that can used for work or personal resource | [optional] 
**IndexingEncryptedStoresOrItemsBlocked** | Pointer to **bool** | This switch is for the Windows Search Indexer, to allow or disallow indexing of items | [optional] 
**SmbAutoEncryptedFileExtensions** | Pointer to [**[]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection**](anyOf&lt;microsoft.graph.windowsInformationProtectionResourceCollection&gt;.md) | Specifies a list of file extensions, so that files with these extensions are encrypted when copying from an SMB share within the corporate boundary | [optional] 
**IsAssigned** | Pointer to **bool** | Indicates if the policy is deployed to any inclusion groups or not. | [optional] 
**ProtectedAppLockerFiles** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionAppLockerFile**](microsoft.graph.windowsInformationProtectionAppLockerFile.md) |  | [optional] 
**ExemptAppLockerFiles** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionAppLockerFile**](microsoft.graph.windowsInformationProtectionAppLockerFile.md) |  | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphTargetedManagedAppPolicyAssignment**](microsoft.graph.targetedManagedAppPolicyAssignment.md) |  | [optional] 
**RevokeOnMdmHandoffDisabled** | Pointer to **bool** | New property in RS2, pending documentation | [optional] 
**MdmEnrollmentUrl** | Pointer to **string** | Enrollment url for the MDM | [optional] 
**WindowsHelloForBusinessBlocked** | Pointer to **bool** | Boolean value that sets Windows Hello for Business as a method for signing into Windows. | [optional] 
**PinMinimumLength** | Pointer to **int32** | Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest. | [optional] 
**PinUppercaseLetters** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements**](anyOf&lt;microsoft.graph.windowsInformationProtectionPinCharacterRequirements&gt;.md) | Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow. | [optional] 
**PinLowercaseLetters** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements**](anyOf&lt;microsoft.graph.windowsInformationProtectionPinCharacterRequirements&gt;.md) | Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow. | [optional] 
**PinSpecialCharacters** | Pointer to [**AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements**](anyOf&lt;microsoft.graph.windowsInformationProtectionPinCharacterRequirements&gt;.md) | Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! \&quot; # $ % &amp; &#39; ( ) * + , - . / : ; &lt; &#x3D; &gt; ? @ [ \\ ] ^ _ &#x60; { | } ~. Default is NotAllow. | [optional] 
**PinExpirationDays** | Pointer to **int32** | Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user&#39;s PIN will never expire. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**NumberOfPastPinsRemembered** | Pointer to **int32** | Integer value that specifies the number of past PINs that can be associated to a user account that can&#39;t be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0. | [optional] 
**PasswordMaximumAttemptCount** | Pointer to **int32** | The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 &lt;&#x3D; X &lt;&#x3D; 16 for desktop and 0 &lt;&#x3D; X &lt;&#x3D; 999 for mobile devices. | [optional] 
**MinutesOfInactivityBeforeDeviceLock** | Pointer to **int32** | Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 &lt;&#x3D; X &lt;&#x3D; 999. | [optional] 
**DaysWithoutContactBeforeUnenroll** | Pointer to **int32** | Offline interval before app data is wiped (days)  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetEnforcementLevel

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnforcementLevel() AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel`

GetEnforcementLevel returns the EnforcementLevel field if non-nil, zero value otherwise.

### GetEnforcementLevelOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnforcementLevelOk() (AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel, bool)`

GetEnforcementLevelOk returns a tuple with the EnforcementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnforcementLevel

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnforcementLevel() bool`

HasEnforcementLevel returns a boolean if a field has been set.

### SetEnforcementLevel

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnforcementLevel(v AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel)`

SetEnforcementLevel gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel and assigns it to the EnforcementLevel field.

### GetEnterpriseDomain

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseDomain() string`

GetEnterpriseDomain returns the EnterpriseDomain field if non-nil, zero value otherwise.

### GetEnterpriseDomainOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseDomainOk() (string, bool)`

GetEnterpriseDomainOk returns a tuple with the EnterpriseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseDomain

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnterpriseDomain() bool`

HasEnterpriseDomain returns a boolean if a field has been set.

### SetEnterpriseDomain

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseDomain(v string)`

SetEnterpriseDomain gets a reference to the given string and assigns it to the EnterpriseDomain field.

### SetEnterpriseDomainExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseDomainExplicitNull(b bool)`

SetEnterpriseDomainExplicitNull (un)sets EnterpriseDomain to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseDomain value is set to nil even if false is passed
### GetEnterpriseProtectedDomainNames

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseProtectedDomainNames() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProtectedDomainNames returns the EnterpriseProtectedDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseProtectedDomainNamesOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseProtectedDomainNamesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProtectedDomainNamesOk returns a tuple with the EnterpriseProtectedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProtectedDomainNames

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnterpriseProtectedDomainNames() bool`

HasEnterpriseProtectedDomainNames returns a boolean if a field has been set.

### SetEnterpriseProtectedDomainNames

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseProtectedDomainNames(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProtectedDomainNames gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseProtectedDomainNames field.

### GetProtectionUnderLockConfigRequired

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetProtectionUnderLockConfigRequired() bool`

GetProtectionUnderLockConfigRequired returns the ProtectionUnderLockConfigRequired field if non-nil, zero value otherwise.

### GetProtectionUnderLockConfigRequiredOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetProtectionUnderLockConfigRequiredOk() (bool, bool)`

GetProtectionUnderLockConfigRequiredOk returns a tuple with the ProtectionUnderLockConfigRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtectionUnderLockConfigRequired

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasProtectionUnderLockConfigRequired() bool`

HasProtectionUnderLockConfigRequired returns a boolean if a field has been set.

### SetProtectionUnderLockConfigRequired

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetProtectionUnderLockConfigRequired(v bool)`

SetProtectionUnderLockConfigRequired gets a reference to the given bool and assigns it to the ProtectionUnderLockConfigRequired field.

### GetDataRecoveryCertificate

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetDataRecoveryCertificate() AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate`

GetDataRecoveryCertificate returns the DataRecoveryCertificate field if non-nil, zero value otherwise.

### GetDataRecoveryCertificateOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetDataRecoveryCertificateOk() (AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate, bool)`

GetDataRecoveryCertificateOk returns a tuple with the DataRecoveryCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataRecoveryCertificate

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasDataRecoveryCertificate() bool`

HasDataRecoveryCertificate returns a boolean if a field has been set.

### SetDataRecoveryCertificate

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetDataRecoveryCertificate(v AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate)`

SetDataRecoveryCertificate gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate and assigns it to the DataRecoveryCertificate field.

### SetDataRecoveryCertificateExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetDataRecoveryCertificateExplicitNull(b bool)`

SetDataRecoveryCertificateExplicitNull (un)sets DataRecoveryCertificate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DataRecoveryCertificate value is set to nil even if false is passed
### GetRevokeOnUnenrollDisabled

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetRevokeOnUnenrollDisabled() bool`

GetRevokeOnUnenrollDisabled returns the RevokeOnUnenrollDisabled field if non-nil, zero value otherwise.

### GetRevokeOnUnenrollDisabledOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetRevokeOnUnenrollDisabledOk() (bool, bool)`

GetRevokeOnUnenrollDisabledOk returns a tuple with the RevokeOnUnenrollDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRevokeOnUnenrollDisabled

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasRevokeOnUnenrollDisabled() bool`

HasRevokeOnUnenrollDisabled returns a boolean if a field has been set.

### SetRevokeOnUnenrollDisabled

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetRevokeOnUnenrollDisabled(v bool)`

SetRevokeOnUnenrollDisabled gets a reference to the given bool and assigns it to the RevokeOnUnenrollDisabled field.

### GetRightsManagementServicesTemplateId

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetRightsManagementServicesTemplateId() string`

GetRightsManagementServicesTemplateId returns the RightsManagementServicesTemplateId field if non-nil, zero value otherwise.

### GetRightsManagementServicesTemplateIdOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetRightsManagementServicesTemplateIdOk() (string, bool)`

GetRightsManagementServicesTemplateIdOk returns a tuple with the RightsManagementServicesTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRightsManagementServicesTemplateId

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasRightsManagementServicesTemplateId() bool`

HasRightsManagementServicesTemplateId returns a boolean if a field has been set.

### SetRightsManagementServicesTemplateId

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetRightsManagementServicesTemplateId(v string)`

SetRightsManagementServicesTemplateId gets a reference to the given string and assigns it to the RightsManagementServicesTemplateId field.

### SetRightsManagementServicesTemplateIdExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetRightsManagementServicesTemplateIdExplicitNull(b bool)`

SetRightsManagementServicesTemplateIdExplicitNull (un)sets RightsManagementServicesTemplateId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RightsManagementServicesTemplateId value is set to nil even if false is passed
### GetAzureRightsManagementServicesAllowed

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetAzureRightsManagementServicesAllowed() bool`

GetAzureRightsManagementServicesAllowed returns the AzureRightsManagementServicesAllowed field if non-nil, zero value otherwise.

### GetAzureRightsManagementServicesAllowedOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetAzureRightsManagementServicesAllowedOk() (bool, bool)`

GetAzureRightsManagementServicesAllowedOk returns a tuple with the AzureRightsManagementServicesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureRightsManagementServicesAllowed

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasAzureRightsManagementServicesAllowed() bool`

HasAzureRightsManagementServicesAllowed returns a boolean if a field has been set.

### SetAzureRightsManagementServicesAllowed

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetAzureRightsManagementServicesAllowed(v bool)`

SetAzureRightsManagementServicesAllowed gets a reference to the given bool and assigns it to the AzureRightsManagementServicesAllowed field.

### GetIconsVisible

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetIconsVisible() bool`

GetIconsVisible returns the IconsVisible field if non-nil, zero value otherwise.

### GetIconsVisibleOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetIconsVisibleOk() (bool, bool)`

GetIconsVisibleOk returns a tuple with the IconsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIconsVisible

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasIconsVisible() bool`

HasIconsVisible returns a boolean if a field has been set.

### SetIconsVisible

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetIconsVisible(v bool)`

SetIconsVisible gets a reference to the given bool and assigns it to the IconsVisible field.

### GetProtectedApps

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetProtectedApps() []AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetProtectedApps returns the ProtectedApps field if non-nil, zero value otherwise.

### GetProtectedAppsOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetProtectedAppsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetProtectedAppsOk returns a tuple with the ProtectedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtectedApps

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasProtectedApps() bool`

HasProtectedApps returns a boolean if a field has been set.

### SetProtectedApps

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetProtectedApps(v []AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetProtectedApps gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionApp and assigns it to the ProtectedApps field.

### GetExemptApps

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetExemptApps() []AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetExemptApps returns the ExemptApps field if non-nil, zero value otherwise.

### GetExemptAppsOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetExemptAppsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetExemptAppsOk returns a tuple with the ExemptApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExemptApps

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasExemptApps() bool`

HasExemptApps returns a boolean if a field has been set.

### SetExemptApps

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetExemptApps(v []AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetExemptApps gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionApp and assigns it to the ExemptApps field.

### GetEnterpriseNetworkDomainNames

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseNetworkDomainNames() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseNetworkDomainNames returns the EnterpriseNetworkDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseNetworkDomainNamesOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseNetworkDomainNamesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseNetworkDomainNamesOk returns a tuple with the EnterpriseNetworkDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseNetworkDomainNames

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnterpriseNetworkDomainNames() bool`

HasEnterpriseNetworkDomainNames returns a boolean if a field has been set.

### SetEnterpriseNetworkDomainNames

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseNetworkDomainNames(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseNetworkDomainNames gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseNetworkDomainNames field.

### GetEnterpriseProxiedDomains

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseProxiedDomains() []AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection`

GetEnterpriseProxiedDomains returns the EnterpriseProxiedDomains field if non-nil, zero value otherwise.

### GetEnterpriseProxiedDomainsOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseProxiedDomainsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection, bool)`

GetEnterpriseProxiedDomainsOk returns a tuple with the EnterpriseProxiedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProxiedDomains

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnterpriseProxiedDomains() bool`

HasEnterpriseProxiedDomains returns a boolean if a field has been set.

### SetEnterpriseProxiedDomains

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseProxiedDomains(v []AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection)`

SetEnterpriseProxiedDomains gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection and assigns it to the EnterpriseProxiedDomains field.

### GetEnterpriseIPRanges

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseIPRanges() []AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection`

GetEnterpriseIPRanges returns the EnterpriseIPRanges field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseIPRangesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection, bool)`

GetEnterpriseIPRangesOk returns a tuple with the EnterpriseIPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseIPRanges

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnterpriseIPRanges() bool`

HasEnterpriseIPRanges returns a boolean if a field has been set.

### SetEnterpriseIPRanges

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseIPRanges(v []AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection)`

SetEnterpriseIPRanges gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection and assigns it to the EnterpriseIPRanges field.

### GetEnterpriseIPRangesAreAuthoritative

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseIPRangesAreAuthoritative() bool`

GetEnterpriseIPRangesAreAuthoritative returns the EnterpriseIPRangesAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesAreAuthoritativeOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseIPRangesAreAuthoritativeOk() (bool, bool)`

GetEnterpriseIPRangesAreAuthoritativeOk returns a tuple with the EnterpriseIPRangesAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseIPRangesAreAuthoritative

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnterpriseIPRangesAreAuthoritative() bool`

HasEnterpriseIPRangesAreAuthoritative returns a boolean if a field has been set.

### SetEnterpriseIPRangesAreAuthoritative

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseIPRangesAreAuthoritative(v bool)`

SetEnterpriseIPRangesAreAuthoritative gets a reference to the given bool and assigns it to the EnterpriseIPRangesAreAuthoritative field.

### GetEnterpriseProxyServers

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseProxyServers() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProxyServers returns the EnterpriseProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseProxyServersOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProxyServersOk returns a tuple with the EnterpriseProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProxyServers

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnterpriseProxyServers() bool`

HasEnterpriseProxyServers returns a boolean if a field has been set.

### SetEnterpriseProxyServers

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseProxyServers(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProxyServers gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseProxyServers field.

### GetEnterpriseInternalProxyServers

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseInternalProxyServers() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseInternalProxyServers returns the EnterpriseInternalProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseInternalProxyServersOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseInternalProxyServersOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseInternalProxyServersOk returns a tuple with the EnterpriseInternalProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseInternalProxyServers

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnterpriseInternalProxyServers() bool`

HasEnterpriseInternalProxyServers returns a boolean if a field has been set.

### SetEnterpriseInternalProxyServers

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseInternalProxyServers(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseInternalProxyServers gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseInternalProxyServers field.

### GetEnterpriseProxyServersAreAuthoritative

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseProxyServersAreAuthoritative() bool`

GetEnterpriseProxyServersAreAuthoritative returns the EnterpriseProxyServersAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersAreAuthoritativeOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetEnterpriseProxyServersAreAuthoritativeOk() (bool, bool)`

GetEnterpriseProxyServersAreAuthoritativeOk returns a tuple with the EnterpriseProxyServersAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProxyServersAreAuthoritative

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasEnterpriseProxyServersAreAuthoritative() bool`

HasEnterpriseProxyServersAreAuthoritative returns a boolean if a field has been set.

### SetEnterpriseProxyServersAreAuthoritative

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetEnterpriseProxyServersAreAuthoritative(v bool)`

SetEnterpriseProxyServersAreAuthoritative gets a reference to the given bool and assigns it to the EnterpriseProxyServersAreAuthoritative field.

### GetNeutralDomainResources

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetNeutralDomainResources() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetNeutralDomainResources returns the NeutralDomainResources field if non-nil, zero value otherwise.

### GetNeutralDomainResourcesOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetNeutralDomainResourcesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetNeutralDomainResourcesOk returns a tuple with the NeutralDomainResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNeutralDomainResources

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasNeutralDomainResources() bool`

HasNeutralDomainResources returns a boolean if a field has been set.

### SetNeutralDomainResources

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetNeutralDomainResources(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetNeutralDomainResources gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the NeutralDomainResources field.

### GetIndexingEncryptedStoresOrItemsBlocked

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetIndexingEncryptedStoresOrItemsBlocked() bool`

GetIndexingEncryptedStoresOrItemsBlocked returns the IndexingEncryptedStoresOrItemsBlocked field if non-nil, zero value otherwise.

### GetIndexingEncryptedStoresOrItemsBlockedOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetIndexingEncryptedStoresOrItemsBlockedOk() (bool, bool)`

GetIndexingEncryptedStoresOrItemsBlockedOk returns a tuple with the IndexingEncryptedStoresOrItemsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndexingEncryptedStoresOrItemsBlocked

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasIndexingEncryptedStoresOrItemsBlocked() bool`

HasIndexingEncryptedStoresOrItemsBlocked returns a boolean if a field has been set.

### SetIndexingEncryptedStoresOrItemsBlocked

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetIndexingEncryptedStoresOrItemsBlocked(v bool)`

SetIndexingEncryptedStoresOrItemsBlocked gets a reference to the given bool and assigns it to the IndexingEncryptedStoresOrItemsBlocked field.

### GetSmbAutoEncryptedFileExtensions

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetSmbAutoEncryptedFileExtensions() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetSmbAutoEncryptedFileExtensions returns the SmbAutoEncryptedFileExtensions field if non-nil, zero value otherwise.

### GetSmbAutoEncryptedFileExtensionsOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetSmbAutoEncryptedFileExtensionsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetSmbAutoEncryptedFileExtensionsOk returns a tuple with the SmbAutoEncryptedFileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmbAutoEncryptedFileExtensions

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasSmbAutoEncryptedFileExtensions() bool`

HasSmbAutoEncryptedFileExtensions returns a boolean if a field has been set.

### SetSmbAutoEncryptedFileExtensions

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetSmbAutoEncryptedFileExtensions(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetSmbAutoEncryptedFileExtensions gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the SmbAutoEncryptedFileExtensions field.

### GetIsAssigned

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetIsAssignedOk() (bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAssigned

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### SetIsAssigned

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetIsAssigned(v bool)`

SetIsAssigned gets a reference to the given bool and assigns it to the IsAssigned field.

### GetProtectedAppLockerFiles

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetProtectedAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetProtectedAppLockerFiles returns the ProtectedAppLockerFiles field if non-nil, zero value otherwise.

### GetProtectedAppLockerFilesOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetProtectedAppLockerFilesOk() ([]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetProtectedAppLockerFilesOk returns a tuple with the ProtectedAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtectedAppLockerFiles

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasProtectedAppLockerFiles() bool`

HasProtectedAppLockerFiles returns a boolean if a field has been set.

### SetProtectedAppLockerFiles

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetProtectedAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetProtectedAppLockerFiles gets a reference to the given []MicrosoftGraphWindowsInformationProtectionAppLockerFile and assigns it to the ProtectedAppLockerFiles field.

### GetExemptAppLockerFiles

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetExemptAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetExemptAppLockerFiles returns the ExemptAppLockerFiles field if non-nil, zero value otherwise.

### GetExemptAppLockerFilesOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetExemptAppLockerFilesOk() ([]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetExemptAppLockerFilesOk returns a tuple with the ExemptAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExemptAppLockerFiles

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasExemptAppLockerFiles() bool`

HasExemptAppLockerFiles returns a boolean if a field has been set.

### SetExemptAppLockerFiles

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetExemptAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetExemptAppLockerFiles gets a reference to the given []MicrosoftGraphWindowsInformationProtectionAppLockerFile and assigns it to the ExemptAppLockerFiles field.

### GetAssignments

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetAssignmentsOk() ([]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.

### GetRevokeOnMdmHandoffDisabled

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabled() bool`

GetRevokeOnMdmHandoffDisabled returns the RevokeOnMdmHandoffDisabled field if non-nil, zero value otherwise.

### GetRevokeOnMdmHandoffDisabledOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetRevokeOnMdmHandoffDisabledOk() (bool, bool)`

GetRevokeOnMdmHandoffDisabledOk returns a tuple with the RevokeOnMdmHandoffDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRevokeOnMdmHandoffDisabled

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasRevokeOnMdmHandoffDisabled() bool`

HasRevokeOnMdmHandoffDisabled returns a boolean if a field has been set.

### SetRevokeOnMdmHandoffDisabled

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetRevokeOnMdmHandoffDisabled(v bool)`

SetRevokeOnMdmHandoffDisabled gets a reference to the given bool and assigns it to the RevokeOnMdmHandoffDisabled field.

### GetMdmEnrollmentUrl

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetMdmEnrollmentUrl() string`

GetMdmEnrollmentUrl returns the MdmEnrollmentUrl field if non-nil, zero value otherwise.

### GetMdmEnrollmentUrlOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetMdmEnrollmentUrlOk() (string, bool)`

GetMdmEnrollmentUrlOk returns a tuple with the MdmEnrollmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMdmEnrollmentUrl

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasMdmEnrollmentUrl() bool`

HasMdmEnrollmentUrl returns a boolean if a field has been set.

### SetMdmEnrollmentUrl

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetMdmEnrollmentUrl(v string)`

SetMdmEnrollmentUrl gets a reference to the given string and assigns it to the MdmEnrollmentUrl field.

### SetMdmEnrollmentUrlExplicitNull

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetMdmEnrollmentUrlExplicitNull(b bool)`

SetMdmEnrollmentUrlExplicitNull (un)sets MdmEnrollmentUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MdmEnrollmentUrl value is set to nil even if false is passed
### GetWindowsHelloForBusinessBlocked

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlocked() bool`

GetWindowsHelloForBusinessBlocked returns the WindowsHelloForBusinessBlocked field if non-nil, zero value otherwise.

### GetWindowsHelloForBusinessBlockedOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetWindowsHelloForBusinessBlockedOk() (bool, bool)`

GetWindowsHelloForBusinessBlockedOk returns a tuple with the WindowsHelloForBusinessBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsHelloForBusinessBlocked

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasWindowsHelloForBusinessBlocked() bool`

HasWindowsHelloForBusinessBlocked returns a boolean if a field has been set.

### SetWindowsHelloForBusinessBlocked

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetWindowsHelloForBusinessBlocked(v bool)`

SetWindowsHelloForBusinessBlocked gets a reference to the given bool and assigns it to the WindowsHelloForBusinessBlocked field.

### GetPinMinimumLength

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinMinimumLength() int32`

GetPinMinimumLength returns the PinMinimumLength field if non-nil, zero value otherwise.

### GetPinMinimumLengthOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinMinimumLengthOk() (int32, bool)`

GetPinMinimumLengthOk returns a tuple with the PinMinimumLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinMinimumLength

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasPinMinimumLength() bool`

HasPinMinimumLength returns a boolean if a field has been set.

### SetPinMinimumLength

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetPinMinimumLength(v int32)`

SetPinMinimumLength gets a reference to the given int32 and assigns it to the PinMinimumLength field.

### GetPinUppercaseLetters

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinUppercaseLetters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements`

GetPinUppercaseLetters returns the PinUppercaseLetters field if non-nil, zero value otherwise.

### GetPinUppercaseLettersOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinUppercaseLettersOk() (AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool)`

GetPinUppercaseLettersOk returns a tuple with the PinUppercaseLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinUppercaseLetters

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasPinUppercaseLetters() bool`

HasPinUppercaseLetters returns a boolean if a field has been set.

### SetPinUppercaseLetters

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetPinUppercaseLetters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements)`

SetPinUppercaseLetters gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements and assigns it to the PinUppercaseLetters field.

### GetPinLowercaseLetters

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinLowercaseLetters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements`

GetPinLowercaseLetters returns the PinLowercaseLetters field if non-nil, zero value otherwise.

### GetPinLowercaseLettersOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinLowercaseLettersOk() (AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool)`

GetPinLowercaseLettersOk returns a tuple with the PinLowercaseLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinLowercaseLetters

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasPinLowercaseLetters() bool`

HasPinLowercaseLetters returns a boolean if a field has been set.

### SetPinLowercaseLetters

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetPinLowercaseLetters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements)`

SetPinLowercaseLetters gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements and assigns it to the PinLowercaseLetters field.

### GetPinSpecialCharacters

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinSpecialCharacters() AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements`

GetPinSpecialCharacters returns the PinSpecialCharacters field if non-nil, zero value otherwise.

### GetPinSpecialCharactersOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinSpecialCharactersOk() (AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements, bool)`

GetPinSpecialCharactersOk returns a tuple with the PinSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinSpecialCharacters

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasPinSpecialCharacters() bool`

HasPinSpecialCharacters returns a boolean if a field has been set.

### SetPinSpecialCharacters

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetPinSpecialCharacters(v AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements)`

SetPinSpecialCharacters gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionPinCharacterRequirements and assigns it to the PinSpecialCharacters field.

### GetPinExpirationDays

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinExpirationDays() int32`

GetPinExpirationDays returns the PinExpirationDays field if non-nil, zero value otherwise.

### GetPinExpirationDaysOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPinExpirationDaysOk() (int32, bool)`

GetPinExpirationDaysOk returns a tuple with the PinExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPinExpirationDays

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasPinExpirationDays() bool`

HasPinExpirationDays returns a boolean if a field has been set.

### SetPinExpirationDays

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetPinExpirationDays(v int32)`

SetPinExpirationDays gets a reference to the given int32 and assigns it to the PinExpirationDays field.

### GetNumberOfPastPinsRemembered

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetNumberOfPastPinsRemembered() int32`

GetNumberOfPastPinsRemembered returns the NumberOfPastPinsRemembered field if non-nil, zero value otherwise.

### GetNumberOfPastPinsRememberedOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetNumberOfPastPinsRememberedOk() (int32, bool)`

GetNumberOfPastPinsRememberedOk returns a tuple with the NumberOfPastPinsRemembered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberOfPastPinsRemembered

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasNumberOfPastPinsRemembered() bool`

HasNumberOfPastPinsRemembered returns a boolean if a field has been set.

### SetNumberOfPastPinsRemembered

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetNumberOfPastPinsRemembered(v int32)`

SetNumberOfPastPinsRemembered gets a reference to the given int32 and assigns it to the NumberOfPastPinsRemembered field.

### GetPasswordMaximumAttemptCount

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCount() int32`

GetPasswordMaximumAttemptCount returns the PasswordMaximumAttemptCount field if non-nil, zero value otherwise.

### GetPasswordMaximumAttemptCountOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetPasswordMaximumAttemptCountOk() (int32, bool)`

GetPasswordMaximumAttemptCountOk returns a tuple with the PasswordMaximumAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPasswordMaximumAttemptCount

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasPasswordMaximumAttemptCount() bool`

HasPasswordMaximumAttemptCount returns a boolean if a field has been set.

### SetPasswordMaximumAttemptCount

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetPasswordMaximumAttemptCount(v int32)`

SetPasswordMaximumAttemptCount gets a reference to the given int32 and assigns it to the PasswordMaximumAttemptCount field.

### GetMinutesOfInactivityBeforeDeviceLock

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLock() int32`

GetMinutesOfInactivityBeforeDeviceLock returns the MinutesOfInactivityBeforeDeviceLock field if non-nil, zero value otherwise.

### GetMinutesOfInactivityBeforeDeviceLockOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetMinutesOfInactivityBeforeDeviceLockOk() (int32, bool)`

GetMinutesOfInactivityBeforeDeviceLockOk returns a tuple with the MinutesOfInactivityBeforeDeviceLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinutesOfInactivityBeforeDeviceLock

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasMinutesOfInactivityBeforeDeviceLock() bool`

HasMinutesOfInactivityBeforeDeviceLock returns a boolean if a field has been set.

### SetMinutesOfInactivityBeforeDeviceLock

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetMinutesOfInactivityBeforeDeviceLock(v int32)`

SetMinutesOfInactivityBeforeDeviceLock gets a reference to the given int32 and assigns it to the MinutesOfInactivityBeforeDeviceLock field.

### GetDaysWithoutContactBeforeUnenroll

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenroll() int32`

GetDaysWithoutContactBeforeUnenroll returns the DaysWithoutContactBeforeUnenroll field if non-nil, zero value otherwise.

### GetDaysWithoutContactBeforeUnenrollOk

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) GetDaysWithoutContactBeforeUnenrollOk() (int32, bool)`

GetDaysWithoutContactBeforeUnenrollOk returns a tuple with the DaysWithoutContactBeforeUnenroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDaysWithoutContactBeforeUnenroll

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) HasDaysWithoutContactBeforeUnenroll() bool`

HasDaysWithoutContactBeforeUnenroll returns a boolean if a field has been set.

### SetDaysWithoutContactBeforeUnenroll

`func (o *MicrosoftGraphWindowsInformationProtectionPolicy) SetDaysWithoutContactBeforeUnenroll(v int32)`

SetDaysWithoutContactBeforeUnenroll gets a reference to the given int32 and assigns it to the DaysWithoutContactBeforeUnenroll field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MicrosoftGraphMdmWindowsInformationProtectionPolicy

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

## Methods

### GetId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetEnforcementLevel

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnforcementLevel() AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel`

GetEnforcementLevel returns the EnforcementLevel field if non-nil, zero value otherwise.

### GetEnforcementLevelOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnforcementLevelOk() (AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel, bool)`

GetEnforcementLevelOk returns a tuple with the EnforcementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnforcementLevel

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnforcementLevel() bool`

HasEnforcementLevel returns a boolean if a field has been set.

### SetEnforcementLevel

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnforcementLevel(v AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel)`

SetEnforcementLevel gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel and assigns it to the EnforcementLevel field.

### GetEnterpriseDomain

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseDomain() string`

GetEnterpriseDomain returns the EnterpriseDomain field if non-nil, zero value otherwise.

### GetEnterpriseDomainOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseDomainOk() (string, bool)`

GetEnterpriseDomainOk returns a tuple with the EnterpriseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseDomain

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseDomain() bool`

HasEnterpriseDomain returns a boolean if a field has been set.

### SetEnterpriseDomain

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseDomain(v string)`

SetEnterpriseDomain gets a reference to the given string and assigns it to the EnterpriseDomain field.

### SetEnterpriseDomainExplicitNull

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseDomainExplicitNull(b bool)`

SetEnterpriseDomainExplicitNull (un)sets EnterpriseDomain to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseDomain value is set to nil even if false is passed
### GetEnterpriseProtectedDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProtectedDomainNames() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProtectedDomainNames returns the EnterpriseProtectedDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseProtectedDomainNamesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProtectedDomainNamesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProtectedDomainNamesOk returns a tuple with the EnterpriseProtectedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProtectedDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseProtectedDomainNames() bool`

HasEnterpriseProtectedDomainNames returns a boolean if a field has been set.

### SetEnterpriseProtectedDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseProtectedDomainNames(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProtectedDomainNames gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseProtectedDomainNames field.

### GetProtectionUnderLockConfigRequired

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectionUnderLockConfigRequired() bool`

GetProtectionUnderLockConfigRequired returns the ProtectionUnderLockConfigRequired field if non-nil, zero value otherwise.

### GetProtectionUnderLockConfigRequiredOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectionUnderLockConfigRequiredOk() (bool, bool)`

GetProtectionUnderLockConfigRequiredOk returns a tuple with the ProtectionUnderLockConfigRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtectionUnderLockConfigRequired

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasProtectionUnderLockConfigRequired() bool`

HasProtectionUnderLockConfigRequired returns a boolean if a field has been set.

### SetProtectionUnderLockConfigRequired

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetProtectionUnderLockConfigRequired(v bool)`

SetProtectionUnderLockConfigRequired gets a reference to the given bool and assigns it to the ProtectionUnderLockConfigRequired field.

### GetDataRecoveryCertificate

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDataRecoveryCertificate() AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate`

GetDataRecoveryCertificate returns the DataRecoveryCertificate field if non-nil, zero value otherwise.

### GetDataRecoveryCertificateOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetDataRecoveryCertificateOk() (AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate, bool)`

GetDataRecoveryCertificateOk returns a tuple with the DataRecoveryCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataRecoveryCertificate

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasDataRecoveryCertificate() bool`

HasDataRecoveryCertificate returns a boolean if a field has been set.

### SetDataRecoveryCertificate

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDataRecoveryCertificate(v AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate)`

SetDataRecoveryCertificate gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate and assigns it to the DataRecoveryCertificate field.

### SetDataRecoveryCertificateExplicitNull

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetDataRecoveryCertificateExplicitNull(b bool)`

SetDataRecoveryCertificateExplicitNull (un)sets DataRecoveryCertificate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DataRecoveryCertificate value is set to nil even if false is passed
### GetRevokeOnUnenrollDisabled

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetRevokeOnUnenrollDisabled() bool`

GetRevokeOnUnenrollDisabled returns the RevokeOnUnenrollDisabled field if non-nil, zero value otherwise.

### GetRevokeOnUnenrollDisabledOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetRevokeOnUnenrollDisabledOk() (bool, bool)`

GetRevokeOnUnenrollDisabledOk returns a tuple with the RevokeOnUnenrollDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRevokeOnUnenrollDisabled

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasRevokeOnUnenrollDisabled() bool`

HasRevokeOnUnenrollDisabled returns a boolean if a field has been set.

### SetRevokeOnUnenrollDisabled

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetRevokeOnUnenrollDisabled(v bool)`

SetRevokeOnUnenrollDisabled gets a reference to the given bool and assigns it to the RevokeOnUnenrollDisabled field.

### GetRightsManagementServicesTemplateId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetRightsManagementServicesTemplateId() string`

GetRightsManagementServicesTemplateId returns the RightsManagementServicesTemplateId field if non-nil, zero value otherwise.

### GetRightsManagementServicesTemplateIdOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetRightsManagementServicesTemplateIdOk() (string, bool)`

GetRightsManagementServicesTemplateIdOk returns a tuple with the RightsManagementServicesTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRightsManagementServicesTemplateId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasRightsManagementServicesTemplateId() bool`

HasRightsManagementServicesTemplateId returns a boolean if a field has been set.

### SetRightsManagementServicesTemplateId

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetRightsManagementServicesTemplateId(v string)`

SetRightsManagementServicesTemplateId gets a reference to the given string and assigns it to the RightsManagementServicesTemplateId field.

### SetRightsManagementServicesTemplateIdExplicitNull

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetRightsManagementServicesTemplateIdExplicitNull(b bool)`

SetRightsManagementServicesTemplateIdExplicitNull (un)sets RightsManagementServicesTemplateId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RightsManagementServicesTemplateId value is set to nil even if false is passed
### GetAzureRightsManagementServicesAllowed

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetAzureRightsManagementServicesAllowed() bool`

GetAzureRightsManagementServicesAllowed returns the AzureRightsManagementServicesAllowed field if non-nil, zero value otherwise.

### GetAzureRightsManagementServicesAllowedOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetAzureRightsManagementServicesAllowedOk() (bool, bool)`

GetAzureRightsManagementServicesAllowedOk returns a tuple with the AzureRightsManagementServicesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureRightsManagementServicesAllowed

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasAzureRightsManagementServicesAllowed() bool`

HasAzureRightsManagementServicesAllowed returns a boolean if a field has been set.

### SetAzureRightsManagementServicesAllowed

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetAzureRightsManagementServicesAllowed(v bool)`

SetAzureRightsManagementServicesAllowed gets a reference to the given bool and assigns it to the AzureRightsManagementServicesAllowed field.

### GetIconsVisible

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIconsVisible() bool`

GetIconsVisible returns the IconsVisible field if non-nil, zero value otherwise.

### GetIconsVisibleOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIconsVisibleOk() (bool, bool)`

GetIconsVisibleOk returns a tuple with the IconsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIconsVisible

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasIconsVisible() bool`

HasIconsVisible returns a boolean if a field has been set.

### SetIconsVisible

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetIconsVisible(v bool)`

SetIconsVisible gets a reference to the given bool and assigns it to the IconsVisible field.

### GetProtectedApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectedApps() []AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetProtectedApps returns the ProtectedApps field if non-nil, zero value otherwise.

### GetProtectedAppsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectedAppsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetProtectedAppsOk returns a tuple with the ProtectedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtectedApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasProtectedApps() bool`

HasProtectedApps returns a boolean if a field has been set.

### SetProtectedApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetProtectedApps(v []AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetProtectedApps gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionApp and assigns it to the ProtectedApps field.

### GetExemptApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetExemptApps() []AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetExemptApps returns the ExemptApps field if non-nil, zero value otherwise.

### GetExemptAppsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetExemptAppsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetExemptAppsOk returns a tuple with the ExemptApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExemptApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasExemptApps() bool`

HasExemptApps returns a boolean if a field has been set.

### SetExemptApps

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetExemptApps(v []AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetExemptApps gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionApp and assigns it to the ExemptApps field.

### GetEnterpriseNetworkDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseNetworkDomainNames() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseNetworkDomainNames returns the EnterpriseNetworkDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseNetworkDomainNamesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseNetworkDomainNamesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseNetworkDomainNamesOk returns a tuple with the EnterpriseNetworkDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseNetworkDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseNetworkDomainNames() bool`

HasEnterpriseNetworkDomainNames returns a boolean if a field has been set.

### SetEnterpriseNetworkDomainNames

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseNetworkDomainNames(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseNetworkDomainNames gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseNetworkDomainNames field.

### GetEnterpriseProxiedDomains

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxiedDomains() []AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection`

GetEnterpriseProxiedDomains returns the EnterpriseProxiedDomains field if non-nil, zero value otherwise.

### GetEnterpriseProxiedDomainsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxiedDomainsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection, bool)`

GetEnterpriseProxiedDomainsOk returns a tuple with the EnterpriseProxiedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProxiedDomains

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseProxiedDomains() bool`

HasEnterpriseProxiedDomains returns a boolean if a field has been set.

### SetEnterpriseProxiedDomains

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseProxiedDomains(v []AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection)`

SetEnterpriseProxiedDomains gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection and assigns it to the EnterpriseProxiedDomains field.

### GetEnterpriseIPRanges

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseIPRanges() []AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection`

GetEnterpriseIPRanges returns the EnterpriseIPRanges field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseIPRangesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection, bool)`

GetEnterpriseIPRangesOk returns a tuple with the EnterpriseIPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseIPRanges

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseIPRanges() bool`

HasEnterpriseIPRanges returns a boolean if a field has been set.

### SetEnterpriseIPRanges

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseIPRanges(v []AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection)`

SetEnterpriseIPRanges gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection and assigns it to the EnterpriseIPRanges field.

### GetEnterpriseIPRangesAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseIPRangesAreAuthoritative() bool`

GetEnterpriseIPRangesAreAuthoritative returns the EnterpriseIPRangesAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesAreAuthoritativeOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseIPRangesAreAuthoritativeOk() (bool, bool)`

GetEnterpriseIPRangesAreAuthoritativeOk returns a tuple with the EnterpriseIPRangesAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseIPRangesAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseIPRangesAreAuthoritative() bool`

HasEnterpriseIPRangesAreAuthoritative returns a boolean if a field has been set.

### SetEnterpriseIPRangesAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseIPRangesAreAuthoritative(v bool)`

SetEnterpriseIPRangesAreAuthoritative gets a reference to the given bool and assigns it to the EnterpriseIPRangesAreAuthoritative field.

### GetEnterpriseProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxyServers() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProxyServers returns the EnterpriseProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxyServersOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProxyServersOk returns a tuple with the EnterpriseProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseProxyServers() bool`

HasEnterpriseProxyServers returns a boolean if a field has been set.

### SetEnterpriseProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseProxyServers(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProxyServers gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseProxyServers field.

### GetEnterpriseInternalProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseInternalProxyServers() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseInternalProxyServers returns the EnterpriseInternalProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseInternalProxyServersOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseInternalProxyServersOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseInternalProxyServersOk returns a tuple with the EnterpriseInternalProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseInternalProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseInternalProxyServers() bool`

HasEnterpriseInternalProxyServers returns a boolean if a field has been set.

### SetEnterpriseInternalProxyServers

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseInternalProxyServers(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseInternalProxyServers gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseInternalProxyServers field.

### GetEnterpriseProxyServersAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxyServersAreAuthoritative() bool`

GetEnterpriseProxyServersAreAuthoritative returns the EnterpriseProxyServersAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersAreAuthoritativeOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetEnterpriseProxyServersAreAuthoritativeOk() (bool, bool)`

GetEnterpriseProxyServersAreAuthoritativeOk returns a tuple with the EnterpriseProxyServersAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProxyServersAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasEnterpriseProxyServersAreAuthoritative() bool`

HasEnterpriseProxyServersAreAuthoritative returns a boolean if a field has been set.

### SetEnterpriseProxyServersAreAuthoritative

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetEnterpriseProxyServersAreAuthoritative(v bool)`

SetEnterpriseProxyServersAreAuthoritative gets a reference to the given bool and assigns it to the EnterpriseProxyServersAreAuthoritative field.

### GetNeutralDomainResources

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetNeutralDomainResources() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetNeutralDomainResources returns the NeutralDomainResources field if non-nil, zero value otherwise.

### GetNeutralDomainResourcesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetNeutralDomainResourcesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetNeutralDomainResourcesOk returns a tuple with the NeutralDomainResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNeutralDomainResources

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasNeutralDomainResources() bool`

HasNeutralDomainResources returns a boolean if a field has been set.

### SetNeutralDomainResources

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetNeutralDomainResources(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetNeutralDomainResources gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the NeutralDomainResources field.

### GetIndexingEncryptedStoresOrItemsBlocked

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIndexingEncryptedStoresOrItemsBlocked() bool`

GetIndexingEncryptedStoresOrItemsBlocked returns the IndexingEncryptedStoresOrItemsBlocked field if non-nil, zero value otherwise.

### GetIndexingEncryptedStoresOrItemsBlockedOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIndexingEncryptedStoresOrItemsBlockedOk() (bool, bool)`

GetIndexingEncryptedStoresOrItemsBlockedOk returns a tuple with the IndexingEncryptedStoresOrItemsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndexingEncryptedStoresOrItemsBlocked

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasIndexingEncryptedStoresOrItemsBlocked() bool`

HasIndexingEncryptedStoresOrItemsBlocked returns a boolean if a field has been set.

### SetIndexingEncryptedStoresOrItemsBlocked

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetIndexingEncryptedStoresOrItemsBlocked(v bool)`

SetIndexingEncryptedStoresOrItemsBlocked gets a reference to the given bool and assigns it to the IndexingEncryptedStoresOrItemsBlocked field.

### GetSmbAutoEncryptedFileExtensions

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetSmbAutoEncryptedFileExtensions() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetSmbAutoEncryptedFileExtensions returns the SmbAutoEncryptedFileExtensions field if non-nil, zero value otherwise.

### GetSmbAutoEncryptedFileExtensionsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetSmbAutoEncryptedFileExtensionsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetSmbAutoEncryptedFileExtensionsOk returns a tuple with the SmbAutoEncryptedFileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmbAutoEncryptedFileExtensions

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasSmbAutoEncryptedFileExtensions() bool`

HasSmbAutoEncryptedFileExtensions returns a boolean if a field has been set.

### SetSmbAutoEncryptedFileExtensions

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetSmbAutoEncryptedFileExtensions(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetSmbAutoEncryptedFileExtensions gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the SmbAutoEncryptedFileExtensions field.

### GetIsAssigned

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetIsAssignedOk() (bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAssigned

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### SetIsAssigned

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetIsAssigned(v bool)`

SetIsAssigned gets a reference to the given bool and assigns it to the IsAssigned field.

### GetProtectedAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectedAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetProtectedAppLockerFiles returns the ProtectedAppLockerFiles field if non-nil, zero value otherwise.

### GetProtectedAppLockerFilesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetProtectedAppLockerFilesOk() ([]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetProtectedAppLockerFilesOk returns a tuple with the ProtectedAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtectedAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasProtectedAppLockerFiles() bool`

HasProtectedAppLockerFiles returns a boolean if a field has been set.

### SetProtectedAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetProtectedAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetProtectedAppLockerFiles gets a reference to the given []MicrosoftGraphWindowsInformationProtectionAppLockerFile and assigns it to the ProtectedAppLockerFiles field.

### GetExemptAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetExemptAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetExemptAppLockerFiles returns the ExemptAppLockerFiles field if non-nil, zero value otherwise.

### GetExemptAppLockerFilesOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetExemptAppLockerFilesOk() ([]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetExemptAppLockerFilesOk returns a tuple with the ExemptAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExemptAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasExemptAppLockerFiles() bool`

HasExemptAppLockerFiles returns a boolean if a field has been set.

### SetExemptAppLockerFiles

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetExemptAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetExemptAppLockerFiles gets a reference to the given []MicrosoftGraphWindowsInformationProtectionAppLockerFile and assigns it to the ExemptAppLockerFiles field.

### GetAssignments

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) GetAssignmentsOk() ([]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphMdmWindowsInformationProtectionPolicy) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



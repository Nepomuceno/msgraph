# WindowsInformationProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetEnforcementLevel

`func (o *WindowsInformationProtection) GetEnforcementLevel() AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel`

GetEnforcementLevel returns the EnforcementLevel field if non-nil, zero value otherwise.

### GetEnforcementLevelOk

`func (o *WindowsInformationProtection) GetEnforcementLevelOk() (AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel, bool)`

GetEnforcementLevelOk returns a tuple with the EnforcementLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnforcementLevel

`func (o *WindowsInformationProtection) HasEnforcementLevel() bool`

HasEnforcementLevel returns a boolean if a field has been set.

### SetEnforcementLevel

`func (o *WindowsInformationProtection) SetEnforcementLevel(v AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel)`

SetEnforcementLevel gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionEnforcementLevel and assigns it to the EnforcementLevel field.

### GetEnterpriseDomain

`func (o *WindowsInformationProtection) GetEnterpriseDomain() string`

GetEnterpriseDomain returns the EnterpriseDomain field if non-nil, zero value otherwise.

### GetEnterpriseDomainOk

`func (o *WindowsInformationProtection) GetEnterpriseDomainOk() (string, bool)`

GetEnterpriseDomainOk returns a tuple with the EnterpriseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseDomain

`func (o *WindowsInformationProtection) HasEnterpriseDomain() bool`

HasEnterpriseDomain returns a boolean if a field has been set.

### SetEnterpriseDomain

`func (o *WindowsInformationProtection) SetEnterpriseDomain(v string)`

SetEnterpriseDomain gets a reference to the given string and assigns it to the EnterpriseDomain field.

### SetEnterpriseDomainExplicitNull

`func (o *WindowsInformationProtection) SetEnterpriseDomainExplicitNull(b bool)`

SetEnterpriseDomainExplicitNull (un)sets EnterpriseDomain to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EnterpriseDomain value is set to nil even if false is passed
### GetEnterpriseProtectedDomainNames

`func (o *WindowsInformationProtection) GetEnterpriseProtectedDomainNames() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProtectedDomainNames returns the EnterpriseProtectedDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseProtectedDomainNamesOk

`func (o *WindowsInformationProtection) GetEnterpriseProtectedDomainNamesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProtectedDomainNamesOk returns a tuple with the EnterpriseProtectedDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProtectedDomainNames

`func (o *WindowsInformationProtection) HasEnterpriseProtectedDomainNames() bool`

HasEnterpriseProtectedDomainNames returns a boolean if a field has been set.

### SetEnterpriseProtectedDomainNames

`func (o *WindowsInformationProtection) SetEnterpriseProtectedDomainNames(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProtectedDomainNames gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseProtectedDomainNames field.

### GetProtectionUnderLockConfigRequired

`func (o *WindowsInformationProtection) GetProtectionUnderLockConfigRequired() bool`

GetProtectionUnderLockConfigRequired returns the ProtectionUnderLockConfigRequired field if non-nil, zero value otherwise.

### GetProtectionUnderLockConfigRequiredOk

`func (o *WindowsInformationProtection) GetProtectionUnderLockConfigRequiredOk() (bool, bool)`

GetProtectionUnderLockConfigRequiredOk returns a tuple with the ProtectionUnderLockConfigRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtectionUnderLockConfigRequired

`func (o *WindowsInformationProtection) HasProtectionUnderLockConfigRequired() bool`

HasProtectionUnderLockConfigRequired returns a boolean if a field has been set.

### SetProtectionUnderLockConfigRequired

`func (o *WindowsInformationProtection) SetProtectionUnderLockConfigRequired(v bool)`

SetProtectionUnderLockConfigRequired gets a reference to the given bool and assigns it to the ProtectionUnderLockConfigRequired field.

### GetDataRecoveryCertificate

`func (o *WindowsInformationProtection) GetDataRecoveryCertificate() AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate`

GetDataRecoveryCertificate returns the DataRecoveryCertificate field if non-nil, zero value otherwise.

### GetDataRecoveryCertificateOk

`func (o *WindowsInformationProtection) GetDataRecoveryCertificateOk() (AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate, bool)`

GetDataRecoveryCertificateOk returns a tuple with the DataRecoveryCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDataRecoveryCertificate

`func (o *WindowsInformationProtection) HasDataRecoveryCertificate() bool`

HasDataRecoveryCertificate returns a boolean if a field has been set.

### SetDataRecoveryCertificate

`func (o *WindowsInformationProtection) SetDataRecoveryCertificate(v AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate)`

SetDataRecoveryCertificate gets a reference to the given AnyOfmicrosoftGraphWindowsInformationProtectionDataRecoveryCertificate and assigns it to the DataRecoveryCertificate field.

### SetDataRecoveryCertificateExplicitNull

`func (o *WindowsInformationProtection) SetDataRecoveryCertificateExplicitNull(b bool)`

SetDataRecoveryCertificateExplicitNull (un)sets DataRecoveryCertificate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DataRecoveryCertificate value is set to nil even if false is passed
### GetRevokeOnUnenrollDisabled

`func (o *WindowsInformationProtection) GetRevokeOnUnenrollDisabled() bool`

GetRevokeOnUnenrollDisabled returns the RevokeOnUnenrollDisabled field if non-nil, zero value otherwise.

### GetRevokeOnUnenrollDisabledOk

`func (o *WindowsInformationProtection) GetRevokeOnUnenrollDisabledOk() (bool, bool)`

GetRevokeOnUnenrollDisabledOk returns a tuple with the RevokeOnUnenrollDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRevokeOnUnenrollDisabled

`func (o *WindowsInformationProtection) HasRevokeOnUnenrollDisabled() bool`

HasRevokeOnUnenrollDisabled returns a boolean if a field has been set.

### SetRevokeOnUnenrollDisabled

`func (o *WindowsInformationProtection) SetRevokeOnUnenrollDisabled(v bool)`

SetRevokeOnUnenrollDisabled gets a reference to the given bool and assigns it to the RevokeOnUnenrollDisabled field.

### GetRightsManagementServicesTemplateId

`func (o *WindowsInformationProtection) GetRightsManagementServicesTemplateId() string`

GetRightsManagementServicesTemplateId returns the RightsManagementServicesTemplateId field if non-nil, zero value otherwise.

### GetRightsManagementServicesTemplateIdOk

`func (o *WindowsInformationProtection) GetRightsManagementServicesTemplateIdOk() (string, bool)`

GetRightsManagementServicesTemplateIdOk returns a tuple with the RightsManagementServicesTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRightsManagementServicesTemplateId

`func (o *WindowsInformationProtection) HasRightsManagementServicesTemplateId() bool`

HasRightsManagementServicesTemplateId returns a boolean if a field has been set.

### SetRightsManagementServicesTemplateId

`func (o *WindowsInformationProtection) SetRightsManagementServicesTemplateId(v string)`

SetRightsManagementServicesTemplateId gets a reference to the given string and assigns it to the RightsManagementServicesTemplateId field.

### SetRightsManagementServicesTemplateIdExplicitNull

`func (o *WindowsInformationProtection) SetRightsManagementServicesTemplateIdExplicitNull(b bool)`

SetRightsManagementServicesTemplateIdExplicitNull (un)sets RightsManagementServicesTemplateId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RightsManagementServicesTemplateId value is set to nil even if false is passed
### GetAzureRightsManagementServicesAllowed

`func (o *WindowsInformationProtection) GetAzureRightsManagementServicesAllowed() bool`

GetAzureRightsManagementServicesAllowed returns the AzureRightsManagementServicesAllowed field if non-nil, zero value otherwise.

### GetAzureRightsManagementServicesAllowedOk

`func (o *WindowsInformationProtection) GetAzureRightsManagementServicesAllowedOk() (bool, bool)`

GetAzureRightsManagementServicesAllowedOk returns a tuple with the AzureRightsManagementServicesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAzureRightsManagementServicesAllowed

`func (o *WindowsInformationProtection) HasAzureRightsManagementServicesAllowed() bool`

HasAzureRightsManagementServicesAllowed returns a boolean if a field has been set.

### SetAzureRightsManagementServicesAllowed

`func (o *WindowsInformationProtection) SetAzureRightsManagementServicesAllowed(v bool)`

SetAzureRightsManagementServicesAllowed gets a reference to the given bool and assigns it to the AzureRightsManagementServicesAllowed field.

### GetIconsVisible

`func (o *WindowsInformationProtection) GetIconsVisible() bool`

GetIconsVisible returns the IconsVisible field if non-nil, zero value otherwise.

### GetIconsVisibleOk

`func (o *WindowsInformationProtection) GetIconsVisibleOk() (bool, bool)`

GetIconsVisibleOk returns a tuple with the IconsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIconsVisible

`func (o *WindowsInformationProtection) HasIconsVisible() bool`

HasIconsVisible returns a boolean if a field has been set.

### SetIconsVisible

`func (o *WindowsInformationProtection) SetIconsVisible(v bool)`

SetIconsVisible gets a reference to the given bool and assigns it to the IconsVisible field.

### GetProtectedApps

`func (o *WindowsInformationProtection) GetProtectedApps() []AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetProtectedApps returns the ProtectedApps field if non-nil, zero value otherwise.

### GetProtectedAppsOk

`func (o *WindowsInformationProtection) GetProtectedAppsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetProtectedAppsOk returns a tuple with the ProtectedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtectedApps

`func (o *WindowsInformationProtection) HasProtectedApps() bool`

HasProtectedApps returns a boolean if a field has been set.

### SetProtectedApps

`func (o *WindowsInformationProtection) SetProtectedApps(v []AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetProtectedApps gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionApp and assigns it to the ProtectedApps field.

### GetExemptApps

`func (o *WindowsInformationProtection) GetExemptApps() []AnyOfmicrosoftGraphWindowsInformationProtectionApp`

GetExemptApps returns the ExemptApps field if non-nil, zero value otherwise.

### GetExemptAppsOk

`func (o *WindowsInformationProtection) GetExemptAppsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionApp, bool)`

GetExemptAppsOk returns a tuple with the ExemptApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExemptApps

`func (o *WindowsInformationProtection) HasExemptApps() bool`

HasExemptApps returns a boolean if a field has been set.

### SetExemptApps

`func (o *WindowsInformationProtection) SetExemptApps(v []AnyOfmicrosoftGraphWindowsInformationProtectionApp)`

SetExemptApps gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionApp and assigns it to the ExemptApps field.

### GetEnterpriseNetworkDomainNames

`func (o *WindowsInformationProtection) GetEnterpriseNetworkDomainNames() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseNetworkDomainNames returns the EnterpriseNetworkDomainNames field if non-nil, zero value otherwise.

### GetEnterpriseNetworkDomainNamesOk

`func (o *WindowsInformationProtection) GetEnterpriseNetworkDomainNamesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseNetworkDomainNamesOk returns a tuple with the EnterpriseNetworkDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseNetworkDomainNames

`func (o *WindowsInformationProtection) HasEnterpriseNetworkDomainNames() bool`

HasEnterpriseNetworkDomainNames returns a boolean if a field has been set.

### SetEnterpriseNetworkDomainNames

`func (o *WindowsInformationProtection) SetEnterpriseNetworkDomainNames(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseNetworkDomainNames gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseNetworkDomainNames field.

### GetEnterpriseProxiedDomains

`func (o *WindowsInformationProtection) GetEnterpriseProxiedDomains() []AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection`

GetEnterpriseProxiedDomains returns the EnterpriseProxiedDomains field if non-nil, zero value otherwise.

### GetEnterpriseProxiedDomainsOk

`func (o *WindowsInformationProtection) GetEnterpriseProxiedDomainsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection, bool)`

GetEnterpriseProxiedDomainsOk returns a tuple with the EnterpriseProxiedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProxiedDomains

`func (o *WindowsInformationProtection) HasEnterpriseProxiedDomains() bool`

HasEnterpriseProxiedDomains returns a boolean if a field has been set.

### SetEnterpriseProxiedDomains

`func (o *WindowsInformationProtection) SetEnterpriseProxiedDomains(v []AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection)`

SetEnterpriseProxiedDomains gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionProxiedDomainCollection and assigns it to the EnterpriseProxiedDomains field.

### GetEnterpriseIPRanges

`func (o *WindowsInformationProtection) GetEnterpriseIPRanges() []AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection`

GetEnterpriseIPRanges returns the EnterpriseIPRanges field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesOk

`func (o *WindowsInformationProtection) GetEnterpriseIPRangesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection, bool)`

GetEnterpriseIPRangesOk returns a tuple with the EnterpriseIPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseIPRanges

`func (o *WindowsInformationProtection) HasEnterpriseIPRanges() bool`

HasEnterpriseIPRanges returns a boolean if a field has been set.

### SetEnterpriseIPRanges

`func (o *WindowsInformationProtection) SetEnterpriseIPRanges(v []AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection)`

SetEnterpriseIPRanges gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionIpRangeCollection and assigns it to the EnterpriseIPRanges field.

### GetEnterpriseIPRangesAreAuthoritative

`func (o *WindowsInformationProtection) GetEnterpriseIPRangesAreAuthoritative() bool`

GetEnterpriseIPRangesAreAuthoritative returns the EnterpriseIPRangesAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseIPRangesAreAuthoritativeOk

`func (o *WindowsInformationProtection) GetEnterpriseIPRangesAreAuthoritativeOk() (bool, bool)`

GetEnterpriseIPRangesAreAuthoritativeOk returns a tuple with the EnterpriseIPRangesAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseIPRangesAreAuthoritative

`func (o *WindowsInformationProtection) HasEnterpriseIPRangesAreAuthoritative() bool`

HasEnterpriseIPRangesAreAuthoritative returns a boolean if a field has been set.

### SetEnterpriseIPRangesAreAuthoritative

`func (o *WindowsInformationProtection) SetEnterpriseIPRangesAreAuthoritative(v bool)`

SetEnterpriseIPRangesAreAuthoritative gets a reference to the given bool and assigns it to the EnterpriseIPRangesAreAuthoritative field.

### GetEnterpriseProxyServers

`func (o *WindowsInformationProtection) GetEnterpriseProxyServers() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseProxyServers returns the EnterpriseProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersOk

`func (o *WindowsInformationProtection) GetEnterpriseProxyServersOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseProxyServersOk returns a tuple with the EnterpriseProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProxyServers

`func (o *WindowsInformationProtection) HasEnterpriseProxyServers() bool`

HasEnterpriseProxyServers returns a boolean if a field has been set.

### SetEnterpriseProxyServers

`func (o *WindowsInformationProtection) SetEnterpriseProxyServers(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseProxyServers gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseProxyServers field.

### GetEnterpriseInternalProxyServers

`func (o *WindowsInformationProtection) GetEnterpriseInternalProxyServers() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetEnterpriseInternalProxyServers returns the EnterpriseInternalProxyServers field if non-nil, zero value otherwise.

### GetEnterpriseInternalProxyServersOk

`func (o *WindowsInformationProtection) GetEnterpriseInternalProxyServersOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetEnterpriseInternalProxyServersOk returns a tuple with the EnterpriseInternalProxyServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseInternalProxyServers

`func (o *WindowsInformationProtection) HasEnterpriseInternalProxyServers() bool`

HasEnterpriseInternalProxyServers returns a boolean if a field has been set.

### SetEnterpriseInternalProxyServers

`func (o *WindowsInformationProtection) SetEnterpriseInternalProxyServers(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetEnterpriseInternalProxyServers gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the EnterpriseInternalProxyServers field.

### GetEnterpriseProxyServersAreAuthoritative

`func (o *WindowsInformationProtection) GetEnterpriseProxyServersAreAuthoritative() bool`

GetEnterpriseProxyServersAreAuthoritative returns the EnterpriseProxyServersAreAuthoritative field if non-nil, zero value otherwise.

### GetEnterpriseProxyServersAreAuthoritativeOk

`func (o *WindowsInformationProtection) GetEnterpriseProxyServersAreAuthoritativeOk() (bool, bool)`

GetEnterpriseProxyServersAreAuthoritativeOk returns a tuple with the EnterpriseProxyServersAreAuthoritative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnterpriseProxyServersAreAuthoritative

`func (o *WindowsInformationProtection) HasEnterpriseProxyServersAreAuthoritative() bool`

HasEnterpriseProxyServersAreAuthoritative returns a boolean if a field has been set.

### SetEnterpriseProxyServersAreAuthoritative

`func (o *WindowsInformationProtection) SetEnterpriseProxyServersAreAuthoritative(v bool)`

SetEnterpriseProxyServersAreAuthoritative gets a reference to the given bool and assigns it to the EnterpriseProxyServersAreAuthoritative field.

### GetNeutralDomainResources

`func (o *WindowsInformationProtection) GetNeutralDomainResources() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetNeutralDomainResources returns the NeutralDomainResources field if non-nil, zero value otherwise.

### GetNeutralDomainResourcesOk

`func (o *WindowsInformationProtection) GetNeutralDomainResourcesOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetNeutralDomainResourcesOk returns a tuple with the NeutralDomainResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNeutralDomainResources

`func (o *WindowsInformationProtection) HasNeutralDomainResources() bool`

HasNeutralDomainResources returns a boolean if a field has been set.

### SetNeutralDomainResources

`func (o *WindowsInformationProtection) SetNeutralDomainResources(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetNeutralDomainResources gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the NeutralDomainResources field.

### GetIndexingEncryptedStoresOrItemsBlocked

`func (o *WindowsInformationProtection) GetIndexingEncryptedStoresOrItemsBlocked() bool`

GetIndexingEncryptedStoresOrItemsBlocked returns the IndexingEncryptedStoresOrItemsBlocked field if non-nil, zero value otherwise.

### GetIndexingEncryptedStoresOrItemsBlockedOk

`func (o *WindowsInformationProtection) GetIndexingEncryptedStoresOrItemsBlockedOk() (bool, bool)`

GetIndexingEncryptedStoresOrItemsBlockedOk returns a tuple with the IndexingEncryptedStoresOrItemsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndexingEncryptedStoresOrItemsBlocked

`func (o *WindowsInformationProtection) HasIndexingEncryptedStoresOrItemsBlocked() bool`

HasIndexingEncryptedStoresOrItemsBlocked returns a boolean if a field has been set.

### SetIndexingEncryptedStoresOrItemsBlocked

`func (o *WindowsInformationProtection) SetIndexingEncryptedStoresOrItemsBlocked(v bool)`

SetIndexingEncryptedStoresOrItemsBlocked gets a reference to the given bool and assigns it to the IndexingEncryptedStoresOrItemsBlocked field.

### GetSmbAutoEncryptedFileExtensions

`func (o *WindowsInformationProtection) GetSmbAutoEncryptedFileExtensions() []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection`

GetSmbAutoEncryptedFileExtensions returns the SmbAutoEncryptedFileExtensions field if non-nil, zero value otherwise.

### GetSmbAutoEncryptedFileExtensionsOk

`func (o *WindowsInformationProtection) GetSmbAutoEncryptedFileExtensionsOk() ([]AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection, bool)`

GetSmbAutoEncryptedFileExtensionsOk returns a tuple with the SmbAutoEncryptedFileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSmbAutoEncryptedFileExtensions

`func (o *WindowsInformationProtection) HasSmbAutoEncryptedFileExtensions() bool`

HasSmbAutoEncryptedFileExtensions returns a boolean if a field has been set.

### SetSmbAutoEncryptedFileExtensions

`func (o *WindowsInformationProtection) SetSmbAutoEncryptedFileExtensions(v []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection)`

SetSmbAutoEncryptedFileExtensions gets a reference to the given []AnyOfmicrosoftGraphWindowsInformationProtectionResourceCollection and assigns it to the SmbAutoEncryptedFileExtensions field.

### GetIsAssigned

`func (o *WindowsInformationProtection) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *WindowsInformationProtection) GetIsAssignedOk() (bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAssigned

`func (o *WindowsInformationProtection) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### SetIsAssigned

`func (o *WindowsInformationProtection) SetIsAssigned(v bool)`

SetIsAssigned gets a reference to the given bool and assigns it to the IsAssigned field.

### GetProtectedAppLockerFiles

`func (o *WindowsInformationProtection) GetProtectedAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetProtectedAppLockerFiles returns the ProtectedAppLockerFiles field if non-nil, zero value otherwise.

### GetProtectedAppLockerFilesOk

`func (o *WindowsInformationProtection) GetProtectedAppLockerFilesOk() ([]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetProtectedAppLockerFilesOk returns a tuple with the ProtectedAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtectedAppLockerFiles

`func (o *WindowsInformationProtection) HasProtectedAppLockerFiles() bool`

HasProtectedAppLockerFiles returns a boolean if a field has been set.

### SetProtectedAppLockerFiles

`func (o *WindowsInformationProtection) SetProtectedAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetProtectedAppLockerFiles gets a reference to the given []MicrosoftGraphWindowsInformationProtectionAppLockerFile and assigns it to the ProtectedAppLockerFiles field.

### GetExemptAppLockerFiles

`func (o *WindowsInformationProtection) GetExemptAppLockerFiles() []MicrosoftGraphWindowsInformationProtectionAppLockerFile`

GetExemptAppLockerFiles returns the ExemptAppLockerFiles field if non-nil, zero value otherwise.

### GetExemptAppLockerFilesOk

`func (o *WindowsInformationProtection) GetExemptAppLockerFilesOk() ([]MicrosoftGraphWindowsInformationProtectionAppLockerFile, bool)`

GetExemptAppLockerFilesOk returns a tuple with the ExemptAppLockerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExemptAppLockerFiles

`func (o *WindowsInformationProtection) HasExemptAppLockerFiles() bool`

HasExemptAppLockerFiles returns a boolean if a field has been set.

### SetExemptAppLockerFiles

`func (o *WindowsInformationProtection) SetExemptAppLockerFiles(v []MicrosoftGraphWindowsInformationProtectionAppLockerFile)`

SetExemptAppLockerFiles gets a reference to the given []MicrosoftGraphWindowsInformationProtectionAppLockerFile and assigns it to the ExemptAppLockerFiles field.

### GetAssignments

`func (o *WindowsInformationProtection) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *WindowsInformationProtection) GetAssignmentsOk() ([]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *WindowsInformationProtection) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *WindowsInformationProtection) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



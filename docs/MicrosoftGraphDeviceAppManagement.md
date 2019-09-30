# MicrosoftGraphDeviceAppManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MicrosoftStoreForBusinessLastSuccessfulSyncDateTime** | Pointer to [**time.Time**](time.Time.md) | The last time the apps from the Microsoft Store for Business were synced successfully for the account. | [optional] 
**IsEnabledForMicrosoftStoreForBusiness** | Pointer to **bool** | Whether the account is enabled for syncing applications from the Microsoft Store for Business. | [optional] 
**MicrosoftStoreForBusinessLanguage** | Pointer to **string** | The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is &lt;languagecode2&gt;-&lt;country/regioncode2&gt;, where &lt;languagecode2&gt; is a lowercase two-letter code derived from ISO 639-1 and &lt;country/regioncode2&gt; is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture. | [optional] 
**MicrosoftStoreForBusinessLastCompletedApplicationSyncTime** | Pointer to [**time.Time**](time.Time.md) | The last time an application sync from the Microsoft Store for Business was completed. | [optional] 
**ManagedEBooks** | Pointer to [**[]MicrosoftGraphManagedEBook**](microsoft.graph.managedEBook.md) |  | [optional] 
**MobileApps** | Pointer to [**[]MicrosoftGraphMobileApp**](microsoft.graph.mobileApp.md) |  | [optional] 
**MobileAppCategories** | Pointer to [**[]MicrosoftGraphMobileAppCategory**](microsoft.graph.mobileAppCategory.md) |  | [optional] 
**MobileAppConfigurations** | Pointer to [**[]MicrosoftGraphManagedDeviceMobileAppConfiguration**](microsoft.graph.managedDeviceMobileAppConfiguration.md) |  | [optional] 
**VppTokens** | Pointer to [**[]MicrosoftGraphVppToken**](microsoft.graph.vppToken.md) |  | [optional] 
**ManagedAppPolicies** | Pointer to [**[]MicrosoftGraphManagedAppPolicy**](microsoft.graph.managedAppPolicy.md) |  | [optional] 
**IosManagedAppProtections** | Pointer to [**[]MicrosoftGraphIosManagedAppProtection**](microsoft.graph.iosManagedAppProtection.md) |  | [optional] 
**AndroidManagedAppProtections** | Pointer to [**[]MicrosoftGraphAndroidManagedAppProtection**](microsoft.graph.androidManagedAppProtection.md) |  | [optional] 
**DefaultManagedAppProtections** | Pointer to [**[]MicrosoftGraphDefaultManagedAppProtection**](microsoft.graph.defaultManagedAppProtection.md) |  | [optional] 
**TargetedManagedAppConfigurations** | Pointer to [**[]MicrosoftGraphTargetedManagedAppConfiguration**](microsoft.graph.targetedManagedAppConfiguration.md) |  | [optional] 
**MdmWindowsInformationProtectionPolicies** | Pointer to [**[]MicrosoftGraphMdmWindowsInformationProtectionPolicy**](microsoft.graph.mdmWindowsInformationProtectionPolicy.md) |  | [optional] 
**WindowsInformationProtectionPolicies** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionPolicy**](microsoft.graph.windowsInformationProtectionPolicy.md) |  | [optional] 
**ManagedAppRegistrations** | Pointer to [**[]MicrosoftGraphManagedAppRegistration**](microsoft.graph.managedAppRegistration.md) |  | [optional] 
**ManagedAppStatuses** | Pointer to [**[]MicrosoftGraphManagedAppStatus**](microsoft.graph.managedAppStatus.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceAppManagement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceAppManagement) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceAppManagement) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceAppManagement) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime() time.Time`

GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime returns the MicrosoftStoreForBusinessLastSuccessfulSyncDateTime field if non-nil, zero value otherwise.

### GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTimeOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTimeOk() (time.Time, bool)`

GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTimeOk returns a tuple with the MicrosoftStoreForBusinessLastSuccessfulSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftStoreForBusinessLastSuccessfulSyncDateTime

`func (o *MicrosoftGraphDeviceAppManagement) HasMicrosoftStoreForBusinessLastSuccessfulSyncDateTime() bool`

HasMicrosoftStoreForBusinessLastSuccessfulSyncDateTime returns a boolean if a field has been set.

### SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime

`func (o *MicrosoftGraphDeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(v time.Time)`

SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime gets a reference to the given time.Time and assigns it to the MicrosoftStoreForBusinessLastSuccessfulSyncDateTime field.

### GetIsEnabledForMicrosoftStoreForBusiness

`func (o *MicrosoftGraphDeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness() bool`

GetIsEnabledForMicrosoftStoreForBusiness returns the IsEnabledForMicrosoftStoreForBusiness field if non-nil, zero value otherwise.

### GetIsEnabledForMicrosoftStoreForBusinessOk

`func (o *MicrosoftGraphDeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusinessOk() (bool, bool)`

GetIsEnabledForMicrosoftStoreForBusinessOk returns a tuple with the IsEnabledForMicrosoftStoreForBusiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEnabledForMicrosoftStoreForBusiness

`func (o *MicrosoftGraphDeviceAppManagement) HasIsEnabledForMicrosoftStoreForBusiness() bool`

HasIsEnabledForMicrosoftStoreForBusiness returns a boolean if a field has been set.

### SetIsEnabledForMicrosoftStoreForBusiness

`func (o *MicrosoftGraphDeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(v bool)`

SetIsEnabledForMicrosoftStoreForBusiness gets a reference to the given bool and assigns it to the IsEnabledForMicrosoftStoreForBusiness field.

### GetMicrosoftStoreForBusinessLanguage

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLanguage() string`

GetMicrosoftStoreForBusinessLanguage returns the MicrosoftStoreForBusinessLanguage field if non-nil, zero value otherwise.

### GetMicrosoftStoreForBusinessLanguageOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLanguageOk() (string, bool)`

GetMicrosoftStoreForBusinessLanguageOk returns a tuple with the MicrosoftStoreForBusinessLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftStoreForBusinessLanguage

`func (o *MicrosoftGraphDeviceAppManagement) HasMicrosoftStoreForBusinessLanguage() bool`

HasMicrosoftStoreForBusinessLanguage returns a boolean if a field has been set.

### SetMicrosoftStoreForBusinessLanguage

`func (o *MicrosoftGraphDeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(v string)`

SetMicrosoftStoreForBusinessLanguage gets a reference to the given string and assigns it to the MicrosoftStoreForBusinessLanguage field.

### SetMicrosoftStoreForBusinessLanguageExplicitNull

`func (o *MicrosoftGraphDeviceAppManagement) SetMicrosoftStoreForBusinessLanguageExplicitNull(b bool)`

SetMicrosoftStoreForBusinessLanguageExplicitNull (un)sets MicrosoftStoreForBusinessLanguage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MicrosoftStoreForBusinessLanguage value is set to nil even if false is passed
### GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime() time.Time`

GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime returns the MicrosoftStoreForBusinessLastCompletedApplicationSyncTime field if non-nil, zero value otherwise.

### GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTimeOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTimeOk() (time.Time, bool)`

GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTimeOk returns a tuple with the MicrosoftStoreForBusinessLastCompletedApplicationSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftStoreForBusinessLastCompletedApplicationSyncTime

`func (o *MicrosoftGraphDeviceAppManagement) HasMicrosoftStoreForBusinessLastCompletedApplicationSyncTime() bool`

HasMicrosoftStoreForBusinessLastCompletedApplicationSyncTime returns a boolean if a field has been set.

### SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime

`func (o *MicrosoftGraphDeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(v time.Time)`

SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime gets a reference to the given time.Time and assigns it to the MicrosoftStoreForBusinessLastCompletedApplicationSyncTime field.

### GetManagedEBooks

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedEBooks() []MicrosoftGraphManagedEBook`

GetManagedEBooks returns the ManagedEBooks field if non-nil, zero value otherwise.

### GetManagedEBooksOk

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedEBooksOk() ([]MicrosoftGraphManagedEBook, bool)`

GetManagedEBooksOk returns a tuple with the ManagedEBooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedEBooks

`func (o *MicrosoftGraphDeviceAppManagement) HasManagedEBooks() bool`

HasManagedEBooks returns a boolean if a field has been set.

### SetManagedEBooks

`func (o *MicrosoftGraphDeviceAppManagement) SetManagedEBooks(v []MicrosoftGraphManagedEBook)`

SetManagedEBooks gets a reference to the given []MicrosoftGraphManagedEBook and assigns it to the ManagedEBooks field.

### GetMobileApps

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileApps() []MicrosoftGraphMobileApp`

GetMobileApps returns the MobileApps field if non-nil, zero value otherwise.

### GetMobileAppsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppsOk() ([]MicrosoftGraphMobileApp, bool)`

GetMobileAppsOk returns a tuple with the MobileApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileApps

`func (o *MicrosoftGraphDeviceAppManagement) HasMobileApps() bool`

HasMobileApps returns a boolean if a field has been set.

### SetMobileApps

`func (o *MicrosoftGraphDeviceAppManagement) SetMobileApps(v []MicrosoftGraphMobileApp)`

SetMobileApps gets a reference to the given []MicrosoftGraphMobileApp and assigns it to the MobileApps field.

### GetMobileAppCategories

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppCategories() []MicrosoftGraphMobileAppCategory`

GetMobileAppCategories returns the MobileAppCategories field if non-nil, zero value otherwise.

### GetMobileAppCategoriesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppCategoriesOk() ([]MicrosoftGraphMobileAppCategory, bool)`

GetMobileAppCategoriesOk returns a tuple with the MobileAppCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileAppCategories

`func (o *MicrosoftGraphDeviceAppManagement) HasMobileAppCategories() bool`

HasMobileAppCategories returns a boolean if a field has been set.

### SetMobileAppCategories

`func (o *MicrosoftGraphDeviceAppManagement) SetMobileAppCategories(v []MicrosoftGraphMobileAppCategory)`

SetMobileAppCategories gets a reference to the given []MicrosoftGraphMobileAppCategory and assigns it to the MobileAppCategories field.

### GetMobileAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppConfigurations() []MicrosoftGraphManagedDeviceMobileAppConfiguration`

GetMobileAppConfigurations returns the MobileAppConfigurations field if non-nil, zero value otherwise.

### GetMobileAppConfigurationsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMobileAppConfigurationsOk() ([]MicrosoftGraphManagedDeviceMobileAppConfiguration, bool)`

GetMobileAppConfigurationsOk returns a tuple with the MobileAppConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) HasMobileAppConfigurations() bool`

HasMobileAppConfigurations returns a boolean if a field has been set.

### SetMobileAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) SetMobileAppConfigurations(v []MicrosoftGraphManagedDeviceMobileAppConfiguration)`

SetMobileAppConfigurations gets a reference to the given []MicrosoftGraphManagedDeviceMobileAppConfiguration and assigns it to the MobileAppConfigurations field.

### GetVppTokens

`func (o *MicrosoftGraphDeviceAppManagement) GetVppTokens() []MicrosoftGraphVppToken`

GetVppTokens returns the VppTokens field if non-nil, zero value otherwise.

### GetVppTokensOk

`func (o *MicrosoftGraphDeviceAppManagement) GetVppTokensOk() ([]MicrosoftGraphVppToken, bool)`

GetVppTokensOk returns a tuple with the VppTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVppTokens

`func (o *MicrosoftGraphDeviceAppManagement) HasVppTokens() bool`

HasVppTokens returns a boolean if a field has been set.

### SetVppTokens

`func (o *MicrosoftGraphDeviceAppManagement) SetVppTokens(v []MicrosoftGraphVppToken)`

SetVppTokens gets a reference to the given []MicrosoftGraphVppToken and assigns it to the VppTokens field.

### GetManagedAppPolicies

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppPolicies() []MicrosoftGraphManagedAppPolicy`

GetManagedAppPolicies returns the ManagedAppPolicies field if non-nil, zero value otherwise.

### GetManagedAppPoliciesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppPoliciesOk() ([]MicrosoftGraphManagedAppPolicy, bool)`

GetManagedAppPoliciesOk returns a tuple with the ManagedAppPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedAppPolicies

`func (o *MicrosoftGraphDeviceAppManagement) HasManagedAppPolicies() bool`

HasManagedAppPolicies returns a boolean if a field has been set.

### SetManagedAppPolicies

`func (o *MicrosoftGraphDeviceAppManagement) SetManagedAppPolicies(v []MicrosoftGraphManagedAppPolicy)`

SetManagedAppPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the ManagedAppPolicies field.

### GetIosManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) GetIosManagedAppProtections() []MicrosoftGraphIosManagedAppProtection`

GetIosManagedAppProtections returns the IosManagedAppProtections field if non-nil, zero value otherwise.

### GetIosManagedAppProtectionsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetIosManagedAppProtectionsOk() ([]MicrosoftGraphIosManagedAppProtection, bool)`

GetIosManagedAppProtectionsOk returns a tuple with the IosManagedAppProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIosManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) HasIosManagedAppProtections() bool`

HasIosManagedAppProtections returns a boolean if a field has been set.

### SetIosManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) SetIosManagedAppProtections(v []MicrosoftGraphIosManagedAppProtection)`

SetIosManagedAppProtections gets a reference to the given []MicrosoftGraphIosManagedAppProtection and assigns it to the IosManagedAppProtections field.

### GetAndroidManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) GetAndroidManagedAppProtections() []MicrosoftGraphAndroidManagedAppProtection`

GetAndroidManagedAppProtections returns the AndroidManagedAppProtections field if non-nil, zero value otherwise.

### GetAndroidManagedAppProtectionsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetAndroidManagedAppProtectionsOk() ([]MicrosoftGraphAndroidManagedAppProtection, bool)`

GetAndroidManagedAppProtectionsOk returns a tuple with the AndroidManagedAppProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAndroidManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) HasAndroidManagedAppProtections() bool`

HasAndroidManagedAppProtections returns a boolean if a field has been set.

### SetAndroidManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) SetAndroidManagedAppProtections(v []MicrosoftGraphAndroidManagedAppProtection)`

SetAndroidManagedAppProtections gets a reference to the given []MicrosoftGraphAndroidManagedAppProtection and assigns it to the AndroidManagedAppProtections field.

### GetDefaultManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) GetDefaultManagedAppProtections() []MicrosoftGraphDefaultManagedAppProtection`

GetDefaultManagedAppProtections returns the DefaultManagedAppProtections field if non-nil, zero value otherwise.

### GetDefaultManagedAppProtectionsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetDefaultManagedAppProtectionsOk() ([]MicrosoftGraphDefaultManagedAppProtection, bool)`

GetDefaultManagedAppProtectionsOk returns a tuple with the DefaultManagedAppProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) HasDefaultManagedAppProtections() bool`

HasDefaultManagedAppProtections returns a boolean if a field has been set.

### SetDefaultManagedAppProtections

`func (o *MicrosoftGraphDeviceAppManagement) SetDefaultManagedAppProtections(v []MicrosoftGraphDefaultManagedAppProtection)`

SetDefaultManagedAppProtections gets a reference to the given []MicrosoftGraphDefaultManagedAppProtection and assigns it to the DefaultManagedAppProtections field.

### GetTargetedManagedAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) GetTargetedManagedAppConfigurations() []MicrosoftGraphTargetedManagedAppConfiguration`

GetTargetedManagedAppConfigurations returns the TargetedManagedAppConfigurations field if non-nil, zero value otherwise.

### GetTargetedManagedAppConfigurationsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetTargetedManagedAppConfigurationsOk() ([]MicrosoftGraphTargetedManagedAppConfiguration, bool)`

GetTargetedManagedAppConfigurationsOk returns a tuple with the TargetedManagedAppConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetedManagedAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) HasTargetedManagedAppConfigurations() bool`

HasTargetedManagedAppConfigurations returns a boolean if a field has been set.

### SetTargetedManagedAppConfigurations

`func (o *MicrosoftGraphDeviceAppManagement) SetTargetedManagedAppConfigurations(v []MicrosoftGraphTargetedManagedAppConfiguration)`

SetTargetedManagedAppConfigurations gets a reference to the given []MicrosoftGraphTargetedManagedAppConfiguration and assigns it to the TargetedManagedAppConfigurations field.

### GetMdmWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) GetMdmWindowsInformationProtectionPolicies() []MicrosoftGraphMdmWindowsInformationProtectionPolicy`

GetMdmWindowsInformationProtectionPolicies returns the MdmWindowsInformationProtectionPolicies field if non-nil, zero value otherwise.

### GetMdmWindowsInformationProtectionPoliciesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetMdmWindowsInformationProtectionPoliciesOk() ([]MicrosoftGraphMdmWindowsInformationProtectionPolicy, bool)`

GetMdmWindowsInformationProtectionPoliciesOk returns a tuple with the MdmWindowsInformationProtectionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMdmWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) HasMdmWindowsInformationProtectionPolicies() bool`

HasMdmWindowsInformationProtectionPolicies returns a boolean if a field has been set.

### SetMdmWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(v []MicrosoftGraphMdmWindowsInformationProtectionPolicy)`

SetMdmWindowsInformationProtectionPolicies gets a reference to the given []MicrosoftGraphMdmWindowsInformationProtectionPolicy and assigns it to the MdmWindowsInformationProtectionPolicies field.

### GetWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) GetWindowsInformationProtectionPolicies() []MicrosoftGraphWindowsInformationProtectionPolicy`

GetWindowsInformationProtectionPolicies returns the WindowsInformationProtectionPolicies field if non-nil, zero value otherwise.

### GetWindowsInformationProtectionPoliciesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetWindowsInformationProtectionPoliciesOk() ([]MicrosoftGraphWindowsInformationProtectionPolicy, bool)`

GetWindowsInformationProtectionPoliciesOk returns a tuple with the WindowsInformationProtectionPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) HasWindowsInformationProtectionPolicies() bool`

HasWindowsInformationProtectionPolicies returns a boolean if a field has been set.

### SetWindowsInformationProtectionPolicies

`func (o *MicrosoftGraphDeviceAppManagement) SetWindowsInformationProtectionPolicies(v []MicrosoftGraphWindowsInformationProtectionPolicy)`

SetWindowsInformationProtectionPolicies gets a reference to the given []MicrosoftGraphWindowsInformationProtectionPolicy and assigns it to the WindowsInformationProtectionPolicies field.

### GetManagedAppRegistrations

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppRegistrations() []MicrosoftGraphManagedAppRegistration`

GetManagedAppRegistrations returns the ManagedAppRegistrations field if non-nil, zero value otherwise.

### GetManagedAppRegistrationsOk

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppRegistrationsOk() ([]MicrosoftGraphManagedAppRegistration, bool)`

GetManagedAppRegistrationsOk returns a tuple with the ManagedAppRegistrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedAppRegistrations

`func (o *MicrosoftGraphDeviceAppManagement) HasManagedAppRegistrations() bool`

HasManagedAppRegistrations returns a boolean if a field has been set.

### SetManagedAppRegistrations

`func (o *MicrosoftGraphDeviceAppManagement) SetManagedAppRegistrations(v []MicrosoftGraphManagedAppRegistration)`

SetManagedAppRegistrations gets a reference to the given []MicrosoftGraphManagedAppRegistration and assigns it to the ManagedAppRegistrations field.

### GetManagedAppStatuses

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppStatuses() []MicrosoftGraphManagedAppStatus`

GetManagedAppStatuses returns the ManagedAppStatuses field if non-nil, zero value otherwise.

### GetManagedAppStatusesOk

`func (o *MicrosoftGraphDeviceAppManagement) GetManagedAppStatusesOk() ([]MicrosoftGraphManagedAppStatus, bool)`

GetManagedAppStatusesOk returns a tuple with the ManagedAppStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedAppStatuses

`func (o *MicrosoftGraphDeviceAppManagement) HasManagedAppStatuses() bool`

HasManagedAppStatuses returns a boolean if a field has been set.

### SetManagedAppStatuses

`func (o *MicrosoftGraphDeviceAppManagement) SetManagedAppStatuses(v []MicrosoftGraphManagedAppStatus)`

SetManagedAppStatuses gets a reference to the given []MicrosoftGraphManagedAppStatus and assigns it to the ManagedAppStatuses field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



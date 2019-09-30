# MicrosoftGraphDeviceManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementSettings**](anyOf&lt;microsoft.graph.deviceManagementSettings&gt;.md) | Account level settings. | [optional] 
**IntuneBrand** | Pointer to [**AnyOfmicrosoftGraphIntuneBrand**](anyOf&lt;microsoft.graph.intuneBrand&gt;.md) | intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal. | [optional] 
**SubscriptionState** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementSubscriptionState**](anyOf&lt;microsoft.graph.deviceManagementSubscriptionState&gt;.md) | Tenant mobile device management subscription state. | [optional] 
**TermsAndConditions** | Pointer to [**[]MicrosoftGraphTermsAndConditions**](microsoft.graph.termsAndConditions.md) |  | [optional] 
**DeviceConfigurations** | Pointer to [**[]MicrosoftGraphDeviceConfiguration**](microsoft.graph.deviceConfiguration.md) |  | [optional] 
**DeviceCompliancePolicies** | Pointer to [**[]MicrosoftGraphDeviceCompliancePolicy**](microsoft.graph.deviceCompliancePolicy.md) |  | [optional] 
**SoftwareUpdateStatusSummary** | Pointer to [**AnyOfmicrosoftGraphSoftwareUpdateStatusSummary**](anyOf&lt;microsoft.graph.softwareUpdateStatusSummary&gt;.md) |  | [optional] 
**DeviceCompliancePolicyDeviceStateSummary** | Pointer to [**AnyOfmicrosoftGraphDeviceCompliancePolicyDeviceStateSummary**](anyOf&lt;microsoft.graph.deviceCompliancePolicyDeviceStateSummary&gt;.md) |  | [optional] 
**DeviceCompliancePolicySettingStateSummaries** | Pointer to [**[]MicrosoftGraphDeviceCompliancePolicySettingStateSummary**](microsoft.graph.deviceCompliancePolicySettingStateSummary.md) |  | [optional] 
**DeviceConfigurationDeviceStateSummaries** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationDeviceStateSummary**](anyOf&lt;microsoft.graph.deviceConfigurationDeviceStateSummary&gt;.md) |  | [optional] 
**IosUpdateStatuses** | Pointer to [**[]MicrosoftGraphIosUpdateDeviceStatus**](microsoft.graph.iosUpdateDeviceStatus.md) |  | [optional] 
**DeviceCategories** | Pointer to [**[]MicrosoftGraphDeviceCategory**](microsoft.graph.deviceCategory.md) |  | [optional] 
**ExchangeConnectors** | Pointer to [**[]MicrosoftGraphDeviceManagementExchangeConnector**](microsoft.graph.deviceManagementExchangeConnector.md) |  | [optional] 
**DeviceEnrollmentConfigurations** | Pointer to [**[]MicrosoftGraphDeviceEnrollmentConfiguration**](microsoft.graph.deviceEnrollmentConfiguration.md) |  | [optional] 
**ConditionalAccessSettings** | Pointer to [**AnyOfmicrosoftGraphOnPremisesConditionalAccessSettings**](anyOf&lt;microsoft.graph.onPremisesConditionalAccessSettings&gt;.md) |  | [optional] 
**MobileThreatDefenseConnectors** | Pointer to [**[]MicrosoftGraphMobileThreatDefenseConnector**](microsoft.graph.mobileThreatDefenseConnector.md) |  | [optional] 
**DeviceManagementPartners** | Pointer to [**[]MicrosoftGraphDeviceManagementPartner**](microsoft.graph.deviceManagementPartner.md) |  | [optional] 
**ApplePushNotificationCertificate** | Pointer to [**AnyOfmicrosoftGraphApplePushNotificationCertificate**](anyOf&lt;microsoft.graph.applePushNotificationCertificate&gt;.md) |  | [optional] 
**ManagedDeviceOverview** | Pointer to [**AnyOfmicrosoftGraphManagedDeviceOverview**](anyOf&lt;microsoft.graph.managedDeviceOverview&gt;.md) |  | [optional] 
**DetectedApps** | Pointer to [**[]MicrosoftGraphDetectedApp**](microsoft.graph.detectedApp.md) |  | [optional] 
**ManagedDevices** | Pointer to [**[]MicrosoftGraphManagedDevice**](microsoft.graph.managedDevice.md) |  | [optional] 
**NotificationMessageTemplates** | Pointer to [**[]MicrosoftGraphNotificationMessageTemplate**](microsoft.graph.notificationMessageTemplate.md) |  | [optional] 
**RoleDefinitions** | Pointer to [**[]MicrosoftGraphRoleDefinition**](microsoft.graph.roleDefinition.md) |  | [optional] 
**RoleAssignments** | Pointer to [**[]MicrosoftGraphDeviceAndAppManagementRoleAssignment**](microsoft.graph.deviceAndAppManagementRoleAssignment.md) |  | [optional] 
**ResourceOperations** | Pointer to [**[]MicrosoftGraphResourceOperation**](microsoft.graph.resourceOperation.md) |  | [optional] 
**RemoteAssistancePartners** | Pointer to [**[]MicrosoftGraphRemoteAssistancePartner**](microsoft.graph.remoteAssistancePartner.md) |  | [optional] 
**TelecomExpenseManagementPartners** | Pointer to [**[]MicrosoftGraphTelecomExpenseManagementPartner**](microsoft.graph.telecomExpenseManagementPartner.md) |  | [optional] 
**TroubleshootingEvents** | Pointer to [**[]MicrosoftGraphDeviceManagementTroubleshootingEvent**](microsoft.graph.deviceManagementTroubleshootingEvent.md) |  | [optional] 
**WindowsInformationProtectionAppLearningSummaries** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionAppLearningSummary**](microsoft.graph.windowsInformationProtectionAppLearningSummary.md) |  | [optional] 
**WindowsInformationProtectionNetworkLearningSummaries** | Pointer to [**[]MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary**](microsoft.graph.windowsInformationProtectionNetworkLearningSummary.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphDeviceManagement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceManagement) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphDeviceManagement) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphDeviceManagement) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetSettings

`func (o *MicrosoftGraphDeviceManagement) GetSettings() AnyOfmicrosoftGraphDeviceManagementSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphDeviceManagement) GetSettingsOk() (AnyOfmicrosoftGraphDeviceManagementSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettings

`func (o *MicrosoftGraphDeviceManagement) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettings

`func (o *MicrosoftGraphDeviceManagement) SetSettings(v AnyOfmicrosoftGraphDeviceManagementSettings)`

SetSettings gets a reference to the given AnyOfmicrosoftGraphDeviceManagementSettings and assigns it to the Settings field.

### SetSettingsExplicitNull

`func (o *MicrosoftGraphDeviceManagement) SetSettingsExplicitNull(b bool)`

SetSettingsExplicitNull (un)sets Settings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Settings value is set to nil even if false is passed
### GetIntuneBrand

`func (o *MicrosoftGraphDeviceManagement) GetIntuneBrand() AnyOfmicrosoftGraphIntuneBrand`

GetIntuneBrand returns the IntuneBrand field if non-nil, zero value otherwise.

### GetIntuneBrandOk

`func (o *MicrosoftGraphDeviceManagement) GetIntuneBrandOk() (AnyOfmicrosoftGraphIntuneBrand, bool)`

GetIntuneBrandOk returns a tuple with the IntuneBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntuneBrand

`func (o *MicrosoftGraphDeviceManagement) HasIntuneBrand() bool`

HasIntuneBrand returns a boolean if a field has been set.

### SetIntuneBrand

`func (o *MicrosoftGraphDeviceManagement) SetIntuneBrand(v AnyOfmicrosoftGraphIntuneBrand)`

SetIntuneBrand gets a reference to the given AnyOfmicrosoftGraphIntuneBrand and assigns it to the IntuneBrand field.

### SetIntuneBrandExplicitNull

`func (o *MicrosoftGraphDeviceManagement) SetIntuneBrandExplicitNull(b bool)`

SetIntuneBrandExplicitNull (un)sets IntuneBrand to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IntuneBrand value is set to nil even if false is passed
### GetSubscriptionState

`func (o *MicrosoftGraphDeviceManagement) GetSubscriptionState() AnyOfmicrosoftGraphDeviceManagementSubscriptionState`

GetSubscriptionState returns the SubscriptionState field if non-nil, zero value otherwise.

### GetSubscriptionStateOk

`func (o *MicrosoftGraphDeviceManagement) GetSubscriptionStateOk() (AnyOfmicrosoftGraphDeviceManagementSubscriptionState, bool)`

GetSubscriptionStateOk returns a tuple with the SubscriptionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscriptionState

`func (o *MicrosoftGraphDeviceManagement) HasSubscriptionState() bool`

HasSubscriptionState returns a boolean if a field has been set.

### SetSubscriptionState

`func (o *MicrosoftGraphDeviceManagement) SetSubscriptionState(v AnyOfmicrosoftGraphDeviceManagementSubscriptionState)`

SetSubscriptionState gets a reference to the given AnyOfmicrosoftGraphDeviceManagementSubscriptionState and assigns it to the SubscriptionState field.

### GetTermsAndConditions

`func (o *MicrosoftGraphDeviceManagement) GetTermsAndConditions() []MicrosoftGraphTermsAndConditions`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *MicrosoftGraphDeviceManagement) GetTermsAndConditionsOk() ([]MicrosoftGraphTermsAndConditions, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTermsAndConditions

`func (o *MicrosoftGraphDeviceManagement) HasTermsAndConditions() bool`

HasTermsAndConditions returns a boolean if a field has been set.

### SetTermsAndConditions

`func (o *MicrosoftGraphDeviceManagement) SetTermsAndConditions(v []MicrosoftGraphTermsAndConditions)`

SetTermsAndConditions gets a reference to the given []MicrosoftGraphTermsAndConditions and assigns it to the TermsAndConditions field.

### GetDeviceConfigurations

`func (o *MicrosoftGraphDeviceManagement) GetDeviceConfigurations() []MicrosoftGraphDeviceConfiguration`

GetDeviceConfigurations returns the DeviceConfigurations field if non-nil, zero value otherwise.

### GetDeviceConfigurationsOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceConfigurationsOk() ([]MicrosoftGraphDeviceConfiguration, bool)`

GetDeviceConfigurationsOk returns a tuple with the DeviceConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceConfigurations

`func (o *MicrosoftGraphDeviceManagement) HasDeviceConfigurations() bool`

HasDeviceConfigurations returns a boolean if a field has been set.

### SetDeviceConfigurations

`func (o *MicrosoftGraphDeviceManagement) SetDeviceConfigurations(v []MicrosoftGraphDeviceConfiguration)`

SetDeviceConfigurations gets a reference to the given []MicrosoftGraphDeviceConfiguration and assigns it to the DeviceConfigurations field.

### GetDeviceCompliancePolicies

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicies() []MicrosoftGraphDeviceCompliancePolicy`

GetDeviceCompliancePolicies returns the DeviceCompliancePolicies field if non-nil, zero value otherwise.

### GetDeviceCompliancePoliciesOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePoliciesOk() ([]MicrosoftGraphDeviceCompliancePolicy, bool)`

GetDeviceCompliancePoliciesOk returns a tuple with the DeviceCompliancePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCompliancePolicies

`func (o *MicrosoftGraphDeviceManagement) HasDeviceCompliancePolicies() bool`

HasDeviceCompliancePolicies returns a boolean if a field has been set.

### SetDeviceCompliancePolicies

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCompliancePolicies(v []MicrosoftGraphDeviceCompliancePolicy)`

SetDeviceCompliancePolicies gets a reference to the given []MicrosoftGraphDeviceCompliancePolicy and assigns it to the DeviceCompliancePolicies field.

### GetSoftwareUpdateStatusSummary

`func (o *MicrosoftGraphDeviceManagement) GetSoftwareUpdateStatusSummary() AnyOfmicrosoftGraphSoftwareUpdateStatusSummary`

GetSoftwareUpdateStatusSummary returns the SoftwareUpdateStatusSummary field if non-nil, zero value otherwise.

### GetSoftwareUpdateStatusSummaryOk

`func (o *MicrosoftGraphDeviceManagement) GetSoftwareUpdateStatusSummaryOk() (AnyOfmicrosoftGraphSoftwareUpdateStatusSummary, bool)`

GetSoftwareUpdateStatusSummaryOk returns a tuple with the SoftwareUpdateStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSoftwareUpdateStatusSummary

`func (o *MicrosoftGraphDeviceManagement) HasSoftwareUpdateStatusSummary() bool`

HasSoftwareUpdateStatusSummary returns a boolean if a field has been set.

### SetSoftwareUpdateStatusSummary

`func (o *MicrosoftGraphDeviceManagement) SetSoftwareUpdateStatusSummary(v AnyOfmicrosoftGraphSoftwareUpdateStatusSummary)`

SetSoftwareUpdateStatusSummary gets a reference to the given AnyOfmicrosoftGraphSoftwareUpdateStatusSummary and assigns it to the SoftwareUpdateStatusSummary field.

### SetSoftwareUpdateStatusSummaryExplicitNull

`func (o *MicrosoftGraphDeviceManagement) SetSoftwareUpdateStatusSummaryExplicitNull(b bool)`

SetSoftwareUpdateStatusSummaryExplicitNull (un)sets SoftwareUpdateStatusSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SoftwareUpdateStatusSummary value is set to nil even if false is passed
### GetDeviceCompliancePolicyDeviceStateSummary

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary() AnyOfmicrosoftGraphDeviceCompliancePolicyDeviceStateSummary`

GetDeviceCompliancePolicyDeviceStateSummary returns the DeviceCompliancePolicyDeviceStateSummary field if non-nil, zero value otherwise.

### GetDeviceCompliancePolicyDeviceStateSummaryOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicyDeviceStateSummaryOk() (AnyOfmicrosoftGraphDeviceCompliancePolicyDeviceStateSummary, bool)`

GetDeviceCompliancePolicyDeviceStateSummaryOk returns a tuple with the DeviceCompliancePolicyDeviceStateSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCompliancePolicyDeviceStateSummary

`func (o *MicrosoftGraphDeviceManagement) HasDeviceCompliancePolicyDeviceStateSummary() bool`

HasDeviceCompliancePolicyDeviceStateSummary returns a boolean if a field has been set.

### SetDeviceCompliancePolicyDeviceStateSummary

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(v AnyOfmicrosoftGraphDeviceCompliancePolicyDeviceStateSummary)`

SetDeviceCompliancePolicyDeviceStateSummary gets a reference to the given AnyOfmicrosoftGraphDeviceCompliancePolicyDeviceStateSummary and assigns it to the DeviceCompliancePolicyDeviceStateSummary field.

### SetDeviceCompliancePolicyDeviceStateSummaryExplicitNull

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCompliancePolicyDeviceStateSummaryExplicitNull(b bool)`

SetDeviceCompliancePolicyDeviceStateSummaryExplicitNull (un)sets DeviceCompliancePolicyDeviceStateSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceCompliancePolicyDeviceStateSummary value is set to nil even if false is passed
### GetDeviceCompliancePolicySettingStateSummaries

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicySettingStateSummaries() []MicrosoftGraphDeviceCompliancePolicySettingStateSummary`

GetDeviceCompliancePolicySettingStateSummaries returns the DeviceCompliancePolicySettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceCompliancePolicySettingStateSummariesOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCompliancePolicySettingStateSummariesOk() ([]MicrosoftGraphDeviceCompliancePolicySettingStateSummary, bool)`

GetDeviceCompliancePolicySettingStateSummariesOk returns a tuple with the DeviceCompliancePolicySettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCompliancePolicySettingStateSummaries

`func (o *MicrosoftGraphDeviceManagement) HasDeviceCompliancePolicySettingStateSummaries() bool`

HasDeviceCompliancePolicySettingStateSummaries returns a boolean if a field has been set.

### SetDeviceCompliancePolicySettingStateSummaries

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(v []MicrosoftGraphDeviceCompliancePolicySettingStateSummary)`

SetDeviceCompliancePolicySettingStateSummaries gets a reference to the given []MicrosoftGraphDeviceCompliancePolicySettingStateSummary and assigns it to the DeviceCompliancePolicySettingStateSummaries field.

### GetDeviceConfigurationDeviceStateSummaries

`func (o *MicrosoftGraphDeviceManagement) GetDeviceConfigurationDeviceStateSummaries() AnyOfmicrosoftGraphDeviceConfigurationDeviceStateSummary`

GetDeviceConfigurationDeviceStateSummaries returns the DeviceConfigurationDeviceStateSummaries field if non-nil, zero value otherwise.

### GetDeviceConfigurationDeviceStateSummariesOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceConfigurationDeviceStateSummariesOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceStateSummary, bool)`

GetDeviceConfigurationDeviceStateSummariesOk returns a tuple with the DeviceConfigurationDeviceStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceConfigurationDeviceStateSummaries

`func (o *MicrosoftGraphDeviceManagement) HasDeviceConfigurationDeviceStateSummaries() bool`

HasDeviceConfigurationDeviceStateSummaries returns a boolean if a field has been set.

### SetDeviceConfigurationDeviceStateSummaries

`func (o *MicrosoftGraphDeviceManagement) SetDeviceConfigurationDeviceStateSummaries(v AnyOfmicrosoftGraphDeviceConfigurationDeviceStateSummary)`

SetDeviceConfigurationDeviceStateSummaries gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceStateSummary and assigns it to the DeviceConfigurationDeviceStateSummaries field.

### SetDeviceConfigurationDeviceStateSummariesExplicitNull

`func (o *MicrosoftGraphDeviceManagement) SetDeviceConfigurationDeviceStateSummariesExplicitNull(b bool)`

SetDeviceConfigurationDeviceStateSummariesExplicitNull (un)sets DeviceConfigurationDeviceStateSummaries to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceConfigurationDeviceStateSummaries value is set to nil even if false is passed
### GetIosUpdateStatuses

`func (o *MicrosoftGraphDeviceManagement) GetIosUpdateStatuses() []MicrosoftGraphIosUpdateDeviceStatus`

GetIosUpdateStatuses returns the IosUpdateStatuses field if non-nil, zero value otherwise.

### GetIosUpdateStatusesOk

`func (o *MicrosoftGraphDeviceManagement) GetIosUpdateStatusesOk() ([]MicrosoftGraphIosUpdateDeviceStatus, bool)`

GetIosUpdateStatusesOk returns a tuple with the IosUpdateStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIosUpdateStatuses

`func (o *MicrosoftGraphDeviceManagement) HasIosUpdateStatuses() bool`

HasIosUpdateStatuses returns a boolean if a field has been set.

### SetIosUpdateStatuses

`func (o *MicrosoftGraphDeviceManagement) SetIosUpdateStatuses(v []MicrosoftGraphIosUpdateDeviceStatus)`

SetIosUpdateStatuses gets a reference to the given []MicrosoftGraphIosUpdateDeviceStatus and assigns it to the IosUpdateStatuses field.

### GetDeviceCategories

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCategories() []MicrosoftGraphDeviceCategory`

GetDeviceCategories returns the DeviceCategories field if non-nil, zero value otherwise.

### GetDeviceCategoriesOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceCategoriesOk() ([]MicrosoftGraphDeviceCategory, bool)`

GetDeviceCategoriesOk returns a tuple with the DeviceCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceCategories

`func (o *MicrosoftGraphDeviceManagement) HasDeviceCategories() bool`

HasDeviceCategories returns a boolean if a field has been set.

### SetDeviceCategories

`func (o *MicrosoftGraphDeviceManagement) SetDeviceCategories(v []MicrosoftGraphDeviceCategory)`

SetDeviceCategories gets a reference to the given []MicrosoftGraphDeviceCategory and assigns it to the DeviceCategories field.

### GetExchangeConnectors

`func (o *MicrosoftGraphDeviceManagement) GetExchangeConnectors() []MicrosoftGraphDeviceManagementExchangeConnector`

GetExchangeConnectors returns the ExchangeConnectors field if non-nil, zero value otherwise.

### GetExchangeConnectorsOk

`func (o *MicrosoftGraphDeviceManagement) GetExchangeConnectorsOk() ([]MicrosoftGraphDeviceManagementExchangeConnector, bool)`

GetExchangeConnectorsOk returns a tuple with the ExchangeConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExchangeConnectors

`func (o *MicrosoftGraphDeviceManagement) HasExchangeConnectors() bool`

HasExchangeConnectors returns a boolean if a field has been set.

### SetExchangeConnectors

`func (o *MicrosoftGraphDeviceManagement) SetExchangeConnectors(v []MicrosoftGraphDeviceManagementExchangeConnector)`

SetExchangeConnectors gets a reference to the given []MicrosoftGraphDeviceManagementExchangeConnector and assigns it to the ExchangeConnectors field.

### GetDeviceEnrollmentConfigurations

`func (o *MicrosoftGraphDeviceManagement) GetDeviceEnrollmentConfigurations() []MicrosoftGraphDeviceEnrollmentConfiguration`

GetDeviceEnrollmentConfigurations returns the DeviceEnrollmentConfigurations field if non-nil, zero value otherwise.

### GetDeviceEnrollmentConfigurationsOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceEnrollmentConfigurationsOk() ([]MicrosoftGraphDeviceEnrollmentConfiguration, bool)`

GetDeviceEnrollmentConfigurationsOk returns a tuple with the DeviceEnrollmentConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceEnrollmentConfigurations

`func (o *MicrosoftGraphDeviceManagement) HasDeviceEnrollmentConfigurations() bool`

HasDeviceEnrollmentConfigurations returns a boolean if a field has been set.

### SetDeviceEnrollmentConfigurations

`func (o *MicrosoftGraphDeviceManagement) SetDeviceEnrollmentConfigurations(v []MicrosoftGraphDeviceEnrollmentConfiguration)`

SetDeviceEnrollmentConfigurations gets a reference to the given []MicrosoftGraphDeviceEnrollmentConfiguration and assigns it to the DeviceEnrollmentConfigurations field.

### GetConditionalAccessSettings

`func (o *MicrosoftGraphDeviceManagement) GetConditionalAccessSettings() AnyOfmicrosoftGraphOnPremisesConditionalAccessSettings`

GetConditionalAccessSettings returns the ConditionalAccessSettings field if non-nil, zero value otherwise.

### GetConditionalAccessSettingsOk

`func (o *MicrosoftGraphDeviceManagement) GetConditionalAccessSettingsOk() (AnyOfmicrosoftGraphOnPremisesConditionalAccessSettings, bool)`

GetConditionalAccessSettingsOk returns a tuple with the ConditionalAccessSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConditionalAccessSettings

`func (o *MicrosoftGraphDeviceManagement) HasConditionalAccessSettings() bool`

HasConditionalAccessSettings returns a boolean if a field has been set.

### SetConditionalAccessSettings

`func (o *MicrosoftGraphDeviceManagement) SetConditionalAccessSettings(v AnyOfmicrosoftGraphOnPremisesConditionalAccessSettings)`

SetConditionalAccessSettings gets a reference to the given AnyOfmicrosoftGraphOnPremisesConditionalAccessSettings and assigns it to the ConditionalAccessSettings field.

### SetConditionalAccessSettingsExplicitNull

`func (o *MicrosoftGraphDeviceManagement) SetConditionalAccessSettingsExplicitNull(b bool)`

SetConditionalAccessSettingsExplicitNull (un)sets ConditionalAccessSettings to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConditionalAccessSettings value is set to nil even if false is passed
### GetMobileThreatDefenseConnectors

`func (o *MicrosoftGraphDeviceManagement) GetMobileThreatDefenseConnectors() []MicrosoftGraphMobileThreatDefenseConnector`

GetMobileThreatDefenseConnectors returns the MobileThreatDefenseConnectors field if non-nil, zero value otherwise.

### GetMobileThreatDefenseConnectorsOk

`func (o *MicrosoftGraphDeviceManagement) GetMobileThreatDefenseConnectorsOk() ([]MicrosoftGraphMobileThreatDefenseConnector, bool)`

GetMobileThreatDefenseConnectorsOk returns a tuple with the MobileThreatDefenseConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileThreatDefenseConnectors

`func (o *MicrosoftGraphDeviceManagement) HasMobileThreatDefenseConnectors() bool`

HasMobileThreatDefenseConnectors returns a boolean if a field has been set.

### SetMobileThreatDefenseConnectors

`func (o *MicrosoftGraphDeviceManagement) SetMobileThreatDefenseConnectors(v []MicrosoftGraphMobileThreatDefenseConnector)`

SetMobileThreatDefenseConnectors gets a reference to the given []MicrosoftGraphMobileThreatDefenseConnector and assigns it to the MobileThreatDefenseConnectors field.

### GetDeviceManagementPartners

`func (o *MicrosoftGraphDeviceManagement) GetDeviceManagementPartners() []MicrosoftGraphDeviceManagementPartner`

GetDeviceManagementPartners returns the DeviceManagementPartners field if non-nil, zero value otherwise.

### GetDeviceManagementPartnersOk

`func (o *MicrosoftGraphDeviceManagement) GetDeviceManagementPartnersOk() ([]MicrosoftGraphDeviceManagementPartner, bool)`

GetDeviceManagementPartnersOk returns a tuple with the DeviceManagementPartners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceManagementPartners

`func (o *MicrosoftGraphDeviceManagement) HasDeviceManagementPartners() bool`

HasDeviceManagementPartners returns a boolean if a field has been set.

### SetDeviceManagementPartners

`func (o *MicrosoftGraphDeviceManagement) SetDeviceManagementPartners(v []MicrosoftGraphDeviceManagementPartner)`

SetDeviceManagementPartners gets a reference to the given []MicrosoftGraphDeviceManagementPartner and assigns it to the DeviceManagementPartners field.

### GetApplePushNotificationCertificate

`func (o *MicrosoftGraphDeviceManagement) GetApplePushNotificationCertificate() AnyOfmicrosoftGraphApplePushNotificationCertificate`

GetApplePushNotificationCertificate returns the ApplePushNotificationCertificate field if non-nil, zero value otherwise.

### GetApplePushNotificationCertificateOk

`func (o *MicrosoftGraphDeviceManagement) GetApplePushNotificationCertificateOk() (AnyOfmicrosoftGraphApplePushNotificationCertificate, bool)`

GetApplePushNotificationCertificateOk returns a tuple with the ApplePushNotificationCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplePushNotificationCertificate

`func (o *MicrosoftGraphDeviceManagement) HasApplePushNotificationCertificate() bool`

HasApplePushNotificationCertificate returns a boolean if a field has been set.

### SetApplePushNotificationCertificate

`func (o *MicrosoftGraphDeviceManagement) SetApplePushNotificationCertificate(v AnyOfmicrosoftGraphApplePushNotificationCertificate)`

SetApplePushNotificationCertificate gets a reference to the given AnyOfmicrosoftGraphApplePushNotificationCertificate and assigns it to the ApplePushNotificationCertificate field.

### SetApplePushNotificationCertificateExplicitNull

`func (o *MicrosoftGraphDeviceManagement) SetApplePushNotificationCertificateExplicitNull(b bool)`

SetApplePushNotificationCertificateExplicitNull (un)sets ApplePushNotificationCertificate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ApplePushNotificationCertificate value is set to nil even if false is passed
### GetManagedDeviceOverview

`func (o *MicrosoftGraphDeviceManagement) GetManagedDeviceOverview() AnyOfmicrosoftGraphManagedDeviceOverview`

GetManagedDeviceOverview returns the ManagedDeviceOverview field if non-nil, zero value otherwise.

### GetManagedDeviceOverviewOk

`func (o *MicrosoftGraphDeviceManagement) GetManagedDeviceOverviewOk() (AnyOfmicrosoftGraphManagedDeviceOverview, bool)`

GetManagedDeviceOverviewOk returns a tuple with the ManagedDeviceOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDeviceOverview

`func (o *MicrosoftGraphDeviceManagement) HasManagedDeviceOverview() bool`

HasManagedDeviceOverview returns a boolean if a field has been set.

### SetManagedDeviceOverview

`func (o *MicrosoftGraphDeviceManagement) SetManagedDeviceOverview(v AnyOfmicrosoftGraphManagedDeviceOverview)`

SetManagedDeviceOverview gets a reference to the given AnyOfmicrosoftGraphManagedDeviceOverview and assigns it to the ManagedDeviceOverview field.

### SetManagedDeviceOverviewExplicitNull

`func (o *MicrosoftGraphDeviceManagement) SetManagedDeviceOverviewExplicitNull(b bool)`

SetManagedDeviceOverviewExplicitNull (un)sets ManagedDeviceOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ManagedDeviceOverview value is set to nil even if false is passed
### GetDetectedApps

`func (o *MicrosoftGraphDeviceManagement) GetDetectedApps() []MicrosoftGraphDetectedApp`

GetDetectedApps returns the DetectedApps field if non-nil, zero value otherwise.

### GetDetectedAppsOk

`func (o *MicrosoftGraphDeviceManagement) GetDetectedAppsOk() ([]MicrosoftGraphDetectedApp, bool)`

GetDetectedAppsOk returns a tuple with the DetectedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetectedApps

`func (o *MicrosoftGraphDeviceManagement) HasDetectedApps() bool`

HasDetectedApps returns a boolean if a field has been set.

### SetDetectedApps

`func (o *MicrosoftGraphDeviceManagement) SetDetectedApps(v []MicrosoftGraphDetectedApp)`

SetDetectedApps gets a reference to the given []MicrosoftGraphDetectedApp and assigns it to the DetectedApps field.

### GetManagedDevices

`func (o *MicrosoftGraphDeviceManagement) GetManagedDevices() []MicrosoftGraphManagedDevice`

GetManagedDevices returns the ManagedDevices field if non-nil, zero value otherwise.

### GetManagedDevicesOk

`func (o *MicrosoftGraphDeviceManagement) GetManagedDevicesOk() ([]MicrosoftGraphManagedDevice, bool)`

GetManagedDevicesOk returns a tuple with the ManagedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasManagedDevices

`func (o *MicrosoftGraphDeviceManagement) HasManagedDevices() bool`

HasManagedDevices returns a boolean if a field has been set.

### SetManagedDevices

`func (o *MicrosoftGraphDeviceManagement) SetManagedDevices(v []MicrosoftGraphManagedDevice)`

SetManagedDevices gets a reference to the given []MicrosoftGraphManagedDevice and assigns it to the ManagedDevices field.

### GetNotificationMessageTemplates

`func (o *MicrosoftGraphDeviceManagement) GetNotificationMessageTemplates() []MicrosoftGraphNotificationMessageTemplate`

GetNotificationMessageTemplates returns the NotificationMessageTemplates field if non-nil, zero value otherwise.

### GetNotificationMessageTemplatesOk

`func (o *MicrosoftGraphDeviceManagement) GetNotificationMessageTemplatesOk() ([]MicrosoftGraphNotificationMessageTemplate, bool)`

GetNotificationMessageTemplatesOk returns a tuple with the NotificationMessageTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationMessageTemplates

`func (o *MicrosoftGraphDeviceManagement) HasNotificationMessageTemplates() bool`

HasNotificationMessageTemplates returns a boolean if a field has been set.

### SetNotificationMessageTemplates

`func (o *MicrosoftGraphDeviceManagement) SetNotificationMessageTemplates(v []MicrosoftGraphNotificationMessageTemplate)`

SetNotificationMessageTemplates gets a reference to the given []MicrosoftGraphNotificationMessageTemplate and assigns it to the NotificationMessageTemplates field.

### GetRoleDefinitions

`func (o *MicrosoftGraphDeviceManagement) GetRoleDefinitions() []MicrosoftGraphRoleDefinition`

GetRoleDefinitions returns the RoleDefinitions field if non-nil, zero value otherwise.

### GetRoleDefinitionsOk

`func (o *MicrosoftGraphDeviceManagement) GetRoleDefinitionsOk() ([]MicrosoftGraphRoleDefinition, bool)`

GetRoleDefinitionsOk returns a tuple with the RoleDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoleDefinitions

`func (o *MicrosoftGraphDeviceManagement) HasRoleDefinitions() bool`

HasRoleDefinitions returns a boolean if a field has been set.

### SetRoleDefinitions

`func (o *MicrosoftGraphDeviceManagement) SetRoleDefinitions(v []MicrosoftGraphRoleDefinition)`

SetRoleDefinitions gets a reference to the given []MicrosoftGraphRoleDefinition and assigns it to the RoleDefinitions field.

### GetRoleAssignments

`func (o *MicrosoftGraphDeviceManagement) GetRoleAssignments() []MicrosoftGraphDeviceAndAppManagementRoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *MicrosoftGraphDeviceManagement) GetRoleAssignmentsOk() ([]MicrosoftGraphDeviceAndAppManagementRoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoleAssignments

`func (o *MicrosoftGraphDeviceManagement) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.

### SetRoleAssignments

`func (o *MicrosoftGraphDeviceManagement) SetRoleAssignments(v []MicrosoftGraphDeviceAndAppManagementRoleAssignment)`

SetRoleAssignments gets a reference to the given []MicrosoftGraphDeviceAndAppManagementRoleAssignment and assigns it to the RoleAssignments field.

### GetResourceOperations

`func (o *MicrosoftGraphDeviceManagement) GetResourceOperations() []MicrosoftGraphResourceOperation`

GetResourceOperations returns the ResourceOperations field if non-nil, zero value otherwise.

### GetResourceOperationsOk

`func (o *MicrosoftGraphDeviceManagement) GetResourceOperationsOk() ([]MicrosoftGraphResourceOperation, bool)`

GetResourceOperationsOk returns a tuple with the ResourceOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceOperations

`func (o *MicrosoftGraphDeviceManagement) HasResourceOperations() bool`

HasResourceOperations returns a boolean if a field has been set.

### SetResourceOperations

`func (o *MicrosoftGraphDeviceManagement) SetResourceOperations(v []MicrosoftGraphResourceOperation)`

SetResourceOperations gets a reference to the given []MicrosoftGraphResourceOperation and assigns it to the ResourceOperations field.

### GetRemoteAssistancePartners

`func (o *MicrosoftGraphDeviceManagement) GetRemoteAssistancePartners() []MicrosoftGraphRemoteAssistancePartner`

GetRemoteAssistancePartners returns the RemoteAssistancePartners field if non-nil, zero value otherwise.

### GetRemoteAssistancePartnersOk

`func (o *MicrosoftGraphDeviceManagement) GetRemoteAssistancePartnersOk() ([]MicrosoftGraphRemoteAssistancePartner, bool)`

GetRemoteAssistancePartnersOk returns a tuple with the RemoteAssistancePartners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemoteAssistancePartners

`func (o *MicrosoftGraphDeviceManagement) HasRemoteAssistancePartners() bool`

HasRemoteAssistancePartners returns a boolean if a field has been set.

### SetRemoteAssistancePartners

`func (o *MicrosoftGraphDeviceManagement) SetRemoteAssistancePartners(v []MicrosoftGraphRemoteAssistancePartner)`

SetRemoteAssistancePartners gets a reference to the given []MicrosoftGraphRemoteAssistancePartner and assigns it to the RemoteAssistancePartners field.

### GetTelecomExpenseManagementPartners

`func (o *MicrosoftGraphDeviceManagement) GetTelecomExpenseManagementPartners() []MicrosoftGraphTelecomExpenseManagementPartner`

GetTelecomExpenseManagementPartners returns the TelecomExpenseManagementPartners field if non-nil, zero value otherwise.

### GetTelecomExpenseManagementPartnersOk

`func (o *MicrosoftGraphDeviceManagement) GetTelecomExpenseManagementPartnersOk() ([]MicrosoftGraphTelecomExpenseManagementPartner, bool)`

GetTelecomExpenseManagementPartnersOk returns a tuple with the TelecomExpenseManagementPartners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTelecomExpenseManagementPartners

`func (o *MicrosoftGraphDeviceManagement) HasTelecomExpenseManagementPartners() bool`

HasTelecomExpenseManagementPartners returns a boolean if a field has been set.

### SetTelecomExpenseManagementPartners

`func (o *MicrosoftGraphDeviceManagement) SetTelecomExpenseManagementPartners(v []MicrosoftGraphTelecomExpenseManagementPartner)`

SetTelecomExpenseManagementPartners gets a reference to the given []MicrosoftGraphTelecomExpenseManagementPartner and assigns it to the TelecomExpenseManagementPartners field.

### GetTroubleshootingEvents

`func (o *MicrosoftGraphDeviceManagement) GetTroubleshootingEvents() []MicrosoftGraphDeviceManagementTroubleshootingEvent`

GetTroubleshootingEvents returns the TroubleshootingEvents field if non-nil, zero value otherwise.

### GetTroubleshootingEventsOk

`func (o *MicrosoftGraphDeviceManagement) GetTroubleshootingEventsOk() ([]MicrosoftGraphDeviceManagementTroubleshootingEvent, bool)`

GetTroubleshootingEventsOk returns a tuple with the TroubleshootingEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTroubleshootingEvents

`func (o *MicrosoftGraphDeviceManagement) HasTroubleshootingEvents() bool`

HasTroubleshootingEvents returns a boolean if a field has been set.

### SetTroubleshootingEvents

`func (o *MicrosoftGraphDeviceManagement) SetTroubleshootingEvents(v []MicrosoftGraphDeviceManagementTroubleshootingEvent)`

SetTroubleshootingEvents gets a reference to the given []MicrosoftGraphDeviceManagementTroubleshootingEvent and assigns it to the TroubleshootingEvents field.

### GetWindowsInformationProtectionAppLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) GetWindowsInformationProtectionAppLearningSummaries() []MicrosoftGraphWindowsInformationProtectionAppLearningSummary`

GetWindowsInformationProtectionAppLearningSummaries returns the WindowsInformationProtectionAppLearningSummaries field if non-nil, zero value otherwise.

### GetWindowsInformationProtectionAppLearningSummariesOk

`func (o *MicrosoftGraphDeviceManagement) GetWindowsInformationProtectionAppLearningSummariesOk() ([]MicrosoftGraphWindowsInformationProtectionAppLearningSummary, bool)`

GetWindowsInformationProtectionAppLearningSummariesOk returns a tuple with the WindowsInformationProtectionAppLearningSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsInformationProtectionAppLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) HasWindowsInformationProtectionAppLearningSummaries() bool`

HasWindowsInformationProtectionAppLearningSummaries returns a boolean if a field has been set.

### SetWindowsInformationProtectionAppLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(v []MicrosoftGraphWindowsInformationProtectionAppLearningSummary)`

SetWindowsInformationProtectionAppLearningSummaries gets a reference to the given []MicrosoftGraphWindowsInformationProtectionAppLearningSummary and assigns it to the WindowsInformationProtectionAppLearningSummaries field.

### GetWindowsInformationProtectionNetworkLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries() []MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary`

GetWindowsInformationProtectionNetworkLearningSummaries returns the WindowsInformationProtectionNetworkLearningSummaries field if non-nil, zero value otherwise.

### GetWindowsInformationProtectionNetworkLearningSummariesOk

`func (o *MicrosoftGraphDeviceManagement) GetWindowsInformationProtectionNetworkLearningSummariesOk() ([]MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary, bool)`

GetWindowsInformationProtectionNetworkLearningSummariesOk returns a tuple with the WindowsInformationProtectionNetworkLearningSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWindowsInformationProtectionNetworkLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) HasWindowsInformationProtectionNetworkLearningSummaries() bool`

HasWindowsInformationProtectionNetworkLearningSummaries returns a boolean if a field has been set.

### SetWindowsInformationProtectionNetworkLearningSummaries

`func (o *MicrosoftGraphDeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(v []MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary)`

SetWindowsInformationProtectionNetworkLearningSummaries gets a reference to the given []MicrosoftGraphWindowsInformationProtectionNetworkLearningSummary and assigns it to the WindowsInformationProtectionNetworkLearningSummaries field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



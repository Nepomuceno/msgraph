# DefaultManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetAppDataEncryptionType

`func (o *DefaultManagedAppProtection) GetAppDataEncryptionType() AnyOfmicrosoftGraphManagedAppDataEncryptionType`

GetAppDataEncryptionType returns the AppDataEncryptionType field if non-nil, zero value otherwise.

### GetAppDataEncryptionTypeOk

`func (o *DefaultManagedAppProtection) GetAppDataEncryptionTypeOk() (AnyOfmicrosoftGraphManagedAppDataEncryptionType, bool)`

GetAppDataEncryptionTypeOk returns a tuple with the AppDataEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDataEncryptionType

`func (o *DefaultManagedAppProtection) HasAppDataEncryptionType() bool`

HasAppDataEncryptionType returns a boolean if a field has been set.

### SetAppDataEncryptionType

`func (o *DefaultManagedAppProtection) SetAppDataEncryptionType(v AnyOfmicrosoftGraphManagedAppDataEncryptionType)`

SetAppDataEncryptionType gets a reference to the given AnyOfmicrosoftGraphManagedAppDataEncryptionType and assigns it to the AppDataEncryptionType field.

### GetScreenCaptureBlocked

`func (o *DefaultManagedAppProtection) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *DefaultManagedAppProtection) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *DefaultManagedAppProtection) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *DefaultManagedAppProtection) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetEncryptAppData

`func (o *DefaultManagedAppProtection) GetEncryptAppData() bool`

GetEncryptAppData returns the EncryptAppData field if non-nil, zero value otherwise.

### GetEncryptAppDataOk

`func (o *DefaultManagedAppProtection) GetEncryptAppDataOk() (bool, bool)`

GetEncryptAppDataOk returns a tuple with the EncryptAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEncryptAppData

`func (o *DefaultManagedAppProtection) HasEncryptAppData() bool`

HasEncryptAppData returns a boolean if a field has been set.

### SetEncryptAppData

`func (o *DefaultManagedAppProtection) SetEncryptAppData(v bool)`

SetEncryptAppData gets a reference to the given bool and assigns it to the EncryptAppData field.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *DefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabled returns the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if non-nil, zero value otherwise.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk

`func (o *DefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk() (bool, bool)`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk returns a tuple with the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *DefaultManagedAppProtection) HasDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

HasDisableAppEncryptionIfDeviceEncryptionIsEnabled returns a boolean if a field has been set.

### SetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *DefaultManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(v bool)`

SetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets a reference to the given bool and assigns it to the DisableAppEncryptionIfDeviceEncryptionIsEnabled field.

### GetMinimumRequiredSdkVersion

`func (o *DefaultManagedAppProtection) GetMinimumRequiredSdkVersion() string`

GetMinimumRequiredSdkVersion returns the MinimumRequiredSdkVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredSdkVersionOk

`func (o *DefaultManagedAppProtection) GetMinimumRequiredSdkVersionOk() (string, bool)`

GetMinimumRequiredSdkVersionOk returns a tuple with the MinimumRequiredSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredSdkVersion

`func (o *DefaultManagedAppProtection) HasMinimumRequiredSdkVersion() bool`

HasMinimumRequiredSdkVersion returns a boolean if a field has been set.

### SetMinimumRequiredSdkVersion

`func (o *DefaultManagedAppProtection) SetMinimumRequiredSdkVersion(v string)`

SetMinimumRequiredSdkVersion gets a reference to the given string and assigns it to the MinimumRequiredSdkVersion field.

### SetMinimumRequiredSdkVersionExplicitNull

`func (o *DefaultManagedAppProtection) SetMinimumRequiredSdkVersionExplicitNull(b bool)`

SetMinimumRequiredSdkVersionExplicitNull (un)sets MinimumRequiredSdkVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredSdkVersion value is set to nil even if false is passed
### GetCustomSettings

`func (o *DefaultManagedAppProtection) GetCustomSettings() []MicrosoftGraphKeyValuePair`

GetCustomSettings returns the CustomSettings field if non-nil, zero value otherwise.

### GetCustomSettingsOk

`func (o *DefaultManagedAppProtection) GetCustomSettingsOk() ([]MicrosoftGraphKeyValuePair, bool)`

GetCustomSettingsOk returns a tuple with the CustomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomSettings

`func (o *DefaultManagedAppProtection) HasCustomSettings() bool`

HasCustomSettings returns a boolean if a field has been set.

### SetCustomSettings

`func (o *DefaultManagedAppProtection) SetCustomSettings(v []MicrosoftGraphKeyValuePair)`

SetCustomSettings gets a reference to the given []MicrosoftGraphKeyValuePair and assigns it to the CustomSettings field.

### GetDeployedAppCount

`func (o *DefaultManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *DefaultManagedAppProtection) GetDeployedAppCountOk() (int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeployedAppCount

`func (o *DefaultManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### SetDeployedAppCount

`func (o *DefaultManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount gets a reference to the given int32 and assigns it to the DeployedAppCount field.

### GetMinimumRequiredPatchVersion

`func (o *DefaultManagedAppProtection) GetMinimumRequiredPatchVersion() string`

GetMinimumRequiredPatchVersion returns the MinimumRequiredPatchVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredPatchVersionOk

`func (o *DefaultManagedAppProtection) GetMinimumRequiredPatchVersionOk() (string, bool)`

GetMinimumRequiredPatchVersionOk returns a tuple with the MinimumRequiredPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredPatchVersion

`func (o *DefaultManagedAppProtection) HasMinimumRequiredPatchVersion() bool`

HasMinimumRequiredPatchVersion returns a boolean if a field has been set.

### SetMinimumRequiredPatchVersion

`func (o *DefaultManagedAppProtection) SetMinimumRequiredPatchVersion(v string)`

SetMinimumRequiredPatchVersion gets a reference to the given string and assigns it to the MinimumRequiredPatchVersion field.

### SetMinimumRequiredPatchVersionExplicitNull

`func (o *DefaultManagedAppProtection) SetMinimumRequiredPatchVersionExplicitNull(b bool)`

SetMinimumRequiredPatchVersionExplicitNull (un)sets MinimumRequiredPatchVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredPatchVersion value is set to nil even if false is passed
### GetMinimumWarningPatchVersion

`func (o *DefaultManagedAppProtection) GetMinimumWarningPatchVersion() string`

GetMinimumWarningPatchVersion returns the MinimumWarningPatchVersion field if non-nil, zero value otherwise.

### GetMinimumWarningPatchVersionOk

`func (o *DefaultManagedAppProtection) GetMinimumWarningPatchVersionOk() (string, bool)`

GetMinimumWarningPatchVersionOk returns a tuple with the MinimumWarningPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningPatchVersion

`func (o *DefaultManagedAppProtection) HasMinimumWarningPatchVersion() bool`

HasMinimumWarningPatchVersion returns a boolean if a field has been set.

### SetMinimumWarningPatchVersion

`func (o *DefaultManagedAppProtection) SetMinimumWarningPatchVersion(v string)`

SetMinimumWarningPatchVersion gets a reference to the given string and assigns it to the MinimumWarningPatchVersion field.

### SetMinimumWarningPatchVersionExplicitNull

`func (o *DefaultManagedAppProtection) SetMinimumWarningPatchVersionExplicitNull(b bool)`

SetMinimumWarningPatchVersionExplicitNull (un)sets MinimumWarningPatchVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningPatchVersion value is set to nil even if false is passed
### GetFaceIdBlocked

`func (o *DefaultManagedAppProtection) GetFaceIdBlocked() bool`

GetFaceIdBlocked returns the FaceIdBlocked field if non-nil, zero value otherwise.

### GetFaceIdBlockedOk

`func (o *DefaultManagedAppProtection) GetFaceIdBlockedOk() (bool, bool)`

GetFaceIdBlockedOk returns a tuple with the FaceIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFaceIdBlocked

`func (o *DefaultManagedAppProtection) HasFaceIdBlocked() bool`

HasFaceIdBlocked returns a boolean if a field has been set.

### SetFaceIdBlocked

`func (o *DefaultManagedAppProtection) SetFaceIdBlocked(v bool)`

SetFaceIdBlocked gets a reference to the given bool and assigns it to the FaceIdBlocked field.

### GetApps

`func (o *DefaultManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *DefaultManagedAppProtection) GetAppsOk() ([]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApps

`func (o *DefaultManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetApps

`func (o *DefaultManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps gets a reference to the given []MicrosoftGraphManagedMobileApp and assigns it to the Apps field.

### GetDeploymentSummary

`func (o *DefaultManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *DefaultManagedAppProtection) GetDeploymentSummaryOk() (AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeploymentSummary

`func (o *DefaultManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummary

`func (o *DefaultManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary gets a reference to the given AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary and assigns it to the DeploymentSummary field.

### SetDeploymentSummaryExplicitNull

`func (o *DefaultManagedAppProtection) SetDeploymentSummaryExplicitNull(b bool)`

SetDeploymentSummaryExplicitNull (un)sets DeploymentSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeploymentSummary value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



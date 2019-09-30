# AndroidManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScreenCaptureBlocked** | Pointer to **bool** | Indicates whether a managed user can take screen captures of managed apps | [optional] 
**DisableAppEncryptionIfDeviceEncryptionIsEnabled** | Pointer to **bool** | When this setting is enabled, app level encryption is disabled if device level encryption is enabled | [optional] 
**EncryptAppData** | Pointer to **bool** | Indicates whether application data for managed apps should be encrypted | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**MinimumRequiredPatchVersion** | Pointer to **string** | Define the oldest required Android security patch level a user can have to gain secure access to the app. | [optional] 
**MinimumWarningPatchVersion** | Pointer to **string** | Define the oldest recommended Android security patch level a user can have for secure access to the app. | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) |  | [optional] 

## Methods

### GetScreenCaptureBlocked

`func (o *AndroidManagedAppProtection) GetScreenCaptureBlocked() bool`

GetScreenCaptureBlocked returns the ScreenCaptureBlocked field if non-nil, zero value otherwise.

### GetScreenCaptureBlockedOk

`func (o *AndroidManagedAppProtection) GetScreenCaptureBlockedOk() (bool, bool)`

GetScreenCaptureBlockedOk returns a tuple with the ScreenCaptureBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScreenCaptureBlocked

`func (o *AndroidManagedAppProtection) HasScreenCaptureBlocked() bool`

HasScreenCaptureBlocked returns a boolean if a field has been set.

### SetScreenCaptureBlocked

`func (o *AndroidManagedAppProtection) SetScreenCaptureBlocked(v bool)`

SetScreenCaptureBlocked gets a reference to the given bool and assigns it to the ScreenCaptureBlocked field.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *AndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabled returns the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if non-nil, zero value otherwise.

### GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk

`func (o *AndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk() (bool, bool)`

GetDisableAppEncryptionIfDeviceEncryptionIsEnabledOk returns a tuple with the DisableAppEncryptionIfDeviceEncryptionIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *AndroidManagedAppProtection) HasDisableAppEncryptionIfDeviceEncryptionIsEnabled() bool`

HasDisableAppEncryptionIfDeviceEncryptionIsEnabled returns a boolean if a field has been set.

### SetDisableAppEncryptionIfDeviceEncryptionIsEnabled

`func (o *AndroidManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(v bool)`

SetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets a reference to the given bool and assigns it to the DisableAppEncryptionIfDeviceEncryptionIsEnabled field.

### GetEncryptAppData

`func (o *AndroidManagedAppProtection) GetEncryptAppData() bool`

GetEncryptAppData returns the EncryptAppData field if non-nil, zero value otherwise.

### GetEncryptAppDataOk

`func (o *AndroidManagedAppProtection) GetEncryptAppDataOk() (bool, bool)`

GetEncryptAppDataOk returns a tuple with the EncryptAppData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEncryptAppData

`func (o *AndroidManagedAppProtection) HasEncryptAppData() bool`

HasEncryptAppData returns a boolean if a field has been set.

### SetEncryptAppData

`func (o *AndroidManagedAppProtection) SetEncryptAppData(v bool)`

SetEncryptAppData gets a reference to the given bool and assigns it to the EncryptAppData field.

### GetDeployedAppCount

`func (o *AndroidManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *AndroidManagedAppProtection) GetDeployedAppCountOk() (int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeployedAppCount

`func (o *AndroidManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### SetDeployedAppCount

`func (o *AndroidManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount gets a reference to the given int32 and assigns it to the DeployedAppCount field.

### GetMinimumRequiredPatchVersion

`func (o *AndroidManagedAppProtection) GetMinimumRequiredPatchVersion() string`

GetMinimumRequiredPatchVersion returns the MinimumRequiredPatchVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredPatchVersionOk

`func (o *AndroidManagedAppProtection) GetMinimumRequiredPatchVersionOk() (string, bool)`

GetMinimumRequiredPatchVersionOk returns a tuple with the MinimumRequiredPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredPatchVersion

`func (o *AndroidManagedAppProtection) HasMinimumRequiredPatchVersion() bool`

HasMinimumRequiredPatchVersion returns a boolean if a field has been set.

### SetMinimumRequiredPatchVersion

`func (o *AndroidManagedAppProtection) SetMinimumRequiredPatchVersion(v string)`

SetMinimumRequiredPatchVersion gets a reference to the given string and assigns it to the MinimumRequiredPatchVersion field.

### SetMinimumRequiredPatchVersionExplicitNull

`func (o *AndroidManagedAppProtection) SetMinimumRequiredPatchVersionExplicitNull(b bool)`

SetMinimumRequiredPatchVersionExplicitNull (un)sets MinimumRequiredPatchVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredPatchVersion value is set to nil even if false is passed
### GetMinimumWarningPatchVersion

`func (o *AndroidManagedAppProtection) GetMinimumWarningPatchVersion() string`

GetMinimumWarningPatchVersion returns the MinimumWarningPatchVersion field if non-nil, zero value otherwise.

### GetMinimumWarningPatchVersionOk

`func (o *AndroidManagedAppProtection) GetMinimumWarningPatchVersionOk() (string, bool)`

GetMinimumWarningPatchVersionOk returns a tuple with the MinimumWarningPatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumWarningPatchVersion

`func (o *AndroidManagedAppProtection) HasMinimumWarningPatchVersion() bool`

HasMinimumWarningPatchVersion returns a boolean if a field has been set.

### SetMinimumWarningPatchVersion

`func (o *AndroidManagedAppProtection) SetMinimumWarningPatchVersion(v string)`

SetMinimumWarningPatchVersion gets a reference to the given string and assigns it to the MinimumWarningPatchVersion field.

### SetMinimumWarningPatchVersionExplicitNull

`func (o *AndroidManagedAppProtection) SetMinimumWarningPatchVersionExplicitNull(b bool)`

SetMinimumWarningPatchVersionExplicitNull (un)sets MinimumWarningPatchVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumWarningPatchVersion value is set to nil even if false is passed
### GetApps

`func (o *AndroidManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AndroidManagedAppProtection) GetAppsOk() ([]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApps

`func (o *AndroidManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetApps

`func (o *AndroidManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps gets a reference to the given []MicrosoftGraphManagedMobileApp and assigns it to the Apps field.

### GetDeploymentSummary

`func (o *AndroidManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *AndroidManagedAppProtection) GetDeploymentSummaryOk() (AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeploymentSummary

`func (o *AndroidManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummary

`func (o *AndroidManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary gets a reference to the given AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary and assigns it to the DeploymentSummary field.

### SetDeploymentSummaryExplicitNull

`func (o *AndroidManagedAppProtection) SetDeploymentSummaryExplicitNull(b bool)`

SetDeploymentSummaryExplicitNull (un)sets DeploymentSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeploymentSummary value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



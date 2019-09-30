# IosManagedAppProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDataEncryptionType** | Pointer to [**AnyOfmicrosoftGraphManagedAppDataEncryptionType**](anyOf&lt;microsoft.graph.managedAppDataEncryptionType&gt;.md) | Type of encryption which should be used for data in a managed app. | [optional] 
**MinimumRequiredSdkVersion** | Pointer to **string** | Versions less than the specified version will block the managed app from accessing company data. | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**FaceIdBlocked** | Pointer to **bool** | Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) |  | [optional] 

## Methods

### GetAppDataEncryptionType

`func (o *IosManagedAppProtection) GetAppDataEncryptionType() AnyOfmicrosoftGraphManagedAppDataEncryptionType`

GetAppDataEncryptionType returns the AppDataEncryptionType field if non-nil, zero value otherwise.

### GetAppDataEncryptionTypeOk

`func (o *IosManagedAppProtection) GetAppDataEncryptionTypeOk() (AnyOfmicrosoftGraphManagedAppDataEncryptionType, bool)`

GetAppDataEncryptionTypeOk returns a tuple with the AppDataEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppDataEncryptionType

`func (o *IosManagedAppProtection) HasAppDataEncryptionType() bool`

HasAppDataEncryptionType returns a boolean if a field has been set.

### SetAppDataEncryptionType

`func (o *IosManagedAppProtection) SetAppDataEncryptionType(v AnyOfmicrosoftGraphManagedAppDataEncryptionType)`

SetAppDataEncryptionType gets a reference to the given AnyOfmicrosoftGraphManagedAppDataEncryptionType and assigns it to the AppDataEncryptionType field.

### GetMinimumRequiredSdkVersion

`func (o *IosManagedAppProtection) GetMinimumRequiredSdkVersion() string`

GetMinimumRequiredSdkVersion returns the MinimumRequiredSdkVersion field if non-nil, zero value otherwise.

### GetMinimumRequiredSdkVersionOk

`func (o *IosManagedAppProtection) GetMinimumRequiredSdkVersionOk() (string, bool)`

GetMinimumRequiredSdkVersionOk returns a tuple with the MinimumRequiredSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinimumRequiredSdkVersion

`func (o *IosManagedAppProtection) HasMinimumRequiredSdkVersion() bool`

HasMinimumRequiredSdkVersion returns a boolean if a field has been set.

### SetMinimumRequiredSdkVersion

`func (o *IosManagedAppProtection) SetMinimumRequiredSdkVersion(v string)`

SetMinimumRequiredSdkVersion gets a reference to the given string and assigns it to the MinimumRequiredSdkVersion field.

### SetMinimumRequiredSdkVersionExplicitNull

`func (o *IosManagedAppProtection) SetMinimumRequiredSdkVersionExplicitNull(b bool)`

SetMinimumRequiredSdkVersionExplicitNull (un)sets MinimumRequiredSdkVersion to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MinimumRequiredSdkVersion value is set to nil even if false is passed
### GetDeployedAppCount

`func (o *IosManagedAppProtection) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *IosManagedAppProtection) GetDeployedAppCountOk() (int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeployedAppCount

`func (o *IosManagedAppProtection) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### SetDeployedAppCount

`func (o *IosManagedAppProtection) SetDeployedAppCount(v int32)`

SetDeployedAppCount gets a reference to the given int32 and assigns it to the DeployedAppCount field.

### GetFaceIdBlocked

`func (o *IosManagedAppProtection) GetFaceIdBlocked() bool`

GetFaceIdBlocked returns the FaceIdBlocked field if non-nil, zero value otherwise.

### GetFaceIdBlockedOk

`func (o *IosManagedAppProtection) GetFaceIdBlockedOk() (bool, bool)`

GetFaceIdBlockedOk returns a tuple with the FaceIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFaceIdBlocked

`func (o *IosManagedAppProtection) HasFaceIdBlocked() bool`

HasFaceIdBlocked returns a boolean if a field has been set.

### SetFaceIdBlocked

`func (o *IosManagedAppProtection) SetFaceIdBlocked(v bool)`

SetFaceIdBlocked gets a reference to the given bool and assigns it to the FaceIdBlocked field.

### GetApps

`func (o *IosManagedAppProtection) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *IosManagedAppProtection) GetAppsOk() ([]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApps

`func (o *IosManagedAppProtection) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetApps

`func (o *IosManagedAppProtection) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps gets a reference to the given []MicrosoftGraphManagedMobileApp and assigns it to the Apps field.

### GetDeploymentSummary

`func (o *IosManagedAppProtection) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *IosManagedAppProtection) GetDeploymentSummaryOk() (AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeploymentSummary

`func (o *IosManagedAppProtection) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummary

`func (o *IosManagedAppProtection) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary gets a reference to the given AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary and assigns it to the DeploymentSummary field.

### SetDeploymentSummaryExplicitNull

`func (o *IosManagedAppProtection) SetDeploymentSummaryExplicitNull(b bool)`

SetDeploymentSummaryExplicitNull (un)sets DeploymentSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeploymentSummary value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



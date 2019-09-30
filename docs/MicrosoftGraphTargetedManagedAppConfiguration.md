# MicrosoftGraphTargetedManagedAppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Policy display name. | [optional] 
**Description** | Pointer to **string** | The policy&#39;s description. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date and time the policy was created. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | Last time the policy was modified. | [optional] 
**Version** | Pointer to **string** | Version of the entity. | [optional] 
**CustomSettings** | Pointer to [**[]MicrosoftGraphKeyValuePair**](microsoft.graph.keyValuePair.md) | A set of string key and string value pairs to be sent to apps for users to whom the configuration is scoped, unalterned by this service | [optional] 
**DeployedAppCount** | Pointer to **int32** | Count of apps to which the current policy is deployed. | [optional] 
**IsAssigned** | Pointer to **bool** | Indicates if the policy is deployed to any inclusion groups or not. | [optional] 
**Apps** | Pointer to [**[]MicrosoftGraphManagedMobileApp**](microsoft.graph.managedMobileApp.md) |  | [optional] 
**DeploymentSummary** | Pointer to [**AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummary&gt;.md) |  | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphTargetedManagedAppPolicyAssignment**](microsoft.graph.targetedManagedAppPolicyAssignment.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetDisplayName

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetDescription

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetCreatedDateTime

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetVersion

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed
### GetCustomSettings

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetCustomSettings() []MicrosoftGraphKeyValuePair`

GetCustomSettings returns the CustomSettings field if non-nil, zero value otherwise.

### GetCustomSettingsOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetCustomSettingsOk() ([]MicrosoftGraphKeyValuePair, bool)`

GetCustomSettingsOk returns a tuple with the CustomSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomSettings

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasCustomSettings() bool`

HasCustomSettings returns a boolean if a field has been set.

### SetCustomSettings

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetCustomSettings(v []MicrosoftGraphKeyValuePair)`

SetCustomSettings gets a reference to the given []MicrosoftGraphKeyValuePair and assigns it to the CustomSettings field.

### GetDeployedAppCount

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetDeployedAppCount() int32`

GetDeployedAppCount returns the DeployedAppCount field if non-nil, zero value otherwise.

### GetDeployedAppCountOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetDeployedAppCountOk() (int32, bool)`

GetDeployedAppCountOk returns a tuple with the DeployedAppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeployedAppCount

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasDeployedAppCount() bool`

HasDeployedAppCount returns a boolean if a field has been set.

### SetDeployedAppCount

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetDeployedAppCount(v int32)`

SetDeployedAppCount gets a reference to the given int32 and assigns it to the DeployedAppCount field.

### GetIsAssigned

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetIsAssigned() bool`

GetIsAssigned returns the IsAssigned field if non-nil, zero value otherwise.

### GetIsAssignedOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetIsAssignedOk() (bool, bool)`

GetIsAssignedOk returns a tuple with the IsAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAssigned

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasIsAssigned() bool`

HasIsAssigned returns a boolean if a field has been set.

### SetIsAssigned

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetIsAssigned(v bool)`

SetIsAssigned gets a reference to the given bool and assigns it to the IsAssigned field.

### GetApps

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetApps() []MicrosoftGraphManagedMobileApp`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetAppsOk() ([]MicrosoftGraphManagedMobileApp, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApps

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetApps

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetApps(v []MicrosoftGraphManagedMobileApp)`

SetApps gets a reference to the given []MicrosoftGraphManagedMobileApp and assigns it to the Apps field.

### GetDeploymentSummary

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetDeploymentSummary() AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary`

GetDeploymentSummary returns the DeploymentSummary field if non-nil, zero value otherwise.

### GetDeploymentSummaryOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetDeploymentSummaryOk() (AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary, bool)`

GetDeploymentSummaryOk returns a tuple with the DeploymentSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeploymentSummary

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasDeploymentSummary() bool`

HasDeploymentSummary returns a boolean if a field has been set.

### SetDeploymentSummary

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetDeploymentSummary(v AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary)`

SetDeploymentSummary gets a reference to the given AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummary and assigns it to the DeploymentSummary field.

### SetDeploymentSummaryExplicitNull

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetDeploymentSummaryExplicitNull(b bool)`

SetDeploymentSummaryExplicitNull (un)sets DeploymentSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeploymentSummary value is set to nil even if false is passed
### GetAssignments

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetAssignments() []MicrosoftGraphTargetedManagedAppPolicyAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) GetAssignmentsOk() ([]MicrosoftGraphTargetedManagedAppPolicyAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphTargetedManagedAppConfiguration) SetAssignments(v []MicrosoftGraphTargetedManagedAppPolicyAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphTargetedManagedAppPolicyAssignment and assigns it to the Assignments field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



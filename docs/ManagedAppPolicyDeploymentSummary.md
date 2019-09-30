# ManagedAppPolicyDeploymentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**ConfigurationDeployedUserCount** | Pointer to **int32** |  | [optional] 
**LastRefreshTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ConfigurationDeploymentSummaryPerApp** | Pointer to [**[]AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp**](anyOf&lt;microsoft.graph.managedAppPolicyDeploymentSummaryPerApp&gt;.md) |  | [optional] 
**Version** | Pointer to **string** | Version of the entity. | [optional] 

## Methods

### GetDisplayName

`func (o *ManagedAppPolicyDeploymentSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManagedAppPolicyDeploymentSummary) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *ManagedAppPolicyDeploymentSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *ManagedAppPolicyDeploymentSummary) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### SetDisplayNameExplicitNull

`func (o *ManagedAppPolicyDeploymentSummary) SetDisplayNameExplicitNull(b bool)`

SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DisplayName value is set to nil even if false is passed
### GetConfigurationDeployedUserCount

`func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCount() int32`

GetConfigurationDeployedUserCount returns the ConfigurationDeployedUserCount field if non-nil, zero value otherwise.

### GetConfigurationDeployedUserCountOk

`func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCountOk() (int32, bool)`

GetConfigurationDeployedUserCountOk returns a tuple with the ConfigurationDeployedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationDeployedUserCount

`func (o *ManagedAppPolicyDeploymentSummary) HasConfigurationDeployedUserCount() bool`

HasConfigurationDeployedUserCount returns a boolean if a field has been set.

### SetConfigurationDeployedUserCount

`func (o *ManagedAppPolicyDeploymentSummary) SetConfigurationDeployedUserCount(v int32)`

SetConfigurationDeployedUserCount gets a reference to the given int32 and assigns it to the ConfigurationDeployedUserCount field.

### GetLastRefreshTime

`func (o *ManagedAppPolicyDeploymentSummary) GetLastRefreshTime() time.Time`

GetLastRefreshTime returns the LastRefreshTime field if non-nil, zero value otherwise.

### GetLastRefreshTimeOk

`func (o *ManagedAppPolicyDeploymentSummary) GetLastRefreshTimeOk() (time.Time, bool)`

GetLastRefreshTimeOk returns a tuple with the LastRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastRefreshTime

`func (o *ManagedAppPolicyDeploymentSummary) HasLastRefreshTime() bool`

HasLastRefreshTime returns a boolean if a field has been set.

### SetLastRefreshTime

`func (o *ManagedAppPolicyDeploymentSummary) SetLastRefreshTime(v time.Time)`

SetLastRefreshTime gets a reference to the given time.Time and assigns it to the LastRefreshTime field.

### GetConfigurationDeploymentSummaryPerApp

`func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerApp() []AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp`

GetConfigurationDeploymentSummaryPerApp returns the ConfigurationDeploymentSummaryPerApp field if non-nil, zero value otherwise.

### GetConfigurationDeploymentSummaryPerAppOk

`func (o *ManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerAppOk() ([]AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp, bool)`

GetConfigurationDeploymentSummaryPerAppOk returns a tuple with the ConfigurationDeploymentSummaryPerApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfigurationDeploymentSummaryPerApp

`func (o *ManagedAppPolicyDeploymentSummary) HasConfigurationDeploymentSummaryPerApp() bool`

HasConfigurationDeploymentSummaryPerApp returns a boolean if a field has been set.

### SetConfigurationDeploymentSummaryPerApp

`func (o *ManagedAppPolicyDeploymentSummary) SetConfigurationDeploymentSummaryPerApp(v []AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp)`

SetConfigurationDeploymentSummaryPerApp gets a reference to the given []AnyOfmicrosoftGraphManagedAppPolicyDeploymentSummaryPerApp and assigns it to the ConfigurationDeploymentSummaryPerApp field.

### GetVersion

`func (o *ManagedAppPolicyDeploymentSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedAppPolicyDeploymentSummary) GetVersionOk() (string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *ManagedAppPolicyDeploymentSummary) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *ManagedAppPolicyDeploymentSummary) SetVersion(v string)`

SetVersion gets a reference to the given string and assigns it to the Version field.

### SetVersionExplicitNull

`func (o *ManagedAppPolicyDeploymentSummary) SetVersionExplicitNull(b bool)`

SetVersionExplicitNull (un)sets Version to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Version value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



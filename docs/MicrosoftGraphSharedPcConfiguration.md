# MicrosoftGraphSharedPcConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was created. | [optional] 
**Description** | Pointer to **string** | Admin provided description of the Device Configuration. | [optional] 
**DisplayName** | Pointer to **string** | Admin provided name of the device configuration. | [optional] 
**Version** | Pointer to **int32** | Version of the device configuration. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphDeviceConfigurationAssignment**](microsoft.graph.deviceConfigurationAssignment.md) |  | [optional] 
**DeviceStatuses** | Pointer to [**[]MicrosoftGraphDeviceConfigurationDeviceStatus**](microsoft.graph.deviceConfigurationDeviceStatus.md) |  | [optional] 
**UserStatuses** | Pointer to [**[]MicrosoftGraphDeviceConfigurationUserStatus**](microsoft.graph.deviceConfigurationUserStatus.md) |  | [optional] 
**DeviceStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview**](anyOf&lt;microsoft.graph.deviceConfigurationDeviceOverview&gt;.md) |  | [optional] 
**UserStatusOverview** | Pointer to [**AnyOfmicrosoftGraphDeviceConfigurationUserOverview**](anyOf&lt;microsoft.graph.deviceConfigurationUserOverview&gt;.md) |  | [optional] 
**DeviceSettingStateSummaries** | Pointer to [**[]MicrosoftGraphSettingStateDeviceSummary**](microsoft.graph.settingStateDeviceSummary.md) |  | [optional] 
**AccountManagerPolicy** | Pointer to [**AnyOfmicrosoftGraphSharedPcAccountManagerPolicy**](anyOf&lt;microsoft.graph.sharedPCAccountManagerPolicy&gt;.md) | Specifies how accounts are managed on a shared PC. Only applies when disableAccountManager is false. | [optional] 
**AllowedAccounts** | Pointer to [**AnyOfmicrosoftGraphSharedPcAllowedAccountType**](anyOf&lt;microsoft.graph.sharedPCAllowedAccountType&gt;.md) | Indicates which type of accounts are allowed to use on a shared PC. | [optional] 
**AllowLocalStorage** | Pointer to **bool** | Specifies whether local storage is allowed on a shared PC. | [optional] 
**DisableAccountManager** | Pointer to **bool** | Disables the account manager for shared PC mode. | [optional] 
**DisableEduPolicies** | Pointer to **bool** | Specifies whether the default shared PC education environment policies should be disabled. For Windows 10 RS2 and later, this policy will be applied without setting Enabled to true. | [optional] 
**DisablePowerPolicies** | Pointer to **bool** | Specifies whether the default shared PC power policies should be disabled. | [optional] 
**DisableSignInOnResume** | Pointer to **bool** | Disables the requirement to sign in whenever the device wakes up from sleep mode. | [optional] 
**Enabled** | Pointer to **bool** | Enables shared PC mode and applies the shared pc policies. | [optional] 
**IdleTimeBeforeSleepInSeconds** | Pointer to **int32** | Specifies the time in seconds that a device must sit idle before the PC goes to sleep. Setting this value to 0 prevents the sleep timeout from occurring. | [optional] 
**KioskAppDisplayName** | Pointer to **string** | Specifies the display text for the account shown on the sign-in screen which launches the app specified by SetKioskAppUserModelId. Only applies when KioskAppUserModelId is set. | [optional] 
**KioskAppUserModelId** | Pointer to **string** | Specifies the application user model ID of the app to use with assigned access. | [optional] 
**MaintenanceStartTime** | Pointer to **string** | Specifies the daily start time of maintenance hour. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphSharedPcConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphSharedPcConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphSharedPcConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphSharedPcConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphSharedPcConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphSharedPcConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphSharedPcConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphSharedPcConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSharedPcConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphSharedPcConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphSharedPcConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphSharedPcConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphSharedPcConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphSharedPcConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphSharedPcConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphSharedPcConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphSharedPcConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphSharedPcConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphSharedPcConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphSharedPcConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphSharedPcConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphSharedPcConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphSharedPcConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphSharedPcConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphSharedPcConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphSharedPcConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphSharedPcConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphSharedPcConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphSharedPcConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphSharedPcConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphSharedPcConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphSharedPcConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphSharedPcConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphSharedPcConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphSharedPcConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphSharedPcConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphSharedPcConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphSharedPcConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAccountManagerPolicy

`func (o *MicrosoftGraphSharedPcConfiguration) GetAccountManagerPolicy() AnyOfmicrosoftGraphSharedPcAccountManagerPolicy`

GetAccountManagerPolicy returns the AccountManagerPolicy field if non-nil, zero value otherwise.

### GetAccountManagerPolicyOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetAccountManagerPolicyOk() (AnyOfmicrosoftGraphSharedPcAccountManagerPolicy, bool)`

GetAccountManagerPolicyOk returns a tuple with the AccountManagerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountManagerPolicy

`func (o *MicrosoftGraphSharedPcConfiguration) HasAccountManagerPolicy() bool`

HasAccountManagerPolicy returns a boolean if a field has been set.

### SetAccountManagerPolicy

`func (o *MicrosoftGraphSharedPcConfiguration) SetAccountManagerPolicy(v AnyOfmicrosoftGraphSharedPcAccountManagerPolicy)`

SetAccountManagerPolicy gets a reference to the given AnyOfmicrosoftGraphSharedPcAccountManagerPolicy and assigns it to the AccountManagerPolicy field.

### SetAccountManagerPolicyExplicitNull

`func (o *MicrosoftGraphSharedPcConfiguration) SetAccountManagerPolicyExplicitNull(b bool)`

SetAccountManagerPolicyExplicitNull (un)sets AccountManagerPolicy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountManagerPolicy value is set to nil even if false is passed
### GetAllowedAccounts

`func (o *MicrosoftGraphSharedPcConfiguration) GetAllowedAccounts() AnyOfmicrosoftGraphSharedPcAllowedAccountType`

GetAllowedAccounts returns the AllowedAccounts field if non-nil, zero value otherwise.

### GetAllowedAccountsOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetAllowedAccountsOk() (AnyOfmicrosoftGraphSharedPcAllowedAccountType, bool)`

GetAllowedAccountsOk returns a tuple with the AllowedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedAccounts

`func (o *MicrosoftGraphSharedPcConfiguration) HasAllowedAccounts() bool`

HasAllowedAccounts returns a boolean if a field has been set.

### SetAllowedAccounts

`func (o *MicrosoftGraphSharedPcConfiguration) SetAllowedAccounts(v AnyOfmicrosoftGraphSharedPcAllowedAccountType)`

SetAllowedAccounts gets a reference to the given AnyOfmicrosoftGraphSharedPcAllowedAccountType and assigns it to the AllowedAccounts field.

### GetAllowLocalStorage

`func (o *MicrosoftGraphSharedPcConfiguration) GetAllowLocalStorage() bool`

GetAllowLocalStorage returns the AllowLocalStorage field if non-nil, zero value otherwise.

### GetAllowLocalStorageOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetAllowLocalStorageOk() (bool, bool)`

GetAllowLocalStorageOk returns a tuple with the AllowLocalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowLocalStorage

`func (o *MicrosoftGraphSharedPcConfiguration) HasAllowLocalStorage() bool`

HasAllowLocalStorage returns a boolean if a field has been set.

### SetAllowLocalStorage

`func (o *MicrosoftGraphSharedPcConfiguration) SetAllowLocalStorage(v bool)`

SetAllowLocalStorage gets a reference to the given bool and assigns it to the AllowLocalStorage field.

### GetDisableAccountManager

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisableAccountManager() bool`

GetDisableAccountManager returns the DisableAccountManager field if non-nil, zero value otherwise.

### GetDisableAccountManagerOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisableAccountManagerOk() (bool, bool)`

GetDisableAccountManagerOk returns a tuple with the DisableAccountManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAccountManager

`func (o *MicrosoftGraphSharedPcConfiguration) HasDisableAccountManager() bool`

HasDisableAccountManager returns a boolean if a field has been set.

### SetDisableAccountManager

`func (o *MicrosoftGraphSharedPcConfiguration) SetDisableAccountManager(v bool)`

SetDisableAccountManager gets a reference to the given bool and assigns it to the DisableAccountManager field.

### GetDisableEduPolicies

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisableEduPolicies() bool`

GetDisableEduPolicies returns the DisableEduPolicies field if non-nil, zero value otherwise.

### GetDisableEduPoliciesOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisableEduPoliciesOk() (bool, bool)`

GetDisableEduPoliciesOk returns a tuple with the DisableEduPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableEduPolicies

`func (o *MicrosoftGraphSharedPcConfiguration) HasDisableEduPolicies() bool`

HasDisableEduPolicies returns a boolean if a field has been set.

### SetDisableEduPolicies

`func (o *MicrosoftGraphSharedPcConfiguration) SetDisableEduPolicies(v bool)`

SetDisableEduPolicies gets a reference to the given bool and assigns it to the DisableEduPolicies field.

### GetDisablePowerPolicies

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisablePowerPolicies() bool`

GetDisablePowerPolicies returns the DisablePowerPolicies field if non-nil, zero value otherwise.

### GetDisablePowerPoliciesOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisablePowerPoliciesOk() (bool, bool)`

GetDisablePowerPoliciesOk returns a tuple with the DisablePowerPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisablePowerPolicies

`func (o *MicrosoftGraphSharedPcConfiguration) HasDisablePowerPolicies() bool`

HasDisablePowerPolicies returns a boolean if a field has been set.

### SetDisablePowerPolicies

`func (o *MicrosoftGraphSharedPcConfiguration) SetDisablePowerPolicies(v bool)`

SetDisablePowerPolicies gets a reference to the given bool and assigns it to the DisablePowerPolicies field.

### GetDisableSignInOnResume

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisableSignInOnResume() bool`

GetDisableSignInOnResume returns the DisableSignInOnResume field if non-nil, zero value otherwise.

### GetDisableSignInOnResumeOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetDisableSignInOnResumeOk() (bool, bool)`

GetDisableSignInOnResumeOk returns a tuple with the DisableSignInOnResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableSignInOnResume

`func (o *MicrosoftGraphSharedPcConfiguration) HasDisableSignInOnResume() bool`

HasDisableSignInOnResume returns a boolean if a field has been set.

### SetDisableSignInOnResume

`func (o *MicrosoftGraphSharedPcConfiguration) SetDisableSignInOnResume(v bool)`

SetDisableSignInOnResume gets a reference to the given bool and assigns it to the DisableSignInOnResume field.

### GetEnabled

`func (o *MicrosoftGraphSharedPcConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *MicrosoftGraphSharedPcConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *MicrosoftGraphSharedPcConfiguration) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetIdleTimeBeforeSleepInSeconds

`func (o *MicrosoftGraphSharedPcConfiguration) GetIdleTimeBeforeSleepInSeconds() int32`

GetIdleTimeBeforeSleepInSeconds returns the IdleTimeBeforeSleepInSeconds field if non-nil, zero value otherwise.

### GetIdleTimeBeforeSleepInSecondsOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetIdleTimeBeforeSleepInSecondsOk() (int32, bool)`

GetIdleTimeBeforeSleepInSecondsOk returns a tuple with the IdleTimeBeforeSleepInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdleTimeBeforeSleepInSeconds

`func (o *MicrosoftGraphSharedPcConfiguration) HasIdleTimeBeforeSleepInSeconds() bool`

HasIdleTimeBeforeSleepInSeconds returns a boolean if a field has been set.

### SetIdleTimeBeforeSleepInSeconds

`func (o *MicrosoftGraphSharedPcConfiguration) SetIdleTimeBeforeSleepInSeconds(v int32)`

SetIdleTimeBeforeSleepInSeconds gets a reference to the given int32 and assigns it to the IdleTimeBeforeSleepInSeconds field.

### SetIdleTimeBeforeSleepInSecondsExplicitNull

`func (o *MicrosoftGraphSharedPcConfiguration) SetIdleTimeBeforeSleepInSecondsExplicitNull(b bool)`

SetIdleTimeBeforeSleepInSecondsExplicitNull (un)sets IdleTimeBeforeSleepInSeconds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IdleTimeBeforeSleepInSeconds value is set to nil even if false is passed
### GetKioskAppDisplayName

`func (o *MicrosoftGraphSharedPcConfiguration) GetKioskAppDisplayName() string`

GetKioskAppDisplayName returns the KioskAppDisplayName field if non-nil, zero value otherwise.

### GetKioskAppDisplayNameOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetKioskAppDisplayNameOk() (string, bool)`

GetKioskAppDisplayNameOk returns a tuple with the KioskAppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskAppDisplayName

`func (o *MicrosoftGraphSharedPcConfiguration) HasKioskAppDisplayName() bool`

HasKioskAppDisplayName returns a boolean if a field has been set.

### SetKioskAppDisplayName

`func (o *MicrosoftGraphSharedPcConfiguration) SetKioskAppDisplayName(v string)`

SetKioskAppDisplayName gets a reference to the given string and assigns it to the KioskAppDisplayName field.

### SetKioskAppDisplayNameExplicitNull

`func (o *MicrosoftGraphSharedPcConfiguration) SetKioskAppDisplayNameExplicitNull(b bool)`

SetKioskAppDisplayNameExplicitNull (un)sets KioskAppDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskAppDisplayName value is set to nil even if false is passed
### GetKioskAppUserModelId

`func (o *MicrosoftGraphSharedPcConfiguration) GetKioskAppUserModelId() string`

GetKioskAppUserModelId returns the KioskAppUserModelId field if non-nil, zero value otherwise.

### GetKioskAppUserModelIdOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetKioskAppUserModelIdOk() (string, bool)`

GetKioskAppUserModelIdOk returns a tuple with the KioskAppUserModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskAppUserModelId

`func (o *MicrosoftGraphSharedPcConfiguration) HasKioskAppUserModelId() bool`

HasKioskAppUserModelId returns a boolean if a field has been set.

### SetKioskAppUserModelId

`func (o *MicrosoftGraphSharedPcConfiguration) SetKioskAppUserModelId(v string)`

SetKioskAppUserModelId gets a reference to the given string and assigns it to the KioskAppUserModelId field.

### SetKioskAppUserModelIdExplicitNull

`func (o *MicrosoftGraphSharedPcConfiguration) SetKioskAppUserModelIdExplicitNull(b bool)`

SetKioskAppUserModelIdExplicitNull (un)sets KioskAppUserModelId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskAppUserModelId value is set to nil even if false is passed
### GetMaintenanceStartTime

`func (o *MicrosoftGraphSharedPcConfiguration) GetMaintenanceStartTime() string`

GetMaintenanceStartTime returns the MaintenanceStartTime field if non-nil, zero value otherwise.

### GetMaintenanceStartTimeOk

`func (o *MicrosoftGraphSharedPcConfiguration) GetMaintenanceStartTimeOk() (string, bool)`

GetMaintenanceStartTimeOk returns a tuple with the MaintenanceStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaintenanceStartTime

`func (o *MicrosoftGraphSharedPcConfiguration) HasMaintenanceStartTime() bool`

HasMaintenanceStartTime returns a boolean if a field has been set.

### SetMaintenanceStartTime

`func (o *MicrosoftGraphSharedPcConfiguration) SetMaintenanceStartTime(v string)`

SetMaintenanceStartTime gets a reference to the given string and assigns it to the MaintenanceStartTime field.

### SetMaintenanceStartTimeExplicitNull

`func (o *MicrosoftGraphSharedPcConfiguration) SetMaintenanceStartTimeExplicitNull(b bool)`

SetMaintenanceStartTimeExplicitNull (un)sets MaintenanceStartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaintenanceStartTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



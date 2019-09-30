# SharedPcConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetAccountManagerPolicy

`func (o *SharedPcConfiguration) GetAccountManagerPolicy() AnyOfmicrosoftGraphSharedPcAccountManagerPolicy`

GetAccountManagerPolicy returns the AccountManagerPolicy field if non-nil, zero value otherwise.

### GetAccountManagerPolicyOk

`func (o *SharedPcConfiguration) GetAccountManagerPolicyOk() (AnyOfmicrosoftGraphSharedPcAccountManagerPolicy, bool)`

GetAccountManagerPolicyOk returns a tuple with the AccountManagerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountManagerPolicy

`func (o *SharedPcConfiguration) HasAccountManagerPolicy() bool`

HasAccountManagerPolicy returns a boolean if a field has been set.

### SetAccountManagerPolicy

`func (o *SharedPcConfiguration) SetAccountManagerPolicy(v AnyOfmicrosoftGraphSharedPcAccountManagerPolicy)`

SetAccountManagerPolicy gets a reference to the given AnyOfmicrosoftGraphSharedPcAccountManagerPolicy and assigns it to the AccountManagerPolicy field.

### SetAccountManagerPolicyExplicitNull

`func (o *SharedPcConfiguration) SetAccountManagerPolicyExplicitNull(b bool)`

SetAccountManagerPolicyExplicitNull (un)sets AccountManagerPolicy to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AccountManagerPolicy value is set to nil even if false is passed
### GetAllowedAccounts

`func (o *SharedPcConfiguration) GetAllowedAccounts() AnyOfmicrosoftGraphSharedPcAllowedAccountType`

GetAllowedAccounts returns the AllowedAccounts field if non-nil, zero value otherwise.

### GetAllowedAccountsOk

`func (o *SharedPcConfiguration) GetAllowedAccountsOk() (AnyOfmicrosoftGraphSharedPcAllowedAccountType, bool)`

GetAllowedAccountsOk returns a tuple with the AllowedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowedAccounts

`func (o *SharedPcConfiguration) HasAllowedAccounts() bool`

HasAllowedAccounts returns a boolean if a field has been set.

### SetAllowedAccounts

`func (o *SharedPcConfiguration) SetAllowedAccounts(v AnyOfmicrosoftGraphSharedPcAllowedAccountType)`

SetAllowedAccounts gets a reference to the given AnyOfmicrosoftGraphSharedPcAllowedAccountType and assigns it to the AllowedAccounts field.

### GetAllowLocalStorage

`func (o *SharedPcConfiguration) GetAllowLocalStorage() bool`

GetAllowLocalStorage returns the AllowLocalStorage field if non-nil, zero value otherwise.

### GetAllowLocalStorageOk

`func (o *SharedPcConfiguration) GetAllowLocalStorageOk() (bool, bool)`

GetAllowLocalStorageOk returns a tuple with the AllowLocalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowLocalStorage

`func (o *SharedPcConfiguration) HasAllowLocalStorage() bool`

HasAllowLocalStorage returns a boolean if a field has been set.

### SetAllowLocalStorage

`func (o *SharedPcConfiguration) SetAllowLocalStorage(v bool)`

SetAllowLocalStorage gets a reference to the given bool and assigns it to the AllowLocalStorage field.

### GetDisableAccountManager

`func (o *SharedPcConfiguration) GetDisableAccountManager() bool`

GetDisableAccountManager returns the DisableAccountManager field if non-nil, zero value otherwise.

### GetDisableAccountManagerOk

`func (o *SharedPcConfiguration) GetDisableAccountManagerOk() (bool, bool)`

GetDisableAccountManagerOk returns a tuple with the DisableAccountManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableAccountManager

`func (o *SharedPcConfiguration) HasDisableAccountManager() bool`

HasDisableAccountManager returns a boolean if a field has been set.

### SetDisableAccountManager

`func (o *SharedPcConfiguration) SetDisableAccountManager(v bool)`

SetDisableAccountManager gets a reference to the given bool and assigns it to the DisableAccountManager field.

### GetDisableEduPolicies

`func (o *SharedPcConfiguration) GetDisableEduPolicies() bool`

GetDisableEduPolicies returns the DisableEduPolicies field if non-nil, zero value otherwise.

### GetDisableEduPoliciesOk

`func (o *SharedPcConfiguration) GetDisableEduPoliciesOk() (bool, bool)`

GetDisableEduPoliciesOk returns a tuple with the DisableEduPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableEduPolicies

`func (o *SharedPcConfiguration) HasDisableEduPolicies() bool`

HasDisableEduPolicies returns a boolean if a field has been set.

### SetDisableEduPolicies

`func (o *SharedPcConfiguration) SetDisableEduPolicies(v bool)`

SetDisableEduPolicies gets a reference to the given bool and assigns it to the DisableEduPolicies field.

### GetDisablePowerPolicies

`func (o *SharedPcConfiguration) GetDisablePowerPolicies() bool`

GetDisablePowerPolicies returns the DisablePowerPolicies field if non-nil, zero value otherwise.

### GetDisablePowerPoliciesOk

`func (o *SharedPcConfiguration) GetDisablePowerPoliciesOk() (bool, bool)`

GetDisablePowerPoliciesOk returns a tuple with the DisablePowerPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisablePowerPolicies

`func (o *SharedPcConfiguration) HasDisablePowerPolicies() bool`

HasDisablePowerPolicies returns a boolean if a field has been set.

### SetDisablePowerPolicies

`func (o *SharedPcConfiguration) SetDisablePowerPolicies(v bool)`

SetDisablePowerPolicies gets a reference to the given bool and assigns it to the DisablePowerPolicies field.

### GetDisableSignInOnResume

`func (o *SharedPcConfiguration) GetDisableSignInOnResume() bool`

GetDisableSignInOnResume returns the DisableSignInOnResume field if non-nil, zero value otherwise.

### GetDisableSignInOnResumeOk

`func (o *SharedPcConfiguration) GetDisableSignInOnResumeOk() (bool, bool)`

GetDisableSignInOnResumeOk returns a tuple with the DisableSignInOnResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisableSignInOnResume

`func (o *SharedPcConfiguration) HasDisableSignInOnResume() bool`

HasDisableSignInOnResume returns a boolean if a field has been set.

### SetDisableSignInOnResume

`func (o *SharedPcConfiguration) SetDisableSignInOnResume(v bool)`

SetDisableSignInOnResume gets a reference to the given bool and assigns it to the DisableSignInOnResume field.

### GetEnabled

`func (o *SharedPcConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SharedPcConfiguration) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *SharedPcConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *SharedPcConfiguration) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetIdleTimeBeforeSleepInSeconds

`func (o *SharedPcConfiguration) GetIdleTimeBeforeSleepInSeconds() int32`

GetIdleTimeBeforeSleepInSeconds returns the IdleTimeBeforeSleepInSeconds field if non-nil, zero value otherwise.

### GetIdleTimeBeforeSleepInSecondsOk

`func (o *SharedPcConfiguration) GetIdleTimeBeforeSleepInSecondsOk() (int32, bool)`

GetIdleTimeBeforeSleepInSecondsOk returns a tuple with the IdleTimeBeforeSleepInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIdleTimeBeforeSleepInSeconds

`func (o *SharedPcConfiguration) HasIdleTimeBeforeSleepInSeconds() bool`

HasIdleTimeBeforeSleepInSeconds returns a boolean if a field has been set.

### SetIdleTimeBeforeSleepInSeconds

`func (o *SharedPcConfiguration) SetIdleTimeBeforeSleepInSeconds(v int32)`

SetIdleTimeBeforeSleepInSeconds gets a reference to the given int32 and assigns it to the IdleTimeBeforeSleepInSeconds field.

### SetIdleTimeBeforeSleepInSecondsExplicitNull

`func (o *SharedPcConfiguration) SetIdleTimeBeforeSleepInSecondsExplicitNull(b bool)`

SetIdleTimeBeforeSleepInSecondsExplicitNull (un)sets IdleTimeBeforeSleepInSeconds to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IdleTimeBeforeSleepInSeconds value is set to nil even if false is passed
### GetKioskAppDisplayName

`func (o *SharedPcConfiguration) GetKioskAppDisplayName() string`

GetKioskAppDisplayName returns the KioskAppDisplayName field if non-nil, zero value otherwise.

### GetKioskAppDisplayNameOk

`func (o *SharedPcConfiguration) GetKioskAppDisplayNameOk() (string, bool)`

GetKioskAppDisplayNameOk returns a tuple with the KioskAppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskAppDisplayName

`func (o *SharedPcConfiguration) HasKioskAppDisplayName() bool`

HasKioskAppDisplayName returns a boolean if a field has been set.

### SetKioskAppDisplayName

`func (o *SharedPcConfiguration) SetKioskAppDisplayName(v string)`

SetKioskAppDisplayName gets a reference to the given string and assigns it to the KioskAppDisplayName field.

### SetKioskAppDisplayNameExplicitNull

`func (o *SharedPcConfiguration) SetKioskAppDisplayNameExplicitNull(b bool)`

SetKioskAppDisplayNameExplicitNull (un)sets KioskAppDisplayName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskAppDisplayName value is set to nil even if false is passed
### GetKioskAppUserModelId

`func (o *SharedPcConfiguration) GetKioskAppUserModelId() string`

GetKioskAppUserModelId returns the KioskAppUserModelId field if non-nil, zero value otherwise.

### GetKioskAppUserModelIdOk

`func (o *SharedPcConfiguration) GetKioskAppUserModelIdOk() (string, bool)`

GetKioskAppUserModelIdOk returns a tuple with the KioskAppUserModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKioskAppUserModelId

`func (o *SharedPcConfiguration) HasKioskAppUserModelId() bool`

HasKioskAppUserModelId returns a boolean if a field has been set.

### SetKioskAppUserModelId

`func (o *SharedPcConfiguration) SetKioskAppUserModelId(v string)`

SetKioskAppUserModelId gets a reference to the given string and assigns it to the KioskAppUserModelId field.

### SetKioskAppUserModelIdExplicitNull

`func (o *SharedPcConfiguration) SetKioskAppUserModelIdExplicitNull(b bool)`

SetKioskAppUserModelIdExplicitNull (un)sets KioskAppUserModelId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The KioskAppUserModelId value is set to nil even if false is passed
### GetMaintenanceStartTime

`func (o *SharedPcConfiguration) GetMaintenanceStartTime() string`

GetMaintenanceStartTime returns the MaintenanceStartTime field if non-nil, zero value otherwise.

### GetMaintenanceStartTimeOk

`func (o *SharedPcConfiguration) GetMaintenanceStartTimeOk() (string, bool)`

GetMaintenanceStartTimeOk returns a tuple with the MaintenanceStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaintenanceStartTime

`func (o *SharedPcConfiguration) HasMaintenanceStartTime() bool`

HasMaintenanceStartTime returns a boolean if a field has been set.

### SetMaintenanceStartTime

`func (o *SharedPcConfiguration) SetMaintenanceStartTime(v string)`

SetMaintenanceStartTime gets a reference to the given string and assigns it to the MaintenanceStartTime field.

### SetMaintenanceStartTimeExplicitNull

`func (o *SharedPcConfiguration) SetMaintenanceStartTimeExplicitNull(b bool)`

SetMaintenanceStartTimeExplicitNull (un)sets MaintenanceStartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaintenanceStartTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



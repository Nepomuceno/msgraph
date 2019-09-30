# MicrosoftGraphSharedPcAccountManagerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountDeletionPolicy** | Pointer to [**AnyOfmicrosoftGraphSharedPcAccountDeletionPolicyType**](anyOf&lt;microsoft.graph.sharedPCAccountDeletionPolicyType&gt;.md) | Configures when accounts are deleted. | [optional] 
**CacheAccountsAboveDiskFreePercentage** | Pointer to **int32** | Sets the percentage of available disk space a PC should have before it stops deleting cached shared PC accounts. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100 | [optional] 
**InactiveThresholdDays** | Pointer to **int32** | Specifies when the accounts will start being deleted when they have not been logged on during the specified period, given as number of days. Only applies when AccountDeletionPolicy is DiskSpaceThreshold or DiskSpaceThresholdOrInactiveThreshold. | [optional] 
**RemoveAccountsBelowDiskFreePercentage** | Pointer to **int32** | Sets the percentage of disk space remaining on a PC before cached accounts will be deleted to free disk space. Accounts that have been inactive the longest will be deleted first. Only applies when AccountDeletionPolicy is DiskSpaceThresholdOrInactiveThreshold. Valid values 0 to 100 | [optional] 

## Methods

### GetAccountDeletionPolicy

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) GetAccountDeletionPolicy() AnyOfmicrosoftGraphSharedPcAccountDeletionPolicyType`

GetAccountDeletionPolicy returns the AccountDeletionPolicy field if non-nil, zero value otherwise.

### GetAccountDeletionPolicyOk

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) GetAccountDeletionPolicyOk() (AnyOfmicrosoftGraphSharedPcAccountDeletionPolicyType, bool)`

GetAccountDeletionPolicyOk returns a tuple with the AccountDeletionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountDeletionPolicy

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) HasAccountDeletionPolicy() bool`

HasAccountDeletionPolicy returns a boolean if a field has been set.

### SetAccountDeletionPolicy

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) SetAccountDeletionPolicy(v AnyOfmicrosoftGraphSharedPcAccountDeletionPolicyType)`

SetAccountDeletionPolicy gets a reference to the given AnyOfmicrosoftGraphSharedPcAccountDeletionPolicyType and assigns it to the AccountDeletionPolicy field.

### GetCacheAccountsAboveDiskFreePercentage

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) GetCacheAccountsAboveDiskFreePercentage() int32`

GetCacheAccountsAboveDiskFreePercentage returns the CacheAccountsAboveDiskFreePercentage field if non-nil, zero value otherwise.

### GetCacheAccountsAboveDiskFreePercentageOk

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) GetCacheAccountsAboveDiskFreePercentageOk() (int32, bool)`

GetCacheAccountsAboveDiskFreePercentageOk returns a tuple with the CacheAccountsAboveDiskFreePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCacheAccountsAboveDiskFreePercentage

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) HasCacheAccountsAboveDiskFreePercentage() bool`

HasCacheAccountsAboveDiskFreePercentage returns a boolean if a field has been set.

### SetCacheAccountsAboveDiskFreePercentage

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) SetCacheAccountsAboveDiskFreePercentage(v int32)`

SetCacheAccountsAboveDiskFreePercentage gets a reference to the given int32 and assigns it to the CacheAccountsAboveDiskFreePercentage field.

### SetCacheAccountsAboveDiskFreePercentageExplicitNull

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) SetCacheAccountsAboveDiskFreePercentageExplicitNull(b bool)`

SetCacheAccountsAboveDiskFreePercentageExplicitNull (un)sets CacheAccountsAboveDiskFreePercentage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CacheAccountsAboveDiskFreePercentage value is set to nil even if false is passed
### GetInactiveThresholdDays

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) GetInactiveThresholdDays() int32`

GetInactiveThresholdDays returns the InactiveThresholdDays field if non-nil, zero value otherwise.

### GetInactiveThresholdDaysOk

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) GetInactiveThresholdDaysOk() (int32, bool)`

GetInactiveThresholdDaysOk returns a tuple with the InactiveThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInactiveThresholdDays

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) HasInactiveThresholdDays() bool`

HasInactiveThresholdDays returns a boolean if a field has been set.

### SetInactiveThresholdDays

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) SetInactiveThresholdDays(v int32)`

SetInactiveThresholdDays gets a reference to the given int32 and assigns it to the InactiveThresholdDays field.

### SetInactiveThresholdDaysExplicitNull

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) SetInactiveThresholdDaysExplicitNull(b bool)`

SetInactiveThresholdDaysExplicitNull (un)sets InactiveThresholdDays to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InactiveThresholdDays value is set to nil even if false is passed
### GetRemoveAccountsBelowDiskFreePercentage

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) GetRemoveAccountsBelowDiskFreePercentage() int32`

GetRemoveAccountsBelowDiskFreePercentage returns the RemoveAccountsBelowDiskFreePercentage field if non-nil, zero value otherwise.

### GetRemoveAccountsBelowDiskFreePercentageOk

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) GetRemoveAccountsBelowDiskFreePercentageOk() (int32, bool)`

GetRemoveAccountsBelowDiskFreePercentageOk returns a tuple with the RemoveAccountsBelowDiskFreePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemoveAccountsBelowDiskFreePercentage

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) HasRemoveAccountsBelowDiskFreePercentage() bool`

HasRemoveAccountsBelowDiskFreePercentage returns a boolean if a field has been set.

### SetRemoveAccountsBelowDiskFreePercentage

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) SetRemoveAccountsBelowDiskFreePercentage(v int32)`

SetRemoveAccountsBelowDiskFreePercentage gets a reference to the given int32 and assigns it to the RemoveAccountsBelowDiskFreePercentage field.

### SetRemoveAccountsBelowDiskFreePercentageExplicitNull

`func (o *MicrosoftGraphSharedPcAccountManagerPolicy) SetRemoveAccountsBelowDiskFreePercentageExplicitNull(b bool)`

SetRemoveAccountsBelowDiskFreePercentageExplicitNull (un)sets RemoveAccountsBelowDiskFreePercentage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RemoveAccountsBelowDiskFreePercentage value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



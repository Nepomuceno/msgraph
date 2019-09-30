# MicrosoftGraphDeviceManagementSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceComplianceCheckinThresholdDays** | Pointer to **int32** | The number of days a device is allowed to go without checking in to remain compliant. Valid values 0 to 120 | [optional] 
**IsScheduledActionEnabled** | Pointer to **bool** | Is feature enabled or not for scheduled action for rule. | [optional] 
**SecureByDefault** | Pointer to **bool** | Device should be noncompliant when there is no compliance policy targeted when this is true | [optional] 

## Methods

### GetDeviceComplianceCheckinThresholdDays

`func (o *MicrosoftGraphDeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays() int32`

GetDeviceComplianceCheckinThresholdDays returns the DeviceComplianceCheckinThresholdDays field if non-nil, zero value otherwise.

### GetDeviceComplianceCheckinThresholdDaysOk

`func (o *MicrosoftGraphDeviceManagementSettings) GetDeviceComplianceCheckinThresholdDaysOk() (int32, bool)`

GetDeviceComplianceCheckinThresholdDaysOk returns a tuple with the DeviceComplianceCheckinThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceComplianceCheckinThresholdDays

`func (o *MicrosoftGraphDeviceManagementSettings) HasDeviceComplianceCheckinThresholdDays() bool`

HasDeviceComplianceCheckinThresholdDays returns a boolean if a field has been set.

### SetDeviceComplianceCheckinThresholdDays

`func (o *MicrosoftGraphDeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(v int32)`

SetDeviceComplianceCheckinThresholdDays gets a reference to the given int32 and assigns it to the DeviceComplianceCheckinThresholdDays field.

### GetIsScheduledActionEnabled

`func (o *MicrosoftGraphDeviceManagementSettings) GetIsScheduledActionEnabled() bool`

GetIsScheduledActionEnabled returns the IsScheduledActionEnabled field if non-nil, zero value otherwise.

### GetIsScheduledActionEnabledOk

`func (o *MicrosoftGraphDeviceManagementSettings) GetIsScheduledActionEnabledOk() (bool, bool)`

GetIsScheduledActionEnabledOk returns a tuple with the IsScheduledActionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsScheduledActionEnabled

`func (o *MicrosoftGraphDeviceManagementSettings) HasIsScheduledActionEnabled() bool`

HasIsScheduledActionEnabled returns a boolean if a field has been set.

### SetIsScheduledActionEnabled

`func (o *MicrosoftGraphDeviceManagementSettings) SetIsScheduledActionEnabled(v bool)`

SetIsScheduledActionEnabled gets a reference to the given bool and assigns it to the IsScheduledActionEnabled field.

### GetSecureByDefault

`func (o *MicrosoftGraphDeviceManagementSettings) GetSecureByDefault() bool`

GetSecureByDefault returns the SecureByDefault field if non-nil, zero value otherwise.

### GetSecureByDefaultOk

`func (o *MicrosoftGraphDeviceManagementSettings) GetSecureByDefaultOk() (bool, bool)`

GetSecureByDefaultOk returns a tuple with the SecureByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecureByDefault

`func (o *MicrosoftGraphDeviceManagementSettings) HasSecureByDefault() bool`

HasSecureByDefault returns a boolean if a field has been set.

### SetSecureByDefault

`func (o *MicrosoftGraphDeviceManagementSettings) SetSecureByDefault(v bool)`

SetSecureByDefault gets a reference to the given bool and assigns it to the SecureByDefault field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



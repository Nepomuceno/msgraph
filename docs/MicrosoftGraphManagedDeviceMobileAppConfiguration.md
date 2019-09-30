# MicrosoftGraphManagedDeviceMobileAppConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TargetedMobileApps** | Pointer to **[]string** | the associated app. | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was created. | [optional] 
**Description** | Pointer to **string** | Admin provided description of the Device Configuration. | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) | DateTime the object was last modified. | [optional] 
**DisplayName** | Pointer to **string** | Admin provided name of the device configuration. | [optional] 
**Version** | Pointer to **int32** | Version of the device configuration. | [optional] 
**Assignments** | Pointer to [**[]MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment**](microsoft.graph.managedDeviceMobileAppConfigurationAssignment.md) |  | [optional] 
**DeviceStatuses** | Pointer to [**[]MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus**](microsoft.graph.managedDeviceMobileAppConfigurationDeviceStatus.md) |  | [optional] 
**UserStatuses** | Pointer to [**[]MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus**](microsoft.graph.managedDeviceMobileAppConfigurationUserStatus.md) |  | [optional] 
**DeviceStatusSummary** | Pointer to [**AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary**](anyOf&lt;microsoft.graph.managedDeviceMobileAppConfigurationDeviceSummary&gt;.md) |  | [optional] 
**UserStatusSummary** | Pointer to [**AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary**](anyOf&lt;microsoft.graph.managedDeviceMobileAppConfigurationUserSummary&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetTargetedMobileApps

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetTargetedMobileApps() []string`

GetTargetedMobileApps returns the TargetedMobileApps field if non-nil, zero value otherwise.

### GetTargetedMobileAppsOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetTargetedMobileAppsOk() ([]string, bool)`

GetTargetedMobileAppsOk returns a tuple with the TargetedMobileApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetedMobileApps

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasTargetedMobileApps() bool`

HasTargetedMobileApps returns a boolean if a field has been set.

### SetTargetedMobileApps

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetTargetedMobileApps(v []string)`

SetTargetedMobileApps gets a reference to the given []string and assigns it to the TargetedMobileApps field.

### GetCreatedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetAssignments() []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetAssignmentsOk() ([]MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetAssignments(v []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDeviceStatuses() []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDeviceStatuses(v []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetUserStatuses() []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetUserStatusesOk() ([]MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetUserStatuses(v []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDeviceStatusSummary() AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary`

GetDeviceStatusSummary returns the DeviceStatusSummary field if non-nil, zero value otherwise.

### GetDeviceStatusSummaryOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetDeviceStatusSummaryOk() (AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary, bool)`

GetDeviceStatusSummaryOk returns a tuple with the DeviceStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasDeviceStatusSummary() bool`

HasDeviceStatusSummary returns a boolean if a field has been set.

### SetDeviceStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDeviceStatusSummary(v AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary)`

SetDeviceStatusSummary gets a reference to the given AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary and assigns it to the DeviceStatusSummary field.

### SetDeviceStatusSummaryExplicitNull

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetDeviceStatusSummaryExplicitNull(b bool)`

SetDeviceStatusSummaryExplicitNull (un)sets DeviceStatusSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusSummary value is set to nil even if false is passed
### GetUserStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetUserStatusSummary() AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary`

GetUserStatusSummary returns the UserStatusSummary field if non-nil, zero value otherwise.

### GetUserStatusSummaryOk

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) GetUserStatusSummaryOk() (AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary, bool)`

GetUserStatusSummaryOk returns a tuple with the UserStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) HasUserStatusSummary() bool`

HasUserStatusSummary returns a boolean if a field has been set.

### SetUserStatusSummary

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetUserStatusSummary(v AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary)`

SetUserStatusSummary gets a reference to the given AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary and assigns it to the UserStatusSummary field.

### SetUserStatusSummaryExplicitNull

`func (o *MicrosoftGraphManagedDeviceMobileAppConfiguration) SetUserStatusSummaryExplicitNull(b bool)`

SetUserStatusSummaryExplicitNull (un)sets UserStatusSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusSummary value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



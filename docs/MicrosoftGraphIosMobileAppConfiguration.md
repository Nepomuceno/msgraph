# MicrosoftGraphIosMobileAppConfiguration

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
**EncodedSettingXml** | Pointer to **string** | mdm app configuration Base64 binary. | [optional] 
**Settings** | Pointer to [**[]AnyOfmicrosoftGraphAppConfigurationSettingItem**](anyOf&lt;microsoft.graph.appConfigurationSettingItem&gt;.md) | app configuration setting items. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetTargetedMobileApps

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetTargetedMobileApps() []string`

GetTargetedMobileApps returns the TargetedMobileApps field if non-nil, zero value otherwise.

### GetTargetedMobileAppsOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetTargetedMobileAppsOk() ([]string, bool)`

GetTargetedMobileAppsOk returns a tuple with the TargetedMobileApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTargetedMobileApps

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasTargetedMobileApps() bool`

HasTargetedMobileApps returns a boolean if a field has been set.

### SetTargetedMobileApps

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetTargetedMobileApps(v []string)`

SetTargetedMobileApps gets a reference to the given []string and assigns it to the TargetedMobileApps field.

### GetCreatedDateTime

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetDisplayName

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetAssignments() []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetAssignmentsOk() ([]MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetAssignments(v []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphManagedDeviceMobileAppConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetDeviceStatuses() []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetDeviceStatuses(v []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphManagedDeviceMobileAppConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetUserStatuses() []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetUserStatusesOk() ([]MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetUserStatuses(v []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphManagedDeviceMobileAppConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusSummary

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetDeviceStatusSummary() AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary`

GetDeviceStatusSummary returns the DeviceStatusSummary field if non-nil, zero value otherwise.

### GetDeviceStatusSummaryOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetDeviceStatusSummaryOk() (AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary, bool)`

GetDeviceStatusSummaryOk returns a tuple with the DeviceStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusSummary

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasDeviceStatusSummary() bool`

HasDeviceStatusSummary returns a boolean if a field has been set.

### SetDeviceStatusSummary

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetDeviceStatusSummary(v AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary)`

SetDeviceStatusSummary gets a reference to the given AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationDeviceSummary and assigns it to the DeviceStatusSummary field.

### SetDeviceStatusSummaryExplicitNull

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetDeviceStatusSummaryExplicitNull(b bool)`

SetDeviceStatusSummaryExplicitNull (un)sets DeviceStatusSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusSummary value is set to nil even if false is passed
### GetUserStatusSummary

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetUserStatusSummary() AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary`

GetUserStatusSummary returns the UserStatusSummary field if non-nil, zero value otherwise.

### GetUserStatusSummaryOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetUserStatusSummaryOk() (AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary, bool)`

GetUserStatusSummaryOk returns a tuple with the UserStatusSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusSummary

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasUserStatusSummary() bool`

HasUserStatusSummary returns a boolean if a field has been set.

### SetUserStatusSummary

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetUserStatusSummary(v AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary)`

SetUserStatusSummary gets a reference to the given AnyOfmicrosoftGraphManagedDeviceMobileAppConfigurationUserSummary and assigns it to the UserStatusSummary field.

### SetUserStatusSummaryExplicitNull

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetUserStatusSummaryExplicitNull(b bool)`

SetUserStatusSummaryExplicitNull (un)sets UserStatusSummary to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusSummary value is set to nil even if false is passed
### GetEncodedSettingXml

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetEncodedSettingXml() string`

GetEncodedSettingXml returns the EncodedSettingXml field if non-nil, zero value otherwise.

### GetEncodedSettingXmlOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetEncodedSettingXmlOk() (string, bool)`

GetEncodedSettingXmlOk returns a tuple with the EncodedSettingXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEncodedSettingXml

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasEncodedSettingXml() bool`

HasEncodedSettingXml returns a boolean if a field has been set.

### SetEncodedSettingXml

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetEncodedSettingXml(v string)`

SetEncodedSettingXml gets a reference to the given string and assigns it to the EncodedSettingXml field.

### SetEncodedSettingXmlExplicitNull

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetEncodedSettingXmlExplicitNull(b bool)`

SetEncodedSettingXmlExplicitNull (un)sets EncodedSettingXml to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EncodedSettingXml value is set to nil even if false is passed
### GetSettings

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetSettings() []AnyOfmicrosoftGraphAppConfigurationSettingItem`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MicrosoftGraphIosMobileAppConfiguration) GetSettingsOk() ([]AnyOfmicrosoftGraphAppConfigurationSettingItem, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSettings

`func (o *MicrosoftGraphIosMobileAppConfiguration) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettings

`func (o *MicrosoftGraphIosMobileAppConfiguration) SetSettings(v []AnyOfmicrosoftGraphAppConfigurationSettingItem)`

SetSettings gets a reference to the given []AnyOfmicrosoftGraphAppConfigurationSettingItem and assigns it to the Settings field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



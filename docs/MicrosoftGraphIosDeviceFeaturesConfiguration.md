# MicrosoftGraphIosDeviceFeaturesConfiguration

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
**AssetTagTemplate** | Pointer to **string** | Asset tag information for the device, displayed on the login window and lock screen. | [optional] 
**LockScreenFootnote** | Pointer to **string** | A footnote displayed on the login window and lock screen. Available in iOS 9.3.1 and later. | [optional] 
**HomeScreenDockIcons** | Pointer to [**[]AnyOfmicrosoftGraphIosHomeScreenItem**](anyOf&lt;microsoft.graph.iosHomeScreenItem&gt;.md) | A list of app and folders to appear on the Home Screen Dock. This collection can contain a maximum of 500 elements. | [optional] 
**HomeScreenPages** | Pointer to [**[]AnyOfmicrosoftGraphIosHomeScreenPage**](anyOf&lt;microsoft.graph.iosHomeScreenPage&gt;.md) | A list of pages on the Home Screen. This collection can contain a maximum of 500 elements. | [optional] 
**NotificationSettings** | Pointer to [**[]AnyOfmicrosoftGraphIosNotificationSettings**](anyOf&lt;microsoft.graph.iosNotificationSettings&gt;.md) | Notification settings for each bundle id. Applicable to devices in supervised mode only (iOS 9.3 and later). This collection can contain a maximum of 500 elements. | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetAssetTagTemplate

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetAssetTagTemplate() string`

GetAssetTagTemplate returns the AssetTagTemplate field if non-nil, zero value otherwise.

### GetAssetTagTemplateOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetAssetTagTemplateOk() (string, bool)`

GetAssetTagTemplateOk returns a tuple with the AssetTagTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssetTagTemplate

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasAssetTagTemplate() bool`

HasAssetTagTemplate returns a boolean if a field has been set.

### SetAssetTagTemplate

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetAssetTagTemplate(v string)`

SetAssetTagTemplate gets a reference to the given string and assigns it to the AssetTagTemplate field.

### SetAssetTagTemplateExplicitNull

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetAssetTagTemplateExplicitNull(b bool)`

SetAssetTagTemplateExplicitNull (un)sets AssetTagTemplate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssetTagTemplate value is set to nil even if false is passed
### GetLockScreenFootnote

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetLockScreenFootnote() string`

GetLockScreenFootnote returns the LockScreenFootnote field if non-nil, zero value otherwise.

### GetLockScreenFootnoteOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetLockScreenFootnoteOk() (string, bool)`

GetLockScreenFootnoteOk returns a tuple with the LockScreenFootnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenFootnote

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasLockScreenFootnote() bool`

HasLockScreenFootnote returns a boolean if a field has been set.

### SetLockScreenFootnote

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetLockScreenFootnote(v string)`

SetLockScreenFootnote gets a reference to the given string and assigns it to the LockScreenFootnote field.

### SetLockScreenFootnoteExplicitNull

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetLockScreenFootnoteExplicitNull(b bool)`

SetLockScreenFootnoteExplicitNull (un)sets LockScreenFootnote to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LockScreenFootnote value is set to nil even if false is passed
### GetHomeScreenDockIcons

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetHomeScreenDockIcons() []AnyOfmicrosoftGraphIosHomeScreenItem`

GetHomeScreenDockIcons returns the HomeScreenDockIcons field if non-nil, zero value otherwise.

### GetHomeScreenDockIconsOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetHomeScreenDockIconsOk() ([]AnyOfmicrosoftGraphIosHomeScreenItem, bool)`

GetHomeScreenDockIconsOk returns a tuple with the HomeScreenDockIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomeScreenDockIcons

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasHomeScreenDockIcons() bool`

HasHomeScreenDockIcons returns a boolean if a field has been set.

### SetHomeScreenDockIcons

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetHomeScreenDockIcons(v []AnyOfmicrosoftGraphIosHomeScreenItem)`

SetHomeScreenDockIcons gets a reference to the given []AnyOfmicrosoftGraphIosHomeScreenItem and assigns it to the HomeScreenDockIcons field.

### GetHomeScreenPages

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetHomeScreenPages() []AnyOfmicrosoftGraphIosHomeScreenPage`

GetHomeScreenPages returns the HomeScreenPages field if non-nil, zero value otherwise.

### GetHomeScreenPagesOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetHomeScreenPagesOk() ([]AnyOfmicrosoftGraphIosHomeScreenPage, bool)`

GetHomeScreenPagesOk returns a tuple with the HomeScreenPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomeScreenPages

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasHomeScreenPages() bool`

HasHomeScreenPages returns a boolean if a field has been set.

### SetHomeScreenPages

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetHomeScreenPages(v []AnyOfmicrosoftGraphIosHomeScreenPage)`

SetHomeScreenPages gets a reference to the given []AnyOfmicrosoftGraphIosHomeScreenPage and assigns it to the HomeScreenPages field.

### GetNotificationSettings

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetNotificationSettings() []AnyOfmicrosoftGraphIosNotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) GetNotificationSettingsOk() ([]AnyOfmicrosoftGraphIosNotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationSettings

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### SetNotificationSettings

`func (o *MicrosoftGraphIosDeviceFeaturesConfiguration) SetNotificationSettings(v []AnyOfmicrosoftGraphIosNotificationSettings)`

SetNotificationSettings gets a reference to the given []AnyOfmicrosoftGraphIosNotificationSettings and assigns it to the NotificationSettings field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



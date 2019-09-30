# IosDeviceFeaturesConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetTagTemplate** | Pointer to **string** | Asset tag information for the device, displayed on the login window and lock screen. | [optional] 
**LockScreenFootnote** | Pointer to **string** | A footnote displayed on the login window and lock screen. Available in iOS 9.3.1 and later. | [optional] 
**HomeScreenDockIcons** | Pointer to [**[]AnyOfmicrosoftGraphIosHomeScreenItem**](anyOf&lt;microsoft.graph.iosHomeScreenItem&gt;.md) | A list of app and folders to appear on the Home Screen Dock. This collection can contain a maximum of 500 elements. | [optional] 
**HomeScreenPages** | Pointer to [**[]AnyOfmicrosoftGraphIosHomeScreenPage**](anyOf&lt;microsoft.graph.iosHomeScreenPage&gt;.md) | A list of pages on the Home Screen. This collection can contain a maximum of 500 elements. | [optional] 
**NotificationSettings** | Pointer to [**[]AnyOfmicrosoftGraphIosNotificationSettings**](anyOf&lt;microsoft.graph.iosNotificationSettings&gt;.md) | Notification settings for each bundle id. Applicable to devices in supervised mode only (iOS 9.3 and later). This collection can contain a maximum of 500 elements. | [optional] 

## Methods

### GetAssetTagTemplate

`func (o *IosDeviceFeaturesConfiguration) GetAssetTagTemplate() string`

GetAssetTagTemplate returns the AssetTagTemplate field if non-nil, zero value otherwise.

### GetAssetTagTemplateOk

`func (o *IosDeviceFeaturesConfiguration) GetAssetTagTemplateOk() (string, bool)`

GetAssetTagTemplateOk returns a tuple with the AssetTagTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssetTagTemplate

`func (o *IosDeviceFeaturesConfiguration) HasAssetTagTemplate() bool`

HasAssetTagTemplate returns a boolean if a field has been set.

### SetAssetTagTemplate

`func (o *IosDeviceFeaturesConfiguration) SetAssetTagTemplate(v string)`

SetAssetTagTemplate gets a reference to the given string and assigns it to the AssetTagTemplate field.

### SetAssetTagTemplateExplicitNull

`func (o *IosDeviceFeaturesConfiguration) SetAssetTagTemplateExplicitNull(b bool)`

SetAssetTagTemplateExplicitNull (un)sets AssetTagTemplate to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AssetTagTemplate value is set to nil even if false is passed
### GetLockScreenFootnote

`func (o *IosDeviceFeaturesConfiguration) GetLockScreenFootnote() string`

GetLockScreenFootnote returns the LockScreenFootnote field if non-nil, zero value otherwise.

### GetLockScreenFootnoteOk

`func (o *IosDeviceFeaturesConfiguration) GetLockScreenFootnoteOk() (string, bool)`

GetLockScreenFootnoteOk returns a tuple with the LockScreenFootnote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLockScreenFootnote

`func (o *IosDeviceFeaturesConfiguration) HasLockScreenFootnote() bool`

HasLockScreenFootnote returns a boolean if a field has been set.

### SetLockScreenFootnote

`func (o *IosDeviceFeaturesConfiguration) SetLockScreenFootnote(v string)`

SetLockScreenFootnote gets a reference to the given string and assigns it to the LockScreenFootnote field.

### SetLockScreenFootnoteExplicitNull

`func (o *IosDeviceFeaturesConfiguration) SetLockScreenFootnoteExplicitNull(b bool)`

SetLockScreenFootnoteExplicitNull (un)sets LockScreenFootnote to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LockScreenFootnote value is set to nil even if false is passed
### GetHomeScreenDockIcons

`func (o *IosDeviceFeaturesConfiguration) GetHomeScreenDockIcons() []AnyOfmicrosoftGraphIosHomeScreenItem`

GetHomeScreenDockIcons returns the HomeScreenDockIcons field if non-nil, zero value otherwise.

### GetHomeScreenDockIconsOk

`func (o *IosDeviceFeaturesConfiguration) GetHomeScreenDockIconsOk() ([]AnyOfmicrosoftGraphIosHomeScreenItem, bool)`

GetHomeScreenDockIconsOk returns a tuple with the HomeScreenDockIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomeScreenDockIcons

`func (o *IosDeviceFeaturesConfiguration) HasHomeScreenDockIcons() bool`

HasHomeScreenDockIcons returns a boolean if a field has been set.

### SetHomeScreenDockIcons

`func (o *IosDeviceFeaturesConfiguration) SetHomeScreenDockIcons(v []AnyOfmicrosoftGraphIosHomeScreenItem)`

SetHomeScreenDockIcons gets a reference to the given []AnyOfmicrosoftGraphIosHomeScreenItem and assigns it to the HomeScreenDockIcons field.

### GetHomeScreenPages

`func (o *IosDeviceFeaturesConfiguration) GetHomeScreenPages() []AnyOfmicrosoftGraphIosHomeScreenPage`

GetHomeScreenPages returns the HomeScreenPages field if non-nil, zero value otherwise.

### GetHomeScreenPagesOk

`func (o *IosDeviceFeaturesConfiguration) GetHomeScreenPagesOk() ([]AnyOfmicrosoftGraphIosHomeScreenPage, bool)`

GetHomeScreenPagesOk returns a tuple with the HomeScreenPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHomeScreenPages

`func (o *IosDeviceFeaturesConfiguration) HasHomeScreenPages() bool`

HasHomeScreenPages returns a boolean if a field has been set.

### SetHomeScreenPages

`func (o *IosDeviceFeaturesConfiguration) SetHomeScreenPages(v []AnyOfmicrosoftGraphIosHomeScreenPage)`

SetHomeScreenPages gets a reference to the given []AnyOfmicrosoftGraphIosHomeScreenPage and assigns it to the HomeScreenPages field.

### GetNotificationSettings

`func (o *IosDeviceFeaturesConfiguration) GetNotificationSettings() []AnyOfmicrosoftGraphIosNotificationSettings`

GetNotificationSettings returns the NotificationSettings field if non-nil, zero value otherwise.

### GetNotificationSettingsOk

`func (o *IosDeviceFeaturesConfiguration) GetNotificationSettingsOk() ([]AnyOfmicrosoftGraphIosNotificationSettings, bool)`

GetNotificationSettingsOk returns a tuple with the NotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationSettings

`func (o *IosDeviceFeaturesConfiguration) HasNotificationSettings() bool`

HasNotificationSettings returns a boolean if a field has been set.

### SetNotificationSettings

`func (o *IosDeviceFeaturesConfiguration) SetNotificationSettings(v []AnyOfmicrosoftGraphIosNotificationSettings)`

SetNotificationSettings gets a reference to the given []AnyOfmicrosoftGraphIosNotificationSettings and assigns it to the NotificationSettings field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



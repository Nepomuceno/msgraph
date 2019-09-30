# MicrosoftGraphIosNotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleID** | Pointer to **string** | Bundle id of app to which to apply these notification settings. | [optional] 
**AppName** | Pointer to **string** | Application name to be associated with the bundleID. | [optional] 
**Publisher** | Pointer to **string** | Publisher to be associated with the bundleID. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether notifications are allowed for this app. | [optional] 
**ShowInNotificationCenter** | Pointer to **bool** | Indicates whether notifications can be shown in notification center. | [optional] 
**ShowOnLockScreen** | Pointer to **bool** | Indicates whether notifications can be shown on the lock screen. | [optional] 
**AlertType** | Pointer to [**AnyOfmicrosoftGraphIosNotificationAlertType**](anyOf&lt;microsoft.graph.iosNotificationAlertType&gt;.md) | Indicates the type of alert for notifications for this app. | [optional] 
**BadgesEnabled** | Pointer to **bool** | Indicates whether badges are allowed for this app. | [optional] 
**SoundsEnabled** | Pointer to **bool** | Indicates whether sounds are allowed for this app. | [optional] 

## Methods

### GetBundleID

`func (o *MicrosoftGraphIosNotificationSettings) GetBundleID() string`

GetBundleID returns the BundleID field if non-nil, zero value otherwise.

### GetBundleIDOk

`func (o *MicrosoftGraphIosNotificationSettings) GetBundleIDOk() (string, bool)`

GetBundleIDOk returns a tuple with the BundleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBundleID

`func (o *MicrosoftGraphIosNotificationSettings) HasBundleID() bool`

HasBundleID returns a boolean if a field has been set.

### SetBundleID

`func (o *MicrosoftGraphIosNotificationSettings) SetBundleID(v string)`

SetBundleID gets a reference to the given string and assigns it to the BundleID field.

### GetAppName

`func (o *MicrosoftGraphIosNotificationSettings) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *MicrosoftGraphIosNotificationSettings) GetAppNameOk() (string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAppName

`func (o *MicrosoftGraphIosNotificationSettings) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppName

`func (o *MicrosoftGraphIosNotificationSettings) SetAppName(v string)`

SetAppName gets a reference to the given string and assigns it to the AppName field.

### SetAppNameExplicitNull

`func (o *MicrosoftGraphIosNotificationSettings) SetAppNameExplicitNull(b bool)`

SetAppNameExplicitNull (un)sets AppName to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AppName value is set to nil even if false is passed
### GetPublisher

`func (o *MicrosoftGraphIosNotificationSettings) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphIosNotificationSettings) GetPublisherOk() (string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublisher

`func (o *MicrosoftGraphIosNotificationSettings) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisher

`func (o *MicrosoftGraphIosNotificationSettings) SetPublisher(v string)`

SetPublisher gets a reference to the given string and assigns it to the Publisher field.

### SetPublisherExplicitNull

`func (o *MicrosoftGraphIosNotificationSettings) SetPublisherExplicitNull(b bool)`

SetPublisherExplicitNull (un)sets Publisher to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Publisher value is set to nil even if false is passed
### GetEnabled

`func (o *MicrosoftGraphIosNotificationSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MicrosoftGraphIosNotificationSettings) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *MicrosoftGraphIosNotificationSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *MicrosoftGraphIosNotificationSettings) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### SetEnabledExplicitNull

`func (o *MicrosoftGraphIosNotificationSettings) SetEnabledExplicitNull(b bool)`

SetEnabledExplicitNull (un)sets Enabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Enabled value is set to nil even if false is passed
### GetShowInNotificationCenter

`func (o *MicrosoftGraphIosNotificationSettings) GetShowInNotificationCenter() bool`

GetShowInNotificationCenter returns the ShowInNotificationCenter field if non-nil, zero value otherwise.

### GetShowInNotificationCenterOk

`func (o *MicrosoftGraphIosNotificationSettings) GetShowInNotificationCenterOk() (bool, bool)`

GetShowInNotificationCenterOk returns a tuple with the ShowInNotificationCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowInNotificationCenter

`func (o *MicrosoftGraphIosNotificationSettings) HasShowInNotificationCenter() bool`

HasShowInNotificationCenter returns a boolean if a field has been set.

### SetShowInNotificationCenter

`func (o *MicrosoftGraphIosNotificationSettings) SetShowInNotificationCenter(v bool)`

SetShowInNotificationCenter gets a reference to the given bool and assigns it to the ShowInNotificationCenter field.

### SetShowInNotificationCenterExplicitNull

`func (o *MicrosoftGraphIosNotificationSettings) SetShowInNotificationCenterExplicitNull(b bool)`

SetShowInNotificationCenterExplicitNull (un)sets ShowInNotificationCenter to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ShowInNotificationCenter value is set to nil even if false is passed
### GetShowOnLockScreen

`func (o *MicrosoftGraphIosNotificationSettings) GetShowOnLockScreen() bool`

GetShowOnLockScreen returns the ShowOnLockScreen field if non-nil, zero value otherwise.

### GetShowOnLockScreenOk

`func (o *MicrosoftGraphIosNotificationSettings) GetShowOnLockScreenOk() (bool, bool)`

GetShowOnLockScreenOk returns a tuple with the ShowOnLockScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowOnLockScreen

`func (o *MicrosoftGraphIosNotificationSettings) HasShowOnLockScreen() bool`

HasShowOnLockScreen returns a boolean if a field has been set.

### SetShowOnLockScreen

`func (o *MicrosoftGraphIosNotificationSettings) SetShowOnLockScreen(v bool)`

SetShowOnLockScreen gets a reference to the given bool and assigns it to the ShowOnLockScreen field.

### SetShowOnLockScreenExplicitNull

`func (o *MicrosoftGraphIosNotificationSettings) SetShowOnLockScreenExplicitNull(b bool)`

SetShowOnLockScreenExplicitNull (un)sets ShowOnLockScreen to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ShowOnLockScreen value is set to nil even if false is passed
### GetAlertType

`func (o *MicrosoftGraphIosNotificationSettings) GetAlertType() AnyOfmicrosoftGraphIosNotificationAlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *MicrosoftGraphIosNotificationSettings) GetAlertTypeOk() (AnyOfmicrosoftGraphIosNotificationAlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAlertType

`func (o *MicrosoftGraphIosNotificationSettings) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### SetAlertType

`func (o *MicrosoftGraphIosNotificationSettings) SetAlertType(v AnyOfmicrosoftGraphIosNotificationAlertType)`

SetAlertType gets a reference to the given AnyOfmicrosoftGraphIosNotificationAlertType and assigns it to the AlertType field.

### GetBadgesEnabled

`func (o *MicrosoftGraphIosNotificationSettings) GetBadgesEnabled() bool`

GetBadgesEnabled returns the BadgesEnabled field if non-nil, zero value otherwise.

### GetBadgesEnabledOk

`func (o *MicrosoftGraphIosNotificationSettings) GetBadgesEnabledOk() (bool, bool)`

GetBadgesEnabledOk returns a tuple with the BadgesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBadgesEnabled

`func (o *MicrosoftGraphIosNotificationSettings) HasBadgesEnabled() bool`

HasBadgesEnabled returns a boolean if a field has been set.

### SetBadgesEnabled

`func (o *MicrosoftGraphIosNotificationSettings) SetBadgesEnabled(v bool)`

SetBadgesEnabled gets a reference to the given bool and assigns it to the BadgesEnabled field.

### SetBadgesEnabledExplicitNull

`func (o *MicrosoftGraphIosNotificationSettings) SetBadgesEnabledExplicitNull(b bool)`

SetBadgesEnabledExplicitNull (un)sets BadgesEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BadgesEnabled value is set to nil even if false is passed
### GetSoundsEnabled

`func (o *MicrosoftGraphIosNotificationSettings) GetSoundsEnabled() bool`

GetSoundsEnabled returns the SoundsEnabled field if non-nil, zero value otherwise.

### GetSoundsEnabledOk

`func (o *MicrosoftGraphIosNotificationSettings) GetSoundsEnabledOk() (bool, bool)`

GetSoundsEnabledOk returns a tuple with the SoundsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSoundsEnabled

`func (o *MicrosoftGraphIosNotificationSettings) HasSoundsEnabled() bool`

HasSoundsEnabled returns a boolean if a field has been set.

### SetSoundsEnabled

`func (o *MicrosoftGraphIosNotificationSettings) SetSoundsEnabled(v bool)`

SetSoundsEnabled gets a reference to the given bool and assigns it to the SoundsEnabled field.

### SetSoundsEnabledExplicitNull

`func (o *MicrosoftGraphIosNotificationSettings) SetSoundsEnabledExplicitNull(b bool)`

SetSoundsEnabledExplicitNull (un)sets SoundsEnabled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SoundsEnabled value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



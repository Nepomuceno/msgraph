# MicrosoftGraphWindowsUpdateForBusinessConfiguration

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
**DeliveryOptimizationMode** | Pointer to [**AnyOfmicrosoftGraphWindowsDeliveryOptimizationMode**](anyOf&lt;microsoft.graph.windowsDeliveryOptimizationMode&gt;.md) | Delivery Optimization Mode | [optional] 
**PrereleaseFeatures** | Pointer to [**AnyOfmicrosoftGraphPrereleaseFeatures**](anyOf&lt;microsoft.graph.prereleaseFeatures&gt;.md) | The pre-release features. | [optional] 
**AutomaticUpdateMode** | Pointer to [**AnyOfmicrosoftGraphAutomaticUpdateMode**](anyOf&lt;microsoft.graph.automaticUpdateMode&gt;.md) | Automatic update mode. | [optional] 
**MicrosoftUpdateServiceAllowed** | Pointer to **bool** | Allow Microsoft Update Service | [optional] 
**DriversExcluded** | Pointer to **bool** | Exclude Windows update Drivers | [optional] 
**InstallationSchedule** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Installation schedule | [optional] 
**QualityUpdatesDeferralPeriodInDays** | Pointer to **int32** | Defer Quality Updates by these many days | [optional] 
**FeatureUpdatesDeferralPeriodInDays** | Pointer to **int32** | Defer Feature Updates by these many days | [optional] 
**QualityUpdatesPaused** | Pointer to **bool** | Pause Quality Updates | [optional] 
**FeatureUpdatesPaused** | Pointer to **bool** | Pause Feature Updates | [optional] 
**QualityUpdatesPauseExpiryDateTime** | Pointer to [**time.Time**](time.Time.md) | Quality Updates Pause Expiry datetime | [optional] 
**FeatureUpdatesPauseExpiryDateTime** | Pointer to [**time.Time**](time.Time.md) | Feature Updates Pause Expiry datetime | [optional] 
**BusinessReadyUpdatesOnly** | Pointer to [**AnyOfmicrosoftGraphWindowsUpdateType**](anyOf&lt;microsoft.graph.windowsUpdateType&gt;.md) | Determines which branch devices will receive their updates from | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### GetCreatedDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### GetDescription

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### SetDescriptionExplicitNull

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetDescriptionExplicitNull(b bool)`

SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Description value is set to nil even if false is passed
### GetDisplayName

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDisplayNameOk() (string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDisplayName

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetDisplayName(v string)`

SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.

### GetVersion

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetVersionOk() (int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetVersion(v int32)`

SetVersion gets a reference to the given int32 and assigns it to the Version field.

### GetAssignments

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignments

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### SetAssignments

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment)`

SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.

### GetDeviceStatuses

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus`

GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.

### GetDeviceStatusesOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool)`

GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatuses

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasDeviceStatuses() bool`

HasDeviceStatuses returns a boolean if a field has been set.

### SetDeviceStatuses

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus)`

SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.

### GetUserStatuses

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatuses

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.

### SetUserStatuses

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus)`

SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.

### GetDeviceStatusOverview

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview`

GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.

### GetDeviceStatusOverviewOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool)`

GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceStatusOverview

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasDeviceStatusOverview() bool`

HasDeviceStatusOverview returns a boolean if a field has been set.

### SetDeviceStatusOverview

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview)`

SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.

### SetDeviceStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetDeviceStatusOverviewExplicitNull(b bool)`

SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeviceStatusOverview value is set to nil even if false is passed
### GetUserStatusOverview

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview`

GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.

### GetUserStatusOverviewOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool)`

GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserStatusOverview

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasUserStatusOverview() bool`

HasUserStatusOverview returns a boolean if a field has been set.

### SetUserStatusOverview

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview)`

SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.

### SetUserStatusOverviewExplicitNull

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetUserStatusOverviewExplicitNull(b bool)`

SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UserStatusOverview value is set to nil even if false is passed
### GetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary`

GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.

### GetDeviceSettingStateSummariesOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool)`

GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasDeviceSettingStateSummaries() bool`

HasDeviceSettingStateSummaries returns a boolean if a field has been set.

### SetDeviceSettingStateSummaries

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary)`

SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.

### GetDeliveryOptimizationMode

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDeliveryOptimizationMode() AnyOfmicrosoftGraphWindowsDeliveryOptimizationMode`

GetDeliveryOptimizationMode returns the DeliveryOptimizationMode field if non-nil, zero value otherwise.

### GetDeliveryOptimizationModeOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDeliveryOptimizationModeOk() (AnyOfmicrosoftGraphWindowsDeliveryOptimizationMode, bool)`

GetDeliveryOptimizationModeOk returns a tuple with the DeliveryOptimizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeliveryOptimizationMode

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasDeliveryOptimizationMode() bool`

HasDeliveryOptimizationMode returns a boolean if a field has been set.

### SetDeliveryOptimizationMode

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetDeliveryOptimizationMode(v AnyOfmicrosoftGraphWindowsDeliveryOptimizationMode)`

SetDeliveryOptimizationMode gets a reference to the given AnyOfmicrosoftGraphWindowsDeliveryOptimizationMode and assigns it to the DeliveryOptimizationMode field.

### GetPrereleaseFeatures

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetPrereleaseFeatures() AnyOfmicrosoftGraphPrereleaseFeatures`

GetPrereleaseFeatures returns the PrereleaseFeatures field if non-nil, zero value otherwise.

### GetPrereleaseFeaturesOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetPrereleaseFeaturesOk() (AnyOfmicrosoftGraphPrereleaseFeatures, bool)`

GetPrereleaseFeaturesOk returns a tuple with the PrereleaseFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrereleaseFeatures

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasPrereleaseFeatures() bool`

HasPrereleaseFeatures returns a boolean if a field has been set.

### SetPrereleaseFeatures

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetPrereleaseFeatures(v AnyOfmicrosoftGraphPrereleaseFeatures)`

SetPrereleaseFeatures gets a reference to the given AnyOfmicrosoftGraphPrereleaseFeatures and assigns it to the PrereleaseFeatures field.

### GetAutomaticUpdateMode

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetAutomaticUpdateMode() AnyOfmicrosoftGraphAutomaticUpdateMode`

GetAutomaticUpdateMode returns the AutomaticUpdateMode field if non-nil, zero value otherwise.

### GetAutomaticUpdateModeOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetAutomaticUpdateModeOk() (AnyOfmicrosoftGraphAutomaticUpdateMode, bool)`

GetAutomaticUpdateModeOk returns a tuple with the AutomaticUpdateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutomaticUpdateMode

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasAutomaticUpdateMode() bool`

HasAutomaticUpdateMode returns a boolean if a field has been set.

### SetAutomaticUpdateMode

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetAutomaticUpdateMode(v AnyOfmicrosoftGraphAutomaticUpdateMode)`

SetAutomaticUpdateMode gets a reference to the given AnyOfmicrosoftGraphAutomaticUpdateMode and assigns it to the AutomaticUpdateMode field.

### GetMicrosoftUpdateServiceAllowed

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetMicrosoftUpdateServiceAllowed() bool`

GetMicrosoftUpdateServiceAllowed returns the MicrosoftUpdateServiceAllowed field if non-nil, zero value otherwise.

### GetMicrosoftUpdateServiceAllowedOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetMicrosoftUpdateServiceAllowedOk() (bool, bool)`

GetMicrosoftUpdateServiceAllowedOk returns a tuple with the MicrosoftUpdateServiceAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftUpdateServiceAllowed

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasMicrosoftUpdateServiceAllowed() bool`

HasMicrosoftUpdateServiceAllowed returns a boolean if a field has been set.

### SetMicrosoftUpdateServiceAllowed

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetMicrosoftUpdateServiceAllowed(v bool)`

SetMicrosoftUpdateServiceAllowed gets a reference to the given bool and assigns it to the MicrosoftUpdateServiceAllowed field.

### GetDriversExcluded

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDriversExcluded() bool`

GetDriversExcluded returns the DriversExcluded field if non-nil, zero value otherwise.

### GetDriversExcludedOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetDriversExcludedOk() (bool, bool)`

GetDriversExcludedOk returns a tuple with the DriversExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDriversExcluded

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasDriversExcluded() bool`

HasDriversExcluded returns a boolean if a field has been set.

### SetDriversExcluded

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetDriversExcluded(v bool)`

SetDriversExcluded gets a reference to the given bool and assigns it to the DriversExcluded field.

### GetInstallationSchedule

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetInstallationSchedule() AnyOfobject`

GetInstallationSchedule returns the InstallationSchedule field if non-nil, zero value otherwise.

### GetInstallationScheduleOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetInstallationScheduleOk() (AnyOfobject, bool)`

GetInstallationScheduleOk returns a tuple with the InstallationSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstallationSchedule

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasInstallationSchedule() bool`

HasInstallationSchedule returns a boolean if a field has been set.

### SetInstallationSchedule

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetInstallationSchedule(v AnyOfobject)`

SetInstallationSchedule gets a reference to the given AnyOfobject and assigns it to the InstallationSchedule field.

### SetInstallationScheduleExplicitNull

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetInstallationScheduleExplicitNull(b bool)`

SetInstallationScheduleExplicitNull (un)sets InstallationSchedule to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InstallationSchedule value is set to nil even if false is passed
### GetQualityUpdatesDeferralPeriodInDays

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetQualityUpdatesDeferralPeriodInDays() int32`

GetQualityUpdatesDeferralPeriodInDays returns the QualityUpdatesDeferralPeriodInDays field if non-nil, zero value otherwise.

### GetQualityUpdatesDeferralPeriodInDaysOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetQualityUpdatesDeferralPeriodInDaysOk() (int32, bool)`

GetQualityUpdatesDeferralPeriodInDaysOk returns a tuple with the QualityUpdatesDeferralPeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQualityUpdatesDeferralPeriodInDays

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasQualityUpdatesDeferralPeriodInDays() bool`

HasQualityUpdatesDeferralPeriodInDays returns a boolean if a field has been set.

### SetQualityUpdatesDeferralPeriodInDays

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetQualityUpdatesDeferralPeriodInDays(v int32)`

SetQualityUpdatesDeferralPeriodInDays gets a reference to the given int32 and assigns it to the QualityUpdatesDeferralPeriodInDays field.

### GetFeatureUpdatesDeferralPeriodInDays

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetFeatureUpdatesDeferralPeriodInDays() int32`

GetFeatureUpdatesDeferralPeriodInDays returns the FeatureUpdatesDeferralPeriodInDays field if non-nil, zero value otherwise.

### GetFeatureUpdatesDeferralPeriodInDaysOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetFeatureUpdatesDeferralPeriodInDaysOk() (int32, bool)`

GetFeatureUpdatesDeferralPeriodInDaysOk returns a tuple with the FeatureUpdatesDeferralPeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureUpdatesDeferralPeriodInDays

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasFeatureUpdatesDeferralPeriodInDays() bool`

HasFeatureUpdatesDeferralPeriodInDays returns a boolean if a field has been set.

### SetFeatureUpdatesDeferralPeriodInDays

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetFeatureUpdatesDeferralPeriodInDays(v int32)`

SetFeatureUpdatesDeferralPeriodInDays gets a reference to the given int32 and assigns it to the FeatureUpdatesDeferralPeriodInDays field.

### GetQualityUpdatesPaused

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetQualityUpdatesPaused() bool`

GetQualityUpdatesPaused returns the QualityUpdatesPaused field if non-nil, zero value otherwise.

### GetQualityUpdatesPausedOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetQualityUpdatesPausedOk() (bool, bool)`

GetQualityUpdatesPausedOk returns a tuple with the QualityUpdatesPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQualityUpdatesPaused

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasQualityUpdatesPaused() bool`

HasQualityUpdatesPaused returns a boolean if a field has been set.

### SetQualityUpdatesPaused

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetQualityUpdatesPaused(v bool)`

SetQualityUpdatesPaused gets a reference to the given bool and assigns it to the QualityUpdatesPaused field.

### GetFeatureUpdatesPaused

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPaused() bool`

GetFeatureUpdatesPaused returns the FeatureUpdatesPaused field if non-nil, zero value otherwise.

### GetFeatureUpdatesPausedOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPausedOk() (bool, bool)`

GetFeatureUpdatesPausedOk returns a tuple with the FeatureUpdatesPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureUpdatesPaused

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasFeatureUpdatesPaused() bool`

HasFeatureUpdatesPaused returns a boolean if a field has been set.

### SetFeatureUpdatesPaused

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetFeatureUpdatesPaused(v bool)`

SetFeatureUpdatesPaused gets a reference to the given bool and assigns it to the FeatureUpdatesPaused field.

### GetQualityUpdatesPauseExpiryDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetQualityUpdatesPauseExpiryDateTime() time.Time`

GetQualityUpdatesPauseExpiryDateTime returns the QualityUpdatesPauseExpiryDateTime field if non-nil, zero value otherwise.

### GetQualityUpdatesPauseExpiryDateTimeOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetQualityUpdatesPauseExpiryDateTimeOk() (time.Time, bool)`

GetQualityUpdatesPauseExpiryDateTimeOk returns a tuple with the QualityUpdatesPauseExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQualityUpdatesPauseExpiryDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasQualityUpdatesPauseExpiryDateTime() bool`

HasQualityUpdatesPauseExpiryDateTime returns a boolean if a field has been set.

### SetQualityUpdatesPauseExpiryDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetQualityUpdatesPauseExpiryDateTime(v time.Time)`

SetQualityUpdatesPauseExpiryDateTime gets a reference to the given time.Time and assigns it to the QualityUpdatesPauseExpiryDateTime field.

### GetFeatureUpdatesPauseExpiryDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPauseExpiryDateTime() time.Time`

GetFeatureUpdatesPauseExpiryDateTime returns the FeatureUpdatesPauseExpiryDateTime field if non-nil, zero value otherwise.

### GetFeatureUpdatesPauseExpiryDateTimeOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPauseExpiryDateTimeOk() (time.Time, bool)`

GetFeatureUpdatesPauseExpiryDateTimeOk returns a tuple with the FeatureUpdatesPauseExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureUpdatesPauseExpiryDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasFeatureUpdatesPauseExpiryDateTime() bool`

HasFeatureUpdatesPauseExpiryDateTime returns a boolean if a field has been set.

### SetFeatureUpdatesPauseExpiryDateTime

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetFeatureUpdatesPauseExpiryDateTime(v time.Time)`

SetFeatureUpdatesPauseExpiryDateTime gets a reference to the given time.Time and assigns it to the FeatureUpdatesPauseExpiryDateTime field.

### GetBusinessReadyUpdatesOnly

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetBusinessReadyUpdatesOnly() AnyOfmicrosoftGraphWindowsUpdateType`

GetBusinessReadyUpdatesOnly returns the BusinessReadyUpdatesOnly field if non-nil, zero value otherwise.

### GetBusinessReadyUpdatesOnlyOk

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) GetBusinessReadyUpdatesOnlyOk() (AnyOfmicrosoftGraphWindowsUpdateType, bool)`

GetBusinessReadyUpdatesOnlyOk returns a tuple with the BusinessReadyUpdatesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessReadyUpdatesOnly

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) HasBusinessReadyUpdatesOnly() bool`

HasBusinessReadyUpdatesOnly returns a boolean if a field has been set.

### SetBusinessReadyUpdatesOnly

`func (o *MicrosoftGraphWindowsUpdateForBusinessConfiguration) SetBusinessReadyUpdatesOnly(v AnyOfmicrosoftGraphWindowsUpdateType)`

SetBusinessReadyUpdatesOnly gets a reference to the given AnyOfmicrosoftGraphWindowsUpdateType and assigns it to the BusinessReadyUpdatesOnly field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



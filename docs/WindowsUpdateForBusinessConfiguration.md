# WindowsUpdateForBusinessConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### GetDeliveryOptimizationMode

`func (o *WindowsUpdateForBusinessConfiguration) GetDeliveryOptimizationMode() AnyOfmicrosoftGraphWindowsDeliveryOptimizationMode`

GetDeliveryOptimizationMode returns the DeliveryOptimizationMode field if non-nil, zero value otherwise.

### GetDeliveryOptimizationModeOk

`func (o *WindowsUpdateForBusinessConfiguration) GetDeliveryOptimizationModeOk() (AnyOfmicrosoftGraphWindowsDeliveryOptimizationMode, bool)`

GetDeliveryOptimizationModeOk returns a tuple with the DeliveryOptimizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeliveryOptimizationMode

`func (o *WindowsUpdateForBusinessConfiguration) HasDeliveryOptimizationMode() bool`

HasDeliveryOptimizationMode returns a boolean if a field has been set.

### SetDeliveryOptimizationMode

`func (o *WindowsUpdateForBusinessConfiguration) SetDeliveryOptimizationMode(v AnyOfmicrosoftGraphWindowsDeliveryOptimizationMode)`

SetDeliveryOptimizationMode gets a reference to the given AnyOfmicrosoftGraphWindowsDeliveryOptimizationMode and assigns it to the DeliveryOptimizationMode field.

### GetPrereleaseFeatures

`func (o *WindowsUpdateForBusinessConfiguration) GetPrereleaseFeatures() AnyOfmicrosoftGraphPrereleaseFeatures`

GetPrereleaseFeatures returns the PrereleaseFeatures field if non-nil, zero value otherwise.

### GetPrereleaseFeaturesOk

`func (o *WindowsUpdateForBusinessConfiguration) GetPrereleaseFeaturesOk() (AnyOfmicrosoftGraphPrereleaseFeatures, bool)`

GetPrereleaseFeaturesOk returns a tuple with the PrereleaseFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrereleaseFeatures

`func (o *WindowsUpdateForBusinessConfiguration) HasPrereleaseFeatures() bool`

HasPrereleaseFeatures returns a boolean if a field has been set.

### SetPrereleaseFeatures

`func (o *WindowsUpdateForBusinessConfiguration) SetPrereleaseFeatures(v AnyOfmicrosoftGraphPrereleaseFeatures)`

SetPrereleaseFeatures gets a reference to the given AnyOfmicrosoftGraphPrereleaseFeatures and assigns it to the PrereleaseFeatures field.

### GetAutomaticUpdateMode

`func (o *WindowsUpdateForBusinessConfiguration) GetAutomaticUpdateMode() AnyOfmicrosoftGraphAutomaticUpdateMode`

GetAutomaticUpdateMode returns the AutomaticUpdateMode field if non-nil, zero value otherwise.

### GetAutomaticUpdateModeOk

`func (o *WindowsUpdateForBusinessConfiguration) GetAutomaticUpdateModeOk() (AnyOfmicrosoftGraphAutomaticUpdateMode, bool)`

GetAutomaticUpdateModeOk returns a tuple with the AutomaticUpdateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutomaticUpdateMode

`func (o *WindowsUpdateForBusinessConfiguration) HasAutomaticUpdateMode() bool`

HasAutomaticUpdateMode returns a boolean if a field has been set.

### SetAutomaticUpdateMode

`func (o *WindowsUpdateForBusinessConfiguration) SetAutomaticUpdateMode(v AnyOfmicrosoftGraphAutomaticUpdateMode)`

SetAutomaticUpdateMode gets a reference to the given AnyOfmicrosoftGraphAutomaticUpdateMode and assigns it to the AutomaticUpdateMode field.

### GetMicrosoftUpdateServiceAllowed

`func (o *WindowsUpdateForBusinessConfiguration) GetMicrosoftUpdateServiceAllowed() bool`

GetMicrosoftUpdateServiceAllowed returns the MicrosoftUpdateServiceAllowed field if non-nil, zero value otherwise.

### GetMicrosoftUpdateServiceAllowedOk

`func (o *WindowsUpdateForBusinessConfiguration) GetMicrosoftUpdateServiceAllowedOk() (bool, bool)`

GetMicrosoftUpdateServiceAllowedOk returns a tuple with the MicrosoftUpdateServiceAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMicrosoftUpdateServiceAllowed

`func (o *WindowsUpdateForBusinessConfiguration) HasMicrosoftUpdateServiceAllowed() bool`

HasMicrosoftUpdateServiceAllowed returns a boolean if a field has been set.

### SetMicrosoftUpdateServiceAllowed

`func (o *WindowsUpdateForBusinessConfiguration) SetMicrosoftUpdateServiceAllowed(v bool)`

SetMicrosoftUpdateServiceAllowed gets a reference to the given bool and assigns it to the MicrosoftUpdateServiceAllowed field.

### GetDriversExcluded

`func (o *WindowsUpdateForBusinessConfiguration) GetDriversExcluded() bool`

GetDriversExcluded returns the DriversExcluded field if non-nil, zero value otherwise.

### GetDriversExcludedOk

`func (o *WindowsUpdateForBusinessConfiguration) GetDriversExcludedOk() (bool, bool)`

GetDriversExcludedOk returns a tuple with the DriversExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDriversExcluded

`func (o *WindowsUpdateForBusinessConfiguration) HasDriversExcluded() bool`

HasDriversExcluded returns a boolean if a field has been set.

### SetDriversExcluded

`func (o *WindowsUpdateForBusinessConfiguration) SetDriversExcluded(v bool)`

SetDriversExcluded gets a reference to the given bool and assigns it to the DriversExcluded field.

### GetInstallationSchedule

`func (o *WindowsUpdateForBusinessConfiguration) GetInstallationSchedule() AnyOfobject`

GetInstallationSchedule returns the InstallationSchedule field if non-nil, zero value otherwise.

### GetInstallationScheduleOk

`func (o *WindowsUpdateForBusinessConfiguration) GetInstallationScheduleOk() (AnyOfobject, bool)`

GetInstallationScheduleOk returns a tuple with the InstallationSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstallationSchedule

`func (o *WindowsUpdateForBusinessConfiguration) HasInstallationSchedule() bool`

HasInstallationSchedule returns a boolean if a field has been set.

### SetInstallationSchedule

`func (o *WindowsUpdateForBusinessConfiguration) SetInstallationSchedule(v AnyOfobject)`

SetInstallationSchedule gets a reference to the given AnyOfobject and assigns it to the InstallationSchedule field.

### SetInstallationScheduleExplicitNull

`func (o *WindowsUpdateForBusinessConfiguration) SetInstallationScheduleExplicitNull(b bool)`

SetInstallationScheduleExplicitNull (un)sets InstallationSchedule to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InstallationSchedule value is set to nil even if false is passed
### GetQualityUpdatesDeferralPeriodInDays

`func (o *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesDeferralPeriodInDays() int32`

GetQualityUpdatesDeferralPeriodInDays returns the QualityUpdatesDeferralPeriodInDays field if non-nil, zero value otherwise.

### GetQualityUpdatesDeferralPeriodInDaysOk

`func (o *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesDeferralPeriodInDaysOk() (int32, bool)`

GetQualityUpdatesDeferralPeriodInDaysOk returns a tuple with the QualityUpdatesDeferralPeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQualityUpdatesDeferralPeriodInDays

`func (o *WindowsUpdateForBusinessConfiguration) HasQualityUpdatesDeferralPeriodInDays() bool`

HasQualityUpdatesDeferralPeriodInDays returns a boolean if a field has been set.

### SetQualityUpdatesDeferralPeriodInDays

`func (o *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesDeferralPeriodInDays(v int32)`

SetQualityUpdatesDeferralPeriodInDays gets a reference to the given int32 and assigns it to the QualityUpdatesDeferralPeriodInDays field.

### GetFeatureUpdatesDeferralPeriodInDays

`func (o *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesDeferralPeriodInDays() int32`

GetFeatureUpdatesDeferralPeriodInDays returns the FeatureUpdatesDeferralPeriodInDays field if non-nil, zero value otherwise.

### GetFeatureUpdatesDeferralPeriodInDaysOk

`func (o *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesDeferralPeriodInDaysOk() (int32, bool)`

GetFeatureUpdatesDeferralPeriodInDaysOk returns a tuple with the FeatureUpdatesDeferralPeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureUpdatesDeferralPeriodInDays

`func (o *WindowsUpdateForBusinessConfiguration) HasFeatureUpdatesDeferralPeriodInDays() bool`

HasFeatureUpdatesDeferralPeriodInDays returns a boolean if a field has been set.

### SetFeatureUpdatesDeferralPeriodInDays

`func (o *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesDeferralPeriodInDays(v int32)`

SetFeatureUpdatesDeferralPeriodInDays gets a reference to the given int32 and assigns it to the FeatureUpdatesDeferralPeriodInDays field.

### GetQualityUpdatesPaused

`func (o *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesPaused() bool`

GetQualityUpdatesPaused returns the QualityUpdatesPaused field if non-nil, zero value otherwise.

### GetQualityUpdatesPausedOk

`func (o *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesPausedOk() (bool, bool)`

GetQualityUpdatesPausedOk returns a tuple with the QualityUpdatesPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQualityUpdatesPaused

`func (o *WindowsUpdateForBusinessConfiguration) HasQualityUpdatesPaused() bool`

HasQualityUpdatesPaused returns a boolean if a field has been set.

### SetQualityUpdatesPaused

`func (o *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesPaused(v bool)`

SetQualityUpdatesPaused gets a reference to the given bool and assigns it to the QualityUpdatesPaused field.

### GetFeatureUpdatesPaused

`func (o *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPaused() bool`

GetFeatureUpdatesPaused returns the FeatureUpdatesPaused field if non-nil, zero value otherwise.

### GetFeatureUpdatesPausedOk

`func (o *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPausedOk() (bool, bool)`

GetFeatureUpdatesPausedOk returns a tuple with the FeatureUpdatesPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureUpdatesPaused

`func (o *WindowsUpdateForBusinessConfiguration) HasFeatureUpdatesPaused() bool`

HasFeatureUpdatesPaused returns a boolean if a field has been set.

### SetFeatureUpdatesPaused

`func (o *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesPaused(v bool)`

SetFeatureUpdatesPaused gets a reference to the given bool and assigns it to the FeatureUpdatesPaused field.

### GetQualityUpdatesPauseExpiryDateTime

`func (o *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesPauseExpiryDateTime() time.Time`

GetQualityUpdatesPauseExpiryDateTime returns the QualityUpdatesPauseExpiryDateTime field if non-nil, zero value otherwise.

### GetQualityUpdatesPauseExpiryDateTimeOk

`func (o *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesPauseExpiryDateTimeOk() (time.Time, bool)`

GetQualityUpdatesPauseExpiryDateTimeOk returns a tuple with the QualityUpdatesPauseExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQualityUpdatesPauseExpiryDateTime

`func (o *WindowsUpdateForBusinessConfiguration) HasQualityUpdatesPauseExpiryDateTime() bool`

HasQualityUpdatesPauseExpiryDateTime returns a boolean if a field has been set.

### SetQualityUpdatesPauseExpiryDateTime

`func (o *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesPauseExpiryDateTime(v time.Time)`

SetQualityUpdatesPauseExpiryDateTime gets a reference to the given time.Time and assigns it to the QualityUpdatesPauseExpiryDateTime field.

### GetFeatureUpdatesPauseExpiryDateTime

`func (o *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPauseExpiryDateTime() time.Time`

GetFeatureUpdatesPauseExpiryDateTime returns the FeatureUpdatesPauseExpiryDateTime field if non-nil, zero value otherwise.

### GetFeatureUpdatesPauseExpiryDateTimeOk

`func (o *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPauseExpiryDateTimeOk() (time.Time, bool)`

GetFeatureUpdatesPauseExpiryDateTimeOk returns a tuple with the FeatureUpdatesPauseExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFeatureUpdatesPauseExpiryDateTime

`func (o *WindowsUpdateForBusinessConfiguration) HasFeatureUpdatesPauseExpiryDateTime() bool`

HasFeatureUpdatesPauseExpiryDateTime returns a boolean if a field has been set.

### SetFeatureUpdatesPauseExpiryDateTime

`func (o *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesPauseExpiryDateTime(v time.Time)`

SetFeatureUpdatesPauseExpiryDateTime gets a reference to the given time.Time and assigns it to the FeatureUpdatesPauseExpiryDateTime field.

### GetBusinessReadyUpdatesOnly

`func (o *WindowsUpdateForBusinessConfiguration) GetBusinessReadyUpdatesOnly() AnyOfmicrosoftGraphWindowsUpdateType`

GetBusinessReadyUpdatesOnly returns the BusinessReadyUpdatesOnly field if non-nil, zero value otherwise.

### GetBusinessReadyUpdatesOnlyOk

`func (o *WindowsUpdateForBusinessConfiguration) GetBusinessReadyUpdatesOnlyOk() (AnyOfmicrosoftGraphWindowsUpdateType, bool)`

GetBusinessReadyUpdatesOnlyOk returns a tuple with the BusinessReadyUpdatesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBusinessReadyUpdatesOnly

`func (o *WindowsUpdateForBusinessConfiguration) HasBusinessReadyUpdatesOnly() bool`

HasBusinessReadyUpdatesOnly returns a boolean if a field has been set.

### SetBusinessReadyUpdatesOnly

`func (o *WindowsUpdateForBusinessConfiguration) SetBusinessReadyUpdatesOnly(v AnyOfmicrosoftGraphWindowsUpdateType)`

SetBusinessReadyUpdatesOnly gets a reference to the given AnyOfmicrosoftGraphWindowsUpdateType and assigns it to the BusinessReadyUpdatesOnly field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



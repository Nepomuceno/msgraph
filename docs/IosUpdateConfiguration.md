# IosUpdateConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveHoursStart** | Pointer to **string** | Active Hours Start (active hours mean the time window when updates install should not happen) | [optional] 
**ActiveHoursEnd** | Pointer to **string** | Active Hours End (active hours mean the time window when updates install should not happen) | [optional] 
**ScheduledInstallDays** | Pointer to [**[]AnyOfmicrosoftGraphDayOfWeek**](anyOf&lt;microsoft.graph.dayOfWeek&gt;.md) | Days in week for which active hours are configured. This collection can contain a maximum of 7 elements. | [optional] 
**UtcTimeOffsetInMinutes** | Pointer to **int32** | UTC Time Offset indicated in minutes | [optional] 

## Methods

### GetActiveHoursStart

`func (o *IosUpdateConfiguration) GetActiveHoursStart() string`

GetActiveHoursStart returns the ActiveHoursStart field if non-nil, zero value otherwise.

### GetActiveHoursStartOk

`func (o *IosUpdateConfiguration) GetActiveHoursStartOk() (string, bool)`

GetActiveHoursStartOk returns a tuple with the ActiveHoursStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveHoursStart

`func (o *IosUpdateConfiguration) HasActiveHoursStart() bool`

HasActiveHoursStart returns a boolean if a field has been set.

### SetActiveHoursStart

`func (o *IosUpdateConfiguration) SetActiveHoursStart(v string)`

SetActiveHoursStart gets a reference to the given string and assigns it to the ActiveHoursStart field.

### GetActiveHoursEnd

`func (o *IosUpdateConfiguration) GetActiveHoursEnd() string`

GetActiveHoursEnd returns the ActiveHoursEnd field if non-nil, zero value otherwise.

### GetActiveHoursEndOk

`func (o *IosUpdateConfiguration) GetActiveHoursEndOk() (string, bool)`

GetActiveHoursEndOk returns a tuple with the ActiveHoursEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveHoursEnd

`func (o *IosUpdateConfiguration) HasActiveHoursEnd() bool`

HasActiveHoursEnd returns a boolean if a field has been set.

### SetActiveHoursEnd

`func (o *IosUpdateConfiguration) SetActiveHoursEnd(v string)`

SetActiveHoursEnd gets a reference to the given string and assigns it to the ActiveHoursEnd field.

### GetScheduledInstallDays

`func (o *IosUpdateConfiguration) GetScheduledInstallDays() []AnyOfmicrosoftGraphDayOfWeek`

GetScheduledInstallDays returns the ScheduledInstallDays field if non-nil, zero value otherwise.

### GetScheduledInstallDaysOk

`func (o *IosUpdateConfiguration) GetScheduledInstallDaysOk() ([]AnyOfmicrosoftGraphDayOfWeek, bool)`

GetScheduledInstallDaysOk returns a tuple with the ScheduledInstallDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledInstallDays

`func (o *IosUpdateConfiguration) HasScheduledInstallDays() bool`

HasScheduledInstallDays returns a boolean if a field has been set.

### SetScheduledInstallDays

`func (o *IosUpdateConfiguration) SetScheduledInstallDays(v []AnyOfmicrosoftGraphDayOfWeek)`

SetScheduledInstallDays gets a reference to the given []AnyOfmicrosoftGraphDayOfWeek and assigns it to the ScheduledInstallDays field.

### GetUtcTimeOffsetInMinutes

`func (o *IosUpdateConfiguration) GetUtcTimeOffsetInMinutes() int32`

GetUtcTimeOffsetInMinutes returns the UtcTimeOffsetInMinutes field if non-nil, zero value otherwise.

### GetUtcTimeOffsetInMinutesOk

`func (o *IosUpdateConfiguration) GetUtcTimeOffsetInMinutesOk() (int32, bool)`

GetUtcTimeOffsetInMinutesOk returns a tuple with the UtcTimeOffsetInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUtcTimeOffsetInMinutes

`func (o *IosUpdateConfiguration) HasUtcTimeOffsetInMinutes() bool`

HasUtcTimeOffsetInMinutes returns a boolean if a field has been set.

### SetUtcTimeOffsetInMinutes

`func (o *IosUpdateConfiguration) SetUtcTimeOffsetInMinutes(v int32)`

SetUtcTimeOffsetInMinutes gets a reference to the given int32 and assigns it to the UtcTimeOffsetInMinutes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



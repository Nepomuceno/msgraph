# MicrosoftGraphWindowsUpdateScheduledInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledInstallDay** | Pointer to [**AnyOfmicrosoftGraphWeeklySchedule**](anyOf&lt;microsoft.graph.weeklySchedule&gt;.md) | Scheduled Install Day in week | [optional] 
**ScheduledInstallTime** | Pointer to **string** | Scheduled Install Time during day | [optional] 

## Methods

### GetScheduledInstallDay

`func (o *MicrosoftGraphWindowsUpdateScheduledInstall) GetScheduledInstallDay() AnyOfmicrosoftGraphWeeklySchedule`

GetScheduledInstallDay returns the ScheduledInstallDay field if non-nil, zero value otherwise.

### GetScheduledInstallDayOk

`func (o *MicrosoftGraphWindowsUpdateScheduledInstall) GetScheduledInstallDayOk() (AnyOfmicrosoftGraphWeeklySchedule, bool)`

GetScheduledInstallDayOk returns a tuple with the ScheduledInstallDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledInstallDay

`func (o *MicrosoftGraphWindowsUpdateScheduledInstall) HasScheduledInstallDay() bool`

HasScheduledInstallDay returns a boolean if a field has been set.

### SetScheduledInstallDay

`func (o *MicrosoftGraphWindowsUpdateScheduledInstall) SetScheduledInstallDay(v AnyOfmicrosoftGraphWeeklySchedule)`

SetScheduledInstallDay gets a reference to the given AnyOfmicrosoftGraphWeeklySchedule and assigns it to the ScheduledInstallDay field.

### GetScheduledInstallTime

`func (o *MicrosoftGraphWindowsUpdateScheduledInstall) GetScheduledInstallTime() string`

GetScheduledInstallTime returns the ScheduledInstallTime field if non-nil, zero value otherwise.

### GetScheduledInstallTimeOk

`func (o *MicrosoftGraphWindowsUpdateScheduledInstall) GetScheduledInstallTimeOk() (string, bool)`

GetScheduledInstallTimeOk returns a tuple with the ScheduledInstallTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScheduledInstallTime

`func (o *MicrosoftGraphWindowsUpdateScheduledInstall) HasScheduledInstallTime() bool`

HasScheduledInstallTime returns a boolean if a field has been set.

### SetScheduledInstallTime

`func (o *MicrosoftGraphWindowsUpdateScheduledInstall) SetScheduledInstallTime(v string)`

SetScheduledInstallTime gets a reference to the given string and assigns it to the ScheduledInstallTime field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


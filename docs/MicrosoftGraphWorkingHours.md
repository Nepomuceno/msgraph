# MicrosoftGraphWorkingHours

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysOfWeek** | Pointer to [**[]AnyOfmicrosoftGraphDayOfWeek**](anyOf&lt;microsoft.graph.dayOfWeek&gt;.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to [**AnyOfmicrosoftGraphTimeZoneBase**](anyOf&lt;microsoft.graph.timeZoneBase&gt;.md) |  | [optional] 

## Methods

### GetDaysOfWeek

`func (o *MicrosoftGraphWorkingHours) GetDaysOfWeek() []AnyOfmicrosoftGraphDayOfWeek`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *MicrosoftGraphWorkingHours) GetDaysOfWeekOk() ([]AnyOfmicrosoftGraphDayOfWeek, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDaysOfWeek

`func (o *MicrosoftGraphWorkingHours) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.

### SetDaysOfWeek

`func (o *MicrosoftGraphWorkingHours) SetDaysOfWeek(v []AnyOfmicrosoftGraphDayOfWeek)`

SetDaysOfWeek gets a reference to the given []AnyOfmicrosoftGraphDayOfWeek and assigns it to the DaysOfWeek field.

### GetStartTime

`func (o *MicrosoftGraphWorkingHours) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MicrosoftGraphWorkingHours) GetStartTimeOk() (string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *MicrosoftGraphWorkingHours) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *MicrosoftGraphWorkingHours) SetStartTime(v string)`

SetStartTime gets a reference to the given string and assigns it to the StartTime field.

### SetStartTimeExplicitNull

`func (o *MicrosoftGraphWorkingHours) SetStartTimeExplicitNull(b bool)`

SetStartTimeExplicitNull (un)sets StartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartTime value is set to nil even if false is passed
### GetEndTime

`func (o *MicrosoftGraphWorkingHours) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MicrosoftGraphWorkingHours) GetEndTimeOk() (string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *MicrosoftGraphWorkingHours) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *MicrosoftGraphWorkingHours) SetEndTime(v string)`

SetEndTime gets a reference to the given string and assigns it to the EndTime field.

### SetEndTimeExplicitNull

`func (o *MicrosoftGraphWorkingHours) SetEndTimeExplicitNull(b bool)`

SetEndTimeExplicitNull (un)sets EndTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EndTime value is set to nil even if false is passed
### GetTimeZone

`func (o *MicrosoftGraphWorkingHours) GetTimeZone() AnyOfmicrosoftGraphTimeZoneBase`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MicrosoftGraphWorkingHours) GetTimeZoneOk() (AnyOfmicrosoftGraphTimeZoneBase, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeZone

`func (o *MicrosoftGraphWorkingHours) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZone

`func (o *MicrosoftGraphWorkingHours) SetTimeZone(v AnyOfmicrosoftGraphTimeZoneBase)`

SetTimeZone gets a reference to the given AnyOfmicrosoftGraphTimeZoneBase and assigns it to the TimeZone field.

### SetTimeZoneExplicitNull

`func (o *MicrosoftGraphWorkingHours) SetTimeZoneExplicitNull(b bool)`

SetTimeZoneExplicitNull (un)sets TimeZone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TimeZone value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



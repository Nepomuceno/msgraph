# MicrosoftGraphDaylightTimeZoneOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **string** |  | [optional] 
**DayOccurrence** | Pointer to **int32** |  | [optional] 
**DayOfWeek** | Pointer to [**AnyOfmicrosoftGraphDayOfWeek**](anyOf&lt;microsoft.graph.dayOfWeek&gt;.md) |  | [optional] 
**Month** | Pointer to **int32** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**DaylightBias** | Pointer to **int32** |  | [optional] 

## Methods

### GetTime

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetTimeOk() (string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *MicrosoftGraphDaylightTimeZoneOffset) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetTime(v string)`

SetTime gets a reference to the given string and assigns it to the Time field.

### SetTimeExplicitNull

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetTimeExplicitNull(b bool)`

SetTimeExplicitNull (un)sets Time to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Time value is set to nil even if false is passed
### GetDayOccurrence

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetDayOccurrence() int32`

GetDayOccurrence returns the DayOccurrence field if non-nil, zero value otherwise.

### GetDayOccurrenceOk

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetDayOccurrenceOk() (int32, bool)`

GetDayOccurrenceOk returns a tuple with the DayOccurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDayOccurrence

`func (o *MicrosoftGraphDaylightTimeZoneOffset) HasDayOccurrence() bool`

HasDayOccurrence returns a boolean if a field has been set.

### SetDayOccurrence

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetDayOccurrence(v int32)`

SetDayOccurrence gets a reference to the given int32 and assigns it to the DayOccurrence field.

### SetDayOccurrenceExplicitNull

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetDayOccurrenceExplicitNull(b bool)`

SetDayOccurrenceExplicitNull (un)sets DayOccurrence to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DayOccurrence value is set to nil even if false is passed
### GetDayOfWeek

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetDayOfWeek() AnyOfmicrosoftGraphDayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetDayOfWeekOk() (AnyOfmicrosoftGraphDayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDayOfWeek

`func (o *MicrosoftGraphDaylightTimeZoneOffset) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeek

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetDayOfWeek(v AnyOfmicrosoftGraphDayOfWeek)`

SetDayOfWeek gets a reference to the given AnyOfmicrosoftGraphDayOfWeek and assigns it to the DayOfWeek field.

### SetDayOfWeekExplicitNull

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetDayOfWeekExplicitNull(b bool)`

SetDayOfWeekExplicitNull (un)sets DayOfWeek to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DayOfWeek value is set to nil even if false is passed
### GetMonth

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetMonthOk() (int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMonth

`func (o *MicrosoftGraphDaylightTimeZoneOffset) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### SetMonth

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetMonth(v int32)`

SetMonth gets a reference to the given int32 and assigns it to the Month field.

### SetMonthExplicitNull

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetMonthExplicitNull(b bool)`

SetMonthExplicitNull (un)sets Month to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Month value is set to nil even if false is passed
### GetYear

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetYearOk() (int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasYear

`func (o *MicrosoftGraphDaylightTimeZoneOffset) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYear

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetYear(v int32)`

SetYear gets a reference to the given int32 and assigns it to the Year field.

### SetYearExplicitNull

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetYearExplicitNull(b bool)`

SetYearExplicitNull (un)sets Year to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Year value is set to nil even if false is passed
### GetDaylightBias

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetDaylightBias() int32`

GetDaylightBias returns the DaylightBias field if non-nil, zero value otherwise.

### GetDaylightBiasOk

`func (o *MicrosoftGraphDaylightTimeZoneOffset) GetDaylightBiasOk() (int32, bool)`

GetDaylightBiasOk returns a tuple with the DaylightBias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDaylightBias

`func (o *MicrosoftGraphDaylightTimeZoneOffset) HasDaylightBias() bool`

HasDaylightBias returns a boolean if a field has been set.

### SetDaylightBias

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetDaylightBias(v int32)`

SetDaylightBias gets a reference to the given int32 and assigns it to the DaylightBias field.

### SetDaylightBiasExplicitNull

`func (o *MicrosoftGraphDaylightTimeZoneOffset) SetDaylightBiasExplicitNull(b bool)`

SetDaylightBiasExplicitNull (un)sets DaylightBias to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DaylightBias value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



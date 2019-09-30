# MicrosoftGraphRecurrencePattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AnyOfmicrosoftGraphRecurrencePatternType**](anyOf&lt;microsoft.graph.recurrencePatternType&gt;.md) |  | [optional] 
**Interval** | Pointer to **int32** |  | [optional] 
**Month** | Pointer to **int32** |  | [optional] 
**DayOfMonth** | Pointer to **int32** |  | [optional] 
**DaysOfWeek** | Pointer to [**[]AnyOfmicrosoftGraphDayOfWeek**](anyOf&lt;microsoft.graph.dayOfWeek&gt;.md) |  | [optional] 
**FirstDayOfWeek** | Pointer to [**AnyOfmicrosoftGraphDayOfWeek**](anyOf&lt;microsoft.graph.dayOfWeek&gt;.md) |  | [optional] 
**Index** | Pointer to [**AnyOfmicrosoftGraphWeekIndex**](anyOf&lt;microsoft.graph.weekIndex&gt;.md) |  | [optional] 

## Methods

### GetType

`func (o *MicrosoftGraphRecurrencePattern) GetType() AnyOfmicrosoftGraphRecurrencePatternType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphRecurrencePattern) GetTypeOk() (AnyOfmicrosoftGraphRecurrencePatternType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *MicrosoftGraphRecurrencePattern) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *MicrosoftGraphRecurrencePattern) SetType(v AnyOfmicrosoftGraphRecurrencePatternType)`

SetType gets a reference to the given AnyOfmicrosoftGraphRecurrencePatternType and assigns it to the Type field.

### SetTypeExplicitNull

`func (o *MicrosoftGraphRecurrencePattern) SetTypeExplicitNull(b bool)`

SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Type value is set to nil even if false is passed
### GetInterval

`func (o *MicrosoftGraphRecurrencePattern) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MicrosoftGraphRecurrencePattern) GetIntervalOk() (int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInterval

`func (o *MicrosoftGraphRecurrencePattern) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetInterval

`func (o *MicrosoftGraphRecurrencePattern) SetInterval(v int32)`

SetInterval gets a reference to the given int32 and assigns it to the Interval field.

### GetMonth

`func (o *MicrosoftGraphRecurrencePattern) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *MicrosoftGraphRecurrencePattern) GetMonthOk() (int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMonth

`func (o *MicrosoftGraphRecurrencePattern) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### SetMonth

`func (o *MicrosoftGraphRecurrencePattern) SetMonth(v int32)`

SetMonth gets a reference to the given int32 and assigns it to the Month field.

### GetDayOfMonth

`func (o *MicrosoftGraphRecurrencePattern) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *MicrosoftGraphRecurrencePattern) GetDayOfMonthOk() (int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDayOfMonth

`func (o *MicrosoftGraphRecurrencePattern) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### SetDayOfMonth

`func (o *MicrosoftGraphRecurrencePattern) SetDayOfMonth(v int32)`

SetDayOfMonth gets a reference to the given int32 and assigns it to the DayOfMonth field.

### GetDaysOfWeek

`func (o *MicrosoftGraphRecurrencePattern) GetDaysOfWeek() []AnyOfmicrosoftGraphDayOfWeek`

GetDaysOfWeek returns the DaysOfWeek field if non-nil, zero value otherwise.

### GetDaysOfWeekOk

`func (o *MicrosoftGraphRecurrencePattern) GetDaysOfWeekOk() ([]AnyOfmicrosoftGraphDayOfWeek, bool)`

GetDaysOfWeekOk returns a tuple with the DaysOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDaysOfWeek

`func (o *MicrosoftGraphRecurrencePattern) HasDaysOfWeek() bool`

HasDaysOfWeek returns a boolean if a field has been set.

### SetDaysOfWeek

`func (o *MicrosoftGraphRecurrencePattern) SetDaysOfWeek(v []AnyOfmicrosoftGraphDayOfWeek)`

SetDaysOfWeek gets a reference to the given []AnyOfmicrosoftGraphDayOfWeek and assigns it to the DaysOfWeek field.

### GetFirstDayOfWeek

`func (o *MicrosoftGraphRecurrencePattern) GetFirstDayOfWeek() AnyOfmicrosoftGraphDayOfWeek`

GetFirstDayOfWeek returns the FirstDayOfWeek field if non-nil, zero value otherwise.

### GetFirstDayOfWeekOk

`func (o *MicrosoftGraphRecurrencePattern) GetFirstDayOfWeekOk() (AnyOfmicrosoftGraphDayOfWeek, bool)`

GetFirstDayOfWeekOk returns a tuple with the FirstDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirstDayOfWeek

`func (o *MicrosoftGraphRecurrencePattern) HasFirstDayOfWeek() bool`

HasFirstDayOfWeek returns a boolean if a field has been set.

### SetFirstDayOfWeek

`func (o *MicrosoftGraphRecurrencePattern) SetFirstDayOfWeek(v AnyOfmicrosoftGraphDayOfWeek)`

SetFirstDayOfWeek gets a reference to the given AnyOfmicrosoftGraphDayOfWeek and assigns it to the FirstDayOfWeek field.

### SetFirstDayOfWeekExplicitNull

`func (o *MicrosoftGraphRecurrencePattern) SetFirstDayOfWeekExplicitNull(b bool)`

SetFirstDayOfWeekExplicitNull (un)sets FirstDayOfWeek to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FirstDayOfWeek value is set to nil even if false is passed
### GetIndex

`func (o *MicrosoftGraphRecurrencePattern) GetIndex() AnyOfmicrosoftGraphWeekIndex`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MicrosoftGraphRecurrencePattern) GetIndexOk() (AnyOfmicrosoftGraphWeekIndex, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndex

`func (o *MicrosoftGraphRecurrencePattern) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndex

`func (o *MicrosoftGraphRecurrencePattern) SetIndex(v AnyOfmicrosoftGraphWeekIndex)`

SetIndex gets a reference to the given AnyOfmicrosoftGraphWeekIndex and assigns it to the Index field.

### SetIndexExplicitNull

`func (o *MicrosoftGraphRecurrencePattern) SetIndexExplicitNull(b bool)`

SetIndexExplicitNull (un)sets Index to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Index value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



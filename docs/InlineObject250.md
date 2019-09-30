# InlineObject250

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | Pointer to **[]string** |  | [optional] 
**EndTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**StartTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**AvailabilityViewInterval** | Pointer to **int32** |  | [optional] 

## Methods

### GetSchedules

`func (o *InlineObject250) GetSchedules() []string`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *InlineObject250) GetSchedulesOk() ([]string, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchedules

`func (o *InlineObject250) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### SetSchedules

`func (o *InlineObject250) SetSchedules(v []string)`

SetSchedules gets a reference to the given []string and assigns it to the Schedules field.

### GetEndTime

`func (o *InlineObject250) GetEndTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InlineObject250) GetEndTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEndTime

`func (o *InlineObject250) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTime

`func (o *InlineObject250) SetEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEndTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the EndTime field.

### SetEndTimeExplicitNull

`func (o *InlineObject250) SetEndTimeExplicitNull(b bool)`

SetEndTimeExplicitNull (un)sets EndTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EndTime value is set to nil even if false is passed
### GetStartTime

`func (o *InlineObject250) GetStartTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InlineObject250) GetStartTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartTime

`func (o *InlineObject250) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTime

`func (o *InlineObject250) SetStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetStartTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the StartTime field.

### SetStartTimeExplicitNull

`func (o *InlineObject250) SetStartTimeExplicitNull(b bool)`

SetStartTimeExplicitNull (un)sets StartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StartTime value is set to nil even if false is passed
### GetAvailabilityViewInterval

`func (o *InlineObject250) GetAvailabilityViewInterval() int32`

GetAvailabilityViewInterval returns the AvailabilityViewInterval field if non-nil, zero value otherwise.

### GetAvailabilityViewIntervalOk

`func (o *InlineObject250) GetAvailabilityViewIntervalOk() (int32, bool)`

GetAvailabilityViewIntervalOk returns a tuple with the AvailabilityViewInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvailabilityViewInterval

`func (o *InlineObject250) HasAvailabilityViewInterval() bool`

HasAvailabilityViewInterval returns a boolean if a field has been set.

### SetAvailabilityViewInterval

`func (o *InlineObject250) SetAvailabilityViewInterval(v int32)`

SetAvailabilityViewInterval gets a reference to the given int32 and assigns it to the AvailabilityViewInterval field.

### SetAvailabilityViewIntervalExplicitNull

`func (o *InlineObject250) SetAvailabilityViewIntervalExplicitNull(b bool)`

SetAvailabilityViewIntervalExplicitNull (un)sets AvailabilityViewInterval to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AvailabilityViewInterval value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



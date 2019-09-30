# MicrosoftGraphReminder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** |  | [optional] 
**EventStartTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**EventEndTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**ChangeKey** | Pointer to **string** |  | [optional] 
**EventSubject** | Pointer to **string** |  | [optional] 
**EventLocation** | Pointer to [**AnyOfmicrosoftGraphLocation**](anyOf&lt;microsoft.graph.location&gt;.md) |  | [optional] 
**EventWebLink** | Pointer to **string** |  | [optional] 
**ReminderFireTime** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 

## Methods

### GetEventId

`func (o *MicrosoftGraphReminder) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *MicrosoftGraphReminder) GetEventIdOk() (string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventId

`func (o *MicrosoftGraphReminder) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventId

`func (o *MicrosoftGraphReminder) SetEventId(v string)`

SetEventId gets a reference to the given string and assigns it to the EventId field.

### SetEventIdExplicitNull

`func (o *MicrosoftGraphReminder) SetEventIdExplicitNull(b bool)`

SetEventIdExplicitNull (un)sets EventId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EventId value is set to nil even if false is passed
### GetEventStartTime

`func (o *MicrosoftGraphReminder) GetEventStartTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEventStartTime returns the EventStartTime field if non-nil, zero value otherwise.

### GetEventStartTimeOk

`func (o *MicrosoftGraphReminder) GetEventStartTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEventStartTimeOk returns a tuple with the EventStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventStartTime

`func (o *MicrosoftGraphReminder) HasEventStartTime() bool`

HasEventStartTime returns a boolean if a field has been set.

### SetEventStartTime

`func (o *MicrosoftGraphReminder) SetEventStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEventStartTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the EventStartTime field.

### SetEventStartTimeExplicitNull

`func (o *MicrosoftGraphReminder) SetEventStartTimeExplicitNull(b bool)`

SetEventStartTimeExplicitNull (un)sets EventStartTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EventStartTime value is set to nil even if false is passed
### GetEventEndTime

`func (o *MicrosoftGraphReminder) GetEventEndTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEventEndTime returns the EventEndTime field if non-nil, zero value otherwise.

### GetEventEndTimeOk

`func (o *MicrosoftGraphReminder) GetEventEndTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEventEndTimeOk returns a tuple with the EventEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventEndTime

`func (o *MicrosoftGraphReminder) HasEventEndTime() bool`

HasEventEndTime returns a boolean if a field has been set.

### SetEventEndTime

`func (o *MicrosoftGraphReminder) SetEventEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEventEndTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the EventEndTime field.

### SetEventEndTimeExplicitNull

`func (o *MicrosoftGraphReminder) SetEventEndTimeExplicitNull(b bool)`

SetEventEndTimeExplicitNull (un)sets EventEndTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EventEndTime value is set to nil even if false is passed
### GetChangeKey

`func (o *MicrosoftGraphReminder) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphReminder) GetChangeKeyOk() (string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeKey

`func (o *MicrosoftGraphReminder) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKey

`func (o *MicrosoftGraphReminder) SetChangeKey(v string)`

SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.

### SetChangeKeyExplicitNull

`func (o *MicrosoftGraphReminder) SetChangeKeyExplicitNull(b bool)`

SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChangeKey value is set to nil even if false is passed
### GetEventSubject

`func (o *MicrosoftGraphReminder) GetEventSubject() string`

GetEventSubject returns the EventSubject field if non-nil, zero value otherwise.

### GetEventSubjectOk

`func (o *MicrosoftGraphReminder) GetEventSubjectOk() (string, bool)`

GetEventSubjectOk returns a tuple with the EventSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventSubject

`func (o *MicrosoftGraphReminder) HasEventSubject() bool`

HasEventSubject returns a boolean if a field has been set.

### SetEventSubject

`func (o *MicrosoftGraphReminder) SetEventSubject(v string)`

SetEventSubject gets a reference to the given string and assigns it to the EventSubject field.

### SetEventSubjectExplicitNull

`func (o *MicrosoftGraphReminder) SetEventSubjectExplicitNull(b bool)`

SetEventSubjectExplicitNull (un)sets EventSubject to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EventSubject value is set to nil even if false is passed
### GetEventLocation

`func (o *MicrosoftGraphReminder) GetEventLocation() AnyOfmicrosoftGraphLocation`

GetEventLocation returns the EventLocation field if non-nil, zero value otherwise.

### GetEventLocationOk

`func (o *MicrosoftGraphReminder) GetEventLocationOk() (AnyOfmicrosoftGraphLocation, bool)`

GetEventLocationOk returns a tuple with the EventLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventLocation

`func (o *MicrosoftGraphReminder) HasEventLocation() bool`

HasEventLocation returns a boolean if a field has been set.

### SetEventLocation

`func (o *MicrosoftGraphReminder) SetEventLocation(v AnyOfmicrosoftGraphLocation)`

SetEventLocation gets a reference to the given AnyOfmicrosoftGraphLocation and assigns it to the EventLocation field.

### SetEventLocationExplicitNull

`func (o *MicrosoftGraphReminder) SetEventLocationExplicitNull(b bool)`

SetEventLocationExplicitNull (un)sets EventLocation to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EventLocation value is set to nil even if false is passed
### GetEventWebLink

`func (o *MicrosoftGraphReminder) GetEventWebLink() string`

GetEventWebLink returns the EventWebLink field if non-nil, zero value otherwise.

### GetEventWebLinkOk

`func (o *MicrosoftGraphReminder) GetEventWebLinkOk() (string, bool)`

GetEventWebLinkOk returns a tuple with the EventWebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventWebLink

`func (o *MicrosoftGraphReminder) HasEventWebLink() bool`

HasEventWebLink returns a boolean if a field has been set.

### SetEventWebLink

`func (o *MicrosoftGraphReminder) SetEventWebLink(v string)`

SetEventWebLink gets a reference to the given string and assigns it to the EventWebLink field.

### SetEventWebLinkExplicitNull

`func (o *MicrosoftGraphReminder) SetEventWebLinkExplicitNull(b bool)`

SetEventWebLinkExplicitNull (un)sets EventWebLink to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EventWebLink value is set to nil even if false is passed
### GetReminderFireTime

`func (o *MicrosoftGraphReminder) GetReminderFireTime() AnyOfmicrosoftGraphDateTimeTimeZone`

GetReminderFireTime returns the ReminderFireTime field if non-nil, zero value otherwise.

### GetReminderFireTimeOk

`func (o *MicrosoftGraphReminder) GetReminderFireTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetReminderFireTimeOk returns a tuple with the ReminderFireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReminderFireTime

`func (o *MicrosoftGraphReminder) HasReminderFireTime() bool`

HasReminderFireTime returns a boolean if a field has been set.

### SetReminderFireTime

`func (o *MicrosoftGraphReminder) SetReminderFireTime(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetReminderFireTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the ReminderFireTime field.

### SetReminderFireTimeExplicitNull

`func (o *MicrosoftGraphReminder) SetReminderFireTimeExplicitNull(b bool)`

SetReminderFireTimeExplicitNull (un)sets ReminderFireTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReminderFireTime value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



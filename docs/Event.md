# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalStartTimeZone** | Pointer to **string** |  | [optional] 
**OriginalEndTimeZone** | Pointer to **string** |  | [optional] 
**ResponseStatus** | Pointer to [**AnyOfmicrosoftGraphResponseStatus**](anyOf&lt;microsoft.graph.responseStatus&gt;.md) |  | [optional] 
**ICalUId** | Pointer to **string** |  | [optional] 
**ReminderMinutesBeforeStart** | Pointer to **int32** |  | [optional] 
**IsReminderOn** | Pointer to **bool** |  | [optional] 
**HasAttachments** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Body** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) |  | [optional] 
**BodyPreview** | Pointer to **string** |  | [optional] 
**Importance** | Pointer to [**AnyOfmicrosoftGraphImportance**](anyOf&lt;microsoft.graph.importance&gt;.md) |  | [optional] 
**Sensitivity** | Pointer to [**AnyOfmicrosoftGraphSensitivity**](anyOf&lt;microsoft.graph.sensitivity&gt;.md) |  | [optional] 
**Start** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**OriginalStart** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**End** | Pointer to [**AnyOfmicrosoftGraphDateTimeTimeZone**](anyOf&lt;microsoft.graph.dateTimeTimeZone&gt;.md) |  | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphLocation**](anyOf&lt;microsoft.graph.location&gt;.md) |  | [optional] 
**Locations** | Pointer to [**[]AnyOfmicrosoftGraphLocation**](anyOf&lt;microsoft.graph.location&gt;.md) |  | [optional] 
**IsAllDay** | Pointer to **bool** |  | [optional] 
**IsCancelled** | Pointer to **bool** |  | [optional] 
**IsOrganizer** | Pointer to **bool** |  | [optional] 
**Recurrence** | Pointer to [**AnyOfmicrosoftGraphPatternedRecurrence**](anyOf&lt;microsoft.graph.patternedRecurrence&gt;.md) |  | [optional] 
**ResponseRequested** | Pointer to **bool** |  | [optional] 
**SeriesMasterId** | Pointer to **string** |  | [optional] 
**ShowAs** | Pointer to [**AnyOfmicrosoftGraphFreeBusyStatus**](anyOf&lt;microsoft.graph.freeBusyStatus&gt;.md) |  | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphEventType**](anyOf&lt;microsoft.graph.eventType&gt;.md) |  | [optional] 
**Attendees** | Pointer to [**[]AnyOfmicrosoftGraphAttendee**](anyOf&lt;microsoft.graph.attendee&gt;.md) |  | [optional] 
**Organizer** | Pointer to [**AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**WebLink** | Pointer to **string** |  | [optional] 
**OnlineMeetingUrl** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**[]MicrosoftGraphAttachment**](microsoft.graph.attachment.md) |  | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md) |  | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md) |  | [optional] 
**Calendar** | Pointer to [**AnyOfmicrosoftGraphCalendar**](anyOf&lt;microsoft.graph.calendar&gt;.md) |  | [optional] 
**Instances** | Pointer to [**[]MicrosoftGraphEvent**](microsoft.graph.event.md) |  | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](microsoft.graph.extension.md) |  | [optional] 

## Methods

### GetOriginalStartTimeZone

`func (o *Event) GetOriginalStartTimeZone() string`

GetOriginalStartTimeZone returns the OriginalStartTimeZone field if non-nil, zero value otherwise.

### GetOriginalStartTimeZoneOk

`func (o *Event) GetOriginalStartTimeZoneOk() (string, bool)`

GetOriginalStartTimeZoneOk returns a tuple with the OriginalStartTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOriginalStartTimeZone

`func (o *Event) HasOriginalStartTimeZone() bool`

HasOriginalStartTimeZone returns a boolean if a field has been set.

### SetOriginalStartTimeZone

`func (o *Event) SetOriginalStartTimeZone(v string)`

SetOriginalStartTimeZone gets a reference to the given string and assigns it to the OriginalStartTimeZone field.

### SetOriginalStartTimeZoneExplicitNull

`func (o *Event) SetOriginalStartTimeZoneExplicitNull(b bool)`

SetOriginalStartTimeZoneExplicitNull (un)sets OriginalStartTimeZone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OriginalStartTimeZone value is set to nil even if false is passed
### GetOriginalEndTimeZone

`func (o *Event) GetOriginalEndTimeZone() string`

GetOriginalEndTimeZone returns the OriginalEndTimeZone field if non-nil, zero value otherwise.

### GetOriginalEndTimeZoneOk

`func (o *Event) GetOriginalEndTimeZoneOk() (string, bool)`

GetOriginalEndTimeZoneOk returns a tuple with the OriginalEndTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOriginalEndTimeZone

`func (o *Event) HasOriginalEndTimeZone() bool`

HasOriginalEndTimeZone returns a boolean if a field has been set.

### SetOriginalEndTimeZone

`func (o *Event) SetOriginalEndTimeZone(v string)`

SetOriginalEndTimeZone gets a reference to the given string and assigns it to the OriginalEndTimeZone field.

### SetOriginalEndTimeZoneExplicitNull

`func (o *Event) SetOriginalEndTimeZoneExplicitNull(b bool)`

SetOriginalEndTimeZoneExplicitNull (un)sets OriginalEndTimeZone to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OriginalEndTimeZone value is set to nil even if false is passed
### GetResponseStatus

`func (o *Event) GetResponseStatus() AnyOfmicrosoftGraphResponseStatus`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *Event) GetResponseStatusOk() (AnyOfmicrosoftGraphResponseStatus, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseStatus

`func (o *Event) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### SetResponseStatus

`func (o *Event) SetResponseStatus(v AnyOfmicrosoftGraphResponseStatus)`

SetResponseStatus gets a reference to the given AnyOfmicrosoftGraphResponseStatus and assigns it to the ResponseStatus field.

### SetResponseStatusExplicitNull

`func (o *Event) SetResponseStatusExplicitNull(b bool)`

SetResponseStatusExplicitNull (un)sets ResponseStatus to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResponseStatus value is set to nil even if false is passed
### GetICalUId

`func (o *Event) GetICalUId() string`

GetICalUId returns the ICalUId field if non-nil, zero value otherwise.

### GetICalUIdOk

`func (o *Event) GetICalUIdOk() (string, bool)`

GetICalUIdOk returns a tuple with the ICalUId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasICalUId

`func (o *Event) HasICalUId() bool`

HasICalUId returns a boolean if a field has been set.

### SetICalUId

`func (o *Event) SetICalUId(v string)`

SetICalUId gets a reference to the given string and assigns it to the ICalUId field.

### SetICalUIdExplicitNull

`func (o *Event) SetICalUIdExplicitNull(b bool)`

SetICalUIdExplicitNull (un)sets ICalUId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ICalUId value is set to nil even if false is passed
### GetReminderMinutesBeforeStart

`func (o *Event) GetReminderMinutesBeforeStart() int32`

GetReminderMinutesBeforeStart returns the ReminderMinutesBeforeStart field if non-nil, zero value otherwise.

### GetReminderMinutesBeforeStartOk

`func (o *Event) GetReminderMinutesBeforeStartOk() (int32, bool)`

GetReminderMinutesBeforeStartOk returns a tuple with the ReminderMinutesBeforeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReminderMinutesBeforeStart

`func (o *Event) HasReminderMinutesBeforeStart() bool`

HasReminderMinutesBeforeStart returns a boolean if a field has been set.

### SetReminderMinutesBeforeStart

`func (o *Event) SetReminderMinutesBeforeStart(v int32)`

SetReminderMinutesBeforeStart gets a reference to the given int32 and assigns it to the ReminderMinutesBeforeStart field.

### SetReminderMinutesBeforeStartExplicitNull

`func (o *Event) SetReminderMinutesBeforeStartExplicitNull(b bool)`

SetReminderMinutesBeforeStartExplicitNull (un)sets ReminderMinutesBeforeStart to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReminderMinutesBeforeStart value is set to nil even if false is passed
### GetIsReminderOn

`func (o *Event) GetIsReminderOn() bool`

GetIsReminderOn returns the IsReminderOn field if non-nil, zero value otherwise.

### GetIsReminderOnOk

`func (o *Event) GetIsReminderOnOk() (bool, bool)`

GetIsReminderOnOk returns a tuple with the IsReminderOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReminderOn

`func (o *Event) HasIsReminderOn() bool`

HasIsReminderOn returns a boolean if a field has been set.

### SetIsReminderOn

`func (o *Event) SetIsReminderOn(v bool)`

SetIsReminderOn gets a reference to the given bool and assigns it to the IsReminderOn field.

### SetIsReminderOnExplicitNull

`func (o *Event) SetIsReminderOnExplicitNull(b bool)`

SetIsReminderOnExplicitNull (un)sets IsReminderOn to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsReminderOn value is set to nil even if false is passed
### GetHasAttachments

`func (o *Event) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Event) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *Event) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *Event) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### SetHasAttachmentsExplicitNull

`func (o *Event) SetHasAttachmentsExplicitNull(b bool)`

SetHasAttachmentsExplicitNull (un)sets HasAttachments to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasAttachments value is set to nil even if false is passed
### GetSubject

`func (o *Event) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Event) GetSubjectOk() (string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubject

`func (o *Event) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubject

`func (o *Event) SetSubject(v string)`

SetSubject gets a reference to the given string and assigns it to the Subject field.

### SetSubjectExplicitNull

`func (o *Event) SetSubjectExplicitNull(b bool)`

SetSubjectExplicitNull (un)sets Subject to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Subject value is set to nil even if false is passed
### GetBody

`func (o *Event) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Event) GetBodyOk() (AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *Event) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *Event) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the Body field.

### SetBodyExplicitNull

`func (o *Event) SetBodyExplicitNull(b bool)`

SetBodyExplicitNull (un)sets Body to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Body value is set to nil even if false is passed
### GetBodyPreview

`func (o *Event) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *Event) GetBodyPreviewOk() (string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBodyPreview

`func (o *Event) HasBodyPreview() bool`

HasBodyPreview returns a boolean if a field has been set.

### SetBodyPreview

`func (o *Event) SetBodyPreview(v string)`

SetBodyPreview gets a reference to the given string and assigns it to the BodyPreview field.

### SetBodyPreviewExplicitNull

`func (o *Event) SetBodyPreviewExplicitNull(b bool)`

SetBodyPreviewExplicitNull (un)sets BodyPreview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BodyPreview value is set to nil even if false is passed
### GetImportance

`func (o *Event) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *Event) GetImportanceOk() (AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportance

`func (o *Event) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportance

`func (o *Event) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance gets a reference to the given AnyOfmicrosoftGraphImportance and assigns it to the Importance field.

### SetImportanceExplicitNull

`func (o *Event) SetImportanceExplicitNull(b bool)`

SetImportanceExplicitNull (un)sets Importance to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Importance value is set to nil even if false is passed
### GetSensitivity

`func (o *Event) GetSensitivity() AnyOfmicrosoftGraphSensitivity`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *Event) GetSensitivityOk() (AnyOfmicrosoftGraphSensitivity, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSensitivity

`func (o *Event) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.

### SetSensitivity

`func (o *Event) SetSensitivity(v AnyOfmicrosoftGraphSensitivity)`

SetSensitivity gets a reference to the given AnyOfmicrosoftGraphSensitivity and assigns it to the Sensitivity field.

### SetSensitivityExplicitNull

`func (o *Event) SetSensitivityExplicitNull(b bool)`

SetSensitivityExplicitNull (un)sets Sensitivity to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sensitivity value is set to nil even if false is passed
### GetStart

`func (o *Event) GetStart() AnyOfmicrosoftGraphDateTimeTimeZone`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Event) GetStartOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStart

`func (o *Event) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStart

`func (o *Event) SetStart(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetStart gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the Start field.

### SetStartExplicitNull

`func (o *Event) SetStartExplicitNull(b bool)`

SetStartExplicitNull (un)sets Start to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Start value is set to nil even if false is passed
### GetOriginalStart

`func (o *Event) GetOriginalStart() time.Time`

GetOriginalStart returns the OriginalStart field if non-nil, zero value otherwise.

### GetOriginalStartOk

`func (o *Event) GetOriginalStartOk() (time.Time, bool)`

GetOriginalStartOk returns a tuple with the OriginalStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOriginalStart

`func (o *Event) HasOriginalStart() bool`

HasOriginalStart returns a boolean if a field has been set.

### SetOriginalStart

`func (o *Event) SetOriginalStart(v time.Time)`

SetOriginalStart gets a reference to the given time.Time and assigns it to the OriginalStart field.

### SetOriginalStartExplicitNull

`func (o *Event) SetOriginalStartExplicitNull(b bool)`

SetOriginalStartExplicitNull (un)sets OriginalStart to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OriginalStart value is set to nil even if false is passed
### GetEnd

`func (o *Event) GetEnd() AnyOfmicrosoftGraphDateTimeTimeZone`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Event) GetEndOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnd

`func (o *Event) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEnd

`func (o *Event) SetEnd(v AnyOfmicrosoftGraphDateTimeTimeZone)`

SetEnd gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the End field.

### SetEndExplicitNull

`func (o *Event) SetEndExplicitNull(b bool)`

SetEndExplicitNull (un)sets End to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The End value is set to nil even if false is passed
### GetLocation

`func (o *Event) GetLocation() AnyOfmicrosoftGraphLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Event) GetLocationOk() (AnyOfmicrosoftGraphLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *Event) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *Event) SetLocation(v AnyOfmicrosoftGraphLocation)`

SetLocation gets a reference to the given AnyOfmicrosoftGraphLocation and assigns it to the Location field.

### SetLocationExplicitNull

`func (o *Event) SetLocationExplicitNull(b bool)`

SetLocationExplicitNull (un)sets Location to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Location value is set to nil even if false is passed
### GetLocations

`func (o *Event) GetLocations() []AnyOfmicrosoftGraphLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Event) GetLocationsOk() ([]AnyOfmicrosoftGraphLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocations

`func (o *Event) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocations

`func (o *Event) SetLocations(v []AnyOfmicrosoftGraphLocation)`

SetLocations gets a reference to the given []AnyOfmicrosoftGraphLocation and assigns it to the Locations field.

### GetIsAllDay

`func (o *Event) GetIsAllDay() bool`

GetIsAllDay returns the IsAllDay field if non-nil, zero value otherwise.

### GetIsAllDayOk

`func (o *Event) GetIsAllDayOk() (bool, bool)`

GetIsAllDayOk returns a tuple with the IsAllDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAllDay

`func (o *Event) HasIsAllDay() bool`

HasIsAllDay returns a boolean if a field has been set.

### SetIsAllDay

`func (o *Event) SetIsAllDay(v bool)`

SetIsAllDay gets a reference to the given bool and assigns it to the IsAllDay field.

### SetIsAllDayExplicitNull

`func (o *Event) SetIsAllDayExplicitNull(b bool)`

SetIsAllDayExplicitNull (un)sets IsAllDay to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsAllDay value is set to nil even if false is passed
### GetIsCancelled

`func (o *Event) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *Event) GetIsCancelledOk() (bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsCancelled

`func (o *Event) HasIsCancelled() bool`

HasIsCancelled returns a boolean if a field has been set.

### SetIsCancelled

`func (o *Event) SetIsCancelled(v bool)`

SetIsCancelled gets a reference to the given bool and assigns it to the IsCancelled field.

### SetIsCancelledExplicitNull

`func (o *Event) SetIsCancelledExplicitNull(b bool)`

SetIsCancelledExplicitNull (un)sets IsCancelled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsCancelled value is set to nil even if false is passed
### GetIsOrganizer

`func (o *Event) GetIsOrganizer() bool`

GetIsOrganizer returns the IsOrganizer field if non-nil, zero value otherwise.

### GetIsOrganizerOk

`func (o *Event) GetIsOrganizerOk() (bool, bool)`

GetIsOrganizerOk returns a tuple with the IsOrganizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsOrganizer

`func (o *Event) HasIsOrganizer() bool`

HasIsOrganizer returns a boolean if a field has been set.

### SetIsOrganizer

`func (o *Event) SetIsOrganizer(v bool)`

SetIsOrganizer gets a reference to the given bool and assigns it to the IsOrganizer field.

### SetIsOrganizerExplicitNull

`func (o *Event) SetIsOrganizerExplicitNull(b bool)`

SetIsOrganizerExplicitNull (un)sets IsOrganizer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsOrganizer value is set to nil even if false is passed
### GetRecurrence

`func (o *Event) GetRecurrence() AnyOfmicrosoftGraphPatternedRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *Event) GetRecurrenceOk() (AnyOfmicrosoftGraphPatternedRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecurrence

`func (o *Event) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrence

`func (o *Event) SetRecurrence(v AnyOfmicrosoftGraphPatternedRecurrence)`

SetRecurrence gets a reference to the given AnyOfmicrosoftGraphPatternedRecurrence and assigns it to the Recurrence field.

### SetRecurrenceExplicitNull

`func (o *Event) SetRecurrenceExplicitNull(b bool)`

SetRecurrenceExplicitNull (un)sets Recurrence to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Recurrence value is set to nil even if false is passed
### GetResponseRequested

`func (o *Event) GetResponseRequested() bool`

GetResponseRequested returns the ResponseRequested field if non-nil, zero value otherwise.

### GetResponseRequestedOk

`func (o *Event) GetResponseRequestedOk() (bool, bool)`

GetResponseRequestedOk returns a tuple with the ResponseRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseRequested

`func (o *Event) HasResponseRequested() bool`

HasResponseRequested returns a boolean if a field has been set.

### SetResponseRequested

`func (o *Event) SetResponseRequested(v bool)`

SetResponseRequested gets a reference to the given bool and assigns it to the ResponseRequested field.

### SetResponseRequestedExplicitNull

`func (o *Event) SetResponseRequestedExplicitNull(b bool)`

SetResponseRequestedExplicitNull (un)sets ResponseRequested to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ResponseRequested value is set to nil even if false is passed
### GetSeriesMasterId

`func (o *Event) GetSeriesMasterId() string`

GetSeriesMasterId returns the SeriesMasterId field if non-nil, zero value otherwise.

### GetSeriesMasterIdOk

`func (o *Event) GetSeriesMasterIdOk() (string, bool)`

GetSeriesMasterIdOk returns a tuple with the SeriesMasterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSeriesMasterId

`func (o *Event) HasSeriesMasterId() bool`

HasSeriesMasterId returns a boolean if a field has been set.

### SetSeriesMasterId

`func (o *Event) SetSeriesMasterId(v string)`

SetSeriesMasterId gets a reference to the given string and assigns it to the SeriesMasterId field.

### SetSeriesMasterIdExplicitNull

`func (o *Event) SetSeriesMasterIdExplicitNull(b bool)`

SetSeriesMasterIdExplicitNull (un)sets SeriesMasterId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SeriesMasterId value is set to nil even if false is passed
### GetShowAs

`func (o *Event) GetShowAs() AnyOfmicrosoftGraphFreeBusyStatus`

GetShowAs returns the ShowAs field if non-nil, zero value otherwise.

### GetShowAsOk

`func (o *Event) GetShowAsOk() (AnyOfmicrosoftGraphFreeBusyStatus, bool)`

GetShowAsOk returns a tuple with the ShowAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowAs

`func (o *Event) HasShowAs() bool`

HasShowAs returns a boolean if a field has been set.

### SetShowAs

`func (o *Event) SetShowAs(v AnyOfmicrosoftGraphFreeBusyStatus)`

SetShowAs gets a reference to the given AnyOfmicrosoftGraphFreeBusyStatus and assigns it to the ShowAs field.

### SetShowAsExplicitNull

`func (o *Event) SetShowAsExplicitNull(b bool)`

SetShowAsExplicitNull (un)sets ShowAs to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ShowAs value is set to nil even if false is passed
### GetType

`func (o *Event) GetType() AnyOfmicrosoftGraphEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (AnyOfmicrosoftGraphEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Event) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Event) SetType(v AnyOfmicrosoftGraphEventType)`

SetType gets a reference to the given AnyOfmicrosoftGraphEventType and assigns it to the Type field.

### SetTypeExplicitNull

`func (o *Event) SetTypeExplicitNull(b bool)`

SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Type value is set to nil even if false is passed
### GetAttendees

`func (o *Event) GetAttendees() []AnyOfmicrosoftGraphAttendee`

GetAttendees returns the Attendees field if non-nil, zero value otherwise.

### GetAttendeesOk

`func (o *Event) GetAttendeesOk() ([]AnyOfmicrosoftGraphAttendee, bool)`

GetAttendeesOk returns a tuple with the Attendees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttendees

`func (o *Event) HasAttendees() bool`

HasAttendees returns a boolean if a field has been set.

### SetAttendees

`func (o *Event) SetAttendees(v []AnyOfmicrosoftGraphAttendee)`

SetAttendees gets a reference to the given []AnyOfmicrosoftGraphAttendee and assigns it to the Attendees field.

### GetOrganizer

`func (o *Event) GetOrganizer() AnyOfmicrosoftGraphRecipient`

GetOrganizer returns the Organizer field if non-nil, zero value otherwise.

### GetOrganizerOk

`func (o *Event) GetOrganizerOk() (AnyOfmicrosoftGraphRecipient, bool)`

GetOrganizerOk returns a tuple with the Organizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizer

`func (o *Event) HasOrganizer() bool`

HasOrganizer returns a boolean if a field has been set.

### SetOrganizer

`func (o *Event) SetOrganizer(v AnyOfmicrosoftGraphRecipient)`

SetOrganizer gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the Organizer field.

### SetOrganizerExplicitNull

`func (o *Event) SetOrganizerExplicitNull(b bool)`

SetOrganizerExplicitNull (un)sets Organizer to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Organizer value is set to nil even if false is passed
### GetWebLink

`func (o *Event) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *Event) GetWebLinkOk() (string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebLink

`func (o *Event) HasWebLink() bool`

HasWebLink returns a boolean if a field has been set.

### SetWebLink

`func (o *Event) SetWebLink(v string)`

SetWebLink gets a reference to the given string and assigns it to the WebLink field.

### SetWebLinkExplicitNull

`func (o *Event) SetWebLinkExplicitNull(b bool)`

SetWebLinkExplicitNull (un)sets WebLink to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebLink value is set to nil even if false is passed
### GetOnlineMeetingUrl

`func (o *Event) GetOnlineMeetingUrl() string`

GetOnlineMeetingUrl returns the OnlineMeetingUrl field if non-nil, zero value otherwise.

### GetOnlineMeetingUrlOk

`func (o *Event) GetOnlineMeetingUrlOk() (string, bool)`

GetOnlineMeetingUrlOk returns a tuple with the OnlineMeetingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOnlineMeetingUrl

`func (o *Event) HasOnlineMeetingUrl() bool`

HasOnlineMeetingUrl returns a boolean if a field has been set.

### SetOnlineMeetingUrl

`func (o *Event) SetOnlineMeetingUrl(v string)`

SetOnlineMeetingUrl gets a reference to the given string and assigns it to the OnlineMeetingUrl field.

### SetOnlineMeetingUrlExplicitNull

`func (o *Event) SetOnlineMeetingUrlExplicitNull(b bool)`

SetOnlineMeetingUrlExplicitNull (un)sets OnlineMeetingUrl to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OnlineMeetingUrl value is set to nil even if false is passed
### GetAttachments

`func (o *Event) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Event) GetAttachmentsOk() ([]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttachments

`func (o *Event) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachments

`func (o *Event) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments gets a reference to the given []MicrosoftGraphAttachment and assigns it to the Attachments field.

### GetSingleValueExtendedProperties

`func (o *Event) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Event) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *Event) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *Event) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *Event) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Event) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *Event) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *Event) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetCalendar

`func (o *Event) GetCalendar() AnyOfmicrosoftGraphCalendar`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *Event) GetCalendarOk() (AnyOfmicrosoftGraphCalendar, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCalendar

`func (o *Event) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### SetCalendar

`func (o *Event) SetCalendar(v AnyOfmicrosoftGraphCalendar)`

SetCalendar gets a reference to the given AnyOfmicrosoftGraphCalendar and assigns it to the Calendar field.

### SetCalendarExplicitNull

`func (o *Event) SetCalendarExplicitNull(b bool)`

SetCalendarExplicitNull (un)sets Calendar to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Calendar value is set to nil even if false is passed
### GetInstances

`func (o *Event) GetInstances() []MicrosoftGraphEvent`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *Event) GetInstancesOk() ([]MicrosoftGraphEvent, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInstances

`func (o *Event) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstances

`func (o *Event) SetInstances(v []MicrosoftGraphEvent)`

SetInstances gets a reference to the given []MicrosoftGraphEvent and assigns it to the Instances field.

### GetExtensions

`func (o *Event) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Event) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *Event) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *Event) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



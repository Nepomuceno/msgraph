# MicrosoftGraphMeetingTimeSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**OrganizerAvailability** | Pointer to [**AnyOfmicrosoftGraphFreeBusyStatus**](anyOf&lt;microsoft.graph.freeBusyStatus&gt;.md) |  | [optional] 
**AttendeeAvailability** | Pointer to [**[]AnyOfmicrosoftGraphAttendeeAvailability**](anyOf&lt;microsoft.graph.attendeeAvailability&gt;.md) |  | [optional] 
**Locations** | Pointer to [**[]AnyOfmicrosoftGraphLocation**](anyOf&lt;microsoft.graph.location&gt;.md) |  | [optional] 
**SuggestionReason** | Pointer to **string** |  | [optional] 
**MeetingTimeSlot** | Pointer to [**AnyOfmicrosoftGraphTimeSlot**](anyOf&lt;microsoft.graph.timeSlot&gt;.md) |  | [optional] 

## Methods

### GetConfidence

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetConfidence() AnyOfnumberstringstring`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetConfidenceOk() (AnyOfnumberstringstring, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConfidence

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidence

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetConfidence(v AnyOfnumberstringstring)`

SetConfidence gets a reference to the given AnyOfnumberstringstring and assigns it to the Confidence field.

### SetConfidenceExplicitNull

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetConfidenceExplicitNull(b bool)`

SetConfidenceExplicitNull (un)sets Confidence to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Confidence value is set to nil even if false is passed
### GetOrder

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetOrderOk() (int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrder

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrder

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetOrder(v int32)`

SetOrder gets a reference to the given int32 and assigns it to the Order field.

### SetOrderExplicitNull

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetOrderExplicitNull(b bool)`

SetOrderExplicitNull (un)sets Order to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Order value is set to nil even if false is passed
### GetOrganizerAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetOrganizerAvailability() AnyOfmicrosoftGraphFreeBusyStatus`

GetOrganizerAvailability returns the OrganizerAvailability field if non-nil, zero value otherwise.

### GetOrganizerAvailabilityOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetOrganizerAvailabilityOk() (AnyOfmicrosoftGraphFreeBusyStatus, bool)`

GetOrganizerAvailabilityOk returns a tuple with the OrganizerAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOrganizerAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasOrganizerAvailability() bool`

HasOrganizerAvailability returns a boolean if a field has been set.

### SetOrganizerAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetOrganizerAvailability(v AnyOfmicrosoftGraphFreeBusyStatus)`

SetOrganizerAvailability gets a reference to the given AnyOfmicrosoftGraphFreeBusyStatus and assigns it to the OrganizerAvailability field.

### SetOrganizerAvailabilityExplicitNull

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetOrganizerAvailabilityExplicitNull(b bool)`

SetOrganizerAvailabilityExplicitNull (un)sets OrganizerAvailability to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The OrganizerAvailability value is set to nil even if false is passed
### GetAttendeeAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetAttendeeAvailability() []AnyOfmicrosoftGraphAttendeeAvailability`

GetAttendeeAvailability returns the AttendeeAvailability field if non-nil, zero value otherwise.

### GetAttendeeAvailabilityOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetAttendeeAvailabilityOk() ([]AnyOfmicrosoftGraphAttendeeAvailability, bool)`

GetAttendeeAvailabilityOk returns a tuple with the AttendeeAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttendeeAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasAttendeeAvailability() bool`

HasAttendeeAvailability returns a boolean if a field has been set.

### SetAttendeeAvailability

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetAttendeeAvailability(v []AnyOfmicrosoftGraphAttendeeAvailability)`

SetAttendeeAvailability gets a reference to the given []AnyOfmicrosoftGraphAttendeeAvailability and assigns it to the AttendeeAvailability field.

### GetLocations

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetLocations() []AnyOfmicrosoftGraphLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetLocationsOk() ([]AnyOfmicrosoftGraphLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocations

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocations

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetLocations(v []AnyOfmicrosoftGraphLocation)`

SetLocations gets a reference to the given []AnyOfmicrosoftGraphLocation and assigns it to the Locations field.

### GetSuggestionReason

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetSuggestionReason() string`

GetSuggestionReason returns the SuggestionReason field if non-nil, zero value otherwise.

### GetSuggestionReasonOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetSuggestionReasonOk() (string, bool)`

GetSuggestionReasonOk returns a tuple with the SuggestionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSuggestionReason

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasSuggestionReason() bool`

HasSuggestionReason returns a boolean if a field has been set.

### SetSuggestionReason

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetSuggestionReason(v string)`

SetSuggestionReason gets a reference to the given string and assigns it to the SuggestionReason field.

### SetSuggestionReasonExplicitNull

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetSuggestionReasonExplicitNull(b bool)`

SetSuggestionReasonExplicitNull (un)sets SuggestionReason to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SuggestionReason value is set to nil even if false is passed
### GetMeetingTimeSlot

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetMeetingTimeSlot() AnyOfmicrosoftGraphTimeSlot`

GetMeetingTimeSlot returns the MeetingTimeSlot field if non-nil, zero value otherwise.

### GetMeetingTimeSlotOk

`func (o *MicrosoftGraphMeetingTimeSuggestion) GetMeetingTimeSlotOk() (AnyOfmicrosoftGraphTimeSlot, bool)`

GetMeetingTimeSlotOk returns a tuple with the MeetingTimeSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMeetingTimeSlot

`func (o *MicrosoftGraphMeetingTimeSuggestion) HasMeetingTimeSlot() bool`

HasMeetingTimeSlot returns a boolean if a field has been set.

### SetMeetingTimeSlot

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetMeetingTimeSlot(v AnyOfmicrosoftGraphTimeSlot)`

SetMeetingTimeSlot gets a reference to the given AnyOfmicrosoftGraphTimeSlot and assigns it to the MeetingTimeSlot field.

### SetMeetingTimeSlotExplicitNull

`func (o *MicrosoftGraphMeetingTimeSuggestion) SetMeetingTimeSlotExplicitNull(b bool)`

SetMeetingTimeSlotExplicitNull (un)sets MeetingTimeSlot to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MeetingTimeSlot value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MicrosoftGraphAttendeeAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attendee** | Pointer to [**AnyOfmicrosoftGraphAttendeeBase**](anyOf&lt;microsoft.graph.attendeeBase&gt;.md) |  | [optional] 
**Availability** | Pointer to [**AnyOfmicrosoftGraphFreeBusyStatus**](anyOf&lt;microsoft.graph.freeBusyStatus&gt;.md) |  | [optional] 

## Methods

### GetAttendee

`func (o *MicrosoftGraphAttendeeAvailability) GetAttendee() AnyOfmicrosoftGraphAttendeeBase`

GetAttendee returns the Attendee field if non-nil, zero value otherwise.

### GetAttendeeOk

`func (o *MicrosoftGraphAttendeeAvailability) GetAttendeeOk() (AnyOfmicrosoftGraphAttendeeBase, bool)`

GetAttendeeOk returns a tuple with the Attendee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttendee

`func (o *MicrosoftGraphAttendeeAvailability) HasAttendee() bool`

HasAttendee returns a boolean if a field has been set.

### SetAttendee

`func (o *MicrosoftGraphAttendeeAvailability) SetAttendee(v AnyOfmicrosoftGraphAttendeeBase)`

SetAttendee gets a reference to the given AnyOfmicrosoftGraphAttendeeBase and assigns it to the Attendee field.

### SetAttendeeExplicitNull

`func (o *MicrosoftGraphAttendeeAvailability) SetAttendeeExplicitNull(b bool)`

SetAttendeeExplicitNull (un)sets Attendee to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Attendee value is set to nil even if false is passed
### GetAvailability

`func (o *MicrosoftGraphAttendeeAvailability) GetAvailability() AnyOfmicrosoftGraphFreeBusyStatus`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *MicrosoftGraphAttendeeAvailability) GetAvailabilityOk() (AnyOfmicrosoftGraphFreeBusyStatus, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvailability

`func (o *MicrosoftGraphAttendeeAvailability) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailability

`func (o *MicrosoftGraphAttendeeAvailability) SetAvailability(v AnyOfmicrosoftGraphFreeBusyStatus)`

SetAvailability gets a reference to the given AnyOfmicrosoftGraphFreeBusyStatus and assigns it to the Availability field.

### SetAvailabilityExplicitNull

`func (o *MicrosoftGraphAttendeeAvailability) SetAvailabilityExplicitNull(b bool)`

SetAvailabilityExplicitNull (un)sets Availability to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Availability value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



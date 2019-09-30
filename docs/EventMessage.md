# EventMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeetingMessageType** | Pointer to [**AnyOfmicrosoftGraphMeetingMessageType**](anyOf&lt;microsoft.graph.meetingMessageType&gt;.md) |  | [optional] 
**Event** | Pointer to [**AnyOfmicrosoftGraphEvent**](anyOf&lt;microsoft.graph.event&gt;.md) |  | [optional] 

## Methods

### GetMeetingMessageType

`func (o *EventMessage) GetMeetingMessageType() AnyOfmicrosoftGraphMeetingMessageType`

GetMeetingMessageType returns the MeetingMessageType field if non-nil, zero value otherwise.

### GetMeetingMessageTypeOk

`func (o *EventMessage) GetMeetingMessageTypeOk() (AnyOfmicrosoftGraphMeetingMessageType, bool)`

GetMeetingMessageTypeOk returns a tuple with the MeetingMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMeetingMessageType

`func (o *EventMessage) HasMeetingMessageType() bool`

HasMeetingMessageType returns a boolean if a field has been set.

### SetMeetingMessageType

`func (o *EventMessage) SetMeetingMessageType(v AnyOfmicrosoftGraphMeetingMessageType)`

SetMeetingMessageType gets a reference to the given AnyOfmicrosoftGraphMeetingMessageType and assigns it to the MeetingMessageType field.

### SetMeetingMessageTypeExplicitNull

`func (o *EventMessage) SetMeetingMessageTypeExplicitNull(b bool)`

SetMeetingMessageTypeExplicitNull (un)sets MeetingMessageType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MeetingMessageType value is set to nil even if false is passed
### GetEvent

`func (o *EventMessage) GetEvent() AnyOfmicrosoftGraphEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventMessage) GetEventOk() (AnyOfmicrosoftGraphEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvent

`func (o *EventMessage) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEvent

`func (o *EventMessage) SetEvent(v AnyOfmicrosoftGraphEvent)`

SetEvent gets a reference to the given AnyOfmicrosoftGraphEvent and assigns it to the Event field.

### SetEventExplicitNull

`func (o *EventMessage) SetEventExplicitNull(b bool)`

SetEventExplicitNull (un)sets Event to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Event value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



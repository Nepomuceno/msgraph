/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
import (
	"encoding/json"
)
// EventMessage struct for EventMessage
type EventMessage struct {
	MeetingMessageType *AnyOfmicrosoftGraphMeetingMessageType `json:"meetingMessageType,omitempty"`
	isExplicitNullMeetingMessageType bool `json:"-"`
	Event *AnyOfmicrosoftGraphEvent `json:"event,omitempty"`
	isExplicitNullEvent bool `json:"-"`
}

// GetMeetingMessageType returns the MeetingMessageType field if non-nil, zero value otherwise.
func (o *EventMessage) GetMeetingMessageType() AnyOfmicrosoftGraphMeetingMessageType {
	if o == nil || o.MeetingMessageType == nil {
		var ret AnyOfmicrosoftGraphMeetingMessageType
		return ret
	}
	return *o.MeetingMessageType
}

// GetMeetingMessageTypeOk returns a tuple with the MeetingMessageType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EventMessage) GetMeetingMessageTypeOk() (AnyOfmicrosoftGraphMeetingMessageType, bool) {
	if o == nil || o.MeetingMessageType == nil {
		var ret AnyOfmicrosoftGraphMeetingMessageType
		return ret, false
	}
	return *o.MeetingMessageType, true
}

// HasMeetingMessageType returns a boolean if a field has been set.
func (o *EventMessage) HasMeetingMessageType() bool {
	if o != nil && o.MeetingMessageType != nil {
		return true
	}

	return false
}

// SetMeetingMessageType gets a reference to the given AnyOfmicrosoftGraphMeetingMessageType and assigns it to the MeetingMessageType field.
func (o *EventMessage) SetMeetingMessageType(v AnyOfmicrosoftGraphMeetingMessageType) {
	o.MeetingMessageType = &v
}

// SetMeetingMessageTypeExplicitNull (un)sets MeetingMessageType to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The MeetingMessageType value is set to nil even if false is passed
func (o *EventMessage) SetMeetingMessageTypeExplicitNull(b bool) {
	o.MeetingMessageType = nil
	o.isExplicitNullMeetingMessageType = b
}
// GetEvent returns the Event field if non-nil, zero value otherwise.
func (o *EventMessage) GetEvent() AnyOfmicrosoftGraphEvent {
	if o == nil || o.Event == nil {
		var ret AnyOfmicrosoftGraphEvent
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EventMessage) GetEventOk() (AnyOfmicrosoftGraphEvent, bool) {
	if o == nil || o.Event == nil {
		var ret AnyOfmicrosoftGraphEvent
		return ret, false
	}
	return *o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *EventMessage) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given AnyOfmicrosoftGraphEvent and assigns it to the Event field.
func (o *EventMessage) SetEvent(v AnyOfmicrosoftGraphEvent) {
	o.Event = &v
}

// SetEventExplicitNull (un)sets Event to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Event value is set to nil even if false is passed
func (o *EventMessage) SetEventExplicitNull(b bool) {
	o.Event = nil
	o.isExplicitNullEvent = b
}

// MarshalJSON returns the JSON representation of the model.
func (o EventMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MeetingMessageType == nil {
		if o.isExplicitNullMeetingMessageType {
			toSerialize["meetingMessageType"] = o.MeetingMessageType
		}
	} else {
		toSerialize["meetingMessageType"] = o.MeetingMessageType
	}
	if o.Event == nil {
		if o.isExplicitNullEvent {
			toSerialize["event"] = o.Event
		}
	} else {
		toSerialize["event"] = o.Event
	}
	return json.Marshal(toSerialize)
}



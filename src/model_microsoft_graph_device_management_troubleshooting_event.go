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
	"time"
	"encoding/json"
)
// MicrosoftGraphDeviceManagementTroubleshootingEvent struct for MicrosoftGraphDeviceManagementTroubleshootingEvent
type MicrosoftGraphDeviceManagementTroubleshootingEvent struct {
	Id *string `json:"id,omitempty"`

	// Time when the event occurred .
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`

	// Id used for tracing the failure in the service.
	CorrelationId *string `json:"correlationId,omitempty"`
	isExplicitNullCorrelationId bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetId(v string) {
	o.Id = &v
}

// GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetEventDateTime() time.Time {
	if o == nil || o.EventDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetEventDateTimeOk() (time.Time, bool) {
	if o == nil || o.EventDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EventDateTime, true
}

// HasEventDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasEventDateTime() bool {
	if o != nil && o.EventDateTime != nil {
		return true
	}

	return false
}

// SetEventDateTime gets a reference to the given time.Time and assigns it to the EventDateTime field.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetEventDateTime(v time.Time) {
	o.EventDateTime = &v
}

// GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetCorrelationId() string {
	if o == nil || o.CorrelationId == nil {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) GetCorrelationIdOk() (string, bool) {
	if o == nil || o.CorrelationId == nil {
		var ret string
		return ret, false
	}
	return *o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) HasCorrelationId() bool {
	if o != nil && o.CorrelationId != nil {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// SetCorrelationIdExplicitNull (un)sets CorrelationId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CorrelationId value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceManagementTroubleshootingEvent) SetCorrelationIdExplicitNull(b bool) {
	o.CorrelationId = nil
	o.isExplicitNullCorrelationId = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDeviceManagementTroubleshootingEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EventDateTime != nil {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	if o.CorrelationId == nil {
		if o.isExplicitNullCorrelationId {
			toSerialize["correlationId"] = o.CorrelationId
		}
	} else {
		toSerialize["correlationId"] = o.CorrelationId
	}
	return json.Marshal(toSerialize)
}



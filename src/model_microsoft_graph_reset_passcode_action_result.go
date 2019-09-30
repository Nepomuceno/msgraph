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
// MicrosoftGraphResetPasscodeActionResult struct for MicrosoftGraphResetPasscodeActionResult
type MicrosoftGraphResetPasscodeActionResult struct {
	// Action name
	ActionName *string `json:"actionName,omitempty"`
	isExplicitNullActionName bool `json:"-"`
	// State of the action
	ActionState *AnyOfmicrosoftGraphActionState `json:"actionState,omitempty"`

	// Time the action was initiated
	StartDateTime *time.Time `json:"startDateTime,omitempty"`

	// Time the action state was last updated
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`

	// Newly generated passcode for the device 
	Passcode *string `json:"passcode,omitempty"`
	isExplicitNullPasscode bool `json:"-"`
}

// GetActionName returns the ActionName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResetPasscodeActionResult) GetActionName() string {
	if o == nil || o.ActionName == nil {
		var ret string
		return ret
	}
	return *o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) GetActionNameOk() (string, bool) {
	if o == nil || o.ActionName == nil {
		var ret string
		return ret, false
	}
	return *o.ActionName, true
}

// HasActionName returns a boolean if a field has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) HasActionName() bool {
	if o != nil && o.ActionName != nil {
		return true
	}

	return false
}

// SetActionName gets a reference to the given string and assigns it to the ActionName field.
func (o *MicrosoftGraphResetPasscodeActionResult) SetActionName(v string) {
	o.ActionName = &v
}

// SetActionNameExplicitNull (un)sets ActionName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ActionName value is set to nil even if false is passed
func (o *MicrosoftGraphResetPasscodeActionResult) SetActionNameExplicitNull(b bool) {
	o.ActionName = nil
	o.isExplicitNullActionName = b
}
// GetActionState returns the ActionState field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResetPasscodeActionResult) GetActionState() AnyOfmicrosoftGraphActionState {
	if o == nil || o.ActionState == nil {
		var ret AnyOfmicrosoftGraphActionState
		return ret
	}
	return *o.ActionState
}

// GetActionStateOk returns a tuple with the ActionState field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) GetActionStateOk() (AnyOfmicrosoftGraphActionState, bool) {
	if o == nil || o.ActionState == nil {
		var ret AnyOfmicrosoftGraphActionState
		return ret, false
	}
	return *o.ActionState, true
}

// HasActionState returns a boolean if a field has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) HasActionState() bool {
	if o != nil && o.ActionState != nil {
		return true
	}

	return false
}

// SetActionState gets a reference to the given AnyOfmicrosoftGraphActionState and assigns it to the ActionState field.
func (o *MicrosoftGraphResetPasscodeActionResult) SetActionState(v AnyOfmicrosoftGraphActionState) {
	o.ActionState = &v
}

// GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResetPasscodeActionResult) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) GetStartDateTimeOk() (time.Time, bool) {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *MicrosoftGraphResetPasscodeActionResult) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetLastUpdatedDateTime returns the LastUpdatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResetPasscodeActionResult) GetLastUpdatedDateTime() time.Time {
	if o == nil || o.LastUpdatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDateTime
}

// GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) GetLastUpdatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastUpdatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastUpdatedDateTime, true
}

// HasLastUpdatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) HasLastUpdatedDateTime() bool {
	if o != nil && o.LastUpdatedDateTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedDateTime gets a reference to the given time.Time and assigns it to the LastUpdatedDateTime field.
func (o *MicrosoftGraphResetPasscodeActionResult) SetLastUpdatedDateTime(v time.Time) {
	o.LastUpdatedDateTime = &v
}

// GetPasscode returns the Passcode field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResetPasscodeActionResult) GetPasscode() string {
	if o == nil || o.Passcode == nil {
		var ret string
		return ret
	}
	return *o.Passcode
}

// GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) GetPasscodeOk() (string, bool) {
	if o == nil || o.Passcode == nil {
		var ret string
		return ret, false
	}
	return *o.Passcode, true
}

// HasPasscode returns a boolean if a field has been set.
func (o *MicrosoftGraphResetPasscodeActionResult) HasPasscode() bool {
	if o != nil && o.Passcode != nil {
		return true
	}

	return false
}

// SetPasscode gets a reference to the given string and assigns it to the Passcode field.
func (o *MicrosoftGraphResetPasscodeActionResult) SetPasscode(v string) {
	o.Passcode = &v
}

// SetPasscodeExplicitNull (un)sets Passcode to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Passcode value is set to nil even if false is passed
func (o *MicrosoftGraphResetPasscodeActionResult) SetPasscodeExplicitNull(b bool) {
	o.Passcode = nil
	o.isExplicitNullPasscode = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphResetPasscodeActionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionName == nil {
		if o.isExplicitNullActionName {
			toSerialize["actionName"] = o.ActionName
		}
	} else {
		toSerialize["actionName"] = o.ActionName
	}
	if o.ActionState != nil {
		toSerialize["actionState"] = o.ActionState
	}
	if o.StartDateTime != nil {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if o.LastUpdatedDateTime != nil {
		toSerialize["lastUpdatedDateTime"] = o.LastUpdatedDateTime
	}
	if o.Passcode == nil {
		if o.isExplicitNullPasscode {
			toSerialize["passcode"] = o.Passcode
		}
	} else {
		toSerialize["passcode"] = o.Passcode
	}
	return json.Marshal(toSerialize)
}



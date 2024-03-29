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
// MicrosoftGraphLocateDeviceActionResult struct for MicrosoftGraphLocateDeviceActionResult
type MicrosoftGraphLocateDeviceActionResult struct {
	// Action name
	ActionName *string `json:"actionName,omitempty"`
	isExplicitNullActionName bool `json:"-"`
	// State of the action
	ActionState *AnyOfmicrosoftGraphActionState `json:"actionState,omitempty"`

	// Time the action was initiated
	StartDateTime *time.Time `json:"startDateTime,omitempty"`

	// Time the action state was last updated
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`

	// device location
	DeviceLocation *AnyOfmicrosoftGraphDeviceGeoLocation `json:"deviceLocation,omitempty"`
	isExplicitNullDeviceLocation bool `json:"-"`
}

// GetActionName returns the ActionName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphLocateDeviceActionResult) GetActionName() string {
	if o == nil || o.ActionName == nil {
		var ret string
		return ret
	}
	return *o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) GetActionNameOk() (string, bool) {
	if o == nil || o.ActionName == nil {
		var ret string
		return ret, false
	}
	return *o.ActionName, true
}

// HasActionName returns a boolean if a field has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) HasActionName() bool {
	if o != nil && o.ActionName != nil {
		return true
	}

	return false
}

// SetActionName gets a reference to the given string and assigns it to the ActionName field.
func (o *MicrosoftGraphLocateDeviceActionResult) SetActionName(v string) {
	o.ActionName = &v
}

// SetActionNameExplicitNull (un)sets ActionName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ActionName value is set to nil even if false is passed
func (o *MicrosoftGraphLocateDeviceActionResult) SetActionNameExplicitNull(b bool) {
	o.ActionName = nil
	o.isExplicitNullActionName = b
}
// GetActionState returns the ActionState field if non-nil, zero value otherwise.
func (o *MicrosoftGraphLocateDeviceActionResult) GetActionState() AnyOfmicrosoftGraphActionState {
	if o == nil || o.ActionState == nil {
		var ret AnyOfmicrosoftGraphActionState
		return ret
	}
	return *o.ActionState
}

// GetActionStateOk returns a tuple with the ActionState field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) GetActionStateOk() (AnyOfmicrosoftGraphActionState, bool) {
	if o == nil || o.ActionState == nil {
		var ret AnyOfmicrosoftGraphActionState
		return ret, false
	}
	return *o.ActionState, true
}

// HasActionState returns a boolean if a field has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) HasActionState() bool {
	if o != nil && o.ActionState != nil {
		return true
	}

	return false
}

// SetActionState gets a reference to the given AnyOfmicrosoftGraphActionState and assigns it to the ActionState field.
func (o *MicrosoftGraphLocateDeviceActionResult) SetActionState(v AnyOfmicrosoftGraphActionState) {
	o.ActionState = &v
}

// GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphLocateDeviceActionResult) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) GetStartDateTimeOk() (time.Time, bool) {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *MicrosoftGraphLocateDeviceActionResult) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetLastUpdatedDateTime returns the LastUpdatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphLocateDeviceActionResult) GetLastUpdatedDateTime() time.Time {
	if o == nil || o.LastUpdatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDateTime
}

// GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) GetLastUpdatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastUpdatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastUpdatedDateTime, true
}

// HasLastUpdatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) HasLastUpdatedDateTime() bool {
	if o != nil && o.LastUpdatedDateTime != nil {
		return true
	}

	return false
}

// SetLastUpdatedDateTime gets a reference to the given time.Time and assigns it to the LastUpdatedDateTime field.
func (o *MicrosoftGraphLocateDeviceActionResult) SetLastUpdatedDateTime(v time.Time) {
	o.LastUpdatedDateTime = &v
}

// GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.
func (o *MicrosoftGraphLocateDeviceActionResult) GetDeviceLocation() AnyOfmicrosoftGraphDeviceGeoLocation {
	if o == nil || o.DeviceLocation == nil {
		var ret AnyOfmicrosoftGraphDeviceGeoLocation
		return ret
	}
	return *o.DeviceLocation
}

// GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) GetDeviceLocationOk() (AnyOfmicrosoftGraphDeviceGeoLocation, bool) {
	if o == nil || o.DeviceLocation == nil {
		var ret AnyOfmicrosoftGraphDeviceGeoLocation
		return ret, false
	}
	return *o.DeviceLocation, true
}

// HasDeviceLocation returns a boolean if a field has been set.
func (o *MicrosoftGraphLocateDeviceActionResult) HasDeviceLocation() bool {
	if o != nil && o.DeviceLocation != nil {
		return true
	}

	return false
}

// SetDeviceLocation gets a reference to the given AnyOfmicrosoftGraphDeviceGeoLocation and assigns it to the DeviceLocation field.
func (o *MicrosoftGraphLocateDeviceActionResult) SetDeviceLocation(v AnyOfmicrosoftGraphDeviceGeoLocation) {
	o.DeviceLocation = &v
}

// SetDeviceLocationExplicitNull (un)sets DeviceLocation to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DeviceLocation value is set to nil even if false is passed
func (o *MicrosoftGraphLocateDeviceActionResult) SetDeviceLocationExplicitNull(b bool) {
	o.DeviceLocation = nil
	o.isExplicitNullDeviceLocation = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphLocateDeviceActionResult) MarshalJSON() ([]byte, error) {
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
	if o.DeviceLocation == nil {
		if o.isExplicitNullDeviceLocation {
			toSerialize["deviceLocation"] = o.DeviceLocation
		}
	} else {
		toSerialize["deviceLocation"] = o.DeviceLocation
	}
	return json.Marshal(toSerialize)
}



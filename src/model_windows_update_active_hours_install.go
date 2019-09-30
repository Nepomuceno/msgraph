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
// WindowsUpdateActiveHoursInstall struct for WindowsUpdateActiveHoursInstall
type WindowsUpdateActiveHoursInstall struct {
	// Active Hours Start
	ActiveHoursStart *string `json:"activeHoursStart,omitempty"`

	// Active Hours End
	ActiveHoursEnd *string `json:"activeHoursEnd,omitempty"`

}

// GetActiveHoursStart returns the ActiveHoursStart field if non-nil, zero value otherwise.
func (o *WindowsUpdateActiveHoursInstall) GetActiveHoursStart() string {
	if o == nil || o.ActiveHoursStart == nil {
		var ret string
		return ret
	}
	return *o.ActiveHoursStart
}

// GetActiveHoursStartOk returns a tuple with the ActiveHoursStart field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsUpdateActiveHoursInstall) GetActiveHoursStartOk() (string, bool) {
	if o == nil || o.ActiveHoursStart == nil {
		var ret string
		return ret, false
	}
	return *o.ActiveHoursStart, true
}

// HasActiveHoursStart returns a boolean if a field has been set.
func (o *WindowsUpdateActiveHoursInstall) HasActiveHoursStart() bool {
	if o != nil && o.ActiveHoursStart != nil {
		return true
	}

	return false
}

// SetActiveHoursStart gets a reference to the given string and assigns it to the ActiveHoursStart field.
func (o *WindowsUpdateActiveHoursInstall) SetActiveHoursStart(v string) {
	o.ActiveHoursStart = &v
}

// GetActiveHoursEnd returns the ActiveHoursEnd field if non-nil, zero value otherwise.
func (o *WindowsUpdateActiveHoursInstall) GetActiveHoursEnd() string {
	if o == nil || o.ActiveHoursEnd == nil {
		var ret string
		return ret
	}
	return *o.ActiveHoursEnd
}

// GetActiveHoursEndOk returns a tuple with the ActiveHoursEnd field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsUpdateActiveHoursInstall) GetActiveHoursEndOk() (string, bool) {
	if o == nil || o.ActiveHoursEnd == nil {
		var ret string
		return ret, false
	}
	return *o.ActiveHoursEnd, true
}

// HasActiveHoursEnd returns a boolean if a field has been set.
func (o *WindowsUpdateActiveHoursInstall) HasActiveHoursEnd() bool {
	if o != nil && o.ActiveHoursEnd != nil {
		return true
	}

	return false
}

// SetActiveHoursEnd gets a reference to the given string and assigns it to the ActiveHoursEnd field.
func (o *WindowsUpdateActiveHoursInstall) SetActiveHoursEnd(v string) {
	o.ActiveHoursEnd = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o WindowsUpdateActiveHoursInstall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveHoursStart != nil {
		toSerialize["activeHoursStart"] = o.ActiveHoursStart
	}
	if o.ActiveHoursEnd != nil {
		toSerialize["activeHoursEnd"] = o.ActiveHoursEnd
	}
	return json.Marshal(toSerialize)
}



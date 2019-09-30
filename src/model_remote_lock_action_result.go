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
// RemoteLockActionResult struct for RemoteLockActionResult
type RemoteLockActionResult struct {
	// Pin to unlock the client
	UnlockPin *string `json:"unlockPin,omitempty"`
	isExplicitNullUnlockPin bool `json:"-"`
}

// GetUnlockPin returns the UnlockPin field if non-nil, zero value otherwise.
func (o *RemoteLockActionResult) GetUnlockPin() string {
	if o == nil || o.UnlockPin == nil {
		var ret string
		return ret
	}
	return *o.UnlockPin
}

// GetUnlockPinOk returns a tuple with the UnlockPin field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RemoteLockActionResult) GetUnlockPinOk() (string, bool) {
	if o == nil || o.UnlockPin == nil {
		var ret string
		return ret, false
	}
	return *o.UnlockPin, true
}

// HasUnlockPin returns a boolean if a field has been set.
func (o *RemoteLockActionResult) HasUnlockPin() bool {
	if o != nil && o.UnlockPin != nil {
		return true
	}

	return false
}

// SetUnlockPin gets a reference to the given string and assigns it to the UnlockPin field.
func (o *RemoteLockActionResult) SetUnlockPin(v string) {
	o.UnlockPin = &v
}

// SetUnlockPinExplicitNull (un)sets UnlockPin to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UnlockPin value is set to nil even if false is passed
func (o *RemoteLockActionResult) SetUnlockPinExplicitNull(b bool) {
	o.UnlockPin = nil
	o.isExplicitNullUnlockPin = b
}

// MarshalJSON returns the JSON representation of the model.
func (o RemoteLockActionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnlockPin == nil {
		if o.isExplicitNullUnlockPin {
			toSerialize["unlockPin"] = o.UnlockPin
		}
	} else {
		toSerialize["unlockPin"] = o.UnlockPin
	}
	return json.Marshal(toSerialize)
}


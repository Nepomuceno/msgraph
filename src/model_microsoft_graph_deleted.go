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
// MicrosoftGraphDeleted struct for MicrosoftGraphDeleted
type MicrosoftGraphDeleted struct {
	State *string `json:"state,omitempty"`
	isExplicitNullState bool `json:"-"`
}

// GetState returns the State field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeleted) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeleted) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphDeleted) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MicrosoftGraphDeleted) SetState(v string) {
	o.State = &v
}

// SetStateExplicitNull (un)sets State to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The State value is set to nil even if false is passed
func (o *MicrosoftGraphDeleted) SetStateExplicitNull(b bool) {
	o.State = nil
	o.isExplicitNullState = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDeleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State == nil {
		if o.isExplicitNullState {
			toSerialize["state"] = o.State
		}
	} else {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}



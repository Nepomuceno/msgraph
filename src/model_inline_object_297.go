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
// InlineObject297 struct for InlineObject297
type InlineObject297 struct {
	KeepUserData *bool `json:"keepUserData,omitempty"`

}

// GetKeepUserData returns the KeepUserData field if non-nil, zero value otherwise.
func (o *InlineObject297) GetKeepUserData() bool {
	if o == nil || o.KeepUserData == nil {
		var ret bool
		return ret
	}
	return *o.KeepUserData
}

// GetKeepUserDataOk returns a tuple with the KeepUserData field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject297) GetKeepUserDataOk() (bool, bool) {
	if o == nil || o.KeepUserData == nil {
		var ret bool
		return ret, false
	}
	return *o.KeepUserData, true
}

// HasKeepUserData returns a boolean if a field has been set.
func (o *InlineObject297) HasKeepUserData() bool {
	if o != nil && o.KeepUserData != nil {
		return true
	}

	return false
}

// SetKeepUserData gets a reference to the given bool and assigns it to the KeepUserData field.
func (o *InlineObject297) SetKeepUserData(v bool) {
	o.KeepUserData = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject297) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeepUserData != nil {
		toSerialize["keepUserData"] = o.KeepUserData
	}
	return json.Marshal(toSerialize)
}



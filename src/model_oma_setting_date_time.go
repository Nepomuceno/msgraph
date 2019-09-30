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
// OmaSettingDateTime struct for OmaSettingDateTime
type OmaSettingDateTime struct {
	// Value.
	Value *time.Time `json:"value,omitempty"`

}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *OmaSettingDateTime) GetValue() time.Time {
	if o == nil || o.Value == nil {
		var ret time.Time
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OmaSettingDateTime) GetValueOk() (time.Time, bool) {
	if o == nil || o.Value == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OmaSettingDateTime) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given time.Time and assigns it to the Value field.
func (o *OmaSettingDateTime) SetValue(v time.Time) {
	o.Value = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o OmaSettingDateTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}



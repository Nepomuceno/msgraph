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
// InlineObject968 struct for InlineObject968
type InlineObject968 struct {
	SerialNumber *AnyOfobject `json:"serialNumber,omitempty"`
	isExplicitNullSerialNumber bool `json:"-"`
}

// GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.
func (o *InlineObject968) GetSerialNumber() AnyOfobject {
	if o == nil || o.SerialNumber == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject968) GetSerialNumberOk() (AnyOfobject, bool) {
	if o == nil || o.SerialNumber == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *InlineObject968) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given AnyOfobject and assigns it to the SerialNumber field.
func (o *InlineObject968) SetSerialNumber(v AnyOfobject) {
	o.SerialNumber = &v
}

// SetSerialNumberExplicitNull (un)sets SerialNumber to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SerialNumber value is set to nil even if false is passed
func (o *InlineObject968) SetSerialNumberExplicitNull(b bool) {
	o.SerialNumber = nil
	o.isExplicitNullSerialNumber = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject968) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SerialNumber == nil {
		if o.isExplicitNullSerialNumber {
			toSerialize["serialNumber"] = o.SerialNumber
		}
	} else {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	return json.Marshal(toSerialize)
}



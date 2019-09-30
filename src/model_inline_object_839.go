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
// InlineObject839 struct for InlineObject839
type InlineObject839 struct {
	Angle *AnyOfobject `json:"angle,omitempty"`
	isExplicitNullAngle bool `json:"-"`
}

// GetAngle returns the Angle field if non-nil, zero value otherwise.
func (o *InlineObject839) GetAngle() AnyOfobject {
	if o == nil || o.Angle == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Angle
}

// GetAngleOk returns a tuple with the Angle field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject839) GetAngleOk() (AnyOfobject, bool) {
	if o == nil || o.Angle == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Angle, true
}

// HasAngle returns a boolean if a field has been set.
func (o *InlineObject839) HasAngle() bool {
	if o != nil && o.Angle != nil {
		return true
	}

	return false
}

// SetAngle gets a reference to the given AnyOfobject and assigns it to the Angle field.
func (o *InlineObject839) SetAngle(v AnyOfobject) {
	o.Angle = &v
}

// SetAngleExplicitNull (un)sets Angle to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Angle value is set to nil even if false is passed
func (o *InlineObject839) SetAngleExplicitNull(b bool) {
	o.Angle = nil
	o.isExplicitNullAngle = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject839) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Angle == nil {
		if o.isExplicitNullAngle {
			toSerialize["angle"] = o.Angle
		}
	} else {
		toSerialize["angle"] = o.Angle
	}
	return json.Marshal(toSerialize)
}



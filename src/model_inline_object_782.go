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
// InlineObject782 struct for InlineObject782
type InlineObject782 struct {
	Number1 *AnyOfobject `json:"number1,omitempty"`
	isExplicitNullNumber1 bool `json:"-"`
	Number2 *AnyOfobject `json:"number2,omitempty"`
	isExplicitNullNumber2 bool `json:"-"`
}

// GetNumber1 returns the Number1 field if non-nil, zero value otherwise.
func (o *InlineObject782) GetNumber1() AnyOfobject {
	if o == nil || o.Number1 == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Number1
}

// GetNumber1Ok returns a tuple with the Number1 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject782) GetNumber1Ok() (AnyOfobject, bool) {
	if o == nil || o.Number1 == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Number1, true
}

// HasNumber1 returns a boolean if a field has been set.
func (o *InlineObject782) HasNumber1() bool {
	if o != nil && o.Number1 != nil {
		return true
	}

	return false
}

// SetNumber1 gets a reference to the given AnyOfobject and assigns it to the Number1 field.
func (o *InlineObject782) SetNumber1(v AnyOfobject) {
	o.Number1 = &v
}

// SetNumber1ExplicitNull (un)sets Number1 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Number1 value is set to nil even if false is passed
func (o *InlineObject782) SetNumber1ExplicitNull(b bool) {
	o.Number1 = nil
	o.isExplicitNullNumber1 = b
}
// GetNumber2 returns the Number2 field if non-nil, zero value otherwise.
func (o *InlineObject782) GetNumber2() AnyOfobject {
	if o == nil || o.Number2 == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Number2
}

// GetNumber2Ok returns a tuple with the Number2 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject782) GetNumber2Ok() (AnyOfobject, bool) {
	if o == nil || o.Number2 == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Number2, true
}

// HasNumber2 returns a boolean if a field has been set.
func (o *InlineObject782) HasNumber2() bool {
	if o != nil && o.Number2 != nil {
		return true
	}

	return false
}

// SetNumber2 gets a reference to the given AnyOfobject and assigns it to the Number2 field.
func (o *InlineObject782) SetNumber2(v AnyOfobject) {
	o.Number2 = &v
}

// SetNumber2ExplicitNull (un)sets Number2 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Number2 value is set to nil even if false is passed
func (o *InlineObject782) SetNumber2ExplicitNull(b bool) {
	o.Number2 = nil
	o.isExplicitNullNumber2 = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject782) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number1 == nil {
		if o.isExplicitNullNumber1 {
			toSerialize["number1"] = o.Number1
		}
	} else {
		toSerialize["number1"] = o.Number1
	}
	if o.Number2 == nil {
		if o.isExplicitNullNumber2 {
			toSerialize["number2"] = o.Number2
		}
	} else {
		toSerialize["number2"] = o.Number2
	}
	return json.Marshal(toSerialize)
}


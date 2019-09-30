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
// InlineObject970 struct for InlineObject970
type InlineObject970 struct {
	Number *AnyOfobject `json:"number,omitempty"`
	isExplicitNullNumber bool `json:"-"`
	Divisor *AnyOfobject `json:"divisor,omitempty"`
	isExplicitNullDivisor bool `json:"-"`
}

// GetNumber returns the Number field if non-nil, zero value otherwise.
func (o *InlineObject970) GetNumber() AnyOfobject {
	if o == nil || o.Number == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject970) GetNumberOk() (AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject970) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject970) SetNumber(v AnyOfobject) {
	o.Number = &v
}

// SetNumberExplicitNull (un)sets Number to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Number value is set to nil even if false is passed
func (o *InlineObject970) SetNumberExplicitNull(b bool) {
	o.Number = nil
	o.isExplicitNullNumber = b
}
// GetDivisor returns the Divisor field if non-nil, zero value otherwise.
func (o *InlineObject970) GetDivisor() AnyOfobject {
	if o == nil || o.Divisor == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Divisor
}

// GetDivisorOk returns a tuple with the Divisor field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject970) GetDivisorOk() (AnyOfobject, bool) {
	if o == nil || o.Divisor == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Divisor, true
}

// HasDivisor returns a boolean if a field has been set.
func (o *InlineObject970) HasDivisor() bool {
	if o != nil && o.Divisor != nil {
		return true
	}

	return false
}

// SetDivisor gets a reference to the given AnyOfobject and assigns it to the Divisor field.
func (o *InlineObject970) SetDivisor(v AnyOfobject) {
	o.Divisor = &v
}

// SetDivisorExplicitNull (un)sets Divisor to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Divisor value is set to nil even if false is passed
func (o *InlineObject970) SetDivisorExplicitNull(b bool) {
	o.Divisor = nil
	o.isExplicitNullDivisor = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject970) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number == nil {
		if o.isExplicitNullNumber {
			toSerialize["number"] = o.Number
		}
	} else {
		toSerialize["number"] = o.Number
	}
	if o.Divisor == nil {
		if o.isExplicitNullDivisor {
			toSerialize["divisor"] = o.Divisor
		}
	} else {
		toSerialize["divisor"] = o.Divisor
	}
	return json.Marshal(toSerialize)
}



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
// InlineObject953 struct for InlineObject953
type InlineObject953 struct {
	Number *AnyOfobject `json:"number,omitempty"`
	isExplicitNullNumber bool `json:"-"`
	Base *AnyOfobject `json:"base,omitempty"`
	isExplicitNullBase bool `json:"-"`
}

// GetNumber returns the Number field if non-nil, zero value otherwise.
func (o *InlineObject953) GetNumber() AnyOfobject {
	if o == nil || o.Number == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject953) GetNumberOk() (AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject953) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject953) SetNumber(v AnyOfobject) {
	o.Number = &v
}

// SetNumberExplicitNull (un)sets Number to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Number value is set to nil even if false is passed
func (o *InlineObject953) SetNumberExplicitNull(b bool) {
	o.Number = nil
	o.isExplicitNullNumber = b
}
// GetBase returns the Base field if non-nil, zero value otherwise.
func (o *InlineObject953) GetBase() AnyOfobject {
	if o == nil || o.Base == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject953) GetBaseOk() (AnyOfobject, bool) {
	if o == nil || o.Base == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *InlineObject953) HasBase() bool {
	if o != nil && o.Base != nil {
		return true
	}

	return false
}

// SetBase gets a reference to the given AnyOfobject and assigns it to the Base field.
func (o *InlineObject953) SetBase(v AnyOfobject) {
	o.Base = &v
}

// SetBaseExplicitNull (un)sets Base to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Base value is set to nil even if false is passed
func (o *InlineObject953) SetBaseExplicitNull(b bool) {
	o.Base = nil
	o.isExplicitNullBase = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject953) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number == nil {
		if o.isExplicitNullNumber {
			toSerialize["number"] = o.Number
		}
	} else {
		toSerialize["number"] = o.Number
	}
	if o.Base == nil {
		if o.isExplicitNullBase {
			toSerialize["base"] = o.Base
		}
	} else {
		toSerialize["base"] = o.Base
	}
	return json.Marshal(toSerialize)
}



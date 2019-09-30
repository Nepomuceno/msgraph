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
// InlineObject770 struct for InlineObject770
type InlineObject770 struct {
	X *AnyOfobject `json:"x,omitempty"`
	isExplicitNullX bool `json:"-"`
	N *AnyOfobject `json:"n,omitempty"`
	isExplicitNullN bool `json:"-"`
}

// GetX returns the X field if non-nil, zero value otherwise.
func (o *InlineObject770) GetX() AnyOfobject {
	if o == nil || o.X == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject770) GetXOk() (AnyOfobject, bool) {
	if o == nil || o.X == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject770) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject770) SetX(v AnyOfobject) {
	o.X = &v
}

// SetXExplicitNull (un)sets X to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The X value is set to nil even if false is passed
func (o *InlineObject770) SetXExplicitNull(b bool) {
	o.X = nil
	o.isExplicitNullX = b
}
// GetN returns the N field if non-nil, zero value otherwise.
func (o *InlineObject770) GetN() AnyOfobject {
	if o == nil || o.N == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject770) GetNOk() (AnyOfobject, bool) {
	if o == nil || o.N == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *InlineObject770) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given AnyOfobject and assigns it to the N field.
func (o *InlineObject770) SetN(v AnyOfobject) {
	o.N = &v
}

// SetNExplicitNull (un)sets N to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The N value is set to nil even if false is passed
func (o *InlineObject770) SetNExplicitNull(b bool) {
	o.N = nil
	o.isExplicitNullN = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject770) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X == nil {
		if o.isExplicitNullX {
			toSerialize["x"] = o.X
		}
	} else {
		toSerialize["x"] = o.X
	}
	if o.N == nil {
		if o.isExplicitNullN {
			toSerialize["n"] = o.N
		}
	} else {
		toSerialize["n"] = o.N
	}
	return json.Marshal(toSerialize)
}



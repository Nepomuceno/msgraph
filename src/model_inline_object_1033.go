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
// InlineObject1033 struct for InlineObject1033
type InlineObject1033 struct {
	Nper *AnyOfobject `json:"nper,omitempty"`
	isExplicitNullNper bool `json:"-"`
	Pv *AnyOfobject `json:"pv,omitempty"`
	isExplicitNullPv bool `json:"-"`
	Fv *AnyOfobject `json:"fv,omitempty"`
	isExplicitNullFv bool `json:"-"`
}

// GetNper returns the Nper field if non-nil, zero value otherwise.
func (o *InlineObject1033) GetNper() AnyOfobject {
	if o == nil || o.Nper == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Nper
}

// GetNperOk returns a tuple with the Nper field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1033) GetNperOk() (AnyOfobject, bool) {
	if o == nil || o.Nper == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Nper, true
}

// HasNper returns a boolean if a field has been set.
func (o *InlineObject1033) HasNper() bool {
	if o != nil && o.Nper != nil {
		return true
	}

	return false
}

// SetNper gets a reference to the given AnyOfobject and assigns it to the Nper field.
func (o *InlineObject1033) SetNper(v AnyOfobject) {
	o.Nper = &v
}

// SetNperExplicitNull (un)sets Nper to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Nper value is set to nil even if false is passed
func (o *InlineObject1033) SetNperExplicitNull(b bool) {
	o.Nper = nil
	o.isExplicitNullNper = b
}
// GetPv returns the Pv field if non-nil, zero value otherwise.
func (o *InlineObject1033) GetPv() AnyOfobject {
	if o == nil || o.Pv == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Pv
}

// GetPvOk returns a tuple with the Pv field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1033) GetPvOk() (AnyOfobject, bool) {
	if o == nil || o.Pv == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Pv, true
}

// HasPv returns a boolean if a field has been set.
func (o *InlineObject1033) HasPv() bool {
	if o != nil && o.Pv != nil {
		return true
	}

	return false
}

// SetPv gets a reference to the given AnyOfobject and assigns it to the Pv field.
func (o *InlineObject1033) SetPv(v AnyOfobject) {
	o.Pv = &v
}

// SetPvExplicitNull (un)sets Pv to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Pv value is set to nil even if false is passed
func (o *InlineObject1033) SetPvExplicitNull(b bool) {
	o.Pv = nil
	o.isExplicitNullPv = b
}
// GetFv returns the Fv field if non-nil, zero value otherwise.
func (o *InlineObject1033) GetFv() AnyOfobject {
	if o == nil || o.Fv == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Fv
}

// GetFvOk returns a tuple with the Fv field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1033) GetFvOk() (AnyOfobject, bool) {
	if o == nil || o.Fv == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Fv, true
}

// HasFv returns a boolean if a field has been set.
func (o *InlineObject1033) HasFv() bool {
	if o != nil && o.Fv != nil {
		return true
	}

	return false
}

// SetFv gets a reference to the given AnyOfobject and assigns it to the Fv field.
func (o *InlineObject1033) SetFv(v AnyOfobject) {
	o.Fv = &v
}

// SetFvExplicitNull (un)sets Fv to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Fv value is set to nil even if false is passed
func (o *InlineObject1033) SetFvExplicitNull(b bool) {
	o.Fv = nil
	o.isExplicitNullFv = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1033) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nper == nil {
		if o.isExplicitNullNper {
			toSerialize["nper"] = o.Nper
		}
	} else {
		toSerialize["nper"] = o.Nper
	}
	if o.Pv == nil {
		if o.isExplicitNullPv {
			toSerialize["pv"] = o.Pv
		}
	} else {
		toSerialize["pv"] = o.Pv
	}
	if o.Fv == nil {
		if o.isExplicitNullFv {
			toSerialize["fv"] = o.Fv
		}
	} else {
		toSerialize["fv"] = o.Fv
	}
	return json.Marshal(toSerialize)
}



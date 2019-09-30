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
// InlineObject996 struct for InlineObject996
type InlineObject996 struct {
	Rate *AnyOfobject `json:"rate,omitempty"`
	isExplicitNullRate bool `json:"-"`
	Pv *AnyOfobject `json:"pv,omitempty"`
	isExplicitNullPv bool `json:"-"`
	Fv *AnyOfobject `json:"fv,omitempty"`
	isExplicitNullFv bool `json:"-"`
}

// GetRate returns the Rate field if non-nil, zero value otherwise.
func (o *InlineObject996) GetRate() AnyOfobject {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject996) GetRateOk() (AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject996) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject996) SetRate(v AnyOfobject) {
	o.Rate = &v
}

// SetRateExplicitNull (un)sets Rate to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Rate value is set to nil even if false is passed
func (o *InlineObject996) SetRateExplicitNull(b bool) {
	o.Rate = nil
	o.isExplicitNullRate = b
}
// GetPv returns the Pv field if non-nil, zero value otherwise.
func (o *InlineObject996) GetPv() AnyOfobject {
	if o == nil || o.Pv == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Pv
}

// GetPvOk returns a tuple with the Pv field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject996) GetPvOk() (AnyOfobject, bool) {
	if o == nil || o.Pv == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Pv, true
}

// HasPv returns a boolean if a field has been set.
func (o *InlineObject996) HasPv() bool {
	if o != nil && o.Pv != nil {
		return true
	}

	return false
}

// SetPv gets a reference to the given AnyOfobject and assigns it to the Pv field.
func (o *InlineObject996) SetPv(v AnyOfobject) {
	o.Pv = &v
}

// SetPvExplicitNull (un)sets Pv to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Pv value is set to nil even if false is passed
func (o *InlineObject996) SetPvExplicitNull(b bool) {
	o.Pv = nil
	o.isExplicitNullPv = b
}
// GetFv returns the Fv field if non-nil, zero value otherwise.
func (o *InlineObject996) GetFv() AnyOfobject {
	if o == nil || o.Fv == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Fv
}

// GetFvOk returns a tuple with the Fv field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject996) GetFvOk() (AnyOfobject, bool) {
	if o == nil || o.Fv == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Fv, true
}

// HasFv returns a boolean if a field has been set.
func (o *InlineObject996) HasFv() bool {
	if o != nil && o.Fv != nil {
		return true
	}

	return false
}

// SetFv gets a reference to the given AnyOfobject and assigns it to the Fv field.
func (o *InlineObject996) SetFv(v AnyOfobject) {
	o.Fv = &v
}

// SetFvExplicitNull (un)sets Fv to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Fv value is set to nil even if false is passed
func (o *InlineObject996) SetFvExplicitNull(b bool) {
	o.Fv = nil
	o.isExplicitNullFv = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject996) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rate == nil {
		if o.isExplicitNullRate {
			toSerialize["rate"] = o.Rate
		}
	} else {
		toSerialize["rate"] = o.Rate
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


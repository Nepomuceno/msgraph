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
// InlineObject942 struct for InlineObject942
type InlineObject942 struct {
	Rate *AnyOfobject `json:"rate,omitempty"`
	isExplicitNullRate bool `json:"-"`
	Per *AnyOfobject `json:"per,omitempty"`
	isExplicitNullPer bool `json:"-"`
	Nper *AnyOfobject `json:"nper,omitempty"`
	isExplicitNullNper bool `json:"-"`
	Pv *AnyOfobject `json:"pv,omitempty"`
	isExplicitNullPv bool `json:"-"`
}

// GetRate returns the Rate field if non-nil, zero value otherwise.
func (o *InlineObject942) GetRate() AnyOfobject {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject942) GetRateOk() (AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject942) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject942) SetRate(v AnyOfobject) {
	o.Rate = &v
}

// SetRateExplicitNull (un)sets Rate to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Rate value is set to nil even if false is passed
func (o *InlineObject942) SetRateExplicitNull(b bool) {
	o.Rate = nil
	o.isExplicitNullRate = b
}
// GetPer returns the Per field if non-nil, zero value otherwise.
func (o *InlineObject942) GetPer() AnyOfobject {
	if o == nil || o.Per == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Per
}

// GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject942) GetPerOk() (AnyOfobject, bool) {
	if o == nil || o.Per == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Per, true
}

// HasPer returns a boolean if a field has been set.
func (o *InlineObject942) HasPer() bool {
	if o != nil && o.Per != nil {
		return true
	}

	return false
}

// SetPer gets a reference to the given AnyOfobject and assigns it to the Per field.
func (o *InlineObject942) SetPer(v AnyOfobject) {
	o.Per = &v
}

// SetPerExplicitNull (un)sets Per to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Per value is set to nil even if false is passed
func (o *InlineObject942) SetPerExplicitNull(b bool) {
	o.Per = nil
	o.isExplicitNullPer = b
}
// GetNper returns the Nper field if non-nil, zero value otherwise.
func (o *InlineObject942) GetNper() AnyOfobject {
	if o == nil || o.Nper == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Nper
}

// GetNperOk returns a tuple with the Nper field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject942) GetNperOk() (AnyOfobject, bool) {
	if o == nil || o.Nper == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Nper, true
}

// HasNper returns a boolean if a field has been set.
func (o *InlineObject942) HasNper() bool {
	if o != nil && o.Nper != nil {
		return true
	}

	return false
}

// SetNper gets a reference to the given AnyOfobject and assigns it to the Nper field.
func (o *InlineObject942) SetNper(v AnyOfobject) {
	o.Nper = &v
}

// SetNperExplicitNull (un)sets Nper to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Nper value is set to nil even if false is passed
func (o *InlineObject942) SetNperExplicitNull(b bool) {
	o.Nper = nil
	o.isExplicitNullNper = b
}
// GetPv returns the Pv field if non-nil, zero value otherwise.
func (o *InlineObject942) GetPv() AnyOfobject {
	if o == nil || o.Pv == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Pv
}

// GetPvOk returns a tuple with the Pv field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject942) GetPvOk() (AnyOfobject, bool) {
	if o == nil || o.Pv == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Pv, true
}

// HasPv returns a boolean if a field has been set.
func (o *InlineObject942) HasPv() bool {
	if o != nil && o.Pv != nil {
		return true
	}

	return false
}

// SetPv gets a reference to the given AnyOfobject and assigns it to the Pv field.
func (o *InlineObject942) SetPv(v AnyOfobject) {
	o.Pv = &v
}

// SetPvExplicitNull (un)sets Pv to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Pv value is set to nil even if false is passed
func (o *InlineObject942) SetPvExplicitNull(b bool) {
	o.Pv = nil
	o.isExplicitNullPv = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject942) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rate == nil {
		if o.isExplicitNullRate {
			toSerialize["rate"] = o.Rate
		}
	} else {
		toSerialize["rate"] = o.Rate
	}
	if o.Per == nil {
		if o.isExplicitNullPer {
			toSerialize["per"] = o.Per
		}
	} else {
		toSerialize["per"] = o.Per
	}
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
	return json.Marshal(toSerialize)
}


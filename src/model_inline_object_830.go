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
// InlineObject830 struct for InlineObject830
type InlineObject830 struct {
	Cost *AnyOfobject `json:"cost,omitempty"`
	isExplicitNullCost bool `json:"-"`
	Salvage *AnyOfobject `json:"salvage,omitempty"`
	isExplicitNullSalvage bool `json:"-"`
	Life *AnyOfobject `json:"life,omitempty"`
	isExplicitNullLife bool `json:"-"`
	Period *AnyOfobject `json:"period,omitempty"`
	isExplicitNullPeriod bool `json:"-"`
	Month *AnyOfobject `json:"month,omitempty"`
	isExplicitNullMonth bool `json:"-"`
}

// GetCost returns the Cost field if non-nil, zero value otherwise.
func (o *InlineObject830) GetCost() AnyOfobject {
	if o == nil || o.Cost == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject830) GetCostOk() (AnyOfobject, bool) {
	if o == nil || o.Cost == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *InlineObject830) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given AnyOfobject and assigns it to the Cost field.
func (o *InlineObject830) SetCost(v AnyOfobject) {
	o.Cost = &v
}

// SetCostExplicitNull (un)sets Cost to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Cost value is set to nil even if false is passed
func (o *InlineObject830) SetCostExplicitNull(b bool) {
	o.Cost = nil
	o.isExplicitNullCost = b
}
// GetSalvage returns the Salvage field if non-nil, zero value otherwise.
func (o *InlineObject830) GetSalvage() AnyOfobject {
	if o == nil || o.Salvage == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Salvage
}

// GetSalvageOk returns a tuple with the Salvage field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject830) GetSalvageOk() (AnyOfobject, bool) {
	if o == nil || o.Salvage == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Salvage, true
}

// HasSalvage returns a boolean if a field has been set.
func (o *InlineObject830) HasSalvage() bool {
	if o != nil && o.Salvage != nil {
		return true
	}

	return false
}

// SetSalvage gets a reference to the given AnyOfobject and assigns it to the Salvage field.
func (o *InlineObject830) SetSalvage(v AnyOfobject) {
	o.Salvage = &v
}

// SetSalvageExplicitNull (un)sets Salvage to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Salvage value is set to nil even if false is passed
func (o *InlineObject830) SetSalvageExplicitNull(b bool) {
	o.Salvage = nil
	o.isExplicitNullSalvage = b
}
// GetLife returns the Life field if non-nil, zero value otherwise.
func (o *InlineObject830) GetLife() AnyOfobject {
	if o == nil || o.Life == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Life
}

// GetLifeOk returns a tuple with the Life field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject830) GetLifeOk() (AnyOfobject, bool) {
	if o == nil || o.Life == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Life, true
}

// HasLife returns a boolean if a field has been set.
func (o *InlineObject830) HasLife() bool {
	if o != nil && o.Life != nil {
		return true
	}

	return false
}

// SetLife gets a reference to the given AnyOfobject and assigns it to the Life field.
func (o *InlineObject830) SetLife(v AnyOfobject) {
	o.Life = &v
}

// SetLifeExplicitNull (un)sets Life to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Life value is set to nil even if false is passed
func (o *InlineObject830) SetLifeExplicitNull(b bool) {
	o.Life = nil
	o.isExplicitNullLife = b
}
// GetPeriod returns the Period field if non-nil, zero value otherwise.
func (o *InlineObject830) GetPeriod() AnyOfobject {
	if o == nil || o.Period == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject830) GetPeriodOk() (AnyOfobject, bool) {
	if o == nil || o.Period == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *InlineObject830) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given AnyOfobject and assigns it to the Period field.
func (o *InlineObject830) SetPeriod(v AnyOfobject) {
	o.Period = &v
}

// SetPeriodExplicitNull (un)sets Period to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Period value is set to nil even if false is passed
func (o *InlineObject830) SetPeriodExplicitNull(b bool) {
	o.Period = nil
	o.isExplicitNullPeriod = b
}
// GetMonth returns the Month field if non-nil, zero value otherwise.
func (o *InlineObject830) GetMonth() AnyOfobject {
	if o == nil || o.Month == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject830) GetMonthOk() (AnyOfobject, bool) {
	if o == nil || o.Month == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *InlineObject830) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given AnyOfobject and assigns it to the Month field.
func (o *InlineObject830) SetMonth(v AnyOfobject) {
	o.Month = &v
}

// SetMonthExplicitNull (un)sets Month to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Month value is set to nil even if false is passed
func (o *InlineObject830) SetMonthExplicitNull(b bool) {
	o.Month = nil
	o.isExplicitNullMonth = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject830) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cost == nil {
		if o.isExplicitNullCost {
			toSerialize["cost"] = o.Cost
		}
	} else {
		toSerialize["cost"] = o.Cost
	}
	if o.Salvage == nil {
		if o.isExplicitNullSalvage {
			toSerialize["salvage"] = o.Salvage
		}
	} else {
		toSerialize["salvage"] = o.Salvage
	}
	if o.Life == nil {
		if o.isExplicitNullLife {
			toSerialize["life"] = o.Life
		}
	} else {
		toSerialize["life"] = o.Life
	}
	if o.Period == nil {
		if o.isExplicitNullPeriod {
			toSerialize["period"] = o.Period
		}
	} else {
		toSerialize["period"] = o.Period
	}
	if o.Month == nil {
		if o.isExplicitNullMonth {
			toSerialize["month"] = o.Month
		}
	} else {
		toSerialize["month"] = o.Month
	}
	return json.Marshal(toSerialize)
}


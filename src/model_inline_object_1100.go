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
// InlineObject1100 struct for InlineObject1100
type InlineObject1100 struct {
	Settlement *AnyOfobject `json:"settlement,omitempty"`
	isExplicitNullSettlement bool `json:"-"`
	Maturity *AnyOfobject `json:"maturity,omitempty"`
	isExplicitNullMaturity bool `json:"-"`
	Rate *AnyOfobject `json:"rate,omitempty"`
	isExplicitNullRate bool `json:"-"`
	Pr *AnyOfobject `json:"pr,omitempty"`
	isExplicitNullPr bool `json:"-"`
	Redemption *AnyOfobject `json:"redemption,omitempty"`
	isExplicitNullRedemption bool `json:"-"`
	Frequency *AnyOfobject `json:"frequency,omitempty"`
	isExplicitNullFrequency bool `json:"-"`
	Basis *AnyOfobject `json:"basis,omitempty"`
	isExplicitNullBasis bool `json:"-"`
}

// GetSettlement returns the Settlement field if non-nil, zero value otherwise.
func (o *InlineObject1100) GetSettlement() AnyOfobject {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1100) GetSettlementOk() (AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1100) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1100) SetSettlement(v AnyOfobject) {
	o.Settlement = &v
}

// SetSettlementExplicitNull (un)sets Settlement to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Settlement value is set to nil even if false is passed
func (o *InlineObject1100) SetSettlementExplicitNull(b bool) {
	o.Settlement = nil
	o.isExplicitNullSettlement = b
}
// GetMaturity returns the Maturity field if non-nil, zero value otherwise.
func (o *InlineObject1100) GetMaturity() AnyOfobject {
	if o == nil || o.Maturity == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1100) GetMaturityOk() (AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1100) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1100) SetMaturity(v AnyOfobject) {
	o.Maturity = &v
}

// SetMaturityExplicitNull (un)sets Maturity to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Maturity value is set to nil even if false is passed
func (o *InlineObject1100) SetMaturityExplicitNull(b bool) {
	o.Maturity = nil
	o.isExplicitNullMaturity = b
}
// GetRate returns the Rate field if non-nil, zero value otherwise.
func (o *InlineObject1100) GetRate() AnyOfobject {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1100) GetRateOk() (AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1100) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1100) SetRate(v AnyOfobject) {
	o.Rate = &v
}

// SetRateExplicitNull (un)sets Rate to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Rate value is set to nil even if false is passed
func (o *InlineObject1100) SetRateExplicitNull(b bool) {
	o.Rate = nil
	o.isExplicitNullRate = b
}
// GetPr returns the Pr field if non-nil, zero value otherwise.
func (o *InlineObject1100) GetPr() AnyOfobject {
	if o == nil || o.Pr == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Pr
}

// GetPrOk returns a tuple with the Pr field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1100) GetPrOk() (AnyOfobject, bool) {
	if o == nil || o.Pr == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Pr, true
}

// HasPr returns a boolean if a field has been set.
func (o *InlineObject1100) HasPr() bool {
	if o != nil && o.Pr != nil {
		return true
	}

	return false
}

// SetPr gets a reference to the given AnyOfobject and assigns it to the Pr field.
func (o *InlineObject1100) SetPr(v AnyOfobject) {
	o.Pr = &v
}

// SetPrExplicitNull (un)sets Pr to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Pr value is set to nil even if false is passed
func (o *InlineObject1100) SetPrExplicitNull(b bool) {
	o.Pr = nil
	o.isExplicitNullPr = b
}
// GetRedemption returns the Redemption field if non-nil, zero value otherwise.
func (o *InlineObject1100) GetRedemption() AnyOfobject {
	if o == nil || o.Redemption == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Redemption
}

// GetRedemptionOk returns a tuple with the Redemption field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1100) GetRedemptionOk() (AnyOfobject, bool) {
	if o == nil || o.Redemption == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Redemption, true
}

// HasRedemption returns a boolean if a field has been set.
func (o *InlineObject1100) HasRedemption() bool {
	if o != nil && o.Redemption != nil {
		return true
	}

	return false
}

// SetRedemption gets a reference to the given AnyOfobject and assigns it to the Redemption field.
func (o *InlineObject1100) SetRedemption(v AnyOfobject) {
	o.Redemption = &v
}

// SetRedemptionExplicitNull (un)sets Redemption to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Redemption value is set to nil even if false is passed
func (o *InlineObject1100) SetRedemptionExplicitNull(b bool) {
	o.Redemption = nil
	o.isExplicitNullRedemption = b
}
// GetFrequency returns the Frequency field if non-nil, zero value otherwise.
func (o *InlineObject1100) GetFrequency() AnyOfobject {
	if o == nil || o.Frequency == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1100) GetFrequencyOk() (AnyOfobject, bool) {
	if o == nil || o.Frequency == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *InlineObject1100) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given AnyOfobject and assigns it to the Frequency field.
func (o *InlineObject1100) SetFrequency(v AnyOfobject) {
	o.Frequency = &v
}

// SetFrequencyExplicitNull (un)sets Frequency to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Frequency value is set to nil even if false is passed
func (o *InlineObject1100) SetFrequencyExplicitNull(b bool) {
	o.Frequency = nil
	o.isExplicitNullFrequency = b
}
// GetBasis returns the Basis field if non-nil, zero value otherwise.
func (o *InlineObject1100) GetBasis() AnyOfobject {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Basis
}

// GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1100) GetBasisOk() (AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject1100) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject1100) SetBasis(v AnyOfobject) {
	o.Basis = &v
}

// SetBasisExplicitNull (un)sets Basis to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Basis value is set to nil even if false is passed
func (o *InlineObject1100) SetBasisExplicitNull(b bool) {
	o.Basis = nil
	o.isExplicitNullBasis = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1100) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settlement == nil {
		if o.isExplicitNullSettlement {
			toSerialize["settlement"] = o.Settlement
		}
	} else {
		toSerialize["settlement"] = o.Settlement
	}
	if o.Maturity == nil {
		if o.isExplicitNullMaturity {
			toSerialize["maturity"] = o.Maturity
		}
	} else {
		toSerialize["maturity"] = o.Maturity
	}
	if o.Rate == nil {
		if o.isExplicitNullRate {
			toSerialize["rate"] = o.Rate
		}
	} else {
		toSerialize["rate"] = o.Rate
	}
	if o.Pr == nil {
		if o.isExplicitNullPr {
			toSerialize["pr"] = o.Pr
		}
	} else {
		toSerialize["pr"] = o.Pr
	}
	if o.Redemption == nil {
		if o.isExplicitNullRedemption {
			toSerialize["redemption"] = o.Redemption
		}
	} else {
		toSerialize["redemption"] = o.Redemption
	}
	if o.Frequency == nil {
		if o.isExplicitNullFrequency {
			toSerialize["frequency"] = o.Frequency
		}
	} else {
		toSerialize["frequency"] = o.Frequency
	}
	if o.Basis == nil {
		if o.isExplicitNullBasis {
			toSerialize["basis"] = o.Basis
		}
	} else {
		toSerialize["basis"] = o.Basis
	}
	return json.Marshal(toSerialize)
}



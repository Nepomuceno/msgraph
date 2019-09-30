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
// InlineObject815 struct for InlineObject815
type InlineObject815 struct {
	Settlement *AnyOfobject `json:"settlement,omitempty"`
	isExplicitNullSettlement bool `json:"-"`
	Maturity *AnyOfobject `json:"maturity,omitempty"`
	isExplicitNullMaturity bool `json:"-"`
	Frequency *AnyOfobject `json:"frequency,omitempty"`
	isExplicitNullFrequency bool `json:"-"`
	Basis *AnyOfobject `json:"basis,omitempty"`
	isExplicitNullBasis bool `json:"-"`
}

// GetSettlement returns the Settlement field if non-nil, zero value otherwise.
func (o *InlineObject815) GetSettlement() AnyOfobject {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject815) GetSettlementOk() (AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject815) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject815) SetSettlement(v AnyOfobject) {
	o.Settlement = &v
}

// SetSettlementExplicitNull (un)sets Settlement to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Settlement value is set to nil even if false is passed
func (o *InlineObject815) SetSettlementExplicitNull(b bool) {
	o.Settlement = nil
	o.isExplicitNullSettlement = b
}
// GetMaturity returns the Maturity field if non-nil, zero value otherwise.
func (o *InlineObject815) GetMaturity() AnyOfobject {
	if o == nil || o.Maturity == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject815) GetMaturityOk() (AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject815) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject815) SetMaturity(v AnyOfobject) {
	o.Maturity = &v
}

// SetMaturityExplicitNull (un)sets Maturity to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Maturity value is set to nil even if false is passed
func (o *InlineObject815) SetMaturityExplicitNull(b bool) {
	o.Maturity = nil
	o.isExplicitNullMaturity = b
}
// GetFrequency returns the Frequency field if non-nil, zero value otherwise.
func (o *InlineObject815) GetFrequency() AnyOfobject {
	if o == nil || o.Frequency == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject815) GetFrequencyOk() (AnyOfobject, bool) {
	if o == nil || o.Frequency == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *InlineObject815) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given AnyOfobject and assigns it to the Frequency field.
func (o *InlineObject815) SetFrequency(v AnyOfobject) {
	o.Frequency = &v
}

// SetFrequencyExplicitNull (un)sets Frequency to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Frequency value is set to nil even if false is passed
func (o *InlineObject815) SetFrequencyExplicitNull(b bool) {
	o.Frequency = nil
	o.isExplicitNullFrequency = b
}
// GetBasis returns the Basis field if non-nil, zero value otherwise.
func (o *InlineObject815) GetBasis() AnyOfobject {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Basis
}

// GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject815) GetBasisOk() (AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject815) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject815) SetBasis(v AnyOfobject) {
	o.Basis = &v
}

// SetBasisExplicitNull (un)sets Basis to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Basis value is set to nil even if false is passed
func (o *InlineObject815) SetBasisExplicitNull(b bool) {
	o.Basis = nil
	o.isExplicitNullBasis = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject815) MarshalJSON() ([]byte, error) {
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


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
// InlineObject1022 struct for InlineObject1022
type InlineObject1022 struct {
	Settlement *AnyOfobject `json:"settlement,omitempty"`
	isExplicitNullSettlement bool `json:"-"`
	Maturity *AnyOfobject `json:"maturity,omitempty"`
	isExplicitNullMaturity bool `json:"-"`
	Investment *AnyOfobject `json:"investment,omitempty"`
	isExplicitNullInvestment bool `json:"-"`
	Discount *AnyOfobject `json:"discount,omitempty"`
	isExplicitNullDiscount bool `json:"-"`
	Basis *AnyOfobject `json:"basis,omitempty"`
	isExplicitNullBasis bool `json:"-"`
}

// GetSettlement returns the Settlement field if non-nil, zero value otherwise.
func (o *InlineObject1022) GetSettlement() AnyOfobject {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1022) GetSettlementOk() (AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1022) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1022) SetSettlement(v AnyOfobject) {
	o.Settlement = &v
}

// SetSettlementExplicitNull (un)sets Settlement to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Settlement value is set to nil even if false is passed
func (o *InlineObject1022) SetSettlementExplicitNull(b bool) {
	o.Settlement = nil
	o.isExplicitNullSettlement = b
}
// GetMaturity returns the Maturity field if non-nil, zero value otherwise.
func (o *InlineObject1022) GetMaturity() AnyOfobject {
	if o == nil || o.Maturity == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1022) GetMaturityOk() (AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1022) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1022) SetMaturity(v AnyOfobject) {
	o.Maturity = &v
}

// SetMaturityExplicitNull (un)sets Maturity to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Maturity value is set to nil even if false is passed
func (o *InlineObject1022) SetMaturityExplicitNull(b bool) {
	o.Maturity = nil
	o.isExplicitNullMaturity = b
}
// GetInvestment returns the Investment field if non-nil, zero value otherwise.
func (o *InlineObject1022) GetInvestment() AnyOfobject {
	if o == nil || o.Investment == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Investment
}

// GetInvestmentOk returns a tuple with the Investment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1022) GetInvestmentOk() (AnyOfobject, bool) {
	if o == nil || o.Investment == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Investment, true
}

// HasInvestment returns a boolean if a field has been set.
func (o *InlineObject1022) HasInvestment() bool {
	if o != nil && o.Investment != nil {
		return true
	}

	return false
}

// SetInvestment gets a reference to the given AnyOfobject and assigns it to the Investment field.
func (o *InlineObject1022) SetInvestment(v AnyOfobject) {
	o.Investment = &v
}

// SetInvestmentExplicitNull (un)sets Investment to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Investment value is set to nil even if false is passed
func (o *InlineObject1022) SetInvestmentExplicitNull(b bool) {
	o.Investment = nil
	o.isExplicitNullInvestment = b
}
// GetDiscount returns the Discount field if non-nil, zero value otherwise.
func (o *InlineObject1022) GetDiscount() AnyOfobject {
	if o == nil || o.Discount == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1022) GetDiscountOk() (AnyOfobject, bool) {
	if o == nil || o.Discount == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *InlineObject1022) HasDiscount() bool {
	if o != nil && o.Discount != nil {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given AnyOfobject and assigns it to the Discount field.
func (o *InlineObject1022) SetDiscount(v AnyOfobject) {
	o.Discount = &v
}

// SetDiscountExplicitNull (un)sets Discount to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Discount value is set to nil even if false is passed
func (o *InlineObject1022) SetDiscountExplicitNull(b bool) {
	o.Discount = nil
	o.isExplicitNullDiscount = b
}
// GetBasis returns the Basis field if non-nil, zero value otherwise.
func (o *InlineObject1022) GetBasis() AnyOfobject {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Basis
}

// GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1022) GetBasisOk() (AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject1022) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject1022) SetBasis(v AnyOfobject) {
	o.Basis = &v
}

// SetBasisExplicitNull (un)sets Basis to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Basis value is set to nil even if false is passed
func (o *InlineObject1022) SetBasisExplicitNull(b bool) {
	o.Basis = nil
	o.isExplicitNullBasis = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1022) MarshalJSON() ([]byte, error) {
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
	if o.Investment == nil {
		if o.isExplicitNullInvestment {
			toSerialize["investment"] = o.Investment
		}
	} else {
		toSerialize["investment"] = o.Investment
	}
	if o.Discount == nil {
		if o.isExplicitNullDiscount {
			toSerialize["discount"] = o.Discount
		}
	} else {
		toSerialize["discount"] = o.Discount
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



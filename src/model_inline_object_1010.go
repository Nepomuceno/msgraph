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
// InlineObject1010 struct for InlineObject1010
type InlineObject1010 struct {
	Settlement *AnyOfobject `json:"settlement,omitempty"`
	isExplicitNullSettlement bool `json:"-"`
	Maturity *AnyOfobject `json:"maturity,omitempty"`
	isExplicitNullMaturity bool `json:"-"`
	Issue *AnyOfobject `json:"issue,omitempty"`
	isExplicitNullIssue bool `json:"-"`
	Rate *AnyOfobject `json:"rate,omitempty"`
	isExplicitNullRate bool `json:"-"`
	Yld *AnyOfobject `json:"yld,omitempty"`
	isExplicitNullYld bool `json:"-"`
	Basis *AnyOfobject `json:"basis,omitempty"`
	isExplicitNullBasis bool `json:"-"`
}

// GetSettlement returns the Settlement field if non-nil, zero value otherwise.
func (o *InlineObject1010) GetSettlement() AnyOfobject {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1010) GetSettlementOk() (AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject1010) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject1010) SetSettlement(v AnyOfobject) {
	o.Settlement = &v
}

// SetSettlementExplicitNull (un)sets Settlement to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Settlement value is set to nil even if false is passed
func (o *InlineObject1010) SetSettlementExplicitNull(b bool) {
	o.Settlement = nil
	o.isExplicitNullSettlement = b
}
// GetMaturity returns the Maturity field if non-nil, zero value otherwise.
func (o *InlineObject1010) GetMaturity() AnyOfobject {
	if o == nil || o.Maturity == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Maturity
}

// GetMaturityOk returns a tuple with the Maturity field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1010) GetMaturityOk() (AnyOfobject, bool) {
	if o == nil || o.Maturity == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Maturity, true
}

// HasMaturity returns a boolean if a field has been set.
func (o *InlineObject1010) HasMaturity() bool {
	if o != nil && o.Maturity != nil {
		return true
	}

	return false
}

// SetMaturity gets a reference to the given AnyOfobject and assigns it to the Maturity field.
func (o *InlineObject1010) SetMaturity(v AnyOfobject) {
	o.Maturity = &v
}

// SetMaturityExplicitNull (un)sets Maturity to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Maturity value is set to nil even if false is passed
func (o *InlineObject1010) SetMaturityExplicitNull(b bool) {
	o.Maturity = nil
	o.isExplicitNullMaturity = b
}
// GetIssue returns the Issue field if non-nil, zero value otherwise.
func (o *InlineObject1010) GetIssue() AnyOfobject {
	if o == nil || o.Issue == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Issue
}

// GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1010) GetIssueOk() (AnyOfobject, bool) {
	if o == nil || o.Issue == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *InlineObject1010) HasIssue() bool {
	if o != nil && o.Issue != nil {
		return true
	}

	return false
}

// SetIssue gets a reference to the given AnyOfobject and assigns it to the Issue field.
func (o *InlineObject1010) SetIssue(v AnyOfobject) {
	o.Issue = &v
}

// SetIssueExplicitNull (un)sets Issue to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Issue value is set to nil even if false is passed
func (o *InlineObject1010) SetIssueExplicitNull(b bool) {
	o.Issue = nil
	o.isExplicitNullIssue = b
}
// GetRate returns the Rate field if non-nil, zero value otherwise.
func (o *InlineObject1010) GetRate() AnyOfobject {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1010) GetRateOk() (AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject1010) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject1010) SetRate(v AnyOfobject) {
	o.Rate = &v
}

// SetRateExplicitNull (un)sets Rate to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Rate value is set to nil even if false is passed
func (o *InlineObject1010) SetRateExplicitNull(b bool) {
	o.Rate = nil
	o.isExplicitNullRate = b
}
// GetYld returns the Yld field if non-nil, zero value otherwise.
func (o *InlineObject1010) GetYld() AnyOfobject {
	if o == nil || o.Yld == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Yld
}

// GetYldOk returns a tuple with the Yld field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1010) GetYldOk() (AnyOfobject, bool) {
	if o == nil || o.Yld == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Yld, true
}

// HasYld returns a boolean if a field has been set.
func (o *InlineObject1010) HasYld() bool {
	if o != nil && o.Yld != nil {
		return true
	}

	return false
}

// SetYld gets a reference to the given AnyOfobject and assigns it to the Yld field.
func (o *InlineObject1010) SetYld(v AnyOfobject) {
	o.Yld = &v
}

// SetYldExplicitNull (un)sets Yld to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Yld value is set to nil even if false is passed
func (o *InlineObject1010) SetYldExplicitNull(b bool) {
	o.Yld = nil
	o.isExplicitNullYld = b
}
// GetBasis returns the Basis field if non-nil, zero value otherwise.
func (o *InlineObject1010) GetBasis() AnyOfobject {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Basis
}

// GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1010) GetBasisOk() (AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject1010) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject1010) SetBasis(v AnyOfobject) {
	o.Basis = &v
}

// SetBasisExplicitNull (un)sets Basis to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Basis value is set to nil even if false is passed
func (o *InlineObject1010) SetBasisExplicitNull(b bool) {
	o.Basis = nil
	o.isExplicitNullBasis = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1010) MarshalJSON() ([]byte, error) {
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
	if o.Issue == nil {
		if o.isExplicitNullIssue {
			toSerialize["issue"] = o.Issue
		}
	} else {
		toSerialize["issue"] = o.Issue
	}
	if o.Rate == nil {
		if o.isExplicitNullRate {
			toSerialize["rate"] = o.Rate
		}
	} else {
		toSerialize["rate"] = o.Rate
	}
	if o.Yld == nil {
		if o.isExplicitNullYld {
			toSerialize["yld"] = o.Yld
		}
	} else {
		toSerialize["yld"] = o.Yld
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



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
// InlineObject746 struct for InlineObject746
type InlineObject746 struct {
	Issue *AnyOfobject `json:"issue,omitempty"`
	isExplicitNullIssue bool `json:"-"`
	FirstInterest *AnyOfobject `json:"firstInterest,omitempty"`
	isExplicitNullFirstInterest bool `json:"-"`
	Settlement *AnyOfobject `json:"settlement,omitempty"`
	isExplicitNullSettlement bool `json:"-"`
	Rate *AnyOfobject `json:"rate,omitempty"`
	isExplicitNullRate bool `json:"-"`
	Par *AnyOfobject `json:"par,omitempty"`
	isExplicitNullPar bool `json:"-"`
	Frequency *AnyOfobject `json:"frequency,omitempty"`
	isExplicitNullFrequency bool `json:"-"`
	Basis *AnyOfobject `json:"basis,omitempty"`
	isExplicitNullBasis bool `json:"-"`
	CalcMethod *AnyOfobject `json:"calcMethod,omitempty"`
	isExplicitNullCalcMethod bool `json:"-"`
}

// GetIssue returns the Issue field if non-nil, zero value otherwise.
func (o *InlineObject746) GetIssue() AnyOfobject {
	if o == nil || o.Issue == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Issue
}

// GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject746) GetIssueOk() (AnyOfobject, bool) {
	if o == nil || o.Issue == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *InlineObject746) HasIssue() bool {
	if o != nil && o.Issue != nil {
		return true
	}

	return false
}

// SetIssue gets a reference to the given AnyOfobject and assigns it to the Issue field.
func (o *InlineObject746) SetIssue(v AnyOfobject) {
	o.Issue = &v
}

// SetIssueExplicitNull (un)sets Issue to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Issue value is set to nil even if false is passed
func (o *InlineObject746) SetIssueExplicitNull(b bool) {
	o.Issue = nil
	o.isExplicitNullIssue = b
}
// GetFirstInterest returns the FirstInterest field if non-nil, zero value otherwise.
func (o *InlineObject746) GetFirstInterest() AnyOfobject {
	if o == nil || o.FirstInterest == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.FirstInterest
}

// GetFirstInterestOk returns a tuple with the FirstInterest field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject746) GetFirstInterestOk() (AnyOfobject, bool) {
	if o == nil || o.FirstInterest == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.FirstInterest, true
}

// HasFirstInterest returns a boolean if a field has been set.
func (o *InlineObject746) HasFirstInterest() bool {
	if o != nil && o.FirstInterest != nil {
		return true
	}

	return false
}

// SetFirstInterest gets a reference to the given AnyOfobject and assigns it to the FirstInterest field.
func (o *InlineObject746) SetFirstInterest(v AnyOfobject) {
	o.FirstInterest = &v
}

// SetFirstInterestExplicitNull (un)sets FirstInterest to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The FirstInterest value is set to nil even if false is passed
func (o *InlineObject746) SetFirstInterestExplicitNull(b bool) {
	o.FirstInterest = nil
	o.isExplicitNullFirstInterest = b
}
// GetSettlement returns the Settlement field if non-nil, zero value otherwise.
func (o *InlineObject746) GetSettlement() AnyOfobject {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Settlement
}

// GetSettlementOk returns a tuple with the Settlement field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject746) GetSettlementOk() (AnyOfobject, bool) {
	if o == nil || o.Settlement == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Settlement, true
}

// HasSettlement returns a boolean if a field has been set.
func (o *InlineObject746) HasSettlement() bool {
	if o != nil && o.Settlement != nil {
		return true
	}

	return false
}

// SetSettlement gets a reference to the given AnyOfobject and assigns it to the Settlement field.
func (o *InlineObject746) SetSettlement(v AnyOfobject) {
	o.Settlement = &v
}

// SetSettlementExplicitNull (un)sets Settlement to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Settlement value is set to nil even if false is passed
func (o *InlineObject746) SetSettlementExplicitNull(b bool) {
	o.Settlement = nil
	o.isExplicitNullSettlement = b
}
// GetRate returns the Rate field if non-nil, zero value otherwise.
func (o *InlineObject746) GetRate() AnyOfobject {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject746) GetRateOk() (AnyOfobject, bool) {
	if o == nil || o.Rate == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *InlineObject746) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given AnyOfobject and assigns it to the Rate field.
func (o *InlineObject746) SetRate(v AnyOfobject) {
	o.Rate = &v
}

// SetRateExplicitNull (un)sets Rate to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Rate value is set to nil even if false is passed
func (o *InlineObject746) SetRateExplicitNull(b bool) {
	o.Rate = nil
	o.isExplicitNullRate = b
}
// GetPar returns the Par field if non-nil, zero value otherwise.
func (o *InlineObject746) GetPar() AnyOfobject {
	if o == nil || o.Par == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Par
}

// GetParOk returns a tuple with the Par field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject746) GetParOk() (AnyOfobject, bool) {
	if o == nil || o.Par == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Par, true
}

// HasPar returns a boolean if a field has been set.
func (o *InlineObject746) HasPar() bool {
	if o != nil && o.Par != nil {
		return true
	}

	return false
}

// SetPar gets a reference to the given AnyOfobject and assigns it to the Par field.
func (o *InlineObject746) SetPar(v AnyOfobject) {
	o.Par = &v
}

// SetParExplicitNull (un)sets Par to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Par value is set to nil even if false is passed
func (o *InlineObject746) SetParExplicitNull(b bool) {
	o.Par = nil
	o.isExplicitNullPar = b
}
// GetFrequency returns the Frequency field if non-nil, zero value otherwise.
func (o *InlineObject746) GetFrequency() AnyOfobject {
	if o == nil || o.Frequency == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject746) GetFrequencyOk() (AnyOfobject, bool) {
	if o == nil || o.Frequency == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *InlineObject746) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given AnyOfobject and assigns it to the Frequency field.
func (o *InlineObject746) SetFrequency(v AnyOfobject) {
	o.Frequency = &v
}

// SetFrequencyExplicitNull (un)sets Frequency to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Frequency value is set to nil even if false is passed
func (o *InlineObject746) SetFrequencyExplicitNull(b bool) {
	o.Frequency = nil
	o.isExplicitNullFrequency = b
}
// GetBasis returns the Basis field if non-nil, zero value otherwise.
func (o *InlineObject746) GetBasis() AnyOfobject {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Basis
}

// GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject746) GetBasisOk() (AnyOfobject, bool) {
	if o == nil || o.Basis == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *InlineObject746) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given AnyOfobject and assigns it to the Basis field.
func (o *InlineObject746) SetBasis(v AnyOfobject) {
	o.Basis = &v
}

// SetBasisExplicitNull (un)sets Basis to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Basis value is set to nil even if false is passed
func (o *InlineObject746) SetBasisExplicitNull(b bool) {
	o.Basis = nil
	o.isExplicitNullBasis = b
}
// GetCalcMethod returns the CalcMethod field if non-nil, zero value otherwise.
func (o *InlineObject746) GetCalcMethod() AnyOfobject {
	if o == nil || o.CalcMethod == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.CalcMethod
}

// GetCalcMethodOk returns a tuple with the CalcMethod field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject746) GetCalcMethodOk() (AnyOfobject, bool) {
	if o == nil || o.CalcMethod == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.CalcMethod, true
}

// HasCalcMethod returns a boolean if a field has been set.
func (o *InlineObject746) HasCalcMethod() bool {
	if o != nil && o.CalcMethod != nil {
		return true
	}

	return false
}

// SetCalcMethod gets a reference to the given AnyOfobject and assigns it to the CalcMethod field.
func (o *InlineObject746) SetCalcMethod(v AnyOfobject) {
	o.CalcMethod = &v
}

// SetCalcMethodExplicitNull (un)sets CalcMethod to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CalcMethod value is set to nil even if false is passed
func (o *InlineObject746) SetCalcMethodExplicitNull(b bool) {
	o.CalcMethod = nil
	o.isExplicitNullCalcMethod = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject746) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Issue == nil {
		if o.isExplicitNullIssue {
			toSerialize["issue"] = o.Issue
		}
	} else {
		toSerialize["issue"] = o.Issue
	}
	if o.FirstInterest == nil {
		if o.isExplicitNullFirstInterest {
			toSerialize["firstInterest"] = o.FirstInterest
		}
	} else {
		toSerialize["firstInterest"] = o.FirstInterest
	}
	if o.Settlement == nil {
		if o.isExplicitNullSettlement {
			toSerialize["settlement"] = o.Settlement
		}
	} else {
		toSerialize["settlement"] = o.Settlement
	}
	if o.Rate == nil {
		if o.isExplicitNullRate {
			toSerialize["rate"] = o.Rate
		}
	} else {
		toSerialize["rate"] = o.Rate
	}
	if o.Par == nil {
		if o.isExplicitNullPar {
			toSerialize["par"] = o.Par
		}
	} else {
		toSerialize["par"] = o.Par
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
	if o.CalcMethod == nil {
		if o.isExplicitNullCalcMethod {
			toSerialize["calcMethod"] = o.CalcMethod
		}
	} else {
		toSerialize["calcMethod"] = o.CalcMethod
	}
	return json.Marshal(toSerialize)
}



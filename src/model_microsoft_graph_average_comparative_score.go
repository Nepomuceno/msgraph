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
// MicrosoftGraphAverageComparativeScore struct for MicrosoftGraphAverageComparativeScore
type MicrosoftGraphAverageComparativeScore struct {
	AverageScore *AnyOfnumberstringstring `json:"averageScore,omitempty"`
	isExplicitNullAverageScore bool `json:"-"`
	Basis *string `json:"basis,omitempty"`
	isExplicitNullBasis bool `json:"-"`
}

// GetAverageScore returns the AverageScore field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAverageComparativeScore) GetAverageScore() AnyOfnumberstringstring {
	if o == nil || o.AverageScore == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.AverageScore
}

// GetAverageScoreOk returns a tuple with the AverageScore field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAverageComparativeScore) GetAverageScoreOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.AverageScore == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.AverageScore, true
}

// HasAverageScore returns a boolean if a field has been set.
func (o *MicrosoftGraphAverageComparativeScore) HasAverageScore() bool {
	if o != nil && o.AverageScore != nil {
		return true
	}

	return false
}

// SetAverageScore gets a reference to the given AnyOfnumberstringstring and assigns it to the AverageScore field.
func (o *MicrosoftGraphAverageComparativeScore) SetAverageScore(v AnyOfnumberstringstring) {
	o.AverageScore = &v
}

// SetAverageScoreExplicitNull (un)sets AverageScore to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AverageScore value is set to nil even if false is passed
func (o *MicrosoftGraphAverageComparativeScore) SetAverageScoreExplicitNull(b bool) {
	o.AverageScore = nil
	o.isExplicitNullAverageScore = b
}
// GetBasis returns the Basis field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAverageComparativeScore) GetBasis() string {
	if o == nil || o.Basis == nil {
		var ret string
		return ret
	}
	return *o.Basis
}

// GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAverageComparativeScore) GetBasisOk() (string, bool) {
	if o == nil || o.Basis == nil {
		var ret string
		return ret, false
	}
	return *o.Basis, true
}

// HasBasis returns a boolean if a field has been set.
func (o *MicrosoftGraphAverageComparativeScore) HasBasis() bool {
	if o != nil && o.Basis != nil {
		return true
	}

	return false
}

// SetBasis gets a reference to the given string and assigns it to the Basis field.
func (o *MicrosoftGraphAverageComparativeScore) SetBasis(v string) {
	o.Basis = &v
}

// SetBasisExplicitNull (un)sets Basis to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Basis value is set to nil even if false is passed
func (o *MicrosoftGraphAverageComparativeScore) SetBasisExplicitNull(b bool) {
	o.Basis = nil
	o.isExplicitNullBasis = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphAverageComparativeScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AverageScore == nil {
		if o.isExplicitNullAverageScore {
			toSerialize["averageScore"] = o.AverageScore
		}
	} else {
		toSerialize["averageScore"] = o.AverageScore
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


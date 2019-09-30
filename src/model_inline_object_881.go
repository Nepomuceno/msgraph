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
// InlineObject881 struct for InlineObject881
type InlineObject881 struct {
	Number *AnyOfobject `json:"number,omitempty"`
	isExplicitNullNumber bool `json:"-"`
	Significance *AnyOfobject `json:"significance,omitempty"`
	isExplicitNullSignificance bool `json:"-"`
}

// GetNumber returns the Number field if non-nil, zero value otherwise.
func (o *InlineObject881) GetNumber() AnyOfobject {
	if o == nil || o.Number == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject881) GetNumberOk() (AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject881) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject881) SetNumber(v AnyOfobject) {
	o.Number = &v
}

// SetNumberExplicitNull (un)sets Number to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Number value is set to nil even if false is passed
func (o *InlineObject881) SetNumberExplicitNull(b bool) {
	o.Number = nil
	o.isExplicitNullNumber = b
}
// GetSignificance returns the Significance field if non-nil, zero value otherwise.
func (o *InlineObject881) GetSignificance() AnyOfobject {
	if o == nil || o.Significance == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Significance
}

// GetSignificanceOk returns a tuple with the Significance field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject881) GetSignificanceOk() (AnyOfobject, bool) {
	if o == nil || o.Significance == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Significance, true
}

// HasSignificance returns a boolean if a field has been set.
func (o *InlineObject881) HasSignificance() bool {
	if o != nil && o.Significance != nil {
		return true
	}

	return false
}

// SetSignificance gets a reference to the given AnyOfobject and assigns it to the Significance field.
func (o *InlineObject881) SetSignificance(v AnyOfobject) {
	o.Significance = &v
}

// SetSignificanceExplicitNull (un)sets Significance to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Significance value is set to nil even if false is passed
func (o *InlineObject881) SetSignificanceExplicitNull(b bool) {
	o.Significance = nil
	o.isExplicitNullSignificance = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject881) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number == nil {
		if o.isExplicitNullNumber {
			toSerialize["number"] = o.Number
		}
	} else {
		toSerialize["number"] = o.Number
	}
	if o.Significance == nil {
		if o.isExplicitNullSignificance {
			toSerialize["significance"] = o.Significance
		}
	} else {
		toSerialize["significance"] = o.Significance
	}
	return json.Marshal(toSerialize)
}



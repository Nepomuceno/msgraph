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
// InlineObject787 struct for InlineObject787
type InlineObject787 struct {
	Number *AnyOfobject `json:"number,omitempty"`
	isExplicitNullNumber bool `json:"-"`
	Significance *AnyOfobject `json:"significance,omitempty"`
	isExplicitNullSignificance bool `json:"-"`
	Mode *AnyOfobject `json:"mode,omitempty"`
	isExplicitNullMode bool `json:"-"`
}

// GetNumber returns the Number field if non-nil, zero value otherwise.
func (o *InlineObject787) GetNumber() AnyOfobject {
	if o == nil || o.Number == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject787) GetNumberOk() (AnyOfobject, bool) {
	if o == nil || o.Number == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *InlineObject787) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given AnyOfobject and assigns it to the Number field.
func (o *InlineObject787) SetNumber(v AnyOfobject) {
	o.Number = &v
}

// SetNumberExplicitNull (un)sets Number to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Number value is set to nil even if false is passed
func (o *InlineObject787) SetNumberExplicitNull(b bool) {
	o.Number = nil
	o.isExplicitNullNumber = b
}
// GetSignificance returns the Significance field if non-nil, zero value otherwise.
func (o *InlineObject787) GetSignificance() AnyOfobject {
	if o == nil || o.Significance == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Significance
}

// GetSignificanceOk returns a tuple with the Significance field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject787) GetSignificanceOk() (AnyOfobject, bool) {
	if o == nil || o.Significance == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Significance, true
}

// HasSignificance returns a boolean if a field has been set.
func (o *InlineObject787) HasSignificance() bool {
	if o != nil && o.Significance != nil {
		return true
	}

	return false
}

// SetSignificance gets a reference to the given AnyOfobject and assigns it to the Significance field.
func (o *InlineObject787) SetSignificance(v AnyOfobject) {
	o.Significance = &v
}

// SetSignificanceExplicitNull (un)sets Significance to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Significance value is set to nil even if false is passed
func (o *InlineObject787) SetSignificanceExplicitNull(b bool) {
	o.Significance = nil
	o.isExplicitNullSignificance = b
}
// GetMode returns the Mode field if non-nil, zero value otherwise.
func (o *InlineObject787) GetMode() AnyOfobject {
	if o == nil || o.Mode == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject787) GetModeOk() (AnyOfobject, bool) {
	if o == nil || o.Mode == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *InlineObject787) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given AnyOfobject and assigns it to the Mode field.
func (o *InlineObject787) SetMode(v AnyOfobject) {
	o.Mode = &v
}

// SetModeExplicitNull (un)sets Mode to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Mode value is set to nil even if false is passed
func (o *InlineObject787) SetModeExplicitNull(b bool) {
	o.Mode = nil
	o.isExplicitNullMode = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject787) MarshalJSON() ([]byte, error) {
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
	if o.Mode == nil {
		if o.isExplicitNullMode {
			toSerialize["mode"] = o.Mode
		}
	} else {
		toSerialize["mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}



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
// InlineObject1092 struct for InlineObject1092
type InlineObject1092 struct {
	X *AnyOfobject `json:"x,omitempty"`
	isExplicitNullX bool `json:"-"`
	Alpha *AnyOfobject `json:"alpha,omitempty"`
	isExplicitNullAlpha bool `json:"-"`
	Beta *AnyOfobject `json:"beta,omitempty"`
	isExplicitNullBeta bool `json:"-"`
	Cumulative *AnyOfobject `json:"cumulative,omitempty"`
	isExplicitNullCumulative bool `json:"-"`
}

// GetX returns the X field if non-nil, zero value otherwise.
func (o *InlineObject1092) GetX() AnyOfobject {
	if o == nil || o.X == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1092) GetXOk() (AnyOfobject, bool) {
	if o == nil || o.X == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1092) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1092) SetX(v AnyOfobject) {
	o.X = &v
}

// SetXExplicitNull (un)sets X to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The X value is set to nil even if false is passed
func (o *InlineObject1092) SetXExplicitNull(b bool) {
	o.X = nil
	o.isExplicitNullX = b
}
// GetAlpha returns the Alpha field if non-nil, zero value otherwise.
func (o *InlineObject1092) GetAlpha() AnyOfobject {
	if o == nil || o.Alpha == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Alpha
}

// GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1092) GetAlphaOk() (AnyOfobject, bool) {
	if o == nil || o.Alpha == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Alpha, true
}

// HasAlpha returns a boolean if a field has been set.
func (o *InlineObject1092) HasAlpha() bool {
	if o != nil && o.Alpha != nil {
		return true
	}

	return false
}

// SetAlpha gets a reference to the given AnyOfobject and assigns it to the Alpha field.
func (o *InlineObject1092) SetAlpha(v AnyOfobject) {
	o.Alpha = &v
}

// SetAlphaExplicitNull (un)sets Alpha to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Alpha value is set to nil even if false is passed
func (o *InlineObject1092) SetAlphaExplicitNull(b bool) {
	o.Alpha = nil
	o.isExplicitNullAlpha = b
}
// GetBeta returns the Beta field if non-nil, zero value otherwise.
func (o *InlineObject1092) GetBeta() AnyOfobject {
	if o == nil || o.Beta == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Beta
}

// GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1092) GetBetaOk() (AnyOfobject, bool) {
	if o == nil || o.Beta == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Beta, true
}

// HasBeta returns a boolean if a field has been set.
func (o *InlineObject1092) HasBeta() bool {
	if o != nil && o.Beta != nil {
		return true
	}

	return false
}

// SetBeta gets a reference to the given AnyOfobject and assigns it to the Beta field.
func (o *InlineObject1092) SetBeta(v AnyOfobject) {
	o.Beta = &v
}

// SetBetaExplicitNull (un)sets Beta to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Beta value is set to nil even if false is passed
func (o *InlineObject1092) SetBetaExplicitNull(b bool) {
	o.Beta = nil
	o.isExplicitNullBeta = b
}
// GetCumulative returns the Cumulative field if non-nil, zero value otherwise.
func (o *InlineObject1092) GetCumulative() AnyOfobject {
	if o == nil || o.Cumulative == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Cumulative
}

// GetCumulativeOk returns a tuple with the Cumulative field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1092) GetCumulativeOk() (AnyOfobject, bool) {
	if o == nil || o.Cumulative == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Cumulative, true
}

// HasCumulative returns a boolean if a field has been set.
func (o *InlineObject1092) HasCumulative() bool {
	if o != nil && o.Cumulative != nil {
		return true
	}

	return false
}

// SetCumulative gets a reference to the given AnyOfobject and assigns it to the Cumulative field.
func (o *InlineObject1092) SetCumulative(v AnyOfobject) {
	o.Cumulative = &v
}

// SetCumulativeExplicitNull (un)sets Cumulative to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Cumulative value is set to nil even if false is passed
func (o *InlineObject1092) SetCumulativeExplicitNull(b bool) {
	o.Cumulative = nil
	o.isExplicitNullCumulative = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1092) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X == nil {
		if o.isExplicitNullX {
			toSerialize["x"] = o.X
		}
	} else {
		toSerialize["x"] = o.X
	}
	if o.Alpha == nil {
		if o.isExplicitNullAlpha {
			toSerialize["alpha"] = o.Alpha
		}
	} else {
		toSerialize["alpha"] = o.Alpha
	}
	if o.Beta == nil {
		if o.isExplicitNullBeta {
			toSerialize["beta"] = o.Beta
		}
	} else {
		toSerialize["beta"] = o.Beta
	}
	if o.Cumulative == nil {
		if o.isExplicitNullCumulative {
			toSerialize["cumulative"] = o.Cumulative
		}
	} else {
		toSerialize["cumulative"] = o.Cumulative
	}
	return json.Marshal(toSerialize)
}



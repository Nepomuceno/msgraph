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
// InlineObject1005 struct for InlineObject1005
type InlineObject1005 struct {
	X *AnyOfobject `json:"x,omitempty"`
	isExplicitNullX bool `json:"-"`
	Mean *AnyOfobject `json:"mean,omitempty"`
	isExplicitNullMean bool `json:"-"`
	Cumulative *AnyOfobject `json:"cumulative,omitempty"`
	isExplicitNullCumulative bool `json:"-"`
}

// GetX returns the X field if non-nil, zero value otherwise.
func (o *InlineObject1005) GetX() AnyOfobject {
	if o == nil || o.X == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1005) GetXOk() (AnyOfobject, bool) {
	if o == nil || o.X == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject1005) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject1005) SetX(v AnyOfobject) {
	o.X = &v
}

// SetXExplicitNull (un)sets X to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The X value is set to nil even if false is passed
func (o *InlineObject1005) SetXExplicitNull(b bool) {
	o.X = nil
	o.isExplicitNullX = b
}
// GetMean returns the Mean field if non-nil, zero value otherwise.
func (o *InlineObject1005) GetMean() AnyOfobject {
	if o == nil || o.Mean == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Mean
}

// GetMeanOk returns a tuple with the Mean field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1005) GetMeanOk() (AnyOfobject, bool) {
	if o == nil || o.Mean == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *InlineObject1005) HasMean() bool {
	if o != nil && o.Mean != nil {
		return true
	}

	return false
}

// SetMean gets a reference to the given AnyOfobject and assigns it to the Mean field.
func (o *InlineObject1005) SetMean(v AnyOfobject) {
	o.Mean = &v
}

// SetMeanExplicitNull (un)sets Mean to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Mean value is set to nil even if false is passed
func (o *InlineObject1005) SetMeanExplicitNull(b bool) {
	o.Mean = nil
	o.isExplicitNullMean = b
}
// GetCumulative returns the Cumulative field if non-nil, zero value otherwise.
func (o *InlineObject1005) GetCumulative() AnyOfobject {
	if o == nil || o.Cumulative == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Cumulative
}

// GetCumulativeOk returns a tuple with the Cumulative field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1005) GetCumulativeOk() (AnyOfobject, bool) {
	if o == nil || o.Cumulative == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Cumulative, true
}

// HasCumulative returns a boolean if a field has been set.
func (o *InlineObject1005) HasCumulative() bool {
	if o != nil && o.Cumulative != nil {
		return true
	}

	return false
}

// SetCumulative gets a reference to the given AnyOfobject and assigns it to the Cumulative field.
func (o *InlineObject1005) SetCumulative(v AnyOfobject) {
	o.Cumulative = &v
}

// SetCumulativeExplicitNull (un)sets Cumulative to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Cumulative value is set to nil even if false is passed
func (o *InlineObject1005) SetCumulativeExplicitNull(b bool) {
	o.Cumulative = nil
	o.isExplicitNullCumulative = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1005) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X == nil {
		if o.isExplicitNullX {
			toSerialize["x"] = o.X
		}
	} else {
		toSerialize["x"] = o.X
	}
	if o.Mean == nil {
		if o.isExplicitNullMean {
			toSerialize["mean"] = o.Mean
		}
	} else {
		toSerialize["mean"] = o.Mean
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


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
// InlineObject999 struct for InlineObject999
type InlineObject999 struct {
	Array *AnyOfobject `json:"array,omitempty"`
	isExplicitNullArray bool `json:"-"`
	X *AnyOfobject `json:"x,omitempty"`
	isExplicitNullX bool `json:"-"`
	Significance *AnyOfobject `json:"significance,omitempty"`
	isExplicitNullSignificance bool `json:"-"`
}

// GetArray returns the Array field if non-nil, zero value otherwise.
func (o *InlineObject999) GetArray() AnyOfobject {
	if o == nil || o.Array == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject999) GetArrayOk() (AnyOfobject, bool) {
	if o == nil || o.Array == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *InlineObject999) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given AnyOfobject and assigns it to the Array field.
func (o *InlineObject999) SetArray(v AnyOfobject) {
	o.Array = &v
}

// SetArrayExplicitNull (un)sets Array to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Array value is set to nil even if false is passed
func (o *InlineObject999) SetArrayExplicitNull(b bool) {
	o.Array = nil
	o.isExplicitNullArray = b
}
// GetX returns the X field if non-nil, zero value otherwise.
func (o *InlineObject999) GetX() AnyOfobject {
	if o == nil || o.X == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject999) GetXOk() (AnyOfobject, bool) {
	if o == nil || o.X == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *InlineObject999) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given AnyOfobject and assigns it to the X field.
func (o *InlineObject999) SetX(v AnyOfobject) {
	o.X = &v
}

// SetXExplicitNull (un)sets X to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The X value is set to nil even if false is passed
func (o *InlineObject999) SetXExplicitNull(b bool) {
	o.X = nil
	o.isExplicitNullX = b
}
// GetSignificance returns the Significance field if non-nil, zero value otherwise.
func (o *InlineObject999) GetSignificance() AnyOfobject {
	if o == nil || o.Significance == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Significance
}

// GetSignificanceOk returns a tuple with the Significance field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject999) GetSignificanceOk() (AnyOfobject, bool) {
	if o == nil || o.Significance == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Significance, true
}

// HasSignificance returns a boolean if a field has been set.
func (o *InlineObject999) HasSignificance() bool {
	if o != nil && o.Significance != nil {
		return true
	}

	return false
}

// SetSignificance gets a reference to the given AnyOfobject and assigns it to the Significance field.
func (o *InlineObject999) SetSignificance(v AnyOfobject) {
	o.Significance = &v
}

// SetSignificanceExplicitNull (un)sets Significance to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Significance value is set to nil even if false is passed
func (o *InlineObject999) SetSignificanceExplicitNull(b bool) {
	o.Significance = nil
	o.isExplicitNullSignificance = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject999) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Array == nil {
		if o.isExplicitNullArray {
			toSerialize["array"] = o.Array
		}
	} else {
		toSerialize["array"] = o.Array
	}
	if o.X == nil {
		if o.isExplicitNullX {
			toSerialize["x"] = o.X
		}
	} else {
		toSerialize["x"] = o.X
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


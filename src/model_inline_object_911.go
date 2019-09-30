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
// InlineObject911 struct for InlineObject911
type InlineObject911 struct {
	Inumber1 *AnyOfobject `json:"inumber1,omitempty"`
	isExplicitNullInumber1 bool `json:"-"`
	Inumber2 *AnyOfobject `json:"inumber2,omitempty"`
	isExplicitNullInumber2 bool `json:"-"`
}

// GetInumber1 returns the Inumber1 field if non-nil, zero value otherwise.
func (o *InlineObject911) GetInumber1() AnyOfobject {
	if o == nil || o.Inumber1 == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Inumber1
}

// GetInumber1Ok returns a tuple with the Inumber1 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject911) GetInumber1Ok() (AnyOfobject, bool) {
	if o == nil || o.Inumber1 == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Inumber1, true
}

// HasInumber1 returns a boolean if a field has been set.
func (o *InlineObject911) HasInumber1() bool {
	if o != nil && o.Inumber1 != nil {
		return true
	}

	return false
}

// SetInumber1 gets a reference to the given AnyOfobject and assigns it to the Inumber1 field.
func (o *InlineObject911) SetInumber1(v AnyOfobject) {
	o.Inumber1 = &v
}

// SetInumber1ExplicitNull (un)sets Inumber1 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Inumber1 value is set to nil even if false is passed
func (o *InlineObject911) SetInumber1ExplicitNull(b bool) {
	o.Inumber1 = nil
	o.isExplicitNullInumber1 = b
}
// GetInumber2 returns the Inumber2 field if non-nil, zero value otherwise.
func (o *InlineObject911) GetInumber2() AnyOfobject {
	if o == nil || o.Inumber2 == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Inumber2
}

// GetInumber2Ok returns a tuple with the Inumber2 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject911) GetInumber2Ok() (AnyOfobject, bool) {
	if o == nil || o.Inumber2 == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Inumber2, true
}

// HasInumber2 returns a boolean if a field has been set.
func (o *InlineObject911) HasInumber2() bool {
	if o != nil && o.Inumber2 != nil {
		return true
	}

	return false
}

// SetInumber2 gets a reference to the given AnyOfobject and assigns it to the Inumber2 field.
func (o *InlineObject911) SetInumber2(v AnyOfobject) {
	o.Inumber2 = &v
}

// SetInumber2ExplicitNull (un)sets Inumber2 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Inumber2 value is set to nil even if false is passed
func (o *InlineObject911) SetInumber2ExplicitNull(b bool) {
	o.Inumber2 = nil
	o.isExplicitNullInumber2 = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject911) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Inumber1 == nil {
		if o.isExplicitNullInumber1 {
			toSerialize["inumber1"] = o.Inumber1
		}
	} else {
		toSerialize["inumber1"] = o.Inumber1
	}
	if o.Inumber2 == nil {
		if o.isExplicitNullInumber2 {
			toSerialize["inumber2"] = o.Inumber2
		}
	} else {
		toSerialize["inumber2"] = o.Inumber2
	}
	return json.Marshal(toSerialize)
}



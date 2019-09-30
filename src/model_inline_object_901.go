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
// InlineObject901 struct for InlineObject901
type InlineObject901 struct {
	LogicalTest *AnyOfobject `json:"logicalTest,omitempty"`
	isExplicitNullLogicalTest bool `json:"-"`
	ValueIfTrue *AnyOfobject `json:"valueIfTrue,omitempty"`
	isExplicitNullValueIfTrue bool `json:"-"`
	ValueIfFalse *AnyOfobject `json:"valueIfFalse,omitempty"`
	isExplicitNullValueIfFalse bool `json:"-"`
}

// GetLogicalTest returns the LogicalTest field if non-nil, zero value otherwise.
func (o *InlineObject901) GetLogicalTest() AnyOfobject {
	if o == nil || o.LogicalTest == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.LogicalTest
}

// GetLogicalTestOk returns a tuple with the LogicalTest field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject901) GetLogicalTestOk() (AnyOfobject, bool) {
	if o == nil || o.LogicalTest == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.LogicalTest, true
}

// HasLogicalTest returns a boolean if a field has been set.
func (o *InlineObject901) HasLogicalTest() bool {
	if o != nil && o.LogicalTest != nil {
		return true
	}

	return false
}

// SetLogicalTest gets a reference to the given AnyOfobject and assigns it to the LogicalTest field.
func (o *InlineObject901) SetLogicalTest(v AnyOfobject) {
	o.LogicalTest = &v
}

// SetLogicalTestExplicitNull (un)sets LogicalTest to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LogicalTest value is set to nil even if false is passed
func (o *InlineObject901) SetLogicalTestExplicitNull(b bool) {
	o.LogicalTest = nil
	o.isExplicitNullLogicalTest = b
}
// GetValueIfTrue returns the ValueIfTrue field if non-nil, zero value otherwise.
func (o *InlineObject901) GetValueIfTrue() AnyOfobject {
	if o == nil || o.ValueIfTrue == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.ValueIfTrue
}

// GetValueIfTrueOk returns a tuple with the ValueIfTrue field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject901) GetValueIfTrueOk() (AnyOfobject, bool) {
	if o == nil || o.ValueIfTrue == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.ValueIfTrue, true
}

// HasValueIfTrue returns a boolean if a field has been set.
func (o *InlineObject901) HasValueIfTrue() bool {
	if o != nil && o.ValueIfTrue != nil {
		return true
	}

	return false
}

// SetValueIfTrue gets a reference to the given AnyOfobject and assigns it to the ValueIfTrue field.
func (o *InlineObject901) SetValueIfTrue(v AnyOfobject) {
	o.ValueIfTrue = &v
}

// SetValueIfTrueExplicitNull (un)sets ValueIfTrue to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ValueIfTrue value is set to nil even if false is passed
func (o *InlineObject901) SetValueIfTrueExplicitNull(b bool) {
	o.ValueIfTrue = nil
	o.isExplicitNullValueIfTrue = b
}
// GetValueIfFalse returns the ValueIfFalse field if non-nil, zero value otherwise.
func (o *InlineObject901) GetValueIfFalse() AnyOfobject {
	if o == nil || o.ValueIfFalse == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.ValueIfFalse
}

// GetValueIfFalseOk returns a tuple with the ValueIfFalse field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject901) GetValueIfFalseOk() (AnyOfobject, bool) {
	if o == nil || o.ValueIfFalse == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.ValueIfFalse, true
}

// HasValueIfFalse returns a boolean if a field has been set.
func (o *InlineObject901) HasValueIfFalse() bool {
	if o != nil && o.ValueIfFalse != nil {
		return true
	}

	return false
}

// SetValueIfFalse gets a reference to the given AnyOfobject and assigns it to the ValueIfFalse field.
func (o *InlineObject901) SetValueIfFalse(v AnyOfobject) {
	o.ValueIfFalse = &v
}

// SetValueIfFalseExplicitNull (un)sets ValueIfFalse to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ValueIfFalse value is set to nil even if false is passed
func (o *InlineObject901) SetValueIfFalseExplicitNull(b bool) {
	o.ValueIfFalse = nil
	o.isExplicitNullValueIfFalse = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject901) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogicalTest == nil {
		if o.isExplicitNullLogicalTest {
			toSerialize["logicalTest"] = o.LogicalTest
		}
	} else {
		toSerialize["logicalTest"] = o.LogicalTest
	}
	if o.ValueIfTrue == nil {
		if o.isExplicitNullValueIfTrue {
			toSerialize["valueIfTrue"] = o.ValueIfTrue
		}
	} else {
		toSerialize["valueIfTrue"] = o.ValueIfTrue
	}
	if o.ValueIfFalse == nil {
		if o.isExplicitNullValueIfFalse {
			toSerialize["valueIfFalse"] = o.ValueIfFalse
		}
	} else {
		toSerialize["valueIfFalse"] = o.ValueIfFalse
	}
	return json.Marshal(toSerialize)
}



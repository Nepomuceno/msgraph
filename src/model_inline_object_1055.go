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
// InlineObject1055 struct for InlineObject1055
type InlineObject1055 struct {
	FunctionNum *AnyOfobject `json:"functionNum,omitempty"`
	isExplicitNullFunctionNum bool `json:"-"`
	Values *AnyOfobject `json:"values,omitempty"`
	isExplicitNullValues bool `json:"-"`
}

// GetFunctionNum returns the FunctionNum field if non-nil, zero value otherwise.
func (o *InlineObject1055) GetFunctionNum() AnyOfobject {
	if o == nil || o.FunctionNum == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.FunctionNum
}

// GetFunctionNumOk returns a tuple with the FunctionNum field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1055) GetFunctionNumOk() (AnyOfobject, bool) {
	if o == nil || o.FunctionNum == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.FunctionNum, true
}

// HasFunctionNum returns a boolean if a field has been set.
func (o *InlineObject1055) HasFunctionNum() bool {
	if o != nil && o.FunctionNum != nil {
		return true
	}

	return false
}

// SetFunctionNum gets a reference to the given AnyOfobject and assigns it to the FunctionNum field.
func (o *InlineObject1055) SetFunctionNum(v AnyOfobject) {
	o.FunctionNum = &v
}

// SetFunctionNumExplicitNull (un)sets FunctionNum to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The FunctionNum value is set to nil even if false is passed
func (o *InlineObject1055) SetFunctionNumExplicitNull(b bool) {
	o.FunctionNum = nil
	o.isExplicitNullFunctionNum = b
}
// GetValues returns the Values field if non-nil, zero value otherwise.
func (o *InlineObject1055) GetValues() AnyOfobject {
	if o == nil || o.Values == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1055) GetValuesOk() (AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1055) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1055) SetValues(v AnyOfobject) {
	o.Values = &v
}

// SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Values value is set to nil even if false is passed
func (o *InlineObject1055) SetValuesExplicitNull(b bool) {
	o.Values = nil
	o.isExplicitNullValues = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1055) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FunctionNum == nil {
		if o.isExplicitNullFunctionNum {
			toSerialize["functionNum"] = o.FunctionNum
		}
	} else {
		toSerialize["functionNum"] = o.FunctionNum
	}
	if o.Values == nil {
		if o.isExplicitNullValues {
			toSerialize["values"] = o.Values
		}
	} else {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}


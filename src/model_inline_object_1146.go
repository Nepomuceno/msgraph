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
// InlineObject1146 struct for InlineObject1146
type InlineObject1146 struct {
	Index *int32 `json:"index,omitempty"`
	isExplicitNullIndex bool `json:"-"`
	Values *AnyOfobject `json:"values,omitempty"`
	isExplicitNullValues bool `json:"-"`
}

// GetIndex returns the Index field if non-nil, zero value otherwise.
func (o *InlineObject1146) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1146) GetIndexOk() (int32, bool) {
	if o == nil || o.Index == nil {
		var ret int32
		return ret, false
	}
	return *o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *InlineObject1146) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *InlineObject1146) SetIndex(v int32) {
	o.Index = &v
}

// SetIndexExplicitNull (un)sets Index to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Index value is set to nil even if false is passed
func (o *InlineObject1146) SetIndexExplicitNull(b bool) {
	o.Index = nil
	o.isExplicitNullIndex = b
}
// GetValues returns the Values field if non-nil, zero value otherwise.
func (o *InlineObject1146) GetValues() AnyOfobject {
	if o == nil || o.Values == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1146) GetValuesOk() (AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1146) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1146) SetValues(v AnyOfobject) {
	o.Values = &v
}

// SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Values value is set to nil even if false is passed
func (o *InlineObject1146) SetValuesExplicitNull(b bool) {
	o.Values = nil
	o.isExplicitNullValues = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1146) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index == nil {
		if o.isExplicitNullIndex {
			toSerialize["index"] = o.Index
		}
	} else {
		toSerialize["index"] = o.Index
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



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
// InlineObject1056 struct for InlineObject1056
type InlineObject1056 struct {
	Values *AnyOfobject `json:"values,omitempty"`
	isExplicitNullValues bool `json:"-"`
}

// GetValues returns the Values field if non-nil, zero value otherwise.
func (o *InlineObject1056) GetValues() AnyOfobject {
	if o == nil || o.Values == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1056) GetValuesOk() (AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject1056) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *InlineObject1056) SetValues(v AnyOfobject) {
	o.Values = &v
}

// SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Values value is set to nil even if false is passed
func (o *InlineObject1056) SetValuesExplicitNull(b bool) {
	o.Values = nil
	o.isExplicitNullValues = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1056) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values == nil {
		if o.isExplicitNullValues {
			toSerialize["values"] = o.Values
		}
	} else {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}



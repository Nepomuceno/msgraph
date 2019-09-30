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
// InlineObject811 struct for InlineObject811
type InlineObject811 struct {
	Range *AnyOfobject `json:"range,omitempty"`
	isExplicitNullRange bool `json:"-"`
}

// GetRange returns the Range field if non-nil, zero value otherwise.
func (o *InlineObject811) GetRange() AnyOfobject {
	if o == nil || o.Range == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject811) GetRangeOk() (AnyOfobject, bool) {
	if o == nil || o.Range == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *InlineObject811) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given AnyOfobject and assigns it to the Range field.
func (o *InlineObject811) SetRange(v AnyOfobject) {
	o.Range = &v
}

// SetRangeExplicitNull (un)sets Range to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Range value is set to nil even if false is passed
func (o *InlineObject811) SetRangeExplicitNull(b bool) {
	o.Range = nil
	o.isExplicitNullRange = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject811) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Range == nil {
		if o.isExplicitNullRange {
			toSerialize["range"] = o.Range
		}
	} else {
		toSerialize["range"] = o.Range
	}
	return json.Marshal(toSerialize)
}



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
// InlineObject1109 struct for InlineObject1109
type InlineObject1109 struct {
	StartCell *AnyOfobject `json:"startCell,omitempty"`
	isExplicitNullStartCell bool `json:"-"`
	EndCell *AnyOfobject `json:"endCell,omitempty"`
	isExplicitNullEndCell bool `json:"-"`
}

// GetStartCell returns the StartCell field if non-nil, zero value otherwise.
func (o *InlineObject1109) GetStartCell() AnyOfobject {
	if o == nil || o.StartCell == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.StartCell
}

// GetStartCellOk returns a tuple with the StartCell field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1109) GetStartCellOk() (AnyOfobject, bool) {
	if o == nil || o.StartCell == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.StartCell, true
}

// HasStartCell returns a boolean if a field has been set.
func (o *InlineObject1109) HasStartCell() bool {
	if o != nil && o.StartCell != nil {
		return true
	}

	return false
}

// SetStartCell gets a reference to the given AnyOfobject and assigns it to the StartCell field.
func (o *InlineObject1109) SetStartCell(v AnyOfobject) {
	o.StartCell = &v
}

// SetStartCellExplicitNull (un)sets StartCell to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The StartCell value is set to nil even if false is passed
func (o *InlineObject1109) SetStartCellExplicitNull(b bool) {
	o.StartCell = nil
	o.isExplicitNullStartCell = b
}
// GetEndCell returns the EndCell field if non-nil, zero value otherwise.
func (o *InlineObject1109) GetEndCell() AnyOfobject {
	if o == nil || o.EndCell == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.EndCell
}

// GetEndCellOk returns a tuple with the EndCell field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1109) GetEndCellOk() (AnyOfobject, bool) {
	if o == nil || o.EndCell == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.EndCell, true
}

// HasEndCell returns a boolean if a field has been set.
func (o *InlineObject1109) HasEndCell() bool {
	if o != nil && o.EndCell != nil {
		return true
	}

	return false
}

// SetEndCell gets a reference to the given AnyOfobject and assigns it to the EndCell field.
func (o *InlineObject1109) SetEndCell(v AnyOfobject) {
	o.EndCell = &v
}

// SetEndCellExplicitNull (un)sets EndCell to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The EndCell value is set to nil even if false is passed
func (o *InlineObject1109) SetEndCellExplicitNull(b bool) {
	o.EndCell = nil
	o.isExplicitNullEndCell = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1109) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartCell == nil {
		if o.isExplicitNullStartCell {
			toSerialize["startCell"] = o.StartCell
		}
	} else {
		toSerialize["startCell"] = o.StartCell
	}
	if o.EndCell == nil {
		if o.isExplicitNullEndCell {
			toSerialize["endCell"] = o.EndCell
		}
	} else {
		toSerialize["endCell"] = o.EndCell
	}
	return json.Marshal(toSerialize)
}



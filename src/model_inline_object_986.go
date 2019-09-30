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
// InlineObject986 struct for InlineObject986
type InlineObject986 struct {
	Text *AnyOfobject `json:"text,omitempty"`
	isExplicitNullText bool `json:"-"`
	DecimalSeparator *AnyOfobject `json:"decimalSeparator,omitempty"`
	isExplicitNullDecimalSeparator bool `json:"-"`
	GroupSeparator *AnyOfobject `json:"groupSeparator,omitempty"`
	isExplicitNullGroupSeparator bool `json:"-"`
}

// GetText returns the Text field if non-nil, zero value otherwise.
func (o *InlineObject986) GetText() AnyOfobject {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject986) GetTextOk() (AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *InlineObject986) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *InlineObject986) SetText(v AnyOfobject) {
	o.Text = &v
}

// SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Text value is set to nil even if false is passed
func (o *InlineObject986) SetTextExplicitNull(b bool) {
	o.Text = nil
	o.isExplicitNullText = b
}
// GetDecimalSeparator returns the DecimalSeparator field if non-nil, zero value otherwise.
func (o *InlineObject986) GetDecimalSeparator() AnyOfobject {
	if o == nil || o.DecimalSeparator == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.DecimalSeparator
}

// GetDecimalSeparatorOk returns a tuple with the DecimalSeparator field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject986) GetDecimalSeparatorOk() (AnyOfobject, bool) {
	if o == nil || o.DecimalSeparator == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.DecimalSeparator, true
}

// HasDecimalSeparator returns a boolean if a field has been set.
func (o *InlineObject986) HasDecimalSeparator() bool {
	if o != nil && o.DecimalSeparator != nil {
		return true
	}

	return false
}

// SetDecimalSeparator gets a reference to the given AnyOfobject and assigns it to the DecimalSeparator field.
func (o *InlineObject986) SetDecimalSeparator(v AnyOfobject) {
	o.DecimalSeparator = &v
}

// SetDecimalSeparatorExplicitNull (un)sets DecimalSeparator to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DecimalSeparator value is set to nil even if false is passed
func (o *InlineObject986) SetDecimalSeparatorExplicitNull(b bool) {
	o.DecimalSeparator = nil
	o.isExplicitNullDecimalSeparator = b
}
// GetGroupSeparator returns the GroupSeparator field if non-nil, zero value otherwise.
func (o *InlineObject986) GetGroupSeparator() AnyOfobject {
	if o == nil || o.GroupSeparator == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.GroupSeparator
}

// GetGroupSeparatorOk returns a tuple with the GroupSeparator field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject986) GetGroupSeparatorOk() (AnyOfobject, bool) {
	if o == nil || o.GroupSeparator == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.GroupSeparator, true
}

// HasGroupSeparator returns a boolean if a field has been set.
func (o *InlineObject986) HasGroupSeparator() bool {
	if o != nil && o.GroupSeparator != nil {
		return true
	}

	return false
}

// SetGroupSeparator gets a reference to the given AnyOfobject and assigns it to the GroupSeparator field.
func (o *InlineObject986) SetGroupSeparator(v AnyOfobject) {
	o.GroupSeparator = &v
}

// SetGroupSeparatorExplicitNull (un)sets GroupSeparator to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The GroupSeparator value is set to nil even if false is passed
func (o *InlineObject986) SetGroupSeparatorExplicitNull(b bool) {
	o.GroupSeparator = nil
	o.isExplicitNullGroupSeparator = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject986) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text == nil {
		if o.isExplicitNullText {
			toSerialize["text"] = o.Text
		}
	} else {
		toSerialize["text"] = o.Text
	}
	if o.DecimalSeparator == nil {
		if o.isExplicitNullDecimalSeparator {
			toSerialize["decimalSeparator"] = o.DecimalSeparator
		}
	} else {
		toSerialize["decimalSeparator"] = o.DecimalSeparator
	}
	if o.GroupSeparator == nil {
		if o.isExplicitNullGroupSeparator {
			toSerialize["groupSeparator"] = o.GroupSeparator
		}
	} else {
		toSerialize["groupSeparator"] = o.GroupSeparator
	}
	return json.Marshal(toSerialize)
}



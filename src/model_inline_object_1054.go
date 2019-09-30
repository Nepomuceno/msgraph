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
// InlineObject1054 struct for InlineObject1054
type InlineObject1054 struct {
	Text *AnyOfobject `json:"text,omitempty"`
	isExplicitNullText bool `json:"-"`
	OldText *AnyOfobject `json:"oldText,omitempty"`
	isExplicitNullOldText bool `json:"-"`
	NewText *AnyOfobject `json:"newText,omitempty"`
	isExplicitNullNewText bool `json:"-"`
	InstanceNum *AnyOfobject `json:"instanceNum,omitempty"`
	isExplicitNullInstanceNum bool `json:"-"`
}

// GetText returns the Text field if non-nil, zero value otherwise.
func (o *InlineObject1054) GetText() AnyOfobject {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1054) GetTextOk() (AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *InlineObject1054) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *InlineObject1054) SetText(v AnyOfobject) {
	o.Text = &v
}

// SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Text value is set to nil even if false is passed
func (o *InlineObject1054) SetTextExplicitNull(b bool) {
	o.Text = nil
	o.isExplicitNullText = b
}
// GetOldText returns the OldText field if non-nil, zero value otherwise.
func (o *InlineObject1054) GetOldText() AnyOfobject {
	if o == nil || o.OldText == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.OldText
}

// GetOldTextOk returns a tuple with the OldText field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1054) GetOldTextOk() (AnyOfobject, bool) {
	if o == nil || o.OldText == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.OldText, true
}

// HasOldText returns a boolean if a field has been set.
func (o *InlineObject1054) HasOldText() bool {
	if o != nil && o.OldText != nil {
		return true
	}

	return false
}

// SetOldText gets a reference to the given AnyOfobject and assigns it to the OldText field.
func (o *InlineObject1054) SetOldText(v AnyOfobject) {
	o.OldText = &v
}

// SetOldTextExplicitNull (un)sets OldText to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OldText value is set to nil even if false is passed
func (o *InlineObject1054) SetOldTextExplicitNull(b bool) {
	o.OldText = nil
	o.isExplicitNullOldText = b
}
// GetNewText returns the NewText field if non-nil, zero value otherwise.
func (o *InlineObject1054) GetNewText() AnyOfobject {
	if o == nil || o.NewText == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.NewText
}

// GetNewTextOk returns a tuple with the NewText field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1054) GetNewTextOk() (AnyOfobject, bool) {
	if o == nil || o.NewText == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.NewText, true
}

// HasNewText returns a boolean if a field has been set.
func (o *InlineObject1054) HasNewText() bool {
	if o != nil && o.NewText != nil {
		return true
	}

	return false
}

// SetNewText gets a reference to the given AnyOfobject and assigns it to the NewText field.
func (o *InlineObject1054) SetNewText(v AnyOfobject) {
	o.NewText = &v
}

// SetNewTextExplicitNull (un)sets NewText to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The NewText value is set to nil even if false is passed
func (o *InlineObject1054) SetNewTextExplicitNull(b bool) {
	o.NewText = nil
	o.isExplicitNullNewText = b
}
// GetInstanceNum returns the InstanceNum field if non-nil, zero value otherwise.
func (o *InlineObject1054) GetInstanceNum() AnyOfobject {
	if o == nil || o.InstanceNum == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.InstanceNum
}

// GetInstanceNumOk returns a tuple with the InstanceNum field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1054) GetInstanceNumOk() (AnyOfobject, bool) {
	if o == nil || o.InstanceNum == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.InstanceNum, true
}

// HasInstanceNum returns a boolean if a field has been set.
func (o *InlineObject1054) HasInstanceNum() bool {
	if o != nil && o.InstanceNum != nil {
		return true
	}

	return false
}

// SetInstanceNum gets a reference to the given AnyOfobject and assigns it to the InstanceNum field.
func (o *InlineObject1054) SetInstanceNum(v AnyOfobject) {
	o.InstanceNum = &v
}

// SetInstanceNumExplicitNull (un)sets InstanceNum to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The InstanceNum value is set to nil even if false is passed
func (o *InlineObject1054) SetInstanceNumExplicitNull(b bool) {
	o.InstanceNum = nil
	o.isExplicitNullInstanceNum = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1054) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text == nil {
		if o.isExplicitNullText {
			toSerialize["text"] = o.Text
		}
	} else {
		toSerialize["text"] = o.Text
	}
	if o.OldText == nil {
		if o.isExplicitNullOldText {
			toSerialize["oldText"] = o.OldText
		}
	} else {
		toSerialize["oldText"] = o.OldText
	}
	if o.NewText == nil {
		if o.isExplicitNullNewText {
			toSerialize["newText"] = o.NewText
		}
	} else {
		toSerialize["newText"] = o.NewText
	}
	if o.InstanceNum == nil {
		if o.isExplicitNullInstanceNum {
			toSerialize["instanceNum"] = o.InstanceNum
		}
	} else {
		toSerialize["instanceNum"] = o.InstanceNum
	}
	return json.Marshal(toSerialize)
}



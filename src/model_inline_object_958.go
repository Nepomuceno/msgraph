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
// InlineObject958 struct for InlineObject958
type InlineObject958 struct {
	Text *AnyOfobject `json:"text,omitempty"`
	isExplicitNullText bool `json:"-"`
}

// GetText returns the Text field if non-nil, zero value otherwise.
func (o *InlineObject958) GetText() AnyOfobject {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject958) GetTextOk() (AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *InlineObject958) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *InlineObject958) SetText(v AnyOfobject) {
	o.Text = &v
}

// SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Text value is set to nil even if false is passed
func (o *InlineObject958) SetTextExplicitNull(b bool) {
	o.Text = nil
	o.isExplicitNullText = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject958) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text == nil {
		if o.isExplicitNullText {
			toSerialize["text"] = o.Text
		}
	} else {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}



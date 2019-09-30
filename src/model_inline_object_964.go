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
// InlineObject964 struct for InlineObject964
type InlineObject964 struct {
	Text *AnyOfobject `json:"text,omitempty"`
	isExplicitNullText bool `json:"-"`
	StartNum *AnyOfobject `json:"startNum,omitempty"`
	isExplicitNullStartNum bool `json:"-"`
	NumChars *AnyOfobject `json:"numChars,omitempty"`
	isExplicitNullNumChars bool `json:"-"`
}

// GetText returns the Text field if non-nil, zero value otherwise.
func (o *InlineObject964) GetText() AnyOfobject {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject964) GetTextOk() (AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *InlineObject964) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *InlineObject964) SetText(v AnyOfobject) {
	o.Text = &v
}

// SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Text value is set to nil even if false is passed
func (o *InlineObject964) SetTextExplicitNull(b bool) {
	o.Text = nil
	o.isExplicitNullText = b
}
// GetStartNum returns the StartNum field if non-nil, zero value otherwise.
func (o *InlineObject964) GetStartNum() AnyOfobject {
	if o == nil || o.StartNum == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.StartNum
}

// GetStartNumOk returns a tuple with the StartNum field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject964) GetStartNumOk() (AnyOfobject, bool) {
	if o == nil || o.StartNum == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.StartNum, true
}

// HasStartNum returns a boolean if a field has been set.
func (o *InlineObject964) HasStartNum() bool {
	if o != nil && o.StartNum != nil {
		return true
	}

	return false
}

// SetStartNum gets a reference to the given AnyOfobject and assigns it to the StartNum field.
func (o *InlineObject964) SetStartNum(v AnyOfobject) {
	o.StartNum = &v
}

// SetStartNumExplicitNull (un)sets StartNum to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The StartNum value is set to nil even if false is passed
func (o *InlineObject964) SetStartNumExplicitNull(b bool) {
	o.StartNum = nil
	o.isExplicitNullStartNum = b
}
// GetNumChars returns the NumChars field if non-nil, zero value otherwise.
func (o *InlineObject964) GetNumChars() AnyOfobject {
	if o == nil || o.NumChars == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.NumChars
}

// GetNumCharsOk returns a tuple with the NumChars field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject964) GetNumCharsOk() (AnyOfobject, bool) {
	if o == nil || o.NumChars == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.NumChars, true
}

// HasNumChars returns a boolean if a field has been set.
func (o *InlineObject964) HasNumChars() bool {
	if o != nil && o.NumChars != nil {
		return true
	}

	return false
}

// SetNumChars gets a reference to the given AnyOfobject and assigns it to the NumChars field.
func (o *InlineObject964) SetNumChars(v AnyOfobject) {
	o.NumChars = &v
}

// SetNumCharsExplicitNull (un)sets NumChars to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The NumChars value is set to nil even if false is passed
func (o *InlineObject964) SetNumCharsExplicitNull(b bool) {
	o.NumChars = nil
	o.isExplicitNullNumChars = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject964) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text == nil {
		if o.isExplicitNullText {
			toSerialize["text"] = o.Text
		}
	} else {
		toSerialize["text"] = o.Text
	}
	if o.StartNum == nil {
		if o.isExplicitNullStartNum {
			toSerialize["startNum"] = o.StartNum
		}
	} else {
		toSerialize["startNum"] = o.StartNum
	}
	if o.NumChars == nil {
		if o.isExplicitNullNumChars {
			toSerialize["numChars"] = o.NumChars
		}
	} else {
		toSerialize["numChars"] = o.NumChars
	}
	return json.Marshal(toSerialize)
}


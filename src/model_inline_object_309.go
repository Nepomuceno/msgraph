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
// InlineObject309 struct for InlineObject309
type InlineObject309 struct {
	Message *AnyOfmicrosoftGraphMessage `json:"Message,omitempty"`
	isExplicitNullMessage bool `json:"-"`
	Comment *string `json:"Comment,omitempty"`
	isExplicitNullComment bool `json:"-"`
}

// GetMessage returns the Message field if non-nil, zero value otherwise.
func (o *InlineObject309) GetMessage() AnyOfmicrosoftGraphMessage {
	if o == nil || o.Message == nil {
		var ret AnyOfmicrosoftGraphMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject309) GetMessageOk() (AnyOfmicrosoftGraphMessage, bool) {
	if o == nil || o.Message == nil {
		var ret AnyOfmicrosoftGraphMessage
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObject309) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given AnyOfmicrosoftGraphMessage and assigns it to the Message field.
func (o *InlineObject309) SetMessage(v AnyOfmicrosoftGraphMessage) {
	o.Message = &v
}

// SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Message value is set to nil even if false is passed
func (o *InlineObject309) SetMessageExplicitNull(b bool) {
	o.Message = nil
	o.isExplicitNullMessage = b
}
// GetComment returns the Comment field if non-nil, zero value otherwise.
func (o *InlineObject309) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject309) GetCommentOk() (string, bool) {
	if o == nil || o.Comment == nil {
		var ret string
		return ret, false
	}
	return *o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject309) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineObject309) SetComment(v string) {
	o.Comment = &v
}

// SetCommentExplicitNull (un)sets Comment to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Comment value is set to nil even if false is passed
func (o *InlineObject309) SetCommentExplicitNull(b bool) {
	o.Comment = nil
	o.isExplicitNullComment = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject309) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message == nil {
		if o.isExplicitNullMessage {
			toSerialize["Message"] = o.Message
		}
	} else {
		toSerialize["Message"] = o.Message
	}
	if o.Comment == nil {
		if o.isExplicitNullComment {
			toSerialize["Comment"] = o.Comment
		}
	} else {
		toSerialize["Comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}



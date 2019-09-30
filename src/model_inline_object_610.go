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
// InlineObject610 struct for InlineObject610
type InlineObject610 struct {
	ToRecipients *[]AnyOfmicrosoftGraphRecipient `json:"ToRecipients,omitempty"`

	Message *AnyOfmicrosoftGraphMessage `json:"Message,omitempty"`
	isExplicitNullMessage bool `json:"-"`
	Comment *string `json:"Comment,omitempty"`
	isExplicitNullComment bool `json:"-"`
}

// GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.
func (o *InlineObject610) GetToRecipients() []AnyOfmicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []AnyOfmicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject610) GetToRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		var ret []AnyOfmicrosoftGraphRecipient
		return ret, false
	}
	return *o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *InlineObject610) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *InlineObject610) SetToRecipients(v []AnyOfmicrosoftGraphRecipient) {
	o.ToRecipients = &v
}

// GetMessage returns the Message field if non-nil, zero value otherwise.
func (o *InlineObject610) GetMessage() AnyOfmicrosoftGraphMessage {
	if o == nil || o.Message == nil {
		var ret AnyOfmicrosoftGraphMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject610) GetMessageOk() (AnyOfmicrosoftGraphMessage, bool) {
	if o == nil || o.Message == nil {
		var ret AnyOfmicrosoftGraphMessage
		return ret, false
	}
	return *o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObject610) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given AnyOfmicrosoftGraphMessage and assigns it to the Message field.
func (o *InlineObject610) SetMessage(v AnyOfmicrosoftGraphMessage) {
	o.Message = &v
}

// SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Message value is set to nil even if false is passed
func (o *InlineObject610) SetMessageExplicitNull(b bool) {
	o.Message = nil
	o.isExplicitNullMessage = b
}
// GetComment returns the Comment field if non-nil, zero value otherwise.
func (o *InlineObject610) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject610) GetCommentOk() (string, bool) {
	if o == nil || o.Comment == nil {
		var ret string
		return ret, false
	}
	return *o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject610) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineObject610) SetComment(v string) {
	o.Comment = &v
}

// SetCommentExplicitNull (un)sets Comment to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Comment value is set to nil even if false is passed
func (o *InlineObject610) SetCommentExplicitNull(b bool) {
	o.Comment = nil
	o.isExplicitNullComment = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject610) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ToRecipients != nil {
		toSerialize["ToRecipients"] = o.ToRecipients
	}
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


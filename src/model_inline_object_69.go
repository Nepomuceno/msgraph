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
// InlineObject69 struct for InlineObject69
type InlineObject69 struct {
	Comment *string `json:"Comment,omitempty"`
	isExplicitNullComment bool `json:"-"`
	ToRecipients *[]MicrosoftGraphRecipient `json:"ToRecipients,omitempty"`

}

// GetComment returns the Comment field if non-nil, zero value otherwise.
func (o *InlineObject69) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject69) GetCommentOk() (string, bool) {
	if o == nil || o.Comment == nil {
		var ret string
		return ret, false
	}
	return *o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject69) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineObject69) SetComment(v string) {
	o.Comment = &v
}

// SetCommentExplicitNull (un)sets Comment to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Comment value is set to nil even if false is passed
func (o *InlineObject69) SetCommentExplicitNull(b bool) {
	o.Comment = nil
	o.isExplicitNullComment = b
}
// GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.
func (o *InlineObject69) GetToRecipients() []MicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject69) GetToRecipientsOk() ([]MicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret, false
	}
	return *o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *InlineObject69) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *InlineObject69) SetToRecipients(v []MicrosoftGraphRecipient) {
	o.ToRecipients = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject69) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment == nil {
		if o.isExplicitNullComment {
			toSerialize["Comment"] = o.Comment
		}
	} else {
		toSerialize["Comment"] = o.Comment
	}
	if o.ToRecipients != nil {
		toSerialize["ToRecipients"] = o.ToRecipients
	}
	return json.Marshal(toSerialize)
}


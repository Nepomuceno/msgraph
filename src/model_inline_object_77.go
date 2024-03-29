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
// InlineObject77 struct for InlineObject77
type InlineObject77 struct {
	SendResponse *bool `json:"SendResponse,omitempty"`
	isExplicitNullSendResponse bool `json:"-"`
	Comment *string `json:"Comment,omitempty"`
	isExplicitNullComment bool `json:"-"`
}

// GetSendResponse returns the SendResponse field if non-nil, zero value otherwise.
func (o *InlineObject77) GetSendResponse() bool {
	if o == nil || o.SendResponse == nil {
		var ret bool
		return ret
	}
	return *o.SendResponse
}

// GetSendResponseOk returns a tuple with the SendResponse field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject77) GetSendResponseOk() (bool, bool) {
	if o == nil || o.SendResponse == nil {
		var ret bool
		return ret, false
	}
	return *o.SendResponse, true
}

// HasSendResponse returns a boolean if a field has been set.
func (o *InlineObject77) HasSendResponse() bool {
	if o != nil && o.SendResponse != nil {
		return true
	}

	return false
}

// SetSendResponse gets a reference to the given bool and assigns it to the SendResponse field.
func (o *InlineObject77) SetSendResponse(v bool) {
	o.SendResponse = &v
}

// SetSendResponseExplicitNull (un)sets SendResponse to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SendResponse value is set to nil even if false is passed
func (o *InlineObject77) SetSendResponseExplicitNull(b bool) {
	o.SendResponse = nil
	o.isExplicitNullSendResponse = b
}
// GetComment returns the Comment field if non-nil, zero value otherwise.
func (o *InlineObject77) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject77) GetCommentOk() (string, bool) {
	if o == nil || o.Comment == nil {
		var ret string
		return ret, false
	}
	return *o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject77) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineObject77) SetComment(v string) {
	o.Comment = &v
}

// SetCommentExplicitNull (un)sets Comment to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Comment value is set to nil even if false is passed
func (o *InlineObject77) SetCommentExplicitNull(b bool) {
	o.Comment = nil
	o.isExplicitNullComment = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject77) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SendResponse == nil {
		if o.isExplicitNullSendResponse {
			toSerialize["SendResponse"] = o.SendResponse
		}
	} else {
		toSerialize["SendResponse"] = o.SendResponse
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



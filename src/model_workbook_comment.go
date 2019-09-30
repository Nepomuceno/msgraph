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
// WorkbookComment struct for WorkbookComment
type WorkbookComment struct {
	Content *string `json:"content,omitempty"`
	isExplicitNullContent bool `json:"-"`
	ContentType *string `json:"contentType,omitempty"`

	Replies *[]MicrosoftGraphWorkbookCommentReply `json:"replies,omitempty"`

}

// GetContent returns the Content field if non-nil, zero value otherwise.
func (o *WorkbookComment) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookComment) GetContentOk() (string, bool) {
	if o == nil || o.Content == nil {
		var ret string
		return ret, false
	}
	return *o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *WorkbookComment) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *WorkbookComment) SetContent(v string) {
	o.Content = &v
}

// SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Content value is set to nil even if false is passed
func (o *WorkbookComment) SetContentExplicitNull(b bool) {
	o.Content = nil
	o.isExplicitNullContent = b
}
// GetContentType returns the ContentType field if non-nil, zero value otherwise.
func (o *WorkbookComment) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookComment) GetContentTypeOk() (string, bool) {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret, false
	}
	return *o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *WorkbookComment) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *WorkbookComment) SetContentType(v string) {
	o.ContentType = &v
}

// GetReplies returns the Replies field if non-nil, zero value otherwise.
func (o *WorkbookComment) GetReplies() []MicrosoftGraphWorkbookCommentReply {
	if o == nil || o.Replies == nil {
		var ret []MicrosoftGraphWorkbookCommentReply
		return ret
	}
	return *o.Replies
}

// GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookComment) GetRepliesOk() ([]MicrosoftGraphWorkbookCommentReply, bool) {
	if o == nil || o.Replies == nil {
		var ret []MicrosoftGraphWorkbookCommentReply
		return ret, false
	}
	return *o.Replies, true
}

// HasReplies returns a boolean if a field has been set.
func (o *WorkbookComment) HasReplies() bool {
	if o != nil && o.Replies != nil {
		return true
	}

	return false
}

// SetReplies gets a reference to the given []MicrosoftGraphWorkbookCommentReply and assigns it to the Replies field.
func (o *WorkbookComment) SetReplies(v []MicrosoftGraphWorkbookCommentReply) {
	o.Replies = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o WorkbookComment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content == nil {
		if o.isExplicitNullContent {
			toSerialize["content"] = o.Content
		}
	} else {
		toSerialize["content"] = o.Content
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Replies != nil {
		toSerialize["replies"] = o.Replies
	}
	return json.Marshal(toSerialize)
}



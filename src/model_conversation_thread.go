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
	"time"
	"encoding/json"
)
// ConversationThread struct for ConversationThread
type ConversationThread struct {
	ToRecipients *[]MicrosoftGraphRecipient `json:"toRecipients,omitempty"`

	Topic *string `json:"topic,omitempty"`

	HasAttachments *bool `json:"hasAttachments,omitempty"`

	LastDeliveredDateTime *time.Time `json:"lastDeliveredDateTime,omitempty"`

	UniqueSenders *[]string `json:"uniqueSenders,omitempty"`

	CcRecipients *[]MicrosoftGraphRecipient `json:"ccRecipients,omitempty"`

	Preview *string `json:"preview,omitempty"`

	IsLocked *bool `json:"isLocked,omitempty"`

	Posts *[]MicrosoftGraphPost `json:"posts,omitempty"`

}

// GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.
func (o *ConversationThread) GetToRecipients() []MicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetToRecipientsOk() ([]MicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret, false
	}
	return *o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *ConversationThread) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *ConversationThread) SetToRecipients(v []MicrosoftGraphRecipient) {
	o.ToRecipients = &v
}

// GetTopic returns the Topic field if non-nil, zero value otherwise.
func (o *ConversationThread) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetTopicOk() (string, bool) {
	if o == nil || o.Topic == nil {
		var ret string
		return ret, false
	}
	return *o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *ConversationThread) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *ConversationThread) SetTopic(v string) {
	o.Topic = &v
}

// GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.
func (o *ConversationThread) GetHasAttachments() bool {
	if o == nil || o.HasAttachments == nil {
		var ret bool
		return ret
	}
	return *o.HasAttachments
}

// GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetHasAttachmentsOk() (bool, bool) {
	if o == nil || o.HasAttachments == nil {
		var ret bool
		return ret, false
	}
	return *o.HasAttachments, true
}

// HasHasAttachments returns a boolean if a field has been set.
func (o *ConversationThread) HasHasAttachments() bool {
	if o != nil && o.HasAttachments != nil {
		return true
	}

	return false
}

// SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.
func (o *ConversationThread) SetHasAttachments(v bool) {
	o.HasAttachments = &v
}

// GetLastDeliveredDateTime returns the LastDeliveredDateTime field if non-nil, zero value otherwise.
func (o *ConversationThread) GetLastDeliveredDateTime() time.Time {
	if o == nil || o.LastDeliveredDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastDeliveredDateTime
}

// GetLastDeliveredDateTimeOk returns a tuple with the LastDeliveredDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetLastDeliveredDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastDeliveredDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastDeliveredDateTime, true
}

// HasLastDeliveredDateTime returns a boolean if a field has been set.
func (o *ConversationThread) HasLastDeliveredDateTime() bool {
	if o != nil && o.LastDeliveredDateTime != nil {
		return true
	}

	return false
}

// SetLastDeliveredDateTime gets a reference to the given time.Time and assigns it to the LastDeliveredDateTime field.
func (o *ConversationThread) SetLastDeliveredDateTime(v time.Time) {
	o.LastDeliveredDateTime = &v
}

// GetUniqueSenders returns the UniqueSenders field if non-nil, zero value otherwise.
func (o *ConversationThread) GetUniqueSenders() []string {
	if o == nil || o.UniqueSenders == nil {
		var ret []string
		return ret
	}
	return *o.UniqueSenders
}

// GetUniqueSendersOk returns a tuple with the UniqueSenders field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetUniqueSendersOk() ([]string, bool) {
	if o == nil || o.UniqueSenders == nil {
		var ret []string
		return ret, false
	}
	return *o.UniqueSenders, true
}

// HasUniqueSenders returns a boolean if a field has been set.
func (o *ConversationThread) HasUniqueSenders() bool {
	if o != nil && o.UniqueSenders != nil {
		return true
	}

	return false
}

// SetUniqueSenders gets a reference to the given []string and assigns it to the UniqueSenders field.
func (o *ConversationThread) SetUniqueSenders(v []string) {
	o.UniqueSenders = &v
}

// GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.
func (o *ConversationThread) GetCcRecipients() []MicrosoftGraphRecipient {
	if o == nil || o.CcRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret
	}
	return *o.CcRecipients
}

// GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetCcRecipientsOk() ([]MicrosoftGraphRecipient, bool) {
	if o == nil || o.CcRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret, false
	}
	return *o.CcRecipients, true
}

// HasCcRecipients returns a boolean if a field has been set.
func (o *ConversationThread) HasCcRecipients() bool {
	if o != nil && o.CcRecipients != nil {
		return true
	}

	return false
}

// SetCcRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the CcRecipients field.
func (o *ConversationThread) SetCcRecipients(v []MicrosoftGraphRecipient) {
	o.CcRecipients = &v
}

// GetPreview returns the Preview field if non-nil, zero value otherwise.
func (o *ConversationThread) GetPreview() string {
	if o == nil || o.Preview == nil {
		var ret string
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetPreviewOk() (string, bool) {
	if o == nil || o.Preview == nil {
		var ret string
		return ret, false
	}
	return *o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *ConversationThread) HasPreview() bool {
	if o != nil && o.Preview != nil {
		return true
	}

	return false
}

// SetPreview gets a reference to the given string and assigns it to the Preview field.
func (o *ConversationThread) SetPreview(v string) {
	o.Preview = &v
}

// GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.
func (o *ConversationThread) GetIsLocked() bool {
	if o == nil || o.IsLocked == nil {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetIsLockedOk() (bool, bool) {
	if o == nil || o.IsLocked == nil {
		var ret bool
		return ret, false
	}
	return *o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *ConversationThread) HasIsLocked() bool {
	if o != nil && o.IsLocked != nil {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *ConversationThread) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetPosts returns the Posts field if non-nil, zero value otherwise.
func (o *ConversationThread) GetPosts() []MicrosoftGraphPost {
	if o == nil || o.Posts == nil {
		var ret []MicrosoftGraphPost
		return ret
	}
	return *o.Posts
}

// GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetPostsOk() ([]MicrosoftGraphPost, bool) {
	if o == nil || o.Posts == nil {
		var ret []MicrosoftGraphPost
		return ret, false
	}
	return *o.Posts, true
}

// HasPosts returns a boolean if a field has been set.
func (o *ConversationThread) HasPosts() bool {
	if o != nil && o.Posts != nil {
		return true
	}

	return false
}

// SetPosts gets a reference to the given []MicrosoftGraphPost and assigns it to the Posts field.
func (o *ConversationThread) SetPosts(v []MicrosoftGraphPost) {
	o.Posts = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o ConversationThread) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ToRecipients != nil {
		toSerialize["toRecipients"] = o.ToRecipients
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.HasAttachments != nil {
		toSerialize["hasAttachments"] = o.HasAttachments
	}
	if o.LastDeliveredDateTime != nil {
		toSerialize["lastDeliveredDateTime"] = o.LastDeliveredDateTime
	}
	if o.UniqueSenders != nil {
		toSerialize["uniqueSenders"] = o.UniqueSenders
	}
	if o.CcRecipients != nil {
		toSerialize["ccRecipients"] = o.CcRecipients
	}
	if o.Preview != nil {
		toSerialize["preview"] = o.Preview
	}
	if o.IsLocked != nil {
		toSerialize["isLocked"] = o.IsLocked
	}
	if o.Posts != nil {
		toSerialize["posts"] = o.Posts
	}
	return json.Marshal(toSerialize)
}



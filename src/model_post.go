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
// Post struct for Post
type Post struct {
	Body *AnyOfmicrosoftGraphItemBody `json:"body,omitempty"`
	isExplicitNullBody bool `json:"-"`
	ReceivedDateTime *time.Time `json:"receivedDateTime,omitempty"`

	HasAttachments *bool `json:"hasAttachments,omitempty"`

	From *MicrosoftGraphRecipient `json:"from,omitempty"`

	Sender *AnyOfmicrosoftGraphRecipient `json:"sender,omitempty"`
	isExplicitNullSender bool `json:"-"`
	ConversationThreadId *string `json:"conversationThreadId,omitempty"`
	isExplicitNullConversationThreadId bool `json:"-"`
	NewParticipants *[]MicrosoftGraphRecipient `json:"newParticipants,omitempty"`

	ConversationId *string `json:"conversationId,omitempty"`
	isExplicitNullConversationId bool `json:"-"`
	InReplyTo *AnyOfmicrosoftGraphPost `json:"inReplyTo,omitempty"`
	isExplicitNullInReplyTo bool `json:"-"`
	SingleValueExtendedProperties *[]MicrosoftGraphSingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	MultiValueExtendedProperties *[]MicrosoftGraphMultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	Extensions *[]MicrosoftGraphExtension `json:"extensions,omitempty"`

	Attachments *[]MicrosoftGraphAttachment `json:"attachments,omitempty"`

}

// GetBody returns the Body field if non-nil, zero value otherwise.
func (o *Post) GetBody() AnyOfmicrosoftGraphItemBody {
	if o == nil || o.Body == nil {
		var ret AnyOfmicrosoftGraphItemBody
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetBodyOk() (AnyOfmicrosoftGraphItemBody, bool) {
	if o == nil || o.Body == nil {
		var ret AnyOfmicrosoftGraphItemBody
		return ret, false
	}
	return *o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *Post) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the Body field.
func (o *Post) SetBody(v AnyOfmicrosoftGraphItemBody) {
	o.Body = &v
}

// SetBodyExplicitNull (un)sets Body to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Body value is set to nil even if false is passed
func (o *Post) SetBodyExplicitNull(b bool) {
	o.Body = nil
	o.isExplicitNullBody = b
}
// GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.
func (o *Post) GetReceivedDateTime() time.Time {
	if o == nil || o.ReceivedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ReceivedDateTime
}

// GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetReceivedDateTimeOk() (time.Time, bool) {
	if o == nil || o.ReceivedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ReceivedDateTime, true
}

// HasReceivedDateTime returns a boolean if a field has been set.
func (o *Post) HasReceivedDateTime() bool {
	if o != nil && o.ReceivedDateTime != nil {
		return true
	}

	return false
}

// SetReceivedDateTime gets a reference to the given time.Time and assigns it to the ReceivedDateTime field.
func (o *Post) SetReceivedDateTime(v time.Time) {
	o.ReceivedDateTime = &v
}

// GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.
func (o *Post) GetHasAttachments() bool {
	if o == nil || o.HasAttachments == nil {
		var ret bool
		return ret
	}
	return *o.HasAttachments
}

// GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetHasAttachmentsOk() (bool, bool) {
	if o == nil || o.HasAttachments == nil {
		var ret bool
		return ret, false
	}
	return *o.HasAttachments, true
}

// HasHasAttachments returns a boolean if a field has been set.
func (o *Post) HasHasAttachments() bool {
	if o != nil && o.HasAttachments != nil {
		return true
	}

	return false
}

// SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.
func (o *Post) SetHasAttachments(v bool) {
	o.HasAttachments = &v
}

// GetFrom returns the From field if non-nil, zero value otherwise.
func (o *Post) GetFrom() MicrosoftGraphRecipient {
	if o == nil || o.From == nil {
		var ret MicrosoftGraphRecipient
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetFromOk() (MicrosoftGraphRecipient, bool) {
	if o == nil || o.From == nil {
		var ret MicrosoftGraphRecipient
		return ret, false
	}
	return *o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *Post) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given MicrosoftGraphRecipient and assigns it to the From field.
func (o *Post) SetFrom(v MicrosoftGraphRecipient) {
	o.From = &v
}

// GetSender returns the Sender field if non-nil, zero value otherwise.
func (o *Post) GetSender() AnyOfmicrosoftGraphRecipient {
	if o == nil || o.Sender == nil {
		var ret AnyOfmicrosoftGraphRecipient
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetSenderOk() (AnyOfmicrosoftGraphRecipient, bool) {
	if o == nil || o.Sender == nil {
		var ret AnyOfmicrosoftGraphRecipient
		return ret, false
	}
	return *o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *Post) HasSender() bool {
	if o != nil && o.Sender != nil {
		return true
	}

	return false
}

// SetSender gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the Sender field.
func (o *Post) SetSender(v AnyOfmicrosoftGraphRecipient) {
	o.Sender = &v
}

// SetSenderExplicitNull (un)sets Sender to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Sender value is set to nil even if false is passed
func (o *Post) SetSenderExplicitNull(b bool) {
	o.Sender = nil
	o.isExplicitNullSender = b
}
// GetConversationThreadId returns the ConversationThreadId field if non-nil, zero value otherwise.
func (o *Post) GetConversationThreadId() string {
	if o == nil || o.ConversationThreadId == nil {
		var ret string
		return ret
	}
	return *o.ConversationThreadId
}

// GetConversationThreadIdOk returns a tuple with the ConversationThreadId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetConversationThreadIdOk() (string, bool) {
	if o == nil || o.ConversationThreadId == nil {
		var ret string
		return ret, false
	}
	return *o.ConversationThreadId, true
}

// HasConversationThreadId returns a boolean if a field has been set.
func (o *Post) HasConversationThreadId() bool {
	if o != nil && o.ConversationThreadId != nil {
		return true
	}

	return false
}

// SetConversationThreadId gets a reference to the given string and assigns it to the ConversationThreadId field.
func (o *Post) SetConversationThreadId(v string) {
	o.ConversationThreadId = &v
}

// SetConversationThreadIdExplicitNull (un)sets ConversationThreadId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ConversationThreadId value is set to nil even if false is passed
func (o *Post) SetConversationThreadIdExplicitNull(b bool) {
	o.ConversationThreadId = nil
	o.isExplicitNullConversationThreadId = b
}
// GetNewParticipants returns the NewParticipants field if non-nil, zero value otherwise.
func (o *Post) GetNewParticipants() []MicrosoftGraphRecipient {
	if o == nil || o.NewParticipants == nil {
		var ret []MicrosoftGraphRecipient
		return ret
	}
	return *o.NewParticipants
}

// GetNewParticipantsOk returns a tuple with the NewParticipants field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetNewParticipantsOk() ([]MicrosoftGraphRecipient, bool) {
	if o == nil || o.NewParticipants == nil {
		var ret []MicrosoftGraphRecipient
		return ret, false
	}
	return *o.NewParticipants, true
}

// HasNewParticipants returns a boolean if a field has been set.
func (o *Post) HasNewParticipants() bool {
	if o != nil && o.NewParticipants != nil {
		return true
	}

	return false
}

// SetNewParticipants gets a reference to the given []MicrosoftGraphRecipient and assigns it to the NewParticipants field.
func (o *Post) SetNewParticipants(v []MicrosoftGraphRecipient) {
	o.NewParticipants = &v
}

// GetConversationId returns the ConversationId field if non-nil, zero value otherwise.
func (o *Post) GetConversationId() string {
	if o == nil || o.ConversationId == nil {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetConversationIdOk() (string, bool) {
	if o == nil || o.ConversationId == nil {
		var ret string
		return ret, false
	}
	return *o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *Post) HasConversationId() bool {
	if o != nil && o.ConversationId != nil {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *Post) SetConversationId(v string) {
	o.ConversationId = &v
}

// SetConversationIdExplicitNull (un)sets ConversationId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ConversationId value is set to nil even if false is passed
func (o *Post) SetConversationIdExplicitNull(b bool) {
	o.ConversationId = nil
	o.isExplicitNullConversationId = b
}
// GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.
func (o *Post) GetInReplyTo() AnyOfmicrosoftGraphPost {
	if o == nil || o.InReplyTo == nil {
		var ret AnyOfmicrosoftGraphPost
		return ret
	}
	return *o.InReplyTo
}

// GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetInReplyToOk() (AnyOfmicrosoftGraphPost, bool) {
	if o == nil || o.InReplyTo == nil {
		var ret AnyOfmicrosoftGraphPost
		return ret, false
	}
	return *o.InReplyTo, true
}

// HasInReplyTo returns a boolean if a field has been set.
func (o *Post) HasInReplyTo() bool {
	if o != nil && o.InReplyTo != nil {
		return true
	}

	return false
}

// SetInReplyTo gets a reference to the given AnyOfmicrosoftGraphPost and assigns it to the InReplyTo field.
func (o *Post) SetInReplyTo(v AnyOfmicrosoftGraphPost) {
	o.InReplyTo = &v
}

// SetInReplyToExplicitNull (un)sets InReplyTo to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The InReplyTo value is set to nil even if false is passed
func (o *Post) SetInReplyToExplicitNull(b bool) {
	o.InReplyTo = nil
	o.isExplicitNullInReplyTo = b
}
// GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.
func (o *Post) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty {
	if o == nil || o.SingleValueExtendedProperties == nil {
		var ret []MicrosoftGraphSingleValueLegacyExtendedProperty
		return ret
	}
	return *o.SingleValueExtendedProperties
}

// GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool) {
	if o == nil || o.SingleValueExtendedProperties == nil {
		var ret []MicrosoftGraphSingleValueLegacyExtendedProperty
		return ret, false
	}
	return *o.SingleValueExtendedProperties, true
}

// HasSingleValueExtendedProperties returns a boolean if a field has been set.
func (o *Post) HasSingleValueExtendedProperties() bool {
	if o != nil && o.SingleValueExtendedProperties != nil {
		return true
	}

	return false
}

// SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.
func (o *Post) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty) {
	o.SingleValueExtendedProperties = &v
}

// GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.
func (o *Post) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty {
	if o == nil || o.MultiValueExtendedProperties == nil {
		var ret []MicrosoftGraphMultiValueLegacyExtendedProperty
		return ret
	}
	return *o.MultiValueExtendedProperties
}

// GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool) {
	if o == nil || o.MultiValueExtendedProperties == nil {
		var ret []MicrosoftGraphMultiValueLegacyExtendedProperty
		return ret, false
	}
	return *o.MultiValueExtendedProperties, true
}

// HasMultiValueExtendedProperties returns a boolean if a field has been set.
func (o *Post) HasMultiValueExtendedProperties() bool {
	if o != nil && o.MultiValueExtendedProperties != nil {
		return true
	}

	return false
}

// SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.
func (o *Post) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty) {
	o.MultiValueExtendedProperties = &v
}

// GetExtensions returns the Extensions field if non-nil, zero value otherwise.
func (o *Post) GetExtensions() []MicrosoftGraphExtension {
	if o == nil || o.Extensions == nil {
		var ret []MicrosoftGraphExtension
		return ret
	}
	return *o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetExtensionsOk() ([]MicrosoftGraphExtension, bool) {
	if o == nil || o.Extensions == nil {
		var ret []MicrosoftGraphExtension
		return ret, false
	}
	return *o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Post) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.
func (o *Post) SetExtensions(v []MicrosoftGraphExtension) {
	o.Extensions = &v
}

// GetAttachments returns the Attachments field if non-nil, zero value otherwise.
func (o *Post) GetAttachments() []MicrosoftGraphAttachment {
	if o == nil || o.Attachments == nil {
		var ret []MicrosoftGraphAttachment
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Post) GetAttachmentsOk() ([]MicrosoftGraphAttachment, bool) {
	if o == nil || o.Attachments == nil {
		var ret []MicrosoftGraphAttachment
		return ret, false
	}
	return *o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *Post) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []MicrosoftGraphAttachment and assigns it to the Attachments field.
func (o *Post) SetAttachments(v []MicrosoftGraphAttachment) {
	o.Attachments = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o Post) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body == nil {
		if o.isExplicitNullBody {
			toSerialize["body"] = o.Body
		}
	} else {
		toSerialize["body"] = o.Body
	}
	if o.ReceivedDateTime != nil {
		toSerialize["receivedDateTime"] = o.ReceivedDateTime
	}
	if o.HasAttachments != nil {
		toSerialize["hasAttachments"] = o.HasAttachments
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Sender == nil {
		if o.isExplicitNullSender {
			toSerialize["sender"] = o.Sender
		}
	} else {
		toSerialize["sender"] = o.Sender
	}
	if o.ConversationThreadId == nil {
		if o.isExplicitNullConversationThreadId {
			toSerialize["conversationThreadId"] = o.ConversationThreadId
		}
	} else {
		toSerialize["conversationThreadId"] = o.ConversationThreadId
	}
	if o.NewParticipants != nil {
		toSerialize["newParticipants"] = o.NewParticipants
	}
	if o.ConversationId == nil {
		if o.isExplicitNullConversationId {
			toSerialize["conversationId"] = o.ConversationId
		}
	} else {
		toSerialize["conversationId"] = o.ConversationId
	}
	if o.InReplyTo == nil {
		if o.isExplicitNullInReplyTo {
			toSerialize["inReplyTo"] = o.InReplyTo
		}
	} else {
		toSerialize["inReplyTo"] = o.InReplyTo
	}
	if o.SingleValueExtendedProperties != nil {
		toSerialize["singleValueExtendedProperties"] = o.SingleValueExtendedProperties
	}
	if o.MultiValueExtendedProperties != nil {
		toSerialize["multiValueExtendedProperties"] = o.MultiValueExtendedProperties
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}


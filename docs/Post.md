# Post

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) |  | [optional] 
**ReceivedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**HasAttachments** | Pointer to **bool** |  | [optional] 
**From** | Pointer to [**MicrosoftGraphRecipient**](microsoft.graph.recipient.md) |  | [optional] 
**Sender** | Pointer to [**AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**ConversationThreadId** | Pointer to **string** |  | [optional] 
**NewParticipants** | Pointer to [**[]MicrosoftGraphRecipient**](microsoft.graph.recipient.md) |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**InReplyTo** | Pointer to [**AnyOfmicrosoftGraphPost**](anyOf&lt;microsoft.graph.post&gt;.md) |  | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md) |  | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md) |  | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](microsoft.graph.extension.md) |  | [optional] 
**Attachments** | Pointer to [**[]MicrosoftGraphAttachment**](microsoft.graph.attachment.md) |  | [optional] 

## Methods

### GetBody

`func (o *Post) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Post) GetBodyOk() (AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *Post) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *Post) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the Body field.

### SetBodyExplicitNull

`func (o *Post) SetBodyExplicitNull(b bool)`

SetBodyExplicitNull (un)sets Body to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Body value is set to nil even if false is passed
### GetReceivedDateTime

`func (o *Post) GetReceivedDateTime() time.Time`

GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.

### GetReceivedDateTimeOk

`func (o *Post) GetReceivedDateTimeOk() (time.Time, bool)`

GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReceivedDateTime

`func (o *Post) HasReceivedDateTime() bool`

HasReceivedDateTime returns a boolean if a field has been set.

### SetReceivedDateTime

`func (o *Post) SetReceivedDateTime(v time.Time)`

SetReceivedDateTime gets a reference to the given time.Time and assigns it to the ReceivedDateTime field.

### GetHasAttachments

`func (o *Post) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Post) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *Post) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *Post) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### GetFrom

`func (o *Post) GetFrom() MicrosoftGraphRecipient`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Post) GetFromOk() (MicrosoftGraphRecipient, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrom

`func (o *Post) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFrom

`func (o *Post) SetFrom(v MicrosoftGraphRecipient)`

SetFrom gets a reference to the given MicrosoftGraphRecipient and assigns it to the From field.

### GetSender

`func (o *Post) GetSender() AnyOfmicrosoftGraphRecipient`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Post) GetSenderOk() (AnyOfmicrosoftGraphRecipient, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSender

`func (o *Post) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSender

`func (o *Post) SetSender(v AnyOfmicrosoftGraphRecipient)`

SetSender gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the Sender field.

### SetSenderExplicitNull

`func (o *Post) SetSenderExplicitNull(b bool)`

SetSenderExplicitNull (un)sets Sender to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sender value is set to nil even if false is passed
### GetConversationThreadId

`func (o *Post) GetConversationThreadId() string`

GetConversationThreadId returns the ConversationThreadId field if non-nil, zero value otherwise.

### GetConversationThreadIdOk

`func (o *Post) GetConversationThreadIdOk() (string, bool)`

GetConversationThreadIdOk returns a tuple with the ConversationThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversationThreadId

`func (o *Post) HasConversationThreadId() bool`

HasConversationThreadId returns a boolean if a field has been set.

### SetConversationThreadId

`func (o *Post) SetConversationThreadId(v string)`

SetConversationThreadId gets a reference to the given string and assigns it to the ConversationThreadId field.

### SetConversationThreadIdExplicitNull

`func (o *Post) SetConversationThreadIdExplicitNull(b bool)`

SetConversationThreadIdExplicitNull (un)sets ConversationThreadId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConversationThreadId value is set to nil even if false is passed
### GetNewParticipants

`func (o *Post) GetNewParticipants() []MicrosoftGraphRecipient`

GetNewParticipants returns the NewParticipants field if non-nil, zero value otherwise.

### GetNewParticipantsOk

`func (o *Post) GetNewParticipantsOk() ([]MicrosoftGraphRecipient, bool)`

GetNewParticipantsOk returns a tuple with the NewParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewParticipants

`func (o *Post) HasNewParticipants() bool`

HasNewParticipants returns a boolean if a field has been set.

### SetNewParticipants

`func (o *Post) SetNewParticipants(v []MicrosoftGraphRecipient)`

SetNewParticipants gets a reference to the given []MicrosoftGraphRecipient and assigns it to the NewParticipants field.

### GetConversationId

`func (o *Post) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *Post) GetConversationIdOk() (string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversationId

`func (o *Post) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationId

`func (o *Post) SetConversationId(v string)`

SetConversationId gets a reference to the given string and assigns it to the ConversationId field.

### SetConversationIdExplicitNull

`func (o *Post) SetConversationIdExplicitNull(b bool)`

SetConversationIdExplicitNull (un)sets ConversationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConversationId value is set to nil even if false is passed
### GetInReplyTo

`func (o *Post) GetInReplyTo() AnyOfmicrosoftGraphPost`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *Post) GetInReplyToOk() (AnyOfmicrosoftGraphPost, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInReplyTo

`func (o *Post) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.

### SetInReplyTo

`func (o *Post) SetInReplyTo(v AnyOfmicrosoftGraphPost)`

SetInReplyTo gets a reference to the given AnyOfmicrosoftGraphPost and assigns it to the InReplyTo field.

### SetInReplyToExplicitNull

`func (o *Post) SetInReplyToExplicitNull(b bool)`

SetInReplyToExplicitNull (un)sets InReplyTo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InReplyTo value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *Post) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Post) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *Post) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *Post) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *Post) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Post) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *Post) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *Post) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetExtensions

`func (o *Post) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Post) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *Post) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *Post) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.

### GetAttachments

`func (o *Post) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Post) GetAttachmentsOk() ([]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttachments

`func (o *Post) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachments

`func (o *Post) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments gets a reference to the given []MicrosoftGraphAttachment and assigns it to the Attachments field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



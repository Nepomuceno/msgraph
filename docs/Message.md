# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceivedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**SentDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**HasAttachments** | Pointer to **bool** |  | [optional] 
**InternetMessageId** | Pointer to **string** |  | [optional] 
**InternetMessageHeaders** | Pointer to [**[]AnyOfmicrosoftGraphInternetMessageHeader**](anyOf&lt;microsoft.graph.internetMessageHeader&gt;.md) |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Body** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) |  | [optional] 
**BodyPreview** | Pointer to **string** |  | [optional] 
**Importance** | Pointer to [**AnyOfmicrosoftGraphImportance**](anyOf&lt;microsoft.graph.importance&gt;.md) |  | [optional] 
**ParentFolderId** | Pointer to **string** |  | [optional] 
**Sender** | Pointer to [**AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**From** | Pointer to [**AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**ToRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**CcRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**BccRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**ReplyTo** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**UniqueBody** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) |  | [optional] 
**IsDeliveryReceiptRequested** | Pointer to **bool** |  | [optional] 
**IsReadReceiptRequested** | Pointer to **bool** |  | [optional] 
**IsRead** | Pointer to **bool** |  | [optional] 
**IsDraft** | Pointer to **bool** |  | [optional] 
**WebLink** | Pointer to **string** |  | [optional] 
**InferenceClassification** | Pointer to [**AnyOfmicrosoftGraphInferenceClassificationType**](anyOf&lt;microsoft.graph.inferenceClassificationType&gt;.md) |  | [optional] 
**Flag** | Pointer to [**AnyOfmicrosoftGraphFollowupFlag**](anyOf&lt;microsoft.graph.followupFlag&gt;.md) |  | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](microsoft.graph.singleValueLegacyExtendedProperty.md) |  | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](microsoft.graph.multiValueLegacyExtendedProperty.md) |  | [optional] 
**Attachments** | Pointer to [**[]MicrosoftGraphAttachment**](microsoft.graph.attachment.md) |  | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](microsoft.graph.extension.md) |  | [optional] 

## Methods

### GetReceivedDateTime

`func (o *Message) GetReceivedDateTime() time.Time`

GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.

### GetReceivedDateTimeOk

`func (o *Message) GetReceivedDateTimeOk() (time.Time, bool)`

GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReceivedDateTime

`func (o *Message) HasReceivedDateTime() bool`

HasReceivedDateTime returns a boolean if a field has been set.

### SetReceivedDateTime

`func (o *Message) SetReceivedDateTime(v time.Time)`

SetReceivedDateTime gets a reference to the given time.Time and assigns it to the ReceivedDateTime field.

### SetReceivedDateTimeExplicitNull

`func (o *Message) SetReceivedDateTimeExplicitNull(b bool)`

SetReceivedDateTimeExplicitNull (un)sets ReceivedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReceivedDateTime value is set to nil even if false is passed
### GetSentDateTime

`func (o *Message) GetSentDateTime() time.Time`

GetSentDateTime returns the SentDateTime field if non-nil, zero value otherwise.

### GetSentDateTimeOk

`func (o *Message) GetSentDateTimeOk() (time.Time, bool)`

GetSentDateTimeOk returns a tuple with the SentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSentDateTime

`func (o *Message) HasSentDateTime() bool`

HasSentDateTime returns a boolean if a field has been set.

### SetSentDateTime

`func (o *Message) SetSentDateTime(v time.Time)`

SetSentDateTime gets a reference to the given time.Time and assigns it to the SentDateTime field.

### SetSentDateTimeExplicitNull

`func (o *Message) SetSentDateTimeExplicitNull(b bool)`

SetSentDateTimeExplicitNull (un)sets SentDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SentDateTime value is set to nil even if false is passed
### GetHasAttachments

`func (o *Message) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Message) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *Message) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *Message) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### SetHasAttachmentsExplicitNull

`func (o *Message) SetHasAttachmentsExplicitNull(b bool)`

SetHasAttachmentsExplicitNull (un)sets HasAttachments to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasAttachments value is set to nil even if false is passed
### GetInternetMessageId

`func (o *Message) GetInternetMessageId() string`

GetInternetMessageId returns the InternetMessageId field if non-nil, zero value otherwise.

### GetInternetMessageIdOk

`func (o *Message) GetInternetMessageIdOk() (string, bool)`

GetInternetMessageIdOk returns a tuple with the InternetMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetMessageId

`func (o *Message) HasInternetMessageId() bool`

HasInternetMessageId returns a boolean if a field has been set.

### SetInternetMessageId

`func (o *Message) SetInternetMessageId(v string)`

SetInternetMessageId gets a reference to the given string and assigns it to the InternetMessageId field.

### SetInternetMessageIdExplicitNull

`func (o *Message) SetInternetMessageIdExplicitNull(b bool)`

SetInternetMessageIdExplicitNull (un)sets InternetMessageId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InternetMessageId value is set to nil even if false is passed
### GetInternetMessageHeaders

`func (o *Message) GetInternetMessageHeaders() []AnyOfmicrosoftGraphInternetMessageHeader`

GetInternetMessageHeaders returns the InternetMessageHeaders field if non-nil, zero value otherwise.

### GetInternetMessageHeadersOk

`func (o *Message) GetInternetMessageHeadersOk() ([]AnyOfmicrosoftGraphInternetMessageHeader, bool)`

GetInternetMessageHeadersOk returns a tuple with the InternetMessageHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetMessageHeaders

`func (o *Message) HasInternetMessageHeaders() bool`

HasInternetMessageHeaders returns a boolean if a field has been set.

### SetInternetMessageHeaders

`func (o *Message) SetInternetMessageHeaders(v []AnyOfmicrosoftGraphInternetMessageHeader)`

SetInternetMessageHeaders gets a reference to the given []AnyOfmicrosoftGraphInternetMessageHeader and assigns it to the InternetMessageHeaders field.

### GetSubject

`func (o *Message) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Message) GetSubjectOk() (string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubject

`func (o *Message) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubject

`func (o *Message) SetSubject(v string)`

SetSubject gets a reference to the given string and assigns it to the Subject field.

### SetSubjectExplicitNull

`func (o *Message) SetSubjectExplicitNull(b bool)`

SetSubjectExplicitNull (un)sets Subject to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Subject value is set to nil even if false is passed
### GetBody

`func (o *Message) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Message) GetBodyOk() (AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *Message) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *Message) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the Body field.

### SetBodyExplicitNull

`func (o *Message) SetBodyExplicitNull(b bool)`

SetBodyExplicitNull (un)sets Body to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Body value is set to nil even if false is passed
### GetBodyPreview

`func (o *Message) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *Message) GetBodyPreviewOk() (string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBodyPreview

`func (o *Message) HasBodyPreview() bool`

HasBodyPreview returns a boolean if a field has been set.

### SetBodyPreview

`func (o *Message) SetBodyPreview(v string)`

SetBodyPreview gets a reference to the given string and assigns it to the BodyPreview field.

### SetBodyPreviewExplicitNull

`func (o *Message) SetBodyPreviewExplicitNull(b bool)`

SetBodyPreviewExplicitNull (un)sets BodyPreview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BodyPreview value is set to nil even if false is passed
### GetImportance

`func (o *Message) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *Message) GetImportanceOk() (AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportance

`func (o *Message) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportance

`func (o *Message) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance gets a reference to the given AnyOfmicrosoftGraphImportance and assigns it to the Importance field.

### SetImportanceExplicitNull

`func (o *Message) SetImportanceExplicitNull(b bool)`

SetImportanceExplicitNull (un)sets Importance to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Importance value is set to nil even if false is passed
### GetParentFolderId

`func (o *Message) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *Message) GetParentFolderIdOk() (string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentFolderId

`func (o *Message) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderId

`func (o *Message) SetParentFolderId(v string)`

SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.

### SetParentFolderIdExplicitNull

`func (o *Message) SetParentFolderIdExplicitNull(b bool)`

SetParentFolderIdExplicitNull (un)sets ParentFolderId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentFolderId value is set to nil even if false is passed
### GetSender

`func (o *Message) GetSender() AnyOfmicrosoftGraphRecipient`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Message) GetSenderOk() (AnyOfmicrosoftGraphRecipient, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSender

`func (o *Message) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSender

`func (o *Message) SetSender(v AnyOfmicrosoftGraphRecipient)`

SetSender gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the Sender field.

### SetSenderExplicitNull

`func (o *Message) SetSenderExplicitNull(b bool)`

SetSenderExplicitNull (un)sets Sender to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sender value is set to nil even if false is passed
### GetFrom

`func (o *Message) GetFrom() AnyOfmicrosoftGraphRecipient`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Message) GetFromOk() (AnyOfmicrosoftGraphRecipient, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrom

`func (o *Message) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFrom

`func (o *Message) SetFrom(v AnyOfmicrosoftGraphRecipient)`

SetFrom gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the From field.

### SetFromExplicitNull

`func (o *Message) SetFromExplicitNull(b bool)`

SetFromExplicitNull (un)sets From to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The From value is set to nil even if false is passed
### GetToRecipients

`func (o *Message) GetToRecipients() []AnyOfmicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *Message) GetToRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToRecipients

`func (o *Message) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### SetToRecipients

`func (o *Message) SetToRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetToRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ToRecipients field.

### GetCcRecipients

`func (o *Message) GetCcRecipients() []AnyOfmicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *Message) GetCcRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCcRecipients

`func (o *Message) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### SetCcRecipients

`func (o *Message) SetCcRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetCcRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the CcRecipients field.

### GetBccRecipients

`func (o *Message) GetBccRecipients() []AnyOfmicrosoftGraphRecipient`

GetBccRecipients returns the BccRecipients field if non-nil, zero value otherwise.

### GetBccRecipientsOk

`func (o *Message) GetBccRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetBccRecipientsOk returns a tuple with the BccRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBccRecipients

`func (o *Message) HasBccRecipients() bool`

HasBccRecipients returns a boolean if a field has been set.

### SetBccRecipients

`func (o *Message) SetBccRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetBccRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the BccRecipients field.

### GetReplyTo

`func (o *Message) GetReplyTo() []AnyOfmicrosoftGraphRecipient`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *Message) GetReplyToOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReplyTo

`func (o *Message) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### SetReplyTo

`func (o *Message) SetReplyTo(v []AnyOfmicrosoftGraphRecipient)`

SetReplyTo gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ReplyTo field.

### GetConversationId

`func (o *Message) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *Message) GetConversationIdOk() (string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversationId

`func (o *Message) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationId

`func (o *Message) SetConversationId(v string)`

SetConversationId gets a reference to the given string and assigns it to the ConversationId field.

### SetConversationIdExplicitNull

`func (o *Message) SetConversationIdExplicitNull(b bool)`

SetConversationIdExplicitNull (un)sets ConversationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConversationId value is set to nil even if false is passed
### GetUniqueBody

`func (o *Message) GetUniqueBody() AnyOfmicrosoftGraphItemBody`

GetUniqueBody returns the UniqueBody field if non-nil, zero value otherwise.

### GetUniqueBodyOk

`func (o *Message) GetUniqueBodyOk() (AnyOfmicrosoftGraphItemBody, bool)`

GetUniqueBodyOk returns a tuple with the UniqueBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueBody

`func (o *Message) HasUniqueBody() bool`

HasUniqueBody returns a boolean if a field has been set.

### SetUniqueBody

`func (o *Message) SetUniqueBody(v AnyOfmicrosoftGraphItemBody)`

SetUniqueBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the UniqueBody field.

### SetUniqueBodyExplicitNull

`func (o *Message) SetUniqueBodyExplicitNull(b bool)`

SetUniqueBodyExplicitNull (un)sets UniqueBody to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UniqueBody value is set to nil even if false is passed
### GetIsDeliveryReceiptRequested

`func (o *Message) GetIsDeliveryReceiptRequested() bool`

GetIsDeliveryReceiptRequested returns the IsDeliveryReceiptRequested field if non-nil, zero value otherwise.

### GetIsDeliveryReceiptRequestedOk

`func (o *Message) GetIsDeliveryReceiptRequestedOk() (bool, bool)`

GetIsDeliveryReceiptRequestedOk returns a tuple with the IsDeliveryReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDeliveryReceiptRequested

`func (o *Message) HasIsDeliveryReceiptRequested() bool`

HasIsDeliveryReceiptRequested returns a boolean if a field has been set.

### SetIsDeliveryReceiptRequested

`func (o *Message) SetIsDeliveryReceiptRequested(v bool)`

SetIsDeliveryReceiptRequested gets a reference to the given bool and assigns it to the IsDeliveryReceiptRequested field.

### SetIsDeliveryReceiptRequestedExplicitNull

`func (o *Message) SetIsDeliveryReceiptRequestedExplicitNull(b bool)`

SetIsDeliveryReceiptRequestedExplicitNull (un)sets IsDeliveryReceiptRequested to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDeliveryReceiptRequested value is set to nil even if false is passed
### GetIsReadReceiptRequested

`func (o *Message) GetIsReadReceiptRequested() bool`

GetIsReadReceiptRequested returns the IsReadReceiptRequested field if non-nil, zero value otherwise.

### GetIsReadReceiptRequestedOk

`func (o *Message) GetIsReadReceiptRequestedOk() (bool, bool)`

GetIsReadReceiptRequestedOk returns a tuple with the IsReadReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReadReceiptRequested

`func (o *Message) HasIsReadReceiptRequested() bool`

HasIsReadReceiptRequested returns a boolean if a field has been set.

### SetIsReadReceiptRequested

`func (o *Message) SetIsReadReceiptRequested(v bool)`

SetIsReadReceiptRequested gets a reference to the given bool and assigns it to the IsReadReceiptRequested field.

### SetIsReadReceiptRequestedExplicitNull

`func (o *Message) SetIsReadReceiptRequestedExplicitNull(b bool)`

SetIsReadReceiptRequestedExplicitNull (un)sets IsReadReceiptRequested to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsReadReceiptRequested value is set to nil even if false is passed
### GetIsRead

`func (o *Message) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *Message) GetIsReadOk() (bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsRead

`func (o *Message) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### SetIsRead

`func (o *Message) SetIsRead(v bool)`

SetIsRead gets a reference to the given bool and assigns it to the IsRead field.

### SetIsReadExplicitNull

`func (o *Message) SetIsReadExplicitNull(b bool)`

SetIsReadExplicitNull (un)sets IsRead to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsRead value is set to nil even if false is passed
### GetIsDraft

`func (o *Message) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *Message) GetIsDraftOk() (bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDraft

`func (o *Message) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### SetIsDraft

`func (o *Message) SetIsDraft(v bool)`

SetIsDraft gets a reference to the given bool and assigns it to the IsDraft field.

### SetIsDraftExplicitNull

`func (o *Message) SetIsDraftExplicitNull(b bool)`

SetIsDraftExplicitNull (un)sets IsDraft to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDraft value is set to nil even if false is passed
### GetWebLink

`func (o *Message) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *Message) GetWebLinkOk() (string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebLink

`func (o *Message) HasWebLink() bool`

HasWebLink returns a boolean if a field has been set.

### SetWebLink

`func (o *Message) SetWebLink(v string)`

SetWebLink gets a reference to the given string and assigns it to the WebLink field.

### SetWebLinkExplicitNull

`func (o *Message) SetWebLinkExplicitNull(b bool)`

SetWebLinkExplicitNull (un)sets WebLink to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebLink value is set to nil even if false is passed
### GetInferenceClassification

`func (o *Message) GetInferenceClassification() AnyOfmicrosoftGraphInferenceClassificationType`

GetInferenceClassification returns the InferenceClassification field if non-nil, zero value otherwise.

### GetInferenceClassificationOk

`func (o *Message) GetInferenceClassificationOk() (AnyOfmicrosoftGraphInferenceClassificationType, bool)`

GetInferenceClassificationOk returns a tuple with the InferenceClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInferenceClassification

`func (o *Message) HasInferenceClassification() bool`

HasInferenceClassification returns a boolean if a field has been set.

### SetInferenceClassification

`func (o *Message) SetInferenceClassification(v AnyOfmicrosoftGraphInferenceClassificationType)`

SetInferenceClassification gets a reference to the given AnyOfmicrosoftGraphInferenceClassificationType and assigns it to the InferenceClassification field.

### SetInferenceClassificationExplicitNull

`func (o *Message) SetInferenceClassificationExplicitNull(b bool)`

SetInferenceClassificationExplicitNull (un)sets InferenceClassification to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InferenceClassification value is set to nil even if false is passed
### GetFlag

`func (o *Message) GetFlag() AnyOfmicrosoftGraphFollowupFlag`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *Message) GetFlagOk() (AnyOfmicrosoftGraphFollowupFlag, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlag

`func (o *Message) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### SetFlag

`func (o *Message) SetFlag(v AnyOfmicrosoftGraphFollowupFlag)`

SetFlag gets a reference to the given AnyOfmicrosoftGraphFollowupFlag and assigns it to the Flag field.

### SetFlagExplicitNull

`func (o *Message) SetFlagExplicitNull(b bool)`

SetFlagExplicitNull (un)sets Flag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Flag value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *Message) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *Message) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *Message) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *Message) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *Message) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *Message) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *Message) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *Message) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetAttachments

`func (o *Message) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Message) GetAttachmentsOk() ([]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttachments

`func (o *Message) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachments

`func (o *Message) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments gets a reference to the given []MicrosoftGraphAttachment and assigns it to the Attachments field.

### GetExtensions

`func (o *Message) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Message) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *Message) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *Message) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



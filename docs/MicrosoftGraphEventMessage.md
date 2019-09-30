# MicrosoftGraphEventMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ChangeKey** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
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
**MeetingMessageType** | Pointer to [**AnyOfmicrosoftGraphMeetingMessageType**](anyOf&lt;microsoft.graph.meetingMessageType&gt;.md) |  | [optional] 
**Event** | Pointer to [**AnyOfmicrosoftGraphEvent**](anyOf&lt;microsoft.graph.event&gt;.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphEventMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEventMessage) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphEventMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphEventMessage) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphEventMessage) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphEventMessage) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphEventMessage) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphEventMessage) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphEventMessage) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphEventMessage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphEventMessage) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphEventMessage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphEventMessage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphEventMessage) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetChangeKey

`func (o *MicrosoftGraphEventMessage) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphEventMessage) GetChangeKeyOk() (string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeKey

`func (o *MicrosoftGraphEventMessage) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKey

`func (o *MicrosoftGraphEventMessage) SetChangeKey(v string)`

SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.

### SetChangeKeyExplicitNull

`func (o *MicrosoftGraphEventMessage) SetChangeKeyExplicitNull(b bool)`

SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChangeKey value is set to nil even if false is passed
### GetCategories

`func (o *MicrosoftGraphEventMessage) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphEventMessage) GetCategoriesOk() ([]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphEventMessage) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphEventMessage) SetCategories(v []string)`

SetCategories gets a reference to the given []string and assigns it to the Categories field.

### GetReceivedDateTime

`func (o *MicrosoftGraphEventMessage) GetReceivedDateTime() time.Time`

GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.

### GetReceivedDateTimeOk

`func (o *MicrosoftGraphEventMessage) GetReceivedDateTimeOk() (time.Time, bool)`

GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReceivedDateTime

`func (o *MicrosoftGraphEventMessage) HasReceivedDateTime() bool`

HasReceivedDateTime returns a boolean if a field has been set.

### SetReceivedDateTime

`func (o *MicrosoftGraphEventMessage) SetReceivedDateTime(v time.Time)`

SetReceivedDateTime gets a reference to the given time.Time and assigns it to the ReceivedDateTime field.

### SetReceivedDateTimeExplicitNull

`func (o *MicrosoftGraphEventMessage) SetReceivedDateTimeExplicitNull(b bool)`

SetReceivedDateTimeExplicitNull (un)sets ReceivedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReceivedDateTime value is set to nil even if false is passed
### GetSentDateTime

`func (o *MicrosoftGraphEventMessage) GetSentDateTime() time.Time`

GetSentDateTime returns the SentDateTime field if non-nil, zero value otherwise.

### GetSentDateTimeOk

`func (o *MicrosoftGraphEventMessage) GetSentDateTimeOk() (time.Time, bool)`

GetSentDateTimeOk returns a tuple with the SentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSentDateTime

`func (o *MicrosoftGraphEventMessage) HasSentDateTime() bool`

HasSentDateTime returns a boolean if a field has been set.

### SetSentDateTime

`func (o *MicrosoftGraphEventMessage) SetSentDateTime(v time.Time)`

SetSentDateTime gets a reference to the given time.Time and assigns it to the SentDateTime field.

### SetSentDateTimeExplicitNull

`func (o *MicrosoftGraphEventMessage) SetSentDateTimeExplicitNull(b bool)`

SetSentDateTimeExplicitNull (un)sets SentDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SentDateTime value is set to nil even if false is passed
### GetHasAttachments

`func (o *MicrosoftGraphEventMessage) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphEventMessage) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *MicrosoftGraphEventMessage) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *MicrosoftGraphEventMessage) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### SetHasAttachmentsExplicitNull

`func (o *MicrosoftGraphEventMessage) SetHasAttachmentsExplicitNull(b bool)`

SetHasAttachmentsExplicitNull (un)sets HasAttachments to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasAttachments value is set to nil even if false is passed
### GetInternetMessageId

`func (o *MicrosoftGraphEventMessage) GetInternetMessageId() string`

GetInternetMessageId returns the InternetMessageId field if non-nil, zero value otherwise.

### GetInternetMessageIdOk

`func (o *MicrosoftGraphEventMessage) GetInternetMessageIdOk() (string, bool)`

GetInternetMessageIdOk returns a tuple with the InternetMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetMessageId

`func (o *MicrosoftGraphEventMessage) HasInternetMessageId() bool`

HasInternetMessageId returns a boolean if a field has been set.

### SetInternetMessageId

`func (o *MicrosoftGraphEventMessage) SetInternetMessageId(v string)`

SetInternetMessageId gets a reference to the given string and assigns it to the InternetMessageId field.

### SetInternetMessageIdExplicitNull

`func (o *MicrosoftGraphEventMessage) SetInternetMessageIdExplicitNull(b bool)`

SetInternetMessageIdExplicitNull (un)sets InternetMessageId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InternetMessageId value is set to nil even if false is passed
### GetInternetMessageHeaders

`func (o *MicrosoftGraphEventMessage) GetInternetMessageHeaders() []AnyOfmicrosoftGraphInternetMessageHeader`

GetInternetMessageHeaders returns the InternetMessageHeaders field if non-nil, zero value otherwise.

### GetInternetMessageHeadersOk

`func (o *MicrosoftGraphEventMessage) GetInternetMessageHeadersOk() ([]AnyOfmicrosoftGraphInternetMessageHeader, bool)`

GetInternetMessageHeadersOk returns a tuple with the InternetMessageHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetMessageHeaders

`func (o *MicrosoftGraphEventMessage) HasInternetMessageHeaders() bool`

HasInternetMessageHeaders returns a boolean if a field has been set.

### SetInternetMessageHeaders

`func (o *MicrosoftGraphEventMessage) SetInternetMessageHeaders(v []AnyOfmicrosoftGraphInternetMessageHeader)`

SetInternetMessageHeaders gets a reference to the given []AnyOfmicrosoftGraphInternetMessageHeader and assigns it to the InternetMessageHeaders field.

### GetSubject

`func (o *MicrosoftGraphEventMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MicrosoftGraphEventMessage) GetSubjectOk() (string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubject

`func (o *MicrosoftGraphEventMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubject

`func (o *MicrosoftGraphEventMessage) SetSubject(v string)`

SetSubject gets a reference to the given string and assigns it to the Subject field.

### SetSubjectExplicitNull

`func (o *MicrosoftGraphEventMessage) SetSubjectExplicitNull(b bool)`

SetSubjectExplicitNull (un)sets Subject to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Subject value is set to nil even if false is passed
### GetBody

`func (o *MicrosoftGraphEventMessage) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MicrosoftGraphEventMessage) GetBodyOk() (AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *MicrosoftGraphEventMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *MicrosoftGraphEventMessage) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the Body field.

### SetBodyExplicitNull

`func (o *MicrosoftGraphEventMessage) SetBodyExplicitNull(b bool)`

SetBodyExplicitNull (un)sets Body to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Body value is set to nil even if false is passed
### GetBodyPreview

`func (o *MicrosoftGraphEventMessage) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *MicrosoftGraphEventMessage) GetBodyPreviewOk() (string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBodyPreview

`func (o *MicrosoftGraphEventMessage) HasBodyPreview() bool`

HasBodyPreview returns a boolean if a field has been set.

### SetBodyPreview

`func (o *MicrosoftGraphEventMessage) SetBodyPreview(v string)`

SetBodyPreview gets a reference to the given string and assigns it to the BodyPreview field.

### SetBodyPreviewExplicitNull

`func (o *MicrosoftGraphEventMessage) SetBodyPreviewExplicitNull(b bool)`

SetBodyPreviewExplicitNull (un)sets BodyPreview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BodyPreview value is set to nil even if false is passed
### GetImportance

`func (o *MicrosoftGraphEventMessage) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *MicrosoftGraphEventMessage) GetImportanceOk() (AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportance

`func (o *MicrosoftGraphEventMessage) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportance

`func (o *MicrosoftGraphEventMessage) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance gets a reference to the given AnyOfmicrosoftGraphImportance and assigns it to the Importance field.

### SetImportanceExplicitNull

`func (o *MicrosoftGraphEventMessage) SetImportanceExplicitNull(b bool)`

SetImportanceExplicitNull (un)sets Importance to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Importance value is set to nil even if false is passed
### GetParentFolderId

`func (o *MicrosoftGraphEventMessage) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MicrosoftGraphEventMessage) GetParentFolderIdOk() (string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentFolderId

`func (o *MicrosoftGraphEventMessage) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderId

`func (o *MicrosoftGraphEventMessage) SetParentFolderId(v string)`

SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.

### SetParentFolderIdExplicitNull

`func (o *MicrosoftGraphEventMessage) SetParentFolderIdExplicitNull(b bool)`

SetParentFolderIdExplicitNull (un)sets ParentFolderId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentFolderId value is set to nil even if false is passed
### GetSender

`func (o *MicrosoftGraphEventMessage) GetSender() AnyOfmicrosoftGraphRecipient`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MicrosoftGraphEventMessage) GetSenderOk() (AnyOfmicrosoftGraphRecipient, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSender

`func (o *MicrosoftGraphEventMessage) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSender

`func (o *MicrosoftGraphEventMessage) SetSender(v AnyOfmicrosoftGraphRecipient)`

SetSender gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the Sender field.

### SetSenderExplicitNull

`func (o *MicrosoftGraphEventMessage) SetSenderExplicitNull(b bool)`

SetSenderExplicitNull (un)sets Sender to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sender value is set to nil even if false is passed
### GetFrom

`func (o *MicrosoftGraphEventMessage) GetFrom() AnyOfmicrosoftGraphRecipient`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MicrosoftGraphEventMessage) GetFromOk() (AnyOfmicrosoftGraphRecipient, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrom

`func (o *MicrosoftGraphEventMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFrom

`func (o *MicrosoftGraphEventMessage) SetFrom(v AnyOfmicrosoftGraphRecipient)`

SetFrom gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the From field.

### SetFromExplicitNull

`func (o *MicrosoftGraphEventMessage) SetFromExplicitNull(b bool)`

SetFromExplicitNull (un)sets From to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The From value is set to nil even if false is passed
### GetToRecipients

`func (o *MicrosoftGraphEventMessage) GetToRecipients() []AnyOfmicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *MicrosoftGraphEventMessage) GetToRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToRecipients

`func (o *MicrosoftGraphEventMessage) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### SetToRecipients

`func (o *MicrosoftGraphEventMessage) SetToRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetToRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ToRecipients field.

### GetCcRecipients

`func (o *MicrosoftGraphEventMessage) GetCcRecipients() []AnyOfmicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *MicrosoftGraphEventMessage) GetCcRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCcRecipients

`func (o *MicrosoftGraphEventMessage) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### SetCcRecipients

`func (o *MicrosoftGraphEventMessage) SetCcRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetCcRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the CcRecipients field.

### GetBccRecipients

`func (o *MicrosoftGraphEventMessage) GetBccRecipients() []AnyOfmicrosoftGraphRecipient`

GetBccRecipients returns the BccRecipients field if non-nil, zero value otherwise.

### GetBccRecipientsOk

`func (o *MicrosoftGraphEventMessage) GetBccRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetBccRecipientsOk returns a tuple with the BccRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBccRecipients

`func (o *MicrosoftGraphEventMessage) HasBccRecipients() bool`

HasBccRecipients returns a boolean if a field has been set.

### SetBccRecipients

`func (o *MicrosoftGraphEventMessage) SetBccRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetBccRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the BccRecipients field.

### GetReplyTo

`func (o *MicrosoftGraphEventMessage) GetReplyTo() []AnyOfmicrosoftGraphRecipient`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MicrosoftGraphEventMessage) GetReplyToOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReplyTo

`func (o *MicrosoftGraphEventMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### SetReplyTo

`func (o *MicrosoftGraphEventMessage) SetReplyTo(v []AnyOfmicrosoftGraphRecipient)`

SetReplyTo gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ReplyTo field.

### GetConversationId

`func (o *MicrosoftGraphEventMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *MicrosoftGraphEventMessage) GetConversationIdOk() (string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversationId

`func (o *MicrosoftGraphEventMessage) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationId

`func (o *MicrosoftGraphEventMessage) SetConversationId(v string)`

SetConversationId gets a reference to the given string and assigns it to the ConversationId field.

### SetConversationIdExplicitNull

`func (o *MicrosoftGraphEventMessage) SetConversationIdExplicitNull(b bool)`

SetConversationIdExplicitNull (un)sets ConversationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConversationId value is set to nil even if false is passed
### GetUniqueBody

`func (o *MicrosoftGraphEventMessage) GetUniqueBody() AnyOfmicrosoftGraphItemBody`

GetUniqueBody returns the UniqueBody field if non-nil, zero value otherwise.

### GetUniqueBodyOk

`func (o *MicrosoftGraphEventMessage) GetUniqueBodyOk() (AnyOfmicrosoftGraphItemBody, bool)`

GetUniqueBodyOk returns a tuple with the UniqueBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueBody

`func (o *MicrosoftGraphEventMessage) HasUniqueBody() bool`

HasUniqueBody returns a boolean if a field has been set.

### SetUniqueBody

`func (o *MicrosoftGraphEventMessage) SetUniqueBody(v AnyOfmicrosoftGraphItemBody)`

SetUniqueBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the UniqueBody field.

### SetUniqueBodyExplicitNull

`func (o *MicrosoftGraphEventMessage) SetUniqueBodyExplicitNull(b bool)`

SetUniqueBodyExplicitNull (un)sets UniqueBody to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UniqueBody value is set to nil even if false is passed
### GetIsDeliveryReceiptRequested

`func (o *MicrosoftGraphEventMessage) GetIsDeliveryReceiptRequested() bool`

GetIsDeliveryReceiptRequested returns the IsDeliveryReceiptRequested field if non-nil, zero value otherwise.

### GetIsDeliveryReceiptRequestedOk

`func (o *MicrosoftGraphEventMessage) GetIsDeliveryReceiptRequestedOk() (bool, bool)`

GetIsDeliveryReceiptRequestedOk returns a tuple with the IsDeliveryReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDeliveryReceiptRequested

`func (o *MicrosoftGraphEventMessage) HasIsDeliveryReceiptRequested() bool`

HasIsDeliveryReceiptRequested returns a boolean if a field has been set.

### SetIsDeliveryReceiptRequested

`func (o *MicrosoftGraphEventMessage) SetIsDeliveryReceiptRequested(v bool)`

SetIsDeliveryReceiptRequested gets a reference to the given bool and assigns it to the IsDeliveryReceiptRequested field.

### SetIsDeliveryReceiptRequestedExplicitNull

`func (o *MicrosoftGraphEventMessage) SetIsDeliveryReceiptRequestedExplicitNull(b bool)`

SetIsDeliveryReceiptRequestedExplicitNull (un)sets IsDeliveryReceiptRequested to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDeliveryReceiptRequested value is set to nil even if false is passed
### GetIsReadReceiptRequested

`func (o *MicrosoftGraphEventMessage) GetIsReadReceiptRequested() bool`

GetIsReadReceiptRequested returns the IsReadReceiptRequested field if non-nil, zero value otherwise.

### GetIsReadReceiptRequestedOk

`func (o *MicrosoftGraphEventMessage) GetIsReadReceiptRequestedOk() (bool, bool)`

GetIsReadReceiptRequestedOk returns a tuple with the IsReadReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReadReceiptRequested

`func (o *MicrosoftGraphEventMessage) HasIsReadReceiptRequested() bool`

HasIsReadReceiptRequested returns a boolean if a field has been set.

### SetIsReadReceiptRequested

`func (o *MicrosoftGraphEventMessage) SetIsReadReceiptRequested(v bool)`

SetIsReadReceiptRequested gets a reference to the given bool and assigns it to the IsReadReceiptRequested field.

### SetIsReadReceiptRequestedExplicitNull

`func (o *MicrosoftGraphEventMessage) SetIsReadReceiptRequestedExplicitNull(b bool)`

SetIsReadReceiptRequestedExplicitNull (un)sets IsReadReceiptRequested to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsReadReceiptRequested value is set to nil even if false is passed
### GetIsRead

`func (o *MicrosoftGraphEventMessage) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *MicrosoftGraphEventMessage) GetIsReadOk() (bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsRead

`func (o *MicrosoftGraphEventMessage) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### SetIsRead

`func (o *MicrosoftGraphEventMessage) SetIsRead(v bool)`

SetIsRead gets a reference to the given bool and assigns it to the IsRead field.

### SetIsReadExplicitNull

`func (o *MicrosoftGraphEventMessage) SetIsReadExplicitNull(b bool)`

SetIsReadExplicitNull (un)sets IsRead to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsRead value is set to nil even if false is passed
### GetIsDraft

`func (o *MicrosoftGraphEventMessage) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *MicrosoftGraphEventMessage) GetIsDraftOk() (bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDraft

`func (o *MicrosoftGraphEventMessage) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### SetIsDraft

`func (o *MicrosoftGraphEventMessage) SetIsDraft(v bool)`

SetIsDraft gets a reference to the given bool and assigns it to the IsDraft field.

### SetIsDraftExplicitNull

`func (o *MicrosoftGraphEventMessage) SetIsDraftExplicitNull(b bool)`

SetIsDraftExplicitNull (un)sets IsDraft to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDraft value is set to nil even if false is passed
### GetWebLink

`func (o *MicrosoftGraphEventMessage) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *MicrosoftGraphEventMessage) GetWebLinkOk() (string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebLink

`func (o *MicrosoftGraphEventMessage) HasWebLink() bool`

HasWebLink returns a boolean if a field has been set.

### SetWebLink

`func (o *MicrosoftGraphEventMessage) SetWebLink(v string)`

SetWebLink gets a reference to the given string and assigns it to the WebLink field.

### SetWebLinkExplicitNull

`func (o *MicrosoftGraphEventMessage) SetWebLinkExplicitNull(b bool)`

SetWebLinkExplicitNull (un)sets WebLink to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebLink value is set to nil even if false is passed
### GetInferenceClassification

`func (o *MicrosoftGraphEventMessage) GetInferenceClassification() AnyOfmicrosoftGraphInferenceClassificationType`

GetInferenceClassification returns the InferenceClassification field if non-nil, zero value otherwise.

### GetInferenceClassificationOk

`func (o *MicrosoftGraphEventMessage) GetInferenceClassificationOk() (AnyOfmicrosoftGraphInferenceClassificationType, bool)`

GetInferenceClassificationOk returns a tuple with the InferenceClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInferenceClassification

`func (o *MicrosoftGraphEventMessage) HasInferenceClassification() bool`

HasInferenceClassification returns a boolean if a field has been set.

### SetInferenceClassification

`func (o *MicrosoftGraphEventMessage) SetInferenceClassification(v AnyOfmicrosoftGraphInferenceClassificationType)`

SetInferenceClassification gets a reference to the given AnyOfmicrosoftGraphInferenceClassificationType and assigns it to the InferenceClassification field.

### SetInferenceClassificationExplicitNull

`func (o *MicrosoftGraphEventMessage) SetInferenceClassificationExplicitNull(b bool)`

SetInferenceClassificationExplicitNull (un)sets InferenceClassification to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InferenceClassification value is set to nil even if false is passed
### GetFlag

`func (o *MicrosoftGraphEventMessage) GetFlag() AnyOfmicrosoftGraphFollowupFlag`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *MicrosoftGraphEventMessage) GetFlagOk() (AnyOfmicrosoftGraphFollowupFlag, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlag

`func (o *MicrosoftGraphEventMessage) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### SetFlag

`func (o *MicrosoftGraphEventMessage) SetFlag(v AnyOfmicrosoftGraphFollowupFlag)`

SetFlag gets a reference to the given AnyOfmicrosoftGraphFollowupFlag and assigns it to the Flag field.

### SetFlagExplicitNull

`func (o *MicrosoftGraphEventMessage) SetFlagExplicitNull(b bool)`

SetFlagExplicitNull (un)sets Flag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Flag value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphEventMessage) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphEventMessage) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphEventMessage) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphEventMessage) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphEventMessage) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphEventMessage) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphEventMessage) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphEventMessage) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetAttachments

`func (o *MicrosoftGraphEventMessage) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MicrosoftGraphEventMessage) GetAttachmentsOk() ([]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttachments

`func (o *MicrosoftGraphEventMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachments

`func (o *MicrosoftGraphEventMessage) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments gets a reference to the given []MicrosoftGraphAttachment and assigns it to the Attachments field.

### GetExtensions

`func (o *MicrosoftGraphEventMessage) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphEventMessage) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *MicrosoftGraphEventMessage) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *MicrosoftGraphEventMessage) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.

### GetMeetingMessageType

`func (o *MicrosoftGraphEventMessage) GetMeetingMessageType() AnyOfmicrosoftGraphMeetingMessageType`

GetMeetingMessageType returns the MeetingMessageType field if non-nil, zero value otherwise.

### GetMeetingMessageTypeOk

`func (o *MicrosoftGraphEventMessage) GetMeetingMessageTypeOk() (AnyOfmicrosoftGraphMeetingMessageType, bool)`

GetMeetingMessageTypeOk returns a tuple with the MeetingMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMeetingMessageType

`func (o *MicrosoftGraphEventMessage) HasMeetingMessageType() bool`

HasMeetingMessageType returns a boolean if a field has been set.

### SetMeetingMessageType

`func (o *MicrosoftGraphEventMessage) SetMeetingMessageType(v AnyOfmicrosoftGraphMeetingMessageType)`

SetMeetingMessageType gets a reference to the given AnyOfmicrosoftGraphMeetingMessageType and assigns it to the MeetingMessageType field.

### SetMeetingMessageTypeExplicitNull

`func (o *MicrosoftGraphEventMessage) SetMeetingMessageTypeExplicitNull(b bool)`

SetMeetingMessageTypeExplicitNull (un)sets MeetingMessageType to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MeetingMessageType value is set to nil even if false is passed
### GetEvent

`func (o *MicrosoftGraphEventMessage) GetEvent() AnyOfmicrosoftGraphEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *MicrosoftGraphEventMessage) GetEventOk() (AnyOfmicrosoftGraphEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvent

`func (o *MicrosoftGraphEventMessage) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEvent

`func (o *MicrosoftGraphEventMessage) SetEvent(v AnyOfmicrosoftGraphEvent)`

SetEvent gets a reference to the given AnyOfmicrosoftGraphEvent and assigns it to the Event field.

### SetEventExplicitNull

`func (o *MicrosoftGraphEventMessage) SetEventExplicitNull(b bool)`

SetEventExplicitNull (un)sets Event to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Event value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



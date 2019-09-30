# MicrosoftGraphMessage

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

## Methods

### GetId

`func (o *MicrosoftGraphMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphMessage) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphMessage) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphMessage) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphMessage) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphMessage) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphMessage) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphMessage) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphMessage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphMessage) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphMessage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphMessage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphMessage) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetChangeKey

`func (o *MicrosoftGraphMessage) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphMessage) GetChangeKeyOk() (string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeKey

`func (o *MicrosoftGraphMessage) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKey

`func (o *MicrosoftGraphMessage) SetChangeKey(v string)`

SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.

### SetChangeKeyExplicitNull

`func (o *MicrosoftGraphMessage) SetChangeKeyExplicitNull(b bool)`

SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChangeKey value is set to nil even if false is passed
### GetCategories

`func (o *MicrosoftGraphMessage) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphMessage) GetCategoriesOk() ([]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphMessage) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphMessage) SetCategories(v []string)`

SetCategories gets a reference to the given []string and assigns it to the Categories field.

### GetReceivedDateTime

`func (o *MicrosoftGraphMessage) GetReceivedDateTime() time.Time`

GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.

### GetReceivedDateTimeOk

`func (o *MicrosoftGraphMessage) GetReceivedDateTimeOk() (time.Time, bool)`

GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReceivedDateTime

`func (o *MicrosoftGraphMessage) HasReceivedDateTime() bool`

HasReceivedDateTime returns a boolean if a field has been set.

### SetReceivedDateTime

`func (o *MicrosoftGraphMessage) SetReceivedDateTime(v time.Time)`

SetReceivedDateTime gets a reference to the given time.Time and assigns it to the ReceivedDateTime field.

### SetReceivedDateTimeExplicitNull

`func (o *MicrosoftGraphMessage) SetReceivedDateTimeExplicitNull(b bool)`

SetReceivedDateTimeExplicitNull (un)sets ReceivedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ReceivedDateTime value is set to nil even if false is passed
### GetSentDateTime

`func (o *MicrosoftGraphMessage) GetSentDateTime() time.Time`

GetSentDateTime returns the SentDateTime field if non-nil, zero value otherwise.

### GetSentDateTimeOk

`func (o *MicrosoftGraphMessage) GetSentDateTimeOk() (time.Time, bool)`

GetSentDateTimeOk returns a tuple with the SentDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSentDateTime

`func (o *MicrosoftGraphMessage) HasSentDateTime() bool`

HasSentDateTime returns a boolean if a field has been set.

### SetSentDateTime

`func (o *MicrosoftGraphMessage) SetSentDateTime(v time.Time)`

SetSentDateTime gets a reference to the given time.Time and assigns it to the SentDateTime field.

### SetSentDateTimeExplicitNull

`func (o *MicrosoftGraphMessage) SetSentDateTimeExplicitNull(b bool)`

SetSentDateTimeExplicitNull (un)sets SentDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SentDateTime value is set to nil even if false is passed
### GetHasAttachments

`func (o *MicrosoftGraphMessage) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphMessage) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *MicrosoftGraphMessage) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *MicrosoftGraphMessage) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### SetHasAttachmentsExplicitNull

`func (o *MicrosoftGraphMessage) SetHasAttachmentsExplicitNull(b bool)`

SetHasAttachmentsExplicitNull (un)sets HasAttachments to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasAttachments value is set to nil even if false is passed
### GetInternetMessageId

`func (o *MicrosoftGraphMessage) GetInternetMessageId() string`

GetInternetMessageId returns the InternetMessageId field if non-nil, zero value otherwise.

### GetInternetMessageIdOk

`func (o *MicrosoftGraphMessage) GetInternetMessageIdOk() (string, bool)`

GetInternetMessageIdOk returns a tuple with the InternetMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetMessageId

`func (o *MicrosoftGraphMessage) HasInternetMessageId() bool`

HasInternetMessageId returns a boolean if a field has been set.

### SetInternetMessageId

`func (o *MicrosoftGraphMessage) SetInternetMessageId(v string)`

SetInternetMessageId gets a reference to the given string and assigns it to the InternetMessageId field.

### SetInternetMessageIdExplicitNull

`func (o *MicrosoftGraphMessage) SetInternetMessageIdExplicitNull(b bool)`

SetInternetMessageIdExplicitNull (un)sets InternetMessageId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InternetMessageId value is set to nil even if false is passed
### GetInternetMessageHeaders

`func (o *MicrosoftGraphMessage) GetInternetMessageHeaders() []AnyOfmicrosoftGraphInternetMessageHeader`

GetInternetMessageHeaders returns the InternetMessageHeaders field if non-nil, zero value otherwise.

### GetInternetMessageHeadersOk

`func (o *MicrosoftGraphMessage) GetInternetMessageHeadersOk() ([]AnyOfmicrosoftGraphInternetMessageHeader, bool)`

GetInternetMessageHeadersOk returns a tuple with the InternetMessageHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetMessageHeaders

`func (o *MicrosoftGraphMessage) HasInternetMessageHeaders() bool`

HasInternetMessageHeaders returns a boolean if a field has been set.

### SetInternetMessageHeaders

`func (o *MicrosoftGraphMessage) SetInternetMessageHeaders(v []AnyOfmicrosoftGraphInternetMessageHeader)`

SetInternetMessageHeaders gets a reference to the given []AnyOfmicrosoftGraphInternetMessageHeader and assigns it to the InternetMessageHeaders field.

### GetSubject

`func (o *MicrosoftGraphMessage) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MicrosoftGraphMessage) GetSubjectOk() (string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubject

`func (o *MicrosoftGraphMessage) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubject

`func (o *MicrosoftGraphMessage) SetSubject(v string)`

SetSubject gets a reference to the given string and assigns it to the Subject field.

### SetSubjectExplicitNull

`func (o *MicrosoftGraphMessage) SetSubjectExplicitNull(b bool)`

SetSubjectExplicitNull (un)sets Subject to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Subject value is set to nil even if false is passed
### GetBody

`func (o *MicrosoftGraphMessage) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MicrosoftGraphMessage) GetBodyOk() (AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *MicrosoftGraphMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *MicrosoftGraphMessage) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the Body field.

### SetBodyExplicitNull

`func (o *MicrosoftGraphMessage) SetBodyExplicitNull(b bool)`

SetBodyExplicitNull (un)sets Body to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Body value is set to nil even if false is passed
### GetBodyPreview

`func (o *MicrosoftGraphMessage) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *MicrosoftGraphMessage) GetBodyPreviewOk() (string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBodyPreview

`func (o *MicrosoftGraphMessage) HasBodyPreview() bool`

HasBodyPreview returns a boolean if a field has been set.

### SetBodyPreview

`func (o *MicrosoftGraphMessage) SetBodyPreview(v string)`

SetBodyPreview gets a reference to the given string and assigns it to the BodyPreview field.

### SetBodyPreviewExplicitNull

`func (o *MicrosoftGraphMessage) SetBodyPreviewExplicitNull(b bool)`

SetBodyPreviewExplicitNull (un)sets BodyPreview to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The BodyPreview value is set to nil even if false is passed
### GetImportance

`func (o *MicrosoftGraphMessage) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *MicrosoftGraphMessage) GetImportanceOk() (AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportance

`func (o *MicrosoftGraphMessage) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportance

`func (o *MicrosoftGraphMessage) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance gets a reference to the given AnyOfmicrosoftGraphImportance and assigns it to the Importance field.

### SetImportanceExplicitNull

`func (o *MicrosoftGraphMessage) SetImportanceExplicitNull(b bool)`

SetImportanceExplicitNull (un)sets Importance to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Importance value is set to nil even if false is passed
### GetParentFolderId

`func (o *MicrosoftGraphMessage) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MicrosoftGraphMessage) GetParentFolderIdOk() (string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParentFolderId

`func (o *MicrosoftGraphMessage) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderId

`func (o *MicrosoftGraphMessage) SetParentFolderId(v string)`

SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.

### SetParentFolderIdExplicitNull

`func (o *MicrosoftGraphMessage) SetParentFolderIdExplicitNull(b bool)`

SetParentFolderIdExplicitNull (un)sets ParentFolderId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ParentFolderId value is set to nil even if false is passed
### GetSender

`func (o *MicrosoftGraphMessage) GetSender() AnyOfmicrosoftGraphRecipient`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MicrosoftGraphMessage) GetSenderOk() (AnyOfmicrosoftGraphRecipient, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSender

`func (o *MicrosoftGraphMessage) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSender

`func (o *MicrosoftGraphMessage) SetSender(v AnyOfmicrosoftGraphRecipient)`

SetSender gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the Sender field.

### SetSenderExplicitNull

`func (o *MicrosoftGraphMessage) SetSenderExplicitNull(b bool)`

SetSenderExplicitNull (un)sets Sender to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sender value is set to nil even if false is passed
### GetFrom

`func (o *MicrosoftGraphMessage) GetFrom() AnyOfmicrosoftGraphRecipient`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MicrosoftGraphMessage) GetFromOk() (AnyOfmicrosoftGraphRecipient, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrom

`func (o *MicrosoftGraphMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFrom

`func (o *MicrosoftGraphMessage) SetFrom(v AnyOfmicrosoftGraphRecipient)`

SetFrom gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the From field.

### SetFromExplicitNull

`func (o *MicrosoftGraphMessage) SetFromExplicitNull(b bool)`

SetFromExplicitNull (un)sets From to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The From value is set to nil even if false is passed
### GetToRecipients

`func (o *MicrosoftGraphMessage) GetToRecipients() []AnyOfmicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *MicrosoftGraphMessage) GetToRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToRecipients

`func (o *MicrosoftGraphMessage) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### SetToRecipients

`func (o *MicrosoftGraphMessage) SetToRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetToRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ToRecipients field.

### GetCcRecipients

`func (o *MicrosoftGraphMessage) GetCcRecipients() []AnyOfmicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *MicrosoftGraphMessage) GetCcRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCcRecipients

`func (o *MicrosoftGraphMessage) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### SetCcRecipients

`func (o *MicrosoftGraphMessage) SetCcRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetCcRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the CcRecipients field.

### GetBccRecipients

`func (o *MicrosoftGraphMessage) GetBccRecipients() []AnyOfmicrosoftGraphRecipient`

GetBccRecipients returns the BccRecipients field if non-nil, zero value otherwise.

### GetBccRecipientsOk

`func (o *MicrosoftGraphMessage) GetBccRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetBccRecipientsOk returns a tuple with the BccRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBccRecipients

`func (o *MicrosoftGraphMessage) HasBccRecipients() bool`

HasBccRecipients returns a boolean if a field has been set.

### SetBccRecipients

`func (o *MicrosoftGraphMessage) SetBccRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetBccRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the BccRecipients field.

### GetReplyTo

`func (o *MicrosoftGraphMessage) GetReplyTo() []AnyOfmicrosoftGraphRecipient`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *MicrosoftGraphMessage) GetReplyToOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReplyTo

`func (o *MicrosoftGraphMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### SetReplyTo

`func (o *MicrosoftGraphMessage) SetReplyTo(v []AnyOfmicrosoftGraphRecipient)`

SetReplyTo gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ReplyTo field.

### GetConversationId

`func (o *MicrosoftGraphMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *MicrosoftGraphMessage) GetConversationIdOk() (string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversationId

`func (o *MicrosoftGraphMessage) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationId

`func (o *MicrosoftGraphMessage) SetConversationId(v string)`

SetConversationId gets a reference to the given string and assigns it to the ConversationId field.

### SetConversationIdExplicitNull

`func (o *MicrosoftGraphMessage) SetConversationIdExplicitNull(b bool)`

SetConversationIdExplicitNull (un)sets ConversationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConversationId value is set to nil even if false is passed
### GetUniqueBody

`func (o *MicrosoftGraphMessage) GetUniqueBody() AnyOfmicrosoftGraphItemBody`

GetUniqueBody returns the UniqueBody field if non-nil, zero value otherwise.

### GetUniqueBodyOk

`func (o *MicrosoftGraphMessage) GetUniqueBodyOk() (AnyOfmicrosoftGraphItemBody, bool)`

GetUniqueBodyOk returns a tuple with the UniqueBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueBody

`func (o *MicrosoftGraphMessage) HasUniqueBody() bool`

HasUniqueBody returns a boolean if a field has been set.

### SetUniqueBody

`func (o *MicrosoftGraphMessage) SetUniqueBody(v AnyOfmicrosoftGraphItemBody)`

SetUniqueBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the UniqueBody field.

### SetUniqueBodyExplicitNull

`func (o *MicrosoftGraphMessage) SetUniqueBodyExplicitNull(b bool)`

SetUniqueBodyExplicitNull (un)sets UniqueBody to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The UniqueBody value is set to nil even if false is passed
### GetIsDeliveryReceiptRequested

`func (o *MicrosoftGraphMessage) GetIsDeliveryReceiptRequested() bool`

GetIsDeliveryReceiptRequested returns the IsDeliveryReceiptRequested field if non-nil, zero value otherwise.

### GetIsDeliveryReceiptRequestedOk

`func (o *MicrosoftGraphMessage) GetIsDeliveryReceiptRequestedOk() (bool, bool)`

GetIsDeliveryReceiptRequestedOk returns a tuple with the IsDeliveryReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDeliveryReceiptRequested

`func (o *MicrosoftGraphMessage) HasIsDeliveryReceiptRequested() bool`

HasIsDeliveryReceiptRequested returns a boolean if a field has been set.

### SetIsDeliveryReceiptRequested

`func (o *MicrosoftGraphMessage) SetIsDeliveryReceiptRequested(v bool)`

SetIsDeliveryReceiptRequested gets a reference to the given bool and assigns it to the IsDeliveryReceiptRequested field.

### SetIsDeliveryReceiptRequestedExplicitNull

`func (o *MicrosoftGraphMessage) SetIsDeliveryReceiptRequestedExplicitNull(b bool)`

SetIsDeliveryReceiptRequestedExplicitNull (un)sets IsDeliveryReceiptRequested to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDeliveryReceiptRequested value is set to nil even if false is passed
### GetIsReadReceiptRequested

`func (o *MicrosoftGraphMessage) GetIsReadReceiptRequested() bool`

GetIsReadReceiptRequested returns the IsReadReceiptRequested field if non-nil, zero value otherwise.

### GetIsReadReceiptRequestedOk

`func (o *MicrosoftGraphMessage) GetIsReadReceiptRequestedOk() (bool, bool)`

GetIsReadReceiptRequestedOk returns a tuple with the IsReadReceiptRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReadReceiptRequested

`func (o *MicrosoftGraphMessage) HasIsReadReceiptRequested() bool`

HasIsReadReceiptRequested returns a boolean if a field has been set.

### SetIsReadReceiptRequested

`func (o *MicrosoftGraphMessage) SetIsReadReceiptRequested(v bool)`

SetIsReadReceiptRequested gets a reference to the given bool and assigns it to the IsReadReceiptRequested field.

### SetIsReadReceiptRequestedExplicitNull

`func (o *MicrosoftGraphMessage) SetIsReadReceiptRequestedExplicitNull(b bool)`

SetIsReadReceiptRequestedExplicitNull (un)sets IsReadReceiptRequested to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsReadReceiptRequested value is set to nil even if false is passed
### GetIsRead

`func (o *MicrosoftGraphMessage) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *MicrosoftGraphMessage) GetIsReadOk() (bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsRead

`func (o *MicrosoftGraphMessage) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### SetIsRead

`func (o *MicrosoftGraphMessage) SetIsRead(v bool)`

SetIsRead gets a reference to the given bool and assigns it to the IsRead field.

### SetIsReadExplicitNull

`func (o *MicrosoftGraphMessage) SetIsReadExplicitNull(b bool)`

SetIsReadExplicitNull (un)sets IsRead to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsRead value is set to nil even if false is passed
### GetIsDraft

`func (o *MicrosoftGraphMessage) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *MicrosoftGraphMessage) GetIsDraftOk() (bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsDraft

`func (o *MicrosoftGraphMessage) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### SetIsDraft

`func (o *MicrosoftGraphMessage) SetIsDraft(v bool)`

SetIsDraft gets a reference to the given bool and assigns it to the IsDraft field.

### SetIsDraftExplicitNull

`func (o *MicrosoftGraphMessage) SetIsDraftExplicitNull(b bool)`

SetIsDraftExplicitNull (un)sets IsDraft to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsDraft value is set to nil even if false is passed
### GetWebLink

`func (o *MicrosoftGraphMessage) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *MicrosoftGraphMessage) GetWebLinkOk() (string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebLink

`func (o *MicrosoftGraphMessage) HasWebLink() bool`

HasWebLink returns a boolean if a field has been set.

### SetWebLink

`func (o *MicrosoftGraphMessage) SetWebLink(v string)`

SetWebLink gets a reference to the given string and assigns it to the WebLink field.

### SetWebLinkExplicitNull

`func (o *MicrosoftGraphMessage) SetWebLinkExplicitNull(b bool)`

SetWebLinkExplicitNull (un)sets WebLink to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WebLink value is set to nil even if false is passed
### GetInferenceClassification

`func (o *MicrosoftGraphMessage) GetInferenceClassification() AnyOfmicrosoftGraphInferenceClassificationType`

GetInferenceClassification returns the InferenceClassification field if non-nil, zero value otherwise.

### GetInferenceClassificationOk

`func (o *MicrosoftGraphMessage) GetInferenceClassificationOk() (AnyOfmicrosoftGraphInferenceClassificationType, bool)`

GetInferenceClassificationOk returns a tuple with the InferenceClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInferenceClassification

`func (o *MicrosoftGraphMessage) HasInferenceClassification() bool`

HasInferenceClassification returns a boolean if a field has been set.

### SetInferenceClassification

`func (o *MicrosoftGraphMessage) SetInferenceClassification(v AnyOfmicrosoftGraphInferenceClassificationType)`

SetInferenceClassification gets a reference to the given AnyOfmicrosoftGraphInferenceClassificationType and assigns it to the InferenceClassification field.

### SetInferenceClassificationExplicitNull

`func (o *MicrosoftGraphMessage) SetInferenceClassificationExplicitNull(b bool)`

SetInferenceClassificationExplicitNull (un)sets InferenceClassification to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InferenceClassification value is set to nil even if false is passed
### GetFlag

`func (o *MicrosoftGraphMessage) GetFlag() AnyOfmicrosoftGraphFollowupFlag`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *MicrosoftGraphMessage) GetFlagOk() (AnyOfmicrosoftGraphFollowupFlag, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlag

`func (o *MicrosoftGraphMessage) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### SetFlag

`func (o *MicrosoftGraphMessage) SetFlag(v AnyOfmicrosoftGraphFollowupFlag)`

SetFlag gets a reference to the given AnyOfmicrosoftGraphFollowupFlag and assigns it to the Flag field.

### SetFlagExplicitNull

`func (o *MicrosoftGraphMessage) SetFlagExplicitNull(b bool)`

SetFlagExplicitNull (un)sets Flag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Flag value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphMessage) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphMessage) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphMessage) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphMessage) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphMessage) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphMessage) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphMessage) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphMessage) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetAttachments

`func (o *MicrosoftGraphMessage) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MicrosoftGraphMessage) GetAttachmentsOk() ([]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttachments

`func (o *MicrosoftGraphMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachments

`func (o *MicrosoftGraphMessage) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments gets a reference to the given []MicrosoftGraphAttachment and assigns it to the Attachments field.

### GetExtensions

`func (o *MicrosoftGraphMessage) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphMessage) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *MicrosoftGraphMessage) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *MicrosoftGraphMessage) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



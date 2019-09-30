# MicrosoftGraphPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastModifiedDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ChangeKey** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPost) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphPost) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphPost) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCreatedDateTime

`func (o *MicrosoftGraphPost) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphPost) GetCreatedDateTimeOk() (time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedDateTime

`func (o *MicrosoftGraphPost) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphPost) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.

### SetCreatedDateTimeExplicitNull

`func (o *MicrosoftGraphPost) SetCreatedDateTimeExplicitNull(b bool)`

SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CreatedDateTime value is set to nil even if false is passed
### GetLastModifiedDateTime

`func (o *MicrosoftGraphPost) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphPost) GetLastModifiedDateTimeOk() (time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphPost) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphPost) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.

### SetLastModifiedDateTimeExplicitNull

`func (o *MicrosoftGraphPost) SetLastModifiedDateTimeExplicitNull(b bool)`

SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LastModifiedDateTime value is set to nil even if false is passed
### GetChangeKey

`func (o *MicrosoftGraphPost) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *MicrosoftGraphPost) GetChangeKeyOk() (string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeKey

`func (o *MicrosoftGraphPost) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKey

`func (o *MicrosoftGraphPost) SetChangeKey(v string)`

SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.

### SetChangeKeyExplicitNull

`func (o *MicrosoftGraphPost) SetChangeKeyExplicitNull(b bool)`

SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChangeKey value is set to nil even if false is passed
### GetCategories

`func (o *MicrosoftGraphPost) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphPost) GetCategoriesOk() ([]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphPost) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphPost) SetCategories(v []string)`

SetCategories gets a reference to the given []string and assigns it to the Categories field.

### GetBody

`func (o *MicrosoftGraphPost) GetBody() AnyOfmicrosoftGraphItemBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MicrosoftGraphPost) GetBodyOk() (AnyOfmicrosoftGraphItemBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *MicrosoftGraphPost) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *MicrosoftGraphPost) SetBody(v AnyOfmicrosoftGraphItemBody)`

SetBody gets a reference to the given AnyOfmicrosoftGraphItemBody and assigns it to the Body field.

### SetBodyExplicitNull

`func (o *MicrosoftGraphPost) SetBodyExplicitNull(b bool)`

SetBodyExplicitNull (un)sets Body to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Body value is set to nil even if false is passed
### GetReceivedDateTime

`func (o *MicrosoftGraphPost) GetReceivedDateTime() time.Time`

GetReceivedDateTime returns the ReceivedDateTime field if non-nil, zero value otherwise.

### GetReceivedDateTimeOk

`func (o *MicrosoftGraphPost) GetReceivedDateTimeOk() (time.Time, bool)`

GetReceivedDateTimeOk returns a tuple with the ReceivedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReceivedDateTime

`func (o *MicrosoftGraphPost) HasReceivedDateTime() bool`

HasReceivedDateTime returns a boolean if a field has been set.

### SetReceivedDateTime

`func (o *MicrosoftGraphPost) SetReceivedDateTime(v time.Time)`

SetReceivedDateTime gets a reference to the given time.Time and assigns it to the ReceivedDateTime field.

### GetHasAttachments

`func (o *MicrosoftGraphPost) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphPost) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *MicrosoftGraphPost) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *MicrosoftGraphPost) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### GetFrom

`func (o *MicrosoftGraphPost) GetFrom() MicrosoftGraphRecipient`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MicrosoftGraphPost) GetFromOk() (MicrosoftGraphRecipient, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFrom

`func (o *MicrosoftGraphPost) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFrom

`func (o *MicrosoftGraphPost) SetFrom(v MicrosoftGraphRecipient)`

SetFrom gets a reference to the given MicrosoftGraphRecipient and assigns it to the From field.

### GetSender

`func (o *MicrosoftGraphPost) GetSender() AnyOfmicrosoftGraphRecipient`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MicrosoftGraphPost) GetSenderOk() (AnyOfmicrosoftGraphRecipient, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSender

`func (o *MicrosoftGraphPost) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSender

`func (o *MicrosoftGraphPost) SetSender(v AnyOfmicrosoftGraphRecipient)`

SetSender gets a reference to the given AnyOfmicrosoftGraphRecipient and assigns it to the Sender field.

### SetSenderExplicitNull

`func (o *MicrosoftGraphPost) SetSenderExplicitNull(b bool)`

SetSenderExplicitNull (un)sets Sender to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sender value is set to nil even if false is passed
### GetConversationThreadId

`func (o *MicrosoftGraphPost) GetConversationThreadId() string`

GetConversationThreadId returns the ConversationThreadId field if non-nil, zero value otherwise.

### GetConversationThreadIdOk

`func (o *MicrosoftGraphPost) GetConversationThreadIdOk() (string, bool)`

GetConversationThreadIdOk returns a tuple with the ConversationThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversationThreadId

`func (o *MicrosoftGraphPost) HasConversationThreadId() bool`

HasConversationThreadId returns a boolean if a field has been set.

### SetConversationThreadId

`func (o *MicrosoftGraphPost) SetConversationThreadId(v string)`

SetConversationThreadId gets a reference to the given string and assigns it to the ConversationThreadId field.

### SetConversationThreadIdExplicitNull

`func (o *MicrosoftGraphPost) SetConversationThreadIdExplicitNull(b bool)`

SetConversationThreadIdExplicitNull (un)sets ConversationThreadId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConversationThreadId value is set to nil even if false is passed
### GetNewParticipants

`func (o *MicrosoftGraphPost) GetNewParticipants() []MicrosoftGraphRecipient`

GetNewParticipants returns the NewParticipants field if non-nil, zero value otherwise.

### GetNewParticipantsOk

`func (o *MicrosoftGraphPost) GetNewParticipantsOk() ([]MicrosoftGraphRecipient, bool)`

GetNewParticipantsOk returns a tuple with the NewParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNewParticipants

`func (o *MicrosoftGraphPost) HasNewParticipants() bool`

HasNewParticipants returns a boolean if a field has been set.

### SetNewParticipants

`func (o *MicrosoftGraphPost) SetNewParticipants(v []MicrosoftGraphRecipient)`

SetNewParticipants gets a reference to the given []MicrosoftGraphRecipient and assigns it to the NewParticipants field.

### GetConversationId

`func (o *MicrosoftGraphPost) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *MicrosoftGraphPost) GetConversationIdOk() (string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConversationId

`func (o *MicrosoftGraphPost) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationId

`func (o *MicrosoftGraphPost) SetConversationId(v string)`

SetConversationId gets a reference to the given string and assigns it to the ConversationId field.

### SetConversationIdExplicitNull

`func (o *MicrosoftGraphPost) SetConversationIdExplicitNull(b bool)`

SetConversationIdExplicitNull (un)sets ConversationId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ConversationId value is set to nil even if false is passed
### GetInReplyTo

`func (o *MicrosoftGraphPost) GetInReplyTo() AnyOfmicrosoftGraphPost`

GetInReplyTo returns the InReplyTo field if non-nil, zero value otherwise.

### GetInReplyToOk

`func (o *MicrosoftGraphPost) GetInReplyToOk() (AnyOfmicrosoftGraphPost, bool)`

GetInReplyToOk returns a tuple with the InReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInReplyTo

`func (o *MicrosoftGraphPost) HasInReplyTo() bool`

HasInReplyTo returns a boolean if a field has been set.

### SetInReplyTo

`func (o *MicrosoftGraphPost) SetInReplyTo(v AnyOfmicrosoftGraphPost)`

SetInReplyTo gets a reference to the given AnyOfmicrosoftGraphPost and assigns it to the InReplyTo field.

### SetInReplyToExplicitNull

`func (o *MicrosoftGraphPost) SetInReplyToExplicitNull(b bool)`

SetInReplyToExplicitNull (un)sets InReplyTo to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The InReplyTo value is set to nil even if false is passed
### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphPost) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphPost) GetSingleValueExtendedPropertiesOk() ([]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphPost) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphPost) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphPost) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphPost) GetMultiValueExtendedPropertiesOk() ([]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphPost) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphPost) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.

### GetExtensions

`func (o *MicrosoftGraphPost) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphPost) GetExtensionsOk() ([]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtensions

`func (o *MicrosoftGraphPost) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensions

`func (o *MicrosoftGraphPost) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.

### GetAttachments

`func (o *MicrosoftGraphPost) GetAttachments() []MicrosoftGraphAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *MicrosoftGraphPost) GetAttachmentsOk() ([]MicrosoftGraphAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttachments

`func (o *MicrosoftGraphPost) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachments

`func (o *MicrosoftGraphPost) SetAttachments(v []MicrosoftGraphAttachment)`

SetAttachments gets a reference to the given []MicrosoftGraphAttachment and assigns it to the Attachments field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



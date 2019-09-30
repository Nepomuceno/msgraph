# MicrosoftGraphMessageRulePredicates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** |  | [optional] 
**SubjectContains** | Pointer to **[]string** |  | [optional] 
**BodyContains** | Pointer to **[]string** |  | [optional] 
**BodyOrSubjectContains** | Pointer to **[]string** |  | [optional] 
**SenderContains** | Pointer to **[]string** |  | [optional] 
**RecipientContains** | Pointer to **[]string** |  | [optional] 
**HeaderContains** | Pointer to **[]string** |  | [optional] 
**MessageActionFlag** | Pointer to [**AnyOfmicrosoftGraphMessageActionFlag**](anyOf&lt;microsoft.graph.messageActionFlag&gt;.md) |  | [optional] 
**Importance** | Pointer to [**AnyOfmicrosoftGraphImportance**](anyOf&lt;microsoft.graph.importance&gt;.md) |  | [optional] 
**Sensitivity** | Pointer to [**AnyOfmicrosoftGraphSensitivity**](anyOf&lt;microsoft.graph.sensitivity&gt;.md) |  | [optional] 
**FromAddresses** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**SentToAddresses** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**SentToMe** | Pointer to **bool** |  | [optional] 
**SentOnlyToMe** | Pointer to **bool** |  | [optional] 
**SentCcMe** | Pointer to **bool** |  | [optional] 
**SentToOrCcMe** | Pointer to **bool** |  | [optional] 
**NotSentToMe** | Pointer to **bool** |  | [optional] 
**HasAttachments** | Pointer to **bool** |  | [optional] 
**IsApprovalRequest** | Pointer to **bool** |  | [optional] 
**IsAutomaticForward** | Pointer to **bool** |  | [optional] 
**IsAutomaticReply** | Pointer to **bool** |  | [optional] 
**IsEncrypted** | Pointer to **bool** |  | [optional] 
**IsMeetingRequest** | Pointer to **bool** |  | [optional] 
**IsMeetingResponse** | Pointer to **bool** |  | [optional] 
**IsNonDeliveryReport** | Pointer to **bool** |  | [optional] 
**IsPermissionControlled** | Pointer to **bool** |  | [optional] 
**IsReadReceipt** | Pointer to **bool** |  | [optional] 
**IsSigned** | Pointer to **bool** |  | [optional] 
**IsVoicemail** | Pointer to **bool** |  | [optional] 
**WithinSizeRange** | Pointer to [**AnyOfmicrosoftGraphSizeRange**](anyOf&lt;microsoft.graph.sizeRange&gt;.md) |  | [optional] 

## Methods

### GetCategories

`func (o *MicrosoftGraphMessageRulePredicates) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphMessageRulePredicates) GetCategoriesOk() ([]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCategories

`func (o *MicrosoftGraphMessageRulePredicates) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategories

`func (o *MicrosoftGraphMessageRulePredicates) SetCategories(v []string)`

SetCategories gets a reference to the given []string and assigns it to the Categories field.

### GetSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) GetSubjectContains() []string`

GetSubjectContains returns the SubjectContains field if non-nil, zero value otherwise.

### GetSubjectContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSubjectContainsOk() ([]string, bool)`

GetSubjectContainsOk returns a tuple with the SubjectContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) HasSubjectContains() bool`

HasSubjectContains returns a boolean if a field has been set.

### SetSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) SetSubjectContains(v []string)`

SetSubjectContains gets a reference to the given []string and assigns it to the SubjectContains field.

### GetBodyContains

`func (o *MicrosoftGraphMessageRulePredicates) GetBodyContains() []string`

GetBodyContains returns the BodyContains field if non-nil, zero value otherwise.

### GetBodyContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetBodyContainsOk() ([]string, bool)`

GetBodyContainsOk returns a tuple with the BodyContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBodyContains

`func (o *MicrosoftGraphMessageRulePredicates) HasBodyContains() bool`

HasBodyContains returns a boolean if a field has been set.

### SetBodyContains

`func (o *MicrosoftGraphMessageRulePredicates) SetBodyContains(v []string)`

SetBodyContains gets a reference to the given []string and assigns it to the BodyContains field.

### GetBodyOrSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) GetBodyOrSubjectContains() []string`

GetBodyOrSubjectContains returns the BodyOrSubjectContains field if non-nil, zero value otherwise.

### GetBodyOrSubjectContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetBodyOrSubjectContainsOk() ([]string, bool)`

GetBodyOrSubjectContainsOk returns a tuple with the BodyOrSubjectContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBodyOrSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) HasBodyOrSubjectContains() bool`

HasBodyOrSubjectContains returns a boolean if a field has been set.

### SetBodyOrSubjectContains

`func (o *MicrosoftGraphMessageRulePredicates) SetBodyOrSubjectContains(v []string)`

SetBodyOrSubjectContains gets a reference to the given []string and assigns it to the BodyOrSubjectContains field.

### GetSenderContains

`func (o *MicrosoftGraphMessageRulePredicates) GetSenderContains() []string`

GetSenderContains returns the SenderContains field if non-nil, zero value otherwise.

### GetSenderContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSenderContainsOk() ([]string, bool)`

GetSenderContainsOk returns a tuple with the SenderContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSenderContains

`func (o *MicrosoftGraphMessageRulePredicates) HasSenderContains() bool`

HasSenderContains returns a boolean if a field has been set.

### SetSenderContains

`func (o *MicrosoftGraphMessageRulePredicates) SetSenderContains(v []string)`

SetSenderContains gets a reference to the given []string and assigns it to the SenderContains field.

### GetRecipientContains

`func (o *MicrosoftGraphMessageRulePredicates) GetRecipientContains() []string`

GetRecipientContains returns the RecipientContains field if non-nil, zero value otherwise.

### GetRecipientContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetRecipientContainsOk() ([]string, bool)`

GetRecipientContainsOk returns a tuple with the RecipientContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientContains

`func (o *MicrosoftGraphMessageRulePredicates) HasRecipientContains() bool`

HasRecipientContains returns a boolean if a field has been set.

### SetRecipientContains

`func (o *MicrosoftGraphMessageRulePredicates) SetRecipientContains(v []string)`

SetRecipientContains gets a reference to the given []string and assigns it to the RecipientContains field.

### GetHeaderContains

`func (o *MicrosoftGraphMessageRulePredicates) GetHeaderContains() []string`

GetHeaderContains returns the HeaderContains field if non-nil, zero value otherwise.

### GetHeaderContainsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetHeaderContainsOk() ([]string, bool)`

GetHeaderContainsOk returns a tuple with the HeaderContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeaderContains

`func (o *MicrosoftGraphMessageRulePredicates) HasHeaderContains() bool`

HasHeaderContains returns a boolean if a field has been set.

### SetHeaderContains

`func (o *MicrosoftGraphMessageRulePredicates) SetHeaderContains(v []string)`

SetHeaderContains gets a reference to the given []string and assigns it to the HeaderContains field.

### GetMessageActionFlag

`func (o *MicrosoftGraphMessageRulePredicates) GetMessageActionFlag() AnyOfmicrosoftGraphMessageActionFlag`

GetMessageActionFlag returns the MessageActionFlag field if non-nil, zero value otherwise.

### GetMessageActionFlagOk

`func (o *MicrosoftGraphMessageRulePredicates) GetMessageActionFlagOk() (AnyOfmicrosoftGraphMessageActionFlag, bool)`

GetMessageActionFlagOk returns a tuple with the MessageActionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessageActionFlag

`func (o *MicrosoftGraphMessageRulePredicates) HasMessageActionFlag() bool`

HasMessageActionFlag returns a boolean if a field has been set.

### SetMessageActionFlag

`func (o *MicrosoftGraphMessageRulePredicates) SetMessageActionFlag(v AnyOfmicrosoftGraphMessageActionFlag)`

SetMessageActionFlag gets a reference to the given AnyOfmicrosoftGraphMessageActionFlag and assigns it to the MessageActionFlag field.

### SetMessageActionFlagExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetMessageActionFlagExplicitNull(b bool)`

SetMessageActionFlagExplicitNull (un)sets MessageActionFlag to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MessageActionFlag value is set to nil even if false is passed
### GetImportance

`func (o *MicrosoftGraphMessageRulePredicates) GetImportance() AnyOfmicrosoftGraphImportance`

GetImportance returns the Importance field if non-nil, zero value otherwise.

### GetImportanceOk

`func (o *MicrosoftGraphMessageRulePredicates) GetImportanceOk() (AnyOfmicrosoftGraphImportance, bool)`

GetImportanceOk returns a tuple with the Importance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportance

`func (o *MicrosoftGraphMessageRulePredicates) HasImportance() bool`

HasImportance returns a boolean if a field has been set.

### SetImportance

`func (o *MicrosoftGraphMessageRulePredicates) SetImportance(v AnyOfmicrosoftGraphImportance)`

SetImportance gets a reference to the given AnyOfmicrosoftGraphImportance and assigns it to the Importance field.

### SetImportanceExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetImportanceExplicitNull(b bool)`

SetImportanceExplicitNull (un)sets Importance to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Importance value is set to nil even if false is passed
### GetSensitivity

`func (o *MicrosoftGraphMessageRulePredicates) GetSensitivity() AnyOfmicrosoftGraphSensitivity`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSensitivityOk() (AnyOfmicrosoftGraphSensitivity, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSensitivity

`func (o *MicrosoftGraphMessageRulePredicates) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.

### SetSensitivity

`func (o *MicrosoftGraphMessageRulePredicates) SetSensitivity(v AnyOfmicrosoftGraphSensitivity)`

SetSensitivity gets a reference to the given AnyOfmicrosoftGraphSensitivity and assigns it to the Sensitivity field.

### SetSensitivityExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetSensitivityExplicitNull(b bool)`

SetSensitivityExplicitNull (un)sets Sensitivity to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sensitivity value is set to nil even if false is passed
### GetFromAddresses

`func (o *MicrosoftGraphMessageRulePredicates) GetFromAddresses() []AnyOfmicrosoftGraphRecipient`

GetFromAddresses returns the FromAddresses field if non-nil, zero value otherwise.

### GetFromAddressesOk

`func (o *MicrosoftGraphMessageRulePredicates) GetFromAddressesOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetFromAddressesOk returns a tuple with the FromAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFromAddresses

`func (o *MicrosoftGraphMessageRulePredicates) HasFromAddresses() bool`

HasFromAddresses returns a boolean if a field has been set.

### SetFromAddresses

`func (o *MicrosoftGraphMessageRulePredicates) SetFromAddresses(v []AnyOfmicrosoftGraphRecipient)`

SetFromAddresses gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the FromAddresses field.

### GetSentToAddresses

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToAddresses() []AnyOfmicrosoftGraphRecipient`

GetSentToAddresses returns the SentToAddresses field if non-nil, zero value otherwise.

### GetSentToAddressesOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToAddressesOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetSentToAddressesOk returns a tuple with the SentToAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSentToAddresses

`func (o *MicrosoftGraphMessageRulePredicates) HasSentToAddresses() bool`

HasSentToAddresses returns a boolean if a field has been set.

### SetSentToAddresses

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToAddresses(v []AnyOfmicrosoftGraphRecipient)`

SetSentToAddresses gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the SentToAddresses field.

### GetSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToMe() bool`

GetSentToMe returns the SentToMe field if non-nil, zero value otherwise.

### GetSentToMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToMeOk() (bool, bool)`

GetSentToMeOk returns a tuple with the SentToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) HasSentToMe() bool`

HasSentToMe returns a boolean if a field has been set.

### SetSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToMe(v bool)`

SetSentToMe gets a reference to the given bool and assigns it to the SentToMe field.

### SetSentToMeExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToMeExplicitNull(b bool)`

SetSentToMeExplicitNull (un)sets SentToMe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SentToMe value is set to nil even if false is passed
### GetSentOnlyToMe

`func (o *MicrosoftGraphMessageRulePredicates) GetSentOnlyToMe() bool`

GetSentOnlyToMe returns the SentOnlyToMe field if non-nil, zero value otherwise.

### GetSentOnlyToMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentOnlyToMeOk() (bool, bool)`

GetSentOnlyToMeOk returns a tuple with the SentOnlyToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSentOnlyToMe

`func (o *MicrosoftGraphMessageRulePredicates) HasSentOnlyToMe() bool`

HasSentOnlyToMe returns a boolean if a field has been set.

### SetSentOnlyToMe

`func (o *MicrosoftGraphMessageRulePredicates) SetSentOnlyToMe(v bool)`

SetSentOnlyToMe gets a reference to the given bool and assigns it to the SentOnlyToMe field.

### SetSentOnlyToMeExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetSentOnlyToMeExplicitNull(b bool)`

SetSentOnlyToMeExplicitNull (un)sets SentOnlyToMe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SentOnlyToMe value is set to nil even if false is passed
### GetSentCcMe

`func (o *MicrosoftGraphMessageRulePredicates) GetSentCcMe() bool`

GetSentCcMe returns the SentCcMe field if non-nil, zero value otherwise.

### GetSentCcMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentCcMeOk() (bool, bool)`

GetSentCcMeOk returns a tuple with the SentCcMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSentCcMe

`func (o *MicrosoftGraphMessageRulePredicates) HasSentCcMe() bool`

HasSentCcMe returns a boolean if a field has been set.

### SetSentCcMe

`func (o *MicrosoftGraphMessageRulePredicates) SetSentCcMe(v bool)`

SetSentCcMe gets a reference to the given bool and assigns it to the SentCcMe field.

### SetSentCcMeExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetSentCcMeExplicitNull(b bool)`

SetSentCcMeExplicitNull (un)sets SentCcMe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SentCcMe value is set to nil even if false is passed
### GetSentToOrCcMe

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToOrCcMe() bool`

GetSentToOrCcMe returns the SentToOrCcMe field if non-nil, zero value otherwise.

### GetSentToOrCcMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetSentToOrCcMeOk() (bool, bool)`

GetSentToOrCcMeOk returns a tuple with the SentToOrCcMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSentToOrCcMe

`func (o *MicrosoftGraphMessageRulePredicates) HasSentToOrCcMe() bool`

HasSentToOrCcMe returns a boolean if a field has been set.

### SetSentToOrCcMe

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToOrCcMe(v bool)`

SetSentToOrCcMe gets a reference to the given bool and assigns it to the SentToOrCcMe field.

### SetSentToOrCcMeExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetSentToOrCcMeExplicitNull(b bool)`

SetSentToOrCcMeExplicitNull (un)sets SentToOrCcMe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SentToOrCcMe value is set to nil even if false is passed
### GetNotSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) GetNotSentToMe() bool`

GetNotSentToMe returns the NotSentToMe field if non-nil, zero value otherwise.

### GetNotSentToMeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetNotSentToMeOk() (bool, bool)`

GetNotSentToMeOk returns a tuple with the NotSentToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) HasNotSentToMe() bool`

HasNotSentToMe returns a boolean if a field has been set.

### SetNotSentToMe

`func (o *MicrosoftGraphMessageRulePredicates) SetNotSentToMe(v bool)`

SetNotSentToMe gets a reference to the given bool and assigns it to the NotSentToMe field.

### SetNotSentToMeExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetNotSentToMeExplicitNull(b bool)`

SetNotSentToMeExplicitNull (un)sets NotSentToMe to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NotSentToMe value is set to nil even if false is passed
### GetHasAttachments

`func (o *MicrosoftGraphMessageRulePredicates) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphMessageRulePredicates) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *MicrosoftGraphMessageRulePredicates) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *MicrosoftGraphMessageRulePredicates) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### SetHasAttachmentsExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetHasAttachmentsExplicitNull(b bool)`

SetHasAttachmentsExplicitNull (un)sets HasAttachments to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The HasAttachments value is set to nil even if false is passed
### GetIsApprovalRequest

`func (o *MicrosoftGraphMessageRulePredicates) GetIsApprovalRequest() bool`

GetIsApprovalRequest returns the IsApprovalRequest field if non-nil, zero value otherwise.

### GetIsApprovalRequestOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsApprovalRequestOk() (bool, bool)`

GetIsApprovalRequestOk returns a tuple with the IsApprovalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsApprovalRequest

`func (o *MicrosoftGraphMessageRulePredicates) HasIsApprovalRequest() bool`

HasIsApprovalRequest returns a boolean if a field has been set.

### SetIsApprovalRequest

`func (o *MicrosoftGraphMessageRulePredicates) SetIsApprovalRequest(v bool)`

SetIsApprovalRequest gets a reference to the given bool and assigns it to the IsApprovalRequest field.

### SetIsApprovalRequestExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsApprovalRequestExplicitNull(b bool)`

SetIsApprovalRequestExplicitNull (un)sets IsApprovalRequest to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsApprovalRequest value is set to nil even if false is passed
### GetIsAutomaticForward

`func (o *MicrosoftGraphMessageRulePredicates) GetIsAutomaticForward() bool`

GetIsAutomaticForward returns the IsAutomaticForward field if non-nil, zero value otherwise.

### GetIsAutomaticForwardOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsAutomaticForwardOk() (bool, bool)`

GetIsAutomaticForwardOk returns a tuple with the IsAutomaticForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAutomaticForward

`func (o *MicrosoftGraphMessageRulePredicates) HasIsAutomaticForward() bool`

HasIsAutomaticForward returns a boolean if a field has been set.

### SetIsAutomaticForward

`func (o *MicrosoftGraphMessageRulePredicates) SetIsAutomaticForward(v bool)`

SetIsAutomaticForward gets a reference to the given bool and assigns it to the IsAutomaticForward field.

### SetIsAutomaticForwardExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsAutomaticForwardExplicitNull(b bool)`

SetIsAutomaticForwardExplicitNull (un)sets IsAutomaticForward to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsAutomaticForward value is set to nil even if false is passed
### GetIsAutomaticReply

`func (o *MicrosoftGraphMessageRulePredicates) GetIsAutomaticReply() bool`

GetIsAutomaticReply returns the IsAutomaticReply field if non-nil, zero value otherwise.

### GetIsAutomaticReplyOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsAutomaticReplyOk() (bool, bool)`

GetIsAutomaticReplyOk returns a tuple with the IsAutomaticReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAutomaticReply

`func (o *MicrosoftGraphMessageRulePredicates) HasIsAutomaticReply() bool`

HasIsAutomaticReply returns a boolean if a field has been set.

### SetIsAutomaticReply

`func (o *MicrosoftGraphMessageRulePredicates) SetIsAutomaticReply(v bool)`

SetIsAutomaticReply gets a reference to the given bool and assigns it to the IsAutomaticReply field.

### SetIsAutomaticReplyExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsAutomaticReplyExplicitNull(b bool)`

SetIsAutomaticReplyExplicitNull (un)sets IsAutomaticReply to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsAutomaticReply value is set to nil even if false is passed
### GetIsEncrypted

`func (o *MicrosoftGraphMessageRulePredicates) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsEncryptedOk() (bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsEncrypted

`func (o *MicrosoftGraphMessageRulePredicates) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncrypted

`func (o *MicrosoftGraphMessageRulePredicates) SetIsEncrypted(v bool)`

SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.

### SetIsEncryptedExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsEncryptedExplicitNull(b bool)`

SetIsEncryptedExplicitNull (un)sets IsEncrypted to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsEncrypted value is set to nil even if false is passed
### GetIsMeetingRequest

`func (o *MicrosoftGraphMessageRulePredicates) GetIsMeetingRequest() bool`

GetIsMeetingRequest returns the IsMeetingRequest field if non-nil, zero value otherwise.

### GetIsMeetingRequestOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsMeetingRequestOk() (bool, bool)`

GetIsMeetingRequestOk returns a tuple with the IsMeetingRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsMeetingRequest

`func (o *MicrosoftGraphMessageRulePredicates) HasIsMeetingRequest() bool`

HasIsMeetingRequest returns a boolean if a field has been set.

### SetIsMeetingRequest

`func (o *MicrosoftGraphMessageRulePredicates) SetIsMeetingRequest(v bool)`

SetIsMeetingRequest gets a reference to the given bool and assigns it to the IsMeetingRequest field.

### SetIsMeetingRequestExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsMeetingRequestExplicitNull(b bool)`

SetIsMeetingRequestExplicitNull (un)sets IsMeetingRequest to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsMeetingRequest value is set to nil even if false is passed
### GetIsMeetingResponse

`func (o *MicrosoftGraphMessageRulePredicates) GetIsMeetingResponse() bool`

GetIsMeetingResponse returns the IsMeetingResponse field if non-nil, zero value otherwise.

### GetIsMeetingResponseOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsMeetingResponseOk() (bool, bool)`

GetIsMeetingResponseOk returns a tuple with the IsMeetingResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsMeetingResponse

`func (o *MicrosoftGraphMessageRulePredicates) HasIsMeetingResponse() bool`

HasIsMeetingResponse returns a boolean if a field has been set.

### SetIsMeetingResponse

`func (o *MicrosoftGraphMessageRulePredicates) SetIsMeetingResponse(v bool)`

SetIsMeetingResponse gets a reference to the given bool and assigns it to the IsMeetingResponse field.

### SetIsMeetingResponseExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsMeetingResponseExplicitNull(b bool)`

SetIsMeetingResponseExplicitNull (un)sets IsMeetingResponse to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsMeetingResponse value is set to nil even if false is passed
### GetIsNonDeliveryReport

`func (o *MicrosoftGraphMessageRulePredicates) GetIsNonDeliveryReport() bool`

GetIsNonDeliveryReport returns the IsNonDeliveryReport field if non-nil, zero value otherwise.

### GetIsNonDeliveryReportOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsNonDeliveryReportOk() (bool, bool)`

GetIsNonDeliveryReportOk returns a tuple with the IsNonDeliveryReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsNonDeliveryReport

`func (o *MicrosoftGraphMessageRulePredicates) HasIsNonDeliveryReport() bool`

HasIsNonDeliveryReport returns a boolean if a field has been set.

### SetIsNonDeliveryReport

`func (o *MicrosoftGraphMessageRulePredicates) SetIsNonDeliveryReport(v bool)`

SetIsNonDeliveryReport gets a reference to the given bool and assigns it to the IsNonDeliveryReport field.

### SetIsNonDeliveryReportExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsNonDeliveryReportExplicitNull(b bool)`

SetIsNonDeliveryReportExplicitNull (un)sets IsNonDeliveryReport to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsNonDeliveryReport value is set to nil even if false is passed
### GetIsPermissionControlled

`func (o *MicrosoftGraphMessageRulePredicates) GetIsPermissionControlled() bool`

GetIsPermissionControlled returns the IsPermissionControlled field if non-nil, zero value otherwise.

### GetIsPermissionControlledOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsPermissionControlledOk() (bool, bool)`

GetIsPermissionControlledOk returns a tuple with the IsPermissionControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsPermissionControlled

`func (o *MicrosoftGraphMessageRulePredicates) HasIsPermissionControlled() bool`

HasIsPermissionControlled returns a boolean if a field has been set.

### SetIsPermissionControlled

`func (o *MicrosoftGraphMessageRulePredicates) SetIsPermissionControlled(v bool)`

SetIsPermissionControlled gets a reference to the given bool and assigns it to the IsPermissionControlled field.

### SetIsPermissionControlledExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsPermissionControlledExplicitNull(b bool)`

SetIsPermissionControlledExplicitNull (un)sets IsPermissionControlled to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsPermissionControlled value is set to nil even if false is passed
### GetIsReadReceipt

`func (o *MicrosoftGraphMessageRulePredicates) GetIsReadReceipt() bool`

GetIsReadReceipt returns the IsReadReceipt field if non-nil, zero value otherwise.

### GetIsReadReceiptOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsReadReceiptOk() (bool, bool)`

GetIsReadReceiptOk returns a tuple with the IsReadReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReadReceipt

`func (o *MicrosoftGraphMessageRulePredicates) HasIsReadReceipt() bool`

HasIsReadReceipt returns a boolean if a field has been set.

### SetIsReadReceipt

`func (o *MicrosoftGraphMessageRulePredicates) SetIsReadReceipt(v bool)`

SetIsReadReceipt gets a reference to the given bool and assigns it to the IsReadReceipt field.

### SetIsReadReceiptExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsReadReceiptExplicitNull(b bool)`

SetIsReadReceiptExplicitNull (un)sets IsReadReceipt to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsReadReceipt value is set to nil even if false is passed
### GetIsSigned

`func (o *MicrosoftGraphMessageRulePredicates) GetIsSigned() bool`

GetIsSigned returns the IsSigned field if non-nil, zero value otherwise.

### GetIsSignedOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsSignedOk() (bool, bool)`

GetIsSignedOk returns a tuple with the IsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsSigned

`func (o *MicrosoftGraphMessageRulePredicates) HasIsSigned() bool`

HasIsSigned returns a boolean if a field has been set.

### SetIsSigned

`func (o *MicrosoftGraphMessageRulePredicates) SetIsSigned(v bool)`

SetIsSigned gets a reference to the given bool and assigns it to the IsSigned field.

### SetIsSignedExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsSignedExplicitNull(b bool)`

SetIsSignedExplicitNull (un)sets IsSigned to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsSigned value is set to nil even if false is passed
### GetIsVoicemail

`func (o *MicrosoftGraphMessageRulePredicates) GetIsVoicemail() bool`

GetIsVoicemail returns the IsVoicemail field if non-nil, zero value otherwise.

### GetIsVoicemailOk

`func (o *MicrosoftGraphMessageRulePredicates) GetIsVoicemailOk() (bool, bool)`

GetIsVoicemailOk returns a tuple with the IsVoicemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsVoicemail

`func (o *MicrosoftGraphMessageRulePredicates) HasIsVoicemail() bool`

HasIsVoicemail returns a boolean if a field has been set.

### SetIsVoicemail

`func (o *MicrosoftGraphMessageRulePredicates) SetIsVoicemail(v bool)`

SetIsVoicemail gets a reference to the given bool and assigns it to the IsVoicemail field.

### SetIsVoicemailExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetIsVoicemailExplicitNull(b bool)`

SetIsVoicemailExplicitNull (un)sets IsVoicemail to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsVoicemail value is set to nil even if false is passed
### GetWithinSizeRange

`func (o *MicrosoftGraphMessageRulePredicates) GetWithinSizeRange() AnyOfmicrosoftGraphSizeRange`

GetWithinSizeRange returns the WithinSizeRange field if non-nil, zero value otherwise.

### GetWithinSizeRangeOk

`func (o *MicrosoftGraphMessageRulePredicates) GetWithinSizeRangeOk() (AnyOfmicrosoftGraphSizeRange, bool)`

GetWithinSizeRangeOk returns a tuple with the WithinSizeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWithinSizeRange

`func (o *MicrosoftGraphMessageRulePredicates) HasWithinSizeRange() bool`

HasWithinSizeRange returns a boolean if a field has been set.

### SetWithinSizeRange

`func (o *MicrosoftGraphMessageRulePredicates) SetWithinSizeRange(v AnyOfmicrosoftGraphSizeRange)`

SetWithinSizeRange gets a reference to the given AnyOfmicrosoftGraphSizeRange and assigns it to the WithinSizeRange field.

### SetWithinSizeRangeExplicitNull

`func (o *MicrosoftGraphMessageRulePredicates) SetWithinSizeRangeExplicitNull(b bool)`

SetWithinSizeRangeExplicitNull (un)sets WithinSizeRange to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The WithinSizeRange value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



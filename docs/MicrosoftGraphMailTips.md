# MicrosoftGraphMailTips

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) |  | [optional] 
**AutomaticReplies** | Pointer to [**AnyOfmicrosoftGraphAutomaticRepliesMailTips**](anyOf&lt;microsoft.graph.automaticRepliesMailTips&gt;.md) |  | [optional] 
**MailboxFull** | Pointer to **bool** |  | [optional] 
**CustomMailTip** | Pointer to **string** |  | [optional] 
**ExternalMemberCount** | Pointer to **int32** |  | [optional] 
**TotalMemberCount** | Pointer to **int32** |  | [optional] 
**DeliveryRestricted** | Pointer to **bool** |  | [optional] 
**IsModerated** | Pointer to **bool** |  | [optional] 
**RecipientScope** | Pointer to [**AnyOfmicrosoftGraphRecipientScopeType**](anyOf&lt;microsoft.graph.recipientScopeType&gt;.md) |  | [optional] 
**RecipientSuggestions** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**MaxMessageSize** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to [**AnyOfmicrosoftGraphMailTipsError**](anyOf&lt;microsoft.graph.mailTipsError&gt;.md) |  | [optional] 

## Methods

### GetEmailAddress

`func (o *MicrosoftGraphMailTips) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphMailTips) GetEmailAddressOk() (AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmailAddress

`func (o *MicrosoftGraphMailTips) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddress

`func (o *MicrosoftGraphMailTips) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddress gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the EmailAddress field.

### SetEmailAddressExplicitNull

`func (o *MicrosoftGraphMailTips) SetEmailAddressExplicitNull(b bool)`

SetEmailAddressExplicitNull (un)sets EmailAddress to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The EmailAddress value is set to nil even if false is passed
### GetAutomaticReplies

`func (o *MicrosoftGraphMailTips) GetAutomaticReplies() AnyOfmicrosoftGraphAutomaticRepliesMailTips`

GetAutomaticReplies returns the AutomaticReplies field if non-nil, zero value otherwise.

### GetAutomaticRepliesOk

`func (o *MicrosoftGraphMailTips) GetAutomaticRepliesOk() (AnyOfmicrosoftGraphAutomaticRepliesMailTips, bool)`

GetAutomaticRepliesOk returns a tuple with the AutomaticReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutomaticReplies

`func (o *MicrosoftGraphMailTips) HasAutomaticReplies() bool`

HasAutomaticReplies returns a boolean if a field has been set.

### SetAutomaticReplies

`func (o *MicrosoftGraphMailTips) SetAutomaticReplies(v AnyOfmicrosoftGraphAutomaticRepliesMailTips)`

SetAutomaticReplies gets a reference to the given AnyOfmicrosoftGraphAutomaticRepliesMailTips and assigns it to the AutomaticReplies field.

### SetAutomaticRepliesExplicitNull

`func (o *MicrosoftGraphMailTips) SetAutomaticRepliesExplicitNull(b bool)`

SetAutomaticRepliesExplicitNull (un)sets AutomaticReplies to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AutomaticReplies value is set to nil even if false is passed
### GetMailboxFull

`func (o *MicrosoftGraphMailTips) GetMailboxFull() bool`

GetMailboxFull returns the MailboxFull field if non-nil, zero value otherwise.

### GetMailboxFullOk

`func (o *MicrosoftGraphMailTips) GetMailboxFullOk() (bool, bool)`

GetMailboxFullOk returns a tuple with the MailboxFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMailboxFull

`func (o *MicrosoftGraphMailTips) HasMailboxFull() bool`

HasMailboxFull returns a boolean if a field has been set.

### SetMailboxFull

`func (o *MicrosoftGraphMailTips) SetMailboxFull(v bool)`

SetMailboxFull gets a reference to the given bool and assigns it to the MailboxFull field.

### SetMailboxFullExplicitNull

`func (o *MicrosoftGraphMailTips) SetMailboxFullExplicitNull(b bool)`

SetMailboxFullExplicitNull (un)sets MailboxFull to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MailboxFull value is set to nil even if false is passed
### GetCustomMailTip

`func (o *MicrosoftGraphMailTips) GetCustomMailTip() string`

GetCustomMailTip returns the CustomMailTip field if non-nil, zero value otherwise.

### GetCustomMailTipOk

`func (o *MicrosoftGraphMailTips) GetCustomMailTipOk() (string, bool)`

GetCustomMailTipOk returns a tuple with the CustomMailTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomMailTip

`func (o *MicrosoftGraphMailTips) HasCustomMailTip() bool`

HasCustomMailTip returns a boolean if a field has been set.

### SetCustomMailTip

`func (o *MicrosoftGraphMailTips) SetCustomMailTip(v string)`

SetCustomMailTip gets a reference to the given string and assigns it to the CustomMailTip field.

### SetCustomMailTipExplicitNull

`func (o *MicrosoftGraphMailTips) SetCustomMailTipExplicitNull(b bool)`

SetCustomMailTipExplicitNull (un)sets CustomMailTip to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CustomMailTip value is set to nil even if false is passed
### GetExternalMemberCount

`func (o *MicrosoftGraphMailTips) GetExternalMemberCount() int32`

GetExternalMemberCount returns the ExternalMemberCount field if non-nil, zero value otherwise.

### GetExternalMemberCountOk

`func (o *MicrosoftGraphMailTips) GetExternalMemberCountOk() (int32, bool)`

GetExternalMemberCountOk returns a tuple with the ExternalMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExternalMemberCount

`func (o *MicrosoftGraphMailTips) HasExternalMemberCount() bool`

HasExternalMemberCount returns a boolean if a field has been set.

### SetExternalMemberCount

`func (o *MicrosoftGraphMailTips) SetExternalMemberCount(v int32)`

SetExternalMemberCount gets a reference to the given int32 and assigns it to the ExternalMemberCount field.

### SetExternalMemberCountExplicitNull

`func (o *MicrosoftGraphMailTips) SetExternalMemberCountExplicitNull(b bool)`

SetExternalMemberCountExplicitNull (un)sets ExternalMemberCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ExternalMemberCount value is set to nil even if false is passed
### GetTotalMemberCount

`func (o *MicrosoftGraphMailTips) GetTotalMemberCount() int32`

GetTotalMemberCount returns the TotalMemberCount field if non-nil, zero value otherwise.

### GetTotalMemberCountOk

`func (o *MicrosoftGraphMailTips) GetTotalMemberCountOk() (int32, bool)`

GetTotalMemberCountOk returns a tuple with the TotalMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalMemberCount

`func (o *MicrosoftGraphMailTips) HasTotalMemberCount() bool`

HasTotalMemberCount returns a boolean if a field has been set.

### SetTotalMemberCount

`func (o *MicrosoftGraphMailTips) SetTotalMemberCount(v int32)`

SetTotalMemberCount gets a reference to the given int32 and assigns it to the TotalMemberCount field.

### SetTotalMemberCountExplicitNull

`func (o *MicrosoftGraphMailTips) SetTotalMemberCountExplicitNull(b bool)`

SetTotalMemberCountExplicitNull (un)sets TotalMemberCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The TotalMemberCount value is set to nil even if false is passed
### GetDeliveryRestricted

`func (o *MicrosoftGraphMailTips) GetDeliveryRestricted() bool`

GetDeliveryRestricted returns the DeliveryRestricted field if non-nil, zero value otherwise.

### GetDeliveryRestrictedOk

`func (o *MicrosoftGraphMailTips) GetDeliveryRestrictedOk() (bool, bool)`

GetDeliveryRestrictedOk returns a tuple with the DeliveryRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeliveryRestricted

`func (o *MicrosoftGraphMailTips) HasDeliveryRestricted() bool`

HasDeliveryRestricted returns a boolean if a field has been set.

### SetDeliveryRestricted

`func (o *MicrosoftGraphMailTips) SetDeliveryRestricted(v bool)`

SetDeliveryRestricted gets a reference to the given bool and assigns it to the DeliveryRestricted field.

### SetDeliveryRestrictedExplicitNull

`func (o *MicrosoftGraphMailTips) SetDeliveryRestrictedExplicitNull(b bool)`

SetDeliveryRestrictedExplicitNull (un)sets DeliveryRestricted to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The DeliveryRestricted value is set to nil even if false is passed
### GetIsModerated

`func (o *MicrosoftGraphMailTips) GetIsModerated() bool`

GetIsModerated returns the IsModerated field if non-nil, zero value otherwise.

### GetIsModeratedOk

`func (o *MicrosoftGraphMailTips) GetIsModeratedOk() (bool, bool)`

GetIsModeratedOk returns a tuple with the IsModerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsModerated

`func (o *MicrosoftGraphMailTips) HasIsModerated() bool`

HasIsModerated returns a boolean if a field has been set.

### SetIsModerated

`func (o *MicrosoftGraphMailTips) SetIsModerated(v bool)`

SetIsModerated gets a reference to the given bool and assigns it to the IsModerated field.

### SetIsModeratedExplicitNull

`func (o *MicrosoftGraphMailTips) SetIsModeratedExplicitNull(b bool)`

SetIsModeratedExplicitNull (un)sets IsModerated to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The IsModerated value is set to nil even if false is passed
### GetRecipientScope

`func (o *MicrosoftGraphMailTips) GetRecipientScope() AnyOfmicrosoftGraphRecipientScopeType`

GetRecipientScope returns the RecipientScope field if non-nil, zero value otherwise.

### GetRecipientScopeOk

`func (o *MicrosoftGraphMailTips) GetRecipientScopeOk() (AnyOfmicrosoftGraphRecipientScopeType, bool)`

GetRecipientScopeOk returns a tuple with the RecipientScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientScope

`func (o *MicrosoftGraphMailTips) HasRecipientScope() bool`

HasRecipientScope returns a boolean if a field has been set.

### SetRecipientScope

`func (o *MicrosoftGraphMailTips) SetRecipientScope(v AnyOfmicrosoftGraphRecipientScopeType)`

SetRecipientScope gets a reference to the given AnyOfmicrosoftGraphRecipientScopeType and assigns it to the RecipientScope field.

### SetRecipientScopeExplicitNull

`func (o *MicrosoftGraphMailTips) SetRecipientScopeExplicitNull(b bool)`

SetRecipientScopeExplicitNull (un)sets RecipientScope to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RecipientScope value is set to nil even if false is passed
### GetRecipientSuggestions

`func (o *MicrosoftGraphMailTips) GetRecipientSuggestions() []AnyOfmicrosoftGraphRecipient`

GetRecipientSuggestions returns the RecipientSuggestions field if non-nil, zero value otherwise.

### GetRecipientSuggestionsOk

`func (o *MicrosoftGraphMailTips) GetRecipientSuggestionsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetRecipientSuggestionsOk returns a tuple with the RecipientSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientSuggestions

`func (o *MicrosoftGraphMailTips) HasRecipientSuggestions() bool`

HasRecipientSuggestions returns a boolean if a field has been set.

### SetRecipientSuggestions

`func (o *MicrosoftGraphMailTips) SetRecipientSuggestions(v []AnyOfmicrosoftGraphRecipient)`

SetRecipientSuggestions gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the RecipientSuggestions field.

### GetMaxMessageSize

`func (o *MicrosoftGraphMailTips) GetMaxMessageSize() int32`

GetMaxMessageSize returns the MaxMessageSize field if non-nil, zero value otherwise.

### GetMaxMessageSizeOk

`func (o *MicrosoftGraphMailTips) GetMaxMessageSizeOk() (int32, bool)`

GetMaxMessageSizeOk returns a tuple with the MaxMessageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxMessageSize

`func (o *MicrosoftGraphMailTips) HasMaxMessageSize() bool`

HasMaxMessageSize returns a boolean if a field has been set.

### SetMaxMessageSize

`func (o *MicrosoftGraphMailTips) SetMaxMessageSize(v int32)`

SetMaxMessageSize gets a reference to the given int32 and assigns it to the MaxMessageSize field.

### SetMaxMessageSizeExplicitNull

`func (o *MicrosoftGraphMailTips) SetMaxMessageSizeExplicitNull(b bool)`

SetMaxMessageSizeExplicitNull (un)sets MaxMessageSize to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MaxMessageSize value is set to nil even if false is passed
### GetError

`func (o *MicrosoftGraphMailTips) GetError() AnyOfmicrosoftGraphMailTipsError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphMailTips) GetErrorOk() (AnyOfmicrosoftGraphMailTipsError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasError

`func (o *MicrosoftGraphMailTips) HasError() bool`

HasError returns a boolean if a field has been set.

### SetError

`func (o *MicrosoftGraphMailTips) SetError(v AnyOfmicrosoftGraphMailTipsError)`

SetError gets a reference to the given AnyOfmicrosoftGraphMailTipsError and assigns it to the Error field.

### SetErrorExplicitNull

`func (o *MicrosoftGraphMailTips) SetErrorExplicitNull(b bool)`

SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Error value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



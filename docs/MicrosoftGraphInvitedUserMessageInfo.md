# MicrosoftGraphInvitedUserMessageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CcRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**MessageLanguage** | Pointer to **string** |  | [optional] 
**CustomizedMessageBody** | Pointer to **string** |  | [optional] 

## Methods

### GetCcRecipients

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetCcRecipients() []AnyOfmicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetCcRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCcRecipients

`func (o *MicrosoftGraphInvitedUserMessageInfo) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### SetCcRecipients

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetCcRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetCcRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the CcRecipients field.

### GetMessageLanguage

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetMessageLanguage() string`

GetMessageLanguage returns the MessageLanguage field if non-nil, zero value otherwise.

### GetMessageLanguageOk

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetMessageLanguageOk() (string, bool)`

GetMessageLanguageOk returns a tuple with the MessageLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessageLanguage

`func (o *MicrosoftGraphInvitedUserMessageInfo) HasMessageLanguage() bool`

HasMessageLanguage returns a boolean if a field has been set.

### SetMessageLanguage

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetMessageLanguage(v string)`

SetMessageLanguage gets a reference to the given string and assigns it to the MessageLanguage field.

### SetMessageLanguageExplicitNull

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetMessageLanguageExplicitNull(b bool)`

SetMessageLanguageExplicitNull (un)sets MessageLanguage to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MessageLanguage value is set to nil even if false is passed
### GetCustomizedMessageBody

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetCustomizedMessageBody() string`

GetCustomizedMessageBody returns the CustomizedMessageBody field if non-nil, zero value otherwise.

### GetCustomizedMessageBodyOk

`func (o *MicrosoftGraphInvitedUserMessageInfo) GetCustomizedMessageBodyOk() (string, bool)`

GetCustomizedMessageBodyOk returns a tuple with the CustomizedMessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomizedMessageBody

`func (o *MicrosoftGraphInvitedUserMessageInfo) HasCustomizedMessageBody() bool`

HasCustomizedMessageBody returns a boolean if a field has been set.

### SetCustomizedMessageBody

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetCustomizedMessageBody(v string)`

SetCustomizedMessageBody gets a reference to the given string and assigns it to the CustomizedMessageBody field.

### SetCustomizedMessageBodyExplicitNull

`func (o *MicrosoftGraphInvitedUserMessageInfo) SetCustomizedMessageBodyExplicitNull(b bool)`

SetCustomizedMessageBodyExplicitNull (un)sets CustomizedMessageBody to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CustomizedMessageBody value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



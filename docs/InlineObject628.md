# InlineObject628

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToRecipients** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**Message** | Pointer to [**AnyOfmicrosoftGraphMessage**](anyOf&lt;microsoft.graph.message&gt;.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### GetToRecipients

`func (o *InlineObject628) GetToRecipients() []AnyOfmicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *InlineObject628) GetToRecipientsOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToRecipients

`func (o *InlineObject628) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### SetToRecipients

`func (o *InlineObject628) SetToRecipients(v []AnyOfmicrosoftGraphRecipient)`

SetToRecipients gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ToRecipients field.

### GetMessage

`func (o *InlineObject628) GetMessage() AnyOfmicrosoftGraphMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject628) GetMessageOk() (AnyOfmicrosoftGraphMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *InlineObject628) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *InlineObject628) SetMessage(v AnyOfmicrosoftGraphMessage)`

SetMessage gets a reference to the given AnyOfmicrosoftGraphMessage and assigns it to the Message field.

### SetMessageExplicitNull

`func (o *InlineObject628) SetMessageExplicitNull(b bool)`

SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Message value is set to nil even if false is passed
### GetComment

`func (o *InlineObject628) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineObject628) GetCommentOk() (string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComment

`func (o *InlineObject628) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetComment

`func (o *InlineObject628) SetComment(v string)`

SetComment gets a reference to the given string and assigns it to the Comment field.

### SetCommentExplicitNull

`func (o *InlineObject628) SetCommentExplicitNull(b bool)`

SetCommentExplicitNull (un)sets Comment to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Comment value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



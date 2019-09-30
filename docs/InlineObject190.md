# InlineObject190

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ToRecipients** | Pointer to [**[]MicrosoftGraphRecipient**](microsoft.graph.recipient.md) |  | [optional] 

## Methods

### GetComment

`func (o *InlineObject190) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineObject190) GetCommentOk() (string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComment

`func (o *InlineObject190) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetComment

`func (o *InlineObject190) SetComment(v string)`

SetComment gets a reference to the given string and assigns it to the Comment field.

### SetCommentExplicitNull

`func (o *InlineObject190) SetCommentExplicitNull(b bool)`

SetCommentExplicitNull (un)sets Comment to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Comment value is set to nil even if false is passed
### GetToRecipients

`func (o *InlineObject190) GetToRecipients() []MicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *InlineObject190) GetToRecipientsOk() ([]MicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToRecipients

`func (o *InlineObject190) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### SetToRecipients

`func (o *InlineObject190) SetToRecipients(v []MicrosoftGraphRecipient)`

SetToRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the ToRecipients field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



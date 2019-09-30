# ConversationThread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToRecipients** | Pointer to [**[]MicrosoftGraphRecipient**](microsoft.graph.recipient.md) |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**HasAttachments** | Pointer to **bool** |  | [optional] 
**LastDeliveredDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**UniqueSenders** | Pointer to **[]string** |  | [optional] 
**CcRecipients** | Pointer to [**[]MicrosoftGraphRecipient**](microsoft.graph.recipient.md) |  | [optional] 
**Preview** | Pointer to **string** |  | [optional] 
**IsLocked** | Pointer to **bool** |  | [optional] 
**Posts** | Pointer to [**[]MicrosoftGraphPost**](microsoft.graph.post.md) |  | [optional] 

## Methods

### GetToRecipients

`func (o *ConversationThread) GetToRecipients() []MicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *ConversationThread) GetToRecipientsOk() ([]MicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToRecipients

`func (o *ConversationThread) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### SetToRecipients

`func (o *ConversationThread) SetToRecipients(v []MicrosoftGraphRecipient)`

SetToRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the ToRecipients field.

### GetTopic

`func (o *ConversationThread) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *ConversationThread) GetTopicOk() (string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTopic

`func (o *ConversationThread) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopic

`func (o *ConversationThread) SetTopic(v string)`

SetTopic gets a reference to the given string and assigns it to the Topic field.

### GetHasAttachments

`func (o *ConversationThread) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *ConversationThread) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *ConversationThread) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *ConversationThread) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### GetLastDeliveredDateTime

`func (o *ConversationThread) GetLastDeliveredDateTime() time.Time`

GetLastDeliveredDateTime returns the LastDeliveredDateTime field if non-nil, zero value otherwise.

### GetLastDeliveredDateTimeOk

`func (o *ConversationThread) GetLastDeliveredDateTimeOk() (time.Time, bool)`

GetLastDeliveredDateTimeOk returns a tuple with the LastDeliveredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastDeliveredDateTime

`func (o *ConversationThread) HasLastDeliveredDateTime() bool`

HasLastDeliveredDateTime returns a boolean if a field has been set.

### SetLastDeliveredDateTime

`func (o *ConversationThread) SetLastDeliveredDateTime(v time.Time)`

SetLastDeliveredDateTime gets a reference to the given time.Time and assigns it to the LastDeliveredDateTime field.

### GetUniqueSenders

`func (o *ConversationThread) GetUniqueSenders() []string`

GetUniqueSenders returns the UniqueSenders field if non-nil, zero value otherwise.

### GetUniqueSendersOk

`func (o *ConversationThread) GetUniqueSendersOk() ([]string, bool)`

GetUniqueSendersOk returns a tuple with the UniqueSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueSenders

`func (o *ConversationThread) HasUniqueSenders() bool`

HasUniqueSenders returns a boolean if a field has been set.

### SetUniqueSenders

`func (o *ConversationThread) SetUniqueSenders(v []string)`

SetUniqueSenders gets a reference to the given []string and assigns it to the UniqueSenders field.

### GetCcRecipients

`func (o *ConversationThread) GetCcRecipients() []MicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *ConversationThread) GetCcRecipientsOk() ([]MicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCcRecipients

`func (o *ConversationThread) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### SetCcRecipients

`func (o *ConversationThread) SetCcRecipients(v []MicrosoftGraphRecipient)`

SetCcRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the CcRecipients field.

### GetPreview

`func (o *ConversationThread) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *ConversationThread) GetPreviewOk() (string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreview

`func (o *ConversationThread) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### SetPreview

`func (o *ConversationThread) SetPreview(v string)`

SetPreview gets a reference to the given string and assigns it to the Preview field.

### GetIsLocked

`func (o *ConversationThread) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ConversationThread) GetIsLockedOk() (bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsLocked

`func (o *ConversationThread) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLocked

`func (o *ConversationThread) SetIsLocked(v bool)`

SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.

### GetPosts

`func (o *ConversationThread) GetPosts() []MicrosoftGraphPost`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *ConversationThread) GetPostsOk() ([]MicrosoftGraphPost, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosts

`func (o *ConversationThread) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### SetPosts

`func (o *ConversationThread) SetPosts(v []MicrosoftGraphPost)`

SetPosts gets a reference to the given []MicrosoftGraphPost and assigns it to the Posts field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



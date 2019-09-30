# MicrosoftGraphConversationThread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphConversationThread) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphConversationThread) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphConversationThread) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphConversationThread) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetToRecipients

`func (o *MicrosoftGraphConversationThread) GetToRecipients() []MicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *MicrosoftGraphConversationThread) GetToRecipientsOk() ([]MicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToRecipients

`func (o *MicrosoftGraphConversationThread) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### SetToRecipients

`func (o *MicrosoftGraphConversationThread) SetToRecipients(v []MicrosoftGraphRecipient)`

SetToRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the ToRecipients field.

### GetTopic

`func (o *MicrosoftGraphConversationThread) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MicrosoftGraphConversationThread) GetTopicOk() (string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTopic

`func (o *MicrosoftGraphConversationThread) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopic

`func (o *MicrosoftGraphConversationThread) SetTopic(v string)`

SetTopic gets a reference to the given string and assigns it to the Topic field.

### GetHasAttachments

`func (o *MicrosoftGraphConversationThread) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphConversationThread) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *MicrosoftGraphConversationThread) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *MicrosoftGraphConversationThread) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### GetLastDeliveredDateTime

`func (o *MicrosoftGraphConversationThread) GetLastDeliveredDateTime() time.Time`

GetLastDeliveredDateTime returns the LastDeliveredDateTime field if non-nil, zero value otherwise.

### GetLastDeliveredDateTimeOk

`func (o *MicrosoftGraphConversationThread) GetLastDeliveredDateTimeOk() (time.Time, bool)`

GetLastDeliveredDateTimeOk returns a tuple with the LastDeliveredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastDeliveredDateTime

`func (o *MicrosoftGraphConversationThread) HasLastDeliveredDateTime() bool`

HasLastDeliveredDateTime returns a boolean if a field has been set.

### SetLastDeliveredDateTime

`func (o *MicrosoftGraphConversationThread) SetLastDeliveredDateTime(v time.Time)`

SetLastDeliveredDateTime gets a reference to the given time.Time and assigns it to the LastDeliveredDateTime field.

### GetUniqueSenders

`func (o *MicrosoftGraphConversationThread) GetUniqueSenders() []string`

GetUniqueSenders returns the UniqueSenders field if non-nil, zero value otherwise.

### GetUniqueSendersOk

`func (o *MicrosoftGraphConversationThread) GetUniqueSendersOk() ([]string, bool)`

GetUniqueSendersOk returns a tuple with the UniqueSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueSenders

`func (o *MicrosoftGraphConversationThread) HasUniqueSenders() bool`

HasUniqueSenders returns a boolean if a field has been set.

### SetUniqueSenders

`func (o *MicrosoftGraphConversationThread) SetUniqueSenders(v []string)`

SetUniqueSenders gets a reference to the given []string and assigns it to the UniqueSenders field.

### GetCcRecipients

`func (o *MicrosoftGraphConversationThread) GetCcRecipients() []MicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *MicrosoftGraphConversationThread) GetCcRecipientsOk() ([]MicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCcRecipients

`func (o *MicrosoftGraphConversationThread) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### SetCcRecipients

`func (o *MicrosoftGraphConversationThread) SetCcRecipients(v []MicrosoftGraphRecipient)`

SetCcRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the CcRecipients field.

### GetPreview

`func (o *MicrosoftGraphConversationThread) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *MicrosoftGraphConversationThread) GetPreviewOk() (string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreview

`func (o *MicrosoftGraphConversationThread) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### SetPreview

`func (o *MicrosoftGraphConversationThread) SetPreview(v string)`

SetPreview gets a reference to the given string and assigns it to the Preview field.

### GetIsLocked

`func (o *MicrosoftGraphConversationThread) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *MicrosoftGraphConversationThread) GetIsLockedOk() (bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsLocked

`func (o *MicrosoftGraphConversationThread) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### SetIsLocked

`func (o *MicrosoftGraphConversationThread) SetIsLocked(v bool)`

SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.

### GetPosts

`func (o *MicrosoftGraphConversationThread) GetPosts() []MicrosoftGraphPost`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *MicrosoftGraphConversationThread) GetPostsOk() ([]MicrosoftGraphPost, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosts

`func (o *MicrosoftGraphConversationThread) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### SetPosts

`func (o *MicrosoftGraphConversationThread) SetPosts(v []MicrosoftGraphPost)`

SetPosts gets a reference to the given []MicrosoftGraphPost and assigns it to the Posts field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



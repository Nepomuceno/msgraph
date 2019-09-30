# MicrosoftGraphConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**HasAttachments** | Pointer to **bool** |  | [optional] 
**LastDeliveredDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**UniqueSenders** | Pointer to **[]string** |  | [optional] 
**Preview** | Pointer to **string** |  | [optional] 
**Threads** | Pointer to [**[]MicrosoftGraphConversationThread**](microsoft.graph.conversationThread.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphConversation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphConversation) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphConversation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphConversation) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetTopic

`func (o *MicrosoftGraphConversation) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MicrosoftGraphConversation) GetTopicOk() (string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTopic

`func (o *MicrosoftGraphConversation) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopic

`func (o *MicrosoftGraphConversation) SetTopic(v string)`

SetTopic gets a reference to the given string and assigns it to the Topic field.

### GetHasAttachments

`func (o *MicrosoftGraphConversation) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphConversation) GetHasAttachmentsOk() (bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHasAttachments

`func (o *MicrosoftGraphConversation) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### SetHasAttachments

`func (o *MicrosoftGraphConversation) SetHasAttachments(v bool)`

SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.

### GetLastDeliveredDateTime

`func (o *MicrosoftGraphConversation) GetLastDeliveredDateTime() time.Time`

GetLastDeliveredDateTime returns the LastDeliveredDateTime field if non-nil, zero value otherwise.

### GetLastDeliveredDateTimeOk

`func (o *MicrosoftGraphConversation) GetLastDeliveredDateTimeOk() (time.Time, bool)`

GetLastDeliveredDateTimeOk returns a tuple with the LastDeliveredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastDeliveredDateTime

`func (o *MicrosoftGraphConversation) HasLastDeliveredDateTime() bool`

HasLastDeliveredDateTime returns a boolean if a field has been set.

### SetLastDeliveredDateTime

`func (o *MicrosoftGraphConversation) SetLastDeliveredDateTime(v time.Time)`

SetLastDeliveredDateTime gets a reference to the given time.Time and assigns it to the LastDeliveredDateTime field.

### GetUniqueSenders

`func (o *MicrosoftGraphConversation) GetUniqueSenders() []string`

GetUniqueSenders returns the UniqueSenders field if non-nil, zero value otherwise.

### GetUniqueSendersOk

`func (o *MicrosoftGraphConversation) GetUniqueSendersOk() ([]string, bool)`

GetUniqueSendersOk returns a tuple with the UniqueSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUniqueSenders

`func (o *MicrosoftGraphConversation) HasUniqueSenders() bool`

HasUniqueSenders returns a boolean if a field has been set.

### SetUniqueSenders

`func (o *MicrosoftGraphConversation) SetUniqueSenders(v []string)`

SetUniqueSenders gets a reference to the given []string and assigns it to the UniqueSenders field.

### GetPreview

`func (o *MicrosoftGraphConversation) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *MicrosoftGraphConversation) GetPreviewOk() (string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreview

`func (o *MicrosoftGraphConversation) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### SetPreview

`func (o *MicrosoftGraphConversation) SetPreview(v string)`

SetPreview gets a reference to the given string and assigns it to the Preview field.

### GetThreads

`func (o *MicrosoftGraphConversation) GetThreads() []MicrosoftGraphConversationThread`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *MicrosoftGraphConversation) GetThreadsOk() ([]MicrosoftGraphConversationThread, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasThreads

`func (o *MicrosoftGraphConversation) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### SetThreads

`func (o *MicrosoftGraphConversation) SetThreads(v []MicrosoftGraphConversationThread)`

SetThreads gets a reference to the given []MicrosoftGraphConversationThread and assigns it to the Threads field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



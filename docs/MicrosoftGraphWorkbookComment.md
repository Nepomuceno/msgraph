# MicrosoftGraphWorkbookComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Replies** | Pointer to [**[]MicrosoftGraphWorkbookCommentReply**](microsoft.graph.workbookCommentReply.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookComment) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookComment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookComment) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetContent

`func (o *MicrosoftGraphWorkbookComment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphWorkbookComment) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *MicrosoftGraphWorkbookComment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *MicrosoftGraphWorkbookComment) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *MicrosoftGraphWorkbookComment) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed
### GetContentType

`func (o *MicrosoftGraphWorkbookComment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphWorkbookComment) GetContentTypeOk() (string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentType

`func (o *MicrosoftGraphWorkbookComment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentType

`func (o *MicrosoftGraphWorkbookComment) SetContentType(v string)`

SetContentType gets a reference to the given string and assigns it to the ContentType field.

### GetReplies

`func (o *MicrosoftGraphWorkbookComment) GetReplies() []MicrosoftGraphWorkbookCommentReply`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *MicrosoftGraphWorkbookComment) GetRepliesOk() ([]MicrosoftGraphWorkbookCommentReply, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReplies

`func (o *MicrosoftGraphWorkbookComment) HasReplies() bool`

HasReplies returns a boolean if a field has been set.

### SetReplies

`func (o *MicrosoftGraphWorkbookComment) SetReplies(v []MicrosoftGraphWorkbookCommentReply)`

SetReplies gets a reference to the given []MicrosoftGraphWorkbookCommentReply and assigns it to the Replies field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



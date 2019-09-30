# WorkbookComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Replies** | Pointer to [**[]MicrosoftGraphWorkbookCommentReply**](microsoft.graph.workbookCommentReply.md) |  | [optional] 

## Methods

### GetContent

`func (o *WorkbookComment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WorkbookComment) GetContentOk() (string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContent

`func (o *WorkbookComment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContent

`func (o *WorkbookComment) SetContent(v string)`

SetContent gets a reference to the given string and assigns it to the Content field.

### SetContentExplicitNull

`func (o *WorkbookComment) SetContentExplicitNull(b bool)`

SetContentExplicitNull (un)sets Content to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Content value is set to nil even if false is passed
### GetContentType

`func (o *WorkbookComment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WorkbookComment) GetContentTypeOk() (string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasContentType

`func (o *WorkbookComment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentType

`func (o *WorkbookComment) SetContentType(v string)`

SetContentType gets a reference to the given string and assigns it to the ContentType field.

### GetReplies

`func (o *WorkbookComment) GetReplies() []MicrosoftGraphWorkbookCommentReply`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *WorkbookComment) GetRepliesOk() ([]MicrosoftGraphWorkbookCommentReply, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReplies

`func (o *WorkbookComment) HasReplies() bool`

HasReplies returns a boolean if a field has been set.

### SetReplies

`func (o *WorkbookComment) SetReplies(v []MicrosoftGraphWorkbookCommentReply)`

SetReplies gets a reference to the given []MicrosoftGraphWorkbookCommentReply and assigns it to the Replies field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



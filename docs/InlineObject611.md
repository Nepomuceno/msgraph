# InlineObject611

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to [**AnyOfmicrosoftGraphMessage**](anyOf&lt;microsoft.graph.message&gt;.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### GetMessage

`func (o *InlineObject611) GetMessage() AnyOfmicrosoftGraphMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject611) GetMessageOk() (AnyOfmicrosoftGraphMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *InlineObject611) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *InlineObject611) SetMessage(v AnyOfmicrosoftGraphMessage)`

SetMessage gets a reference to the given AnyOfmicrosoftGraphMessage and assigns it to the Message field.

### SetMessageExplicitNull

`func (o *InlineObject611) SetMessageExplicitNull(b bool)`

SetMessageExplicitNull (un)sets Message to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Message value is set to nil even if false is passed
### GetComment

`func (o *InlineObject611) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InlineObject611) GetCommentOk() (string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComment

`func (o *InlineObject611) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetComment

`func (o *InlineObject611) SetComment(v string)`

SetComment gets a reference to the given string and assigns it to the Comment field.

### SetCommentExplicitNull

`func (o *InlineObject611) SetCommentExplicitNull(b bool)`

SetCommentExplicitNull (un)sets Comment to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Comment value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



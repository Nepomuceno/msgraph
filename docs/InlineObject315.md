# InlineObject315

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to [**MicrosoftGraphMessage**](microsoft.graph.message.md) |  | [optional] 
**SaveToSentItems** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### GetMessage

`func (o *InlineObject315) GetMessage() MicrosoftGraphMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject315) GetMessageOk() (MicrosoftGraphMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *InlineObject315) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *InlineObject315) SetMessage(v MicrosoftGraphMessage)`

SetMessage gets a reference to the given MicrosoftGraphMessage and assigns it to the Message field.

### GetSaveToSentItems

`func (o *InlineObject315) GetSaveToSentItems() bool`

GetSaveToSentItems returns the SaveToSentItems field if non-nil, zero value otherwise.

### GetSaveToSentItemsOk

`func (o *InlineObject315) GetSaveToSentItemsOk() (bool, bool)`

GetSaveToSentItemsOk returns a tuple with the SaveToSentItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSaveToSentItems

`func (o *InlineObject315) HasSaveToSentItems() bool`

HasSaveToSentItems returns a boolean if a field has been set.

### SetSaveToSentItems

`func (o *InlineObject315) SetSaveToSentItems(v bool)`

SetSaveToSentItems gets a reference to the given bool and assigns it to the SaveToSentItems field.

### SetSaveToSentItemsExplicitNull

`func (o *InlineObject315) SetSaveToSentItemsExplicitNull(b bool)`

SetSaveToSentItemsExplicitNull (un)sets SaveToSentItems to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The SaveToSentItems value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



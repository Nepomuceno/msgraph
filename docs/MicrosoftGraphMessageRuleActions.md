# MicrosoftGraphMessageRuleActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveToFolder** | Pointer to **string** |  | [optional] 
**CopyToFolder** | Pointer to **string** |  | [optional] 
**Delete** | Pointer to **bool** |  | [optional] 
**PermanentDelete** | Pointer to **bool** |  | [optional] 
**MarkAsRead** | Pointer to **bool** |  | [optional] 
**MarkImportance** | Pointer to [**AnyOfmicrosoftGraphImportance**](anyOf&lt;microsoft.graph.importance&gt;.md) |  | [optional] 
**ForwardTo** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**ForwardAsAttachmentTo** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**RedirectTo** | Pointer to [**[]AnyOfmicrosoftGraphRecipient**](anyOf&lt;microsoft.graph.recipient&gt;.md) |  | [optional] 
**AssignCategories** | Pointer to **[]string** |  | [optional] 
**StopProcessingRules** | Pointer to **bool** |  | [optional] 

## Methods

### GetMoveToFolder

`func (o *MicrosoftGraphMessageRuleActions) GetMoveToFolder() string`

GetMoveToFolder returns the MoveToFolder field if non-nil, zero value otherwise.

### GetMoveToFolderOk

`func (o *MicrosoftGraphMessageRuleActions) GetMoveToFolderOk() (string, bool)`

GetMoveToFolderOk returns a tuple with the MoveToFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMoveToFolder

`func (o *MicrosoftGraphMessageRuleActions) HasMoveToFolder() bool`

HasMoveToFolder returns a boolean if a field has been set.

### SetMoveToFolder

`func (o *MicrosoftGraphMessageRuleActions) SetMoveToFolder(v string)`

SetMoveToFolder gets a reference to the given string and assigns it to the MoveToFolder field.

### SetMoveToFolderExplicitNull

`func (o *MicrosoftGraphMessageRuleActions) SetMoveToFolderExplicitNull(b bool)`

SetMoveToFolderExplicitNull (un)sets MoveToFolder to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MoveToFolder value is set to nil even if false is passed
### GetCopyToFolder

`func (o *MicrosoftGraphMessageRuleActions) GetCopyToFolder() string`

GetCopyToFolder returns the CopyToFolder field if non-nil, zero value otherwise.

### GetCopyToFolderOk

`func (o *MicrosoftGraphMessageRuleActions) GetCopyToFolderOk() (string, bool)`

GetCopyToFolderOk returns a tuple with the CopyToFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCopyToFolder

`func (o *MicrosoftGraphMessageRuleActions) HasCopyToFolder() bool`

HasCopyToFolder returns a boolean if a field has been set.

### SetCopyToFolder

`func (o *MicrosoftGraphMessageRuleActions) SetCopyToFolder(v string)`

SetCopyToFolder gets a reference to the given string and assigns it to the CopyToFolder field.

### SetCopyToFolderExplicitNull

`func (o *MicrosoftGraphMessageRuleActions) SetCopyToFolderExplicitNull(b bool)`

SetCopyToFolderExplicitNull (un)sets CopyToFolder to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CopyToFolder value is set to nil even if false is passed
### GetDelete

`func (o *MicrosoftGraphMessageRuleActions) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *MicrosoftGraphMessageRuleActions) GetDeleteOk() (bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDelete

`func (o *MicrosoftGraphMessageRuleActions) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDelete

`func (o *MicrosoftGraphMessageRuleActions) SetDelete(v bool)`

SetDelete gets a reference to the given bool and assigns it to the Delete field.

### SetDeleteExplicitNull

`func (o *MicrosoftGraphMessageRuleActions) SetDeleteExplicitNull(b bool)`

SetDeleteExplicitNull (un)sets Delete to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Delete value is set to nil even if false is passed
### GetPermanentDelete

`func (o *MicrosoftGraphMessageRuleActions) GetPermanentDelete() bool`

GetPermanentDelete returns the PermanentDelete field if non-nil, zero value otherwise.

### GetPermanentDeleteOk

`func (o *MicrosoftGraphMessageRuleActions) GetPermanentDeleteOk() (bool, bool)`

GetPermanentDeleteOk returns a tuple with the PermanentDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermanentDelete

`func (o *MicrosoftGraphMessageRuleActions) HasPermanentDelete() bool`

HasPermanentDelete returns a boolean if a field has been set.

### SetPermanentDelete

`func (o *MicrosoftGraphMessageRuleActions) SetPermanentDelete(v bool)`

SetPermanentDelete gets a reference to the given bool and assigns it to the PermanentDelete field.

### SetPermanentDeleteExplicitNull

`func (o *MicrosoftGraphMessageRuleActions) SetPermanentDeleteExplicitNull(b bool)`

SetPermanentDeleteExplicitNull (un)sets PermanentDelete to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The PermanentDelete value is set to nil even if false is passed
### GetMarkAsRead

`func (o *MicrosoftGraphMessageRuleActions) GetMarkAsRead() bool`

GetMarkAsRead returns the MarkAsRead field if non-nil, zero value otherwise.

### GetMarkAsReadOk

`func (o *MicrosoftGraphMessageRuleActions) GetMarkAsReadOk() (bool, bool)`

GetMarkAsReadOk returns a tuple with the MarkAsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMarkAsRead

`func (o *MicrosoftGraphMessageRuleActions) HasMarkAsRead() bool`

HasMarkAsRead returns a boolean if a field has been set.

### SetMarkAsRead

`func (o *MicrosoftGraphMessageRuleActions) SetMarkAsRead(v bool)`

SetMarkAsRead gets a reference to the given bool and assigns it to the MarkAsRead field.

### SetMarkAsReadExplicitNull

`func (o *MicrosoftGraphMessageRuleActions) SetMarkAsReadExplicitNull(b bool)`

SetMarkAsReadExplicitNull (un)sets MarkAsRead to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MarkAsRead value is set to nil even if false is passed
### GetMarkImportance

`func (o *MicrosoftGraphMessageRuleActions) GetMarkImportance() AnyOfmicrosoftGraphImportance`

GetMarkImportance returns the MarkImportance field if non-nil, zero value otherwise.

### GetMarkImportanceOk

`func (o *MicrosoftGraphMessageRuleActions) GetMarkImportanceOk() (AnyOfmicrosoftGraphImportance, bool)`

GetMarkImportanceOk returns a tuple with the MarkImportance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMarkImportance

`func (o *MicrosoftGraphMessageRuleActions) HasMarkImportance() bool`

HasMarkImportance returns a boolean if a field has been set.

### SetMarkImportance

`func (o *MicrosoftGraphMessageRuleActions) SetMarkImportance(v AnyOfmicrosoftGraphImportance)`

SetMarkImportance gets a reference to the given AnyOfmicrosoftGraphImportance and assigns it to the MarkImportance field.

### SetMarkImportanceExplicitNull

`func (o *MicrosoftGraphMessageRuleActions) SetMarkImportanceExplicitNull(b bool)`

SetMarkImportanceExplicitNull (un)sets MarkImportance to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The MarkImportance value is set to nil even if false is passed
### GetForwardTo

`func (o *MicrosoftGraphMessageRuleActions) GetForwardTo() []AnyOfmicrosoftGraphRecipient`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *MicrosoftGraphMessageRuleActions) GetForwardToOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasForwardTo

`func (o *MicrosoftGraphMessageRuleActions) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### SetForwardTo

`func (o *MicrosoftGraphMessageRuleActions) SetForwardTo(v []AnyOfmicrosoftGraphRecipient)`

SetForwardTo gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ForwardTo field.

### GetForwardAsAttachmentTo

`func (o *MicrosoftGraphMessageRuleActions) GetForwardAsAttachmentTo() []AnyOfmicrosoftGraphRecipient`

GetForwardAsAttachmentTo returns the ForwardAsAttachmentTo field if non-nil, zero value otherwise.

### GetForwardAsAttachmentToOk

`func (o *MicrosoftGraphMessageRuleActions) GetForwardAsAttachmentToOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetForwardAsAttachmentToOk returns a tuple with the ForwardAsAttachmentTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasForwardAsAttachmentTo

`func (o *MicrosoftGraphMessageRuleActions) HasForwardAsAttachmentTo() bool`

HasForwardAsAttachmentTo returns a boolean if a field has been set.

### SetForwardAsAttachmentTo

`func (o *MicrosoftGraphMessageRuleActions) SetForwardAsAttachmentTo(v []AnyOfmicrosoftGraphRecipient)`

SetForwardAsAttachmentTo gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the ForwardAsAttachmentTo field.

### GetRedirectTo

`func (o *MicrosoftGraphMessageRuleActions) GetRedirectTo() []AnyOfmicrosoftGraphRecipient`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *MicrosoftGraphMessageRuleActions) GetRedirectToOk() ([]AnyOfmicrosoftGraphRecipient, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRedirectTo

`func (o *MicrosoftGraphMessageRuleActions) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### SetRedirectTo

`func (o *MicrosoftGraphMessageRuleActions) SetRedirectTo(v []AnyOfmicrosoftGraphRecipient)`

SetRedirectTo gets a reference to the given []AnyOfmicrosoftGraphRecipient and assigns it to the RedirectTo field.

### GetAssignCategories

`func (o *MicrosoftGraphMessageRuleActions) GetAssignCategories() []string`

GetAssignCategories returns the AssignCategories field if non-nil, zero value otherwise.

### GetAssignCategoriesOk

`func (o *MicrosoftGraphMessageRuleActions) GetAssignCategoriesOk() ([]string, bool)`

GetAssignCategoriesOk returns a tuple with the AssignCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAssignCategories

`func (o *MicrosoftGraphMessageRuleActions) HasAssignCategories() bool`

HasAssignCategories returns a boolean if a field has been set.

### SetAssignCategories

`func (o *MicrosoftGraphMessageRuleActions) SetAssignCategories(v []string)`

SetAssignCategories gets a reference to the given []string and assigns it to the AssignCategories field.

### GetStopProcessingRules

`func (o *MicrosoftGraphMessageRuleActions) GetStopProcessingRules() bool`

GetStopProcessingRules returns the StopProcessingRules field if non-nil, zero value otherwise.

### GetStopProcessingRulesOk

`func (o *MicrosoftGraphMessageRuleActions) GetStopProcessingRulesOk() (bool, bool)`

GetStopProcessingRulesOk returns a tuple with the StopProcessingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStopProcessingRules

`func (o *MicrosoftGraphMessageRuleActions) HasStopProcessingRules() bool`

HasStopProcessingRules returns a boolean if a field has been set.

### SetStopProcessingRules

`func (o *MicrosoftGraphMessageRuleActions) SetStopProcessingRules(v bool)`

SetStopProcessingRules gets a reference to the given bool and assigns it to the StopProcessingRules field.

### SetStopProcessingRulesExplicitNull

`func (o *MicrosoftGraphMessageRuleActions) SetStopProcessingRulesExplicitNull(b bool)`

SetStopProcessingRulesExplicitNull (un)sets StopProcessingRules to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The StopProcessingRules value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



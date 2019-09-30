# MicrosoftGraphFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildCount** | Pointer to **int32** |  | [optional] 
**View** | Pointer to [**AnyOfmicrosoftGraphFolderView**](anyOf&lt;microsoft.graph.folderView&gt;.md) |  | [optional] 

## Methods

### GetChildCount

`func (o *MicrosoftGraphFolder) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *MicrosoftGraphFolder) GetChildCountOk() (int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChildCount

`func (o *MicrosoftGraphFolder) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### SetChildCount

`func (o *MicrosoftGraphFolder) SetChildCount(v int32)`

SetChildCount gets a reference to the given int32 and assigns it to the ChildCount field.

### SetChildCountExplicitNull

`func (o *MicrosoftGraphFolder) SetChildCountExplicitNull(b bool)`

SetChildCountExplicitNull (un)sets ChildCount to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ChildCount value is set to nil even if false is passed
### GetView

`func (o *MicrosoftGraphFolder) GetView() AnyOfmicrosoftGraphFolderView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *MicrosoftGraphFolder) GetViewOk() (AnyOfmicrosoftGraphFolderView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasView

`func (o *MicrosoftGraphFolder) HasView() bool`

HasView returns a boolean if a field has been set.

### SetView

`func (o *MicrosoftGraphFolder) SetView(v AnyOfmicrosoftGraphFolderView)`

SetView gets a reference to the given AnyOfmicrosoftGraphFolderView and assigns it to the View field.

### SetViewExplicitNull

`func (o *MicrosoftGraphFolder) SetViewExplicitNull(b bool)`

SetViewExplicitNull (un)sets View to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The View value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



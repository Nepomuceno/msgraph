# ItemAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**AnyOfmicrosoftGraphOutlookItem**](anyOf&lt;microsoft.graph.outlookItem&gt;.md) |  | [optional] 

## Methods

### GetItem

`func (o *ItemAttachment) GetItem() AnyOfmicrosoftGraphOutlookItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ItemAttachment) GetItemOk() (AnyOfmicrosoftGraphOutlookItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasItem

`func (o *ItemAttachment) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItem

`func (o *ItemAttachment) SetItem(v AnyOfmicrosoftGraphOutlookItem)`

SetItem gets a reference to the given AnyOfmicrosoftGraphOutlookItem and assigns it to the Item field.

### SetItemExplicitNull

`func (o *ItemAttachment) SetItemExplicitNull(b bool)`

SetItemExplicitNull (un)sets Item to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Item value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



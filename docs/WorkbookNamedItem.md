# WorkbookNamedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) |  | [optional] 

## Methods

### GetComment

`func (o *WorkbookNamedItem) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WorkbookNamedItem) GetCommentOk() (string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComment

`func (o *WorkbookNamedItem) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetComment

`func (o *WorkbookNamedItem) SetComment(v string)`

SetComment gets a reference to the given string and assigns it to the Comment field.

### SetCommentExplicitNull

`func (o *WorkbookNamedItem) SetCommentExplicitNull(b bool)`

SetCommentExplicitNull (un)sets Comment to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Comment value is set to nil even if false is passed
### GetName

`func (o *WorkbookNamedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookNamedItem) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *WorkbookNamedItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *WorkbookNamedItem) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *WorkbookNamedItem) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetScope

`func (o *WorkbookNamedItem) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *WorkbookNamedItem) GetScopeOk() (string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScope

`func (o *WorkbookNamedItem) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScope

`func (o *WorkbookNamedItem) SetScope(v string)`

SetScope gets a reference to the given string and assigns it to the Scope field.

### GetType

`func (o *WorkbookNamedItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkbookNamedItem) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *WorkbookNamedItem) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *WorkbookNamedItem) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### SetTypeExplicitNull

`func (o *WorkbookNamedItem) SetTypeExplicitNull(b bool)`

SetTypeExplicitNull (un)sets Type to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Type value is set to nil even if false is passed
### GetValue

`func (o *WorkbookNamedItem) GetValue() AnyOfobject`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkbookNamedItem) GetValueOk() (AnyOfobject, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *WorkbookNamedItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *WorkbookNamedItem) SetValue(v AnyOfobject)`

SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.

### SetValueExplicitNull

`func (o *WorkbookNamedItem) SetValueExplicitNull(b bool)`

SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Value value is set to nil even if false is passed
### GetVisible

`func (o *WorkbookNamedItem) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *WorkbookNamedItem) GetVisibleOk() (bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVisible

`func (o *WorkbookNamedItem) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### SetVisible

`func (o *WorkbookNamedItem) SetVisible(v bool)`

SetVisible gets a reference to the given bool and assigns it to the Visible field.

### GetWorksheet

`func (o *WorkbookNamedItem) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *WorkbookNamedItem) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorksheet

`func (o *WorkbookNamedItem) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheet

`func (o *WorkbookNamedItem) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.

### SetWorksheetExplicitNull

`func (o *WorkbookNamedItem) SetWorksheetExplicitNull(b bool)`

SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Worksheet value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



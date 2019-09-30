# WorkbookPivotTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) |  | [optional] 

## Methods

### GetName

`func (o *WorkbookPivotTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookPivotTable) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *WorkbookPivotTable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *WorkbookPivotTable) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *WorkbookPivotTable) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetWorksheet

`func (o *WorkbookPivotTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *WorkbookPivotTable) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorksheet

`func (o *WorkbookPivotTable) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheet

`func (o *WorkbookPivotTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.

### SetWorksheetExplicitNull

`func (o *WorkbookPivotTable) SetWorksheetExplicitNull(b bool)`

SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Worksheet value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



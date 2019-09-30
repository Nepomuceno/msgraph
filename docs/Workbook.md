# Workbook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**AnyOfmicrosoftGraphWorkbookApplication**](anyOf&lt;microsoft.graph.workbookApplication&gt;.md) |  | [optional] 
**Names** | Pointer to [**[]MicrosoftGraphWorkbookNamedItem**](microsoft.graph.workbookNamedItem.md) |  | [optional] 
**Tables** | Pointer to [**[]MicrosoftGraphWorkbookTable**](microsoft.graph.workbookTable.md) |  | [optional] 
**Worksheets** | Pointer to [**[]MicrosoftGraphWorkbookWorksheet**](microsoft.graph.workbookWorksheet.md) |  | [optional] 
**Comments** | Pointer to [**[]MicrosoftGraphWorkbookComment**](microsoft.graph.workbookComment.md) |  | [optional] 
**Functions** | Pointer to [**AnyOfmicrosoftGraphWorkbookFunctions**](anyOf&lt;microsoft.graph.workbookFunctions&gt;.md) |  | [optional] 

## Methods

### GetApplication

`func (o *Workbook) GetApplication() AnyOfmicrosoftGraphWorkbookApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Workbook) GetApplicationOk() (AnyOfmicrosoftGraphWorkbookApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplication

`func (o *Workbook) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplication

`func (o *Workbook) SetApplication(v AnyOfmicrosoftGraphWorkbookApplication)`

SetApplication gets a reference to the given AnyOfmicrosoftGraphWorkbookApplication and assigns it to the Application field.

### SetApplicationExplicitNull

`func (o *Workbook) SetApplicationExplicitNull(b bool)`

SetApplicationExplicitNull (un)sets Application to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Application value is set to nil even if false is passed
### GetNames

`func (o *Workbook) GetNames() []MicrosoftGraphWorkbookNamedItem`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Workbook) GetNamesOk() ([]MicrosoftGraphWorkbookNamedItem, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNames

`func (o *Workbook) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNames

`func (o *Workbook) SetNames(v []MicrosoftGraphWorkbookNamedItem)`

SetNames gets a reference to the given []MicrosoftGraphWorkbookNamedItem and assigns it to the Names field.

### GetTables

`func (o *Workbook) GetTables() []MicrosoftGraphWorkbookTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *Workbook) GetTablesOk() ([]MicrosoftGraphWorkbookTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTables

`func (o *Workbook) HasTables() bool`

HasTables returns a boolean if a field has been set.

### SetTables

`func (o *Workbook) SetTables(v []MicrosoftGraphWorkbookTable)`

SetTables gets a reference to the given []MicrosoftGraphWorkbookTable and assigns it to the Tables field.

### GetWorksheets

`func (o *Workbook) GetWorksheets() []MicrosoftGraphWorkbookWorksheet`

GetWorksheets returns the Worksheets field if non-nil, zero value otherwise.

### GetWorksheetsOk

`func (o *Workbook) GetWorksheetsOk() ([]MicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetsOk returns a tuple with the Worksheets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorksheets

`func (o *Workbook) HasWorksheets() bool`

HasWorksheets returns a boolean if a field has been set.

### SetWorksheets

`func (o *Workbook) SetWorksheets(v []MicrosoftGraphWorkbookWorksheet)`

SetWorksheets gets a reference to the given []MicrosoftGraphWorkbookWorksheet and assigns it to the Worksheets field.

### GetComments

`func (o *Workbook) GetComments() []MicrosoftGraphWorkbookComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Workbook) GetCommentsOk() ([]MicrosoftGraphWorkbookComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComments

`func (o *Workbook) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetComments

`func (o *Workbook) SetComments(v []MicrosoftGraphWorkbookComment)`

SetComments gets a reference to the given []MicrosoftGraphWorkbookComment and assigns it to the Comments field.

### GetFunctions

`func (o *Workbook) GetFunctions() AnyOfmicrosoftGraphWorkbookFunctions`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *Workbook) GetFunctionsOk() (AnyOfmicrosoftGraphWorkbookFunctions, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFunctions

`func (o *Workbook) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.

### SetFunctions

`func (o *Workbook) SetFunctions(v AnyOfmicrosoftGraphWorkbookFunctions)`

SetFunctions gets a reference to the given AnyOfmicrosoftGraphWorkbookFunctions and assigns it to the Functions field.

### SetFunctionsExplicitNull

`func (o *Workbook) SetFunctionsExplicitNull(b bool)`

SetFunctionsExplicitNull (un)sets Functions to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Functions value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



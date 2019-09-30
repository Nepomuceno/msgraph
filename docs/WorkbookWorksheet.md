# WorkbookWorksheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Charts** | Pointer to [**[]MicrosoftGraphWorkbookChart**](microsoft.graph.workbookChart.md) |  | [optional] 
**Names** | Pointer to [**[]MicrosoftGraphWorkbookNamedItem**](microsoft.graph.workbookNamedItem.md) |  | [optional] 
**PivotTables** | Pointer to [**[]MicrosoftGraphWorkbookPivotTable**](microsoft.graph.workbookPivotTable.md) |  | [optional] 
**Protection** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheetProtection**](anyOf&lt;microsoft.graph.workbookWorksheetProtection&gt;.md) |  | [optional] 
**Tables** | Pointer to [**[]MicrosoftGraphWorkbookTable**](microsoft.graph.workbookTable.md) |  | [optional] 

## Methods

### GetName

`func (o *WorkbookWorksheet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookWorksheet) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *WorkbookWorksheet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *WorkbookWorksheet) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *WorkbookWorksheet) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetPosition

`func (o *WorkbookWorksheet) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *WorkbookWorksheet) GetPositionOk() (int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosition

`func (o *WorkbookWorksheet) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPosition

`func (o *WorkbookWorksheet) SetPosition(v int32)`

SetPosition gets a reference to the given int32 and assigns it to the Position field.

### GetVisibility

`func (o *WorkbookWorksheet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *WorkbookWorksheet) GetVisibilityOk() (string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVisibility

`func (o *WorkbookWorksheet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibility

`func (o *WorkbookWorksheet) SetVisibility(v string)`

SetVisibility gets a reference to the given string and assigns it to the Visibility field.

### GetCharts

`func (o *WorkbookWorksheet) GetCharts() []MicrosoftGraphWorkbookChart`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *WorkbookWorksheet) GetChartsOk() ([]MicrosoftGraphWorkbookChart, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCharts

`func (o *WorkbookWorksheet) HasCharts() bool`

HasCharts returns a boolean if a field has been set.

### SetCharts

`func (o *WorkbookWorksheet) SetCharts(v []MicrosoftGraphWorkbookChart)`

SetCharts gets a reference to the given []MicrosoftGraphWorkbookChart and assigns it to the Charts field.

### GetNames

`func (o *WorkbookWorksheet) GetNames() []MicrosoftGraphWorkbookNamedItem`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *WorkbookWorksheet) GetNamesOk() ([]MicrosoftGraphWorkbookNamedItem, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNames

`func (o *WorkbookWorksheet) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNames

`func (o *WorkbookWorksheet) SetNames(v []MicrosoftGraphWorkbookNamedItem)`

SetNames gets a reference to the given []MicrosoftGraphWorkbookNamedItem and assigns it to the Names field.

### GetPivotTables

`func (o *WorkbookWorksheet) GetPivotTables() []MicrosoftGraphWorkbookPivotTable`

GetPivotTables returns the PivotTables field if non-nil, zero value otherwise.

### GetPivotTablesOk

`func (o *WorkbookWorksheet) GetPivotTablesOk() ([]MicrosoftGraphWorkbookPivotTable, bool)`

GetPivotTablesOk returns a tuple with the PivotTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPivotTables

`func (o *WorkbookWorksheet) HasPivotTables() bool`

HasPivotTables returns a boolean if a field has been set.

### SetPivotTables

`func (o *WorkbookWorksheet) SetPivotTables(v []MicrosoftGraphWorkbookPivotTable)`

SetPivotTables gets a reference to the given []MicrosoftGraphWorkbookPivotTable and assigns it to the PivotTables field.

### GetProtection

`func (o *WorkbookWorksheet) GetProtection() AnyOfmicrosoftGraphWorkbookWorksheetProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *WorkbookWorksheet) GetProtectionOk() (AnyOfmicrosoftGraphWorkbookWorksheetProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtection

`func (o *WorkbookWorksheet) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### SetProtection

`func (o *WorkbookWorksheet) SetProtection(v AnyOfmicrosoftGraphWorkbookWorksheetProtection)`

SetProtection gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheetProtection and assigns it to the Protection field.

### SetProtectionExplicitNull

`func (o *WorkbookWorksheet) SetProtectionExplicitNull(b bool)`

SetProtectionExplicitNull (un)sets Protection to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Protection value is set to nil even if false is passed
### GetTables

`func (o *WorkbookWorksheet) GetTables() []MicrosoftGraphWorkbookTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *WorkbookWorksheet) GetTablesOk() ([]MicrosoftGraphWorkbookTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTables

`func (o *WorkbookWorksheet) HasTables() bool`

HasTables returns a boolean if a field has been set.

### SetTables

`func (o *WorkbookWorksheet) SetTables(v []MicrosoftGraphWorkbookTable)`

SetTables gets a reference to the given []MicrosoftGraphWorkbookTable and assigns it to the Tables field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MicrosoftGraphWorkbookWorksheet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Charts** | Pointer to [**[]MicrosoftGraphWorkbookChart**](microsoft.graph.workbookChart.md) |  | [optional] 
**Names** | Pointer to [**[]MicrosoftGraphWorkbookNamedItem**](microsoft.graph.workbookNamedItem.md) |  | [optional] 
**PivotTables** | Pointer to [**[]MicrosoftGraphWorkbookPivotTable**](microsoft.graph.workbookPivotTable.md) |  | [optional] 
**Protection** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheetProtection**](anyOf&lt;microsoft.graph.workbookWorksheetProtection&gt;.md) |  | [optional] 
**Tables** | Pointer to [**[]MicrosoftGraphWorkbookTable**](microsoft.graph.workbookTable.md) |  | [optional] 

## Methods

### GetId

`func (o *MicrosoftGraphWorkbookWorksheet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookWorksheet) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookWorksheet) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetName

`func (o *MicrosoftGraphWorkbookWorksheet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MicrosoftGraphWorkbookWorksheet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MicrosoftGraphWorkbookWorksheet) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *MicrosoftGraphWorkbookWorksheet) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetPosition

`func (o *MicrosoftGraphWorkbookWorksheet) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetPositionOk() (int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPosition

`func (o *MicrosoftGraphWorkbookWorksheet) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPosition

`func (o *MicrosoftGraphWorkbookWorksheet) SetPosition(v int32)`

SetPosition gets a reference to the given int32 and assigns it to the Position field.

### GetVisibility

`func (o *MicrosoftGraphWorkbookWorksheet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetVisibilityOk() (string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVisibility

`func (o *MicrosoftGraphWorkbookWorksheet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibility

`func (o *MicrosoftGraphWorkbookWorksheet) SetVisibility(v string)`

SetVisibility gets a reference to the given string and assigns it to the Visibility field.

### GetCharts

`func (o *MicrosoftGraphWorkbookWorksheet) GetCharts() []MicrosoftGraphWorkbookChart`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetChartsOk() ([]MicrosoftGraphWorkbookChart, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCharts

`func (o *MicrosoftGraphWorkbookWorksheet) HasCharts() bool`

HasCharts returns a boolean if a field has been set.

### SetCharts

`func (o *MicrosoftGraphWorkbookWorksheet) SetCharts(v []MicrosoftGraphWorkbookChart)`

SetCharts gets a reference to the given []MicrosoftGraphWorkbookChart and assigns it to the Charts field.

### GetNames

`func (o *MicrosoftGraphWorkbookWorksheet) GetNames() []MicrosoftGraphWorkbookNamedItem`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetNamesOk() ([]MicrosoftGraphWorkbookNamedItem, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNames

`func (o *MicrosoftGraphWorkbookWorksheet) HasNames() bool`

HasNames returns a boolean if a field has been set.

### SetNames

`func (o *MicrosoftGraphWorkbookWorksheet) SetNames(v []MicrosoftGraphWorkbookNamedItem)`

SetNames gets a reference to the given []MicrosoftGraphWorkbookNamedItem and assigns it to the Names field.

### GetPivotTables

`func (o *MicrosoftGraphWorkbookWorksheet) GetPivotTables() []MicrosoftGraphWorkbookPivotTable`

GetPivotTables returns the PivotTables field if non-nil, zero value otherwise.

### GetPivotTablesOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetPivotTablesOk() ([]MicrosoftGraphWorkbookPivotTable, bool)`

GetPivotTablesOk returns a tuple with the PivotTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPivotTables

`func (o *MicrosoftGraphWorkbookWorksheet) HasPivotTables() bool`

HasPivotTables returns a boolean if a field has been set.

### SetPivotTables

`func (o *MicrosoftGraphWorkbookWorksheet) SetPivotTables(v []MicrosoftGraphWorkbookPivotTable)`

SetPivotTables gets a reference to the given []MicrosoftGraphWorkbookPivotTable and assigns it to the PivotTables field.

### GetProtection

`func (o *MicrosoftGraphWorkbookWorksheet) GetProtection() AnyOfmicrosoftGraphWorkbookWorksheetProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetProtectionOk() (AnyOfmicrosoftGraphWorkbookWorksheetProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtection

`func (o *MicrosoftGraphWorkbookWorksheet) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### SetProtection

`func (o *MicrosoftGraphWorkbookWorksheet) SetProtection(v AnyOfmicrosoftGraphWorkbookWorksheetProtection)`

SetProtection gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheetProtection and assigns it to the Protection field.

### SetProtectionExplicitNull

`func (o *MicrosoftGraphWorkbookWorksheet) SetProtectionExplicitNull(b bool)`

SetProtectionExplicitNull (un)sets Protection to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Protection value is set to nil even if false is passed
### GetTables

`func (o *MicrosoftGraphWorkbookWorksheet) GetTables() []MicrosoftGraphWorkbookTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *MicrosoftGraphWorkbookWorksheet) GetTablesOk() ([]MicrosoftGraphWorkbookTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTables

`func (o *MicrosoftGraphWorkbookWorksheet) HasTables() bool`

HasTables returns a boolean if a field has been set.

### SetTables

`func (o *MicrosoftGraphWorkbookWorksheet) SetTables(v []MicrosoftGraphWorkbookTable)`

SetTables gets a reference to the given []MicrosoftGraphWorkbookTable and assigns it to the Tables field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



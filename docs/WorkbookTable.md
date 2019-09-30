# WorkbookTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HighlightFirstColumn** | Pointer to **bool** |  | [optional] 
**HighlightLastColumn** | Pointer to **bool** |  | [optional] 
**LegacyId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ShowBandedColumns** | Pointer to **bool** |  | [optional] 
**ShowBandedRows** | Pointer to **bool** |  | [optional] 
**ShowFilterButton** | Pointer to **bool** |  | [optional] 
**ShowHeaders** | Pointer to **bool** |  | [optional] 
**ShowTotals** | Pointer to **bool** |  | [optional] 
**Style** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]MicrosoftGraphWorkbookTableColumn**](microsoft.graph.workbookTableColumn.md) |  | [optional] 
**Rows** | Pointer to [**[]MicrosoftGraphWorkbookTableRow**](microsoft.graph.workbookTableRow.md) |  | [optional] 
**Sort** | Pointer to [**AnyOfmicrosoftGraphWorkbookTableSort**](anyOf&lt;microsoft.graph.workbookTableSort&gt;.md) |  | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) |  | [optional] 

## Methods

### GetHighlightFirstColumn

`func (o *WorkbookTable) GetHighlightFirstColumn() bool`

GetHighlightFirstColumn returns the HighlightFirstColumn field if non-nil, zero value otherwise.

### GetHighlightFirstColumnOk

`func (o *WorkbookTable) GetHighlightFirstColumnOk() (bool, bool)`

GetHighlightFirstColumnOk returns a tuple with the HighlightFirstColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHighlightFirstColumn

`func (o *WorkbookTable) HasHighlightFirstColumn() bool`

HasHighlightFirstColumn returns a boolean if a field has been set.

### SetHighlightFirstColumn

`func (o *WorkbookTable) SetHighlightFirstColumn(v bool)`

SetHighlightFirstColumn gets a reference to the given bool and assigns it to the HighlightFirstColumn field.

### GetHighlightLastColumn

`func (o *WorkbookTable) GetHighlightLastColumn() bool`

GetHighlightLastColumn returns the HighlightLastColumn field if non-nil, zero value otherwise.

### GetHighlightLastColumnOk

`func (o *WorkbookTable) GetHighlightLastColumnOk() (bool, bool)`

GetHighlightLastColumnOk returns a tuple with the HighlightLastColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHighlightLastColumn

`func (o *WorkbookTable) HasHighlightLastColumn() bool`

HasHighlightLastColumn returns a boolean if a field has been set.

### SetHighlightLastColumn

`func (o *WorkbookTable) SetHighlightLastColumn(v bool)`

SetHighlightLastColumn gets a reference to the given bool and assigns it to the HighlightLastColumn field.

### GetLegacyId

`func (o *WorkbookTable) GetLegacyId() string`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *WorkbookTable) GetLegacyIdOk() (string, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLegacyId

`func (o *WorkbookTable) HasLegacyId() bool`

HasLegacyId returns a boolean if a field has been set.

### SetLegacyId

`func (o *WorkbookTable) SetLegacyId(v string)`

SetLegacyId gets a reference to the given string and assigns it to the LegacyId field.

### SetLegacyIdExplicitNull

`func (o *WorkbookTable) SetLegacyIdExplicitNull(b bool)`

SetLegacyIdExplicitNull (un)sets LegacyId to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The LegacyId value is set to nil even if false is passed
### GetName

`func (o *WorkbookTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkbookTable) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *WorkbookTable) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *WorkbookTable) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### SetNameExplicitNull

`func (o *WorkbookTable) SetNameExplicitNull(b bool)`

SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Name value is set to nil even if false is passed
### GetShowBandedColumns

`func (o *WorkbookTable) GetShowBandedColumns() bool`

GetShowBandedColumns returns the ShowBandedColumns field if non-nil, zero value otherwise.

### GetShowBandedColumnsOk

`func (o *WorkbookTable) GetShowBandedColumnsOk() (bool, bool)`

GetShowBandedColumnsOk returns a tuple with the ShowBandedColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowBandedColumns

`func (o *WorkbookTable) HasShowBandedColumns() bool`

HasShowBandedColumns returns a boolean if a field has been set.

### SetShowBandedColumns

`func (o *WorkbookTable) SetShowBandedColumns(v bool)`

SetShowBandedColumns gets a reference to the given bool and assigns it to the ShowBandedColumns field.

### GetShowBandedRows

`func (o *WorkbookTable) GetShowBandedRows() bool`

GetShowBandedRows returns the ShowBandedRows field if non-nil, zero value otherwise.

### GetShowBandedRowsOk

`func (o *WorkbookTable) GetShowBandedRowsOk() (bool, bool)`

GetShowBandedRowsOk returns a tuple with the ShowBandedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowBandedRows

`func (o *WorkbookTable) HasShowBandedRows() bool`

HasShowBandedRows returns a boolean if a field has been set.

### SetShowBandedRows

`func (o *WorkbookTable) SetShowBandedRows(v bool)`

SetShowBandedRows gets a reference to the given bool and assigns it to the ShowBandedRows field.

### GetShowFilterButton

`func (o *WorkbookTable) GetShowFilterButton() bool`

GetShowFilterButton returns the ShowFilterButton field if non-nil, zero value otherwise.

### GetShowFilterButtonOk

`func (o *WorkbookTable) GetShowFilterButtonOk() (bool, bool)`

GetShowFilterButtonOk returns a tuple with the ShowFilterButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowFilterButton

`func (o *WorkbookTable) HasShowFilterButton() bool`

HasShowFilterButton returns a boolean if a field has been set.

### SetShowFilterButton

`func (o *WorkbookTable) SetShowFilterButton(v bool)`

SetShowFilterButton gets a reference to the given bool and assigns it to the ShowFilterButton field.

### GetShowHeaders

`func (o *WorkbookTable) GetShowHeaders() bool`

GetShowHeaders returns the ShowHeaders field if non-nil, zero value otherwise.

### GetShowHeadersOk

`func (o *WorkbookTable) GetShowHeadersOk() (bool, bool)`

GetShowHeadersOk returns a tuple with the ShowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowHeaders

`func (o *WorkbookTable) HasShowHeaders() bool`

HasShowHeaders returns a boolean if a field has been set.

### SetShowHeaders

`func (o *WorkbookTable) SetShowHeaders(v bool)`

SetShowHeaders gets a reference to the given bool and assigns it to the ShowHeaders field.

### GetShowTotals

`func (o *WorkbookTable) GetShowTotals() bool`

GetShowTotals returns the ShowTotals field if non-nil, zero value otherwise.

### GetShowTotalsOk

`func (o *WorkbookTable) GetShowTotalsOk() (bool, bool)`

GetShowTotalsOk returns a tuple with the ShowTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasShowTotals

`func (o *WorkbookTable) HasShowTotals() bool`

HasShowTotals returns a boolean if a field has been set.

### SetShowTotals

`func (o *WorkbookTable) SetShowTotals(v bool)`

SetShowTotals gets a reference to the given bool and assigns it to the ShowTotals field.

### GetStyle

`func (o *WorkbookTable) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *WorkbookTable) GetStyleOk() (string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStyle

`func (o *WorkbookTable) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyle

`func (o *WorkbookTable) SetStyle(v string)`

SetStyle gets a reference to the given string and assigns it to the Style field.

### SetStyleExplicitNull

`func (o *WorkbookTable) SetStyleExplicitNull(b bool)`

SetStyleExplicitNull (un)sets Style to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Style value is set to nil even if false is passed
### GetColumns

`func (o *WorkbookTable) GetColumns() []MicrosoftGraphWorkbookTableColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *WorkbookTable) GetColumnsOk() ([]MicrosoftGraphWorkbookTableColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumns

`func (o *WorkbookTable) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumns

`func (o *WorkbookTable) SetColumns(v []MicrosoftGraphWorkbookTableColumn)`

SetColumns gets a reference to the given []MicrosoftGraphWorkbookTableColumn and assigns it to the Columns field.

### GetRows

`func (o *WorkbookTable) GetRows() []MicrosoftGraphWorkbookTableRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WorkbookTable) GetRowsOk() ([]MicrosoftGraphWorkbookTableRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRows

`func (o *WorkbookTable) HasRows() bool`

HasRows returns a boolean if a field has been set.

### SetRows

`func (o *WorkbookTable) SetRows(v []MicrosoftGraphWorkbookTableRow)`

SetRows gets a reference to the given []MicrosoftGraphWorkbookTableRow and assigns it to the Rows field.

### GetSort

`func (o *WorkbookTable) GetSort() AnyOfmicrosoftGraphWorkbookTableSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *WorkbookTable) GetSortOk() (AnyOfmicrosoftGraphWorkbookTableSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSort

`func (o *WorkbookTable) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSort

`func (o *WorkbookTable) SetSort(v AnyOfmicrosoftGraphWorkbookTableSort)`

SetSort gets a reference to the given AnyOfmicrosoftGraphWorkbookTableSort and assigns it to the Sort field.

### SetSortExplicitNull

`func (o *WorkbookTable) SetSortExplicitNull(b bool)`

SetSortExplicitNull (un)sets Sort to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sort value is set to nil even if false is passed
### GetWorksheet

`func (o *WorkbookTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *WorkbookTable) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorksheet

`func (o *WorkbookTable) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheet

`func (o *WorkbookTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.

### SetWorksheetExplicitNull

`func (o *WorkbookTable) SetWorksheetExplicitNull(b bool)`

SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Worksheet value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WorkbookRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AddressLocal** | Pointer to **string** |  | [optional] 
**CellCount** | Pointer to **int32** |  | [optional] 
**ColumnCount** | Pointer to **int32** |  | [optional] 
**ColumnHidden** | Pointer to **bool** |  | [optional] 
**ColumnIndex** | Pointer to **int32** |  | [optional] 
**Formulas** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**FormulasLocal** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**FormulasR1C1** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**NumberFormat** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**RowCount** | Pointer to **int32** |  | [optional] 
**RowHidden** | Pointer to **bool** |  | [optional] 
**RowIndex** | Pointer to **int32** |  | [optional] 
**Text** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**ValueTypes** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphWorkbookRangeFormat**](anyOf&lt;microsoft.graph.workbookRangeFormat&gt;.md) |  | [optional] 
**Sort** | Pointer to [**AnyOfmicrosoftGraphWorkbookRangeSort**](anyOf&lt;microsoft.graph.workbookRangeSort&gt;.md) |  | [optional] 
**Worksheet** | Pointer to [**AnyOfmicrosoftGraphWorkbookWorksheet**](anyOf&lt;microsoft.graph.workbookWorksheet&gt;.md) |  | [optional] 

## Methods

### GetAddress

`func (o *WorkbookRange) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WorkbookRange) GetAddressOk() (string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddress

`func (o *WorkbookRange) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddress

`func (o *WorkbookRange) SetAddress(v string)`

SetAddress gets a reference to the given string and assigns it to the Address field.

### SetAddressExplicitNull

`func (o *WorkbookRange) SetAddressExplicitNull(b bool)`

SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Address value is set to nil even if false is passed
### GetAddressLocal

`func (o *WorkbookRange) GetAddressLocal() string`

GetAddressLocal returns the AddressLocal field if non-nil, zero value otherwise.

### GetAddressLocalOk

`func (o *WorkbookRange) GetAddressLocalOk() (string, bool)`

GetAddressLocalOk returns a tuple with the AddressLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddressLocal

`func (o *WorkbookRange) HasAddressLocal() bool`

HasAddressLocal returns a boolean if a field has been set.

### SetAddressLocal

`func (o *WorkbookRange) SetAddressLocal(v string)`

SetAddressLocal gets a reference to the given string and assigns it to the AddressLocal field.

### SetAddressLocalExplicitNull

`func (o *WorkbookRange) SetAddressLocalExplicitNull(b bool)`

SetAddressLocalExplicitNull (un)sets AddressLocal to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AddressLocal value is set to nil even if false is passed
### GetCellCount

`func (o *WorkbookRange) GetCellCount() int32`

GetCellCount returns the CellCount field if non-nil, zero value otherwise.

### GetCellCountOk

`func (o *WorkbookRange) GetCellCountOk() (int32, bool)`

GetCellCountOk returns a tuple with the CellCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellCount

`func (o *WorkbookRange) HasCellCount() bool`

HasCellCount returns a boolean if a field has been set.

### SetCellCount

`func (o *WorkbookRange) SetCellCount(v int32)`

SetCellCount gets a reference to the given int32 and assigns it to the CellCount field.

### GetColumnCount

`func (o *WorkbookRange) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *WorkbookRange) GetColumnCountOk() (int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnCount

`func (o *WorkbookRange) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### SetColumnCount

`func (o *WorkbookRange) SetColumnCount(v int32)`

SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.

### GetColumnHidden

`func (o *WorkbookRange) GetColumnHidden() bool`

GetColumnHidden returns the ColumnHidden field if non-nil, zero value otherwise.

### GetColumnHiddenOk

`func (o *WorkbookRange) GetColumnHiddenOk() (bool, bool)`

GetColumnHiddenOk returns a tuple with the ColumnHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnHidden

`func (o *WorkbookRange) HasColumnHidden() bool`

HasColumnHidden returns a boolean if a field has been set.

### SetColumnHidden

`func (o *WorkbookRange) SetColumnHidden(v bool)`

SetColumnHidden gets a reference to the given bool and assigns it to the ColumnHidden field.

### SetColumnHiddenExplicitNull

`func (o *WorkbookRange) SetColumnHiddenExplicitNull(b bool)`

SetColumnHiddenExplicitNull (un)sets ColumnHidden to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ColumnHidden value is set to nil even if false is passed
### GetColumnIndex

`func (o *WorkbookRange) GetColumnIndex() int32`

GetColumnIndex returns the ColumnIndex field if non-nil, zero value otherwise.

### GetColumnIndexOk

`func (o *WorkbookRange) GetColumnIndexOk() (int32, bool)`

GetColumnIndexOk returns a tuple with the ColumnIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnIndex

`func (o *WorkbookRange) HasColumnIndex() bool`

HasColumnIndex returns a boolean if a field has been set.

### SetColumnIndex

`func (o *WorkbookRange) SetColumnIndex(v int32)`

SetColumnIndex gets a reference to the given int32 and assigns it to the ColumnIndex field.

### GetFormulas

`func (o *WorkbookRange) GetFormulas() AnyOfobject`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *WorkbookRange) GetFormulasOk() (AnyOfobject, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulas

`func (o *WorkbookRange) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### SetFormulas

`func (o *WorkbookRange) SetFormulas(v AnyOfobject)`

SetFormulas gets a reference to the given AnyOfobject and assigns it to the Formulas field.

### SetFormulasExplicitNull

`func (o *WorkbookRange) SetFormulasExplicitNull(b bool)`

SetFormulasExplicitNull (un)sets Formulas to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Formulas value is set to nil even if false is passed
### GetFormulasLocal

`func (o *WorkbookRange) GetFormulasLocal() AnyOfobject`

GetFormulasLocal returns the FormulasLocal field if non-nil, zero value otherwise.

### GetFormulasLocalOk

`func (o *WorkbookRange) GetFormulasLocalOk() (AnyOfobject, bool)`

GetFormulasLocalOk returns a tuple with the FormulasLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulasLocal

`func (o *WorkbookRange) HasFormulasLocal() bool`

HasFormulasLocal returns a boolean if a field has been set.

### SetFormulasLocal

`func (o *WorkbookRange) SetFormulasLocal(v AnyOfobject)`

SetFormulasLocal gets a reference to the given AnyOfobject and assigns it to the FormulasLocal field.

### SetFormulasLocalExplicitNull

`func (o *WorkbookRange) SetFormulasLocalExplicitNull(b bool)`

SetFormulasLocalExplicitNull (un)sets FormulasLocal to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FormulasLocal value is set to nil even if false is passed
### GetFormulasR1C1

`func (o *WorkbookRange) GetFormulasR1C1() AnyOfobject`

GetFormulasR1C1 returns the FormulasR1C1 field if non-nil, zero value otherwise.

### GetFormulasR1C1Ok

`func (o *WorkbookRange) GetFormulasR1C1Ok() (AnyOfobject, bool)`

GetFormulasR1C1Ok returns a tuple with the FormulasR1C1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulasR1C1

`func (o *WorkbookRange) HasFormulasR1C1() bool`

HasFormulasR1C1 returns a boolean if a field has been set.

### SetFormulasR1C1

`func (o *WorkbookRange) SetFormulasR1C1(v AnyOfobject)`

SetFormulasR1C1 gets a reference to the given AnyOfobject and assigns it to the FormulasR1C1 field.

### SetFormulasR1C1ExplicitNull

`func (o *WorkbookRange) SetFormulasR1C1ExplicitNull(b bool)`

SetFormulasR1C1ExplicitNull (un)sets FormulasR1C1 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FormulasR1C1 value is set to nil even if false is passed
### GetHidden

`func (o *WorkbookRange) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *WorkbookRange) GetHiddenOk() (bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHidden

`func (o *WorkbookRange) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHidden

`func (o *WorkbookRange) SetHidden(v bool)`

SetHidden gets a reference to the given bool and assigns it to the Hidden field.

### SetHiddenExplicitNull

`func (o *WorkbookRange) SetHiddenExplicitNull(b bool)`

SetHiddenExplicitNull (un)sets Hidden to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Hidden value is set to nil even if false is passed
### GetNumberFormat

`func (o *WorkbookRange) GetNumberFormat() AnyOfobject`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *WorkbookRange) GetNumberFormatOk() (AnyOfobject, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberFormat

`func (o *WorkbookRange) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### SetNumberFormat

`func (o *WorkbookRange) SetNumberFormat(v AnyOfobject)`

SetNumberFormat gets a reference to the given AnyOfobject and assigns it to the NumberFormat field.

### SetNumberFormatExplicitNull

`func (o *WorkbookRange) SetNumberFormatExplicitNull(b bool)`

SetNumberFormatExplicitNull (un)sets NumberFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NumberFormat value is set to nil even if false is passed
### GetRowCount

`func (o *WorkbookRange) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *WorkbookRange) GetRowCountOk() (int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowCount

`func (o *WorkbookRange) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCount

`func (o *WorkbookRange) SetRowCount(v int32)`

SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.

### GetRowHidden

`func (o *WorkbookRange) GetRowHidden() bool`

GetRowHidden returns the RowHidden field if non-nil, zero value otherwise.

### GetRowHiddenOk

`func (o *WorkbookRange) GetRowHiddenOk() (bool, bool)`

GetRowHiddenOk returns a tuple with the RowHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowHidden

`func (o *WorkbookRange) HasRowHidden() bool`

HasRowHidden returns a boolean if a field has been set.

### SetRowHidden

`func (o *WorkbookRange) SetRowHidden(v bool)`

SetRowHidden gets a reference to the given bool and assigns it to the RowHidden field.

### SetRowHiddenExplicitNull

`func (o *WorkbookRange) SetRowHiddenExplicitNull(b bool)`

SetRowHiddenExplicitNull (un)sets RowHidden to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RowHidden value is set to nil even if false is passed
### GetRowIndex

`func (o *WorkbookRange) GetRowIndex() int32`

GetRowIndex returns the RowIndex field if non-nil, zero value otherwise.

### GetRowIndexOk

`func (o *WorkbookRange) GetRowIndexOk() (int32, bool)`

GetRowIndexOk returns a tuple with the RowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowIndex

`func (o *WorkbookRange) HasRowIndex() bool`

HasRowIndex returns a boolean if a field has been set.

### SetRowIndex

`func (o *WorkbookRange) SetRowIndex(v int32)`

SetRowIndex gets a reference to the given int32 and assigns it to the RowIndex field.

### GetText

`func (o *WorkbookRange) GetText() AnyOfobject`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WorkbookRange) GetTextOk() (AnyOfobject, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasText

`func (o *WorkbookRange) HasText() bool`

HasText returns a boolean if a field has been set.

### SetText

`func (o *WorkbookRange) SetText(v AnyOfobject)`

SetText gets a reference to the given AnyOfobject and assigns it to the Text field.

### SetTextExplicitNull

`func (o *WorkbookRange) SetTextExplicitNull(b bool)`

SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Text value is set to nil even if false is passed
### GetValueTypes

`func (o *WorkbookRange) GetValueTypes() AnyOfobject`

GetValueTypes returns the ValueTypes field if non-nil, zero value otherwise.

### GetValueTypesOk

`func (o *WorkbookRange) GetValueTypesOk() (AnyOfobject, bool)`

GetValueTypesOk returns a tuple with the ValueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValueTypes

`func (o *WorkbookRange) HasValueTypes() bool`

HasValueTypes returns a boolean if a field has been set.

### SetValueTypes

`func (o *WorkbookRange) SetValueTypes(v AnyOfobject)`

SetValueTypes gets a reference to the given AnyOfobject and assigns it to the ValueTypes field.

### SetValueTypesExplicitNull

`func (o *WorkbookRange) SetValueTypesExplicitNull(b bool)`

SetValueTypesExplicitNull (un)sets ValueTypes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ValueTypes value is set to nil even if false is passed
### GetValues

`func (o *WorkbookRange) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WorkbookRange) GetValuesOk() (AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *WorkbookRange) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *WorkbookRange) SetValues(v AnyOfobject)`

SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.

### SetValuesExplicitNull

`func (o *WorkbookRange) SetValuesExplicitNull(b bool)`

SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Values value is set to nil even if false is passed
### GetFormat

`func (o *WorkbookRange) GetFormat() AnyOfmicrosoftGraphWorkbookRangeFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkbookRange) GetFormatOk() (AnyOfmicrosoftGraphWorkbookRangeFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormat

`func (o *WorkbookRange) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormat

`func (o *WorkbookRange) SetFormat(v AnyOfmicrosoftGraphWorkbookRangeFormat)`

SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFormat and assigns it to the Format field.

### SetFormatExplicitNull

`func (o *WorkbookRange) SetFormatExplicitNull(b bool)`

SetFormatExplicitNull (un)sets Format to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Format value is set to nil even if false is passed
### GetSort

`func (o *WorkbookRange) GetSort() AnyOfmicrosoftGraphWorkbookRangeSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *WorkbookRange) GetSortOk() (AnyOfmicrosoftGraphWorkbookRangeSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSort

`func (o *WorkbookRange) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSort

`func (o *WorkbookRange) SetSort(v AnyOfmicrosoftGraphWorkbookRangeSort)`

SetSort gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeSort and assigns it to the Sort field.

### SetSortExplicitNull

`func (o *WorkbookRange) SetSortExplicitNull(b bool)`

SetSortExplicitNull (un)sets Sort to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sort value is set to nil even if false is passed
### GetWorksheet

`func (o *WorkbookRange) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *WorkbookRange) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorksheet

`func (o *WorkbookRange) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheet

`func (o *WorkbookRange) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.

### SetWorksheetExplicitNull

`func (o *WorkbookRange) SetWorksheetExplicitNull(b bool)`

SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Worksheet value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



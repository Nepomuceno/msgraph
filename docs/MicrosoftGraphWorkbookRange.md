# MicrosoftGraphWorkbookRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphWorkbookRange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookRange) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookRange) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookRange) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetAddress

`func (o *MicrosoftGraphWorkbookRange) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphWorkbookRange) GetAddressOk() (string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddress

`func (o *MicrosoftGraphWorkbookRange) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddress

`func (o *MicrosoftGraphWorkbookRange) SetAddress(v string)`

SetAddress gets a reference to the given string and assigns it to the Address field.

### SetAddressExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetAddressExplicitNull(b bool)`

SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Address value is set to nil even if false is passed
### GetAddressLocal

`func (o *MicrosoftGraphWorkbookRange) GetAddressLocal() string`

GetAddressLocal returns the AddressLocal field if non-nil, zero value otherwise.

### GetAddressLocalOk

`func (o *MicrosoftGraphWorkbookRange) GetAddressLocalOk() (string, bool)`

GetAddressLocalOk returns a tuple with the AddressLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddressLocal

`func (o *MicrosoftGraphWorkbookRange) HasAddressLocal() bool`

HasAddressLocal returns a boolean if a field has been set.

### SetAddressLocal

`func (o *MicrosoftGraphWorkbookRange) SetAddressLocal(v string)`

SetAddressLocal gets a reference to the given string and assigns it to the AddressLocal field.

### SetAddressLocalExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetAddressLocalExplicitNull(b bool)`

SetAddressLocalExplicitNull (un)sets AddressLocal to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The AddressLocal value is set to nil even if false is passed
### GetCellCount

`func (o *MicrosoftGraphWorkbookRange) GetCellCount() int32`

GetCellCount returns the CellCount field if non-nil, zero value otherwise.

### GetCellCountOk

`func (o *MicrosoftGraphWorkbookRange) GetCellCountOk() (int32, bool)`

GetCellCountOk returns a tuple with the CellCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellCount

`func (o *MicrosoftGraphWorkbookRange) HasCellCount() bool`

HasCellCount returns a boolean if a field has been set.

### SetCellCount

`func (o *MicrosoftGraphWorkbookRange) SetCellCount(v int32)`

SetCellCount gets a reference to the given int32 and assigns it to the CellCount field.

### GetColumnCount

`func (o *MicrosoftGraphWorkbookRange) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *MicrosoftGraphWorkbookRange) GetColumnCountOk() (int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnCount

`func (o *MicrosoftGraphWorkbookRange) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### SetColumnCount

`func (o *MicrosoftGraphWorkbookRange) SetColumnCount(v int32)`

SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.

### GetColumnHidden

`func (o *MicrosoftGraphWorkbookRange) GetColumnHidden() bool`

GetColumnHidden returns the ColumnHidden field if non-nil, zero value otherwise.

### GetColumnHiddenOk

`func (o *MicrosoftGraphWorkbookRange) GetColumnHiddenOk() (bool, bool)`

GetColumnHiddenOk returns a tuple with the ColumnHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnHidden

`func (o *MicrosoftGraphWorkbookRange) HasColumnHidden() bool`

HasColumnHidden returns a boolean if a field has been set.

### SetColumnHidden

`func (o *MicrosoftGraphWorkbookRange) SetColumnHidden(v bool)`

SetColumnHidden gets a reference to the given bool and assigns it to the ColumnHidden field.

### SetColumnHiddenExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetColumnHiddenExplicitNull(b bool)`

SetColumnHiddenExplicitNull (un)sets ColumnHidden to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ColumnHidden value is set to nil even if false is passed
### GetColumnIndex

`func (o *MicrosoftGraphWorkbookRange) GetColumnIndex() int32`

GetColumnIndex returns the ColumnIndex field if non-nil, zero value otherwise.

### GetColumnIndexOk

`func (o *MicrosoftGraphWorkbookRange) GetColumnIndexOk() (int32, bool)`

GetColumnIndexOk returns a tuple with the ColumnIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnIndex

`func (o *MicrosoftGraphWorkbookRange) HasColumnIndex() bool`

HasColumnIndex returns a boolean if a field has been set.

### SetColumnIndex

`func (o *MicrosoftGraphWorkbookRange) SetColumnIndex(v int32)`

SetColumnIndex gets a reference to the given int32 and assigns it to the ColumnIndex field.

### GetFormulas

`func (o *MicrosoftGraphWorkbookRange) GetFormulas() AnyOfobject`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *MicrosoftGraphWorkbookRange) GetFormulasOk() (AnyOfobject, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulas

`func (o *MicrosoftGraphWorkbookRange) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### SetFormulas

`func (o *MicrosoftGraphWorkbookRange) SetFormulas(v AnyOfobject)`

SetFormulas gets a reference to the given AnyOfobject and assigns it to the Formulas field.

### SetFormulasExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetFormulasExplicitNull(b bool)`

SetFormulasExplicitNull (un)sets Formulas to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Formulas value is set to nil even if false is passed
### GetFormulasLocal

`func (o *MicrosoftGraphWorkbookRange) GetFormulasLocal() AnyOfobject`

GetFormulasLocal returns the FormulasLocal field if non-nil, zero value otherwise.

### GetFormulasLocalOk

`func (o *MicrosoftGraphWorkbookRange) GetFormulasLocalOk() (AnyOfobject, bool)`

GetFormulasLocalOk returns a tuple with the FormulasLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulasLocal

`func (o *MicrosoftGraphWorkbookRange) HasFormulasLocal() bool`

HasFormulasLocal returns a boolean if a field has been set.

### SetFormulasLocal

`func (o *MicrosoftGraphWorkbookRange) SetFormulasLocal(v AnyOfobject)`

SetFormulasLocal gets a reference to the given AnyOfobject and assigns it to the FormulasLocal field.

### SetFormulasLocalExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetFormulasLocalExplicitNull(b bool)`

SetFormulasLocalExplicitNull (un)sets FormulasLocal to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FormulasLocal value is set to nil even if false is passed
### GetFormulasR1C1

`func (o *MicrosoftGraphWorkbookRange) GetFormulasR1C1() AnyOfobject`

GetFormulasR1C1 returns the FormulasR1C1 field if non-nil, zero value otherwise.

### GetFormulasR1C1Ok

`func (o *MicrosoftGraphWorkbookRange) GetFormulasR1C1Ok() (AnyOfobject, bool)`

GetFormulasR1C1Ok returns a tuple with the FormulasR1C1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulasR1C1

`func (o *MicrosoftGraphWorkbookRange) HasFormulasR1C1() bool`

HasFormulasR1C1 returns a boolean if a field has been set.

### SetFormulasR1C1

`func (o *MicrosoftGraphWorkbookRange) SetFormulasR1C1(v AnyOfobject)`

SetFormulasR1C1 gets a reference to the given AnyOfobject and assigns it to the FormulasR1C1 field.

### SetFormulasR1C1ExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetFormulasR1C1ExplicitNull(b bool)`

SetFormulasR1C1ExplicitNull (un)sets FormulasR1C1 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FormulasR1C1 value is set to nil even if false is passed
### GetHidden

`func (o *MicrosoftGraphWorkbookRange) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MicrosoftGraphWorkbookRange) GetHiddenOk() (bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHidden

`func (o *MicrosoftGraphWorkbookRange) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHidden

`func (o *MicrosoftGraphWorkbookRange) SetHidden(v bool)`

SetHidden gets a reference to the given bool and assigns it to the Hidden field.

### SetHiddenExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetHiddenExplicitNull(b bool)`

SetHiddenExplicitNull (un)sets Hidden to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Hidden value is set to nil even if false is passed
### GetNumberFormat

`func (o *MicrosoftGraphWorkbookRange) GetNumberFormat() AnyOfobject`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *MicrosoftGraphWorkbookRange) GetNumberFormatOk() (AnyOfobject, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberFormat

`func (o *MicrosoftGraphWorkbookRange) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### SetNumberFormat

`func (o *MicrosoftGraphWorkbookRange) SetNumberFormat(v AnyOfobject)`

SetNumberFormat gets a reference to the given AnyOfobject and assigns it to the NumberFormat field.

### SetNumberFormatExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetNumberFormatExplicitNull(b bool)`

SetNumberFormatExplicitNull (un)sets NumberFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NumberFormat value is set to nil even if false is passed
### GetRowCount

`func (o *MicrosoftGraphWorkbookRange) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *MicrosoftGraphWorkbookRange) GetRowCountOk() (int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowCount

`func (o *MicrosoftGraphWorkbookRange) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCount

`func (o *MicrosoftGraphWorkbookRange) SetRowCount(v int32)`

SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.

### GetRowHidden

`func (o *MicrosoftGraphWorkbookRange) GetRowHidden() bool`

GetRowHidden returns the RowHidden field if non-nil, zero value otherwise.

### GetRowHiddenOk

`func (o *MicrosoftGraphWorkbookRange) GetRowHiddenOk() (bool, bool)`

GetRowHiddenOk returns a tuple with the RowHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowHidden

`func (o *MicrosoftGraphWorkbookRange) HasRowHidden() bool`

HasRowHidden returns a boolean if a field has been set.

### SetRowHidden

`func (o *MicrosoftGraphWorkbookRange) SetRowHidden(v bool)`

SetRowHidden gets a reference to the given bool and assigns it to the RowHidden field.

### SetRowHiddenExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetRowHiddenExplicitNull(b bool)`

SetRowHiddenExplicitNull (un)sets RowHidden to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The RowHidden value is set to nil even if false is passed
### GetRowIndex

`func (o *MicrosoftGraphWorkbookRange) GetRowIndex() int32`

GetRowIndex returns the RowIndex field if non-nil, zero value otherwise.

### GetRowIndexOk

`func (o *MicrosoftGraphWorkbookRange) GetRowIndexOk() (int32, bool)`

GetRowIndexOk returns a tuple with the RowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowIndex

`func (o *MicrosoftGraphWorkbookRange) HasRowIndex() bool`

HasRowIndex returns a boolean if a field has been set.

### SetRowIndex

`func (o *MicrosoftGraphWorkbookRange) SetRowIndex(v int32)`

SetRowIndex gets a reference to the given int32 and assigns it to the RowIndex field.

### GetText

`func (o *MicrosoftGraphWorkbookRange) GetText() AnyOfobject`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MicrosoftGraphWorkbookRange) GetTextOk() (AnyOfobject, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasText

`func (o *MicrosoftGraphWorkbookRange) HasText() bool`

HasText returns a boolean if a field has been set.

### SetText

`func (o *MicrosoftGraphWorkbookRange) SetText(v AnyOfobject)`

SetText gets a reference to the given AnyOfobject and assigns it to the Text field.

### SetTextExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetTextExplicitNull(b bool)`

SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Text value is set to nil even if false is passed
### GetValueTypes

`func (o *MicrosoftGraphWorkbookRange) GetValueTypes() AnyOfobject`

GetValueTypes returns the ValueTypes field if non-nil, zero value otherwise.

### GetValueTypesOk

`func (o *MicrosoftGraphWorkbookRange) GetValueTypesOk() (AnyOfobject, bool)`

GetValueTypesOk returns a tuple with the ValueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValueTypes

`func (o *MicrosoftGraphWorkbookRange) HasValueTypes() bool`

HasValueTypes returns a boolean if a field has been set.

### SetValueTypes

`func (o *MicrosoftGraphWorkbookRange) SetValueTypes(v AnyOfobject)`

SetValueTypes gets a reference to the given AnyOfobject and assigns it to the ValueTypes field.

### SetValueTypesExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetValueTypesExplicitNull(b bool)`

SetValueTypesExplicitNull (un)sets ValueTypes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ValueTypes value is set to nil even if false is passed
### GetValues

`func (o *MicrosoftGraphWorkbookRange) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MicrosoftGraphWorkbookRange) GetValuesOk() (AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *MicrosoftGraphWorkbookRange) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *MicrosoftGraphWorkbookRange) SetValues(v AnyOfobject)`

SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.

### SetValuesExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetValuesExplicitNull(b bool)`

SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Values value is set to nil even if false is passed
### GetFormat

`func (o *MicrosoftGraphWorkbookRange) GetFormat() AnyOfmicrosoftGraphWorkbookRangeFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphWorkbookRange) GetFormatOk() (AnyOfmicrosoftGraphWorkbookRangeFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormat

`func (o *MicrosoftGraphWorkbookRange) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormat

`func (o *MicrosoftGraphWorkbookRange) SetFormat(v AnyOfmicrosoftGraphWorkbookRangeFormat)`

SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFormat and assigns it to the Format field.

### SetFormatExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetFormatExplicitNull(b bool)`

SetFormatExplicitNull (un)sets Format to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Format value is set to nil even if false is passed
### GetSort

`func (o *MicrosoftGraphWorkbookRange) GetSort() AnyOfmicrosoftGraphWorkbookRangeSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *MicrosoftGraphWorkbookRange) GetSortOk() (AnyOfmicrosoftGraphWorkbookRangeSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSort

`func (o *MicrosoftGraphWorkbookRange) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSort

`func (o *MicrosoftGraphWorkbookRange) SetSort(v AnyOfmicrosoftGraphWorkbookRangeSort)`

SetSort gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeSort and assigns it to the Sort field.

### SetSortExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetSortExplicitNull(b bool)`

SetSortExplicitNull (un)sets Sort to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Sort value is set to nil even if false is passed
### GetWorksheet

`func (o *MicrosoftGraphWorkbookRange) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet`

GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.

### GetWorksheetOk

`func (o *MicrosoftGraphWorkbookRange) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool)`

GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWorksheet

`func (o *MicrosoftGraphWorkbookRange) HasWorksheet() bool`

HasWorksheet returns a boolean if a field has been set.

### SetWorksheet

`func (o *MicrosoftGraphWorkbookRange) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet)`

SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.

### SetWorksheetExplicitNull

`func (o *MicrosoftGraphWorkbookRange) SetWorksheetExplicitNull(b bool)`

SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Worksheet value is set to nil even if false is passed

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



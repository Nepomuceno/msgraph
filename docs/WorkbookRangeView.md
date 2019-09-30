# WorkbookRangeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellAddresses** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**ColumnCount** | Pointer to **int32** |  | [optional] 
**Formulas** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**FormulasLocal** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**FormulasR1C1** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**NumberFormat** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**RowCount** | Pointer to **int32** |  | [optional] 
**Text** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**ValueTypes** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 
**Rows** | Pointer to [**[]MicrosoftGraphWorkbookRangeView**](microsoft.graph.workbookRangeView.md) |  | [optional] 

## Methods

### GetCellAddresses

`func (o *WorkbookRangeView) GetCellAddresses() AnyOfobject`

GetCellAddresses returns the CellAddresses field if non-nil, zero value otherwise.

### GetCellAddressesOk

`func (o *WorkbookRangeView) GetCellAddressesOk() (AnyOfobject, bool)`

GetCellAddressesOk returns a tuple with the CellAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellAddresses

`func (o *WorkbookRangeView) HasCellAddresses() bool`

HasCellAddresses returns a boolean if a field has been set.

### SetCellAddresses

`func (o *WorkbookRangeView) SetCellAddresses(v AnyOfobject)`

SetCellAddresses gets a reference to the given AnyOfobject and assigns it to the CellAddresses field.

### SetCellAddressesExplicitNull

`func (o *WorkbookRangeView) SetCellAddressesExplicitNull(b bool)`

SetCellAddressesExplicitNull (un)sets CellAddresses to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CellAddresses value is set to nil even if false is passed
### GetColumnCount

`func (o *WorkbookRangeView) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *WorkbookRangeView) GetColumnCountOk() (int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnCount

`func (o *WorkbookRangeView) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### SetColumnCount

`func (o *WorkbookRangeView) SetColumnCount(v int32)`

SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.

### GetFormulas

`func (o *WorkbookRangeView) GetFormulas() AnyOfobject`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *WorkbookRangeView) GetFormulasOk() (AnyOfobject, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulas

`func (o *WorkbookRangeView) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### SetFormulas

`func (o *WorkbookRangeView) SetFormulas(v AnyOfobject)`

SetFormulas gets a reference to the given AnyOfobject and assigns it to the Formulas field.

### SetFormulasExplicitNull

`func (o *WorkbookRangeView) SetFormulasExplicitNull(b bool)`

SetFormulasExplicitNull (un)sets Formulas to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Formulas value is set to nil even if false is passed
### GetFormulasLocal

`func (o *WorkbookRangeView) GetFormulasLocal() AnyOfobject`

GetFormulasLocal returns the FormulasLocal field if non-nil, zero value otherwise.

### GetFormulasLocalOk

`func (o *WorkbookRangeView) GetFormulasLocalOk() (AnyOfobject, bool)`

GetFormulasLocalOk returns a tuple with the FormulasLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulasLocal

`func (o *WorkbookRangeView) HasFormulasLocal() bool`

HasFormulasLocal returns a boolean if a field has been set.

### SetFormulasLocal

`func (o *WorkbookRangeView) SetFormulasLocal(v AnyOfobject)`

SetFormulasLocal gets a reference to the given AnyOfobject and assigns it to the FormulasLocal field.

### SetFormulasLocalExplicitNull

`func (o *WorkbookRangeView) SetFormulasLocalExplicitNull(b bool)`

SetFormulasLocalExplicitNull (un)sets FormulasLocal to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FormulasLocal value is set to nil even if false is passed
### GetFormulasR1C1

`func (o *WorkbookRangeView) GetFormulasR1C1() AnyOfobject`

GetFormulasR1C1 returns the FormulasR1C1 field if non-nil, zero value otherwise.

### GetFormulasR1C1Ok

`func (o *WorkbookRangeView) GetFormulasR1C1Ok() (AnyOfobject, bool)`

GetFormulasR1C1Ok returns a tuple with the FormulasR1C1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulasR1C1

`func (o *WorkbookRangeView) HasFormulasR1C1() bool`

HasFormulasR1C1 returns a boolean if a field has been set.

### SetFormulasR1C1

`func (o *WorkbookRangeView) SetFormulasR1C1(v AnyOfobject)`

SetFormulasR1C1 gets a reference to the given AnyOfobject and assigns it to the FormulasR1C1 field.

### SetFormulasR1C1ExplicitNull

`func (o *WorkbookRangeView) SetFormulasR1C1ExplicitNull(b bool)`

SetFormulasR1C1ExplicitNull (un)sets FormulasR1C1 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FormulasR1C1 value is set to nil even if false is passed
### GetIndex

`func (o *WorkbookRangeView) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WorkbookRangeView) GetIndexOk() (int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndex

`func (o *WorkbookRangeView) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndex

`func (o *WorkbookRangeView) SetIndex(v int32)`

SetIndex gets a reference to the given int32 and assigns it to the Index field.

### GetNumberFormat

`func (o *WorkbookRangeView) GetNumberFormat() AnyOfobject`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *WorkbookRangeView) GetNumberFormatOk() (AnyOfobject, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberFormat

`func (o *WorkbookRangeView) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### SetNumberFormat

`func (o *WorkbookRangeView) SetNumberFormat(v AnyOfobject)`

SetNumberFormat gets a reference to the given AnyOfobject and assigns it to the NumberFormat field.

### SetNumberFormatExplicitNull

`func (o *WorkbookRangeView) SetNumberFormatExplicitNull(b bool)`

SetNumberFormatExplicitNull (un)sets NumberFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NumberFormat value is set to nil even if false is passed
### GetRowCount

`func (o *WorkbookRangeView) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *WorkbookRangeView) GetRowCountOk() (int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowCount

`func (o *WorkbookRangeView) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCount

`func (o *WorkbookRangeView) SetRowCount(v int32)`

SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.

### GetText

`func (o *WorkbookRangeView) GetText() AnyOfobject`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WorkbookRangeView) GetTextOk() (AnyOfobject, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasText

`func (o *WorkbookRangeView) HasText() bool`

HasText returns a boolean if a field has been set.

### SetText

`func (o *WorkbookRangeView) SetText(v AnyOfobject)`

SetText gets a reference to the given AnyOfobject and assigns it to the Text field.

### SetTextExplicitNull

`func (o *WorkbookRangeView) SetTextExplicitNull(b bool)`

SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Text value is set to nil even if false is passed
### GetValueTypes

`func (o *WorkbookRangeView) GetValueTypes() AnyOfobject`

GetValueTypes returns the ValueTypes field if non-nil, zero value otherwise.

### GetValueTypesOk

`func (o *WorkbookRangeView) GetValueTypesOk() (AnyOfobject, bool)`

GetValueTypesOk returns a tuple with the ValueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValueTypes

`func (o *WorkbookRangeView) HasValueTypes() bool`

HasValueTypes returns a boolean if a field has been set.

### SetValueTypes

`func (o *WorkbookRangeView) SetValueTypes(v AnyOfobject)`

SetValueTypes gets a reference to the given AnyOfobject and assigns it to the ValueTypes field.

### SetValueTypesExplicitNull

`func (o *WorkbookRangeView) SetValueTypesExplicitNull(b bool)`

SetValueTypesExplicitNull (un)sets ValueTypes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ValueTypes value is set to nil even if false is passed
### GetValues

`func (o *WorkbookRangeView) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WorkbookRangeView) GetValuesOk() (AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *WorkbookRangeView) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *WorkbookRangeView) SetValues(v AnyOfobject)`

SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.

### SetValuesExplicitNull

`func (o *WorkbookRangeView) SetValuesExplicitNull(b bool)`

SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Values value is set to nil even if false is passed
### GetRows

`func (o *WorkbookRangeView) GetRows() []MicrosoftGraphWorkbookRangeView`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WorkbookRangeView) GetRowsOk() ([]MicrosoftGraphWorkbookRangeView, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRows

`func (o *WorkbookRangeView) HasRows() bool`

HasRows returns a boolean if a field has been set.

### SetRows

`func (o *WorkbookRangeView) SetRows(v []MicrosoftGraphWorkbookRangeView)`

SetRows gets a reference to the given []MicrosoftGraphWorkbookRangeView and assigns it to the Rows field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



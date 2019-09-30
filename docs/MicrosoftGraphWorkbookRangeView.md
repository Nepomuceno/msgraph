# MicrosoftGraphWorkbookRangeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *MicrosoftGraphWorkbookRangeView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookRangeView) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MicrosoftGraphWorkbookRangeView) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MicrosoftGraphWorkbookRangeView) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetCellAddresses

`func (o *MicrosoftGraphWorkbookRangeView) GetCellAddresses() AnyOfobject`

GetCellAddresses returns the CellAddresses field if non-nil, zero value otherwise.

### GetCellAddressesOk

`func (o *MicrosoftGraphWorkbookRangeView) GetCellAddressesOk() (AnyOfobject, bool)`

GetCellAddressesOk returns a tuple with the CellAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCellAddresses

`func (o *MicrosoftGraphWorkbookRangeView) HasCellAddresses() bool`

HasCellAddresses returns a boolean if a field has been set.

### SetCellAddresses

`func (o *MicrosoftGraphWorkbookRangeView) SetCellAddresses(v AnyOfobject)`

SetCellAddresses gets a reference to the given AnyOfobject and assigns it to the CellAddresses field.

### SetCellAddressesExplicitNull

`func (o *MicrosoftGraphWorkbookRangeView) SetCellAddressesExplicitNull(b bool)`

SetCellAddressesExplicitNull (un)sets CellAddresses to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The CellAddresses value is set to nil even if false is passed
### GetColumnCount

`func (o *MicrosoftGraphWorkbookRangeView) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *MicrosoftGraphWorkbookRangeView) GetColumnCountOk() (int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasColumnCount

`func (o *MicrosoftGraphWorkbookRangeView) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### SetColumnCount

`func (o *MicrosoftGraphWorkbookRangeView) SetColumnCount(v int32)`

SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.

### GetFormulas

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulas() AnyOfobject`

GetFormulas returns the Formulas field if non-nil, zero value otherwise.

### GetFormulasOk

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasOk() (AnyOfobject, bool)`

GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulas

`func (o *MicrosoftGraphWorkbookRangeView) HasFormulas() bool`

HasFormulas returns a boolean if a field has been set.

### SetFormulas

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulas(v AnyOfobject)`

SetFormulas gets a reference to the given AnyOfobject and assigns it to the Formulas field.

### SetFormulasExplicitNull

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasExplicitNull(b bool)`

SetFormulasExplicitNull (un)sets Formulas to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Formulas value is set to nil even if false is passed
### GetFormulasLocal

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasLocal() AnyOfobject`

GetFormulasLocal returns the FormulasLocal field if non-nil, zero value otherwise.

### GetFormulasLocalOk

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasLocalOk() (AnyOfobject, bool)`

GetFormulasLocalOk returns a tuple with the FormulasLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulasLocal

`func (o *MicrosoftGraphWorkbookRangeView) HasFormulasLocal() bool`

HasFormulasLocal returns a boolean if a field has been set.

### SetFormulasLocal

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasLocal(v AnyOfobject)`

SetFormulasLocal gets a reference to the given AnyOfobject and assigns it to the FormulasLocal field.

### SetFormulasLocalExplicitNull

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasLocalExplicitNull(b bool)`

SetFormulasLocalExplicitNull (un)sets FormulasLocal to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FormulasLocal value is set to nil even if false is passed
### GetFormulasR1C1

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasR1C1() AnyOfobject`

GetFormulasR1C1 returns the FormulasR1C1 field if non-nil, zero value otherwise.

### GetFormulasR1C1Ok

`func (o *MicrosoftGraphWorkbookRangeView) GetFormulasR1C1Ok() (AnyOfobject, bool)`

GetFormulasR1C1Ok returns a tuple with the FormulasR1C1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFormulasR1C1

`func (o *MicrosoftGraphWorkbookRangeView) HasFormulasR1C1() bool`

HasFormulasR1C1 returns a boolean if a field has been set.

### SetFormulasR1C1

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasR1C1(v AnyOfobject)`

SetFormulasR1C1 gets a reference to the given AnyOfobject and assigns it to the FormulasR1C1 field.

### SetFormulasR1C1ExplicitNull

`func (o *MicrosoftGraphWorkbookRangeView) SetFormulasR1C1ExplicitNull(b bool)`

SetFormulasR1C1ExplicitNull (un)sets FormulasR1C1 to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The FormulasR1C1 value is set to nil even if false is passed
### GetIndex

`func (o *MicrosoftGraphWorkbookRangeView) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MicrosoftGraphWorkbookRangeView) GetIndexOk() (int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIndex

`func (o *MicrosoftGraphWorkbookRangeView) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndex

`func (o *MicrosoftGraphWorkbookRangeView) SetIndex(v int32)`

SetIndex gets a reference to the given int32 and assigns it to the Index field.

### GetNumberFormat

`func (o *MicrosoftGraphWorkbookRangeView) GetNumberFormat() AnyOfobject`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *MicrosoftGraphWorkbookRangeView) GetNumberFormatOk() (AnyOfobject, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberFormat

`func (o *MicrosoftGraphWorkbookRangeView) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### SetNumberFormat

`func (o *MicrosoftGraphWorkbookRangeView) SetNumberFormat(v AnyOfobject)`

SetNumberFormat gets a reference to the given AnyOfobject and assigns it to the NumberFormat field.

### SetNumberFormatExplicitNull

`func (o *MicrosoftGraphWorkbookRangeView) SetNumberFormatExplicitNull(b bool)`

SetNumberFormatExplicitNull (un)sets NumberFormat to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The NumberFormat value is set to nil even if false is passed
### GetRowCount

`func (o *MicrosoftGraphWorkbookRangeView) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *MicrosoftGraphWorkbookRangeView) GetRowCountOk() (int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRowCount

`func (o *MicrosoftGraphWorkbookRangeView) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### SetRowCount

`func (o *MicrosoftGraphWorkbookRangeView) SetRowCount(v int32)`

SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.

### GetText

`func (o *MicrosoftGraphWorkbookRangeView) GetText() AnyOfobject`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MicrosoftGraphWorkbookRangeView) GetTextOk() (AnyOfobject, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasText

`func (o *MicrosoftGraphWorkbookRangeView) HasText() bool`

HasText returns a boolean if a field has been set.

### SetText

`func (o *MicrosoftGraphWorkbookRangeView) SetText(v AnyOfobject)`

SetText gets a reference to the given AnyOfobject and assigns it to the Text field.

### SetTextExplicitNull

`func (o *MicrosoftGraphWorkbookRangeView) SetTextExplicitNull(b bool)`

SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Text value is set to nil even if false is passed
### GetValueTypes

`func (o *MicrosoftGraphWorkbookRangeView) GetValueTypes() AnyOfobject`

GetValueTypes returns the ValueTypes field if non-nil, zero value otherwise.

### GetValueTypesOk

`func (o *MicrosoftGraphWorkbookRangeView) GetValueTypesOk() (AnyOfobject, bool)`

GetValueTypesOk returns a tuple with the ValueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValueTypes

`func (o *MicrosoftGraphWorkbookRangeView) HasValueTypes() bool`

HasValueTypes returns a boolean if a field has been set.

### SetValueTypes

`func (o *MicrosoftGraphWorkbookRangeView) SetValueTypes(v AnyOfobject)`

SetValueTypes gets a reference to the given AnyOfobject and assigns it to the ValueTypes field.

### SetValueTypesExplicitNull

`func (o *MicrosoftGraphWorkbookRangeView) SetValueTypesExplicitNull(b bool)`

SetValueTypesExplicitNull (un)sets ValueTypes to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The ValueTypes value is set to nil even if false is passed
### GetValues

`func (o *MicrosoftGraphWorkbookRangeView) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MicrosoftGraphWorkbookRangeView) GetValuesOk() (AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *MicrosoftGraphWorkbookRangeView) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *MicrosoftGraphWorkbookRangeView) SetValues(v AnyOfobject)`

SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.

### SetValuesExplicitNull

`func (o *MicrosoftGraphWorkbookRangeView) SetValuesExplicitNull(b bool)`

SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
when serializing to JSON (pass true as argument to set this, false to unset)
The Values value is set to nil even if false is passed
### GetRows

`func (o *MicrosoftGraphWorkbookRangeView) GetRows() []MicrosoftGraphWorkbookRangeView`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *MicrosoftGraphWorkbookRangeView) GetRowsOk() ([]MicrosoftGraphWorkbookRangeView, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRows

`func (o *MicrosoftGraphWorkbookRangeView) HasRows() bool`

HasRows returns a boolean if a field has been set.

### SetRows

`func (o *MicrosoftGraphWorkbookRangeView) SetRows(v []MicrosoftGraphWorkbookRangeView)`

SetRows gets a reference to the given []MicrosoftGraphWorkbookRangeView and assigns it to the Rows field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



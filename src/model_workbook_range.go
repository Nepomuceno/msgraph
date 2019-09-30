/*
 * OData Service for namespace microsoft.graph
 *
 * This OData service is located at https://graph.microsoft.com/v1.0/
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msgraph
import (
	"encoding/json"
)
// WorkbookRange struct for WorkbookRange
type WorkbookRange struct {
	Address *string `json:"address,omitempty"`
	isExplicitNullAddress bool `json:"-"`
	AddressLocal *string `json:"addressLocal,omitempty"`
	isExplicitNullAddressLocal bool `json:"-"`
	CellCount *int32 `json:"cellCount,omitempty"`

	ColumnCount *int32 `json:"columnCount,omitempty"`

	ColumnHidden *bool `json:"columnHidden,omitempty"`
	isExplicitNullColumnHidden bool `json:"-"`
	ColumnIndex *int32 `json:"columnIndex,omitempty"`

	Formulas *AnyOfobject `json:"formulas,omitempty"`
	isExplicitNullFormulas bool `json:"-"`
	FormulasLocal *AnyOfobject `json:"formulasLocal,omitempty"`
	isExplicitNullFormulasLocal bool `json:"-"`
	FormulasR1C1 *AnyOfobject `json:"formulasR1C1,omitempty"`
	isExplicitNullFormulasR1C1 bool `json:"-"`
	Hidden *bool `json:"hidden,omitempty"`
	isExplicitNullHidden bool `json:"-"`
	NumberFormat *AnyOfobject `json:"numberFormat,omitempty"`
	isExplicitNullNumberFormat bool `json:"-"`
	RowCount *int32 `json:"rowCount,omitempty"`

	RowHidden *bool `json:"rowHidden,omitempty"`
	isExplicitNullRowHidden bool `json:"-"`
	RowIndex *int32 `json:"rowIndex,omitempty"`

	Text *AnyOfobject `json:"text,omitempty"`
	isExplicitNullText bool `json:"-"`
	ValueTypes *AnyOfobject `json:"valueTypes,omitempty"`
	isExplicitNullValueTypes bool `json:"-"`
	Values *AnyOfobject `json:"values,omitempty"`
	isExplicitNullValues bool `json:"-"`
	Format *AnyOfmicrosoftGraphWorkbookRangeFormat `json:"format,omitempty"`
	isExplicitNullFormat bool `json:"-"`
	Sort *AnyOfmicrosoftGraphWorkbookRangeSort `json:"sort,omitempty"`
	isExplicitNullSort bool `json:"-"`
	Worksheet *AnyOfmicrosoftGraphWorkbookWorksheet `json:"worksheet,omitempty"`
	isExplicitNullWorksheet bool `json:"-"`
}

// GetAddress returns the Address field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetAddressOk() (string, bool) {
	if o == nil || o.Address == nil {
		var ret string
		return ret, false
	}
	return *o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *WorkbookRange) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *WorkbookRange) SetAddress(v string) {
	o.Address = &v
}

// SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Address value is set to nil even if false is passed
func (o *WorkbookRange) SetAddressExplicitNull(b bool) {
	o.Address = nil
	o.isExplicitNullAddress = b
}
// GetAddressLocal returns the AddressLocal field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetAddressLocal() string {
	if o == nil || o.AddressLocal == nil {
		var ret string
		return ret
	}
	return *o.AddressLocal
}

// GetAddressLocalOk returns a tuple with the AddressLocal field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetAddressLocalOk() (string, bool) {
	if o == nil || o.AddressLocal == nil {
		var ret string
		return ret, false
	}
	return *o.AddressLocal, true
}

// HasAddressLocal returns a boolean if a field has been set.
func (o *WorkbookRange) HasAddressLocal() bool {
	if o != nil && o.AddressLocal != nil {
		return true
	}

	return false
}

// SetAddressLocal gets a reference to the given string and assigns it to the AddressLocal field.
func (o *WorkbookRange) SetAddressLocal(v string) {
	o.AddressLocal = &v
}

// SetAddressLocalExplicitNull (un)sets AddressLocal to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AddressLocal value is set to nil even if false is passed
func (o *WorkbookRange) SetAddressLocalExplicitNull(b bool) {
	o.AddressLocal = nil
	o.isExplicitNullAddressLocal = b
}
// GetCellCount returns the CellCount field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetCellCount() int32 {
	if o == nil || o.CellCount == nil {
		var ret int32
		return ret
	}
	return *o.CellCount
}

// GetCellCountOk returns a tuple with the CellCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetCellCountOk() (int32, bool) {
	if o == nil || o.CellCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CellCount, true
}

// HasCellCount returns a boolean if a field has been set.
func (o *WorkbookRange) HasCellCount() bool {
	if o != nil && o.CellCount != nil {
		return true
	}

	return false
}

// SetCellCount gets a reference to the given int32 and assigns it to the CellCount field.
func (o *WorkbookRange) SetCellCount(v int32) {
	o.CellCount = &v
}

// GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetColumnCount() int32 {
	if o == nil || o.ColumnCount == nil {
		var ret int32
		return ret
	}
	return *o.ColumnCount
}

// GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetColumnCountOk() (int32, bool) {
	if o == nil || o.ColumnCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ColumnCount, true
}

// HasColumnCount returns a boolean if a field has been set.
func (o *WorkbookRange) HasColumnCount() bool {
	if o != nil && o.ColumnCount != nil {
		return true
	}

	return false
}

// SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.
func (o *WorkbookRange) SetColumnCount(v int32) {
	o.ColumnCount = &v
}

// GetColumnHidden returns the ColumnHidden field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetColumnHidden() bool {
	if o == nil || o.ColumnHidden == nil {
		var ret bool
		return ret
	}
	return *o.ColumnHidden
}

// GetColumnHiddenOk returns a tuple with the ColumnHidden field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetColumnHiddenOk() (bool, bool) {
	if o == nil || o.ColumnHidden == nil {
		var ret bool
		return ret, false
	}
	return *o.ColumnHidden, true
}

// HasColumnHidden returns a boolean if a field has been set.
func (o *WorkbookRange) HasColumnHidden() bool {
	if o != nil && o.ColumnHidden != nil {
		return true
	}

	return false
}

// SetColumnHidden gets a reference to the given bool and assigns it to the ColumnHidden field.
func (o *WorkbookRange) SetColumnHidden(v bool) {
	o.ColumnHidden = &v
}

// SetColumnHiddenExplicitNull (un)sets ColumnHidden to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ColumnHidden value is set to nil even if false is passed
func (o *WorkbookRange) SetColumnHiddenExplicitNull(b bool) {
	o.ColumnHidden = nil
	o.isExplicitNullColumnHidden = b
}
// GetColumnIndex returns the ColumnIndex field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetColumnIndex() int32 {
	if o == nil || o.ColumnIndex == nil {
		var ret int32
		return ret
	}
	return *o.ColumnIndex
}

// GetColumnIndexOk returns a tuple with the ColumnIndex field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetColumnIndexOk() (int32, bool) {
	if o == nil || o.ColumnIndex == nil {
		var ret int32
		return ret, false
	}
	return *o.ColumnIndex, true
}

// HasColumnIndex returns a boolean if a field has been set.
func (o *WorkbookRange) HasColumnIndex() bool {
	if o != nil && o.ColumnIndex != nil {
		return true
	}

	return false
}

// SetColumnIndex gets a reference to the given int32 and assigns it to the ColumnIndex field.
func (o *WorkbookRange) SetColumnIndex(v int32) {
	o.ColumnIndex = &v
}

// GetFormulas returns the Formulas field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetFormulas() AnyOfobject {
	if o == nil || o.Formulas == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetFormulasOk() (AnyOfobject, bool) {
	if o == nil || o.Formulas == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *WorkbookRange) HasFormulas() bool {
	if o != nil && o.Formulas != nil {
		return true
	}

	return false
}

// SetFormulas gets a reference to the given AnyOfobject and assigns it to the Formulas field.
func (o *WorkbookRange) SetFormulas(v AnyOfobject) {
	o.Formulas = &v
}

// SetFormulasExplicitNull (un)sets Formulas to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Formulas value is set to nil even if false is passed
func (o *WorkbookRange) SetFormulasExplicitNull(b bool) {
	o.Formulas = nil
	o.isExplicitNullFormulas = b
}
// GetFormulasLocal returns the FormulasLocal field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetFormulasLocal() AnyOfobject {
	if o == nil || o.FormulasLocal == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.FormulasLocal
}

// GetFormulasLocalOk returns a tuple with the FormulasLocal field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetFormulasLocalOk() (AnyOfobject, bool) {
	if o == nil || o.FormulasLocal == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.FormulasLocal, true
}

// HasFormulasLocal returns a boolean if a field has been set.
func (o *WorkbookRange) HasFormulasLocal() bool {
	if o != nil && o.FormulasLocal != nil {
		return true
	}

	return false
}

// SetFormulasLocal gets a reference to the given AnyOfobject and assigns it to the FormulasLocal field.
func (o *WorkbookRange) SetFormulasLocal(v AnyOfobject) {
	o.FormulasLocal = &v
}

// SetFormulasLocalExplicitNull (un)sets FormulasLocal to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The FormulasLocal value is set to nil even if false is passed
func (o *WorkbookRange) SetFormulasLocalExplicitNull(b bool) {
	o.FormulasLocal = nil
	o.isExplicitNullFormulasLocal = b
}
// GetFormulasR1C1 returns the FormulasR1C1 field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetFormulasR1C1() AnyOfobject {
	if o == nil || o.FormulasR1C1 == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.FormulasR1C1
}

// GetFormulasR1C1Ok returns a tuple with the FormulasR1C1 field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetFormulasR1C1Ok() (AnyOfobject, bool) {
	if o == nil || o.FormulasR1C1 == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.FormulasR1C1, true
}

// HasFormulasR1C1 returns a boolean if a field has been set.
func (o *WorkbookRange) HasFormulasR1C1() bool {
	if o != nil && o.FormulasR1C1 != nil {
		return true
	}

	return false
}

// SetFormulasR1C1 gets a reference to the given AnyOfobject and assigns it to the FormulasR1C1 field.
func (o *WorkbookRange) SetFormulasR1C1(v AnyOfobject) {
	o.FormulasR1C1 = &v
}

// SetFormulasR1C1ExplicitNull (un)sets FormulasR1C1 to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The FormulasR1C1 value is set to nil even if false is passed
func (o *WorkbookRange) SetFormulasR1C1ExplicitNull(b bool) {
	o.FormulasR1C1 = nil
	o.isExplicitNullFormulasR1C1 = b
}
// GetHidden returns the Hidden field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetHiddenOk() (bool, bool) {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret, false
	}
	return *o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *WorkbookRange) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *WorkbookRange) SetHidden(v bool) {
	o.Hidden = &v
}

// SetHiddenExplicitNull (un)sets Hidden to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Hidden value is set to nil even if false is passed
func (o *WorkbookRange) SetHiddenExplicitNull(b bool) {
	o.Hidden = nil
	o.isExplicitNullHidden = b
}
// GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetNumberFormat() AnyOfobject {
	if o == nil || o.NumberFormat == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetNumberFormatOk() (AnyOfobject, bool) {
	if o == nil || o.NumberFormat == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *WorkbookRange) HasNumberFormat() bool {
	if o != nil && o.NumberFormat != nil {
		return true
	}

	return false
}

// SetNumberFormat gets a reference to the given AnyOfobject and assigns it to the NumberFormat field.
func (o *WorkbookRange) SetNumberFormat(v AnyOfobject) {
	o.NumberFormat = &v
}

// SetNumberFormatExplicitNull (un)sets NumberFormat to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The NumberFormat value is set to nil even if false is passed
func (o *WorkbookRange) SetNumberFormatExplicitNull(b bool) {
	o.NumberFormat = nil
	o.isExplicitNullNumberFormat = b
}
// GetRowCount returns the RowCount field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetRowCount() int32 {
	if o == nil || o.RowCount == nil {
		var ret int32
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetRowCountOk() (int32, bool) {
	if o == nil || o.RowCount == nil {
		var ret int32
		return ret, false
	}
	return *o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *WorkbookRange) HasRowCount() bool {
	if o != nil && o.RowCount != nil {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *WorkbookRange) SetRowCount(v int32) {
	o.RowCount = &v
}

// GetRowHidden returns the RowHidden field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetRowHidden() bool {
	if o == nil || o.RowHidden == nil {
		var ret bool
		return ret
	}
	return *o.RowHidden
}

// GetRowHiddenOk returns a tuple with the RowHidden field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetRowHiddenOk() (bool, bool) {
	if o == nil || o.RowHidden == nil {
		var ret bool
		return ret, false
	}
	return *o.RowHidden, true
}

// HasRowHidden returns a boolean if a field has been set.
func (o *WorkbookRange) HasRowHidden() bool {
	if o != nil && o.RowHidden != nil {
		return true
	}

	return false
}

// SetRowHidden gets a reference to the given bool and assigns it to the RowHidden field.
func (o *WorkbookRange) SetRowHidden(v bool) {
	o.RowHidden = &v
}

// SetRowHiddenExplicitNull (un)sets RowHidden to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The RowHidden value is set to nil even if false is passed
func (o *WorkbookRange) SetRowHiddenExplicitNull(b bool) {
	o.RowHidden = nil
	o.isExplicitNullRowHidden = b
}
// GetRowIndex returns the RowIndex field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetRowIndex() int32 {
	if o == nil || o.RowIndex == nil {
		var ret int32
		return ret
	}
	return *o.RowIndex
}

// GetRowIndexOk returns a tuple with the RowIndex field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetRowIndexOk() (int32, bool) {
	if o == nil || o.RowIndex == nil {
		var ret int32
		return ret, false
	}
	return *o.RowIndex, true
}

// HasRowIndex returns a boolean if a field has been set.
func (o *WorkbookRange) HasRowIndex() bool {
	if o != nil && o.RowIndex != nil {
		return true
	}

	return false
}

// SetRowIndex gets a reference to the given int32 and assigns it to the RowIndex field.
func (o *WorkbookRange) SetRowIndex(v int32) {
	o.RowIndex = &v
}

// GetText returns the Text field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetText() AnyOfobject {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetTextOk() (AnyOfobject, bool) {
	if o == nil || o.Text == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WorkbookRange) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given AnyOfobject and assigns it to the Text field.
func (o *WorkbookRange) SetText(v AnyOfobject) {
	o.Text = &v
}

// SetTextExplicitNull (un)sets Text to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Text value is set to nil even if false is passed
func (o *WorkbookRange) SetTextExplicitNull(b bool) {
	o.Text = nil
	o.isExplicitNullText = b
}
// GetValueTypes returns the ValueTypes field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetValueTypes() AnyOfobject {
	if o == nil || o.ValueTypes == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.ValueTypes
}

// GetValueTypesOk returns a tuple with the ValueTypes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetValueTypesOk() (AnyOfobject, bool) {
	if o == nil || o.ValueTypes == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.ValueTypes, true
}

// HasValueTypes returns a boolean if a field has been set.
func (o *WorkbookRange) HasValueTypes() bool {
	if o != nil && o.ValueTypes != nil {
		return true
	}

	return false
}

// SetValueTypes gets a reference to the given AnyOfobject and assigns it to the ValueTypes field.
func (o *WorkbookRange) SetValueTypes(v AnyOfobject) {
	o.ValueTypes = &v
}

// SetValueTypesExplicitNull (un)sets ValueTypes to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ValueTypes value is set to nil even if false is passed
func (o *WorkbookRange) SetValueTypesExplicitNull(b bool) {
	o.ValueTypes = nil
	o.isExplicitNullValueTypes = b
}
// GetValues returns the Values field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetValues() AnyOfobject {
	if o == nil || o.Values == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetValuesOk() (AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *WorkbookRange) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *WorkbookRange) SetValues(v AnyOfobject) {
	o.Values = &v
}

// SetValuesExplicitNull (un)sets Values to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Values value is set to nil even if false is passed
func (o *WorkbookRange) SetValuesExplicitNull(b bool) {
	o.Values = nil
	o.isExplicitNullValues = b
}
// GetFormat returns the Format field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetFormat() AnyOfmicrosoftGraphWorkbookRangeFormat {
	if o == nil || o.Format == nil {
		var ret AnyOfmicrosoftGraphWorkbookRangeFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetFormatOk() (AnyOfmicrosoftGraphWorkbookRangeFormat, bool) {
	if o == nil || o.Format == nil {
		var ret AnyOfmicrosoftGraphWorkbookRangeFormat
		return ret, false
	}
	return *o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WorkbookRange) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFormat and assigns it to the Format field.
func (o *WorkbookRange) SetFormat(v AnyOfmicrosoftGraphWorkbookRangeFormat) {
	o.Format = &v
}

// SetFormatExplicitNull (un)sets Format to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Format value is set to nil even if false is passed
func (o *WorkbookRange) SetFormatExplicitNull(b bool) {
	o.Format = nil
	o.isExplicitNullFormat = b
}
// GetSort returns the Sort field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetSort() AnyOfmicrosoftGraphWorkbookRangeSort {
	if o == nil || o.Sort == nil {
		var ret AnyOfmicrosoftGraphWorkbookRangeSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetSortOk() (AnyOfmicrosoftGraphWorkbookRangeSort, bool) {
	if o == nil || o.Sort == nil {
		var ret AnyOfmicrosoftGraphWorkbookRangeSort
		return ret, false
	}
	return *o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *WorkbookRange) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeSort and assigns it to the Sort field.
func (o *WorkbookRange) SetSort(v AnyOfmicrosoftGraphWorkbookRangeSort) {
	o.Sort = &v
}

// SetSortExplicitNull (un)sets Sort to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Sort value is set to nil even if false is passed
func (o *WorkbookRange) SetSortExplicitNull(b bool) {
	o.Sort = nil
	o.isExplicitNullSort = b
}
// GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.
func (o *WorkbookRange) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return *o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRange) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret, false
	}
	return *o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *WorkbookRange) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *WorkbookRange) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = &v
}

// SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Worksheet value is set to nil even if false is passed
func (o *WorkbookRange) SetWorksheetExplicitNull(b bool) {
	o.Worksheet = nil
	o.isExplicitNullWorksheet = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address == nil {
		if o.isExplicitNullAddress {
			toSerialize["address"] = o.Address
		}
	} else {
		toSerialize["address"] = o.Address
	}
	if o.AddressLocal == nil {
		if o.isExplicitNullAddressLocal {
			toSerialize["addressLocal"] = o.AddressLocal
		}
	} else {
		toSerialize["addressLocal"] = o.AddressLocal
	}
	if o.CellCount != nil {
		toSerialize["cellCount"] = o.CellCount
	}
	if o.ColumnCount != nil {
		toSerialize["columnCount"] = o.ColumnCount
	}
	if o.ColumnHidden == nil {
		if o.isExplicitNullColumnHidden {
			toSerialize["columnHidden"] = o.ColumnHidden
		}
	} else {
		toSerialize["columnHidden"] = o.ColumnHidden
	}
	if o.ColumnIndex != nil {
		toSerialize["columnIndex"] = o.ColumnIndex
	}
	if o.Formulas == nil {
		if o.isExplicitNullFormulas {
			toSerialize["formulas"] = o.Formulas
		}
	} else {
		toSerialize["formulas"] = o.Formulas
	}
	if o.FormulasLocal == nil {
		if o.isExplicitNullFormulasLocal {
			toSerialize["formulasLocal"] = o.FormulasLocal
		}
	} else {
		toSerialize["formulasLocal"] = o.FormulasLocal
	}
	if o.FormulasR1C1 == nil {
		if o.isExplicitNullFormulasR1C1 {
			toSerialize["formulasR1C1"] = o.FormulasR1C1
		}
	} else {
		toSerialize["formulasR1C1"] = o.FormulasR1C1
	}
	if o.Hidden == nil {
		if o.isExplicitNullHidden {
			toSerialize["hidden"] = o.Hidden
		}
	} else {
		toSerialize["hidden"] = o.Hidden
	}
	if o.NumberFormat == nil {
		if o.isExplicitNullNumberFormat {
			toSerialize["numberFormat"] = o.NumberFormat
		}
	} else {
		toSerialize["numberFormat"] = o.NumberFormat
	}
	if o.RowCount != nil {
		toSerialize["rowCount"] = o.RowCount
	}
	if o.RowHidden == nil {
		if o.isExplicitNullRowHidden {
			toSerialize["rowHidden"] = o.RowHidden
		}
	} else {
		toSerialize["rowHidden"] = o.RowHidden
	}
	if o.RowIndex != nil {
		toSerialize["rowIndex"] = o.RowIndex
	}
	if o.Text == nil {
		if o.isExplicitNullText {
			toSerialize["text"] = o.Text
		}
	} else {
		toSerialize["text"] = o.Text
	}
	if o.ValueTypes == nil {
		if o.isExplicitNullValueTypes {
			toSerialize["valueTypes"] = o.ValueTypes
		}
	} else {
		toSerialize["valueTypes"] = o.ValueTypes
	}
	if o.Values == nil {
		if o.isExplicitNullValues {
			toSerialize["values"] = o.Values
		}
	} else {
		toSerialize["values"] = o.Values
	}
	if o.Format == nil {
		if o.isExplicitNullFormat {
			toSerialize["format"] = o.Format
		}
	} else {
		toSerialize["format"] = o.Format
	}
	if o.Sort == nil {
		if o.isExplicitNullSort {
			toSerialize["sort"] = o.Sort
		}
	} else {
		toSerialize["sort"] = o.Sort
	}
	if o.Worksheet == nil {
		if o.isExplicitNullWorksheet {
			toSerialize["worksheet"] = o.Worksheet
		}
	} else {
		toSerialize["worksheet"] = o.Worksheet
	}
	return json.Marshal(toSerialize)
}


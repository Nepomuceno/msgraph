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
// WorkbookTable struct for WorkbookTable
type WorkbookTable struct {
	HighlightFirstColumn *bool `json:"highlightFirstColumn,omitempty"`

	HighlightLastColumn *bool `json:"highlightLastColumn,omitempty"`

	LegacyId *string `json:"legacyId,omitempty"`
	isExplicitNullLegacyId bool `json:"-"`
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	ShowBandedColumns *bool `json:"showBandedColumns,omitempty"`

	ShowBandedRows *bool `json:"showBandedRows,omitempty"`

	ShowFilterButton *bool `json:"showFilterButton,omitempty"`

	ShowHeaders *bool `json:"showHeaders,omitempty"`

	ShowTotals *bool `json:"showTotals,omitempty"`

	Style *string `json:"style,omitempty"`
	isExplicitNullStyle bool `json:"-"`
	Columns *[]MicrosoftGraphWorkbookTableColumn `json:"columns,omitempty"`

	Rows *[]MicrosoftGraphWorkbookTableRow `json:"rows,omitempty"`

	Sort *AnyOfmicrosoftGraphWorkbookTableSort `json:"sort,omitempty"`
	isExplicitNullSort bool `json:"-"`
	Worksheet *AnyOfmicrosoftGraphWorkbookWorksheet `json:"worksheet,omitempty"`
	isExplicitNullWorksheet bool `json:"-"`
}

// GetHighlightFirstColumn returns the HighlightFirstColumn field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetHighlightFirstColumn() bool {
	if o == nil || o.HighlightFirstColumn == nil {
		var ret bool
		return ret
	}
	return *o.HighlightFirstColumn
}

// GetHighlightFirstColumnOk returns a tuple with the HighlightFirstColumn field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetHighlightFirstColumnOk() (bool, bool) {
	if o == nil || o.HighlightFirstColumn == nil {
		var ret bool
		return ret, false
	}
	return *o.HighlightFirstColumn, true
}

// HasHighlightFirstColumn returns a boolean if a field has been set.
func (o *WorkbookTable) HasHighlightFirstColumn() bool {
	if o != nil && o.HighlightFirstColumn != nil {
		return true
	}

	return false
}

// SetHighlightFirstColumn gets a reference to the given bool and assigns it to the HighlightFirstColumn field.
func (o *WorkbookTable) SetHighlightFirstColumn(v bool) {
	o.HighlightFirstColumn = &v
}

// GetHighlightLastColumn returns the HighlightLastColumn field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetHighlightLastColumn() bool {
	if o == nil || o.HighlightLastColumn == nil {
		var ret bool
		return ret
	}
	return *o.HighlightLastColumn
}

// GetHighlightLastColumnOk returns a tuple with the HighlightLastColumn field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetHighlightLastColumnOk() (bool, bool) {
	if o == nil || o.HighlightLastColumn == nil {
		var ret bool
		return ret, false
	}
	return *o.HighlightLastColumn, true
}

// HasHighlightLastColumn returns a boolean if a field has been set.
func (o *WorkbookTable) HasHighlightLastColumn() bool {
	if o != nil && o.HighlightLastColumn != nil {
		return true
	}

	return false
}

// SetHighlightLastColumn gets a reference to the given bool and assigns it to the HighlightLastColumn field.
func (o *WorkbookTable) SetHighlightLastColumn(v bool) {
	o.HighlightLastColumn = &v
}

// GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetLegacyId() string {
	if o == nil || o.LegacyId == nil {
		var ret string
		return ret
	}
	return *o.LegacyId
}

// GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetLegacyIdOk() (string, bool) {
	if o == nil || o.LegacyId == nil {
		var ret string
		return ret, false
	}
	return *o.LegacyId, true
}

// HasLegacyId returns a boolean if a field has been set.
func (o *WorkbookTable) HasLegacyId() bool {
	if o != nil && o.LegacyId != nil {
		return true
	}

	return false
}

// SetLegacyId gets a reference to the given string and assigns it to the LegacyId field.
func (o *WorkbookTable) SetLegacyId(v string) {
	o.LegacyId = &v
}

// SetLegacyIdExplicitNull (un)sets LegacyId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LegacyId value is set to nil even if false is passed
func (o *WorkbookTable) SetLegacyIdExplicitNull(b bool) {
	o.LegacyId = nil
	o.isExplicitNullLegacyId = b
}
// GetName returns the Name field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookTable) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkbookTable) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *WorkbookTable) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetShowBandedColumns returns the ShowBandedColumns field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetShowBandedColumns() bool {
	if o == nil || o.ShowBandedColumns == nil {
		var ret bool
		return ret
	}
	return *o.ShowBandedColumns
}

// GetShowBandedColumnsOk returns a tuple with the ShowBandedColumns field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowBandedColumnsOk() (bool, bool) {
	if o == nil || o.ShowBandedColumns == nil {
		var ret bool
		return ret, false
	}
	return *o.ShowBandedColumns, true
}

// HasShowBandedColumns returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowBandedColumns() bool {
	if o != nil && o.ShowBandedColumns != nil {
		return true
	}

	return false
}

// SetShowBandedColumns gets a reference to the given bool and assigns it to the ShowBandedColumns field.
func (o *WorkbookTable) SetShowBandedColumns(v bool) {
	o.ShowBandedColumns = &v
}

// GetShowBandedRows returns the ShowBandedRows field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetShowBandedRows() bool {
	if o == nil || o.ShowBandedRows == nil {
		var ret bool
		return ret
	}
	return *o.ShowBandedRows
}

// GetShowBandedRowsOk returns a tuple with the ShowBandedRows field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowBandedRowsOk() (bool, bool) {
	if o == nil || o.ShowBandedRows == nil {
		var ret bool
		return ret, false
	}
	return *o.ShowBandedRows, true
}

// HasShowBandedRows returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowBandedRows() bool {
	if o != nil && o.ShowBandedRows != nil {
		return true
	}

	return false
}

// SetShowBandedRows gets a reference to the given bool and assigns it to the ShowBandedRows field.
func (o *WorkbookTable) SetShowBandedRows(v bool) {
	o.ShowBandedRows = &v
}

// GetShowFilterButton returns the ShowFilterButton field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetShowFilterButton() bool {
	if o == nil || o.ShowFilterButton == nil {
		var ret bool
		return ret
	}
	return *o.ShowFilterButton
}

// GetShowFilterButtonOk returns a tuple with the ShowFilterButton field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowFilterButtonOk() (bool, bool) {
	if o == nil || o.ShowFilterButton == nil {
		var ret bool
		return ret, false
	}
	return *o.ShowFilterButton, true
}

// HasShowFilterButton returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowFilterButton() bool {
	if o != nil && o.ShowFilterButton != nil {
		return true
	}

	return false
}

// SetShowFilterButton gets a reference to the given bool and assigns it to the ShowFilterButton field.
func (o *WorkbookTable) SetShowFilterButton(v bool) {
	o.ShowFilterButton = &v
}

// GetShowHeaders returns the ShowHeaders field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetShowHeaders() bool {
	if o == nil || o.ShowHeaders == nil {
		var ret bool
		return ret
	}
	return *o.ShowHeaders
}

// GetShowHeadersOk returns a tuple with the ShowHeaders field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowHeadersOk() (bool, bool) {
	if o == nil || o.ShowHeaders == nil {
		var ret bool
		return ret, false
	}
	return *o.ShowHeaders, true
}

// HasShowHeaders returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowHeaders() bool {
	if o != nil && o.ShowHeaders != nil {
		return true
	}

	return false
}

// SetShowHeaders gets a reference to the given bool and assigns it to the ShowHeaders field.
func (o *WorkbookTable) SetShowHeaders(v bool) {
	o.ShowHeaders = &v
}

// GetShowTotals returns the ShowTotals field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetShowTotals() bool {
	if o == nil || o.ShowTotals == nil {
		var ret bool
		return ret
	}
	return *o.ShowTotals
}

// GetShowTotalsOk returns a tuple with the ShowTotals field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetShowTotalsOk() (bool, bool) {
	if o == nil || o.ShowTotals == nil {
		var ret bool
		return ret, false
	}
	return *o.ShowTotals, true
}

// HasShowTotals returns a boolean if a field has been set.
func (o *WorkbookTable) HasShowTotals() bool {
	if o != nil && o.ShowTotals != nil {
		return true
	}

	return false
}

// SetShowTotals gets a reference to the given bool and assigns it to the ShowTotals field.
func (o *WorkbookTable) SetShowTotals(v bool) {
	o.ShowTotals = &v
}

// GetStyle returns the Style field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetStyle() string {
	if o == nil || o.Style == nil {
		var ret string
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetStyleOk() (string, bool) {
	if o == nil || o.Style == nil {
		var ret string
		return ret, false
	}
	return *o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *WorkbookTable) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given string and assigns it to the Style field.
func (o *WorkbookTable) SetStyle(v string) {
	o.Style = &v
}

// SetStyleExplicitNull (un)sets Style to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Style value is set to nil even if false is passed
func (o *WorkbookTable) SetStyleExplicitNull(b bool) {
	o.Style = nil
	o.isExplicitNullStyle = b
}
// GetColumns returns the Columns field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetColumns() []MicrosoftGraphWorkbookTableColumn {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphWorkbookTableColumn
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetColumnsOk() ([]MicrosoftGraphWorkbookTableColumn, bool) {
	if o == nil || o.Columns == nil {
		var ret []MicrosoftGraphWorkbookTableColumn
		return ret, false
	}
	return *o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *WorkbookTable) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []MicrosoftGraphWorkbookTableColumn and assigns it to the Columns field.
func (o *WorkbookTable) SetColumns(v []MicrosoftGraphWorkbookTableColumn) {
	o.Columns = &v
}

// GetRows returns the Rows field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetRows() []MicrosoftGraphWorkbookTableRow {
	if o == nil || o.Rows == nil {
		var ret []MicrosoftGraphWorkbookTableRow
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetRowsOk() ([]MicrosoftGraphWorkbookTableRow, bool) {
	if o == nil || o.Rows == nil {
		var ret []MicrosoftGraphWorkbookTableRow
		return ret, false
	}
	return *o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *WorkbookTable) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given []MicrosoftGraphWorkbookTableRow and assigns it to the Rows field.
func (o *WorkbookTable) SetRows(v []MicrosoftGraphWorkbookTableRow) {
	o.Rows = &v
}

// GetSort returns the Sort field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetSort() AnyOfmicrosoftGraphWorkbookTableSort {
	if o == nil || o.Sort == nil {
		var ret AnyOfmicrosoftGraphWorkbookTableSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetSortOk() (AnyOfmicrosoftGraphWorkbookTableSort, bool) {
	if o == nil || o.Sort == nil {
		var ret AnyOfmicrosoftGraphWorkbookTableSort
		return ret, false
	}
	return *o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *WorkbookTable) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given AnyOfmicrosoftGraphWorkbookTableSort and assigns it to the Sort field.
func (o *WorkbookTable) SetSort(v AnyOfmicrosoftGraphWorkbookTableSort) {
	o.Sort = &v
}

// SetSortExplicitNull (un)sets Sort to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Sort value is set to nil even if false is passed
func (o *WorkbookTable) SetSortExplicitNull(b bool) {
	o.Sort = nil
	o.isExplicitNullSort = b
}
// GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.
func (o *WorkbookTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return *o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTable) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret, false
	}
	return *o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *WorkbookTable) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *WorkbookTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = &v
}

// SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Worksheet value is set to nil even if false is passed
func (o *WorkbookTable) SetWorksheetExplicitNull(b bool) {
	o.Worksheet = nil
	o.isExplicitNullWorksheet = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HighlightFirstColumn != nil {
		toSerialize["highlightFirstColumn"] = o.HighlightFirstColumn
	}
	if o.HighlightLastColumn != nil {
		toSerialize["highlightLastColumn"] = o.HighlightLastColumn
	}
	if o.LegacyId == nil {
		if o.isExplicitNullLegacyId {
			toSerialize["legacyId"] = o.LegacyId
		}
	} else {
		toSerialize["legacyId"] = o.LegacyId
	}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.ShowBandedColumns != nil {
		toSerialize["showBandedColumns"] = o.ShowBandedColumns
	}
	if o.ShowBandedRows != nil {
		toSerialize["showBandedRows"] = o.ShowBandedRows
	}
	if o.ShowFilterButton != nil {
		toSerialize["showFilterButton"] = o.ShowFilterButton
	}
	if o.ShowHeaders != nil {
		toSerialize["showHeaders"] = o.ShowHeaders
	}
	if o.ShowTotals != nil {
		toSerialize["showTotals"] = o.ShowTotals
	}
	if o.Style == nil {
		if o.isExplicitNullStyle {
			toSerialize["style"] = o.Style
		}
	} else {
		toSerialize["style"] = o.Style
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
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


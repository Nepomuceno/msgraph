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
// WorkbookRangeFormat struct for WorkbookRangeFormat
type WorkbookRangeFormat struct {
	ColumnWidth *AnyOfnumberstringstring `json:"columnWidth,omitempty"`
	isExplicitNullColumnWidth bool `json:"-"`
	HorizontalAlignment *string `json:"horizontalAlignment,omitempty"`
	isExplicitNullHorizontalAlignment bool `json:"-"`
	RowHeight *AnyOfnumberstringstring `json:"rowHeight,omitempty"`
	isExplicitNullRowHeight bool `json:"-"`
	VerticalAlignment *string `json:"verticalAlignment,omitempty"`
	isExplicitNullVerticalAlignment bool `json:"-"`
	WrapText *bool `json:"wrapText,omitempty"`
	isExplicitNullWrapText bool `json:"-"`
	Borders *[]MicrosoftGraphWorkbookRangeBorder `json:"borders,omitempty"`

	Fill *AnyOfmicrosoftGraphWorkbookRangeFill `json:"fill,omitempty"`
	isExplicitNullFill bool `json:"-"`
	Font *AnyOfmicrosoftGraphWorkbookRangeFont `json:"font,omitempty"`
	isExplicitNullFont bool `json:"-"`
	Protection *AnyOfmicrosoftGraphWorkbookFormatProtection `json:"protection,omitempty"`
	isExplicitNullProtection bool `json:"-"`
}

// GetColumnWidth returns the ColumnWidth field if non-nil, zero value otherwise.
func (o *WorkbookRangeFormat) GetColumnWidth() AnyOfnumberstringstring {
	if o == nil || o.ColumnWidth == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.ColumnWidth
}

// GetColumnWidthOk returns a tuple with the ColumnWidth field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetColumnWidthOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.ColumnWidth == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.ColumnWidth, true
}

// HasColumnWidth returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasColumnWidth() bool {
	if o != nil && o.ColumnWidth != nil {
		return true
	}

	return false
}

// SetColumnWidth gets a reference to the given AnyOfnumberstringstring and assigns it to the ColumnWidth field.
func (o *WorkbookRangeFormat) SetColumnWidth(v AnyOfnumberstringstring) {
	o.ColumnWidth = &v
}

// SetColumnWidthExplicitNull (un)sets ColumnWidth to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ColumnWidth value is set to nil even if false is passed
func (o *WorkbookRangeFormat) SetColumnWidthExplicitNull(b bool) {
	o.ColumnWidth = nil
	o.isExplicitNullColumnWidth = b
}
// GetHorizontalAlignment returns the HorizontalAlignment field if non-nil, zero value otherwise.
func (o *WorkbookRangeFormat) GetHorizontalAlignment() string {
	if o == nil || o.HorizontalAlignment == nil {
		var ret string
		return ret
	}
	return *o.HorizontalAlignment
}

// GetHorizontalAlignmentOk returns a tuple with the HorizontalAlignment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetHorizontalAlignmentOk() (string, bool) {
	if o == nil || o.HorizontalAlignment == nil {
		var ret string
		return ret, false
	}
	return *o.HorizontalAlignment, true
}

// HasHorizontalAlignment returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasHorizontalAlignment() bool {
	if o != nil && o.HorizontalAlignment != nil {
		return true
	}

	return false
}

// SetHorizontalAlignment gets a reference to the given string and assigns it to the HorizontalAlignment field.
func (o *WorkbookRangeFormat) SetHorizontalAlignment(v string) {
	o.HorizontalAlignment = &v
}

// SetHorizontalAlignmentExplicitNull (un)sets HorizontalAlignment to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The HorizontalAlignment value is set to nil even if false is passed
func (o *WorkbookRangeFormat) SetHorizontalAlignmentExplicitNull(b bool) {
	o.HorizontalAlignment = nil
	o.isExplicitNullHorizontalAlignment = b
}
// GetRowHeight returns the RowHeight field if non-nil, zero value otherwise.
func (o *WorkbookRangeFormat) GetRowHeight() AnyOfnumberstringstring {
	if o == nil || o.RowHeight == nil {
		var ret AnyOfnumberstringstring
		return ret
	}
	return *o.RowHeight
}

// GetRowHeightOk returns a tuple with the RowHeight field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetRowHeightOk() (AnyOfnumberstringstring, bool) {
	if o == nil || o.RowHeight == nil {
		var ret AnyOfnumberstringstring
		return ret, false
	}
	return *o.RowHeight, true
}

// HasRowHeight returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasRowHeight() bool {
	if o != nil && o.RowHeight != nil {
		return true
	}

	return false
}

// SetRowHeight gets a reference to the given AnyOfnumberstringstring and assigns it to the RowHeight field.
func (o *WorkbookRangeFormat) SetRowHeight(v AnyOfnumberstringstring) {
	o.RowHeight = &v
}

// SetRowHeightExplicitNull (un)sets RowHeight to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The RowHeight value is set to nil even if false is passed
func (o *WorkbookRangeFormat) SetRowHeightExplicitNull(b bool) {
	o.RowHeight = nil
	o.isExplicitNullRowHeight = b
}
// GetVerticalAlignment returns the VerticalAlignment field if non-nil, zero value otherwise.
func (o *WorkbookRangeFormat) GetVerticalAlignment() string {
	if o == nil || o.VerticalAlignment == nil {
		var ret string
		return ret
	}
	return *o.VerticalAlignment
}

// GetVerticalAlignmentOk returns a tuple with the VerticalAlignment field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetVerticalAlignmentOk() (string, bool) {
	if o == nil || o.VerticalAlignment == nil {
		var ret string
		return ret, false
	}
	return *o.VerticalAlignment, true
}

// HasVerticalAlignment returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasVerticalAlignment() bool {
	if o != nil && o.VerticalAlignment != nil {
		return true
	}

	return false
}

// SetVerticalAlignment gets a reference to the given string and assigns it to the VerticalAlignment field.
func (o *WorkbookRangeFormat) SetVerticalAlignment(v string) {
	o.VerticalAlignment = &v
}

// SetVerticalAlignmentExplicitNull (un)sets VerticalAlignment to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The VerticalAlignment value is set to nil even if false is passed
func (o *WorkbookRangeFormat) SetVerticalAlignmentExplicitNull(b bool) {
	o.VerticalAlignment = nil
	o.isExplicitNullVerticalAlignment = b
}
// GetWrapText returns the WrapText field if non-nil, zero value otherwise.
func (o *WorkbookRangeFormat) GetWrapText() bool {
	if o == nil || o.WrapText == nil {
		var ret bool
		return ret
	}
	return *o.WrapText
}

// GetWrapTextOk returns a tuple with the WrapText field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetWrapTextOk() (bool, bool) {
	if o == nil || o.WrapText == nil {
		var ret bool
		return ret, false
	}
	return *o.WrapText, true
}

// HasWrapText returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasWrapText() bool {
	if o != nil && o.WrapText != nil {
		return true
	}

	return false
}

// SetWrapText gets a reference to the given bool and assigns it to the WrapText field.
func (o *WorkbookRangeFormat) SetWrapText(v bool) {
	o.WrapText = &v
}

// SetWrapTextExplicitNull (un)sets WrapText to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The WrapText value is set to nil even if false is passed
func (o *WorkbookRangeFormat) SetWrapTextExplicitNull(b bool) {
	o.WrapText = nil
	o.isExplicitNullWrapText = b
}
// GetBorders returns the Borders field if non-nil, zero value otherwise.
func (o *WorkbookRangeFormat) GetBorders() []MicrosoftGraphWorkbookRangeBorder {
	if o == nil || o.Borders == nil {
		var ret []MicrosoftGraphWorkbookRangeBorder
		return ret
	}
	return *o.Borders
}

// GetBordersOk returns a tuple with the Borders field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetBordersOk() ([]MicrosoftGraphWorkbookRangeBorder, bool) {
	if o == nil || o.Borders == nil {
		var ret []MicrosoftGraphWorkbookRangeBorder
		return ret, false
	}
	return *o.Borders, true
}

// HasBorders returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasBorders() bool {
	if o != nil && o.Borders != nil {
		return true
	}

	return false
}

// SetBorders gets a reference to the given []MicrosoftGraphWorkbookRangeBorder and assigns it to the Borders field.
func (o *WorkbookRangeFormat) SetBorders(v []MicrosoftGraphWorkbookRangeBorder) {
	o.Borders = &v
}

// GetFill returns the Fill field if non-nil, zero value otherwise.
func (o *WorkbookRangeFormat) GetFill() AnyOfmicrosoftGraphWorkbookRangeFill {
	if o == nil || o.Fill == nil {
		var ret AnyOfmicrosoftGraphWorkbookRangeFill
		return ret
	}
	return *o.Fill
}

// GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetFillOk() (AnyOfmicrosoftGraphWorkbookRangeFill, bool) {
	if o == nil || o.Fill == nil {
		var ret AnyOfmicrosoftGraphWorkbookRangeFill
		return ret, false
	}
	return *o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasFill() bool {
	if o != nil && o.Fill != nil {
		return true
	}

	return false
}

// SetFill gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFill and assigns it to the Fill field.
func (o *WorkbookRangeFormat) SetFill(v AnyOfmicrosoftGraphWorkbookRangeFill) {
	o.Fill = &v
}

// SetFillExplicitNull (un)sets Fill to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Fill value is set to nil even if false is passed
func (o *WorkbookRangeFormat) SetFillExplicitNull(b bool) {
	o.Fill = nil
	o.isExplicitNullFill = b
}
// GetFont returns the Font field if non-nil, zero value otherwise.
func (o *WorkbookRangeFormat) GetFont() AnyOfmicrosoftGraphWorkbookRangeFont {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookRangeFont
		return ret
	}
	return *o.Font
}

// GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetFontOk() (AnyOfmicrosoftGraphWorkbookRangeFont, bool) {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookRangeFont
		return ret, false
	}
	return *o.Font, true
}

// HasFont returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasFont() bool {
	if o != nil && o.Font != nil {
		return true
	}

	return false
}

// SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookRangeFont and assigns it to the Font field.
func (o *WorkbookRangeFormat) SetFont(v AnyOfmicrosoftGraphWorkbookRangeFont) {
	o.Font = &v
}

// SetFontExplicitNull (un)sets Font to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Font value is set to nil even if false is passed
func (o *WorkbookRangeFormat) SetFontExplicitNull(b bool) {
	o.Font = nil
	o.isExplicitNullFont = b
}
// GetProtection returns the Protection field if non-nil, zero value otherwise.
func (o *WorkbookRangeFormat) GetProtection() AnyOfmicrosoftGraphWorkbookFormatProtection {
	if o == nil || o.Protection == nil {
		var ret AnyOfmicrosoftGraphWorkbookFormatProtection
		return ret
	}
	return *o.Protection
}

// GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookRangeFormat) GetProtectionOk() (AnyOfmicrosoftGraphWorkbookFormatProtection, bool) {
	if o == nil || o.Protection == nil {
		var ret AnyOfmicrosoftGraphWorkbookFormatProtection
		return ret, false
	}
	return *o.Protection, true
}

// HasProtection returns a boolean if a field has been set.
func (o *WorkbookRangeFormat) HasProtection() bool {
	if o != nil && o.Protection != nil {
		return true
	}

	return false
}

// SetProtection gets a reference to the given AnyOfmicrosoftGraphWorkbookFormatProtection and assigns it to the Protection field.
func (o *WorkbookRangeFormat) SetProtection(v AnyOfmicrosoftGraphWorkbookFormatProtection) {
	o.Protection = &v
}

// SetProtectionExplicitNull (un)sets Protection to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Protection value is set to nil even if false is passed
func (o *WorkbookRangeFormat) SetProtectionExplicitNull(b bool) {
	o.Protection = nil
	o.isExplicitNullProtection = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookRangeFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColumnWidth == nil {
		if o.isExplicitNullColumnWidth {
			toSerialize["columnWidth"] = o.ColumnWidth
		}
	} else {
		toSerialize["columnWidth"] = o.ColumnWidth
	}
	if o.HorizontalAlignment == nil {
		if o.isExplicitNullHorizontalAlignment {
			toSerialize["horizontalAlignment"] = o.HorizontalAlignment
		}
	} else {
		toSerialize["horizontalAlignment"] = o.HorizontalAlignment
	}
	if o.RowHeight == nil {
		if o.isExplicitNullRowHeight {
			toSerialize["rowHeight"] = o.RowHeight
		}
	} else {
		toSerialize["rowHeight"] = o.RowHeight
	}
	if o.VerticalAlignment == nil {
		if o.isExplicitNullVerticalAlignment {
			toSerialize["verticalAlignment"] = o.VerticalAlignment
		}
	} else {
		toSerialize["verticalAlignment"] = o.VerticalAlignment
	}
	if o.WrapText == nil {
		if o.isExplicitNullWrapText {
			toSerialize["wrapText"] = o.WrapText
		}
	} else {
		toSerialize["wrapText"] = o.WrapText
	}
	if o.Borders != nil {
		toSerialize["borders"] = o.Borders
	}
	if o.Fill == nil {
		if o.isExplicitNullFill {
			toSerialize["fill"] = o.Fill
		}
	} else {
		toSerialize["fill"] = o.Fill
	}
	if o.Font == nil {
		if o.isExplicitNullFont {
			toSerialize["font"] = o.Font
		}
	} else {
		toSerialize["font"] = o.Font
	}
	if o.Protection == nil {
		if o.isExplicitNullProtection {
			toSerialize["protection"] = o.Protection
		}
	} else {
		toSerialize["protection"] = o.Protection
	}
	return json.Marshal(toSerialize)
}


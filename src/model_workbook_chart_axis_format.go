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
// WorkbookChartAxisFormat struct for WorkbookChartAxisFormat
type WorkbookChartAxisFormat struct {
	Font *AnyOfmicrosoftGraphWorkbookChartFont `json:"font,omitempty"`
	isExplicitNullFont bool `json:"-"`
	Line *AnyOfmicrosoftGraphWorkbookChartLineFormat `json:"line,omitempty"`
	isExplicitNullLine bool `json:"-"`
}

// GetFont returns the Font field if non-nil, zero value otherwise.
func (o *WorkbookChartAxisFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret
	}
	return *o.Font
}

// GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChartAxisFormat) GetFontOk() (AnyOfmicrosoftGraphWorkbookChartFont, bool) {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret, false
	}
	return *o.Font, true
}

// HasFont returns a boolean if a field has been set.
func (o *WorkbookChartAxisFormat) HasFont() bool {
	if o != nil && o.Font != nil {
		return true
	}

	return false
}

// SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFont and assigns it to the Font field.
func (o *WorkbookChartAxisFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont) {
	o.Font = &v
}

// SetFontExplicitNull (un)sets Font to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Font value is set to nil even if false is passed
func (o *WorkbookChartAxisFormat) SetFontExplicitNull(b bool) {
	o.Font = nil
	o.isExplicitNullFont = b
}
// GetLine returns the Line field if non-nil, zero value otherwise.
func (o *WorkbookChartAxisFormat) GetLine() AnyOfmicrosoftGraphWorkbookChartLineFormat {
	if o == nil || o.Line == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartLineFormat
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChartAxisFormat) GetLineOk() (AnyOfmicrosoftGraphWorkbookChartLineFormat, bool) {
	if o == nil || o.Line == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartLineFormat
		return ret, false
	}
	return *o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *WorkbookChartAxisFormat) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given AnyOfmicrosoftGraphWorkbookChartLineFormat and assigns it to the Line field.
func (o *WorkbookChartAxisFormat) SetLine(v AnyOfmicrosoftGraphWorkbookChartLineFormat) {
	o.Line = &v
}

// SetLineExplicitNull (un)sets Line to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Line value is set to nil even if false is passed
func (o *WorkbookChartAxisFormat) SetLineExplicitNull(b bool) {
	o.Line = nil
	o.isExplicitNullLine = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookChartAxisFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Font == nil {
		if o.isExplicitNullFont {
			toSerialize["font"] = o.Font
		}
	} else {
		toSerialize["font"] = o.Font
	}
	if o.Line == nil {
		if o.isExplicitNullLine {
			toSerialize["line"] = o.Line
		}
	} else {
		toSerialize["line"] = o.Line
	}
	return json.Marshal(toSerialize)
}



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
// WorkbookChartAxisTitleFormat struct for WorkbookChartAxisTitleFormat
type WorkbookChartAxisTitleFormat struct {
	Font *AnyOfmicrosoftGraphWorkbookChartFont `json:"font,omitempty"`
	isExplicitNullFont bool `json:"-"`
}

// GetFont returns the Font field if non-nil, zero value otherwise.
func (o *WorkbookChartAxisTitleFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret
	}
	return *o.Font
}

// GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChartAxisTitleFormat) GetFontOk() (AnyOfmicrosoftGraphWorkbookChartFont, bool) {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret, false
	}
	return *o.Font, true
}

// HasFont returns a boolean if a field has been set.
func (o *WorkbookChartAxisTitleFormat) HasFont() bool {
	if o != nil && o.Font != nil {
		return true
	}

	return false
}

// SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFont and assigns it to the Font field.
func (o *WorkbookChartAxisTitleFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont) {
	o.Font = &v
}

// SetFontExplicitNull (un)sets Font to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Font value is set to nil even if false is passed
func (o *WorkbookChartAxisTitleFormat) SetFontExplicitNull(b bool) {
	o.Font = nil
	o.isExplicitNullFont = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookChartAxisTitleFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Font == nil {
		if o.isExplicitNullFont {
			toSerialize["font"] = o.Font
		}
	} else {
		toSerialize["font"] = o.Font
	}
	return json.Marshal(toSerialize)
}



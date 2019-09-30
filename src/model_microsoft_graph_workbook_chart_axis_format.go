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
// MicrosoftGraphWorkbookChartAxisFormat struct for MicrosoftGraphWorkbookChartAxisFormat
type MicrosoftGraphWorkbookChartAxisFormat struct {
	Id *string `json:"id,omitempty"`

	Font *AnyOfmicrosoftGraphWorkbookChartFont `json:"font,omitempty"`
	isExplicitNullFont bool `json:"-"`
	Line *AnyOfmicrosoftGraphWorkbookChartLineFormat `json:"line,omitempty"`
	isExplicitNullLine bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartAxisFormat) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartAxisFormat) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxisFormat) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartAxisFormat) SetId(v string) {
	o.Id = &v
}

// GetFont returns the Font field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartAxisFormat) GetFont() AnyOfmicrosoftGraphWorkbookChartFont {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret
	}
	return *o.Font
}

// GetFontOk returns a tuple with the Font field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartAxisFormat) GetFontOk() (AnyOfmicrosoftGraphWorkbookChartFont, bool) {
	if o == nil || o.Font == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFont
		return ret, false
	}
	return *o.Font, true
}

// HasFont returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxisFormat) HasFont() bool {
	if o != nil && o.Font != nil {
		return true
	}

	return false
}

// SetFont gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFont and assigns it to the Font field.
func (o *MicrosoftGraphWorkbookChartAxisFormat) SetFont(v AnyOfmicrosoftGraphWorkbookChartFont) {
	o.Font = &v
}

// SetFontExplicitNull (un)sets Font to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Font value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartAxisFormat) SetFontExplicitNull(b bool) {
	o.Font = nil
	o.isExplicitNullFont = b
}
// GetLine returns the Line field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartAxisFormat) GetLine() AnyOfmicrosoftGraphWorkbookChartLineFormat {
	if o == nil || o.Line == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartLineFormat
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartAxisFormat) GetLineOk() (AnyOfmicrosoftGraphWorkbookChartLineFormat, bool) {
	if o == nil || o.Line == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartLineFormat
		return ret, false
	}
	return *o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartAxisFormat) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given AnyOfmicrosoftGraphWorkbookChartLineFormat and assigns it to the Line field.
func (o *MicrosoftGraphWorkbookChartAxisFormat) SetLine(v AnyOfmicrosoftGraphWorkbookChartLineFormat) {
	o.Line = &v
}

// SetLineExplicitNull (un)sets Line to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Line value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartAxisFormat) SetLineExplicitNull(b bool) {
	o.Line = nil
	o.isExplicitNullLine = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphWorkbookChartAxisFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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


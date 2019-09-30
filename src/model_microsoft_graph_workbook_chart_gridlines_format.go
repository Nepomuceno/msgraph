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
// MicrosoftGraphWorkbookChartGridlinesFormat struct for MicrosoftGraphWorkbookChartGridlinesFormat
type MicrosoftGraphWorkbookChartGridlinesFormat struct {
	Id *string `json:"id,omitempty"`

	Line *AnyOfmicrosoftGraphWorkbookChartLineFormat `json:"line,omitempty"`
	isExplicitNullLine bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartGridlinesFormat) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartGridlinesFormat) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartGridlinesFormat) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookChartGridlinesFormat) SetId(v string) {
	o.Id = &v
}

// GetLine returns the Line field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookChartGridlinesFormat) GetLine() AnyOfmicrosoftGraphWorkbookChartLineFormat {
	if o == nil || o.Line == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartLineFormat
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookChartGridlinesFormat) GetLineOk() (AnyOfmicrosoftGraphWorkbookChartLineFormat, bool) {
	if o == nil || o.Line == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartLineFormat
		return ret, false
	}
	return *o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookChartGridlinesFormat) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given AnyOfmicrosoftGraphWorkbookChartLineFormat and assigns it to the Line field.
func (o *MicrosoftGraphWorkbookChartGridlinesFormat) SetLine(v AnyOfmicrosoftGraphWorkbookChartLineFormat) {
	o.Line = &v
}

// SetLineExplicitNull (un)sets Line to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Line value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookChartGridlinesFormat) SetLineExplicitNull(b bool) {
	o.Line = nil
	o.isExplicitNullLine = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphWorkbookChartGridlinesFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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


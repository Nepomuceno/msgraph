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
// WorkbookChartPoint struct for WorkbookChartPoint
type WorkbookChartPoint struct {
	Value *AnyOfobject `json:"value,omitempty"`
	isExplicitNullValue bool `json:"-"`
	Format *AnyOfmicrosoftGraphWorkbookChartPointFormat `json:"format,omitempty"`
	isExplicitNullFormat bool `json:"-"`
}

// GetValue returns the Value field if non-nil, zero value otherwise.
func (o *WorkbookChartPoint) GetValue() AnyOfobject {
	if o == nil || o.Value == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChartPoint) GetValueOk() (AnyOfobject, bool) {
	if o == nil || o.Value == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkbookChartPoint) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given AnyOfobject and assigns it to the Value field.
func (o *WorkbookChartPoint) SetValue(v AnyOfobject) {
	o.Value = &v
}

// SetValueExplicitNull (un)sets Value to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Value value is set to nil even if false is passed
func (o *WorkbookChartPoint) SetValueExplicitNull(b bool) {
	o.Value = nil
	o.isExplicitNullValue = b
}
// GetFormat returns the Format field if non-nil, zero value otherwise.
func (o *WorkbookChartPoint) GetFormat() AnyOfmicrosoftGraphWorkbookChartPointFormat {
	if o == nil || o.Format == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartPointFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChartPoint) GetFormatOk() (AnyOfmicrosoftGraphWorkbookChartPointFormat, bool) {
	if o == nil || o.Format == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartPointFormat
		return ret, false
	}
	return *o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WorkbookChartPoint) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given AnyOfmicrosoftGraphWorkbookChartPointFormat and assigns it to the Format field.
func (o *WorkbookChartPoint) SetFormat(v AnyOfmicrosoftGraphWorkbookChartPointFormat) {
	o.Format = &v
}

// SetFormatExplicitNull (un)sets Format to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Format value is set to nil even if false is passed
func (o *WorkbookChartPoint) SetFormatExplicitNull(b bool) {
	o.Format = nil
	o.isExplicitNullFormat = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookChartPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value == nil {
		if o.isExplicitNullValue {
			toSerialize["value"] = o.Value
		}
	} else {
		toSerialize["value"] = o.Value
	}
	if o.Format == nil {
		if o.isExplicitNullFormat {
			toSerialize["format"] = o.Format
		}
	} else {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}



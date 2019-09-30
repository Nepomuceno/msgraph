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
// WorkbookChartPointFormat struct for WorkbookChartPointFormat
type WorkbookChartPointFormat struct {
	Fill *AnyOfmicrosoftGraphWorkbookChartFill `json:"fill,omitempty"`
	isExplicitNullFill bool `json:"-"`
}

// GetFill returns the Fill field if non-nil, zero value otherwise.
func (o *WorkbookChartPointFormat) GetFill() AnyOfmicrosoftGraphWorkbookChartFill {
	if o == nil || o.Fill == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFill
		return ret
	}
	return *o.Fill
}

// GetFillOk returns a tuple with the Fill field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookChartPointFormat) GetFillOk() (AnyOfmicrosoftGraphWorkbookChartFill, bool) {
	if o == nil || o.Fill == nil {
		var ret AnyOfmicrosoftGraphWorkbookChartFill
		return ret, false
	}
	return *o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *WorkbookChartPointFormat) HasFill() bool {
	if o != nil && o.Fill != nil {
		return true
	}

	return false
}

// SetFill gets a reference to the given AnyOfmicrosoftGraphWorkbookChartFill and assigns it to the Fill field.
func (o *WorkbookChartPointFormat) SetFill(v AnyOfmicrosoftGraphWorkbookChartFill) {
	o.Fill = &v
}

// SetFillExplicitNull (un)sets Fill to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Fill value is set to nil even if false is passed
func (o *WorkbookChartPointFormat) SetFillExplicitNull(b bool) {
	o.Fill = nil
	o.isExplicitNullFill = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookChartPointFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fill == nil {
		if o.isExplicitNullFill {
			toSerialize["fill"] = o.Fill
		}
	} else {
		toSerialize["fill"] = o.Fill
	}
	return json.Marshal(toSerialize)
}



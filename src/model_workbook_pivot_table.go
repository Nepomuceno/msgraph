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
// WorkbookPivotTable struct for WorkbookPivotTable
type WorkbookPivotTable struct {
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	Worksheet *AnyOfmicrosoftGraphWorkbookWorksheet `json:"worksheet,omitempty"`
	isExplicitNullWorksheet bool `json:"-"`
}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *WorkbookPivotTable) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookPivotTable) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookPivotTable) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkbookPivotTable) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *WorkbookPivotTable) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetWorksheet returns the Worksheet field if non-nil, zero value otherwise.
func (o *WorkbookPivotTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return *o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookPivotTable) GetWorksheetOk() (AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret, false
	}
	return *o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *WorkbookPivotTable) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *WorkbookPivotTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = &v
}

// SetWorksheetExplicitNull (un)sets Worksheet to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Worksheet value is set to nil even if false is passed
func (o *WorkbookPivotTable) SetWorksheetExplicitNull(b bool) {
	o.Worksheet = nil
	o.isExplicitNullWorksheet = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookPivotTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
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


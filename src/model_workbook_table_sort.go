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
// WorkbookTableSort struct for WorkbookTableSort
type WorkbookTableSort struct {
	Fields *[]AnyOfmicrosoftGraphWorkbookSortField `json:"fields,omitempty"`

	MatchCase *bool `json:"matchCase,omitempty"`

	Method *string `json:"method,omitempty"`

}

// GetFields returns the Fields field if non-nil, zero value otherwise.
func (o *WorkbookTableSort) GetFields() []AnyOfmicrosoftGraphWorkbookSortField {
	if o == nil || o.Fields == nil {
		var ret []AnyOfmicrosoftGraphWorkbookSortField
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTableSort) GetFieldsOk() ([]AnyOfmicrosoftGraphWorkbookSortField, bool) {
	if o == nil || o.Fields == nil {
		var ret []AnyOfmicrosoftGraphWorkbookSortField
		return ret, false
	}
	return *o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *WorkbookTableSort) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []AnyOfmicrosoftGraphWorkbookSortField and assigns it to the Fields field.
func (o *WorkbookTableSort) SetFields(v []AnyOfmicrosoftGraphWorkbookSortField) {
	o.Fields = &v
}

// GetMatchCase returns the MatchCase field if non-nil, zero value otherwise.
func (o *WorkbookTableSort) GetMatchCase() bool {
	if o == nil || o.MatchCase == nil {
		var ret bool
		return ret
	}
	return *o.MatchCase
}

// GetMatchCaseOk returns a tuple with the MatchCase field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTableSort) GetMatchCaseOk() (bool, bool) {
	if o == nil || o.MatchCase == nil {
		var ret bool
		return ret, false
	}
	return *o.MatchCase, true
}

// HasMatchCase returns a boolean if a field has been set.
func (o *WorkbookTableSort) HasMatchCase() bool {
	if o != nil && o.MatchCase != nil {
		return true
	}

	return false
}

// SetMatchCase gets a reference to the given bool and assigns it to the MatchCase field.
func (o *WorkbookTableSort) SetMatchCase(v bool) {
	o.MatchCase = &v
}

// GetMethod returns the Method field if non-nil, zero value otherwise.
func (o *WorkbookTableSort) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTableSort) GetMethodOk() (string, bool) {
	if o == nil || o.Method == nil {
		var ret string
		return ret, false
	}
	return *o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *WorkbookTableSort) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *WorkbookTableSort) SetMethod(v string) {
	o.Method = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o WorkbookTableSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.MatchCase != nil {
		toSerialize["matchCase"] = o.MatchCase
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}



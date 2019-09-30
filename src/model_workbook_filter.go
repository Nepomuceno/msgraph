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
// WorkbookFilter struct for WorkbookFilter
type WorkbookFilter struct {
	Criteria *AnyOfmicrosoftGraphWorkbookFilterCriteria `json:"criteria,omitempty"`
	isExplicitNullCriteria bool `json:"-"`
}

// GetCriteria returns the Criteria field if non-nil, zero value otherwise.
func (o *WorkbookFilter) GetCriteria() AnyOfmicrosoftGraphWorkbookFilterCriteria {
	if o == nil || o.Criteria == nil {
		var ret AnyOfmicrosoftGraphWorkbookFilterCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookFilter) GetCriteriaOk() (AnyOfmicrosoftGraphWorkbookFilterCriteria, bool) {
	if o == nil || o.Criteria == nil {
		var ret AnyOfmicrosoftGraphWorkbookFilterCriteria
		return ret, false
	}
	return *o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *WorkbookFilter) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AnyOfmicrosoftGraphWorkbookFilterCriteria and assigns it to the Criteria field.
func (o *WorkbookFilter) SetCriteria(v AnyOfmicrosoftGraphWorkbookFilterCriteria) {
	o.Criteria = &v
}

// SetCriteriaExplicitNull (un)sets Criteria to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Criteria value is set to nil even if false is passed
func (o *WorkbookFilter) SetCriteriaExplicitNull(b bool) {
	o.Criteria = nil
	o.isExplicitNullCriteria = b
}

// MarshalJSON returns the JSON representation of the model.
func (o WorkbookFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Criteria == nil {
		if o.isExplicitNullCriteria {
			toSerialize["criteria"] = o.Criteria
		}
	} else {
		toSerialize["criteria"] = o.Criteria
	}
	return json.Marshal(toSerialize)
}



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
// MicrosoftGraphWorkbookFilter struct for MicrosoftGraphWorkbookFilter
type MicrosoftGraphWorkbookFilter struct {
	Id *string `json:"id,omitempty"`

	Criteria *AnyOfmicrosoftGraphWorkbookFilterCriteria `json:"criteria,omitempty"`
	isExplicitNullCriteria bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookFilter) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookFilter) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookFilter) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookFilter) SetId(v string) {
	o.Id = &v
}

// GetCriteria returns the Criteria field if non-nil, zero value otherwise.
func (o *MicrosoftGraphWorkbookFilter) GetCriteria() AnyOfmicrosoftGraphWorkbookFilterCriteria {
	if o == nil || o.Criteria == nil {
		var ret AnyOfmicrosoftGraphWorkbookFilterCriteria
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookFilter) GetCriteriaOk() (AnyOfmicrosoftGraphWorkbookFilterCriteria, bool) {
	if o == nil || o.Criteria == nil {
		var ret AnyOfmicrosoftGraphWorkbookFilterCriteria
		return ret, false
	}
	return *o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookFilter) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AnyOfmicrosoftGraphWorkbookFilterCriteria and assigns it to the Criteria field.
func (o *MicrosoftGraphWorkbookFilter) SetCriteria(v AnyOfmicrosoftGraphWorkbookFilterCriteria) {
	o.Criteria = &v
}

// SetCriteriaExplicitNull (un)sets Criteria to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Criteria value is set to nil even if false is passed
func (o *MicrosoftGraphWorkbookFilter) SetCriteriaExplicitNull(b bool) {
	o.Criteria = nil
	o.isExplicitNullCriteria = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphWorkbookFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Criteria == nil {
		if o.isExplicitNullCriteria {
			toSerialize["criteria"] = o.Criteria
		}
	} else {
		toSerialize["criteria"] = o.Criteria
	}
	return json.Marshal(toSerialize)
}


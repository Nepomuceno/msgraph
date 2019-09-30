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
// InlineObject851 struct for InlineObject851
type InlineObject851 struct {
	Database *AnyOfobject `json:"database,omitempty"`
	isExplicitNullDatabase bool `json:"-"`
	Field *AnyOfobject `json:"field,omitempty"`
	isExplicitNullField bool `json:"-"`
	Criteria *AnyOfobject `json:"criteria,omitempty"`
	isExplicitNullCriteria bool `json:"-"`
}

// GetDatabase returns the Database field if non-nil, zero value otherwise.
func (o *InlineObject851) GetDatabase() AnyOfobject {
	if o == nil || o.Database == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject851) GetDatabaseOk() (AnyOfobject, bool) {
	if o == nil || o.Database == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *InlineObject851) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given AnyOfobject and assigns it to the Database field.
func (o *InlineObject851) SetDatabase(v AnyOfobject) {
	o.Database = &v
}

// SetDatabaseExplicitNull (un)sets Database to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Database value is set to nil even if false is passed
func (o *InlineObject851) SetDatabaseExplicitNull(b bool) {
	o.Database = nil
	o.isExplicitNullDatabase = b
}
// GetField returns the Field field if non-nil, zero value otherwise.
func (o *InlineObject851) GetField() AnyOfobject {
	if o == nil || o.Field == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject851) GetFieldOk() (AnyOfobject, bool) {
	if o == nil || o.Field == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *InlineObject851) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given AnyOfobject and assigns it to the Field field.
func (o *InlineObject851) SetField(v AnyOfobject) {
	o.Field = &v
}

// SetFieldExplicitNull (un)sets Field to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Field value is set to nil even if false is passed
func (o *InlineObject851) SetFieldExplicitNull(b bool) {
	o.Field = nil
	o.isExplicitNullField = b
}
// GetCriteria returns the Criteria field if non-nil, zero value otherwise.
func (o *InlineObject851) GetCriteria() AnyOfobject {
	if o == nil || o.Criteria == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject851) GetCriteriaOk() (AnyOfobject, bool) {
	if o == nil || o.Criteria == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *InlineObject851) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AnyOfobject and assigns it to the Criteria field.
func (o *InlineObject851) SetCriteria(v AnyOfobject) {
	o.Criteria = &v
}

// SetCriteriaExplicitNull (un)sets Criteria to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Criteria value is set to nil even if false is passed
func (o *InlineObject851) SetCriteriaExplicitNull(b bool) {
	o.Criteria = nil
	o.isExplicitNullCriteria = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject851) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Database == nil {
		if o.isExplicitNullDatabase {
			toSerialize["database"] = o.Database
		}
	} else {
		toSerialize["database"] = o.Database
	}
	if o.Field == nil {
		if o.isExplicitNullField {
			toSerialize["field"] = o.Field
		}
	} else {
		toSerialize["field"] = o.Field
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


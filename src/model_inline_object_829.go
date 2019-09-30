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
// InlineObject829 struct for InlineObject829
type InlineObject829 struct {
	StartDate *AnyOfobject `json:"startDate,omitempty"`
	isExplicitNullStartDate bool `json:"-"`
	EndDate *AnyOfobject `json:"endDate,omitempty"`
	isExplicitNullEndDate bool `json:"-"`
	Method *AnyOfobject `json:"method,omitempty"`
	isExplicitNullMethod bool `json:"-"`
}

// GetStartDate returns the StartDate field if non-nil, zero value otherwise.
func (o *InlineObject829) GetStartDate() AnyOfobject {
	if o == nil || o.StartDate == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject829) GetStartDateOk() (AnyOfobject, bool) {
	if o == nil || o.StartDate == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InlineObject829) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given AnyOfobject and assigns it to the StartDate field.
func (o *InlineObject829) SetStartDate(v AnyOfobject) {
	o.StartDate = &v
}

// SetStartDateExplicitNull (un)sets StartDate to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The StartDate value is set to nil even if false is passed
func (o *InlineObject829) SetStartDateExplicitNull(b bool) {
	o.StartDate = nil
	o.isExplicitNullStartDate = b
}
// GetEndDate returns the EndDate field if non-nil, zero value otherwise.
func (o *InlineObject829) GetEndDate() AnyOfobject {
	if o == nil || o.EndDate == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject829) GetEndDateOk() (AnyOfobject, bool) {
	if o == nil || o.EndDate == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InlineObject829) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given AnyOfobject and assigns it to the EndDate field.
func (o *InlineObject829) SetEndDate(v AnyOfobject) {
	o.EndDate = &v
}

// SetEndDateExplicitNull (un)sets EndDate to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The EndDate value is set to nil even if false is passed
func (o *InlineObject829) SetEndDateExplicitNull(b bool) {
	o.EndDate = nil
	o.isExplicitNullEndDate = b
}
// GetMethod returns the Method field if non-nil, zero value otherwise.
func (o *InlineObject829) GetMethod() AnyOfobject {
	if o == nil || o.Method == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject829) GetMethodOk() (AnyOfobject, bool) {
	if o == nil || o.Method == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineObject829) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given AnyOfobject and assigns it to the Method field.
func (o *InlineObject829) SetMethod(v AnyOfobject) {
	o.Method = &v
}

// SetMethodExplicitNull (un)sets Method to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Method value is set to nil even if false is passed
func (o *InlineObject829) SetMethodExplicitNull(b bool) {
	o.Method = nil
	o.isExplicitNullMethod = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject829) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate == nil {
		if o.isExplicitNullStartDate {
			toSerialize["startDate"] = o.StartDate
		}
	} else {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate == nil {
		if o.isExplicitNullEndDate {
			toSerialize["endDate"] = o.EndDate
		}
	} else {
		toSerialize["endDate"] = o.EndDate
	}
	if o.Method == nil {
		if o.isExplicitNullMethod {
			toSerialize["method"] = o.Method
		}
	} else {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}



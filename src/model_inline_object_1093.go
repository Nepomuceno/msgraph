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
// InlineObject1093 struct for InlineObject1093
type InlineObject1093 struct {
	StartDate *AnyOfobject `json:"startDate,omitempty"`
	isExplicitNullStartDate bool `json:"-"`
	Days *AnyOfobject `json:"days,omitempty"`
	isExplicitNullDays bool `json:"-"`
	Holidays *AnyOfobject `json:"holidays,omitempty"`
	isExplicitNullHolidays bool `json:"-"`
}

// GetStartDate returns the StartDate field if non-nil, zero value otherwise.
func (o *InlineObject1093) GetStartDate() AnyOfobject {
	if o == nil || o.StartDate == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1093) GetStartDateOk() (AnyOfobject, bool) {
	if o == nil || o.StartDate == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InlineObject1093) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given AnyOfobject and assigns it to the StartDate field.
func (o *InlineObject1093) SetStartDate(v AnyOfobject) {
	o.StartDate = &v
}

// SetStartDateExplicitNull (un)sets StartDate to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The StartDate value is set to nil even if false is passed
func (o *InlineObject1093) SetStartDateExplicitNull(b bool) {
	o.StartDate = nil
	o.isExplicitNullStartDate = b
}
// GetDays returns the Days field if non-nil, zero value otherwise.
func (o *InlineObject1093) GetDays() AnyOfobject {
	if o == nil || o.Days == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1093) GetDaysOk() (AnyOfobject, bool) {
	if o == nil || o.Days == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *InlineObject1093) HasDays() bool {
	if o != nil && o.Days != nil {
		return true
	}

	return false
}

// SetDays gets a reference to the given AnyOfobject and assigns it to the Days field.
func (o *InlineObject1093) SetDays(v AnyOfobject) {
	o.Days = &v
}

// SetDaysExplicitNull (un)sets Days to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Days value is set to nil even if false is passed
func (o *InlineObject1093) SetDaysExplicitNull(b bool) {
	o.Days = nil
	o.isExplicitNullDays = b
}
// GetHolidays returns the Holidays field if non-nil, zero value otherwise.
func (o *InlineObject1093) GetHolidays() AnyOfobject {
	if o == nil || o.Holidays == nil {
		var ret AnyOfobject
		return ret
	}
	return *o.Holidays
}

// GetHolidaysOk returns a tuple with the Holidays field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1093) GetHolidaysOk() (AnyOfobject, bool) {
	if o == nil || o.Holidays == nil {
		var ret AnyOfobject
		return ret, false
	}
	return *o.Holidays, true
}

// HasHolidays returns a boolean if a field has been set.
func (o *InlineObject1093) HasHolidays() bool {
	if o != nil && o.Holidays != nil {
		return true
	}

	return false
}

// SetHolidays gets a reference to the given AnyOfobject and assigns it to the Holidays field.
func (o *InlineObject1093) SetHolidays(v AnyOfobject) {
	o.Holidays = &v
}

// SetHolidaysExplicitNull (un)sets Holidays to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Holidays value is set to nil even if false is passed
func (o *InlineObject1093) SetHolidaysExplicitNull(b bool) {
	o.Holidays = nil
	o.isExplicitNullHolidays = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject1093) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate == nil {
		if o.isExplicitNullStartDate {
			toSerialize["startDate"] = o.StartDate
		}
	} else {
		toSerialize["startDate"] = o.StartDate
	}
	if o.Days == nil {
		if o.isExplicitNullDays {
			toSerialize["days"] = o.Days
		}
	} else {
		toSerialize["days"] = o.Days
	}
	if o.Holidays == nil {
		if o.isExplicitNullHolidays {
			toSerialize["holidays"] = o.Holidays
		}
	} else {
		toSerialize["holidays"] = o.Holidays
	}
	return json.Marshal(toSerialize)
}


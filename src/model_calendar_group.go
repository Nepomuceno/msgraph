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
// CalendarGroup struct for CalendarGroup
type CalendarGroup struct {
	Name *string `json:"name,omitempty"`
	isExplicitNullName bool `json:"-"`
	ClassId *string `json:"classId,omitempty"`
	isExplicitNullClassId bool `json:"-"`
	ChangeKey *string `json:"changeKey,omitempty"`
	isExplicitNullChangeKey bool `json:"-"`
	Calendars *[]MicrosoftGraphCalendar `json:"calendars,omitempty"`

}

// GetName returns the Name field if non-nil, zero value otherwise.
func (o *CalendarGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CalendarGroup) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CalendarGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CalendarGroup) SetName(v string) {
	o.Name = &v
}

// SetNameExplicitNull (un)sets Name to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Name value is set to nil even if false is passed
func (o *CalendarGroup) SetNameExplicitNull(b bool) {
	o.Name = nil
	o.isExplicitNullName = b
}
// GetClassId returns the ClassId field if non-nil, zero value otherwise.
func (o *CalendarGroup) GetClassId() string {
	if o == nil || o.ClassId == nil {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CalendarGroup) GetClassIdOk() (string, bool) {
	if o == nil || o.ClassId == nil {
		var ret string
		return ret, false
	}
	return *o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *CalendarGroup) HasClassId() bool {
	if o != nil && o.ClassId != nil {
		return true
	}

	return false
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *CalendarGroup) SetClassId(v string) {
	o.ClassId = &v
}

// SetClassIdExplicitNull (un)sets ClassId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ClassId value is set to nil even if false is passed
func (o *CalendarGroup) SetClassIdExplicitNull(b bool) {
	o.ClassId = nil
	o.isExplicitNullClassId = b
}
// GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.
func (o *CalendarGroup) GetChangeKey() string {
	if o == nil || o.ChangeKey == nil {
		var ret string
		return ret
	}
	return *o.ChangeKey
}

// GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CalendarGroup) GetChangeKeyOk() (string, bool) {
	if o == nil || o.ChangeKey == nil {
		var ret string
		return ret, false
	}
	return *o.ChangeKey, true
}

// HasChangeKey returns a boolean if a field has been set.
func (o *CalendarGroup) HasChangeKey() bool {
	if o != nil && o.ChangeKey != nil {
		return true
	}

	return false
}

// SetChangeKey gets a reference to the given string and assigns it to the ChangeKey field.
func (o *CalendarGroup) SetChangeKey(v string) {
	o.ChangeKey = &v
}

// SetChangeKeyExplicitNull (un)sets ChangeKey to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ChangeKey value is set to nil even if false is passed
func (o *CalendarGroup) SetChangeKeyExplicitNull(b bool) {
	o.ChangeKey = nil
	o.isExplicitNullChangeKey = b
}
// GetCalendars returns the Calendars field if non-nil, zero value otherwise.
func (o *CalendarGroup) GetCalendars() []MicrosoftGraphCalendar {
	if o == nil || o.Calendars == nil {
		var ret []MicrosoftGraphCalendar
		return ret
	}
	return *o.Calendars
}

// GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CalendarGroup) GetCalendarsOk() ([]MicrosoftGraphCalendar, bool) {
	if o == nil || o.Calendars == nil {
		var ret []MicrosoftGraphCalendar
		return ret, false
	}
	return *o.Calendars, true
}

// HasCalendars returns a boolean if a field has been set.
func (o *CalendarGroup) HasCalendars() bool {
	if o != nil && o.Calendars != nil {
		return true
	}

	return false
}

// SetCalendars gets a reference to the given []MicrosoftGraphCalendar and assigns it to the Calendars field.
func (o *CalendarGroup) SetCalendars(v []MicrosoftGraphCalendar) {
	o.Calendars = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o CalendarGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name == nil {
		if o.isExplicitNullName {
			toSerialize["name"] = o.Name
		}
	} else {
		toSerialize["name"] = o.Name
	}
	if o.ClassId == nil {
		if o.isExplicitNullClassId {
			toSerialize["classId"] = o.ClassId
		}
	} else {
		toSerialize["classId"] = o.ClassId
	}
	if o.ChangeKey == nil {
		if o.isExplicitNullChangeKey {
			toSerialize["changeKey"] = o.ChangeKey
		}
	} else {
		toSerialize["changeKey"] = o.ChangeKey
	}
	if o.Calendars != nil {
		toSerialize["calendars"] = o.Calendars
	}
	return json.Marshal(toSerialize)
}


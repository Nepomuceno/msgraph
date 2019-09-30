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
// MicrosoftGraphEducationTeacher struct for MicrosoftGraphEducationTeacher
type MicrosoftGraphEducationTeacher struct {
	TeacherNumber *string `json:"teacherNumber,omitempty"`
	isExplicitNullTeacherNumber bool `json:"-"`
	ExternalId *string `json:"externalId,omitempty"`
	isExplicitNullExternalId bool `json:"-"`
}

// GetTeacherNumber returns the TeacherNumber field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationTeacher) GetTeacherNumber() string {
	if o == nil || o.TeacherNumber == nil {
		var ret string
		return ret
	}
	return *o.TeacherNumber
}

// GetTeacherNumberOk returns a tuple with the TeacherNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationTeacher) GetTeacherNumberOk() (string, bool) {
	if o == nil || o.TeacherNumber == nil {
		var ret string
		return ret, false
	}
	return *o.TeacherNumber, true
}

// HasTeacherNumber returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationTeacher) HasTeacherNumber() bool {
	if o != nil && o.TeacherNumber != nil {
		return true
	}

	return false
}

// SetTeacherNumber gets a reference to the given string and assigns it to the TeacherNumber field.
func (o *MicrosoftGraphEducationTeacher) SetTeacherNumber(v string) {
	o.TeacherNumber = &v
}

// SetTeacherNumberExplicitNull (un)sets TeacherNumber to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The TeacherNumber value is set to nil even if false is passed
func (o *MicrosoftGraphEducationTeacher) SetTeacherNumberExplicitNull(b bool) {
	o.TeacherNumber = nil
	o.isExplicitNullTeacherNumber = b
}
// GetExternalId returns the ExternalId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationTeacher) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationTeacher) GetExternalIdOk() (string, bool) {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret, false
	}
	return *o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationTeacher) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *MicrosoftGraphEducationTeacher) SetExternalId(v string) {
	o.ExternalId = &v
}

// SetExternalIdExplicitNull (un)sets ExternalId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ExternalId value is set to nil even if false is passed
func (o *MicrosoftGraphEducationTeacher) SetExternalIdExplicitNull(b bool) {
	o.ExternalId = nil
	o.isExplicitNullExternalId = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphEducationTeacher) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TeacherNumber == nil {
		if o.isExplicitNullTeacherNumber {
			toSerialize["teacherNumber"] = o.TeacherNumber
		}
	} else {
		toSerialize["teacherNumber"] = o.TeacherNumber
	}
	if o.ExternalId == nil {
		if o.isExplicitNullExternalId {
			toSerialize["externalId"] = o.ExternalId
		}
	} else {
		toSerialize["externalId"] = o.ExternalId
	}
	return json.Marshal(toSerialize)
}



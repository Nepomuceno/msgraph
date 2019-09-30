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
// EducationSchool struct for EducationSchool
type EducationSchool struct {
	PrincipalEmail *string `json:"principalEmail,omitempty"`
	isExplicitNullPrincipalEmail bool `json:"-"`
	PrincipalName *string `json:"principalName,omitempty"`
	isExplicitNullPrincipalName bool `json:"-"`
	ExternalPrincipalId *string `json:"externalPrincipalId,omitempty"`
	isExplicitNullExternalPrincipalId bool `json:"-"`
	LowestGrade *string `json:"lowestGrade,omitempty"`
	isExplicitNullLowestGrade bool `json:"-"`
	HighestGrade *string `json:"highestGrade,omitempty"`
	isExplicitNullHighestGrade bool `json:"-"`
	SchoolNumber *string `json:"schoolNumber,omitempty"`
	isExplicitNullSchoolNumber bool `json:"-"`
	ExternalId *string `json:"externalId,omitempty"`
	isExplicitNullExternalId bool `json:"-"`
	Phone *string `json:"phone,omitempty"`
	isExplicitNullPhone bool `json:"-"`
	Fax *string `json:"fax,omitempty"`
	isExplicitNullFax bool `json:"-"`
	CreatedBy *AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	isExplicitNullCreatedBy bool `json:"-"`
	Address *AnyOfmicrosoftGraphPhysicalAddress `json:"address,omitempty"`
	isExplicitNullAddress bool `json:"-"`
	Classes *[]MicrosoftGraphEducationClass `json:"classes,omitempty"`

	Users *[]MicrosoftGraphEducationUser `json:"users,omitempty"`

}

// GetPrincipalEmail returns the PrincipalEmail field if non-nil, zero value otherwise.
func (o *EducationSchool) GetPrincipalEmail() string {
	if o == nil || o.PrincipalEmail == nil {
		var ret string
		return ret
	}
	return *o.PrincipalEmail
}

// GetPrincipalEmailOk returns a tuple with the PrincipalEmail field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetPrincipalEmailOk() (string, bool) {
	if o == nil || o.PrincipalEmail == nil {
		var ret string
		return ret, false
	}
	return *o.PrincipalEmail, true
}

// HasPrincipalEmail returns a boolean if a field has been set.
func (o *EducationSchool) HasPrincipalEmail() bool {
	if o != nil && o.PrincipalEmail != nil {
		return true
	}

	return false
}

// SetPrincipalEmail gets a reference to the given string and assigns it to the PrincipalEmail field.
func (o *EducationSchool) SetPrincipalEmail(v string) {
	o.PrincipalEmail = &v
}

// SetPrincipalEmailExplicitNull (un)sets PrincipalEmail to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PrincipalEmail value is set to nil even if false is passed
func (o *EducationSchool) SetPrincipalEmailExplicitNull(b bool) {
	o.PrincipalEmail = nil
	o.isExplicitNullPrincipalEmail = b
}
// GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.
func (o *EducationSchool) GetPrincipalName() string {
	if o == nil || o.PrincipalName == nil {
		var ret string
		return ret
	}
	return *o.PrincipalName
}

// GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetPrincipalNameOk() (string, bool) {
	if o == nil || o.PrincipalName == nil {
		var ret string
		return ret, false
	}
	return *o.PrincipalName, true
}

// HasPrincipalName returns a boolean if a field has been set.
func (o *EducationSchool) HasPrincipalName() bool {
	if o != nil && o.PrincipalName != nil {
		return true
	}

	return false
}

// SetPrincipalName gets a reference to the given string and assigns it to the PrincipalName field.
func (o *EducationSchool) SetPrincipalName(v string) {
	o.PrincipalName = &v
}

// SetPrincipalNameExplicitNull (un)sets PrincipalName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PrincipalName value is set to nil even if false is passed
func (o *EducationSchool) SetPrincipalNameExplicitNull(b bool) {
	o.PrincipalName = nil
	o.isExplicitNullPrincipalName = b
}
// GetExternalPrincipalId returns the ExternalPrincipalId field if non-nil, zero value otherwise.
func (o *EducationSchool) GetExternalPrincipalId() string {
	if o == nil || o.ExternalPrincipalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalPrincipalId
}

// GetExternalPrincipalIdOk returns a tuple with the ExternalPrincipalId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetExternalPrincipalIdOk() (string, bool) {
	if o == nil || o.ExternalPrincipalId == nil {
		var ret string
		return ret, false
	}
	return *o.ExternalPrincipalId, true
}

// HasExternalPrincipalId returns a boolean if a field has been set.
func (o *EducationSchool) HasExternalPrincipalId() bool {
	if o != nil && o.ExternalPrincipalId != nil {
		return true
	}

	return false
}

// SetExternalPrincipalId gets a reference to the given string and assigns it to the ExternalPrincipalId field.
func (o *EducationSchool) SetExternalPrincipalId(v string) {
	o.ExternalPrincipalId = &v
}

// SetExternalPrincipalIdExplicitNull (un)sets ExternalPrincipalId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ExternalPrincipalId value is set to nil even if false is passed
func (o *EducationSchool) SetExternalPrincipalIdExplicitNull(b bool) {
	o.ExternalPrincipalId = nil
	o.isExplicitNullExternalPrincipalId = b
}
// GetLowestGrade returns the LowestGrade field if non-nil, zero value otherwise.
func (o *EducationSchool) GetLowestGrade() string {
	if o == nil || o.LowestGrade == nil {
		var ret string
		return ret
	}
	return *o.LowestGrade
}

// GetLowestGradeOk returns a tuple with the LowestGrade field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetLowestGradeOk() (string, bool) {
	if o == nil || o.LowestGrade == nil {
		var ret string
		return ret, false
	}
	return *o.LowestGrade, true
}

// HasLowestGrade returns a boolean if a field has been set.
func (o *EducationSchool) HasLowestGrade() bool {
	if o != nil && o.LowestGrade != nil {
		return true
	}

	return false
}

// SetLowestGrade gets a reference to the given string and assigns it to the LowestGrade field.
func (o *EducationSchool) SetLowestGrade(v string) {
	o.LowestGrade = &v
}

// SetLowestGradeExplicitNull (un)sets LowestGrade to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LowestGrade value is set to nil even if false is passed
func (o *EducationSchool) SetLowestGradeExplicitNull(b bool) {
	o.LowestGrade = nil
	o.isExplicitNullLowestGrade = b
}
// GetHighestGrade returns the HighestGrade field if non-nil, zero value otherwise.
func (o *EducationSchool) GetHighestGrade() string {
	if o == nil || o.HighestGrade == nil {
		var ret string
		return ret
	}
	return *o.HighestGrade
}

// GetHighestGradeOk returns a tuple with the HighestGrade field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetHighestGradeOk() (string, bool) {
	if o == nil || o.HighestGrade == nil {
		var ret string
		return ret, false
	}
	return *o.HighestGrade, true
}

// HasHighestGrade returns a boolean if a field has been set.
func (o *EducationSchool) HasHighestGrade() bool {
	if o != nil && o.HighestGrade != nil {
		return true
	}

	return false
}

// SetHighestGrade gets a reference to the given string and assigns it to the HighestGrade field.
func (o *EducationSchool) SetHighestGrade(v string) {
	o.HighestGrade = &v
}

// SetHighestGradeExplicitNull (un)sets HighestGrade to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The HighestGrade value is set to nil even if false is passed
func (o *EducationSchool) SetHighestGradeExplicitNull(b bool) {
	o.HighestGrade = nil
	o.isExplicitNullHighestGrade = b
}
// GetSchoolNumber returns the SchoolNumber field if non-nil, zero value otherwise.
func (o *EducationSchool) GetSchoolNumber() string {
	if o == nil || o.SchoolNumber == nil {
		var ret string
		return ret
	}
	return *o.SchoolNumber
}

// GetSchoolNumberOk returns a tuple with the SchoolNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetSchoolNumberOk() (string, bool) {
	if o == nil || o.SchoolNumber == nil {
		var ret string
		return ret, false
	}
	return *o.SchoolNumber, true
}

// HasSchoolNumber returns a boolean if a field has been set.
func (o *EducationSchool) HasSchoolNumber() bool {
	if o != nil && o.SchoolNumber != nil {
		return true
	}

	return false
}

// SetSchoolNumber gets a reference to the given string and assigns it to the SchoolNumber field.
func (o *EducationSchool) SetSchoolNumber(v string) {
	o.SchoolNumber = &v
}

// SetSchoolNumberExplicitNull (un)sets SchoolNumber to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SchoolNumber value is set to nil even if false is passed
func (o *EducationSchool) SetSchoolNumberExplicitNull(b bool) {
	o.SchoolNumber = nil
	o.isExplicitNullSchoolNumber = b
}
// GetExternalId returns the ExternalId field if non-nil, zero value otherwise.
func (o *EducationSchool) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetExternalIdOk() (string, bool) {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret, false
	}
	return *o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *EducationSchool) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *EducationSchool) SetExternalId(v string) {
	o.ExternalId = &v
}

// SetExternalIdExplicitNull (un)sets ExternalId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ExternalId value is set to nil even if false is passed
func (o *EducationSchool) SetExternalIdExplicitNull(b bool) {
	o.ExternalId = nil
	o.isExplicitNullExternalId = b
}
// GetPhone returns the Phone field if non-nil, zero value otherwise.
func (o *EducationSchool) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetPhoneOk() (string, bool) {
	if o == nil || o.Phone == nil {
		var ret string
		return ret, false
	}
	return *o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *EducationSchool) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *EducationSchool) SetPhone(v string) {
	o.Phone = &v
}

// SetPhoneExplicitNull (un)sets Phone to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Phone value is set to nil even if false is passed
func (o *EducationSchool) SetPhoneExplicitNull(b bool) {
	o.Phone = nil
	o.isExplicitNullPhone = b
}
// GetFax returns the Fax field if non-nil, zero value otherwise.
func (o *EducationSchool) GetFax() string {
	if o == nil || o.Fax == nil {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetFaxOk() (string, bool) {
	if o == nil || o.Fax == nil {
		var ret string
		return ret, false
	}
	return *o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *EducationSchool) HasFax() bool {
	if o != nil && o.Fax != nil {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *EducationSchool) SetFax(v string) {
	o.Fax = &v
}

// SetFaxExplicitNull (un)sets Fax to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Fax value is set to nil even if false is passed
func (o *EducationSchool) SetFaxExplicitNull(b bool) {
	o.Fax = nil
	o.isExplicitNullFax = b
}
// GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.
func (o *EducationSchool) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *EducationSchool) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *EducationSchool) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = &v
}

// SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedBy value is set to nil even if false is passed
func (o *EducationSchool) SetCreatedByExplicitNull(b bool) {
	o.CreatedBy = nil
	o.isExplicitNullCreatedBy = b
}
// GetAddress returns the Address field if non-nil, zero value otherwise.
func (o *EducationSchool) GetAddress() AnyOfmicrosoftGraphPhysicalAddress {
	if o == nil || o.Address == nil {
		var ret AnyOfmicrosoftGraphPhysicalAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetAddressOk() (AnyOfmicrosoftGraphPhysicalAddress, bool) {
	if o == nil || o.Address == nil {
		var ret AnyOfmicrosoftGraphPhysicalAddress
		return ret, false
	}
	return *o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *EducationSchool) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the Address field.
func (o *EducationSchool) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress) {
	o.Address = &v
}

// SetAddressExplicitNull (un)sets Address to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Address value is set to nil even if false is passed
func (o *EducationSchool) SetAddressExplicitNull(b bool) {
	o.Address = nil
	o.isExplicitNullAddress = b
}
// GetClasses returns the Classes field if non-nil, zero value otherwise.
func (o *EducationSchool) GetClasses() []MicrosoftGraphEducationClass {
	if o == nil || o.Classes == nil {
		var ret []MicrosoftGraphEducationClass
		return ret
	}
	return *o.Classes
}

// GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetClassesOk() ([]MicrosoftGraphEducationClass, bool) {
	if o == nil || o.Classes == nil {
		var ret []MicrosoftGraphEducationClass
		return ret, false
	}
	return *o.Classes, true
}

// HasClasses returns a boolean if a field has been set.
func (o *EducationSchool) HasClasses() bool {
	if o != nil && o.Classes != nil {
		return true
	}

	return false
}

// SetClasses gets a reference to the given []MicrosoftGraphEducationClass and assigns it to the Classes field.
func (o *EducationSchool) SetClasses(v []MicrosoftGraphEducationClass) {
	o.Classes = &v
}

// GetUsers returns the Users field if non-nil, zero value otherwise.
func (o *EducationSchool) GetUsers() []MicrosoftGraphEducationUser {
	if o == nil || o.Users == nil {
		var ret []MicrosoftGraphEducationUser
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *EducationSchool) GetUsersOk() ([]MicrosoftGraphEducationUser, bool) {
	if o == nil || o.Users == nil {
		var ret []MicrosoftGraphEducationUser
		return ret, false
	}
	return *o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *EducationSchool) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Users field.
func (o *EducationSchool) SetUsers(v []MicrosoftGraphEducationUser) {
	o.Users = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o EducationSchool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrincipalEmail == nil {
		if o.isExplicitNullPrincipalEmail {
			toSerialize["principalEmail"] = o.PrincipalEmail
		}
	} else {
		toSerialize["principalEmail"] = o.PrincipalEmail
	}
	if o.PrincipalName == nil {
		if o.isExplicitNullPrincipalName {
			toSerialize["principalName"] = o.PrincipalName
		}
	} else {
		toSerialize["principalName"] = o.PrincipalName
	}
	if o.ExternalPrincipalId == nil {
		if o.isExplicitNullExternalPrincipalId {
			toSerialize["externalPrincipalId"] = o.ExternalPrincipalId
		}
	} else {
		toSerialize["externalPrincipalId"] = o.ExternalPrincipalId
	}
	if o.LowestGrade == nil {
		if o.isExplicitNullLowestGrade {
			toSerialize["lowestGrade"] = o.LowestGrade
		}
	} else {
		toSerialize["lowestGrade"] = o.LowestGrade
	}
	if o.HighestGrade == nil {
		if o.isExplicitNullHighestGrade {
			toSerialize["highestGrade"] = o.HighestGrade
		}
	} else {
		toSerialize["highestGrade"] = o.HighestGrade
	}
	if o.SchoolNumber == nil {
		if o.isExplicitNullSchoolNumber {
			toSerialize["schoolNumber"] = o.SchoolNumber
		}
	} else {
		toSerialize["schoolNumber"] = o.SchoolNumber
	}
	if o.ExternalId == nil {
		if o.isExplicitNullExternalId {
			toSerialize["externalId"] = o.ExternalId
		}
	} else {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Phone == nil {
		if o.isExplicitNullPhone {
			toSerialize["phone"] = o.Phone
		}
	} else {
		toSerialize["phone"] = o.Phone
	}
	if o.Fax == nil {
		if o.isExplicitNullFax {
			toSerialize["fax"] = o.Fax
		}
	} else {
		toSerialize["fax"] = o.Fax
	}
	if o.CreatedBy == nil {
		if o.isExplicitNullCreatedBy {
			toSerialize["createdBy"] = o.CreatedBy
		}
	} else {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Address == nil {
		if o.isExplicitNullAddress {
			toSerialize["address"] = o.Address
		}
	} else {
		toSerialize["address"] = o.Address
	}
	if o.Classes != nil {
		toSerialize["classes"] = o.Classes
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}


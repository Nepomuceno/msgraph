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
// MicrosoftGraphEducationClass struct for MicrosoftGraphEducationClass
type MicrosoftGraphEducationClass struct {
	Id *string `json:"id,omitempty"`

	DisplayName *string `json:"displayName,omitempty"`

	MailNickname *string `json:"mailNickname,omitempty"`

	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	CreatedBy *AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	isExplicitNullCreatedBy bool `json:"-"`
	ClassCode *string `json:"classCode,omitempty"`
	isExplicitNullClassCode bool `json:"-"`
	ExternalName *string `json:"externalName,omitempty"`
	isExplicitNullExternalName bool `json:"-"`
	ExternalId *string `json:"externalId,omitempty"`
	isExplicitNullExternalId bool `json:"-"`
	ExternalSource *AnyOfmicrosoftGraphEducationExternalSource `json:"externalSource,omitempty"`
	isExplicitNullExternalSource bool `json:"-"`
	Term *AnyOfmicrosoftGraphEducationTerm `json:"term,omitempty"`
	isExplicitNullTerm bool `json:"-"`
	Schools *[]MicrosoftGraphEducationSchool `json:"schools,omitempty"`

	Members *[]MicrosoftGraphEducationUser `json:"members,omitempty"`

	Teachers *[]MicrosoftGraphEducationUser `json:"teachers,omitempty"`

	Group *AnyOfmicrosoftGraphGroup `json:"group,omitempty"`
	isExplicitNullGroup bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphEducationClass) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphEducationClass) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetMailNickname returns the MailNickname field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetMailNickname() string {
	if o == nil || o.MailNickname == nil {
		var ret string
		return ret
	}
	return *o.MailNickname
}

// GetMailNicknameOk returns a tuple with the MailNickname field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetMailNicknameOk() (string, bool) {
	if o == nil || o.MailNickname == nil {
		var ret string
		return ret, false
	}
	return *o.MailNickname, true
}

// HasMailNickname returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasMailNickname() bool {
	if o != nil && o.MailNickname != nil {
		return true
	}

	return false
}

// SetMailNickname gets a reference to the given string and assigns it to the MailNickname field.
func (o *MicrosoftGraphEducationClass) SetMailNickname(v string) {
	o.MailNickname = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphEducationClass) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphEducationClass) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetCreatedByOk() (AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret, false
	}
	return *o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphEducationClass) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = &v
}

// SetCreatedByExplicitNull (un)sets CreatedBy to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedBy value is set to nil even if false is passed
func (o *MicrosoftGraphEducationClass) SetCreatedByExplicitNull(b bool) {
	o.CreatedBy = nil
	o.isExplicitNullCreatedBy = b
}
// GetClassCode returns the ClassCode field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetClassCodeOk() (string, bool) {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret, false
	}
	return *o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasClassCode() bool {
	if o != nil && o.ClassCode != nil {
		return true
	}

	return false
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *MicrosoftGraphEducationClass) SetClassCode(v string) {
	o.ClassCode = &v
}

// SetClassCodeExplicitNull (un)sets ClassCode to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ClassCode value is set to nil even if false is passed
func (o *MicrosoftGraphEducationClass) SetClassCodeExplicitNull(b bool) {
	o.ClassCode = nil
	o.isExplicitNullClassCode = b
}
// GetExternalName returns the ExternalName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetExternalName() string {
	if o == nil || o.ExternalName == nil {
		var ret string
		return ret
	}
	return *o.ExternalName
}

// GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetExternalNameOk() (string, bool) {
	if o == nil || o.ExternalName == nil {
		var ret string
		return ret, false
	}
	return *o.ExternalName, true
}

// HasExternalName returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasExternalName() bool {
	if o != nil && o.ExternalName != nil {
		return true
	}

	return false
}

// SetExternalName gets a reference to the given string and assigns it to the ExternalName field.
func (o *MicrosoftGraphEducationClass) SetExternalName(v string) {
	o.ExternalName = &v
}

// SetExternalNameExplicitNull (un)sets ExternalName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ExternalName value is set to nil even if false is passed
func (o *MicrosoftGraphEducationClass) SetExternalNameExplicitNull(b bool) {
	o.ExternalName = nil
	o.isExplicitNullExternalName = b
}
// GetExternalId returns the ExternalId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetExternalIdOk() (string, bool) {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret, false
	}
	return *o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *MicrosoftGraphEducationClass) SetExternalId(v string) {
	o.ExternalId = &v
}

// SetExternalIdExplicitNull (un)sets ExternalId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ExternalId value is set to nil even if false is passed
func (o *MicrosoftGraphEducationClass) SetExternalIdExplicitNull(b bool) {
	o.ExternalId = nil
	o.isExplicitNullExternalId = b
}
// GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetExternalSource() AnyOfmicrosoftGraphEducationExternalSource {
	if o == nil || o.ExternalSource == nil {
		var ret AnyOfmicrosoftGraphEducationExternalSource
		return ret
	}
	return *o.ExternalSource
}

// GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetExternalSourceOk() (AnyOfmicrosoftGraphEducationExternalSource, bool) {
	if o == nil || o.ExternalSource == nil {
		var ret AnyOfmicrosoftGraphEducationExternalSource
		return ret, false
	}
	return *o.ExternalSource, true
}

// HasExternalSource returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasExternalSource() bool {
	if o != nil && o.ExternalSource != nil {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given AnyOfmicrosoftGraphEducationExternalSource and assigns it to the ExternalSource field.
func (o *MicrosoftGraphEducationClass) SetExternalSource(v AnyOfmicrosoftGraphEducationExternalSource) {
	o.ExternalSource = &v
}

// SetExternalSourceExplicitNull (un)sets ExternalSource to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ExternalSource value is set to nil even if false is passed
func (o *MicrosoftGraphEducationClass) SetExternalSourceExplicitNull(b bool) {
	o.ExternalSource = nil
	o.isExplicitNullExternalSource = b
}
// GetTerm returns the Term field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetTerm() AnyOfmicrosoftGraphEducationTerm {
	if o == nil || o.Term == nil {
		var ret AnyOfmicrosoftGraphEducationTerm
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetTermOk() (AnyOfmicrosoftGraphEducationTerm, bool) {
	if o == nil || o.Term == nil {
		var ret AnyOfmicrosoftGraphEducationTerm
		return ret, false
	}
	return *o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasTerm() bool {
	if o != nil && o.Term != nil {
		return true
	}

	return false
}

// SetTerm gets a reference to the given AnyOfmicrosoftGraphEducationTerm and assigns it to the Term field.
func (o *MicrosoftGraphEducationClass) SetTerm(v AnyOfmicrosoftGraphEducationTerm) {
	o.Term = &v
}

// SetTermExplicitNull (un)sets Term to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Term value is set to nil even if false is passed
func (o *MicrosoftGraphEducationClass) SetTermExplicitNull(b bool) {
	o.Term = nil
	o.isExplicitNullTerm = b
}
// GetSchools returns the Schools field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetSchools() []MicrosoftGraphEducationSchool {
	if o == nil || o.Schools == nil {
		var ret []MicrosoftGraphEducationSchool
		return ret
	}
	return *o.Schools
}

// GetSchoolsOk returns a tuple with the Schools field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetSchoolsOk() ([]MicrosoftGraphEducationSchool, bool) {
	if o == nil || o.Schools == nil {
		var ret []MicrosoftGraphEducationSchool
		return ret, false
	}
	return *o.Schools, true
}

// HasSchools returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasSchools() bool {
	if o != nil && o.Schools != nil {
		return true
	}

	return false
}

// SetSchools gets a reference to the given []MicrosoftGraphEducationSchool and assigns it to the Schools field.
func (o *MicrosoftGraphEducationClass) SetSchools(v []MicrosoftGraphEducationSchool) {
	o.Schools = &v
}

// GetMembers returns the Members field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetMembers() []MicrosoftGraphEducationUser {
	if o == nil || o.Members == nil {
		var ret []MicrosoftGraphEducationUser
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetMembersOk() ([]MicrosoftGraphEducationUser, bool) {
	if o == nil || o.Members == nil {
		var ret []MicrosoftGraphEducationUser
		return ret, false
	}
	return *o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Members field.
func (o *MicrosoftGraphEducationClass) SetMembers(v []MicrosoftGraphEducationUser) {
	o.Members = &v
}

// GetTeachers returns the Teachers field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetTeachers() []MicrosoftGraphEducationUser {
	if o == nil || o.Teachers == nil {
		var ret []MicrosoftGraphEducationUser
		return ret
	}
	return *o.Teachers
}

// GetTeachersOk returns a tuple with the Teachers field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetTeachersOk() ([]MicrosoftGraphEducationUser, bool) {
	if o == nil || o.Teachers == nil {
		var ret []MicrosoftGraphEducationUser
		return ret, false
	}
	return *o.Teachers, true
}

// HasTeachers returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasTeachers() bool {
	if o != nil && o.Teachers != nil {
		return true
	}

	return false
}

// SetTeachers gets a reference to the given []MicrosoftGraphEducationUser and assigns it to the Teachers field.
func (o *MicrosoftGraphEducationClass) SetTeachers(v []MicrosoftGraphEducationUser) {
	o.Teachers = &v
}

// GetGroup returns the Group field if non-nil, zero value otherwise.
func (o *MicrosoftGraphEducationClass) GetGroup() AnyOfmicrosoftGraphGroup {
	if o == nil || o.Group == nil {
		var ret AnyOfmicrosoftGraphGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationClass) GetGroupOk() (AnyOfmicrosoftGraphGroup, bool) {
	if o == nil || o.Group == nil {
		var ret AnyOfmicrosoftGraphGroup
		return ret, false
	}
	return *o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationClass) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given AnyOfmicrosoftGraphGroup and assigns it to the Group field.
func (o *MicrosoftGraphEducationClass) SetGroup(v AnyOfmicrosoftGraphGroup) {
	o.Group = &v
}

// SetGroupExplicitNull (un)sets Group to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Group value is set to nil even if false is passed
func (o *MicrosoftGraphEducationClass) SetGroupExplicitNull(b bool) {
	o.Group = nil
	o.isExplicitNullGroup = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphEducationClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.MailNickname != nil {
		toSerialize["mailNickname"] = o.MailNickname
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.CreatedBy == nil {
		if o.isExplicitNullCreatedBy {
			toSerialize["createdBy"] = o.CreatedBy
		}
	} else {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.ClassCode == nil {
		if o.isExplicitNullClassCode {
			toSerialize["classCode"] = o.ClassCode
		}
	} else {
		toSerialize["classCode"] = o.ClassCode
	}
	if o.ExternalName == nil {
		if o.isExplicitNullExternalName {
			toSerialize["externalName"] = o.ExternalName
		}
	} else {
		toSerialize["externalName"] = o.ExternalName
	}
	if o.ExternalId == nil {
		if o.isExplicitNullExternalId {
			toSerialize["externalId"] = o.ExternalId
		}
	} else {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.ExternalSource == nil {
		if o.isExplicitNullExternalSource {
			toSerialize["externalSource"] = o.ExternalSource
		}
	} else {
		toSerialize["externalSource"] = o.ExternalSource
	}
	if o.Term == nil {
		if o.isExplicitNullTerm {
			toSerialize["term"] = o.Term
		}
	} else {
		toSerialize["term"] = o.Term
	}
	if o.Schools != nil {
		toSerialize["schools"] = o.Schools
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Teachers != nil {
		toSerialize["teachers"] = o.Teachers
	}
	if o.Group == nil {
		if o.isExplicitNullGroup {
			toSerialize["group"] = o.Group
		}
	} else {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}



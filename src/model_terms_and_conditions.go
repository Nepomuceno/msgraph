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
	"time"
	"encoding/json"
)
// TermsAndConditions A termsAndConditions entity represents the metadata and contents of a given Terms and Conditions (T&C) policy. T&C policies’ contents are presented to users upon their first attempt to enroll into Intune and subsequently upon edits where an administrator has required re-acceptance. They enable administrators to communicate the provisions to which a user must agree in order to have devices enrolled into Intune.
type TermsAndConditions struct {
	// DateTime the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`

	// Administrator-supplied name for the T&C policy. 
	DisplayName *string `json:"displayName,omitempty"`

	// Administrator-supplied description of the T&C policy.
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	// Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&C policy.
	Title *string `json:"title,omitempty"`
	isExplicitNullTitle bool `json:"-"`
	// Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&C policy.
	BodyText *string `json:"bodyText,omitempty"`
	isExplicitNullBodyText bool `json:"-"`
	// Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&C policy. This is shown to the user on prompts to accept the T&C policy.
	AcceptanceStatement *string `json:"acceptanceStatement,omitempty"`
	isExplicitNullAcceptanceStatement bool `json:"-"`
	// Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&C policy.
	Version *int32 `json:"version,omitempty"`

	Assignments *[]MicrosoftGraphTermsAndConditionsAssignment `json:"assignments,omitempty"`

	AcceptanceStatuses *[]MicrosoftGraphTermsAndConditionsAcceptanceStatus `json:"acceptanceStatuses,omitempty"`

}

// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *TermsAndConditions) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *TermsAndConditions) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *TermsAndConditions) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *TermsAndConditions) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TermsAndConditions) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TermsAndConditions) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TermsAndConditions) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TermsAndConditions) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *TermsAndConditions) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetTitle returns the Title field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TermsAndConditions) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TermsAndConditions) SetTitle(v string) {
	o.Title = &v
}

// SetTitleExplicitNull (un)sets Title to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Title value is set to nil even if false is passed
func (o *TermsAndConditions) SetTitleExplicitNull(b bool) {
	o.Title = nil
	o.isExplicitNullTitle = b
}
// GetBodyText returns the BodyText field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetBodyText() string {
	if o == nil || o.BodyText == nil {
		var ret string
		return ret
	}
	return *o.BodyText
}

// GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetBodyTextOk() (string, bool) {
	if o == nil || o.BodyText == nil {
		var ret string
		return ret, false
	}
	return *o.BodyText, true
}

// HasBodyText returns a boolean if a field has been set.
func (o *TermsAndConditions) HasBodyText() bool {
	if o != nil && o.BodyText != nil {
		return true
	}

	return false
}

// SetBodyText gets a reference to the given string and assigns it to the BodyText field.
func (o *TermsAndConditions) SetBodyText(v string) {
	o.BodyText = &v
}

// SetBodyTextExplicitNull (un)sets BodyText to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The BodyText value is set to nil even if false is passed
func (o *TermsAndConditions) SetBodyTextExplicitNull(b bool) {
	o.BodyText = nil
	o.isExplicitNullBodyText = b
}
// GetAcceptanceStatement returns the AcceptanceStatement field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetAcceptanceStatement() string {
	if o == nil || o.AcceptanceStatement == nil {
		var ret string
		return ret
	}
	return *o.AcceptanceStatement
}

// GetAcceptanceStatementOk returns a tuple with the AcceptanceStatement field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetAcceptanceStatementOk() (string, bool) {
	if o == nil || o.AcceptanceStatement == nil {
		var ret string
		return ret, false
	}
	return *o.AcceptanceStatement, true
}

// HasAcceptanceStatement returns a boolean if a field has been set.
func (o *TermsAndConditions) HasAcceptanceStatement() bool {
	if o != nil && o.AcceptanceStatement != nil {
		return true
	}

	return false
}

// SetAcceptanceStatement gets a reference to the given string and assigns it to the AcceptanceStatement field.
func (o *TermsAndConditions) SetAcceptanceStatement(v string) {
	o.AcceptanceStatement = &v
}

// SetAcceptanceStatementExplicitNull (un)sets AcceptanceStatement to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AcceptanceStatement value is set to nil even if false is passed
func (o *TermsAndConditions) SetAcceptanceStatementExplicitNull(b bool) {
	o.AcceptanceStatement = nil
	o.isExplicitNullAcceptanceStatement = b
}
// GetVersion returns the Version field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetVersionOk() (int32, bool) {
	if o == nil || o.Version == nil {
		var ret int32
		return ret, false
	}
	return *o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TermsAndConditions) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *TermsAndConditions) SetVersion(v int32) {
	o.Version = &v
}

// GetAssignments returns the Assignments field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetAssignments() []MicrosoftGraphTermsAndConditionsAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphTermsAndConditionsAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetAssignmentsOk() ([]MicrosoftGraphTermsAndConditionsAssignment, bool) {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphTermsAndConditionsAssignment
		return ret, false
	}
	return *o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *TermsAndConditions) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphTermsAndConditionsAssignment and assigns it to the Assignments field.
func (o *TermsAndConditions) SetAssignments(v []MicrosoftGraphTermsAndConditionsAssignment) {
	o.Assignments = &v
}

// GetAcceptanceStatuses returns the AcceptanceStatuses field if non-nil, zero value otherwise.
func (o *TermsAndConditions) GetAcceptanceStatuses() []MicrosoftGraphTermsAndConditionsAcceptanceStatus {
	if o == nil || o.AcceptanceStatuses == nil {
		var ret []MicrosoftGraphTermsAndConditionsAcceptanceStatus
		return ret
	}
	return *o.AcceptanceStatuses
}

// GetAcceptanceStatusesOk returns a tuple with the AcceptanceStatuses field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetAcceptanceStatusesOk() ([]MicrosoftGraphTermsAndConditionsAcceptanceStatus, bool) {
	if o == nil || o.AcceptanceStatuses == nil {
		var ret []MicrosoftGraphTermsAndConditionsAcceptanceStatus
		return ret, false
	}
	return *o.AcceptanceStatuses, true
}

// HasAcceptanceStatuses returns a boolean if a field has been set.
func (o *TermsAndConditions) HasAcceptanceStatuses() bool {
	if o != nil && o.AcceptanceStatuses != nil {
		return true
	}

	return false
}

// SetAcceptanceStatuses gets a reference to the given []MicrosoftGraphTermsAndConditionsAcceptanceStatus and assigns it to the AcceptanceStatuses field.
func (o *TermsAndConditions) SetAcceptanceStatuses(v []MicrosoftGraphTermsAndConditionsAcceptanceStatus) {
	o.AcceptanceStatuses = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o TermsAndConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.Title == nil {
		if o.isExplicitNullTitle {
			toSerialize["title"] = o.Title
		}
	} else {
		toSerialize["title"] = o.Title
	}
	if o.BodyText == nil {
		if o.isExplicitNullBodyText {
			toSerialize["bodyText"] = o.BodyText
		}
	} else {
		toSerialize["bodyText"] = o.BodyText
	}
	if o.AcceptanceStatement == nil {
		if o.isExplicitNullAcceptanceStatement {
			toSerialize["acceptanceStatement"] = o.AcceptanceStatement
		}
	} else {
		toSerialize["acceptanceStatement"] = o.AcceptanceStatement
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	if o.AcceptanceStatuses != nil {
		toSerialize["acceptanceStatuses"] = o.AcceptanceStatuses
	}
	return json.Marshal(toSerialize)
}



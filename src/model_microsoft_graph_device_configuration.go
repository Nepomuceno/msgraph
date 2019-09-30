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
// MicrosoftGraphDeviceConfiguration struct for MicrosoftGraphDeviceConfiguration
type MicrosoftGraphDeviceConfiguration struct {
	Id *string `json:"id,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description *string `json:"description,omitempty"`
	isExplicitNullDescription bool `json:"-"`
	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// Version of the device configuration.
	Version *int32 `json:"version,omitempty"`

	Assignments *[]MicrosoftGraphDeviceConfigurationAssignment `json:"assignments,omitempty"`

	DeviceStatuses *[]MicrosoftGraphDeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	UserStatuses *[]MicrosoftGraphDeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	DeviceStatusOverview *AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`
	isExplicitNullDeviceStatusOverview bool `json:"-"`
	UserStatusOverview *AnyOfmicrosoftGraphDeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`
	isExplicitNullUserStatusOverview bool `json:"-"`
	DeviceSettingStateSummaries *[]MicrosoftGraphSettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDeviceConfiguration) SetId(v string) {
	o.Id = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphDeviceConfiguration) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphDeviceConfiguration) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MicrosoftGraphDeviceConfiguration) SetDescription(v string) {
	o.Description = &v
}

// SetDescriptionExplicitNull (un)sets Description to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Description value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfiguration) SetDescriptionExplicitNull(b bool) {
	o.Description = nil
	o.isExplicitNullDescription = b
}
// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphDeviceConfiguration) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetVersion returns the Version field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetVersionOk() (int32, bool) {
	if o == nil || o.Version == nil {
		var ret int32
		return ret, false
	}
	return *o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *MicrosoftGraphDeviceConfiguration) SetVersion(v int32) {
	o.Version = &v
}

// GetAssignments returns the Assignments field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetAssignments() []MicrosoftGraphDeviceConfigurationAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphDeviceConfigurationAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetAssignmentsOk() ([]MicrosoftGraphDeviceConfigurationAssignment, bool) {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphDeviceConfigurationAssignment
		return ret, false
	}
	return *o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphDeviceConfigurationAssignment and assigns it to the Assignments field.
func (o *MicrosoftGraphDeviceConfiguration) SetAssignments(v []MicrosoftGraphDeviceConfigurationAssignment) {
	o.Assignments = &v
}

// GetDeviceStatuses returns the DeviceStatuses field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceStatuses() []MicrosoftGraphDeviceConfigurationDeviceStatus {
	if o == nil || o.DeviceStatuses == nil {
		var ret []MicrosoftGraphDeviceConfigurationDeviceStatus
		return ret
	}
	return *o.DeviceStatuses
}

// GetDeviceStatusesOk returns a tuple with the DeviceStatuses field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceStatusesOk() ([]MicrosoftGraphDeviceConfigurationDeviceStatus, bool) {
	if o == nil || o.DeviceStatuses == nil {
		var ret []MicrosoftGraphDeviceConfigurationDeviceStatus
		return ret, false
	}
	return *o.DeviceStatuses, true
}

// HasDeviceStatuses returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDeviceStatuses() bool {
	if o != nil && o.DeviceStatuses != nil {
		return true
	}

	return false
}

// SetDeviceStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationDeviceStatus and assigns it to the DeviceStatuses field.
func (o *MicrosoftGraphDeviceConfiguration) SetDeviceStatuses(v []MicrosoftGraphDeviceConfigurationDeviceStatus) {
	o.DeviceStatuses = &v
}

// GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetUserStatuses() []MicrosoftGraphDeviceConfigurationUserStatus {
	if o == nil || o.UserStatuses == nil {
		var ret []MicrosoftGraphDeviceConfigurationUserStatus
		return ret
	}
	return *o.UserStatuses
}

// GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetUserStatusesOk() ([]MicrosoftGraphDeviceConfigurationUserStatus, bool) {
	if o == nil || o.UserStatuses == nil {
		var ret []MicrosoftGraphDeviceConfigurationUserStatus
		return ret, false
	}
	return *o.UserStatuses, true
}

// HasUserStatuses returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasUserStatuses() bool {
	if o != nil && o.UserStatuses != nil {
		return true
	}

	return false
}

// SetUserStatuses gets a reference to the given []MicrosoftGraphDeviceConfigurationUserStatus and assigns it to the UserStatuses field.
func (o *MicrosoftGraphDeviceConfiguration) SetUserStatuses(v []MicrosoftGraphDeviceConfigurationUserStatus) {
	o.UserStatuses = &v
}

// GetDeviceStatusOverview returns the DeviceStatusOverview field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview {
	if o == nil || o.DeviceStatusOverview == nil {
		var ret AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview
		return ret
	}
	return *o.DeviceStatusOverview
}

// GetDeviceStatusOverviewOk returns a tuple with the DeviceStatusOverview field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview, bool) {
	if o == nil || o.DeviceStatusOverview == nil {
		var ret AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview
		return ret, false
	}
	return *o.DeviceStatusOverview, true
}

// HasDeviceStatusOverview returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDeviceStatusOverview() bool {
	if o != nil && o.DeviceStatusOverview != nil {
		return true
	}

	return false
}

// SetDeviceStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview and assigns it to the DeviceStatusOverview field.
func (o *MicrosoftGraphDeviceConfiguration) SetDeviceStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationDeviceOverview) {
	o.DeviceStatusOverview = &v
}

// SetDeviceStatusOverviewExplicitNull (un)sets DeviceStatusOverview to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DeviceStatusOverview value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfiguration) SetDeviceStatusOverviewExplicitNull(b bool) {
	o.DeviceStatusOverview = nil
	o.isExplicitNullDeviceStatusOverview = b
}
// GetUserStatusOverview returns the UserStatusOverview field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetUserStatusOverview() AnyOfmicrosoftGraphDeviceConfigurationUserOverview {
	if o == nil || o.UserStatusOverview == nil {
		var ret AnyOfmicrosoftGraphDeviceConfigurationUserOverview
		return ret
	}
	return *o.UserStatusOverview
}

// GetUserStatusOverviewOk returns a tuple with the UserStatusOverview field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetUserStatusOverviewOk() (AnyOfmicrosoftGraphDeviceConfigurationUserOverview, bool) {
	if o == nil || o.UserStatusOverview == nil {
		var ret AnyOfmicrosoftGraphDeviceConfigurationUserOverview
		return ret, false
	}
	return *o.UserStatusOverview, true
}

// HasUserStatusOverview returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasUserStatusOverview() bool {
	if o != nil && o.UserStatusOverview != nil {
		return true
	}

	return false
}

// SetUserStatusOverview gets a reference to the given AnyOfmicrosoftGraphDeviceConfigurationUserOverview and assigns it to the UserStatusOverview field.
func (o *MicrosoftGraphDeviceConfiguration) SetUserStatusOverview(v AnyOfmicrosoftGraphDeviceConfigurationUserOverview) {
	o.UserStatusOverview = &v
}

// SetUserStatusOverviewExplicitNull (un)sets UserStatusOverview to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UserStatusOverview value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceConfiguration) SetUserStatusOverviewExplicitNull(b bool) {
	o.UserStatusOverview = nil
	o.isExplicitNullUserStatusOverview = b
}
// GetDeviceSettingStateSummaries returns the DeviceSettingStateSummaries field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceSettingStateSummaries() []MicrosoftGraphSettingStateDeviceSummary {
	if o == nil || o.DeviceSettingStateSummaries == nil {
		var ret []MicrosoftGraphSettingStateDeviceSummary
		return ret
	}
	return *o.DeviceSettingStateSummaries
}

// GetDeviceSettingStateSummariesOk returns a tuple with the DeviceSettingStateSummaries field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceConfiguration) GetDeviceSettingStateSummariesOk() ([]MicrosoftGraphSettingStateDeviceSummary, bool) {
	if o == nil || o.DeviceSettingStateSummaries == nil {
		var ret []MicrosoftGraphSettingStateDeviceSummary
		return ret, false
	}
	return *o.DeviceSettingStateSummaries, true
}

// HasDeviceSettingStateSummaries returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceConfiguration) HasDeviceSettingStateSummaries() bool {
	if o != nil && o.DeviceSettingStateSummaries != nil {
		return true
	}

	return false
}

// SetDeviceSettingStateSummaries gets a reference to the given []MicrosoftGraphSettingStateDeviceSummary and assigns it to the DeviceSettingStateSummaries field.
func (o *MicrosoftGraphDeviceConfiguration) SetDeviceSettingStateSummaries(v []MicrosoftGraphSettingStateDeviceSummary) {
	o.DeviceSettingStateSummaries = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDeviceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description == nil {
		if o.isExplicitNullDescription {
			toSerialize["description"] = o.Description
		}
	} else {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	if o.DeviceStatuses != nil {
		toSerialize["deviceStatuses"] = o.DeviceStatuses
	}
	if o.UserStatuses != nil {
		toSerialize["userStatuses"] = o.UserStatuses
	}
	if o.DeviceStatusOverview == nil {
		if o.isExplicitNullDeviceStatusOverview {
			toSerialize["deviceStatusOverview"] = o.DeviceStatusOverview
		}
	} else {
		toSerialize["deviceStatusOverview"] = o.DeviceStatusOverview
	}
	if o.UserStatusOverview == nil {
		if o.isExplicitNullUserStatusOverview {
			toSerialize["userStatusOverview"] = o.UserStatusOverview
		}
	} else {
		toSerialize["userStatusOverview"] = o.UserStatusOverview
	}
	if o.DeviceSettingStateSummaries != nil {
		toSerialize["deviceSettingStateSummaries"] = o.DeviceSettingStateSummaries
	}
	return json.Marshal(toSerialize)
}



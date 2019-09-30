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
// SoftwareUpdateStatusSummary struct for SoftwareUpdateStatusSummary
type SoftwareUpdateStatusSummary struct {
	// The name of the policy.
	DisplayName *string `json:"displayName,omitempty"`
	isExplicitNullDisplayName bool `json:"-"`
	// Number of compliant devices.
	CompliantDeviceCount *int32 `json:"compliantDeviceCount,omitempty"`

	// Number of non compliant devices.
	NonCompliantDeviceCount *int32 `json:"nonCompliantDeviceCount,omitempty"`

	// Number of remediated devices.
	RemediatedDeviceCount *int32 `json:"remediatedDeviceCount,omitempty"`

	// Number of devices had error.
	ErrorDeviceCount *int32 `json:"errorDeviceCount,omitempty"`

	// Number of unknown devices.
	UnknownDeviceCount *int32 `json:"unknownDeviceCount,omitempty"`

	// Number of conflict devices.
	ConflictDeviceCount *int32 `json:"conflictDeviceCount,omitempty"`

	// Number of not applicable devices.
	NotApplicableDeviceCount *int32 `json:"notApplicableDeviceCount,omitempty"`

	// Number of compliant users.
	CompliantUserCount *int32 `json:"compliantUserCount,omitempty"`

	// Number of non compliant users.
	NonCompliantUserCount *int32 `json:"nonCompliantUserCount,omitempty"`

	// Number of remediated users.
	RemediatedUserCount *int32 `json:"remediatedUserCount,omitempty"`

	// Number of users had error.
	ErrorUserCount *int32 `json:"errorUserCount,omitempty"`

	// Number of unknown users.
	UnknownUserCount *int32 `json:"unknownUserCount,omitempty"`

	// Number of conflict users.
	ConflictUserCount *int32 `json:"conflictUserCount,omitempty"`

	// Number of not applicable users.
	NotApplicableUserCount *int32 `json:"notApplicableUserCount,omitempty"`

}

// GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SoftwareUpdateStatusSummary) SetDisplayName(v string) {
	o.DisplayName = &v
}

// SetDisplayNameExplicitNull (un)sets DisplayName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DisplayName value is set to nil even if false is passed
func (o *SoftwareUpdateStatusSummary) SetDisplayNameExplicitNull(b bool) {
	o.DisplayName = nil
	o.isExplicitNullDisplayName = b
}
// GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetCompliantDeviceCount() int32 {
	if o == nil || o.CompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.CompliantDeviceCount
}

// GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetCompliantDeviceCountOk() (int32, bool) {
	if o == nil || o.CompliantDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CompliantDeviceCount, true
}

// HasCompliantDeviceCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasCompliantDeviceCount() bool {
	if o != nil && o.CompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.
func (o *SoftwareUpdateStatusSummary) SetCompliantDeviceCount(v int32) {
	o.CompliantDeviceCount = &v
}

// GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCount() int32 {
	if o == nil || o.NonCompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NonCompliantDeviceCount
}

// GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCountOk() (int32, bool) {
	if o == nil || o.NonCompliantDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.NonCompliantDeviceCount, true
}

// HasNonCompliantDeviceCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasNonCompliantDeviceCount() bool {
	if o != nil && o.NonCompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.
func (o *SoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(v int32) {
	o.NonCompliantDeviceCount = &v
}

// GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetRemediatedDeviceCount() int32 {
	if o == nil || o.RemediatedDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.RemediatedDeviceCount
}

// GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetRemediatedDeviceCountOk() (int32, bool) {
	if o == nil || o.RemediatedDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.RemediatedDeviceCount, true
}

// HasRemediatedDeviceCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasRemediatedDeviceCount() bool {
	if o != nil && o.RemediatedDeviceCount != nil {
		return true
	}

	return false
}

// SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.
func (o *SoftwareUpdateStatusSummary) SetRemediatedDeviceCount(v int32) {
	o.RemediatedDeviceCount = &v
}

// GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetErrorDeviceCount() int32 {
	if o == nil || o.ErrorDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorDeviceCount
}

// GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetErrorDeviceCountOk() (int32, bool) {
	if o == nil || o.ErrorDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ErrorDeviceCount, true
}

// HasErrorDeviceCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasErrorDeviceCount() bool {
	if o != nil && o.ErrorDeviceCount != nil {
		return true
	}

	return false
}

// SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.
func (o *SoftwareUpdateStatusSummary) SetErrorDeviceCount(v int32) {
	o.ErrorDeviceCount = &v
}

// GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetUnknownDeviceCount() int32 {
	if o == nil || o.UnknownDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.UnknownDeviceCount
}

// GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetUnknownDeviceCountOk() (int32, bool) {
	if o == nil || o.UnknownDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.UnknownDeviceCount, true
}

// HasUnknownDeviceCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasUnknownDeviceCount() bool {
	if o != nil && o.UnknownDeviceCount != nil {
		return true
	}

	return false
}

// SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.
func (o *SoftwareUpdateStatusSummary) SetUnknownDeviceCount(v int32) {
	o.UnknownDeviceCount = &v
}

// GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetConflictDeviceCount() int32 {
	if o == nil || o.ConflictDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ConflictDeviceCount
}

// GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetConflictDeviceCountOk() (int32, bool) {
	if o == nil || o.ConflictDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ConflictDeviceCount, true
}

// HasConflictDeviceCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasConflictDeviceCount() bool {
	if o != nil && o.ConflictDeviceCount != nil {
		return true
	}

	return false
}

// SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.
func (o *SoftwareUpdateStatusSummary) SetConflictDeviceCount(v int32) {
	o.ConflictDeviceCount = &v
}

// GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCount() int32 {
	if o == nil || o.NotApplicableDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableDeviceCount
}

// GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCountOk() (int32, bool) {
	if o == nil || o.NotApplicableDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.NotApplicableDeviceCount, true
}

// HasNotApplicableDeviceCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasNotApplicableDeviceCount() bool {
	if o != nil && o.NotApplicableDeviceCount != nil {
		return true
	}

	return false
}

// SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.
func (o *SoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(v int32) {
	o.NotApplicableDeviceCount = &v
}

// GetCompliantUserCount returns the CompliantUserCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetCompliantUserCount() int32 {
	if o == nil || o.CompliantUserCount == nil {
		var ret int32
		return ret
	}
	return *o.CompliantUserCount
}

// GetCompliantUserCountOk returns a tuple with the CompliantUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetCompliantUserCountOk() (int32, bool) {
	if o == nil || o.CompliantUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CompliantUserCount, true
}

// HasCompliantUserCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasCompliantUserCount() bool {
	if o != nil && o.CompliantUserCount != nil {
		return true
	}

	return false
}

// SetCompliantUserCount gets a reference to the given int32 and assigns it to the CompliantUserCount field.
func (o *SoftwareUpdateStatusSummary) SetCompliantUserCount(v int32) {
	o.CompliantUserCount = &v
}

// GetNonCompliantUserCount returns the NonCompliantUserCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetNonCompliantUserCount() int32 {
	if o == nil || o.NonCompliantUserCount == nil {
		var ret int32
		return ret
	}
	return *o.NonCompliantUserCount
}

// GetNonCompliantUserCountOk returns a tuple with the NonCompliantUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetNonCompliantUserCountOk() (int32, bool) {
	if o == nil || o.NonCompliantUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.NonCompliantUserCount, true
}

// HasNonCompliantUserCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasNonCompliantUserCount() bool {
	if o != nil && o.NonCompliantUserCount != nil {
		return true
	}

	return false
}

// SetNonCompliantUserCount gets a reference to the given int32 and assigns it to the NonCompliantUserCount field.
func (o *SoftwareUpdateStatusSummary) SetNonCompliantUserCount(v int32) {
	o.NonCompliantUserCount = &v
}

// GetRemediatedUserCount returns the RemediatedUserCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetRemediatedUserCount() int32 {
	if o == nil || o.RemediatedUserCount == nil {
		var ret int32
		return ret
	}
	return *o.RemediatedUserCount
}

// GetRemediatedUserCountOk returns a tuple with the RemediatedUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetRemediatedUserCountOk() (int32, bool) {
	if o == nil || o.RemediatedUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.RemediatedUserCount, true
}

// HasRemediatedUserCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasRemediatedUserCount() bool {
	if o != nil && o.RemediatedUserCount != nil {
		return true
	}

	return false
}

// SetRemediatedUserCount gets a reference to the given int32 and assigns it to the RemediatedUserCount field.
func (o *SoftwareUpdateStatusSummary) SetRemediatedUserCount(v int32) {
	o.RemediatedUserCount = &v
}

// GetErrorUserCount returns the ErrorUserCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetErrorUserCount() int32 {
	if o == nil || o.ErrorUserCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorUserCount
}

// GetErrorUserCountOk returns a tuple with the ErrorUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetErrorUserCountOk() (int32, bool) {
	if o == nil || o.ErrorUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ErrorUserCount, true
}

// HasErrorUserCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasErrorUserCount() bool {
	if o != nil && o.ErrorUserCount != nil {
		return true
	}

	return false
}

// SetErrorUserCount gets a reference to the given int32 and assigns it to the ErrorUserCount field.
func (o *SoftwareUpdateStatusSummary) SetErrorUserCount(v int32) {
	o.ErrorUserCount = &v
}

// GetUnknownUserCount returns the UnknownUserCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetUnknownUserCount() int32 {
	if o == nil || o.UnknownUserCount == nil {
		var ret int32
		return ret
	}
	return *o.UnknownUserCount
}

// GetUnknownUserCountOk returns a tuple with the UnknownUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetUnknownUserCountOk() (int32, bool) {
	if o == nil || o.UnknownUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.UnknownUserCount, true
}

// HasUnknownUserCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasUnknownUserCount() bool {
	if o != nil && o.UnknownUserCount != nil {
		return true
	}

	return false
}

// SetUnknownUserCount gets a reference to the given int32 and assigns it to the UnknownUserCount field.
func (o *SoftwareUpdateStatusSummary) SetUnknownUserCount(v int32) {
	o.UnknownUserCount = &v
}

// GetConflictUserCount returns the ConflictUserCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetConflictUserCount() int32 {
	if o == nil || o.ConflictUserCount == nil {
		var ret int32
		return ret
	}
	return *o.ConflictUserCount
}

// GetConflictUserCountOk returns a tuple with the ConflictUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetConflictUserCountOk() (int32, bool) {
	if o == nil || o.ConflictUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ConflictUserCount, true
}

// HasConflictUserCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasConflictUserCount() bool {
	if o != nil && o.ConflictUserCount != nil {
		return true
	}

	return false
}

// SetConflictUserCount gets a reference to the given int32 and assigns it to the ConflictUserCount field.
func (o *SoftwareUpdateStatusSummary) SetConflictUserCount(v int32) {
	o.ConflictUserCount = &v
}

// GetNotApplicableUserCount returns the NotApplicableUserCount field if non-nil, zero value otherwise.
func (o *SoftwareUpdateStatusSummary) GetNotApplicableUserCount() int32 {
	if o == nil || o.NotApplicableUserCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableUserCount
}

// GetNotApplicableUserCountOk returns a tuple with the NotApplicableUserCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUpdateStatusSummary) GetNotApplicableUserCountOk() (int32, bool) {
	if o == nil || o.NotApplicableUserCount == nil {
		var ret int32
		return ret, false
	}
	return *o.NotApplicableUserCount, true
}

// HasNotApplicableUserCount returns a boolean if a field has been set.
func (o *SoftwareUpdateStatusSummary) HasNotApplicableUserCount() bool {
	if o != nil && o.NotApplicableUserCount != nil {
		return true
	}

	return false
}

// SetNotApplicableUserCount gets a reference to the given int32 and assigns it to the NotApplicableUserCount field.
func (o *SoftwareUpdateStatusSummary) SetNotApplicableUserCount(v int32) {
	o.NotApplicableUserCount = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o SoftwareUpdateStatusSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName == nil {
		if o.isExplicitNullDisplayName {
			toSerialize["displayName"] = o.DisplayName
		}
	} else {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.CompliantDeviceCount != nil {
		toSerialize["compliantDeviceCount"] = o.CompliantDeviceCount
	}
	if o.NonCompliantDeviceCount != nil {
		toSerialize["nonCompliantDeviceCount"] = o.NonCompliantDeviceCount
	}
	if o.RemediatedDeviceCount != nil {
		toSerialize["remediatedDeviceCount"] = o.RemediatedDeviceCount
	}
	if o.ErrorDeviceCount != nil {
		toSerialize["errorDeviceCount"] = o.ErrorDeviceCount
	}
	if o.UnknownDeviceCount != nil {
		toSerialize["unknownDeviceCount"] = o.UnknownDeviceCount
	}
	if o.ConflictDeviceCount != nil {
		toSerialize["conflictDeviceCount"] = o.ConflictDeviceCount
	}
	if o.NotApplicableDeviceCount != nil {
		toSerialize["notApplicableDeviceCount"] = o.NotApplicableDeviceCount
	}
	if o.CompliantUserCount != nil {
		toSerialize["compliantUserCount"] = o.CompliantUserCount
	}
	if o.NonCompliantUserCount != nil {
		toSerialize["nonCompliantUserCount"] = o.NonCompliantUserCount
	}
	if o.RemediatedUserCount != nil {
		toSerialize["remediatedUserCount"] = o.RemediatedUserCount
	}
	if o.ErrorUserCount != nil {
		toSerialize["errorUserCount"] = o.ErrorUserCount
	}
	if o.UnknownUserCount != nil {
		toSerialize["unknownUserCount"] = o.UnknownUserCount
	}
	if o.ConflictUserCount != nil {
		toSerialize["conflictUserCount"] = o.ConflictUserCount
	}
	if o.NotApplicableUserCount != nil {
		toSerialize["notApplicableUserCount"] = o.NotApplicableUserCount
	}
	return json.Marshal(toSerialize)
}


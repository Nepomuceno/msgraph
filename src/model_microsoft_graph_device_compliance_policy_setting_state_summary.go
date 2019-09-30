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
// MicrosoftGraphDeviceCompliancePolicySettingStateSummary struct for MicrosoftGraphDeviceCompliancePolicySettingStateSummary
type MicrosoftGraphDeviceCompliancePolicySettingStateSummary struct {
	Id *string `json:"id,omitempty"`

	// The setting class name and property name.
	Setting *string `json:"setting,omitempty"`
	isExplicitNullSetting bool `json:"-"`
	// Name of the setting.
	SettingName *string `json:"settingName,omitempty"`
	isExplicitNullSettingName bool `json:"-"`
	// Setting platform
	PlatformType *AnyOfmicrosoftGraphPolicyPlatformType `json:"platformType,omitempty"`

	// Number of unknown devices
	UnknownDeviceCount *int32 `json:"unknownDeviceCount,omitempty"`

	// Number of not applicable devices
	NotApplicableDeviceCount *int32 `json:"notApplicableDeviceCount,omitempty"`

	// Number of compliant devices
	CompliantDeviceCount *int32 `json:"compliantDeviceCount,omitempty"`

	// Number of remediated devices
	RemediatedDeviceCount *int32 `json:"remediatedDeviceCount,omitempty"`

	// Number of NonCompliant devices
	NonCompliantDeviceCount *int32 `json:"nonCompliantDeviceCount,omitempty"`

	// Number of error devices
	ErrorDeviceCount *int32 `json:"errorDeviceCount,omitempty"`

	// Number of conflict devices
	ConflictDeviceCount *int32 `json:"conflictDeviceCount,omitempty"`

	DeviceComplianceSettingStates *[]MicrosoftGraphDeviceComplianceSettingState `json:"deviceComplianceSettingStates,omitempty"`

}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetId(v string) {
	o.Id = &v
}

// GetSetting returns the Setting field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetSetting() string {
	if o == nil || o.Setting == nil {
		var ret string
		return ret
	}
	return *o.Setting
}

// GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetSettingOk() (string, bool) {
	if o == nil || o.Setting == nil {
		var ret string
		return ret, false
	}
	return *o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasSetting() bool {
	if o != nil && o.Setting != nil {
		return true
	}

	return false
}

// SetSetting gets a reference to the given string and assigns it to the Setting field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetSetting(v string) {
	o.Setting = &v
}

// SetSettingExplicitNull (un)sets Setting to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Setting value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetSettingExplicitNull(b bool) {
	o.Setting = nil
	o.isExplicitNullSetting = b
}
// GetSettingName returns the SettingName field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetSettingName() string {
	if o == nil || o.SettingName == nil {
		var ret string
		return ret
	}
	return *o.SettingName
}

// GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetSettingNameOk() (string, bool) {
	if o == nil || o.SettingName == nil {
		var ret string
		return ret, false
	}
	return *o.SettingName, true
}

// HasSettingName returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasSettingName() bool {
	if o != nil && o.SettingName != nil {
		return true
	}

	return false
}

// SetSettingName gets a reference to the given string and assigns it to the SettingName field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetSettingName(v string) {
	o.SettingName = &v
}

// SetSettingNameExplicitNull (un)sets SettingName to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The SettingName value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetSettingNameExplicitNull(b bool) {
	o.SettingName = nil
	o.isExplicitNullSettingName = b
}
// GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetPlatformType() AnyOfmicrosoftGraphPolicyPlatformType {
	if o == nil || o.PlatformType == nil {
		var ret AnyOfmicrosoftGraphPolicyPlatformType
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetPlatformTypeOk() (AnyOfmicrosoftGraphPolicyPlatformType, bool) {
	if o == nil || o.PlatformType == nil {
		var ret AnyOfmicrosoftGraphPolicyPlatformType
		return ret, false
	}
	return *o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given AnyOfmicrosoftGraphPolicyPlatformType and assigns it to the PlatformType field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetPlatformType(v AnyOfmicrosoftGraphPolicyPlatformType) {
	o.PlatformType = &v
}

// GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetUnknownDeviceCount() int32 {
	if o == nil || o.UnknownDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.UnknownDeviceCount
}

// GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetUnknownDeviceCountOk() (int32, bool) {
	if o == nil || o.UnknownDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.UnknownDeviceCount, true
}

// HasUnknownDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasUnknownDeviceCount() bool {
	if o != nil && o.UnknownDeviceCount != nil {
		return true
	}

	return false
}

// SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetUnknownDeviceCount(v int32) {
	o.UnknownDeviceCount = &v
}

// GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetNotApplicableDeviceCount() int32 {
	if o == nil || o.NotApplicableDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableDeviceCount
}

// GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetNotApplicableDeviceCountOk() (int32, bool) {
	if o == nil || o.NotApplicableDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.NotApplicableDeviceCount, true
}

// HasNotApplicableDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasNotApplicableDeviceCount() bool {
	if o != nil && o.NotApplicableDeviceCount != nil {
		return true
	}

	return false
}

// SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetNotApplicableDeviceCount(v int32) {
	o.NotApplicableDeviceCount = &v
}

// GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetCompliantDeviceCount() int32 {
	if o == nil || o.CompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.CompliantDeviceCount
}

// GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetCompliantDeviceCountOk() (int32, bool) {
	if o == nil || o.CompliantDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CompliantDeviceCount, true
}

// HasCompliantDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasCompliantDeviceCount() bool {
	if o != nil && o.CompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetCompliantDeviceCount(v int32) {
	o.CompliantDeviceCount = &v
}

// GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetRemediatedDeviceCount() int32 {
	if o == nil || o.RemediatedDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.RemediatedDeviceCount
}

// GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetRemediatedDeviceCountOk() (int32, bool) {
	if o == nil || o.RemediatedDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.RemediatedDeviceCount, true
}

// HasRemediatedDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasRemediatedDeviceCount() bool {
	if o != nil && o.RemediatedDeviceCount != nil {
		return true
	}

	return false
}

// SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetRemediatedDeviceCount(v int32) {
	o.RemediatedDeviceCount = &v
}

// GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetNonCompliantDeviceCount() int32 {
	if o == nil || o.NonCompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NonCompliantDeviceCount
}

// GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetNonCompliantDeviceCountOk() (int32, bool) {
	if o == nil || o.NonCompliantDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.NonCompliantDeviceCount, true
}

// HasNonCompliantDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasNonCompliantDeviceCount() bool {
	if o != nil && o.NonCompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetNonCompliantDeviceCount(v int32) {
	o.NonCompliantDeviceCount = &v
}

// GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetErrorDeviceCount() int32 {
	if o == nil || o.ErrorDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorDeviceCount
}

// GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetErrorDeviceCountOk() (int32, bool) {
	if o == nil || o.ErrorDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ErrorDeviceCount, true
}

// HasErrorDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasErrorDeviceCount() bool {
	if o != nil && o.ErrorDeviceCount != nil {
		return true
	}

	return false
}

// SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetErrorDeviceCount(v int32) {
	o.ErrorDeviceCount = &v
}

// GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetConflictDeviceCount() int32 {
	if o == nil || o.ConflictDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ConflictDeviceCount
}

// GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetConflictDeviceCountOk() (int32, bool) {
	if o == nil || o.ConflictDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ConflictDeviceCount, true
}

// HasConflictDeviceCount returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasConflictDeviceCount() bool {
	if o != nil && o.ConflictDeviceCount != nil {
		return true
	}

	return false
}

// SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetConflictDeviceCount(v int32) {
	o.ConflictDeviceCount = &v
}

// GetDeviceComplianceSettingStates returns the DeviceComplianceSettingStates field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetDeviceComplianceSettingStates() []MicrosoftGraphDeviceComplianceSettingState {
	if o == nil || o.DeviceComplianceSettingStates == nil {
		var ret []MicrosoftGraphDeviceComplianceSettingState
		return ret
	}
	return *o.DeviceComplianceSettingStates
}

// GetDeviceComplianceSettingStatesOk returns a tuple with the DeviceComplianceSettingStates field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) GetDeviceComplianceSettingStatesOk() ([]MicrosoftGraphDeviceComplianceSettingState, bool) {
	if o == nil || o.DeviceComplianceSettingStates == nil {
		var ret []MicrosoftGraphDeviceComplianceSettingState
		return ret, false
	}
	return *o.DeviceComplianceSettingStates, true
}

// HasDeviceComplianceSettingStates returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) HasDeviceComplianceSettingStates() bool {
	if o != nil && o.DeviceComplianceSettingStates != nil {
		return true
	}

	return false
}

// SetDeviceComplianceSettingStates gets a reference to the given []MicrosoftGraphDeviceComplianceSettingState and assigns it to the DeviceComplianceSettingStates field.
func (o *MicrosoftGraphDeviceCompliancePolicySettingStateSummary) SetDeviceComplianceSettingStates(v []MicrosoftGraphDeviceComplianceSettingState) {
	o.DeviceComplianceSettingStates = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDeviceCompliancePolicySettingStateSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Setting == nil {
		if o.isExplicitNullSetting {
			toSerialize["setting"] = o.Setting
		}
	} else {
		toSerialize["setting"] = o.Setting
	}
	if o.SettingName == nil {
		if o.isExplicitNullSettingName {
			toSerialize["settingName"] = o.SettingName
		}
	} else {
		toSerialize["settingName"] = o.SettingName
	}
	if o.PlatformType != nil {
		toSerialize["platformType"] = o.PlatformType
	}
	if o.UnknownDeviceCount != nil {
		toSerialize["unknownDeviceCount"] = o.UnknownDeviceCount
	}
	if o.NotApplicableDeviceCount != nil {
		toSerialize["notApplicableDeviceCount"] = o.NotApplicableDeviceCount
	}
	if o.CompliantDeviceCount != nil {
		toSerialize["compliantDeviceCount"] = o.CompliantDeviceCount
	}
	if o.RemediatedDeviceCount != nil {
		toSerialize["remediatedDeviceCount"] = o.RemediatedDeviceCount
	}
	if o.NonCompliantDeviceCount != nil {
		toSerialize["nonCompliantDeviceCount"] = o.NonCompliantDeviceCount
	}
	if o.ErrorDeviceCount != nil {
		toSerialize["errorDeviceCount"] = o.ErrorDeviceCount
	}
	if o.ConflictDeviceCount != nil {
		toSerialize["conflictDeviceCount"] = o.ConflictDeviceCount
	}
	if o.DeviceComplianceSettingStates != nil {
		toSerialize["deviceComplianceSettingStates"] = o.DeviceComplianceSettingStates
	}
	return json.Marshal(toSerialize)
}



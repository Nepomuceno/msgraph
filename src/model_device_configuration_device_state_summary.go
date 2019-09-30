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
// DeviceConfigurationDeviceStateSummary struct for DeviceConfigurationDeviceStateSummary
type DeviceConfigurationDeviceStateSummary struct {
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

}

// GetUnknownDeviceCount returns the UnknownDeviceCount field if non-nil, zero value otherwise.
func (o *DeviceConfigurationDeviceStateSummary) GetUnknownDeviceCount() int32 {
	if o == nil || o.UnknownDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.UnknownDeviceCount
}

// GetUnknownDeviceCountOk returns a tuple with the UnknownDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationDeviceStateSummary) GetUnknownDeviceCountOk() (int32, bool) {
	if o == nil || o.UnknownDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.UnknownDeviceCount, true
}

// HasUnknownDeviceCount returns a boolean if a field has been set.
func (o *DeviceConfigurationDeviceStateSummary) HasUnknownDeviceCount() bool {
	if o != nil && o.UnknownDeviceCount != nil {
		return true
	}

	return false
}

// SetUnknownDeviceCount gets a reference to the given int32 and assigns it to the UnknownDeviceCount field.
func (o *DeviceConfigurationDeviceStateSummary) SetUnknownDeviceCount(v int32) {
	o.UnknownDeviceCount = &v
}

// GetNotApplicableDeviceCount returns the NotApplicableDeviceCount field if non-nil, zero value otherwise.
func (o *DeviceConfigurationDeviceStateSummary) GetNotApplicableDeviceCount() int32 {
	if o == nil || o.NotApplicableDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableDeviceCount
}

// GetNotApplicableDeviceCountOk returns a tuple with the NotApplicableDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationDeviceStateSummary) GetNotApplicableDeviceCountOk() (int32, bool) {
	if o == nil || o.NotApplicableDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.NotApplicableDeviceCount, true
}

// HasNotApplicableDeviceCount returns a boolean if a field has been set.
func (o *DeviceConfigurationDeviceStateSummary) HasNotApplicableDeviceCount() bool {
	if o != nil && o.NotApplicableDeviceCount != nil {
		return true
	}

	return false
}

// SetNotApplicableDeviceCount gets a reference to the given int32 and assigns it to the NotApplicableDeviceCount field.
func (o *DeviceConfigurationDeviceStateSummary) SetNotApplicableDeviceCount(v int32) {
	o.NotApplicableDeviceCount = &v
}

// GetCompliantDeviceCount returns the CompliantDeviceCount field if non-nil, zero value otherwise.
func (o *DeviceConfigurationDeviceStateSummary) GetCompliantDeviceCount() int32 {
	if o == nil || o.CompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.CompliantDeviceCount
}

// GetCompliantDeviceCountOk returns a tuple with the CompliantDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationDeviceStateSummary) GetCompliantDeviceCountOk() (int32, bool) {
	if o == nil || o.CompliantDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.CompliantDeviceCount, true
}

// HasCompliantDeviceCount returns a boolean if a field has been set.
func (o *DeviceConfigurationDeviceStateSummary) HasCompliantDeviceCount() bool {
	if o != nil && o.CompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetCompliantDeviceCount gets a reference to the given int32 and assigns it to the CompliantDeviceCount field.
func (o *DeviceConfigurationDeviceStateSummary) SetCompliantDeviceCount(v int32) {
	o.CompliantDeviceCount = &v
}

// GetRemediatedDeviceCount returns the RemediatedDeviceCount field if non-nil, zero value otherwise.
func (o *DeviceConfigurationDeviceStateSummary) GetRemediatedDeviceCount() int32 {
	if o == nil || o.RemediatedDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.RemediatedDeviceCount
}

// GetRemediatedDeviceCountOk returns a tuple with the RemediatedDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationDeviceStateSummary) GetRemediatedDeviceCountOk() (int32, bool) {
	if o == nil || o.RemediatedDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.RemediatedDeviceCount, true
}

// HasRemediatedDeviceCount returns a boolean if a field has been set.
func (o *DeviceConfigurationDeviceStateSummary) HasRemediatedDeviceCount() bool {
	if o != nil && o.RemediatedDeviceCount != nil {
		return true
	}

	return false
}

// SetRemediatedDeviceCount gets a reference to the given int32 and assigns it to the RemediatedDeviceCount field.
func (o *DeviceConfigurationDeviceStateSummary) SetRemediatedDeviceCount(v int32) {
	o.RemediatedDeviceCount = &v
}

// GetNonCompliantDeviceCount returns the NonCompliantDeviceCount field if non-nil, zero value otherwise.
func (o *DeviceConfigurationDeviceStateSummary) GetNonCompliantDeviceCount() int32 {
	if o == nil || o.NonCompliantDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.NonCompliantDeviceCount
}

// GetNonCompliantDeviceCountOk returns a tuple with the NonCompliantDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationDeviceStateSummary) GetNonCompliantDeviceCountOk() (int32, bool) {
	if o == nil || o.NonCompliantDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.NonCompliantDeviceCount, true
}

// HasNonCompliantDeviceCount returns a boolean if a field has been set.
func (o *DeviceConfigurationDeviceStateSummary) HasNonCompliantDeviceCount() bool {
	if o != nil && o.NonCompliantDeviceCount != nil {
		return true
	}

	return false
}

// SetNonCompliantDeviceCount gets a reference to the given int32 and assigns it to the NonCompliantDeviceCount field.
func (o *DeviceConfigurationDeviceStateSummary) SetNonCompliantDeviceCount(v int32) {
	o.NonCompliantDeviceCount = &v
}

// GetErrorDeviceCount returns the ErrorDeviceCount field if non-nil, zero value otherwise.
func (o *DeviceConfigurationDeviceStateSummary) GetErrorDeviceCount() int32 {
	if o == nil || o.ErrorDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorDeviceCount
}

// GetErrorDeviceCountOk returns a tuple with the ErrorDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationDeviceStateSummary) GetErrorDeviceCountOk() (int32, bool) {
	if o == nil || o.ErrorDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ErrorDeviceCount, true
}

// HasErrorDeviceCount returns a boolean if a field has been set.
func (o *DeviceConfigurationDeviceStateSummary) HasErrorDeviceCount() bool {
	if o != nil && o.ErrorDeviceCount != nil {
		return true
	}

	return false
}

// SetErrorDeviceCount gets a reference to the given int32 and assigns it to the ErrorDeviceCount field.
func (o *DeviceConfigurationDeviceStateSummary) SetErrorDeviceCount(v int32) {
	o.ErrorDeviceCount = &v
}

// GetConflictDeviceCount returns the ConflictDeviceCount field if non-nil, zero value otherwise.
func (o *DeviceConfigurationDeviceStateSummary) GetConflictDeviceCount() int32 {
	if o == nil || o.ConflictDeviceCount == nil {
		var ret int32
		return ret
	}
	return *o.ConflictDeviceCount
}

// GetConflictDeviceCountOk returns a tuple with the ConflictDeviceCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfigurationDeviceStateSummary) GetConflictDeviceCountOk() (int32, bool) {
	if o == nil || o.ConflictDeviceCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ConflictDeviceCount, true
}

// HasConflictDeviceCount returns a boolean if a field has been set.
func (o *DeviceConfigurationDeviceStateSummary) HasConflictDeviceCount() bool {
	if o != nil && o.ConflictDeviceCount != nil {
		return true
	}

	return false
}

// SetConflictDeviceCount gets a reference to the given int32 and assigns it to the ConflictDeviceCount field.
func (o *DeviceConfigurationDeviceStateSummary) SetConflictDeviceCount(v int32) {
	o.ConflictDeviceCount = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o DeviceConfigurationDeviceStateSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}


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
// ManagedDeviceMobileAppConfigurationDeviceSummary Contains properties, inherited properties and actions for an MDM mobile app configuration device status summary.
type ManagedDeviceMobileAppConfigurationDeviceSummary struct {
	// Number of pending devices
	PendingCount *int32 `json:"pendingCount,omitempty"`

	// Number of not applicable devices
	NotApplicableCount *int32 `json:"notApplicableCount,omitempty"`

	// Number of succeeded devices
	SuccessCount *int32 `json:"successCount,omitempty"`

	// Number of error devices
	ErrorCount *int32 `json:"errorCount,omitempty"`

	// Number of failed devices
	FailedCount *int32 `json:"failedCount,omitempty"`

	// Last update time
	LastUpdateDateTime *time.Time `json:"lastUpdateDateTime,omitempty"`

	// Version of the policy for that overview
	ConfigurationVersion *int32 `json:"configurationVersion,omitempty"`

}

// GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetPendingCount() int32 {
	if o == nil || o.PendingCount == nil {
		var ret int32
		return ret
	}
	return *o.PendingCount
}

// GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetPendingCountOk() (int32, bool) {
	if o == nil || o.PendingCount == nil {
		var ret int32
		return ret, false
	}
	return *o.PendingCount, true
}

// HasPendingCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) HasPendingCount() bool {
	if o != nil && o.PendingCount != nil {
		return true
	}

	return false
}

// SetPendingCount gets a reference to the given int32 and assigns it to the PendingCount field.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) SetPendingCount(v int32) {
	o.PendingCount = &v
}

// GetNotApplicableCount returns the NotApplicableCount field if non-nil, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetNotApplicableCount() int32 {
	if o == nil || o.NotApplicableCount == nil {
		var ret int32
		return ret
	}
	return *o.NotApplicableCount
}

// GetNotApplicableCountOk returns a tuple with the NotApplicableCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetNotApplicableCountOk() (int32, bool) {
	if o == nil || o.NotApplicableCount == nil {
		var ret int32
		return ret, false
	}
	return *o.NotApplicableCount, true
}

// HasNotApplicableCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) HasNotApplicableCount() bool {
	if o != nil && o.NotApplicableCount != nil {
		return true
	}

	return false
}

// SetNotApplicableCount gets a reference to the given int32 and assigns it to the NotApplicableCount field.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) SetNotApplicableCount(v int32) {
	o.NotApplicableCount = &v
}

// GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetSuccessCount() int32 {
	if o == nil || o.SuccessCount == nil {
		var ret int32
		return ret
	}
	return *o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetSuccessCountOk() (int32, bool) {
	if o == nil || o.SuccessCount == nil {
		var ret int32
		return ret, false
	}
	return *o.SuccessCount, true
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) HasSuccessCount() bool {
	if o != nil && o.SuccessCount != nil {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given int32 and assigns it to the SuccessCount field.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) SetSuccessCount(v int32) {
	o.SuccessCount = &v
}

// GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetErrorCount() int32 {
	if o == nil || o.ErrorCount == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetErrorCountOk() (int32, bool) {
	if o == nil || o.ErrorCount == nil {
		var ret int32
		return ret, false
	}
	return *o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) HasErrorCount() bool {
	if o != nil && o.ErrorCount != nil {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) SetErrorCount(v int32) {
	o.ErrorCount = &v
}

// GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetFailedCount() int32 {
	if o == nil || o.FailedCount == nil {
		var ret int32
		return ret
	}
	return *o.FailedCount
}

// GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetFailedCountOk() (int32, bool) {
	if o == nil || o.FailedCount == nil {
		var ret int32
		return ret, false
	}
	return *o.FailedCount, true
}

// HasFailedCount returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) HasFailedCount() bool {
	if o != nil && o.FailedCount != nil {
		return true
	}

	return false
}

// SetFailedCount gets a reference to the given int32 and assigns it to the FailedCount field.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) SetFailedCount(v int32) {
	o.FailedCount = &v
}

// GetLastUpdateDateTime returns the LastUpdateDateTime field if non-nil, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetLastUpdateDateTime() time.Time {
	if o == nil || o.LastUpdateDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateDateTime
}

// GetLastUpdateDateTimeOk returns a tuple with the LastUpdateDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetLastUpdateDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastUpdateDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastUpdateDateTime, true
}

// HasLastUpdateDateTime returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) HasLastUpdateDateTime() bool {
	if o != nil && o.LastUpdateDateTime != nil {
		return true
	}

	return false
}

// SetLastUpdateDateTime gets a reference to the given time.Time and assigns it to the LastUpdateDateTime field.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) SetLastUpdateDateTime(v time.Time) {
	o.LastUpdateDateTime = &v
}

// GetConfigurationVersion returns the ConfigurationVersion field if non-nil, zero value otherwise.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetConfigurationVersion() int32 {
	if o == nil || o.ConfigurationVersion == nil {
		var ret int32
		return ret
	}
	return *o.ConfigurationVersion
}

// GetConfigurationVersionOk returns a tuple with the ConfigurationVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) GetConfigurationVersionOk() (int32, bool) {
	if o == nil || o.ConfigurationVersion == nil {
		var ret int32
		return ret, false
	}
	return *o.ConfigurationVersion, true
}

// HasConfigurationVersion returns a boolean if a field has been set.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) HasConfigurationVersion() bool {
	if o != nil && o.ConfigurationVersion != nil {
		return true
	}

	return false
}

// SetConfigurationVersion gets a reference to the given int32 and assigns it to the ConfigurationVersion field.
func (o *ManagedDeviceMobileAppConfigurationDeviceSummary) SetConfigurationVersion(v int32) {
	o.ConfigurationVersion = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o ManagedDeviceMobileAppConfigurationDeviceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PendingCount != nil {
		toSerialize["pendingCount"] = o.PendingCount
	}
	if o.NotApplicableCount != nil {
		toSerialize["notApplicableCount"] = o.NotApplicableCount
	}
	if o.SuccessCount != nil {
		toSerialize["successCount"] = o.SuccessCount
	}
	if o.ErrorCount != nil {
		toSerialize["errorCount"] = o.ErrorCount
	}
	if o.FailedCount != nil {
		toSerialize["failedCount"] = o.FailedCount
	}
	if o.LastUpdateDateTime != nil {
		toSerialize["lastUpdateDateTime"] = o.LastUpdateDateTime
	}
	if o.ConfigurationVersion != nil {
		toSerialize["configurationVersion"] = o.ConfigurationVersion
	}
	return json.Marshal(toSerialize)
}



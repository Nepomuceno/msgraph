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
// WindowsUpdateScheduledInstall struct for WindowsUpdateScheduledInstall
type WindowsUpdateScheduledInstall struct {
	// Scheduled Install Day in week
	ScheduledInstallDay *AnyOfmicrosoftGraphWeeklySchedule `json:"scheduledInstallDay,omitempty"`

	// Scheduled Install Time during day
	ScheduledInstallTime *string `json:"scheduledInstallTime,omitempty"`

}

// GetScheduledInstallDay returns the ScheduledInstallDay field if non-nil, zero value otherwise.
func (o *WindowsUpdateScheduledInstall) GetScheduledInstallDay() AnyOfmicrosoftGraphWeeklySchedule {
	if o == nil || o.ScheduledInstallDay == nil {
		var ret AnyOfmicrosoftGraphWeeklySchedule
		return ret
	}
	return *o.ScheduledInstallDay
}

// GetScheduledInstallDayOk returns a tuple with the ScheduledInstallDay field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsUpdateScheduledInstall) GetScheduledInstallDayOk() (AnyOfmicrosoftGraphWeeklySchedule, bool) {
	if o == nil || o.ScheduledInstallDay == nil {
		var ret AnyOfmicrosoftGraphWeeklySchedule
		return ret, false
	}
	return *o.ScheduledInstallDay, true
}

// HasScheduledInstallDay returns a boolean if a field has been set.
func (o *WindowsUpdateScheduledInstall) HasScheduledInstallDay() bool {
	if o != nil && o.ScheduledInstallDay != nil {
		return true
	}

	return false
}

// SetScheduledInstallDay gets a reference to the given AnyOfmicrosoftGraphWeeklySchedule and assigns it to the ScheduledInstallDay field.
func (o *WindowsUpdateScheduledInstall) SetScheduledInstallDay(v AnyOfmicrosoftGraphWeeklySchedule) {
	o.ScheduledInstallDay = &v
}

// GetScheduledInstallTime returns the ScheduledInstallTime field if non-nil, zero value otherwise.
func (o *WindowsUpdateScheduledInstall) GetScheduledInstallTime() string {
	if o == nil || o.ScheduledInstallTime == nil {
		var ret string
		return ret
	}
	return *o.ScheduledInstallTime
}

// GetScheduledInstallTimeOk returns a tuple with the ScheduledInstallTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WindowsUpdateScheduledInstall) GetScheduledInstallTimeOk() (string, bool) {
	if o == nil || o.ScheduledInstallTime == nil {
		var ret string
		return ret, false
	}
	return *o.ScheduledInstallTime, true
}

// HasScheduledInstallTime returns a boolean if a field has been set.
func (o *WindowsUpdateScheduledInstall) HasScheduledInstallTime() bool {
	if o != nil && o.ScheduledInstallTime != nil {
		return true
	}

	return false
}

// SetScheduledInstallTime gets a reference to the given string and assigns it to the ScheduledInstallTime field.
func (o *WindowsUpdateScheduledInstall) SetScheduledInstallTime(v string) {
	o.ScheduledInstallTime = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o WindowsUpdateScheduledInstall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScheduledInstallDay != nil {
		toSerialize["scheduledInstallDay"] = o.ScheduledInstallDay
	}
	if o.ScheduledInstallTime != nil {
		toSerialize["scheduledInstallTime"] = o.ScheduledInstallTime
	}
	return json.Marshal(toSerialize)
}



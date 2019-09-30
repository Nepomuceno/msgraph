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
// InlineObject32 struct for InlineObject32
type InlineObject32 struct {
	Schedules *[]string `json:"Schedules,omitempty"`

	EndTime *AnyOfmicrosoftGraphDateTimeTimeZone `json:"EndTime,omitempty"`
	isExplicitNullEndTime bool `json:"-"`
	StartTime *AnyOfmicrosoftGraphDateTimeTimeZone `json:"StartTime,omitempty"`
	isExplicitNullStartTime bool `json:"-"`
	AvailabilityViewInterval *int32 `json:"AvailabilityViewInterval,omitempty"`
	isExplicitNullAvailabilityViewInterval bool `json:"-"`
}

// GetSchedules returns the Schedules field if non-nil, zero value otherwise.
func (o *InlineObject32) GetSchedules() []string {
	if o == nil || o.Schedules == nil {
		var ret []string
		return ret
	}
	return *o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject32) GetSchedulesOk() ([]string, bool) {
	if o == nil || o.Schedules == nil {
		var ret []string
		return ret, false
	}
	return *o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *InlineObject32) HasSchedules() bool {
	if o != nil && o.Schedules != nil {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []string and assigns it to the Schedules field.
func (o *InlineObject32) SetSchedules(v []string) {
	o.Schedules = &v
}

// GetEndTime returns the EndTime field if non-nil, zero value otherwise.
func (o *InlineObject32) GetEndTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil || o.EndTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject32) GetEndTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.EndTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret, false
	}
	return *o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *InlineObject32) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the EndTime field.
func (o *InlineObject32) SetEndTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.EndTime = &v
}

// SetEndTimeExplicitNull (un)sets EndTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The EndTime value is set to nil even if false is passed
func (o *InlineObject32) SetEndTimeExplicitNull(b bool) {
	o.EndTime = nil
	o.isExplicitNullEndTime = b
}
// GetStartTime returns the StartTime field if non-nil, zero value otherwise.
func (o *InlineObject32) GetStartTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil || o.StartTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject32) GetStartTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.StartTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret, false
	}
	return *o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *InlineObject32) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the StartTime field.
func (o *InlineObject32) SetStartTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.StartTime = &v
}

// SetStartTimeExplicitNull (un)sets StartTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The StartTime value is set to nil even if false is passed
func (o *InlineObject32) SetStartTimeExplicitNull(b bool) {
	o.StartTime = nil
	o.isExplicitNullStartTime = b
}
// GetAvailabilityViewInterval returns the AvailabilityViewInterval field if non-nil, zero value otherwise.
func (o *InlineObject32) GetAvailabilityViewInterval() int32 {
	if o == nil || o.AvailabilityViewInterval == nil {
		var ret int32
		return ret
	}
	return *o.AvailabilityViewInterval
}

// GetAvailabilityViewIntervalOk returns a tuple with the AvailabilityViewInterval field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject32) GetAvailabilityViewIntervalOk() (int32, bool) {
	if o == nil || o.AvailabilityViewInterval == nil {
		var ret int32
		return ret, false
	}
	return *o.AvailabilityViewInterval, true
}

// HasAvailabilityViewInterval returns a boolean if a field has been set.
func (o *InlineObject32) HasAvailabilityViewInterval() bool {
	if o != nil && o.AvailabilityViewInterval != nil {
		return true
	}

	return false
}

// SetAvailabilityViewInterval gets a reference to the given int32 and assigns it to the AvailabilityViewInterval field.
func (o *InlineObject32) SetAvailabilityViewInterval(v int32) {
	o.AvailabilityViewInterval = &v
}

// SetAvailabilityViewIntervalExplicitNull (un)sets AvailabilityViewInterval to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AvailabilityViewInterval value is set to nil even if false is passed
func (o *InlineObject32) SetAvailabilityViewIntervalExplicitNull(b bool) {
	o.AvailabilityViewInterval = nil
	o.isExplicitNullAvailabilityViewInterval = b
}

// MarshalJSON returns the JSON representation of the model.
func (o InlineObject32) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schedules != nil {
		toSerialize["Schedules"] = o.Schedules
	}
	if o.EndTime == nil {
		if o.isExplicitNullEndTime {
			toSerialize["EndTime"] = o.EndTime
		}
	} else {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.StartTime == nil {
		if o.isExplicitNullStartTime {
			toSerialize["StartTime"] = o.StartTime
		}
	} else {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.AvailabilityViewInterval == nil {
		if o.isExplicitNullAvailabilityViewInterval {
			toSerialize["AvailabilityViewInterval"] = o.AvailabilityViewInterval
		}
	} else {
		toSerialize["AvailabilityViewInterval"] = o.AvailabilityViewInterval
	}
	return json.Marshal(toSerialize)
}



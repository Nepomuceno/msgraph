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
// MicrosoftGraphAutomaticRepliesSetting struct for MicrosoftGraphAutomaticRepliesSetting
type MicrosoftGraphAutomaticRepliesSetting struct {
	Status *AnyOfmicrosoftGraphAutomaticRepliesStatus `json:"status,omitempty"`
	isExplicitNullStatus bool `json:"-"`
	ExternalAudience *AnyOfmicrosoftGraphExternalAudienceScope `json:"externalAudience,omitempty"`
	isExplicitNullExternalAudience bool `json:"-"`
	ScheduledStartDateTime *AnyOfmicrosoftGraphDateTimeTimeZone `json:"scheduledStartDateTime,omitempty"`
	isExplicitNullScheduledStartDateTime bool `json:"-"`
	ScheduledEndDateTime *AnyOfmicrosoftGraphDateTimeTimeZone `json:"scheduledEndDateTime,omitempty"`
	isExplicitNullScheduledEndDateTime bool `json:"-"`
	InternalReplyMessage *string `json:"internalReplyMessage,omitempty"`
	isExplicitNullInternalReplyMessage bool `json:"-"`
	ExternalReplyMessage *string `json:"externalReplyMessage,omitempty"`
	isExplicitNullExternalReplyMessage bool `json:"-"`
}

// GetStatus returns the Status field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetStatus() AnyOfmicrosoftGraphAutomaticRepliesStatus {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphAutomaticRepliesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetStatusOk() (AnyOfmicrosoftGraphAutomaticRepliesStatus, bool) {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphAutomaticRepliesStatus
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphAutomaticRepliesStatus and assigns it to the Status field.
func (o *MicrosoftGraphAutomaticRepliesSetting) SetStatus(v AnyOfmicrosoftGraphAutomaticRepliesStatus) {
	o.Status = &v
}

// SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Status value is set to nil even if false is passed
func (o *MicrosoftGraphAutomaticRepliesSetting) SetStatusExplicitNull(b bool) {
	o.Status = nil
	o.isExplicitNullStatus = b
}
// GetExternalAudience returns the ExternalAudience field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalAudience() AnyOfmicrosoftGraphExternalAudienceScope {
	if o == nil || o.ExternalAudience == nil {
		var ret AnyOfmicrosoftGraphExternalAudienceScope
		return ret
	}
	return *o.ExternalAudience
}

// GetExternalAudienceOk returns a tuple with the ExternalAudience field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalAudienceOk() (AnyOfmicrosoftGraphExternalAudienceScope, bool) {
	if o == nil || o.ExternalAudience == nil {
		var ret AnyOfmicrosoftGraphExternalAudienceScope
		return ret, false
	}
	return *o.ExternalAudience, true
}

// HasExternalAudience returns a boolean if a field has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) HasExternalAudience() bool {
	if o != nil && o.ExternalAudience != nil {
		return true
	}

	return false
}

// SetExternalAudience gets a reference to the given AnyOfmicrosoftGraphExternalAudienceScope and assigns it to the ExternalAudience field.
func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalAudience(v AnyOfmicrosoftGraphExternalAudienceScope) {
	o.ExternalAudience = &v
}

// SetExternalAudienceExplicitNull (un)sets ExternalAudience to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ExternalAudience value is set to nil even if false is passed
func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalAudienceExplicitNull(b bool) {
	o.ExternalAudience = nil
	o.isExplicitNullExternalAudience = b
}
// GetScheduledStartDateTime returns the ScheduledStartDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledStartDateTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil || o.ScheduledStartDateTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return *o.ScheduledStartDateTime
}

// GetScheduledStartDateTimeOk returns a tuple with the ScheduledStartDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledStartDateTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.ScheduledStartDateTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret, false
	}
	return *o.ScheduledStartDateTime, true
}

// HasScheduledStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) HasScheduledStartDateTime() bool {
	if o != nil && o.ScheduledStartDateTime != nil {
		return true
	}

	return false
}

// SetScheduledStartDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the ScheduledStartDateTime field.
func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledStartDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.ScheduledStartDateTime = &v
}

// SetScheduledStartDateTimeExplicitNull (un)sets ScheduledStartDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ScheduledStartDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledStartDateTimeExplicitNull(b bool) {
	o.ScheduledStartDateTime = nil
	o.isExplicitNullScheduledStartDateTime = b
}
// GetScheduledEndDateTime returns the ScheduledEndDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledEndDateTime() AnyOfmicrosoftGraphDateTimeTimeZone {
	if o == nil || o.ScheduledEndDateTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret
	}
	return *o.ScheduledEndDateTime
}

// GetScheduledEndDateTimeOk returns a tuple with the ScheduledEndDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetScheduledEndDateTimeOk() (AnyOfmicrosoftGraphDateTimeTimeZone, bool) {
	if o == nil || o.ScheduledEndDateTime == nil {
		var ret AnyOfmicrosoftGraphDateTimeTimeZone
		return ret, false
	}
	return *o.ScheduledEndDateTime, true
}

// HasScheduledEndDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) HasScheduledEndDateTime() bool {
	if o != nil && o.ScheduledEndDateTime != nil {
		return true
	}

	return false
}

// SetScheduledEndDateTime gets a reference to the given AnyOfmicrosoftGraphDateTimeTimeZone and assigns it to the ScheduledEndDateTime field.
func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledEndDateTime(v AnyOfmicrosoftGraphDateTimeTimeZone) {
	o.ScheduledEndDateTime = &v
}

// SetScheduledEndDateTimeExplicitNull (un)sets ScheduledEndDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ScheduledEndDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphAutomaticRepliesSetting) SetScheduledEndDateTimeExplicitNull(b bool) {
	o.ScheduledEndDateTime = nil
	o.isExplicitNullScheduledEndDateTime = b
}
// GetInternalReplyMessage returns the InternalReplyMessage field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetInternalReplyMessage() string {
	if o == nil || o.InternalReplyMessage == nil {
		var ret string
		return ret
	}
	return *o.InternalReplyMessage
}

// GetInternalReplyMessageOk returns a tuple with the InternalReplyMessage field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetInternalReplyMessageOk() (string, bool) {
	if o == nil || o.InternalReplyMessage == nil {
		var ret string
		return ret, false
	}
	return *o.InternalReplyMessage, true
}

// HasInternalReplyMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) HasInternalReplyMessage() bool {
	if o != nil && o.InternalReplyMessage != nil {
		return true
	}

	return false
}

// SetInternalReplyMessage gets a reference to the given string and assigns it to the InternalReplyMessage field.
func (o *MicrosoftGraphAutomaticRepliesSetting) SetInternalReplyMessage(v string) {
	o.InternalReplyMessage = &v
}

// SetInternalReplyMessageExplicitNull (un)sets InternalReplyMessage to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The InternalReplyMessage value is set to nil even if false is passed
func (o *MicrosoftGraphAutomaticRepliesSetting) SetInternalReplyMessageExplicitNull(b bool) {
	o.InternalReplyMessage = nil
	o.isExplicitNullInternalReplyMessage = b
}
// GetExternalReplyMessage returns the ExternalReplyMessage field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalReplyMessage() string {
	if o == nil || o.ExternalReplyMessage == nil {
		var ret string
		return ret
	}
	return *o.ExternalReplyMessage
}

// GetExternalReplyMessageOk returns a tuple with the ExternalReplyMessage field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) GetExternalReplyMessageOk() (string, bool) {
	if o == nil || o.ExternalReplyMessage == nil {
		var ret string
		return ret, false
	}
	return *o.ExternalReplyMessage, true
}

// HasExternalReplyMessage returns a boolean if a field has been set.
func (o *MicrosoftGraphAutomaticRepliesSetting) HasExternalReplyMessage() bool {
	if o != nil && o.ExternalReplyMessage != nil {
		return true
	}

	return false
}

// SetExternalReplyMessage gets a reference to the given string and assigns it to the ExternalReplyMessage field.
func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalReplyMessage(v string) {
	o.ExternalReplyMessage = &v
}

// SetExternalReplyMessageExplicitNull (un)sets ExternalReplyMessage to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ExternalReplyMessage value is set to nil even if false is passed
func (o *MicrosoftGraphAutomaticRepliesSetting) SetExternalReplyMessageExplicitNull(b bool) {
	o.ExternalReplyMessage = nil
	o.isExplicitNullExternalReplyMessage = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphAutomaticRepliesSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status == nil {
		if o.isExplicitNullStatus {
			toSerialize["status"] = o.Status
		}
	} else {
		toSerialize["status"] = o.Status
	}
	if o.ExternalAudience == nil {
		if o.isExplicitNullExternalAudience {
			toSerialize["externalAudience"] = o.ExternalAudience
		}
	} else {
		toSerialize["externalAudience"] = o.ExternalAudience
	}
	if o.ScheduledStartDateTime == nil {
		if o.isExplicitNullScheduledStartDateTime {
			toSerialize["scheduledStartDateTime"] = o.ScheduledStartDateTime
		}
	} else {
		toSerialize["scheduledStartDateTime"] = o.ScheduledStartDateTime
	}
	if o.ScheduledEndDateTime == nil {
		if o.isExplicitNullScheduledEndDateTime {
			toSerialize["scheduledEndDateTime"] = o.ScheduledEndDateTime
		}
	} else {
		toSerialize["scheduledEndDateTime"] = o.ScheduledEndDateTime
	}
	if o.InternalReplyMessage == nil {
		if o.isExplicitNullInternalReplyMessage {
			toSerialize["internalReplyMessage"] = o.InternalReplyMessage
		}
	} else {
		toSerialize["internalReplyMessage"] = o.InternalReplyMessage
	}
	if o.ExternalReplyMessage == nil {
		if o.isExplicitNullExternalReplyMessage {
			toSerialize["externalReplyMessage"] = o.ExternalReplyMessage
		}
	} else {
		toSerialize["externalReplyMessage"] = o.ExternalReplyMessage
	}
	return json.Marshal(toSerialize)
}



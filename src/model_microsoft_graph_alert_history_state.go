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
// MicrosoftGraphAlertHistoryState struct for MicrosoftGraphAlertHistoryState
type MicrosoftGraphAlertHistoryState struct {
	AppId *string `json:"appId,omitempty"`
	isExplicitNullAppId bool `json:"-"`
	AssignedTo *string `json:"assignedTo,omitempty"`
	isExplicitNullAssignedTo bool `json:"-"`
	Comments *[]string `json:"comments,omitempty"`

	Feedback *AnyOfmicrosoftGraphAlertFeedback `json:"feedback,omitempty"`
	isExplicitNullFeedback bool `json:"-"`
	Status *AnyOfmicrosoftGraphAlertStatus `json:"status,omitempty"`
	isExplicitNullStatus bool `json:"-"`
	UpdatedDateTime *time.Time `json:"updatedDateTime,omitempty"`
	isExplicitNullUpdatedDateTime bool `json:"-"`
	User *string `json:"user,omitempty"`
	isExplicitNullUser bool `json:"-"`
}

// GetAppId returns the AppId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAlertHistoryState) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAlertHistoryState) GetAppIdOk() (string, bool) {
	if o == nil || o.AppId == nil {
		var ret string
		return ret, false
	}
	return *o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *MicrosoftGraphAlertHistoryState) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *MicrosoftGraphAlertHistoryState) SetAppId(v string) {
	o.AppId = &v
}

// SetAppIdExplicitNull (un)sets AppId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AppId value is set to nil even if false is passed
func (o *MicrosoftGraphAlertHistoryState) SetAppIdExplicitNull(b bool) {
	o.AppId = nil
	o.isExplicitNullAppId = b
}
// GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAlertHistoryState) GetAssignedTo() string {
	if o == nil || o.AssignedTo == nil {
		var ret string
		return ret
	}
	return *o.AssignedTo
}

// GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAlertHistoryState) GetAssignedToOk() (string, bool) {
	if o == nil || o.AssignedTo == nil {
		var ret string
		return ret, false
	}
	return *o.AssignedTo, true
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *MicrosoftGraphAlertHistoryState) HasAssignedTo() bool {
	if o != nil && o.AssignedTo != nil {
		return true
	}

	return false
}

// SetAssignedTo gets a reference to the given string and assigns it to the AssignedTo field.
func (o *MicrosoftGraphAlertHistoryState) SetAssignedTo(v string) {
	o.AssignedTo = &v
}

// SetAssignedToExplicitNull (un)sets AssignedTo to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The AssignedTo value is set to nil even if false is passed
func (o *MicrosoftGraphAlertHistoryState) SetAssignedToExplicitNull(b bool) {
	o.AssignedTo = nil
	o.isExplicitNullAssignedTo = b
}
// GetComments returns the Comments field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAlertHistoryState) GetComments() []string {
	if o == nil || o.Comments == nil {
		var ret []string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAlertHistoryState) GetCommentsOk() ([]string, bool) {
	if o == nil || o.Comments == nil {
		var ret []string
		return ret, false
	}
	return *o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *MicrosoftGraphAlertHistoryState) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given []string and assigns it to the Comments field.
func (o *MicrosoftGraphAlertHistoryState) SetComments(v []string) {
	o.Comments = &v
}

// GetFeedback returns the Feedback field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAlertHistoryState) GetFeedback() AnyOfmicrosoftGraphAlertFeedback {
	if o == nil || o.Feedback == nil {
		var ret AnyOfmicrosoftGraphAlertFeedback
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAlertHistoryState) GetFeedbackOk() (AnyOfmicrosoftGraphAlertFeedback, bool) {
	if o == nil || o.Feedback == nil {
		var ret AnyOfmicrosoftGraphAlertFeedback
		return ret, false
	}
	return *o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *MicrosoftGraphAlertHistoryState) HasFeedback() bool {
	if o != nil && o.Feedback != nil {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given AnyOfmicrosoftGraphAlertFeedback and assigns it to the Feedback field.
func (o *MicrosoftGraphAlertHistoryState) SetFeedback(v AnyOfmicrosoftGraphAlertFeedback) {
	o.Feedback = &v
}

// SetFeedbackExplicitNull (un)sets Feedback to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Feedback value is set to nil even if false is passed
func (o *MicrosoftGraphAlertHistoryState) SetFeedbackExplicitNull(b bool) {
	o.Feedback = nil
	o.isExplicitNullFeedback = b
}
// GetStatus returns the Status field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAlertHistoryState) GetStatus() AnyOfmicrosoftGraphAlertStatus {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphAlertStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAlertHistoryState) GetStatusOk() (AnyOfmicrosoftGraphAlertStatus, bool) {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphAlertStatus
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphAlertHistoryState) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphAlertStatus and assigns it to the Status field.
func (o *MicrosoftGraphAlertHistoryState) SetStatus(v AnyOfmicrosoftGraphAlertStatus) {
	o.Status = &v
}

// SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Status value is set to nil even if false is passed
func (o *MicrosoftGraphAlertHistoryState) SetStatusExplicitNull(b bool) {
	o.Status = nil
	o.isExplicitNullStatus = b
}
// GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAlertHistoryState) GetUpdatedDateTime() time.Time {
	if o == nil || o.UpdatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedDateTime
}

// GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAlertHistoryState) GetUpdatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.UpdatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.UpdatedDateTime, true
}

// HasUpdatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphAlertHistoryState) HasUpdatedDateTime() bool {
	if o != nil && o.UpdatedDateTime != nil {
		return true
	}

	return false
}

// SetUpdatedDateTime gets a reference to the given time.Time and assigns it to the UpdatedDateTime field.
func (o *MicrosoftGraphAlertHistoryState) SetUpdatedDateTime(v time.Time) {
	o.UpdatedDateTime = &v
}

// SetUpdatedDateTimeExplicitNull (un)sets UpdatedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The UpdatedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphAlertHistoryState) SetUpdatedDateTimeExplicitNull(b bool) {
	o.UpdatedDateTime = nil
	o.isExplicitNullUpdatedDateTime = b
}
// GetUser returns the User field if non-nil, zero value otherwise.
func (o *MicrosoftGraphAlertHistoryState) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAlertHistoryState) GetUserOk() (string, bool) {
	if o == nil || o.User == nil {
		var ret string
		return ret, false
	}
	return *o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MicrosoftGraphAlertHistoryState) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *MicrosoftGraphAlertHistoryState) SetUser(v string) {
	o.User = &v
}

// SetUserExplicitNull (un)sets User to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The User value is set to nil even if false is passed
func (o *MicrosoftGraphAlertHistoryState) SetUserExplicitNull(b bool) {
	o.User = nil
	o.isExplicitNullUser = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphAlertHistoryState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId == nil {
		if o.isExplicitNullAppId {
			toSerialize["appId"] = o.AppId
		}
	} else {
		toSerialize["appId"] = o.AppId
	}
	if o.AssignedTo == nil {
		if o.isExplicitNullAssignedTo {
			toSerialize["assignedTo"] = o.AssignedTo
		}
	} else {
		toSerialize["assignedTo"] = o.AssignedTo
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Feedback == nil {
		if o.isExplicitNullFeedback {
			toSerialize["feedback"] = o.Feedback
		}
	} else {
		toSerialize["feedback"] = o.Feedback
	}
	if o.Status == nil {
		if o.isExplicitNullStatus {
			toSerialize["status"] = o.Status
		}
	} else {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedDateTime == nil {
		if o.isExplicitNullUpdatedDateTime {
			toSerialize["updatedDateTime"] = o.UpdatedDateTime
		}
	} else {
		toSerialize["updatedDateTime"] = o.UpdatedDateTime
	}
	if o.User == nil {
		if o.isExplicitNullUser {
			toSerialize["user"] = o.User
		}
	} else {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}



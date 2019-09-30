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
// ItemActivityStat struct for ItemActivityStat
type ItemActivityStat struct {
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	isExplicitNullStartDateTime bool `json:"-"`
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	isExplicitNullEndDateTime bool `json:"-"`
	Access *AnyOfmicrosoftGraphItemActionStat `json:"access,omitempty"`
	isExplicitNullAccess bool `json:"-"`
	Create *AnyOfmicrosoftGraphItemActionStat `json:"create,omitempty"`
	isExplicitNullCreate bool `json:"-"`
	Delete *AnyOfmicrosoftGraphItemActionStat `json:"delete,omitempty"`
	isExplicitNullDelete bool `json:"-"`
	Edit *AnyOfmicrosoftGraphItemActionStat `json:"edit,omitempty"`
	isExplicitNullEdit bool `json:"-"`
	Move *AnyOfmicrosoftGraphItemActionStat `json:"move,omitempty"`
	isExplicitNullMove bool `json:"-"`
	IsTrending *bool `json:"isTrending,omitempty"`
	isExplicitNullIsTrending bool `json:"-"`
	IncompleteData *AnyOfmicrosoftGraphIncompleteData `json:"incompleteData,omitempty"`
	isExplicitNullIncompleteData bool `json:"-"`
	Activities *[]MicrosoftGraphItemActivity `json:"activities,omitempty"`

}

// GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetStartDateTimeOk() (time.Time, bool) {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *ItemActivityStat) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *ItemActivityStat) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// SetStartDateTimeExplicitNull (un)sets StartDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The StartDateTime value is set to nil even if false is passed
func (o *ItemActivityStat) SetStartDateTimeExplicitNull(b bool) {
	o.StartDateTime = nil
	o.isExplicitNullStartDateTime = b
}
// GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetEndDateTimeOk() (time.Time, bool) {
	if o == nil || o.EndDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *ItemActivityStat) HasEndDateTime() bool {
	if o != nil && o.EndDateTime != nil {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.
func (o *ItemActivityStat) SetEndDateTime(v time.Time) {
	o.EndDateTime = &v
}

// SetEndDateTimeExplicitNull (un)sets EndDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The EndDateTime value is set to nil even if false is passed
func (o *ItemActivityStat) SetEndDateTimeExplicitNull(b bool) {
	o.EndDateTime = nil
	o.isExplicitNullEndDateTime = b
}
// GetAccess returns the Access field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetAccess() AnyOfmicrosoftGraphItemActionStat {
	if o == nil || o.Access == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetAccessOk() (AnyOfmicrosoftGraphItemActionStat, bool) {
	if o == nil || o.Access == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret, false
	}
	return *o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ItemActivityStat) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Access field.
func (o *ItemActivityStat) SetAccess(v AnyOfmicrosoftGraphItemActionStat) {
	o.Access = &v
}

// SetAccessExplicitNull (un)sets Access to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Access value is set to nil even if false is passed
func (o *ItemActivityStat) SetAccessExplicitNull(b bool) {
	o.Access = nil
	o.isExplicitNullAccess = b
}
// GetCreate returns the Create field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetCreate() AnyOfmicrosoftGraphItemActionStat {
	if o == nil || o.Create == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetCreateOk() (AnyOfmicrosoftGraphItemActionStat, bool) {
	if o == nil || o.Create == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret, false
	}
	return *o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ItemActivityStat) HasCreate() bool {
	if o != nil && o.Create != nil {
		return true
	}

	return false
}

// SetCreate gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Create field.
func (o *ItemActivityStat) SetCreate(v AnyOfmicrosoftGraphItemActionStat) {
	o.Create = &v
}

// SetCreateExplicitNull (un)sets Create to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Create value is set to nil even if false is passed
func (o *ItemActivityStat) SetCreateExplicitNull(b bool) {
	o.Create = nil
	o.isExplicitNullCreate = b
}
// GetDelete returns the Delete field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetDelete() AnyOfmicrosoftGraphItemActionStat {
	if o == nil || o.Delete == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetDeleteOk() (AnyOfmicrosoftGraphItemActionStat, bool) {
	if o == nil || o.Delete == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret, false
	}
	return *o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *ItemActivityStat) HasDelete() bool {
	if o != nil && o.Delete != nil {
		return true
	}

	return false
}

// SetDelete gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Delete field.
func (o *ItemActivityStat) SetDelete(v AnyOfmicrosoftGraphItemActionStat) {
	o.Delete = &v
}

// SetDeleteExplicitNull (un)sets Delete to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Delete value is set to nil even if false is passed
func (o *ItemActivityStat) SetDeleteExplicitNull(b bool) {
	o.Delete = nil
	o.isExplicitNullDelete = b
}
// GetEdit returns the Edit field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetEdit() AnyOfmicrosoftGraphItemActionStat {
	if o == nil || o.Edit == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret
	}
	return *o.Edit
}

// GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetEditOk() (AnyOfmicrosoftGraphItemActionStat, bool) {
	if o == nil || o.Edit == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret, false
	}
	return *o.Edit, true
}

// HasEdit returns a boolean if a field has been set.
func (o *ItemActivityStat) HasEdit() bool {
	if o != nil && o.Edit != nil {
		return true
	}

	return false
}

// SetEdit gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Edit field.
func (o *ItemActivityStat) SetEdit(v AnyOfmicrosoftGraphItemActionStat) {
	o.Edit = &v
}

// SetEditExplicitNull (un)sets Edit to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Edit value is set to nil even if false is passed
func (o *ItemActivityStat) SetEditExplicitNull(b bool) {
	o.Edit = nil
	o.isExplicitNullEdit = b
}
// GetMove returns the Move field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetMove() AnyOfmicrosoftGraphItemActionStat {
	if o == nil || o.Move == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret
	}
	return *o.Move
}

// GetMoveOk returns a tuple with the Move field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetMoveOk() (AnyOfmicrosoftGraphItemActionStat, bool) {
	if o == nil || o.Move == nil {
		var ret AnyOfmicrosoftGraphItemActionStat
		return ret, false
	}
	return *o.Move, true
}

// HasMove returns a boolean if a field has been set.
func (o *ItemActivityStat) HasMove() bool {
	if o != nil && o.Move != nil {
		return true
	}

	return false
}

// SetMove gets a reference to the given AnyOfmicrosoftGraphItemActionStat and assigns it to the Move field.
func (o *ItemActivityStat) SetMove(v AnyOfmicrosoftGraphItemActionStat) {
	o.Move = &v
}

// SetMoveExplicitNull (un)sets Move to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Move value is set to nil even if false is passed
func (o *ItemActivityStat) SetMoveExplicitNull(b bool) {
	o.Move = nil
	o.isExplicitNullMove = b
}
// GetIsTrending returns the IsTrending field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetIsTrending() bool {
	if o == nil || o.IsTrending == nil {
		var ret bool
		return ret
	}
	return *o.IsTrending
}

// GetIsTrendingOk returns a tuple with the IsTrending field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetIsTrendingOk() (bool, bool) {
	if o == nil || o.IsTrending == nil {
		var ret bool
		return ret, false
	}
	return *o.IsTrending, true
}

// HasIsTrending returns a boolean if a field has been set.
func (o *ItemActivityStat) HasIsTrending() bool {
	if o != nil && o.IsTrending != nil {
		return true
	}

	return false
}

// SetIsTrending gets a reference to the given bool and assigns it to the IsTrending field.
func (o *ItemActivityStat) SetIsTrending(v bool) {
	o.IsTrending = &v
}

// SetIsTrendingExplicitNull (un)sets IsTrending to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The IsTrending value is set to nil even if false is passed
func (o *ItemActivityStat) SetIsTrendingExplicitNull(b bool) {
	o.IsTrending = nil
	o.isExplicitNullIsTrending = b
}
// GetIncompleteData returns the IncompleteData field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetIncompleteData() AnyOfmicrosoftGraphIncompleteData {
	if o == nil || o.IncompleteData == nil {
		var ret AnyOfmicrosoftGraphIncompleteData
		return ret
	}
	return *o.IncompleteData
}

// GetIncompleteDataOk returns a tuple with the IncompleteData field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetIncompleteDataOk() (AnyOfmicrosoftGraphIncompleteData, bool) {
	if o == nil || o.IncompleteData == nil {
		var ret AnyOfmicrosoftGraphIncompleteData
		return ret, false
	}
	return *o.IncompleteData, true
}

// HasIncompleteData returns a boolean if a field has been set.
func (o *ItemActivityStat) HasIncompleteData() bool {
	if o != nil && o.IncompleteData != nil {
		return true
	}

	return false
}

// SetIncompleteData gets a reference to the given AnyOfmicrosoftGraphIncompleteData and assigns it to the IncompleteData field.
func (o *ItemActivityStat) SetIncompleteData(v AnyOfmicrosoftGraphIncompleteData) {
	o.IncompleteData = &v
}

// SetIncompleteDataExplicitNull (un)sets IncompleteData to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The IncompleteData value is set to nil even if false is passed
func (o *ItemActivityStat) SetIncompleteDataExplicitNull(b bool) {
	o.IncompleteData = nil
	o.isExplicitNullIncompleteData = b
}
// GetActivities returns the Activities field if non-nil, zero value otherwise.
func (o *ItemActivityStat) GetActivities() []MicrosoftGraphItemActivity {
	if o == nil || o.Activities == nil {
		var ret []MicrosoftGraphItemActivity
		return ret
	}
	return *o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ItemActivityStat) GetActivitiesOk() ([]MicrosoftGraphItemActivity, bool) {
	if o == nil || o.Activities == nil {
		var ret []MicrosoftGraphItemActivity
		return ret, false
	}
	return *o.Activities, true
}

// HasActivities returns a boolean if a field has been set.
func (o *ItemActivityStat) HasActivities() bool {
	if o != nil && o.Activities != nil {
		return true
	}

	return false
}

// SetActivities gets a reference to the given []MicrosoftGraphItemActivity and assigns it to the Activities field.
func (o *ItemActivityStat) SetActivities(v []MicrosoftGraphItemActivity) {
	o.Activities = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o ItemActivityStat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDateTime == nil {
		if o.isExplicitNullStartDateTime {
			toSerialize["startDateTime"] = o.StartDateTime
		}
	} else {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if o.EndDateTime == nil {
		if o.isExplicitNullEndDateTime {
			toSerialize["endDateTime"] = o.EndDateTime
		}
	} else {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	if o.Access == nil {
		if o.isExplicitNullAccess {
			toSerialize["access"] = o.Access
		}
	} else {
		toSerialize["access"] = o.Access
	}
	if o.Create == nil {
		if o.isExplicitNullCreate {
			toSerialize["create"] = o.Create
		}
	} else {
		toSerialize["create"] = o.Create
	}
	if o.Delete == nil {
		if o.isExplicitNullDelete {
			toSerialize["delete"] = o.Delete
		}
	} else {
		toSerialize["delete"] = o.Delete
	}
	if o.Edit == nil {
		if o.isExplicitNullEdit {
			toSerialize["edit"] = o.Edit
		}
	} else {
		toSerialize["edit"] = o.Edit
	}
	if o.Move == nil {
		if o.isExplicitNullMove {
			toSerialize["move"] = o.Move
		}
	} else {
		toSerialize["move"] = o.Move
	}
	if o.IsTrending == nil {
		if o.isExplicitNullIsTrending {
			toSerialize["isTrending"] = o.IsTrending
		}
	} else {
		toSerialize["isTrending"] = o.IsTrending
	}
	if o.IncompleteData == nil {
		if o.isExplicitNullIncompleteData {
			toSerialize["incompleteData"] = o.IncompleteData
		}
	} else {
		toSerialize["incompleteData"] = o.IncompleteData
	}
	if o.Activities != nil {
		toSerialize["activities"] = o.Activities
	}
	return json.Marshal(toSerialize)
}


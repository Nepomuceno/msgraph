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
// TeamsAsyncOperation struct for TeamsAsyncOperation
type TeamsAsyncOperation struct {
	OperationType *AnyOfmicrosoftGraphTeamsAsyncOperationType `json:"operationType,omitempty"`

	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`

	Status *AnyOfmicrosoftGraphTeamsAsyncOperationStatus `json:"status,omitempty"`

	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`

	AttemptsCount *int32 `json:"attemptsCount,omitempty"`

	TargetResourceId *string `json:"targetResourceId,omitempty"`
	isExplicitNullTargetResourceId bool `json:"-"`
	TargetResourceLocation *string `json:"targetResourceLocation,omitempty"`
	isExplicitNullTargetResourceLocation bool `json:"-"`
	Error *AnyOfmicrosoftGraphOperationError `json:"error,omitempty"`
	isExplicitNullError bool `json:"-"`
}

// GetOperationType returns the OperationType field if non-nil, zero value otherwise.
func (o *TeamsAsyncOperation) GetOperationType() AnyOfmicrosoftGraphTeamsAsyncOperationType {
	if o == nil || o.OperationType == nil {
		var ret AnyOfmicrosoftGraphTeamsAsyncOperationType
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeamsAsyncOperation) GetOperationTypeOk() (AnyOfmicrosoftGraphTeamsAsyncOperationType, bool) {
	if o == nil || o.OperationType == nil {
		var ret AnyOfmicrosoftGraphTeamsAsyncOperationType
		return ret, false
	}
	return *o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *TeamsAsyncOperation) HasOperationType() bool {
	if o != nil && o.OperationType != nil {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given AnyOfmicrosoftGraphTeamsAsyncOperationType and assigns it to the OperationType field.
func (o *TeamsAsyncOperation) SetOperationType(v AnyOfmicrosoftGraphTeamsAsyncOperationType) {
	o.OperationType = &v
}

// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *TeamsAsyncOperation) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeamsAsyncOperation) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *TeamsAsyncOperation) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *TeamsAsyncOperation) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetStatus returns the Status field if non-nil, zero value otherwise.
func (o *TeamsAsyncOperation) GetStatus() AnyOfmicrosoftGraphTeamsAsyncOperationStatus {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphTeamsAsyncOperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeamsAsyncOperation) GetStatusOk() (AnyOfmicrosoftGraphTeamsAsyncOperationStatus, bool) {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphTeamsAsyncOperationStatus
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TeamsAsyncOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphTeamsAsyncOperationStatus and assigns it to the Status field.
func (o *TeamsAsyncOperation) SetStatus(v AnyOfmicrosoftGraphTeamsAsyncOperationStatus) {
	o.Status = &v
}

// GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.
func (o *TeamsAsyncOperation) GetLastActionDateTime() time.Time {
	if o == nil || o.LastActionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActionDateTime
}

// GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeamsAsyncOperation) GetLastActionDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastActionDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastActionDateTime, true
}

// HasLastActionDateTime returns a boolean if a field has been set.
func (o *TeamsAsyncOperation) HasLastActionDateTime() bool {
	if o != nil && o.LastActionDateTime != nil {
		return true
	}

	return false
}

// SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.
func (o *TeamsAsyncOperation) SetLastActionDateTime(v time.Time) {
	o.LastActionDateTime = &v
}

// GetAttemptsCount returns the AttemptsCount field if non-nil, zero value otherwise.
func (o *TeamsAsyncOperation) GetAttemptsCount() int32 {
	if o == nil || o.AttemptsCount == nil {
		var ret int32
		return ret
	}
	return *o.AttemptsCount
}

// GetAttemptsCountOk returns a tuple with the AttemptsCount field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeamsAsyncOperation) GetAttemptsCountOk() (int32, bool) {
	if o == nil || o.AttemptsCount == nil {
		var ret int32
		return ret, false
	}
	return *o.AttemptsCount, true
}

// HasAttemptsCount returns a boolean if a field has been set.
func (o *TeamsAsyncOperation) HasAttemptsCount() bool {
	if o != nil && o.AttemptsCount != nil {
		return true
	}

	return false
}

// SetAttemptsCount gets a reference to the given int32 and assigns it to the AttemptsCount field.
func (o *TeamsAsyncOperation) SetAttemptsCount(v int32) {
	o.AttemptsCount = &v
}

// GetTargetResourceId returns the TargetResourceId field if non-nil, zero value otherwise.
func (o *TeamsAsyncOperation) GetTargetResourceId() string {
	if o == nil || o.TargetResourceId == nil {
		var ret string
		return ret
	}
	return *o.TargetResourceId
}

// GetTargetResourceIdOk returns a tuple with the TargetResourceId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeamsAsyncOperation) GetTargetResourceIdOk() (string, bool) {
	if o == nil || o.TargetResourceId == nil {
		var ret string
		return ret, false
	}
	return *o.TargetResourceId, true
}

// HasTargetResourceId returns a boolean if a field has been set.
func (o *TeamsAsyncOperation) HasTargetResourceId() bool {
	if o != nil && o.TargetResourceId != nil {
		return true
	}

	return false
}

// SetTargetResourceId gets a reference to the given string and assigns it to the TargetResourceId field.
func (o *TeamsAsyncOperation) SetTargetResourceId(v string) {
	o.TargetResourceId = &v
}

// SetTargetResourceIdExplicitNull (un)sets TargetResourceId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The TargetResourceId value is set to nil even if false is passed
func (o *TeamsAsyncOperation) SetTargetResourceIdExplicitNull(b bool) {
	o.TargetResourceId = nil
	o.isExplicitNullTargetResourceId = b
}
// GetTargetResourceLocation returns the TargetResourceLocation field if non-nil, zero value otherwise.
func (o *TeamsAsyncOperation) GetTargetResourceLocation() string {
	if o == nil || o.TargetResourceLocation == nil {
		var ret string
		return ret
	}
	return *o.TargetResourceLocation
}

// GetTargetResourceLocationOk returns a tuple with the TargetResourceLocation field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeamsAsyncOperation) GetTargetResourceLocationOk() (string, bool) {
	if o == nil || o.TargetResourceLocation == nil {
		var ret string
		return ret, false
	}
	return *o.TargetResourceLocation, true
}

// HasTargetResourceLocation returns a boolean if a field has been set.
func (o *TeamsAsyncOperation) HasTargetResourceLocation() bool {
	if o != nil && o.TargetResourceLocation != nil {
		return true
	}

	return false
}

// SetTargetResourceLocation gets a reference to the given string and assigns it to the TargetResourceLocation field.
func (o *TeamsAsyncOperation) SetTargetResourceLocation(v string) {
	o.TargetResourceLocation = &v
}

// SetTargetResourceLocationExplicitNull (un)sets TargetResourceLocation to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The TargetResourceLocation value is set to nil even if false is passed
func (o *TeamsAsyncOperation) SetTargetResourceLocationExplicitNull(b bool) {
	o.TargetResourceLocation = nil
	o.isExplicitNullTargetResourceLocation = b
}
// GetError returns the Error field if non-nil, zero value otherwise.
func (o *TeamsAsyncOperation) GetError() AnyOfmicrosoftGraphOperationError {
	if o == nil || o.Error == nil {
		var ret AnyOfmicrosoftGraphOperationError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TeamsAsyncOperation) GetErrorOk() (AnyOfmicrosoftGraphOperationError, bool) {
	if o == nil || o.Error == nil {
		var ret AnyOfmicrosoftGraphOperationError
		return ret, false
	}
	return *o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TeamsAsyncOperation) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given AnyOfmicrosoftGraphOperationError and assigns it to the Error field.
func (o *TeamsAsyncOperation) SetError(v AnyOfmicrosoftGraphOperationError) {
	o.Error = &v
}

// SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Error value is set to nil even if false is passed
func (o *TeamsAsyncOperation) SetErrorExplicitNull(b bool) {
	o.Error = nil
	o.isExplicitNullError = b
}

// MarshalJSON returns the JSON representation of the model.
func (o TeamsAsyncOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OperationType != nil {
		toSerialize["operationType"] = o.OperationType
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.LastActionDateTime != nil {
		toSerialize["lastActionDateTime"] = o.LastActionDateTime
	}
	if o.AttemptsCount != nil {
		toSerialize["attemptsCount"] = o.AttemptsCount
	}
	if o.TargetResourceId == nil {
		if o.isExplicitNullTargetResourceId {
			toSerialize["targetResourceId"] = o.TargetResourceId
		}
	} else {
		toSerialize["targetResourceId"] = o.TargetResourceId
	}
	if o.TargetResourceLocation == nil {
		if o.isExplicitNullTargetResourceLocation {
			toSerialize["targetResourceLocation"] = o.TargetResourceLocation
		}
	} else {
		toSerialize["targetResourceLocation"] = o.TargetResourceLocation
	}
	if o.Error == nil {
		if o.isExplicitNullError {
			toSerialize["error"] = o.Error
		}
	} else {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}



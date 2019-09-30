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
// MicrosoftGraphOperation struct for MicrosoftGraphOperation
type MicrosoftGraphOperation struct {
	Id *string `json:"id,omitempty"`

	Status *AnyOfmicrosoftGraphOperationStatus `json:"status,omitempty"`
	isExplicitNullStatus bool `json:"-"`
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	isExplicitNullCreatedDateTime bool `json:"-"`
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	isExplicitNullLastActionDateTime bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOperation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOperation) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOperation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOperation) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphOperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOperation) GetStatusOk() (AnyOfmicrosoftGraphOperationStatus, bool) {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphOperationStatus
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphOperationStatus and assigns it to the Status field.
func (o *MicrosoftGraphOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus) {
	o.Status = &v
}

// SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Status value is set to nil even if false is passed
func (o *MicrosoftGraphOperation) SetStatusExplicitNull(b bool) {
	o.Status = nil
	o.isExplicitNullStatus = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOperation) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOperation) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOperation) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphOperation) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOperation) SetCreatedDateTimeExplicitNull(b bool) {
	o.CreatedDateTime = nil
	o.isExplicitNullCreatedDateTime = b
}
// GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOperation) GetLastActionDateTime() time.Time {
	if o == nil || o.LastActionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActionDateTime
}

// GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOperation) GetLastActionDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastActionDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastActionDateTime, true
}

// HasLastActionDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOperation) HasLastActionDateTime() bool {
	if o != nil && o.LastActionDateTime != nil {
		return true
	}

	return false
}

// SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.
func (o *MicrosoftGraphOperation) SetLastActionDateTime(v time.Time) {
	o.LastActionDateTime = &v
}

// SetLastActionDateTimeExplicitNull (un)sets LastActionDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastActionDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOperation) SetLastActionDateTimeExplicitNull(b bool) {
	o.LastActionDateTime = nil
	o.isExplicitNullLastActionDateTime = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status == nil {
		if o.isExplicitNullStatus {
			toSerialize["status"] = o.Status
		}
	} else {
		toSerialize["status"] = o.Status
	}
	if o.CreatedDateTime == nil {
		if o.isExplicitNullCreatedDateTime {
			toSerialize["createdDateTime"] = o.CreatedDateTime
		}
	} else {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.LastActionDateTime == nil {
		if o.isExplicitNullLastActionDateTime {
			toSerialize["lastActionDateTime"] = o.LastActionDateTime
		}
	} else {
		toSerialize["lastActionDateTime"] = o.LastActionDateTime
	}
	return json.Marshal(toSerialize)
}



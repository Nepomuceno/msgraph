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
// MicrosoftGraphDomainState struct for MicrosoftGraphDomainState
type MicrosoftGraphDomainState struct {
	Status *string `json:"status,omitempty"`
	isExplicitNullStatus bool `json:"-"`
	Operation *string `json:"operation,omitempty"`
	isExplicitNullOperation bool `json:"-"`
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	isExplicitNullLastActionDateTime bool `json:"-"`
}

// GetStatus returns the Status field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDomainState) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomainState) GetStatusOk() (string, bool) {
	if o == nil || o.Status == nil {
		var ret string
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainState) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MicrosoftGraphDomainState) SetStatus(v string) {
	o.Status = &v
}

// SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Status value is set to nil even if false is passed
func (o *MicrosoftGraphDomainState) SetStatusExplicitNull(b bool) {
	o.Status = nil
	o.isExplicitNullStatus = b
}
// GetOperation returns the Operation field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDomainState) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomainState) GetOperationOk() (string, bool) {
	if o == nil || o.Operation == nil {
		var ret string
		return ret, false
	}
	return *o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainState) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *MicrosoftGraphDomainState) SetOperation(v string) {
	o.Operation = &v
}

// SetOperationExplicitNull (un)sets Operation to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Operation value is set to nil even if false is passed
func (o *MicrosoftGraphDomainState) SetOperationExplicitNull(b bool) {
	o.Operation = nil
	o.isExplicitNullOperation = b
}
// GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDomainState) GetLastActionDateTime() time.Time {
	if o == nil || o.LastActionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActionDateTime
}

// GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDomainState) GetLastActionDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastActionDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastActionDateTime, true
}

// HasLastActionDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDomainState) HasLastActionDateTime() bool {
	if o != nil && o.LastActionDateTime != nil {
		return true
	}

	return false
}

// SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.
func (o *MicrosoftGraphDomainState) SetLastActionDateTime(v time.Time) {
	o.LastActionDateTime = &v
}

// SetLastActionDateTimeExplicitNull (un)sets LastActionDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastActionDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphDomainState) SetLastActionDateTimeExplicitNull(b bool) {
	o.LastActionDateTime = nil
	o.isExplicitNullLastActionDateTime = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDomainState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status == nil {
		if o.isExplicitNullStatus {
			toSerialize["status"] = o.Status
		}
	} else {
		toSerialize["status"] = o.Status
	}
	if o.Operation == nil {
		if o.isExplicitNullOperation {
			toSerialize["operation"] = o.Operation
		}
	} else {
		toSerialize["operation"] = o.Operation
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


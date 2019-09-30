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
// MicrosoftGraphOnenoteOperation struct for MicrosoftGraphOnenoteOperation
type MicrosoftGraphOnenoteOperation struct {
	Id *string `json:"id,omitempty"`

	Status *AnyOfmicrosoftGraphOperationStatus `json:"status,omitempty"`
	isExplicitNullStatus bool `json:"-"`
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	isExplicitNullCreatedDateTime bool `json:"-"`
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	isExplicitNullLastActionDateTime bool `json:"-"`
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	isExplicitNullResourceLocation bool `json:"-"`
	ResourceId *string `json:"resourceId,omitempty"`
	isExplicitNullResourceId bool `json:"-"`
	Error *AnyOfmicrosoftGraphOnenoteOperationError `json:"error,omitempty"`
	isExplicitNullError bool `json:"-"`
	PercentComplete *string `json:"percentComplete,omitempty"`
	isExplicitNullPercentComplete bool `json:"-"`
}

// GetId returns the Id field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperation) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphOnenoteOperation) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperation) GetStatus() AnyOfmicrosoftGraphOperationStatus {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphOperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperation) GetStatusOk() (AnyOfmicrosoftGraphOperationStatus, bool) {
	if o == nil || o.Status == nil {
		var ret AnyOfmicrosoftGraphOperationStatus
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphOperationStatus and assigns it to the Status field.
func (o *MicrosoftGraphOnenoteOperation) SetStatus(v AnyOfmicrosoftGraphOperationStatus) {
	o.Status = &v
}

// SetStatusExplicitNull (un)sets Status to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Status value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteOperation) SetStatusExplicitNull(b bool) {
	o.Status = nil
	o.isExplicitNullStatus = b
}
// GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperation) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperation) GetCreatedDateTimeOk() (time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphOnenoteOperation) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// SetCreatedDateTimeExplicitNull (un)sets CreatedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The CreatedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteOperation) SetCreatedDateTimeExplicitNull(b bool) {
	o.CreatedDateTime = nil
	o.isExplicitNullCreatedDateTime = b
}
// GetLastActionDateTime returns the LastActionDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperation) GetLastActionDateTime() time.Time {
	if o == nil || o.LastActionDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActionDateTime
}

// GetLastActionDateTimeOk returns a tuple with the LastActionDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperation) GetLastActionDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastActionDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastActionDateTime, true
}

// HasLastActionDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasLastActionDateTime() bool {
	if o != nil && o.LastActionDateTime != nil {
		return true
	}

	return false
}

// SetLastActionDateTime gets a reference to the given time.Time and assigns it to the LastActionDateTime field.
func (o *MicrosoftGraphOnenoteOperation) SetLastActionDateTime(v time.Time) {
	o.LastActionDateTime = &v
}

// SetLastActionDateTimeExplicitNull (un)sets LastActionDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastActionDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteOperation) SetLastActionDateTimeExplicitNull(b bool) {
	o.LastActionDateTime = nil
	o.isExplicitNullLastActionDateTime = b
}
// GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperation) GetResourceLocation() string {
	if o == nil || o.ResourceLocation == nil {
		var ret string
		return ret
	}
	return *o.ResourceLocation
}

// GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperation) GetResourceLocationOk() (string, bool) {
	if o == nil || o.ResourceLocation == nil {
		var ret string
		return ret, false
	}
	return *o.ResourceLocation, true
}

// HasResourceLocation returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasResourceLocation() bool {
	if o != nil && o.ResourceLocation != nil {
		return true
	}

	return false
}

// SetResourceLocation gets a reference to the given string and assigns it to the ResourceLocation field.
func (o *MicrosoftGraphOnenoteOperation) SetResourceLocation(v string) {
	o.ResourceLocation = &v
}

// SetResourceLocationExplicitNull (un)sets ResourceLocation to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ResourceLocation value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteOperation) SetResourceLocationExplicitNull(b bool) {
	o.ResourceLocation = nil
	o.isExplicitNullResourceLocation = b
}
// GetResourceId returns the ResourceId field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperation) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperation) GetResourceIdOk() (string, bool) {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret, false
	}
	return *o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *MicrosoftGraphOnenoteOperation) SetResourceId(v string) {
	o.ResourceId = &v
}

// SetResourceIdExplicitNull (un)sets ResourceId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ResourceId value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteOperation) SetResourceIdExplicitNull(b bool) {
	o.ResourceId = nil
	o.isExplicitNullResourceId = b
}
// GetError returns the Error field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperation) GetError() AnyOfmicrosoftGraphOnenoteOperationError {
	if o == nil || o.Error == nil {
		var ret AnyOfmicrosoftGraphOnenoteOperationError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperation) GetErrorOk() (AnyOfmicrosoftGraphOnenoteOperationError, bool) {
	if o == nil || o.Error == nil {
		var ret AnyOfmicrosoftGraphOnenoteOperationError
		return ret, false
	}
	return *o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given AnyOfmicrosoftGraphOnenoteOperationError and assigns it to the Error field.
func (o *MicrosoftGraphOnenoteOperation) SetError(v AnyOfmicrosoftGraphOnenoteOperationError) {
	o.Error = &v
}

// SetErrorExplicitNull (un)sets Error to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Error value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteOperation) SetErrorExplicitNull(b bool) {
	o.Error = nil
	o.isExplicitNullError = b
}
// GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.
func (o *MicrosoftGraphOnenoteOperation) GetPercentComplete() string {
	if o == nil || o.PercentComplete == nil {
		var ret string
		return ret
	}
	return *o.PercentComplete
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphOnenoteOperation) GetPercentCompleteOk() (string, bool) {
	if o == nil || o.PercentComplete == nil {
		var ret string
		return ret, false
	}
	return *o.PercentComplete, true
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *MicrosoftGraphOnenoteOperation) HasPercentComplete() bool {
	if o != nil && o.PercentComplete != nil {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given string and assigns it to the PercentComplete field.
func (o *MicrosoftGraphOnenoteOperation) SetPercentComplete(v string) {
	o.PercentComplete = &v
}

// SetPercentCompleteExplicitNull (un)sets PercentComplete to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The PercentComplete value is set to nil even if false is passed
func (o *MicrosoftGraphOnenoteOperation) SetPercentCompleteExplicitNull(b bool) {
	o.PercentComplete = nil
	o.isExplicitNullPercentComplete = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphOnenoteOperation) MarshalJSON() ([]byte, error) {
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
	if o.ResourceLocation == nil {
		if o.isExplicitNullResourceLocation {
			toSerialize["resourceLocation"] = o.ResourceLocation
		}
	} else {
		toSerialize["resourceLocation"] = o.ResourceLocation
	}
	if o.ResourceId == nil {
		if o.isExplicitNullResourceId {
			toSerialize["resourceId"] = o.ResourceId
		}
	} else {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.Error == nil {
		if o.isExplicitNullError {
			toSerialize["error"] = o.Error
		}
	} else {
		toSerialize["error"] = o.Error
	}
	if o.PercentComplete == nil {
		if o.isExplicitNullPercentComplete {
			toSerialize["percentComplete"] = o.PercentComplete
		}
	} else {
		toSerialize["percentComplete"] = o.PercentComplete
	}
	return json.Marshal(toSerialize)
}



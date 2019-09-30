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
// MicrosoftGraphResponseStatus struct for MicrosoftGraphResponseStatus
type MicrosoftGraphResponseStatus struct {
	Response *AnyOfmicrosoftGraphResponseType `json:"response,omitempty"`
	isExplicitNullResponse bool `json:"-"`
	Time *time.Time `json:"time,omitempty"`
	isExplicitNullTime bool `json:"-"`
}

// GetResponse returns the Response field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResponseStatus) GetResponse() AnyOfmicrosoftGraphResponseType {
	if o == nil || o.Response == nil {
		var ret AnyOfmicrosoftGraphResponseType
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResponseStatus) GetResponseOk() (AnyOfmicrosoftGraphResponseType, bool) {
	if o == nil || o.Response == nil {
		var ret AnyOfmicrosoftGraphResponseType
		return ret, false
	}
	return *o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *MicrosoftGraphResponseStatus) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given AnyOfmicrosoftGraphResponseType and assigns it to the Response field.
func (o *MicrosoftGraphResponseStatus) SetResponse(v AnyOfmicrosoftGraphResponseType) {
	o.Response = &v
}

// SetResponseExplicitNull (un)sets Response to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Response value is set to nil even if false is passed
func (o *MicrosoftGraphResponseStatus) SetResponseExplicitNull(b bool) {
	o.Response = nil
	o.isExplicitNullResponse = b
}
// GetTime returns the Time field if non-nil, zero value otherwise.
func (o *MicrosoftGraphResponseStatus) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResponseStatus) GetTimeOk() (time.Time, bool) {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret, false
	}
	return *o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *MicrosoftGraphResponseStatus) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *MicrosoftGraphResponseStatus) SetTime(v time.Time) {
	o.Time = &v
}

// SetTimeExplicitNull (un)sets Time to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The Time value is set to nil even if false is passed
func (o *MicrosoftGraphResponseStatus) SetTimeExplicitNull(b bool) {
	o.Time = nil
	o.isExplicitNullTime = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphResponseStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Response == nil {
		if o.isExplicitNullResponse {
			toSerialize["response"] = o.Response
		}
	} else {
		toSerialize["response"] = o.Response
	}
	if o.Time == nil {
		if o.isExplicitNullTime {
			toSerialize["time"] = o.Time
		}
	} else {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}



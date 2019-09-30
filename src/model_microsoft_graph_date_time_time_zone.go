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
// MicrosoftGraphDateTimeTimeZone struct for MicrosoftGraphDateTimeTimeZone
type MicrosoftGraphDateTimeTimeZone struct {
	DateTime *string `json:"dateTime,omitempty"`

	TimeZone *string `json:"timeZone,omitempty"`
	isExplicitNullTimeZone bool `json:"-"`
}

// GetDateTime returns the DateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDateTimeTimeZone) GetDateTime() string {
	if o == nil || o.DateTime == nil {
		var ret string
		return ret
	}
	return *o.DateTime
}

// GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDateTimeTimeZone) GetDateTimeOk() (string, bool) {
	if o == nil || o.DateTime == nil {
		var ret string
		return ret, false
	}
	return *o.DateTime, true
}

// HasDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDateTimeTimeZone) HasDateTime() bool {
	if o != nil && o.DateTime != nil {
		return true
	}

	return false
}

// SetDateTime gets a reference to the given string and assigns it to the DateTime field.
func (o *MicrosoftGraphDateTimeTimeZone) SetDateTime(v string) {
	o.DateTime = &v
}

// GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDateTimeTimeZone) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDateTimeTimeZone) GetTimeZoneOk() (string, bool) {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret, false
	}
	return *o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *MicrosoftGraphDateTimeTimeZone) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *MicrosoftGraphDateTimeTimeZone) SetTimeZone(v string) {
	o.TimeZone = &v
}

// SetTimeZoneExplicitNull (un)sets TimeZone to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The TimeZone value is set to nil even if false is passed
func (o *MicrosoftGraphDateTimeTimeZone) SetTimeZoneExplicitNull(b bool) {
	o.TimeZone = nil
	o.isExplicitNullTimeZone = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDateTimeTimeZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateTime != nil {
		toSerialize["dateTime"] = o.DateTime
	}
	if o.TimeZone == nil {
		if o.isExplicitNullTimeZone {
			toSerialize["timeZone"] = o.TimeZone
		}
	} else {
		toSerialize["timeZone"] = o.TimeZone
	}
	return json.Marshal(toSerialize)
}



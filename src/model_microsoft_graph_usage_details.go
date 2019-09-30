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
// MicrosoftGraphUsageDetails struct for MicrosoftGraphUsageDetails
type MicrosoftGraphUsageDetails struct {
	LastAccessedDateTime *time.Time `json:"lastAccessedDateTime,omitempty"`
	isExplicitNullLastAccessedDateTime bool `json:"-"`
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	isExplicitNullLastModifiedDateTime bool `json:"-"`
}

// GetLastAccessedDateTime returns the LastAccessedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphUsageDetails) GetLastAccessedDateTime() time.Time {
	if o == nil || o.LastAccessedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessedDateTime
}

// GetLastAccessedDateTimeOk returns a tuple with the LastAccessedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUsageDetails) GetLastAccessedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastAccessedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastAccessedDateTime, true
}

// HasLastAccessedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphUsageDetails) HasLastAccessedDateTime() bool {
	if o != nil && o.LastAccessedDateTime != nil {
		return true
	}

	return false
}

// SetLastAccessedDateTime gets a reference to the given time.Time and assigns it to the LastAccessedDateTime field.
func (o *MicrosoftGraphUsageDetails) SetLastAccessedDateTime(v time.Time) {
	o.LastAccessedDateTime = &v
}

// SetLastAccessedDateTimeExplicitNull (un)sets LastAccessedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastAccessedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphUsageDetails) SetLastAccessedDateTimeExplicitNull(b bool) {
	o.LastAccessedDateTime = nil
	o.isExplicitNullLastAccessedDateTime = b
}
// GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.
func (o *MicrosoftGraphUsageDetails) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUsageDetails) GetLastModifiedDateTimeOk() (time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphUsageDetails) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphUsageDetails) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// SetLastModifiedDateTimeExplicitNull (un)sets LastModifiedDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The LastModifiedDateTime value is set to nil even if false is passed
func (o *MicrosoftGraphUsageDetails) SetLastModifiedDateTimeExplicitNull(b bool) {
	o.LastModifiedDateTime = nil
	o.isExplicitNullLastModifiedDateTime = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphUsageDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastAccessedDateTime == nil {
		if o.isExplicitNullLastAccessedDateTime {
			toSerialize["lastAccessedDateTime"] = o.LastAccessedDateTime
		}
	} else {
		toSerialize["lastAccessedDateTime"] = o.LastAccessedDateTime
	}
	if o.LastModifiedDateTime == nil {
		if o.isExplicitNullLastModifiedDateTime {
			toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
		}
	} else {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	return json.Marshal(toSerialize)
}


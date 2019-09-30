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
// MicrosoftGraphDeviceEnrollmentPlatformRestriction struct for MicrosoftGraphDeviceEnrollmentPlatformRestriction
type MicrosoftGraphDeviceEnrollmentPlatformRestriction struct {
	// Block the platform from enrolling
	PlatformBlocked *bool `json:"platformBlocked,omitempty"`

	// Block personally owned devices from enrolling
	PersonalDeviceEnrollmentBlocked *bool `json:"personalDeviceEnrollmentBlocked,omitempty"`

	// Min OS version supported
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	isExplicitNullOsMinimumVersion bool `json:"-"`
	// Max OS version supported
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	isExplicitNullOsMaximumVersion bool `json:"-"`
}

// GetPlatformBlocked returns the PlatformBlocked field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetPlatformBlocked() bool {
	if o == nil || o.PlatformBlocked == nil {
		var ret bool
		return ret
	}
	return *o.PlatformBlocked
}

// GetPlatformBlockedOk returns a tuple with the PlatformBlocked field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetPlatformBlockedOk() (bool, bool) {
	if o == nil || o.PlatformBlocked == nil {
		var ret bool
		return ret, false
	}
	return *o.PlatformBlocked, true
}

// HasPlatformBlocked returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) HasPlatformBlocked() bool {
	if o != nil && o.PlatformBlocked != nil {
		return true
	}

	return false
}

// SetPlatformBlocked gets a reference to the given bool and assigns it to the PlatformBlocked field.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetPlatformBlocked(v bool) {
	o.PlatformBlocked = &v
}

// GetPersonalDeviceEnrollmentBlocked returns the PersonalDeviceEnrollmentBlocked field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetPersonalDeviceEnrollmentBlocked() bool {
	if o == nil || o.PersonalDeviceEnrollmentBlocked == nil {
		var ret bool
		return ret
	}
	return *o.PersonalDeviceEnrollmentBlocked
}

// GetPersonalDeviceEnrollmentBlockedOk returns a tuple with the PersonalDeviceEnrollmentBlocked field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetPersonalDeviceEnrollmentBlockedOk() (bool, bool) {
	if o == nil || o.PersonalDeviceEnrollmentBlocked == nil {
		var ret bool
		return ret, false
	}
	return *o.PersonalDeviceEnrollmentBlocked, true
}

// HasPersonalDeviceEnrollmentBlocked returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) HasPersonalDeviceEnrollmentBlocked() bool {
	if o != nil && o.PersonalDeviceEnrollmentBlocked != nil {
		return true
	}

	return false
}

// SetPersonalDeviceEnrollmentBlocked gets a reference to the given bool and assigns it to the PersonalDeviceEnrollmentBlocked field.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetPersonalDeviceEnrollmentBlocked(v bool) {
	o.PersonalDeviceEnrollmentBlocked = &v
}

// GetOsMinimumVersion returns the OsMinimumVersion field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetOsMinimumVersion() string {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMinimumVersion
}

// GetOsMinimumVersionOk returns a tuple with the OsMinimumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetOsMinimumVersionOk() (string, bool) {
	if o == nil || o.OsMinimumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMinimumVersion, true
}

// HasOsMinimumVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) HasOsMinimumVersion() bool {
	if o != nil && o.OsMinimumVersion != nil {
		return true
	}

	return false
}

// SetOsMinimumVersion gets a reference to the given string and assigns it to the OsMinimumVersion field.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetOsMinimumVersion(v string) {
	o.OsMinimumVersion = &v
}

// SetOsMinimumVersionExplicitNull (un)sets OsMinimumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMinimumVersion value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetOsMinimumVersionExplicitNull(b bool) {
	o.OsMinimumVersion = nil
	o.isExplicitNullOsMinimumVersion = b
}
// GetOsMaximumVersion returns the OsMaximumVersion field if non-nil, zero value otherwise.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetOsMaximumVersion() string {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret
	}
	return *o.OsMaximumVersion
}

// GetOsMaximumVersionOk returns a tuple with the OsMaximumVersion field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) GetOsMaximumVersionOk() (string, bool) {
	if o == nil || o.OsMaximumVersion == nil {
		var ret string
		return ret, false
	}
	return *o.OsMaximumVersion, true
}

// HasOsMaximumVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) HasOsMaximumVersion() bool {
	if o != nil && o.OsMaximumVersion != nil {
		return true
	}

	return false
}

// SetOsMaximumVersion gets a reference to the given string and assigns it to the OsMaximumVersion field.
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetOsMaximumVersion(v string) {
	o.OsMaximumVersion = &v
}

// SetOsMaximumVersionExplicitNull (un)sets OsMaximumVersion to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The OsMaximumVersion value is set to nil even if false is passed
func (o *MicrosoftGraphDeviceEnrollmentPlatformRestriction) SetOsMaximumVersionExplicitNull(b bool) {
	o.OsMaximumVersion = nil
	o.isExplicitNullOsMaximumVersion = b
}

// MarshalJSON returns the JSON representation of the model.
func (o MicrosoftGraphDeviceEnrollmentPlatformRestriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PlatformBlocked != nil {
		toSerialize["platformBlocked"] = o.PlatformBlocked
	}
	if o.PersonalDeviceEnrollmentBlocked != nil {
		toSerialize["personalDeviceEnrollmentBlocked"] = o.PersonalDeviceEnrollmentBlocked
	}
	if o.OsMinimumVersion == nil {
		if o.isExplicitNullOsMinimumVersion {
			toSerialize["osMinimumVersion"] = o.OsMinimumVersion
		}
	} else {
		toSerialize["osMinimumVersion"] = o.OsMinimumVersion
	}
	if o.OsMaximumVersion == nil {
		if o.isExplicitNullOsMaximumVersion {
			toSerialize["osMaximumVersion"] = o.OsMaximumVersion
		}
	} else {
		toSerialize["osMaximumVersion"] = o.OsMaximumVersion
	}
	return json.Marshal(toSerialize)
}



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
// ManagedIosLobApp Contains properties and inherited properties for Managed iOS Line Of Business apps.
type ManagedIosLobApp struct {
	// The Identity Name.
	BundleId *string `json:"bundleId,omitempty"`
	isExplicitNullBundleId bool `json:"-"`
	ApplicableDeviceType *MicrosoftGraphIosDeviceType `json:"applicableDeviceType,omitempty"`

	// The value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *AnyOfmicrosoftGraphIosMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	isExplicitNullMinimumSupportedOperatingSystem bool `json:"-"`
	// The expiration time.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	isExplicitNullExpirationDateTime bool `json:"-"`
	// The version number of managed iOS Line of Business (LoB) app.
	VersionNumber *string `json:"versionNumber,omitempty"`
	isExplicitNullVersionNumber bool `json:"-"`
	// The build number of managed iOS Line of Business (LoB) app.
	BuildNumber *string `json:"buildNumber,omitempty"`
	isExplicitNullBuildNumber bool `json:"-"`
}

// GetBundleId returns the BundleId field if non-nil, zero value otherwise.
func (o *ManagedIosLobApp) GetBundleId() string {
	if o == nil || o.BundleId == nil {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedIosLobApp) GetBundleIdOk() (string, bool) {
	if o == nil || o.BundleId == nil {
		var ret string
		return ret, false
	}
	return *o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *ManagedIosLobApp) HasBundleId() bool {
	if o != nil && o.BundleId != nil {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *ManagedIosLobApp) SetBundleId(v string) {
	o.BundleId = &v
}

// SetBundleIdExplicitNull (un)sets BundleId to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The BundleId value is set to nil even if false is passed
func (o *ManagedIosLobApp) SetBundleIdExplicitNull(b bool) {
	o.BundleId = nil
	o.isExplicitNullBundleId = b
}
// GetApplicableDeviceType returns the ApplicableDeviceType field if non-nil, zero value otherwise.
func (o *ManagedIosLobApp) GetApplicableDeviceType() MicrosoftGraphIosDeviceType {
	if o == nil || o.ApplicableDeviceType == nil {
		var ret MicrosoftGraphIosDeviceType
		return ret
	}
	return *o.ApplicableDeviceType
}

// GetApplicableDeviceTypeOk returns a tuple with the ApplicableDeviceType field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedIosLobApp) GetApplicableDeviceTypeOk() (MicrosoftGraphIosDeviceType, bool) {
	if o == nil || o.ApplicableDeviceType == nil {
		var ret MicrosoftGraphIosDeviceType
		return ret, false
	}
	return *o.ApplicableDeviceType, true
}

// HasApplicableDeviceType returns a boolean if a field has been set.
func (o *ManagedIosLobApp) HasApplicableDeviceType() bool {
	if o != nil && o.ApplicableDeviceType != nil {
		return true
	}

	return false
}

// SetApplicableDeviceType gets a reference to the given MicrosoftGraphIosDeviceType and assigns it to the ApplicableDeviceType field.
func (o *ManagedIosLobApp) SetApplicableDeviceType(v MicrosoftGraphIosDeviceType) {
	o.ApplicableDeviceType = &v
}

// GetMinimumSupportedOperatingSystem returns the MinimumSupportedOperatingSystem field if non-nil, zero value otherwise.
func (o *ManagedIosLobApp) GetMinimumSupportedOperatingSystem() AnyOfmicrosoftGraphIosMinimumOperatingSystem {
	if o == nil || o.MinimumSupportedOperatingSystem == nil {
		var ret AnyOfmicrosoftGraphIosMinimumOperatingSystem
		return ret
	}
	return *o.MinimumSupportedOperatingSystem
}

// GetMinimumSupportedOperatingSystemOk returns a tuple with the MinimumSupportedOperatingSystem field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedIosLobApp) GetMinimumSupportedOperatingSystemOk() (AnyOfmicrosoftGraphIosMinimumOperatingSystem, bool) {
	if o == nil || o.MinimumSupportedOperatingSystem == nil {
		var ret AnyOfmicrosoftGraphIosMinimumOperatingSystem
		return ret, false
	}
	return *o.MinimumSupportedOperatingSystem, true
}

// HasMinimumSupportedOperatingSystem returns a boolean if a field has been set.
func (o *ManagedIosLobApp) HasMinimumSupportedOperatingSystem() bool {
	if o != nil && o.MinimumSupportedOperatingSystem != nil {
		return true
	}

	return false
}

// SetMinimumSupportedOperatingSystem gets a reference to the given AnyOfmicrosoftGraphIosMinimumOperatingSystem and assigns it to the MinimumSupportedOperatingSystem field.
func (o *ManagedIosLobApp) SetMinimumSupportedOperatingSystem(v AnyOfmicrosoftGraphIosMinimumOperatingSystem) {
	o.MinimumSupportedOperatingSystem = &v
}

// SetMinimumSupportedOperatingSystemExplicitNull (un)sets MinimumSupportedOperatingSystem to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The MinimumSupportedOperatingSystem value is set to nil even if false is passed
func (o *ManagedIosLobApp) SetMinimumSupportedOperatingSystemExplicitNull(b bool) {
	o.MinimumSupportedOperatingSystem = nil
	o.isExplicitNullMinimumSupportedOperatingSystem = b
}
// GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.
func (o *ManagedIosLobApp) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedIosLobApp) GetExpirationDateTimeOk() (time.Time, bool) {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *ManagedIosLobApp) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime != nil {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.
func (o *ManagedIosLobApp) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime = &v
}

// SetExpirationDateTimeExplicitNull (un)sets ExpirationDateTime to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ExpirationDateTime value is set to nil even if false is passed
func (o *ManagedIosLobApp) SetExpirationDateTimeExplicitNull(b bool) {
	o.ExpirationDateTime = nil
	o.isExplicitNullExpirationDateTime = b
}
// GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.
func (o *ManagedIosLobApp) GetVersionNumber() string {
	if o == nil || o.VersionNumber == nil {
		var ret string
		return ret
	}
	return *o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedIosLobApp) GetVersionNumberOk() (string, bool) {
	if o == nil || o.VersionNumber == nil {
		var ret string
		return ret, false
	}
	return *o.VersionNumber, true
}

// HasVersionNumber returns a boolean if a field has been set.
func (o *ManagedIosLobApp) HasVersionNumber() bool {
	if o != nil && o.VersionNumber != nil {
		return true
	}

	return false
}

// SetVersionNumber gets a reference to the given string and assigns it to the VersionNumber field.
func (o *ManagedIosLobApp) SetVersionNumber(v string) {
	o.VersionNumber = &v
}

// SetVersionNumberExplicitNull (un)sets VersionNumber to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The VersionNumber value is set to nil even if false is passed
func (o *ManagedIosLobApp) SetVersionNumberExplicitNull(b bool) {
	o.VersionNumber = nil
	o.isExplicitNullVersionNumber = b
}
// GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.
func (o *ManagedIosLobApp) GetBuildNumber() string {
	if o == nil || o.BuildNumber == nil {
		var ret string
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ManagedIosLobApp) GetBuildNumberOk() (string, bool) {
	if o == nil || o.BuildNumber == nil {
		var ret string
		return ret, false
	}
	return *o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *ManagedIosLobApp) HasBuildNumber() bool {
	if o != nil && o.BuildNumber != nil {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given string and assigns it to the BuildNumber field.
func (o *ManagedIosLobApp) SetBuildNumber(v string) {
	o.BuildNumber = &v
}

// SetBuildNumberExplicitNull (un)sets BuildNumber to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The BuildNumber value is set to nil even if false is passed
func (o *ManagedIosLobApp) SetBuildNumberExplicitNull(b bool) {
	o.BuildNumber = nil
	o.isExplicitNullBuildNumber = b
}

// MarshalJSON returns the JSON representation of the model.
func (o ManagedIosLobApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BundleId == nil {
		if o.isExplicitNullBundleId {
			toSerialize["bundleId"] = o.BundleId
		}
	} else {
		toSerialize["bundleId"] = o.BundleId
	}
	if o.ApplicableDeviceType != nil {
		toSerialize["applicableDeviceType"] = o.ApplicableDeviceType
	}
	if o.MinimumSupportedOperatingSystem == nil {
		if o.isExplicitNullMinimumSupportedOperatingSystem {
			toSerialize["minimumSupportedOperatingSystem"] = o.MinimumSupportedOperatingSystem
		}
	} else {
		toSerialize["minimumSupportedOperatingSystem"] = o.MinimumSupportedOperatingSystem
	}
	if o.ExpirationDateTime == nil {
		if o.isExplicitNullExpirationDateTime {
			toSerialize["expirationDateTime"] = o.ExpirationDateTime
		}
	} else {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if o.VersionNumber == nil {
		if o.isExplicitNullVersionNumber {
			toSerialize["versionNumber"] = o.VersionNumber
		}
	} else {
		toSerialize["versionNumber"] = o.VersionNumber
	}
	if o.BuildNumber == nil {
		if o.isExplicitNullBuildNumber {
			toSerialize["buildNumber"] = o.BuildNumber
		}
	} else {
		toSerialize["buildNumber"] = o.BuildNumber
	}
	return json.Marshal(toSerialize)
}

